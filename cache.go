package main

import (
	"context"
	"errors"
	"fmt"
	"slices"

	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	toolscache "k8s.io/client-go/tools/cache"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func mergeTransformFunc(ff ...toolscache.TransformFunc) toolscache.TransformFunc {
	return func(i interface{}) (interface{}, error) {
		var err error

		for _, f := range ff {
			if i, err = f(i); err != nil {
				return nil, err
			}
		}
		return i, nil
	}
}

func trimRole() toolscache.TransformFunc {
	return mergeTransformFunc(
		cache.TransformStripManagedFields(),
		func(i interface{}) (interface{}, error) {
			r, ok := i.(*rbacv1.Role)
			if !ok {
				return nil, fmt.Errorf("error caching Role: expected Role received %T", i)
			}

			r.Rules = filterNamespacesRelatedPolicyRules(r.Rules)
			if len(r.Rules) == 0 {
				return nil, nil
			}
			return r, nil
		},
	)
}

func trimClusterRole() toolscache.TransformFunc {
	return mergeTransformFunc(
		cache.TransformStripManagedFields(),
		func(i interface{}) (interface{}, error) {
			cr, ok := i.(*rbacv1.ClusterRole)
			if !ok {
				return nil, fmt.Errorf("error caching ClusterRole: expected a ClusterRole received %T", i)
			}

			// can't define at this time if it will relate to namespaces, so let's keep it
			if cr.AggregationRule != nil && cr.AggregationRule.ClusterRoleSelectors != nil {
				return i, nil
			}

			cr.Rules = filterNamespacesRelatedPolicyRules(cr.Rules)
			if len(cr.Rules) == 0 {
				return nil, nil
			}
			return cr, nil
		},
	)
}

func filterNamespacesRelatedPolicyRules(pp []rbacv1.PolicyRule) []rbacv1.PolicyRule {
	fr := []rbacv1.PolicyRule{}
	for _, r := range pp {
		if slices.Contains(r.APIGroups, "") &&
			slices.Contains(r.Resources, "namespaces") &&
			slices.ContainsFunc(r.Verbs, func(v string) bool {
				return v == "get" || v == "list"
			}) {

			fr = append(fr, r)
		}
	}
	return fr
}

func BuildAndStartCache(ctx context.Context, cfg *rest.Config) (cache.Cache, error) {
	s := runtime.NewScheme()
	if err := corev1.AddToScheme(s); err != nil {
		return nil, err
	}
	if err := rbacv1.AddToScheme(s); err != nil {
		return nil, err
	}
	oo := []client.Object{
		&corev1.Namespace{},
		&rbacv1.RoleBinding{},
		&rbacv1.ClusterRole{},
		&rbacv1.ClusterRoleBinding{},
		&rbacv1.Role{},
	}
	c, err := cache.New(cfg, cache.Options{
		Scheme: s,
		ByObject: map[client.Object]cache.ByObject{
			&corev1.Namespace{}: {
				Transform: cache.TransformStripManagedFields(),
			},
			&rbacv1.ClusterRole{}: {
				Transform: trimClusterRole(),
			},
			&rbacv1.ClusterRoleBinding{}: {
				Transform: cache.TransformStripManagedFields(),
			},
			&rbacv1.RoleBinding{}: {
				Transform: cache.TransformStripManagedFields(),
			},
			&rbacv1.Role{}: {
				Transform: trimRole(),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	for _, o := range oo {
		_, err := c.GetInformer(ctx, o)
		if err != nil {
			return nil, fmt.Errorf("error starting cache: getting informer for %s: %w", o.GetObjectKind().GroupVersionKind().String(), err)
		}
	}

	go func() {
		if err := c.Start(ctx); err != nil {
			panic(err)
		}
	}()
	if !c.WaitForCacheSync(ctx) {
		return nil, errors.New("error starting the cache")
	}

	return c, nil
}

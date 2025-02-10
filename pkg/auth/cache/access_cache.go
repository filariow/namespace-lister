package cache

import (
	"slices"
	"strings"
	"sync/atomic"

	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
)

type store struct {
	accessMap  map[rbacv1.Subject][]string
	namespaces []corev1.Namespace
}

// stores data
type AccessCache struct {
	data atomic.Pointer[store]
}

func NewAccessCache() *AccessCache {
	c := AccessCache{
		data: atomic.Pointer[store]{},
	}

	c.data.Store(&store{
		accessMap:  map[rbacv1.Subject][]string{},
		namespaces: []corev1.Namespace{},
	})
	return &c
}

func (c *AccessCache) List(subject rbacv1.Subject) []corev1.Namespace {
	// load cache data
	m := c.data.Load()
	if m == nil {
		return nil
	}

	// build the namespace list for current subject
	namespacesName := m.accessMap[subject]
	namespaces := make([]corev1.Namespace, 0, len(namespacesName))
	for _, namespaceName := range namespacesName {
		i, found := slices.BinarySearchFunc(m.namespaces, namespaceName, func(lookup corev1.Namespace, target string) int {
			return strings.Compare(lookup.GetName(), target)
		})
		if !found {
			continue
		}

		namespaces = append(namespaces, m.namespaces[i])
	}

	return slices.Clip(namespaces)
}

func (c *AccessCache) Restock(nn corev1.NamespaceList, data *map[rbacv1.Subject][]corev1.Namespace) error {
	// sort namespaces so that we can then use binary search
	namespaces := nn.Items
	slices.SortFunc(namespaces, func(n1, n2 corev1.Namespace) int {
		return strings.Compare(n1.GetName(), n2.GetName())
	})

	// build accessMap
	accessMap := make(map[rbacv1.Subject][]string, len(*data))
	for k, v := range *data {
		accessMap[k] = project(v, func(n corev1.Namespace) string {
			return n.GetName()
		})
	}

	// store new data
	c.data.Store(&store{
		namespaces: namespaces,
		accessMap:  accessMap,
	})

	return nil
}

func project[S ~[]E, E, T any](s S, projectionFunc func(E) T) []T {
	tt := make([]T, len(s))
	for i, e := range s {
		tt[i] = projectionFunc(e)
	}
	return tt
}

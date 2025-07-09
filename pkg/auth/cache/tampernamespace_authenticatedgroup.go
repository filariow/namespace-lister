package cache

import (
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
)

const (
	AccessLabelKey         = "konflux-ci.dev/access"
	AccessLabelValueShared = "community"

	SystemAuthenticatedGroup = "system:authenticated"
)

// TamperNamespaceWithSharedAccessVirtualLabel if the Subject is the well-known `system:authenticated` Group,
// this function overrides the namespace label `konflux-ci.dev/access` to `community`
func TamperNamespaceWithSharedAccessVirtualLabel(s rbacv1.Subject, ns corev1.Namespace) corev1.Namespace {
	if s.Kind == rbacv1.GroupKind &&
		s.APIGroup == rbacv1.GroupName &&
		s.Name == SystemAuthenticatedGroup {
		// retrieve labels
		ll := ns.GetLabels()

		// override AccessLabel Key
		ll[AccessLabelKey] = AccessLabelValueShared

		// update labels
		ns.Labels = ll
	}

	// return namespace
	return ns
}

package authenticated

import (
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
)

// TamperNamespaceWithSharedAccessVirtualLabel if the Subject is the well-known `system:authenticated` Group,
// this function overrides the namespace label `konflux-ci.dev/access` to `community`
func TamperNamespaceWithSharedAccessVirtualLabel(s rbacv1.Subject, ns corev1.Namespace) corev1.Namespace {
	if !isSystemAuthenticatedGroup(s) {
		return ns
	}

	// retrieve labels
	ll := ns.GetLabels()

	// override AccessLabel Key
	ll[AccessLabelKey] = AccessLabelValueShared

	// update labels
	ns.Labels = ll

	// return namespace
	return ns
}

func isSystemAuthenticatedGroup(s rbacv1.Subject) bool {
	return s.Kind == rbacv1.GroupKind &&
		s.APIGroup == rbacv1.GroupName &&
		s.Name == SystemAuthenticatedGroup
}

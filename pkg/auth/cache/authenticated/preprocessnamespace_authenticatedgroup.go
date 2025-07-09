package authenticated

import corev1 "k8s.io/api/core/v1"

// PreprocessNamespaceRemoveSharedAccessVirtualLabel removes the `konflux-ci.dev/access` label from the namespace
func PreprocessNamespaceRemoveSharedAccessVirtualLabel(ns corev1.Namespace) corev1.Namespace {
	// retrieve labels
	ll := ns.GetLabels()

	// override AccessLabel Key
	delete(ll, AccessLabelKey)

	// update labels
	ns.Labels = ll

	return ns
}

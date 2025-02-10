package cache_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/konflux-ci/namespace-lister/pkg/auth/cache"
)

var _ = Describe("AuthCache", func() {
	enn := []corev1.Namespace{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:        "myns",
				Labels:      map[string]string{"key": "value"},
				Annotations: map[string]string{"key": "value"},
			},
		},
	}

	It("returns an empty result if it is empty", func() {
		// given
		emptyCache := cache.NewAccessCache()

		// when
		nn := emptyCache.List(rbacv1.Subject{})

		// then
		Expect(nn).To(BeEmpty())
	})

	It("matches subjects", func(ctx context.Context) {
		// given
		sub := rbacv1.Subject{Kind: "User", Name: "myuser"}
		nnl := corev1.NamespaceList{Items: enn}
		c := cache.NewAccessCache()

		Expect(c.Restock(nnl, &map[rbacv1.Subject][]corev1.Namespace{sub: enn})).To(Succeed())

		// when
		nn := c.List(sub)

		// then
		Expect(nn).To(BeEquivalentTo(enn))
	})
})

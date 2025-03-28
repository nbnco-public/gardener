// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	context "context"

	corev1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	scheme "github.com/gardener/gardener/pkg/client/core/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// SecretBindingsGetter has a method to return a SecretBindingInterface.
// A group's client should implement this interface.
type SecretBindingsGetter interface {
	SecretBindings(namespace string) SecretBindingInterface
}

// SecretBindingInterface has methods to work with SecretBinding resources.
type SecretBindingInterface interface {
	Create(ctx context.Context, secretBinding *corev1beta1.SecretBinding, opts v1.CreateOptions) (*corev1beta1.SecretBinding, error)
	Update(ctx context.Context, secretBinding *corev1beta1.SecretBinding, opts v1.UpdateOptions) (*corev1beta1.SecretBinding, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*corev1beta1.SecretBinding, error)
	List(ctx context.Context, opts v1.ListOptions) (*corev1beta1.SecretBindingList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *corev1beta1.SecretBinding, err error)
	SecretBindingExpansion
}

// secretBindings implements SecretBindingInterface
type secretBindings struct {
	*gentype.ClientWithList[*corev1beta1.SecretBinding, *corev1beta1.SecretBindingList]
}

// newSecretBindings returns a SecretBindings
func newSecretBindings(c *CoreV1beta1Client, namespace string) *secretBindings {
	return &secretBindings{
		gentype.NewClientWithList[*corev1beta1.SecretBinding, *corev1beta1.SecretBindingList](
			"secretbindings",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *corev1beta1.SecretBinding { return &corev1beta1.SecretBinding{} },
			func() *corev1beta1.SecretBindingList { return &corev1beta1.SecretBindingList{} },
		),
	}
}

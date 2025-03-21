// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/gardener/gardener/pkg/apis/settings/v1alpha1"
	settingsv1alpha1 "github.com/gardener/gardener/pkg/client/settings/clientset/versioned/typed/settings/v1alpha1"
	gentype "k8s.io/client-go/gentype"
)

// fakeOpenIDConnectPresets implements OpenIDConnectPresetInterface
type fakeOpenIDConnectPresets struct {
	*gentype.FakeClientWithList[*v1alpha1.OpenIDConnectPreset, *v1alpha1.OpenIDConnectPresetList]
	Fake *FakeSettingsV1alpha1
}

func newFakeOpenIDConnectPresets(fake *FakeSettingsV1alpha1, namespace string) settingsv1alpha1.OpenIDConnectPresetInterface {
	return &fakeOpenIDConnectPresets{
		gentype.NewFakeClientWithList[*v1alpha1.OpenIDConnectPreset, *v1alpha1.OpenIDConnectPresetList](
			fake.Fake,
			namespace,
			v1alpha1.SchemeGroupVersion.WithResource("openidconnectpresets"),
			v1alpha1.SchemeGroupVersion.WithKind("OpenIDConnectPreset"),
			func() *v1alpha1.OpenIDConnectPreset { return &v1alpha1.OpenIDConnectPreset{} },
			func() *v1alpha1.OpenIDConnectPresetList { return &v1alpha1.OpenIDConnectPresetList{} },
			func(dst, src *v1alpha1.OpenIDConnectPresetList) { dst.ListMeta = src.ListMeta },
			func(list *v1alpha1.OpenIDConnectPresetList) []*v1alpha1.OpenIDConnectPreset {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1alpha1.OpenIDConnectPresetList, items []*v1alpha1.OpenIDConnectPreset) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}

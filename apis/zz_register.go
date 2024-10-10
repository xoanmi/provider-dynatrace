// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/xoanmi/provider-dynatrace/apis/alerting/v1alpha1"
	v1alpha1dashboard "github.com/xoanmi/provider-dynatrace/apis/dashboard/v1alpha1"
	v1alpha1event "github.com/xoanmi/provider-dynatrace/apis/event/v1alpha1"
	v1alpha1notification "github.com/xoanmi/provider-dynatrace/apis/notification/v1alpha1"
	v1alpha1apis "github.com/xoanmi/provider-dynatrace/apis/v1alpha1"
	v1beta1 "github.com/xoanmi/provider-dynatrace/apis/v1beta1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1dashboard.SchemeBuilder.AddToScheme,
		v1alpha1event.SchemeBuilder.AddToScheme,
		v1alpha1notification.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1beta1.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}

//go:build !ignore_autogenerated

/*
SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and secret-generator-cop contributors
SPDX-License-Identifier: Apache-2.0
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretGenerator) DeepCopyInto(out *SecretGenerator) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretGenerator.
func (in *SecretGenerator) DeepCopy() *SecretGenerator {
	if in == nil {
		return nil
	}
	out := new(SecretGenerator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretGenerator) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretGeneratorList) DeepCopyInto(out *SecretGeneratorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SecretGenerator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretGeneratorList.
func (in *SecretGeneratorList) DeepCopy() *SecretGeneratorList {
	if in == nil {
		return nil
	}
	out := new(SecretGeneratorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretGeneratorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretGeneratorSpec) DeepCopyInto(out *SecretGeneratorSpec) {
	*out = *in
	out.Spec = in.Spec
	out.Image = in.Image
	in.KubernetesProperties.DeepCopyInto(&out.KubernetesProperties)
	if in.ObjectSelector != nil {
		in, out := &in.ObjectSelector, &out.ObjectSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.NamespaceSelector != nil {
		in, out := &in.NamespaceSelector, &out.NamespaceSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretGeneratorSpec.
func (in *SecretGeneratorSpec) DeepCopy() *SecretGeneratorSpec {
	if in == nil {
		return nil
	}
	out := new(SecretGeneratorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretGeneratorStatus) DeepCopyInto(out *SecretGeneratorStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretGeneratorStatus.
func (in *SecretGeneratorStatus) DeepCopy() *SecretGeneratorStatus {
	if in == nil {
		return nil
	}
	out := new(SecretGeneratorStatus)
	in.DeepCopyInto(out)
	return out
}

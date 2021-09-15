// +build !ignore_autogenerated

/*
 * Copyright 2019 The original author or authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha2

import (
	v1alpha1 "github.com/pivotal/kpack/pkg/apis/core/v1alpha1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Binding) DeepCopyInto(out *Binding) {
	*out = *in
	if in.MetadataRef != nil {
		in, out := &in.MetadataRef, &out.MetadataRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Binding.
func (in *Binding) DeepCopy() *Binding {
	if in == nil {
		return nil
	}
	out := new(Binding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Bindings) DeepCopyInto(out *Bindings) {
	{
		in := &in
		*out = make(Bindings, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bindings.
func (in Bindings) DeepCopy() Bindings {
	if in == nil {
		return nil
	}
	out := new(Bindings)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Build) DeepCopyInto(out *Build) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Build.
func (in *Build) DeepCopy() *Build {
	if in == nil {
		return nil
	}
	out := new(Build)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Build) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildCache) DeepCopyInto(out *BuildCache) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildCache.
func (in *BuildCache) DeepCopy() *BuildCache {
	if in == nil {
		return nil
	}
	out := new(BuildCache)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildCacheConfig) DeepCopyInto(out *BuildCacheConfig) {
	*out = *in
	if in.Volume != nil {
		in, out := &in.Volume, &out.Volume
		*out = new(BuildPersistentVolumeCache)
		**out = **in
	}
	if in.Registry != nil {
		in, out := &in.Registry, &out.Registry
		*out = new(RegistryCache)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildCacheConfig.
func (in *BuildCacheConfig) DeepCopy() *BuildCacheConfig {
	if in == nil {
		return nil
	}
	out := new(BuildCacheConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildList) DeepCopyInto(out *BuildList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Build, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildList.
func (in *BuildList) DeepCopy() *BuildList {
	if in == nil {
		return nil
	}
	out := new(BuildList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BuildList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildPersistentVolumeCache) DeepCopyInto(out *BuildPersistentVolumeCache) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildPersistentVolumeCache.
func (in *BuildPersistentVolumeCache) DeepCopy() *BuildPersistentVolumeCache {
	if in == nil {
		return nil
	}
	out := new(BuildPersistentVolumeCache)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildPodBuilderConfig) DeepCopyInto(out *BuildPodBuilderConfig) {
	*out = *in
	if in.PlatformAPIs != nil {
		in, out := &in.PlatformAPIs, &out.PlatformAPIs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildPodBuilderConfig.
func (in *BuildPodBuilderConfig) DeepCopy() *BuildPodBuilderConfig {
	if in == nil {
		return nil
	}
	out := new(BuildPodBuilderConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildPodImages) DeepCopyInto(out *BuildPodImages) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildPodImages.
func (in *BuildPodImages) DeepCopy() *BuildPodImages {
	if in == nil {
		return nil
	}
	out := new(BuildPodImages)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildSpec) DeepCopyInto(out *BuildSpec) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Builder.DeepCopyInto(&out.Builder)
	in.Source.DeepCopyInto(&out.Source)
	if in.Cache != nil {
		in, out := &in.Cache, &out.Cache
		*out = new(BuildCacheConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Bindings != nil {
		in, out := &in.Bindings, &out.Bindings
		*out = make(v1alpha1.Bindings, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.LastBuild != nil {
		in, out := &in.LastBuild, &out.LastBuild
		*out = new(LastBuild)
		**out = **in
	}
	if in.Notary != nil {
		in, out := &in.Notary, &out.Notary
		*out = new(v1alpha1.NotaryConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildSpec.
func (in *BuildSpec) DeepCopy() *BuildSpec {
	if in == nil {
		return nil
	}
	out := new(BuildSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildStack) DeepCopyInto(out *BuildStack) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildStack.
func (in *BuildStack) DeepCopy() *BuildStack {
	if in == nil {
		return nil
	}
	out := new(BuildStack)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildStatus) DeepCopyInto(out *BuildStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.BuildMetadata != nil {
		in, out := &in.BuildMetadata, &out.BuildMetadata
		*out = make(v1alpha1.BuildpackMetadataList, len(*in))
		copy(*out, *in)
	}
	out.Stack = in.Stack
	if in.StepStates != nil {
		in, out := &in.StepStates, &out.StepStates
		*out = make([]v1.ContainerState, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StepsCompleted != nil {
		in, out := &in.StepsCompleted, &out.StepsCompleted
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildStatus.
func (in *BuildStatus) DeepCopy() *BuildStatus {
	if in == nil {
		return nil
	}
	out := new(BuildStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Builder) DeepCopyInto(out *Builder) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Builder.
func (in *Builder) DeepCopy() *Builder {
	if in == nil {
		return nil
	}
	out := new(Builder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObjectMetaAccessor is an autogenerated deepcopy function, copying the receiver, creating a new metav1.ObjectMetaAccessor.
func (in *Builder) DeepCopyObjectMetaAccessor() metav1.ObjectMetaAccessor {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Builder) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuilderList) DeepCopyInto(out *BuilderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Builder, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuilderList.
func (in *BuilderList) DeepCopy() *BuilderList {
	if in == nil {
		return nil
	}
	out := new(BuilderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BuilderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuilderRecord) DeepCopyInto(out *BuilderRecord) {
	*out = *in
	out.Stack = in.Stack
	if in.Buildpacks != nil {
		in, out := &in.Buildpacks, &out.Buildpacks
		*out = make(v1alpha1.BuildpackMetadataList, len(*in))
		copy(*out, *in)
	}
	if in.Order != nil {
		in, out := &in.Order, &out.Order
		*out = make([]v1alpha1.OrderEntry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuilderRecord.
func (in *BuilderRecord) DeepCopy() *BuilderRecord {
	if in == nil {
		return nil
	}
	out := new(BuilderRecord)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuilderSpec) DeepCopyInto(out *BuilderSpec) {
	*out = *in
	out.Stack = in.Stack
	out.Store = in.Store
	if in.Order != nil {
		in, out := &in.Order, &out.Order
		*out = make([]v1alpha1.OrderEntry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuilderSpec.
func (in *BuilderSpec) DeepCopy() *BuilderSpec {
	if in == nil {
		return nil
	}
	out := new(BuilderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuilderStatus) DeepCopyInto(out *BuilderStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.BuilderMetadata != nil {
		in, out := &in.BuilderMetadata, &out.BuilderMetadata
		*out = make(v1alpha1.BuildpackMetadataList, len(*in))
		copy(*out, *in)
	}
	if in.Order != nil {
		in, out := &in.Order, &out.Order
		*out = make([]v1alpha1.OrderEntry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Stack = in.Stack
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuilderStatus.
func (in *BuilderStatus) DeepCopy() *BuilderStatus {
	if in == nil {
		return nil
	}
	out := new(BuilderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBuilder) DeepCopyInto(out *ClusterBuilder) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBuilder.
func (in *ClusterBuilder) DeepCopy() *ClusterBuilder {
	if in == nil {
		return nil
	}
	out := new(ClusterBuilder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObjectMetaAccessor is an autogenerated deepcopy function, copying the receiver, creating a new metav1.ObjectMetaAccessor.
func (in *ClusterBuilder) DeepCopyObjectMetaAccessor() metav1.ObjectMetaAccessor {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterBuilder) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBuilderList) DeepCopyInto(out *ClusterBuilderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterBuilder, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBuilderList.
func (in *ClusterBuilderList) DeepCopy() *ClusterBuilderList {
	if in == nil {
		return nil
	}
	out := new(ClusterBuilderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterBuilderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBuilderSpec) DeepCopyInto(out *ClusterBuilderSpec) {
	*out = *in
	in.BuilderSpec.DeepCopyInto(&out.BuilderSpec)
	out.ServiceAccountRef = in.ServiceAccountRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBuilderSpec.
func (in *ClusterBuilderSpec) DeepCopy() *ClusterBuilderSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterBuilderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStack) DeepCopyInto(out *ClusterStack) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStack.
func (in *ClusterStack) DeepCopy() *ClusterStack {
	if in == nil {
		return nil
	}
	out := new(ClusterStack)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObjectMetaAccessor is an autogenerated deepcopy function, copying the receiver, creating a new metav1.ObjectMetaAccessor.
func (in *ClusterStack) DeepCopyObjectMetaAccessor() metav1.ObjectMetaAccessor {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterStack) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStackList) DeepCopyInto(out *ClusterStackList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterStack, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStackList.
func (in *ClusterStackList) DeepCopy() *ClusterStackList {
	if in == nil {
		return nil
	}
	out := new(ClusterStackList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterStackList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStackSpec) DeepCopyInto(out *ClusterStackSpec) {
	*out = *in
	out.BuildImage = in.BuildImage
	out.RunImage = in.RunImage
	if in.ServiceAccountRef != nil {
		in, out := &in.ServiceAccountRef, &out.ServiceAccountRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStackSpec.
func (in *ClusterStackSpec) DeepCopy() *ClusterStackSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterStackSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStackSpecImage) DeepCopyInto(out *ClusterStackSpecImage) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStackSpecImage.
func (in *ClusterStackSpecImage) DeepCopy() *ClusterStackSpecImage {
	if in == nil {
		return nil
	}
	out := new(ClusterStackSpecImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStackStatus) DeepCopyInto(out *ClusterStackStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	in.ResolvedClusterStack.DeepCopyInto(&out.ResolvedClusterStack)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStackStatus.
func (in *ClusterStackStatus) DeepCopy() *ClusterStackStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStackStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStackStatusImage) DeepCopyInto(out *ClusterStackStatusImage) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStackStatusImage.
func (in *ClusterStackStatusImage) DeepCopy() *ClusterStackStatusImage {
	if in == nil {
		return nil
	}
	out := new(ClusterStackStatusImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStore) DeepCopyInto(out *ClusterStore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStore.
func (in *ClusterStore) DeepCopy() *ClusterStore {
	if in == nil {
		return nil
	}
	out := new(ClusterStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObjectMetaAccessor is an autogenerated deepcopy function, copying the receiver, creating a new metav1.ObjectMetaAccessor.
func (in *ClusterStore) DeepCopyObjectMetaAccessor() metav1.ObjectMetaAccessor {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterStore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStoreList) DeepCopyInto(out *ClusterStoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterStore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStoreList.
func (in *ClusterStoreList) DeepCopy() *ClusterStoreList {
	if in == nil {
		return nil
	}
	out := new(ClusterStoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterStoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStoreSpec) DeepCopyInto(out *ClusterStoreSpec) {
	*out = *in
	if in.Sources != nil {
		in, out := &in.Sources, &out.Sources
		*out = make([]v1alpha1.StoreImage, len(*in))
		copy(*out, *in)
	}
	if in.ServiceAccountRef != nil {
		in, out := &in.ServiceAccountRef, &out.ServiceAccountRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStoreSpec.
func (in *ClusterStoreSpec) DeepCopy() *ClusterStoreSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterStoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStoreStatus) DeepCopyInto(out *ClusterStoreStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.Buildpacks != nil {
		in, out := &in.Buildpacks, &out.Buildpacks
		*out = make([]v1alpha1.StoreBuildpack, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStoreStatus.
func (in *ClusterStoreStatus) DeepCopy() *ClusterStoreStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image) DeepCopyInto(out *Image) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image.
func (in *Image) DeepCopy() *Image {
	if in == nil {
		return nil
	}
	out := new(Image)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Image) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageBuilder) DeepCopyInto(out *ImageBuilder) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageBuilder.
func (in *ImageBuilder) DeepCopy() *ImageBuilder {
	if in == nil {
		return nil
	}
	out := new(ImageBuilder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageCacheConfig) DeepCopyInto(out *ImageCacheConfig) {
	*out = *in
	if in.Volume != nil {
		in, out := &in.Volume, &out.Volume
		*out = new(ImagePersistentVolumeCache)
		(*in).DeepCopyInto(*out)
	}
	if in.Registry != nil {
		in, out := &in.Registry, &out.Registry
		*out = new(RegistryCache)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageCacheConfig.
func (in *ImageCacheConfig) DeepCopy() *ImageCacheConfig {
	if in == nil {
		return nil
	}
	out := new(ImageCacheConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageList) DeepCopyInto(out *ImageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Image, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageList.
func (in *ImageList) DeepCopy() *ImageList {
	if in == nil {
		return nil
	}
	out := new(ImageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagePersistentVolumeCache) DeepCopyInto(out *ImagePersistentVolumeCache) {
	*out = *in
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		x := (*in).DeepCopy()
		*out = &x
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagePersistentVolumeCache.
func (in *ImagePersistentVolumeCache) DeepCopy() *ImagePersistentVolumeCache {
	if in == nil {
		return nil
	}
	out := new(ImagePersistentVolumeCache)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSpec) DeepCopyInto(out *ImageSpec) {
	*out = *in
	out.Builder = in.Builder
	in.Source.DeepCopyInto(&out.Source)
	if in.Cache != nil {
		in, out := &in.Cache, &out.Cache
		*out = new(ImageCacheConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.FailedBuildHistoryLimit != nil {
		in, out := &in.FailedBuildHistoryLimit, &out.FailedBuildHistoryLimit
		*out = new(int64)
		**out = **in
	}
	if in.SuccessBuildHistoryLimit != nil {
		in, out := &in.SuccessBuildHistoryLimit, &out.SuccessBuildHistoryLimit
		*out = new(int64)
		**out = **in
	}
	if in.Build != nil {
		in, out := &in.Build, &out.Build
		*out = new(v1alpha1.ImageBuild)
		(*in).DeepCopyInto(*out)
	}
	if in.Notary != nil {
		in, out := &in.Notary, &out.Notary
		*out = new(v1alpha1.NotaryConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.AdditionalTags != nil {
		in, out := &in.AdditionalTags, &out.AdditionalTags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSpec.
func (in *ImageSpec) DeepCopy() *ImageSpec {
	if in == nil {
		return nil
	}
	out := new(ImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageStatus) DeepCopyInto(out *ImageStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageStatus.
func (in *ImageStatus) DeepCopy() *ImageStatus {
	if in == nil {
		return nil
	}
	out := new(ImageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LastBuild) DeepCopyInto(out *LastBuild) {
	*out = *in
	out.Cache = in.Cache
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LastBuild.
func (in *LastBuild) DeepCopy() *LastBuild {
	if in == nil {
		return nil
	}
	out := new(LastBuild)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedBuilderSpec) DeepCopyInto(out *NamespacedBuilderSpec) {
	*out = *in
	in.BuilderSpec.DeepCopyInto(&out.BuilderSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedBuilderSpec.
func (in *NamespacedBuilderSpec) DeepCopy() *NamespacedBuilderSpec {
	if in == nil {
		return nil
	}
	out := new(NamespacedBuilderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryCache) DeepCopyInto(out *RegistryCache) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryCache.
func (in *RegistryCache) DeepCopy() *RegistryCache {
	if in == nil {
		return nil
	}
	out := new(RegistryCache)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolvedClusterStack) DeepCopyInto(out *ResolvedClusterStack) {
	*out = *in
	out.BuildImage = in.BuildImage
	out.RunImage = in.RunImage
	if in.Mixins != nil {
		in, out := &in.Mixins, &out.Mixins
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolvedClusterStack.
func (in *ResolvedClusterStack) DeepCopy() *ResolvedClusterStack {
	if in == nil {
		return nil
	}
	out := new(ResolvedClusterStack)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceResolver) DeepCopyInto(out *SourceResolver) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceResolver.
func (in *SourceResolver) DeepCopy() *SourceResolver {
	if in == nil {
		return nil
	}
	out := new(SourceResolver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SourceResolver) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceResolverList) DeepCopyInto(out *SourceResolverList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SourceResolver, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceResolverList.
func (in *SourceResolverList) DeepCopy() *SourceResolverList {
	if in == nil {
		return nil
	}
	out := new(SourceResolverList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SourceResolverList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceResolverSpec) DeepCopyInto(out *SourceResolverSpec) {
	*out = *in
	in.Source.DeepCopyInto(&out.Source)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceResolverSpec.
func (in *SourceResolverSpec) DeepCopy() *SourceResolverSpec {
	if in == nil {
		return nil
	}
	out := new(SourceResolverSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceResolverStatus) DeepCopyInto(out *SourceResolverStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	in.Source.DeepCopyInto(&out.Source)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceResolverStatus.
func (in *SourceResolverStatus) DeepCopy() *SourceResolverStatus {
	if in == nil {
		return nil
	}
	out := new(SourceResolverStatus)
	in.DeepCopyInto(out)
	return out
}

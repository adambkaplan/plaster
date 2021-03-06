// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	v1 "github.com/openshift/api/build/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlasterBuild) DeepCopyInto(out *PlasterBuild) {
	*out = *in
	if in.Dockerfile != nil {
		in, out := &in.Dockerfile, &out.Dockerfile
		*out = new(PlasterDockerfileBuild)
		**out = **in
	}
	if in.Language != nil {
		in, out := &in.Language, &out.Language
		*out = new(PlasterLanguageBuild)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlasterBuild.
func (in *PlasterBuild) DeepCopy() *PlasterBuild {
	if in == nil {
		return nil
	}
	out := new(PlasterBuild)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlasterDockerfileBuild) DeepCopyInto(out *PlasterDockerfileBuild) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlasterDockerfileBuild.
func (in *PlasterDockerfileBuild) DeepCopy() *PlasterDockerfileBuild {
	if in == nil {
		return nil
	}
	out := new(PlasterDockerfileBuild)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlasterLanguageBuild) DeepCopyInto(out *PlasterLanguageBuild) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlasterLanguageBuild.
func (in *PlasterLanguageBuild) DeepCopy() *PlasterLanguageBuild {
	if in == nil {
		return nil
	}
	out := new(PlasterLanguageBuild)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlasterProject) DeepCopyInto(out *PlasterProject) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlasterProject.
func (in *PlasterProject) DeepCopy() *PlasterProject {
	if in == nil {
		return nil
	}
	out := new(PlasterProject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PlasterProject) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlasterProjectCondition) DeepCopyInto(out *PlasterProjectCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlasterProjectCondition.
func (in *PlasterProjectCondition) DeepCopy() *PlasterProjectCondition {
	if in == nil {
		return nil
	}
	out := new(PlasterProjectCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlasterProjectList) DeepCopyInto(out *PlasterProjectList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PlasterProject, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlasterProjectList.
func (in *PlasterProjectList) DeepCopy() *PlasterProjectList {
	if in == nil {
		return nil
	}
	out := new(PlasterProjectList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PlasterProjectList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlasterProjectSpec) DeepCopyInto(out *PlasterProjectSpec) {
	*out = *in
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(PlasterSource)
		(*in).DeepCopyInto(*out)
	}
	if in.Build != nil {
		in, out := &in.Build, &out.Build
		*out = new(PlasterBuild)
		(*in).DeepCopyInto(*out)
	}
	if in.Tests != nil {
		in, out := &in.Tests, &out.Tests
		*out = make([]PlasterTest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlasterProjectSpec.
func (in *PlasterProjectSpec) DeepCopy() *PlasterProjectSpec {
	if in == nil {
		return nil
	}
	out := new(PlasterProjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlasterProjectStatus) DeepCopyInto(out *PlasterProjectStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]PlasterProjectCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlasterProjectStatus.
func (in *PlasterProjectStatus) DeepCopy() *PlasterProjectStatus {
	if in == nil {
		return nil
	}
	out := new(PlasterProjectStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlasterSource) DeepCopyInto(out *PlasterSource) {
	*out = *in
	if in.Git != nil {
		in, out := &in.Git, &out.Git
		*out = new(v1.GitBuildSource)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlasterSource.
func (in *PlasterSource) DeepCopy() *PlasterSource {
	if in == nil {
		return nil
	}
	out := new(PlasterSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlasterTest) DeepCopyInto(out *PlasterTest) {
	*out = *in
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlasterTest.
func (in *PlasterTest) DeepCopy() *PlasterTest {
	if in == nil {
		return nil
	}
	out := new(PlasterTest)
	in.DeepCopyInto(out)
	return out
}

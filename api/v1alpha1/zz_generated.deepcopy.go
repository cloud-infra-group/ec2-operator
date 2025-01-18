//go:build !ignore_autogenerated

/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSVPCEndpointService) DeepCopyInto(out *AWSVPCEndpointService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSVPCEndpointService.
func (in *AWSVPCEndpointService) DeepCopy() *AWSVPCEndpointService {
	if in == nil {
		return nil
	}
	out := new(AWSVPCEndpointService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSVPCEndpointService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSVPCEndpointServiceList) DeepCopyInto(out *AWSVPCEndpointServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AWSVPCEndpointService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSVPCEndpointServiceList.
func (in *AWSVPCEndpointServiceList) DeepCopy() *AWSVPCEndpointServiceList {
	if in == nil {
		return nil
	}
	out := new(AWSVPCEndpointServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSVPCEndpointServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSVPCEndpointServicePrincipal) DeepCopyInto(out *AWSVPCEndpointServicePrincipal) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSVPCEndpointServicePrincipal.
func (in *AWSVPCEndpointServicePrincipal) DeepCopy() *AWSVPCEndpointServicePrincipal {
	if in == nil {
		return nil
	}
	out := new(AWSVPCEndpointServicePrincipal)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSVPCEndpointServicePrincipal) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSVPCEndpointServicePrincipalList) DeepCopyInto(out *AWSVPCEndpointServicePrincipalList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AWSVPCEndpointServicePrincipal, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSVPCEndpointServicePrincipalList.
func (in *AWSVPCEndpointServicePrincipalList) DeepCopy() *AWSVPCEndpointServicePrincipalList {
	if in == nil {
		return nil
	}
	out := new(AWSVPCEndpointServicePrincipalList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSVPCEndpointServicePrincipalList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSVPCEndpointServicePrincipalSpec) DeepCopyInto(out *AWSVPCEndpointServicePrincipalSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSVPCEndpointServicePrincipalSpec.
func (in *AWSVPCEndpointServicePrincipalSpec) DeepCopy() *AWSVPCEndpointServicePrincipalSpec {
	if in == nil {
		return nil
	}
	out := new(AWSVPCEndpointServicePrincipalSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSVPCEndpointServicePrincipalStatus) DeepCopyInto(out *AWSVPCEndpointServicePrincipalStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSVPCEndpointServicePrincipalStatus.
func (in *AWSVPCEndpointServicePrincipalStatus) DeepCopy() *AWSVPCEndpointServicePrincipalStatus {
	if in == nil {
		return nil
	}
	out := new(AWSVPCEndpointServicePrincipalStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSVPCEndpointServiceSpec) DeepCopyInto(out *AWSVPCEndpointServiceSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSVPCEndpointServiceSpec.
func (in *AWSVPCEndpointServiceSpec) DeepCopy() *AWSVPCEndpointServiceSpec {
	if in == nil {
		return nil
	}
	out := new(AWSVPCEndpointServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSVPCEndpointServiceStatus) DeepCopyInto(out *AWSVPCEndpointServiceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSVPCEndpointServiceStatus.
func (in *AWSVPCEndpointServiceStatus) DeepCopy() *AWSVPCEndpointServiceStatus {
	if in == nil {
		return nil
	}
	out := new(AWSVPCEndpointServiceStatus)
	in.DeepCopyInto(out)
	return out
}

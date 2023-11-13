//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023 Always Bespoke LLC

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequestLine) DeepCopyInto(out *RequestLine) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequestLine.
func (in *RequestLine) DeepCopy() *RequestLine {
	if in == nil {
		return nil
	}
	out := new(RequestLine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SIPCall) DeepCopyInto(out *SIPCall) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SIPCall.
func (in *SIPCall) DeepCopy() *SIPCall {
	if in == nil {
		return nil
	}
	out := new(SIPCall)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SIPCall) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SIPCallList) DeepCopyInto(out *SIPCallList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SIPCall, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SIPCallList.
func (in *SIPCallList) DeepCopy() *SIPCallList {
	if in == nil {
		return nil
	}
	out := new(SIPCallList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SIPCallList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SIPCallMessage) DeepCopyInto(out *SIPCallMessage) {
	*out = *in
	in.Created.DeepCopyInto(&out.Created)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SIPCallMessage.
func (in *SIPCallMessage) DeepCopy() *SIPCallMessage {
	if in == nil {
		return nil
	}
	out := new(SIPCallMessage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SIPCallSpec) DeepCopyInto(out *SIPCallSpec) {
	*out = *in
	if in.Messages != nil {
		in, out := &in.Messages, &out.Messages
		*out = make([]SIPCallMessage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SIPCallSpec.
func (in *SIPCallSpec) DeepCopy() *SIPCallSpec {
	if in == nil {
		return nil
	}
	out := new(SIPCallSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SIPCallStatus) DeepCopyInto(out *SIPCallStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SIPCallStatus.
func (in *SIPCallStatus) DeepCopy() *SIPCallStatus {
	if in == nil {
		return nil
	}
	out := new(SIPCallStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SIPMessage) DeepCopyInto(out *SIPMessage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SIPMessage.
func (in *SIPMessage) DeepCopy() *SIPMessage {
	if in == nil {
		return nil
	}
	out := new(SIPMessage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SIPMessage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SIPMessageList) DeepCopyInto(out *SIPMessageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SIPMessage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SIPMessageList.
func (in *SIPMessageList) DeepCopy() *SIPMessageList {
	if in == nil {
		return nil
	}
	out := new(SIPMessageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SIPMessageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SIPMessageSpec) DeepCopyInto(out *SIPMessageSpec) {
	*out = *in
	if in.RequestLine != nil {
		in, out := &in.RequestLine, &out.RequestLine
		*out = new(RequestLine)
		**out = **in
	}
	if in.StatusLine != nil {
		in, out := &in.StatusLine, &out.StatusLine
		*out = new(StatusLine)
		**out = **in
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.Body != nil {
		in, out := &in.Body, &out.Body
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SIPMessageSpec.
func (in *SIPMessageSpec) DeepCopy() *SIPMessageSpec {
	if in == nil {
		return nil
	}
	out := new(SIPMessageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SIPMessageStatus) DeepCopyInto(out *SIPMessageStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SIPMessageStatus.
func (in *SIPMessageStatus) DeepCopy() *SIPMessageStatus {
	if in == nil {
		return nil
	}
	out := new(SIPMessageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusLine) DeepCopyInto(out *StatusLine) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusLine.
func (in *StatusLine) DeepCopy() *StatusLine {
	if in == nil {
		return nil
	}
	out := new(StatusLine)
	in.DeepCopyInto(out)
	return out
}

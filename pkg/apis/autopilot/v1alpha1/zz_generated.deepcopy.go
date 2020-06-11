// +build !ignore_autogenerated

/*
Copyright 2019 Openstorage.org

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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionPreviewObjectMetadata) DeepCopyInto(out *ActionPreviewObjectMetadata) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionPreviewObjectMetadata.
func (in *ActionPreviewObjectMetadata) DeepCopy() *ActionPreviewObjectMetadata {
	if in == nil {
		return nil
	}
	out := new(ActionPreviewObjectMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutopilotActionApproval) DeepCopyInto(out *AutopilotActionApproval) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Action.DeepCopyInto(&out.Action)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutopilotActionApproval.
func (in *AutopilotActionApproval) DeepCopy() *AutopilotActionApproval {
	if in == nil {
		return nil
	}
	out := new(AutopilotActionApproval)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutopilotActionPreview) DeepCopyInto(out *AutopilotActionPreview) {
	*out = *in
	in.ObjectMetadata.DeepCopyInto(&out.ObjectMetadata)
	in.RuleAction.DeepCopyInto(&out.RuleAction)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutopilotActionPreview.
func (in *AutopilotActionPreview) DeepCopy() *AutopilotActionPreview {
	if in == nil {
		return nil
	}
	out := new(AutopilotActionPreview)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutopilotRule) DeepCopyInto(out *AutopilotRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutopilotRule.
func (in *AutopilotRule) DeepCopy() *AutopilotRule {
	if in == nil {
		return nil
	}
	out := new(AutopilotRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AutopilotRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutopilotRuleList) DeepCopyInto(out *AutopilotRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AutopilotRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutopilotRuleList.
func (in *AutopilotRuleList) DeepCopy() *AutopilotRuleList {
	if in == nil {
		return nil
	}
	out := new(AutopilotRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AutopilotRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutopilotRuleObject) DeepCopyInto(out *AutopilotRuleObject) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutopilotRuleObject.
func (in *AutopilotRuleObject) DeepCopy() *AutopilotRuleObject {
	if in == nil {
		return nil
	}
	out := new(AutopilotRuleObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AutopilotRuleObject) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutopilotRuleObjectList) DeepCopyInto(out *AutopilotRuleObjectList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AutopilotRuleObject, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutopilotRuleObjectList.
func (in *AutopilotRuleObjectList) DeepCopy() *AutopilotRuleObjectList {
	if in == nil {
		return nil
	}
	out := new(AutopilotRuleObjectList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AutopilotRuleObjectList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutopilotRuleObjectSpec) DeepCopyInto(out *AutopilotRuleObjectSpec) {
	*out = *in
	if in.ActionApprovals != nil {
		in, out := &in.ActionApprovals, &out.ActionApprovals
		*out = make([]*AutopilotActionApproval, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(AutopilotActionApproval)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutopilotRuleObjectSpec.
func (in *AutopilotRuleObjectSpec) DeepCopy() *AutopilotRuleObjectSpec {
	if in == nil {
		return nil
	}
	out := new(AutopilotRuleObjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutopilotRuleObjectStatus) DeepCopyInto(out *AutopilotRuleObjectStatus) {
	*out = *in
	in.CurrentStatus.DeepCopyInto(&out.CurrentStatus)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]*AutopilotRuleObjectStatusItem, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(AutopilotRuleObjectStatusItem)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutopilotRuleObjectStatus.
func (in *AutopilotRuleObjectStatus) DeepCopy() *AutopilotRuleObjectStatus {
	if in == nil {
		return nil
	}
	out := new(AutopilotRuleObjectStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutopilotRuleObjectStatusItem) DeepCopyInto(out *AutopilotRuleObjectStatusItem) {
	*out = *in
	in.LastProcessTimestamp.DeepCopyInto(&out.LastProcessTimestamp)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutopilotRuleObjectStatusItem.
func (in *AutopilotRuleObjectStatusItem) DeepCopy() *AutopilotRuleObjectStatusItem {
	if in == nil {
		return nil
	}
	out := new(AutopilotRuleObjectStatusItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutopilotRuleSpec) DeepCopyInto(out *AutopilotRuleSpec) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	in.NamespaceSelector.DeepCopyInto(&out.NamespaceSelector)
	in.Conditions.DeepCopyInto(&out.Conditions)
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]*RuleAction, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(RuleAction)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutopilotRuleSpec.
func (in *AutopilotRuleSpec) DeepCopy() *AutopilotRuleSpec {
	if in == nil {
		return nil
	}
	out := new(AutopilotRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LabelSelectorRequirement) DeepCopyInto(out *LabelSelectorRequirement) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LabelSelectorRequirement.
func (in *LabelSelectorRequirement) DeepCopy() *LabelSelectorRequirement {
	if in == nil {
		return nil
	}
	out := new(LabelSelectorRequirement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleAction) DeepCopyInto(out *RuleAction) {
	*out = *in
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleAction.
func (in *RuleAction) DeepCopy() *RuleAction {
	if in == nil {
		return nil
	}
	out := new(RuleAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleConditions) DeepCopyInto(out *RuleConditions) {
	*out = *in
	if in.Expressions != nil {
		in, out := &in.Expressions, &out.Expressions
		*out = make([]*LabelSelectorRequirement, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(LabelSelectorRequirement)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleConditions.
func (in *RuleConditions) DeepCopy() *RuleConditions {
	if in == nil {
		return nil
	}
	out := new(RuleConditions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleObjectSelector) DeepCopyInto(out *RuleObjectSelector) {
	*out = *in
	in.LabelSelector.DeepCopyInto(&out.LabelSelector)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleObjectSelector.
func (in *RuleObjectSelector) DeepCopy() *RuleObjectSelector {
	if in == nil {
		return nil
	}
	out := new(RuleObjectSelector)
	in.DeepCopyInto(out)
	return out
}

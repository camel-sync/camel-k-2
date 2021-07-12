// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"encoding/json"
	"github.com/apache/camel-k/pkg/apis/camel/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthorizationSpec) DeepCopyInto(out *AuthorizationSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthorizationSpec.
func (in *AuthorizationSpec) DeepCopy() *AuthorizationSpec {
	if in == nil {
		return nil
	}
	out := new(AuthorizationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BeanProperties) DeepCopyInto(out *BeanProperties) {
	*out = *in
	if in.RawMessage != nil {
		in, out := &in.RawMessage, &out.RawMessage
		*out = make(v1.RawMessage, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BeanProperties.
func (in *BeanProperties) DeepCopy() *BeanProperties {
	if in == nil {
		return nil
	}
	out := new(BeanProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Endpoint) DeepCopyInto(out *Endpoint) {
	*out = *in
	if in.Ref != nil {
		in, out := &in.Ref, &out.Ref
		*out = new(corev1.ObjectReference)
		**out = **in
	}
	if in.URI != nil {
		in, out := &in.URI, &out.URI
		*out = new(string)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(EndpointProperties)
		(*in).DeepCopyInto(*out)
	}
	if in.Types != nil {
		in, out := &in.Types, &out.Types
		*out = make(map[EventSlot]EventTypeSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Endpoint.
func (in *Endpoint) DeepCopy() *Endpoint {
	if in == nil {
		return nil
	}
	out := new(Endpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointProperties) DeepCopyInto(out *EndpointProperties) {
	*out = *in
	if in.RawMessage != nil {
		in, out := &in.RawMessage, &out.RawMessage
		*out = make(v1.RawMessage, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointProperties.
func (in *EndpointProperties) DeepCopy() *EndpointProperties {
	if in == nil {
		return nil
	}
	out := new(EndpointProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ErrorHandlerBean) DeepCopyInto(out *ErrorHandlerBean) {
	*out = *in
	out.ErrorHandlerNone = in.ErrorHandlerNone
	if in.BeanType != nil {
		in, out := &in.BeanType, &out.BeanType
		*out = new(string)
		**out = **in
	}
	if in.BeanProperties != nil {
		in, out := &in.BeanProperties, &out.BeanProperties
		*out = new(BeanProperties)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrorHandlerBean.
func (in *ErrorHandlerBean) DeepCopy() *ErrorHandlerBean {
	if in == nil {
		return nil
	}
	out := new(ErrorHandlerBean)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ErrorHandlerDeadLetterChannel) DeepCopyInto(out *ErrorHandlerDeadLetterChannel) {
	*out = *in
	in.ErrorHandlerLog.DeepCopyInto(&out.ErrorHandlerLog)
	if in.DLCEndpoint != nil {
		in, out := &in.DLCEndpoint, &out.DLCEndpoint
		*out = new(Endpoint)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrorHandlerDeadLetterChannel.
func (in *ErrorHandlerDeadLetterChannel) DeepCopy() *ErrorHandlerDeadLetterChannel {
	if in == nil {
		return nil
	}
	out := new(ErrorHandlerDeadLetterChannel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ErrorHandlerLog) DeepCopyInto(out *ErrorHandlerLog) {
	*out = *in
	out.ErrorHandlerNone = in.ErrorHandlerNone
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = new(ErrorHandlerParameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrorHandlerLog.
func (in *ErrorHandlerLog) DeepCopy() *ErrorHandlerLog {
	if in == nil {
		return nil
	}
	out := new(ErrorHandlerLog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ErrorHandlerNone) DeepCopyInto(out *ErrorHandlerNone) {
	*out = *in
	out.baseErrorHandler = in.baseErrorHandler
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrorHandlerNone.
func (in *ErrorHandlerNone) DeepCopy() *ErrorHandlerNone {
	if in == nil {
		return nil
	}
	out := new(ErrorHandlerNone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ErrorHandlerParameters) DeepCopyInto(out *ErrorHandlerParameters) {
	*out = *in
	if in.RawMessage != nil {
		in, out := &in.RawMessage, &out.RawMessage
		*out = make(v1.RawMessage, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrorHandlerParameters.
func (in *ErrorHandlerParameters) DeepCopy() *ErrorHandlerParameters {
	if in == nil {
		return nil
	}
	out := new(ErrorHandlerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ErrorHandlerRef) DeepCopyInto(out *ErrorHandlerRef) {
	*out = *in
	out.baseErrorHandler = in.baseErrorHandler
	if in.RawMessage != nil {
		in, out := &in.RawMessage, &out.RawMessage
		*out = make(v1.RawMessage, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrorHandlerRef.
func (in *ErrorHandlerRef) DeepCopy() *ErrorHandlerRef {
	if in == nil {
		return nil
	}
	out := new(ErrorHandlerRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ErrorHandlerSpec) DeepCopyInto(out *ErrorHandlerSpec) {
	*out = *in
	if in.RawMessage != nil {
		in, out := &in.RawMessage, &out.RawMessage
		*out = make(v1.RawMessage, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrorHandlerSpec.
func (in *ErrorHandlerSpec) DeepCopy() *ErrorHandlerSpec {
	if in == nil {
		return nil
	}
	out := new(ErrorHandlerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventTypeSpec) DeepCopyInto(out *EventTypeSpec) {
	*out = *in
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = new(JSONSchemaProps)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventTypeSpec.
func (in *EventTypeSpec) DeepCopy() *EventTypeSpec {
	if in == nil {
		return nil
	}
	out := new(EventTypeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalDocumentation) DeepCopyInto(out *ExternalDocumentation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalDocumentation.
func (in *ExternalDocumentation) DeepCopy() *ExternalDocumentation {
	if in == nil {
		return nil
	}
	out := new(ExternalDocumentation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JSON) DeepCopyInto(out *JSON) {
	*out = *in
	if in.RawMessage != nil {
		in, out := &in.RawMessage, &out.RawMessage
		*out = make(RawMessage, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JSON.
func (in *JSON) DeepCopy() *JSON {
	if in == nil {
		return nil
	}
	out := new(JSON)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JSONSchemaProp) DeepCopyInto(out *JSONSchemaProp) {
	*out = *in
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(JSON)
		(*in).DeepCopyInto(*out)
	}
	if in.Maximum != nil {
		in, out := &in.Maximum, &out.Maximum
		*out = new(json.Number)
		**out = **in
	}
	if in.Minimum != nil {
		in, out := &in.Minimum, &out.Minimum
		*out = new(json.Number)
		**out = **in
	}
	if in.MaxLength != nil {
		in, out := &in.MaxLength, &out.MaxLength
		*out = new(int64)
		**out = **in
	}
	if in.MinLength != nil {
		in, out := &in.MinLength, &out.MinLength
		*out = new(int64)
		**out = **in
	}
	if in.MaxItems != nil {
		in, out := &in.MaxItems, &out.MaxItems
		*out = new(int64)
		**out = **in
	}
	if in.MinItems != nil {
		in, out := &in.MinItems, &out.MinItems
		*out = new(int64)
		**out = **in
	}
	if in.MaxProperties != nil {
		in, out := &in.MaxProperties, &out.MaxProperties
		*out = new(int64)
		**out = **in
	}
	if in.MinProperties != nil {
		in, out := &in.MinProperties, &out.MinProperties
		*out = new(int64)
		**out = **in
	}
	if in.MultipleOf != nil {
		in, out := &in.MultipleOf, &out.MultipleOf
		*out = new(json.Number)
		**out = **in
	}
	if in.Enum != nil {
		in, out := &in.Enum, &out.Enum
		*out = make([]JSON, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Example != nil {
		in, out := &in.Example, &out.Example
		*out = new(JSON)
		(*in).DeepCopyInto(*out)
	}
	if in.XDescriptors != nil {
		in, out := &in.XDescriptors, &out.XDescriptors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JSONSchemaProp.
func (in *JSONSchemaProp) DeepCopy() *JSONSchemaProp {
	if in == nil {
		return nil
	}
	out := new(JSONSchemaProp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JSONSchemaProps) DeepCopyInto(out *JSONSchemaProps) {
	*out = *in
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[string]JSONSchemaProp, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Required != nil {
		in, out := &in.Required, &out.Required
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Example != nil {
		in, out := &in.Example, &out.Example
		*out = new(JSON)
		(*in).DeepCopyInto(*out)
	}
	if in.ExternalDocs != nil {
		in, out := &in.ExternalDocs, &out.ExternalDocs
		*out = new(ExternalDocumentation)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JSONSchemaProps.
func (in *JSONSchemaProps) DeepCopy() *JSONSchemaProps {
	if in == nil {
		return nil
	}
	out := new(JSONSchemaProps)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kamelet) DeepCopyInto(out *Kamelet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kamelet.
func (in *Kamelet) DeepCopy() *Kamelet {
	if in == nil {
		return nil
	}
	out := new(Kamelet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Kamelet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KameletBinding) DeepCopyInto(out *KameletBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KameletBinding.
func (in *KameletBinding) DeepCopy() *KameletBinding {
	if in == nil {
		return nil
	}
	out := new(KameletBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KameletBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KameletBindingCondition) DeepCopyInto(out *KameletBindingCondition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KameletBindingCondition.
func (in *KameletBindingCondition) DeepCopy() *KameletBindingCondition {
	if in == nil {
		return nil
	}
	out := new(KameletBindingCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KameletBindingList) DeepCopyInto(out *KameletBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KameletBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KameletBindingList.
func (in *KameletBindingList) DeepCopy() *KameletBindingList {
	if in == nil {
		return nil
	}
	out := new(KameletBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KameletBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KameletBindingSpec) DeepCopyInto(out *KameletBindingSpec) {
	*out = *in
	if in.Integration != nil {
		in, out := &in.Integration, &out.Integration
		*out = new(v1.IntegrationSpec)
		(*in).DeepCopyInto(*out)
	}
	in.Source.DeepCopyInto(&out.Source)
	in.Sink.DeepCopyInto(&out.Sink)
	if in.ErrorHandler != nil {
		in, out := &in.ErrorHandler, &out.ErrorHandler
		*out = new(ErrorHandlerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]Endpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KameletBindingSpec.
func (in *KameletBindingSpec) DeepCopy() *KameletBindingSpec {
	if in == nil {
		return nil
	}
	out := new(KameletBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KameletBindingStatus) DeepCopyInto(out *KameletBindingStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]KameletBindingCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KameletBindingStatus.
func (in *KameletBindingStatus) DeepCopy() *KameletBindingStatus {
	if in == nil {
		return nil
	}
	out := new(KameletBindingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KameletCondition) DeepCopyInto(out *KameletCondition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KameletCondition.
func (in *KameletCondition) DeepCopy() *KameletCondition {
	if in == nil {
		return nil
	}
	out := new(KameletCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KameletList) DeepCopyInto(out *KameletList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Kamelet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KameletList.
func (in *KameletList) DeepCopy() *KameletList {
	if in == nil {
		return nil
	}
	out := new(KameletList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KameletList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KameletProperty) DeepCopyInto(out *KameletProperty) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KameletProperty.
func (in *KameletProperty) DeepCopy() *KameletProperty {
	if in == nil {
		return nil
	}
	out := new(KameletProperty)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KameletSpec) DeepCopyInto(out *KameletSpec) {
	*out = *in
	if in.Definition != nil {
		in, out := &in.Definition, &out.Definition
		*out = new(JSONSchemaProps)
		(*in).DeepCopyInto(*out)
	}
	if in.Sources != nil {
		in, out := &in.Sources, &out.Sources
		*out = make([]v1.SourceSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Flow != nil {
		in, out := &in.Flow, &out.Flow
		*out = new(v1.Flow)
		(*in).DeepCopyInto(*out)
	}
	if in.Authorization != nil {
		in, out := &in.Authorization, &out.Authorization
		*out = new(AuthorizationSpec)
		**out = **in
	}
	if in.Types != nil {
		in, out := &in.Types, &out.Types
		*out = make(map[EventSlot]EventTypeSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Dependencies != nil {
		in, out := &in.Dependencies, &out.Dependencies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KameletSpec.
func (in *KameletSpec) DeepCopy() *KameletSpec {
	if in == nil {
		return nil
	}
	out := new(KameletSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KameletStatus) DeepCopyInto(out *KameletStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]KameletCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make([]KameletProperty, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KameletStatus.
func (in *KameletStatus) DeepCopy() *KameletStatus {
	if in == nil {
		return nil
	}
	out := new(KameletStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in RawMessage) DeepCopyInto(out *RawMessage) {
	{
		in := &in
		*out = make(RawMessage, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RawMessage.
func (in RawMessage) DeepCopy() RawMessage {
	if in == nil {
		return nil
	}
	out := new(RawMessage)
	in.DeepCopyInto(out)
	return *out
}

package loader

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoiface"
)

var (
	// _ proto.Message        = (*reflectedMessage)(nil)
	_ protoreflect.Message = (*reflectedMessage)(nil)
)

// reflectedMessage
type reflectedMessage struct {
}

// MessageType
//  @param messageType
//  @return protoreflect.Message
//
func MessageType(messageType string) protoreflect.Message {
	return &reflectedMessage{}
}

// Descriptor
//  @receiver r
//  @return protoreflect.MessageDescriptor
//
func (r *reflectedMessage) Descriptor() protoreflect.MessageDescriptor {
	// TODO implement me
	panic("implement me")
}

// Type
//  @receiver r
//  @return protoreflect.MessageType
//
func (r *reflectedMessage) Type() protoreflect.MessageType {
	// TODO implement me
	panic("implement me")
}

// New
//  @receiver r
//  @return protoreflect.Message
//
func (r *reflectedMessage) New() protoreflect.Message {
	// TODO implement me
	panic("implement me")
}

// Interface
//  @receiver r
//  @return protoreflect.ProtoMessage
//
func (r *reflectedMessage) Interface() protoreflect.ProtoMessage {
	// TODO implement me
	panic("implement me")
}

func (r *reflectedMessage) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	// TODO implement me
	panic("implement me")
}

func (r *reflectedMessage) Has(descriptor protoreflect.FieldDescriptor) bool {
	// TODO implement me
	panic("implement me")
}

func (r *reflectedMessage) Clear(descriptor protoreflect.FieldDescriptor) {
	// TODO implement me
	panic("implement me")
}

func (r *reflectedMessage) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	// TODO implement me
	panic("implement me")
}

func (r *reflectedMessage) Set(descriptor protoreflect.FieldDescriptor, value protoreflect.Value) {
	// TODO implement me
	panic("implement me")
}

func (r *reflectedMessage) Mutable(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	// TODO implement me
	panic("implement me")
}

func (r *reflectedMessage) NewField(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	// TODO implement me
	panic("implement me")
}

func (r *reflectedMessage) WhichOneof(descriptor protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	// TODO implement me
	panic("implement me")
}

func (r *reflectedMessage) GetUnknown() protoreflect.RawFields {
	// TODO implement me
	panic("implement me")
}

func (r *reflectedMessage) SetUnknown(fields protoreflect.RawFields) {
	// TODO implement me
	panic("implement me")
}

func (r *reflectedMessage) IsValid() bool {
	// TODO implement me
	panic("implement me")
}

func (r *reflectedMessage) ProtoMethods() *protoiface.Methods {
	// TODO implement me
	panic("implement me")
}

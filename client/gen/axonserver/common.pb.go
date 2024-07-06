// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: common.proto

package axonserver

import (
	_ "github.com/golang/protobuf/ptypes/any"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// An enumeration of possible keys for processing instructions.
type ProcessingKey int32

const (
	// key indicating that the attached value should be used for consistent routing.
	ProcessingKey_ROUTING_KEY ProcessingKey = 0
	// key indicating that the attached value indicates relative priority of this message.
	ProcessingKey_PRIORITY ProcessingKey = 1
	// key indicating that the accompanied message has a finite validity. The attached value contains the number of milliseconds.
	ProcessingKey_TIMEOUT ProcessingKey = 2
	// key indicating that the requester expects at most the given number of results from this message. Use -1 for unlimited.
	ProcessingKey_NR_OF_RESULTS ProcessingKey = 3
	// key indicating whether Axon Server supports streaming.
	ProcessingKey_SERVER_SUPPORTS_STREAMING ProcessingKey = 7
	// key indicating whether Client supports streaming.
	ProcessingKey_CLIENT_SUPPORTS_STREAMING ProcessingKey = 8
)

// Enum value maps for ProcessingKey.
var (
	ProcessingKey_name = map[int32]string{
		0: "ROUTING_KEY",
		1: "PRIORITY",
		2: "TIMEOUT",
		3: "NR_OF_RESULTS",
		7: "SERVER_SUPPORTS_STREAMING",
		8: "CLIENT_SUPPORTS_STREAMING",
	}
	ProcessingKey_value = map[string]int32{
		"ROUTING_KEY":               0,
		"PRIORITY":                  1,
		"TIMEOUT":                   2,
		"NR_OF_RESULTS":             3,
		"SERVER_SUPPORTS_STREAMING": 7,
		"CLIENT_SUPPORTS_STREAMING": 8,
	}
)

func (x ProcessingKey) Enum() *ProcessingKey {
	p := new(ProcessingKey)
	*p = x
	return p
}

func (x ProcessingKey) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProcessingKey) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[0].Descriptor()
}

func (ProcessingKey) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[0]
}

func (x ProcessingKey) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProcessingKey.Descriptor instead.
func (ProcessingKey) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

// Defines status values for a scheduled task
type TaskStatus int32

const (
	// Task is scheduled for execution
	TaskStatus_SCHEDULED TaskStatus = 0
	// Task execution completed successfully
	TaskStatus_COMPLETED TaskStatus = 1
	// Task execution failed with non transient exception
	TaskStatus_FAILED TaskStatus = 2
	// Task execution is in progress
	TaskStatus_RUNNING TaskStatus = 3
	// Task execution is in progress
	TaskStatus_CANCELLED TaskStatus = 4
)

// Enum value maps for TaskStatus.
var (
	TaskStatus_name = map[int32]string{
		0: "SCHEDULED",
		1: "COMPLETED",
		2: "FAILED",
		3: "RUNNING",
		4: "CANCELLED",
	}
	TaskStatus_value = map[string]int32{
		"SCHEDULED": 0,
		"COMPLETED": 1,
		"FAILED":    2,
		"RUNNING":   3,
		"CANCELLED": 4,
	}
)

func (x TaskStatus) Enum() *TaskStatus {
	p := new(TaskStatus)
	*p = x
	return p
}

func (x TaskStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[1].Descriptor()
}

func (TaskStatus) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[1]
}

func (x TaskStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskStatus.Descriptor instead.
func (TaskStatus) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

// Describes a serialized object
type SerializedObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type identifier of the serialized object.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// The revision of the serialized form of the given type.
	Revision string `protobuf:"bytes,2,opt,name=revision,proto3" json:"revision,omitempty"`
	// The actual data representing the object in serialized form.
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SerializedObject) Reset() {
	*x = SerializedObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SerializedObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SerializedObject) ProtoMessage() {}

func (x *SerializedObject) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SerializedObject.ProtoReflect.Descriptor instead.
func (*SerializedObject) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *SerializedObject) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SerializedObject) GetRevision() string {
	if x != nil {
		return x.Revision
	}
	return ""
}

func (x *SerializedObject) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// The value of a MetaData entry.
type MetaDataValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The data of the MetaData entry, depending on the type of data it contains.
	// Types that are assignable to Data:
	//
	//	*MetaDataValue_TextValue
	//	*MetaDataValue_NumberValue
	//	*MetaDataValue_BooleanValue
	//	*MetaDataValue_DoubleValue
	//	*MetaDataValue_BytesValue
	Data isMetaDataValue_Data `protobuf_oneof:"data"`
}

func (x *MetaDataValue) Reset() {
	*x = MetaDataValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetaDataValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaDataValue) ProtoMessage() {}

func (x *MetaDataValue) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaDataValue.ProtoReflect.Descriptor instead.
func (*MetaDataValue) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (m *MetaDataValue) GetData() isMetaDataValue_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *MetaDataValue) GetTextValue() string {
	if x, ok := x.GetData().(*MetaDataValue_TextValue); ok {
		return x.TextValue
	}
	return ""
}

func (x *MetaDataValue) GetNumberValue() int64 {
	if x, ok := x.GetData().(*MetaDataValue_NumberValue); ok {
		return x.NumberValue
	}
	return 0
}

func (x *MetaDataValue) GetBooleanValue() bool {
	if x, ok := x.GetData().(*MetaDataValue_BooleanValue); ok {
		return x.BooleanValue
	}
	return false
}

func (x *MetaDataValue) GetDoubleValue() float64 {
	if x, ok := x.GetData().(*MetaDataValue_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (x *MetaDataValue) GetBytesValue() *SerializedObject {
	if x, ok := x.GetData().(*MetaDataValue_BytesValue); ok {
		return x.BytesValue
	}
	return nil
}

type isMetaDataValue_Data interface {
	isMetaDataValue_Data()
}

type MetaDataValue_TextValue struct {
	// The text value of the Meta Data entry.
	TextValue string `protobuf:"bytes,1,opt,name=text_value,json=textValue,proto3,oneof"`
}

type MetaDataValue_NumberValue struct {
	// The numeric value of the Meta Data entry.
	NumberValue int64 `protobuf:"zigzag64,2,opt,name=number_value,json=numberValue,proto3,oneof"`
}

type MetaDataValue_BooleanValue struct {
	// The boolean value of the Meta Data entry.
	BooleanValue bool `protobuf:"varint,3,opt,name=boolean_value,json=booleanValue,proto3,oneof"`
}

type MetaDataValue_DoubleValue struct {
	// The floating point value of the Meta Data entry.
	DoubleValue float64 `protobuf:"fixed64,4,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type MetaDataValue_BytesValue struct {
	// The binary value of the Meta Data entry.
	BytesValue *SerializedObject `protobuf:"bytes,5,opt,name=bytes_value,json=bytesValue,proto3,oneof"`
}

func (*MetaDataValue_TextValue) isMetaDataValue_Data() {}

func (*MetaDataValue_NumberValue) isMetaDataValue_Data() {}

func (*MetaDataValue_BooleanValue) isMetaDataValue_Data() {}

func (*MetaDataValue_DoubleValue) isMetaDataValue_Data() {}

func (*MetaDataValue_BytesValue) isMetaDataValue_Data() {}

// An instruction for routing components when routing or processing a message.
type ProcessingInstruction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of processing message.
	Key ProcessingKey `protobuf:"varint,1,opt,name=key,proto3,enum=io.axoniq.axonserver.grpc.ProcessingKey" json:"key,omitempty"`
	// The value associated with the processing key.
	Value *MetaDataValue `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ProcessingInstruction) Reset() {
	*x = ProcessingInstruction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessingInstruction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessingInstruction) ProtoMessage() {}

func (x *ProcessingInstruction) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessingInstruction.ProtoReflect.Descriptor instead.
func (*ProcessingInstruction) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

func (x *ProcessingInstruction) GetKey() ProcessingKey {
	if x != nil {
		return x.Key
	}
	return ProcessingKey_ROUTING_KEY
}

func (x *ProcessingInstruction) GetValue() *MetaDataValue {
	if x != nil {
		return x.Value
	}
	return nil
}

// Message containing details of an error
type ErrorMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A human readable message explaining the error
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	// A description of the location (client component, server) where the error occurred
	Location string `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	// A collection of messages providing more details about root causes of the error
	Details []string `protobuf:"bytes,3,rep,name=details,proto3" json:"details,omitempty"`
	// An Error Code identifying the type of error
	ErrorCode string `protobuf:"bytes,4,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
}

func (x *ErrorMessage) Reset() {
	*x = ErrorMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorMessage) ProtoMessage() {}

func (x *ErrorMessage) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorMessage.ProtoReflect.Descriptor instead.
func (*ErrorMessage) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{3}
}

func (x *ErrorMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ErrorMessage) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *ErrorMessage) GetDetails() []string {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *ErrorMessage) GetErrorCode() string {
	if x != nil {
		return x.ErrorCode
	}
	return ""
}

// Message used for Flow Control instruction, providing the counterpart with additional permits for sending messages
type FlowControl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ClientID of the component providing additional permits
	ClientId string `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// The number of permits to provide
	Permits int64 `protobuf:"varint,3,opt,name=permits,proto3" json:"permits,omitempty"`
}

func (x *FlowControl) Reset() {
	*x = FlowControl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlowControl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowControl) ProtoMessage() {}

func (x *FlowControl) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowControl.ProtoReflect.Descriptor instead.
func (*FlowControl) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{4}
}

func (x *FlowControl) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *FlowControl) GetPermits() int64 {
	if x != nil {
		return x.Permits
	}
	return 0
}

// Message describing instruction acknowledgement
type InstructionAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The identifier of the instruction
	InstructionId string `protobuf:"bytes,1,opt,name=instruction_id,json=instructionId,proto3" json:"instruction_id,omitempty"`
	// Indicator whether the instruction was acknowledged successfully
	Success bool `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	// Set if instruction acknowledgement failed.
	Error *ErrorMessage `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *InstructionAck) Reset() {
	*x = InstructionAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstructionAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstructionAck) ProtoMessage() {}

func (x *InstructionAck) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstructionAck.ProtoReflect.Descriptor instead.
func (*InstructionAck) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{5}
}

func (x *InstructionAck) GetInstructionId() string {
	if x != nil {
		return x.InstructionId
	}
	return ""
}

func (x *InstructionAck) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *InstructionAck) GetError() *ErrorMessage {
	if x != nil {
		return x.Error
	}
	return nil
}

// Message describing the result of the execution of an instruction
type InstructionResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The identifier of the instruction
	InstructionId string `protobuf:"bytes,1,opt,name=instruction_id,json=instructionId,proto3" json:"instruction_id,omitempty"`
	// Indicator whether the instruction was processed successfully
	Success bool `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	// Cause of instruction handling failure.
	Error *ErrorMessage `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *InstructionResult) Reset() {
	*x = InstructionResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstructionResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstructionResult) ProtoMessage() {}

func (x *InstructionResult) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstructionResult.ProtoReflect.Descriptor instead.
func (*InstructionResult) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{6}
}

func (x *InstructionResult) GetInstructionId() string {
	if x != nil {
		return x.InstructionId
	}
	return ""
}

func (x *InstructionResult) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *InstructionResult) GetError() *ErrorMessage {
	if x != nil {
		return x.Error
	}
	return nil
}

type Component struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Component string `protobuf:"bytes,1,opt,name=component,proto3" json:"component,omitempty"`
}

func (x *Component) Reset() {
	*x = Component{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Component) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Component) ProtoMessage() {}

func (x *Component) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Component.ProtoReflect.Descriptor instead.
func (*Component) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{7}
}

func (x *Component) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

type Principal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *Principal) Reset() {
	*x = Principal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Principal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Principal) ProtoMessage() {}

func (x *Principal) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Principal.ProtoReflect.Descriptor instead.
func (*Principal) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{8}
}

func (x *Principal) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19,
	0x69, 0x6f, 0x2e, 0x61, 0x78, 0x6f, 0x6e, 0x69, 0x71, 0x2e, 0x61, 0x78, 0x6f, 0x6e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x10, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a,
	0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xf9, 0x01, 0x0a,
	0x0d, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f,
	0x0a, 0x0a, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x74, 0x65, 0x78, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x23, 0x0a, 0x0c, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x12, 0x48, 0x00, 0x52, 0x0b, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x25, 0x0a, 0x0d, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0c, 0x62,
	0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x64,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x01, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x4e, 0x0a, 0x0b, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x78, 0x6f, 0x6e, 0x69,
	0x71, 0x2e, 0x61, 0x78, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x48, 0x00, 0x52, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x93, 0x01, 0x0a, 0x15, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x28, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x78, 0x6f, 0x6e, 0x69, 0x71, 0x2e, 0x61, 0x78, 0x6f, 0x6e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3e,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x69, 0x6f, 0x2e, 0x61, 0x78, 0x6f, 0x6e, 0x69, 0x71, 0x2e, 0x61, 0x78, 0x6f, 0x6e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61,
	0x74, 0x61, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x7d,
	0x0a, 0x0c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x4a, 0x0a,
	0x0b, 0x46, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x74, 0x73, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x22, 0x90, 0x01, 0x0a, 0x0e, 0x49, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x6b, 0x12, 0x25, 0x0a, 0x0e,
	0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x3d, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x69,
	0x6f, 0x2e, 0x61, 0x78, 0x6f, 0x6e, 0x69, 0x71, 0x2e, 0x61, 0x78, 0x6f, 0x6e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x93, 0x01, 0x0a,
	0x11, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x3d, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x78, 0x6f, 0x6e, 0x69, 0x71, 0x2e, 0x61,
	0x78, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0x29, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x22, 0x1f, 0x0a,
	0x09, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x2a, 0x9e,
	0x01, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79,
	0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x4f, 0x55, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x4b, 0x45, 0x59, 0x10,
	0x00, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x52, 0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x10, 0x01, 0x12,
	0x0b, 0x0a, 0x07, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d,
	0x4e, 0x52, 0x5f, 0x4f, 0x46, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x53, 0x10, 0x03, 0x12,
	0x1d, 0x0a, 0x19, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52,
	0x54, 0x53, 0x5f, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x49, 0x4e, 0x47, 0x10, 0x07, 0x12, 0x1d,
	0x0a, 0x19, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54,
	0x53, 0x5f, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x49, 0x4e, 0x47, 0x10, 0x08, 0x22, 0x04, 0x08,
	0x04, 0x10, 0x04, 0x22, 0x04, 0x08, 0x05, 0x10, 0x05, 0x22, 0x04, 0x08, 0x06, 0x10, 0x06, 0x2a,
	0x52, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0d, 0x0a,
	0x09, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09,
	0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49,
	0x4e, 0x47, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45,
	0x44, 0x10, 0x04, 0x42, 0x19, 0x50, 0x01, 0x5a, 0x15, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x61, 0x78, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_common_proto_goTypes = []interface{}{
	(ProcessingKey)(0),            // 0: io.axoniq.axonserver.grpc.ProcessingKey
	(TaskStatus)(0),               // 1: io.axoniq.axonserver.grpc.TaskStatus
	(*SerializedObject)(nil),      // 2: io.axoniq.axonserver.grpc.SerializedObject
	(*MetaDataValue)(nil),         // 3: io.axoniq.axonserver.grpc.MetaDataValue
	(*ProcessingInstruction)(nil), // 4: io.axoniq.axonserver.grpc.ProcessingInstruction
	(*ErrorMessage)(nil),          // 5: io.axoniq.axonserver.grpc.ErrorMessage
	(*FlowControl)(nil),           // 6: io.axoniq.axonserver.grpc.FlowControl
	(*InstructionAck)(nil),        // 7: io.axoniq.axonserver.grpc.InstructionAck
	(*InstructionResult)(nil),     // 8: io.axoniq.axonserver.grpc.InstructionResult
	(*Component)(nil),             // 9: io.axoniq.axonserver.grpc.Component
	(*Principal)(nil),             // 10: io.axoniq.axonserver.grpc.Principal
}
var file_common_proto_depIdxs = []int32{
	2, // 0: io.axoniq.axonserver.grpc.MetaDataValue.bytes_value:type_name -> io.axoniq.axonserver.grpc.SerializedObject
	0, // 1: io.axoniq.axonserver.grpc.ProcessingInstruction.key:type_name -> io.axoniq.axonserver.grpc.ProcessingKey
	3, // 2: io.axoniq.axonserver.grpc.ProcessingInstruction.value:type_name -> io.axoniq.axonserver.grpc.MetaDataValue
	5, // 3: io.axoniq.axonserver.grpc.InstructionAck.error:type_name -> io.axoniq.axonserver.grpc.ErrorMessage
	5, // 4: io.axoniq.axonserver.grpc.InstructionResult.error:type_name -> io.axoniq.axonserver.grpc.ErrorMessage
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SerializedObject); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetaDataValue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessingInstruction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlowControl); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstructionAck); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstructionResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Component); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Principal); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_common_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*MetaDataValue_TextValue)(nil),
		(*MetaDataValue_NumberValue)(nil),
		(*MetaDataValue_BooleanValue)(nil),
		(*MetaDataValue_DoubleValue)(nil),
		(*MetaDataValue_BytesValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		EnumInfos:         file_common_proto_enumTypes,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
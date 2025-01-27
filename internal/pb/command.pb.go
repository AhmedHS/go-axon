// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: command.proto

package pb

import (
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

// An instruction from the components that provides the Command Handler towards AxonServer.
type CommandProviderOutbound struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The instruction for AxonServer
	// Types that are assignable to Request:
	//
	//	*CommandProviderOutbound_Subscribe
	//	*CommandProviderOutbound_Unsubscribe
	//	*CommandProviderOutbound_FlowControl
	//	*CommandProviderOutbound_CommandResponse
	//	*CommandProviderOutbound_Ack
	Request isCommandProviderOutbound_Request `protobuf_oneof:"request"`
	// Instruction identifier. If this identifier is set, this instruction will be acknowledged via inbound stream
	InstructionId string `protobuf:"bytes,6,opt,name=instruction_id,json=instructionId,proto3" json:"instruction_id,omitempty"`
}

func (x *CommandProviderOutbound) Reset() {
	*x = CommandProviderOutbound{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandProviderOutbound) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandProviderOutbound) ProtoMessage() {}

func (x *CommandProviderOutbound) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandProviderOutbound.ProtoReflect.Descriptor instead.
func (*CommandProviderOutbound) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{0}
}

func (m *CommandProviderOutbound) GetRequest() isCommandProviderOutbound_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *CommandProviderOutbound) GetSubscribe() *CommandSubscription {
	if x, ok := x.GetRequest().(*CommandProviderOutbound_Subscribe); ok {
		return x.Subscribe
	}
	return nil
}

func (x *CommandProviderOutbound) GetUnsubscribe() *CommandSubscription {
	if x, ok := x.GetRequest().(*CommandProviderOutbound_Unsubscribe); ok {
		return x.Unsubscribe
	}
	return nil
}

func (x *CommandProviderOutbound) GetFlowControl() *FlowControl {
	if x, ok := x.GetRequest().(*CommandProviderOutbound_FlowControl); ok {
		return x.FlowControl
	}
	return nil
}

func (x *CommandProviderOutbound) GetCommandResponse() *CommandResponse {
	if x, ok := x.GetRequest().(*CommandProviderOutbound_CommandResponse); ok {
		return x.CommandResponse
	}
	return nil
}

func (x *CommandProviderOutbound) GetAck() *InstructionAck {
	if x, ok := x.GetRequest().(*CommandProviderOutbound_Ack); ok {
		return x.Ack
	}
	return nil
}

func (x *CommandProviderOutbound) GetInstructionId() string {
	if x != nil {
		return x.InstructionId
	}
	return ""
}

type isCommandProviderOutbound_Request interface {
	isCommandProviderOutbound_Request()
}

type CommandProviderOutbound_Subscribe struct {
	// Instruction to subscribe this component as handler of a specific type of command
	Subscribe *CommandSubscription `protobuf:"bytes,1,opt,name=subscribe,proto3,oneof"`
}

type CommandProviderOutbound_Unsubscribe struct {
	// Instruction to unsubscribe this component as handler of a specific type of command
	Unsubscribe *CommandSubscription `protobuf:"bytes,2,opt,name=unsubscribe,proto3,oneof"`
}

type CommandProviderOutbound_FlowControl struct {
	// Instruction to increase the number of instructions AxonServer may send to this component
	FlowControl *FlowControl `protobuf:"bytes,3,opt,name=flow_control,json=flowControl,proto3,oneof"`
}

type CommandProviderOutbound_CommandResponse struct {
	// Sends a result of Command processing
	CommandResponse *CommandResponse `protobuf:"bytes,4,opt,name=command_response,json=commandResponse,proto3,oneof"`
}

type CommandProviderOutbound_Ack struct {
	// Acknowledgement of previously sent instruction via inbound stream
	Ack *InstructionAck `protobuf:"bytes,5,opt,name=ack,proto3,oneof"`
}

func (*CommandProviderOutbound_Subscribe) isCommandProviderOutbound_Request() {}

func (*CommandProviderOutbound_Unsubscribe) isCommandProviderOutbound_Request() {}

func (*CommandProviderOutbound_FlowControl) isCommandProviderOutbound_Request() {}

func (*CommandProviderOutbound_CommandResponse) isCommandProviderOutbound_Request() {}

func (*CommandProviderOutbound_Ack) isCommandProviderOutbound_Request() {}

// An instruction or confirmation from AxonServer towards the component that provides the Command Handler
type CommandProviderInbound struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The instruction from AxonServer for this component
	// Types that are assignable to Request:
	//
	//	*CommandProviderInbound_Ack
	//	*CommandProviderInbound_Command
	Request isCommandProviderInbound_Request `protobuf_oneof:"request"`
	// Instruction identifier. If this identifier is set, this instruction will be acknowledged via outbound stream
	InstructionId string `protobuf:"bytes,3,opt,name=instruction_id,json=instructionId,proto3" json:"instruction_id,omitempty"`
}

func (x *CommandProviderInbound) Reset() {
	*x = CommandProviderInbound{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandProviderInbound) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandProviderInbound) ProtoMessage() {}

func (x *CommandProviderInbound) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandProviderInbound.ProtoReflect.Descriptor instead.
func (*CommandProviderInbound) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{1}
}

func (m *CommandProviderInbound) GetRequest() isCommandProviderInbound_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *CommandProviderInbound) GetAck() *InstructionAck {
	if x, ok := x.GetRequest().(*CommandProviderInbound_Ack); ok {
		return x.Ack
	}
	return nil
}

func (x *CommandProviderInbound) GetCommand() *Command {
	if x, ok := x.GetRequest().(*CommandProviderInbound_Command); ok {
		return x.Command
	}
	return nil
}

func (x *CommandProviderInbound) GetInstructionId() string {
	if x != nil {
		return x.InstructionId
	}
	return ""
}

type isCommandProviderInbound_Request interface {
	isCommandProviderInbound_Request()
}

type CommandProviderInbound_Ack struct {
	// Acknowledgement of previously sent instruction via outbound stream
	Ack *InstructionAck `protobuf:"bytes,1,opt,name=ack,proto3,oneof"`
}

type CommandProviderInbound_Command struct {
	// A command for this component to process
	Command *Command `protobuf:"bytes,2,opt,name=command,proto3,oneof"`
}

func (*CommandProviderInbound_Ack) isCommandProviderInbound_Request() {}

func (*CommandProviderInbound_Command) isCommandProviderInbound_Request() {}

// A message representing a Command that needs to be routed to a component capable of handling it
type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique identifier of the Command Message
	MessageIdentifier string `protobuf:"bytes,1,opt,name=message_identifier,json=messageIdentifier,proto3" json:"message_identifier,omitempty"`
	// The name of the command, used for routing it to a destination capable of handling it
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The time at which the command was dispatched
	Timestamp int64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// The payload of the Command, providing details on the instructions for the recipient
	Payload *SerializedObject `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	// Meta Data entries of the Command Message, providing contextual information to the recipient
	MetaData map[string]*MetaDataValue `protobuf:"bytes,5,rep,name=meta_data,json=metaData,proto3" json:"meta_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Instructions for AxonServer when routing this Command Message
	ProcessingInstructions []*ProcessingInstruction `protobuf:"bytes,6,rep,name=processing_instructions,json=processingInstructions,proto3" json:"processing_instructions,omitempty"`
	// The unique identifier of the component dispatching this message
	ClientId string `protobuf:"bytes,7,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// The name/type of the component dispatching this message
	ComponentName string `protobuf:"bytes,8,opt,name=component_name,json=componentName,proto3" json:"component_name,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{2}
}

func (x *Command) GetMessageIdentifier() string {
	if x != nil {
		return x.MessageIdentifier
	}
	return ""
}

func (x *Command) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Command) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Command) GetPayload() *SerializedObject {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *Command) GetMetaData() map[string]*MetaDataValue {
	if x != nil {
		return x.MetaData
	}
	return nil
}

func (x *Command) GetProcessingInstructions() []*ProcessingInstruction {
	if x != nil {
		return x.ProcessingInstructions
	}
	return nil
}

func (x *Command) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *Command) GetComponentName() string {
	if x != nil {
		return x.ComponentName
	}
	return ""
}

// Message representing the result of Command Handler execution
type CommandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique identifier of the response message
	MessageIdentifier string `protobuf:"bytes,1,opt,name=message_identifier,json=messageIdentifier,proto3" json:"message_identifier,omitempty"`
	// An error code describing the error, if any
	ErrorCode string `protobuf:"bytes,2,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	// A detailed description of the error
	ErrorMessage *ErrorMessage `protobuf:"bytes,3,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	// The payload to provide as a result to the dispatcher
	Payload *SerializedObject `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	// Any meta data entries providing contextual information back to the dispatcher
	MetaData map[string]*MetaDataValue `protobuf:"bytes,5,rep,name=meta_data,json=metaData,proto3" json:"meta_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Instructions for AxonServer when routing this Command Response Message
	ProcessingInstructions []*ProcessingInstruction `protobuf:"bytes,6,rep,name=processing_instructions,json=processingInstructions,proto3" json:"processing_instructions,omitempty"`
	// The unique identifier of the Command Message for which this is the response
	RequestIdentifier string `protobuf:"bytes,7,opt,name=request_identifier,json=requestIdentifier,proto3" json:"request_identifier,omitempty"`
}

func (x *CommandResponse) Reset() {
	*x = CommandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandResponse) ProtoMessage() {}

func (x *CommandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandResponse.ProtoReflect.Descriptor instead.
func (*CommandResponse) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{3}
}

func (x *CommandResponse) GetMessageIdentifier() string {
	if x != nil {
		return x.MessageIdentifier
	}
	return ""
}

func (x *CommandResponse) GetErrorCode() string {
	if x != nil {
		return x.ErrorCode
	}
	return ""
}

func (x *CommandResponse) GetErrorMessage() *ErrorMessage {
	if x != nil {
		return x.ErrorMessage
	}
	return nil
}

func (x *CommandResponse) GetPayload() *SerializedObject {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *CommandResponse) GetMetaData() map[string]*MetaDataValue {
	if x != nil {
		return x.MetaData
	}
	return nil
}

func (x *CommandResponse) GetProcessingInstructions() []*ProcessingInstruction {
	if x != nil {
		return x.ProcessingInstructions
	}
	return nil
}

func (x *CommandResponse) GetRequestIdentifier() string {
	if x != nil {
		return x.RequestIdentifier
	}
	return ""
}

// Message describing a component's capability of handling a command type
type CommandSubscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A unique identifier for this subscription. This identifier is returned in Acknowledgements to allow
	// pipelining of subscription messages
	MessageId string `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	// The name of the command the component can handle
	Command string `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
	// The name/type of the component handling the command
	ComponentName string `protobuf:"bytes,3,opt,name=component_name,json=componentName,proto3" json:"component_name,omitempty"`
	// The unique identifier of the component instance subscribing
	ClientId string `protobuf:"bytes,4,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// A number that represents the client's relative load capacity compared to other clients.
	// This information is interpreted by Axon Server in relation to the other connected nodes' values.
	// Used to balance the dispatching of commands. If set to 0, Axon Server consider 100 as default value.
	LoadFactor int32 `protobuf:"varint,5,opt,name=load_factor,json=loadFactor,proto3" json:"load_factor,omitempty"`
}

func (x *CommandSubscription) Reset() {
	*x = CommandSubscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandSubscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandSubscription) ProtoMessage() {}

func (x *CommandSubscription) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandSubscription.ProtoReflect.Descriptor instead.
func (*CommandSubscription) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{4}
}

func (x *CommandSubscription) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *CommandSubscription) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *CommandSubscription) GetComponentName() string {
	if x != nil {
		return x.ComponentName
	}
	return ""
}

func (x *CommandSubscription) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *CommandSubscription) GetLoadFactor() int32 {
	if x != nil {
		return x.LoadFactor
	}
	return 0
}

var File_command_proto protoreflect.FileDescriptor

var file_command_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x21, 0x69, 0x6f, 0x2e, 0x61, 0x78, 0x6f, 0x6e, 0x69, 0x71, 0x2e, 0x61, 0x78, 0x6f, 0x6e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xec, 0x03, 0x0a, 0x17, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x56, 0x0a, 0x09,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x36, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x78, 0x6f, 0x6e, 0x69, 0x71, 0x2e, 0x61, 0x78, 0x6f, 0x6e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x09, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x12, 0x5a, 0x0a, 0x0b, 0x75, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x69, 0x6f, 0x2e, 0x61,
	0x78, 0x6f, 0x6e, 0x69, 0x71, 0x2e, 0x61, 0x78, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x48, 0x00, 0x52, 0x0b, 0x75, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x12, 0x4b, 0x0a, 0x0c, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x78, 0x6f, 0x6e,
	0x69, 0x71, 0x2e, 0x61, 0x78, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x48, 0x00,
	0x52, 0x0b, 0x66, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x5f, 0x0a,
	0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x78, 0x6f,
	0x6e, 0x69, 0x71, 0x2e, 0x61, 0x78, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d,
	0x0a, 0x03, 0x61, 0x63, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x69, 0x6f,
	0x2e, 0x61, 0x78, 0x6f, 0x6e, 0x69, 0x71, 0x2e, 0x61, 0x78, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x41, 0x63, 0x6b, 0x48, 0x00, 0x52, 0x03, 0x61, 0x63, 0x6b, 0x12, 0x25, 0x0a,
	0x0e, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0xd1, 0x01, 0x0a, 0x16, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x49, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x3d, 0x0a, 0x03, 0x61, 0x63,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x78, 0x6f,
	0x6e, 0x69, 0x71, 0x2e, 0x61, 0x78, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x63, 0x6b, 0x48, 0x00, 0x52, 0x03, 0x61, 0x63, 0x6b, 0x12, 0x46, 0x0a, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x69, 0x6f, 0x2e,
	0x61, 0x78, 0x6f, 0x6e, 0x69, 0x71, 0x2e, 0x61, 0x78, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x9e, 0x04, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12,
	0x2d, 0x0a, 0x12, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x45, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2b, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x78, 0x6f, 0x6e, 0x69, 0x71, 0x2e, 0x61, 0x78,
	0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x55, 0x0a, 0x09, 0x6d, 0x65, 0x74, 0x61, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x69, 0x6f, 0x2e,
	0x61, 0x78, 0x6f, 0x6e, 0x69, 0x71, 0x2e, 0x61, 0x78, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x69,
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x30, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x78, 0x6f, 0x6e, 0x69, 0x71, 0x2e, 0x61, 0x78, 0x6f, 0x6e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x16, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x65, 0x0a,
	0x0d, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x3e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x78, 0x6f, 0x6e, 0x69, 0x71, 0x2e, 0x61, 0x78, 0x6f, 0x6e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x44, 0x61, 0x74, 0x61, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0xd4, 0x04, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x4c, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x69, 0x6f, 0x2e, 0x61, 0x78, 0x6f, 0x6e, 0x69, 0x71, 0x2e, 0x61, 0x78, 0x6f, 0x6e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x45, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x78, 0x6f, 0x6e, 0x69,
	0x71, 0x2e, 0x61, 0x78, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x5d, 0x0a, 0x09, 0x6d,
	0x65, 0x74, 0x61, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x40,
	0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x78, 0x6f, 0x6e, 0x69, 0x71, 0x2e, 0x61, 0x78, 0x6f, 0x6e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x69, 0x0a, 0x17, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x69, 0x6f,
	0x2e, 0x61, 0x78, 0x6f, 0x6e, 0x69, 0x71, 0x2e, 0x61, 0x78, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69,
	0x6e, 0x67, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x16, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x1a, 0x65, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x78, 0x6f, 0x6e,
	0x69, 0x71, 0x2e, 0x61, 0x78, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xb3, 0x01, 0x0a, 0x13,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x25, 0x0a, 0x0e,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x32, 0x8a, 0x02, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x89, 0x01, 0x0a, 0x0a, 0x4f, 0x70, 0x65, 0x6e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x12, 0x3a, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x78, 0x6f, 0x6e, 0x69, 0x71, 0x2e,
	0x61, 0x78, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x1a,
	0x39, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x78, 0x6f, 0x6e, 0x69, 0x71, 0x2e, 0x61, 0x78, 0x6f, 0x6e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x49, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01,
	0x12, 0x6c, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x12, 0x2a, 0x2e, 0x69,
	0x6f, 0x2e, 0x61, 0x78, 0x6f, 0x6e, 0x69, 0x71, 0x2e, 0x61, 0x78, 0x6f, 0x6e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x32, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x78,
	0x6f, 0x6e, 0x69, 0x71, 0x2e, 0x61, 0x78, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f,
	0x50, 0x01, 0x5a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_command_proto_rawDescOnce sync.Once
	file_command_proto_rawDescData = file_command_proto_rawDesc
)

func file_command_proto_rawDescGZIP() []byte {
	file_command_proto_rawDescOnce.Do(func() {
		file_command_proto_rawDescData = protoimpl.X.CompressGZIP(file_command_proto_rawDescData)
	})
	return file_command_proto_rawDescData
}

var file_command_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_command_proto_goTypes = []interface{}{
	(*CommandProviderOutbound)(nil), // 0: io.axoniq.axonserver.grpc.command.CommandProviderOutbound
	(*CommandProviderInbound)(nil),  // 1: io.axoniq.axonserver.grpc.command.CommandProviderInbound
	(*Command)(nil),                 // 2: io.axoniq.axonserver.grpc.command.Command
	(*CommandResponse)(nil),         // 3: io.axoniq.axonserver.grpc.command.CommandResponse
	(*CommandSubscription)(nil),     // 4: io.axoniq.axonserver.grpc.command.CommandSubscription
	nil,                             // 5: io.axoniq.axonserver.grpc.command.Command.MetaDataEntry
	nil,                             // 6: io.axoniq.axonserver.grpc.command.CommandResponse.MetaDataEntry
	(*FlowControl)(nil),             // 7: io.axoniq.axonserver.grpc.FlowControl
	(*InstructionAck)(nil),          // 8: io.axoniq.axonserver.grpc.InstructionAck
	(*SerializedObject)(nil),        // 9: io.axoniq.axonserver.grpc.SerializedObject
	(*ProcessingInstruction)(nil),   // 10: io.axoniq.axonserver.grpc.ProcessingInstruction
	(*ErrorMessage)(nil),            // 11: io.axoniq.axonserver.grpc.ErrorMessage
	(*MetaDataValue)(nil),           // 12: io.axoniq.axonserver.grpc.MetaDataValue
}
var file_command_proto_depIdxs = []int32{
	4,  // 0: io.axoniq.axonserver.grpc.command.CommandProviderOutbound.subscribe:type_name -> io.axoniq.axonserver.grpc.command.CommandSubscription
	4,  // 1: io.axoniq.axonserver.grpc.command.CommandProviderOutbound.unsubscribe:type_name -> io.axoniq.axonserver.grpc.command.CommandSubscription
	7,  // 2: io.axoniq.axonserver.grpc.command.CommandProviderOutbound.flow_control:type_name -> io.axoniq.axonserver.grpc.FlowControl
	3,  // 3: io.axoniq.axonserver.grpc.command.CommandProviderOutbound.command_response:type_name -> io.axoniq.axonserver.grpc.command.CommandResponse
	8,  // 4: io.axoniq.axonserver.grpc.command.CommandProviderOutbound.ack:type_name -> io.axoniq.axonserver.grpc.InstructionAck
	8,  // 5: io.axoniq.axonserver.grpc.command.CommandProviderInbound.ack:type_name -> io.axoniq.axonserver.grpc.InstructionAck
	2,  // 6: io.axoniq.axonserver.grpc.command.CommandProviderInbound.command:type_name -> io.axoniq.axonserver.grpc.command.Command
	9,  // 7: io.axoniq.axonserver.grpc.command.Command.payload:type_name -> io.axoniq.axonserver.grpc.SerializedObject
	5,  // 8: io.axoniq.axonserver.grpc.command.Command.meta_data:type_name -> io.axoniq.axonserver.grpc.command.Command.MetaDataEntry
	10, // 9: io.axoniq.axonserver.grpc.command.Command.processing_instructions:type_name -> io.axoniq.axonserver.grpc.ProcessingInstruction
	11, // 10: io.axoniq.axonserver.grpc.command.CommandResponse.error_message:type_name -> io.axoniq.axonserver.grpc.ErrorMessage
	9,  // 11: io.axoniq.axonserver.grpc.command.CommandResponse.payload:type_name -> io.axoniq.axonserver.grpc.SerializedObject
	6,  // 12: io.axoniq.axonserver.grpc.command.CommandResponse.meta_data:type_name -> io.axoniq.axonserver.grpc.command.CommandResponse.MetaDataEntry
	10, // 13: io.axoniq.axonserver.grpc.command.CommandResponse.processing_instructions:type_name -> io.axoniq.axonserver.grpc.ProcessingInstruction
	12, // 14: io.axoniq.axonserver.grpc.command.Command.MetaDataEntry.value:type_name -> io.axoniq.axonserver.grpc.MetaDataValue
	12, // 15: io.axoniq.axonserver.grpc.command.CommandResponse.MetaDataEntry.value:type_name -> io.axoniq.axonserver.grpc.MetaDataValue
	0,  // 16: io.axoniq.axonserver.grpc.command.CommandService.OpenStream:input_type -> io.axoniq.axonserver.grpc.command.CommandProviderOutbound
	2,  // 17: io.axoniq.axonserver.grpc.command.CommandService.Dispatch:input_type -> io.axoniq.axonserver.grpc.command.Command
	1,  // 18: io.axoniq.axonserver.grpc.command.CommandService.OpenStream:output_type -> io.axoniq.axonserver.grpc.command.CommandProviderInbound
	3,  // 19: io.axoniq.axonserver.grpc.command.CommandService.Dispatch:output_type -> io.axoniq.axonserver.grpc.command.CommandResponse
	18, // [18:20] is the sub-list for method output_type
	16, // [16:18] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_command_proto_init() }
func file_command_proto_init() {
	if File_command_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_command_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandProviderOutbound); i {
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
		file_command_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandProviderInbound); i {
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
		file_command_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
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
		file_command_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandResponse); i {
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
		file_command_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandSubscription); i {
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
	file_command_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*CommandProviderOutbound_Subscribe)(nil),
		(*CommandProviderOutbound_Unsubscribe)(nil),
		(*CommandProviderOutbound_FlowControl)(nil),
		(*CommandProviderOutbound_CommandResponse)(nil),
		(*CommandProviderOutbound_Ack)(nil),
	}
	file_command_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*CommandProviderInbound_Ack)(nil),
		(*CommandProviderInbound_Command)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_command_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_command_proto_goTypes,
		DependencyIndexes: file_command_proto_depIdxs,
		MessageInfos:      file_command_proto_msgTypes,
	}.Build()
	File_command_proto = out.File
	file_command_proto_rawDesc = nil
	file_command_proto_goTypes = nil
	file_command_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agentpb/agent.proto

package agentpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	inventorypb "github.com/percona/pmm/api/inventorypb"
	qanpb "github.com/percona/pmm/api/qanpb"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Type represents Agent type.
// TODO Replace with inventory.AgentType. https://jira.percona.com/browse/PMM-3786
type Type int32

const (
	Type_TYPE_INVALID               Type = 0
	Type_PMM_AGENT                  Type = 1
	Type_NODE_EXPORTER              Type = 2
	Type_MYSQLD_EXPORTER            Type = 3
	Type_MONGODB_EXPORTER           Type = 4
	Type_POSTGRES_EXPORTER          Type = 5
	Type_QAN_MYSQL_PERFSCHEMA_AGENT Type = 6
	Type_QAN_MYSQL_SLOWLOG_AGENT    Type = 7
	Type_QAN_MONGODB_PROFILER_AGENT Type = 8
	Type_RDS_EXPORTER               Type = 9
)

var Type_name = map[int32]string{
	0: "TYPE_INVALID",
	1: "PMM_AGENT",
	2: "NODE_EXPORTER",
	3: "MYSQLD_EXPORTER",
	4: "MONGODB_EXPORTER",
	5: "POSTGRES_EXPORTER",
	6: "QAN_MYSQL_PERFSCHEMA_AGENT",
	7: "QAN_MYSQL_SLOWLOG_AGENT",
	8: "QAN_MONGODB_PROFILER_AGENT",
	9: "RDS_EXPORTER",
}

var Type_value = map[string]int32{
	"TYPE_INVALID":               0,
	"PMM_AGENT":                  1,
	"NODE_EXPORTER":              2,
	"MYSQLD_EXPORTER":            3,
	"MONGODB_EXPORTER":           4,
	"POSTGRES_EXPORTER":          5,
	"QAN_MYSQL_PERFSCHEMA_AGENT": 6,
	"QAN_MYSQL_SLOWLOG_AGENT":    7,
	"QAN_MONGODB_PROFILER_AGENT": 8,
	"RDS_EXPORTER":               9,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0e8f797f46bba6fc, []int{0}
}

// Ping is a AgentMessage/ServerMessage for checking connectivity, latency and clock drift.
type Ping struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e8f797f46bba6fc, []int{0}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

// Pong is an AgentMessage/ServerMessage with current time for measuring clock drift.
type Pong struct {
	CurrentTime          *timestamp.Timestamp `protobuf:"bytes,1,opt,name=current_time,json=currentTime,proto3" json:"current_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e8f797f46bba6fc, []int{1}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetCurrentTime() *timestamp.Timestamp {
	if m != nil {
		return m.CurrentTime
	}
	return nil
}

// QANCollectRequest is an AgentMessage for sending QAN data for qan-api.
type QANCollectRequest struct {
	Message              *qanpb.CollectRequest `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *QANCollectRequest) Reset()         { *m = QANCollectRequest{} }
func (m *QANCollectRequest) String() string { return proto.CompactTextString(m) }
func (*QANCollectRequest) ProtoMessage()    {}
func (*QANCollectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e8f797f46bba6fc, []int{2}
}

func (m *QANCollectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QANCollectRequest.Unmarshal(m, b)
}
func (m *QANCollectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QANCollectRequest.Marshal(b, m, deterministic)
}
func (m *QANCollectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QANCollectRequest.Merge(m, src)
}
func (m *QANCollectRequest) XXX_Size() int {
	return xxx_messageInfo_QANCollectRequest.Size(m)
}
func (m *QANCollectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QANCollectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QANCollectRequest proto.InternalMessageInfo

func (m *QANCollectRequest) GetMessage() *qanpb.CollectRequest {
	if m != nil {
		return m.Message
	}
	return nil
}

// QANCollectResponse is a ServerMessage for QAN data acceptance.
type QANCollectResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QANCollectResponse) Reset()         { *m = QANCollectResponse{} }
func (m *QANCollectResponse) String() string { return proto.CompactTextString(m) }
func (*QANCollectResponse) ProtoMessage()    {}
func (*QANCollectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e8f797f46bba6fc, []int{3}
}

func (m *QANCollectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QANCollectResponse.Unmarshal(m, b)
}
func (m *QANCollectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QANCollectResponse.Marshal(b, m, deterministic)
}
func (m *QANCollectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QANCollectResponse.Merge(m, src)
}
func (m *QANCollectResponse) XXX_Size() int {
	return xxx_messageInfo_QANCollectResponse.Size(m)
}
func (m *QANCollectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QANCollectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QANCollectResponse proto.InternalMessageInfo

// StateChangedRequest is an AgentMessage describing actual agent status.
type StateChangedRequest struct {
	AgentId              string                  `protobuf:"bytes,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	Status               inventorypb.AgentStatus `protobuf:"varint,2,opt,name=status,proto3,enum=inventory.AgentStatus" json:"status,omitempty"`
	ListenPort           uint32                  `protobuf:"varint,3,opt,name=listen_port,json=listenPort,proto3" json:"listen_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *StateChangedRequest) Reset()         { *m = StateChangedRequest{} }
func (m *StateChangedRequest) String() string { return proto.CompactTextString(m) }
func (*StateChangedRequest) ProtoMessage()    {}
func (*StateChangedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e8f797f46bba6fc, []int{4}
}

func (m *StateChangedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateChangedRequest.Unmarshal(m, b)
}
func (m *StateChangedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateChangedRequest.Marshal(b, m, deterministic)
}
func (m *StateChangedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateChangedRequest.Merge(m, src)
}
func (m *StateChangedRequest) XXX_Size() int {
	return xxx_messageInfo_StateChangedRequest.Size(m)
}
func (m *StateChangedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StateChangedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StateChangedRequest proto.InternalMessageInfo

func (m *StateChangedRequest) GetAgentId() string {
	if m != nil {
		return m.AgentId
	}
	return ""
}

func (m *StateChangedRequest) GetStatus() inventorypb.AgentStatus {
	if m != nil {
		return m.Status
	}
	return inventorypb.AgentStatus_AGENT_STATUS_INVALID
}

func (m *StateChangedRequest) GetListenPort() uint32 {
	if m != nil {
		return m.ListenPort
	}
	return 0
}

// StateChangedResponse is a ServerMessage for StateChangedRequest acceptance.
type StateChangedResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StateChangedResponse) Reset()         { *m = StateChangedResponse{} }
func (m *StateChangedResponse) String() string { return proto.CompactTextString(m) }
func (*StateChangedResponse) ProtoMessage()    {}
func (*StateChangedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e8f797f46bba6fc, []int{5}
}

func (m *StateChangedResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateChangedResponse.Unmarshal(m, b)
}
func (m *StateChangedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateChangedResponse.Marshal(b, m, deterministic)
}
func (m *StateChangedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateChangedResponse.Merge(m, src)
}
func (m *StateChangedResponse) XXX_Size() int {
	return xxx_messageInfo_StateChangedResponse.Size(m)
}
func (m *StateChangedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StateChangedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StateChangedResponse proto.InternalMessageInfo

// SetStateRequest is a ServerMessage asking pmm-agent to run agents according to desired state.
type SetStateRequest struct {
	AgentProcesses       map[string]*SetStateRequest_AgentProcess `protobuf:"bytes,1,rep,name=agent_processes,json=agentProcesses,proto3" json:"agent_processes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	BuiltinAgents        map[string]*SetStateRequest_BuiltinAgent `protobuf:"bytes,2,rep,name=builtin_agents,json=builtinAgents,proto3" json:"builtin_agents,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *SetStateRequest) Reset()         { *m = SetStateRequest{} }
func (m *SetStateRequest) String() string { return proto.CompactTextString(m) }
func (*SetStateRequest) ProtoMessage()    {}
func (*SetStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e8f797f46bba6fc, []int{6}
}

func (m *SetStateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetStateRequest.Unmarshal(m, b)
}
func (m *SetStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetStateRequest.Marshal(b, m, deterministic)
}
func (m *SetStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetStateRequest.Merge(m, src)
}
func (m *SetStateRequest) XXX_Size() int {
	return xxx_messageInfo_SetStateRequest.Size(m)
}
func (m *SetStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetStateRequest proto.InternalMessageInfo

func (m *SetStateRequest) GetAgentProcesses() map[string]*SetStateRequest_AgentProcess {
	if m != nil {
		return m.AgentProcesses
	}
	return nil
}

func (m *SetStateRequest) GetBuiltinAgents() map[string]*SetStateRequest_BuiltinAgent {
	if m != nil {
		return m.BuiltinAgents
	}
	return nil
}

// AgentProcess describes desired configuration of a single agent process started by pmm-agent.
type SetStateRequest_AgentProcess struct {
	Type                 Type              `protobuf:"varint,1,opt,name=type,proto3,enum=agent.Type" json:"type,omitempty"`
	TemplateLeftDelim    string            `protobuf:"bytes,2,opt,name=template_left_delim,json=templateLeftDelim,proto3" json:"template_left_delim,omitempty"`
	TemplateRightDelim   string            `protobuf:"bytes,3,opt,name=template_right_delim,json=templateRightDelim,proto3" json:"template_right_delim,omitempty"`
	Args                 []string          `protobuf:"bytes,4,rep,name=args,proto3" json:"args,omitempty"`
	Env                  []string          `protobuf:"bytes,5,rep,name=env,proto3" json:"env,omitempty"`
	TextFiles            map[string]string `protobuf:"bytes,6,rep,name=text_files,json=textFiles,proto3" json:"text_files,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SetStateRequest_AgentProcess) Reset()         { *m = SetStateRequest_AgentProcess{} }
func (m *SetStateRequest_AgentProcess) String() string { return proto.CompactTextString(m) }
func (*SetStateRequest_AgentProcess) ProtoMessage()    {}
func (*SetStateRequest_AgentProcess) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e8f797f46bba6fc, []int{6, 0}
}

func (m *SetStateRequest_AgentProcess) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetStateRequest_AgentProcess.Unmarshal(m, b)
}
func (m *SetStateRequest_AgentProcess) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetStateRequest_AgentProcess.Marshal(b, m, deterministic)
}
func (m *SetStateRequest_AgentProcess) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetStateRequest_AgentProcess.Merge(m, src)
}
func (m *SetStateRequest_AgentProcess) XXX_Size() int {
	return xxx_messageInfo_SetStateRequest_AgentProcess.Size(m)
}
func (m *SetStateRequest_AgentProcess) XXX_DiscardUnknown() {
	xxx_messageInfo_SetStateRequest_AgentProcess.DiscardUnknown(m)
}

var xxx_messageInfo_SetStateRequest_AgentProcess proto.InternalMessageInfo

func (m *SetStateRequest_AgentProcess) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_TYPE_INVALID
}

func (m *SetStateRequest_AgentProcess) GetTemplateLeftDelim() string {
	if m != nil {
		return m.TemplateLeftDelim
	}
	return ""
}

func (m *SetStateRequest_AgentProcess) GetTemplateRightDelim() string {
	if m != nil {
		return m.TemplateRightDelim
	}
	return ""
}

func (m *SetStateRequest_AgentProcess) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *SetStateRequest_AgentProcess) GetEnv() []string {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *SetStateRequest_AgentProcess) GetTextFiles() map[string]string {
	if m != nil {
		return m.TextFiles
	}
	return nil
}

// BuiltinAgent describes desired configuration of a single built-in agent for pmm-agent.
type SetStateRequest_BuiltinAgent struct {
	Type                 Type     `protobuf:"varint,1,opt,name=type,proto3,enum=agent.Type" json:"type,omitempty"`
	Dsn                  string   `protobuf:"bytes,2,opt,name=dsn,proto3" json:"dsn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetStateRequest_BuiltinAgent) Reset()         { *m = SetStateRequest_BuiltinAgent{} }
func (m *SetStateRequest_BuiltinAgent) String() string { return proto.CompactTextString(m) }
func (*SetStateRequest_BuiltinAgent) ProtoMessage()    {}
func (*SetStateRequest_BuiltinAgent) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e8f797f46bba6fc, []int{6, 2}
}

func (m *SetStateRequest_BuiltinAgent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetStateRequest_BuiltinAgent.Unmarshal(m, b)
}
func (m *SetStateRequest_BuiltinAgent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetStateRequest_BuiltinAgent.Marshal(b, m, deterministic)
}
func (m *SetStateRequest_BuiltinAgent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetStateRequest_BuiltinAgent.Merge(m, src)
}
func (m *SetStateRequest_BuiltinAgent) XXX_Size() int {
	return xxx_messageInfo_SetStateRequest_BuiltinAgent.Size(m)
}
func (m *SetStateRequest_BuiltinAgent) XXX_DiscardUnknown() {
	xxx_messageInfo_SetStateRequest_BuiltinAgent.DiscardUnknown(m)
}

var xxx_messageInfo_SetStateRequest_BuiltinAgent proto.InternalMessageInfo

func (m *SetStateRequest_BuiltinAgent) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_TYPE_INVALID
}

func (m *SetStateRequest_BuiltinAgent) GetDsn() string {
	if m != nil {
		return m.Dsn
	}
	return ""
}

// SetStateResponse is an AgentMessage for SetStateRequest acceptance.
type SetStateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetStateResponse) Reset()         { *m = SetStateResponse{} }
func (m *SetStateResponse) String() string { return proto.CompactTextString(m) }
func (*SetStateResponse) ProtoMessage()    {}
func (*SetStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e8f797f46bba6fc, []int{7}
}

func (m *SetStateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetStateResponse.Unmarshal(m, b)
}
func (m *SetStateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetStateResponse.Marshal(b, m, deterministic)
}
func (m *SetStateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetStateResponse.Merge(m, src)
}
func (m *SetStateResponse) XXX_Size() int {
	return xxx_messageInfo_SetStateResponse.Size(m)
}
func (m *SetStateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetStateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetStateResponse proto.InternalMessageInfo

type AgentMessage struct {
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to Payload:
	//	*AgentMessage_Ping
	//	*AgentMessage_StateChanged
	//	*AgentMessage_QanCollect
	//	*AgentMessage_Pong
	//	*AgentMessage_SetState
	Payload              isAgentMessage_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *AgentMessage) Reset()         { *m = AgentMessage{} }
func (m *AgentMessage) String() string { return proto.CompactTextString(m) }
func (*AgentMessage) ProtoMessage()    {}
func (*AgentMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e8f797f46bba6fc, []int{8}
}

func (m *AgentMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentMessage.Unmarshal(m, b)
}
func (m *AgentMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentMessage.Marshal(b, m, deterministic)
}
func (m *AgentMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentMessage.Merge(m, src)
}
func (m *AgentMessage) XXX_Size() int {
	return xxx_messageInfo_AgentMessage.Size(m)
}
func (m *AgentMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AgentMessage proto.InternalMessageInfo

func (m *AgentMessage) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type isAgentMessage_Payload interface {
	isAgentMessage_Payload()
}

type AgentMessage_Ping struct {
	Ping *Ping `protobuf:"bytes,2,opt,name=ping,proto3,oneof"`
}

type AgentMessage_StateChanged struct {
	StateChanged *StateChangedRequest `protobuf:"bytes,3,opt,name=state_changed,json=stateChanged,proto3,oneof"`
}

type AgentMessage_QanCollect struct {
	QanCollect *QANCollectRequest `protobuf:"bytes,4,opt,name=qan_collect,json=qanCollect,proto3,oneof"`
}

type AgentMessage_Pong struct {
	Pong *Pong `protobuf:"bytes,8,opt,name=pong,proto3,oneof"`
}

type AgentMessage_SetState struct {
	SetState *SetStateResponse `protobuf:"bytes,9,opt,name=set_state,json=setState,proto3,oneof"`
}

func (*AgentMessage_Ping) isAgentMessage_Payload() {}

func (*AgentMessage_StateChanged) isAgentMessage_Payload() {}

func (*AgentMessage_QanCollect) isAgentMessage_Payload() {}

func (*AgentMessage_Pong) isAgentMessage_Payload() {}

func (*AgentMessage_SetState) isAgentMessage_Payload() {}

func (m *AgentMessage) GetPayload() isAgentMessage_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *AgentMessage) GetPing() *Ping {
	if x, ok := m.GetPayload().(*AgentMessage_Ping); ok {
		return x.Ping
	}
	return nil
}

func (m *AgentMessage) GetStateChanged() *StateChangedRequest {
	if x, ok := m.GetPayload().(*AgentMessage_StateChanged); ok {
		return x.StateChanged
	}
	return nil
}

func (m *AgentMessage) GetQanCollect() *QANCollectRequest {
	if x, ok := m.GetPayload().(*AgentMessage_QanCollect); ok {
		return x.QanCollect
	}
	return nil
}

func (m *AgentMessage) GetPong() *Pong {
	if x, ok := m.GetPayload().(*AgentMessage_Pong); ok {
		return x.Pong
	}
	return nil
}

func (m *AgentMessage) GetSetState() *SetStateResponse {
	if x, ok := m.GetPayload().(*AgentMessage_SetState); ok {
		return x.SetState
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AgentMessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AgentMessage_Ping)(nil),
		(*AgentMessage_StateChanged)(nil),
		(*AgentMessage_QanCollect)(nil),
		(*AgentMessage_Pong)(nil),
		(*AgentMessage_SetState)(nil),
	}
}

type ServerMessage struct {
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to Payload:
	//	*ServerMessage_Pong
	//	*ServerMessage_StateChanged
	//	*ServerMessage_QanCollect
	//	*ServerMessage_Ping
	//	*ServerMessage_SetState
	Payload              isServerMessage_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ServerMessage) Reset()         { *m = ServerMessage{} }
func (m *ServerMessage) String() string { return proto.CompactTextString(m) }
func (*ServerMessage) ProtoMessage()    {}
func (*ServerMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e8f797f46bba6fc, []int{9}
}

func (m *ServerMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerMessage.Unmarshal(m, b)
}
func (m *ServerMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerMessage.Marshal(b, m, deterministic)
}
func (m *ServerMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerMessage.Merge(m, src)
}
func (m *ServerMessage) XXX_Size() int {
	return xxx_messageInfo_ServerMessage.Size(m)
}
func (m *ServerMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ServerMessage proto.InternalMessageInfo

func (m *ServerMessage) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type isServerMessage_Payload interface {
	isServerMessage_Payload()
}

type ServerMessage_Pong struct {
	Pong *Pong `protobuf:"bytes,2,opt,name=pong,proto3,oneof"`
}

type ServerMessage_StateChanged struct {
	StateChanged *StateChangedResponse `protobuf:"bytes,3,opt,name=state_changed,json=stateChanged,proto3,oneof"`
}

type ServerMessage_QanCollect struct {
	QanCollect *QANCollectResponse `protobuf:"bytes,4,opt,name=qan_collect,json=qanCollect,proto3,oneof"`
}

type ServerMessage_Ping struct {
	Ping *Ping `protobuf:"bytes,8,opt,name=ping,proto3,oneof"`
}

type ServerMessage_SetState struct {
	SetState *SetStateRequest `protobuf:"bytes,9,opt,name=set_state,json=setState,proto3,oneof"`
}

func (*ServerMessage_Pong) isServerMessage_Payload() {}

func (*ServerMessage_StateChanged) isServerMessage_Payload() {}

func (*ServerMessage_QanCollect) isServerMessage_Payload() {}

func (*ServerMessage_Ping) isServerMessage_Payload() {}

func (*ServerMessage_SetState) isServerMessage_Payload() {}

func (m *ServerMessage) GetPayload() isServerMessage_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ServerMessage) GetPong() *Pong {
	if x, ok := m.GetPayload().(*ServerMessage_Pong); ok {
		return x.Pong
	}
	return nil
}

func (m *ServerMessage) GetStateChanged() *StateChangedResponse {
	if x, ok := m.GetPayload().(*ServerMessage_StateChanged); ok {
		return x.StateChanged
	}
	return nil
}

func (m *ServerMessage) GetQanCollect() *QANCollectResponse {
	if x, ok := m.GetPayload().(*ServerMessage_QanCollect); ok {
		return x.QanCollect
	}
	return nil
}

func (m *ServerMessage) GetPing() *Ping {
	if x, ok := m.GetPayload().(*ServerMessage_Ping); ok {
		return x.Ping
	}
	return nil
}

func (m *ServerMessage) GetSetState() *SetStateRequest {
	if x, ok := m.GetPayload().(*ServerMessage_SetState); ok {
		return x.SetState
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ServerMessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ServerMessage_Pong)(nil),
		(*ServerMessage_StateChanged)(nil),
		(*ServerMessage_QanCollect)(nil),
		(*ServerMessage_Ping)(nil),
		(*ServerMessage_SetState)(nil),
	}
}

func init() {
	proto.RegisterEnum("agent.Type", Type_name, Type_value)
	proto.RegisterType((*Ping)(nil), "agent.Ping")
	proto.RegisterType((*Pong)(nil), "agent.Pong")
	proto.RegisterType((*QANCollectRequest)(nil), "agent.QANCollectRequest")
	proto.RegisterType((*QANCollectResponse)(nil), "agent.QANCollectResponse")
	proto.RegisterType((*StateChangedRequest)(nil), "agent.StateChangedRequest")
	proto.RegisterType((*StateChangedResponse)(nil), "agent.StateChangedResponse")
	proto.RegisterType((*SetStateRequest)(nil), "agent.SetStateRequest")
	proto.RegisterMapType((map[string]*SetStateRequest_AgentProcess)(nil), "agent.SetStateRequest.AgentProcessesEntry")
	proto.RegisterMapType((map[string]*SetStateRequest_BuiltinAgent)(nil), "agent.SetStateRequest.BuiltinAgentsEntry")
	proto.RegisterType((*SetStateRequest_AgentProcess)(nil), "agent.SetStateRequest.AgentProcess")
	proto.RegisterMapType((map[string]string)(nil), "agent.SetStateRequest.AgentProcess.TextFilesEntry")
	proto.RegisterType((*SetStateRequest_BuiltinAgent)(nil), "agent.SetStateRequest.BuiltinAgent")
	proto.RegisterType((*SetStateResponse)(nil), "agent.SetStateResponse")
	proto.RegisterType((*AgentMessage)(nil), "agent.AgentMessage")
	proto.RegisterType((*ServerMessage)(nil), "agent.ServerMessage")
}

func init() { proto.RegisterFile("agentpb/agent.proto", fileDescriptor_0e8f797f46bba6fc) }

var fileDescriptor_0e8f797f46bba6fc = []byte{
	// 965 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xdd, 0x6e, 0xe2, 0x46,
	0x14, 0x8e, 0x81, 0x40, 0x7c, 0x08, 0x84, 0x0c, 0x6c, 0xd6, 0xcb, 0x4a, 0x9b, 0x94, 0xde, 0xd0,
	0x95, 0xea, 0xac, 0xa8, 0x5a, 0x6d, 0xdb, 0xed, 0x05, 0x04, 0xe7, 0x47, 0xe2, 0xc7, 0x19, 0xa3,
	0xb6, 0xdb, 0x1b, 0xcb, 0x84, 0x89, 0xd7, 0xaa, 0x33, 0x36, 0xf6, 0x10, 0x2d, 0x97, 0x7d, 0x88,
	0x3e, 0x43, 0xd5, 0x37, 0xe9, 0xbb, 0xf4, 0x25, 0x2a, 0xcf, 0x8c, 0xc1, 0x2c, 0x6c, 0xbb, 0xea,
	0x55, 0x3c, 0xe7, 0xe7, 0x3b, 0xdf, 0x7c, 0x73, 0xce, 0x09, 0x50, 0x77, 0x5c, 0x42, 0x59, 0x38,
	0x3d, 0xe7, 0x7f, 0xf5, 0x30, 0x0a, 0x58, 0x80, 0xf6, 0xf9, 0xa1, 0x79, 0xea, 0x06, 0x81, 0xeb,
	0x93, 0x73, 0x6e, 0x9c, 0x2e, 0xee, 0xcf, 0x99, 0xf7, 0x40, 0x62, 0xe6, 0x3c, 0x84, 0x22, 0xae,
	0xf9, 0xc2, 0xa3, 0x8f, 0x84, 0xb2, 0x20, 0x5a, 0xa6, 0x00, 0x76, 0xcc, 0x1c, 0xb6, 0x88, 0xa5,
	0xff, 0xc9, 0xdc, 0xa1, 0xe1, 0xf4, 0xfc, 0x2e, 0xf0, 0x7d, 0x72, 0xc7, 0x82, 0x48, 0x98, 0x5b,
	0x45, 0x28, 0x98, 0x1e, 0x75, 0x5b, 0x06, 0x14, 0xcc, 0x80, 0xba, 0xe8, 0x07, 0x38, 0xbc, 0x5b,
	0x44, 0x51, 0x92, 0x9e, 0x54, 0xd0, 0x94, 0x33, 0xa5, 0x5d, 0xee, 0x34, 0x75, 0x51, 0x5e, 0x4f,
	0xcb, 0xeb, 0x93, 0xb4, 0x3c, 0x2e, 0xcb, 0xf8, 0xc4, 0xd2, 0xea, 0xc1, 0xf1, 0x6d, 0x77, 0x74,
	0x21, 0x8a, 0x60, 0x32, 0x5f, 0x90, 0x98, 0xa1, 0x2f, 0xa1, 0xf4, 0x40, 0xe2, 0xd8, 0x71, 0x53,
	0xb8, 0xba, 0x3e, 0x77, 0xa8, 0xbe, 0x19, 0x85, 0xd3, 0x98, 0x56, 0x03, 0x50, 0x16, 0x23, 0x0e,
	0x03, 0x1a, 0x93, 0xd6, 0x6f, 0x0a, 0xd4, 0x2d, 0xe6, 0x30, 0x72, 0xf1, 0xce, 0xa1, 0x2e, 0x99,
	0xa5, 0xe0, 0xcf, 0xe0, 0x40, 0xdc, 0xd6, 0x9b, 0x71, 0x74, 0x15, 0x97, 0xf8, 0xf9, 0x66, 0x86,
	0x74, 0x28, 0x0a, 0x09, 0xb4, 0xdc, 0x99, 0xd2, 0xae, 0x76, 0x4e, 0xf4, 0x95, 0x46, 0x7a, 0x37,
	0x89, 0xb1, 0xb8, 0x17, 0xcb, 0x28, 0x74, 0x0a, 0x65, 0xdf, 0x8b, 0x19, 0xa1, 0x76, 0x18, 0x44,
	0x4c, 0xcb, 0x9f, 0x29, 0xed, 0x0a, 0x06, 0x61, 0x32, 0x83, 0x88, 0xb5, 0x4e, 0xa0, 0xb1, 0x49,
	0x41, 0x72, 0xfb, 0xbd, 0x08, 0x47, 0x16, 0xe1, 0x70, 0x24, 0xe5, 0x65, 0xc1, 0x91, 0xe0, 0x15,
	0x46, 0xc1, 0x1d, 0x89, 0x63, 0x12, 0x6b, 0xca, 0x59, 0xbe, 0x5d, 0xee, 0xbc, 0xd4, 0xc5, 0xf3,
	0x7e, 0x90, 0x20, 0x18, 0x99, 0x69, 0xb0, 0x41, 0x59, 0xb4, 0xc4, 0x55, 0x67, 0xc3, 0x88, 0x4c,
	0xa8, 0x4e, 0x17, 0x9e, 0xcf, 0x3c, 0x6a, 0x73, 0x4f, 0x72, 0xb3, 0x04, 0xf3, 0x8b, 0x8f, 0x60,
	0xf6, 0x44, 0x30, 0x87, 0x96, 0x90, 0x95, 0x69, 0xd6, 0xd6, 0xfc, 0x2b, 0x07, 0x87, 0xd9, 0xca,
	0xe8, 0x14, 0x0a, 0x6c, 0x19, 0x8a, 0x97, 0xaa, 0x76, 0xca, 0x12, 0x78, 0xb2, 0x0c, 0x09, 0xe6,
	0x0e, 0xa4, 0x43, 0x9d, 0x91, 0x87, 0xd0, 0x77, 0x18, 0xb1, 0x7d, 0x72, 0xcf, 0xec, 0x19, 0xf1,
	0xbd, 0x07, 0x2e, 0xb1, 0x8a, 0x8f, 0x53, 0xd7, 0x80, 0xdc, 0xb3, 0x7e, 0xe2, 0x40, 0xaf, 0xa0,
	0xb1, 0x8a, 0x8f, 0x3c, 0xf7, 0x5d, 0x9a, 0x90, 0xe7, 0x09, 0x28, 0xf5, 0xe1, 0xc4, 0x25, 0x32,
	0x10, 0x14, 0x9c, 0xc8, 0x8d, 0xb5, 0xc2, 0x59, 0xbe, 0xad, 0x62, 0xfe, 0x8d, 0x6a, 0x90, 0x27,
	0xf4, 0x51, 0xdb, 0xe7, 0xa6, 0xe4, 0x13, 0xdd, 0x02, 0x30, 0xf2, 0x9e, 0xd9, 0xf7, 0x9e, 0x4f,
	0x62, 0xad, 0xc8, 0x75, 0xe8, 0x7c, 0x82, 0xb6, 0xfa, 0x84, 0xbc, 0x67, 0x97, 0x49, 0x92, 0x10,
	0x44, 0x65, 0xe9, 0xb9, 0xf9, 0x06, 0xaa, 0x9b, 0xce, 0xa4, 0xec, 0xaf, 0x64, 0x29, 0x1b, 0x2b,
	0xf9, 0x44, 0x0d, 0xd8, 0x7f, 0x74, 0xfc, 0x05, 0x91, 0x17, 0x16, 0x87, 0xef, 0x72, 0xaf, 0x95,
	0xe6, 0x3d, 0xd4, 0x77, 0xbc, 0xe1, 0x0e, 0x88, 0x6f, 0xb3, 0x10, 0xe5, 0xce, 0xe7, 0x9f, 0x40,
	0x3a, 0x5b, 0xa7, 0x0b, 0x87, 0xd9, 0x77, 0xfd, 0xef, 0x17, 0xab, 0x41, 0x7e, 0x16, 0x53, 0x49,
	0x38, 0xf9, 0x6c, 0x12, 0x40, 0xdb, 0xad, 0xf1, 0xff, 0x99, 0x66, 0xb1, 0x32, 0x4c, 0x5b, 0x08,
	0x6a, 0xeb, 0x50, 0x39, 0x2b, 0x7f, 0xa4, 0x0d, 0x37, 0x14, 0xe3, 0x8e, 0xaa, 0x90, 0x93, 0xa3,
	0x5b, 0xc1, 0x39, 0x6f, 0x86, 0x3e, 0x83, 0x42, 0xe8, 0x51, 0x57, 0x96, 0x4c, 0xaf, 0x93, 0x2c,
	0xa9, 0xeb, 0x3d, 0xcc, 0x5d, 0xa8, 0x0b, 0x95, 0x64, 0x64, 0x89, 0x7d, 0x27, 0x06, 0x91, 0xf7,
	0x52, 0xb2, 0xa5, 0x24, 0xbd, 0xed, 0x35, 0x71, 0xbd, 0x87, 0x0f, 0xe3, 0x8c, 0x19, 0x7d, 0x0f,
	0xe5, 0xb9, 0x43, 0x6d, 0xb9, 0x0e, 0xb5, 0x02, 0x07, 0xd0, 0x24, 0xc0, 0xd6, 0x0a, 0xbb, 0xde,
	0xc3, 0x30, 0x77, 0xa8, 0x34, 0x72, 0x8a, 0x01, 0x75, 0xb5, 0x83, 0x4d, 0x8a, 0x81, 0xa4, 0x98,
	0xec, 0xd1, 0x6f, 0x40, 0x8d, 0x89, 0x58, 0xc1, 0x44, 0x53, 0x79, 0xdc, 0xd3, 0x2d, 0xf5, 0x84,
	0x24, 0xd7, 0x7b, 0xf8, 0x20, 0x96, 0xb6, 0x9e, 0x0a, 0xa5, 0xd0, 0x59, 0xfa, 0x81, 0x33, 0x6b,
	0xfd, 0x99, 0x83, 0x8a, 0x45, 0xa2, 0x47, 0x12, 0xfd, 0x9b, 0x54, 0xc1, 0xb6, 0x54, 0x59, 0x1e,
	0xbd, 0xdd, 0x52, 0x3d, 0xdf, 0x29, 0xd5, 0x8a, 0xcf, 0xa6, 0x56, 0x6f, 0x76, 0x69, 0xf5, 0x6c,
	0x87, 0x56, 0xab, 0xfc, 0x0f, 0xc5, 0xf2, 0xb6, 0xc5, 0xca, 0xbe, 0xe7, 0xd7, 0xdb, 0x62, 0x9d,
	0xec, 0x6e, 0xb5, 0x8f, 0x68, 0xf5, 0xf2, 0x6f, 0x05, 0x0a, 0x13, 0xd1, 0xeb, 0x87, 0x93, 0xb7,
	0xa6, 0x61, 0xdf, 0x8c, 0x7e, 0xec, 0x0e, 0x6e, 0xfa, 0xb5, 0x3d, 0x54, 0x01, 0xd5, 0x1c, 0x0e,
	0xed, 0xee, 0x95, 0x31, 0x9a, 0xd4, 0x14, 0x74, 0x0c, 0x95, 0xd1, 0xb8, 0x6f, 0xd8, 0xc6, 0xcf,
	0xe6, 0x18, 0x4f, 0x0c, 0x5c, 0xcb, 0xa1, 0x3a, 0x1c, 0x0d, 0xdf, 0x5a, 0xb7, 0x83, 0xfe, 0xda,
	0x98, 0x47, 0x0d, 0xa8, 0x0d, 0xc7, 0xa3, 0xab, 0x71, 0xbf, 0xb7, 0xb6, 0x16, 0xd0, 0x13, 0x38,
	0x36, 0xc7, 0xd6, 0xe4, 0x0a, 0x1b, 0xd6, 0xda, 0xbc, 0x8f, 0x5e, 0x40, 0xf3, 0xb6, 0x3b, 0xb2,
	0x39, 0x8a, 0x6d, 0x1a, 0xf8, 0xd2, 0xba, 0xb8, 0x36, 0x86, 0x5d, 0x59, 0xb4, 0x88, 0x9e, 0xc3,
	0xd3, 0xb5, 0xdf, 0x1a, 0x8c, 0x7f, 0x1a, 0x8c, 0xaf, 0xa4, 0xb3, 0xb4, 0x4a, 0x96, 0xd5, 0x4c,
	0x3c, 0xbe, 0xbc, 0x19, 0x18, 0x58, 0xfa, 0x0f, 0x92, 0x2b, 0xe1, 0x7e, 0xa6, 0x9c, 0xda, 0xe9,
	0xc2, 0xbe, 0x18, 0xfd, 0xd7, 0x50, 0xba, 0x08, 0x28, 0x4d, 0x64, 0xae, 0x4b, 0xc1, 0xb2, 0xb3,
	0xd5, 0x6c, 0xac, 0x54, 0xcc, 0xb4, 0x51, 0x5b, 0x79, 0xa5, 0xf4, 0xd4, 0x5f, 0x4a, 0xf2, 0xd7,
	0xc6, 0xb4, 0xc8, 0xff, 0xa9, 0x7f, 0xf5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x00, 0x3e,
	0x5e, 0x7f, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentClient interface {
	// Connect establishes two-way communication channel between pmm-agent and pmm-managed.
	Connect(ctx context.Context, opts ...grpc.CallOption) (Agent_ConnectClient, error)
}

type agentClient struct {
	cc *grpc.ClientConn
}

func NewAgentClient(cc *grpc.ClientConn) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) Connect(ctx context.Context, opts ...grpc.CallOption) (Agent_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Agent_serviceDesc.Streams[0], "/agent.Agent/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentConnectClient{stream}
	return x, nil
}

type Agent_ConnectClient interface {
	Send(*AgentMessage) error
	Recv() (*ServerMessage, error)
	grpc.ClientStream
}

type agentConnectClient struct {
	grpc.ClientStream
}

func (x *agentConnectClient) Send(m *AgentMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *agentConnectClient) Recv() (*ServerMessage, error) {
	m := new(ServerMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AgentServer is the server API for Agent service.
type AgentServer interface {
	// Connect establishes two-way communication channel between pmm-agent and pmm-managed.
	Connect(Agent_ConnectServer) error
}

func RegisterAgentServer(s *grpc.Server, srv AgentServer) {
	s.RegisterService(&_Agent_serviceDesc, srv)
}

func _Agent_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentServer).Connect(&agentConnectServer{stream})
}

type Agent_ConnectServer interface {
	Send(*ServerMessage) error
	Recv() (*AgentMessage, error)
	grpc.ServerStream
}

type agentConnectServer struct {
	grpc.ServerStream
}

func (x *agentConnectServer) Send(m *ServerMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *agentConnectServer) Recv() (*AgentMessage, error) {
	m := new(AgentMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Agent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "agent.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _Agent_Connect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "agentpb/agent.proto",
}

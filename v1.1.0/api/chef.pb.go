// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/ingest/request/chef.proto

package request

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _struct "github.com/golang/protobuf/ptypes/struct"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Run struct {
	// 1 through 15 are for frequently occuring fields
	// Reserving for shared fields between run_start and run_converge mesages.
	Id                   string           `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	RunId                string           `protobuf:"bytes,2,opt,name=run_id,json=runId" json:"run_id,omitempty"`
	EntityUuid           string           `protobuf:"bytes,3,opt,name=entity_uuid,json=entityUuid" json:"entity_uuid,omitempty"`
	MessageVersion       string           `protobuf:"bytes,4,opt,name=message_version,json=messageVersion" json:"message_version,omitempty"`
	MessageType          string           `protobuf:"bytes,5,opt,name=message_type,json=messageType" json:"message_type,omitempty"`
	NodeName             string           `protobuf:"bytes,6,opt,name=node_name,json=nodeName" json:"node_name,omitempty"`
	OrganizationName     string           `protobuf:"bytes,7,opt,name=organization_name,json=organizationName" json:"organization_name,omitempty"`
	StartTime            string           `protobuf:"bytes,8,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	ChefServerFqdn       string           `protobuf:"bytes,9,opt,name=chef_server_fqdn,json=chefServerFqdn" json:"chef_server_fqdn,omitempty"`
	EndTime              string           `protobuf:"bytes,16,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	Status               string           `protobuf:"bytes,17,opt,name=status" json:"status,omitempty"`
	TotalResourceCount   int32            `protobuf:"varint,18,opt,name=total_resource_count,json=totalResourceCount" json:"total_resource_count,omitempty"`
	UpdatedResourceCount int32            `protobuf:"varint,19,opt,name=updated_resource_count,json=updatedResourceCount" json:"updated_resource_count,omitempty"`
	Source               string           `protobuf:"bytes,20,opt,name=source" json:"source,omitempty"`
	ExpandedRunList      *ExpandedRunList `protobuf:"bytes,21,opt,name=expanded_run_list,json=expandedRunList" json:"expanded_run_list,omitempty"`
	Resources            []*Resource      `protobuf:"bytes,22,rep,name=resources" json:"resources,omitempty"`
	RunList              []string         `protobuf:"bytes,23,rep,name=run_list,json=runList" json:"run_list,omitempty"`
	Node                 *_struct.Struct  `protobuf:"bytes,24,opt,name=node" json:"node,omitempty"`
	Error                *Error           `protobuf:"bytes,25,opt,name=error" json:"error,omitempty"`
	PolicyName           string           `protobuf:"bytes,26,opt,name=policy_name,json=policyName" json:"policy_name,omitempty"`
	PolicyGroup          string           `protobuf:"bytes,27,opt,name=policy_group,json=policyGroup" json:"policy_group,omitempty"`
	Deprecations         []*Deprecation   `protobuf:"bytes,28,rep,name=deprecations" json:"deprecations,omitempty"`
	Tags                 []string         `protobuf:"bytes,29,rep,name=tags" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Run) Reset()         { *m = Run{} }
func (m *Run) String() string { return proto.CompactTextString(m) }
func (*Run) ProtoMessage()    {}
func (*Run) Descriptor() ([]byte, []int) {
	return fileDescriptor_chef_32acb3734a93ad7a, []int{0}
}
func (m *Run) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Run.Unmarshal(m, b)
}
func (m *Run) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Run.Marshal(b, m, deterministic)
}
func (dst *Run) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Run.Merge(dst, src)
}
func (m *Run) XXX_Size() int {
	return xxx_messageInfo_Run.Size(m)
}
func (m *Run) XXX_DiscardUnknown() {
	xxx_messageInfo_Run.DiscardUnknown(m)
}

var xxx_messageInfo_Run proto.InternalMessageInfo

func (m *Run) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Run) GetRunId() string {
	if m != nil {
		return m.RunId
	}
	return ""
}

func (m *Run) GetEntityUuid() string {
	if m != nil {
		return m.EntityUuid
	}
	return ""
}

func (m *Run) GetMessageVersion() string {
	if m != nil {
		return m.MessageVersion
	}
	return ""
}

func (m *Run) GetMessageType() string {
	if m != nil {
		return m.MessageType
	}
	return ""
}

func (m *Run) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Run) GetOrganizationName() string {
	if m != nil {
		return m.OrganizationName
	}
	return ""
}

func (m *Run) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *Run) GetChefServerFqdn() string {
	if m != nil {
		return m.ChefServerFqdn
	}
	return ""
}

func (m *Run) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *Run) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Run) GetTotalResourceCount() int32 {
	if m != nil {
		return m.TotalResourceCount
	}
	return 0
}

func (m *Run) GetUpdatedResourceCount() int32 {
	if m != nil {
		return m.UpdatedResourceCount
	}
	return 0
}

func (m *Run) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Run) GetExpandedRunList() *ExpandedRunList {
	if m != nil {
		return m.ExpandedRunList
	}
	return nil
}

func (m *Run) GetResources() []*Resource {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *Run) GetRunList() []string {
	if m != nil {
		return m.RunList
	}
	return nil
}

func (m *Run) GetNode() *_struct.Struct {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *Run) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Run) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

func (m *Run) GetPolicyGroup() string {
	if m != nil {
		return m.PolicyGroup
	}
	return ""
}

func (m *Run) GetDeprecations() []*Deprecation {
	if m != nil {
		return m.Deprecations
	}
	return nil
}

func (m *Run) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type Deprecation struct {
	Message              string   `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	Location             string   `protobuf:"bytes,3,opt,name=location" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Deprecation) Reset()         { *m = Deprecation{} }
func (m *Deprecation) String() string { return proto.CompactTextString(m) }
func (*Deprecation) ProtoMessage()    {}
func (*Deprecation) Descriptor() ([]byte, []int) {
	return fileDescriptor_chef_32acb3734a93ad7a, []int{1}
}
func (m *Deprecation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Deprecation.Unmarshal(m, b)
}
func (m *Deprecation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Deprecation.Marshal(b, m, deterministic)
}
func (dst *Deprecation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deprecation.Merge(dst, src)
}
func (m *Deprecation) XXX_Size() int {
	return xxx_messageInfo_Deprecation.Size(m)
}
func (m *Deprecation) XXX_DiscardUnknown() {
	xxx_messageInfo_Deprecation.DiscardUnknown(m)
}

var xxx_messageInfo_Deprecation proto.InternalMessageInfo

func (m *Deprecation) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Deprecation) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Deprecation) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

type ExpandedRunList struct {
	Id                   string     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	RunList              []*RunList `protobuf:"bytes,2,rep,name=run_list,json=runList" json:"run_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ExpandedRunList) Reset()         { *m = ExpandedRunList{} }
func (m *ExpandedRunList) String() string { return proto.CompactTextString(m) }
func (*ExpandedRunList) ProtoMessage()    {}
func (*ExpandedRunList) Descriptor() ([]byte, []int) {
	return fileDescriptor_chef_32acb3734a93ad7a, []int{2}
}
func (m *ExpandedRunList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExpandedRunList.Unmarshal(m, b)
}
func (m *ExpandedRunList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExpandedRunList.Marshal(b, m, deterministic)
}
func (dst *ExpandedRunList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExpandedRunList.Merge(dst, src)
}
func (m *ExpandedRunList) XXX_Size() int {
	return xxx_messageInfo_ExpandedRunList.Size(m)
}
func (m *ExpandedRunList) XXX_DiscardUnknown() {
	xxx_messageInfo_ExpandedRunList.DiscardUnknown(m)
}

var xxx_messageInfo_ExpandedRunList proto.InternalMessageInfo

func (m *ExpandedRunList) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ExpandedRunList) GetRunList() []*RunList {
	if m != nil {
		return m.RunList
	}
	return nil
}

type RunList struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Version              string   `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
	Skipped              bool     `protobuf:"varint,4,opt,name=skipped" json:"skipped,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunList) Reset()         { *m = RunList{} }
func (m *RunList) String() string { return proto.CompactTextString(m) }
func (*RunList) ProtoMessage()    {}
func (*RunList) Descriptor() ([]byte, []int) {
	return fileDescriptor_chef_32acb3734a93ad7a, []int{3}
}
func (m *RunList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunList.Unmarshal(m, b)
}
func (m *RunList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunList.Marshal(b, m, deterministic)
}
func (dst *RunList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunList.Merge(dst, src)
}
func (m *RunList) XXX_Size() int {
	return xxx_messageInfo_RunList.Size(m)
}
func (m *RunList) XXX_DiscardUnknown() {
	xxx_messageInfo_RunList.DiscardUnknown(m)
}

var xxx_messageInfo_RunList proto.InternalMessageInfo

func (m *RunList) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *RunList) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RunList) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *RunList) GetSkipped() bool {
	if m != nil {
		return m.Skipped
	}
	return false
}

type Resource struct {
	Type                 string          `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Name                 string          `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Id                   string          `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	After                *_struct.Struct `protobuf:"bytes,4,opt,name=after" json:"after,omitempty"`
	Before               *_struct.Struct `protobuf:"bytes,5,opt,name=before" json:"before,omitempty"`
	Duration             string          `protobuf:"bytes,6,opt,name=duration" json:"duration,omitempty"`
	Delta                string          `protobuf:"bytes,7,opt,name=delta" json:"delta,omitempty"`
	CookbookName         string          `protobuf:"bytes,8,opt,name=cookbook_name,json=cookbookName" json:"cookbook_name,omitempty"`
	CookbookVersion      string          `protobuf:"bytes,9,opt,name=cookbook_version,json=cookbookVersion" json:"cookbook_version,omitempty"`
	IgnoreFailure        bool            `protobuf:"varint,10,opt,name=ignore_failure,json=ignoreFailure" json:"ignore_failure,omitempty"`
	Status               string          `protobuf:"bytes,11,opt,name=status" json:"status,omitempty"`
	RecipeName           string          `protobuf:"bytes,12,opt,name=recipe_name,json=recipeName" json:"recipe_name,omitempty"`
	Conditional          string          `protobuf:"bytes,16,opt,name=conditional" json:"conditional,omitempty"`
	Result               string          `protobuf:"bytes,17,opt,name=result" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_chef_32acb3734a93ad7a, []int{4}
}
func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (dst *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(dst, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Resource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Resource) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Resource) GetAfter() *_struct.Struct {
	if m != nil {
		return m.After
	}
	return nil
}

func (m *Resource) GetBefore() *_struct.Struct {
	if m != nil {
		return m.Before
	}
	return nil
}

func (m *Resource) GetDuration() string {
	if m != nil {
		return m.Duration
	}
	return ""
}

func (m *Resource) GetDelta() string {
	if m != nil {
		return m.Delta
	}
	return ""
}

func (m *Resource) GetCookbookName() string {
	if m != nil {
		return m.CookbookName
	}
	return ""
}

func (m *Resource) GetCookbookVersion() string {
	if m != nil {
		return m.CookbookVersion
	}
	return ""
}

func (m *Resource) GetIgnoreFailure() bool {
	if m != nil {
		return m.IgnoreFailure
	}
	return false
}

func (m *Resource) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Resource) GetRecipeName() string {
	if m != nil {
		return m.RecipeName
	}
	return ""
}

func (m *Resource) GetConditional() string {
	if m != nil {
		return m.Conditional
	}
	return ""
}

func (m *Resource) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type Error struct {
	Class                string       `protobuf:"bytes,1,opt,name=class" json:"class,omitempty"`
	Message              string       `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	Backtrace            []string     `protobuf:"bytes,16,rep,name=backtrace" json:"backtrace,omitempty"`
	Description          *Description `protobuf:"bytes,17,opt,name=description" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_chef_32acb3734a93ad7a, []int{5}
}
func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (dst *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(dst, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetClass() string {
	if m != nil {
		return m.Class
	}
	return ""
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Error) GetBacktrace() []string {
	if m != nil {
		return m.Backtrace
	}
	return nil
}

func (m *Error) GetDescription() *Description {
	if m != nil {
		return m.Description
	}
	return nil
}

type Description struct {
	Title                string            `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Sections             []*_struct.Struct `protobuf:"bytes,2,rep,name=sections" json:"sections,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Description) Reset()         { *m = Description{} }
func (m *Description) String() string { return proto.CompactTextString(m) }
func (*Description) ProtoMessage()    {}
func (*Description) Descriptor() ([]byte, []int) {
	return fileDescriptor_chef_32acb3734a93ad7a, []int{6}
}
func (m *Description) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Description.Unmarshal(m, b)
}
func (m *Description) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Description.Marshal(b, m, deterministic)
}
func (dst *Description) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Description.Merge(dst, src)
}
func (m *Description) XXX_Size() int {
	return xxx_messageInfo_Description.Size(m)
}
func (m *Description) XXX_DiscardUnknown() {
	xxx_messageInfo_Description.DiscardUnknown(m)
}

var xxx_messageInfo_Description proto.InternalMessageInfo

func (m *Description) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Description) GetSections() []*_struct.Struct {
	if m != nil {
		return m.Sections
	}
	return nil
}

func init() {
	proto.RegisterType((*Run)(nil), "chef.automate.api.ingest.request.Run")
	proto.RegisterType((*Deprecation)(nil), "chef.automate.api.ingest.request.Deprecation")
	proto.RegisterType((*ExpandedRunList)(nil), "chef.automate.api.ingest.request.ExpandedRunList")
	proto.RegisterType((*RunList)(nil), "chef.automate.api.ingest.request.RunList")
	proto.RegisterType((*Resource)(nil), "chef.automate.api.ingest.request.Resource")
	proto.RegisterType((*Error)(nil), "chef.automate.api.ingest.request.Error")
	proto.RegisterType((*Description)(nil), "chef.automate.api.ingest.request.Description")
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/ingest/request/chef.proto", fileDescriptor_chef_32acb3734a93ad7a)
}

var fileDescriptor_chef_32acb3734a93ad7a = []byte{
	// 949 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0x6f, 0x6f, 0x1b, 0xc5,
	0x13, 0x56, 0xe2, 0x38, 0xb1, 0xc7, 0x69, 0xfe, 0xec, 0x2f, 0x4d, 0xb7, 0x69, 0xaa, 0xba, 0xfe,
	0x09, 0x35, 0xa5, 0x8a, 0x0d, 0x29, 0x6f, 0x78, 0xc1, 0x0b, 0x68, 0x29, 0x20, 0x21, 0x10, 0xd7,
	0x16, 0x21, 0x24, 0x74, 0x5a, 0xdf, 0x8d, 0xaf, 0xab, 0x9c, 0x77, 0x2f, 0xfb, 0xa7, 0x10, 0x3e,
	0x10, 0x1f, 0x82, 0x17, 0x7c, 0x36, 0xb4, 0xb3, 0x7b, 0x8e, 0x5d, 0x44, 0x13, 0x5e, 0xf9, 0xe6,
	0xd9, 0x67, 0x66, 0x67, 0x67, 0x9e, 0x19, 0x19, 0x3e, 0x2d, 0xf4, 0xbc, 0xd1, 0x0a, 0x95, 0xb3,
	0x13, 0xe1, 0x9d, 0x9e, 0x0b, 0x87, 0xa7, 0x95, 0x70, 0xf8, 0xab, 0xb8, 0x9c, 0x88, 0x46, 0x4e,
	0xa4, 0xaa, 0xd0, 0xba, 0x89, 0xc1, 0x0b, 0x1f, 0x7e, 0x8b, 0x37, 0x38, 0x1b, 0x37, 0x46, 0x3b,
	0xcd, 0x86, 0xf4, 0xdd, 0x3a, 0x8d, 0x45, 0x23, 0xc7, 0x91, 0x3c, 0x4e, 0xe4, 0xa3, 0xe3, 0x4a,
	0xeb, 0xaa, 0xc6, 0x09, 0xf1, 0xa7, 0x7e, 0x36, 0xb1, 0xce, 0xf8, 0xc2, 0x45, 0xff, 0xd1, 0x9f,
	0x5b, 0xd0, 0xc9, 0xbc, 0x62, 0x3b, 0xb0, 0x2e, 0x4b, 0xbe, 0x36, 0x5c, 0x3b, 0xe9, 0x67, 0xeb,
	0xb2, 0x64, 0xb7, 0x61, 0xd3, 0x78, 0x95, 0xcb, 0x92, 0xaf, 0x13, 0xd6, 0x35, 0x5e, 0x7d, 0x53,
	0xb2, 0x07, 0x30, 0x40, 0xe5, 0xa4, 0xbb, 0xcc, 0xbd, 0x97, 0x25, 0xef, 0xd0, 0x19, 0x44, 0xe8,
	0xb5, 0x97, 0x25, 0x7b, 0x04, 0xbb, 0x73, 0xb4, 0x56, 0x54, 0x98, 0xbf, 0x45, 0x63, 0xa5, 0x56,
	0x7c, 0x83, 0x48, 0x3b, 0x09, 0xfe, 0x31, 0xa2, 0xec, 0x21, 0x6c, 0xb7, 0x44, 0x77, 0xd9, 0x20,
	0xef, 0x12, 0x6b, 0x90, 0xb0, 0x57, 0x97, 0x0d, 0xb2, 0x7b, 0xd0, 0x57, 0xba, 0xc4, 0x5c, 0x89,
	0x39, 0xf2, 0x4d, 0x3a, 0xef, 0x05, 0xe0, 0x3b, 0x31, 0x47, 0xf6, 0x04, 0xf6, 0xb5, 0xa9, 0x84,
	0x92, 0xbf, 0x0b, 0x27, 0xb5, 0x8a, 0xa4, 0x2d, 0x22, 0xed, 0x2d, 0x1f, 0x10, 0xf9, 0x3e, 0x80,
	0x75, 0xc2, 0xb8, 0xdc, 0xc9, 0x39, 0xf2, 0x1e, 0xb1, 0xfa, 0x84, 0xbc, 0x92, 0x73, 0x64, 0x27,
	0xb0, 0x17, 0xca, 0x98, 0x5b, 0x34, 0x6f, 0xd1, 0xe4, 0xb3, 0x8b, 0x52, 0xf1, 0x7e, 0xcc, 0x3a,
	0xe0, 0x2f, 0x09, 0x7e, 0x71, 0x51, 0x2a, 0x76, 0x17, 0x7a, 0xa8, 0xca, 0x18, 0x66, 0x8f, 0x18,
	0x5b, 0xa8, 0x4a, 0x0a, 0x72, 0x08, 0x9b, 0xd6, 0x09, 0xe7, 0x2d, 0xdf, 0xa7, 0x83, 0x64, 0xb1,
	0x8f, 0xe0, 0xc0, 0x69, 0x27, 0xea, 0xdc, 0xa0, 0xd5, 0xde, 0x14, 0x98, 0x17, 0xda, 0x2b, 0xc7,
	0xd9, 0x70, 0xed, 0xa4, 0x9b, 0x31, 0x3a, 0xcb, 0xd2, 0xd1, 0xb3, 0x70, 0xc2, 0x3e, 0x81, 0x43,
	0xdf, 0x94, 0xc2, 0x61, 0xf9, 0xae, 0xcf, 0xff, 0xc8, 0xe7, 0x20, 0x9d, 0xae, 0x7a, 0x85, 0xfb,
	0xc9, 0xe4, 0x07, 0xe9, 0x7e, 0xb2, 0xd8, 0x2f, 0xb0, 0x8f, 0xbf, 0x35, 0x42, 0x95, 0x21, 0x9c,
	0x57, 0x79, 0x2d, 0xad, 0xe3, 0xb7, 0x87, 0x6b, 0x27, 0x83, 0xb3, 0x8f, 0xc7, 0xd7, 0xa9, 0x67,
	0xfc, 0x65, 0x72, 0xcd, 0xbc, 0xfa, 0x56, 0x5a, 0x97, 0xed, 0xe2, 0x2a, 0xc0, 0xbe, 0x86, 0x7e,
	0x9b, 0xa4, 0xe5, 0x87, 0xc3, 0xce, 0xc9, 0xe0, 0xec, 0xc3, 0xeb, 0xc3, 0xb6, 0xa9, 0x67, 0x57,
	0xce, 0xa1, 0xb6, 0x8b, 0xfc, 0xee, 0x0c, 0x3b, 0xa1, 0xb6, 0x26, 0x5d, 0xf2, 0x04, 0x36, 0x42,
	0xe3, 0x39, 0xa7, 0xb4, 0xef, 0x8c, 0xa3, 0xa4, 0xc7, 0xad, 0xa4, 0xc7, 0x2f, 0x49, 0xd2, 0x19,
	0x91, 0xd8, 0x67, 0xd0, 0x45, 0x63, 0xb4, 0xe1, 0x77, 0x89, 0xfd, 0xe8, 0x06, 0x8f, 0x0c, 0xf4,
	0x2c, 0x7a, 0x05, 0x89, 0x37, 0xba, 0x96, 0xc5, 0x65, 0x94, 0xd4, 0x51, 0x94, 0x78, 0x84, 0x48,
	0x4c, 0x0f, 0x61, 0x3b, 0x11, 0x2a, 0xa3, 0x7d, 0xc3, 0xef, 0x45, 0xe5, 0x46, 0xec, 0xab, 0x00,
	0xb1, 0x1f, 0x60, 0xbb, 0xc4, 0xc6, 0x60, 0x41, 0x12, 0xb4, 0xfc, 0x98, 0xea, 0x72, 0x7a, 0x7d,
	0x26, 0xcf, 0xaf, 0xbc, 0xb2, 0x95, 0x10, 0x8c, 0xc1, 0x86, 0x13, 0x95, 0xe5, 0xf7, 0xa9, 0x32,
	0xf4, 0x3d, 0x7a, 0x0d, 0x83, 0x25, 0x07, 0xc6, 0x61, 0x2b, 0x8d, 0x4f, 0x1a, 0xe4, 0xd6, 0x64,
	0x7b, 0xd0, 0xf1, 0xa6, 0x4e, 0xa3, 0x1c, 0x3e, 0xd9, 0x11, 0xf4, 0x6a, 0x1d, 0xfd, 0xd2, 0x14,
	0x2f, 0xec, 0x51, 0x05, 0xbb, 0xef, 0xb4, 0xfd, 0x1f, 0xeb, 0xe1, 0xf9, 0x52, 0xaf, 0xd6, 0xe9,
	0x71, 0x8f, 0x6f, 0xd0, 0xf4, 0xa4, 0xa1, 0xb6, 0xad, 0x23, 0x84, 0xad, 0xf6, 0x82, 0xf0, 0xbc,
	0xb0, 0x06, 0xe2, 0x15, 0xf4, 0x1d, 0x30, 0x6a, 0x41, 0x4c, 0x9b, 0xbe, 0xc3, 0x1b, 0xdb, 0xbd,
	0x12, 0xd3, 0x6e, 0xcd, 0x70, 0x62, 0xcf, 0x65, 0xd3, 0x60, 0x49, 0x1b, 0xa7, 0x97, 0xb5, 0xe6,
	0xe8, 0xaf, 0x0e, 0xf4, 0x5a, 0xc1, 0xdd, 0xf8, 0xa2, 0xf8, 0xe2, 0xce, 0xe2, 0xc5, 0xa7, 0xd0,
	0x15, 0x33, 0x87, 0x86, 0x82, 0xbf, 0x47, 0x83, 0x91, 0xc5, 0x26, 0xb0, 0x39, 0xc5, 0x99, 0x36,
	0x71, 0xb1, 0xbd, 0x87, 0x9f, 0x68, 0xa1, 0x21, 0xa5, 0x37, 0xb1, 0x21, 0x69, 0xd7, 0xb5, 0x36,
	0x3b, 0x80, 0x6e, 0x89, 0xb5, 0x13, 0x69, 0xbf, 0x45, 0x83, 0xfd, 0x1f, 0x6e, 0x15, 0x5a, 0x9f,
	0x4f, 0xb5, 0x3e, 0x8f, 0x52, 0x8d, 0x7b, 0x6d, 0xbb, 0x05, 0x49, 0xac, 0x8f, 0x61, 0x6f, 0x41,
	0x6a, 0x0b, 0x17, 0x57, 0xdb, 0x6e, 0x8b, 0xb7, 0x1b, 0xf9, 0x03, 0xd8, 0x91, 0x95, 0xd2, 0x06,
	0xf3, 0x99, 0x90, 0xb5, 0x37, 0xc8, 0x81, 0xea, 0x78, 0x2b, 0xa2, 0x2f, 0x22, 0xb8, 0xb4, 0xe7,
	0x06, 0x2b, 0x7b, 0xee, 0x01, 0x0c, 0x0c, 0x16, 0xb2, 0x49, 0xfb, 0x7a, 0x3b, 0xce, 0x4d, 0x84,
	0x28, 0x95, 0x21, 0x0c, 0x0a, 0xad, 0x4a, 0x19, 0x9e, 0x24, 0xea, 0xb4, 0x3e, 0x97, 0xa1, 0x10,
	0xda, 0xa0, 0xf5, 0xb5, 0x6b, 0x57, 0x68, 0xb4, 0x46, 0x7f, 0xac, 0x41, 0x97, 0x66, 0x34, 0x54,
	0xa2, 0xa8, 0x85, 0xb5, 0xa9, 0x7d, 0xd1, 0x58, 0x16, 0xfe, 0xfa, 0xaa, 0xf0, 0x8f, 0xa1, 0x3f,
	0x15, 0xc5, 0xb9, 0x33, 0xa2, 0x08, 0x0b, 0x3b, 0x8c, 0xce, 0x15, 0xc0, 0xbe, 0x87, 0x41, 0x89,
	0xb6, 0x30, 0xb2, 0xa1, 0xb2, 0xef, 0x53, 0xa7, 0x6e, 0x34, 0xa5, 0x0b, 0xa7, 0x6c, 0x39, 0xc2,
	0xe8, 0xa7, 0x30, 0x90, 0x0b, 0x33, 0x64, 0xeb, 0xa4, 0xab, 0x5b, 0xb1, 0x45, 0x83, 0x3d, 0x85,
	0x9e, 0xc5, 0x22, 0x2e, 0x86, 0x38, 0x3b, 0xff, 0x2a, 0x8e, 0x05, 0xf1, 0x8b, 0x67, 0x3f, 0x7f,
	0x5e, 0x49, 0xf7, 0xc6, 0x4f, 0xc7, 0x85, 0x9e, 0xd3, 0x1f, 0x80, 0x89, 0x38, 0x9b, 0xfc, 0xb7,
	0xff, 0x0d, 0xd3, 0x4d, 0x8a, 0xff, 0xf4, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x2e, 0x86,
	0x1e, 0x70, 0x08, 0x00, 0x00,
}

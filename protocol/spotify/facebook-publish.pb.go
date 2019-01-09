// Code generated by protoc-gen-go. DO NOT EDIT.
// source: facebook-publish.proto

package spotify

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EventReply struct {
	Queued               *int32     `protobuf:"varint,1,opt,name=queued" json:"queued,omitempty"`
	Retry                *RetryInfo `protobuf:"bytes,2,opt,name=retry" json:"retry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *EventReply) Reset()         { *m = EventReply{} }
func (m *EventReply) String() string { return proto.CompactTextString(m) }
func (*EventReply) ProtoMessage()    {}
func (*EventReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_facebook_publish_cf0009184a9a4cf7, []int{0}
}
func (m *EventReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventReply.Unmarshal(m, b)
}
func (m *EventReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventReply.Marshal(b, m, deterministic)
}
func (dst *EventReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventReply.Merge(dst, src)
}
func (m *EventReply) XXX_Size() int {
	return xxx_messageInfo_EventReply.Size(m)
}
func (m *EventReply) XXX_DiscardUnknown() {
	xxx_messageInfo_EventReply.DiscardUnknown(m)
}

var xxx_messageInfo_EventReply proto.InternalMessageInfo

func (m *EventReply) GetQueued() int32 {
	if m != nil && m.Queued != nil {
		return *m.Queued
	}
	return 0
}

func (m *EventReply) GetRetry() *RetryInfo {
	if m != nil {
		return m.Retry
	}
	return nil
}

type RetryInfo struct {
	RetryDelay           *int32   `protobuf:"varint,1,opt,name=retry_delay,json=retryDelay" json:"retry_delay,omitempty"`
	MaxRetry             *int32   `protobuf:"varint,2,opt,name=max_retry,json=maxRetry" json:"max_retry,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetryInfo) Reset()         { *m = RetryInfo{} }
func (m *RetryInfo) String() string { return proto.CompactTextString(m) }
func (*RetryInfo) ProtoMessage()    {}
func (*RetryInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_facebook_publish_cf0009184a9a4cf7, []int{1}
}
func (m *RetryInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetryInfo.Unmarshal(m, b)
}
func (m *RetryInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetryInfo.Marshal(b, m, deterministic)
}
func (dst *RetryInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetryInfo.Merge(dst, src)
}
func (m *RetryInfo) XXX_Size() int {
	return xxx_messageInfo_RetryInfo.Size(m)
}
func (m *RetryInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RetryInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RetryInfo proto.InternalMessageInfo

func (m *RetryInfo) GetRetryDelay() int32 {
	if m != nil && m.RetryDelay != nil {
		return *m.RetryDelay
	}
	return 0
}

func (m *RetryInfo) GetMaxRetry() int32 {
	if m != nil && m.MaxRetry != nil {
		return *m.MaxRetry
	}
	return 0
}

type Id struct {
	Uri                  *string  `protobuf:"bytes,1,opt,name=uri" json:"uri,omitempty"`
	StartTime            *int64   `protobuf:"varint,2,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Id) Reset()         { *m = Id{} }
func (m *Id) String() string { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()    {}
func (*Id) Descriptor() ([]byte, []int) {
	return fileDescriptor_facebook_publish_cf0009184a9a4cf7, []int{2}
}
func (m *Id) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Id.Unmarshal(m, b)
}
func (m *Id) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Id.Marshal(b, m, deterministic)
}
func (dst *Id) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Id.Merge(dst, src)
}
func (m *Id) XXX_Size() int {
	return xxx_messageInfo_Id.Size(m)
}
func (m *Id) XXX_DiscardUnknown() {
	xxx_messageInfo_Id.DiscardUnknown(m)
}

var xxx_messageInfo_Id proto.InternalMessageInfo

func (m *Id) GetUri() string {
	if m != nil && m.Uri != nil {
		return *m.Uri
	}
	return ""
}

func (m *Id) GetStartTime() int64 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return 0
}

type Start struct {
	Length               *int32   `protobuf:"varint,1,opt,name=length" json:"length,omitempty"`
	ContextUri           *string  `protobuf:"bytes,2,opt,name=context_uri,json=contextUri" json:"context_uri,omitempty"`
	EndTime              *int64   `protobuf:"varint,3,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Start) Reset()         { *m = Start{} }
func (m *Start) String() string { return proto.CompactTextString(m) }
func (*Start) ProtoMessage()    {}
func (*Start) Descriptor() ([]byte, []int) {
	return fileDescriptor_facebook_publish_cf0009184a9a4cf7, []int{3}
}
func (m *Start) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Start.Unmarshal(m, b)
}
func (m *Start) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Start.Marshal(b, m, deterministic)
}
func (dst *Start) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Start.Merge(dst, src)
}
func (m *Start) XXX_Size() int {
	return xxx_messageInfo_Start.Size(m)
}
func (m *Start) XXX_DiscardUnknown() {
	xxx_messageInfo_Start.DiscardUnknown(m)
}

var xxx_messageInfo_Start proto.InternalMessageInfo

func (m *Start) GetLength() int32 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

func (m *Start) GetContextUri() string {
	if m != nil && m.ContextUri != nil {
		return *m.ContextUri
	}
	return ""
}

func (m *Start) GetEndTime() int64 {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return 0
}

type Seek struct {
	EndTime              *int64   `protobuf:"varint,1,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Seek) Reset()         { *m = Seek{} }
func (m *Seek) String() string { return proto.CompactTextString(m) }
func (*Seek) ProtoMessage()    {}
func (*Seek) Descriptor() ([]byte, []int) {
	return fileDescriptor_facebook_publish_cf0009184a9a4cf7, []int{4}
}
func (m *Seek) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Seek.Unmarshal(m, b)
}
func (m *Seek) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Seek.Marshal(b, m, deterministic)
}
func (dst *Seek) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Seek.Merge(dst, src)
}
func (m *Seek) XXX_Size() int {
	return xxx_messageInfo_Seek.Size(m)
}
func (m *Seek) XXX_DiscardUnknown() {
	xxx_messageInfo_Seek.DiscardUnknown(m)
}

var xxx_messageInfo_Seek proto.InternalMessageInfo

func (m *Seek) GetEndTime() int64 {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return 0
}

type Pause struct {
	SecondsPlayed        *int32   `protobuf:"varint,1,opt,name=seconds_played,json=secondsPlayed" json:"seconds_played,omitempty"`
	EndTime              *int64   `protobuf:"varint,2,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pause) Reset()         { *m = Pause{} }
func (m *Pause) String() string { return proto.CompactTextString(m) }
func (*Pause) ProtoMessage()    {}
func (*Pause) Descriptor() ([]byte, []int) {
	return fileDescriptor_facebook_publish_cf0009184a9a4cf7, []int{5}
}
func (m *Pause) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pause.Unmarshal(m, b)
}
func (m *Pause) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pause.Marshal(b, m, deterministic)
}
func (dst *Pause) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pause.Merge(dst, src)
}
func (m *Pause) XXX_Size() int {
	return xxx_messageInfo_Pause.Size(m)
}
func (m *Pause) XXX_DiscardUnknown() {
	xxx_messageInfo_Pause.DiscardUnknown(m)
}

var xxx_messageInfo_Pause proto.InternalMessageInfo

func (m *Pause) GetSecondsPlayed() int32 {
	if m != nil && m.SecondsPlayed != nil {
		return *m.SecondsPlayed
	}
	return 0
}

func (m *Pause) GetEndTime() int64 {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return 0
}

type Resume struct {
	SecondsPlayed        *int32   `protobuf:"varint,1,opt,name=seconds_played,json=secondsPlayed" json:"seconds_played,omitempty"`
	EndTime              *int64   `protobuf:"varint,2,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resume) Reset()         { *m = Resume{} }
func (m *Resume) String() string { return proto.CompactTextString(m) }
func (*Resume) ProtoMessage()    {}
func (*Resume) Descriptor() ([]byte, []int) {
	return fileDescriptor_facebook_publish_cf0009184a9a4cf7, []int{6}
}
func (m *Resume) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resume.Unmarshal(m, b)
}
func (m *Resume) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resume.Marshal(b, m, deterministic)
}
func (dst *Resume) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resume.Merge(dst, src)
}
func (m *Resume) XXX_Size() int {
	return xxx_messageInfo_Resume.Size(m)
}
func (m *Resume) XXX_DiscardUnknown() {
	xxx_messageInfo_Resume.DiscardUnknown(m)
}

var xxx_messageInfo_Resume proto.InternalMessageInfo

func (m *Resume) GetSecondsPlayed() int32 {
	if m != nil && m.SecondsPlayed != nil {
		return *m.SecondsPlayed
	}
	return 0
}

func (m *Resume) GetEndTime() int64 {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return 0
}

type End struct {
	SecondsPlayed        *int32   `protobuf:"varint,1,opt,name=seconds_played,json=secondsPlayed" json:"seconds_played,omitempty"`
	EndTime              *int64   `protobuf:"varint,2,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *End) Reset()         { *m = End{} }
func (m *End) String() string { return proto.CompactTextString(m) }
func (*End) ProtoMessage()    {}
func (*End) Descriptor() ([]byte, []int) {
	return fileDescriptor_facebook_publish_cf0009184a9a4cf7, []int{7}
}
func (m *End) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_End.Unmarshal(m, b)
}
func (m *End) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_End.Marshal(b, m, deterministic)
}
func (dst *End) XXX_Merge(src proto.Message) {
	xxx_messageInfo_End.Merge(dst, src)
}
func (m *End) XXX_Size() int {
	return xxx_messageInfo_End.Size(m)
}
func (m *End) XXX_DiscardUnknown() {
	xxx_messageInfo_End.DiscardUnknown(m)
}

var xxx_messageInfo_End proto.InternalMessageInfo

func (m *End) GetSecondsPlayed() int32 {
	if m != nil && m.SecondsPlayed != nil {
		return *m.SecondsPlayed
	}
	return 0
}

func (m *End) GetEndTime() int64 {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return 0
}

type Event struct {
	Id                   *Id      `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Start                *Start   `protobuf:"bytes,2,opt,name=start" json:"start,omitempty"`
	Seek                 *Seek    `protobuf:"bytes,3,opt,name=seek" json:"seek,omitempty"`
	Pause                *Pause   `protobuf:"bytes,4,opt,name=pause" json:"pause,omitempty"`
	Resume               *Resume  `protobuf:"bytes,5,opt,name=resume" json:"resume,omitempty"`
	End                  *End     `protobuf:"bytes,6,opt,name=end" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_facebook_publish_cf0009184a9a4cf7, []int{8}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (dst *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(dst, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() *Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Event) GetStart() *Start {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *Event) GetSeek() *Seek {
	if m != nil {
		return m.Seek
	}
	return nil
}

func (m *Event) GetPause() *Pause {
	if m != nil {
		return m.Pause
	}
	return nil
}

func (m *Event) GetResume() *Resume {
	if m != nil {
		return m.Resume
	}
	return nil
}

func (m *Event) GetEnd() *End {
	if m != nil {
		return m.End
	}
	return nil
}

func init() {
	proto.RegisterType((*EventReply)(nil), "spotify.EventReply")
	proto.RegisterType((*RetryInfo)(nil), "spotify.RetryInfo")
	proto.RegisterType((*Id)(nil), "spotify.Id")
	proto.RegisterType((*Start)(nil), "spotify.Start")
	proto.RegisterType((*Seek)(nil), "spotify.Seek")
	proto.RegisterType((*Pause)(nil), "spotify.Pause")
	proto.RegisterType((*Resume)(nil), "spotify.Resume")
	proto.RegisterType((*End)(nil), "spotify.End")
	proto.RegisterType((*Event)(nil), "spotify.Event")
}

func init() {
	proto.RegisterFile("facebook-publish.proto", fileDescriptor_facebook_publish_cf0009184a9a4cf7)
}

var fileDescriptor_facebook_publish_cf0009184a9a4cf7 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x41, 0x6f, 0xd4, 0x30,
	0x10, 0x85, 0x95, 0x64, 0xbd, 0xed, 0x4e, 0x68, 0x41, 0x3e, 0x54, 0x41, 0x15, 0xd0, 0x46, 0x20,
	0xf6, 0xc2, 0x1e, 0x56, 0xe2, 0x1f, 0xb0, 0x42, 0xe1, 0x80, 0x2a, 0x17, 0x4e, 0x1c, 0xa2, 0x74,
	0x3d, 0x4b, 0xad, 0x24, 0x76, 0x88, 0x1d, 0xb4, 0xf9, 0xaf, 0xfc, 0x18, 0xe4, 0x49, 0x48, 0x49,
	0xcf, 0xbd, 0xc5, 0xef, 0xbd, 0xf9, 0x14, 0x3f, 0x0f, 0x5c, 0x1c, 0x8a, 0x3d, 0xde, 0x19, 0x53,
	0x7e, 0x68, 0xba, 0xbb, 0x4a, 0xd9, 0xfb, 0x4d, 0xd3, 0x1a, 0x67, 0xf8, 0x89, 0x6d, 0x8c, 0x53,
	0x87, 0x3e, 0xfd, 0x0a, 0xb0, 0xfb, 0x8d, 0xda, 0x09, 0x6c, 0xaa, 0x9e, 0x5f, 0xc0, 0xf2, 0x57,
	0x87, 0x1d, 0xca, 0x24, 0xb8, 0x0a, 0xd6, 0x4c, 0x8c, 0x27, 0xbe, 0x06, 0xd6, 0xa2, 0x6b, 0xfb,
	0x24, 0xbc, 0x0a, 0xd6, 0xf1, 0x96, 0x6f, 0xc6, 0xf1, 0x8d, 0xf0, 0x6a, 0xa6, 0x0f, 0x46, 0x0c,
	0x81, 0x34, 0x83, 0xd5, 0xa4, 0xf1, 0x37, 0x10, 0x93, 0x9a, 0x4b, 0xac, 0x8a, 0x7e, 0x64, 0x02,
	0x49, 0x9f, 0xbc, 0xc2, 0x2f, 0x61, 0x55, 0x17, 0xc7, 0xfc, 0x81, 0xcd, 0xc4, 0x69, 0x5d, 0x1c,
	0x89, 0x90, 0x7e, 0x84, 0x30, 0x93, 0xfc, 0x05, 0x44, 0x5d, 0xab, 0x68, 0x76, 0x25, 0xfc, 0x27,
	0x7f, 0x05, 0x60, 0x5d, 0xd1, 0xba, 0xdc, 0xa9, 0x1a, 0x69, 0x2a, 0x12, 0x2b, 0x52, 0xbe, 0xa9,
	0x1a, 0xd3, 0x1f, 0xc0, 0x6e, 0xfd, 0xc1, 0x5f, 0xa6, 0x42, 0xfd, 0xd3, 0xdd, 0xff, 0xbb, 0xcc,
	0x70, 0xf2, 0x7f, 0xb5, 0x37, 0xda, 0xe1, 0xd1, 0xe5, 0x9e, 0x1c, 0x12, 0x19, 0x46, 0xe9, 0x7b,
	0xab, 0xf8, 0x4b, 0x38, 0x45, 0x2d, 0x07, 0x7c, 0x44, 0xf8, 0x13, 0xd4, 0x92, 0xe0, 0xd7, 0xb0,
	0xb8, 0x45, 0x2c, 0x67, 0x91, 0x60, 0x1e, 0xc9, 0x80, 0xdd, 0x14, 0x9d, 0x45, 0xfe, 0x0e, 0xce,
	0x2d, 0xee, 0x8d, 0x96, 0x36, 0x6f, 0xaa, 0xa2, 0x9f, 0x4a, 0x3d, 0x1b, 0xd5, 0x1b, 0x12, 0x67,
	0xa8, 0x70, 0x8e, 0xfa, 0x02, 0x4b, 0x81, 0xb6, 0xab, 0x9f, 0x82, 0xf5, 0x19, 0xa2, 0x9d, 0x96,
	0x4f, 0x00, 0xfa, 0x13, 0x00, 0xa3, 0x95, 0xe1, 0x97, 0x10, 0xaa, 0x61, 0x3e, 0xde, 0xc6, 0xd3,
	0x4a, 0x64, 0x52, 0x84, 0x4a, 0xf2, 0xb7, 0xc0, 0xe8, 0x4d, 0xc6, 0x95, 0x39, 0x9f, 0x7c, 0x7a,
	0x1c, 0x31, 0x98, 0xfc, 0x1a, 0x16, 0x16, 0xb1, 0xa4, 0x9a, 0xe3, 0xed, 0xd9, 0x43, 0x08, 0xb1,
	0x14, 0x64, 0x79, 0x50, 0xe3, 0xfb, 0x4c, 0x16, 0x8f, 0x40, 0xd4, 0xb2, 0x18, 0x4c, 0xfe, 0x1e,
	0x96, 0x2d, 0x55, 0x95, 0x30, 0x8a, 0x3d, 0xff, 0x6f, 0x45, 0xbd, 0x2c, 0x46, 0x9b, 0xbf, 0x86,
	0x08, 0xb5, 0x4c, 0x96, 0x94, 0x7a, 0x36, 0xa5, 0x76, 0x5a, 0x0a, 0x6f, 0xfc, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0x4e, 0x5f, 0xd7, 0xda, 0x32, 0x03, 0x00, 0x00,
}
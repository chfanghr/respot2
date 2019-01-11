// Code generated by protoc-gen-go. DO NOT EDIT.
// source: playlist4content.proto

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

type Item struct {
	Uri                  *string         `protobuf:"bytes,1,opt,name=uri" json:"uri,omitempty"`
	Attributes           *ItemAttributes `protobuf:"bytes,2,opt,name=attributes" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_playlist4content_195b69ac1b08267a, []int{0}
}
func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (dst *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(dst, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetUri() string {
	if m != nil && m.Uri != nil {
		return *m.Uri
	}
	return ""
}

func (m *Item) GetAttributes() *ItemAttributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type ListItems struct {
	Pos                  *int32   `protobuf:"varint,1,opt,name=pos" json:"pos,omitempty"`
	Truncated            *bool    `protobuf:"varint,2,opt,name=truncated" json:"truncated,omitempty"`
	Items                []*Item  `protobuf:"bytes,3,rep,name=items" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListItems) Reset()         { *m = ListItems{} }
func (m *ListItems) String() string { return proto.CompactTextString(m) }
func (*ListItems) ProtoMessage()    {}
func (*ListItems) Descriptor() ([]byte, []int) {
	return fileDescriptor_playlist4content_195b69ac1b08267a, []int{1}
}
func (m *ListItems) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListItems.Unmarshal(m, b)
}
func (m *ListItems) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListItems.Marshal(b, m, deterministic)
}
func (dst *ListItems) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListItems.Merge(dst, src)
}
func (m *ListItems) XXX_Size() int {
	return xxx_messageInfo_ListItems.Size(m)
}
func (m *ListItems) XXX_DiscardUnknown() {
	xxx_messageInfo_ListItems.DiscardUnknown(m)
}

var xxx_messageInfo_ListItems proto.InternalMessageInfo

func (m *ListItems) GetPos() int32 {
	if m != nil && m.Pos != nil {
		return *m.Pos
	}
	return 0
}

func (m *ListItems) GetTruncated() bool {
	if m != nil && m.Truncated != nil {
		return *m.Truncated
	}
	return false
}

func (m *ListItems) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

type ContentRange struct {
	Pos                  *int32   `protobuf:"varint,1,opt,name=pos" json:"pos,omitempty"`
	Length               *int32   `protobuf:"varint,2,opt,name=length" json:"length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContentRange) Reset()         { *m = ContentRange{} }
func (m *ContentRange) String() string { return proto.CompactTextString(m) }
func (*ContentRange) ProtoMessage()    {}
func (*ContentRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_playlist4content_195b69ac1b08267a, []int{2}
}
func (m *ContentRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentRange.Unmarshal(m, b)
}
func (m *ContentRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentRange.Marshal(b, m, deterministic)
}
func (dst *ContentRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentRange.Merge(dst, src)
}
func (m *ContentRange) XXX_Size() int {
	return xxx_messageInfo_ContentRange.Size(m)
}
func (m *ContentRange) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentRange.DiscardUnknown(m)
}

var xxx_messageInfo_ContentRange proto.InternalMessageInfo

func (m *ContentRange) GetPos() int32 {
	if m != nil && m.Pos != nil {
		return *m.Pos
	}
	return 0
}

func (m *ContentRange) GetLength() int32 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

type ListContentSelection struct {
	WantRevision          *bool                  `protobuf:"varint,1,opt,name=wantRevision" json:"wantRevision,omitempty"`
	WantLength            *bool                  `protobuf:"varint,2,opt,name=wantLength" json:"wantLength,omitempty"`
	WantAttributes        *bool                  `protobuf:"varint,3,opt,name=wantAttributes" json:"wantAttributes,omitempty"`
	WantChecksum          *bool                  `protobuf:"varint,4,opt,name=wantChecksum" json:"wantChecksum,omitempty"`
	WantContent           *bool                  `protobuf:"varint,5,opt,name=wantContent" json:"wantContent,omitempty"`
	ContentRange          *ContentRange          `protobuf:"bytes,6,opt,name=contentRange" json:"contentRange,omitempty"`
	WantDiff              *bool                  `protobuf:"varint,7,opt,name=wantDiff" json:"wantDiff,omitempty"`
	BaseRevision          []byte                 `protobuf:"bytes,8,opt,name=baseRevision" json:"baseRevision,omitempty"`
	HintRevision          []byte                 `protobuf:"bytes,9,opt,name=hintRevision" json:"hintRevision,omitempty"`
	WantNothingIfUpToDate *bool                  `protobuf:"varint,10,opt,name=wantNothingIfUpToDate" json:"wantNothingIfUpToDate,omitempty"`
	WantResolveAction     *bool                  `protobuf:"varint,12,opt,name=wantResolveAction" json:"wantResolveAction,omitempty"`
	Issues                []*ClientIssue         `protobuf:"bytes,13,rep,name=issues" json:"issues,omitempty"`
	ResolveAction         []*ClientResolveAction `protobuf:"bytes,14,rep,name=resolveAction" json:"resolveAction,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}               `json:"-"`
	XXX_unrecognized      []byte                 `json:"-"`
	XXX_sizecache         int32                  `json:"-"`
}

func (m *ListContentSelection) Reset()         { *m = ListContentSelection{} }
func (m *ListContentSelection) String() string { return proto.CompactTextString(m) }
func (*ListContentSelection) ProtoMessage()    {}
func (*ListContentSelection) Descriptor() ([]byte, []int) {
	return fileDescriptor_playlist4content_195b69ac1b08267a, []int{3}
}
func (m *ListContentSelection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListContentSelection.Unmarshal(m, b)
}
func (m *ListContentSelection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListContentSelection.Marshal(b, m, deterministic)
}
func (dst *ListContentSelection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListContentSelection.Merge(dst, src)
}
func (m *ListContentSelection) XXX_Size() int {
	return xxx_messageInfo_ListContentSelection.Size(m)
}
func (m *ListContentSelection) XXX_DiscardUnknown() {
	xxx_messageInfo_ListContentSelection.DiscardUnknown(m)
}

var xxx_messageInfo_ListContentSelection proto.InternalMessageInfo

func (m *ListContentSelection) GetWantRevision() bool {
	if m != nil && m.WantRevision != nil {
		return *m.WantRevision
	}
	return false
}

func (m *ListContentSelection) GetWantLength() bool {
	if m != nil && m.WantLength != nil {
		return *m.WantLength
	}
	return false
}

func (m *ListContentSelection) GetWantAttributes() bool {
	if m != nil && m.WantAttributes != nil {
		return *m.WantAttributes
	}
	return false
}

func (m *ListContentSelection) GetWantChecksum() bool {
	if m != nil && m.WantChecksum != nil {
		return *m.WantChecksum
	}
	return false
}

func (m *ListContentSelection) GetWantContent() bool {
	if m != nil && m.WantContent != nil {
		return *m.WantContent
	}
	return false
}

func (m *ListContentSelection) GetContentRange() *ContentRange {
	if m != nil {
		return m.ContentRange
	}
	return nil
}

func (m *ListContentSelection) GetWantDiff() bool {
	if m != nil && m.WantDiff != nil {
		return *m.WantDiff
	}
	return false
}

func (m *ListContentSelection) GetBaseRevision() []byte {
	if m != nil {
		return m.BaseRevision
	}
	return nil
}

func (m *ListContentSelection) GetHintRevision() []byte {
	if m != nil {
		return m.HintRevision
	}
	return nil
}

func (m *ListContentSelection) GetWantNothingIfUpToDate() bool {
	if m != nil && m.WantNothingIfUpToDate != nil {
		return *m.WantNothingIfUpToDate
	}
	return false
}

func (m *ListContentSelection) GetWantResolveAction() bool {
	if m != nil && m.WantResolveAction != nil {
		return *m.WantResolveAction
	}
	return false
}

func (m *ListContentSelection) GetIssues() []*ClientIssue {
	if m != nil {
		return m.Issues
	}
	return nil
}

func (m *ListContentSelection) GetResolveAction() []*ClientResolveAction {
	if m != nil {
		return m.ResolveAction
	}
	return nil
}

func init() {
	proto.RegisterType((*Item)(nil), "spotify.Item")
	proto.RegisterType((*ListItems)(nil), "spotify.ListItems")
	proto.RegisterType((*ContentRange)(nil), "spotify.ContentRange")
	proto.RegisterType((*ListContentSelection)(nil), "spotify.ListContentSelection")
}

func init() {
	proto.RegisterFile("playlist4content.proto", fileDescriptor_playlist4content_195b69ac1b08267a)
}

var fileDescriptor_playlist4content_195b69ac1b08267a = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x5b, 0x6f, 0xd3, 0x30,
	0x14, 0x56, 0xe9, 0x65, 0xe9, 0x69, 0x3a, 0x81, 0x59, 0x87, 0x55, 0x4d, 0x28, 0x0a, 0x12, 0xea,
	0xc3, 0xd4, 0x87, 0x69, 0x12, 0xf0, 0x38, 0xb6, 0x97, 0x4a, 0x13, 0x12, 0x06, 0xde, 0xf1, 0xc2,
	0x69, 0x6b, 0x91, 0xda, 0x51, 0x7c, 0x32, 0xb4, 0x1f, 0xc7, 0x7f, 0x43, 0xb6, 0xb3, 0x5c, 0xc6,
	0xde, 0x72, 0xbe, 0xf3, 0x5d, 0x9c, 0xcf, 0x86, 0xd3, 0x22, 0x97, 0x0f, 0xb9, 0xb2, 0x74, 0x99,
	0x19, 0x4d, 0xa8, 0x69, 0x5d, 0x94, 0x86, 0x0c, 0x3b, 0xb2, 0x85, 0x21, 0xb5, 0x7d, 0x58, 0xbe,
	0x6e, 0x08, 0x07, 0x24, 0x19, 0xb6, 0xcb, 0x45, 0x03, 0x2a, 0x6b, 0x2b, 0xb4, 0x01, 0x4e, 0xbf,
	0xc2, 0x68, 0x43, 0x78, 0x60, 0x2f, 0x61, 0x58, 0x95, 0x8a, 0x0f, 0x92, 0xc1, 0x6a, 0x2a, 0xdc,
	0x27, 0xfb, 0x00, 0x20, 0x89, 0x4a, 0x75, 0x57, 0x11, 0x5a, 0xfe, 0x22, 0x19, 0xac, 0x66, 0x17,
	0x6f, 0xd6, 0x75, 0xc6, 0xda, 0x89, 0xae, 0x9a, 0xb5, 0xe8, 0x50, 0xd3, 0x9f, 0x30, 0xbd, 0x55,
	0x96, 0x1c, 0xc3, 0x3a, 0xdf, 0xc2, 0x58, 0xef, 0x3b, 0x16, 0xee, 0x93, 0x9d, 0xc1, 0x94, 0xca,
	0x4a, 0x67, 0x92, 0xf0, 0x97, 0xb7, 0x8d, 0x44, 0x0b, 0xb0, 0x77, 0x30, 0x56, 0x4e, 0xc8, 0x87,
	0xc9, 0x70, 0x35, 0xbb, 0x98, 0xf7, 0x02, 0x45, 0xd8, 0xa5, 0x1f, 0x21, 0xbe, 0x0e, 0xbf, 0x2e,
	0xa4, 0xde, 0xe1, 0x33, 0x21, 0xa7, 0x30, 0xc9, 0x51, 0xef, 0x68, 0xef, 0x13, 0xc6, 0xa2, 0x9e,
	0xd2, 0xbf, 0x23, 0x38, 0x71, 0x87, 0xab, 0xe5, 0xdf, 0x30, 0xc7, 0x8c, 0x94, 0xd1, 0x2c, 0x85,
	0xf8, 0x8f, 0xd4, 0x24, 0xf0, 0x5e, 0x59, 0x65, 0xb4, 0xf7, 0x8a, 0x44, 0x0f, 0x63, 0x6f, 0x01,
	0xdc, 0x7c, 0xdb, 0x1a, 0x47, 0xa2, 0x83, 0xb0, 0xf7, 0x70, 0xec, 0xa6, 0xb6, 0x16, 0x3e, 0xf4,
	0x9c, 0x27, 0xe8, 0x63, 0xd6, 0xf5, 0x1e, 0xb3, 0xdf, 0xb6, 0x3a, 0xf0, 0x51, 0x9b, 0xf5, 0x88,
	0xb1, 0x04, 0x66, 0x7e, 0x0e, 0xe7, 0xe4, 0x63, 0x4f, 0xe9, 0x42, 0xec, 0x13, 0xc4, 0x59, 0xa7,
	0x04, 0x3e, 0xf1, 0x37, 0xb4, 0x68, 0x0a, 0xeb, 0x36, 0x24, 0x7a, 0x54, 0xb6, 0x84, 0xc8, 0x39,
	0xdd, 0xa8, 0xed, 0x96, 0x1f, 0x79, 0xe7, 0x66, 0x76, 0x87, 0xbb, 0x93, 0x16, 0x9b, 0x22, 0xa2,
	0x64, 0xb0, 0x8a, 0x45, 0x0f, 0x73, 0x9c, 0xbd, 0xea, 0x94, 0x35, 0x0d, 0x9c, 0x2e, 0xc6, 0x2e,
	0x61, 0xe1, 0x3c, 0xbf, 0x18, 0xda, 0x2b, 0xbd, 0xdb, 0x6c, 0x7f, 0x14, 0xdf, 0xcd, 0x8d, 0x24,
	0xe4, 0xe0, 0x03, 0x9f, 0x5f, 0xb2, 0x73, 0x78, 0x15, 0x2a, 0xb7, 0x26, 0xbf, 0xc7, 0x2b, 0x7f,
	0x37, 0x3c, 0xf6, 0x8a, 0xff, 0x17, 0xec, 0x1c, 0x26, 0xe1, 0x31, 0xf3, 0xb9, 0x7f, 0x2d, 0x27,
	0xed, 0xcf, 0xe7, 0x0a, 0x35, 0x6d, 0xdc, 0x52, 0xd4, 0x1c, 0xf6, 0x19, 0xe6, 0x65, 0xcf, 0xf7,
	0xd8, 0x8b, 0xce, 0x9e, 0x88, 0x7a, 0x11, 0xa2, 0x2f, 0xf9, 0x17, 0x00, 0x00, 0xff, 0xff, 0x2c,
	0x9c, 0xcd, 0x6f, 0x7c, 0x03, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: playlist4meta.proto

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

type DownloadFormat_Codec int32

const (
	DownloadFormat_CODEC_UNKNOWN  DownloadFormat_Codec = 0
	DownloadFormat_OGG_VORBIS     DownloadFormat_Codec = 1
	DownloadFormat_FLAC           DownloadFormat_Codec = 2
	DownloadFormat_MPEG_1_LAYER_3 DownloadFormat_Codec = 3
)

var DownloadFormat_Codec_name = map[int32]string{
	0: "CODEC_UNKNOWN",
	1: "OGG_VORBIS",
	2: "FLAC",
	3: "MPEG_1_LAYER_3",
}
var DownloadFormat_Codec_value = map[string]int32{
	"CODEC_UNKNOWN":  0,
	"OGG_VORBIS":     1,
	"FLAC":           2,
	"MPEG_1_LAYER_3": 3,
}

func (x DownloadFormat_Codec) Enum() *DownloadFormat_Codec {
	p := new(DownloadFormat_Codec)
	*p = x
	return p
}
func (x DownloadFormat_Codec) String() string {
	return proto.EnumName(DownloadFormat_Codec_name, int32(x))
}
func (x *DownloadFormat_Codec) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DownloadFormat_Codec_value, data, "DownloadFormat_Codec")
	if err != nil {
		return err
	}
	*x = DownloadFormat_Codec(value)
	return nil
}
func (DownloadFormat_Codec) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_playlist4meta_5e2ad56a040bc0d8, []int{1, 0}
}

type ListChecksum struct {
	Version              *int32   `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Sha1                 []byte   `protobuf:"bytes,4,opt,name=sha1" json:"sha1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListChecksum) Reset()         { *m = ListChecksum{} }
func (m *ListChecksum) String() string { return proto.CompactTextString(m) }
func (*ListChecksum) ProtoMessage()    {}
func (*ListChecksum) Descriptor() ([]byte, []int) {
	return fileDescriptor_playlist4meta_5e2ad56a040bc0d8, []int{0}
}
func (m *ListChecksum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListChecksum.Unmarshal(m, b)
}
func (m *ListChecksum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListChecksum.Marshal(b, m, deterministic)
}
func (dst *ListChecksum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListChecksum.Merge(dst, src)
}
func (m *ListChecksum) XXX_Size() int {
	return xxx_messageInfo_ListChecksum.Size(m)
}
func (m *ListChecksum) XXX_DiscardUnknown() {
	xxx_messageInfo_ListChecksum.DiscardUnknown(m)
}

var xxx_messageInfo_ListChecksum proto.InternalMessageInfo

func (m *ListChecksum) GetVersion() int32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *ListChecksum) GetSha1() []byte {
	if m != nil {
		return m.Sha1
	}
	return nil
}

type DownloadFormat struct {
	Codec                *DownloadFormat_Codec `protobuf:"varint,1,opt,name=codec,enum=spotify.DownloadFormat_Codec" json:"codec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DownloadFormat) Reset()         { *m = DownloadFormat{} }
func (m *DownloadFormat) String() string { return proto.CompactTextString(m) }
func (*DownloadFormat) ProtoMessage()    {}
func (*DownloadFormat) Descriptor() ([]byte, []int) {
	return fileDescriptor_playlist4meta_5e2ad56a040bc0d8, []int{1}
}
func (m *DownloadFormat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadFormat.Unmarshal(m, b)
}
func (m *DownloadFormat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadFormat.Marshal(b, m, deterministic)
}
func (dst *DownloadFormat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadFormat.Merge(dst, src)
}
func (m *DownloadFormat) XXX_Size() int {
	return xxx_messageInfo_DownloadFormat.Size(m)
}
func (m *DownloadFormat) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadFormat.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadFormat proto.InternalMessageInfo

func (m *DownloadFormat) GetCodec() DownloadFormat_Codec {
	if m != nil && m.Codec != nil {
		return *m.Codec
	}
	return DownloadFormat_CODEC_UNKNOWN
}

type ListAttributes struct {
	Name                    *string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description             *string  `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Picture                 []byte   `protobuf:"bytes,3,opt,name=picture" json:"picture,omitempty"`
	Collaborative           *bool    `protobuf:"varint,4,opt,name=collaborative" json:"collaborative,omitempty"`
	Pl3Version              *string  `protobuf:"bytes,5,opt,name=pl3_version,json=pl3Version" json:"pl3_version,omitempty"`
	DeletedByOwner          *bool    `protobuf:"varint,6,opt,name=deleted_by_owner,json=deletedByOwner" json:"deleted_by_owner,omitempty"`
	RestrictedCollaborative *bool    `protobuf:"varint,7,opt,name=restricted_collaborative,json=restrictedCollaborative" json:"restricted_collaborative,omitempty"`
	DeprecatedClientId      *int64   `protobuf:"varint,8,opt,name=deprecated_client_id,json=deprecatedClientId" json:"deprecated_client_id,omitempty"`
	PublicStarred           *bool    `protobuf:"varint,9,opt,name=public_starred,json=publicStarred" json:"public_starred,omitempty"`
	ClientId                *string  `protobuf:"bytes,10,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *ListAttributes) Reset()         { *m = ListAttributes{} }
func (m *ListAttributes) String() string { return proto.CompactTextString(m) }
func (*ListAttributes) ProtoMessage()    {}
func (*ListAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_playlist4meta_5e2ad56a040bc0d8, []int{2}
}
func (m *ListAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAttributes.Unmarshal(m, b)
}
func (m *ListAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAttributes.Marshal(b, m, deterministic)
}
func (dst *ListAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAttributes.Merge(dst, src)
}
func (m *ListAttributes) XXX_Size() int {
	return xxx_messageInfo_ListAttributes.Size(m)
}
func (m *ListAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_ListAttributes proto.InternalMessageInfo

func (m *ListAttributes) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ListAttributes) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *ListAttributes) GetPicture() []byte {
	if m != nil {
		return m.Picture
	}
	return nil
}

func (m *ListAttributes) GetCollaborative() bool {
	if m != nil && m.Collaborative != nil {
		return *m.Collaborative
	}
	return false
}

func (m *ListAttributes) GetPl3Version() string {
	if m != nil && m.Pl3Version != nil {
		return *m.Pl3Version
	}
	return ""
}

func (m *ListAttributes) GetDeletedByOwner() bool {
	if m != nil && m.DeletedByOwner != nil {
		return *m.DeletedByOwner
	}
	return false
}

func (m *ListAttributes) GetRestrictedCollaborative() bool {
	if m != nil && m.RestrictedCollaborative != nil {
		return *m.RestrictedCollaborative
	}
	return false
}

func (m *ListAttributes) GetDeprecatedClientId() int64 {
	if m != nil && m.DeprecatedClientId != nil {
		return *m.DeprecatedClientId
	}
	return 0
}

func (m *ListAttributes) GetPublicStarred() bool {
	if m != nil && m.PublicStarred != nil {
		return *m.PublicStarred
	}
	return false
}

func (m *ListAttributes) GetClientId() string {
	if m != nil && m.ClientId != nil {
		return *m.ClientId
	}
	return ""
}

type ItemAttributes struct {
	AddedBy              *string         `protobuf:"bytes,1,opt,name=added_by,json=addedBy" json:"added_by,omitempty"`
	Timestamp            *int64          `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Message              *string         `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	Seen                 *bool           `protobuf:"varint,4,opt,name=seen" json:"seen,omitempty"`
	DownloadCount        *int64          `protobuf:"varint,5,opt,name=download_count,json=downloadCount" json:"download_count,omitempty"`
	DownloadFormat       *DownloadFormat `protobuf:"bytes,6,opt,name=download_format,json=downloadFormat" json:"download_format,omitempty"`
	SevendigitalId       *string         `protobuf:"bytes,7,opt,name=sevendigital_id,json=sevendigitalId" json:"sevendigital_id,omitempty"`
	SevendigitalLeft     *int64          `protobuf:"varint,8,opt,name=sevendigital_left,json=sevendigitalLeft" json:"sevendigital_left,omitempty"`
	SeenAt               *int64          `protobuf:"varint,9,opt,name=seen_at,json=seenAt" json:"seen_at,omitempty"`
	Public               *bool           `protobuf:"varint,10,opt,name=public" json:"public,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ItemAttributes) Reset()         { *m = ItemAttributes{} }
func (m *ItemAttributes) String() string { return proto.CompactTextString(m) }
func (*ItemAttributes) ProtoMessage()    {}
func (*ItemAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_playlist4meta_5e2ad56a040bc0d8, []int{3}
}
func (m *ItemAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemAttributes.Unmarshal(m, b)
}
func (m *ItemAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemAttributes.Marshal(b, m, deterministic)
}
func (dst *ItemAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemAttributes.Merge(dst, src)
}
func (m *ItemAttributes) XXX_Size() int {
	return xxx_messageInfo_ItemAttributes.Size(m)
}
func (m *ItemAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_ItemAttributes proto.InternalMessageInfo

func (m *ItemAttributes) GetAddedBy() string {
	if m != nil && m.AddedBy != nil {
		return *m.AddedBy
	}
	return ""
}

func (m *ItemAttributes) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *ItemAttributes) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *ItemAttributes) GetSeen() bool {
	if m != nil && m.Seen != nil {
		return *m.Seen
	}
	return false
}

func (m *ItemAttributes) GetDownloadCount() int64 {
	if m != nil && m.DownloadCount != nil {
		return *m.DownloadCount
	}
	return 0
}

func (m *ItemAttributes) GetDownloadFormat() *DownloadFormat {
	if m != nil {
		return m.DownloadFormat
	}
	return nil
}

func (m *ItemAttributes) GetSevendigitalId() string {
	if m != nil && m.SevendigitalId != nil {
		return *m.SevendigitalId
	}
	return ""
}

func (m *ItemAttributes) GetSevendigitalLeft() int64 {
	if m != nil && m.SevendigitalLeft != nil {
		return *m.SevendigitalLeft
	}
	return 0
}

func (m *ItemAttributes) GetSeenAt() int64 {
	if m != nil && m.SeenAt != nil {
		return *m.SeenAt
	}
	return 0
}

func (m *ItemAttributes) GetPublic() bool {
	if m != nil && m.Public != nil {
		return *m.Public
	}
	return false
}

type StringAttribute struct {
	Key                  *string  `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value                *string  `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringAttribute) Reset()         { *m = StringAttribute{} }
func (m *StringAttribute) String() string { return proto.CompactTextString(m) }
func (*StringAttribute) ProtoMessage()    {}
func (*StringAttribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_playlist4meta_5e2ad56a040bc0d8, []int{4}
}
func (m *StringAttribute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringAttribute.Unmarshal(m, b)
}
func (m *StringAttribute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringAttribute.Marshal(b, m, deterministic)
}
func (dst *StringAttribute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringAttribute.Merge(dst, src)
}
func (m *StringAttribute) XXX_Size() int {
	return xxx_messageInfo_StringAttribute.Size(m)
}
func (m *StringAttribute) XXX_DiscardUnknown() {
	xxx_messageInfo_StringAttribute.DiscardUnknown(m)
}

var xxx_messageInfo_StringAttribute proto.InternalMessageInfo

func (m *StringAttribute) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *StringAttribute) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type StringAttributes struct {
	Attribute            []*StringAttribute `protobuf:"bytes,1,rep,name=attribute" json:"attribute,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *StringAttributes) Reset()         { *m = StringAttributes{} }
func (m *StringAttributes) String() string { return proto.CompactTextString(m) }
func (*StringAttributes) ProtoMessage()    {}
func (*StringAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_playlist4meta_5e2ad56a040bc0d8, []int{5}
}
func (m *StringAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringAttributes.Unmarshal(m, b)
}
func (m *StringAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringAttributes.Marshal(b, m, deterministic)
}
func (dst *StringAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringAttributes.Merge(dst, src)
}
func (m *StringAttributes) XXX_Size() int {
	return xxx_messageInfo_StringAttributes.Size(m)
}
func (m *StringAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_StringAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_StringAttributes proto.InternalMessageInfo

func (m *StringAttributes) GetAttribute() []*StringAttribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

func init() {
	proto.RegisterType((*ListChecksum)(nil), "spotify.ListChecksum")
	proto.RegisterType((*DownloadFormat)(nil), "spotify.DownloadFormat")
	proto.RegisterType((*ListAttributes)(nil), "spotify.ListAttributes")
	proto.RegisterType((*ItemAttributes)(nil), "spotify.ItemAttributes")
	proto.RegisterType((*StringAttribute)(nil), "spotify.StringAttribute")
	proto.RegisterType((*StringAttributes)(nil), "spotify.StringAttributes")
	proto.RegisterEnum("spotify.DownloadFormat_Codec", DownloadFormat_Codec_name, DownloadFormat_Codec_value)
}

func init() { proto.RegisterFile("playlist4meta.proto", fileDescriptor_playlist4meta_5e2ad56a040bc0d8) }

var fileDescriptor_playlist4meta_5e2ad56a040bc0d8 = []byte{
	// 637 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x6d, 0x6b, 0x13, 0x41,
	0x10, 0x36, 0xbd, 0xa6, 0x49, 0x26, 0xed, 0xf5, 0xba, 0x16, 0x7b, 0xa2, 0x62, 0x08, 0x8a, 0x01,
	0x21, 0xd8, 0x46, 0x84, 0x82, 0x1f, 0x4c, 0xaf, 0x2f, 0x46, 0x63, 0x23, 0x5b, 0xac, 0xf8, 0xe9,
	0xd8, 0xdc, 0x4e, 0xda, 0xa5, 0xf7, 0xc6, 0xee, 0x26, 0x25, 0x5f, 0xfd, 0x03, 0xfe, 0x4c, 0xff,
	0x86, 0xec, 0xe6, 0xf2, 0x56, 0xf0, 0xdb, 0x3e, 0xcf, 0x3c, 0x33, 0x3b, 0xf3, 0xcc, 0xc0, 0xe3,
	0x3c, 0x66, 0xd3, 0x58, 0x28, 0xfd, 0x3e, 0x41, 0xcd, 0xda, 0xb9, 0xcc, 0x74, 0x46, 0x2a, 0x2a,
	0xcf, 0xb4, 0x18, 0x4d, 0x9b, 0x1f, 0x61, 0xbb, 0x2f, 0x94, 0x0e, 0x6e, 0x31, 0xba, 0x53, 0xe3,
	0x84, 0xf8, 0x50, 0x99, 0xa0, 0x54, 0x22, 0x4b, 0xfd, 0x52, 0xa3, 0xd4, 0x2a, 0xd3, 0x39, 0x24,
	0x04, 0x36, 0xd5, 0x2d, 0x3b, 0xf4, 0x37, 0x1b, 0xa5, 0xd6, 0x36, 0xb5, 0xef, 0xe6, 0x9f, 0x12,
	0xb8, 0xa7, 0xd9, 0x7d, 0x1a, 0x67, 0x8c, 0x9f, 0x67, 0x32, 0x61, 0x9a, 0x74, 0xa0, 0x1c, 0x65,
	0x1c, 0x23, 0x9b, 0xee, 0x1e, 0xbd, 0x68, 0x17, 0x3f, 0xb5, 0xd7, 0x75, 0xed, 0xc0, 0x88, 0xe8,
	0x4c, 0xdb, 0xfc, 0x0c, 0x65, 0x8b, 0xc9, 0x1e, 0xec, 0x04, 0x83, 0xd3, 0xb3, 0x20, 0xfc, 0x71,
	0xf9, 0xf5, 0x72, 0xf0, 0xf3, 0xd2, 0x7b, 0x44, 0x5c, 0x80, 0xc1, 0xc5, 0x45, 0x78, 0x3d, 0xa0,
	0x27, 0xbd, 0x2b, 0xaf, 0x44, 0xaa, 0xb0, 0x79, 0xde, 0xef, 0x06, 0xde, 0x06, 0x21, 0xe0, 0x7e,
	0xfb, 0x7e, 0x76, 0x11, 0x1e, 0x86, 0xfd, 0xee, 0xaf, 0x33, 0x1a, 0x76, 0x3c, 0xa7, 0xf9, 0xdb,
	0x01, 0xd7, 0x0c, 0xd4, 0xd5, 0x5a, 0x8a, 0xe1, 0x58, 0xa3, 0x32, 0x8d, 0xa7, 0x2c, 0x41, 0xdb,
	0x50, 0x8d, 0xda, 0x37, 0x69, 0x40, 0x9d, 0xa3, 0x8a, 0xa4, 0xc8, 0xb5, 0x19, 0x75, 0xc3, 0x86,
	0x56, 0x29, 0x63, 0x44, 0x2e, 0x22, 0x3d, 0x96, 0xe8, 0x3b, 0x76, 0xe2, 0x39, 0x24, 0xaf, 0x60,
	0x27, 0xca, 0xe2, 0x98, 0x0d, 0x33, 0xc9, 0xb4, 0x98, 0xa0, 0x75, 0xa4, 0x4a, 0xd7, 0x49, 0xf2,
	0x12, 0xea, 0x79, 0xdc, 0x09, 0xe7, 0x66, 0x96, 0xed, 0x0f, 0x90, 0xc7, 0x9d, 0xeb, 0xc2, 0xcf,
	0x16, 0x78, 0x1c, 0x63, 0xd4, 0xc8, 0xc3, 0xe1, 0x34, 0xcc, 0xee, 0x53, 0x94, 0xfe, 0x96, 0xad,
	0xe4, 0x16, 0xfc, 0xc9, 0x74, 0x60, 0x58, 0x72, 0x0c, 0xbe, 0x44, 0xa5, 0xa5, 0x88, 0x8c, 0x78,
	0xfd, 0xef, 0x8a, 0xcd, 0x38, 0x58, 0xc6, 0x83, 0xb5, 0x2e, 0xde, 0xc1, 0x3e, 0xc7, 0x5c, 0x62,
	0xc4, 0x6c, 0x6a, 0x2c, 0x30, 0xd5, 0xa1, 0xe0, 0x7e, 0xb5, 0x51, 0x6a, 0x39, 0x94, 0x2c, 0x63,
	0x81, 0x0d, 0xf5, 0x38, 0x79, 0x0d, 0x6e, 0x3e, 0x1e, 0xc6, 0x22, 0x0a, 0x95, 0x66, 0x52, 0x22,
	0xf7, 0x6b, 0xb3, 0xf1, 0x66, 0xec, 0xd5, 0x8c, 0x24, 0xcf, 0xa0, 0xb6, 0xac, 0x06, 0x76, 0xb8,
	0x6a, 0x54, 0xd4, 0x68, 0xfe, 0xdd, 0x00, 0xb7, 0xa7, 0x31, 0x59, 0x59, 0xc2, 0x53, 0xa8, 0x32,
	0xce, 0xed, 0xac, 0xc5, 0x22, 0x2a, 0x16, 0x9f, 0x4c, 0xc9, 0x73, 0xa8, 0x69, 0x91, 0xa0, 0xd2,
	0x2c, 0xc9, 0xed, 0x26, 0x1c, 0xba, 0x24, 0xcc, 0x1e, 0x12, 0x54, 0x8a, 0xdd, 0xcc, 0xf6, 0x50,
	0xa3, 0x73, 0x68, 0x0f, 0x12, 0x31, 0x2d, 0xec, 0xb7, 0x6f, 0xd3, 0x3d, 0x2f, 0xee, 0x2c, 0x8c,
	0xb2, 0x71, 0xaa, 0xad, 0xf1, 0x0e, 0xdd, 0x99, 0xb3, 0x81, 0x21, 0xc9, 0x27, 0xd8, 0x5d, 0xc8,
	0x46, 0xf6, 0x1e, 0xad, 0xf5, 0xf5, 0xa3, 0x83, 0xff, 0x9c, 0x2b, 0x5d, 0x94, 0x2d, 0xce, 0xfc,
	0x0d, 0xec, 0x2a, 0x9c, 0x60, 0xca, 0xc5, 0x8d, 0xd0, 0x2c, 0x36, 0x2e, 0x54, 0x6c, 0x7b, 0xee,
	0x2a, 0xdd, 0xe3, 0xe4, 0x2d, 0xec, 0xad, 0x09, 0x63, 0x1c, 0xe9, 0xc2, 0x7e, 0x6f, 0x35, 0xd0,
	0xc7, 0x91, 0x26, 0x07, 0x50, 0x31, 0x63, 0x84, 0x4c, 0x5b, 0xd7, 0x1d, 0xba, 0x65, 0x60, 0x57,
	0x93, 0x27, 0xb0, 0x35, 0xf3, 0xdf, 0x7a, 0x5d, 0xa5, 0x05, 0x6a, 0x1e, 0xc3, 0xee, 0x95, 0x96,
	0x22, 0xbd, 0x59, 0x58, 0x4d, 0x3c, 0x70, 0xee, 0x70, 0x6e, 0xb2, 0x79, 0x92, 0x7d, 0x28, 0x4f,
	0x58, 0x3c, 0xc6, 0xe2, 0xcc, 0x67, 0xa0, 0xf9, 0x05, 0xbc, 0x07, 0xa9, 0x8a, 0x7c, 0x80, 0x1a,
	0x9b, 0x23, 0xbf, 0xd4, 0x70, 0x5a, 0xf5, 0x23, 0x7f, 0xe1, 0xc8, 0x03, 0x35, 0x5d, 0x4a, 0xff,
	0x05, 0x00, 0x00, 0xff, 0xff, 0x90, 0x12, 0xa3, 0xb6, 0x64, 0x04, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: social.proto

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

type DecorationData struct {
	Username             *string  `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	FullName             *string  `protobuf:"bytes,2,opt,name=full_name,json=fullName" json:"full_name,omitempty"`
	ImageUrl             *string  `protobuf:"bytes,3,opt,name=image_url,json=imageUrl" json:"image_url,omitempty"`
	LargeImageUrl        *string  `protobuf:"bytes,5,opt,name=large_image_url,json=largeImageUrl" json:"large_image_url,omitempty"`
	FirstName            *string  `protobuf:"bytes,6,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName             *string  `protobuf:"bytes,7,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	FacebookUid          *string  `protobuf:"bytes,8,opt,name=facebook_uid,json=facebookUid" json:"facebook_uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DecorationData) Reset()         { *m = DecorationData{} }
func (m *DecorationData) String() string { return proto.CompactTextString(m) }
func (*DecorationData) ProtoMessage()    {}
func (*DecorationData) Descriptor() ([]byte, []int) {
	return fileDescriptor_social_4bc983cff226ee1e, []int{0}
}
func (m *DecorationData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecorationData.Unmarshal(m, b)
}
func (m *DecorationData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecorationData.Marshal(b, m, deterministic)
}
func (dst *DecorationData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecorationData.Merge(dst, src)
}
func (m *DecorationData) XXX_Size() int {
	return xxx_messageInfo_DecorationData.Size(m)
}
func (m *DecorationData) XXX_DiscardUnknown() {
	xxx_messageInfo_DecorationData.DiscardUnknown(m)
}

var xxx_messageInfo_DecorationData proto.InternalMessageInfo

func (m *DecorationData) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

func (m *DecorationData) GetFullName() string {
	if m != nil && m.FullName != nil {
		return *m.FullName
	}
	return ""
}

func (m *DecorationData) GetImageUrl() string {
	if m != nil && m.ImageUrl != nil {
		return *m.ImageUrl
	}
	return ""
}

func (m *DecorationData) GetLargeImageUrl() string {
	if m != nil && m.LargeImageUrl != nil {
		return *m.LargeImageUrl
	}
	return ""
}

func (m *DecorationData) GetFirstName() string {
	if m != nil && m.FirstName != nil {
		return *m.FirstName
	}
	return ""
}

func (m *DecorationData) GetLastName() string {
	if m != nil && m.LastName != nil {
		return *m.LastName
	}
	return ""
}

func (m *DecorationData) GetFacebookUid() string {
	if m != nil && m.FacebookUid != nil {
		return *m.FacebookUid
	}
	return ""
}

func init() {
	proto.RegisterType((*DecorationData)(nil), "spotify.DecorationData")
}

func init() { proto.RegisterFile("social.proto", fileDescriptor_social_4bc983cff226ee1e) }

var fileDescriptor_social_4bc983cff226ee1e = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8d, 0xbd, 0x4a, 0xc4, 0x40,
	0x14, 0x46, 0x89, 0xa2, 0x49, 0xae, 0x51, 0x61, 0xaa, 0xa0, 0x08, 0x6a, 0x21, 0x56, 0x3e, 0x45,
	0x1a, 0x1b, 0x0b, 0x21, 0xf5, 0x70, 0x37, 0x99, 0x09, 0xc3, 0x4e, 0x72, 0xc3, 0xfc, 0x14, 0xfb,
	0xd0, 0xfb, 0x0e, 0x4b, 0xee, 0xec, 0xec, 0x96, 0xdf, 0x39, 0x07, 0x3e, 0x68, 0x3c, 0x0d, 0x06,
	0xed, 0xcf, 0xea, 0x28, 0x90, 0x28, 0xfd, 0x4a, 0xc1, 0xe8, 0xc3, 0xe7, 0xb1, 0x80, 0xa7, 0x4e,
	0x0d, 0xe4, 0x30, 0x18, 0x5a, 0x3a, 0x0c, 0x28, 0x5e, 0xa0, 0x8a, 0x5e, 0xb9, 0x05, 0x67, 0xd5,
	0x16, 0xef, 0xc5, 0x77, 0xfd, 0x7f, 0xd9, 0xe2, 0x15, 0x6a, 0x1d, 0xad, 0x95, 0x2c, 0x6f, 0x92,
	0xdc, 0xc0, 0xdf, 0x59, 0x9a, 0x19, 0x27, 0x25, 0xa3, 0xb3, 0xed, 0x6d, 0x92, 0x0c, 0x7a, 0x67,
	0xc5, 0x17, 0x3c, 0x5b, 0x74, 0x93, 0x92, 0xd7, 0xe4, 0x8e, 0x93, 0x47, 0xc6, 0xbf, 0xb9, 0x7b,
	0x03, 0xd0, 0xc6, 0xf9, 0x90, 0x2e, 0xee, 0x39, 0xa9, 0x99, 0xe4, 0x0f, 0x8b, 0xd9, 0x96, 0xe9,
	0x63, 0x03, 0x2c, 0x3f, 0xa0, 0xd1, 0x38, 0xa8, 0x1d, 0xd1, 0x5e, 0x46, 0x33, 0xb6, 0x15, 0xfb,
	0x87, 0xcc, 0x7a, 0x33, 0x9e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x2a, 0xe9, 0x2f, 0x07, 0x01,
	0x00, 0x00,
}

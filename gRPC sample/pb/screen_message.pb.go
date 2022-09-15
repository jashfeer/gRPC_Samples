// Code generated by protoc-gen-go. DO NOT EDIT.
// source: screen_message.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Screen_Panel int32

const (
	Screen_UNKOWN Screen_Panel = 0
	Screen_IPS    Screen_Panel = 1
	Screen_OLED   Screen_Panel = 2
)

var Screen_Panel_name = map[int32]string{
	0: "UNKOWN",
	1: "IPS",
	2: "OLED",
}

var Screen_Panel_value = map[string]int32{
	"UNKOWN": 0,
	"IPS":    1,
	"OLED":   2,
}

func (x Screen_Panel) String() string {
	return proto.EnumName(Screen_Panel_name, int32(x))
}

func (Screen_Panel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8a2814cd02b8aba7, []int{0, 0}
}

type Screen struct {
	SizeInch             float32            `protobuf:"fixed32,1,opt,name=size_inch,json=sizeInch,proto3" json:"size_inch,omitempty"`
	Resolution           *Screen_Resolution `protobuf:"bytes,2,opt,name=resolution,proto3" json:"resolution,omitempty"`
	Panel                Screen_Panel       `protobuf:"varint,3,opt,name=panel,proto3,enum=pb.Screen_Panel" json:"panel,omitempty"`
	Multitouch           bool               `protobuf:"varint,4,opt,name=multitouch,proto3" json:"multitouch,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Screen) Reset()         { *m = Screen{} }
func (m *Screen) String() string { return proto.CompactTextString(m) }
func (*Screen) ProtoMessage()    {}
func (*Screen) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a2814cd02b8aba7, []int{0}
}

func (m *Screen) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Screen.Unmarshal(m, b)
}
func (m *Screen) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Screen.Marshal(b, m, deterministic)
}
func (m *Screen) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Screen.Merge(m, src)
}
func (m *Screen) XXX_Size() int {
	return xxx_messageInfo_Screen.Size(m)
}
func (m *Screen) XXX_DiscardUnknown() {
	xxx_messageInfo_Screen.DiscardUnknown(m)
}

var xxx_messageInfo_Screen proto.InternalMessageInfo

func (m *Screen) GetSizeInch() float32 {
	if m != nil {
		return m.SizeInch
	}
	return 0
}

func (m *Screen) GetResolution() *Screen_Resolution {
	if m != nil {
		return m.Resolution
	}
	return nil
}

func (m *Screen) GetPanel() Screen_Panel {
	if m != nil {
		return m.Panel
	}
	return Screen_UNKOWN
}

func (m *Screen) GetMultitouch() bool {
	if m != nil {
		return m.Multitouch
	}
	return false
}

type Screen_Resolution struct {
	Width                uint32   `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Height               uint32   `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Screen_Resolution) Reset()         { *m = Screen_Resolution{} }
func (m *Screen_Resolution) String() string { return proto.CompactTextString(m) }
func (*Screen_Resolution) ProtoMessage()    {}
func (*Screen_Resolution) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a2814cd02b8aba7, []int{0, 0}
}

func (m *Screen_Resolution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Screen_Resolution.Unmarshal(m, b)
}
func (m *Screen_Resolution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Screen_Resolution.Marshal(b, m, deterministic)
}
func (m *Screen_Resolution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Screen_Resolution.Merge(m, src)
}
func (m *Screen_Resolution) XXX_Size() int {
	return xxx_messageInfo_Screen_Resolution.Size(m)
}
func (m *Screen_Resolution) XXX_DiscardUnknown() {
	xxx_messageInfo_Screen_Resolution.DiscardUnknown(m)
}

var xxx_messageInfo_Screen_Resolution proto.InternalMessageInfo

func (m *Screen_Resolution) GetWidth() uint32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Screen_Resolution) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func init() {
	proto.RegisterEnum("pb.Screen_Panel", Screen_Panel_name, Screen_Panel_value)
	proto.RegisterType((*Screen)(nil), "pb.Screen")
	proto.RegisterType((*Screen_Resolution)(nil), "pb.Screen.Resolution")
}

func init() { proto.RegisterFile("screen_message.proto", fileDescriptor_8a2814cd02b8aba7) }

var fileDescriptor_8a2814cd02b8aba7 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4b, 0xfb, 0x40,
	0x14, 0xc4, 0xff, 0xbb, 0x6d, 0xf2, 0x8f, 0x23, 0x95, 0xf0, 0xa8, 0x12, 0x14, 0x24, 0xf4, 0x50,
	0x72, 0xca, 0xa1, 0xe2, 0xc5, 0xb3, 0x1e, 0x8a, 0xd2, 0x96, 0x2d, 0xe2, 0xb1, 0x34, 0x71, 0xe9,
	0x2e, 0xa4, 0xbb, 0x21, 0xbb, 0x41, 0xf0, 0x13, 0xf8, 0xb1, 0xa5, 0x89, 0x68, 0x8e, 0x33, 0xbf,
	0x37, 0xcc, 0xf0, 0x30, 0x75, 0x65, 0x23, 0xa5, 0xd9, 0x1d, 0xa5, 0x73, 0xfb, 0x83, 0xcc, 0xeb,
	0xc6, 0x7a, 0x4b, 0xbc, 0x2e, 0x66, 0x5f, 0x1c, 0xe1, 0xb6, 0x83, 0x74, 0x83, 0x33, 0xa7, 0x3f,
	0xe5, 0x4e, 0x9b, 0x52, 0x25, 0x2c, 0x65, 0x19, 0x17, 0xd1, 0xc9, 0x58, 0x9a, 0x52, 0xd1, 0x3d,
	0xd0, 0x48, 0x67, 0xab, 0xd6, 0x6b, 0x6b, 0x12, 0x9e, 0xb2, 0xec, 0x7c, 0x71, 0x99, 0xd7, 0x45,
	0xde, 0x87, 0x73, 0xf1, 0x0b, 0xc5, 0xe0, 0x90, 0xe6, 0x08, 0xea, 0xbd, 0x91, 0x55, 0x32, 0x4a,
	0x59, 0x76, 0xb1, 0x88, 0x07, 0x89, 0xcd, 0xc9, 0x17, 0x3d, 0xa6, 0x5b, 0xe0, 0xd8, 0x56, 0x5e,
	0x7b, 0xdb, 0x96, 0x2a, 0x19, 0xa7, 0x2c, 0x8b, 0xc4, 0xc0, 0xb9, 0x7e, 0x00, 0xfe, 0x1a, 0x68,
	0x8a, 0xe0, 0x43, 0xbf, 0xfb, 0x7e, 0xe5, 0x44, 0xf4, 0x82, 0xae, 0x10, 0x2a, 0xa9, 0x0f, 0xca,
	0x77, 0xf3, 0x26, 0xe2, 0x47, 0xcd, 0xe6, 0x08, 0xba, 0x2e, 0x02, 0xc2, 0xd7, 0xd5, 0xf3, 0xfa,
	0x6d, 0x15, 0xff, 0xa3, 0xff, 0x18, 0x2d, 0x37, 0xdb, 0x98, 0x51, 0x84, 0xf1, 0xfa, 0xe5, 0xe9,
	0x31, 0xe6, 0x45, 0xd8, 0x7d, 0xe5, 0xee, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x60, 0xdd, 0x16, 0x27,
	0x2d, 0x01, 0x00, 0x00,
}

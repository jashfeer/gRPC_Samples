// Code generated by protoc-gen-go. DO NOT EDIT.
// source: laptop_message.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Laptop struct {
	Id       string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Brand    string     `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand,omitempty"`
	Name     string     `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Cpu      *CPU       `protobuf:"bytes,4,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Ram      *Memory    `protobuf:"bytes,5,opt,name=ram,proto3" json:"ram,omitempty"`
	Gpus     []*GPU     `protobuf:"bytes,6,rep,name=gpus,proto3" json:"gpus,omitempty"`
	Storage  []*Storage `protobuf:"bytes,7,rep,name=storage,proto3" json:"storage,omitempty"`
	Screen   *Screen    `protobuf:"bytes,8,opt,name=screen,proto3" json:"screen,omitempty"`
	Keyboard *Keyboard  `protobuf:"bytes,9,opt,name=keyboard,proto3" json:"keyboard,omitempty"`
	// Types that are valid to be assigned to Weight:
	//	*Laptop_WeightKg
	//	*Laptop_WeightLb
	Weight               isLaptop_Weight      `protobuf_oneof:"weight"`
	PriceUsd             float64              `protobuf:"fixed64,12,opt,name=price_usd,json=priceUsd,proto3" json:"price_usd,omitempty"`
	ReleaseYear          uint32               `protobuf:"varint,13,opt,name=release_year,json=releaseYear,proto3" json:"release_year,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Laptop) Reset()         { *m = Laptop{} }
func (m *Laptop) String() string { return proto.CompactTextString(m) }
func (*Laptop) ProtoMessage()    {}
func (*Laptop) Descriptor() ([]byte, []int) {
	return fileDescriptor_07a3824d46f4b731, []int{0}
}

func (m *Laptop) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Laptop.Unmarshal(m, b)
}
func (m *Laptop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Laptop.Marshal(b, m, deterministic)
}
func (m *Laptop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Laptop.Merge(m, src)
}
func (m *Laptop) XXX_Size() int {
	return xxx_messageInfo_Laptop.Size(m)
}
func (m *Laptop) XXX_DiscardUnknown() {
	xxx_messageInfo_Laptop.DiscardUnknown(m)
}

var xxx_messageInfo_Laptop proto.InternalMessageInfo

func (m *Laptop) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Laptop) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *Laptop) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Laptop) GetCpu() *CPU {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *Laptop) GetRam() *Memory {
	if m != nil {
		return m.Ram
	}
	return nil
}

func (m *Laptop) GetGpus() []*GPU {
	if m != nil {
		return m.Gpus
	}
	return nil
}

func (m *Laptop) GetStorage() []*Storage {
	if m != nil {
		return m.Storage
	}
	return nil
}

func (m *Laptop) GetScreen() *Screen {
	if m != nil {
		return m.Screen
	}
	return nil
}

func (m *Laptop) GetKeyboard() *Keyboard {
	if m != nil {
		return m.Keyboard
	}
	return nil
}

type isLaptop_Weight interface {
	isLaptop_Weight()
}

type Laptop_WeightKg struct {
	WeightKg float64 `protobuf:"fixed64,10,opt,name=weight_kg,json=weightKg,proto3,oneof"`
}

type Laptop_WeightLb struct {
	WeightLb float64 `protobuf:"fixed64,11,opt,name=weight_lb,json=weightLb,proto3,oneof"`
}

func (*Laptop_WeightKg) isLaptop_Weight() {}

func (*Laptop_WeightLb) isLaptop_Weight() {}

func (m *Laptop) GetWeight() isLaptop_Weight {
	if m != nil {
		return m.Weight
	}
	return nil
}

func (m *Laptop) GetWeightKg() float64 {
	if x, ok := m.GetWeight().(*Laptop_WeightKg); ok {
		return x.WeightKg
	}
	return 0
}

func (m *Laptop) GetWeightLb() float64 {
	if x, ok := m.GetWeight().(*Laptop_WeightLb); ok {
		return x.WeightLb
	}
	return 0
}

func (m *Laptop) GetPriceUsd() float64 {
	if m != nil {
		return m.PriceUsd
	}
	return 0
}

func (m *Laptop) GetReleaseYear() uint32 {
	if m != nil {
		return m.ReleaseYear
	}
	return 0
}

func (m *Laptop) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Laptop) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Laptop_WeightKg)(nil),
		(*Laptop_WeightLb)(nil),
	}
}

func init() {
	proto.RegisterType((*Laptop)(nil), "pb.Laptop")
}

func init() { proto.RegisterFile("laptop_message.proto", fileDescriptor_07a3824d46f4b731) }

var fileDescriptor_07a3824d46f4b731 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x49, 0xd3, 0xa5, 0xc9, 0x6b, 0xb7, 0x83, 0x55, 0xc0, 0x74, 0x20, 0xc2, 0x24, 0xa4,
	0x9c, 0x32, 0x09, 0x4e, 0x1c, 0x81, 0x03, 0x48, 0x1b, 0xd2, 0xe4, 0xd1, 0x03, 0xa7, 0xc8, 0xae,
	0x1f, 0x26, 0x5a, 0x53, 0x5b, 0xb6, 0x23, 0xd4, 0x6f, 0xc9, 0x47, 0x42, 0x71, 0x12, 0x58, 0x73,
	0xf3, 0xfb, 0xfd, 0xfe, 0x96, 0x9f, 0xdf, 0x83, 0xf5, 0x9e, 0x1b, 0xaf, 0x4d, 0xd5, 0xa0, 0x73,
	0x5c, 0x61, 0x69, 0xac, 0xf6, 0x9a, 0xcc, 0x8c, 0xd8, 0x3c, 0x37, 0x56, 0xef, 0xd0, 0x39, 0x6d,
	0x4f, 0xe5, 0x66, 0xdd, 0x60, 0xa3, 0xed, 0x71, 0x42, 0x9f, 0x3a, 0xaf, 0x2d, 0x57, 0x38, 0x0d,
	0xbb, 0x9d, 0x45, 0x3c, 0x4c, 0xe8, 0xb3, 0x07, 0x3c, 0x0a, 0xcd, 0xad, 0x9c, 0xf0, 0xd7, 0x4a,
	0x6b, 0xb5, 0xc7, 0xeb, 0x50, 0x89, 0xf6, 0xe7, 0xb5, 0xaf, 0x1b, 0x74, 0x9e, 0x37, 0xa6, 0x0f,
	0x5c, 0xfd, 0x89, 0x21, 0xb9, 0x0d, 0x1d, 0x93, 0x0b, 0x98, 0xd5, 0x92, 0x46, 0x79, 0x54, 0x64,
	0x6c, 0x56, 0x4b, 0xb2, 0x86, 0x33, 0x61, 0xf9, 0x41, 0xd2, 0x59, 0x40, 0x7d, 0x41, 0x08, 0xcc,
	0x0f, 0xbc, 0x41, 0x1a, 0x07, 0x18, 0xce, 0xe4, 0x05, 0xc4, 0x3b, 0xd3, 0xd2, 0x79, 0x1e, 0x15,
	0xcb, 0x77, 0x8b, 0xd2, 0x88, 0xf2, 0xf3, 0xdd, 0x96, 0x75, 0x8c, 0xbc, 0x84, 0xd8, 0xf2, 0x86,
	0x9e, 0x05, 0x05, 0x9d, 0xfa, 0x16, 0x3e, 0xcb, 0x3a, 0x4c, 0x2e, 0x61, 0xae, 0x4c, 0xeb, 0x68,
	0x92, 0xc7, 0xe3, 0xcd, 0x2f, 0x77, 0x5b, 0x16, 0x20, 0x79, 0x0b, 0x8b, 0x61, 0x04, 0x74, 0x11,
	0xfc, 0xb2, 0xf3, 0xf7, 0x3d, 0x62, 0xa3, 0x23, 0x57, 0x90, 0xf4, 0x23, 0xa1, 0xe9, 0xff, 0x47,
	0xee, 0x03, 0x61, 0x83, 0x21, 0x05, 0xa4, 0xe3, 0x80, 0x68, 0x16, 0x52, 0xab, 0x2e, 0x75, 0x33,
	0x30, 0xf6, 0xcf, 0x92, 0x57, 0x90, 0xfd, 0xc6, 0x5a, 0xfd, 0xf2, 0xd5, 0x83, 0xa2, 0x90, 0x47,
	0x45, 0xf4, 0xf5, 0x09, 0x4b, 0x7b, 0x74, 0xa3, 0x1e, 0xe9, 0xbd, 0xa0, 0xcb, 0x53, 0x7d, 0x2b,
	0xc8, 0x25, 0x64, 0xc6, 0xd6, 0x3b, 0xac, 0x5a, 0x27, 0xe9, 0xaa, 0xd3, 0x2c, 0x0d, 0x60, 0xeb,
	0x24, 0x79, 0x03, 0x2b, 0x8b, 0x7b, 0xe4, 0x0e, 0xab, 0x23, 0x72, 0x4b, 0xcf, 0xf3, 0xa8, 0x38,
	0x67, 0xcb, 0x81, 0xfd, 0x40, 0x6e, 0xc9, 0x07, 0x80, 0xd6, 0x48, 0xee, 0x51, 0x56, 0xdc, 0xd3,
	0x8b, 0xd0, 0xe9, 0xa6, 0xec, 0x77, 0x58, 0x8e, 0x3b, 0x2c, 0xbf, 0x8f, 0x3b, 0x64, 0xd9, 0x90,
	0xfe, 0xe8, 0x3f, 0xa5, 0x90, 0xf4, 0x6d, 0x88, 0x24, 0x04, 0xdf, 0xff, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0xee, 0x5d, 0xff, 0x4a, 0x8a, 0x02, 0x00, 0x00,
}

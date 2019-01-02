// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/common.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	types "github.com/openconfig/gnoi/types"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RemoteDownload_Protocol int32

const (
	RemoteDownload_UNKNOWN RemoteDownload_Protocol = 0
	RemoteDownload_SFTP    RemoteDownload_Protocol = 1
	RemoteDownload_HTTP    RemoteDownload_Protocol = 2
	RemoteDownload_HTTPS   RemoteDownload_Protocol = 3
	RemoteDownload_SCP     RemoteDownload_Protocol = 4
)

var RemoteDownload_Protocol_name = map[int32]string{
	0: "UNKNOWN",
	1: "SFTP",
	2: "HTTP",
	3: "HTTPS",
	4: "SCP",
}

var RemoteDownload_Protocol_value = map[string]int32{
	"UNKNOWN": 0,
	"SFTP":    1,
	"HTTP":    2,
	"HTTPS":   3,
	"SCP":     4,
}

func (x RemoteDownload_Protocol) String() string {
	return proto.EnumName(RemoteDownload_Protocol_name, int32(x))
}

func (RemoteDownload_Protocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0, 0}
}

// RemoteDownload defines the details for a device to initiate a file transfer
// from or to a remote location.
type RemoteDownload struct {
	// The path information containing where to download the data from or to.
	// For HTTP(S), this will be the URL (i.e. foo.com/file.tbz2).
	// For SFTP and SCP, this will be the address:/path/to/file
	// (i.e. host.foo.com:/bar/baz).
	Path                 string                  `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Protocol             RemoteDownload_Protocol `protobuf:"varint,2,opt,name=protocol,proto3,enum=common.RemoteDownload_Protocol" json:"protocol,omitempty"`
	Credentials          *types.Credentials      `protobuf:"bytes,3,opt,name=credentials,proto3" json:"credentials,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *RemoteDownload) Reset()         { *m = RemoteDownload{} }
func (m *RemoteDownload) String() string { return proto.CompactTextString(m) }
func (*RemoteDownload) ProtoMessage()    {}
func (*RemoteDownload) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

func (m *RemoteDownload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoteDownload.Unmarshal(m, b)
}
func (m *RemoteDownload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoteDownload.Marshal(b, m, deterministic)
}
func (m *RemoteDownload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteDownload.Merge(m, src)
}
func (m *RemoteDownload) XXX_Size() int {
	return xxx_messageInfo_RemoteDownload.Size(m)
}
func (m *RemoteDownload) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteDownload.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteDownload proto.InternalMessageInfo

func (m *RemoteDownload) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *RemoteDownload) GetProtocol() RemoteDownload_Protocol {
	if m != nil {
		return m.Protocol
	}
	return RemoteDownload_UNKNOWN
}

func (m *RemoteDownload) GetCredentials() *types.Credentials {
	if m != nil {
		return m.Credentials
	}
	return nil
}

func init() {
	proto.RegisterEnum("common.RemoteDownload_Protocol", RemoteDownload_Protocol_name, RemoteDownload_Protocol_value)
	proto.RegisterType((*RemoteDownload)(nil), "common.RemoteDownload")
}

func init() { proto.RegisterFile("common/common.proto", fileDescriptor_8f954d82c0b891f6) }

var fileDescriptor_8f954d82c0b891f6 = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x87, 0x50, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x6c, 0x10, 0x9e, 0x94,
	0x4e, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x7e, 0x41, 0x6a, 0x5e,
	0x72, 0x7e, 0x5e, 0x5a, 0x66, 0xba, 0x7e, 0x7a, 0x5e, 0x7e, 0xa6, 0x7e, 0x49, 0x65, 0x41, 0x6a,
	0x31, 0x84, 0x84, 0xe8, 0x52, 0xba, 0xc1, 0xc8, 0xc5, 0x17, 0x94, 0x9a, 0x9b, 0x5f, 0x92, 0xea,
	0x92, 0x5f, 0x9e, 0x97, 0x93, 0x9f, 0x98, 0x22, 0x24, 0xc4, 0xc5, 0x52, 0x90, 0x58, 0x92, 0x21,
	0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x66, 0x0b, 0x59, 0x73, 0x71, 0x80, 0xd5, 0x27, 0xe7,
	0xe7, 0x48, 0x30, 0x29, 0x30, 0x6a, 0xf0, 0x19, 0xc9, 0xeb, 0x41, 0x6d, 0x47, 0xd5, 0xad, 0x17,
	0x00, 0x55, 0x16, 0x04, 0xd7, 0x20, 0x64, 0xc2, 0xc5, 0x9d, 0x5c, 0x94, 0x9a, 0x92, 0x9a, 0x57,
	0x92, 0x99, 0x98, 0x53, 0x2c, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0xa4, 0x07, 0x71, 0x86,
	0x33, 0x42, 0x26, 0x08, 0x59, 0x99, 0x92, 0x3d, 0x17, 0x07, 0xcc, 0x2c, 0x21, 0x6e, 0x2e, 0xf6,
	0x50, 0x3f, 0x6f, 0x3f, 0xff, 0x70, 0x3f, 0x01, 0x06, 0x21, 0x0e, 0x2e, 0x96, 0x60, 0xb7, 0x90,
	0x00, 0x01, 0x46, 0x10, 0xcb, 0x23, 0x24, 0x24, 0x40, 0x80, 0x49, 0x88, 0x93, 0x8b, 0x15, 0xc4,
	0x0a, 0x16, 0x60, 0x16, 0x62, 0xe7, 0x62, 0x0e, 0x76, 0x0e, 0x10, 0x60, 0x49, 0x62, 0x03, 0x3b,
	0xc0, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x3d, 0xeb, 0xfc, 0x2e, 0x01, 0x00, 0x00,
}

// Code generated by protoc-gen-go.
// source: github.com/micro/go-platform/db/proto/trace.proto
// DO NOT EDIT!

/*
Package trace is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/go-platform/db/proto/trace.proto

It has these top-level messages:
	Database
	Record
*/
package trace

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Database struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Table string `protobuf:"bytes,2,opt,name=table" json:"table,omitempty"`
}

func (m *Database) Reset()                    { *m = Database{} }
func (m *Database) String() string            { return proto.CompactTextString(m) }
func (*Database) ProtoMessage()               {}
func (*Database) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Record struct {
	Id       string            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Created  int64             `protobuf:"varint,2,opt,name=created" json:"created,omitempty"`
	Updated  int64             `protobuf:"varint,3,opt,name=updated" json:"updated,omitempty"`
	Metadata map[string]string `protobuf:"bytes,4,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Bytes    string            `protobuf:"bytes,5,opt,name=bytes" json:"bytes,omitempty"`
}

func (m *Record) Reset()                    { *m = Record{} }
func (m *Record) String() string            { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()               {}
func (*Record) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Record) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*Database)(nil), "Database")
	proto.RegisterType((*Record)(nil), "Record")
}

var fileDescriptor0 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x8f, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0xc9, 0xa6, 0xad, 0x3a, 0xb5, 0x08, 0x0b, 0x42, 0xf0, 0x54, 0x7a, 0xb1, 0x17, 0x13,
	0xd4, 0x8b, 0x78, 0xd6, 0xa3, 0x17, 0xff, 0xc1, 0x24, 0x19, 0xeb, 0xe2, 0xa6, 0x59, 0xb2, 0xb3,
	0xc2, 0xfe, 0x1a, 0xff, 0xaa, 0x9b, 0xb8, 0x20, 0xbd, 0xcd, 0x7c, 0xf3, 0x78, 0xef, 0x0d, 0xdc,
	0x1f, 0x1a, 0xfe, 0x1c, 0xac, 0x76, 0x31, 0x98, 0xd0, 0xb8, 0x14, 0xcd, 0x21, 0xde, 0x75, 0x2d,
	0xf2, 0x47, 0x4c, 0xc1, 0x78, 0x6b, 0xba, 0x14, 0x39, 0x1a, 0x4e, 0xe8, 0x48, 0x97, 0x79, 0x77,
	0x0b, 0xe7, 0x2f, 0xc8, 0x68, 0xb1, 0xa7, 0xfa, 0x12, 0x16, 0x47, 0x0c, 0xa4, 0xc4, 0x56, 0xec,
	0x2f, 0xea, 0x0d, 0x2c, 0x27, 0xde, 0x92, 0xaa, 0xf2, 0xba, 0xfb, 0x11, 0xb0, 0x7a, 0x27, 0x17,
	0x93, 0xaf, 0x01, 0xaa, 0xc6, 0xcf, 0xaa, 0x2b, 0x38, 0x73, 0x89, 0x90, 0xc9, 0x17, 0x9d, 0xcc,
	0x60, 0xe8, 0x7c, 0x01, 0xb2, 0x80, 0x29, 0x21, 0x10, 0xe3, 0x84, 0x50, 0x2d, 0xb6, 0x72, 0xbf,
	0x7e, 0xb8, 0xd6, 0x7f, 0x46, 0xfa, 0x6d, 0xe6, 0xaf, 0x47, 0x4e, 0x63, 0x0e, 0xb4, 0x23, 0x53,
	0xaf, 0x96, 0xd9, 0xf9, 0xc6, 0xc0, 0xe6, 0xf4, 0xbe, 0x06, 0xf9, 0x45, 0xe3, 0x7f, 0xbb, 0x6f,
	0x6c, 0x87, 0xb9, 0xdd, 0x73, 0xf5, 0x24, 0xec, 0xaa, 0x7c, 0xf4, 0xf8, 0x1b, 0x00, 0x00, 0xff,
	0xff, 0xd3, 0x5a, 0x37, 0xe9, 0x06, 0x01, 0x00, 0x00,
}

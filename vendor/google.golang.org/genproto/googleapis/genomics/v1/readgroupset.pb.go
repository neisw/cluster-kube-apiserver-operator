// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/genomics/v1/readgroupset.proto

package genomics // import "google.golang.org/genproto/googleapis/genomics/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _struct "github.com/golang/protobuf/ptypes/struct"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A read group set is a logical collection of read groups, which are
// collections of reads produced by a sequencer. A read group set typically
// models reads corresponding to one sample, sequenced one way, and aligned one
// way.
//
// * A read group set belongs to one dataset.
// * A read group belongs to one read group set.
// * A read belongs to one read group.
//
// For more genomics resource definitions, see [Fundamentals of Google
// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
type ReadGroupSet struct {
	// The server-generated read group set ID, unique for all read group sets.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The dataset to which this read group set belongs.
	DatasetId string `protobuf:"bytes,2,opt,name=dataset_id,json=datasetId,proto3" json:"dataset_id,omitempty"`
	// The reference set to which the reads in this read group set are aligned.
	ReferenceSetId string `protobuf:"bytes,3,opt,name=reference_set_id,json=referenceSetId,proto3" json:"reference_set_id,omitempty"`
	// The read group set name. By default this will be initialized to the sample
	// name of the sequenced data contained in this set.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The filename of the original source file for this read group set, if any.
	Filename string `protobuf:"bytes,5,opt,name=filename,proto3" json:"filename,omitempty"`
	// The read groups in this set. There are typically 1-10 read groups in a read
	// group set.
	ReadGroups []*ReadGroup `protobuf:"bytes,6,rep,name=read_groups,json=readGroups,proto3" json:"read_groups,omitempty"`
	// A map of additional read group set information.
	Info                 map[string]*_struct.ListValue `protobuf:"bytes,7,rep,name=info,proto3" json:"info,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *ReadGroupSet) Reset()         { *m = ReadGroupSet{} }
func (m *ReadGroupSet) String() string { return proto.CompactTextString(m) }
func (*ReadGroupSet) ProtoMessage()    {}
func (*ReadGroupSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_readgroupset_8a06b992947ff227, []int{0}
}
func (m *ReadGroupSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadGroupSet.Unmarshal(m, b)
}
func (m *ReadGroupSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadGroupSet.Marshal(b, m, deterministic)
}
func (dst *ReadGroupSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadGroupSet.Merge(dst, src)
}
func (m *ReadGroupSet) XXX_Size() int {
	return xxx_messageInfo_ReadGroupSet.Size(m)
}
func (m *ReadGroupSet) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadGroupSet.DiscardUnknown(m)
}

var xxx_messageInfo_ReadGroupSet proto.InternalMessageInfo

func (m *ReadGroupSet) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ReadGroupSet) GetDatasetId() string {
	if m != nil {
		return m.DatasetId
	}
	return ""
}

func (m *ReadGroupSet) GetReferenceSetId() string {
	if m != nil {
		return m.ReferenceSetId
	}
	return ""
}

func (m *ReadGroupSet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReadGroupSet) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *ReadGroupSet) GetReadGroups() []*ReadGroup {
	if m != nil {
		return m.ReadGroups
	}
	return nil
}

func (m *ReadGroupSet) GetInfo() map[string]*_struct.ListValue {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*ReadGroupSet)(nil), "google.genomics.v1.ReadGroupSet")
	proto.RegisterMapType((map[string]*_struct.ListValue)(nil), "google.genomics.v1.ReadGroupSet.InfoEntry")
}

func init() {
	proto.RegisterFile("google/genomics/v1/readgroupset.proto", fileDescriptor_readgroupset_8a06b992947ff227)
}

var fileDescriptor_readgroupset_8a06b992947ff227 = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x4f, 0x8b, 0xdb, 0x30,
	0x10, 0xc5, 0xb1, 0xf3, 0xa7, 0xcd, 0xa4, 0x84, 0x54, 0x87, 0x62, 0x4c, 0x03, 0x21, 0x50, 0x08,
	0x3d, 0xc8, 0x4d, 0x7a, 0x29, 0x29, 0xe4, 0x10, 0x28, 0x25, 0xb0, 0x87, 0x60, 0xc3, 0x1e, 0xf6,
	0x12, 0x14, 0x7b, 0x6c, 0xc4, 0x3a, 0x92, 0x91, 0xe4, 0x40, 0xbe, 0xf3, 0x7e, 0x80, 0x3d, 0x2e,
	0x96, 0xff, 0x10, 0xd8, 0x25, 0xb7, 0xd1, 0xd3, 0xef, 0x8d, 0x46, 0x6f, 0xe0, 0x47, 0x26, 0x65,
	0x96, 0x63, 0x90, 0xa1, 0x90, 0x67, 0x1e, 0xeb, 0xe0, 0xb2, 0x0a, 0x14, 0xb2, 0x24, 0x53, 0xb2,
	0x2c, 0x34, 0x1a, 0x5a, 0x28, 0x69, 0x24, 0x21, 0x35, 0x46, 0x5b, 0x8c, 0x5e, 0x56, 0xfe, 0xf7,
	0xc6, 0xca, 0x0a, 0x1e, 0x30, 0x21, 0xa4, 0x61, 0x86, 0x4b, 0xa1, 0x6b, 0x87, 0xbf, 0xb8, 0xd7,
	0xb8, 0x61, 0xda, 0x0e, 0xf6, 0x74, 0x2a, 0xd3, 0x40, 0x1b, 0x55, 0xc6, 0xcd, 0x9b, 0x8b, 0x17,
	0x17, 0xbe, 0x84, 0xc8, 0x92, 0xff, 0x95, 0x23, 0x42, 0x43, 0x26, 0xe0, 0xf2, 0xc4, 0x73, 0xe6,
	0xce, 0x72, 0x14, 0xba, 0x3c, 0x21, 0x33, 0x80, 0x84, 0x19, 0xa6, 0xd1, 0x1c, 0x79, 0xe2, 0xb9,
	0x56, 0x1f, 0x35, 0xca, 0x3e, 0x21, 0x4b, 0x98, 0x2a, 0x4c, 0x51, 0xa1, 0x88, 0xf1, 0xd8, 0x40,
	0x3d, 0x0b, 0x4d, 0x3a, 0x3d, 0xb2, 0x24, 0x81, 0xbe, 0x60, 0x67, 0xf4, 0xfa, 0xf6, 0xd6, 0xd6,
	0xc4, 0x87, 0xcf, 0x29, 0xcf, 0xd1, 0xea, 0x03, 0xab, 0x77, 0x67, 0xb2, 0x85, 0x71, 0xf5, 0x95,
	0x63, 0x1d, 0x92, 0x37, 0x9c, 0xf7, 0x96, 0xe3, 0xf5, 0x8c, 0xbe, 0xcf, 0x88, 0x76, 0xf3, 0x87,
	0xa0, 0xda, 0x52, 0x93, 0x2d, 0xf4, 0xb9, 0x48, 0xa5, 0xf7, 0xc9, 0x1a, 0x7f, 0xde, 0x35, 0x46,
	0x68, 0xe8, 0x5e, 0xa4, 0xf2, 0x9f, 0x30, 0xea, 0x1a, 0x5a, 0x9f, 0x1f, 0xc1, 0xa8, 0x93, 0xc8,
	0x14, 0x7a, 0xcf, 0x78, 0x6d, 0x62, 0xa9, 0x4a, 0xf2, 0x0b, 0x06, 0x17, 0x96, 0x97, 0x68, 0x23,
	0x19, 0xaf, 0xfd, 0xb6, 0x7f, 0x1b, 0x33, 0x7d, 0xe0, 0xda, 0x3c, 0x56, 0x44, 0x58, 0x83, 0x1b,
	0xf7, 0x8f, 0xb3, 0xcb, 0xe1, 0x5b, 0x2c, 0xcf, 0x1f, 0xcc, 0xb2, 0xfb, 0x7a, 0x3b, 0xcc, 0xa1,
	0x6a, 0x72, 0x70, 0x9e, 0x36, 0x2d, 0x28, 0x73, 0x26, 0x32, 0x2a, 0x55, 0x56, 0xad, 0xda, 0x3e,
	0x11, 0xd4, 0x57, 0xac, 0xe0, 0xfa, 0x76, 0xfd, 0x7f, 0xdb, 0xfa, 0xd5, 0x71, 0x4e, 0x43, 0x4b,
	0xfe, 0x7e, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x1f, 0xa9, 0x2f, 0xa5, 0x80, 0x02, 0x00, 0x00,
}

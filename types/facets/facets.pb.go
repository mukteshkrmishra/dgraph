// Code generated by protoc-gen-gogo.
// source: types/facets/facets.proto
// DO NOT EDIT!

/*
	Package facets is a generated protocol buffer package.

	It is generated from these files:
		types/facets/facets.proto

	It has these top-level messages:
		Facet
*/
package facets

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Facet_ValType int32

const (
	Facet_STRING   Facet_ValType = 0
	Facet_INT32    Facet_ValType = 1
	Facet_FLOAT    Facet_ValType = 2
	Facet_BOOL     Facet_ValType = 3
	Facet_DATE     Facet_ValType = 4
	Facet_DATETIME Facet_ValType = 5
)

var Facet_ValType_name = map[int32]string{
	0: "STRING",
	1: "INT32",
	2: "FLOAT",
	3: "BOOL",
	4: "DATE",
	5: "DATETIME",
}
var Facet_ValType_value = map[string]int32{
	"STRING":   0,
	"INT32":    1,
	"FLOAT":    2,
	"BOOL":     3,
	"DATE":     4,
	"DATETIME": 5,
}

func (x Facet_ValType) String() string {
	return proto.EnumName(Facet_ValType_name, int32(x))
}
func (Facet_ValType) EnumDescriptor() ([]byte, []int) { return fileDescriptorFacets, []int{0, 0} }

type Facet struct {
	Key     string        `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value   []byte        `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	ValType Facet_ValType `protobuf:"varint,3,opt,name=val_type,json=valType,proto3,enum=facets.Facet_ValType" json:"val_type,omitempty"`
}

func (m *Facet) Reset()                    { *m = Facet{} }
func (m *Facet) String() string            { return proto.CompactTextString(m) }
func (*Facet) ProtoMessage()               {}
func (*Facet) Descriptor() ([]byte, []int) { return fileDescriptorFacets, []int{0} }

func (m *Facet) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Facet) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Facet) GetValType() Facet_ValType {
	if m != nil {
		return m.ValType
	}
	return Facet_STRING
}

func init() {
	proto.RegisterType((*Facet)(nil), "facets.Facet")
	proto.RegisterEnum("facets.Facet_ValType", Facet_ValType_name, Facet_ValType_value)
}
func (m *Facet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Facet) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintFacets(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if len(m.Value) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintFacets(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	if m.ValType != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintFacets(dAtA, i, uint64(m.ValType))
	}
	return i, nil
}

func encodeFixed64Facets(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Facets(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintFacets(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Facet) Size() (n int) {
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovFacets(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovFacets(uint64(l))
	}
	if m.ValType != 0 {
		n += 1 + sovFacets(uint64(m.ValType))
	}
	return n
}

func sovFacets(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFacets(x uint64) (n int) {
	return sovFacets(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Facet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFacets
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Facet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Facet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFacets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFacets
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFacets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFacets
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValType", wireType)
			}
			m.ValType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFacets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValType |= (Facet_ValType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFacets(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFacets
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipFacets(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFacets
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFacets
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFacets
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthFacets
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFacets
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipFacets(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthFacets = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFacets   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("types/facets/facets.proto", fileDescriptorFacets) }

var fileDescriptorFacets = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0xa9, 0x2c, 0x48,
	0x2d, 0xd6, 0x4f, 0x4b, 0x4c, 0x4e, 0x2d, 0x81, 0x51, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42,
	0x6c, 0x10, 0x9e, 0xd2, 0x06, 0x46, 0x2e, 0x56, 0x37, 0x10, 0x53, 0x48, 0x80, 0x8b, 0x39, 0x3b,
	0xb5, 0x52, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc4, 0x14, 0x12, 0xe1, 0x62, 0x2d, 0x4b,
	0xcc, 0x29, 0x4d, 0x95, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x09, 0x82, 0x70, 0x84, 0x0c, 0xb8, 0x38,
	0xca, 0x12, 0x73, 0xe2, 0x41, 0x46, 0x4b, 0x30, 0x2b, 0x30, 0x6a, 0xf0, 0x19, 0x89, 0xea, 0x41,
	0x8d, 0x06, 0x1b, 0xa4, 0x17, 0x96, 0x98, 0x13, 0x52, 0x59, 0x90, 0x1a, 0xc4, 0x5e, 0x06, 0x61,
	0x28, 0xf9, 0x72, 0xb1, 0x43, 0xc5, 0x84, 0xb8, 0xb8, 0xd8, 0x82, 0x43, 0x82, 0x3c, 0xfd, 0xdc,
	0x05, 0x18, 0x84, 0x38, 0xb9, 0x58, 0x3d, 0xfd, 0x42, 0x8c, 0x8d, 0x04, 0x18, 0x41, 0x4c, 0x37,
	0x1f, 0x7f, 0xc7, 0x10, 0x01, 0x26, 0x21, 0x0e, 0x2e, 0x16, 0x27, 0x7f, 0x7f, 0x1f, 0x01, 0x66,
	0x10, 0xcb, 0xc5, 0x31, 0xc4, 0x55, 0x80, 0x45, 0x88, 0x87, 0x8b, 0x03, 0xc4, 0x0a, 0xf1, 0xf4,
	0x75, 0x15, 0x60, 0x75, 0x12, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f,
	0xe4, 0x18, 0x67, 0x3c, 0x96, 0x63, 0x48, 0x62, 0x03, 0xfb, 0xc9, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0xde, 0xb0, 0xfc, 0x22, 0xf0, 0x00, 0x00, 0x00,
}

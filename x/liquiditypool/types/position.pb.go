// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sunrise/liquiditypool/position.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/cosmos/gogoproto/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Position struct {
	Id     uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Sender string `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (m *Position) Reset()         { *m = Position{} }
func (m *Position) String() string { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()    {}
func (*Position) Descriptor() ([]byte, []int) {
	return fileDescriptor_27571095523a5953, []int{0}
}
func (m *Position) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Position) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Position.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Position) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Position.Merge(m, src)
}
func (m *Position) XXX_Size() int {
	return m.Size()
}
func (m *Position) XXX_DiscardUnknown() {
	xxx_messageInfo_Position.DiscardUnknown(m)
}

var xxx_messageInfo_Position proto.InternalMessageInfo

func (m *Position) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Position) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func init() {
	proto.RegisterType((*Position)(nil), "sunrise.liquiditypool.Position")
}

func init() {
	proto.RegisterFile("sunrise/liquiditypool/position.proto", fileDescriptor_27571095523a5953)
}

var fileDescriptor_27571095523a5953 = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x29, 0x2e, 0xcd, 0x2b,
	0xca, 0x2c, 0x4e, 0xd5, 0xcf, 0xc9, 0x2c, 0x2c, 0xcd, 0x4c, 0xc9, 0x2c, 0xa9, 0x2c, 0xc8, 0xcf,
	0xcf, 0xd1, 0x2f, 0xc8, 0x2f, 0xce, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x12, 0x85, 0xaa, 0xd2, 0x43, 0x51, 0xa5, 0x64, 0xc4, 0xc5, 0x11, 0x00, 0x55, 0x28, 0xc4,
	0xc7, 0xc5, 0x94, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x12, 0xc4, 0x94, 0x99, 0x22, 0x24,
	0xc6, 0xc5, 0x56, 0x9c, 0x9a, 0x97, 0x92, 0x5a, 0x24, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x19, 0x04,
	0xe5, 0x39, 0xf9, 0x9f, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c,
	0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x69, 0x7a,
	0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0xd4, 0xbe, 0x9c, 0xc4, 0xca, 0xd4,
	0x22, 0x18, 0x47, 0xbf, 0x02, 0xcd, 0x91, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0x27,
	0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x12, 0xb0, 0x9c, 0x91, 0xca, 0x00, 0x00, 0x00,
}

func (m *Position) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Position) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Position) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintPosition(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintPosition(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPosition(dAtA []byte, offset int, v uint64) int {
	offset -= sovPosition(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Position) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovPosition(uint64(m.Id))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovPosition(uint64(l))
	}
	return n
}

func sovPosition(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPosition(x uint64) (n int) {
	return sovPosition(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Position) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPosition
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Position: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Position: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPosition
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPosition
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPosition
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPosition
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPosition(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPosition
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
func skipPosition(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPosition
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
					return 0, ErrIntOverflowPosition
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPosition
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
			if length < 0 {
				return 0, ErrInvalidLengthPosition
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPosition
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPosition
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPosition        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPosition          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPosition = fmt.Errorf("proto: unexpected end of group")
)

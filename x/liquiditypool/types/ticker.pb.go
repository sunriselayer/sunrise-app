// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sunrise/liquiditypool/ticker.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	cosmossdk_io_math "cosmossdk.io/math"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

type TickInfo struct {
	PoolId         uint64                                      `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	TickIndex      int64                                       `protobuf:"varint,2,opt,name=tick_index,json=tickIndex,proto3" json:"tick_index,omitempty"`
	LiquidityGross cosmossdk_io_math.LegacyDec                 `protobuf:"bytes,3,opt,name=liquidity_gross,json=liquidityGross,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"liquidity_gross" yaml:"liquidity_gross"`
	LiquidityNet   cosmossdk_io_math.LegacyDec                 `protobuf:"bytes,4,opt,name=liquidity_net,json=liquidityNet,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"liquidity_net" yaml:"liquidity_net"`
	FeeGrowth      github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,5,rep,name=fee_growth,json=feeGrowth,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"fee_growth"`
}

func (m *TickInfo) Reset()         { *m = TickInfo{} }
func (m *TickInfo) String() string { return proto.CompactTextString(m) }
func (*TickInfo) ProtoMessage()    {}
func (*TickInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_675149ba84fd9eb6, []int{0}
}
func (m *TickInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TickInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TickInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TickInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TickInfo.Merge(m, src)
}
func (m *TickInfo) XXX_Size() int {
	return m.Size()
}
func (m *TickInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TickInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TickInfo proto.InternalMessageInfo

func (m *TickInfo) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *TickInfo) GetTickIndex() int64 {
	if m != nil {
		return m.TickIndex
	}
	return 0
}

func (m *TickInfo) GetFeeGrowth() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.FeeGrowth
	}
	return nil
}

func init() {
	proto.RegisterType((*TickInfo)(nil), "sunrise.liquiditypool.TickInfo")
}

func init() {
	proto.RegisterFile("sunrise/liquiditypool/ticker.proto", fileDescriptor_675149ba84fd9eb6)
}

var fileDescriptor_675149ba84fd9eb6 = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcd, 0xaa, 0xd3, 0x40,
	0x14, 0xc7, 0x93, 0xdb, 0xeb, 0xd5, 0x8e, 0x5f, 0x10, 0xae, 0x1a, 0xaa, 0x26, 0x21, 0xab, 0x80,
	0x38, 0x43, 0x2d, 0x6e, 0x14, 0x37, 0xb5, 0x50, 0x0a, 0xa2, 0x10, 0x5c, 0xb9, 0xa9, 0xc9, 0xe4,
	0x24, 0x1d, 0xf2, 0x31, 0x31, 0x33, 0xd5, 0xe6, 0x2d, 0x7c, 0x0e, 0x9f, 0xa4, 0xcb, 0x2e, 0xc5,
	0x45, 0x95, 0xf6, 0x0d, 0x7c, 0x00, 0x91, 0x49, 0xd2, 0xaa, 0x5d, 0xb9, 0x9a, 0x39, 0x67, 0xfe,
	0xe7, 0xff, 0x3b, 0x9c, 0x33, 0xc8, 0x15, 0xcb, 0xa2, 0x62, 0x02, 0x48, 0xc6, 0x3e, 0x2c, 0x59,
	0xc4, 0x64, 0x5d, 0x72, 0x9e, 0x11, 0xc9, 0x68, 0x0a, 0x15, 0x2e, 0x2b, 0x2e, 0xb9, 0x71, 0xa7,
	0xd3, 0xe0, 0x7f, 0x34, 0x83, 0xcb, 0x84, 0x27, 0xbc, 0x51, 0x10, 0x75, 0x6b, 0xc5, 0x03, 0x8b,
	0x72, 0x91, 0x73, 0x41, 0xc2, 0x40, 0x00, 0xf9, 0x38, 0x0c, 0x41, 0x06, 0x43, 0x42, 0x39, 0x2b,
	0xda, 0x77, 0xf7, 0xd7, 0x19, 0xba, 0xf6, 0x96, 0xd1, 0x74, 0x56, 0xc4, 0xdc, 0xb8, 0x87, 0xae,
	0x2a, 0xab, 0x39, 0x8b, 0x4c, 0xdd, 0xd1, 0xbd, 0x73, 0xff, 0x42, 0x85, 0xb3, 0xc8, 0x78, 0x88,
	0x90, 0x6a, 0x61, 0xce, 0x8a, 0x08, 0x56, 0xe6, 0x99, 0xa3, 0x7b, 0x3d, 0xbf, 0x2f, 0x9b, 0xb2,
	0x08, 0x56, 0x46, 0x8c, 0x6e, 0x1f, 0x7b, 0x99, 0x27, 0x15, 0x17, 0xc2, 0xec, 0x39, 0xba, 0xd7,
	0x1f, 0xbf, 0x58, 0x6f, 0x6d, 0xed, 0xdb, 0xd6, 0xbe, 0xdf, 0x76, 0x21, 0xa2, 0x14, 0x33, 0x4e,
	0xf2, 0x40, 0x2e, 0xf0, 0x2b, 0x48, 0x02, 0x5a, 0x4f, 0x80, 0xfe, 0xdc, 0xda, 0x77, 0xeb, 0x20,
	0xcf, 0x9e, 0xb9, 0x27, 0x1e, 0xae, 0x7f, 0xeb, 0x98, 0x99, 0xaa, 0x84, 0xf1, 0x1e, 0xdd, 0xfc,
	0xa3, 0x29, 0x40, 0x9a, 0xe7, 0x0d, 0xe5, 0xf9, 0xff, 0x51, 0x2e, 0x4f, 0x29, 0x05, 0x48, 0xd7,
	0xbf, 0x71, 0x8c, 0x5f, 0x83, 0x34, 0x4a, 0x84, 0x62, 0x00, 0xc5, 0xff, 0x24, 0x17, 0xe6, 0x15,
	0xa7, 0xe7, 0x5d, 0x7f, 0xf2, 0x00, 0xb7, 0xbe, 0x58, 0xcd, 0x10, 0x77, 0x33, 0xc4, 0x13, 0xa0,
	0x2f, 0x39, 0x2b, 0xc6, 0x23, 0x05, 0xff, 0xf2, 0xdd, 0x7e, 0x94, 0x30, 0xb9, 0x58, 0x86, 0x98,
	0xf2, 0x9c, 0x74, 0x33, 0x6f, 0x8f, 0xc7, 0x22, 0x4a, 0x89, 0xac, 0x4b, 0x10, 0x87, 0x1a, 0xe1,
	0xf7, 0x63, 0x80, 0x69, 0xc3, 0x18, 0xbf, 0x59, 0xef, 0x2c, 0x7d, 0xb3, 0xb3, 0xf4, 0x1f, 0x3b,
	0x4b, 0xff, 0xbc, 0xb7, 0xb4, 0xcd, 0xde, 0xd2, 0xbe, 0xee, 0x2d, 0xed, 0xdd, 0xd3, 0xbf, 0x1c,
	0xbb, 0x95, 0x67, 0x41, 0x0d, 0xd5, 0x21, 0x20, 0xab, 0xd3, 0x5f, 0xa2, 0x20, 0xe1, 0x45, 0xb3,
	0xd8, 0xd1, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xeb, 0xdf, 0x07, 0x50, 0x4b, 0x02, 0x00, 0x00,
}

func (m *TickInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TickInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TickInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FeeGrowth) > 0 {
		for iNdEx := len(m.FeeGrowth) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FeeGrowth[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTicker(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	{
		size := m.LiquidityNet.Size()
		i -= size
		if _, err := m.LiquidityNet.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTicker(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.LiquidityGross.Size()
		i -= size
		if _, err := m.LiquidityGross.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTicker(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.TickIndex != 0 {
		i = encodeVarintTicker(dAtA, i, uint64(m.TickIndex))
		i--
		dAtA[i] = 0x10
	}
	if m.PoolId != 0 {
		i = encodeVarintTicker(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTicker(dAtA []byte, offset int, v uint64) int {
	offset -= sovTicker(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TickInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovTicker(uint64(m.PoolId))
	}
	if m.TickIndex != 0 {
		n += 1 + sovTicker(uint64(m.TickIndex))
	}
	l = m.LiquidityGross.Size()
	n += 1 + l + sovTicker(uint64(l))
	l = m.LiquidityNet.Size()
	n += 1 + l + sovTicker(uint64(l))
	if len(m.FeeGrowth) > 0 {
		for _, e := range m.FeeGrowth {
			l = e.Size()
			n += 1 + l + sovTicker(uint64(l))
		}
	}
	return n
}

func sovTicker(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTicker(x uint64) (n int) {
	return sovTicker(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TickInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTicker
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
			return fmt.Errorf("proto: TickInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TickInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TickIndex", wireType)
			}
			m.TickIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TickIndex |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidityGross", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicker
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
				return ErrInvalidLengthTicker
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LiquidityGross.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidityNet", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicker
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
				return ErrInvalidLengthTicker
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LiquidityNet.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeGrowth", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTicker
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTicker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeGrowth = append(m.FeeGrowth, types.DecCoin{})
			if err := m.FeeGrowth[len(m.FeeGrowth)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTicker(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTicker
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
func skipTicker(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTicker
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
					return 0, ErrIntOverflowTicker
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
					return 0, ErrIntOverflowTicker
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
				return 0, ErrInvalidLengthTicker
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTicker
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTicker
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTicker        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTicker          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTicker = fmt.Errorf("proto: unexpected end of group")
)

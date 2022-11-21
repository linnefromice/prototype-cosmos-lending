// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lending/lending/pair_pool.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type PairPool struct {
	Id                         uint64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Address                    string     `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	PoolId                     uint64     `protobuf:"varint,3,opt,name=poolId,proto3" json:"poolId,omitempty"`
	AssetLiquidity             types.Coin `protobuf:"bytes,4,opt,name=assetLiquidity,proto3" json:"assetLiquidity"`
	AssetLpCoinDenom           string     `protobuf:"bytes,5,opt,name=assetLpCoinDenom,proto3" json:"assetLpCoinDenom,omitempty"`
	AssetTotalNormalDeposited  uint64     `protobuf:"varint,6,opt,name=assetTotalNormalDeposited,proto3" json:"assetTotalNormalDeposited,omitempty"`
	AssetTotalConlyDeposited   uint64     `protobuf:"varint,7,opt,name=assetTotalConlyDeposited,proto3" json:"assetTotalConlyDeposited,omitempty"`
	AssetTotalBorrowed         uint64     `protobuf:"varint,8,opt,name=assetTotalBorrowed,proto3" json:"assetTotalBorrowed,omitempty"`
	ShadowLiquidity            types.Coin `protobuf:"bytes,9,opt,name=shadowLiquidity,proto3" json:"shadowLiquidity"`
	ShadowLpCoinDenom          string     `protobuf:"bytes,10,opt,name=shadowLpCoinDenom,proto3" json:"shadowLpCoinDenom,omitempty"`
	ShadowTotalNormalDeposited uint64     `protobuf:"varint,11,opt,name=shadowTotalNormalDeposited,proto3" json:"shadowTotalNormalDeposited,omitempty"`
	ShadowTotalConlyDeposited  uint64     `protobuf:"varint,12,opt,name=shadowTotalConlyDeposited,proto3" json:"shadowTotalConlyDeposited,omitempty"`
	ShadowTotalBorrowed        uint64     `protobuf:"varint,13,opt,name=shadowTotalBorrowed,proto3" json:"shadowTotalBorrowed,omitempty"`
	LastUpdated                uint64     `protobuf:"varint,14,opt,name=lastUpdated,proto3" json:"lastUpdated,omitempty"`
}

func (m *PairPool) Reset()         { *m = PairPool{} }
func (m *PairPool) String() string { return proto.CompactTextString(m) }
func (*PairPool) ProtoMessage()    {}
func (*PairPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_77ac034e4c5b84fb, []int{0}
}
func (m *PairPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PairPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PairPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PairPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PairPool.Merge(m, src)
}
func (m *PairPool) XXX_Size() int {
	return m.Size()
}
func (m *PairPool) XXX_DiscardUnknown() {
	xxx_messageInfo_PairPool.DiscardUnknown(m)
}

var xxx_messageInfo_PairPool proto.InternalMessageInfo

func (m *PairPool) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PairPool) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *PairPool) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *PairPool) GetAssetLiquidity() types.Coin {
	if m != nil {
		return m.AssetLiquidity
	}
	return types.Coin{}
}

func (m *PairPool) GetAssetLpCoinDenom() string {
	if m != nil {
		return m.AssetLpCoinDenom
	}
	return ""
}

func (m *PairPool) GetAssetTotalNormalDeposited() uint64 {
	if m != nil {
		return m.AssetTotalNormalDeposited
	}
	return 0
}

func (m *PairPool) GetAssetTotalConlyDeposited() uint64 {
	if m != nil {
		return m.AssetTotalConlyDeposited
	}
	return 0
}

func (m *PairPool) GetAssetTotalBorrowed() uint64 {
	if m != nil {
		return m.AssetTotalBorrowed
	}
	return 0
}

func (m *PairPool) GetShadowLiquidity() types.Coin {
	if m != nil {
		return m.ShadowLiquidity
	}
	return types.Coin{}
}

func (m *PairPool) GetShadowLpCoinDenom() string {
	if m != nil {
		return m.ShadowLpCoinDenom
	}
	return ""
}

func (m *PairPool) GetShadowTotalNormalDeposited() uint64 {
	if m != nil {
		return m.ShadowTotalNormalDeposited
	}
	return 0
}

func (m *PairPool) GetShadowTotalConlyDeposited() uint64 {
	if m != nil {
		return m.ShadowTotalConlyDeposited
	}
	return 0
}

func (m *PairPool) GetShadowTotalBorrowed() uint64 {
	if m != nil {
		return m.ShadowTotalBorrowed
	}
	return 0
}

func (m *PairPool) GetLastUpdated() uint64 {
	if m != nil {
		return m.LastUpdated
	}
	return 0
}

func init() {
	proto.RegisterType((*PairPool)(nil), "linnefromice.lending.lending.PairPool")
}

func init() { proto.RegisterFile("lending/lending/pair_pool.proto", fileDescriptor_77ac034e4c5b84fb) }

var fileDescriptor_77ac034e4c5b84fb = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4f, 0x6b, 0x13, 0x41,
	0x14, 0xcf, 0xc6, 0x34, 0x6d, 0x5f, 0x34, 0xea, 0x28, 0x32, 0x0d, 0xb2, 0x0d, 0x9e, 0x82, 0xc8,
	0xac, 0xd5, 0x9b, 0x88, 0x87, 0xb4, 0x20, 0x01, 0x91, 0x12, 0xf4, 0xe2, 0x45, 0x26, 0x99, 0x31,
	0x1d, 0x98, 0x9d, 0xb7, 0xce, 0x4c, 0xad, 0xf9, 0x16, 0x7e, 0x1a, 0x3f, 0x43, 0x8f, 0x3d, 0x7a,
	0x12, 0x49, 0xbe, 0x88, 0xec, 0x6c, 0x92, 0x5d, 0xf2, 0x47, 0x3c, 0xcd, 0x7b, 0xbf, 0x3f, 0xcc,
	0x8f, 0x1f, 0x3c, 0x38, 0xd6, 0xd2, 0x08, 0x65, 0x26, 0xc9, 0xf2, 0xcd, 0xb8, 0xb2, 0x9f, 0x33,
	0x44, 0xcd, 0x32, 0x8b, 0x1e, 0xc9, 0x63, 0xad, 0x8c, 0x91, 0x5f, 0x2c, 0xa6, 0x6a, 0x2c, 0xd9,
	0x42, 0xb5, 0x7c, 0x3b, 0x0f, 0x27, 0x38, 0xc1, 0x20, 0x4c, 0xf2, 0xa9, 0xf0, 0x74, 0xe2, 0x31,
	0xba, 0x14, 0x5d, 0x32, 0xe2, 0x4e, 0x26, 0xdf, 0x4e, 0x46, 0xd2, 0xf3, 0x93, 0x64, 0x8c, 0xca,
	0x14, 0xfc, 0x93, 0x9f, 0x7b, 0x70, 0x70, 0xce, 0x95, 0x3d, 0x47, 0xd4, 0xa4, 0x0d, 0x75, 0x25,
	0x68, 0xd4, 0x8d, 0x7a, 0x8d, 0x61, 0x5d, 0x09, 0x42, 0x61, 0x9f, 0x0b, 0x61, 0xa5, 0x73, 0xb4,
	0xde, 0x8d, 0x7a, 0x87, 0xc3, 0xe5, 0x4a, 0x1e, 0x41, 0x33, 0x0f, 0x36, 0x10, 0xf4, 0x56, 0x50,
	0x2f, 0x36, 0xf2, 0x16, 0xda, 0xdc, 0x39, 0xe9, 0xdf, 0xa9, 0xaf, 0x97, 0x4a, 0x28, 0x3f, 0xa5,
	0x8d, 0x6e, 0xd4, 0x6b, 0xbd, 0x38, 0x62, 0x45, 0x0e, 0x96, 0xe7, 0x60, 0x8b, 0x1c, 0xec, 0x14,
	0x95, 0xe9, 0x37, 0xae, 0x7f, 0x1f, 0xd7, 0x86, 0x6b, 0x36, 0xf2, 0x14, 0xee, 0x15, 0x48, 0x96,
	0x8b, 0xce, 0xa4, 0xc1, 0x94, 0xee, 0x85, 0x0c, 0x1b, 0x38, 0x79, 0x0d, 0x47, 0x01, 0xfb, 0x80,
	0x9e, 0xeb, 0xf7, 0x68, 0x53, 0xae, 0xcf, 0x64, 0x86, 0x4e, 0x79, 0x29, 0x68, 0x33, 0xe4, 0xdb,
	0x2d, 0x20, 0xaf, 0x80, 0x96, 0xe4, 0x29, 0x1a, 0x3d, 0x2d, 0xcd, 0xfb, 0xc1, 0xbc, 0x93, 0x27,
	0x0c, 0x48, 0xc9, 0xf5, 0xd1, 0x5a, 0xbc, 0x92, 0x82, 0x1e, 0x04, 0xd7, 0x16, 0x86, 0x0c, 0xe0,
	0xae, 0xbb, 0xe0, 0x02, 0xaf, 0xca, 0x7e, 0x0e, 0xff, 0xaf, 0x9f, 0x75, 0x1f, 0x79, 0x06, 0xf7,
	0x17, 0x50, 0xa5, 0x21, 0x08, 0x0d, 0x6d, 0x12, 0xe4, 0x0d, 0x74, 0x0a, 0x70, 0x6b, 0x47, 0xad,
	0x10, 0xf8, 0x1f, 0x8a, 0xbc, 0xe2, 0x0a, 0xbb, 0xd6, 0xd2, 0xed, 0xa2, 0xe2, 0x9d, 0x02, 0xf2,
	0x1c, 0x1e, 0x54, 0xc8, 0x55, 0x4f, 0x77, 0x82, 0x6f, 0x1b, 0x45, 0xba, 0xd0, 0xd2, 0xdc, 0xf9,
	0x8f, 0x99, 0xe0, 0xf9, 0x0f, 0xed, 0xa0, 0xac, 0x42, 0xfd, 0xc1, 0xf5, 0x2c, 0x8e, 0x6e, 0x66,
	0x71, 0xf4, 0x67, 0x16, 0x47, 0x3f, 0xe6, 0x71, 0xed, 0x66, 0x1e, 0xd7, 0x7e, 0xcd, 0xe3, 0xda,
	0xa7, 0x64, 0xa2, 0xfc, 0xc5, 0xe5, 0x88, 0x8d, 0x31, 0x4d, 0xaa, 0x17, 0xb3, 0xba, 0xab, 0xef,
	0xab, 0xc9, 0x4f, 0x33, 0xe9, 0x46, 0xcd, 0x70, 0x0a, 0x2f, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff,
	0x9a, 0x45, 0xd7, 0xfd, 0x81, 0x03, 0x00, 0x00,
}

func (m *PairPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PairPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PairPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastUpdated != 0 {
		i = encodeVarintPairPool(dAtA, i, uint64(m.LastUpdated))
		i--
		dAtA[i] = 0x70
	}
	if m.ShadowTotalBorrowed != 0 {
		i = encodeVarintPairPool(dAtA, i, uint64(m.ShadowTotalBorrowed))
		i--
		dAtA[i] = 0x68
	}
	if m.ShadowTotalConlyDeposited != 0 {
		i = encodeVarintPairPool(dAtA, i, uint64(m.ShadowTotalConlyDeposited))
		i--
		dAtA[i] = 0x60
	}
	if m.ShadowTotalNormalDeposited != 0 {
		i = encodeVarintPairPool(dAtA, i, uint64(m.ShadowTotalNormalDeposited))
		i--
		dAtA[i] = 0x58
	}
	if len(m.ShadowLpCoinDenom) > 0 {
		i -= len(m.ShadowLpCoinDenom)
		copy(dAtA[i:], m.ShadowLpCoinDenom)
		i = encodeVarintPairPool(dAtA, i, uint64(len(m.ShadowLpCoinDenom)))
		i--
		dAtA[i] = 0x52
	}
	{
		size, err := m.ShadowLiquidity.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPairPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if m.AssetTotalBorrowed != 0 {
		i = encodeVarintPairPool(dAtA, i, uint64(m.AssetTotalBorrowed))
		i--
		dAtA[i] = 0x40
	}
	if m.AssetTotalConlyDeposited != 0 {
		i = encodeVarintPairPool(dAtA, i, uint64(m.AssetTotalConlyDeposited))
		i--
		dAtA[i] = 0x38
	}
	if m.AssetTotalNormalDeposited != 0 {
		i = encodeVarintPairPool(dAtA, i, uint64(m.AssetTotalNormalDeposited))
		i--
		dAtA[i] = 0x30
	}
	if len(m.AssetLpCoinDenom) > 0 {
		i -= len(m.AssetLpCoinDenom)
		copy(dAtA[i:], m.AssetLpCoinDenom)
		i = encodeVarintPairPool(dAtA, i, uint64(len(m.AssetLpCoinDenom)))
		i--
		dAtA[i] = 0x2a
	}
	{
		size, err := m.AssetLiquidity.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPairPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.PoolId != 0 {
		i = encodeVarintPairPool(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintPairPool(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintPairPool(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPairPool(dAtA []byte, offset int, v uint64) int {
	offset -= sovPairPool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PairPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovPairPool(uint64(m.Id))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovPairPool(uint64(l))
	}
	if m.PoolId != 0 {
		n += 1 + sovPairPool(uint64(m.PoolId))
	}
	l = m.AssetLiquidity.Size()
	n += 1 + l + sovPairPool(uint64(l))
	l = len(m.AssetLpCoinDenom)
	if l > 0 {
		n += 1 + l + sovPairPool(uint64(l))
	}
	if m.AssetTotalNormalDeposited != 0 {
		n += 1 + sovPairPool(uint64(m.AssetTotalNormalDeposited))
	}
	if m.AssetTotalConlyDeposited != 0 {
		n += 1 + sovPairPool(uint64(m.AssetTotalConlyDeposited))
	}
	if m.AssetTotalBorrowed != 0 {
		n += 1 + sovPairPool(uint64(m.AssetTotalBorrowed))
	}
	l = m.ShadowLiquidity.Size()
	n += 1 + l + sovPairPool(uint64(l))
	l = len(m.ShadowLpCoinDenom)
	if l > 0 {
		n += 1 + l + sovPairPool(uint64(l))
	}
	if m.ShadowTotalNormalDeposited != 0 {
		n += 1 + sovPairPool(uint64(m.ShadowTotalNormalDeposited))
	}
	if m.ShadowTotalConlyDeposited != 0 {
		n += 1 + sovPairPool(uint64(m.ShadowTotalConlyDeposited))
	}
	if m.ShadowTotalBorrowed != 0 {
		n += 1 + sovPairPool(uint64(m.ShadowTotalBorrowed))
	}
	if m.LastUpdated != 0 {
		n += 1 + sovPairPool(uint64(m.LastUpdated))
	}
	return n
}

func sovPairPool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPairPool(x uint64) (n int) {
	return sovPairPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PairPool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPairPool
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
			return fmt.Errorf("proto: PairPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PairPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPairPool
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
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPairPool
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
				return ErrInvalidLengthPairPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPairPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPairPool
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetLiquidity", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPairPool
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
				return ErrInvalidLengthPairPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPairPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AssetLiquidity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetLpCoinDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPairPool
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
				return ErrInvalidLengthPairPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPairPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetLpCoinDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetTotalNormalDeposited", wireType)
			}
			m.AssetTotalNormalDeposited = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPairPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetTotalNormalDeposited |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetTotalConlyDeposited", wireType)
			}
			m.AssetTotalConlyDeposited = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPairPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetTotalConlyDeposited |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetTotalBorrowed", wireType)
			}
			m.AssetTotalBorrowed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPairPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetTotalBorrowed |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShadowLiquidity", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPairPool
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
				return ErrInvalidLengthPairPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPairPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ShadowLiquidity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShadowLpCoinDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPairPool
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
				return ErrInvalidLengthPairPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPairPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ShadowLpCoinDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShadowTotalNormalDeposited", wireType)
			}
			m.ShadowTotalNormalDeposited = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPairPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ShadowTotalNormalDeposited |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShadowTotalConlyDeposited", wireType)
			}
			m.ShadowTotalConlyDeposited = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPairPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ShadowTotalConlyDeposited |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShadowTotalBorrowed", wireType)
			}
			m.ShadowTotalBorrowed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPairPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ShadowTotalBorrowed |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastUpdated", wireType)
			}
			m.LastUpdated = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPairPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastUpdated |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPairPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPairPool
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
func skipPairPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPairPool
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
					return 0, ErrIntOverflowPairPool
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
					return 0, ErrIntOverflowPairPool
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
				return 0, ErrInvalidLengthPairPool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPairPool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPairPool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPairPool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPairPool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPairPool = fmt.Errorf("proto: unexpected end of group")
)

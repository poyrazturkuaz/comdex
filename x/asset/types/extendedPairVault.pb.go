// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/asset/v1beta1/extendedPairVault.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_gogo_protobuf_types "github.com/cosmos/gogoproto/types"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ExtendedPairVault struct {
	Id                  uint64                                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AppId               uint64                                 `protobuf:"varint,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty" yaml:"app_id"`
	PairId              uint64                                 `protobuf:"varint,3,opt,name=pair_id,json=pairId,proto3" json:"pair_id,omitempty" yaml:"pair_id"`
	StabilityFee        github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=stability_fee,json=stabilityFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"stability_fee" yaml:"stability_fee"`
	ClosingFee          github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=closing_fee,json=closingFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"closing_fee" yaml:"closing_fee"`
	LiquidationPenalty  github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=liquidation_penalty,json=liquidationPenalty,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"liquidation_penalty" yaml:"liquidation_penalty"`
	DrawDownFee         github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,7,opt,name=draw_down_fee,json=drawDownFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"draw_down_fee" yaml:"draw_down_fee"`
	IsVaultActive       bool                                   `protobuf:"varint,8,opt,name=is_vault_active,json=isVaultActive,proto3" json:"is_vault_active,omitempty" yaml:"active_flag"`
	DebtCeiling         github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,9,opt,name=debt_ceiling,json=debtCeiling,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"debt_ceiling" yaml:"debt_ceiling"`
	DebtFloor           github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,10,opt,name=debt_floor,json=debtFloor,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"debt_floor" yaml:"debt_floor"`
	IsStableMintVault   bool                                   `protobuf:"varint,11,opt,name=is_stable_mint_vault,json=isStableMintVault,proto3" json:"is_stable_mint_vault,omitempty" yaml:"is_stable_mint_vault"`
	MinCr               github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,12,opt,name=min_cr,json=minCr,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"min_cr" yaml:"min_cr"`
	PairName            string                                 `protobuf:"bytes,13,opt,name=pair_name,json=pairName,proto3" json:"pair_name,omitempty" yaml:"pair_name"`
	AssetOutOraclePrice bool                                   `protobuf:"varint,14,opt,name=asset_out_oracle_price,json=assetOutOraclePrice,proto3" json:"asset_out_oracle_price,omitempty" yaml:"asset_out_oracle_price"`
	AssetOutPrice       uint64                                 `protobuf:"varint,15,opt,name=asset_out_price,json=assetOutPrice,proto3" json:"asset_out_price,omitempty" yaml:"asset_out_price"`
	MinUsdValueLeft     uint64                                 `protobuf:"varint,16,opt,name=min_usd_value_left,json=minUsdValueLeft,proto3" json:"min_usd_value_left,omitempty" yaml:"min_usd_value_left"`
	BlockHeight         int64                                  `protobuf:"varint,17,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty" yaml:"block_height"`
	BlockTime           time.Time                              `protobuf:"bytes,18,opt,name=block_time,json=blockTime,proto3,stdtime" json:"block_time" yaml:"block_time"`
}

func (m *ExtendedPairVault) Reset()         { *m = ExtendedPairVault{} }
func (m *ExtendedPairVault) String() string { return proto.CompactTextString(m) }
func (*ExtendedPairVault) ProtoMessage()    {}
func (*ExtendedPairVault) Descriptor() ([]byte, []int) {
	return fileDescriptor_23dd38fcddb231cd, []int{0}
}
func (m *ExtendedPairVault) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExtendedPairVault) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExtendedPairVault.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExtendedPairVault) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtendedPairVault.Merge(m, src)
}
func (m *ExtendedPairVault) XXX_Size() int {
	return m.Size()
}
func (m *ExtendedPairVault) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtendedPairVault.DiscardUnknown(m)
}

var xxx_messageInfo_ExtendedPairVault proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ExtendedPairVault)(nil), "comdex.asset.v1beta1.ExtendedPairVault")
}

func init() {
	proto.RegisterFile("comdex/asset/v1beta1/extendedPairVault.proto", fileDescriptor_23dd38fcddb231cd)
}

var fileDescriptor_23dd38fcddb231cd = []byte{
	// 820 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xc1, 0x72, 0xe3, 0x34,
	0x18, 0x8e, 0xbb, 0xdb, 0x6c, 0xa3, 0x34, 0xed, 0x46, 0xcd, 0x14, 0x13, 0xa6, 0x71, 0xd0, 0x81,
	0xc9, 0x0c, 0x6c, 0x3c, 0x85, 0xdb, 0x1e, 0x60, 0x48, 0xbb, 0x1d, 0xca, 0x2c, 0x6c, 0x11, 0xd0,
	0x61, 0xb8, 0x68, 0x64, 0x5b, 0x71, 0x45, 0x65, 0xcb, 0xd8, 0x4a, 0xbb, 0x3d, 0xf0, 0x0e, 0xfb,
	0x0c, 0x9c, 0x78, 0x94, 0x1e, 0xf7, 0xc8, 0x70, 0x30, 0xd0, 0xbe, 0x81, 0x9f, 0x80, 0x91, 0xe4,
	0x92, 0xa4, 0xd3, 0x4b, 0x66, 0x4f, 0xd1, 0xff, 0xfd, 0x5f, 0xbe, 0xef, 0x97, 0xf4, 0xeb, 0x37,
	0xf8, 0x24, 0x94, 0x49, 0xc4, 0x5e, 0xfb, 0xb4, 0x28, 0x98, 0xf2, 0x2f, 0xf6, 0x03, 0xa6, 0xe8,
	0xbe, 0xcf, 0x5e, 0x2b, 0x96, 0x46, 0x2c, 0x3a, 0xa1, 0x3c, 0x3f, 0xa5, 0x33, 0xa1, 0xc6, 0x59,
	0x2e, 0x95, 0x84, 0x3d, 0xcb, 0x1e, 0x1b, 0xf6, 0xb8, 0x66, 0xf7, 0x7b, 0xb1, 0x8c, 0xa5, 0x21,
	0xf8, 0x7a, 0x65, 0xb9, 0x7d, 0x2f, 0x96, 0x32, 0x16, 0xcc, 0x37, 0x51, 0x30, 0x9b, 0xfa, 0x8a,
	0x27, 0xac, 0x50, 0x34, 0xc9, 0x2c, 0x01, 0xfd, 0xde, 0x06, 0xdd, 0x17, 0xf7, 0x8d, 0xe0, 0x16,
	0x58, 0xe3, 0x91, 0xeb, 0x0c, 0x9d, 0xd1, 0x63, 0xbc, 0xc6, 0x23, 0x38, 0x02, 0x4d, 0x9a, 0x65,
	0x84, 0x47, 0xee, 0x9a, 0xc6, 0x26, 0xdd, 0xaa, 0xf4, 0x3a, 0x57, 0x34, 0x11, 0xcf, 0x91, 0xc5,
	0x11, 0x5e, 0xa7, 0x59, 0x76, 0x1c, 0xc1, 0x8f, 0xc1, 0x93, 0x8c, 0xf2, 0x5c, 0x53, 0x1f, 0x19,
	0x2a, 0xac, 0x4a, 0x6f, 0xcb, 0x52, 0xeb, 0x04, 0xc2, 0x4d, 0xbd, 0x3a, 0x8e, 0xe0, 0x39, 0xe8,
	0x14, 0x8a, 0x06, 0x5c, 0x70, 0x75, 0x45, 0xa6, 0x8c, 0xb9, 0x8f, 0x87, 0xce, 0xa8, 0x35, 0x39,
	0xba, 0x2e, 0xbd, 0xc6, 0x5f, 0xa5, 0xf7, 0x51, 0xcc, 0xd5, 0xd9, 0x2c, 0x18, 0x87, 0x32, 0xf1,
	0x43, 0x59, 0x24, 0xb2, 0xa8, 0x7f, 0x9e, 0x15, 0xd1, 0xb9, 0xaf, 0xae, 0x32, 0x56, 0x8c, 0x0f,
	0x59, 0x58, 0x95, 0x5e, 0xcf, 0x1a, 0x2c, 0x89, 0x21, 0xbc, 0xf9, 0x7f, 0x7c, 0xc4, 0x18, 0x64,
	0xa0, 0x1d, 0x0a, 0x59, 0xf0, 0x34, 0x36, 0x56, 0xeb, 0xc6, 0xea, 0x70, 0x65, 0x2b, 0x68, 0xad,
	0x16, 0xa4, 0x10, 0x06, 0x75, 0xa4, 0x6d, 0x7e, 0x03, 0x3b, 0x82, 0xff, 0x3a, 0xe3, 0x11, 0x55,
	0x5c, 0xa6, 0x24, 0x63, 0x29, 0x15, 0xea, 0xca, 0x6d, 0x1a, 0xbb, 0x97, 0x2b, 0xdb, 0xf5, 0xad,
	0xdd, 0x03, 0x92, 0x08, 0xc3, 0x05, 0xf4, 0xc4, 0x82, 0xf0, 0x17, 0xd0, 0x89, 0x72, 0x7a, 0x49,
	0x22, 0x79, 0x99, 0x9a, 0x7d, 0x3e, 0x79, 0xb7, 0x23, 0x5d, 0x12, 0x43, 0xb8, 0xad, 0xe3, 0x43,
	0x79, 0x99, 0xea, 0xad, 0x7e, 0x0e, 0xb6, 0x79, 0x41, 0x2e, 0x74, 0xc7, 0x10, 0x1a, 0x2a, 0x7e,
	0xc1, 0xdc, 0x8d, 0xa1, 0x33, 0xda, 0x98, 0xec, 0xce, 0xcf, 0xc9, 0xe2, 0x64, 0x2a, 0x68, 0x8c,
	0x70, 0x87, 0x17, 0xa6, 0xbf, 0xbe, 0x34, 0x20, 0x3c, 0x03, 0x9b, 0x11, 0x0b, 0x14, 0x09, 0x19,
	0x17, 0x3c, 0x8d, 0xdd, 0x96, 0x29, 0xf5, 0xc5, 0x0a, 0xa5, 0x1e, 0xa7, 0xaa, 0x2a, 0xbd, 0x9d,
	0xba, 0xd4, 0x05, 0x2d, 0x5d, 0x29, 0x0b, 0xd4, 0x81, 0x8d, 0x60, 0x00, 0x80, 0xc9, 0x4e, 0x85,
	0x94, 0xb9, 0x0b, 0x8c, 0xcf, 0xc1, 0xca, 0x3e, 0xdd, 0x05, 0x1f, 0xa3, 0x84, 0x70, 0x4b, 0x07,
	0x47, 0x7a, 0x0d, 0x4f, 0x40, 0x8f, 0x17, 0x44, 0xb7, 0x9c, 0x60, 0x24, 0xe1, 0xa9, 0xb2, 0x27,
	0xe3, 0xb6, 0xcd, 0x91, 0x78, 0x55, 0xe9, 0x7d, 0x60, 0xff, 0xff, 0x10, 0x0b, 0xe1, 0x2e, 0x2f,
	0xbe, 0x37, 0xe8, 0x37, 0x3c, 0x55, 0xf6, 0x15, 0x9e, 0x82, 0x66, 0xc2, 0x53, 0x12, 0xe6, 0xee,
	0xa6, 0xa9, 0xf8, 0x8b, 0x95, 0x2f, 0xb1, 0x7e, 0xa3, 0x56, 0x05, 0xe1, 0xf5, 0x84, 0xa7, 0x07,
	0x39, 0xdc, 0x07, 0x2d, 0xf3, 0x14, 0x53, 0x9a, 0x30, 0xb7, 0x63, 0xa4, 0x7b, 0x55, 0xe9, 0x3d,
	0x5d, 0x78, 0xa5, 0x3a, 0x85, 0xf0, 0x86, 0x5e, 0x7f, 0x4b, 0x13, 0x06, 0x4f, 0xc1, 0xae, 0x19,
	0x37, 0x44, 0xce, 0x14, 0x91, 0x39, 0x0d, 0x05, 0x23, 0x59, 0xce, 0x43, 0xe6, 0x6e, 0x99, 0xed,
	0x7d, 0x58, 0x95, 0xde, 0x5e, 0x7d, 0xe3, 0x0f, 0xf2, 0x10, 0xde, 0x31, 0x89, 0x57, 0x33, 0xf5,
	0xca, 0xc0, 0x27, 0x1a, 0x85, 0x13, 0xb0, 0x3d, 0xe7, 0x5b, 0xc1, 0x6d, 0x33, 0x36, 0xfa, 0x55,
	0xe9, 0xed, 0xde, 0x17, 0xac, 0x95, 0x3a, 0x77, 0x4a, 0x56, 0xe3, 0x6b, 0x00, 0xf5, 0x06, 0x67,
	0x45, 0x44, 0x2e, 0xa8, 0x98, 0x31, 0x22, 0xd8, 0x54, 0xb9, 0x4f, 0x8d, 0xcc, 0x5e, 0x55, 0x7a,
	0xef, 0xcf, 0x0f, 0x61, 0x99, 0x83, 0xf0, 0x76, 0xc2, 0xd3, 0x1f, 0x8b, 0xe8, 0x54, 0x43, 0x2f,
	0xd9, 0x54, 0xc1, 0xe7, 0x60, 0x33, 0x10, 0x32, 0x3c, 0x27, 0x67, 0x8c, 0xc7, 0x67, 0xca, 0xed,
	0x0e, 0x9d, 0xd1, 0xa3, 0xc9, 0x7b, 0xf3, 0x26, 0x5b, 0xcc, 0x22, 0xdc, 0x36, 0xe1, 0x57, 0x26,
	0x82, 0x3f, 0x01, 0x60, 0xb3, 0x7a, 0xc6, 0xba, 0x70, 0xe8, 0x8c, 0xda, 0x9f, 0xf6, 0xc7, 0x76,
	0x00, 0x8f, 0xef, 0x06, 0xf0, 0xf8, 0x87, 0xbb, 0x01, 0x3c, 0xd9, 0xd3, 0xd7, 0x39, 0x6f, 0xab,
	0xf9, 0x7f, 0xd1, 0x9b, 0xbf, 0x3d, 0x07, 0xb7, 0x0c, 0xa0, 0xe9, 0x93, 0xef, 0xae, 0xff, 0x1d,
	0x34, 0xfe, 0xb8, 0x19, 0x34, 0xae, 0x6f, 0x06, 0xce, 0xdb, 0x9b, 0x81, 0xf3, 0xcf, 0xcd, 0xc0,
	0x79, 0x73, 0x3b, 0x68, 0xbc, 0xbd, 0x1d, 0x34, 0xfe, 0xbc, 0x1d, 0x34, 0x7e, 0xf6, 0x97, 0x5a,
	0x42, 0x7f, 0x1e, 0x9e, 0xc9, 0xe9, 0x94, 0x87, 0x9c, 0x8a, 0x3a, 0xf6, 0xef, 0x3e, 0x2f, 0xa6,
	0x3f, 0x82, 0xa6, 0x29, 0xe8, 0xb3, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x98, 0x15, 0x0f,
	0x7b, 0x06, 0x00, 0x00,
}

func (m *ExtendedPairVault) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExtendedPairVault) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExtendedPairVault) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.BlockTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.BlockTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintExtendedPairVault(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1
	i--
	dAtA[i] = 0x92
	if m.BlockHeight != 0 {
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x88
	}
	if m.MinUsdValueLeft != 0 {
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(m.MinUsdValueLeft))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if m.AssetOutPrice != 0 {
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(m.AssetOutPrice))
		i--
		dAtA[i] = 0x78
	}
	if m.AssetOutOraclePrice {
		i--
		if m.AssetOutOraclePrice {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x70
	}
	if len(m.PairName) > 0 {
		i -= len(m.PairName)
		copy(dAtA[i:], m.PairName)
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(len(m.PairName)))
		i--
		dAtA[i] = 0x6a
	}
	{
		size := m.MinCr.Size()
		i -= size
		if _, err := m.MinCr.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	if m.IsStableMintVault {
		i--
		if m.IsStableMintVault {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x58
	}
	{
		size := m.DebtFloor.Size()
		i -= size
		if _, err := m.DebtFloor.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	{
		size := m.DebtCeiling.Size()
		i -= size
		if _, err := m.DebtCeiling.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if m.IsVaultActive {
		i--
		if m.IsVaultActive {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	{
		size := m.DrawDownFee.Size()
		i -= size
		if _, err := m.DrawDownFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.LiquidationPenalty.Size()
		i -= size
		if _, err := m.LiquidationPenalty.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.ClosingFee.Size()
		i -= size
		if _, err := m.ClosingFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.StabilityFee.Size()
		i -= size
		if _, err := m.StabilityFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.PairId != 0 {
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(m.PairId))
		i--
		dAtA[i] = 0x18
	}
	if m.AppId != 0 {
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(m.AppId))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintExtendedPairVault(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintExtendedPairVault(dAtA []byte, offset int, v uint64) int {
	offset -= sovExtendedPairVault(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ExtendedPairVault) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovExtendedPairVault(uint64(m.Id))
	}
	if m.AppId != 0 {
		n += 1 + sovExtendedPairVault(uint64(m.AppId))
	}
	if m.PairId != 0 {
		n += 1 + sovExtendedPairVault(uint64(m.PairId))
	}
	l = m.StabilityFee.Size()
	n += 1 + l + sovExtendedPairVault(uint64(l))
	l = m.ClosingFee.Size()
	n += 1 + l + sovExtendedPairVault(uint64(l))
	l = m.LiquidationPenalty.Size()
	n += 1 + l + sovExtendedPairVault(uint64(l))
	l = m.DrawDownFee.Size()
	n += 1 + l + sovExtendedPairVault(uint64(l))
	if m.IsVaultActive {
		n += 2
	}
	l = m.DebtCeiling.Size()
	n += 1 + l + sovExtendedPairVault(uint64(l))
	l = m.DebtFloor.Size()
	n += 1 + l + sovExtendedPairVault(uint64(l))
	if m.IsStableMintVault {
		n += 2
	}
	l = m.MinCr.Size()
	n += 1 + l + sovExtendedPairVault(uint64(l))
	l = len(m.PairName)
	if l > 0 {
		n += 1 + l + sovExtendedPairVault(uint64(l))
	}
	if m.AssetOutOraclePrice {
		n += 2
	}
	if m.AssetOutPrice != 0 {
		n += 1 + sovExtendedPairVault(uint64(m.AssetOutPrice))
	}
	if m.MinUsdValueLeft != 0 {
		n += 2 + sovExtendedPairVault(uint64(m.MinUsdValueLeft))
	}
	if m.BlockHeight != 0 {
		n += 2 + sovExtendedPairVault(uint64(m.BlockHeight))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.BlockTime)
	n += 2 + l + sovExtendedPairVault(uint64(l))
	return n
}

func sovExtendedPairVault(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozExtendedPairVault(x uint64) (n int) {
	return sovExtendedPairVault(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExtendedPairVault) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExtendedPairVault
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
			return fmt.Errorf("proto: ExtendedPairVault: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExtendedPairVault: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppId", wireType)
			}
			m.AppId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AppId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairId", wireType)
			}
			m.PairId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PairId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StabilityFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
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
				return ErrInvalidLengthExtendedPairVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtendedPairVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StabilityFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClosingFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
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
				return ErrInvalidLengthExtendedPairVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtendedPairVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ClosingFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidationPenalty", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
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
				return ErrInvalidLengthExtendedPairVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtendedPairVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LiquidationPenalty.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DrawDownFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
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
				return ErrInvalidLengthExtendedPairVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtendedPairVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DrawDownFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsVaultActive", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsVaultActive = bool(v != 0)
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DebtCeiling", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
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
				return ErrInvalidLengthExtendedPairVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtendedPairVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DebtCeiling.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DebtFloor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
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
				return ErrInvalidLengthExtendedPairVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtendedPairVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DebtFloor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsStableMintVault", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsStableMintVault = bool(v != 0)
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinCr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
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
				return ErrInvalidLengthExtendedPairVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtendedPairVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinCr.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
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
				return ErrInvalidLengthExtendedPairVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtendedPairVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PairName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetOutOraclePrice", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AssetOutOraclePrice = bool(v != 0)
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetOutPrice", wireType)
			}
			m.AssetOutPrice = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetOutPrice |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinUsdValueLeft", wireType)
			}
			m.MinUsdValueLeft = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinUsdValueLeft |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 17:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 18:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedPairVault
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
				return ErrInvalidLengthExtendedPairVault
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExtendedPairVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.BlockTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExtendedPairVault(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExtendedPairVault
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
func skipExtendedPairVault(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExtendedPairVault
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
					return 0, ErrIntOverflowExtendedPairVault
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
					return 0, ErrIntOverflowExtendedPairVault
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
				return 0, ErrInvalidLengthExtendedPairVault
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupExtendedPairVault
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthExtendedPairVault
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthExtendedPairVault        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExtendedPairVault          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupExtendedPairVault = fmt.Errorf("proto: unexpected end of group")
)

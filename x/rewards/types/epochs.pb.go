// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/rewards/v1beta1/epochs.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_gogo_protobuf_types "github.com/cosmos/gogoproto/types"
	_ "github.com/golang/protobuf/ptypes/duration"
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

type EpochInfo struct {
	StartTime               time.Time     `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time" yaml:"start_time"`
	Duration                time.Duration `protobuf:"bytes,3,opt,name=duration,proto3,stdduration" json:"duration,omitempty" yaml:"duration"`
	CurrentEpoch            int64         `protobuf:"varint,4,opt,name=current_epoch,json=currentEpoch,proto3" json:"current_epoch,omitempty"`
	CurrentEpochStartTime   time.Time     `protobuf:"bytes,5,opt,name=current_epoch_start_time,json=currentEpochStartTime,proto3,stdtime" json:"current_epoch_start_time" yaml:"current_epoch_start_time"`
	CurrentEpochStartHeight int64         `protobuf:"varint,6,opt,name=current_epoch_start_height,json=currentEpochStartHeight,proto3" json:"current_epoch_start_height,omitempty"`
}

func (m *EpochInfo) Reset()         { *m = EpochInfo{} }
func (m *EpochInfo) String() string { return proto.CompactTextString(m) }
func (*EpochInfo) ProtoMessage()    {}
func (*EpochInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f11e4baf2d58b38, []int{0}
}
func (m *EpochInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EpochInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EpochInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EpochInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EpochInfo.Merge(m, src)
}
func (m *EpochInfo) XXX_Size() int {
	return m.Size()
}
func (m *EpochInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_EpochInfo.DiscardUnknown(m)
}

var xxx_messageInfo_EpochInfo proto.InternalMessageInfo

func (m *EpochInfo) GetStartTime() time.Time {
	if m != nil {
		return m.StartTime
	}
	return time.Time{}
}

func (m *EpochInfo) GetDuration() time.Duration {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *EpochInfo) GetCurrentEpoch() int64 {
	if m != nil {
		return m.CurrentEpoch
	}
	return 0
}

func (m *EpochInfo) GetCurrentEpochStartTime() time.Time {
	if m != nil {
		return m.CurrentEpochStartTime
	}
	return time.Time{}
}

func (m *EpochInfo) GetCurrentEpochStartHeight() int64 {
	if m != nil {
		return m.CurrentEpochStartHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*EpochInfo)(nil), "comdex.rewards.v1beta1.EpochInfo")
}

func init() {
	proto.RegisterFile("comdex/rewards/v1beta1/epochs.proto", fileDescriptor_6f11e4baf2d58b38)
}

var fileDescriptor_6f11e4baf2d58b38 = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xbd, 0x4e, 0xfb, 0x30,
	0x14, 0xc5, 0x93, 0x7f, 0xff, 0x54, 0xd4, 0x80, 0x10, 0x11, 0x1f, 0x21, 0x12, 0x49, 0x95, 0x2e,
	0x95, 0x80, 0x58, 0x81, 0x0d, 0xb6, 0x0a, 0x24, 0x10, 0x5b, 0x61, 0x40, 0x2c, 0x55, 0x92, 0x3a,
	0x1f, 0x52, 0x53, 0x47, 0x89, 0x03, 0x74, 0xe3, 0x11, 0x3a, 0xf2, 0x20, 0x3c, 0x44, 0xc7, 0x8e,
	0x4c, 0x01, 0xb5, 0x1b, 0x63, 0x9f, 0x00, 0xc5, 0x76, 0xda, 0x42, 0x41, 0x6c, 0xb6, 0xef, 0xef,
	0x9e, 0xe3, 0x73, 0x75, 0x41, 0xcd, 0xc1, 0x61, 0x1b, 0x3d, 0xc2, 0x18, 0x3d, 0x58, 0x71, 0x3b,
	0x81, 0xf7, 0xa6, 0x8d, 0x88, 0x65, 0x42, 0x14, 0x61, 0xc7, 0x4f, 0x8c, 0x28, 0xc6, 0x04, 0x4b,
	0xdb, 0x0c, 0x32, 0x38, 0x64, 0x70, 0x48, 0xd9, 0xf4, 0xb0, 0x87, 0x29, 0x02, 0xf3, 0x13, 0xa3,
	0x15, 0xcd, 0xc3, 0xd8, 0xeb, 0x20, 0x48, 0x6f, 0x76, 0xea, 0x42, 0x12, 0x84, 0x28, 0x21, 0x56,
	0x18, 0x71, 0x40, 0xfd, 0x0e, 0xb4, 0xd3, 0xd8, 0x22, 0x01, 0xee, 0xb2, 0xba, 0xfe, 0x52, 0x02,
	0x95, 0xf3, 0xdc, 0xff, 0xb2, 0xeb, 0x62, 0xe9, 0x16, 0x80, 0x84, 0x58, 0x31, 0x69, 0xe5, 0x32,
	0xf2, 0xbf, 0xaa, 0x58, 0x5f, 0x39, 0x52, 0x0c, 0x26, 0x61, 0x14, 0x12, 0xc6, 0x4d, 0xe1, 0xd1,
	0xd8, 0x1b, 0x64, 0x9a, 0x30, 0xc9, 0xb4, 0x8d, 0x9e, 0x15, 0x76, 0x4e, 0xf4, 0x59, 0xaf, 0xde,
	0x7f, 0xd3, 0xc4, 0x66, 0x85, 0x3e, 0xe4, 0xb8, 0xe4, 0x83, 0xe5, 0xc2, 0x59, 0x2e, 0x51, 0xdd,
	0xdd, 0x05, 0xdd, 0x33, 0x0e, 0x34, 0xcc, 0x5c, 0xf6, 0x23, 0xd3, 0xa4, 0xa2, 0xe5, 0x00, 0x87,
	0x01, 0x41, 0x61, 0x44, 0x7a, 0x93, 0x4c, 0x5b, 0x67, 0x66, 0x45, 0x4d, 0x7f, 0xce, 0xad, 0xa6,
	0xea, 0x52, 0x0d, 0xac, 0x39, 0x69, 0x1c, 0xa3, 0x2e, 0x69, 0xd1, 0xc1, 0xca, 0xff, 0xab, 0x62,
	0xbd, 0xd4, 0x5c, 0xe5, 0x8f, 0x34, 0xac, 0xf4, 0x24, 0x02, 0xf9, 0x0b, 0xd5, 0x9a, 0xcb, 0xbd,
	0xf4, 0x67, 0xee, 0x7d, 0x9e, 0x5b, 0x63, 0x5f, 0xf9, 0x4d, 0x89, 0x4d, 0x61, 0x6b, 0xde, 0xf9,
	0x7a, 0x3a, 0x91, 0x53, 0xa0, 0xfc, 0xd4, 0xe7, 0xa3, 0xc0, 0xf3, 0x89, 0x5c, 0xa6, 0x9f, 0xde,
	0x59, 0x68, 0xbd, 0xa0, 0xe5, 0xc6, 0xd5, 0x60, 0xa4, 0x8a, 0xc3, 0x91, 0x2a, 0xbe, 0x8f, 0x54,
	0xb1, 0x3f, 0x56, 0x85, 0xe1, 0x58, 0x15, 0x5e, 0xc7, 0xaa, 0x70, 0x67, 0x7a, 0x01, 0xf1, 0x53,
	0xdb, 0x70, 0x70, 0x08, 0xd9, 0x2a, 0x1d, 0x62, 0xd7, 0x0d, 0x9c, 0xc0, 0xea, 0xf0, 0x3b, 0x9c,
	0x6d, 0x20, 0xe9, 0x45, 0x28, 0xb1, 0xcb, 0x34, 0xe1, 0xf1, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xd9, 0x4a, 0x93, 0xc1, 0xa0, 0x02, 0x00, 0x00,
}

func (m *EpochInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EpochInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EpochInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CurrentEpochStartHeight != 0 {
		i = encodeVarintEpochs(dAtA, i, uint64(m.CurrentEpochStartHeight))
		i--
		dAtA[i] = 0x30
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CurrentEpochStartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CurrentEpochStartTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintEpochs(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x2a
	if m.CurrentEpoch != 0 {
		i = encodeVarintEpochs(dAtA, i, uint64(m.CurrentEpoch))
		i--
		dAtA[i] = 0x20
	}
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Duration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintEpochs(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintEpochs(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}

func encodeVarintEpochs(dAtA []byte, offset int, v uint64) int {
	offset -= sovEpochs(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EpochInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovEpochs(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration)
	n += 1 + l + sovEpochs(uint64(l))
	if m.CurrentEpoch != 0 {
		n += 1 + sovEpochs(uint64(m.CurrentEpoch))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CurrentEpochStartTime)
	n += 1 + l + sovEpochs(uint64(l))
	if m.CurrentEpochStartHeight != 0 {
		n += 1 + sovEpochs(uint64(m.CurrentEpochStartHeight))
	}
	return n
}

func sovEpochs(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEpochs(x uint64) (n int) {
	return sovEpochs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EpochInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEpochs
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
			return fmt.Errorf("proto: EpochInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EpochInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochs
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
				return ErrInvalidLengthEpochs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEpochs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochs
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
				return ErrInvalidLengthEpochs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEpochs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Duration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentEpoch", wireType)
			}
			m.CurrentEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentEpoch |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentEpochStartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochs
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
				return ErrInvalidLengthEpochs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEpochs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CurrentEpochStartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentEpochStartHeight", wireType)
			}
			m.CurrentEpochStartHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentEpochStartHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEpochs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEpochs
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
func skipEpochs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEpochs
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
					return 0, ErrIntOverflowEpochs
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
					return 0, ErrIntOverflowEpochs
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
				return 0, ErrInvalidLengthEpochs
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEpochs
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEpochs
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEpochs        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEpochs          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEpochs = fmt.Errorf("proto: unexpected end of group")
)

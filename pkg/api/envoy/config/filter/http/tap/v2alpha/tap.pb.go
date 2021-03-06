// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/http/tap/v2alpha/tap.proto

package envoy_config_filter_http_tap_v2alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v2alpha "github.com/datawire/ambassador/pkg/api/envoy/config/common/tap/v2alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// Top level configuration for the tap filter.
type Tap struct {
	// Common configuration for the HTTP tap filter.
	CommonConfig         *v2alpha.CommonExtensionConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *Tap) Reset()         { *m = Tap{} }
func (m *Tap) String() string { return proto.CompactTextString(m) }
func (*Tap) ProtoMessage()    {}
func (*Tap) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee77d938a401b876, []int{0}
}
func (m *Tap) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Tap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Tap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Tap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tap.Merge(m, src)
}
func (m *Tap) XXX_Size() int {
	return m.Size()
}
func (m *Tap) XXX_DiscardUnknown() {
	xxx_messageInfo_Tap.DiscardUnknown(m)
}

var xxx_messageInfo_Tap proto.InternalMessageInfo

func (m *Tap) GetCommonConfig() *v2alpha.CommonExtensionConfig {
	if m != nil {
		return m.CommonConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*Tap)(nil), "envoy.config.filter.http.tap.v2alpha.Tap")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/tap/v2alpha/tap.proto", fileDescriptor_ee77d938a401b876)
}

var fileDescriptor_ee77d938a401b876 = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4b, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2,
	0xcf, 0x28, 0x29, 0x29, 0xd0, 0x2f, 0x49, 0x2c, 0xd0, 0x2f, 0x33, 0x4a, 0xcc, 0x29, 0xc8, 0x48,
	0x04, 0xb1, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x54, 0xc0, 0xea, 0xf5, 0x20, 0xea, 0xf5,
	0x20, 0xea, 0xf5, 0x40, 0xea, 0xf5, 0x40, 0x6a, 0xa0, 0xea, 0xa5, 0x74, 0x50, 0x4c, 0x4d, 0xce,
	0xcf, 0xcd, 0xcd, 0xcf, 0x43, 0x31, 0x10, 0x22, 0x04, 0x31, 0x53, 0x4a, 0xae, 0x34, 0xa5, 0x20,
	0x51, 0x3f, 0x31, 0x2f, 0x2f, 0xbf, 0x24, 0xb1, 0x24, 0x33, 0x3f, 0xaf, 0x58, 0x3f, 0x37, 0x33,
	0xbd, 0x28, 0xb1, 0x24, 0x15, 0x2a, 0x2f, 0x8b, 0x21, 0x5f, 0x5c, 0x92, 0x58, 0x52, 0x5a, 0x0c,
	0x95, 0x16, 0x2f, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0x49, 0xd5, 0x87, 0x31, 0x20, 0x12, 0x4a,
	0x39, 0x5c, 0xcc, 0x21, 0x89, 0x05, 0x42, 0xa9, 0x5c, 0xbc, 0x10, 0xeb, 0xe2, 0x21, 0xee, 0x91,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36, 0x32, 0xd3, 0x43, 0xf1, 0x0a, 0xd4, 0x45, 0x48, 0xbe, 0xd0,
	0x73, 0x06, 0x0b, 0xb9, 0x56, 0x94, 0xa4, 0xe6, 0x15, 0x67, 0xe6, 0xe7, 0x39, 0x83, 0x15, 0x3a,
	0x71, 0xfc, 0x72, 0x62, 0xed, 0x62, 0x64, 0x12, 0x60, 0x0c, 0xe2, 0x81, 0xe8, 0x81, 0x8a, 0x37,
	0x30, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x9f, 0x66,
	0xfc, 0xeb, 0x67, 0x55, 0x83, 0x05, 0x57, 0x2a, 0xcc, 0x8c, 0x62, 0x68, 0x90, 0x15, 0x23, 0x85,
	0x99, 0xf1, 0xae, 0x86, 0x13, 0x17, 0xd9, 0x98, 0x04, 0x18, 0xb8, 0x8c, 0x32, 0xf3, 0x21, 0x8e,
	0x2a, 0x28, 0xca, 0xaf, 0xa8, 0xd4, 0x23, 0x26, 0xa8, 0x9d, 0x38, 0x42, 0x12, 0x0b, 0x02, 0x40,
	0xde, 0x0d, 0x60, 0x4c, 0x62, 0x03, 0xfb, 0xdb, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x71, 0x4c,
	0xbc, 0x82, 0xd5, 0x01, 0x00, 0x00,
}

func (m *Tap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Tap) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Tap) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.CommonConfig != nil {
		{
			size, err := m.CommonConfig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTap(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTap(dAtA []byte, offset int, v uint64) int {
	offset -= sovTap(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Tap) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CommonConfig != nil {
		l = m.CommonConfig.Size()
		n += 1 + l + sovTap(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTap(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTap(x uint64) (n int) {
	return sovTap(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Tap) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTap
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
			return fmt.Errorf("proto: Tap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Tap: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommonConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTap
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
				return ErrInvalidLengthTap
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CommonConfig == nil {
				m.CommonConfig = &v2alpha.CommonExtensionConfig{}
			}
			if err := m.CommonConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTap
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTap
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTap(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTap
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
					return 0, ErrIntOverflowTap
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
					return 0, ErrIntOverflowTap
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
				return 0, ErrInvalidLengthTap
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTap
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTap
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTap        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTap          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTap = fmt.Errorf("proto: unexpected end of group")
)

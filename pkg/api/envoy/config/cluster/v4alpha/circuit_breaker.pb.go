// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/cluster/v4alpha/circuit_breaker.proto

package envoy_config_cluster_v4alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v4alpha "github.com/datawire/ambassador/pkg/api/envoy/config/core/v4alpha"
	v3 "github.com/datawire/ambassador/pkg/api/envoy/type/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

// :ref:`Circuit breaking<arch_overview_circuit_break>` settings can be
// specified individually for each defined priority.
type CircuitBreakers struct {
	// If multiple :ref:`Thresholds<envoy_api_msg_config.cluster.v4alpha.CircuitBreakers.Thresholds>`
	// are defined with the same :ref:`RoutingPriority<envoy_api_enum_config.core.v4alpha.RoutingPriority>`,
	// the first one in the list is used. If no Thresholds is defined for a given
	// :ref:`RoutingPriority<envoy_api_enum_config.core.v4alpha.RoutingPriority>`, the default values
	// are used.
	Thresholds           []*CircuitBreakers_Thresholds `protobuf:"bytes,1,rep,name=thresholds,proto3" json:"thresholds,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *CircuitBreakers) Reset()         { *m = CircuitBreakers{} }
func (m *CircuitBreakers) String() string { return proto.CompactTextString(m) }
func (*CircuitBreakers) ProtoMessage()    {}
func (*CircuitBreakers) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e788e4bcf428d89, []int{0}
}
func (m *CircuitBreakers) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CircuitBreakers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CircuitBreakers.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CircuitBreakers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CircuitBreakers.Merge(m, src)
}
func (m *CircuitBreakers) XXX_Size() int {
	return m.Size()
}
func (m *CircuitBreakers) XXX_DiscardUnknown() {
	xxx_messageInfo_CircuitBreakers.DiscardUnknown(m)
}

var xxx_messageInfo_CircuitBreakers proto.InternalMessageInfo

func (m *CircuitBreakers) GetThresholds() []*CircuitBreakers_Thresholds {
	if m != nil {
		return m.Thresholds
	}
	return nil
}

// A Thresholds defines CircuitBreaker settings for a
// :ref:`RoutingPriority<envoy_api_enum_config.core.v4alpha.RoutingPriority>`.
// [#next-free-field: 9]
type CircuitBreakers_Thresholds struct {
	// The :ref:`RoutingPriority<envoy_api_enum_config.core.v4alpha.RoutingPriority>`
	// the specified CircuitBreaker settings apply to.
	Priority v4alpha.RoutingPriority `protobuf:"varint,1,opt,name=priority,proto3,enum=envoy.config.core.v4alpha.RoutingPriority" json:"priority,omitempty"`
	// The maximum number of connections that Envoy will make to the upstream
	// cluster. If not specified, the default is 1024.
	MaxConnections *types.UInt32Value `protobuf:"bytes,2,opt,name=max_connections,json=maxConnections,proto3" json:"max_connections,omitempty"`
	// The maximum number of pending requests that Envoy will allow to the
	// upstream cluster. If not specified, the default is 1024.
	MaxPendingRequests *types.UInt32Value `protobuf:"bytes,3,opt,name=max_pending_requests,json=maxPendingRequests,proto3" json:"max_pending_requests,omitempty"`
	// The maximum number of parallel requests that Envoy will make to the
	// upstream cluster. If not specified, the default is 1024.
	MaxRequests *types.UInt32Value `protobuf:"bytes,4,opt,name=max_requests,json=maxRequests,proto3" json:"max_requests,omitempty"`
	// The maximum number of parallel retries that Envoy will allow to the
	// upstream cluster. If not specified, the default is 3.
	MaxRetries *types.UInt32Value `protobuf:"bytes,5,opt,name=max_retries,json=maxRetries,proto3" json:"max_retries,omitempty"`
	// Specifies a limit on concurrent retries in relation to the number of active requests. This
	// parameter is optional.
	//
	// .. note::
	//
	//    If this field is set, the retry budget will override any configured retry circuit
	//    breaker.
	RetryBudget *CircuitBreakers_Thresholds_RetryBudget `protobuf:"bytes,8,opt,name=retry_budget,json=retryBudget,proto3" json:"retry_budget,omitempty"`
	// If track_remaining is true, then stats will be published that expose
	// the number of resources remaining until the circuit breakers open. If
	// not specified, the default is false.
	//
	// .. note::
	//
	//    If a retry budget is used in lieu of the max_retries circuit breaker,
	//    the remaining retry resources remaining will not be tracked.
	TrackRemaining bool `protobuf:"varint,6,opt,name=track_remaining,json=trackRemaining,proto3" json:"track_remaining,omitempty"`
	// The maximum number of connection pools per cluster that Envoy will concurrently support at
	// once. If not specified, the default is unlimited. Set this for clusters which create a
	// large number of connection pools. See
	// :ref:`Circuit Breaking <arch_overview_circuit_break_cluster_maximum_connection_pools>` for
	// more details.
	MaxConnectionPools   *types.UInt32Value `protobuf:"bytes,7,opt,name=max_connection_pools,json=maxConnectionPools,proto3" json:"max_connection_pools,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CircuitBreakers_Thresholds) Reset()         { *m = CircuitBreakers_Thresholds{} }
func (m *CircuitBreakers_Thresholds) String() string { return proto.CompactTextString(m) }
func (*CircuitBreakers_Thresholds) ProtoMessage()    {}
func (*CircuitBreakers_Thresholds) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e788e4bcf428d89, []int{0, 0}
}
func (m *CircuitBreakers_Thresholds) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CircuitBreakers_Thresholds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CircuitBreakers_Thresholds.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CircuitBreakers_Thresholds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CircuitBreakers_Thresholds.Merge(m, src)
}
func (m *CircuitBreakers_Thresholds) XXX_Size() int {
	return m.Size()
}
func (m *CircuitBreakers_Thresholds) XXX_DiscardUnknown() {
	xxx_messageInfo_CircuitBreakers_Thresholds.DiscardUnknown(m)
}

var xxx_messageInfo_CircuitBreakers_Thresholds proto.InternalMessageInfo

func (m *CircuitBreakers_Thresholds) GetPriority() v4alpha.RoutingPriority {
	if m != nil {
		return m.Priority
	}
	return v4alpha.RoutingPriority_DEFAULT
}

func (m *CircuitBreakers_Thresholds) GetMaxConnections() *types.UInt32Value {
	if m != nil {
		return m.MaxConnections
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetMaxPendingRequests() *types.UInt32Value {
	if m != nil {
		return m.MaxPendingRequests
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetMaxRequests() *types.UInt32Value {
	if m != nil {
		return m.MaxRequests
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetMaxRetries() *types.UInt32Value {
	if m != nil {
		return m.MaxRetries
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetRetryBudget() *CircuitBreakers_Thresholds_RetryBudget {
	if m != nil {
		return m.RetryBudget
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetTrackRemaining() bool {
	if m != nil {
		return m.TrackRemaining
	}
	return false
}

func (m *CircuitBreakers_Thresholds) GetMaxConnectionPools() *types.UInt32Value {
	if m != nil {
		return m.MaxConnectionPools
	}
	return nil
}

type CircuitBreakers_Thresholds_RetryBudget struct {
	// Specifies the limit on concurrent retries as a percentage of the sum of active requests and
	// active pending requests. For example, if there are 100 active requests and the
	// budget_percent is set to 25, there may be 25 active retries.
	//
	// This parameter is optional. Defaults to 20%.
	BudgetPercent *v3.Percent `protobuf:"bytes,1,opt,name=budget_percent,json=budgetPercent,proto3" json:"budget_percent,omitempty"`
	// Specifies the minimum retry concurrency allowed for the retry budget. The limit on the
	// number of active retries may never go below this number.
	//
	// This parameter is optional. Defaults to 3.
	MinRetryConcurrency  *types.UInt32Value `protobuf:"bytes,2,opt,name=min_retry_concurrency,json=minRetryConcurrency,proto3" json:"min_retry_concurrency,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CircuitBreakers_Thresholds_RetryBudget) Reset() {
	*m = CircuitBreakers_Thresholds_RetryBudget{}
}
func (m *CircuitBreakers_Thresholds_RetryBudget) String() string { return proto.CompactTextString(m) }
func (*CircuitBreakers_Thresholds_RetryBudget) ProtoMessage()    {}
func (*CircuitBreakers_Thresholds_RetryBudget) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e788e4bcf428d89, []int{0, 0, 0}
}
func (m *CircuitBreakers_Thresholds_RetryBudget) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CircuitBreakers_Thresholds_RetryBudget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CircuitBreakers_Thresholds_RetryBudget.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CircuitBreakers_Thresholds_RetryBudget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CircuitBreakers_Thresholds_RetryBudget.Merge(m, src)
}
func (m *CircuitBreakers_Thresholds_RetryBudget) XXX_Size() int {
	return m.Size()
}
func (m *CircuitBreakers_Thresholds_RetryBudget) XXX_DiscardUnknown() {
	xxx_messageInfo_CircuitBreakers_Thresholds_RetryBudget.DiscardUnknown(m)
}

var xxx_messageInfo_CircuitBreakers_Thresholds_RetryBudget proto.InternalMessageInfo

func (m *CircuitBreakers_Thresholds_RetryBudget) GetBudgetPercent() *v3.Percent {
	if m != nil {
		return m.BudgetPercent
	}
	return nil
}

func (m *CircuitBreakers_Thresholds_RetryBudget) GetMinRetryConcurrency() *types.UInt32Value {
	if m != nil {
		return m.MinRetryConcurrency
	}
	return nil
}

func init() {
	proto.RegisterType((*CircuitBreakers)(nil), "envoy.config.cluster.v4alpha.CircuitBreakers")
	proto.RegisterType((*CircuitBreakers_Thresholds)(nil), "envoy.config.cluster.v4alpha.CircuitBreakers.Thresholds")
	proto.RegisterType((*CircuitBreakers_Thresholds_RetryBudget)(nil), "envoy.config.cluster.v4alpha.CircuitBreakers.Thresholds.RetryBudget")
}

func init() {
	proto.RegisterFile("envoy/config/cluster/v4alpha/circuit_breaker.proto", fileDescriptor_7e788e4bcf428d89)
}

var fileDescriptor_7e788e4bcf428d89 = []byte{
	// 626 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x4d, 0x6f, 0xd3, 0x30,
	0x18, 0xc7, 0x95, 0x8e, 0x95, 0xca, 0x19, 0xdd, 0x94, 0xf1, 0x12, 0x95, 0x51, 0x15, 0x84, 0xb4,
	0xb2, 0x83, 0x23, 0xb5, 0x3b, 0xc0, 0xa4, 0x01, 0xea, 0xd8, 0x81, 0x0b, 0x8a, 0x22, 0xde, 0x6e,
	0x91, 0x9b, 0x3e, 0xcb, 0xac, 0x25, 0x76, 0xb0, 0x9d, 0xd2, 0xdc, 0x10, 0x27, 0x3e, 0x02, 0xe2,
	0xa3, 0x70, 0x47, 0xda, 0x11, 0xbe, 0x01, 0x9a, 0xc4, 0x97, 0xd8, 0x09, 0x25, 0x4e, 0x9b, 0x6d,
	0xa0, 0x29, 0x70, 0xab, 0xf3, 0xf8, 0xf7, 0x4b, 0x9e, 0xbf, 0x1f, 0x17, 0x0d, 0x80, 0x4d, 0x79,
	0xe6, 0x04, 0x9c, 0x1d, 0xd0, 0xd0, 0x09, 0xa2, 0x54, 0x2a, 0x10, 0xce, 0x74, 0x9b, 0x44, 0xc9,
	0x21, 0x71, 0x02, 0x2a, 0x82, 0x94, 0x2a, 0x7f, 0x2c, 0x80, 0x1c, 0x81, 0xc0, 0x89, 0xe0, 0x8a,
	0x5b, 0x1b, 0x05, 0x83, 0x35, 0x83, 0x4b, 0x06, 0x97, 0x4c, 0xe7, 0xfe, 0x79, 0x23, 0x17, 0xb0,
	0xd0, 0x8d, 0x89, 0x04, 0xed, 0xe8, 0xdc, 0xd6, 0xbb, 0x54, 0x96, 0x80, 0x33, 0x1d, 0x3a, 0x09,
	0x88, 0x00, 0x98, 0x2a, 0x8b, 0xdd, 0x90, 0xf3, 0x30, 0x02, 0xa7, 0x58, 0x8d, 0xd3, 0x03, 0xe7,
	0xbd, 0x20, 0x49, 0x02, 0x42, 0x96, 0xf5, 0x3b, 0xe9, 0x24, 0x21, 0x0e, 0x61, 0x8c, 0x2b, 0xa2,
	0x28, 0x67, 0xd2, 0x91, 0x8a, 0xa8, 0x74, 0x5e, 0xbe, 0xfb, 0x47, 0x79, 0x0a, 0x42, 0x52, 0xce,
	0x28, 0x0b, 0xcb, 0x2d, 0xb7, 0xa6, 0x24, 0xa2, 0x13, 0xa2, 0xc0, 0x99, 0xff, 0xd0, 0x85, 0x7b,
	0x9f, 0x5b, 0x68, 0x75, 0x4f, 0x77, 0x3d, 0xd2, 0x4d, 0x4b, 0xeb, 0x2d, 0x42, 0xea, 0x50, 0x80,
	0x3c, 0xe4, 0xd1, 0x44, 0xda, 0x46, 0x6f, 0xa9, 0x6f, 0x0e, 0x1e, 0xe2, 0xcb, 0x42, 0xc0, 0x17,
	0x14, 0xf8, 0xe5, 0x82, 0xf7, 0xce, 0xb8, 0x3a, 0xa7, 0x4d, 0x84, 0xaa, 0x92, 0xe5, 0xa2, 0x56,
	0x22, 0x28, 0x17, 0x54, 0x65, 0xb6, 0xd1, 0x33, 0xfa, 0xed, 0xc1, 0xd6, 0x85, 0xd7, 0x70, 0x01,
	0x8b, 0x77, 0x78, 0x3c, 0x55, 0x94, 0x85, 0x6e, 0x49, 0x8c, 0x5a, 0xa7, 0xa3, 0xe5, 0x8f, 0x46,
	0x63, 0xcd, 0xf0, 0x16, 0x16, 0x6b, 0x1f, 0xad, 0xc6, 0x64, 0xe6, 0x07, 0x9c, 0x31, 0x08, 0x8a,
	0x2c, 0xec, 0x46, 0xcf, 0xe8, 0x9b, 0x83, 0x0d, 0xac, 0x33, 0xc6, 0xf3, 0x8c, 0xf1, 0xab, 0xe7,
	0x4c, 0x0d, 0x07, 0xaf, 0x49, 0x94, 0x82, 0xd7, 0x8e, 0xc9, 0x6c, 0xaf, 0x62, 0xac, 0x17, 0xe8,
	0x7a, 0xae, 0x49, 0x80, 0x4d, 0x28, 0x0b, 0x7d, 0x01, 0xef, 0x52, 0x90, 0x4a, 0xda, 0x4b, 0x35,
	0x5c, 0x56, 0x4c, 0x66, 0xae, 0x06, 0xbd, 0x92, 0xb3, 0x9e, 0xa0, 0x95, 0xdc, 0xb7, 0xf0, 0x5c,
	0xa9, 0xe1, 0x31, 0x63, 0x32, 0x5b, 0x08, 0x76, 0x91, 0xa9, 0x05, 0x4a, 0x50, 0x90, 0xf6, 0x72,
	0x0d, 0x1e, 0x15, 0x7c, 0xb1, 0xdf, 0x0a, 0xd1, 0x4a, 0x8e, 0x66, 0xfe, 0x38, 0x9d, 0x84, 0xa0,
	0xec, 0x56, 0xc1, 0x3f, 0xfb, 0xdf, 0x33, 0xc5, 0xb9, 0x37, 0x1b, 0x15, 0x2e, 0xcf, 0x14, 0xd5,
	0xc2, 0xda, 0x44, 0xab, 0x4a, 0x90, 0xe0, 0xc8, 0x17, 0x10, 0x13, 0x9a, 0x0f, 0xa0, 0xdd, 0xec,
	0x19, 0xfd, 0x96, 0xd7, 0x2e, 0x1e, 0x7b, 0xf3, 0xa7, 0xf3, 0x84, 0xab, 0x83, 0xf2, 0x13, 0xce,
	0x23, 0x69, 0x5f, 0xad, 0x99, 0x70, 0x75, 0x5a, 0x6e, 0xce, 0x75, 0x7e, 0x19, 0xc8, 0x3c, 0xf3,
	0x55, 0xd6, 0x2e, 0x6a, 0xeb, 0x5e, 0xfd, 0xf2, 0xaa, 0x15, 0x03, 0x66, 0x0e, 0x6e, 0x96, 0x3d,
	0xe7, 0x17, 0x11, 0x4f, 0x87, 0xd8, 0xd5, 0x55, 0xef, 0x9a, 0xde, 0x5d, 0x2e, 0x2d, 0x17, 0xdd,
	0x88, 0x29, 0xf3, 0x75, 0x68, 0x01, 0x67, 0x41, 0x2a, 0x04, 0xb0, 0x20, 0xab, 0x35, 0x4d, 0xeb,
	0x31, 0x65, 0xc5, 0xb7, 0xec, 0x55, 0xe0, 0xce, 0xfe, 0x97, 0x6f, 0x9f, 0xba, 0x4f, 0xd1, 0xe3,
	0xbf, 0x47, 0x3e, 0xac, 0x99, 0xf6, 0xce, 0xa3, 0x5c, 0xb3, 0x5d, 0xfe, 0x8d, 0xfd, 0x93, 0x66,
	0x07, 0xe7, 0xe8, 0x03, 0xb4, 0x59, 0x13, 0x1d, 0xbd, 0x39, 0x3e, 0xe9, 0x1a, 0xdf, 0x4f, 0xba,
	0xc6, 0xcf, 0x93, 0xae, 0xf1, 0xf5, 0xc3, 0xf1, 0x8f, 0x66, 0x63, 0xad, 0x81, 0xb6, 0x28, 0xd7,
	0x11, 0x26, 0x82, 0xcf, 0xb2, 0x4b, 0x27, 0x68, 0xb4, 0x7e, 0x5e, 0xe9, 0xe6, 0x61, 0xb9, 0xc6,
	0xb8, 0x59, 0xa4, 0x36, 0xfc, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xa6, 0xec, 0x2d, 0xac, 0x8c, 0x05,
	0x00, 0x00,
}

func (m *CircuitBreakers) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CircuitBreakers) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CircuitBreakers) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Thresholds) > 0 {
		for iNdEx := len(m.Thresholds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Thresholds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCircuitBreaker(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *CircuitBreakers_Thresholds) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CircuitBreakers_Thresholds) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CircuitBreakers_Thresholds) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.RetryBudget != nil {
		{
			size, err := m.RetryBudget.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCircuitBreaker(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.MaxConnectionPools != nil {
		{
			size, err := m.MaxConnectionPools.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCircuitBreaker(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.TrackRemaining {
		i--
		if m.TrackRemaining {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.MaxRetries != nil {
		{
			size, err := m.MaxRetries.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCircuitBreaker(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.MaxRequests != nil {
		{
			size, err := m.MaxRequests.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCircuitBreaker(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.MaxPendingRequests != nil {
		{
			size, err := m.MaxPendingRequests.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCircuitBreaker(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.MaxConnections != nil {
		{
			size, err := m.MaxConnections.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCircuitBreaker(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Priority != 0 {
		i = encodeVarintCircuitBreaker(dAtA, i, uint64(m.Priority))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CircuitBreakers_Thresholds_RetryBudget) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CircuitBreakers_Thresholds_RetryBudget) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CircuitBreakers_Thresholds_RetryBudget) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.MinRetryConcurrency != nil {
		{
			size, err := m.MinRetryConcurrency.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCircuitBreaker(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.BudgetPercent != nil {
		{
			size, err := m.BudgetPercent.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCircuitBreaker(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCircuitBreaker(dAtA []byte, offset int, v uint64) int {
	offset -= sovCircuitBreaker(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CircuitBreakers) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Thresholds) > 0 {
		for _, e := range m.Thresholds {
			l = e.Size()
			n += 1 + l + sovCircuitBreaker(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CircuitBreakers_Thresholds) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Priority != 0 {
		n += 1 + sovCircuitBreaker(uint64(m.Priority))
	}
	if m.MaxConnections != nil {
		l = m.MaxConnections.Size()
		n += 1 + l + sovCircuitBreaker(uint64(l))
	}
	if m.MaxPendingRequests != nil {
		l = m.MaxPendingRequests.Size()
		n += 1 + l + sovCircuitBreaker(uint64(l))
	}
	if m.MaxRequests != nil {
		l = m.MaxRequests.Size()
		n += 1 + l + sovCircuitBreaker(uint64(l))
	}
	if m.MaxRetries != nil {
		l = m.MaxRetries.Size()
		n += 1 + l + sovCircuitBreaker(uint64(l))
	}
	if m.TrackRemaining {
		n += 2
	}
	if m.MaxConnectionPools != nil {
		l = m.MaxConnectionPools.Size()
		n += 1 + l + sovCircuitBreaker(uint64(l))
	}
	if m.RetryBudget != nil {
		l = m.RetryBudget.Size()
		n += 1 + l + sovCircuitBreaker(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CircuitBreakers_Thresholds_RetryBudget) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BudgetPercent != nil {
		l = m.BudgetPercent.Size()
		n += 1 + l + sovCircuitBreaker(uint64(l))
	}
	if m.MinRetryConcurrency != nil {
		l = m.MinRetryConcurrency.Size()
		n += 1 + l + sovCircuitBreaker(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCircuitBreaker(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCircuitBreaker(x uint64) (n int) {
	return sovCircuitBreaker(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CircuitBreakers) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCircuitBreaker
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
			return fmt.Errorf("proto: CircuitBreakers: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CircuitBreakers: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Thresholds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCircuitBreaker
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
				return ErrInvalidLengthCircuitBreaker
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCircuitBreaker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Thresholds = append(m.Thresholds, &CircuitBreakers_Thresholds{})
			if err := m.Thresholds[len(m.Thresholds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCircuitBreaker(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCircuitBreaker
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCircuitBreaker
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
func (m *CircuitBreakers_Thresholds) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCircuitBreaker
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
			return fmt.Errorf("proto: Thresholds: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Thresholds: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Priority", wireType)
			}
			m.Priority = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCircuitBreaker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Priority |= v4alpha.RoutingPriority(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxConnections", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCircuitBreaker
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
				return ErrInvalidLengthCircuitBreaker
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCircuitBreaker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxConnections == nil {
				m.MaxConnections = &types.UInt32Value{}
			}
			if err := m.MaxConnections.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxPendingRequests", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCircuitBreaker
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
				return ErrInvalidLengthCircuitBreaker
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCircuitBreaker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxPendingRequests == nil {
				m.MaxPendingRequests = &types.UInt32Value{}
			}
			if err := m.MaxPendingRequests.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxRequests", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCircuitBreaker
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
				return ErrInvalidLengthCircuitBreaker
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCircuitBreaker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxRequests == nil {
				m.MaxRequests = &types.UInt32Value{}
			}
			if err := m.MaxRequests.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxRetries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCircuitBreaker
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
				return ErrInvalidLengthCircuitBreaker
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCircuitBreaker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxRetries == nil {
				m.MaxRetries = &types.UInt32Value{}
			}
			if err := m.MaxRetries.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrackRemaining", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCircuitBreaker
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
			m.TrackRemaining = bool(v != 0)
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxConnectionPools", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCircuitBreaker
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
				return ErrInvalidLengthCircuitBreaker
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCircuitBreaker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxConnectionPools == nil {
				m.MaxConnectionPools = &types.UInt32Value{}
			}
			if err := m.MaxConnectionPools.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RetryBudget", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCircuitBreaker
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
				return ErrInvalidLengthCircuitBreaker
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCircuitBreaker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RetryBudget == nil {
				m.RetryBudget = &CircuitBreakers_Thresholds_RetryBudget{}
			}
			if err := m.RetryBudget.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCircuitBreaker(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCircuitBreaker
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCircuitBreaker
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
func (m *CircuitBreakers_Thresholds_RetryBudget) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCircuitBreaker
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
			return fmt.Errorf("proto: RetryBudget: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RetryBudget: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BudgetPercent", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCircuitBreaker
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
				return ErrInvalidLengthCircuitBreaker
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCircuitBreaker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BudgetPercent == nil {
				m.BudgetPercent = &v3.Percent{}
			}
			if err := m.BudgetPercent.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinRetryConcurrency", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCircuitBreaker
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
				return ErrInvalidLengthCircuitBreaker
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCircuitBreaker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MinRetryConcurrency == nil {
				m.MinRetryConcurrency = &types.UInt32Value{}
			}
			if err := m.MinRetryConcurrency.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCircuitBreaker(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCircuitBreaker
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCircuitBreaker
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
func skipCircuitBreaker(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCircuitBreaker
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
					return 0, ErrIntOverflowCircuitBreaker
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
					return 0, ErrIntOverflowCircuitBreaker
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
				return 0, ErrInvalidLengthCircuitBreaker
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCircuitBreaker
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCircuitBreaker
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCircuitBreaker        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCircuitBreaker          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCircuitBreaker = fmt.Errorf("proto: unexpected end of group")
)
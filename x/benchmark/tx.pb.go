// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/benchmark/v1/tx.proto

package benchmark

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgLoadTestOps defines a message containing a sequence of load test operations.
type MsgLoadTest struct {
	Caller []byte `protobuf:"bytes,1,opt,name=caller,proto3" json:"caller,omitempty"`
	Ops    []*Op  `protobuf:"bytes,2,rep,name=ops,proto3" json:"ops,omitempty"`
}

func (m *MsgLoadTest) Reset()         { *m = MsgLoadTest{} }
func (m *MsgLoadTest) String() string { return proto.CompactTextString(m) }
func (*MsgLoadTest) ProtoMessage()    {}
func (*MsgLoadTest) Descriptor() ([]byte, []int) {
	return fileDescriptor_481e3d7f7138b75b, []int{0}
}
func (m *MsgLoadTest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgLoadTest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgLoadTest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgLoadTest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgLoadTest.Merge(m, src)
}
func (m *MsgLoadTest) XXX_Size() int {
	return m.Size()
}
func (m *MsgLoadTest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgLoadTest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgLoadTest proto.InternalMessageInfo

func (m *MsgLoadTest) GetCaller() []byte {
	if m != nil {
		return m.Caller
	}
	return nil
}

func (m *MsgLoadTest) GetOps() []*Op {
	if m != nil {
		return m.Ops
	}
	return nil
}

// MsgLoadTestResponse defines a message containing the results of a load test operation.
type MsgLoadTestResponse struct {
	TotalTime   uint64 `protobuf:"varint,1,opt,name=total_time,json=totalTime,proto3" json:"total_time,omitempty"`
	TotalErrors uint64 `protobuf:"varint,2,opt,name=total_errors,json=totalErrors,proto3" json:"total_errors,omitempty"`
}

func (m *MsgLoadTestResponse) Reset()         { *m = MsgLoadTestResponse{} }
func (m *MsgLoadTestResponse) String() string { return proto.CompactTextString(m) }
func (*MsgLoadTestResponse) ProtoMessage()    {}
func (*MsgLoadTestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_481e3d7f7138b75b, []int{1}
}
func (m *MsgLoadTestResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgLoadTestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgLoadTestResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgLoadTestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgLoadTestResponse.Merge(m, src)
}
func (m *MsgLoadTestResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgLoadTestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgLoadTestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgLoadTestResponse proto.InternalMessageInfo

func (m *MsgLoadTestResponse) GetTotalTime() uint64 {
	if m != nil {
		return m.TotalTime
	}
	return 0
}

func (m *MsgLoadTestResponse) GetTotalErrors() uint64 {
	if m != nil {
		return m.TotalErrors
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgLoadTest)(nil), "cosmos.benchmark.v1.MsgLoadTest")
	proto.RegisterType((*MsgLoadTestResponse)(nil), "cosmos.benchmark.v1.MsgLoadTestResponse")
}

func init() { proto.RegisterFile("cosmos/benchmark/v1/tx.proto", fileDescriptor_481e3d7f7138b75b) }

var fileDescriptor_481e3d7f7138b75b = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x49, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xd6, 0x4f, 0x4a, 0xcd, 0x4b, 0xce, 0xc8, 0x4d, 0x2c, 0xca, 0xd6, 0x2f, 0x33, 0xd4,
	0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x86, 0xc8, 0xea, 0xc1, 0x65, 0xf5,
	0xca, 0x0c, 0xa5, 0x94, 0xb1, 0x69, 0x41, 0xa8, 0x00, 0xeb, 0x94, 0x12, 0x87, 0x2a, 0xca, 0x2d,
	0x4e, 0x07, 0x49, 0xe7, 0x16, 0xa7, 0x43, 0x25, 0x04, 0x13, 0x73, 0x33, 0xf3, 0xf2, 0xf5, 0xc1,
	0x24, 0x44, 0x48, 0xa9, 0x9d, 0x91, 0x8b, 0xdb, 0xb7, 0x38, 0xdd, 0x27, 0x3f, 0x31, 0x25, 0x24,
	0xb5, 0xb8, 0x44, 0x48, 0x8c, 0x8b, 0x2d, 0x39, 0x31, 0x27, 0x27, 0xb5, 0x48, 0x82, 0x51, 0x81,
	0x51, 0x83, 0x27, 0x08, 0xca, 0x13, 0xd2, 0xe4, 0x62, 0xce, 0x2f, 0x28, 0x96, 0x60, 0x52, 0x60,
	0xd6, 0xe0, 0x36, 0x12, 0xd7, 0xc3, 0xe2, 0x36, 0x3d, 0xff, 0x82, 0x20, 0x90, 0x1a, 0x2b, 0xd3,
	0xa6, 0xe7, 0x1b, 0xb4, 0xa0, 0xfa, 0xba, 0x9e, 0x6f, 0xd0, 0x52, 0x85, 0x28, 0xd6, 0x2d, 0x4e,
	0xc9, 0xd6, 0xaf, 0x40, 0x75, 0x39, 0x92, 0xcd, 0x4a, 0xe1, 0x5c, 0xc2, 0x48, 0xdc, 0xa0, 0xd4,
	0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x59, 0x2e, 0xae, 0x92, 0xfc, 0x92, 0xc4, 0x9c, 0xf8,
	0x92, 0xcc, 0xdc, 0x54, 0xb0, 0xa3, 0x58, 0x82, 0x38, 0xc1, 0x22, 0x21, 0x99, 0xb9, 0xa9, 0x42,
	0x8a, 0x5c, 0x3c, 0x10, 0xe9, 0xd4, 0xa2, 0xa2, 0xfc, 0x22, 0x90, 0x03, 0x41, 0x0a, 0xb8, 0xc1,
	0x62, 0xae, 0x60, 0x21, 0xa3, 0x14, 0x2e, 0x66, 0xdf, 0xe2, 0x74, 0xa1, 0x30, 0x2e, 0x0e, 0xb8,
	0x2f, 0x15, 0xb0, 0x7a, 0x00, 0xc9, 0x7a, 0x29, 0x0d, 0x42, 0x2a, 0x60, 0x0e, 0x94, 0x62, 0x6d,
	0x78, 0xbe, 0x41, 0x8b, 0xd1, 0xc9, 0xe8, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f,
	0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18,
	0xa2, 0x24, 0x20, 0x26, 0x15, 0xa7, 0x64, 0xeb, 0x65, 0xe6, 0x23, 0x87, 0x40, 0x12, 0x1b, 0x38,
	0x0e, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x47, 0x7e, 0x92, 0xc2, 0x09, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// LoadTest defines a method for executing a sequence of load test operations.
	LoadTest(ctx context.Context, in *MsgLoadTest, opts ...grpc.CallOption) (*MsgLoadTestResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) LoadTest(ctx context.Context, in *MsgLoadTest, opts ...grpc.CallOption) (*MsgLoadTestResponse, error) {
	out := new(MsgLoadTestResponse)
	err := c.cc.Invoke(ctx, "/cosmos.benchmark.v1.Msg/LoadTest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// LoadTest defines a method for executing a sequence of load test operations.
	LoadTest(context.Context, *MsgLoadTest) (*MsgLoadTestResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) LoadTest(ctx context.Context, req *MsgLoadTest) (*MsgLoadTestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadTest not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_LoadTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgLoadTest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).LoadTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.benchmark.v1.Msg/LoadTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).LoadTest(ctx, req.(*MsgLoadTest))
	}
	return interceptor(ctx, in, info, handler)
}

var Msg_serviceDesc = _Msg_serviceDesc
var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.benchmark.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoadTest",
			Handler:    _Msg_LoadTest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/benchmark/v1/tx.proto",
}

func (m *MsgLoadTest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgLoadTest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgLoadTest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Ops) > 0 {
		for iNdEx := len(m.Ops) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Ops[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Caller) > 0 {
		i -= len(m.Caller)
		copy(dAtA[i:], m.Caller)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Caller)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgLoadTestResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgLoadTestResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgLoadTestResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TotalErrors != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.TotalErrors))
		i--
		dAtA[i] = 0x10
	}
	if m.TotalTime != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.TotalTime))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgLoadTest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Caller)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Ops) > 0 {
		for _, e := range m.Ops {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgLoadTestResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TotalTime != 0 {
		n += 1 + sovTx(uint64(m.TotalTime))
	}
	if m.TotalErrors != 0 {
		n += 1 + sovTx(uint64(m.TotalErrors))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgLoadTest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgLoadTest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgLoadTest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Caller", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Caller = append(m.Caller[:0], dAtA[iNdEx:postIndex]...)
			if m.Caller == nil {
				m.Caller = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ops", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ops = append(m.Ops, &Op{})
			if err := m.Ops[len(m.Ops)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgLoadTestResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgLoadTestResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgLoadTestResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalTime", wireType)
			}
			m.TotalTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalErrors", wireType)
			}
			m.TotalErrors = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalErrors |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/p2p/peer/messages/sync_messages.proto

package messages

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type SyncRequest struct {
}

func (m *SyncRequest) Reset()      { *m = SyncRequest{} }
func (*SyncRequest) ProtoMessage() {}
func (*SyncRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_sync_messages_4163a8f77fa64dea, []int{0}
}
func (m *SyncRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SyncRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SyncRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SyncRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncRequest.Merge(dst, src)
}
func (m *SyncRequest) XXX_Size() int {
	return m.Size()
}
func (m *SyncRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SyncRequest proto.InternalMessageInfo

type SyncResponse struct {
	SignedMessage []string `protobuf:"bytes,1,rep,name=signed_message,json=signedMessage" json:"signed_message,omitempty"`
}

func (m *SyncResponse) Reset()      { *m = SyncResponse{} }
func (*SyncResponse) ProtoMessage() {}
func (*SyncResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_sync_messages_4163a8f77fa64dea, []int{1}
}
func (m *SyncResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SyncResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SyncResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SyncResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncResponse.Merge(dst, src)
}
func (m *SyncResponse) XXX_Size() int {
	return m.Size()
}
func (m *SyncResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SyncResponse proto.InternalMessageInfo

func (m *SyncResponse) GetSignedMessage() []string {
	if m != nil {
		return m.SignedMessage
	}
	return nil
}

func init() {
	proto.RegisterType((*SyncRequest)(nil), "messages.SyncRequest")
	proto.RegisterType((*SyncResponse)(nil), "messages.SyncResponse")
}
func (this *SyncRequest) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*SyncRequest)
	if !ok {
		that2, ok := that.(SyncRequest)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *SyncRequest")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *SyncRequest but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *SyncRequest but is not nil && this == nil")
	}
	return nil
}
func (this *SyncRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SyncRequest)
	if !ok {
		that2, ok := that.(SyncRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *SyncResponse) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*SyncResponse)
	if !ok {
		that2, ok := that.(SyncResponse)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *SyncResponse")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *SyncResponse but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *SyncResponse but is not nil && this == nil")
	}
	if len(this.SignedMessage) != len(that1.SignedMessage) {
		return fmt.Errorf("SignedMessage this(%v) Not Equal that(%v)", len(this.SignedMessage), len(that1.SignedMessage))
	}
	for i := range this.SignedMessage {
		if this.SignedMessage[i] != that1.SignedMessage[i] {
			return fmt.Errorf("SignedMessage this[%v](%v) Not Equal that[%v](%v)", i, this.SignedMessage[i], i, that1.SignedMessage[i])
		}
	}
	return nil
}
func (this *SyncResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SyncResponse)
	if !ok {
		that2, ok := that.(SyncResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.SignedMessage) != len(that1.SignedMessage) {
		return false
	}
	for i := range this.SignedMessage {
		if this.SignedMessage[i] != that1.SignedMessage[i] {
			return false
		}
	}
	return true
}
func (this *SyncRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&messages.SyncRequest{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SyncResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&messages.SyncResponse{")
	s = append(s, "SignedMessage: "+fmt.Sprintf("%#v", this.SignedMessage)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringSyncMessages(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *SyncRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SyncRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *SyncResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SyncResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.SignedMessage) > 0 {
		for _, s := range m.SignedMessage {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func encodeVarintSyncMessages(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedSyncRequest(r randySyncMessages, easy bool) *SyncRequest {
	this := &SyncRequest{}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedSyncResponse(r randySyncMessages, easy bool) *SyncResponse {
	this := &SyncResponse{}
	v1 := r.Intn(10)
	this.SignedMessage = make([]string, v1)
	for i := 0; i < v1; i++ {
		this.SignedMessage[i] = string(randStringSyncMessages(r))
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randySyncMessages interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneSyncMessages(r randySyncMessages) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringSyncMessages(r randySyncMessages) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneSyncMessages(r)
	}
	return string(tmps)
}
func randUnrecognizedSyncMessages(r randySyncMessages, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldSyncMessages(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldSyncMessages(dAtA []byte, r randySyncMessages, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateSyncMessages(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateSyncMessages(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateSyncMessages(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateSyncMessages(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateSyncMessages(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateSyncMessages(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateSyncMessages(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *SyncRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *SyncResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.SignedMessage) > 0 {
		for _, s := range m.SignedMessage {
			l = len(s)
			n += 1 + l + sovSyncMessages(uint64(l))
		}
	}
	return n
}

func sovSyncMessages(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSyncMessages(x uint64) (n int) {
	return sovSyncMessages(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *SyncRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SyncRequest{`,
		`}`,
	}, "")
	return s
}
func (this *SyncResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SyncResponse{`,
		`SignedMessage:` + fmt.Sprintf("%v", this.SignedMessage) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringSyncMessages(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *SyncRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSyncMessages
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
			return fmt.Errorf("proto: SyncRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SyncRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipSyncMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSyncMessages
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
func (m *SyncResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSyncMessages
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
			return fmt.Errorf("proto: SyncResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SyncResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignedMessage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncMessages
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
				return ErrInvalidLengthSyncMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SignedMessage = append(m.SignedMessage, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSyncMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSyncMessages
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
func skipSyncMessages(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSyncMessages
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
					return 0, ErrIntOverflowSyncMessages
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
					return 0, ErrIntOverflowSyncMessages
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
				return 0, ErrInvalidLengthSyncMessages
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSyncMessages
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
				next, err := skipSyncMessages(dAtA[start:])
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
	ErrInvalidLengthSyncMessages = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSyncMessages   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("pkg/p2p/peer/messages/sync_messages.proto", fileDescriptor_sync_messages_4163a8f77fa64dea)
}

var fileDescriptor_sync_messages_4163a8f77fa64dea = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2c, 0xc8, 0x4e, 0xd7,
	0x2f, 0x30, 0x2a, 0xd0, 0x2f, 0x48, 0x4d, 0x2d, 0xd2, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f,
	0x2d, 0xd6, 0x2f, 0xae, 0xcc, 0x4b, 0x8e, 0x87, 0xf1, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85,
	0x38, 0x60, 0x7c, 0x29, 0xdd, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd,
	0xf4, 0xfc, 0xf4, 0x7c, 0x7d, 0xb0, 0x82, 0xa4, 0xd2, 0x34, 0x30, 0x0f, 0xcc, 0x01, 0xb3, 0x20,
	0x1a, 0x95, 0x78, 0xb9, 0xb8, 0x83, 0x2b, 0xf3, 0x92, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b,
	0x94, 0x4c, 0xb9, 0x78, 0x20, 0xdc, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x55, 0x2e, 0xbe,
	0xe2, 0xcc, 0xf4, 0xbc, 0xd4, 0x14, 0x98, 0x85, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0x9c, 0x41, 0xbc,
	0x10, 0x51, 0x5f, 0x88, 0xa0, 0x93, 0xc9, 0x8d, 0x87, 0x72, 0x0c, 0x0f, 0x1e, 0xca, 0x31, 0x7e,
	0x78, 0x28, 0xc7, 0xf8, 0xe3, 0xa1, 0x1c, 0x63, 0xc3, 0x23, 0x39, 0xc6, 0x15, 0x8f, 0xe4, 0x18,
	0x77, 0x3c, 0x92, 0x63, 0x3c, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4,
	0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x24, 0x36,
	0xb0, 0x13, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x05, 0x94, 0xf4, 0xe8, 0x00, 0x00,
	0x00,
}

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pochuman/transaction_data.proto

package types

import (
	fmt "fmt"
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

type TransactionData struct {
	Index              string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	OriginChain        string `protobuf:"bytes,2,opt,name=originChain,proto3" json:"originChain,omitempty"`
	OriginAddress      string `protobuf:"bytes,3,opt,name=originAddress,proto3" json:"originAddress,omitempty"`
	TargetChain        string `protobuf:"bytes,4,opt,name=targetChain,proto3" json:"targetChain,omitempty"`
	TargetAddress      string `protobuf:"bytes,5,opt,name=targetAddress,proto3" json:"targetAddress,omitempty"`
	Amount             string `protobuf:"bytes,6,opt,name=amount,proto3" json:"amount,omitempty"`
	Time               string `protobuf:"bytes,7,opt,name=time,proto3" json:"time,omitempty"`
	Creator            string `protobuf:"bytes,8,opt,name=creator,proto3" json:"creator,omitempty"`
	Status             string `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	ConfirmedBlockHash string `protobuf:"bytes,10,opt,name=confirmedBlockHash,proto3" json:"confirmedBlockHash,omitempty"`
	SignedKey          string `protobuf:"bytes,11,opt,name=signedKey,proto3" json:"signedKey,omitempty"`
	Fee                string `protobuf:"bytes,12,opt,name=fee,proto3" json:"fee,omitempty"`
}

func (m *TransactionData) Reset()         { *m = TransactionData{} }
func (m *TransactionData) String() string { return proto.CompactTextString(m) }
func (*TransactionData) ProtoMessage()    {}
func (*TransactionData) Descriptor() ([]byte, []int) {
	return fileDescriptor_39f3589d73def3b5, []int{0}
}
func (m *TransactionData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TransactionData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransactionData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TransactionData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionData.Merge(m, src)
}
func (m *TransactionData) XXX_Size() int {
	return m.Size()
}
func (m *TransactionData) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionData.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionData proto.InternalMessageInfo

func (m *TransactionData) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *TransactionData) GetOriginChain() string {
	if m != nil {
		return m.OriginChain
	}
	return ""
}

func (m *TransactionData) GetOriginAddress() string {
	if m != nil {
		return m.OriginAddress
	}
	return ""
}

func (m *TransactionData) GetTargetChain() string {
	if m != nil {
		return m.TargetChain
	}
	return ""
}

func (m *TransactionData) GetTargetAddress() string {
	if m != nil {
		return m.TargetAddress
	}
	return ""
}

func (m *TransactionData) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *TransactionData) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *TransactionData) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *TransactionData) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *TransactionData) GetConfirmedBlockHash() string {
	if m != nil {
		return m.ConfirmedBlockHash
	}
	return ""
}

func (m *TransactionData) GetSignedKey() string {
	if m != nil {
		return m.SignedKey
	}
	return ""
}

func (m *TransactionData) GetFee() string {
	if m != nil {
		return m.Fee
	}
	return ""
}

func init() {
	proto.RegisterType((*TransactionData)(nil), "vigorousdeveloper.pochuman.pochuman.TransactionData")
}

func init() { proto.RegisterFile("pochuman/transaction_data.proto", fileDescriptor_39f3589d73def3b5) }

var fileDescriptor_39f3589d73def3b5 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xbd, 0x4e, 0xf3, 0x30,
	0x14, 0x86, 0x9b, 0xfe, 0x7e, 0x75, 0x3f, 0x04, 0xb2, 0x10, 0xf2, 0x80, 0x42, 0x05, 0x0c, 0x2c,
	0x24, 0x03, 0x03, 0x33, 0xa5, 0x03, 0x12, 0x13, 0x08, 0x31, 0xb0, 0x20, 0x37, 0x39, 0x4d, 0x2c,
	0x1a, 0x9f, 0xc8, 0x76, 0xaa, 0xf6, 0x2e, 0xb8, 0x2c, 0xc6, 0x8e, 0x8c, 0xa8, 0xbd, 0x06, 0x76,
	0x14, 0xbb, 0xa1, 0xad, 0xc4, 0x76, 0xde, 0xe7, 0x1c, 0x3f, 0xcb, 0x6b, 0x72, 0x92, 0x63, 0x94,
	0x16, 0x19, 0x97, 0xa1, 0x51, 0x5c, 0x6a, 0x1e, 0x19, 0x81, 0xf2, 0x35, 0xe6, 0x86, 0x07, 0xb9,
	0x42, 0x83, 0xf4, 0x6c, 0x2a, 0x12, 0x54, 0x58, 0xe8, 0x18, 0xa6, 0x30, 0xc1, 0x1c, 0x54, 0x50,
	0x3d, 0xf9, 0x1d, 0x4e, 0xbf, 0xeb, 0x64, 0xff, 0x69, 0xf3, 0x7e, 0xc8, 0x0d, 0xa7, 0x87, 0xa4,
	0x25, 0x64, 0x0c, 0x33, 0xe6, 0xf5, 0xbd, 0x8b, 0xee, 0xa3, 0x0b, 0xb4, 0x4f, 0x7a, 0xa8, 0x44,
	0x22, 0xe4, 0x6d, 0xca, 0x85, 0x64, 0x75, 0xbb, 0xdb, 0x46, 0xf4, 0x9c, 0xec, 0xb9, 0x78, 0x13,
	0xc7, 0x0a, 0xb4, 0x66, 0x0d, 0x7b, 0xb3, 0x0b, 0x4b, 0x8f, 0xe1, 0x2a, 0x01, 0xe3, 0x3c, 0x4d,
	0xe7, 0xd9, 0x42, 0xa5, 0xc7, 0xc5, 0xca, 0xd3, 0x72, 0x9e, 0x1d, 0x48, 0x8f, 0x48, 0x9b, 0x67,
	0x58, 0x48, 0xc3, 0xda, 0x76, 0xbd, 0x4e, 0x94, 0x92, 0xa6, 0x11, 0x19, 0xb0, 0x8e, 0xa5, 0x76,
	0xa6, 0x8c, 0x74, 0x22, 0x05, 0xdc, 0xa0, 0x62, 0xff, 0x2c, 0xae, 0x62, 0x69, 0xd1, 0x86, 0x9b,
	0x42, 0xb3, 0xae, 0xb3, 0xb8, 0x44, 0x03, 0x42, 0x23, 0x94, 0x63, 0xa1, 0x32, 0x88, 0x07, 0x13,
	0x8c, 0xde, 0xee, 0xb8, 0x4e, 0x19, 0xb1, 0x37, 0x7f, 0x6c, 0xe8, 0x31, 0xe9, 0x6a, 0x91, 0x48,
	0x88, 0xef, 0x61, 0xce, 0x7a, 0xf6, 0x6c, 0x03, 0xe8, 0x01, 0x69, 0x8c, 0x01, 0xd8, 0x7f, 0xcb,
	0xcb, 0x71, 0xf0, 0xf0, 0xb1, 0xf4, 0xbd, 0xc5, 0xd2, 0xf7, 0xbe, 0x96, 0xbe, 0xf7, 0xbe, 0xf2,
	0x6b, 0x8b, 0x95, 0x5f, 0xfb, 0x5c, 0xf9, 0xb5, 0x97, 0xeb, 0x44, 0x98, 0xb4, 0x18, 0x05, 0x11,
	0x66, 0xe1, 0xf3, 0xba, 0xc1, 0x61, 0xd5, 0x60, 0x98, 0x63, 0x74, 0xe9, 0x5a, 0x9f, 0x85, 0x9b,
	0x0f, 0x30, 0xcf, 0x41, 0x8f, 0xda, 0xb6, 0xf6, 0xab, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb3,
	0x35, 0x77, 0xb9, 0x19, 0x02, 0x00, 0x00,
}

func (m *TransactionData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransactionData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TransactionData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Fee) > 0 {
		i -= len(m.Fee)
		copy(dAtA[i:], m.Fee)
		i = encodeVarintTransactionData(dAtA, i, uint64(len(m.Fee)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.SignedKey) > 0 {
		i -= len(m.SignedKey)
		copy(dAtA[i:], m.SignedKey)
		i = encodeVarintTransactionData(dAtA, i, uint64(len(m.SignedKey)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.ConfirmedBlockHash) > 0 {
		i -= len(m.ConfirmedBlockHash)
		copy(dAtA[i:], m.ConfirmedBlockHash)
		i = encodeVarintTransactionData(dAtA, i, uint64(len(m.ConfirmedBlockHash)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintTransactionData(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTransactionData(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Time) > 0 {
		i -= len(m.Time)
		copy(dAtA[i:], m.Time)
		i = encodeVarintTransactionData(dAtA, i, uint64(len(m.Time)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintTransactionData(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.TargetAddress) > 0 {
		i -= len(m.TargetAddress)
		copy(dAtA[i:], m.TargetAddress)
		i = encodeVarintTransactionData(dAtA, i, uint64(len(m.TargetAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.TargetChain) > 0 {
		i -= len(m.TargetChain)
		copy(dAtA[i:], m.TargetChain)
		i = encodeVarintTransactionData(dAtA, i, uint64(len(m.TargetChain)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.OriginAddress) > 0 {
		i -= len(m.OriginAddress)
		copy(dAtA[i:], m.OriginAddress)
		i = encodeVarintTransactionData(dAtA, i, uint64(len(m.OriginAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.OriginChain) > 0 {
		i -= len(m.OriginChain)
		copy(dAtA[i:], m.OriginChain)
		i = encodeVarintTransactionData(dAtA, i, uint64(len(m.OriginChain)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintTransactionData(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTransactionData(dAtA []byte, offset int, v uint64) int {
	offset -= sovTransactionData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TransactionData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovTransactionData(uint64(l))
	}
	l = len(m.OriginChain)
	if l > 0 {
		n += 1 + l + sovTransactionData(uint64(l))
	}
	l = len(m.OriginAddress)
	if l > 0 {
		n += 1 + l + sovTransactionData(uint64(l))
	}
	l = len(m.TargetChain)
	if l > 0 {
		n += 1 + l + sovTransactionData(uint64(l))
	}
	l = len(m.TargetAddress)
	if l > 0 {
		n += 1 + l + sovTransactionData(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovTransactionData(uint64(l))
	}
	l = len(m.Time)
	if l > 0 {
		n += 1 + l + sovTransactionData(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTransactionData(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovTransactionData(uint64(l))
	}
	l = len(m.ConfirmedBlockHash)
	if l > 0 {
		n += 1 + l + sovTransactionData(uint64(l))
	}
	l = len(m.SignedKey)
	if l > 0 {
		n += 1 + l + sovTransactionData(uint64(l))
	}
	l = len(m.Fee)
	if l > 0 {
		n += 1 + l + sovTransactionData(uint64(l))
	}
	return n
}

func sovTransactionData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTransactionData(x uint64) (n int) {
	return sovTransactionData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TransactionData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransactionData
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
			return fmt.Errorf("proto: TransactionData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TransactionData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransactionData
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
				return ErrInvalidLengthTransactionData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransactionData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginChain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransactionData
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
				return ErrInvalidLengthTransactionData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransactionData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OriginChain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransactionData
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
				return ErrInvalidLengthTransactionData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransactionData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OriginAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetChain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransactionData
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
				return ErrInvalidLengthTransactionData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransactionData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TargetChain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransactionData
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
				return ErrInvalidLengthTransactionData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransactionData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TargetAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransactionData
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
				return ErrInvalidLengthTransactionData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransactionData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransactionData
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
				return ErrInvalidLengthTransactionData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransactionData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Time = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransactionData
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
				return ErrInvalidLengthTransactionData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransactionData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransactionData
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
				return ErrInvalidLengthTransactionData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransactionData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConfirmedBlockHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransactionData
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
				return ErrInvalidLengthTransactionData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransactionData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConfirmedBlockHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignedKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransactionData
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
				return ErrInvalidLengthTransactionData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransactionData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SignedKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransactionData
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
				return ErrInvalidLengthTransactionData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransactionData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTransactionData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTransactionData
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
func skipTransactionData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTransactionData
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
					return 0, ErrIntOverflowTransactionData
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
					return 0, ErrIntOverflowTransactionData
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
				return 0, ErrInvalidLengthTransactionData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTransactionData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTransactionData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTransactionData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTransactionData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTransactionData = fmt.Errorf("proto: unexpected end of group")
)

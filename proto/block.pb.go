// Code generated by protoc-gen-go.
// source: block.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	block.proto
	message.proto
	transaction.proto

It has these top-level messages:
	BlockHeader
	Block
	GetBlocksMsg
	StatusMsg
	GetInvMsg
	OnBlockMsg
	OnTransactionMsg
	GetDataMsg
	OutPoint
	TxIn
	TxOut
	ContractSpec
	TxHeader
	Transaction
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type BlockHeader struct {
	PreviousHash  []byte `protobuf:"bytes,1,opt,name=previousHash,proto3" json:"previousHash,omitempty"`
	TxsMerkleHash []byte `protobuf:"bytes,2,opt,name=txsMerkleHash,proto3" json:"txsMerkleHash,omitempty"`
	StateHash     []byte `protobuf:"bytes,3,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	TimeStamp     uint32 `protobuf:"varint,4,opt,name=timeStamp" json:"timeStamp,omitempty"`
	Height        uint32 `protobuf:"varint,5,opt,name=height" json:"height,omitempty"`
	Nonce         uint32 `protobuf:"varint,6,opt,name=nonce" json:"nonce,omitempty"`
}

func (m *BlockHeader) Reset()                    { *m = BlockHeader{} }
func (m *BlockHeader) String() string            { return proto1.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()               {}
func (*BlockHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BlockHeader) GetPreviousHash() []byte {
	if m != nil {
		return m.PreviousHash
	}
	return nil
}

func (m *BlockHeader) GetTxsMerkleHash() []byte {
	if m != nil {
		return m.TxsMerkleHash
	}
	return nil
}

func (m *BlockHeader) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *BlockHeader) GetTimeStamp() uint32 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

func (m *BlockHeader) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BlockHeader) GetNonce() uint32 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

type Block struct {
	Header       *BlockHeader   `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Transactions []*Transaction `protobuf:"bytes,2,rep,name=transactions" json:"transactions,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto1.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Block) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func init() {
	proto1.RegisterType((*BlockHeader)(nil), "proto.BlockHeader")
	proto1.RegisterType((*Block)(nil), "proto.Block")
}

func init() { proto1.RegisterFile("block.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0x6b, 0x02, 0x4e, 0xd2, 0x83, 0x83, 0xc8, 0x22, 0x1e, 0x42, 0xf0, 0x50, 0x3c,
	0xf4, 0xa0, 0xe0, 0x03, 0x78, 0xea, 0xc5, 0x4b, 0xf4, 0x05, 0xb6, 0x71, 0x30, 0x4b, 0xda, 0xdd,
	0xb0, 0x3b, 0x8a, 0xaf, 0xe7, 0x9b, 0x95, 0x9d, 0x2d, 0x24, 0x39, 0x2d, 0xf3, 0xfd, 0xdf, 0xc0,
	0xfe, 0x03, 0xe5, 0xe1, 0xe8, 0xba, 0x61, 0x37, 0x7a, 0xc7, 0x0e, 0x73, 0x79, 0xee, 0x6f, 0xd8,
	0x6b, 0x1b, 0x74, 0xc7, 0xc6, 0xd9, 0x94, 0x34, 0xff, 0x19, 0x94, 0x6f, 0xd1, 0xdc, 0x93, 0xfe,
	0x22, 0x8f, 0x0d, 0x54, 0xa3, 0xa7, 0x5f, 0xe3, 0x7e, 0xc2, 0x5e, 0x87, 0x5e, 0x65, 0x75, 0xb6,
	0xad, 0xda, 0x05, 0xc3, 0x47, 0xd8, 0xf0, 0x5f, 0x78, 0x27, 0x3f, 0x1c, 0x49, 0xa4, 0x95, 0x48,
	0x4b, 0x88, 0x0f, 0x70, 0x1d, 0x58, 0x73, 0x32, 0xd6, 0x62, 0x4c, 0x20, 0xa6, 0x6c, 0x4e, 0xf4,
	0xc1, 0xfa, 0x34, 0xaa, 0xab, 0x3a, 0xdb, 0x6e, 0xda, 0x09, 0xe0, 0x1d, 0x14, 0x3d, 0x99, 0xef,
	0x9e, 0x55, 0x2e, 0xd1, 0x65, 0xc2, 0x5b, 0xc8, 0xad, 0xb3, 0x1d, 0xa9, 0x42, 0x70, 0x1a, 0x9a,
	0x01, 0x72, 0xa9, 0x80, 0x4f, 0x71, 0x2d, 0xd6, 0x90, 0x6f, 0x97, 0xcf, 0x98, 0x4a, 0xee, 0x66,
	0x05, 0xdb, 0x8b, 0x81, 0xaf, 0x50, 0xcd, 0xae, 0x11, 0xd4, 0xaa, 0x5e, 0xcf, 0x36, 0x3e, 0xa7,
	0xa8, 0x5d, 0x78, 0x87, 0x42, 0x84, 0x97, 0x73, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe2, 0xbd, 0x44,
	0x82, 0x60, 0x01, 0x00, 0x00,
}

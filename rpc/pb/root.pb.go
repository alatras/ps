// Code generated by protoc-gen-go. DO NOT EDIT.
// source: root.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BigInt struct {
	Values               []byte   `protobuf:"bytes,1,opt,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BigInt) Reset()         { *m = BigInt{} }
func (m *BigInt) String() string { return proto.CompactTextString(m) }
func (*BigInt) ProtoMessage()    {}
func (*BigInt) Descriptor() ([]byte, []int) {
	return fileDescriptor_08a043f6ee9336a8, []int{0}
}
func (m *BigInt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BigInt.Unmarshal(m, b)
}
func (m *BigInt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BigInt.Marshal(b, m, deterministic)
}
func (m *BigInt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BigInt.Merge(m, src)
}
func (m *BigInt) XXX_Size() int {
	return xxx_messageInfo_BigInt.Size(m)
}
func (m *BigInt) XXX_DiscardUnknown() {
	xxx_messageInfo_BigInt.DiscardUnknown(m)
}

var xxx_messageInfo_BigInt proto.InternalMessageInfo

func (m *BigInt) GetValues() []byte {
	if m != nil {
		return m.Values
	}
	return nil
}

type Input struct {
	BlockNum             uint64   `protobuf:"varint,1,opt,name=blockNum,proto3" json:"blockNum,omitempty"`
	TxIdx                uint32   `protobuf:"varint,2,opt,name=txIdx,proto3" json:"txIdx,omitempty"`
	OutIdx               uint32   `protobuf:"varint,3,opt,name=outIdx,proto3" json:"outIdx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Input) Reset()         { *m = Input{} }
func (m *Input) String() string { return proto.CompactTextString(m) }
func (*Input) ProtoMessage()    {}
func (*Input) Descriptor() ([]byte, []int) {
	return fileDescriptor_08a043f6ee9336a8, []int{1}
}
func (m *Input) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Input.Unmarshal(m, b)
}
func (m *Input) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Input.Marshal(b, m, deterministic)
}
func (m *Input) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Input.Merge(m, src)
}
func (m *Input) XXX_Size() int {
	return xxx_messageInfo_Input.Size(m)
}
func (m *Input) XXX_DiscardUnknown() {
	xxx_messageInfo_Input.DiscardUnknown(m)
}

var xxx_messageInfo_Input proto.InternalMessageInfo

func (m *Input) GetBlockNum() uint64 {
	if m != nil {
		return m.BlockNum
	}
	return 0
}

func (m *Input) GetTxIdx() uint32 {
	if m != nil {
		return m.TxIdx
	}
	return 0
}

func (m *Input) GetOutIdx() uint32 {
	if m != nil {
		return m.OutIdx
	}
	return 0
}

type Output struct {
	NewOwner             []byte   `protobuf:"bytes,1,opt,name=newOwner,proto3" json:"newOwner,omitempty"`
	Amount               *BigInt  `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Output) Reset()         { *m = Output{} }
func (m *Output) String() string { return proto.CompactTextString(m) }
func (*Output) ProtoMessage()    {}
func (*Output) Descriptor() ([]byte, []int) {
	return fileDescriptor_08a043f6ee9336a8, []int{2}
}
func (m *Output) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Output.Unmarshal(m, b)
}
func (m *Output) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Output.Marshal(b, m, deterministic)
}
func (m *Output) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Output.Merge(m, src)
}
func (m *Output) XXX_Size() int {
	return xxx_messageInfo_Output.Size(m)
}
func (m *Output) XXX_DiscardUnknown() {
	xxx_messageInfo_Output.DiscardUnknown(m)
}

var xxx_messageInfo_Output proto.InternalMessageInfo

func (m *Output) GetNewOwner() []byte {
	if m != nil {
		return m.NewOwner
	}
	return nil
}

func (m *Output) GetAmount() *BigInt {
	if m != nil {
		return m.Amount
	}
	return nil
}

type BlockHeader struct {
	MerkleRoot           []byte   `protobuf:"bytes,1,opt,name=merkleRoot,proto3" json:"merkleRoot,omitempty"`
	RlpMerkleRoot        []byte   `protobuf:"bytes,2,opt,name=rlpMerkleRoot,proto3" json:"rlpMerkleRoot,omitempty"`
	PrevHash             []byte   `protobuf:"bytes,3,opt,name=prevHash,proto3" json:"prevHash,omitempty"`
	Number               uint64   `protobuf:"varint,4,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockHeader) Reset()         { *m = BlockHeader{} }
func (m *BlockHeader) String() string { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()    {}
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_08a043f6ee9336a8, []int{3}
}
func (m *BlockHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockHeader.Unmarshal(m, b)
}
func (m *BlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockHeader.Marshal(b, m, deterministic)
}
func (m *BlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeader.Merge(m, src)
}
func (m *BlockHeader) XXX_Size() int {
	return xxx_messageInfo_BlockHeader.Size(m)
}
func (m *BlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeader proto.InternalMessageInfo

func (m *BlockHeader) GetMerkleRoot() []byte {
	if m != nil {
		return m.MerkleRoot
	}
	return nil
}

func (m *BlockHeader) GetRlpMerkleRoot() []byte {
	if m != nil {
		return m.RlpMerkleRoot
	}
	return nil
}

func (m *BlockHeader) GetPrevHash() []byte {
	if m != nil {
		return m.PrevHash
	}
	return nil
}

func (m *BlockHeader) GetNumber() uint64 {
	if m != nil {
		return m.Number
	}
	return 0
}

type Block struct {
	Header               *BlockHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Hash                 []byte       `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_08a043f6ee9336a8, []int{4}
}
func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type Transaction struct {
	Input0               *Input   `protobuf:"bytes,1,opt,name=input0,proto3" json:"input0,omitempty"`
	Sig0                 []byte   `protobuf:"bytes,2,opt,name=sig0,proto3" json:"sig0,omitempty"`
	Input1               *Input   `protobuf:"bytes,3,opt,name=input1,proto3" json:"input1,omitempty"`
	Sig1                 []byte   `protobuf:"bytes,4,opt,name=sig1,proto3" json:"sig1,omitempty"`
	Output0              *Output  `protobuf:"bytes,5,opt,name=output0,proto3" json:"output0,omitempty"`
	Output1              *Output  `protobuf:"bytes,6,opt,name=output1,proto3" json:"output1,omitempty"`
	Fee                  *BigInt  `protobuf:"bytes,7,opt,name=fee,proto3" json:"fee,omitempty"`
	BlockNum             uint64   `protobuf:"varint,8,opt,name=BlockNum,proto3" json:"BlockNum,omitempty"`
	TxIdx                uint32   `protobuf:"varint,9,opt,name=TxIdx,proto3" json:"TxIdx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_08a043f6ee9336a8, []int{5}
}
func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetInput0() *Input {
	if m != nil {
		return m.Input0
	}
	return nil
}

func (m *Transaction) GetSig0() []byte {
	if m != nil {
		return m.Sig0
	}
	return nil
}

func (m *Transaction) GetInput1() *Input {
	if m != nil {
		return m.Input1
	}
	return nil
}

func (m *Transaction) GetSig1() []byte {
	if m != nil {
		return m.Sig1
	}
	return nil
}

func (m *Transaction) GetOutput0() *Output {
	if m != nil {
		return m.Output0
	}
	return nil
}

func (m *Transaction) GetOutput1() *Output {
	if m != nil {
		return m.Output1
	}
	return nil
}

func (m *Transaction) GetFee() *BigInt {
	if m != nil {
		return m.Fee
	}
	return nil
}

func (m *Transaction) GetBlockNum() uint64 {
	if m != nil {
		return m.BlockNum
	}
	return 0
}

func (m *Transaction) GetTxIdx() uint32 {
	if m != nil {
		return m.TxIdx
	}
	return 0
}

type GetBalanceRequest struct {
	Address              []byte   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBalanceRequest) Reset()         { *m = GetBalanceRequest{} }
func (m *GetBalanceRequest) String() string { return proto.CompactTextString(m) }
func (*GetBalanceRequest) ProtoMessage()    {}
func (*GetBalanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_08a043f6ee9336a8, []int{6}
}
func (m *GetBalanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBalanceRequest.Unmarshal(m, b)
}
func (m *GetBalanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBalanceRequest.Marshal(b, m, deterministic)
}
func (m *GetBalanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBalanceRequest.Merge(m, src)
}
func (m *GetBalanceRequest) XXX_Size() int {
	return xxx_messageInfo_GetBalanceRequest.Size(m)
}
func (m *GetBalanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBalanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBalanceRequest proto.InternalMessageInfo

func (m *GetBalanceRequest) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

type GetBalanceResponse struct {
	Balance              *BigInt  `protobuf:"bytes,1,opt,name=balance,proto3" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBalanceResponse) Reset()         { *m = GetBalanceResponse{} }
func (m *GetBalanceResponse) String() string { return proto.CompactTextString(m) }
func (*GetBalanceResponse) ProtoMessage()    {}
func (*GetBalanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_08a043f6ee9336a8, []int{7}
}
func (m *GetBalanceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBalanceResponse.Unmarshal(m, b)
}
func (m *GetBalanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBalanceResponse.Marshal(b, m, deterministic)
}
func (m *GetBalanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBalanceResponse.Merge(m, src)
}
func (m *GetBalanceResponse) XXX_Size() int {
	return xxx_messageInfo_GetBalanceResponse.Size(m)
}
func (m *GetBalanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBalanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBalanceResponse proto.InternalMessageInfo

func (m *GetBalanceResponse) GetBalance() *BigInt {
	if m != nil {
		return m.Balance
	}
	return nil
}

type GetUTXOsRequest struct {
	Address              []byte   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUTXOsRequest) Reset()         { *m = GetUTXOsRequest{} }
func (m *GetUTXOsRequest) String() string { return proto.CompactTextString(m) }
func (*GetUTXOsRequest) ProtoMessage()    {}
func (*GetUTXOsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_08a043f6ee9336a8, []int{8}
}
func (m *GetUTXOsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUTXOsRequest.Unmarshal(m, b)
}
func (m *GetUTXOsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUTXOsRequest.Marshal(b, m, deterministic)
}
func (m *GetUTXOsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUTXOsRequest.Merge(m, src)
}
func (m *GetUTXOsRequest) XXX_Size() int {
	return xxx_messageInfo_GetUTXOsRequest.Size(m)
}
func (m *GetUTXOsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUTXOsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUTXOsRequest proto.InternalMessageInfo

func (m *GetUTXOsRequest) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

type GetUTXOsResponse struct {
	Transactions         []*Transaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetUTXOsResponse) Reset()         { *m = GetUTXOsResponse{} }
func (m *GetUTXOsResponse) String() string { return proto.CompactTextString(m) }
func (*GetUTXOsResponse) ProtoMessage()    {}
func (*GetUTXOsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_08a043f6ee9336a8, []int{9}
}
func (m *GetUTXOsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUTXOsResponse.Unmarshal(m, b)
}
func (m *GetUTXOsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUTXOsResponse.Marshal(b, m, deterministic)
}
func (m *GetUTXOsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUTXOsResponse.Merge(m, src)
}
func (m *GetUTXOsResponse) XXX_Size() int {
	return xxx_messageInfo_GetUTXOsResponse.Size(m)
}
func (m *GetUTXOsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUTXOsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUTXOsResponse proto.InternalMessageInfo

func (m *GetUTXOsResponse) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

type GetBlockRequest struct {
	Number               uint64   `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlockRequest) Reset()         { *m = GetBlockRequest{} }
func (m *GetBlockRequest) String() string { return proto.CompactTextString(m) }
func (*GetBlockRequest) ProtoMessage()    {}
func (*GetBlockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_08a043f6ee9336a8, []int{10}
}
func (m *GetBlockRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockRequest.Unmarshal(m, b)
}
func (m *GetBlockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockRequest.Marshal(b, m, deterministic)
}
func (m *GetBlockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockRequest.Merge(m, src)
}
func (m *GetBlockRequest) XXX_Size() int {
	return xxx_messageInfo_GetBlockRequest.Size(m)
}
func (m *GetBlockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockRequest proto.InternalMessageInfo

func (m *GetBlockRequest) GetNumber() uint64 {
	if m != nil {
		return m.Number
	}
	return 0
}

type GetBlockResponse struct {
	Block                *Block         `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	Transactions         []*Transaction `protobuf:"bytes,2,rep,name=transactions,proto3" json:"transactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetBlockResponse) Reset()         { *m = GetBlockResponse{} }
func (m *GetBlockResponse) String() string { return proto.CompactTextString(m) }
func (*GetBlockResponse) ProtoMessage()    {}
func (*GetBlockResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_08a043f6ee9336a8, []int{11}
}
func (m *GetBlockResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockResponse.Unmarshal(m, b)
}
func (m *GetBlockResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockResponse.Marshal(b, m, deterministic)
}
func (m *GetBlockResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockResponse.Merge(m, src)
}
func (m *GetBlockResponse) XXX_Size() int {
	return xxx_messageInfo_GetBlockResponse.Size(m)
}
func (m *GetBlockResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockResponse proto.InternalMessageInfo

func (m *GetBlockResponse) GetBlock() *Block {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *GetBlockResponse) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

type SendRequest struct {
	Transaction          *Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	From                 []byte       `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To                   []byte       `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Amount               *BigInt      `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SendRequest) Reset()         { *m = SendRequest{} }
func (m *SendRequest) String() string { return proto.CompactTextString(m) }
func (*SendRequest) ProtoMessage()    {}
func (*SendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_08a043f6ee9336a8, []int{12}
}
func (m *SendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendRequest.Unmarshal(m, b)
}
func (m *SendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendRequest.Marshal(b, m, deterministic)
}
func (m *SendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendRequest.Merge(m, src)
}
func (m *SendRequest) XXX_Size() int {
	return xxx_messageInfo_SendRequest.Size(m)
}
func (m *SendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendRequest proto.InternalMessageInfo

func (m *SendRequest) GetTransaction() *Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func (m *SendRequest) GetFrom() []byte {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *SendRequest) GetTo() []byte {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *SendRequest) GetAmount() *BigInt {
	if m != nil {
		return m.Amount
	}
	return nil
}

type SendResponse struct {
	Transaction          *Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SendResponse) Reset()         { *m = SendResponse{} }
func (m *SendResponse) String() string { return proto.CompactTextString(m) }
func (*SendResponse) ProtoMessage()    {}
func (*SendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_08a043f6ee9336a8, []int{13}
}
func (m *SendResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendResponse.Unmarshal(m, b)
}
func (m *SendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendResponse.Marshal(b, m, deterministic)
}
func (m *SendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendResponse.Merge(m, src)
}
func (m *SendResponse) XXX_Size() int {
	return xxx_messageInfo_SendResponse.Size(m)
}
func (m *SendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendResponse proto.InternalMessageInfo

func (m *SendResponse) GetTransaction() *Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func init() {
	proto.RegisterType((*BigInt)(nil), "pb.BigInt")
	proto.RegisterType((*Input)(nil), "pb.Input")
	proto.RegisterType((*Output)(nil), "pb.Output")
	proto.RegisterType((*BlockHeader)(nil), "pb.BlockHeader")
	proto.RegisterType((*Block)(nil), "pb.Block")
	proto.RegisterType((*Transaction)(nil), "pb.Transaction")
	proto.RegisterType((*GetBalanceRequest)(nil), "pb.GetBalanceRequest")
	proto.RegisterType((*GetBalanceResponse)(nil), "pb.GetBalanceResponse")
	proto.RegisterType((*GetUTXOsRequest)(nil), "pb.GetUTXOsRequest")
	proto.RegisterType((*GetUTXOsResponse)(nil), "pb.GetUTXOsResponse")
	proto.RegisterType((*GetBlockRequest)(nil), "pb.GetBlockRequest")
	proto.RegisterType((*GetBlockResponse)(nil), "pb.GetBlockResponse")
	proto.RegisterType((*SendRequest)(nil), "pb.SendRequest")
	proto.RegisterType((*SendResponse)(nil), "pb.SendResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RootClient is the client API for Root service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RootClient interface {
	GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error)
	GetUTXOs(ctx context.Context, in *GetUTXOsRequest, opts ...grpc.CallOption) (*GetUTXOsResponse, error)
	GetBlock(ctx context.Context, in *GetBlockRequest, opts ...grpc.CallOption) (*GetBlockResponse, error)
	Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error)
}

type rootClient struct {
	cc *grpc.ClientConn
}

func NewRootClient(cc *grpc.ClientConn) RootClient {
	return &rootClient{cc}
}

func (c *rootClient) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error) {
	out := new(GetBalanceResponse)
	err := c.cc.Invoke(ctx, "/pb.Root/GetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rootClient) GetUTXOs(ctx context.Context, in *GetUTXOsRequest, opts ...grpc.CallOption) (*GetUTXOsResponse, error) {
	out := new(GetUTXOsResponse)
	err := c.cc.Invoke(ctx, "/pb.Root/GetUTXOs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rootClient) GetBlock(ctx context.Context, in *GetBlockRequest, opts ...grpc.CallOption) (*GetBlockResponse, error) {
	out := new(GetBlockResponse)
	err := c.cc.Invoke(ctx, "/pb.Root/GetBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rootClient) Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/pb.Root/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RootServer is the server API for Root service.
type RootServer interface {
	GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error)
	GetUTXOs(context.Context, *GetUTXOsRequest) (*GetUTXOsResponse, error)
	GetBlock(context.Context, *GetBlockRequest) (*GetBlockResponse, error)
	Send(context.Context, *SendRequest) (*SendResponse, error)
}

func RegisterRootServer(s *grpc.Server, srv RootServer) {
	s.RegisterService(&_Root_serviceDesc, srv)
}

func _Root_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RootServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Root/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RootServer).GetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Root_GetUTXOs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUTXOsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RootServer).GetUTXOs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Root/GetUTXOs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RootServer).GetUTXOs(ctx, req.(*GetUTXOsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Root_GetBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RootServer).GetBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Root/GetBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RootServer).GetBlock(ctx, req.(*GetBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Root_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RootServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Root/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RootServer).Send(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Root_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Root",
	HandlerType: (*RootServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBalance",
			Handler:    _Root_GetBalance_Handler,
		},
		{
			MethodName: "GetUTXOs",
			Handler:    _Root_GetUTXOs_Handler,
		},
		{
			MethodName: "GetBlock",
			Handler:    _Root_GetBlock_Handler,
		},
		{
			MethodName: "Send",
			Handler:    _Root_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "root.proto",
}

func init() { proto.RegisterFile("root.proto", fileDescriptor_08a043f6ee9336a8) }

var fileDescriptor_08a043f6ee9336a8 = []byte{
	// 629 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0x5d, 0xc7, 0x6d, 0xc7, 0x29, 0x2d, 0x4b, 0xa9, 0xac, 0x08, 0x41, 0x59, 0x55, 0xa2,
	0xa8, 0xa2, 0x8a, 0xd3, 0x03, 0x12, 0x12, 0x07, 0x22, 0xa4, 0x24, 0x07, 0x88, 0x30, 0x41, 0xe2,
	0x6a, 0x27, 0xdb, 0x26, 0x6a, 0xe2, 0x35, 0xf6, 0xba, 0xed, 0x1f, 0xc0, 0x1f, 0xf0, 0x83, 0x7c,
	0x08, 0xda, 0xd9, 0xdd, 0x78, 0xd3, 0x50, 0x01, 0x37, 0xcf, 0xcc, 0x9b, 0x99, 0xb7, 0x33, 0x6f,
	0x0c, 0x50, 0x70, 0x2e, 0xce, 0xf2, 0x82, 0x0b, 0x4e, 0xdc, 0x3c, 0xa5, 0x47, 0xe0, 0x77, 0x67,
	0x97, 0x83, 0x4c, 0x90, 0x43, 0xf0, 0xaf, 0x93, 0x79, 0xc5, 0xca, 0xd0, 0x39, 0x72, 0x4e, 0x9a,
	0xb1, 0xb6, 0xe8, 0x27, 0x68, 0x0c, 0xb2, 0xbc, 0x12, 0xa4, 0x05, 0xdb, 0xe9, 0x9c, 0x8f, 0xaf,
	0x3e, 0x56, 0x0b, 0x84, 0x78, 0xf1, 0xd2, 0x26, 0x07, 0xd0, 0x10, 0xb7, 0x83, 0xc9, 0x6d, 0xe8,
	0x1e, 0x39, 0x27, 0xbb, 0xb1, 0x32, 0x64, 0x49, 0x5e, 0x09, 0xe9, 0xde, 0x44, 0xb7, 0xb6, 0x68,
	0x1f, 0xfc, 0x61, 0x25, 0x74, 0xcd, 0x8c, 0xdd, 0x0c, 0x6f, 0x32, 0x56, 0xe8, 0xb6, 0x4b, 0x9b,
	0x50, 0xf0, 0x93, 0x05, 0xaf, 0x32, 0x81, 0x45, 0x83, 0x0e, 0x9c, 0xe5, 0xe9, 0x99, 0x22, 0x1b,
	0xeb, 0x08, 0xfd, 0xee, 0x40, 0xd0, 0x95, 0x24, 0xfa, 0x2c, 0x99, 0xb0, 0x82, 0x3c, 0x05, 0x58,
	0xb0, 0xe2, 0x6a, 0xce, 0x62, 0xce, 0x85, 0xae, 0x68, 0x79, 0xc8, 0x31, 0xec, 0x16, 0xf3, 0xfc,
	0x43, 0x0d, 0x71, 0x11, 0xb2, 0xea, 0x94, 0xac, 0xf2, 0x82, 0x5d, 0xf7, 0x93, 0x72, 0x8a, 0xcc,
	0x9b, 0xf1, 0xd2, 0x96, 0x6f, 0xca, 0xaa, 0x45, 0xca, 0x8a, 0xd0, 0xc3, 0x19, 0x68, 0x8b, 0xbe,
	0x87, 0x06, 0x12, 0x21, 0x2f, 0xc0, 0x9f, 0x22, 0x19, 0x6c, 0x1f, 0x74, 0xf6, 0x90, 0x76, 0xcd,
	0x31, 0xd6, 0x61, 0x42, 0xc0, 0x9b, 0xca, 0x0e, 0x8a, 0x02, 0x7e, 0xd3, 0x9f, 0x2e, 0x04, 0xa3,
	0x22, 0xc9, 0xca, 0x64, 0x2c, 0x66, 0x3c, 0x23, 0xcf, 0xc1, 0x9f, 0xc9, 0xe1, 0xb7, 0x75, 0xb1,
	0x1d, 0x59, 0x0c, 0xd7, 0x11, 0xeb, 0x80, 0x2c, 0x53, 0xce, 0x2e, 0xdb, 0xa6, 0x8c, 0xfc, 0x5e,
	0xa6, 0x45, 0x48, 0xff, 0x0f, 0x69, 0x91, 0x4e, 0x8b, 0xf0, 0x15, 0x2a, 0x2d, 0x22, 0xc7, 0xb0,
	0xc5, 0x71, 0x2f, 0xed, 0xb0, 0x51, 0x8f, 0x5c, 0xad, 0x2a, 0x36, 0xa1, 0x1a, 0x15, 0x85, 0xfe,
	0x7d, 0xa8, 0x88, 0x3c, 0x81, 0xcd, 0x0b, 0xc6, 0xc2, 0xad, 0xb5, 0xd5, 0x49, 0xb7, 0x9c, 0x70,
	0xd7, 0x68, 0x69, 0x5b, 0x69, 0xa9, 0x6b, 0x69, 0x69, 0x84, 0x5a, 0xda, 0x51, 0x5a, 0x42, 0x83,
	0xbe, 0x82, 0x87, 0x3d, 0x26, 0xba, 0xc9, 0x3c, 0xc9, 0xc6, 0x2c, 0x66, 0xdf, 0x2a, 0x56, 0x0a,
	0x12, 0xc2, 0x56, 0x32, 0x99, 0x14, 0xac, 0x34, 0xa2, 0x35, 0x26, 0x7d, 0x03, 0xc4, 0x86, 0x97,
	0x39, 0xcf, 0x4a, 0x26, 0xa9, 0xa7, 0xca, 0xa5, 0xe7, 0x69, 0x13, 0x33, 0x21, 0x7a, 0x0a, 0x7b,
	0x3d, 0x26, 0xbe, 0x8c, 0xbe, 0x0e, 0xcb, 0xbf, 0x37, 0xea, 0xc1, 0x7e, 0x0d, 0xd6, 0x6d, 0xce,
	0xa1, 0x29, 0xea, 0x25, 0xca, 0x94, 0x4d, 0x23, 0x04, 0x6b, 0xb9, 0xf1, 0x0a, 0x88, 0xbe, 0xc4,
	0xae, 0x38, 0x05, 0xd3, 0xb5, 0xd6, 0x9a, 0xb3, 0xa2, 0xb5, 0x29, 0xf6, 0xd4, 0x50, 0xdd, 0xf3,
	0x19, 0x34, 0xf0, 0x1a, 0x6d, 0xa1, 0x28, 0x84, 0xf2, 0xaf, 0x91, 0x72, 0xff, 0x85, 0xd4, 0x0f,
	0x07, 0x82, 0xcf, 0x2c, 0x9b, 0x18, 0x46, 0x11, 0x04, 0x56, 0xdc, 0x56, 0xb8, 0x5d, 0xc3, 0xc6,
	0x48, 0xa1, 0x5d, 0x14, 0x7c, 0x61, 0xf4, 0x29, 0xbf, 0xc9, 0x03, 0x70, 0x05, 0xd7, 0xa7, 0xe5,
	0x0a, 0x6e, 0x9d, 0xba, 0x77, 0xef, 0xa9, 0xbf, 0x83, 0xa6, 0x62, 0xa2, 0x1f, 0xfc, 0xff, 0x54,
	0x3a, 0xbf, 0x1c, 0xf0, 0xf0, 0xc0, 0xdf, 0x02, 0xd4, 0xea, 0x20, 0x8f, 0x65, 0xd2, 0x9a, 0xb8,
	0x5a, 0x87, 0x77, 0xdd, 0xaa, 0x31, 0xdd, 0x20, 0xaf, 0x61, 0xdb, 0xec, 0x9c, 0x3c, 0xd2, 0x28,
	0x5b, 0x2e, 0xad, 0x83, 0x55, 0xe7, 0x9d, 0x44, 0xf5, 0x9f, 0x30, 0x89, 0xf6, 0xc6, 0x97, 0x89,
	0x2b, 0xbb, 0xa5, 0x1b, 0xe4, 0x14, 0x3c, 0xf9, 0x78, 0x82, 0xef, 0xb3, 0x16, 0xd2, 0xda, 0xaf,
	0x1d, 0x06, 0x9c, 0xfa, 0xf8, 0x7b, 0x3f, 0xff, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x66, 0x31, 0x1f,
	0xb0, 0xec, 0x05, 0x00, 0x00,
}

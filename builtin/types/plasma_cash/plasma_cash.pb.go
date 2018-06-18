// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/loomnetwork/go-loom/builtin/types/plasma_cash/plasma_cash.proto

/*
Package plasma_cash is a generated protocol buffer package.

It is generated from these files:
	github.com/loomnetwork/go-loom/builtin/types/plasma_cash/plasma_cash.proto

It has these top-level messages:
	PlasmaBlock
	PlasmaTx
	ListTransactionsRequest
	ListTransactionsResponse
	SubmitBlockToMainnetRequest
	SubmitBlockToMainnetResponse
	PlasmaTxRequest
	PlasmaTxResponse
	DepositRequest
	DepositResponse
	GetProofRequest
	GetProofResponse
	Params
	PlasmaCashInitRequest
*/
package plasma_cash

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import types "github.com/loomnetwork/go-loom/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type PlasmaBlock struct {
	Uid          *types.BigUInt `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
	Transactions []*PlasmaTx    `protobuf:"bytes,2,rep,name=transactions" json:"transactions,omitempty"`
	Signature    []byte         `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	MerkleHash   []byte         `protobuf:"bytes,4,opt,name=merkle_hash,json=merkleHash,proto3" json:"merkle_hash,omitempty"`
	Hash         []byte         `protobuf:"bytes,5,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *PlasmaBlock) Reset()                    { *m = PlasmaBlock{} }
func (m *PlasmaBlock) String() string            { return proto.CompactTextString(m) }
func (*PlasmaBlock) ProtoMessage()               {}
func (*PlasmaBlock) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{0} }

func (m *PlasmaBlock) GetUid() *types.BigUInt {
	if m != nil {
		return m.Uid
	}
	return nil
}

func (m *PlasmaBlock) GetTransactions() []*PlasmaTx {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func (m *PlasmaBlock) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *PlasmaBlock) GetMerkleHash() []byte {
	if m != nil {
		return m.MerkleHash
	}
	return nil
}

func (m *PlasmaBlock) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type PlasmaTx struct {
	Uid           uint64         `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	PreviousBlock *types.BigUInt `protobuf:"bytes,2,opt,name=previous_block,json=previousBlock" json:"previous_block,omitempty"`
	Denomination  *types.BigUInt `protobuf:"bytes,3,opt,name=denomination" json:"denomination,omitempty"`
	NewOwner      *types.Address `protobuf:"bytes,4,opt,name=new_owner,json=newOwner" json:"new_owner,omitempty"`
	Signature     []byte         `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
	Hash          []byte         `protobuf:"bytes,6,opt,name=hash,proto3" json:"hash,omitempty"`
	MerkleHash    []byte         `protobuf:"bytes,7,opt,name=merkle_hash,json=merkleHash,proto3" json:"merkle_hash,omitempty"`
	Sender        *types.Address `protobuf:"bytes,8,opt,name=sender" json:"sender,omitempty"`
}

func (m *PlasmaTx) Reset()                    { *m = PlasmaTx{} }
func (m *PlasmaTx) String() string            { return proto.CompactTextString(m) }
func (*PlasmaTx) ProtoMessage()               {}
func (*PlasmaTx) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{1} }

func (m *PlasmaTx) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *PlasmaTx) GetPreviousBlock() *types.BigUInt {
	if m != nil {
		return m.PreviousBlock
	}
	return nil
}

func (m *PlasmaTx) GetDenomination() *types.BigUInt {
	if m != nil {
		return m.Denomination
	}
	return nil
}

func (m *PlasmaTx) GetNewOwner() *types.Address {
	if m != nil {
		return m.NewOwner
	}
	return nil
}

func (m *PlasmaTx) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *PlasmaTx) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *PlasmaTx) GetMerkleHash() []byte {
	if m != nil {
		return m.MerkleHash
	}
	return nil
}

func (m *PlasmaTx) GetSender() *types.Address {
	if m != nil {
		return m.Sender
	}
	return nil
}

type ListTransactionsRequest struct {
	BlockHeight *types.BigUInt `protobuf:"bytes,1,opt,name=block_height,json=blockHeight" json:"block_height,omitempty"`
}

func (m *ListTransactionsRequest) Reset()         { *m = ListTransactionsRequest{} }
func (m *ListTransactionsRequest) String() string { return proto.CompactTextString(m) }
func (*ListTransactionsRequest) ProtoMessage()    {}
func (*ListTransactionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorPlasmaCash, []int{2}
}

func (m *ListTransactionsRequest) GetBlockHeight() *types.BigUInt {
	if m != nil {
		return m.BlockHeight
	}
	return nil
}

type ListTransactionsResponse struct {
	Transactions []*PlasmaTx `protobuf:"bytes,1,rep,name=transactions" json:"transactions,omitempty"`
}

func (m *ListTransactionsResponse) Reset()         { *m = ListTransactionsResponse{} }
func (m *ListTransactionsResponse) String() string { return proto.CompactTextString(m) }
func (*ListTransactionsResponse) ProtoMessage()    {}
func (*ListTransactionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorPlasmaCash, []int{3}
}

func (m *ListTransactionsResponse) GetTransactions() []*PlasmaTx {
	if m != nil {
		return m.Transactions
	}
	return nil
}

// This only originates from the validator
type SubmitBlockToMainnetRequest struct {
}

func (m *SubmitBlockToMainnetRequest) Reset()         { *m = SubmitBlockToMainnetRequest{} }
func (m *SubmitBlockToMainnetRequest) String() string { return proto.CompactTextString(m) }
func (*SubmitBlockToMainnetRequest) ProtoMessage()    {}
func (*SubmitBlockToMainnetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorPlasmaCash, []int{4}
}

type SubmitBlockToMainnetResponse struct {
}

func (m *SubmitBlockToMainnetResponse) Reset()         { *m = SubmitBlockToMainnetResponse{} }
func (m *SubmitBlockToMainnetResponse) String() string { return proto.CompactTextString(m) }
func (*SubmitBlockToMainnetResponse) ProtoMessage()    {}
func (*SubmitBlockToMainnetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorPlasmaCash, []int{5}
}

type PlasmaTxRequest struct {
	Plasmatx *PlasmaTx `protobuf:"bytes,1,opt,name=plasmatx" json:"plasmatx,omitempty"`
}

func (m *PlasmaTxRequest) Reset()                    { *m = PlasmaTxRequest{} }
func (m *PlasmaTxRequest) String() string            { return proto.CompactTextString(m) }
func (*PlasmaTxRequest) ProtoMessage()               {}
func (*PlasmaTxRequest) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{6} }

func (m *PlasmaTxRequest) GetPlasmatx() *PlasmaTx {
	if m != nil {
		return m.Plasmatx
	}
	return nil
}

type PlasmaTxResponse struct {
}

func (m *PlasmaTxResponse) Reset()                    { *m = PlasmaTxResponse{} }
func (m *PlasmaTxResponse) String() string            { return proto.CompactTextString(m) }
func (*PlasmaTxResponse) ProtoMessage()               {}
func (*PlasmaTxResponse) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{7} }

// This only originates from the validator
type DepositRequest struct {
	Plasmatx *PlasmaTx `protobuf:"bytes,1,opt,name=plasmatx" json:"plasmatx,omitempty"`
}

func (m *DepositRequest) Reset()                    { *m = DepositRequest{} }
func (m *DepositRequest) String() string            { return proto.CompactTextString(m) }
func (*DepositRequest) ProtoMessage()               {}
func (*DepositRequest) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{8} }

func (m *DepositRequest) GetPlasmatx() *PlasmaTx {
	if m != nil {
		return m.Plasmatx
	}
	return nil
}

type DepositResponse struct {
}

func (m *DepositResponse) Reset()                    { *m = DepositResponse{} }
func (m *DepositResponse) String() string            { return proto.CompactTextString(m) }
func (*DepositResponse) ProtoMessage()               {}
func (*DepositResponse) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{9} }

type GetProofRequest struct {
	BlockHeight *types.BigUInt `protobuf:"bytes,1,opt,name=block_height,json=blockHeight" json:"block_height,omitempty"`
	Uid         *types.BigUInt `protobuf:"bytes,2,opt,name=uid" json:"uid,omitempty"`
}

func (m *GetProofRequest) Reset()                    { *m = GetProofRequest{} }
func (m *GetProofRequest) String() string            { return proto.CompactTextString(m) }
func (*GetProofRequest) ProtoMessage()               {}
func (*GetProofRequest) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{10} }

func (m *GetProofRequest) GetBlockHeight() *types.BigUInt {
	if m != nil {
		return m.BlockHeight
	}
	return nil
}

func (m *GetProofRequest) GetUid() *types.BigUInt {
	if m != nil {
		return m.Uid
	}
	return nil
}

type GetProofResponse struct {
	Proof []byte `protobuf:"bytes,1,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (m *GetProofResponse) Reset()                    { *m = GetProofResponse{} }
func (m *GetProofResponse) String() string            { return proto.CompactTextString(m) }
func (*GetProofResponse) ProtoMessage()               {}
func (*GetProofResponse) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{11} }

func (m *GetProofResponse) GetProof() []byte {
	if m != nil {
		return m.Proof
	}
	return nil
}

// Initialization of state from Genesis.json
type Params struct {
	BlockInterval uint64 `protobuf:"varint,1,opt,name=block_interval,json=blockInterval,proto3" json:"block_interval,omitempty"`
}

func (m *Params) Reset()                    { *m = Params{} }
func (m *Params) String() string            { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()               {}
func (*Params) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{12} }

func (m *Params) GetBlockInterval() uint64 {
	if m != nil {
		return m.BlockInterval
	}
	return 0
}

type PlasmaCashInitRequest struct {
	Params *Params `protobuf:"bytes,1,opt,name=params" json:"params,omitempty"`
}

func (m *PlasmaCashInitRequest) Reset()                    { *m = PlasmaCashInitRequest{} }
func (m *PlasmaCashInitRequest) String() string            { return proto.CompactTextString(m) }
func (*PlasmaCashInitRequest) ProtoMessage()               {}
func (*PlasmaCashInitRequest) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{13} }

func (m *PlasmaCashInitRequest) GetParams() *Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func init() {
	proto.RegisterType((*PlasmaBlock)(nil), "PlasmaBlock")
	proto.RegisterType((*PlasmaTx)(nil), "PlasmaTx")
	proto.RegisterType((*ListTransactionsRequest)(nil), "ListTransactionsRequest")
	proto.RegisterType((*ListTransactionsResponse)(nil), "ListTransactionsResponse")
	proto.RegisterType((*SubmitBlockToMainnetRequest)(nil), "SubmitBlockToMainnetRequest")
	proto.RegisterType((*SubmitBlockToMainnetResponse)(nil), "SubmitBlockToMainnetResponse")
	proto.RegisterType((*PlasmaTxRequest)(nil), "PlasmaTxRequest")
	proto.RegisterType((*PlasmaTxResponse)(nil), "PlasmaTxResponse")
	proto.RegisterType((*DepositRequest)(nil), "DepositRequest")
	proto.RegisterType((*DepositResponse)(nil), "DepositResponse")
	proto.RegisterType((*GetProofRequest)(nil), "GetProofRequest")
	proto.RegisterType((*GetProofResponse)(nil), "GetProofResponse")
	proto.RegisterType((*Params)(nil), "Params")
	proto.RegisterType((*PlasmaCashInitRequest)(nil), "PlasmaCashInitRequest")
}

func init() {
	proto.RegisterFile("github.com/loomnetwork/go-loom/builtin/types/plasma_cash/plasma_cash.proto", fileDescriptorPlasmaCash)
}

var fileDescriptorPlasmaCash = []byte{
	// 589 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x18, 0x55, 0xbb, 0xad, 0xcb, 0xbe, 0x66, 0x7f, 0x16, 0x88, 0x68, 0x0c, 0x5a, 0x45, 0x9a, 0x54,
	0x09, 0xda, 0xa0, 0xed, 0x82, 0xdd, 0x52, 0x10, 0xac, 0x88, 0x9f, 0x2a, 0x94, 0x1b, 0x6e, 0x2a,
	0xa7, 0xf5, 0x12, 0xab, 0x89, 0x1d, 0x6c, 0x67, 0x1d, 0xcf, 0xc4, 0x0b, 0xf0, 0x34, 0xbb, 0xd8,
	0x93, 0xa0, 0x38, 0x6e, 0x93, 0xae, 0x13, 0x08, 0x6e, 0x2a, 0x7f, 0xe7, 0xd8, 0xe7, 0xfb, 0xce,
	0x89, 0x5d, 0x78, 0x1f, 0x52, 0x15, 0x65, 0x41, 0x6f, 0xc2, 0x13, 0x2f, 0xe6, 0x3c, 0x61, 0x44,
	0xcd, 0xb9, 0x98, 0x79, 0x21, 0xef, 0xe6, 0xa5, 0x17, 0x64, 0x34, 0x56, 0x94, 0x79, 0xea, 0x47,
	0x4a, 0xa4, 0x97, 0xc6, 0x58, 0x26, 0x78, 0x3c, 0xc1, 0x32, 0xaa, 0xae, 0x7b, 0xa9, 0xe0, 0x8a,
	0x1f, 0x75, 0x2b, 0x5a, 0x21, 0x0f, 0xb9, 0xa7, 0xe1, 0x20, 0xbb, 0xd4, 0x95, 0x2e, 0xf4, 0xca,
	0x6c, 0x7f, 0xf1, 0x97, 0xd6, 0x45, 0x4b, 0xfd, 0x5b, 0x9c, 0x70, 0x7f, 0xd6, 0xa0, 0x39, 0xd4,
	0x6d, 0xfb, 0x31, 0x9f, 0xcc, 0xd0, 0x11, 0x6c, 0x64, 0x74, 0xea, 0xd4, 0xda, 0xb5, 0x4e, 0xf3,
	0xd4, 0xea, 0xf5, 0x69, 0xf8, 0x75, 0xc0, 0x94, 0x9f, 0x83, 0xa8, 0x0b, 0xb6, 0x12, 0x98, 0x49,
	0x3c, 0x51, 0x94, 0x33, 0xe9, 0xd4, 0xdb, 0x1b, 0x9d, 0xe6, 0xe9, 0x4e, 0xaf, 0x38, 0x3f, 0xba,
	0xf6, 0x57, 0x68, 0x74, 0x0c, 0x3b, 0x92, 0x86, 0x0c, 0xab, 0x4c, 0x10, 0x67, 0xa3, 0x5d, 0xeb,
	0xd8, 0x7e, 0x09, 0xa0, 0x16, 0x34, 0x13, 0x22, 0x66, 0x31, 0x19, 0x47, 0x58, 0x46, 0xce, 0xa6,
	0xe6, 0xa1, 0x80, 0x2e, 0xb0, 0x8c, 0x10, 0x82, 0x4d, 0xcd, 0x6c, 0x69, 0x46, 0xaf, 0xdd, 0x5f,
	0x75, 0xb0, 0x16, 0xdd, 0xd0, 0x41, 0x39, 0xea, 0x66, 0x31, 0x60, 0x1f, 0xf6, 0x52, 0x41, 0xae,
	0x28, 0xcf, 0xe4, 0x38, 0xc8, 0xed, 0x38, 0xf5, 0x55, 0x1f, 0xfd, 0xc3, 0xdb, 0x9b, 0xd6, 0xee,
	0xd0, 0xec, 0xd1, 0x8e, 0xfd, 0xdd, 0xb4, 0x5a, 0xa2, 0xe7, 0x60, 0x4f, 0x09, 0xe3, 0x09, 0x65,
	0x38, 0xb7, 0xa1, 0x07, 0xaf, 0x26, 0xb1, 0xc2, 0xa2, 0x33, 0xd8, 0x61, 0x64, 0x3e, 0xe6, 0x73,
	0x46, 0x84, 0xf6, 0x90, 0x6f, 0x7d, 0x35, 0x9d, 0x0a, 0x22, 0x65, 0xdf, 0xbe, 0xbd, 0x69, 0x59,
	0x9f, 0xc8, 0xfc, 0x73, 0xce, 0xfa, 0x16, 0x33, 0xab, 0xd5, 0x60, 0xb6, 0xee, 0x06, 0xb3, 0xf0,
	0xdd, 0x28, 0x7d, 0xdf, 0x0d, 0x6b, 0x7b, 0x2d, 0xac, 0x36, 0x34, 0x24, 0x61, 0x53, 0x22, 0x1c,
	0x6b, 0x75, 0x08, 0xdf, 0xe0, 0xee, 0x5b, 0x78, 0xf4, 0x81, 0x4a, 0x35, 0xaa, 0x7c, 0x21, 0x9f,
	0x7c, 0xcf, 0x88, 0x54, 0xe8, 0x19, 0xd8, 0x3a, 0xad, 0x71, 0x44, 0x68, 0x18, 0xa9, 0xb5, 0x8f,
	0xdf, 0xd4, 0xec, 0x85, 0x26, 0xdd, 0x01, 0x38, 0xeb, 0x3a, 0x32, 0xe5, 0x4c, 0x92, 0xb5, 0x0b,
	0x52, 0xfb, 0xe3, 0x05, 0x71, 0x9f, 0xc0, 0xe3, 0x2f, 0x59, 0x90, 0x50, 0xa5, 0x93, 0x1f, 0xf1,
	0x8f, 0x98, 0x32, 0x46, 0x94, 0x19, 0xcb, 0x7d, 0x0a, 0xc7, 0xf7, 0xd3, 0x45, 0x37, 0xf7, 0x1c,
	0xf6, 0x97, 0xc2, 0xc6, 0xc9, 0x09, 0x58, 0xc5, 0x1b, 0x52, 0xd7, 0xc6, 0x45, 0xa5, 0xf9, 0x92,
	0x72, 0x11, 0x1c, 0x94, 0x27, 0x8d, 0xda, 0x4b, 0xd8, 0x7b, 0x43, 0x52, 0x2e, 0xa9, 0xfa, 0x47,
	0xb1, 0x43, 0xd8, 0x5f, 0x1e, 0x34, 0x5a, 0xdf, 0x60, 0xff, 0x1d, 0x51, 0x43, 0xc1, 0xf9, 0xe5,
	0xff, 0x64, 0xbc, 0x78, 0x84, 0xf5, 0x7b, 0x1e, 0xa1, 0xdb, 0x81, 0x83, 0x52, 0xdb, 0xe4, 0xfe,
	0x00, 0xb6, 0xd2, 0x1c, 0xd0, 0xaa, 0xb6, 0x5f, 0x14, 0xae, 0x07, 0x8d, 0x21, 0x16, 0x38, 0x91,
	0xe8, 0x04, 0xf6, 0x8a, 0xe6, 0x94, 0x29, 0x22, 0xae, 0x70, 0x6c, 0x1e, 0xcd, 0xae, 0x46, 0x07,
	0x06, 0x74, 0xcf, 0xe1, 0x61, 0xe1, 0xef, 0x35, 0x96, 0xd1, 0x80, 0x95, 0x49, 0xb4, 0xa0, 0x91,
	0x6a, 0x25, 0x33, 0xf6, 0x76, 0xaf, 0x10, 0xf6, 0x0d, 0x1c, 0x34, 0xf4, 0x9f, 0xc9, 0xd9, 0xef,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x41, 0xda, 0x3b, 0xaa, 0xfb, 0x04, 0x00, 0x00,
}

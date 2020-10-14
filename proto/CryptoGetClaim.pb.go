// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.3
// source: proto/CryptoGetClaim.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Get a single claim attached to an account, or return null if it does not exist.
type CryptoGetClaimQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header    *QueryHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`       // Standard info sent from client to node, including the signed payment, and what kind of response is requested (cost, state proof, both, or neither).
	AccountID *AccountID   `protobuf:"bytes,2,opt,name=accountID,proto3" json:"accountID,omitempty"` // The account ID to which the claim was attached
	Hash      []byte       `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`           // The hash of the claim
}

func (x *CryptoGetClaimQuery) Reset() {
	*x = CryptoGetClaimQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_CryptoGetClaim_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptoGetClaimQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptoGetClaimQuery) ProtoMessage() {}

func (x *CryptoGetClaimQuery) ProtoReflect() protoreflect.Message {
	mi := &file_proto_CryptoGetClaim_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptoGetClaimQuery.ProtoReflect.Descriptor instead.
func (*CryptoGetClaimQuery) Descriptor() ([]byte, []int) {
	return file_proto_CryptoGetClaim_proto_rawDescGZIP(), []int{0}
}

func (x *CryptoGetClaimQuery) GetHeader() *QueryHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *CryptoGetClaimQuery) GetAccountID() *AccountID {
	if x != nil {
		return x.AccountID
	}
	return nil
}

func (x *CryptoGetClaimQuery) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

// Response when the client sends the node CryptoGetClaimQuery. If the claim exists, there can be a state proof for that single claim. If the claim doesn't exist, then the state proof must be obtained for the account as a whole, which lists all the attached claims, which then proves that any claim not on the list must not exist.
type CryptoGetClaimResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"` //Standard response from node to client, including the requested fields: cost, or state proof, or both, or neither
	Claim  *Claim          `protobuf:"bytes,2,opt,name=claim,proto3" json:"claim,omitempty"`   // The claim (account, hash, keys), or null if there is no Claim with the given hash attached to the given account
}

func (x *CryptoGetClaimResponse) Reset() {
	*x = CryptoGetClaimResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_CryptoGetClaim_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptoGetClaimResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptoGetClaimResponse) ProtoMessage() {}

func (x *CryptoGetClaimResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_CryptoGetClaim_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptoGetClaimResponse.ProtoReflect.Descriptor instead.
func (*CryptoGetClaimResponse) Descriptor() ([]byte, []int) {
	return file_proto_CryptoGetClaim_proto_rawDescGZIP(), []int{1}
}

func (x *CryptoGetClaimResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *CryptoGetClaimResponse) GetClaim() *Claim {
	if x != nil {
		return x.Claim
	}
	return nil
}

var File_proto_CryptoGetClaim_proto protoreflect.FileDescriptor

var file_proto_CryptoGetClaim_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x47, 0x65,
	0x74, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x42, 0x61, 0x73, 0x69, 0x63,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x41, 0x64,
	0x64, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x01, 0x0a,
	0x13, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x2a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x2e, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x22, 0x6b, 0x0a, 0x16, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x47, 0x65,
	0x74, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d,
	0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x22, 0x0a,
	0x05, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x69,
	0x6d, 0x42, 0x48, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x2e,
	0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61,
	0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x2d, 0x73,
	0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_CryptoGetClaim_proto_rawDescOnce sync.Once
	file_proto_CryptoGetClaim_proto_rawDescData = file_proto_CryptoGetClaim_proto_rawDesc
)

func file_proto_CryptoGetClaim_proto_rawDescGZIP() []byte {
	file_proto_CryptoGetClaim_proto_rawDescOnce.Do(func() {
		file_proto_CryptoGetClaim_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_CryptoGetClaim_proto_rawDescData)
	})
	return file_proto_CryptoGetClaim_proto_rawDescData
}

var file_proto_CryptoGetClaim_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_CryptoGetClaim_proto_goTypes = []interface{}{
	(*CryptoGetClaimQuery)(nil),    // 0: proto.CryptoGetClaimQuery
	(*CryptoGetClaimResponse)(nil), // 1: proto.CryptoGetClaimResponse
	(*QueryHeader)(nil),            // 2: proto.QueryHeader
	(*AccountID)(nil),              // 3: proto.AccountID
	(*ResponseHeader)(nil),         // 4: proto.ResponseHeader
	(*Claim)(nil),                  // 5: proto.Claim
}
var file_proto_CryptoGetClaim_proto_depIdxs = []int32{
	2, // 0: proto.CryptoGetClaimQuery.header:type_name -> proto.QueryHeader
	3, // 1: proto.CryptoGetClaimQuery.accountID:type_name -> proto.AccountID
	4, // 2: proto.CryptoGetClaimResponse.header:type_name -> proto.ResponseHeader
	5, // 3: proto.CryptoGetClaimResponse.claim:type_name -> proto.Claim
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_CryptoGetClaim_proto_init() }
func file_proto_CryptoGetClaim_proto_init() {
	if File_proto_CryptoGetClaim_proto != nil {
		return
	}
	file_proto_BasicTypes_proto_init()
	file_proto_QueryHeader_proto_init()
	file_proto_ResponseHeader_proto_init()
	file_proto_CryptoAddClaim_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_CryptoGetClaim_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptoGetClaimQuery); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_CryptoGetClaim_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptoGetClaimResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_CryptoGetClaim_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_CryptoGetClaim_proto_goTypes,
		DependencyIndexes: file_proto_CryptoGetClaim_proto_depIdxs,
		MessageInfos:      file_proto_CryptoGetClaim_proto_msgTypes,
	}.Build()
	File_proto_CryptoGetClaim_proto = out.File
	file_proto_CryptoGetClaim_proto_rawDesc = nil
	file_proto_CryptoGetClaim_proto_goTypes = nil
	file_proto_CryptoGetClaim_proto_depIdxs = nil
}
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: proto/TokenService.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
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

var File_proto_TokenService_proto protoreflect.FileDescriptor

var file_proto_TokenService_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe4, 0x06, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3b, 0x0a, 0x09, 0x62, 0x75, 0x72, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a,
	0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x10,
	0x77, 0x69, 0x70, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x44, 0x0a, 0x12, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x14, 0x75, 0x6e, 0x66, 0x72, 0x65, 0x65,
	0x7a, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48,
	0x0a, 0x16, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x4b, 0x79, 0x63, 0x54, 0x6f, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x19, 0x72, 0x65, 0x76, 0x6f,
	0x6b, 0x65, 0x4b, 0x79, 0x63, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0f, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x10, 0x64, 0x69, 0x73, 0x73,
	0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x12, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x0c,
	0x67, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x49, 0x0a, 0x1a, 0x63,
	0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f,
	0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x32,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_TokenService_proto_goTypes = []interface{}{
	(*Transaction)(nil),         // 0: proto.Transaction
	(*Query)(nil),               // 1: proto.Query
	(*TransactionResponse)(nil), // 2: proto.TransactionResponse
	(*Response)(nil),            // 3: proto.Response
}
var file_proto_TokenService_proto_depIdxs = []int32{
	0,  // 0: proto.TokenService.createToken:input_type -> proto.Transaction
	0,  // 1: proto.TokenService.updateToken:input_type -> proto.Transaction
	0,  // 2: proto.TokenService.mintToken:input_type -> proto.Transaction
	0,  // 3: proto.TokenService.burnToken:input_type -> proto.Transaction
	0,  // 4: proto.TokenService.deleteToken:input_type -> proto.Transaction
	0,  // 5: proto.TokenService.wipeTokenAccount:input_type -> proto.Transaction
	0,  // 6: proto.TokenService.freezeTokenAccount:input_type -> proto.Transaction
	0,  // 7: proto.TokenService.unfreezeTokenAccount:input_type -> proto.Transaction
	0,  // 8: proto.TokenService.grantKycToTokenAccount:input_type -> proto.Transaction
	0,  // 9: proto.TokenService.revokeKycFromTokenAccount:input_type -> proto.Transaction
	0,  // 10: proto.TokenService.associateTokens:input_type -> proto.Transaction
	0,  // 11: proto.TokenService.dissociateTokens:input_type -> proto.Transaction
	1,  // 12: proto.TokenService.getTokenInfo:input_type -> proto.Query
	2,  // 13: proto.TokenService.createToken:output_type -> proto.TransactionResponse
	2,  // 14: proto.TokenService.updateToken:output_type -> proto.TransactionResponse
	2,  // 15: proto.TokenService.mintToken:output_type -> proto.TransactionResponse
	2,  // 16: proto.TokenService.burnToken:output_type -> proto.TransactionResponse
	2,  // 17: proto.TokenService.deleteToken:output_type -> proto.TransactionResponse
	2,  // 18: proto.TokenService.wipeTokenAccount:output_type -> proto.TransactionResponse
	2,  // 19: proto.TokenService.freezeTokenAccount:output_type -> proto.TransactionResponse
	2,  // 20: proto.TokenService.unfreezeTokenAccount:output_type -> proto.TransactionResponse
	2,  // 21: proto.TokenService.grantKycToTokenAccount:output_type -> proto.TransactionResponse
	2,  // 22: proto.TokenService.revokeKycFromTokenAccount:output_type -> proto.TransactionResponse
	2,  // 23: proto.TokenService.associateTokens:output_type -> proto.TransactionResponse
	2,  // 24: proto.TokenService.dissociateTokens:output_type -> proto.TransactionResponse
	3,  // 25: proto.TokenService.getTokenInfo:output_type -> proto.Response
	13, // [13:26] is the sub-list for method output_type
	0,  // [0:13] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_proto_TokenService_proto_init() }
func file_proto_TokenService_proto_init() {
	if File_proto_TokenService_proto != nil {
		return
	}
	file_proto_Query_proto_init()
	file_proto_Response_proto_init()
	file_proto_TransactionResponse_proto_init()
	file_proto_Transaction_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_TokenService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_TokenService_proto_goTypes,
		DependencyIndexes: file_proto_TokenService_proto_depIdxs,
	}.Build()
	File_proto_TokenService_proto = out.File
	file_proto_TokenService_proto_rawDesc = nil
	file_proto_TokenService_proto_goTypes = nil
	file_proto_TokenService_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TokenServiceClient is the client API for TokenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TokenServiceClient interface {
	// Creates a new Token by submitting the transaction
	CreateToken(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	// Updates the account by submitting the transaction
	UpdateToken(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	// Mints an amount of the token to the defined treasury account
	MintToken(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	// Burns an amount of the token from the defined treasury account
	BurnToken(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	// Deletes a Token
	DeleteToken(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	// Wipes the provided amount of tokens from the specified Account ID
	WipeTokenAccount(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	// Freezes the transfer of tokens to or from the specified Account ID
	FreezeTokenAccount(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	// Unfreezes the transfer of tokens to or from the specified Account ID
	UnfreezeTokenAccount(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	// Flags the provided Account ID as having gone through KYC
	GrantKycToTokenAccount(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	// Removes the KYC flag of the provided Account ID
	RevokeKycFromTokenAccount(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	// Associates tokens to an account
	AssociateTokens(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	// Dissociates tokens from an account
	DissociateTokens(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	// Retrieves the metadata of a token
	GetTokenInfo(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
}

type tokenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTokenServiceClient(cc grpc.ClientConnInterface) TokenServiceClient {
	return &tokenServiceClient{cc}
}

func (c *tokenServiceClient) CreateToken(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.TokenService/createToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) UpdateToken(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.TokenService/updateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) MintToken(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.TokenService/mintToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) BurnToken(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.TokenService/burnToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) DeleteToken(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.TokenService/deleteToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) WipeTokenAccount(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.TokenService/wipeTokenAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) FreezeTokenAccount(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.TokenService/freezeTokenAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) UnfreezeTokenAccount(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.TokenService/unfreezeTokenAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) GrantKycToTokenAccount(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.TokenService/grantKycToTokenAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) RevokeKycFromTokenAccount(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.TokenService/revokeKycFromTokenAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) AssociateTokens(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.TokenService/associateTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) DissociateTokens(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.TokenService/dissociateTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) GetTokenInfo(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.TokenService/getTokenInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenServiceServer is the server API for TokenService service.
type TokenServiceServer interface {
	// Creates a new Token by submitting the transaction
	CreateToken(context.Context, *Transaction) (*TransactionResponse, error)
	// Updates the account by submitting the transaction
	UpdateToken(context.Context, *Transaction) (*TransactionResponse, error)
	// Mints an amount of the token to the defined treasury account
	MintToken(context.Context, *Transaction) (*TransactionResponse, error)
	// Burns an amount of the token from the defined treasury account
	BurnToken(context.Context, *Transaction) (*TransactionResponse, error)
	// Deletes a Token
	DeleteToken(context.Context, *Transaction) (*TransactionResponse, error)
	// Wipes the provided amount of tokens from the specified Account ID
	WipeTokenAccount(context.Context, *Transaction) (*TransactionResponse, error)
	// Freezes the transfer of tokens to or from the specified Account ID
	FreezeTokenAccount(context.Context, *Transaction) (*TransactionResponse, error)
	// Unfreezes the transfer of tokens to or from the specified Account ID
	UnfreezeTokenAccount(context.Context, *Transaction) (*TransactionResponse, error)
	// Flags the provided Account ID as having gone through KYC
	GrantKycToTokenAccount(context.Context, *Transaction) (*TransactionResponse, error)
	// Removes the KYC flag of the provided Account ID
	RevokeKycFromTokenAccount(context.Context, *Transaction) (*TransactionResponse, error)
	// Associates tokens to an account
	AssociateTokens(context.Context, *Transaction) (*TransactionResponse, error)
	// Dissociates tokens from an account
	DissociateTokens(context.Context, *Transaction) (*TransactionResponse, error)
	// Retrieves the metadata of a token
	GetTokenInfo(context.Context, *Query) (*Response, error)
}

// UnimplementedTokenServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTokenServiceServer struct {
}

func (*UnimplementedTokenServiceServer) CreateToken(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateToken not implemented")
}
func (*UnimplementedTokenServiceServer) UpdateToken(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateToken not implemented")
}
func (*UnimplementedTokenServiceServer) MintToken(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintToken not implemented")
}
func (*UnimplementedTokenServiceServer) BurnToken(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BurnToken not implemented")
}
func (*UnimplementedTokenServiceServer) DeleteToken(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteToken not implemented")
}
func (*UnimplementedTokenServiceServer) WipeTokenAccount(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WipeTokenAccount not implemented")
}
func (*UnimplementedTokenServiceServer) FreezeTokenAccount(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FreezeTokenAccount not implemented")
}
func (*UnimplementedTokenServiceServer) UnfreezeTokenAccount(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnfreezeTokenAccount not implemented")
}
func (*UnimplementedTokenServiceServer) GrantKycToTokenAccount(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GrantKycToTokenAccount not implemented")
}
func (*UnimplementedTokenServiceServer) RevokeKycFromTokenAccount(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeKycFromTokenAccount not implemented")
}
func (*UnimplementedTokenServiceServer) AssociateTokens(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssociateTokens not implemented")
}
func (*UnimplementedTokenServiceServer) DissociateTokens(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DissociateTokens not implemented")
}
func (*UnimplementedTokenServiceServer) GetTokenInfo(context.Context, *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenInfo not implemented")
}

func RegisterTokenServiceServer(s *grpc.Server, srv TokenServiceServer) {
	s.RegisterService(&_TokenService_serviceDesc, srv)
}

func _TokenService_CreateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).CreateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TokenService/CreateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).CreateToken(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_UpdateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).UpdateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TokenService/UpdateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).UpdateToken(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_MintToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).MintToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TokenService/MintToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).MintToken(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_BurnToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).BurnToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TokenService/BurnToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).BurnToken(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_DeleteToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).DeleteToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TokenService/DeleteToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).DeleteToken(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_WipeTokenAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).WipeTokenAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TokenService/WipeTokenAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).WipeTokenAccount(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_FreezeTokenAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).FreezeTokenAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TokenService/FreezeTokenAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).FreezeTokenAccount(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_UnfreezeTokenAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).UnfreezeTokenAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TokenService/UnfreezeTokenAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).UnfreezeTokenAccount(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_GrantKycToTokenAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).GrantKycToTokenAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TokenService/GrantKycToTokenAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).GrantKycToTokenAccount(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_RevokeKycFromTokenAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).RevokeKycFromTokenAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TokenService/RevokeKycFromTokenAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).RevokeKycFromTokenAccount(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_AssociateTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).AssociateTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TokenService/AssociateTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).AssociateTokens(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_DissociateTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).DissociateTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TokenService/DissociateTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).DissociateTokens(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_GetTokenInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).GetTokenInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TokenService/GetTokenInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).GetTokenInfo(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

var _TokenService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TokenService",
	HandlerType: (*TokenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createToken",
			Handler:    _TokenService_CreateToken_Handler,
		},
		{
			MethodName: "updateToken",
			Handler:    _TokenService_UpdateToken_Handler,
		},
		{
			MethodName: "mintToken",
			Handler:    _TokenService_MintToken_Handler,
		},
		{
			MethodName: "burnToken",
			Handler:    _TokenService_BurnToken_Handler,
		},
		{
			MethodName: "deleteToken",
			Handler:    _TokenService_DeleteToken_Handler,
		},
		{
			MethodName: "wipeTokenAccount",
			Handler:    _TokenService_WipeTokenAccount_Handler,
		},
		{
			MethodName: "freezeTokenAccount",
			Handler:    _TokenService_FreezeTokenAccount_Handler,
		},
		{
			MethodName: "unfreezeTokenAccount",
			Handler:    _TokenService_UnfreezeTokenAccount_Handler,
		},
		{
			MethodName: "grantKycToTokenAccount",
			Handler:    _TokenService_GrantKycToTokenAccount_Handler,
		},
		{
			MethodName: "revokeKycFromTokenAccount",
			Handler:    _TokenService_RevokeKycFromTokenAccount_Handler,
		},
		{
			MethodName: "associateTokens",
			Handler:    _TokenService_AssociateTokens_Handler,
		},
		{
			MethodName: "dissociateTokens",
			Handler:    _TokenService_DissociateTokens_Handler,
		},
		{
			MethodName: "getTokenInfo",
			Handler:    _TokenService_GetTokenInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/TokenService.proto",
}

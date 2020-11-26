// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: proto/Transaction.proto

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

// A single signed transaction, including all its signatures. The SignatureList will have a Signature for each Key in the transaction, either explicit or implicit, in the order that they appear in the transaction. For example, a CryptoTransfer will first have a Signature corresponding to the Key for the paying account, followed by a Signature corresponding to the Key for each account that is sending or receiving cryptocurrency in the transfer. Each Transaction should not have more than 50 levels.
// The SignatureList field is deprecated and succeeded by SignatureMap.
type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SignedTransactionBytes []byte `protobuf:"bytes,5,opt,name=signedTransactionBytes,proto3" json:"signedTransactionBytes,omitempty"` // SignedTransaction serialized into bytes
	// Deprecated: Do not use.
	BodyBytes []byte `protobuf:"bytes,4,opt,name=bodyBytes,proto3" json:"bodyBytes,omitempty"` // TransactionBody serialized into bytes, which needs to be signed
	// Deprecated: Do not use.
	SigMap *SignatureMap `protobuf:"bytes,3,opt,name=sigMap,proto3" json:"sigMap,omitempty"` // The signatures on the body with the new format, to authorize the transaction
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_Transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_proto_Transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_proto_Transaction_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetSignedTransactionBytes() []byte {
	if x != nil {
		return x.SignedTransactionBytes
	}
	return nil
}

// Deprecated: Do not use.
func (x *Transaction) GetBodyBytes() []byte {
	if x != nil {
		return x.BodyBytes
	}
	return nil
}

// Deprecated: Do not use.
func (x *Transaction) GetSigMap() *SignatureMap {
	if x != nil {
		return x.SigMap
	}
	return nil
}

var File_proto_Transaction_proto protoreflect.FileDescriptor

var file_proto_Transaction_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x42, 0x61,
	0x73, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x6f, 0x64, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x01, 0x0a, 0x0b,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x16, 0x73,
	0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x16, 0x73, 0x69, 0x67,
	0x6e, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x09, 0x62, 0x6f, 0x64, 0x79, 0x42, 0x79, 0x74, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x02, 0x18, 0x01, 0x52, 0x09, 0x62, 0x6f, 0x64, 0x79,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x4d, 0x61, 0x70, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4d, 0x61, 0x70, 0x42, 0x02, 0x18, 0x01, 0x52, 0x06,
	0x73, 0x69, 0x67, 0x4d, 0x61, 0x70, 0x42, 0x4b, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x68, 0x65, 0x64,
	0x65, 0x72, 0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_Transaction_proto_rawDescOnce sync.Once
	file_proto_Transaction_proto_rawDescData = file_proto_Transaction_proto_rawDesc
)

func file_proto_Transaction_proto_rawDescGZIP() []byte {
	file_proto_Transaction_proto_rawDescOnce.Do(func() {
		file_proto_Transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_Transaction_proto_rawDescData)
	})
	return file_proto_Transaction_proto_rawDescData
}

var file_proto_Transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_Transaction_proto_goTypes = []interface{}{
	(*Transaction)(nil),  // 0: proto.Transaction
	(*SignatureMap)(nil), // 1: proto.SignatureMap
}
var file_proto_Transaction_proto_depIdxs = []int32{
	1, // 0: proto.Transaction.sigMap:type_name -> proto.SignatureMap
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_Transaction_proto_init() }
func file_proto_Transaction_proto_init() {
	if File_proto_Transaction_proto != nil {
		return
	}
	file_proto_Duration_proto_init()
	file_proto_BasicTypes_proto_init()
	file_proto_TransactionBody_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_Transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
			RawDescriptor: file_proto_Transaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_Transaction_proto_goTypes,
		DependencyIndexes: file_proto_Transaction_proto_depIdxs,
		MessageInfos:      file_proto_Transaction_proto_msgTypes,
	}.Build()
	File_proto_Transaction_proto = out.File
	file_proto_Transaction_proto_rawDesc = nil
	file_proto_Transaction_proto_goTypes = nil
	file_proto_Transaction_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: proto/FileGetInfo.proto

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

// Get all of the information about a file, except for its contents. When a file expires, it no longer exists, and there will be no info about it, and the fileInfo field will be blank. If a transaction or smart contract deletes the file, but it has not yet expired, then the fileInfo field will be non-empty, the deleted field will be true, its size will be 0, and its contents will be empty.
type FileGetInfoQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *QueryHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"` // Standard info sent from client to node, including the signed payment, and what kind of response is requested (cost, state proof, both, or neither).
	FileID *FileID      `protobuf:"bytes,2,opt,name=fileID,proto3" json:"fileID,omitempty"` // The file ID of the file for which information is requested
}

func (x *FileGetInfoQuery) Reset() {
	*x = FileGetInfoQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_FileGetInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileGetInfoQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileGetInfoQuery) ProtoMessage() {}

func (x *FileGetInfoQuery) ProtoReflect() protoreflect.Message {
	mi := &file_proto_FileGetInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileGetInfoQuery.ProtoReflect.Descriptor instead.
func (*FileGetInfoQuery) Descriptor() ([]byte, []int) {
	return file_proto_FileGetInfo_proto_rawDescGZIP(), []int{0}
}

func (x *FileGetInfoQuery) GetHeader() *QueryHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *FileGetInfoQuery) GetFileID() *FileID {
	if x != nil {
		return x.FileID
	}
	return nil
}

// Response when the client sends the node FileGetInfoQuery
type FileGetInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header   *ResponseHeader               `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`     //Standard response from node to client, including the requested fields: cost, or state proof, or both, or neither
	FileInfo *FileGetInfoResponse_FileInfo `protobuf:"bytes,2,opt,name=fileInfo,proto3" json:"fileInfo,omitempty"` // The information about the file
}

func (x *FileGetInfoResponse) Reset() {
	*x = FileGetInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_FileGetInfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileGetInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileGetInfoResponse) ProtoMessage() {}

func (x *FileGetInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_FileGetInfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileGetInfoResponse.ProtoReflect.Descriptor instead.
func (*FileGetInfoResponse) Descriptor() ([]byte, []int) {
	return file_proto_FileGetInfo_proto_rawDescGZIP(), []int{1}
}

func (x *FileGetInfoResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *FileGetInfoResponse) GetFileInfo() *FileGetInfoResponse_FileInfo {
	if x != nil {
		return x.FileInfo
	}
	return nil
}

type FileGetInfoResponse_FileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileID         *FileID    `protobuf:"bytes,1,opt,name=fileID,proto3" json:"fileID,omitempty"`                 // The file ID of the file for which information is requested
	Size           int64      `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`                    // Number of bytes in contents
	ExpirationTime *Timestamp `protobuf:"bytes,3,opt,name=expirationTime,proto3" json:"expirationTime,omitempty"` // The current time at which this account is set to expire
	Deleted        bool       `protobuf:"varint,4,opt,name=deleted,proto3" json:"deleted,omitempty"`              // True if deleted but not yet expired
	Keys           *KeyList   `protobuf:"bytes,5,opt,name=keys,proto3" json:"keys,omitempty"`                     // One of these keys must sign in order to modify or delete the file
}

func (x *FileGetInfoResponse_FileInfo) Reset() {
	*x = FileGetInfoResponse_FileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_FileGetInfo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileGetInfoResponse_FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileGetInfoResponse_FileInfo) ProtoMessage() {}

func (x *FileGetInfoResponse_FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_FileGetInfo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileGetInfoResponse_FileInfo.ProtoReflect.Descriptor instead.
func (*FileGetInfoResponse_FileInfo) Descriptor() ([]byte, []int) {
	return file_proto_FileGetInfo_proto_rawDescGZIP(), []int{1, 0}
}

func (x *FileGetInfoResponse_FileInfo) GetFileID() *FileID {
	if x != nil {
		return x.FileID
	}
	return nil
}

func (x *FileGetInfoResponse_FileInfo) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *FileGetInfoResponse_FileInfo) GetExpirationTime() *Timestamp {
	if x != nil {
		return x.ExpirationTime
	}
	return nil
}

func (x *FileGetInfoResponse_FileInfo) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

func (x *FileGetInfoResponse_FileInfo) GetKeys() *KeyList {
	if x != nil {
		return x.Keys
	}
	return nil
}

var File_proto_FileGetInfo_proto protoreflect.FileDescriptor

var file_proto_FileGetInfo_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x46, 0x69, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x42,
	0x61, 0x73, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a, 0x10, 0x46, 0x69, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x49, 0x44, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x22, 0xc5, 0x02, 0x0a, 0x13,
	0x46, 0x69, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x3f, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x1a, 0xbd, 0x01, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x25, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x52,
	0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x38, 0x0a, 0x0e, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12,
	0x22, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x04, 0x6b,
	0x65, 0x79, 0x73, 0x42, 0x4b, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61,
	0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_FileGetInfo_proto_rawDescOnce sync.Once
	file_proto_FileGetInfo_proto_rawDescData = file_proto_FileGetInfo_proto_rawDesc
)

func file_proto_FileGetInfo_proto_rawDescGZIP() []byte {
	file_proto_FileGetInfo_proto_rawDescOnce.Do(func() {
		file_proto_FileGetInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_FileGetInfo_proto_rawDescData)
	})
	return file_proto_FileGetInfo_proto_rawDescData
}

var file_proto_FileGetInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_FileGetInfo_proto_goTypes = []interface{}{
	(*FileGetInfoQuery)(nil),             // 0: proto.FileGetInfoQuery
	(*FileGetInfoResponse)(nil),          // 1: proto.FileGetInfoResponse
	(*FileGetInfoResponse_FileInfo)(nil), // 2: proto.FileGetInfoResponse.FileInfo
	(*QueryHeader)(nil),                  // 3: proto.QueryHeader
	(*FileID)(nil),                       // 4: proto.FileID
	(*ResponseHeader)(nil),               // 5: proto.ResponseHeader
	(*Timestamp)(nil),                    // 6: proto.Timestamp
	(*KeyList)(nil),                      // 7: proto.KeyList
}
var file_proto_FileGetInfo_proto_depIdxs = []int32{
	3, // 0: proto.FileGetInfoQuery.header:type_name -> proto.QueryHeader
	4, // 1: proto.FileGetInfoQuery.fileID:type_name -> proto.FileID
	5, // 2: proto.FileGetInfoResponse.header:type_name -> proto.ResponseHeader
	2, // 3: proto.FileGetInfoResponse.fileInfo:type_name -> proto.FileGetInfoResponse.FileInfo
	4, // 4: proto.FileGetInfoResponse.FileInfo.fileID:type_name -> proto.FileID
	6, // 5: proto.FileGetInfoResponse.FileInfo.expirationTime:type_name -> proto.Timestamp
	7, // 6: proto.FileGetInfoResponse.FileInfo.keys:type_name -> proto.KeyList
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_proto_FileGetInfo_proto_init() }
func file_proto_FileGetInfo_proto_init() {
	if File_proto_FileGetInfo_proto != nil {
		return
	}
	file_proto_Timestamp_proto_init()
	file_proto_BasicTypes_proto_init()
	file_proto_QueryHeader_proto_init()
	file_proto_ResponseHeader_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_FileGetInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileGetInfoQuery); i {
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
		file_proto_FileGetInfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileGetInfoResponse); i {
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
		file_proto_FileGetInfo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileGetInfoResponse_FileInfo); i {
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
			RawDescriptor: file_proto_FileGetInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_FileGetInfo_proto_goTypes,
		DependencyIndexes: file_proto_FileGetInfo_proto_depIdxs,
		MessageInfos:      file_proto_FileGetInfo_proto_msgTypes,
	}.Build()
	File_proto_FileGetInfo_proto = out.File
	file_proto_FileGetInfo_proto_rawDesc = nil
	file_proto_FileGetInfo_proto_goTypes = nil
	file_proto_FileGetInfo_proto_depIdxs = nil
}

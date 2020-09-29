// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: proto/Freeze.proto

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

// Set the freezing period in which the platform will stop creating events and accepting transactions. This is used before safely shut down the platform for maintenance.
type FreezeTransactionBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartHour  int32   `protobuf:"varint,1,opt,name=startHour,proto3" json:"startHour,omitempty"`  // The start hour (in UTC time), a value between 0 and 23
	StartMin   int32   `protobuf:"varint,2,opt,name=startMin,proto3" json:"startMin,omitempty"`    // The start minute (in UTC time), a value between 0 and 59
	EndHour    int32   `protobuf:"varint,3,opt,name=endHour,proto3" json:"endHour,omitempty"`      // The end hour (in UTC time), a value between 0 and 23
	EndMin     int32   `protobuf:"varint,4,opt,name=endMin,proto3" json:"endMin,omitempty"`        // The end minute (in UTC time), a value between 0 and 59
	UpdateFile *FileID `protobuf:"bytes,5,opt,name=updateFile,proto3" json:"updateFile,omitempty"` // The ID of the file needs to be updated during a freeze transaction
}

func (x *FreezeTransactionBody) Reset() {
	*x = FreezeTransactionBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_Freeze_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FreezeTransactionBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FreezeTransactionBody) ProtoMessage() {}

func (x *FreezeTransactionBody) ProtoReflect() protoreflect.Message {
	mi := &file_proto_Freeze_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FreezeTransactionBody.ProtoReflect.Descriptor instead.
func (*FreezeTransactionBody) Descriptor() ([]byte, []int) {
	return file_proto_Freeze_proto_rawDescGZIP(), []int{0}
}

func (x *FreezeTransactionBody) GetStartHour() int32 {
	if x != nil {
		return x.StartHour
	}
	return 0
}

func (x *FreezeTransactionBody) GetStartMin() int32 {
	if x != nil {
		return x.StartMin
	}
	return 0
}

func (x *FreezeTransactionBody) GetEndHour() int32 {
	if x != nil {
		return x.EndHour
	}
	return 0
}

func (x *FreezeTransactionBody) GetEndMin() int32 {
	if x != nil {
		return x.EndMin
	}
	return 0
}

func (x *FreezeTransactionBody) GetUpdateFile() *FileID {
	if x != nil {
		return x.UpdateFile
	}
	return nil
}

var File_proto_Freeze_proto protoreflect.FileDescriptor

var file_proto_Freeze_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x46, 0x72, 0x65, 0x65, 0x7a, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x42, 0x61, 0x73, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x01, 0x0a, 0x15, 0x46, 0x72, 0x65, 0x65, 0x7a, 0x65, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x48, 0x6f, 0x75, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x48, 0x6f, 0x75, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x4d, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x4d, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x48, 0x6f,
	0x75, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x48, 0x6f, 0x75,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x64, 0x4d, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x65, 0x6e, 0x64, 0x4d, 0x69, 0x6e, 0x12, 0x2d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x52, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x42, 0x50, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e,
	0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01,
	0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73,
	0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x2d, 0x73, 0x64,
	0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_Freeze_proto_rawDescOnce sync.Once
	file_proto_Freeze_proto_rawDescData = file_proto_Freeze_proto_rawDesc
)

func file_proto_Freeze_proto_rawDescGZIP() []byte {
	file_proto_Freeze_proto_rawDescOnce.Do(func() {
		file_proto_Freeze_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_Freeze_proto_rawDescData)
	})
	return file_proto_Freeze_proto_rawDescData
}

var file_proto_Freeze_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_Freeze_proto_goTypes = []interface{}{
	(*FreezeTransactionBody)(nil), // 0: proto.FreezeTransactionBody
	(*FileID)(nil),                // 1: proto.FileID
}
var file_proto_Freeze_proto_depIdxs = []int32{
	1, // 0: proto.FreezeTransactionBody.updateFile:type_name -> proto.FileID
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_Freeze_proto_init() }
func file_proto_Freeze_proto_init() {
	if File_proto_Freeze_proto != nil {
		return
	}
	file_proto_BasicTypes_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_Freeze_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FreezeTransactionBody); i {
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
			RawDescriptor: file_proto_Freeze_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_Freeze_proto_goTypes,
		DependencyIndexes: file_proto_Freeze_proto_depIdxs,
		MessageInfos:      file_proto_Freeze_proto_msgTypes,
	}.Build()
	File_proto_Freeze_proto = out.File
	file_proto_Freeze_proto_rawDesc = nil
	file_proto_Freeze_proto_goTypes = nil
	file_proto_Freeze_proto_depIdxs = nil
}

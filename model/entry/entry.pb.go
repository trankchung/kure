// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: entry.proto

package entry

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

type Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	URL      string `protobuf:"bytes,4,opt,name=URL,proto3" json:"URL,omitempty"`
	Notes    string `protobuf:"bytes,5,opt,name=Notes,proto3" json:"Notes,omitempty"`
	Expires  string `protobuf:"bytes,6,opt,name=Expires,proto3" json:"Expires,omitempty"`
}

func (x *Entry) Reset() {
	*x = Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entry) ProtoMessage() {}

func (x *Entry) ProtoReflect() protoreflect.Message {
	mi := &file_entry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entry.ProtoReflect.Descriptor instead.
func (*Entry) Descriptor() ([]byte, []int) {
	return file_entry_proto_rawDescGZIP(), []int{0}
}

func (x *Entry) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Entry) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Entry) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Entry) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

func (x *Entry) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

func (x *Entry) GetExpires() string {
	if x != nil {
		return x.Expires
	}
	return ""
}

var File_entry_proto protoreflect.FileDescriptor

var file_entry_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01,
	0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x4c, 0x12, 0x14, 0x0a, 0x05, 0x4e, 0x6f, 0x74, 0x65, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entry_proto_rawDescOnce sync.Once
	file_entry_proto_rawDescData = file_entry_proto_rawDesc
)

func file_entry_proto_rawDescGZIP() []byte {
	file_entry_proto_rawDescOnce.Do(func() {
		file_entry_proto_rawDescData = protoimpl.X.CompressGZIP(file_entry_proto_rawDescData)
	})
	return file_entry_proto_rawDescData
}

var file_entry_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_entry_proto_goTypes = []interface{}{
	(*Entry)(nil), // 0: Entry
}
var file_entry_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_entry_proto_init() }
func file_entry_proto_init() {
	if File_entry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_entry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entry); i {
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
			RawDescriptor: file_entry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entry_proto_goTypes,
		DependencyIndexes: file_entry_proto_depIdxs,
		MessageInfos:      file_entry_proto_msgTypes,
	}.Build()
	File_entry_proto = out.File
	file_entry_proto_rawDesc = nil
	file_entry_proto_goTypes = nil
	file_entry_proto_depIdxs = nil
}
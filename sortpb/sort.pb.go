// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: sort/sort.proto

package sortpb

import (
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

type Sort int32

const (
	Sort_DESC Sort = 0
	Sort_ASC  Sort = 1
)

// Enum value maps for Sort.
var (
	Sort_name = map[int32]string{
		0: "DESC",
		1: "ASC",
	}
	Sort_value = map[string]int32{
		"DESC": 0,
		"ASC":  1,
	}
)

func (x Sort) Enum() *Sort {
	p := new(Sort)
	*p = x
	return p
}

func (x Sort) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Sort) Descriptor() protoreflect.EnumDescriptor {
	return file_sort_sort_proto_enumTypes[0].Descriptor()
}

func (Sort) Type() protoreflect.EnumType {
	return &file_sort_sort_proto_enumTypes[0]
}

func (x Sort) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Sort.Descriptor instead.
func (Sort) EnumDescriptor() ([]byte, []int) {
	return file_sort_sort_proto_rawDescGZIP(), []int{0}
}

var File_sort_sort_proto protoreflect.FileDescriptor

var file_sort_sort_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x6f, 0x72, 0x74, 0x2f, 0x73, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x2a, 0x19, 0x0a, 0x04, 0x53, 0x6f, 0x72, 0x74, 0x12,
	0x08, 0x0a, 0x04, 0x44, 0x45, 0x53, 0x43, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x53, 0x43,
	0x10, 0x01, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x62, 0x75, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x6f, 0x72, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_sort_sort_proto_rawDescOnce sync.Once
	file_sort_sort_proto_rawDescData = file_sort_sort_proto_rawDesc
)

func file_sort_sort_proto_rawDescGZIP() []byte {
	file_sort_sort_proto_rawDescOnce.Do(func() {
		file_sort_sort_proto_rawDescData = protoimpl.X.CompressGZIP(file_sort_sort_proto_rawDescData)
	})
	return file_sort_sort_proto_rawDescData
}

var file_sort_sort_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sort_sort_proto_goTypes = []interface{}{
	(Sort)(0), // 0: sort.Sort
}
var file_sort_sort_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sort_sort_proto_init() }
func file_sort_sort_proto_init() {
	if File_sort_sort_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sort_sort_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sort_sort_proto_goTypes,
		DependencyIndexes: file_sort_sort_proto_depIdxs,
		EnumInfos:         file_sort_sort_proto_enumTypes,
	}.Build()
	File_sort_sort_proto = out.File
	file_sort_sort_proto_rawDesc = nil
	file_sort_sort_proto_goTypes = nil
	file_sort_sort_proto_depIdxs = nil
}

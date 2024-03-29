// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: protoc/book.proto

package protoc

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

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     string  `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title  string  `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Author string  `protobuf:"bytes,3,opt,name=Author,proto3" json:"Author,omitempty"`
	Year   uint32  `protobuf:"varint,4,opt,name=Year,proto3" json:"Year,omitempty"`
	Size   uint32  `protobuf:"varint,5,opt,name=Size,proto3" json:"Size,omitempty"`
	Rate   float32 `protobuf:"fixed32,6,opt,name=Rate,proto3" json:"Rate,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_book_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_book_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_protoc_book_proto_rawDescGZIP(), []int{0}
}

func (x *Book) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Book) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Book) GetYear() uint32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Book) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Book) GetRate() float32 {
	if x != nil {
		return x.Rate
	}
	return 0
}

type BooksSlice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books []*Book `protobuf:"bytes,1,rep,name=Books,proto3" json:"Books,omitempty"`
}

func (x *BooksSlice) Reset() {
	*x = BooksSlice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_book_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BooksSlice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BooksSlice) ProtoMessage() {}

func (x *BooksSlice) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_book_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BooksSlice.ProtoReflect.Descriptor instead.
func (*BooksSlice) Descriptor() ([]byte, []int) {
	return file_protoc_book_proto_rawDescGZIP(), []int{1}
}

func (x *BooksSlice) GetBooks() []*Book {
	if x != nil {
		return x.Books
	}
	return nil
}

var File_protoc_book_proto protoreflect.FileDescriptor

var file_protoc_book_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x22, 0x80, 0x01, 0x0a, 0x04,
	0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x59, 0x65, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x59, 0x65, 0x61, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x61,
	0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x52, 0x61, 0x74, 0x65, 0x22, 0x30,
	0x0a, 0x0a, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x05,
	0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x05, 0x42, 0x6f, 0x6f, 0x6b, 0x73,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protoc_book_proto_rawDescOnce sync.Once
	file_protoc_book_proto_rawDescData = file_protoc_book_proto_rawDesc
)

func file_protoc_book_proto_rawDescGZIP() []byte {
	file_protoc_book_proto_rawDescOnce.Do(func() {
		file_protoc_book_proto_rawDescData = protoimpl.X.CompressGZIP(file_protoc_book_proto_rawDescData)
	})
	return file_protoc_book_proto_rawDescData
}

var file_protoc_book_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protoc_book_proto_goTypes = []interface{}{
	(*Book)(nil),       // 0: protoc.Book
	(*BooksSlice)(nil), // 1: protoc.BooksSlice
}
var file_protoc_book_proto_depIdxs = []int32{
	0, // 0: protoc.BooksSlice.Books:type_name -> protoc.Book
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protoc_book_proto_init() }
func file_protoc_book_proto_init() {
	if File_protoc_book_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protoc_book_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
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
		file_protoc_book_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BooksSlice); i {
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
			RawDescriptor: file_protoc_book_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protoc_book_proto_goTypes,
		DependencyIndexes: file_protoc_book_proto_depIdxs,
		MessageInfos:      file_protoc_book_proto_msgTypes,
	}.Build()
	File_protoc_book_proto = out.File
	file_protoc_book_proto_rawDesc = nil
	file_protoc_book_proto_goTypes = nil
	file_protoc_book_proto_depIdxs = nil
}

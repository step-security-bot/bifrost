// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.19.3
// source: github.com/aperturerobotics/bifrost/transport/common/pconn/pconn.proto

package pconn

import (
	reflect "reflect"
	sync "sync"

	quic "github.com/aperturerobotics/bifrost/transport/common/quic"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Opts are extra options for the packet conn.
type Opts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Quic are the quic protocol options.
	Quic *quic.Opts `protobuf:"bytes,1,opt,name=quic,proto3" json:"quic,omitempty"`
	// Verbose turns on verbose debug logging.
	Verbose bool `protobuf:"varint,2,opt,name=verbose,proto3" json:"verbose,omitempty"`
}

func (x *Opts) Reset() {
	*x = Opts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_aperturerobotics_bifrost_transport_common_pconn_pconn_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Opts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Opts) ProtoMessage() {}

func (x *Opts) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_aperturerobotics_bifrost_transport_common_pconn_pconn_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Opts.ProtoReflect.Descriptor instead.
func (*Opts) Descriptor() ([]byte, []int) {
	return file_github_com_aperturerobotics_bifrost_transport_common_pconn_pconn_proto_rawDescGZIP(), []int{0}
}

func (x *Opts) GetQuic() *quic.Opts {
	if x != nil {
		return x.Quic
	}
	return nil
}

func (x *Opts) GetVerbose() bool {
	if x != nil {
		return x.Verbose
	}
	return false
}

var File_github_com_aperturerobotics_bifrost_transport_common_pconn_pconn_proto protoreflect.FileDescriptor

var file_github_com_aperturerobotics_bifrost_transport_common_pconn_pconn_proto_rawDesc = []byte{
	0x0a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x65,
	0x72, 0x74, 0x75, 0x72, 0x65, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x62, 0x69,
	0x66, 0x72, 0x6f, 0x73, 0x74, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x63, 0x6f, 0x6e, 0x6e, 0x2f, 0x70, 0x63, 0x6f,
	0x6e, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x63, 0x6f, 0x6e, 0x6e, 0x1a,
	0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x65, 0x72,
	0x74, 0x75, 0x72, 0x65, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x62, 0x69, 0x66,
	0x72, 0x6f, 0x73, 0x74, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x71, 0x75, 0x69, 0x63, 0x2f, 0x71, 0x75, 0x69, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x04, 0x4f, 0x70, 0x74, 0x73, 0x12, 0x28, 0x0a,
	0x04, 0x71, 0x75, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x71, 0x75, 0x69, 0x63, 0x2e, 0x4f, 0x70, 0x74,
	0x73, 0x52, 0x04, 0x71, 0x75, 0x69, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x62, 0x6f,
	0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x76, 0x65, 0x72, 0x62, 0x6f, 0x73,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_aperturerobotics_bifrost_transport_common_pconn_pconn_proto_rawDescOnce sync.Once
	file_github_com_aperturerobotics_bifrost_transport_common_pconn_pconn_proto_rawDescData = file_github_com_aperturerobotics_bifrost_transport_common_pconn_pconn_proto_rawDesc
)

func file_github_com_aperturerobotics_bifrost_transport_common_pconn_pconn_proto_rawDescGZIP() []byte {
	file_github_com_aperturerobotics_bifrost_transport_common_pconn_pconn_proto_rawDescOnce.Do(func() {
		file_github_com_aperturerobotics_bifrost_transport_common_pconn_pconn_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_aperturerobotics_bifrost_transport_common_pconn_pconn_proto_rawDescData)
	})
	return file_github_com_aperturerobotics_bifrost_transport_common_pconn_pconn_proto_rawDescData
}

var file_github_com_aperturerobotics_bifrost_transport_common_pconn_pconn_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_aperturerobotics_bifrost_transport_common_pconn_pconn_proto_goTypes = []interface{}{
	(*Opts)(nil),      // 0: pconn.Opts
	(*quic.Opts)(nil), // 1: transport.quic.Opts
}
var file_github_com_aperturerobotics_bifrost_transport_common_pconn_pconn_proto_depIdxs = []int32{
	1, // 0: pconn.Opts.quic:type_name -> transport.quic.Opts
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_github_com_aperturerobotics_bifrost_transport_common_pconn_pconn_proto_init() }
func file_github_com_aperturerobotics_bifrost_transport_common_pconn_pconn_proto_init() {
	if File_github_com_aperturerobotics_bifrost_transport_common_pconn_pconn_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_aperturerobotics_bifrost_transport_common_pconn_pconn_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Opts); i {
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
			RawDescriptor: file_github_com_aperturerobotics_bifrost_transport_common_pconn_pconn_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_aperturerobotics_bifrost_transport_common_pconn_pconn_proto_goTypes,
		DependencyIndexes: file_github_com_aperturerobotics_bifrost_transport_common_pconn_pconn_proto_depIdxs,
		MessageInfos:      file_github_com_aperturerobotics_bifrost_transport_common_pconn_pconn_proto_msgTypes,
	}.Build()
	File_github_com_aperturerobotics_bifrost_transport_common_pconn_pconn_proto = out.File
	file_github_com_aperturerobotics_bifrost_transport_common_pconn_pconn_proto_rawDesc = nil
	file_github_com_aperturerobotics_bifrost_transport_common_pconn_pconn_proto_goTypes = nil
	file_github_com_aperturerobotics_bifrost_transport_common_pconn_pconn_proto_depIdxs = nil
}

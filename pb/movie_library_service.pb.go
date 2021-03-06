// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: movie_library_service.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type AddMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Movie *Movie `protobuf:"bytes,1,opt,name=movie,proto3" json:"movie,omitempty"`
}

func (x *AddMovieRequest) Reset() {
	*x = AddMovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_library_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMovieRequest) ProtoMessage() {}

func (x *AddMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_movie_library_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMovieRequest.ProtoReflect.Descriptor instead.
func (*AddMovieRequest) Descriptor() ([]byte, []int) {
	return file_movie_library_service_proto_rawDescGZIP(), []int{0}
}

func (x *AddMovieRequest) GetMovie() *Movie {
	if x != nil {
		return x.Movie
	}
	return nil
}

type AddMovieResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddMovieResponse) Reset() {
	*x = AddMovieResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_library_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMovieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMovieResponse) ProtoMessage() {}

func (x *AddMovieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_movie_library_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMovieResponse.ProtoReflect.Descriptor instead.
func (*AddMovieResponse) Descriptor() ([]byte, []int) {
	return file_movie_library_service_proto_rawDescGZIP(), []int{1}
}

func (x *AddMovieResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_movie_library_service_proto protoreflect.FileDescriptor

var file_movie_library_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67,
	0x6f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x6c, 0x69,
	0x62, 0x72, 0x61, 0x72, 0x79, 0x1a, 0x13, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x0f, 0x41, 0x64,
	0x64, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a,
	0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67,
	0x6f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x6c, 0x69,
	0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x05, 0x6d, 0x6f, 0x76,
	0x69, 0x65, 0x22, 0x22, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x75, 0x0a, 0x13, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x4c,
	0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a,
	0x09, 0x53, 0x61, 0x76, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x27, 0x2e, 0x67, 0x6f, 0x73,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72,
	0x61, 0x72, 0x79, 0x2e, 0x41, 0x64, 0x64, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x67, 0x6f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x41, 0x64, 0x64,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x20, 0x0a,
	0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x50, 0x01, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_movie_library_service_proto_rawDescOnce sync.Once
	file_movie_library_service_proto_rawDescData = file_movie_library_service_proto_rawDesc
)

func file_movie_library_service_proto_rawDescGZIP() []byte {
	file_movie_library_service_proto_rawDescOnce.Do(func() {
		file_movie_library_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_movie_library_service_proto_rawDescData)
	})
	return file_movie_library_service_proto_rawDescData
}

var file_movie_library_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_movie_library_service_proto_goTypes = []interface{}{
	(*AddMovieRequest)(nil),  // 0: gosample.movie.library.AddMovieRequest
	(*AddMovieResponse)(nil), // 1: gosample.movie.library.AddMovieResponse
	(*Movie)(nil),            // 2: gosample.movie.library.Movie
}
var file_movie_library_service_proto_depIdxs = []int32{
	2, // 0: gosample.movie.library.AddMovieRequest.movie:type_name -> gosample.movie.library.Movie
	0, // 1: gosample.movie.library.MovieLibraryService.SaveMovie:input_type -> gosample.movie.library.AddMovieRequest
	1, // 2: gosample.movie.library.MovieLibraryService.SaveMovie:output_type -> gosample.movie.library.AddMovieResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_movie_library_service_proto_init() }
func file_movie_library_service_proto_init() {
	if File_movie_library_service_proto != nil {
		return
	}
	file_movie_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_movie_library_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddMovieRequest); i {
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
		file_movie_library_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddMovieResponse); i {
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
			RawDescriptor: file_movie_library_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_movie_library_service_proto_goTypes,
		DependencyIndexes: file_movie_library_service_proto_depIdxs,
		MessageInfos:      file_movie_library_service_proto_msgTypes,
	}.Build()
	File_movie_library_service_proto = out.File
	file_movie_library_service_proto_rawDesc = nil
	file_movie_library_service_proto_goTypes = nil
	file_movie_library_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MovieLibraryServiceClient is the client API for MovieLibraryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MovieLibraryServiceClient interface {
	SaveMovie(ctx context.Context, in *AddMovieRequest, opts ...grpc.CallOption) (*AddMovieResponse, error)
}

type movieLibraryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMovieLibraryServiceClient(cc grpc.ClientConnInterface) MovieLibraryServiceClient {
	return &movieLibraryServiceClient{cc}
}

func (c *movieLibraryServiceClient) SaveMovie(ctx context.Context, in *AddMovieRequest, opts ...grpc.CallOption) (*AddMovieResponse, error) {
	out := new(AddMovieResponse)
	err := c.cc.Invoke(ctx, "/gosample.movie.library.MovieLibraryService/SaveMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MovieLibraryServiceServer is the server API for MovieLibraryService service.
type MovieLibraryServiceServer interface {
	SaveMovie(context.Context, *AddMovieRequest) (*AddMovieResponse, error)
}

// UnimplementedMovieLibraryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMovieLibraryServiceServer struct {
}

func (*UnimplementedMovieLibraryServiceServer) SaveMovie(context.Context, *AddMovieRequest) (*AddMovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveMovie not implemented")
}

func RegisterMovieLibraryServiceServer(s *grpc.Server, srv MovieLibraryServiceServer) {
	s.RegisterService(&_MovieLibraryService_serviceDesc, srv)
}

func _MovieLibraryService_SaveMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieLibraryServiceServer).SaveMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gosample.movie.library.MovieLibraryService/SaveMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieLibraryServiceServer).SaveMovie(ctx, req.(*AddMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MovieLibraryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gosample.movie.library.MovieLibraryService",
	HandlerType: (*MovieLibraryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveMovie",
			Handler:    _MovieLibraryService_SaveMovie_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "movie_library_service.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: thumbnails.proto

package proto

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

// The file types to which the thumbnail can get encoded to.
type GetThumbnailRequest_ThumbnailType int32

const (
	GetThumbnailRequest_PNG GetThumbnailRequest_ThumbnailType = 0 // Represents PNG type
	GetThumbnailRequest_JPG GetThumbnailRequest_ThumbnailType = 1 // Represents JPG type
)

// Enum value maps for GetThumbnailRequest_ThumbnailType.
var (
	GetThumbnailRequest_ThumbnailType_name = map[int32]string{
		0: "PNG",
		1: "JPG",
	}
	GetThumbnailRequest_ThumbnailType_value = map[string]int32{
		"PNG": 0,
		"JPG": 1,
	}
)

func (x GetThumbnailRequest_ThumbnailType) Enum() *GetThumbnailRequest_ThumbnailType {
	p := new(GetThumbnailRequest_ThumbnailType)
	*p = x
	return p
}

func (x GetThumbnailRequest_ThumbnailType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetThumbnailRequest_ThumbnailType) Descriptor() protoreflect.EnumDescriptor {
	return file_thumbnails_proto_enumTypes[0].Descriptor()
}

func (GetThumbnailRequest_ThumbnailType) Type() protoreflect.EnumType {
	return &file_thumbnails_proto_enumTypes[0]
}

func (x GetThumbnailRequest_ThumbnailType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetThumbnailRequest_ThumbnailType.Descriptor instead.
func (GetThumbnailRequest_ThumbnailType) EnumDescriptor() ([]byte, []int) {
	return file_thumbnails_proto_rawDescGZIP(), []int{0, 0}
}

// A request to retrieve a thumbnail
type GetThumbnailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The path to the source image
	Filepath string `protobuf:"bytes,1,opt,name=filepath,proto3" json:"filepath,omitempty"`
	// The type to which the thumbnail should get encoded to.
	ThumbnailType GetThumbnailRequest_ThumbnailType `protobuf:"varint,2,opt,name=thumbnail_type,json=thumbnailType,proto3,enum=com.owncloud.ocis.thumbnails.v0.GetThumbnailRequest_ThumbnailType" json:"thumbnail_type,omitempty"`
	// The width of the thumbnail
	Width int32 `protobuf:"varint,3,opt,name=width,proto3" json:"width,omitempty"`
	// The height of the thumbnail
	Height int32 `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
	// Types that are assignable to Source:
	//	*GetThumbnailRequest_WebdavSource
	//	*GetThumbnailRequest_Cs3Source
	Source isGetThumbnailRequest_Source `protobuf_oneof:"source"`
}

func (x *GetThumbnailRequest) Reset() {
	*x = GetThumbnailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thumbnails_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetThumbnailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetThumbnailRequest) ProtoMessage() {}

func (x *GetThumbnailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thumbnails_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetThumbnailRequest.ProtoReflect.Descriptor instead.
func (*GetThumbnailRequest) Descriptor() ([]byte, []int) {
	return file_thumbnails_proto_rawDescGZIP(), []int{0}
}

func (x *GetThumbnailRequest) GetFilepath() string {
	if x != nil {
		return x.Filepath
	}
	return ""
}

func (x *GetThumbnailRequest) GetThumbnailType() GetThumbnailRequest_ThumbnailType {
	if x != nil {
		return x.ThumbnailType
	}
	return GetThumbnailRequest_PNG
}

func (x *GetThumbnailRequest) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *GetThumbnailRequest) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (m *GetThumbnailRequest) GetSource() isGetThumbnailRequest_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (x *GetThumbnailRequest) GetWebdavSource() *WebdavSource {
	if x, ok := x.GetSource().(*GetThumbnailRequest_WebdavSource); ok {
		return x.WebdavSource
	}
	return nil
}

func (x *GetThumbnailRequest) GetCs3Source() *CS3Source {
	if x, ok := x.GetSource().(*GetThumbnailRequest_Cs3Source); ok {
		return x.Cs3Source
	}
	return nil
}

type isGetThumbnailRequest_Source interface {
	isGetThumbnailRequest_Source()
}

type GetThumbnailRequest_WebdavSource struct {
	WebdavSource *WebdavSource `protobuf:"bytes,5,opt,name=webdav_source,json=webdavSource,proto3,oneof"`
}

type GetThumbnailRequest_Cs3Source struct {
	Cs3Source *CS3Source `protobuf:"bytes,6,opt,name=cs3_source,json=cs3Source,proto3,oneof"`
}

func (*GetThumbnailRequest_WebdavSource) isGetThumbnailRequest_Source() {}

func (*GetThumbnailRequest_Cs3Source) isGetThumbnailRequest_Source() {}

type WebdavSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// REQUIRED.
	IsPublicLink bool `protobuf:"varint,2,opt,name=is_public_link,json=isPublicLink,proto3" json:"is_public_link,omitempty"`
	// OPTIONAL.
	WebdavAuthorization string `protobuf:"bytes,3,opt,name=webdav_authorization,json=webdavAuthorization,proto3" json:"webdav_authorization,omitempty"`
	// OPTIONAL.
	RevaAuthorization string `protobuf:"bytes,4,opt,name=reva_authorization,json=revaAuthorization,proto3" json:"reva_authorization,omitempty"`
	// OPTIONAL.
	PublicLinkToken string `protobuf:"bytes,5,opt,name=public_link_token,json=publicLinkToken,proto3" json:"public_link_token,omitempty"`
}

func (x *WebdavSource) Reset() {
	*x = WebdavSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thumbnails_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebdavSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebdavSource) ProtoMessage() {}

func (x *WebdavSource) ProtoReflect() protoreflect.Message {
	mi := &file_thumbnails_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebdavSource.ProtoReflect.Descriptor instead.
func (*WebdavSource) Descriptor() ([]byte, []int) {
	return file_thumbnails_proto_rawDescGZIP(), []int{1}
}

func (x *WebdavSource) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *WebdavSource) GetIsPublicLink() bool {
	if x != nil {
		return x.IsPublicLink
	}
	return false
}

func (x *WebdavSource) GetWebdavAuthorization() string {
	if x != nil {
		return x.WebdavAuthorization
	}
	return ""
}

func (x *WebdavSource) GetRevaAuthorization() string {
	if x != nil {
		return x.RevaAuthorization
	}
	return ""
}

func (x *WebdavSource) GetPublicLinkToken() string {
	if x != nil {
		return x.PublicLinkToken
	}
	return ""
}

type CS3Source struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path          string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Authorization string `protobuf:"bytes,2,opt,name=authorization,proto3" json:"authorization,omitempty"`
}

func (x *CS3Source) Reset() {
	*x = CS3Source{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thumbnails_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CS3Source) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CS3Source) ProtoMessage() {}

func (x *CS3Source) ProtoReflect() protoreflect.Message {
	mi := &file_thumbnails_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CS3Source.ProtoReflect.Descriptor instead.
func (*CS3Source) Descriptor() ([]byte, []int) {
	return file_thumbnails_proto_rawDescGZIP(), []int{2}
}

func (x *CS3Source) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *CS3Source) GetAuthorization() string {
	if x != nil {
		return x.Authorization
	}
	return ""
}

// The service response
type GetThumbnailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The thumbnail as a binary
	Thumbnail []byte `protobuf:"bytes,1,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
	// The mimetype of the thumbnail
	Mimetype string `protobuf:"bytes,2,opt,name=mimetype,proto3" json:"mimetype,omitempty"`
}

func (x *GetThumbnailResponse) Reset() {
	*x = GetThumbnailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thumbnails_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetThumbnailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetThumbnailResponse) ProtoMessage() {}

func (x *GetThumbnailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_thumbnails_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetThumbnailResponse.ProtoReflect.Descriptor instead.
func (*GetThumbnailResponse) Descriptor() ([]byte, []int) {
	return file_thumbnails_proto_rawDescGZIP(), []int{3}
}

func (x *GetThumbnailResponse) GetThumbnail() []byte {
	if x != nil {
		return x.Thumbnail
	}
	return nil
}

func (x *GetThumbnailResponse) GetMimetype() string {
	if x != nil {
		return x.Mimetype
	}
	return ""
}

var File_thumbnails_proto protoreflect.FileDescriptor

var file_thumbnails_proto_rawDesc = []byte{
	0x0a, 0x10, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x77, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6f, 0x63, 0x69, 0x73, 0x2e, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73,
	0x2e, 0x76, 0x30, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x03, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x54, 0x68, 0x75, 0x6d, 0x62,
	0x6e, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x12, 0x69, 0x0a, 0x0e, 0x74, 0x68, 0x75, 0x6d, 0x62,
	0x6e, 0x61, 0x69, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x42, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x77, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f,
	0x63, 0x69, 0x73, 0x2e, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x76,
	0x30, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0d, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x54, 0x0a, 0x0d, 0x77, 0x65, 0x62, 0x64, 0x61, 0x76, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x77,
	0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x63, 0x69, 0x73, 0x2e, 0x74, 0x68, 0x75, 0x6d,
	0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x76, 0x30, 0x2e, 0x57, 0x65, 0x62, 0x64, 0x61, 0x76,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x77, 0x65, 0x62, 0x64, 0x61, 0x76,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0a, 0x63, 0x73, 0x33, 0x5f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x6f, 0x77, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x63, 0x69, 0x73, 0x2e, 0x74,
	0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x76, 0x30, 0x2e, 0x43, 0x53, 0x33,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x00, 0x52, 0x09, 0x63, 0x73, 0x33, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x22, 0x21, 0x0a, 0x0d, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x07, 0x0a,
	0x03, 0x4a, 0x50, 0x47, 0x10, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x22, 0xd4, 0x01, 0x0a, 0x0c, 0x57, 0x65, 0x62, 0x64, 0x61, 0x76, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x31, 0x0a, 0x14, 0x77, 0x65, 0x62,
	0x64, 0x61, 0x76, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x77, 0x65, 0x62, 0x64, 0x61, 0x76, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x12,
	0x72, 0x65, 0x76, 0x61, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x72, 0x65, 0x76, 0x61, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4c, 0x69,
	0x6e, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x45, 0x0a, 0x09, 0x43, 0x53, 0x33, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x50,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62,
	0x6e, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x32, 0x8f, 0x01, 0x0a, 0x10, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7b, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54, 0x68, 0x75, 0x6d,
	0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x77, 0x6e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x63, 0x69, 0x73, 0x2e, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e,
	0x61, 0x69, 0x6c, 0x73, 0x2e, 0x76, 0x30, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x68, 0x75, 0x6d, 0x62,
	0x6e, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x6f, 0x77, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x63, 0x69, 0x73, 0x2e,
	0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x76, 0x30, 0x2e, 0x47, 0x65,
	0x74, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0xe0, 0x02, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6f, 0x77, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x63, 0x69, 0x73, 0x2f,
	0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x30, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x92, 0x41, 0xa4,
	0x02, 0x12, 0xb8, 0x01, 0x0a, 0x22, 0x6f, 0x77, 0x6e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x20, 0x49,
	0x6e, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x65, 0x20, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x20, 0x74, 0x68,
	0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x47, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x20, 0x47, 0x6d, 0x62, 0x48, 0x12, 0x20, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x77,
	0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x63, 0x69, 0x73, 0x1a, 0x14, 0x73, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x40, 0x6f, 0x77, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f,
	0x6d, 0x2a, 0x42, 0x0a, 0x0a, 0x41, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2d, 0x32, 0x2e, 0x30, 0x12,
	0x34, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x77, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x63, 0x69,
	0x73, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x4c, 0x49,
	0x43, 0x45, 0x4e, 0x53, 0x45, 0x32, 0x05, 0x31, 0x2e, 0x30, 0x2e, 0x30, 0x2a, 0x02, 0x01, 0x02,
	0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73,
	0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x6a, 0x73, 0x6f, 0x6e, 0x72, 0x3f, 0x0a, 0x10, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65,
	0x72, 0x20, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x12, 0x2b, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x6f, 0x77, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e,
	0x61, 0x69, 0x6c, 0x73, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_thumbnails_proto_rawDescOnce sync.Once
	file_thumbnails_proto_rawDescData = file_thumbnails_proto_rawDesc
)

func file_thumbnails_proto_rawDescGZIP() []byte {
	file_thumbnails_proto_rawDescOnce.Do(func() {
		file_thumbnails_proto_rawDescData = protoimpl.X.CompressGZIP(file_thumbnails_proto_rawDescData)
	})
	return file_thumbnails_proto_rawDescData
}

var file_thumbnails_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_thumbnails_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_thumbnails_proto_goTypes = []interface{}{
	(GetThumbnailRequest_ThumbnailType)(0), // 0: com.owncloud.ocis.thumbnails.v0.GetThumbnailRequest.ThumbnailType
	(*GetThumbnailRequest)(nil),            // 1: com.owncloud.ocis.thumbnails.v0.GetThumbnailRequest
	(*WebdavSource)(nil),                   // 2: com.owncloud.ocis.thumbnails.v0.WebdavSource
	(*CS3Source)(nil),                      // 3: com.owncloud.ocis.thumbnails.v0.CS3Source
	(*GetThumbnailResponse)(nil),           // 4: com.owncloud.ocis.thumbnails.v0.GetThumbnailResponse
}
var file_thumbnails_proto_depIdxs = []int32{
	0, // 0: com.owncloud.ocis.thumbnails.v0.GetThumbnailRequest.thumbnail_type:type_name -> com.owncloud.ocis.thumbnails.v0.GetThumbnailRequest.ThumbnailType
	2, // 1: com.owncloud.ocis.thumbnails.v0.GetThumbnailRequest.webdav_source:type_name -> com.owncloud.ocis.thumbnails.v0.WebdavSource
	3, // 2: com.owncloud.ocis.thumbnails.v0.GetThumbnailRequest.cs3_source:type_name -> com.owncloud.ocis.thumbnails.v0.CS3Source
	1, // 3: com.owncloud.ocis.thumbnails.v0.ThumbnailService.GetThumbnail:input_type -> com.owncloud.ocis.thumbnails.v0.GetThumbnailRequest
	4, // 4: com.owncloud.ocis.thumbnails.v0.ThumbnailService.GetThumbnail:output_type -> com.owncloud.ocis.thumbnails.v0.GetThumbnailResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_thumbnails_proto_init() }
func file_thumbnails_proto_init() {
	if File_thumbnails_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_thumbnails_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetThumbnailRequest); i {
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
		file_thumbnails_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebdavSource); i {
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
		file_thumbnails_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CS3Source); i {
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
		file_thumbnails_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetThumbnailResponse); i {
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
	file_thumbnails_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*GetThumbnailRequest_WebdavSource)(nil),
		(*GetThumbnailRequest_Cs3Source)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_thumbnails_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_thumbnails_proto_goTypes,
		DependencyIndexes: file_thumbnails_proto_depIdxs,
		EnumInfos:         file_thumbnails_proto_enumTypes,
		MessageInfos:      file_thumbnails_proto_msgTypes,
	}.Build()
	File_thumbnails_proto = out.File
	file_thumbnails_proto_rawDesc = nil
	file_thumbnails_proto_goTypes = nil
	file_thumbnails_proto_depIdxs = nil
}

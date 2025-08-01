// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: review_detail.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FindAllReviewDetailRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          int32                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize      int32                  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Search        string                 `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FindAllReviewDetailRequest) Reset() {
	*x = FindAllReviewDetailRequest{}
	mi := &file_review_detail_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindAllReviewDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllReviewDetailRequest) ProtoMessage() {}

func (x *FindAllReviewDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_review_detail_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllReviewDetailRequest.ProtoReflect.Descriptor instead.
func (*FindAllReviewDetailRequest) Descriptor() ([]byte, []int) {
	return file_review_detail_proto_rawDescGZIP(), []int{0}
}

func (x *FindAllReviewDetailRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *FindAllReviewDetailRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *FindAllReviewDetailRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type FindByIdReviewDetailRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FindByIdReviewDetailRequest) Reset() {
	*x = FindByIdReviewDetailRequest{}
	mi := &file_review_detail_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindByIdReviewDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByIdReviewDetailRequest) ProtoMessage() {}

func (x *FindByIdReviewDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_review_detail_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByIdReviewDetailRequest.ProtoReflect.Descriptor instead.
func (*FindByIdReviewDetailRequest) Descriptor() ([]byte, []int) {
	return file_review_detail_proto_rawDescGZIP(), []int{1}
}

func (x *FindByIdReviewDetailRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CreateReviewDetailRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ReviewId      int32                  `protobuf:"varint,1,opt,name=review_id,json=reviewId,proto3" json:"review_id,omitempty"`
	Type          string                 `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Url           string                 `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Caption       string                 `protobuf:"bytes,4,opt,name=caption,proto3" json:"caption,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateReviewDetailRequest) Reset() {
	*x = CreateReviewDetailRequest{}
	mi := &file_review_detail_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateReviewDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReviewDetailRequest) ProtoMessage() {}

func (x *CreateReviewDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_review_detail_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReviewDetailRequest.ProtoReflect.Descriptor instead.
func (*CreateReviewDetailRequest) Descriptor() ([]byte, []int) {
	return file_review_detail_proto_rawDescGZIP(), []int{2}
}

func (x *CreateReviewDetailRequest) GetReviewId() int32 {
	if x != nil {
		return x.ReviewId
	}
	return 0
}

func (x *CreateReviewDetailRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateReviewDetailRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *CreateReviewDetailRequest) GetCaption() string {
	if x != nil {
		return x.Caption
	}
	return ""
}

type UpdateReviewDetailRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	ReviewDetailId int32                  `protobuf:"varint,1,opt,name=review_detail_id,json=reviewDetailId,proto3" json:"review_detail_id,omitempty"`
	Type           string                 `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Url            string                 `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Caption        string                 `protobuf:"bytes,4,opt,name=caption,proto3" json:"caption,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *UpdateReviewDetailRequest) Reset() {
	*x = UpdateReviewDetailRequest{}
	mi := &file_review_detail_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateReviewDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReviewDetailRequest) ProtoMessage() {}

func (x *UpdateReviewDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_review_detail_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReviewDetailRequest.ProtoReflect.Descriptor instead.
func (*UpdateReviewDetailRequest) Descriptor() ([]byte, []int) {
	return file_review_detail_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateReviewDetailRequest) GetReviewDetailId() int32 {
	if x != nil {
		return x.ReviewDetailId
	}
	return 0
}

func (x *UpdateReviewDetailRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *UpdateReviewDetailRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *UpdateReviewDetailRequest) GetCaption() string {
	if x != nil {
		return x.Caption
	}
	return ""
}

type ReviewDetailsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ReviewId      int32                  `protobuf:"varint,2,opt,name=review_id,json=reviewId,proto3" json:"review_id,omitempty"`
	Type          string                 `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Url           string                 `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	Caption       string                 `protobuf:"bytes,5,opt,name=caption,proto3" json:"caption,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReviewDetailsResponse) Reset() {
	*x = ReviewDetailsResponse{}
	mi := &file_review_detail_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReviewDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewDetailsResponse) ProtoMessage() {}

func (x *ReviewDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_review_detail_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewDetailsResponse.ProtoReflect.Descriptor instead.
func (*ReviewDetailsResponse) Descriptor() ([]byte, []int) {
	return file_review_detail_proto_rawDescGZIP(), []int{4}
}

func (x *ReviewDetailsResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ReviewDetailsResponse) GetReviewId() int32 {
	if x != nil {
		return x.ReviewId
	}
	return 0
}

func (x *ReviewDetailsResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ReviewDetailsResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ReviewDetailsResponse) GetCaption() string {
	if x != nil {
		return x.Caption
	}
	return ""
}

func (x *ReviewDetailsResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ReviewDetailsResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type ReviewDetailsResponseDeleteAt struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Id            int32                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ReviewId      int32                   `protobuf:"varint,2,opt,name=review_id,json=reviewId,proto3" json:"review_id,omitempty"`
	Type          string                  `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Url           string                  `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	Caption       string                  `protobuf:"bytes,5,opt,name=caption,proto3" json:"caption,omitempty"`
	CreatedAt     string                  `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string                  `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt     *wrapperspb.StringValue `protobuf:"bytes,8,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReviewDetailsResponseDeleteAt) Reset() {
	*x = ReviewDetailsResponseDeleteAt{}
	mi := &file_review_detail_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReviewDetailsResponseDeleteAt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewDetailsResponseDeleteAt) ProtoMessage() {}

func (x *ReviewDetailsResponseDeleteAt) ProtoReflect() protoreflect.Message {
	mi := &file_review_detail_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewDetailsResponseDeleteAt.ProtoReflect.Descriptor instead.
func (*ReviewDetailsResponseDeleteAt) Descriptor() ([]byte, []int) {
	return file_review_detail_proto_rawDescGZIP(), []int{5}
}

func (x *ReviewDetailsResponseDeleteAt) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ReviewDetailsResponseDeleteAt) GetReviewId() int32 {
	if x != nil {
		return x.ReviewId
	}
	return 0
}

func (x *ReviewDetailsResponseDeleteAt) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ReviewDetailsResponseDeleteAt) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ReviewDetailsResponseDeleteAt) GetCaption() string {
	if x != nil {
		return x.Caption
	}
	return ""
}

func (x *ReviewDetailsResponseDeleteAt) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ReviewDetailsResponseDeleteAt) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *ReviewDetailsResponseDeleteAt) GetDeletedAt() *wrapperspb.StringValue {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type ApiResponseReviewDetail struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data          *ReviewDetailsResponse `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ApiResponseReviewDetail) Reset() {
	*x = ApiResponseReviewDetail{}
	mi := &file_review_detail_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApiResponseReviewDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiResponseReviewDetail) ProtoMessage() {}

func (x *ApiResponseReviewDetail) ProtoReflect() protoreflect.Message {
	mi := &file_review_detail_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiResponseReviewDetail.ProtoReflect.Descriptor instead.
func (*ApiResponseReviewDetail) Descriptor() ([]byte, []int) {
	return file_review_detail_proto_rawDescGZIP(), []int{6}
}

func (x *ApiResponseReviewDetail) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ApiResponseReviewDetail) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ApiResponseReviewDetail) GetData() *ReviewDetailsResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

type ApiResponseReviewDetailDeleteAt struct {
	state         protoimpl.MessageState         `protogen:"open.v1"`
	Status        string                         `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message       string                         `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data          *ReviewDetailsResponseDeleteAt `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ApiResponseReviewDetailDeleteAt) Reset() {
	*x = ApiResponseReviewDetailDeleteAt{}
	mi := &file_review_detail_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApiResponseReviewDetailDeleteAt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiResponseReviewDetailDeleteAt) ProtoMessage() {}

func (x *ApiResponseReviewDetailDeleteAt) ProtoReflect() protoreflect.Message {
	mi := &file_review_detail_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiResponseReviewDetailDeleteAt.ProtoReflect.Descriptor instead.
func (*ApiResponseReviewDetailDeleteAt) Descriptor() ([]byte, []int) {
	return file_review_detail_proto_rawDescGZIP(), []int{7}
}

func (x *ApiResponseReviewDetailDeleteAt) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ApiResponseReviewDetailDeleteAt) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ApiResponseReviewDetailDeleteAt) GetData() *ReviewDetailsResponseDeleteAt {
	if x != nil {
		return x.Data
	}
	return nil
}

type ApiResponsesReviewDetails struct {
	state         protoimpl.MessageState   `protogen:"open.v1"`
	Status        string                   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message       string                   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data          []*ReviewDetailsResponse `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ApiResponsesReviewDetails) Reset() {
	*x = ApiResponsesReviewDetails{}
	mi := &file_review_detail_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApiResponsesReviewDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiResponsesReviewDetails) ProtoMessage() {}

func (x *ApiResponsesReviewDetails) ProtoReflect() protoreflect.Message {
	mi := &file_review_detail_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiResponsesReviewDetails.ProtoReflect.Descriptor instead.
func (*ApiResponsesReviewDetails) Descriptor() ([]byte, []int) {
	return file_review_detail_proto_rawDescGZIP(), []int{8}
}

func (x *ApiResponsesReviewDetails) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ApiResponsesReviewDetails) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ApiResponsesReviewDetails) GetData() []*ReviewDetailsResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

type ApiResponsePaginationReviewDetails struct {
	state         protoimpl.MessageState   `protogen:"open.v1"`
	Status        string                   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message       string                   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data          []*ReviewDetailsResponse `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	Pagination    *PaginationMeta          `protobuf:"bytes,4,opt,name=pagination,proto3" json:"pagination,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ApiResponsePaginationReviewDetails) Reset() {
	*x = ApiResponsePaginationReviewDetails{}
	mi := &file_review_detail_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApiResponsePaginationReviewDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiResponsePaginationReviewDetails) ProtoMessage() {}

func (x *ApiResponsePaginationReviewDetails) ProtoReflect() protoreflect.Message {
	mi := &file_review_detail_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiResponsePaginationReviewDetails.ProtoReflect.Descriptor instead.
func (*ApiResponsePaginationReviewDetails) Descriptor() ([]byte, []int) {
	return file_review_detail_proto_rawDescGZIP(), []int{9}
}

func (x *ApiResponsePaginationReviewDetails) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ApiResponsePaginationReviewDetails) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ApiResponsePaginationReviewDetails) GetData() []*ReviewDetailsResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ApiResponsePaginationReviewDetails) GetPagination() *PaginationMeta {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type ApiResponsePaginationReviewDetailsDeleteAt struct {
	state         protoimpl.MessageState           `protogen:"open.v1"`
	Status        string                           `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message       string                           `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data          []*ReviewDetailsResponseDeleteAt `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	Pagination    *PaginationMeta                  `protobuf:"bytes,4,opt,name=pagination,proto3" json:"pagination,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ApiResponsePaginationReviewDetailsDeleteAt) Reset() {
	*x = ApiResponsePaginationReviewDetailsDeleteAt{}
	mi := &file_review_detail_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApiResponsePaginationReviewDetailsDeleteAt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiResponsePaginationReviewDetailsDeleteAt) ProtoMessage() {}

func (x *ApiResponsePaginationReviewDetailsDeleteAt) ProtoReflect() protoreflect.Message {
	mi := &file_review_detail_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiResponsePaginationReviewDetailsDeleteAt.ProtoReflect.Descriptor instead.
func (*ApiResponsePaginationReviewDetailsDeleteAt) Descriptor() ([]byte, []int) {
	return file_review_detail_proto_rawDescGZIP(), []int{10}
}

func (x *ApiResponsePaginationReviewDetailsDeleteAt) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ApiResponsePaginationReviewDetailsDeleteAt) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ApiResponsePaginationReviewDetailsDeleteAt) GetData() []*ReviewDetailsResponseDeleteAt {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ApiResponsePaginationReviewDetailsDeleteAt) GetPagination() *PaginationMeta {
	if x != nil {
		return x.Pagination
	}
	return nil
}

var File_review_detail_proto protoreflect.FileDescriptor

const file_review_detail_proto_rawDesc = "" +
	"\n" +
	"\x13review_detail.proto\x12\x02pb\x1a\freview.proto\x1a\x1egoogle/protobuf/wrappers.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\tapi.proto\"e\n" +
	"\x1aFindAllReviewDetailRequest\x12\x12\n" +
	"\x04page\x18\x01 \x01(\x05R\x04page\x12\x1b\n" +
	"\tpage_size\x18\x02 \x01(\x05R\bpageSize\x12\x16\n" +
	"\x06search\x18\x03 \x01(\tR\x06search\"-\n" +
	"\x1bFindByIdReviewDetailRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\"x\n" +
	"\x19CreateReviewDetailRequest\x12\x1b\n" +
	"\treview_id\x18\x01 \x01(\x05R\breviewId\x12\x12\n" +
	"\x04type\x18\x02 \x01(\tR\x04type\x12\x10\n" +
	"\x03url\x18\x03 \x01(\tR\x03url\x12\x18\n" +
	"\acaption\x18\x04 \x01(\tR\acaption\"\x85\x01\n" +
	"\x19UpdateReviewDetailRequest\x12(\n" +
	"\x10review_detail_id\x18\x01 \x01(\x05R\x0ereviewDetailId\x12\x12\n" +
	"\x04type\x18\x02 \x01(\tR\x04type\x12\x10\n" +
	"\x03url\x18\x03 \x01(\tR\x03url\x12\x18\n" +
	"\acaption\x18\x04 \x01(\tR\acaption\"\xc2\x01\n" +
	"\x15ReviewDetailsResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x1b\n" +
	"\treview_id\x18\x02 \x01(\x05R\breviewId\x12\x12\n" +
	"\x04type\x18\x03 \x01(\tR\x04type\x12\x10\n" +
	"\x03url\x18\x04 \x01(\tR\x03url\x12\x18\n" +
	"\acaption\x18\x05 \x01(\tR\acaption\x12\x1d\n" +
	"\n" +
	"created_at\x18\x06 \x01(\tR\tcreatedAt\x12\x1d\n" +
	"\n" +
	"updated_at\x18\a \x01(\tR\tupdatedAt\"\x87\x02\n" +
	"\x1dReviewDetailsResponseDeleteAt\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x1b\n" +
	"\treview_id\x18\x02 \x01(\x05R\breviewId\x12\x12\n" +
	"\x04type\x18\x03 \x01(\tR\x04type\x12\x10\n" +
	"\x03url\x18\x04 \x01(\tR\x03url\x12\x18\n" +
	"\acaption\x18\x05 \x01(\tR\acaption\x12\x1d\n" +
	"\n" +
	"created_at\x18\x06 \x01(\tR\tcreatedAt\x12\x1d\n" +
	"\n" +
	"updated_at\x18\a \x01(\tR\tupdatedAt\x12;\n" +
	"\n" +
	"deleted_at\x18\b \x01(\v2\x1c.google.protobuf.StringValueR\tdeletedAt\"z\n" +
	"\x17ApiResponseReviewDetail\x12\x16\n" +
	"\x06status\x18\x01 \x01(\tR\x06status\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\x12-\n" +
	"\x04data\x18\x03 \x01(\v2\x19.pb.ReviewDetailsResponseR\x04data\"\x8a\x01\n" +
	"\x1fApiResponseReviewDetailDeleteAt\x12\x16\n" +
	"\x06status\x18\x01 \x01(\tR\x06status\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\x125\n" +
	"\x04data\x18\x03 \x01(\v2!.pb.ReviewDetailsResponseDeleteAtR\x04data\"|\n" +
	"\x19ApiResponsesReviewDetails\x12\x16\n" +
	"\x06status\x18\x01 \x01(\tR\x06status\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\x12-\n" +
	"\x04data\x18\x03 \x03(\v2\x19.pb.ReviewDetailsResponseR\x04data\"\xb9\x01\n" +
	"\"ApiResponsePaginationReviewDetails\x12\x16\n" +
	"\x06status\x18\x01 \x01(\tR\x06status\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\x12-\n" +
	"\x04data\x18\x03 \x03(\v2\x19.pb.ReviewDetailsResponseR\x04data\x122\n" +
	"\n" +
	"pagination\x18\x04 \x01(\v2\x12.pb.PaginationMetaR\n" +
	"pagination\"\xc9\x01\n" +
	"*ApiResponsePaginationReviewDetailsDeleteAt\x12\x16\n" +
	"\x06status\x18\x01 \x01(\tR\x06status\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\x125\n" +
	"\x04data\x18\x03 \x03(\v2!.pb.ReviewDetailsResponseDeleteAtR\x04data\x122\n" +
	"\n" +
	"pagination\x18\x04 \x01(\v2\x12.pb.PaginationMetaR\n" +
	"pagination2\xac\a\n" +
	"\x13ReviewDetailService\x12K\n" +
	"\aFindAll\x12\x18.pb.FindAllReviewRequest\x1a&.pb.ApiResponsePaginationReviewDetails\x12H\n" +
	"\bFindById\x12\x1f.pb.FindByIdReviewDetailRequest\x1a\x1b.pb.ApiResponseReviewDetail\x12Z\n" +
	"\fFindByActive\x12\x18.pb.FindAllReviewRequest\x1a..pb.ApiResponsePaginationReviewDetailsDeleteAt\"\x00\x12[\n" +
	"\rFindByTrashed\x12\x18.pb.FindAllReviewRequest\x1a..pb.ApiResponsePaginationReviewDetailsDeleteAt\"\x00\x12D\n" +
	"\x06Create\x12\x1d.pb.CreateReviewDetailRequest\x1a\x1b.pb.ApiResponseReviewDetail\x12D\n" +
	"\x06Update\x12\x1d.pb.UpdateReviewDetailRequest\x1a\x1b.pb.ApiResponseReviewDetail\x12[\n" +
	"\x13TrashedReviewDetail\x12\x1f.pb.FindByIdReviewDetailRequest\x1a#.pb.ApiResponseReviewDetailDeleteAt\x12[\n" +
	"\x13RestoreReviewDetail\x12\x1f.pb.FindByIdReviewDetailRequest\x1a#.pb.ApiResponseReviewDetailDeleteAt\x12[\n" +
	"\x1bDeleteReviewDetailPermanent\x12\x1f.pb.FindByIdReviewDetailRequest\x1a\x1b.pb.ApiResponseReviewDelete\x12L\n" +
	"\x16RestoreAllReviewDetail\x12\x16.google.protobuf.Empty\x1a\x18.pb.ApiResponseReviewAll\"\x00\x12T\n" +
	"\x1eDeleteAllReviewDetailPermanent\x12\x16.google.protobuf.Empty\x1a\x18.pb.ApiResponseReviewAll\"\x00B4Z2github.com/MamangRust/monolith-ecommerce-shared/pbb\x06proto3"

var (
	file_review_detail_proto_rawDescOnce sync.Once
	file_review_detail_proto_rawDescData []byte
)

func file_review_detail_proto_rawDescGZIP() []byte {
	file_review_detail_proto_rawDescOnce.Do(func() {
		file_review_detail_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_review_detail_proto_rawDesc), len(file_review_detail_proto_rawDesc)))
	})
	return file_review_detail_proto_rawDescData
}

var file_review_detail_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_review_detail_proto_goTypes = []any{
	(*FindAllReviewDetailRequest)(nil),                 // 0: pb.FindAllReviewDetailRequest
	(*FindByIdReviewDetailRequest)(nil),                // 1: pb.FindByIdReviewDetailRequest
	(*CreateReviewDetailRequest)(nil),                  // 2: pb.CreateReviewDetailRequest
	(*UpdateReviewDetailRequest)(nil),                  // 3: pb.UpdateReviewDetailRequest
	(*ReviewDetailsResponse)(nil),                      // 4: pb.ReviewDetailsResponse
	(*ReviewDetailsResponseDeleteAt)(nil),              // 5: pb.ReviewDetailsResponseDeleteAt
	(*ApiResponseReviewDetail)(nil),                    // 6: pb.ApiResponseReviewDetail
	(*ApiResponseReviewDetailDeleteAt)(nil),            // 7: pb.ApiResponseReviewDetailDeleteAt
	(*ApiResponsesReviewDetails)(nil),                  // 8: pb.ApiResponsesReviewDetails
	(*ApiResponsePaginationReviewDetails)(nil),         // 9: pb.ApiResponsePaginationReviewDetails
	(*ApiResponsePaginationReviewDetailsDeleteAt)(nil), // 10: pb.ApiResponsePaginationReviewDetailsDeleteAt
	(*wrapperspb.StringValue)(nil),                     // 11: google.protobuf.StringValue
	(*PaginationMeta)(nil),                             // 12: pb.PaginationMeta
	(*FindAllReviewRequest)(nil),                       // 13: pb.FindAllReviewRequest
	(*emptypb.Empty)(nil),                              // 14: google.protobuf.Empty
	(*ApiResponseReviewDelete)(nil),                    // 15: pb.ApiResponseReviewDelete
	(*ApiResponseReviewAll)(nil),                       // 16: pb.ApiResponseReviewAll
}
var file_review_detail_proto_depIdxs = []int32{
	11, // 0: pb.ReviewDetailsResponseDeleteAt.deleted_at:type_name -> google.protobuf.StringValue
	4,  // 1: pb.ApiResponseReviewDetail.data:type_name -> pb.ReviewDetailsResponse
	5,  // 2: pb.ApiResponseReviewDetailDeleteAt.data:type_name -> pb.ReviewDetailsResponseDeleteAt
	4,  // 3: pb.ApiResponsesReviewDetails.data:type_name -> pb.ReviewDetailsResponse
	4,  // 4: pb.ApiResponsePaginationReviewDetails.data:type_name -> pb.ReviewDetailsResponse
	12, // 5: pb.ApiResponsePaginationReviewDetails.pagination:type_name -> pb.PaginationMeta
	5,  // 6: pb.ApiResponsePaginationReviewDetailsDeleteAt.data:type_name -> pb.ReviewDetailsResponseDeleteAt
	12, // 7: pb.ApiResponsePaginationReviewDetailsDeleteAt.pagination:type_name -> pb.PaginationMeta
	13, // 8: pb.ReviewDetailService.FindAll:input_type -> pb.FindAllReviewRequest
	1,  // 9: pb.ReviewDetailService.FindById:input_type -> pb.FindByIdReviewDetailRequest
	13, // 10: pb.ReviewDetailService.FindByActive:input_type -> pb.FindAllReviewRequest
	13, // 11: pb.ReviewDetailService.FindByTrashed:input_type -> pb.FindAllReviewRequest
	2,  // 12: pb.ReviewDetailService.Create:input_type -> pb.CreateReviewDetailRequest
	3,  // 13: pb.ReviewDetailService.Update:input_type -> pb.UpdateReviewDetailRequest
	1,  // 14: pb.ReviewDetailService.TrashedReviewDetail:input_type -> pb.FindByIdReviewDetailRequest
	1,  // 15: pb.ReviewDetailService.RestoreReviewDetail:input_type -> pb.FindByIdReviewDetailRequest
	1,  // 16: pb.ReviewDetailService.DeleteReviewDetailPermanent:input_type -> pb.FindByIdReviewDetailRequest
	14, // 17: pb.ReviewDetailService.RestoreAllReviewDetail:input_type -> google.protobuf.Empty
	14, // 18: pb.ReviewDetailService.DeleteAllReviewDetailPermanent:input_type -> google.protobuf.Empty
	9,  // 19: pb.ReviewDetailService.FindAll:output_type -> pb.ApiResponsePaginationReviewDetails
	6,  // 20: pb.ReviewDetailService.FindById:output_type -> pb.ApiResponseReviewDetail
	10, // 21: pb.ReviewDetailService.FindByActive:output_type -> pb.ApiResponsePaginationReviewDetailsDeleteAt
	10, // 22: pb.ReviewDetailService.FindByTrashed:output_type -> pb.ApiResponsePaginationReviewDetailsDeleteAt
	6,  // 23: pb.ReviewDetailService.Create:output_type -> pb.ApiResponseReviewDetail
	6,  // 24: pb.ReviewDetailService.Update:output_type -> pb.ApiResponseReviewDetail
	7,  // 25: pb.ReviewDetailService.TrashedReviewDetail:output_type -> pb.ApiResponseReviewDetailDeleteAt
	7,  // 26: pb.ReviewDetailService.RestoreReviewDetail:output_type -> pb.ApiResponseReviewDetailDeleteAt
	15, // 27: pb.ReviewDetailService.DeleteReviewDetailPermanent:output_type -> pb.ApiResponseReviewDelete
	16, // 28: pb.ReviewDetailService.RestoreAllReviewDetail:output_type -> pb.ApiResponseReviewAll
	16, // 29: pb.ReviewDetailService.DeleteAllReviewDetailPermanent:output_type -> pb.ApiResponseReviewAll
	19, // [19:30] is the sub-list for method output_type
	8,  // [8:19] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_review_detail_proto_init() }
func file_review_detail_proto_init() {
	if File_review_detail_proto != nil {
		return
	}
	file_review_proto_init()
	file_api_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_review_detail_proto_rawDesc), len(file_review_detail_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_review_detail_proto_goTypes,
		DependencyIndexes: file_review_detail_proto_depIdxs,
		MessageInfos:      file_review_detail_proto_msgTypes,
	}.Build()
	File_review_detail_proto = out.File
	file_review_detail_proto_goTypes = nil
	file_review_detail_proto_depIdxs = nil
}

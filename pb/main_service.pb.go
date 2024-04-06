// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.3
// source: main_service.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListCoursesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize int32   `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	LastCode *string `protobuf:"bytes,3,opt,name=last_code,json=lastCode,proto3,oneof" json:"last_code,omitempty"`
}

func (x *ListCoursesRequest) Reset() {
	*x = ListCoursesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCoursesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCoursesRequest) ProtoMessage() {}

func (x *ListCoursesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_main_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCoursesRequest.ProtoReflect.Descriptor instead.
func (*ListCoursesRequest) Descriptor() ([]byte, []int) {
	return file_main_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListCoursesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListCoursesRequest) GetLastCode() string {
	if x != nil && x.LastCode != nil {
		return *x.LastCode
	}
	return ""
}

type SearchCourseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query    string  `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	PageSize int32   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	LastCode *string `protobuf:"bytes,3,opt,name=last_code,json=lastCode,proto3,oneof" json:"last_code,omitempty"`
}

func (x *SearchCourseRequest) Reset() {
	*x = SearchCourseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchCourseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchCourseRequest) ProtoMessage() {}

func (x *SearchCourseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_main_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchCourseRequest.ProtoReflect.Descriptor instead.
func (*SearchCourseRequest) Descriptor() ([]byte, []int) {
	return file_main_service_proto_rawDescGZIP(), []int{1}
}

func (x *SearchCourseRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *SearchCourseRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *SearchCourseRequest) GetLastCode() string {
	if x != nil && x.LastCode != nil {
		return *x.LastCode
	}
	return ""
}

type CoursesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Courses     []*Course `protobuf:"bytes,1,rep,name=courses,proto3" json:"courses,omitempty"`
	MoreResults bool      `protobuf:"varint,2,opt,name=moreResults,proto3" json:"moreResults,omitempty"`
}

func (x *CoursesResponse) Reset() {
	*x = CoursesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoursesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoursesResponse) ProtoMessage() {}

func (x *CoursesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_main_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoursesResponse.ProtoReflect.Descriptor instead.
func (*CoursesResponse) Descriptor() ([]byte, []int) {
	return file_main_service_proto_rawDescGZIP(), []int{2}
}

func (x *CoursesResponse) GetCourses() []*Course {
	if x != nil {
		return x.Courses
	}
	return nil
}

func (x *CoursesResponse) GetMoreResults() bool {
	if x != nil {
		return x.MoreResults
	}
	return false
}

type GetCourseDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseCode string `protobuf:"bytes,1,opt,name=course_code,json=courseCode,proto3" json:"course_code,omitempty"`
}

func (x *GetCourseDetailRequest) Reset() {
	*x = GetCourseDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCourseDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCourseDetailRequest) ProtoMessage() {}

func (x *GetCourseDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_main_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCourseDetailRequest.ProtoReflect.Descriptor instead.
func (*GetCourseDetailRequest) Descriptor() ([]byte, []int) {
	return file_main_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetCourseDetailRequest) GetCourseCode() string {
	if x != nil {
		return x.CourseCode
	}
	return ""
}

type GetCourseDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Course  *Course        `protobuf:"bytes,1,opt,name=course,proto3" json:"course,omitempty"`
	Reviews []*Review      `protobuf:"bytes,2,rep,name=reviews,proto3" json:"reviews,omitempty"`
	Meta    *AddReviewMeta `protobuf:"bytes,3,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *GetCourseDetailResponse) Reset() {
	*x = GetCourseDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCourseDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCourseDetailResponse) ProtoMessage() {}

func (x *GetCourseDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_main_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCourseDetailResponse.ProtoReflect.Descriptor instead.
func (*GetCourseDetailResponse) Descriptor() ([]byte, []int) {
	return file_main_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetCourseDetailResponse) GetCourse() *Course {
	if x != nil {
		return x.Course
	}
	return nil
}

func (x *GetCourseDetailResponse) GetReviews() []*Review {
	if x != nil {
		return x.Reviews
	}
	return nil
}

func (x *GetCourseDetailResponse) GetMeta() *AddReviewMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

type UpsertTakenCoursesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TakenCourseCodes []string `protobuf:"bytes,1,rep,name=takenCourseCodes,proto3" json:"takenCourseCodes,omitempty"`
}

func (x *UpsertTakenCoursesRequest) Reset() {
	*x = UpsertTakenCoursesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertTakenCoursesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertTakenCoursesRequest) ProtoMessage() {}

func (x *UpsertTakenCoursesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_main_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertTakenCoursesRequest.ProtoReflect.Descriptor instead.
func (*UpsertTakenCoursesRequest) Descriptor() ([]byte, []int) {
	return file_main_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpsertTakenCoursesRequest) GetTakenCourseCodes() []string {
	if x != nil {
		return x.TakenCourseCodes
	}
	return nil
}

type UpdateUserInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      *uint32 `protobuf:"varint,1,opt,name=uid,proto3,oneof" json:"uid,omitempty"`
	FullName *string `protobuf:"bytes,2,opt,name=fullName,proto3,oneof" json:"fullName,omitempty"`
}

func (x *UpdateUserInfoRequest) Reset() {
	*x = UpdateUserInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserInfoRequest) ProtoMessage() {}

func (x *UpdateUserInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_main_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserInfoRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserInfoRequest) Descriptor() ([]byte, []int) {
	return file_main_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateUserInfoRequest) GetUid() uint32 {
	if x != nil && x.Uid != nil {
		return *x.Uid
	}
	return 0
}

func (x *UpdateUserInfoRequest) GetFullName() string {
	if x != nil && x.FullName != nil {
		return *x.FullName
	}
	return ""
}

type GetUpdatedURLsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VersionTimestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=versionTimestamp,proto3,oneof" json:"versionTimestamp,omitempty"`
}

func (x *GetUpdatedURLsRequest) Reset() {
	*x = GetUpdatedURLsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUpdatedURLsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUpdatedURLsRequest) ProtoMessage() {}

func (x *GetUpdatedURLsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_main_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUpdatedURLsRequest.ProtoReflect.Descriptor instead.
func (*GetUpdatedURLsRequest) Descriptor() ([]byte, []int) {
	return file_main_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetUpdatedURLsRequest) GetVersionTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.VersionTimestamp
	}
	return nil
}

type GetUpdatedURLsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LatestURLs *GetUpdatedURLsResponse_URLsList `protobuf:"bytes,1,opt,name=latestURLs,proto3,oneof" json:"latestURLs,omitempty"`
}

func (x *GetUpdatedURLsResponse) Reset() {
	*x = GetUpdatedURLsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUpdatedURLsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUpdatedURLsResponse) ProtoMessage() {}

func (x *GetUpdatedURLsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_main_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUpdatedURLsResponse.ProtoReflect.Descriptor instead.
func (*GetUpdatedURLsResponse) Descriptor() ([]byte, []int) {
	return file_main_service_proto_rawDescGZIP(), []int{8}
}

func (x *GetUpdatedURLsResponse) GetLatestURLs() *GetUpdatedURLsResponse_URLsList {
	if x != nil {
		return x.LatestURLs
	}
	return nil
}

type GetUpdatedURLsResponse_URLsList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VersionTimestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=versionTimestamp,proto3" json:"versionTimestamp,omitempty"`
	URLs             map[string]string      `protobuf:"bytes,2,rep,name=URLs,proto3" json:"URLs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetUpdatedURLsResponse_URLsList) Reset() {
	*x = GetUpdatedURLsResponse_URLsList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUpdatedURLsResponse_URLsList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUpdatedURLsResponse_URLsList) ProtoMessage() {}

func (x *GetUpdatedURLsResponse_URLsList) ProtoReflect() protoreflect.Message {
	mi := &file_main_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUpdatedURLsResponse_URLsList.ProtoReflect.Descriptor instead.
func (*GetUpdatedURLsResponse_URLsList) Descriptor() ([]byte, []int) {
	return file_main_service_proto_rawDescGZIP(), []int{8, 0}
}

func (x *GetUpdatedURLsResponse_URLsList) GetVersionTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.VersionTimestamp
	}
	return nil
}

func (x *GetUpdatedURLsResponse_URLsList) GetURLs() map[string]string {
	if x != nil {
		return x.URLs
	}
	return nil
}

var File_main_service_proto protoreflect.FileDescriptor

var file_main_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x12, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x20, 0x0a, 0x09,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x78, 0x0a, 0x13,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x20, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x6c, 0x61, 0x73,
	0x74, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x5e, 0x0a, 0x0f, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x07, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x6d, 0x6f, 0x72, 0x65, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x39, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x43, 0x6f, 0x64,
	0x65, 0x22, 0x99, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a,
	0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x06,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x73, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x07, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x73, 0x12, 0x2a, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x47, 0x0a,
	0x19, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x54, 0x61, 0x6b, 0x65, 0x6e, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x74, 0x61,
	0x6b, 0x65, 0x6e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x74, 0x61, 0x6b, 0x65, 0x6e, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x64, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x15, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x03,
	0x75, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x75, 0x69, 0x64, 0x42,
	0x0b, 0x0a, 0x09, 0x5f, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x79, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x55, 0x52, 0x4c, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x10, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x10, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x88,
	0x01, 0x01, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xcc, 0x02, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x55, 0x52, 0x4c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0a, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x55, 0x52, 0x4c, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x55, 0x52, 0x4c, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x55, 0x52, 0x4c, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x48, 0x00, 0x52, 0x0a, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x55, 0x52, 0x4c, 0x73, 0x88, 0x01,
	0x01, 0x1a, 0xd3, 0x01, 0x0a, 0x08, 0x55, 0x52, 0x4c, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x46,
	0x0a, 0x10, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x10, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x46, 0x0a, 0x04, 0x55, 0x52, 0x4c, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x55, 0x52, 0x4c, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x55, 0x52, 0x4c, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x55,
	0x52, 0x4c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x55, 0x52, 0x4c, 0x73, 0x1a, 0x37,
	0x0a, 0x09, 0x55, 0x52, 0x4c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6c, 0x61, 0x74, 0x65,
	0x73, 0x74, 0x55, 0x52, 0x4c, 0x73, 0x32, 0xa6, 0x04, 0x0a, 0x0b, 0x4d, 0x61, 0x69, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x55, 0x52, 0x4c, 0x73, 0x12, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x55, 0x52, 0x4c,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x55, 0x52, 0x4c,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x50, 0x0a, 0x12, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x54, 0x61, 0x6b,
	0x65, 0x6e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x12, 0x22, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x54, 0x61, 0x6b, 0x65, 0x6e, 0x43,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x44, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x73, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0d, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x09, 0x41,
	0x64, 0x64, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x19, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x41, 0x64,
	0x64, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x0c, 0x5a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_main_service_proto_rawDescOnce sync.Once
	file_main_service_proto_rawDescData = file_main_service_proto_rawDesc
)

func file_main_service_proto_rawDescGZIP() []byte {
	file_main_service_proto_rawDescOnce.Do(func() {
		file_main_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_main_service_proto_rawDescData)
	})
	return file_main_service_proto_rawDescData
}

var file_main_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_main_service_proto_goTypes = []interface{}{
	(*ListCoursesRequest)(nil),              // 0: service.ListCoursesRequest
	(*SearchCourseRequest)(nil),             // 1: service.SearchCourseRequest
	(*CoursesResponse)(nil),                 // 2: service.CoursesResponse
	(*GetCourseDetailRequest)(nil),          // 3: service.GetCourseDetailRequest
	(*GetCourseDetailResponse)(nil),         // 4: service.GetCourseDetailResponse
	(*UpsertTakenCoursesRequest)(nil),       // 5: service.UpsertTakenCoursesRequest
	(*UpdateUserInfoRequest)(nil),           // 6: service.UpdateUserInfoRequest
	(*GetUpdatedURLsRequest)(nil),           // 7: service.GetUpdatedURLsRequest
	(*GetUpdatedURLsResponse)(nil),          // 8: service.GetUpdatedURLsResponse
	(*GetUpdatedURLsResponse_URLsList)(nil), // 9: service.GetUpdatedURLsResponse.URLsList
	nil,                                     // 10: service.GetUpdatedURLsResponse.URLsList.URLsEntry
	(*Course)(nil),                          // 11: courses.Course
	(*Review)(nil),                          // 12: reviews.Review
	(*AddReviewMeta)(nil),                   // 13: reviews.AddReviewMeta
	(*timestamppb.Timestamp)(nil),           // 14: google.protobuf.Timestamp
	(*AddReviewRequest)(nil),                // 15: reviews.AddReviewRequest
	(*emptypb.Empty)(nil),                   // 16: google.protobuf.Empty
	(*AddReviewResponse)(nil),               // 17: reviews.AddReviewResponse
}
var file_main_service_proto_depIdxs = []int32{
	11, // 0: service.CoursesResponse.courses:type_name -> courses.Course
	11, // 1: service.GetCourseDetailResponse.course:type_name -> courses.Course
	12, // 2: service.GetCourseDetailResponse.reviews:type_name -> reviews.Review
	13, // 3: service.GetCourseDetailResponse.meta:type_name -> reviews.AddReviewMeta
	14, // 4: service.GetUpdatedURLsRequest.versionTimestamp:type_name -> google.protobuf.Timestamp
	9,  // 5: service.GetUpdatedURLsResponse.latestURLs:type_name -> service.GetUpdatedURLsResponse.URLsList
	14, // 6: service.GetUpdatedURLsResponse.URLsList.versionTimestamp:type_name -> google.protobuf.Timestamp
	10, // 7: service.GetUpdatedURLsResponse.URLsList.URLs:type_name -> service.GetUpdatedURLsResponse.URLsList.URLsEntry
	7,  // 8: service.MainService.GetUpdatedURLs:input_type -> service.GetUpdatedURLsRequest
	6,  // 9: service.MainService.UpdateUserInfo:input_type -> service.UpdateUserInfoRequest
	5,  // 10: service.MainService.UpsertTakenCourses:input_type -> service.UpsertTakenCoursesRequest
	0,  // 11: service.MainService.ListCourses:input_type -> service.ListCoursesRequest
	1,  // 12: service.MainService.SearchCourses:input_type -> service.SearchCourseRequest
	3,  // 13: service.MainService.GetCourseDetails:input_type -> service.GetCourseDetailRequest
	15, // 14: service.MainService.AddReview:input_type -> reviews.AddReviewRequest
	8,  // 15: service.MainService.GetUpdatedURLs:output_type -> service.GetUpdatedURLsResponse
	16, // 16: service.MainService.UpdateUserInfo:output_type -> google.protobuf.Empty
	16, // 17: service.MainService.UpsertTakenCourses:output_type -> google.protobuf.Empty
	2,  // 18: service.MainService.ListCourses:output_type -> service.CoursesResponse
	2,  // 19: service.MainService.SearchCourses:output_type -> service.CoursesResponse
	4,  // 20: service.MainService.GetCourseDetails:output_type -> service.GetCourseDetailResponse
	17, // 21: service.MainService.AddReview:output_type -> reviews.AddReviewResponse
	15, // [15:22] is the sub-list for method output_type
	8,  // [8:15] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_main_service_proto_init() }
func file_main_service_proto_init() {
	if File_main_service_proto != nil {
		return
	}
	file_courses_proto_init()
	file_reviews_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_main_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCoursesRequest); i {
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
		file_main_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchCourseRequest); i {
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
		file_main_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoursesResponse); i {
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
		file_main_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCourseDetailRequest); i {
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
		file_main_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCourseDetailResponse); i {
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
		file_main_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertTakenCoursesRequest); i {
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
		file_main_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserInfoRequest); i {
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
		file_main_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUpdatedURLsRequest); i {
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
		file_main_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUpdatedURLsResponse); i {
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
		file_main_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUpdatedURLsResponse_URLsList); i {
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
	file_main_service_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_main_service_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_main_service_proto_msgTypes[6].OneofWrappers = []interface{}{}
	file_main_service_proto_msgTypes[7].OneofWrappers = []interface{}{}
	file_main_service_proto_msgTypes[8].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_main_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_main_service_proto_goTypes,
		DependencyIndexes: file_main_service_proto_depIdxs,
		MessageInfos:      file_main_service_proto_msgTypes,
	}.Build()
	File_main_service_proto = out.File
	file_main_service_proto_rawDesc = nil
	file_main_service_proto_goTypes = nil
	file_main_service_proto_depIdxs = nil
}

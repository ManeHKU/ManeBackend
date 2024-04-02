// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.3
// source: reviews.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type AddReviewResult int32

const (
	AddReviewResult_SUCCESS                     AddReviewResult = 0
	AddReviewResult_INVALID_COURSE_CODE         AddReviewResult = 1
	AddReviewResult_INVALID_YEAR_TAKEN          AddReviewResult = 2
	AddReviewResult_INVALID_SEMESTER_TAKEN      AddReviewResult = 3
	AddReviewResult_INVALID_CONTENT             AddReviewResult = 4
	AddReviewResult_INVALID_LECTURERS           AddReviewResult = 5
	AddReviewResult_INVALID_RATING              AddReviewResult = 6
	AddReviewResult_ERROR_ALREADY_REVIEWED      AddReviewResult = 7
	AddReviewResult_ERROR_USER_NOT_TAKEN_COURSE AddReviewResult = 8
)

// Enum value maps for AddReviewResult.
var (
	AddReviewResult_name = map[int32]string{
		0: "SUCCESS",
		1: "INVALID_COURSE_CODE",
		2: "INVALID_YEAR_TAKEN",
		3: "INVALID_SEMESTER_TAKEN",
		4: "INVALID_CONTENT",
		5: "INVALID_LECTURERS",
		6: "INVALID_RATING",
		7: "ERROR_ALREADY_REVIEWED",
		8: "ERROR_USER_NOT_TAKEN_COURSE",
	}
	AddReviewResult_value = map[string]int32{
		"SUCCESS":                     0,
		"INVALID_COURSE_CODE":         1,
		"INVALID_YEAR_TAKEN":          2,
		"INVALID_SEMESTER_TAKEN":      3,
		"INVALID_CONTENT":             4,
		"INVALID_LECTURERS":           5,
		"INVALID_RATING":              6,
		"ERROR_ALREADY_REVIEWED":      7,
		"ERROR_USER_NOT_TAKEN_COURSE": 8,
	}
)

func (x AddReviewResult) Enum() *AddReviewResult {
	p := new(AddReviewResult)
	*p = x
	return p
}

func (x AddReviewResult) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AddReviewResult) Descriptor() protoreflect.EnumDescriptor {
	return file_reviews_proto_enumTypes[0].Descriptor()
}

func (AddReviewResult) Type() protoreflect.EnumType {
	return &file_reviews_proto_enumTypes[0]
}

func (x AddReviewResult) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AddReviewResult.Descriptor instead.
func (AddReviewResult) EnumDescriptor() ([]byte, []int) {
	return file_reviews_proto_rawDescGZIP(), []int{0}
}

type AcademicYear int32

const (
	AcademicYear_AY_2018_2019 AcademicYear = 0
	AcademicYear_AY_2019_2020 AcademicYear = 1
	AcademicYear_AY_2020_2021 AcademicYear = 2
	AcademicYear_AY_2021_2022 AcademicYear = 3
	AcademicYear_AY_2022_2023 AcademicYear = 4
	AcademicYear_AY_2023_2024 AcademicYear = 5
	AcademicYear_AY_2024_2025 AcademicYear = 6
	AcademicYear_AY_2025_2026 AcademicYear = 7
	AcademicYear_AY_2026_2027 AcademicYear = 8
	AcademicYear_AY_2027_2028 AcademicYear = 9
	AcademicYear_AY_2028_2029 AcademicYear = 10
	AcademicYear_AY_2029_2030 AcademicYear = 11
)

// Enum value maps for AcademicYear.
var (
	AcademicYear_name = map[int32]string{
		0:  "AY_2018_2019",
		1:  "AY_2019_2020",
		2:  "AY_2020_2021",
		3:  "AY_2021_2022",
		4:  "AY_2022_2023",
		5:  "AY_2023_2024",
		6:  "AY_2024_2025",
		7:  "AY_2025_2026",
		8:  "AY_2026_2027",
		9:  "AY_2027_2028",
		10: "AY_2028_2029",
		11: "AY_2029_2030",
	}
	AcademicYear_value = map[string]int32{
		"AY_2018_2019": 0,
		"AY_2019_2020": 1,
		"AY_2020_2021": 2,
		"AY_2021_2022": 3,
		"AY_2022_2023": 4,
		"AY_2023_2024": 5,
		"AY_2024_2025": 6,
		"AY_2025_2026": 7,
		"AY_2026_2027": 8,
		"AY_2027_2028": 9,
		"AY_2028_2029": 10,
		"AY_2029_2030": 11,
	}
)

func (x AcademicYear) Enum() *AcademicYear {
	p := new(AcademicYear)
	*p = x
	return p
}

func (x AcademicYear) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AcademicYear) Descriptor() protoreflect.EnumDescriptor {
	return file_reviews_proto_enumTypes[1].Descriptor()
}

func (AcademicYear) Type() protoreflect.EnumType {
	return &file_reviews_proto_enumTypes[1]
}

func (x AcademicYear) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AcademicYear.Descriptor instead.
func (AcademicYear) EnumDescriptor() ([]byte, []int) {
	return file_reviews_proto_rawDescGZIP(), []int{1}
}

type Semester int32

const (
	Semester_SEM1   Semester = 0
	Semester_SEM2   Semester = 1
	Semester_SUMMER Semester = 2
)

// Enum value maps for Semester.
var (
	Semester_name = map[int32]string{
		0: "SEM1",
		1: "SEM2",
		2: "SUMMER",
	}
	Semester_value = map[string]int32{
		"SEM1":   0,
		"SEM2":   1,
		"SUMMER": 2,
	}
)

func (x Semester) Enum() *Semester {
	p := new(Semester)
	*p = x
	return p
}

func (x Semester) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Semester) Descriptor() protoreflect.EnumDescriptor {
	return file_reviews_proto_enumTypes[2].Descriptor()
}

func (Semester) Type() protoreflect.EnumType {
	return &file_reviews_proto_enumTypes[2]
}

func (x Semester) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Semester.Descriptor instead.
func (Semester) EnumDescriptor() ([]byte, []int) {
	return file_reviews_proto_rawDescGZIP(), []int{2}
}

type AddReviewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseCode    string       `protobuf:"bytes,1,opt,name=course_code,json=courseCode,proto3" json:"course_code,omitempty"`
	YearTaken     AcademicYear `protobuf:"varint,2,opt,name=year_taken,json=yearTaken,proto3,enum=reviews.AcademicYear" json:"year_taken,omitempty"`
	SemesterTaken Semester     `protobuf:"varint,3,opt,name=semester_taken,json=semesterTaken,proto3,enum=reviews.Semester" json:"semester_taken,omitempty"`
	Content       string       `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Lecturers     []string     `protobuf:"bytes,5,rep,name=lecturers,proto3" json:"lecturers,omitempty"`
	Rating        int32        `protobuf:"varint,6,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *AddReviewRequest) Reset() {
	*x = AddReviewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reviews_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddReviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReviewRequest) ProtoMessage() {}

func (x *AddReviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reviews_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddReviewRequest.ProtoReflect.Descriptor instead.
func (*AddReviewRequest) Descriptor() ([]byte, []int) {
	return file_reviews_proto_rawDescGZIP(), []int{0}
}

func (x *AddReviewRequest) GetCourseCode() string {
	if x != nil {
		return x.CourseCode
	}
	return ""
}

func (x *AddReviewRequest) GetYearTaken() AcademicYear {
	if x != nil {
		return x.YearTaken
	}
	return AcademicYear_AY_2018_2019
}

func (x *AddReviewRequest) GetSemesterTaken() Semester {
	if x != nil {
		return x.SemesterTaken
	}
	return Semester_SEM1
}

func (x *AddReviewRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *AddReviewRequest) GetLecturers() []string {
	if x != nil {
		return x.Lecturers
	}
	return nil
}

func (x *AddReviewRequest) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

type AddReviewResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result       AddReviewResult `protobuf:"varint,1,opt,name=result,proto3,enum=reviews.AddReviewResult" json:"result,omitempty"`
	ErrorMessage *string         `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3,oneof" json:"error_message,omitempty"`
}

func (x *AddReviewResponse) Reset() {
	*x = AddReviewResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reviews_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddReviewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReviewResponse) ProtoMessage() {}

func (x *AddReviewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reviews_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddReviewResponse.ProtoReflect.Descriptor instead.
func (*AddReviewResponse) Descriptor() ([]byte, []int) {
	return file_reviews_proto_rawDescGZIP(), []int{1}
}

func (x *AddReviewResponse) GetResult() AddReviewResult {
	if x != nil {
		return x.Result
	}
	return AddReviewResult_SUCCESS
}

func (x *AddReviewResponse) GetErrorMessage() string {
	if x != nil && x.ErrorMessage != nil {
		return *x.ErrorMessage
	}
	return ""
}

type Review struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseCode    string                 `protobuf:"bytes,1,opt,name=course_code,json=courseCode,proto3" json:"course_code,omitempty"`
	YearTaken     AcademicYear           `protobuf:"varint,2,opt,name=year_taken,json=yearTaken,proto3,enum=reviews.AcademicYear" json:"year_taken,omitempty"`
	SemesterTaken Semester               `protobuf:"varint,3,opt,name=semester_taken,json=semesterTaken,proto3,enum=reviews.Semester" json:"semester_taken,omitempty"`
	Content       string                 `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Rating        int32                  `protobuf:"varint,5,opt,name=rating,proto3" json:"rating,omitempty"`
	Lecturers     []string               `protobuf:"bytes,6,rep,name=lecturers,proto3" json:"lecturers,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Review) Reset() {
	*x = Review{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reviews_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Review) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Review) ProtoMessage() {}

func (x *Review) ProtoReflect() protoreflect.Message {
	mi := &file_reviews_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Review.ProtoReflect.Descriptor instead.
func (*Review) Descriptor() ([]byte, []int) {
	return file_reviews_proto_rawDescGZIP(), []int{2}
}

func (x *Review) GetCourseCode() string {
	if x != nil {
		return x.CourseCode
	}
	return ""
}

func (x *Review) GetYearTaken() AcademicYear {
	if x != nil {
		return x.YearTaken
	}
	return AcademicYear_AY_2018_2019
}

func (x *Review) GetSemesterTaken() Semester {
	if x != nil {
		return x.SemesterTaken
	}
	return Semester_SEM1
}

func (x *Review) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Review) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Review) GetLecturers() []string {
	if x != nil {
		return x.Lecturers
	}
	return nil
}

func (x *Review) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_reviews_proto protoreflect.FileDescriptor

var file_reviews_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf3, 0x01, 0x0a, 0x10, 0x41, 0x64,
	0x64, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x34, 0x0a, 0x0a, 0x79, 0x65, 0x61, 0x72, 0x5f, 0x74, 0x61, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x41, 0x63,
	0x61, 0x64, 0x65, 0x6d, 0x69, 0x63, 0x59, 0x65, 0x61, 0x72, 0x52, 0x09, 0x79, 0x65, 0x61, 0x72,
	0x54, 0x61, 0x6b, 0x65, 0x6e, 0x12, 0x38, 0x0a, 0x0e, 0x73, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x74, 0x61, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x53, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x0d, 0x73, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x54, 0x61, 0x6b, 0x65, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x65, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x65,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22,
	0x81, 0x01, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x28, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01,
	0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0xa4, 0x02, 0x0a, 0x06, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x34, 0x0a, 0x0a, 0x79, 0x65, 0x61, 0x72, 0x5f, 0x74, 0x61, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x41, 0x63,
	0x61, 0x64, 0x65, 0x6d, 0x69, 0x63, 0x59, 0x65, 0x61, 0x72, 0x52, 0x09, 0x79, 0x65, 0x61, 0x72,
	0x54, 0x61, 0x6b, 0x65, 0x6e, 0x12, 0x38, 0x0a, 0x0e, 0x73, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x74, 0x61, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x53, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x0d, 0x73, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x54, 0x61, 0x6b, 0x65, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x73, 0x12,
	0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x2a, 0xe8, 0x01, 0x0a, 0x0f, 0x41,
	0x64, 0x64, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0b,
	0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43, 0x4f, 0x55, 0x52, 0x53, 0x45, 0x5f, 0x43, 0x4f,
	0x44, 0x45, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f,
	0x59, 0x45, 0x41, 0x52, 0x5f, 0x54, 0x41, 0x4b, 0x45, 0x4e, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x53, 0x45, 0x4d, 0x45, 0x53, 0x54, 0x45, 0x52,
	0x5f, 0x54, 0x41, 0x4b, 0x45, 0x4e, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x10, 0x04, 0x12, 0x15, 0x0a,
	0x11, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x4c, 0x45, 0x43, 0x54, 0x55, 0x52, 0x45,
	0x52, 0x53, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f,
	0x52, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x06, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57,
	0x45, 0x44, 0x10, 0x07, 0x12, 0x1f, 0x0a, 0x1b, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x55, 0x53,
	0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x54, 0x41, 0x4b, 0x45, 0x4e, 0x5f, 0x43, 0x4f, 0x55,
	0x52, 0x53, 0x45, 0x10, 0x08, 0x2a, 0xe6, 0x01, 0x0a, 0x0c, 0x41, 0x63, 0x61, 0x64, 0x65, 0x6d,
	0x69, 0x63, 0x59, 0x65, 0x61, 0x72, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x59, 0x5f, 0x32, 0x30, 0x31,
	0x38, 0x5f, 0x32, 0x30, 0x31, 0x39, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x59, 0x5f, 0x32,
	0x30, 0x31, 0x39, 0x5f, 0x32, 0x30, 0x32, 0x30, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x59,
	0x5f, 0x32, 0x30, 0x32, 0x30, 0x5f, 0x32, 0x30, 0x32, 0x31, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c,
	0x41, 0x59, 0x5f, 0x32, 0x30, 0x32, 0x31, 0x5f, 0x32, 0x30, 0x32, 0x32, 0x10, 0x03, 0x12, 0x10,
	0x0a, 0x0c, 0x41, 0x59, 0x5f, 0x32, 0x30, 0x32, 0x32, 0x5f, 0x32, 0x30, 0x32, 0x33, 0x10, 0x04,
	0x12, 0x10, 0x0a, 0x0c, 0x41, 0x59, 0x5f, 0x32, 0x30, 0x32, 0x33, 0x5f, 0x32, 0x30, 0x32, 0x34,
	0x10, 0x05, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x59, 0x5f, 0x32, 0x30, 0x32, 0x34, 0x5f, 0x32, 0x30,
	0x32, 0x35, 0x10, 0x06, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x59, 0x5f, 0x32, 0x30, 0x32, 0x35, 0x5f,
	0x32, 0x30, 0x32, 0x36, 0x10, 0x07, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x59, 0x5f, 0x32, 0x30, 0x32,
	0x36, 0x5f, 0x32, 0x30, 0x32, 0x37, 0x10, 0x08, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x59, 0x5f, 0x32,
	0x30, 0x32, 0x37, 0x5f, 0x32, 0x30, 0x32, 0x38, 0x10, 0x09, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x59,
	0x5f, 0x32, 0x30, 0x32, 0x38, 0x5f, 0x32, 0x30, 0x32, 0x39, 0x10, 0x0a, 0x12, 0x10, 0x0a, 0x0c,
	0x41, 0x59, 0x5f, 0x32, 0x30, 0x32, 0x39, 0x5f, 0x32, 0x30, 0x33, 0x30, 0x10, 0x0b, 0x2a, 0x2a,
	0x0a, 0x08, 0x53, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x45,
	0x4d, 0x31, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x45, 0x4d, 0x32, 0x10, 0x01, 0x12, 0x0a,
	0x0a, 0x06, 0x53, 0x55, 0x4d, 0x4d, 0x45, 0x52, 0x10, 0x02, 0x42, 0x0c, 0x5a, 0x0a, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reviews_proto_rawDescOnce sync.Once
	file_reviews_proto_rawDescData = file_reviews_proto_rawDesc
)

func file_reviews_proto_rawDescGZIP() []byte {
	file_reviews_proto_rawDescOnce.Do(func() {
		file_reviews_proto_rawDescData = protoimpl.X.CompressGZIP(file_reviews_proto_rawDescData)
	})
	return file_reviews_proto_rawDescData
}

var file_reviews_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_reviews_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_reviews_proto_goTypes = []interface{}{
	(AddReviewResult)(0),          // 0: reviews.AddReviewResult
	(AcademicYear)(0),             // 1: reviews.AcademicYear
	(Semester)(0),                 // 2: reviews.Semester
	(*AddReviewRequest)(nil),      // 3: reviews.AddReviewRequest
	(*AddReviewResponse)(nil),     // 4: reviews.AddReviewResponse
	(*Review)(nil),                // 5: reviews.Review
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_reviews_proto_depIdxs = []int32{
	1, // 0: reviews.AddReviewRequest.year_taken:type_name -> reviews.AcademicYear
	2, // 1: reviews.AddReviewRequest.semester_taken:type_name -> reviews.Semester
	0, // 2: reviews.AddReviewResponse.result:type_name -> reviews.AddReviewResult
	1, // 3: reviews.Review.year_taken:type_name -> reviews.AcademicYear
	2, // 4: reviews.Review.semester_taken:type_name -> reviews.Semester
	6, // 5: reviews.Review.created_at:type_name -> google.protobuf.Timestamp
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_reviews_proto_init() }
func file_reviews_proto_init() {
	if File_reviews_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reviews_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddReviewRequest); i {
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
		file_reviews_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddReviewResponse); i {
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
		file_reviews_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Review); i {
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
	file_reviews_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_reviews_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_reviews_proto_goTypes,
		DependencyIndexes: file_reviews_proto_depIdxs,
		EnumInfos:         file_reviews_proto_enumTypes,
		MessageInfos:      file_reviews_proto_msgTypes,
	}.Build()
	File_reviews_proto = out.File
	file_reviews_proto_rawDesc = nil
	file_reviews_proto_goTypes = nil
	file_reviews_proto_depIdxs = nil
}

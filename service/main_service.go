package service

import (
	"ManeBackend/constants"
	"ManeBackend/internal/jwt"
	"ManeBackend/models"
	"ManeBackend/pb"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

type MainService struct {
	jwtManager *jwt.Manager
	pb.UnimplementedMainServiceServer
}

func (s *MainService) GetUpdatedURLs(_ context.Context, request *pb.GetUpdatedURLsRequest) (*pb.GetUpdatedURLsResponse, error) {
	if request.GetVersionTimestamp() == nil {
		return &pb.GetUpdatedURLsResponse{
			LatestURLs: nil,
		}, nil
	}
	println(request.GetVersionTimestamp().AsTime().String())
	ts := request.GetVersionTimestamp().AsTime()
	println(ts.String())
	URLs := make(map[string]string)
	URLs["home"] = "12312f.com"
	LatestUrls := &pb.GetUpdatedURLsResponse_URLsList{
		VersionTimestamp: timestamppb.Now(),
		URLs:             URLs,
	}
	return &pb.GetUpdatedURLsResponse{
		LatestURLs: LatestUrls,
	}, nil
}

func (s *MainService) UpdateUserInfo(context context.Context, request *pb.UpdateUserInfoRequest) (*emptypb.Empty, error) {
	if request.GetUid() == 0 && request.GetFullName() == "" {
		log.Printf("empty request not doing anything")
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}
	jwtClaims := context.Value(constants.JWTClaimsKey).(*jwt.UserClaims)
	userUUID := jwtClaims.Subject
	if userUUID == "" {
		log.Printf("empty uuid or error")
		return &emptypb.Empty{}, nil
	}

	log.Printf("current user id: %v", userUUID)
	if !models.CheckUserExist(userUUID) {
		log.Printf("user not exist")
		return nil, status.Errorf(codes.NotFound, "user not exist")
	}
	log.Printf("user %v exists", userUUID)

	if err := models.UpdateUserInfo(userUUID, request.GetUid(), request.GetFullName()); err != nil {
		return nil, status.Errorf(codes.Aborted, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *MainService) UpsertTakenCourses(context context.Context, request *pb.UpsertTakenCoursesRequest) (*emptypb.Empty, error) {
	log.Printf("Received UpsertTakenCoursesRequest")
	courseCodes := request.GetTakenCourseCodes()
	if courseCodes == nil || len(courseCodes) == 0 {
		log.Printf("empty course codes")
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}
	jwtClaims := context.Value(constants.JWTClaimsKey).(*jwt.UserClaims)
	userUUID := jwtClaims.Subject
	if userUUID == "" {
		log.Printf("empty uuid or error")
		return &emptypb.Empty{}, nil
	}

	log.Printf("current user id: %v", userUUID)

	if err := models.UpsertCourseCodes(userUUID, courseCodes); err != nil {
		log.Printf("error during upsert course codes: %v", err)
		return nil, status.Errorf(codes.Aborted, err.Error())
	}

	log.Printf("no failure during txn, good result")

	return &emptypb.Empty{}, nil
}

func (s *MainService) ListCourses(_ context.Context, request *pb.ListCoursesRequest) (*pb.CoursesResponse, error) {
	log.Printf("Received ListCoursesRequest")
	pageSize := request.GetPageSize()
	if pageSize < 0 {
		log.Printf("invalid page size")
		return nil, status.Errorf(codes.InvalidArgument, "invalid page size")
	} else if pageSize == 0 {
		// Set to default page size
		pageSize = 20
	}
	// Add one more page size to check if there are more courses
	newPageSize := pageSize + 1
	lastCourseCode := request.GetLastCode()
	courses, err := models.GetCourses(newPageSize, lastCourseCode)
	if err != nil {
		log.Printf("error during get courses: %v", err)
		return nil, status.Errorf(codes.Aborted, err.Error())
	}
	if len(courses) == 0 {
		return &pb.CoursesResponse{
			Courses:     nil,
			MoreResults: false,
		}, nil
	}

	moreResults := false
	var outputLength int
	if len(courses) != int(newPageSize) {
		outputLength = len(courses)
	} else {
		outputLength = len(courses) - 1
		moreResults = true
	}

	data := make([]*pb.Course, outputLength)
	for i := range data {
		if courses[i].Description == nil {
			courses[i].Description = new(string)
		}
		data[i] = &pb.Course{
			CourseCode:  courses[i].CourseCode,
			Title:       courses[i].Title,
			Department:  courses[i].Department,
			Description: *courses[i].Description,
			Rating:      courses[i].Rating,
			Offered:     courses[i].Offered,
		}
	}

	return &pb.CoursesResponse{
		Courses:     data,
		MoreResults: moreResults,
	}, nil
}

func (s *MainService) SearchCourses(_ context.Context, request *pb.SearchCourseRequest) (*pb.CoursesResponse, error) {
	log.Printf("Received ListCoursesRequest")
	if request.GetQuery() == "" {
		log.Printf("empty query")
		return nil, status.Errorf(codes.InvalidArgument, "empty query")
	}
	pageSize := request.GetPageSize()
	if pageSize < 0 {
		log.Printf("invalid page size")
		return nil, status.Errorf(codes.InvalidArgument, "invalid page size")
	}
	newPageSize := pageSize + 1

	lastCourseCode := request.GetLastCode()
	query := request.GetQuery()
	courses, err := models.SearchCourses(query, newPageSize, lastCourseCode)
	if err != nil {
		log.Printf("error during get courses: %v", err)
		return nil, status.Errorf(codes.Aborted, err.Error())
	}

	var outputLength int
	var moreResults bool
	if len(courses) != int(newPageSize) {
		outputLength = len(courses)
		moreResults = false
	} else {
		outputLength = len(courses) - 1
		moreResults = true
	}
	data := make([]*pb.Course, outputLength)
	for i := range data {
		if courses[i].Description == nil {
			courses[i].Description = new(string)
		}
		data[i] = &pb.Course{
			CourseCode:  courses[i].CourseCode,
			Title:       courses[i].Title,
			Department:  courses[i].Department,
			Description: *courses[i].Description,
			Rating:      courses[i].Rating,
			Offered:     courses[i].Offered,
		}
	}

	return &pb.CoursesResponse{
		Courses:     data,
		MoreResults: moreResults,
	}, nil
}

func (s *MainService) GetCourseDetails(_ context.Context, request *pb.GetCourseDetailRequest) (*pb.GetCourseDetailResponse, error) {
	courseCode := request.GetCourseCode()
	if courseCode == "" {
		return nil, status.Errorf(codes.InvalidArgument, "cannot get course details with empty course code")
	}
	course, err := models.GetCourseByCode(courseCode)
	if errors.Is(err, pgx.ErrNoRows) {
		return &pb.GetCourseDetailResponse{}, nil
	}
	if err != nil {
		return nil, status.Errorf(codes.Aborted, err.Error())
	}
	if course.Description == nil {
		course.Description = new(string)
	}

	courseReviews, err := models.GetAllCourseReviews(courseCode)
	if err != nil {
		return nil, status.Errorf(codes.Aborted, err.Error())
	}

	return &pb.GetCourseDetailResponse{
		Course: &pb.Course{
			CourseCode:  course.CourseCode,
			Title:       course.Title,
			Department:  course.Department,
			Description: *course.Description,
			Rating:      course.Rating,
			Offered:     course.Offered,
		},
		Reviews: courseReviews,
	}, nil
}

func (s *MainService) AddReview(context context.Context, request *pb.AddReviewRequest) (*pb.AddReviewResponse, error) {
	jwtClaims := context.Value(constants.JWTClaimsKey).(*jwt.UserClaims)
	userUUID := jwtClaims.Subject
	if userUUID == "" {
		log.Printf("empty uuid or error")
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}
	courseCode := request.GetCourseCode()
	if courseCode == "" || len(courseCode) > 10 {
		errorMessage := "Invalid course code"
		return &pb.AddReviewResponse{
			Result:       pb.AddReviewResult_INVALID_COURSE_CODE,
			ErrorMessage: &errorMessage,
		}, nil
	}
	rating := request.GetRating()
	if rating < 0 || rating > 5 {
		errorMessage := "Invalid rating"
		return &pb.AddReviewResponse{
			Result:       pb.AddReviewResult_INVALID_RATING,
			ErrorMessage: &errorMessage,
		}, nil
	}
	lecturers := request.GetLecturers()
	if len(lecturers) < 1 {
		errorMessage := "Please provide at least one lecturer"
		return &pb.AddReviewResponse{
			Result:       pb.AddReviewResult_INVALID_LECTURERS,
			ErrorMessage: &errorMessage,
		}, nil
	}
	yearTaken := request.GetYearTaken()
	if yearTaken > pb.AcademicYear_AY_2023_2024 {
		errorMessage := "Cannot add review for future year"
		return &pb.AddReviewResponse{
			Result:       pb.AddReviewResult_INVALID_YEAR_TAKEN,
			ErrorMessage: &errorMessage,
		}, nil
	}
	content := request.GetContent()
	if len(content) < 10 {
		errorMessage := "Review is too short"
		return &pb.AddReviewResponse{
			Result:       pb.AddReviewResult_INVALID_CONTENT,
			ErrorMessage: &errorMessage,
		}, nil
	}
	err := models.AddCourseReview(userUUID, request)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Message {
			case "review_published_already":
				errorMessage := "Review already published"
				return &pb.AddReviewResponse{
					Result:       pb.AddReviewResult_ERROR_ALREADY_REVIEWED,
					ErrorMessage: &errorMessage,
				}, nil
			case "course_not_taken":
				errorMessage := "User has not taken this course"
				return &pb.AddReviewResponse{
					Result:       pb.AddReviewResult_ERROR_USER_NOT_TAKEN_COURSE,
					ErrorMessage: &errorMessage,
				}, nil
			}
		}
		return nil, status.Errorf(codes.Aborted, err.Error())
	}

	return &pb.AddReviewResponse{
		Result: pb.AddReviewResult_SUCCESS,
	}, nil
}

func NewMainService(jwtManager *jwt.Manager) *MainService {
	return &MainService{jwtManager: jwtManager}
}

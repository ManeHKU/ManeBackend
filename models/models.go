package models

import (
	"ManeBackend/internal/env"
	"ManeBackend/models/types"
	"ManeBackend/pb"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
)

var yearTakenMapToPG = map[pb.AcademicYear]string{
	pb.AcademicYear_AY_2018_2019: "2018-2019",
	pb.AcademicYear_AY_2019_2020: "2019-2020",
	pb.AcademicYear_AY_2020_2021: "2020-2021",
	pb.AcademicYear_AY_2021_2022: "2021-2022",
	pb.AcademicYear_AY_2022_2023: "2022-2023",
	pb.AcademicYear_AY_2023_2024: "2023-2024",
}

var yearTakenMapToProto = map[string]pb.AcademicYear{
	"2018-2019": pb.AcademicYear_AY_2018_2019,
	"2019-2020": pb.AcademicYear_AY_2019_2020,
	"2020-2021": pb.AcademicYear_AY_2020_2021,
	"2021-2022": pb.AcademicYear_AY_2021_2022,
	"2022-2023": pb.AcademicYear_AY_2022_2023,
	"2023-2024": pb.AcademicYear_AY_2023_2024,
}

var dbPool *pgxpool.Pool

type UserInfo struct {
	UUID     string
	FullName string
}

func InitDB() error {
	config, err := pgxpool.ParseConfig(env.GetConfig().DATABASE_URI)
	if err != nil {
		return err
	}
	//certificate := tls.Certificate{Certificate: [][]byte{[]byte(env.DATABASE_CERT)}}
	//config.ConnConfig.TLSConfig = &tls.Config{
	//	Certificates: []tls.Certificate{certificate},
	//}
	dbPool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return err
	}

	return dbPool.Ping(context.Background())
}

func GetAllUsers() (string, error) {
	var name string
	err := dbPool.QueryRow(context.Background(), "SELECT nickname FROM profiles").Scan(&name)
	if err != nil {
		return "", err
	}
	return name, nil
}

func CheckUserExist(uuid string) bool {
	var exist bool
	err := dbPool.QueryRow(context.Background(), "SELECT EXISTS(SELECT id FROM profiles WHERE id = $1)", uuid).Scan(&exist)
	if err != nil || !exist {
		log.Printf("check user exist error: %v", err)
		return false
	}
	return true
}

func UpdateUserInfo(uuid string, uid uint32, fullName string) error {
	ctx := context.Background()
	tx, err := dbPool.Begin(ctx)
	if err != nil {
		return err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		rollBackErr := tx.Rollback(ctx)
		if rollBackErr != nil {
			err = rollBackErr
		}
	}(tx, ctx)

	userExist := CheckUserExist(uuid)
	if !userExist {
		return errors.New("user not exist")
	}

	if fullName != "" {
		_, _ = dbPool.Exec(ctx, "UPDATE profiles SET full_name = $1 WHERE id = $2", fullName, uuid)
	}
	if uid >= 3000000000 && uid <= 4000000000 {
		_, _ = dbPool.Exec(ctx, "UPDATE profiles SET uid = $1 WHERE id = $2", uid, uuid)
	}
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}

func UpsertCourseCodes(uuid string, courseCodes []string) error {
	log.Printf("start upsert course codes")
	ctx := context.Background()
	tx, err := dbPool.Begin(ctx)
	if err != nil {
		return err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		rollBackErr := tx.Rollback(ctx)
		if rollBackErr != nil {
			err = rollBackErr
		}
	}(tx, ctx)

	upsertErrors := make([]error, 0)
	for _, code := range courseCodes {
		_, err = dbPool.Exec(ctx, "INSERT into class_enrolment(user_id, course_code) select $1, $2 WHERE NOT EXISTS (SELECT 1 FROM class_enrolment WHERE user_id = $1 and course_code = $2);", uuid, code)
		if err != nil {
			log.Printf("error for %v: %v", code, err)
			upsertErrors = append(upsertErrors, err)
		}
	}
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	log.Printf("comitted")
	if len(upsertErrors) > 0 {
		errorText := fmt.Sprintf("upsert errors: %v", upsertErrors)
		return errors.New(errorText)
	}
	return nil
}

func GetCourses(pageSize int32, lastCourseCode string) ([]types.Course, error) {
	if pageSize <= 0 {
		return nil, errors.New("invalid page size")
	}
	var rows pgx.Rows
	var err error
	if lastCourseCode == "" {
		rows, err = dbPool.Query(context.Background(), "select course_code, title, department, description, rating, offered from courses_info order by course_code limit $1", pageSize)
		if err != nil {
			log.Printf("error: %v", err)
			return nil, err
		}
	} else {
		rows, err = dbPool.Query(context.Background(), "select course_code, title, department, description, rating, offered from courses_info where course_code > $1 order by course_code limit $2", lastCourseCode, pageSize)
		if err != nil {
			log.Printf("error: %v", err)
			return nil, err
		}
	}
	courses, err := pgx.CollectRows(rows, pgx.RowToStructByName[types.Course])
	if err != nil {
		log.Printf("error when collecting row: %v", err)
		return nil, err
	}

	return courses, nil
}

func SearchCourses(query string, pageSize int32, lastCourseCode string) ([]types.Course, error) {
	if pageSize < 0 {
		return nil, errors.New("invalid page size")
	}
	var rows pgx.Rows
	var err error
	if pageSize == 0 && lastCourseCode == "" {
		rows, err = dbPool.Query(context.Background(), "select course_code, title, department, description, rating, offered from courses_info where course_code ilike '%' || $1 || '%' or description ilike '%' || $1 || '%' or title ilike '%' || $1 || '%' order by course_code", query)
	} else if pageSize != 0 && lastCourseCode != "" {
		rows, err = dbPool.Query(context.Background(), "select course_code, title, department, description, rating, offered from courses_info where (course_code ilike '%' || $1 || '%' or description ilike '%' || $1 || '%' or title ilike '%' || $1 || '%') and course_code > $2 order by course_code limit $3", query, lastCourseCode, pageSize)
	} else if pageSize != 0 {
		rows, err = dbPool.Query(context.Background(), "select course_code, title, department, description, rating, offered from courses_info where (course_code ilike '%' || $1 || '%' or description ilike '%' || $1 || '%' or title ilike '%' || $1 || '%') order by course_code limit $2", query, pageSize)
	}
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}
	courses, err := pgx.CollectRows(rows, pgx.RowToStructByName[types.Course])
	if err != nil {
		log.Printf("error when collecting row: %v", err)
		return nil, err
	}
	return courses, nil
}

func GetCourseByCode(courseCode string) (*types.Course, error) {
	row, err := dbPool.Query(context.Background(), "select course_code, title, department, description, rating, offered from courses_info where course_code = $1", courseCode)
	course, err := pgx.CollectOneRow(row, pgx.RowToStructByName[types.Course])
	if err != nil {
		log.Printf("error when collecting row: %v", err)
		return nil, err
	}
	return &course, nil
}

func AddCourseReview(uuid string, request *pb.AddReviewRequest) error {
	userExist := CheckUserExist(uuid)
	if !userExist {
		return errors.New("user not exist")
	}

	ctx := context.Background()
	tx, err := dbPool.Begin(ctx)
	if err != nil {
		return err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		rollBackErr := tx.Rollback(ctx)
		if rollBackErr != nil {
			err = rollBackErr
		}
	}(tx, ctx)

	yearTaken, ok := yearTakenMapToPG[request.GetYearTaken()]
	if !ok {
		return errors.New("unsupported taken year")
	}
	result, err := dbPool.Exec(ctx, "INSERT INTO course_reviews(author, course_code, lecturers, year_taken, semester_taken, content,rating) VALUES($1, $2, $3, $4, $5, $6, $7)", uuid, request.CourseCode, request.Lecturers, yearTaken, request.SemesterTaken.String(), request.Content, request.Rating)
	if err != nil {
		return err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	if result.Insert() && result.RowsAffected() == 1 {
		return nil
	}
	return errors.New("unknown error")
}

func GetAllCourseReviews(courseCode string) ([]*pb.Review, error) {
	rows, err := dbPool.Query(context.Background(), "select year_taken, semester_taken, content, rating, lecturers, created_at from course_reviews where course_code = $1", courseCode)
	if err != nil {
		return nil, err
	}

	var (
		yearTaken, semesterTaken, content string
		lecturers                         []string
		rating                            int
		createdAt                         pgtype.Timestamptz
	)

	reviews := make([]*pb.Review, 0)
	_, err = pgx.ForEachRow(rows, []any{&yearTaken, &semesterTaken, &content, &rating, &lecturers, &createdAt}, func() error {
		reviews = append(reviews, &pb.Review{
			CourseCode:    courseCode,
			YearTaken:     yearTakenMapToProto[yearTaken],
			SemesterTaken: pb.Semester(pb.Semester_value[semesterTaken]),
			Content:       content,
			Rating:        int32(rating),
			Lecturers:     lecturers,
			CreatedAt:     timestamppb.New(createdAt.Time),
		})
		return nil
	})
	if err != nil {
		return nil, err
	}
	return reviews, nil
}

func CheckUserEnrolledCourse(uuid string, courseCode string) bool {
	var exist bool
	err := dbPool.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM class_enrolment WHERE user_id = $1 and course_code = $2)", uuid, courseCode).Scan(&exist)
	if err != nil || !exist {
		log.Printf("check user enrolled course error: %v", err)
		return false
	}
	return true
}

func CheckUserPublishedReview(uuid string, courseCode string) bool {
	var exist bool
	err := dbPool.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM course_reviews WHERE author = $1 and course_code = $2)", uuid, courseCode).Scan(&exist)
	if err != nil || !exist {
		log.Printf("check user published review error: %v", err)
		return false
	}
	return true
}

func CheckUserIsOrganizerAdmin(uuid string, organizerUUID string) bool {
	var exist bool
	err := dbPool.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM organizer_admins WHERE admin = $1 AND organizer = $2)", uuid, organizerUUID).Scan(&exist)
	if err != nil {
		log.Printf("check user is organizer admin error: %v", err)
		return false
	}
	return exist
}

func AddCampusEvent(uuid string, organizerUUID string, title string, startDate time.Time, endDate time.Time, location string, description string, participantLimit int32, imagePath string, applyInfo *pb.ApplyInfo) error {
	userIsAdmin := CheckUserIsOrganizerAdmin(uuid, organizerUUID)
	if !userIsAdmin {
		return errors.New("user is not admin")
	}

	ctx := context.Background()
	tx, err := dbPool.Begin(ctx)
	if err != nil {
		return err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		rollBackErr := tx.Rollback(ctx)
		if rollBackErr != nil {
			err = rollBackErr
		}
	}(tx, ctx)

	var imagePathPG pgtype.Text
	imagePathPG = pgtype.Text(NewNullString(imagePath))

	var eventId string
	err = dbPool.QueryRow(ctx, "INSERT INTO event_info(organizer, title, start_datetime, end_datetime, location, description, title_image_path, participant_limit) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", organizerUUID, title, startDate, endDate, location, description, imagePathPG, participantLimit).Scan(&eventId)
	if err != nil {
		return err
	}
	_, err = dbPool.Exec(ctx, "INSERT INTO event_apply_info(event_id, info, questions) VALUES($1, $2, $3)", eventId, applyInfo.Info, applyInfo.Questions)
	if err != nil {
		return err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func ListFutureCampusEvents(sortBy pb.SortBy) ([]*pb.ListLatestEventsResponse_FullEventInfo, error) {
	var sortByColumn string
	if sortBy == pb.SortBy_CREATED_AT {
		sortByColumn = "created_at"
	} else if sortBy == pb.SortBy_PARTICIPATION_TIME {
		sortByColumn = "start_datetime"
	}

	eventRows, err := dbPool.Query(context.Background(), "select event_id::text, organizer::text as organizer_id, title, title_image_path, start_datetime, end_datetime, location, event_info.description, status,  event_info.updated_at, organizer.name, organizer.description, event_with_participant_count.participant_limit, event_with_participant_count.participant_count from event_info left join organizer on event_info.organizer = organizer.id left join event_with_participant_count on event_with_participant_count.event_id = event_info.id where status = 'OPEN'::event_status and start_datetime > now() order by $1", sortByColumn)
	if err != nil {
		return nil, err
	}

	eventMaps, err := pgx.CollectRows(eventRows, pgx.RowToMap)
	if err != nil {
		return nil, err
	}

	result := make([]*pb.ListLatestEventsResponse_FullEventInfo, len(eventMaps))
	for i, event := range eventMaps {
		var imagePath *string
		if event["title_image_path"] != nil {
			text := event["title_image_path"].(string)
			imagePath = &text
		} else {
			imagePath = nil
		}
		eventInfo := &pb.EventInfo{
			Id:          event["event_id"].(string),
			OrganizerId: event["organizer_id"].(string),
			Title:       event["title"].(string),
			ImagePath:   imagePath,
			StartTime:   timestamppb.New(event["start_datetime"].(time.Time)),
			EndTime:     timestamppb.New(event["end_datetime"].(time.Time)),
			Location:    event["location"].(string),
			Description: event["description"].(string),
			Status:      pb.EventStatus(pb.EventStatus_value[event["status"].(string)]),
			UpdatedAt:   timestamppb.New(event["updated_at"].(time.Time)),
		}
		var organizationDescription *string
		if event["description"] != nil {
			text := event["description"].(string)
			organizationDescription = &text
		} else {
			organizationDescription = nil
		}
		organizer := &pb.OrganizerInfo{
			Id:          event["organizer_id"].(string),
			Name:        event["name"].(string),
			Description: organizationDescription,
		}
		participation := &pb.EventParticipation{
			Limit:        int32(event["participant_limit"].(int16)),
			CurrentCount: int32(event["participant_count"].(int64)),
		}
		result[i] = &pb.ListLatestEventsResponse_FullEventInfo{
			Event:         eventInfo,
			Organizer:     organizer,
			Participation: participation,
		}
	}

	return result, nil
}

func GetApplyInfo(eventId string, userUUID string) (applyInfo string, questions []string, userApplied bool, err error) {
	err = dbPool.QueryRow(context.Background(), "SELECT event_apply_info.info, event_apply_info.questions FROM event_apply_info WHERE event_id = $1", eventId).Scan(&applyInfo, &questions)
	if err != nil {
		return "", nil, false, err
	}

	err = dbPool.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM event_participants WHERE event_id = $1 AND participant_id = $2)", eventId, userUUID).Scan(&userApplied)
	if err != nil {
		return "", nil, false, err
	}

	return applyInfo, questions, userApplied, nil
}

func CheckOrganizerExist(organizerUUID string) (exist bool) {
	err := dbPool.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM organizer WHERE id = $1)", organizerUUID).Scan(&exist)
	if err != nil {
		log.Printf("check organizer exist error: %v", err)
	}
	return exist
}

func ApplyCampusEvent(userUUID string, eventID string, answers map[string]string) error {
	ctx := context.Background()
	tx, err := dbPool.Begin(ctx)
	if err != nil {
		return err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		rollBackErr := tx.Rollback(ctx)
		if rollBackErr != nil {
			err = rollBackErr
		}
	}(tx, ctx)

	var userApplied bool
	err = dbPool.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM event_participants WHERE event_id = $1 AND participant_id = $2)", eventID, userUUID).Scan(&userApplied)
	if userApplied || err != nil {
		return errors.New("user already applied")
	}

	var isEventFull bool
	err = dbPool.QueryRow(ctx, "SELECT participant_count >= participant_limit FROM event_with_participant_count where event_id = $1", eventID).Scan(&isEventFull)
	if err != nil {
		return err
	} else if isEventFull {
		return errors.New("event is full")
	}

	var isEventOpen bool
	err = dbPool.QueryRow(ctx, "SELECT status = 'OPEN'::event_status FROM event_info where id = $1", eventID).Scan(&isEventOpen)
	if err != nil {
		return err
	}
	if !isEventOpen {
		return errors.New("event is not open")
	}

	answersJSON, _ := json.Marshal(answers)
	_, err = dbPool.Exec(ctx, "INSERT INTO event_participants(event_id, participant_id, answer) VALUES($1, $2, $3)", eventID, userUUID, answersJSON)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}

func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func Close() {
	defer dbPool.Close()
}

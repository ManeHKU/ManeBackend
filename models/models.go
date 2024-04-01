package models

import (
	"ManeBackend/internal/env"
	"ManeBackend/models/types"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

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

func Close() {
	defer dbPool.Close()
}

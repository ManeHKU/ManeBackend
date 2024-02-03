package models

import (
	"ManeBackend/internal/env"
	"context"
	"errors"
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

func Close() {
	defer dbPool.Close()
}

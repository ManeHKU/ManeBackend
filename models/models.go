package models

import (
	"ManeBackend/internal/env"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

var dbPool *pgxpool.Pool

func InitDB(env *env.Config) error {
	config, err := pgxpool.ParseConfig(env.DATABASE_URI)
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

func Close() {
	defer dbPool.Close()
}

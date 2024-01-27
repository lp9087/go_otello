package configs

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type DBConfig struct {
	dbName   string
	host     string
	port     string
	sslMode  string
	user     string
	password string
}

func NewDBConfig(dbName, host, port, sslMode, user, password string) *DBConfig {
	return &DBConfig{dbName: dbName, host: host, port: port, user: user, password: password, sslMode: sslMode}
}

func GetDBConnect(conf *DBConfig) (*sqlx.DB, error) {
	dbString := fmt.Sprintf("host=%s port=%s sslmode=%s dbname=%s user=%s password=%s",
		conf.host,
		conf.port,
		conf.sslMode,
		conf.dbName,
		conf.user,
		conf.password,
	)
	db, err := sqlx.Connect("postgres", dbString)
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	return db, nil
}

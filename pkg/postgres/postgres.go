package postgres

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/lp9087/go_otello_dashboard_api/config"
	"log"
)

type Postgres struct {
	Builder squirrel.StatementBuilderType
	Connect *sqlx.DB
}

func (p *Postgres) Close() {
	if p.Connect != nil {
		p.Connect.Close()
	}
}

func New(dbConf *config.DB) (*Postgres, error) {
	var postgres Postgres
	dbString := fmt.Sprintf("host=%s port=%s sslmode=%s dbname=%s user=%s password=%s",
		dbConf.Host,
		dbConf.Port,
		dbConf.SslMode,
		dbConf.DbName,
		dbConf.User,
		dbConf.Password,
	)
	connect, err := sqlx.Connect("postgres", dbString)
	if err != nil {
		log.Fatalln(err)
	}
	err = connect.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	postgres.Connect = connect
	postgres.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	return &postgres, nil
}

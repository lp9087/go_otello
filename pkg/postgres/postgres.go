package postgres

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
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

func New(dbString string) (*Postgres, error) {
	var postgres Postgres
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

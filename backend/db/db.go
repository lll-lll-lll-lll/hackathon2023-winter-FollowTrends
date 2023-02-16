package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgreSql struct {
	Db *sqlx.DB
}

func NewPostgreSql() *PostgreSql {
	return &PostgreSql{}
}

func (ps *PostgreSql) Init() {
	log.Println("Postgresql 初期化スタート")
	Db, err := sqlx.Open("postgres", "host=postgres user=app_user password=password dbname=app_db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Postgresql　初期化成功")
	ps.Db = Db
}

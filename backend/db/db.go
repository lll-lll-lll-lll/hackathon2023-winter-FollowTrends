package db

import (
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgreSql struct {
	Db *sqlx.DB
}

func NewPostgreSql() *PostgreSql {
	return &PostgreSql{}
}

type Data struct {
	KeyWord string `db:"keyword" json:"keyword"`
	URL     string `db:"url" json:"url"`
}

// SplitedKeyWord 文字列のKeyWordを「,」で分割するメソッド
func (d Data) SplitedKeyWord() []string {
	splited := strings.Split(d.KeyWord, ",")
	return splited
}

func SplitKeyWord(data []Data) [][]string {
	var datas [][]string
	for _, d := range data {
		datas = append(datas, d.SplitedKeyWord())
	}
	return datas
}

// GetData keywordとurlがセットになった構造体のスライスを返す
func (ps *PostgreSql) GetData() ([]Data, error) {
	rows, err := ps.Db.Queryx("SELECT keyword, url FROM test_data")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]Data, 0)
	for rows.Next() {

		var data Data

		err := rows.StructScan(&data)

		if err != nil {
			log.Fatal(err)
		}

		results = append(results, data)
	}
	return results, nil
}

func (ps *PostgreSql) Init() {
	log.Println("Postgresql 初期化スタート")
	Db, err := sqlx.Open("postgres", "host=db user=app_user password=password dbname=app_db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Postgresql　初期化成功")
	ps.Db = Db
}

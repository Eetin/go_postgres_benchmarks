package MyDB

import (
	"database/sql"

	"log"

	"fmt"

	"github.com/Eetin/go_postgres_benchmarks/MyRPC"

	_ "github.com/lib/pq"
)

type DB struct {
	DBHandler *sql.DB
	dburl     string
}

func New(dburl string) (*DB, error) {
	dbh, err := sql.Open("postgres", dburl)
	if err != nil {
		return nil, err
	}
	if err := dbh.Ping(); err != nil {
		return nil, err
	}
	return &DB{dbh, dburl}, nil
}

func (db *DB) InsertProtoData(index int, data *MyRPC.SimpleData) {
	dbdata, err := data.Marshal()
	if err != nil {
		log.Fatal("Marshal err2: ", err)
	}
	query := fmt.Sprintf("INSERT INTO test (index, data) VALUES (%d, $1)", index)
	_, err2 := db.DBHandler.Exec(query, dbdata)
	if err2 != nil {
		log.Fatal("DB err2: ", err2)
	}
}

func (db *DB) GetProtoData(index int, mydata *MyRPC.SimpleData) {
	rows, err := db.DBHandler.Query(fmt.Sprintf("SELECT data FROM test WHERE index=%d", index))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	data := make([]byte, 0)
	for rows.Next() {
		err := rows.Scan(&data)
		if err != nil {
			log.Fatal(err)
		}
	}
	mydata.Unmarshal(data)
}

func (db *DB) GetProtoDataRange(from int, to int) {
	rows, err := db.DBHandler.Query(fmt.Sprintf("SELECT data FROM test"))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	element:=&MyRPC.SimpleData{}
	buf:=make([]byte,element.Size())
	for rows.Next() {
		err := rows.Scan(&buf)
		if err != nil {
			log.Fatal("Simple get: ", err)
		}
		err=element.Unmarshal(buf)
		if err != nil {
			log.Fatal("Simple get: ", err)
		}

	}
}

func (db *DB) InsertData(index int, data *MyRPC.SimpleData) {
	query := fmt.Sprintf("INSERT INTO test2 (index, str1, str2, i64, i32) VALUES (%d, $1, $2, $3, $4)", index)
	_, err2 := db.DBHandler.Exec(query, data.Str1, data.Str2, data.I64, data.I32)
	if err2 != nil {
		log.Fatal("Simple DB err2: ", err2)
	}
}

func (db *DB) GetData(index int, mydata *MyRPC.SimpleData) {
	rows, err := db.DBHandler.Query(fmt.Sprintf("SELECT str1, str2, i64, i32 FROM test2 WHERE index=%d", index))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&mydata.Str1, &mydata.Str2, &mydata.I64, &mydata.I32)
		if err != nil {
			log.Fatal("Simple get: ", err)
		}
	}
}

func (db *DB) GetDataRange(from int, to int) {
	rows, err := db.DBHandler.Query(fmt.Sprintf("SELECT str1, str2, i64, i32 FROM test2"))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	element:=&MyRPC.SimpleData{}

	for rows.Next() {
		err := rows.Scan(&element.Str1, &element.Str2, &element.I64, &element.I32)
		if err != nil {
			log.Fatal("Simple get: ", err)
		}
	}
}

package main

import (
	"fmt"
	"log"
	"github.com/Eetin/go_postgres_benchmarks/MyDB"
	"github.com/Eetin/go_postgres_benchmarks/MyRPC"
)

const USER = "postgres"
const PASS = "postgrespass"
const ADDR = "localhost:5432"
const DB = "postgres"

func main() {

	dburl := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", USER, PASS, ADDR, DB)
	db, err := MyDB.New(dburl)
	if err != nil {
		log.Fatal(err)
	}
	//_, err2 := db.DBHandler.Query("TRUNCATE test, test2")
	//if err2 != nil {
	//	log.Fatal("Truncate err2: ", err2)
	//}
	//simpleData := MyRPC.SimpleData{
	//	"Hello",
	//	"World",
	//	123,
	//	321,
	//}
	//db.InsertProtoData(0, &simpleData)
	//db.GetProtoData(0, &MyRPC.SimpleData{})
	//
	//db.InsertData(0, &simpleData)
	//db.GetData(0, &MyRPC.SimpleData{})
	var dataRange = make([]*MyRPC.SimpleData, 10)
	db.GetDataRange(100, 110, dataRange)
	for i := range dataRange {
		log.Println("data ", i, ": ", dataRange[i])
	}

	var protoDataRange = make([]*MyRPC.SimpleData, 10)
	db.GetProtoDataRange(100, 110, protoDataRange)
	for i := range protoDataRange {
		log.Println("protodata ", i, ": ", protoDataRange[i])
	}
}

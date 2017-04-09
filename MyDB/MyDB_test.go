package MyDB

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"github.com/Eetin/go_postgres_benchmarks/MyRPC"
	"testing"
	"time"
)

const USER = "postgres"
const PASS = "postgrespass"
const ADDR = "localhost:5432"
const DBNAME = "postgres"

var database *DB

func setup() {
	rand.Seed(time.Now().UnixNano())
	connect()
	//clearDB()
	//fillDB(100000)
}

func connect() {
	dburl := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", USER, PASS, ADDR, DBNAME)
	db, err := New(dburl)
	database = db
	if err != nil {
		log.Fatal(err)
	}
}

func clearDB() {
	_, err2 := database.DBHandler.Query("TRUNCATE test, test2")
	if err2 != nil {
		log.Fatal("Truncate err2: ", err2)
	}
}

func fillDB(n int) {
	for i := 0; i < n; i++ {
		data := generateSimpleData()
		database.InsertData(i, data)
		database.InsertProtoData(i, data)
	}
}

func generateRandomString(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func generateSimpleData() *MyRPC.SimpleData {
	data := &MyRPC.SimpleData{
		generateRandomString(100),
		generateRandomString(200),
		rand.Int63(),
		rand.Int31(),
	}
	return data
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

//func BenchmarkDB_GetData(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		database.GetData(0, &MyRPC.SimpleData{})
//	}
//}
//
//func BenchmarkDB_GetProtoData(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		database.GetProtoData(0, &MyRPC.SimpleData{})
//	}
//}


func BenchmarkDB_GetProtoDataRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		database.GetProtoDataRange(4000, 70000)
	}
}

func BenchmarkDB_GetDataRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		database.GetDataRange(4000, 70000)
	}
}

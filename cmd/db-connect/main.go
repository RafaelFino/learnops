package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/lib/pq"
)

//Config base connection config struct
type DbConfig struct {
	Host   string
	User   string
	Pass   string
	DBName string
}

func main() {
	if len(os.Args) < 2 {
		panic(fmt.Errorf("You must set db-config file path as argument"))
	}

	config, err := readConfig(os.Args[1])
	if err != nil {
		log.Fatal("fail to try read db-config")
		panic(err)
	}

	connString := createConnectionString(config)

	db := connect(connString)

	err = testConnection(db)

	if err != nil {
		log.Fatal("db-test fail")
		panic(err)
	}

	log.Printf("Connection Ok!")
}

func readConfig(filepath string) (*DbConfig, error) {
	var config DbConfig

	data, err := ioutil.ReadFile(filepath)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &config)

	if err != nil {
		log.Fatal(err)
	}

	return &config, err
}

func connect(connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func createConnectionString(config *DbConfig) string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable",
		config.User,
		config.Pass,
		config.DBName,
		config.Host,
	)
}

func testConnection(db *sql.DB) error {
	return db.Ping()
}

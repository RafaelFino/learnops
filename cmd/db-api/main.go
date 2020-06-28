package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/rafaelfino/learnops/pkg/banks"
)

//Config base connection config struct
type DbConfig struct {
	Host   string
	User   string
	Pass   string
	DBName string
}

var db *sql.DB

//Main
func main() {
	if len(os.Args) < 2 {
		panic(fmt.Errorf("You must set db-config file path as argument"))
	}

	f, err := os.OpenFile(`/var/log/simple-server`, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
		panic(err)
	}
	defer f.Close()

	log.SetOutput(f)

	config, err := readConfig(os.Args[1])
	if err != nil {
		log.Fatal("fail to try read db-config")
		panic(err)
	}

	connString := createConnectionString(config)

	db = connect(connString)

	err = testConnection(db)

	if err != nil {
		log.Fatal("db-test fail")
		panic(err)
	}

	log.Printf("Connection Ok!")

	http.HandleFunc("/banks", handle)

	log.Fatal(http.ListenAndServe(os.Args[2], nil))
}

//Handlers
func handle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	path := getPath(r.RequestURI)

	loader := banks.New(db)

	ret, err := loader.Load()

	if err != nil {
		log.Printf("fail to try load banks data: %s", err)
	}

	raw, err := json.MarshallIdent(ret, "", "\t")

	if err != nil {
		log.Printf("fail to try load banks data: %s", err)
	}

	fmt.Fprintln(w, string(raw))
	log.Printf(`[%s] request returned in %s`, path, time.Since(start))
}

//Internals
func getPath(path string) string {
	idx := strings.Index(path, `?`)

	if idx > 0 {
		return string(path[0:idx])
	}

	return path
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

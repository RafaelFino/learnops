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

	_ "github.com/lib/pq"
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

	f, err := os.OpenFile(`/var/log/db-api`, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
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

	connect(config)

	http.HandleFunc("/banks", handle)
	http.HandleFunc("/banks-html", handleHTML)

	log.Printf("server started")

	log.Fatal(http.ListenAndServe(os.Args[2], nil))
}

//Handlers
func handle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	path := getPath(r.RequestURI)

	ret, err := LoadBanks()

	if err != nil {
		log.Printf("fail to try load banks data: %s", err)
	}

	raw, err := json.MarshalIndent(ret, "", "\t")

	if err != nil {
		log.Printf("fail to try load banks data: %s", err)
	}

	fmt.Fprintln(w, string(raw))
	log.Printf(`[%s] request returned in %s`, path, time.Since(start))
}

func handleHTML(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	path := getPath(r.RequestURI)

	ret, err := LoadBanks()

	if err != nil {
		log.Printf("fail to try load banks data: %s", err)
	}

	var list string

	for _, i := range ret {
		list += fmt.Sprintf(`
<tr>
	<td>%d</td>
	<td>%s</td>
	<td>%s</td>
</tr>
`,
			i.ID,
			i.Name,
			i.Fullname,
		)
	}

	fmt.Fprintln(w, strings.ReplaceAll(htmlBanks, "##LIST##", list))
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

func connect(config *DbConfig) {
	connString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable",
		config.User,
		config.Pass,
		config.DBName,
		config.Host,
	)

	var err error
	db, err = sql.Open("postgres", connString)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("db-connect fail")
		panic(err)
	}

	log.Printf("connected on %s", config.Host)
}

//Banks
type Bank struct {
	ID       int
	Name     string
	Fullname string
}

const selectBanks = `
SELECT
	BankID,
	BankName,
	BankFullName
FROM
	Banks
ORDER BY
	BankID
`

func LoadBanks() ([]*Bank, error) {
	ret := []*Bank{}

	rows, err := db.Query(selectBanks)

	if err == nil {
		defer rows.Close()

		for rows.Next() {
			bank := &Bank{}
			if err = rows.Scan(
				&bank.ID,
				&bank.Name,
				&bank.Fullname,
			); err == nil {
				ret = append(ret, bank)
			} else {
				return ret, err
			}
		}
	}

	return ret, err
}

const htmlBanks = `
<!DOCTYPE html>
<html>
<head>
<style>
table, th, td {
  border: 1px solid black;
}
</style>
</head>
<body>

<h2>Lista de bancos</h2>

<table style="width:100%">
  <tr>
    <th>ID</th>
    <th>Nome</th> 
    <th>Nome completo</th>
  </tr>
  ##LIST##
</table>

</body>
</html>
`

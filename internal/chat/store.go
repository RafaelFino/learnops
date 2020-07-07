package chat

import (
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(filename string) (*Storage, error) {
	db, err := sql.Open("sqlite3", filename)

	if err != nil {
		return nil, err
	}

	ret := &Storage{db: db}

	err = ret.createTable()

	return ret, err
}

func (s *Storage) createTable() error {
	cmd := `
	CREATE TABLE IF NOT EXISTS messages (		
		subject TEXT DEFAUL NULL,
		nickname TEXT DEFAUL NULL,
		messageId INT DEFAULT 0,
		payload TEXT DEFAULT NULL,
		timestamp TIMESTAMP	DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (subject, nickname, messageId)
	)
	`
	_, err := s.db.Exec(cmd)

	return err
}

func (s *Storage) Write(msg *Message) error {
	cmd := `
	INSERT OR REPLACE INTO 
		messages
	(
		subject,
		nickname,
		messageId,
		payload
	) 
	VALUES
	(
		?,
		?,
		?,
		?
	)
	`

	j, err := json.MarshalIndent(msg, "", "\t")

	_, err = s.db.Exec(cmd, msg.Subject, msg.From, msg.ID, string(j))

	if err != nil {
		log.Fatalf("fail to write message on db, msg: %s error: $s", cmd, err)
	}

	return err
}

func (s *Storage) Read(subject string, nickname string) ([]*Message, error) {
	query := `
	SELECT 
		payload
	FROM
		messages 
	WHERE
		subject = ?
	ORDER BY
		messageId
	LIMIT 200		
	`

	rows, err := s.db.Query(query, subject, nickname)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var ret []*Message

	for rows.Next() && rows.Err() == nil {
		var raw string
		if err = rows.Scan(&raw); err == nil {
			msg := Message{}
			if err = json.Unmarshal([]byte(raw), &msg); err == nil {
				ret = append(ret, &msg)
			}
		} else {
			log.Fatalf("error to try read messages from db. read: %s, error: %s", raw, err)
			return ret, err
		}
	}

	if rows.Err() != nil {
		log.Fatal(err)
	}

	log.Printf("%d newest messages are restored from db\n", len(ret))

	return ret, err
}

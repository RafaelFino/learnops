package banks

import (
	"database/sql"
)

type Banks struct {
	db *sql.DB
}

type Bank struct {
	ID       int
	Name     string
	Fullname string
}

func New(db *sql.DB) *Banks {
	return &Banks{db: db}
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

func (b *Banks) Load() ([]*Bank, error) {
	ret := []*Bank{}

	rows, err := b.db.Query(selectBanks)

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

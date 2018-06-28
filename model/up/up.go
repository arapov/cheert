package up

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

var (
	table = "cheers"
)

type Item struct {
	ID        uint32         `db:"ud"`
	From      string         `db:"from_id"`
	To        string         `db:"to_id"`
	Plus      string         `db:"plus"`
	Note      string         `db:"note"`
	CreatedAt mysql.NullTime `db:"created_at"`
	UpdatedAt mysql.NullTime `db:"updated_at"`
	DeletedAt mysql.NullTime `db:"deleted_at"`
}

type Conection interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

func Create(db Conection, From string, To string, Plus string, Note string) (sql.Result, error) {
	result, err := db.Exec(fmt.Sprintf("INSERT INTO %v (`from`, `to`, `plus`, `note`) VALUES (?,?,?,?)", table),
		From, To, Plus, Note)
	return result, err
}

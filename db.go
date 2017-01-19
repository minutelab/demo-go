package demo

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // we use postgres, we need to iport the library for side effect
)

// ErrNothingDone is returned by DB operations that didn't do anything but there was no error
// Like using ExecOnce to delete a specific row that does not exist
var ErrNothingDone = errors.New("nothing done")

// DB is our main database object
var db *sqlx.DB

// Init initialize the library by supply database object
func Init(d *sql.DB) {
	db = sqlx.NewDb(d, "postgres")
}

// execOnce is use to run a query that is supposed to modify exactly one record (like add)
// it return nil if the query did it, otherwise it return an error
func execOnce(query string, args ...interface{}) error {
	res, err := db.Exec(query, args...)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	switch affected {
	case 1:
		return nil

	case 0:
		return ErrNothingDone

	default:
		return fmt.Errorf("too many lines changed: %d", affected)
	}
}

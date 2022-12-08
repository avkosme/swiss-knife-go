package pg

import (
	"database/sql"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
)

// layout := "2006-01-02T15:04:05-0700"
// createdAt := time.Now().Format(layout)
// updatedAt := &createdAt
// Update
func Update(db *sql.DB, q string, err error, args ...any) (sql.Result, error) {
	erl("", err)
	result, err := db.Exec(
		q,
		args...,
	)
	erl("", err)
	return result, err
}

// erl
func erl(mes string, err error) {
	if err != nil {
		defer sentry.Flush(2 * time.Second)
		sentry.CaptureException(err)
		log.Printf("%q %q", mes, err)
	}
}

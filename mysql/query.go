package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
)

// Finder
type Finder interface {
	FindByQuery(db *sql.DB, q string, limit, offset int, err error) []map[string]any
}

type Creater interface {
	Create(db *sql.DB, q string, err error, args ...any)
}

// FindIn
// var in string
// query := fmt.Sprintf("SELECT id,name from user WHERE phone IN (%s)", in)
func FindIn(db *sql.DB, q string, err error, args ...any) (result []map[string]any) {
	defer func() {
		if r := recover(); r != nil {
			erl("", err)
			defer sentry.Flush(2 * time.Second)
			sentry.CaptureException(err)
		}
	}()
	rows, err := db.Query(
		q,
		args...,
	)

	erl("", err)

	defer rows.Close()

	columns, err := rows.Columns()
	erl("", err)
	row := make([][]byte, len(columns))
	rowPtr := make([]any, len(columns))
	for i := range row {
		rowPtr[i] = &row[i]
	}

	for rows.Next() {
		_ = rows.Scan(rowPtr...)
		rd := map[string]any{}
		for k, v := range row {
			rd[columns[k]] = string(v)
		}
		result = append(result, rd)
	}
	return
}

// layout := "2006-01-02 15:04:05"
// createdAt := time.Now().Format(layout)
// updatedAt := &createdAt
// Create
func Create(db *sql.DB, q string, err error, args ...any) {
	defer func() {
		if r := recover(); r != nil {
			erl("", err)
			defer sentry.Flush(2 * time.Second)
			sentry.CaptureException(err)
		}
	}()
	defer db.Close()
	_, err = db.Exec(
		q,
		args...,
	)
	erl("", err)
}

// FindByQuery
func FindByQuery(db *sql.DB, q string, limit, offset int, err error) (result []map[string]any) {
	if limit > 0 {
		q = fmt.Sprintf("%s LIMIT %d OFFSET %d", q, limit, offset)
	} else {
		q = fmt.Sprintf("%s OFFSET %d", q, offset)
	}
	rows, err := db.Query(q)

	defer func() {
		if r := recover(); r != nil {
			defer sentry.Flush(2 * time.Second)
			sentry.CaptureException(err)
			log.Println(err)
		}
	}()
	defer rows.Close()

	columns, err := rows.Columns()
	erl("", err)
	row := make([][]byte, len(columns))
	rowPtr := make([]any, len(columns))
	for i := range row {
		rowPtr[i] = &row[i]
	}

	for rows.Next() {
		_ = rows.Scan(rowPtr...)
		rd := map[string]any{}
		for k, v := range row {
			rd[columns[k]] = string(v)
		}
		result = append(result, rd)
	}

	return result
}

package mysql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/getsentry/sentry-go"
)

// Finder
type Finder interface {
	FindByQuery(db *sql.DB, q string, limit, offset int, err error) []map[string]any
}

// FindByQuery
func FindByQuery(db *sql.DB, q string, limit, offset int, err error) (result []map[string]any) {
	if limit > 0 {
		q = fmt.Sprintf("%s LIMIT %d OFFSET %d", q, limit, offset)
	} else {
		q = fmt.Sprintf("%s OFFSET %d", q, offset)
	}
	rows, err := db.Query(q)
	erl("", err)

	defer func() {
		if r := recover(); r != nil {
			defer sentry.Flush(2 * time.Second)
			sentry.CaptureException(err)
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

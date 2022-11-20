package mysql

import (
	"database/sql"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/go-sql-driver/mysql"
)

type Mysql struct {
	user                 string
	passwd               string
	net                  string
	addr                 string
	dbname               string
	allownativepasswords bool
	cfg                  mysql.Config
	Db                   *sql.DB
}

// NewMysql
func NewMysql(user, passwd, addr, dbname string) *Mysql {
	return &Mysql{
		user: user, passwd: passwd,
		addr: addr, dbname: dbname,
		net: "tcp", allownativepasswords: true,
	}
}

// Connect
func (m *Mysql) Connect() {
	m.cfg = mysql.Config{
		User:                 m.user,
		Passwd:               m.passwd,
		Net:                  m.net,
		Addr:                 m.addr,
		DBName:               m.dbname,
		AllowNativePasswords: m.allownativepasswords,
	}
	db, err := sql.Open("mysql", m.cfg.FormatDSN())
	if err != nil {
		erl("", err)
	}
	m.Db = db
}

// erl
func erl(mes string, err error) {
	if err != nil {
		defer sentry.Flush(2 * time.Second)
		sentry.CaptureException(err)
		log.Printf("%q %q", mes, err)
	}
}

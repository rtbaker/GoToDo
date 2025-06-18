package mysql

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func NewDB(dsn string) (*sql.DB, error) {
	cfg, err := mysql.ParseDSN(dsn)

	if err != nil {
		return nil, fmt.Errorf("mysql DSN parse error: %s", err)
	}

	// We allow the driver to convert MySQL DateTime to Golang time.Time
	cfg.ParseTime = true

	connector, err := mysql.NewConnector(cfg)
	if err != nil {
		return nil, fmt.Errorf("mysql connector error: %s", err)
	}

	db := sql.OpenDB(connector)

	err = db.Ping()

	if err != nil {
		db.Close()
		return nil, fmt.Errorf("mysql ping error: %s", err)
	}

	return db, nil
}

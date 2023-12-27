package database

import (
	"fmt"
	"golang-ddd/internal/config"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectPostgres(cfg config.DBConfig) (db *sqlx.DB, err error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
			cfg.Host, 
			cfg.Port, 
			cfg.User, 
			cfg.Password, 
			cfg.Name,
		)

	db, err = sqlx.Open("postgres", dsn)

	if err != nil {
		return
	}

	if err = db.Ping(); err != nil {
		return
	}

	db.SetConnMaxIdleTime(time.Duration(cfg.ConnectionPoll.MaxIdleConnection) * time.Second)
	db.SetConnMaxLifetime(time.Duration(cfg.ConnectionPoll.MaxLifetimeConnection) * time.Second)
	db.SetMaxOpenConns(int(cfg.ConnectionPoll.MaxOpenConnection))
	db.SetMaxIdleConns(int(cfg.ConnectionPoll.MaxIdleConnection))

	return

}
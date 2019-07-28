package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func InitDB() {
	db, err = sql.Open("postgres", "host=127.0.0.1 port=5555 user=root password=password dbname=menudb sslmode=disable")
	if err != nil {
		panic(err)
	}

	migration(db)
}

func migration(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS menus(
		id SERIAL PRIMARY KEY SERIAL,
		title VARCHAR NOT NULL,
		cooking_time INTEGER,
		image_url VARCHAR(2083),
		url VARCHAR(2083),
		calorie INTEGER,
		category INTEGER
	);
	`

	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ngongo")
}

func GetDB() *sql.DB {
	return db
}

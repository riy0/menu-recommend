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

func migrationMenus(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS menus(
		id SERIAL PRIMARY KEY SERIAL,
		title VARCHAR NOT NULL,
		cooking_time VARCHAR(255),
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

func migrationIngredients(db *sql.DB) {
	sql = `
	CREATE TABLE IF NOT EXISTS ingredients (
		id SERIAL PRIMARY KEY,
		menu_id integer REFERENCES menus (id) NOT NULL,
		name VARCHAR(255) NOT NULL
	);
	`

	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ngo")
}

func GetDB() *sql.DB {
	return db
}

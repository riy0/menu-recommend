package main

import (
	"github.com/riy0/menu_recommend/db"
	"log"
)

func main() {
	db.InitDB()

	err := FoodCollecting()
	if err != nil {
		log.Fatal(err)
	}

	err := detailCollecting()
	if err != nil {
		log.Fatal(err)
	}
}

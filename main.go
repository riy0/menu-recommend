package main

import (
	"github.com/riy0/menu_recommend/db"
	"github.com/riy0/menu_recommend/router"
)

func main() {
	db.InitDB()
	router.Init()
}

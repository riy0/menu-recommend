package main

import (
	"log"
	"github.com/riy0/menu_recommend/db"
)

func main(){
	db.InitDB()
	err := ebaraFoodCollecting()
	if err != nil{
		log.Fatal(err)
	}
}


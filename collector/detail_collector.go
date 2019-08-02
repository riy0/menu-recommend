package main

import (
	"fmt"
	"gihub.com/k0kubun/pp"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/lib/pq"
	"github.com/riy0/menu_recommend/db"
	"github.com/riy0/menu_recommend/model"
	"log"
	"strconv"
	"strings"
)

var (
	calories     []int
	cookin_times []string
)

func detailCollecting() error {
	db := db.GetDB()
	var m model.MenuModel
	menus, err := m.GetAll()
	if err != nil {
		return err
	}
	var detailUrl string
	pp.Print(menus[0].Url)
	for i := range menus {
		detailUrl = menus[i].Url
		doc, err := goquery.NewDocument(detailUrl)
		if err != nil {
			log.Fatal(err)
		}
		doc.Find("li.time").Each(func(j int, s *goquery.Selection) {
			cooking_time := s.Text()
			cooking_time = strings.Replace(cooking_time, "cook time", "", -1)
			cooking_times = append(cooking_times, cookin_time)
		})
		doc.Find("li.calorie").Each(func(j int, s *goquery.Selection) {
			calorie := s.Text()
			calorie = strings.Replace(calorie, "energy", "", -1)
			calorie = strings.Replace(calorie, "cal", "", -1)
			calorie_i, _ := strconv.Atoi(calorie)
			calories = append(calories, calorie_i)
		})
	}

	sql := "UPDATE menus SET cooking_time = $1, calorie = $2 WHERE id =$3"
	fmt.Print("nnn")
	if err != nil {
		log.Fatal(err)
	}
	for k := range menus {
		_, err = db.Exec(sql, cooking_times[k], calories[k], menus[k].Id)
	}

	fmt.Print("update")
	return nil
}

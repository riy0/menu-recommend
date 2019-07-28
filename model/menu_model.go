package model

import (
	"github.com/riy0/menu_recommend/db"
)

type Menu struct {
	Id          int64  `json:title`
	Title       string `json:title`
	CookingTime int    `json:cooking_time`
	imageUrl    string `json:image_url`
	Url         string `json:url`
	Calorie     int    `json:calorie`
	Category    int    `json:category`
}

type MenuModel struct{}

func (m MenuModel) GetAll() ([]Menu, error) {
	db := db.GetDB()
	var menus []Menu
	sql := "SELECT * FROM menus"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		menu = Menu{}
		err = rows.Scan(&menu.Id, &menu.Title, &menu.CookingTime, &menu.ImageUrl, &menu.Url, &menu.Calorie, &Menu.Category)
		if err != nil {
			return nil, err
		}
		menus = append(menus, menu)
	}
	return menus, nil
}

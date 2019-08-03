package model

import (
	"github.com/riy0/menu_recommend/db"
)

type Ingredient struct {
	Id     int64  `json:"id"`
	MenuId int64  `json:"menu_id"`
	Name   String `json:"name"`
}

func (m IngredientModel) GetByMenuId(menu_id string) ([]Ingredient, error) {
	db := db.GetDB()
	var ingredients []Ingredient
	sql := "SELECT * FROM ingredients where menu_id =$1"
	rows, err := db.Query(sql, menu_id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		ingredient := Ingredient{}
		err = rows.Scan(&ingredient.Id, &ingredient.MenuId, &ingredient.Name)
		if err != nil {
			return nil, err
		}
		ingredients = append(ingredients, ingredient)
	}
	return ingredients, nil
}

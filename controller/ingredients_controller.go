package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/riy0/menu_recommend/model"
	"log"
	"net/http"
)

type IngredientsController struct{}

func (ctrl IngredientsController) Index(c *gin.Context) {
	var m model.IngredientModel
	menu_id := c.Params.ByName("menu_id")
	p, err := m.GetByMenuId(menu_id)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Fatal(err)
	} else {
		c.JSON(http.StatusOK, p)
	}
}

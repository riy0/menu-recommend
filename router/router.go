package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/riy0/menu_recommend/controller"
)

func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())

	o := r.Group("/menus")
	{
		menu := new(controller.MenusController)
		o.GET("", menu.Index)
		o.GET("/:id", menu.Show)
	}

	i := r.Group("/ingredient")
	{
		ingredient := new(controller.IngredientsController)
		i.GET("/:menu_id", ingredient.Index)
	}

	return r
}

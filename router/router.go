package router

import (
	"github.com/gin-gonic/gin"
	"github.com/riy0/menu_recommend/controller"
)

func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()

	o := r.Group("/menus")
	{
		menu := new(controller.MenusController)
		o.GET("", menu.Index)
		o.GET("/:id", menu.Show)
	}

	return r
}

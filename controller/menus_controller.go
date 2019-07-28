package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/riy0/menu_recommend/model"
	"log"
	"net/http"
)

type MenusController struct{}

func (ctrl MenuController) Index(c *gin.Context) {
	var m model.MenuModel
	p, err := m.GetAll()
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Fatal(err)
	} else {
		c.JSON(http.StatusOK, p)
	}
}

func (ctrl MenusController) Show(c *gin.Context) {
	var m model.MenuModel
	id := c.Params.ByName("id")
	p, err := m.GetById(id)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Fatal(err)
	} else {
		c.JSON(http.StatusOK, p)
	}
}

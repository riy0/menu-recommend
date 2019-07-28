package controller

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/riy0/menu_recommend/model"
)

type MenusController struct{}

func (ctrl MenusController) Index(c *gin.Context){
	var m modl.MenuModel
	p, err := m.GetAll()
	if err != nil{
		c.AbortWithStatus(http.StatusBadRequest)
		log.Fatal(err)
	} else {
		c.JSON(http.statusOK, p)
	}
}

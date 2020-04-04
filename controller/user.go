package controller

import (
	"github.com/gin-gonic/gin"
	s "api/services"
	"fmt"
)

type Data struct {
	Url string `json:"url"`
}

func Login(c *gin.Context)  {
	conf, err := s.GetGits()
	if err != nil {
		panic(err)
	}
	var data Data
	data.Url = fmt.Sprintf(conf.Url, conf.Client_id)
	c.JSON(200, gin.H{
		"data": data,
		"msg": "",
	})
}

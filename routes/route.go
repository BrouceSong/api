package routes

import (
	c "api/controller"
	m "api/middleware"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) *gin.Engine {
	r.Use(m.SetHeader())
	v1 := r.Group("/v1")
	{
		v1.GET("/login", c.Login)
		v1.GET("/callback", c.Callback)
		v1.GET("/user", c.GetUser)
		post := v1.Group("/post")
		post.GET("list", c.PostLists)
	}
	return r
}

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	router.SetHTMLTemplate(Tpl)
	router.UseRawPath = true

	handler := Handler{
		Name: "Handler",
	}

	handler.Register(router)

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/home")
	})

	router.Run() //listen and serve on 0.0.0.0:8080
}

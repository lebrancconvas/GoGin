package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRoute(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "User Route accessing complete.")
	})
}


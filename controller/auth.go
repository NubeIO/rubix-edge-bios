package controller

import (
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (inst *Controller) HandleAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorized := auth.Authorize(c.Request)
		if !authorized {
			c.JSON(http.StatusUnauthorized, Message{Message: "unauthorized access"})
			c.Abort()
			return
		}
		c.Next()
		return
	}
}

package controller

import "github.com/gin-gonic/gin"

func (inst *Controller) Ping(c *gin.Context) {
	reposeHandler(Message{Message: "boo-ya"}, err, c)
}

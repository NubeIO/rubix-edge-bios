package controller

import (
	"github.com/gin-gonic/gin"
)

func (inst *Controller) GetProduct(c *gin.Context) {
	data, err := inst.Rubix.App.GetProduct() // https://github.com/NubeIO/lib-command/blob/master/product/product.go#L7
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}

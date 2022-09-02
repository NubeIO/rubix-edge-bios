package controller

import (
	"github.com/gin-gonic/gin"
)

func (inst *Controller) GetProduct(c *gin.Context) {
	data, err := inst.EdgeApp.App.GetProduct() // https://github.com/NubeIO/lib-command/blob/master/product/product.go#L7
	if err != nil {
		responseHandler(data, err, c)
		return
	}
	responseHandler(data, err, c)
}

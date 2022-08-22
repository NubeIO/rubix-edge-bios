package controller

import (
	"github.com/NubeIO/lib-networking/networking"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/rubix-edge/pkg/model"
	"github.com/gin-gonic/gin"
)

type DeviceProduct struct {
	Device     *model.DeviceInfo  `json:"device"`
	Product    *installer.Product `json:"product"`
	Networking []networking.NetworkInterfaces
}

func (inst *Controller) GetDeviceProduct(c *gin.Context) {
	device, err := inst.DB.GetDeviceInfo()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	product, err := inst.Rubix.App.GetProduct() // https://github.com/NubeIO/lib-command/blob/master/product/product.go#L7
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	networks, err := nets.GetNetworks()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	deviceProduct := &DeviceProduct{
		Device:     device,
		Product:    product,
		Networking: networks,
	}
	reposeHandler(deviceProduct, err, c)
}

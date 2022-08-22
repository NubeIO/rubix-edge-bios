package controller

import (
	"github.com/NubeIO/lib-networking/networking"
	"github.com/NubeIO/rubix-edge/service/system"
	"github.com/gin-gonic/gin"
)

var nets = networking.New()

func (inst *Controller) Networking(c *gin.Context) {
	data, err := nets.GetNetworks()
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) GetInterfacesNames(c *gin.Context) {
	data, err := nets.GetInterfacesNames()
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) InternetIP(c *gin.Context) {
	data, err := nets.GetInternetIP()
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) RestartNetworking(c *gin.Context) {
	data, err := inst.System.RestartNetworking()
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) InterfaceUp(c *gin.Context) {
	var m system.NetworkingBody
	err = c.ShouldBindJSON(&m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, err := inst.System.InterfaceUp(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) InterfaceDown(c *gin.Context) {
	var m system.NetworkingBody
	err = c.ShouldBindJSON(&m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, err := inst.System.InterfaceDown(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

package controller

import (
	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/externaltoken"
	interfaces2 "github.com/NubeIO/rubix-edge-bios/interfaces"
	"github.com/gin-gonic/gin"
)

func getBodyTokenCreate(c *gin.Context) (dto *interfaces2.TokenCreate, err error) {
	err = c.ShouldBindJSON(&dto)
	return dto, err
}

func getBodyTokenBlock(ctx *gin.Context) (dto *interfaces2.TokenBlock, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func (inst *Controller) GetTokens(c *gin.Context) {
	q, err := externaltoken.GetExternalTokens()
	responseHandler(q, err, c)
}

func (inst *Controller) GetToken(c *gin.Context) {
	u := c.Param("uuid")
	q, err := externaltoken.GetExternalToken(u)
	responseHandler(q, err, c)
}

func (inst *Controller) GenerateToken(c *gin.Context) {
	body, err := getBodyTokenCreate(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	q, err := externaltoken.CreateExternalToken(&externaltoken.ExternalToken{
		UUID:    uuid.ShortUUID("tok"),
		Name:    body.Name,
		Blocked: *body.Blocked})
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(q, err, c)
}

func (inst *Controller) RegenerateToken(c *gin.Context) {
	u := c.Param("uuid")
	q, err := externaltoken.RegenerateExternalToken(u)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(q, err, c)
}

func (inst *Controller) BlockToken(c *gin.Context) {
	u := c.Param("uuid")
	body, err := getBodyTokenBlock(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	q, err := externaltoken.BlockExternalToken(u, *body.Blocked)
	responseHandler(q, err, c)
}

func (inst *Controller) DeleteToken(c *gin.Context) {
	u := c.Param("uuid")
	q, err := externaltoken.DeleteExternalToken(u)
	responseHandler(q, err, c)
}

package controller

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/externaltoken"
	"github.com/NubeIO/rubix-edge/pkg/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getBodyTokenCreate(c *gin.Context) (dto *interfaces.TokenCreate, err error) {
	err = c.ShouldBindJSON(&dto)
	return dto, err
}

func getBodyTokenBlock(ctx *gin.Context) (dto *interfaces.TokenBlock, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func (inst *Controller) GetTokens(c *gin.Context) {
	q, err := externaltoken.GetExternalTokens()
	reposeHandler(q, err, c)
}

func (inst *Controller) GenerateToken(c *gin.Context) {
	body, err := getBodyTokenCreate(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	u := uuid.New().String()
	q, err := externaltoken.CreateExternalToken(&externaltoken.ExternalToken{
		UUID:    fmt.Sprintf("tok_%s", u),
		Name:    body.Name,
		Blocked: *body.Blocked})
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(q, err, c)
}

func (inst *Controller) RegenerateToken(c *gin.Context) {
	u := c.Param("uuid")
	q, err := externaltoken.RegenerateExternalToken(u)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(q, err, c)
}

func (inst *Controller) BlockToken(c *gin.Context) {
	u := c.Param("uuid")
	body, err := getBodyTokenBlock(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	q, err := externaltoken.BlockExternalToken(u, *body.Blocked)
	reposeHandler(q, err, c)
}

func (inst *Controller) DeleteToken(c *gin.Context) {
	u := c.Param("uuid")
	q, err := externaltoken.DeleteExternalToken(u)
	reposeHandler(q, err, c)
}

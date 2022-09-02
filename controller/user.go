package controller

import (
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/user"
	interfaces2 "github.com/NubeIO/rubix-edge-bios/interfaces"
	nerrors2 "github.com/NubeIO/rubix-edge-bios/nerrors"
	"github.com/gin-gonic/gin"
)

func getBodyUser(c *gin.Context) (dto *user.User, err error) {
	err = c.ShouldBindJSON(&dto)
	return dto, err
}

func (inst *Controller) Login(c *gin.Context) {
	body, err := getBodyUser(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	q, err := user.Login(body)
	if err != nil {
		responseHandler(nil, nerrors2.NewErrUnauthorized(err.Error()), c)
		return
	}
	responseHandler(interfaces2.TokenResponse{AccessToken: q, TokenType: "JWT"}, err, c)
}

func (inst *Controller) UpdateUser(c *gin.Context) {
	body, err := getBodyUser(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	q, err := user.CreateUser(body)
	responseHandler(q, err, c)
}

func (inst *Controller) GetUser(c *gin.Context) {
	q, err := user.GetUser()
	responseHandler(q, err, c)
}

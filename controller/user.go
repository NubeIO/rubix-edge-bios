package controller

import (
	"github.com/NubeIO/nubeio-rubix-lib-auth-go/user"
	"github.com/NubeIO/rubix-edge/pkg/interfaces"
	"github.com/NubeIO/rubix-edge/pkg/nerrors"
	"github.com/gin-gonic/gin"
)

func getBodyUser(c *gin.Context) (dto *user.User, err error) {
	err = c.ShouldBindJSON(&dto)
	return dto, err
}

func (inst *Controller) Login(c *gin.Context) {
	body, err := getBodyUser(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	q, err := user.Login(body)
	if err != nil {
		reposeHandler(nil, nerrors.NewErrUnauthorized(err.Error()), c)
		return
	}
	reposeHandler(interfaces.TokenResponse{AccessToken: q, TokenType: "JWT"}, err, c)
}

func (inst *Controller) UpdateUser(c *gin.Context) {
	body, err := getBodyUser(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	q, err := user.CreateUser(body)
	reposeHandler(q, err, c)
}

func (inst *Controller) GetUser(c *gin.Context) {
	q, err := user.GetUser()
	reposeHandler(q, err, c)
}

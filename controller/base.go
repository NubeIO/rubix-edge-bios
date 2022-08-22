package controller

import (
	"fmt"
	fileutils "github.com/NubeIO/lib-dirs/dirs"
	dbase "github.com/NubeIO/rubix-edge/database"
	"github.com/NubeIO/rubix-edge/service/apps"
	"github.com/NubeIO/rubix-edge/service/system"
	"github.com/gin-gonic/gin"
	"net/http"
)

const nonRoot = 0700
const root = 0777

var fileUtils = fileutils.New()
var filePerm = root

type Controller struct {
	DB     *dbase.DB
	Rubix  *apps.EdgeApps
	System *system.System
}

var err error

type Response struct {
	StatusCode   int         `json:"status_code"`
	ErrorMessage string      `json:"error_message"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data"`
}

func reposeHandler(body interface{}, err error, c *gin.Context, statusCode ...int) {
	var code int
	if err != nil {
		if len(statusCode) > 0 {
			code = statusCode[0]
		} else {
			code = http.StatusNotFound
		}
		msg := Message{
			Message: fmt.Sprintf("rubix-edge: %s", err.Error()),
		}
		c.JSON(code, msg)
	} else {
		if len(statusCode) > 0 {
			code = statusCode[0]
		} else {
			code = http.StatusOK
		}
		c.JSON(code, body)
	}
}

type Message struct {
	Message interface{} `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

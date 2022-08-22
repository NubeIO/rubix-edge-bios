package controller

import (
	"errors"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/system/files"
	"github.com/gin-gonic/gin"
	"os"
	"path"
	"strconv"
)

func (inst *Controller) Unzip(c *gin.Context) {
	source := c.Query("source")
	destination := c.Query("destination")
	perm := c.Query("permission")
	var permission int
	if perm == "" {
		permission = filePerm
	} else {
		permission, err = strconv.Atoi(c.Query("permission"))
		if err != nil {
			permission = filePerm
		}
	}
	pathToZip := source
	if source == "" {
		reposeHandler(nil, errors.New("zip source can not be empty, try /data/zip.zip"), c)
		return
	}
	if destination == "" {
		reposeHandler(nil, errors.New("zip destination can not be empty, try /data/unzip-test"), c)
		return
	}
	zip, err := fileUtils.UnZip(pathToZip, destination, os.FileMode(permission))
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(Message{Message: fmt.Sprintf("unzipped successfully, files count: %d", len(zip))}, err, c)
}

func (inst *Controller) ZipDir(c *gin.Context) {
	source := c.Query("source")
	destination := c.Query("destination")
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	pathToZip := source
	if source == "" {
		reposeHandler(nil, errors.New("zip source can not be empty, try /data/flow-framework"), c)
		return
	}
	if destination == "" {
		reposeHandler(nil, errors.New("zip destination can not be empty, try /home/test/flow-framework.zip"), c)
		return
	}

	exists := fileUtils.DirExists(pathToZip)
	if !exists {
		reposeHandler(nil, errors.New("dir to zip not found"), c)
		return
	}
	err := files.MakeDirectoryIfNotExists(path.Dir(destination))
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	err = fileUtils.RecursiveZip(pathToZip, destination)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(Message{Message: fmt.Sprintf("zip file is created on: %s", destination)}, nil, c)
}

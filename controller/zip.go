package controller

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-files/fileutils"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/system/files"
	"github.com/gin-gonic/gin"
	"os"
	"path"
)

func (inst *Controller) Unzip(c *gin.Context) {
	source := c.Query("source")
	destination := c.Query("destination")
	pathToZip := source
	if source == "" {
		responseHandler(nil, errors.New("zip source can not be empty, try /data/zip.zip"), c)
		return
	}
	if destination == "" {
		responseHandler(nil, errors.New("zip destination can not be empty, try /data/unzip-test"), c)
		return
	}
	zip, err := fileutils.UnZip(pathToZip, destination, os.FileMode(inst.FileMode))
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(Message{Message: fmt.Sprintf("unzipped successfully, files count: %d", len(zip))}, err, c)
}

func (inst *Controller) ZipDir(c *gin.Context) {
	source := c.Query("source")
	destination := c.Query("destination")
	pathToZip := source
	if source == "" {
		responseHandler(nil, errors.New("zip source can not be empty, try /data/flow-framework"), c)
		return
	}
	if destination == "" {
		responseHandler(nil, errors.New("zip destination can not be empty, try /home/test/flow-framework.zip"), c)
		return
	}

	exists := fileutils.DirExists(pathToZip)
	if !exists {
		responseHandler(nil, errors.New("dir to zip not found"), c)
		return
	}
	err := files.MakeDirectoryIfNotExists(path.Dir(destination))
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	err = fileutils.RecursiveZip(pathToZip, destination)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(Message{Message: fmt.Sprintf("zip file is created on: %s", destination)}, nil, c)
}

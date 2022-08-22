package controller

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ListApps apps by listed in the installation (/data/rubix-service/apps/install)
func (inst *Controller) ListApps(c *gin.Context) {
	data, err := inst.Rubix.App.ListApps()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

// ListAppsAndService get all the apps by listed in the installation (/data/rubix-service/apps/install) dir and then check the service
func (inst *Controller) ListAppsAndService(c *gin.Context) {
	data, err := inst.Rubix.App.ListAppsAndService()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

// ListNubeServices list all the services by filtering all the service files with name nubeio
func (inst *Controller) ListNubeServices(c *gin.Context) {
	data, err := inst.Rubix.App.ListNubeServices()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

// ConfirmAppInstalled will check the app is installed and the service exists
func (inst *Controller) ConfirmAppInstalled(c *gin.Context) {
	data, err := inst.Rubix.App.ConfirmAppInstalled(c.Query("name"))
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) GetAppVersion(c *gin.Context) {
	data := inst.Rubix.App.GetAppVersion(c.Query("name"))
	if data == "" {
		reposeHandler(nil, errors.New(fmt.Sprintf("failed to find the app:%s", c.Query("name"))), c)
		return
	}
	reposeHandler(data, nil, c)
}

// AddUploadApp
// upload the build
func (inst *Controller) AddUploadApp(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	m := &installer.Upload{
		Name:    c.Query("name"),
		Version: c.Query("version"),
		Product: c.Query("product"),
		Arch:    c.Query("arch"),
		File:    file,
	}
	data, err := inst.Rubix.App.AddUploadEdgeApp(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

// UploadService
// upload the service file
func (inst *Controller) UploadService(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	m := &installer.Upload{
		Name:    c.Query("name"),
		Version: c.Query("version"),
		File:    file,
	}
	data, err := inst.Rubix.App.UploadServiceFile(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) InstallService(c *gin.Context) {
	var m *installer.Install
	err = c.ShouldBindJSON(&m)
	data, err := inst.Rubix.App.InstallService(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

// UninstallApp full uninstallation of an app
func (inst *Controller) UninstallApp(c *gin.Context) {
	deleteApp, _ := strconv.ParseBool(c.Query("delete"))
	data, err := inst.Rubix.App.UninstallApp(c.Query("name"), deleteApp)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) RemoveApp(c *gin.Context) {
	err := inst.Rubix.App.RemoveApp(c.Query("name"))
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(Message{
		Message: "remove app ok",
	}, nil, c)
}

func (inst *Controller) RemoveAppInstall(c *gin.Context) {
	err := inst.Rubix.App.RemoveAppInstall(c.Query("name"))
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(Message{
		Message: "deleted app ok",
	}, nil, c)
}

// BuildUpload
// upload the build, or plugin
func (inst *Controller) BuildUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, err := inst.Rubix.App.Upload(file)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

// CompareBuildToArch compare the arch and product to the zip build name
func (inst *Controller) CompareBuildToArch(c *gin.Context) {
	zipName := c.Query("build_zip_name")
	product := c.Query("product")
	err := inst.Rubix.App.CompareBuildToArch(zipName, product)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(Message{
		Message: "match ok",
	}, nil, c)
}

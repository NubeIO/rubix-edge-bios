package controller

import (
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/gin-gonic/gin"
	"strconv"
)

/*
RESTORE A BACK-UP
*/

func (inst *Controller) RestoreBackup(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	takeBackup := c.Query("take_backup")
	_takeBackup, _ := strconv.ParseBool(takeBackup)
	m := &installer.RestoreBackup{
		TakeBackup: _takeBackup,
		File:       file,
	}
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, err := inst.Rubix.RestoreBackup(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) RestoreAppBackup(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	takeBackup := c.Query("take_backup")
	_takeBackup, _ := strconv.ParseBool(takeBackup)
	m := &installer.RestoreBackup{
		AppName:    c.Query("app_name"),
		TakeBackup: _takeBackup,
		File:       file,
	}
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, err := inst.Rubix.RestoreAppBackup(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

/*
RUN A BACK-UP
*/

func (inst *Controller) FullBackUp(c *gin.Context) {
	data, err := inst.Rubix.FullBackUp()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) BackupApp(c *gin.Context) {
	appName := c.Query("app_name")
	data, err := inst.Rubix.BackupApp(appName)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

/*
LIST BACK-UPS
*/

func (inst *Controller) ListFullBackups(c *gin.Context) {
	data, err := inst.Rubix.ListFullBackups()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) ListAppBackupsDirs(c *gin.Context) {
	data, err := inst.Rubix.ListAppBackupsDirs()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) ListBackupsByApp(c *gin.Context) {
	appName := c.Query("app_name")
	data, err := inst.Rubix.ListBackupsByApp(appName)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

/*
DELETE BACK-UPS
*/

func (inst *Controller) WipeBackups(c *gin.Context) {
	data, err := inst.Rubix.WipeBackups()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) DeleteAllFullBackups(c *gin.Context) {
	data, err := inst.Rubix.DeleteAllFullBackups()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) DeleteAllAppBackups(c *gin.Context) {
	data, err := inst.Rubix.DeleteAllAppBackups()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) DeleteAppAllBackUpByName(c *gin.Context) {
	appName := c.Query("name")
	data, err := inst.Rubix.DeleteAppAllBackUpByName(appName)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) DeleteAppOneBackUpByName(c *gin.Context) {
	appName := c.Query("name")
	folder := c.Query("folder")
	data, err := inst.Rubix.DeleteAppOneBackUpByName(appName, folder)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

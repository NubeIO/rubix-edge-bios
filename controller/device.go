package controller

import (
	"encoding/json"
	"github.com/NubeIO/rubix-edge-bios/constants"
	"github.com/NubeIO/rubix-edge-bios/pkg/model"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
	"time"
)

func (inst *Controller) GetDeviceInfo(c *gin.Context) {
	deviceInfo, err := inst.GetDeviceInfoFunc()
	reposeHandler(deviceInfo, err, c)
}

func (inst *Controller) UpdateDeviceInfo(c *gin.Context) {
	var deviceInfo *model.DeviceInfo
	err := c.ShouldBindJSON(&deviceInfo)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}

	deviceInfoOld, err := inst.GetDeviceInfoFunc()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}

	deviceInfo.GlobalUUID = deviceInfoOld.GlobalUUID
	deviceInfo.CreatedOn = deviceInfoOld.CreatedOn
	deviceInfo.UpdatedOn = strings.TrimSuffix(time.Now().UTC().Format(time.RFC3339Nano), "Z")

	deviceInfoDefault := model.DeviceInfoDefault{
		DeviceInfoFirstRecord: model.DeviceInfoFirstRecord{
			DeviceInfo: *deviceInfo,
		},
	}
	deviceInfoDefaultRaw, err := json.Marshal(deviceInfoDefault)
	err = os.WriteFile(constants.RubixRegistryFile, deviceInfoDefaultRaw, constants.Permission)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(deviceInfo, err, c)
}

func (inst *Controller) GetDeviceInfoFunc() (*model.DeviceInfo, error) {
	data, err := os.ReadFile(constants.RubixRegistryFile)
	if err != nil {
		return nil, err
	}
	deviceInfoDefault := model.DeviceInfoDefault{}
	err = json.Unmarshal(data, &deviceInfoDefault)
	if err != nil {
		return nil, err
	}
	return &deviceInfoDefault.DeviceInfoFirstRecord.DeviceInfo, nil
}

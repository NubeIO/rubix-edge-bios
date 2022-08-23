package dbase

import (
	"errors"
	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/NubeIO/rubix-edge/pkg/model"
	"time"
)

const deviceInfo = "device info"

func (db *DB) GetDeviceInfo() (*model.DeviceInfo, error) {
	infos, err := db.getAllDeviceInfos()
	if err != nil {
		return nil, err
	}
	if len(infos) == 0 {
		dev, err := db.AddDeviceInfo(&model.DeviceInfo{
			ClientId:    "N/A",
			ClientName:  "N/A",
			SiteId:      "N/A",
			SiteName:    "N/A",
			DeviceId:    "N/A",
			DeviceName:  "N/A",
			SiteAddress: "N/A",
			SiteCity:    "N/A",
			SiteState:   "N/A",
			SiteZip:     "N/A",
			SiteCountry: "N/A",
			SiteLat:     "",
			SiteLon:     "",
			TimeZone:    "",
			CreatedOn:   time.Now(),
			UpdatedOn:   time.Now(),
		})
		if err != nil {
			return dev, err
		}
		return nil, err
	}
	var m *model.DeviceInfo
	if err := db.DB.Where("uuid = ? ", infos[0].UUID).First(&m).Error; err != nil {
		return nil, handleNotFound(deviceInfo)
	}
	return m, nil
}

func (db *DB) getAllDeviceInfos() ([]*model.DeviceInfo, error) {
	var m []*model.DeviceInfo
	if err := db.DB.Find(&m).Error; err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

func (db *DB) AddDeviceInfo(body *model.DeviceInfo) (resp *model.DeviceInfo, err error) {
	infos, err := db.getAllDeviceInfos()
	if err != nil {
		return nil, err
	}
	if len(infos) > 0 {
		return nil, errors.New("device info can only be added once")
	}
	body.UUID = uuid.ShortUUID("eid")
	if err := db.DB.Create(&body).Error; err != nil {
		return nil, err
	} else {
		return body, nil
	}
}

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
			ClientId:    "na",
			ClientName:  "na",
			SiteId:      "na",
			SiteName:    "na",
			DeviceId:    "na",
			DeviceName:  "na",
			SiteAddress: "na",
			SiteCity:    "na",
			SiteState:   "na",
			SiteZip:     "na",
			SiteCountry: "na",
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

func (db *DB) UpdateDeviceInfo(app *model.DeviceInfo) (*model.DeviceInfo, error) {
	infos, err := db.getAllDeviceInfos()
	if err != nil {
		return nil, err
	}
	if len(infos) == 0 {
		return nil, errors.New("device info has not been added yet")
	}
	var m *model.DeviceInfo
	query := db.DB.Where("uuid = ?", infos[0].UUID).Find(&m).Updates(app)
	if query.Error != nil {
		return nil, handleNotFound(deviceInfo)
	} else {
		return app, query.Error
	}
}

func (db *DB) DeleteDeviceInfo(uuid string) (*DeleteMessage, error) {
	var m *model.DeviceInfo
	query := db.DB.Where("uuid = ? ", uuid).Delete(&m)
	return deleteResponse(query)
}

// DropDeviceInfo delete all.
func (db *DB) DropDeviceInfo() (*DeleteMessage, error) {
	var m *model.DeviceInfo
	query := db.DB.Where("1 = 1")
	query.Delete(&m)
	return deleteResponse(query)
}

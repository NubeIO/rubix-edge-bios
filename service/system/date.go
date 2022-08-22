package system

import (
	"fmt"
	"github.com/NubeIO/lib-date/datectl"
	"github.com/NubeIO/lib-date/datelib"
)

type DateBody struct {
	DateTime string `json:"date_time"`
	TimeZone string `json:"time_zone"`
}

func (inst *System) SystemTime() *datelib.Time {
	return datelib.New(&datelib.Date{}).SystemTime()
}

func (inst *System) GenerateTimeSyncConfig(body *datectl.TimeSyncConfig) string {
	return inst.datectl.GenerateTimeSyncConfig(body)
}

func (inst *System) GetHardwareTZ() (string, error) {
	return inst.datectl.GetHardwareTZ()
}

func (inst *System) GetHardwareClock() (*datectl.HardwareClock, error) {
	return inst.datectl.GetHardwareClock()
}

func (inst *System) GetTimeZoneList() ([]string, error) {
	return inst.datectl.GetTimeZoneList()
}

func (inst *System) UpdateTimezone(body DateBody) (*Message, error) {
	err := inst.datectl.UpdateTimezone(body.TimeZone)
	if err != nil {
		return nil, err
	}
	return &Message{
		Message: fmt.Sprintf("updated to %s", body.TimeZone),
	}, nil
}

func (inst *System) SetSystemTime(body DateBody) (*datelib.Time, error) {
	err := inst.datectl.SetSystemTime(body.DateTime)
	if err != nil {
		return nil, err
	}
	return datelib.New(&datelib.Date{}).SystemTime(), nil
}

package system

import (
	"fmt"
	"github.com/NubeIO/lib-dhcpd/dhcpd"
	"strconv"
)

func (inst *System) DHCPPortExists(body NetworkingBody) (*Message, error) {
	exists, err := inst.dhcp.Exists(body.PortName)
	if err != nil {
		return nil, err
	}
	return &Message{
		Message: fmt.Sprintf("%s", strconv.FormatBool(exists)),
	}, nil
}

func (inst *System) DHCPSetAsAuto(body NetworkingBody) (*Message, error) {
	exists, err := inst.dhcp.SetAsAuto(body.PortName)
	if err != nil {
		return nil, err
	}
	msg := fmt.Sprintf("was not able :%s to auto", body.PortName)
	if exists {
		msg = fmt.Sprintf("was able to set interface :%s to auto", body.PortName)
	}
	return &Message{
		Message: msg,
	}, nil
}

func (inst *System) DHCPSetStaticIP(body *dhcpd.SetStaticIP) (string, error) {
	return inst.dhcp.SetStaticIP(body)
}

package system

import (
	pprint "github.com/NubeIO/rubix-edge/pkg/helpers/print"
	"testing"
)

func TestSystem_GetHardwareClock(t *testing.T) {
	clock, err := New(&System{}).GetHardwareClock()
	if err != nil {
		return
	}
	pprint.PrintJOSN(clock)
}

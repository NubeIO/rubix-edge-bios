package system

import (
	pprint "github.com/NubeIO/rubix-edge/pkg/helpers/print"
	"testing"
)

func TestGetFlowToken(t *testing.T) {
	token, err := New(&System{}).GetFlowToken()
	if err != nil {
		return
	}
	pprint.PrintJOSN(token)
}

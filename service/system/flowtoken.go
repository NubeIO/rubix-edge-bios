package system

import (
	"fmt"
	fileutils "github.com/NubeIO/lib-dirs/dirs"
)

type FlowToken struct {
	Token string `json:"token"`
}

// GetFlowToken do a system read to get the flow token
func (inst *System) GetFlowToken() (*FlowToken, error) {
	///data/rubix-service/data/internal_token.txt
	path := "/data/rubix-service/data"
	fileName := "internal_token.txt"
	files, err := fileutils.New().ListFiles(path)
	if err != nil {
		return nil, err
	}
	flowToken := &FlowToken{}
	if len(files) > 0 {
		for _, file := range files {
			if file == fileName {
				readFile, err := fileutils.New().ReadFile(fmt.Sprintf("%s/%s", path, fileName))
				if err != nil {
					return nil, err
				}
				if len(readFile) > 50 {
					flowToken.Token = readFile
				}

			}

		}
	}
	return flowToken, err

}

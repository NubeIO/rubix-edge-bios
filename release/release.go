package release

import (
	"embed"
	"fmt"
)

var VersionFs embed.FS

func GetVersion() string {
	content, err := VersionFs.ReadFile("VERSION")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return "v" + string(content)
}

package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/lib-files/fileutils"
	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/NubeIO/rubix-edge-bios/config"
	"github.com/NubeIO/rubix-edge-bios/constants"
	"github.com/NubeIO/rubix-edge-bios/logger"
	"github.com/NubeIO/rubix-edge-bios/model"
	"github.com/NubeIO/rubix-edge-bios/router"
	"github.com/NubeIO/rubix-registry-go/rubixregistry"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"time"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "starting rubix-edge-bios",
	Long:  "it starts a server for rubix-edge-bios",
	Run:   runServer,
}

func runServer(cmd *cobra.Command, args []string) {
	if err := config.Setup(RootCmd); err != nil {
		fmt.Errorf("error: %s", err) // here log is not setup yet...
	}
	logger.Init()
	if err := os.MkdirAll(config.Config.GetAbsDataDir(), os.FileMode(constants.Permission)); err != nil {
		panic(err)
	}
	rr := rubixregistry.New()
	err := rr.CreateDeviceInfoIfDoesNotExist()
	if err != nil {
		panic(err)
	}
	logger.Logger.Infoln("starting edge-bios...")

	r := router.Setup()

	host := "0.0.0.0"
	port := config.Config.GetPort()
	logger.Logger.Infof("server is starting at %s:%s", host, port)
	logger.Logger.Fatalf("%v", r.Run(fmt.Sprintf("%s:%s", host, port)))
}

func init() {
	RootCmd.AddCommand(serverCmd)
}

func createDeviceInfoIfDoesNotExist() {
	dirExist := fileutils.DirExists(constants.RubixRegistryDir)
	if !dirExist {
		if err := os.MkdirAll(constants.RubixRegistryDir, os.FileMode(constants.Permission)); err != nil {
			panic(err)
		}
	}
	fileExist := fileutils.FileExists(constants.RubixRegistryFile)
	if !fileExist {
		deviceInfoDefault := model.DeviceInfoDefault{}
		currentDate := strings.TrimSuffix(time.Now().UTC().Format(time.RFC3339Nano), "Z")
		deviceInfoDefault.DeviceInfoFirstRecord.DeviceInfo.GlobalUUID = uuid.ShortUUID("glb")
		deviceInfoDefault.DeviceInfoFirstRecord.DeviceInfo.CreatedOn = currentDate
		deviceInfoDefault.DeviceInfoFirstRecord.DeviceInfo.UpdatedOn = currentDate
		deviceInfoDefaultRaw, err := json.Marshal(deviceInfoDefault)
		if err != nil {
			panic(err)
		}
		err = os.WriteFile(constants.RubixRegistryFile, deviceInfoDefaultRaw, constants.Permission)
		if err != nil {
			panic(err)
		}
	}
}

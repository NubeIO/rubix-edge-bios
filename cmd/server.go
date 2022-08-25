package cmd

import (
	"encoding/json"
	"fmt"
	fileutils "github.com/NubeIO/lib-dirs/dirs"
	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/NubeIO/rubix-edge-bios/constants"
	"github.com/NubeIO/rubix-edge-bios/pkg/config"
	"github.com/NubeIO/rubix-edge-bios/pkg/logger"
	"github.com/NubeIO/rubix-edge-bios/pkg/model"
	"github.com/NubeIO/rubix-edge-bios/pkg/router"
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
	createDeviceInfoIfDoesNotExist()
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
	fileUtils := fileutils.New()
	dirExist := fileUtils.DirExists(constants.RubixRegistryDir)
	if !dirExist {
		if err := os.MkdirAll(constants.RubixRegistryDir, os.FileMode(constants.Permission)); err != nil {
			panic(err)
		}
	}
	fileExist := fileUtils.FileExists(constants.RubixRegistryFile)
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

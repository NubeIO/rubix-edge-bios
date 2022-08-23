package cmd

import (
	"fmt"
	"github.com/NubeIO/rubix-edge-bios/pkg/config"
	"github.com/NubeIO/rubix-edge-bios/pkg/logger"
	"github.com/NubeIO/rubix-edge-bios/pkg/router"
	"github.com/spf13/cobra"
	"os"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "system admin for edge28",
	Long:  "pass in the host name and do operation like check arch type of the host",
	Run:   runServer,
}

func runServer(cmd *cobra.Command, args []string) {
	logger.Init()
	if err := config.Setup(RootCmd); err != nil {
		logger.Logger.Errorf("error: %s", err)
	}
	if err := os.MkdirAll(config.Config.GetAbsDataDir(), 0755); err != nil {
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

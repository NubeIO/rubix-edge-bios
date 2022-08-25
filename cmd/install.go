package cmd

import (
	"embed"
	"fmt"
	"github.com/NubeIO/lib-systemctl-go/ctl"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
	"github.com/NubeIO/rubix-edge-bios/pkg/config"
	"github.com/spf13/cobra"
	"os"
	"path"
	"strings"
	"syscall"
)

//go:embed systemd/nubeio-rubix-edge-bios.service
var f embed.FS

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "install rubix-edge-bios",
	Long:  "it installs rubix-edge-bios on your device using systemd",
	Run:   install,
}

func install(cmd *cobra.Command, args []string) {
	const ServiceDir = "/lib/systemd/system"
	const ServiceDirSoftLink = "/etc/systemd/system/multi-user.target.wants"
	const ServiceFileName = "nubeio-rubix-edge-bios.service"

	serviceFile := path.Join(ServiceDir, ServiceFileName)
	symlinkServiceFile := path.Join(ServiceDirSoftLink, ServiceFileName)

	if err := config.Setup(RootCmd); err != nil {
		fmt.Errorf("error: %s", err) // here log is not setup yet...
	}

	fmt.Println("installing edge-bios...")
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	content, err := f.ReadFile("systemd/nubeio-rubix-edge-bios.service")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	systemd := string(content)
	systemd = strings.Replace(systemd, "<working_dir>", wd, -1)
	fmt.Println(fmt.Sprintf("systemd file with working directory: %s", wd))

	fmt.Println(fmt.Sprintf("creating service file: %s...", serviceFile))
	err = os.WriteFile(serviceFile, []byte(systemd), 0644)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("soft un-linking linux service...")
	err = syscall.Unlink(symlinkServiceFile)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("soft linking linux service...")
	syscall.Symlink(serviceFile, symlinkServiceFile)

	fmt.Println("executing installation command...")
	service := ctl.New(ServiceFileName, "<NOT_NEEDED>")
	opts := systemctl.Options{Timeout: 30}
	installOpts := ctl.InstallOpts{
		Options: opts,
	}
	removeOpts := ctl.RemoveOpts{RemoveOpts: opts}
	service.InstallOpts = installOpts
	service.RemoveOpts = removeOpts
	service.Install()
	fmt.Println("successfully executed install command...")
}

func init() {
	RootCmd.AddCommand(installCmd)
}

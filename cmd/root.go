package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "nube-cli",
	Short: "description",
	Long:  `description`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
	}
}

var flgRoot struct {
	prod          bool
	auth          bool
	port          int
	rootDir       string
	appDir        string
	dataDir       string
	configDir     string
	host          string
	ip            string
	sshPort       int
	hostUsername  string
	hostPassword  string
	rubixPort     int
	rubixUsername string
	rubixPassword string
	iface         string
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&flgRoot.prod, "prod", "", false, "prod")
	RootCmd.PersistentFlags().BoolVarP(&flgRoot.prod, "auth", "", true, "auth")
	RootCmd.PersistentFlags().IntVarP(&flgRoot.port, "port", "p", 1661, "port (default 1661)")
	RootCmd.PersistentFlags().StringVarP(&flgRoot.rootDir, "root-dir", "r", "./", "root dir") // in production it will be `/data`
	RootCmd.PersistentFlags().StringVarP(&flgRoot.appDir, "app-dir", "a", "./", "app dir")    // in production it will be `/rubix-edge`
	RootCmd.PersistentFlags().StringVarP(&flgRoot.dataDir, "data-dir", "d", "data", "data dir")
	RootCmd.PersistentFlags().StringVarP(&flgRoot.configDir, "config-dir", "c", "config", "config dir")
	RootCmd.PersistentFlags().StringVarP(&flgRoot.host, "host", "", "RC", "host name (default RC)")
	RootCmd.PersistentFlags().StringVarP(&flgRoot.ip, "ip", "", "192.168.15.10", "host ip (default 192.168.15.10)")
	RootCmd.PersistentFlags().IntVarP(&flgRoot.sshPort, "ssh-port", "", 22, "SSH Port (default 22)")
	RootCmd.PersistentFlags().StringVarP(&flgRoot.iface, "iface", "", "", "pc or host network interface example: eth0")
	RootCmd.PersistentFlags().StringVarP(&flgRoot.hostUsername, "host-user", "", "pi", "host/linux username (default pi)")
	RootCmd.PersistentFlags().StringVarP(&flgRoot.hostPassword, "host-pass", "", "N00BRCRC", "host/linux password")
	RootCmd.PersistentFlags().IntVarP(&flgRoot.rubixPort, "rubix-port", "", 1616, "rubix port (default 1616)")
	RootCmd.PersistentFlags().StringVarP(&flgRoot.rubixUsername, "rubix-user", "", "admin", "rubix username (default admin)")
	RootCmd.PersistentFlags().StringVarP(&flgRoot.rubixPassword, "rubix-pass", "", "N00BWires", "rubix password")
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

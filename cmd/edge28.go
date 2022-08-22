package cmd

import (
	"github.com/spf13/cobra"
)

var edge28CMD = &cobra.Command{
	Use:   "edge",
	Short: "system admin for edge28",
	Long:  `pass in the host name and do operation like check arch type of the host`,
	Run:   runEdge,
}

var flgEdgeIp struct {
	updateIp   bool
	iPAddress  string
	subnetMask string
	gateway    string
	setDHCP    bool
}

func runEdge(cmd *cobra.Command, args []string) {
}

func init() {
	RootCmd.AddCommand(edge28CMD)
	edge28CMD.Flags().BoolVarP(&flgEdgeIp.updateIp, "update-ip", "", false, "update the edge-28 ip address")
	edge28CMD.Flags().StringVarP(&flgEdgeIp.iPAddress, "ip", "", "192.168.15.40", "ip addr")
	edge28CMD.Flags().StringVarP(&flgEdgeIp.subnetMask, "subnet", "", "255.255.255.0", "ip addr")
	edge28CMD.Flags().StringVarP(&flgEdgeIp.gateway, "gateway", "", "", "ip addr")
}

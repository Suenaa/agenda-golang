package cmd

import (
	"fmt"

	"github.com/Suenaa/agenda-go/service"
	"github.com/Suenaa/agenda-go/tools"
	"github.com/spf13/cobra"
	"github.com/Suenaa/agenda-go/logs"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "log out",
	Long:  `log out agenda`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.UserLogout(); err == nil {
			fmt.Println("Success")
			logs.EventLog("current user logs out")
		} else {
			tools.Report(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(logoutCmd)
}

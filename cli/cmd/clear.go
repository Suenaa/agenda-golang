package cmd

import (
	"fmt"

	"github.com/Suenaa/agenda-go/service"
	"github.com/Suenaa/agenda-go/tools"
	"github.com/spf13/cobra"
	"github.com/Suenaa/agenda-go/logs"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "clear all meetings you create",
	Long:  `clear all meetings you create`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.DeleteAllMeeting(); err == nil {
			fmt.Println("Success")
		} else {
			tools.Report(err)
			logs.EventLog("clear all meetings")
		}
	},
}

func init() {
	RootCmd.AddCommand(clearCmd)
}

package cmd

import (
	"errors"
	"fmt"

	"github.com/Suenaa/agenda-go/service"
	"github.com/Suenaa/agenda-go/tools"
	"github.com/spf13/cobra"
	"github.com/Suenaa/agenda-go/logs"
)

// cancelCmd represents the cancel command
var cancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: "cancel a meeting",
	Long:  `cancel a meeting`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		if title == "" {
			tools.Report(errors.New("title required"))
		}
		if err := service.DeleteMeeting(title); err == nil {
			fmt.Println("Success")
			logs.EventLog("cancel a meeting: " + title)
		} else {
			tools.Report(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(cancelCmd)

	cancelCmd.Flags().StringP("title", "t", "", "title of the meeting")
}

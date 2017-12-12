package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/Suenaa/agenda-go/service"
	"github.com/Suenaa/agenda-go/tools"
	"github.com/spf13/cobra"
	"github.com/Suenaa/agenda-go/logs"
)

// dpCmd represents the dp command
var dpCmd = &cobra.Command{
	Use:   "dp",
	Short: "delete participants in a meeting",
	Long:  `delete participants in a meeting`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		participants, _ := cmd.Flags().GetStringSlice("participant")
		if title == "" {
			tools.Report(errors.New("title required"))
		}
		if participants == nil || len(participants) == 0 {
			tools.Report(errors.New("participant(s) required"))
		}
		noError := true
		for _, one := range participants {
			if err := service.DeleteParticipator(title, one); err != nil {
				noError = false
				logs.ErrLog(err)
				fmt.Fprintln(os.Stderr, err)
			}
		}
		if noError {
			fmt.Println("Success")
			logs.EventLog("delete participant(s) in the meeting: " + title)
		}
	},
}

func init() {
	RootCmd.AddCommand(dpCmd)

	dpCmd.Flags().StringP("title", "t", "", "title of the meeting")
	dpCmd.Flags().StringSliceP("participant", "p", nil, "participants want to remove")
}

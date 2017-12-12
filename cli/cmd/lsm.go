package cmd

import (
	"errors"
	
	"github.com/Suenaa/agenda-go/service"
	"github.com/Suenaa/agenda-go/tools"
	"github.com/spf13/cobra"
	"github.com/Suenaa/agenda-go/logs"
)

// lsmCmd represents the lsm command
var lsmCmd = &cobra.Command{
	Use:   "lsm",
	Short: "list all meetings during a period",
	Long:  `list all meetings during a period`,
	Run: func(cmd *cobra.Command, args []string) {
		start, _ := cmd.Flags().GetString("start")
		end, _ := cmd.Flags().GetString("end")
		if start == "" {
			tools.Report(errors.New("start required"))
		}
		if end == "" {
			tools.Report(errors.New("end required"))
		}
		meetings, err := service.QueryMeeting(start, end)
		if err != nil {
			tools.Report(err)
		}
		for _, m := range meetings {
			m.String()
		}
		logs.EventLog("list meetings during " + start + " - " + end)
	},
}

func init() {
	RootCmd.AddCommand(lsmCmd)

	lsmCmd.Flags().StringP("start", "s", "", "start time")
	lsmCmd.Flags().StringP("end", "e", "", "end time")
}

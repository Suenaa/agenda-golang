package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	
	"github.com/Suenaa/agenda-golang/service/tools"
	"github.com/spf13/cobra"
	"github.com/Suenaa/agenda-golang/service/logs"
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
		res, err := http.Get(host + "/meetings/allmeetings?startdate=" +
			start + "&enddate=" + end)	
		if err != nil {
			tools.Report(err)
		} else {
			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)
			tools.Report(err)
			fmt.Println(string(body))
		}
		
		logs.EventLog("list meetings during " + start + " - " + end)
	},
}

func init() {
	RootCmd.AddCommand(lsmCmd)

	lsmCmd.Flags().StringP("start", "s", "", "start time")
	lsmCmd.Flags().StringP("end", "e", "", "end time")
}

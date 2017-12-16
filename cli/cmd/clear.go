package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/Suenaa/agenda-golang/service/tools"
	"github.com/Suenaa/agenda-golang/service/logs"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "clear all meetings you create",
	Long:  `clear all meetings you create`,
	Run: func(cmd *cobra.Command, args []string) {
		req, err := http.NewRequest(http.MethodDelete,host+"/meetings/deletemeetings", nil)
		tools.Report(err)
		client := &http.Client{}
		res, err1 := client.Do(req)
		if err1 == nil {
			fmt.Println("Success")
			defer res.Body.Close()
		} else {
			tools.Report(err)
			logs.EventLog("clear all meetings")
		}
	},
}

func init() {
	RootCmd.AddCommand(clearCmd)
}

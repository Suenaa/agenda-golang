package cmd

import (
	"errors"
	"net/http"
	"fmt"

	"github.com/Suenaa/agenda-golang/service/tools"
	"github.com/spf13/cobra"
	"github.com/Suenaa/agenda-golang/service/logs"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "delete current account",
	Long:  `delete current account`,
	Run: func(cmd *cobra.Command, args []string) {
		password, _ := cmd.Flags().GetString("password")
		if password == "" {
			tools.Report(errors.New("password required"))
		}
		req, err := http.NewRequest(http.MethodDelete,host+"/users/deleteuser?password="+password, nil)
		tools.Report(err)
		client := &http.Client{}
		res, err1 := client.Do(req)
		if err1 == nil {
			fmt.Println("Success")
			logs.EventLog("delete a user")
			defer res.Body.Close()
		} else {
			fmt.Println("no Success")
			tools.Report(err1)
		}
	},
}

func init() {
	RootCmd.AddCommand(delCmd)

	delCmd.Flags().StringP("password", "p", "", "your password")
}

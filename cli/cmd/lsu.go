package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	
	"github.com/spf13/cobra"
	"github.com/Suenaa/agenda-golang/service/tools"
	"github.com/Suenaa/agenda-golang/service/logs"
)

// lsuCmd represents the lsu command
var lsuCmd = &cobra.Command{
	Use:   "lsu",
	Short: "list all users",
	Long:  `list all users`,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := http.Get(host + "/users/allusers")
		tools.Report(err)
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		tools.Report(err)
		var data []map[string]interface{}
		err = json.Unmarshal(body, &data)
		tools.Report(err)
		fmt.Println(string(body))
		logs.EventLog("list all users")
	},
}

func init() {
	RootCmd.AddCommand(lsuCmd)
}

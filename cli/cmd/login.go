package cmd

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Suenaa/agenda-golang/service/tools"
	"github.com/spf13/cobra"
	"github.com/Suenaa/agenda-golang/service/logs"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "log in",
	Long:  `log in agenda`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		if username == "" {
			tools.Report(errors.New("username required"))
		}
		if password == "" {
			tools.Report(errors.New("password required"))
		}
		res, err := http.Get(host + "/user/login?username=" +
			username + "&password=" + password)	
		if err == nil {
			fmt.Println("Success")
			logs.EventLog(username + " log in")
			defer res.Body.Close()
		} else {
			tools.Report(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringP("username", "u", "", "your username")
	loginCmd.Flags().StringP("password", "p", "", "your password")
}

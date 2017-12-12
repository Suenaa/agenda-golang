package cmd

import (
	"errors"
	"fmt"

	"github.com/Suenaa/agenda-go/service"
	"github.com/Suenaa/agenda-go/tools"
	"github.com/spf13/cobra"
	"github.com/Suenaa/agenda-go/logs"
)

// registCmd represents the regist command
var registCmd = &cobra.Command{
	Use:   "regist",
	Short: "regist a new user",
	Long:  `regist a new user`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		phone, _ := cmd.Flags().GetString("telephone")
		if username == "" {
			tools.Report(errors.New("username required"))
		}
		if password == "" {
			tools.Report(errors.New("password required"))
		}
		if email == "" {
			tools.Report(errors.New("email required"))
		}
		if phone == "" {
			tools.Report(errors.New("phone required"))
		}
		err := service.UserRegister(username, password, email, phone)
		if err == nil {
			fmt.Println("Success")
			logs.EventLog(username + " regists")

		} else {
			tools.Report(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(registCmd)

	registCmd.Flags().StringP("username", "u", "", "the username you want")
	registCmd.Flags().StringP("password", "p", "", "the password you want")
	registCmd.Flags().StringP("email", "e", "", "your email address")
	registCmd.Flags().StringP("telephone", "t", "", "your telephone number")
}

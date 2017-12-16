package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"errors"

	"github.com/Suenaa/agenda-golang/service/tools"
	"github.com/spf13/cobra"
	"github.com/Suenaa/agenda-golang/service/logs"
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
		data := struct {
			Username string `json:"username"`
			Password string `json:"password"`
			Email    string `json:"email"`
			Phone    string `json:"phone"`
		}{username, password, email, phone}
		buf, err := json.Marshal(data)
		tools.Report(err)
		res, err := http.Post(host+"/user/register",
			"application/json", bytes.NewBuffer(buf))
		if err == nil {
			fmt.Println("Success")
			logs.EventLog(username + " regists")
			defer res.Body.Close()
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

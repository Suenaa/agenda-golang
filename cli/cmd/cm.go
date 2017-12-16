package cmd

import (
	"errors"
	"bytes"
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/spf13/cobra"

	"github.com/Suenaa/agenda-golang/service/tools"
	"github.com/Suenaa/agenda-golang/service/logs"
)

// cmCmd represents the cm command
var cmCmd = &cobra.Command{
	Use:   "cm",
	Short: "create a meeting",
	Long:  `create a meeting`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		participants, _ := cmd.Flags().GetStringSlice("participant")
		start, _ := cmd.Flags().GetString("start")
		end, _ := cmd.Flags().GetString("end")
		if title == "" {
			tools.Report(errors.New("title required"))
		}
		if participants == nil || len(participants) == 0 {
			tools.Report(errors.New("participant(s) required"))
		}
		if start == "" {
			tools.Report(errors.New("start required"))
		}
		if end == "" {
			tools.Report(errors.New("end required"))
		}
		err := service.CreateMeeting(title, start, end, participants)
		data := struct {
			Title     string   `json:"title"`
			Participators   []string `json:"members"`
			Start string   `json:"starttime"`
			End   string   `json:"endtime"`
		}{title, participants, start, end}
		buf, err := json.Marshal(data)
		tools.Report(err)
		res, err := http.Post(host+"/meetings/newmeeting",
			"application/json", bytes.NewBuffer(buf))
		defer res.Body.Close()
		if err == nil {
			fmt.Println("Success")
			logs.EventLog("create a meeting: " + title)
		} else {
			tools.Report(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(cmCmd)

	cmCmd.Flags().StringP("title", "t", "", "title of the meeting")
	cmCmd.Flags().StringSliceP("participant", "p", nil, "participants of the meeting")
	cmCmd.Flags().StringP("start", "s", "", "when to start the meeting")
	cmCmd.Flags().StringP("end", "e", "", "when to end the meeting")
}

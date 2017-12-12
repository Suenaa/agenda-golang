package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/Suenaa/agenda-golang/service/logs"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "agd",
	Short: "Agenda-go",
	Long:  "A tool for managing meetings",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		logs.ErrLog(err)
		os.Exit(1)
	}
}

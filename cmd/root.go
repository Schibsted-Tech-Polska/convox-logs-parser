package cmd

import (
	"fmt"
	"os"

	"github.com/Schibsted-Tech-Polska/convox-logs-parser/pkg/convox"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
	"strconv"
)

var app string
var rack string
var since string
var filter string
var follow bool

var convoxFormatter convox.Formatter

var rootCmd = &cobra.Command{
	Use:   "convox-logs-parser",
	Short: "Convox logs formatter",
	Long:  `Convox logs formatter.`,
	RunE:  cmdRoot,
}

// Execute is a command required by cobra package in order to run application
// It runs CLI parser and handles further actions
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&app, "app", "a", "", "app name inferred from current directory if not specified")
	rootCmd.PersistentFlags().StringVarP(&rack, "rack", "r", "", "rack name")
	rootCmd.PersistentFlags().StringVarP(&since, "since", "s", "30s", "show logs since a duration (e.g. 10m or 1h2m10s)")
	rootCmd.PersistentFlags().StringVar(&filter, "filter", "", "filter the logs by a given token")
	rootCmd.PersistentFlags().BoolVarP(&follow, "follow", "f", true, "keep streaming new log output")

}

func cmdRoot(cmd *cobra.Command, args []string) error {
	var am []string
	am = append(am, "logs")
	am = append(am, "--follow="+strconv.FormatBool(follow))
	am = append(am, "--since="+since)
	if rack != "" {
		am = append(am, "--rack="+rack)
	}
	if app != "" {
		am = append(am, "--app="+app)
	}

	comm := exec.Command("convox", am...)

	comm.Stdout = &convoxFormatter

	if err := comm.Run(); err != nil {
		log.Fatal(err)
	}

	return nil
}

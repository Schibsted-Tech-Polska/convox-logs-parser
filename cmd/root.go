package cmd

import (
	"fmt"
	"os"

	"github.com/radekl/convox-json-logs/pkg/convox"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

var app string
var rack string
var since string
var filter string
var follow bool

var convoxFormatter convox.Formatter

var rootCmd = &cobra.Command{
	Use:   "convox-json-logs",
	Short: "Convox logs formatter",
	Long:  `Convox logs formatter.`,
	RunE:  cmdRoot,
}

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
	rootCmd.PersistentFlags().BoolVarP(&follow, "follow", "f", false, "keep streaming new log output")

}

func cmdRoot(cmd *cobra.Command, args []string) error {
	var aMap = []string{app, rack, since, filter}
	fmt.Println("Got following params: " + strings.Join(aMap, ", "))

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

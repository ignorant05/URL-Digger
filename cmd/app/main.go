package main

import (
	"fmt"
	"os"
	htmlparser "web-crawler/internal/HTMLparser"

	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
)

var (
	target string
	output map[string][]string
	used   map[string]bool
)

func Init() {
	output = make(map[string][]string)
	used = make(map[string]bool)
}

const version = "1.0.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of the CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version:", version)
	},
}
var rootCmd = &cobra.Command{
	Use:   "udig",
	Short: "Simple URL web-crawler",
	Run: func(cmd *cobra.Command, args []string) {
		url := fmt.Sprintf("%s", target)
		htmlparser.HTMLParser(&used, &output, url)

	},
}

func Execute() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.PersistentFlags().StringVar(
		&target, "target", "", "target")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Init()
	Execute()
	fmt.Println(len(output))
	fmt.Println(len(used))

	t := table.NewWriter()
	t.SetAutoIndex(true)
	t.SetStyle(table.StyleColoredDark)
	t.Style().Options.SeparateColumns = false
	t.AppendHeader(table.Row{"#", "URL", "Children"})
	for key, val := range output {
		t.AppendRow(table.Row{key, ""})
		for _, v := range val {
			t.AppendRow(table.Row{"", v})
		}
	}

	fmt.Println(t.Render())
}

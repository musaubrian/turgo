package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// topicsCmd represents the topics command
var topicsCmd = &cobra.Command{
	Use:   "topics",
	Short: "List all topics",
	Run: func(cmd *cobra.Command, args []string) {
		vals, err := t.ListTopics()
		if err != nil {
			log.Fatal(err)
		}
		t.Tb.AddHeader("#", "TOPICS")
		for i, v := range vals {
			t.Tb.AddLine(i, v)
		}
		t.Tb.Print()

	},
}

func init() {
	rootCmd.AddCommand(topicsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// topicsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// topicsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

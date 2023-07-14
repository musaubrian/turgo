package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List contents of your collection",
	/* Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`, */
	Run: func(cmd *cobra.Command, args []string) {
		topics, URLs, notes, err := t.List()
		if err != nil {
			log.Fatal(err)
		}

		t.Tb.AddHeader("#", "TOPIC", "NOTE", "URL")
		for i := 0; i < len(topics); i++ {
			if len(notes[i]) < 1 {
				t.Tb.AddLine(i, topics[i], "****", URLs[i])
			} else {
				t.Tb.AddLine(i, topics[i], notes[i], URLs[i])
			}
		}
		t.Tb.Print()
	},
}

var filterByTopicCmd = &cobra.Command{
	Use:     "topic",
	Short:   "Filter the results based on the topic",
	Example: "tursgo list topic topic_name",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("NO TOPIC PARSED")
		} else if len(args) > 1 {
			log.Fatal("TOO MANY ARGUMENTS.\nTOPIC MUST BE ONE WORD")
		}
		topic := args[0]
		urls, notes, err := t.FilterByTopic(topic)
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < len(urls); i++ {
			t.Tb.AddHeader("#", "NOTE", "URL")
			if len(notes[i]) < 1 {
				t.Tb.AddLine(i, "****", urls[i])
			} else {
				t.Tb.AddLine(i, notes[i], urls[i])
			}
		}
		t.Tb.Print()

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.AddCommand(filterByTopicCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

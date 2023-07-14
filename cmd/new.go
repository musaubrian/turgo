package cmd

import (
	"log"

	"github.com/musaubrian/turgo/internal/data"
	"github.com/musaubrian/turgo/internal/util"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:     "new",
	Short:   "Create a new bookmark",
	Example: "tursgo new topic_name",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("NO TOPIC PARSED")
		} else if len(args) > 1 {
			log.Fatal("TOO MANY ARGUMENTS.\nTOPICS MUST BE ONE WORD")
		}
		u := util.GetInput("URL")
		note := util.GetInput("A little note")
		c := &data.Collection{
			ID:    util.GenerateULID(),
			Topic: args[0],
			URL:   u,
			Note:  note,
		}

		if err := t.CreateCollection(*c); err != nil {
			log.Fatal(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

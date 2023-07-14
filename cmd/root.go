package cmd

import (
	"log"
	"os"

	"github.com/musaubrian/turgo/internal/data"
	"github.com/spf13/cobra"
)

var t *data.Turgo

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tursgo",
	Short: "CLI bookmark",
	Long: `
Save site URLs you want to visit later

Organize them using topics.
Since you'll have your own db instance, you can extend it to a web app if you wish

This uses turso(https://turso.tech) for hosting the db.

To get set up, follow the steps here -> https://docs.turso.tech/tutorials/get-started-turso-cli/
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	db, err := data.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	t = data.InitTurgo(db)
	if err := t.InitTables(); err != nil {
		log.Fatal(err)
	}
	err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tursgo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

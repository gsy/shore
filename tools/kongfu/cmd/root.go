package cmd

import (
	"os"

	"github.com/gsy/store/tools/kongfu/model"
	"github.com/spf13/cobra"
)

// kongfuCmd represents the base command when called without any subcommands
var kongfuCmd = &cobra.Command{
	Use:   "kongfu",
	Short: "Kongfu is a code generate tool",
	Long: `:Kongfu is a code generate tool

Kongfu is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := kongfuCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.store.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	kongfuCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	kongfuCmd.AddCommand(model.Cmd)
}

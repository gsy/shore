/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/gsy/store/pkg/user/infrastructure/database"
	"github.com/gsy/store/pkg/user/infrastructure/database/model"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// gengormCmd represents the gengorm command
var gengormCmd = &cobra.Command{
	Use:   "gengorm",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gengorm called")

		// 1. create database via config
		db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		// 2. create generator
		g := gen.NewGenerator(gen.Config{
			OutPath: "/Users/tiger/code/store/pkg/user/infrastructure/database/dal/query",
			Mode:    gen.WithDefaultQuery,
		})
		g.UseDB(db)

		// 3. apply & execute
		g.ApplyBasic(model.User{})

		g.ApplyInterface(func(database.UserRepository) {}, model.User{})

		g.Execute()
	},
}

func init() {
	rootCmd.AddCommand(gengormCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gengormCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gengormCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

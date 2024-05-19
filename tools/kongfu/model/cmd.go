package model

import (
	"fmt"

	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var (
	Source  string
	DSN     string
	OutPath string
	Table   string

	Cmd = &cobra.Command{
		Use:   "model gen [OPTIONS]",
		Short: "orm model generate tool",
		Long:  `orm mdoel generate tool`,
	}

	genCmd = &cobra.Command{
		Use:   "gen",
		Short: "generate model using gorm.gen",
		Long:  `generate model using gorm.gen`,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			fmt.Printf("generate models: source: %s, dsn: %s, outpath: %s\n", Source, DSN, OutPath)

			db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
			if err != nil {
				return (err)
			}

			g := gen.NewGenerator(gen.Config{
				OutPath: OutPath,
				Mode:    gen.WithDefaultQuery,
			})
			g.UseDB(db)

			// 3. apply & execute
			model := g.GenerateModel(Table)
			g.ApplyBasic(model)

			// TODO: generate custome method
			// g.ApplyInterface(func(database.UserRepository) {}, model.User{})

			g.Execute()
			return nil
		},
	}
)

func init() {
	genCmd.Flags().StringVarP(&Source, "source", "s", "", "source path for ddl file")
	genCmd.Flags().StringVarP(&DSN, "dsn", "", "", "dsn for database")
	genCmd.Flags().StringVarP(&OutPath, "out", "o", "", "output path")
	genCmd.Flags().StringVarP(&Table, "table", "t", "", "table")
	Cmd.AddCommand(genCmd)
}

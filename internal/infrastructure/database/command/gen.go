package command

// import (
// 	"github.com/gsy/store/pkg/user/adapters/repository"
// 	"github.com/spf13/cobra"
// 	"gorm.io/driver/sqlite"
// 	"gorm.io/gen"
// 	"gorm.io/gorm"
// )

// func Generate(_ *cobra.Command, args []string) (err error) {
// 	// 1. create database via config
// 	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
// 	if err != nil {
// 		panic(err)
// 	}

// 	// 2. create generator
// 	g := gen.NewGenerator(gen.Config{
// 		OutPath: args[0],
// 		Mode:    gen.WithDefaultQuery,
// 	})
// 	g.UseDB(db)

// 	// 3. apply & execute
// 	g.ApplyBasic(model.User{})

// 	g.ApplyInterface(func(database.UserRepository) {}, model.User{})

// 	g.Execute()

// 	return nil
// }

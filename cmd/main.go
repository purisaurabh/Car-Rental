package main

import (
	"fmt"
	"os"

	"github.com/purisaurabh/car-rental/internal/config"
	db "github.com/purisaurabh/car-rental/internal/database"
	"github.com/purisaurabh/car-rental/internal/server"
	"github.com/urfave/cli"
	"go.uber.org/zap"
)

func main() {

	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Printf("Error initializing logger: %v\n", err)
		os.Exit(1)
	}

	// defer the sync call to flush the buffer before the application exits
	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	config.Load()
	cliApp := cli.NewApp()
	cliApp.Name = "Boilerplate Code"
	cliApp.Version = "1.0.0"

	cliApp.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "start_server",
			Action: func(c *cli.Context) error {
				server.StartApiServer()
				return nil
			},
		},
		{
			Name:  "create-migration",
			Usage: "create migration file",
			Action: func(c *cli.Context) error {
				return db.CreateMigrationFile(c.Args().Get(0))
			},
		},
		{
			Name:  "migrate",
			Usage: "run db migration",
			Action: func(c *cli.Context) error {
				return db.RunMigrations()
			},
		},
		{
			Name:  "rollback",
			Usage: "rollback db micgration",
			Action: func(c *cli.Context) error {
				return db.RollbackMigration(c.Args().Get(0))
			},
		},
	}

	if err := cliApp.Run(os.Args); err != nil {
		fmt.Println("Error while running the cli app : ", err)
		panic(err)
	}

}

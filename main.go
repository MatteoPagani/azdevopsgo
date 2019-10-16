package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

var definitions []GeneralStruct
var projects []Project

var app = cli.NewApp()

var configuration Configuration

func main() {

	info()
	commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func info() {
	app.Name = "devops"
	app.Usage = "Get Azure DevOps last builds"
	app.Author = "Matteo Pagani"
	app.Version = "1.0.0"
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "setup",
			Aliases: []string{"s"},
			Usage:   "Setup your Azure Devops project with credentials and default pipeline",
			Action: func(c *cli.Context) {
				askForConfigurationAndSave() // Ask credentials and save them into config.json file
			},
		},
		{
			Name:    "releases",
			Aliases: []string{"r"},
			Usage:   "Get latests releases of current project definition",
			Action: func(c *cli.Context) {
				getLatestReleases()
			},
		},
		{
			Name:    "builds",
			Aliases: []string{"b"},
			Usage:   "Get latests builds of current project definition",
			Action: func(c *cli.Context) {
				readConfigurationFile() // Reads the config.json file and store into configuration var

				if configuration.Project == "" {
					fmt.Println("You must run setup command first")
					return
				}

				getLatestBuilds() // Print in CLI the last N builds of the definition
			},
		},
	}
}

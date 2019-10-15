package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

var definitions []GeneralStruct
var projects []Project

var configuration Configuration

var project string
var definitionId int

func main() {

	app := cli.NewApp()
	app.Name = "devops"
	app.Usage = "Get Azure DevOps last builds"
	app.Action = func(c *cli.Context) error {

		initConfig()

		if configuration.Project.Name == "" || configuration.Project.Definition == 0 {
			askForProject()
		} else {
			project = configuration.Project.Name
			definitionId = configuration.Project.Definition
		}

		getBuildsOfDefinition()

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

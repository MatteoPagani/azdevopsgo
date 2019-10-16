package main

import (
	"log"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/urfave/cli"
)

var buildDefinitions []GeneralStruct
var releaseDefinitions []GeneralStruct

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
				if configurationFileExists() {
					prompt := promptui.Select{
						Label: "Seems like a config file already exists. Do you want to overwrite it?",
						Items: []string{"Yes", "No"},
					}
					_, value, _ := prompt.Run()
					if value == "Yes" {
						askForConfigurationAndSave() // Ask credentials and save them into config.json file
					}
				}
			},
		},
		{
			Name:    "deployments",
			Aliases: []string{"d"},
			Usage:   "Get latest deployments of current project definition",
			Action: func(c *cli.Context) {
				if configurationIsValid() {
					getLatestDeployments()
				}
			},
		},
		{
			Name:    "builds",
			Aliases: []string{"b"},
			Usage:   "Get latest builds of current project definition",
			Action: func(c *cli.Context) {
				if configurationIsValid() {
					getLatestBuilds()
				}
			},
		},
	}
}

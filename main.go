package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

var definitions []GeneralStruct
var projects []Project

var configuration Configuration

var project string
var definitionId int

func initConfig() error {

	file, err := os.Open("config.json")
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		return err
	}

	return nil
}

func askForProject() {
	projectNames := getProjectNames()
	prompt := promptui.Select{
		Label: "Which project?",
		Items: projectNames,
	}

	_, selectedProject, err := prompt.Run()

	project = selectedProject

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	definitionsNames := getProjectDefinitions(project)
	prompt = promptui.Select{
		Label: "Which pipeline?",
		Items: definitionsNames,
	}

	_, definitionChoosed, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	for _, element := range definitions {
		if element.Name == definitionChoosed {
			definitionId = element.Id
		}
	}
}

func main() {

	initConfig()

	if configuration.Project.Name == "" || configuration.Project.Definition == 0 {
		askForProject()
	} else {
		project = configuration.Project.Name
		definitionId = configuration.Project.Definition
	}

	getBuildsOfDefinition()
}

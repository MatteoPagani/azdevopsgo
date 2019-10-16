package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/manifoldco/promptui"
)

func askForSettingAsString(label string) string {
	prompt := promptui.Prompt{
		Label: fmt.Sprintf("Type the %s", label),
	}

	value, _ := prompt.Run()
	return value
}

func askForAnItem(label string, items []string) string {
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}
	_, value, _ := prompt.Run()
	return value
}

func askForConfigurationAndSave() {
	configuration = Configuration{}

	configuration.Username = askForSettingAsString("username")
	configuration.Password = askForSettingAsString("password")
	configuration.Organization = askForSettingAsString("organization name")
	configuration.ApiVersion = askForSettingAsString("API version to use (default 5.1)")
	if configuration.ApiVersion == "" {
		configuration.ApiVersion = "5.1"
	}

	configuration.Project = askForAnItem("Select a project", getProjectNames())

	buildDefinitionName := askForAnItem("Select a build definition", getProjectDefinitions(configuration.Project))
	configuration.BuildDefinition = getDefinitionByName(buildDefinitionName, buildDefinitions).Id

	releaseDefinitionName := askForAnItem("Select a release definition", getProjectReleaseDefinitions(configuration.Project))
	configuration.ReleaseDefinition = getDefinitionByName(releaseDefinitionName, releaseDefinitions).Id

	writeConfigFile(configuration)
}

func writeConfigFile(configuration Configuration) {
	file, _ := json.MarshalIndent(configuration, "", " ")
	_ = ioutil.WriteFile("devops.config.json", file, 0644)
}

func readConfigurationFile() {
	file, _ := ioutil.ReadFile("devops.config.json")

	configuration = Configuration{}

	_ = json.Unmarshal([]byte(file), &configuration)

	if configuration.Project == "" {
		fmt.Println("In order to get builds you must run the setup command first.")
		return
	}
}

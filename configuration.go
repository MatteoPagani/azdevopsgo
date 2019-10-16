package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/manifoldco/promptui"
)

var configFileName = "devops.config.json"

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

func configurationIsValid() bool {
	readConfigurationFile() // Reads the config.json file and store into configuration var

	if configuration.Project == "" {
		fmt.Println("You must run setup command first")
		return false
	}

	return true
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
	_ = ioutil.WriteFile(configFileName, file, 0644)
}

func readConfigurationFile() {
	file, _ := ioutil.ReadFile(configFileName)

	configuration = Configuration{}

	_ = json.Unmarshal([]byte(file), &configuration)

	if configuration.Project == "" {
		fmt.Println("In order to get builds you must run the setup command first.")
		return
	}
}

func configurationFileExists() bool {
	info, err := os.Stat(configFileName)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/manifoldco/promptui"
)

// Name of the configuration file to be stored in the DevOps project folder
var configFileName = "devops.config.json"

// Prompt the user a string setting
func askForSettingAsString(label string) string {
	prompt := promptui.Prompt{
		Label: fmt.Sprintf("Type the %s", label),
	}

	value, _ := prompt.Run()
	return value
}

// Prompt with a fancy select a list of items
func askForAnItem(label string, items []string) string {
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}
	_, value, _ := prompt.Run()
	return value
}

// Check if configuration file is valid (project is not empty)
func configurationIsValid() bool {
	readConfigurationFile() // Reads the config.json file and store into configuration var

	if configuration.Project == "" {
		fmt.Println("You must run setup command first")
		return false
	}

	return true
}

// Ask the user some settings to be stored into the configuration file
// It will be used to make calls to Azure Devops REST API
func askForConfigurationAndSave() {
	configuration = Configuration{}

	configuration.Username = askForSettingAsString("username")
	configuration.Password = askForSettingAsString("password")
	configuration.Organization = askForSettingAsString("organization name")
	configuration.ApiVersion = askForSettingAsString("API version to use (default 5.1)")

	if configuration.ApiVersion == "" {
		configuration.ApiVersion = "5.1"
	}

	configuration.Project = askForAnItem(
		"Select a project",
		getProjectNames(),
	)

	configuration.BuildDefinition = getDefinitionByName(
		askForAnItem(
			"Select a build definition",
			getBuildProjectDefinitions(configuration.Project),
		),
		buildDefinitions).Id

	configuration.ReleaseDefinition = getDefinitionByName(
		askForAnItem(
			"Select a release definition",
			getProjectReleaseDefinitions(configuration.Project),
		),
		releaseDefinitions).Id

	writeConfigFile(configuration)
}

// Simply write a Configuration struct into a configuration json file
func writeConfigFile(configuration Configuration) {
	file, _ := json.MarshalIndent(configuration, "", " ")
	_ = ioutil.WriteFile(configFileName, file, 0644)
}

// Read a the configuration file
// and store settings into a global Configuration var
func readConfigurationFile() {
	file, _ := ioutil.ReadFile(configFileName)

	configuration = Configuration{}

	_ = json.Unmarshal([]byte(file), &configuration)

	if configuration.Project == "" {
		prompt := promptui.Select{
			Label: "You must run the setup command first. Do you want to run it now?",
			Items: []string{"Yes", "No"},
		}
		_, value, _ := prompt.Run()

		if value == "Yes" {
			askForConfigurationAndSave()
		}

		return
	}
}

// Simply check if configuration file exists in the DevOps project folder
func configurationFileExists() bool {
	info, err := os.Stat(configFileName)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

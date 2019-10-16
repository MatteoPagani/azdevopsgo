package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/manifoldco/promptui"
)

func askForConfigurationSetting(label string) string {
	prompt := promptui.Prompt{
		Label: fmt.Sprintf("Type the %s", label),
	}

	value, _ := prompt.Run()
	return value
}

func askForConfigurationAndSave() {
	configuration = Configuration{}

	configuration.Organization = askForConfigurationSetting("organization name")
	configuration.Project = askForConfigurationSetting("project name")
	configuration.ApiVersion = askForConfigurationSetting("API version to use (default 5.1)")
	if configuration.ApiVersion == "" {
		configuration.ApiVersion = "5.1"
	}
	configuration.Username = askForConfigurationSetting("username")
	configuration.Password = askForConfigurationSetting("password")
	configuration.Definition, _ = strconv.Atoi(askForConfigurationSetting("definition ID"))

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

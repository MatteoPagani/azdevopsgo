package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/jedib0t/go-pretty/table"
)

func call(endpoint string) string {
	var username string = configuration.Username
	var passwd string = configuration.Password
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", configuration.BaseUrl, endpoint), nil)
	req.SetBasicAuth(username, passwd)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)

	return s
}

func getProjectNames() []string {
	projectsResult := call(fmt.Sprintf("_apis/projects?api-version=%s", configuration.ApiVersion))
	var res ProjectsResponse
	json.Unmarshal([]byte(projectsResult), &res)

	projects = res.Value

	var projectNames []string

	for _, element := range res.Value {
		projectNames = append(projectNames, element.Name)
	}

	return projectNames
}

func getProjectDefinitions(project string) []string {
	endpoint := fmt.Sprintf("%s/_apis/build/definitions?api-version=%s", project, configuration.ApiVersion)
	buildsResult := call(endpoint)

	var definitionResponse DefinitionsResponse
	json.Unmarshal([]byte(buildsResult), &definitionResponse)

	var definitionsNames []string

	definitions = definitionResponse.Value

	for _, element := range definitions {
		definitionsNames = append(definitionsNames, element.Name)
	}

	return definitionsNames
}

func getBuildsOfDefinition() {

	fmt.Println()
	fmt.Println(fmt.Sprintf("Getting builds of project %s and definition %d", project, definitionId))
	fmt.Println()

	endpoint := fmt.Sprintf("%s/_apis/build/builds?definitions=%d&api-version=%s", project, definitionId, configuration.ApiVersion)
	result := call(endpoint)

	var response BuildsResponse
	json.Unmarshal([]byte(result), &response)

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Started", "Finished", "Branch"})

	i := 0
	for _, element := range response.Value {
		t.AppendRow([]interface{}{element.Id, element.StartTime, element.FinishTime, element.SourceBranch})
		i = i + 1
		if i > 10 {
			break
		}
	}

	t.Render()
}

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

	client := &http.Client{}

	endp := fmt.Sprintf("https://dev.azure.com/%s/%s", configuration.Organization, endpoint)
	req, err := http.NewRequest("GET", endp, nil)
	req.SetBasicAuth(configuration.Username, configuration.Password)
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

	definitions := definitionResponse.Value

	for _, element := range definitions {
		definitionsNames = append(definitionsNames, element.Name)
	}

	return definitionsNames
}

func getLatestReleases() {
	fmt.Println()
	fmt.Println(fmt.Sprintf("Getting releases of project %s and definition %d", configuration.Project, configuration.Definition))
	fmt.Println()

	endpoint := fmt.Sprintf("%s/_apis/release/releases?definitions=%d&api-version=%s", configuration.Project, configuration.Definition, configuration.ApiVersion)
	result := call(endpoint)

	fmt.Println(result)
}

func getLatestBuilds() {

	fmt.Println()
	fmt.Println(fmt.Sprintf("Getting builds of project %s and definition %d", configuration.Project, configuration.Definition))
	fmt.Println()

	endpoint := fmt.Sprintf("%s/_apis/build/builds?definitions=%d&api-version=%s", configuration.Project, configuration.Definition, configuration.ApiVersion)
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

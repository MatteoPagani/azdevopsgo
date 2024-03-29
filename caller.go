package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Call Azure DevOps REST API with VSRM endpoint prefix
func callVsrm(endpoint string) string {
	endp := fmt.Sprintf("https://vsrm.dev.azure.com/%s/%s/%s", configuration.Organization, configuration.Project, endpoint)
	return base(endp)
}

// Call Azure DevOps REST API
func call(endpoint string) string {
	endp := fmt.Sprintf("https://dev.azure.com/%s/%s", configuration.Organization, endpoint)
	return base(endp)
}

// Base function to make REST API calls
func base(endpoint string) string {
	client := &http.Client{}

	req, err := http.NewRequest("GET", endpoint, nil)
	req.SetBasicAuth(configuration.Username, configuration.Password)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)

	return s
}

// Get project names of organization
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

// Get Definition struct by name
func getDefinitionByName(name string, definitions []GeneralStruct) GeneralStruct {
	for _, element := range definitions {
		if element.Name == name {
			return element
		}
	}

	return GeneralStruct{}
}

// Get an array of BuildChange by build ID
func getBuildChangesById(id string) []BuildChange {
	endpoint := fmt.Sprintf("%s/_apis/build/builds/%s/changes?api-version=%s", configuration.Project, id, configuration.ApiVersion)
	buildResult := call(endpoint)

	var res BuildChangeResponse
	json.Unmarshal([]byte(buildResult), &res)

	return res.Value
}

// Get release definitions of a project
func getProjectReleaseDefinitions(project string) []string {
	endpoint := fmt.Sprintf("_apis/release/definitions?api-version=%s", configuration.ApiVersion)
	buildsResult := callVsrm(endpoint)

	var definitionResponse DefinitionsResponse
	json.Unmarshal([]byte(buildsResult), &definitionResponse)

	var definitionsNames []string

	releaseDefinitions = definitionResponse.Value

	for _, element := range releaseDefinitions {
		definitionsNames = append(definitionsNames, element.Name)
	}

	return definitionsNames
}

// Get build definitions of a project
func getBuildProjectDefinitions(project string) []string {
	endpoint := fmt.Sprintf("%s/_apis/build/definitions?api-version=%s", project, configuration.ApiVersion)
	buildsResult := call(endpoint)

	var definitionResponse DefinitionsResponse
	json.Unmarshal([]byte(buildsResult), &definitionResponse)

	var definitionsNames []string

	buildDefinitions = definitionResponse.Value

	for _, element := range buildDefinitions {
		definitionsNames = append(definitionsNames, element.Name)
	}

	return definitionsNames
}

func getLatestDeployments() []Deployment {
	fmt.Println()
	fmt.Println(fmt.Sprintf("Getting deployments of project %s and definition %d", configuration.Project, configuration.ReleaseDefinition))
	fmt.Println()

	endpoint := fmt.Sprintf("_apis/release/deployments?definitionId=%d&api-version=%s", configuration.ReleaseDefinition, configuration.ApiVersion)
	result := callVsrm(endpoint)

	var response ReleasesResponse
	json.Unmarshal([]byte(result), &response)

	return response.Value
}

func getLatestBuilds() []Build {

	fmt.Println()
	fmt.Println(fmt.Sprintf("Getting builds of project %s and definition %d", configuration.Project, configuration.BuildDefinition))
	fmt.Println()

	endpoint := fmt.Sprintf("%s/_apis/build/builds?definitions=%d&api-version=%s", configuration.Project, configuration.BuildDefinition, configuration.ApiVersion)
	result := call(endpoint)

	var response BuildsResponse
	json.Unmarshal([]byte(result), &response)

	return response.Value
}

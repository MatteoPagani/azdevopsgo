package main

import (
	"os"
	"time"

	"github.com/jedib0t/go-pretty/table"
)

// Print a table of latest 10 deployments
func printDeployments(deployments []Deployment) {

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{
		"#",
		"CompletedOn",
		"Environment",
		"Status",
		"Changes",
	})

	i := 0
	for _, element := range deployments {

		buildId := element.Release.Artifacts[0].DefinitionReference.Version.Id
		changes := getBuildChangesById(buildId)

		parsedTime, _ := time.Parse(AzureTimeLayout, element.CompletedOn)

		t.AppendRow([]interface{}{
			element.Id,
			parsedTime.In(desiredLocationForDateTime).Format(DesiredTimeLayout),
			element.ReleaseEnvironment.Name,
			element.DeploymentStatus,
			changes[0].Message,
		})

		i = i + 1
		if i > 10 {
			break
		}
	}

	t.Render()
}

// Print a table of latest 10 builds
func printBuilds(builds []Build) {

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Started", "Finished", "Branch"})

	i := 0
	for _, element := range builds {
		t.AppendRow([]interface{}{element.Id, element.StartTime, element.FinishTime, element.SourceBranch})
		i = i + 1
		if i > 10 {
			break
		}
	}

	t.Render()
}

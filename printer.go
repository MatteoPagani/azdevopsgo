package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/table"
)

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

	loc, _ := time.LoadLocation("Europe/Rome")

	i := 0
	for _, element := range deployments {

		buildId := element.Release.Artifacts[0].DefinitionReference.Version.Id
		changes := getBuildChangesById(buildId)

		fmt.Println(element.CompletedOn)

		parsedTime, _ := time.Parse("2006-01-02T15:04:05Z", element.CompletedOn)

		t.AppendRow([]interface{}{
			element.Id,
			parsedTime.In(loc).Format("02/01/2006 15:04"),
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

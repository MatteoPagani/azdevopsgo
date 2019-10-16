package main

import (
	"os"

	"github.com/jedib0t/go-pretty/table"
)

func printDeployments(deployments []Deployment) {

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{
		"#",
		"StartedOn",
		"CompletedOn",
		"Environment",
		"Status",
	})

	i := 0
	for _, element := range deployments {
		t.AppendRow([]interface{}{
			element.Id,
			element.StartedOn,
			element.CompletedOn,
			element.ReleaseEnvironment.Name,
			element.DeploymentStatus,
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

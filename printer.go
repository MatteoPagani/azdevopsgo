package main

import (
	"fmt"
	"os"
	"strings"
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
	t.AppendHeader(table.Row{"#", "Started", "Finished", "Branch", "Changes"})

	i := 0
	for _, element := range builds {

		lastChangeMessage := ""
		changes := getBuildChangesById(fmt.Sprintf("%d", element.Id))
		if len(changes) > 0 {
			lastChange := changes[0]
			lastChangeMessage = lastChange.Message
		}

		branchName := element.SourceBranch
		if strings.Contains(element.SourceBranch, "/") {
			splittedBranchRefs := strings.Split(element.SourceBranch, "/")
			branchName = splittedBranchRefs[len(splittedBranchRefs)-1]
		}

		startTime := ""
		if element.StartTime != "" {
			parsedStartTime, _ := time.Parse(AzureTimeLayout, element.StartTime)
			startTime = parsedStartTime.In(desiredLocationForDateTime).Format(DesiredTimeLayout)
		}

		finishTime := ""
		if element.FinishTime != "" {
			parsedFinishTime, _ := time.Parse(AzureTimeLayout, element.FinishTime)
			finishTime = parsedFinishTime.In(desiredLocationForDateTime).Format(DesiredTimeLayout)
		}

		t.AppendRow([]interface{}{
			element.Id,
			startTime,
			finishTime,
			branchName,
			lastChangeMessage,
		})
		i = i + 1
		if i > 10 {
			break
		}
	}

	t.Render()
}

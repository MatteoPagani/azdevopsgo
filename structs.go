package main

type ProjectsResponse struct {
	Value []Project
	Count int
}

type DefinitionsResponse struct {
	Value []GeneralStruct
	Count int
}

type BuildsResponse struct {
	Count int
	Value []Build
}

type GeneralStruct struct {
	Id   int
	Name string
}

type Project struct {
	Id             string
	Name           string
	Description    string
	Url            string
	State          string
	Revision       int
	Visibility     string
	LastUpdateTime string
}

type Build struct {
	Id           int
	BuildNumber  string
	StartTime    string
	FinishTime   string
	QueueTime    string
	SourceBranch string
}

type ConfigurationProject struct {
	Name       string
	Definition int
}

type Configuration struct {
	Project      string
	ApiVersion   string
	Username     string
	Password     string
	Definition   int
	Organization string
}

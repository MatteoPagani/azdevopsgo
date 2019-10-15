package main

type ProjectsResponse struct {
	Value []Project `json:"value"`
	Count int       `json:"count"`
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
	Id             string `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Url            string `json:"url"`
	State          string `json:"state"`
	Revision       int    `json:"revision"`
	Visibility     string `json:"visibility"`
	LastUpdateTime string `json:"lastUpdateTime"`
}

type Build struct {
	Id           int
	BuildNumber  string
	StartTime    string
	FinishTime   string
	QueueTime    string
	SourceBranch string
}

type Configuration struct {
	BaseUrl    string
	ApiVersion string
	Username   string
	Password   string
	Project    ConfigurationProject
}

type ConfigurationProject struct {
	Name    string
	Definition int
}

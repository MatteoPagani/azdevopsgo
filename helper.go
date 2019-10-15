package main

func getProjectByName(project string) Project {
	for _, element := range projects {
		if element.Name == project {
			return element
		}
	}

	return Project{}
}

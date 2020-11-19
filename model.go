package jenkinsapi

type JenkinModel struct {
	Classis string `json:"_class"`
}
type JenkinsApi struct {
	JenkinModel
	AssignLabels    map[string]string `json:"assignLabels"`
	Mode            string            `json:"mode"`
	NodeDescription string            `json:"description"`
	NodeName        string            `json:"nodeName"`
	NumExecutors    int               `json:"numExecutors"`
	Jobs            []*JenkinsProject `json:"jobs"`
	Views           []*JenkinsProject `json:"views"`
}

type JenkinsProject struct {
	JenkinModel
	Name  string `json:"name"`
	Url   string `json:"url"`
	Color string `json:"color,omitempty"`
}

type JenkinsJob struct {
	JenkinModel
	Actions         []*JenkinsJobAction `json:"actions"`
	Description     string              `json:"description"`
	DisplayName     string              `json:"displayName"`
	FullDisplayName string              `json:"fullDisplayName"`
	FullName        string              `json:"fullName"`
	Name            string              `json:"name"`
	Url             string              `json:"url"`
	Buildable       bool                `json:"buildable"`
}

type JenkinsJobAction struct {
	JenkinModel
	Params []*JenkinsJobParam `json:"parameterDefinitions"`
}

type JenkinsJobParam struct {
	JenkinModel
	Default     JenkinsJobParamDefaultValue `json:"defaultParameterValue"`
	Description string                      `json:"descritption"`
	Name        string                      `json:"name"`
	Type        string                      `json:"type"`
}

type JenkinsJobParamDefaultValue struct {
	JenkinModel
	Value string `json:"value"`
}

package jenkinsapi

type JenkinsModel struct {
	Classis string `json:"_class"`
}
type JenkinsApi struct {
	JenkinsModel
	AssignLabels    map[string]string `json:"assignLabels"`
	Mode            string            `json:"mode"`
	NodeDescription string            `json:"description"`
	NodeName        string            `json:"nodeName"`
	NumExecutors    int               `json:"numExecutors"`
	Jobs            []*JenkinsProject `json:"jobs"`
	Views           []*JenkinsProject `json:"views"`
}

type JenkinsProject struct {
	JenkinsModel
	Name  string `json:"name"`
	Url   string `json:"url"`
	Color string `json:"color,omitempty"`
}

type JenkinsJob struct {
	JenkinsModel
	Actions               []*JenkinsJobAction `json:"actions"`
	Description           string              `json:"description"`
	DisplayName           string              `json:"displayName"`
	FullDisplayName       string              `json:"fullDisplayName"`
	FullName              string              `json:"fullName"`
	Name                  string              `json:"name"`
	Url                   string              `json:"url"`
	Buildable             bool                `json:"buildable"`
	Builds                []*JenkinsJobBuild  `json:"builds"`
	FirstBuild            *JenkinsJobBuild    `json:"firstBuild"`
	InQueue               bool                `json:"inQueue"`
	LastBuild             *JenkinsJobBuild    `json:"lastBuild"`
	LastCompletedBuild    *JenkinsJobBuild    `json:"lastCompletedBuild"`
	LastFailedBuild       *JenkinsJobBuild    `json:"lastFailedBuild"`
	LastStableBuild       *JenkinsJobBuild    `json:"lastStableBuild"`
	LastSuccessfulBuild   *JenkinsJobBuild    `json:"lastSuccessfulBuild"`
	LastUnstableBuild     *JenkinsJobBuild    `json:"lastUnstableBuild"`
	LastUnsuccessfulBuild *JenkinsJobBuild    `json:"lastUnsuccessfulBuild"`
	NextBuildNumber       int                 `json:"nextBuildNumber"`
	Queue                 *JenkinsJobQueue    `json:"queueItem"`
	ConcurrentBuild       bool                `json:"concurrentBuild"`
	Disabled              bool                `json:"disabled"`
	UpstreamProjects      []*JenkinsJobBuild  `json:"upstreamProjects"`
}

type JenkinsJobQueue struct {
	JenkinsModel
	Blocked                    bool             `json:"blocked"`
	Buildable                  bool             `json:"buildable"`
	Id                         int              `json:"id"`
	InQueueSince               int64            `json:"inQueueSince"`
	Params                     string           `json:"params"`
	Stuck                      bool             `json:"stuck"`
	Task                       *JenkinsJobBuild `json:"task"`
	Url                        string           `json:"url"`
	Why                        string           `json:"why"`
	BuildableStartMilliseconds int64            `json:"buildableStartMilliseconds"`
}

type JenkinsJobBuild struct {
	JenkinsModel
	Number int    `json:"number"`
	Url    string `json:"url"`
	Color  string `json:"color"`
}

type JenkinsJobAction struct {
	JenkinsModel
	Params []*JenkinsJobParam `json:"parameterDefinitions"`
}

type JenkinsJobParam struct {
	JenkinsModel
	Default     JenkinsJobParamDefaultValue `json:"defaultParameterValue"`
	Description string                      `json:"descritption"`
	Name        string                      `json:"name"`
	Type        string                      `json:"type"`
}

type JenkinsJobParamDefaultValue struct {
	JenkinsModel
	Value interface{} `json:"value"`
}

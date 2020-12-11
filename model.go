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

type GenericTriggeResp struct {
	Jobs    map[string]*GenericTriggeJob `json:"jobs"`
	Message string                       `json:"message"`
}

type GenericTriggeJob struct {
	RegexpFilterExpression string            `json:"regexpFilterExpression"`
	Triggered              bool              `json:"triggered"`
	ResolvedVariables      map[string]string `json:"resolvedVariables"`
	RegexpFilterText       string            `json:"regexpFilterText"`
	Id                     int               `json:"id"`
	Url                    string            `json:"url"`
}

type JobResuilt struct {
	JenkinsModel
	Actions           []interface{} `json:"actions"`
	Artifacts         []interface{} `json:"artifacts"`
	Building          bool          `json:"building"`
	Description       string        `json:"description"`
	DisplayName       string        `json:"displayName"`
	Duration          int64         `json:"duration"`
	EstimatedDuration int64         `json:"estimatedDuration"`
	Executor          interface{}   `json:"executor"`
	FullDisplayName   string        `json:"fullDisplayName"`
	Id                int64         `json:"id,string"`
	KeepLog           bool          `json:"keepLog"`
	Number            int64         `json:"number"`
	QueueId           int64         `json:"queueId"`
	Result            string        `json:"result"`
	Timestamp         int64         `json:"timestamp"`
	Url               string        `json:"url"`
	ChangeSets        []interface{} `json:"changeSets"`
	Culprits          []interface{} `json:"culprits"`
	NextBuild         PreviousBuild `json:"nextBuild"`
	PreviousBuild     PreviousBuild `json:"previousBuild"`
}

type PreviousBuild struct {
	Number int64  `json:"number"`
	Url    string `json:"url"`
}

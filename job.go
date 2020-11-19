package jenkinsapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Job struct {
	Jenkins *Jenkins
}

// get all jobs
func (j *Job) GetAll() ([]*JenkinsProject, error) {
	resp, err := j.Jenkins.Do("api/json")
	if err != nil {
		return []*JenkinsProject{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []*JenkinsProject{}, err
	}
	var jenkinsApi JenkinsApi
	if err := json.Unmarshal(body, &jenkinsApi); err != nil {
		return []*JenkinsProject{}, err
	}
	return jenkinsApi.Jobs, nil
}

// get job detail
func (j *Job) Get(JobName string) (*JenkinsJob, error) {
	resp, err := j.Jenkins.Do(fmt.Sprintf("job/%s/api/json", JobName))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var jenkinsJob JenkinsJob
	if err := json.Unmarshal(body, &jenkinsJob); err != nil {
		return nil, err
	}
	return &jenkinsJob, nil
}

// trigger build
func (j *Job) Build(jobName string) error {
	resp, err := j.Jenkins.Do(fmt.Sprintf("job/%s/build?delay=0sec", jobName))
	if err != nil {
		return err
	}
	if resp != nil {
		if resp.StatusCode != http.StatusOK {
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			return errors.New(string(body))
		}
	}
	return nil
}

// trigger bild with paramaters
func (j *Job) BuildWithParamaters(jobName string, params map[string]interface{}) error {
	queryParams := ""
	for key, value := range params {
		queryParams += fmt.Sprintf("&%s=%v", key, value)
	}
	resp, err := j.Jenkins.Do(fmt.Sprintf("job/%s/buildWithParameters?%s", jobName, strings.Trim(queryParams, "&")))
	if err != nil {
		return err
	}
	if resp != nil {
		if resp.StatusCode != http.StatusOK {
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			return errors.New(string(body))
		}
	}
	return nil
}

// get all jobs by view
func (j *Job) GetAllByView(viewName string) ([]*JenkinsProject, error) {
	resp, err := j.Jenkins.Do(fmt.Sprintf("view/%s/api/json", viewName))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var jenkinsApi JenkinsApi
	if err := json.Unmarshal(body, &jenkinsApi); err != nil {
		return nil, err
	}

	return jenkinsApi.Jobs, nil
}

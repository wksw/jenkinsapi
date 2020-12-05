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
func (j *Jenkins) GetJobs() ([]*JenkinsProject, error) {
	resp, err := j.Do("api/json", nil)
	if err != nil {
		return []*JenkinsProject{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []*JenkinsProject{}, err
	}
	defer resp.Body.Close()
	var jenkinsApi JenkinsApi
	if err := json.Unmarshal(body, &jenkinsApi); err != nil {
		return []*JenkinsProject{}, err
	}
	return jenkinsApi.Jobs, nil
}

// get job detail
func (j *Jenkins) GetJob(JobName string) (*JenkinsJob, error) {
	resp, err := j.Do(fmt.Sprintf("job/%s/api/json", JobName), nil)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var jenkinsJob JenkinsJob
	if err := json.Unmarshal(body, &jenkinsJob); err != nil {
		return nil, err
	}
	return &jenkinsJob, nil
}

// trigger build
func (j *Jenkins) Build(jobName string) error {
	resp, err := j.Do(fmt.Sprintf("job/%s/build?delay=0sec", jobName), nil)
	if err != nil {
		return err
	}
	if resp != nil {
		if resp.StatusCode != http.StatusOK {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			return errors.New(string(body))
		}
	}
	return nil
}

// trigger bild with paramaters
func (j *Jenkins) BuildWithParamaters(jobName string, params map[string]interface{}) error {
	queryParams := ""
	for key, value := range params {
		queryParams += fmt.Sprintf("&%s=%v", key, value)
	}
	resp, err := j.Do(fmt.Sprintf("job/%s/buildWithParameters?%s", jobName, strings.Trim(queryParams, "&")), nil)
	if err != nil {
		return err
	}
	if resp != nil {
		if resp.StatusCode != http.StatusOK {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			return errors.New(string(body))
		}
	}
	return nil
}

// get all jobs by view
func (j *Jenkins) GetJobsByView(viewName string) ([]*JenkinsProject, error) {
	resp, err := j.Do(fmt.Sprintf("view/%s/api/json", viewName), nil)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var jenkinsApi JenkinsApi
	if err := json.Unmarshal(body, &jenkinsApi); err != nil {
		return nil, err
	}

	return jenkinsApi.Jobs, nil
}

// Generic trigger
func (j *Jenkins) GenericTrigge(token string, data []byte) (*GenericTriggeResp, error) {
	resp, err := j.Do(fmt.Sprintf("generic-webhook-trigger/invoke?token=%s", token), data)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var genericTriggerResp GenericTriggeResp
	if err := json.Unmarshal(body, &genericTriggerResp); err != nil {
		return nil, err
	}
	return &genericTriggerResp, nil
}

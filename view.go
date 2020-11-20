package jenkinsapi

import (
	"encoding/json"
	"io/ioutil"
)

func (j *Jenkins) GetViews() ([]*JenkinsProject, error) {
	resp, err := j.Do("api/json")
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
	return jenkinsApi.Views, nil
}

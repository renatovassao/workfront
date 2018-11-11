package workfront

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Project struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ProjectSearchResponse struct {
	Data []Project
}

var ProjectURL = BaseURL + "/project"

func SearchProjects(val url.Values) ([]Project, error) {
	client := &http.Client{}

	url := ProjectURL + "/search?" + val.Encode()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("SessionID", login.SessionID)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var search ProjectSearchResponse

	err = json.Unmarshal(body, &search)
	if err != nil {
		return nil, err
	}

	return search.Data, nil

}

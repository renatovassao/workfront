package workfront

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Task struct {
	ID                    string `json:"id"`
	Name                  string `json:"name"`
	PlannedCompletionDate string `json:"plannedCompletionDate"`
	Status                string `json:"status"`
}

type TaskSearchResponse struct {
	Data []Task
}

var TaskURL = BaseURL + "/task"

func SearchTasks(val url.Values) ([]Task, error) {
	client := &http.Client{}

	url := TaskURL + "/search?" + val.Encode()

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

	var search TaskSearchResponse

	err = json.Unmarshal(body, &search)
	if err != nil {
		return nil, err
	}

	return search.Data, nil

}

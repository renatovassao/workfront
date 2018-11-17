package workfront

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type LoginData struct {
	SessionID string `json:"sessionID"`
	UserID    string `json:"userID"`
}

type LoginResponse struct {
	Data *LoginData `json:"data"`
}

const BaseURL = "https://agencia.my.workfront.com/attask/api/v9.0"

var login *LoginData

// Login function takes a username and a password and logs in user in workfront API
func Login(username, password string) (string, error) {
	val := url.Values{"username": {username}, "password": {password}}

	resp, err := http.PostForm(BaseURL+"/login", val)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var lr LoginResponse

	err = json.Unmarshal(body, &lr)
	if err != nil {
		return "", err
	}

	login = lr.Data

	if login == nil {
		return "", errors.New("invalid credentials or workfront API is unavailable")
	}

	return login.UserID, nil
}

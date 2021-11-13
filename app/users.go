package app

import (
	"io/ioutil"
	"net/http"
)

const BASE_URL = "https://jsonplaceholder.typicode.com"

func GetUsers() (string, error) {
	resp, err := http.Get(BASE_URL + "/users")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

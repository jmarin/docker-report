package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type ID struct {
	ImageID string
}

func TagImageID(repo_url string, repo_port string, image_name string, tag string) string {

	url := repo_url + ":" + repo_port + "/v1/repositories/" + image_name + "/tags/" + tag

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return strings.Replace(string(contents), "\"", "", -1)

}

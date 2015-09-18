package main

import (
	//"encoding/json"
	"fmt"
	"github.com/jmarin/docker-report/Godeps/_workspace/src/github.com/fsouza/go-dockerclient"
	"strings"
)

func main() {
	client, _ := docker.NewClientFromEnv()
	containers := Containers(client)
	for _, container := range containers {
		c, _ := client.InspectContainer(container.ID)
		cImage := c.Image
		fullname := c.Config.Image
		namesVersion := strings.Split(fullname, "/")
		var name string
		if len(namesVersion) < 2 {
			name = namesVersion[0]
		} else {
			name = strings.Split(namesVersion[1], ":")[0]
		}
		id := TagImageID("http://awsdevhmdal05", "5000", name, "latest")
		if id != cImage {
			fmt.Printf("ERROR: container's image can't be found in the registry -> " + name + ": " + cImage + "\n")
		}
	}
}

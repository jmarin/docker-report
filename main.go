package main

import (
	//"encoding/json"
	"fmt"
	"github.com/jmarin/docker-report/Godeps/_workspace/src/github.com/fsouza/go-dockerclient"
)

func main() {
	client, _ := docker.NewClientFromEnv()
	containers := Containers(client)
	for _, container := range containers {
		c, _ := client.InspectContainer(container.ID)
		fmt.Printf("%v", c)
	}
}

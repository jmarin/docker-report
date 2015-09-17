package main

import (
	"github.com/jmarin/docker-report/Godeps/_workspace/src/github.com/fsouza/go-dockerclient"
)

func Images(client *docker.Client) (imgs []docker.APIImages) {
	img, _ := client.ListImages(docker.ListImagesOptions{All: false})
	return img
}

func Containers(client *docker.Client) (containers []docker.APIContainers) {
	cont, _ := client.ListContainers(docker.ListContainersOptions{All: false})
	return cont
}

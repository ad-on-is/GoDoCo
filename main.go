package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type DockerCompose struct {
	Services map[string]Service `yaml:"services"`
}

type Service struct {
	Image         string   `yaml:"image"`
	ContainerName string   `yaml:"container_name"`
	DNS           string   `yaml:"dns"`
	User          string   `yaml:"user"`
	Volumes       []string `yaml:"volumes"`
	Environment   []string `yaml:"environment"`
	Links         []string `yaml:"links"`
	Ports         []string `yaml:"ports"`
}

func main() {
	args := os.Args
	down := false
	if len(args) > 1 && args[1] == "down" {
		down = true
	}
	docker_command := ""
	dockerLogs_command := ""
	var dc DockerCompose
	file, _ := ioutil.ReadFile("docker-compose.yml")
	err := yaml.Unmarshal(file, &dc)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}
	for _, service := range dc.Services {

		if down {
			docker_command += "docker stop " + service.ContainerName + "; "
			continue
		}

		docker_command += "docker run -d"
		if service.ContainerName != "" {
			docker_command += " --name " + service.ContainerName
		}
		if service.DNS != "" {
			docker_command += " --dns " + service.DNS
		}
		if service.User != "" {
			docker_command += " --user " + service.User
		}

		if len(service.Volumes) > 0 {
			for _, volume := range service.Volumes {
				docker_command += " -v " + volume
			}
		}

		if len(service.Ports) > 0 {
			for _, port := range service.Ports {
				docker_command += " -p " + port
			}
		}

		if len(service.Environment) > 0 {
			for _, env := range service.Environment {
				docker_command += " -e " + env
			}
		}

		if len(service.Links) > 0 {
			docker_command += " -P"
			for _, link := range service.Links {
				docker_command += " --link " + link
			}
		}

		docker_command += " " + service.Image + ";"
		dockerLogs_command += "docker logs -f " + service.ContainerName + " | sed -e 's/^/[-- " + service.ContainerName + " --]/' & "
	}

	if down {
		docker_command += "docker rm $(docker ps -a -q);"
	}

	fmt.Println(docker_command)
	fmt.Println(dockerLogs_command)

}

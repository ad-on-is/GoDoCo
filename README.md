# GoDoCo
GoLang Docker-Compose

## Info
This script can be used as a temporary solution, if docker-compose doesn't work for whatever reason.

It reads a `docker-compose.yml` file and outputs `docker run` commands with volumes, links, envs, etc...

## Usage
1. Place godoco executable in the same folder as the `docker-compose.yml` is located
2. Run `./godoco up` to get the commands for starting the docker containers
3. Copy output and paste into terminal
4. Run `./godoco down` to get the commands for stoping and deleting containers

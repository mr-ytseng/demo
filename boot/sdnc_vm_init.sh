#!/bin/bash

NEXUS_USERNAME=$(cat /opt/config/nexus_username.txt)
NEXUS_PASSWD=$(cat /opt/config/nexus_password.txt)
export NEXUS_DOCKER_REPO=$(cat /opt/config/nexus_docker_repo.txt)

cd /opt/sdnc
git pull

cd /opt/sdnc/installation/src/main/yaml
docker login -u $NEXUS_USERNAME -p $NEXUS_PASSWD $NEXUS_DOCKER_REPO
/opt/docker/docker-compose up -d

---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/how-to-run-unifi-controller-in-docker-container-a59cd191-2
title: "How To Run UniFi Controller in Docker Container"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: ["2024-01", "2026-07"]
keywords: ["latency", "parameters"]
source: docs/RAG/collect-261001-unifi-ubiquiti/how-to-run-unifi-controller-in-docker-container-a59cd191.md
source_anchor: ""
source_lines: [5, 138]
sha256: 68bdb37be26d077ba8a23c84169de715ddc6273b3e06bd16b8a9f9f0dfd7406c
---

# How To Run UniFi Controller in Docker Container

The main objective of UniFi is to simplify IT operations. It does this with a single stack that ties together networking, communications, security, and access control behind one interface. It allows one to manage deployments from both local and cloud environments.
The UniFi OS is the operating system that hosts the UniFi application suite. There are several products available in the UniFi application suite. These include:
The UniFi Controller is a wireless network management software solution developed by Ubiquiti Networks™. It allows one to manage several wireless networks from its web UI. This tool is ideal for high-density deployments that required low latency and high uptime.
The UniFi Controller can be installed on Linux, Mac OS X, or Windows 10 or 11 by downloading the UniFi Controller software from the Ubiquiti Networks website. But this process requires one to install several dependencies such as Java Runtime Environment e.t.c
In this guide, we will learn how to Run UniFi Controller in Docker Container. This method is preferred since the containers come with all the dependencies bundled and make it easy to run the UniFi Controller.
This guide has been updated to the supported unifi-network-application image. LinuxServer.io retired the older unifi-controller image in January 2024, so containers still based on it run an outdated controller and receive no updates. The supported image no longer bundles a database, which means the setup is now two containers: the Network Application itself and a MongoDB instance it connects to.
Re-tested July 2026 on UniFi Network Application 10.4 with MongoDB 8.0.
Prefer a native package over a container? The controller also installs directly on Ubuntu and Debian.
You will need UniFi network devices to adopt. The UniFi U6 Lite is a popular WiFi 6 access point for home and small office deployments, and the UniFi Switch Lite 8 PoE provides managed switching with built-in PoE to power the APs. If you prefer an all-in-one appliance over running the controller in Docker, the UniFi Dream Router 7 has the controller, router, and switch built in.
Before we begin, you need to have the Docker Engine installed on your system. The guide below can be used to achieve this:
Once installed, ensure that the Docker service is up and running;
sudo systemctl start docker && sudo systemctl enable docker
Add your system user to the Docker group;
sudo usermod -aG docker $USER
newgrp docker
Persistent volumes keep the data across container restarts. The stack needs three things on the host: a config directory for the application, a data directory for MongoDB, and the database init script:
sudo mkdir -p /unifi_data/config /unifi_data/db
Set the permissions:
sudo chmod 775 -R /unifi_data
MongoDB creates the database user for the controller from an init script the first time it starts. Create the script:
sudo vim /unifi_data/init-mongo.sh
Paste in the following. It reads the MONGO_ variables we pass to the database container and creates the user the controller authenticates with:
#!/bin/bash
if which mongosh > /dev/null 2>&1; then
  mongo_init_bin='mongosh'
else
  mongo_init_bin='mongo'
fi
"${mongo_init_bin}" <<EOF
use ${MONGO_AUTHSOURCE}
db.auth("${MONGO_INITDB_ROOT_USERNAME}", "${MONGO_INITDB_ROOT_PASSWORD}")
db.createUser({
  user: "${MONGO_USER}",
  pwd: "${MONGO_PASS}",
  roles: [
    "clusterMonitor",
    { db: "${MONGO_DBNAME}", role: "dbOwner" },
    { db: "${MONGO_DBNAME}_stat", role: "dbOwner" },
    { db: "${MONGO_DBNAME}_audit", role: "dbOwner" },
    { db: "${MONGO_DBNAME}_restore", role: "dbOwner" }
  ]
})
EOF
Make it executable:
sudo chmod +x /unifi_data/init-mongo.sh
The script must exist as a regular file before the first start. If the path is missing when the container starts, Docker creates an empty directory there instead, the init script never runs, and the controller can never authenticate to the database. On RHEL-based systems, keep SELinux enforcing and append :Z to each volume mapping (for example -v /unifi_data/config:/config:Z) so Docker relabels the paths for container access.
Once the Docker engine has been installed, you can easily run the UniFi Controller from the docker command line.
The command has several parameters that include:
default
The supported image needs its MongoDB companion on the same Docker network, so the command-line path is three commands: create the network, start the database, then start the controller. Create the network first:
docker network create unifi-net
Start MongoDB with the init script mounted. The MONGO_ values are read by the init script on first run to create the database user, so change the two passwords to your own. MongoDB 8.0 needs a CPU with AVX support, so on a Proxmox or KVM guest set the VM CPU type to host; the default kvm64 model masks AVX and the database exits on an illegal instruction at startup:
docker run -d --name=unifi-db --net unifi-net 
  -e MONGO_INITDB_ROOT_USERNAME=root 
  -e MONGO_INITDB_ROOT_PASSWORD=Str0ngRootPass 
  -e MONGO_USER=unifi 
  -e MONGO_PASS=Str0ngUnifiPass 
  -e MONGO_DBNAME=unifi 
  -e MONGO_AUTHSOURCE=admin 
  -v /unifi_data/db:/data/db 
  -v /unifi_data/init-mongo.sh:/docker-entrypoint-initdb.d/init-mongo.sh:ro 
  --restart unless-stopped 
  mongo:8.0
MongoDB only runs the init script against an empty data directory. If the database ever starts with the wrong settings, stop it, wipe /unifi_data/db, and start again. Now launch the Network Application, pointing it at the database container. Add any optional ports from the list above that you need:
docker run -d --name=unifi-network-application --net unifi-net 
  -e PUID=1000 
  -e PGID=1000 
  -e TZ=Etc/UTC 
  -e MONGO_USER=unifi 
  -e MONGO_PASS=Str0ngUnifiPass 
  -e MONGO_HOST=unifi-db 
  -e MONGO_PORT=27017 
  -e MONGO_DBNAME=unifi 
  -e MONGO_AUTHSOURCE=admin 
  -e MEM_LIMIT=1024 
  -e MEM_STARTUP=1024 
  -p 8443:8443 
  -p 3478:3478/udp 
  -p 10001:10001/udp 
  -p 8080:8080 
  -v /unifi_data/config:/config 
  --restart unless-stopped 
  lscr.io/linuxserver/unifi-network-application:latest
First startup takes a few minutes while the Java application initialises. Poll the status endpoint until it reports up:
curl -sk https://localhost:8443/status
A fully started controller answers with its version:
{"meta":{"rc":"ok","up":true,"server_version":"10.4.57","uuid":"016f58b1-2e5c-4fe4-97ed-9bd3229ac209"},"data":[]}
You can also run the UniFi Controller using Docker Compose. First, ensure that Docker compose is installed on your system.
The Compose file wires the same two containers together and reads the secrets from an .env file sitting next to it, so create that first. Make a project directory and open the environment file:
mkdir -p ~/unifi && cd ~/unifi
vim .env
Define the two database passwords in it, using your own values:
MONGO_ROOT_PASS=Str0ngRootPass
MONGO_PASS=Str0ngUnifiPass
Then create the Compose file:
vim docker-compose.yml
It defines the database, with the init script mounted, and the controller:
---
services:
  unifi-db:
    image: mongo:8.0
    container_name: unifi-db
    environment:
      - MONGO_INITDB_ROOT_USERNAME=root
      - MONGO_INITDB_ROOT_PASSWORD=${MONGO_ROOT_PASS}
      - MONGO_USER=unifi
      - MONGO_PASS=${MONGO_PASS}
      - MONGO_DBNAME=unifi
      - MONGO_AUTHSOURCE=admin
    volumes:
      - /unifi_data/db:/data/db
      - /unifi_data/init-mongo.sh:/docker-entrypoint-initdb.d/init-mongo.sh:ro
    restart: unless-stopped
  unifi-network-application:
    image: lscr.io/linuxserver/unifi-network-application:latest
    container_name: unifi-network-application
    depends_on:
      - unifi-db
    environment:
      - PUID=1000
      - PGID=1000
      - TZ=Etc/UTC
      - MONGO_USER=unifi
      - MONGO_PASS=${MONGO_PASS}
      - MONGO_HOST=unifi-db
      - MONGO_PORT=27017
      - MONGO_DBNAME=unifi
      - MONGO_AUTHSOURCE=admin
      - MEM_LIMIT=1024
      - MEM_STARTUP=1024
    volumes:
      - /unifi_data/config:/config
    ports:
      - 8443:8443

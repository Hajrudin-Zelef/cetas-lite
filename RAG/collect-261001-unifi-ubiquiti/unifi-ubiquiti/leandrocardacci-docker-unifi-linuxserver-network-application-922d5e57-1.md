---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/leandrocardacci-docker-unifi-linuxserver-network-application-922d5e57-1
title: "leandrocardacci-docker-unifi-linuxserver-network-application-922d5e57"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["intel", "latency"]
source: docs/RAG/collect-261001-unifi-ubiquiti/leandrocardacci-docker-unifi-linuxserver-network-application-922d5e57.md
source_anchor: ""
source_lines: [1, 141]
sha256: 4998c01ed92da5b6f2c2f5ac10fbbb83a913621789de5dc43e78c98dc23d0fe9
---

# leandrocardacci-docker-unifi-linuxserver-network-application-922d5e57

The LinuxServer.io team brings you another container release featuring:
- regular and timely application updates
- easy user mappings (PGID, PUID)
- custom base image with s6 overlay
- weekly base OS updates with common layers across the entire LinuxServer.io ecosystem to minimise space usage, down time and bandwidth
- regular security updates
Find us at:
- Blog - all the things you can do with our containers including How-To guides, opinions and much more!
- Discord - realtime support / chat with the community and the team.
- Discourse - post on our community forum.
- Fleet - an online web interface which displays all of our maintained images.
- GitHub - view the source for all of our repositories.
- Open Collective - please consider helping us by either donating or contributing to our budget
The Unifi-network-application software is a powerful, enterprise wireless software engine ideal for high-density client deployments requiring low latency and high uptime performance.
We utilise the docker manifest for multi-platform awareness. More information is available from docker here and our announcement here.
Simply pulling lscr.io/linuxserver/unifi-network-application:latest should retrieve the correct image for your arch, but you can also pull specific arch images via tags.
The architectures supported by this image are:
| Architecture | Available | Tag | 
|---|---|---|
| x86-64 | ✅ | amd64-<version tag> | 
| arm64 | ✅ | arm64v8-<version tag> | 
| armhf | ❌ |  | 
After setup, the web UI is available at https://ip:8443. The application can be configured, or a backup restored, using the first run wizard.
This container requires an external mongodb database instance.
Starting with version 8.1 of Unifi Network Application, mongodb 3.6 through 7.0 are supported. Starting with version 9.0 of Unifi Network Application, mongodb 8.0 is also supported.
Make sure you pin your database image version and do not use latest, as mongodb does not support automatic upgrades between major versions.
MongoDB >4.4 on X86_64 Hardware needs a CPU with AVX support. Some lower end Intel CPU models like Celeron and Pentium (before Tiger-Lake) more Details: Advanced Vector Extensions - Wikipedia don't support AVX, but you can still use MongoDB 4.4.
If you are using the official mongodb container, you can create your user using an init-mongo.sh file with the following contents (do not modify; copy/paste as is):
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
    { db: "${MONGO_DBNAME}", role: "dbOwner" },
    { db: "${MONGO_DBNAME}_stat", role: "dbOwner" }
  ]
})
EOF
Mount the sh file into your mongodb container, and make sure to set the env vars below with the same values you supplied to the Unifi container.
For example:
  unifi-db:
    image: docker.io/mongo:<version tag>
    container_name: unifi-db
    environment:
      - MONGO_INITDB_ROOT_USERNAME=root
      - MONGO_INITDB_ROOT_PASSWORD=
      - MONGO_USER=unifi
      - MONGO_PASS=
      - MONGO_DBNAME=unifi
      - MONGO_AUTHSOURCE=admin
    volumes:
      - /path/to/data:/data/db
      - /path/to/init-mongo.sh:/docker-entrypoint-initdb.d/init-mongo.sh:ro
    restart: unless-stopped
Note that the init script method will only work on first run. If you start the Mongodb container without an init script it will generate test data automatically and you will have to manually create your databases, or restart with a clean /data/db volume and an init script mounted.
You can also run the commands directly against the database using either mongo (< 6.0) or mongosh (>= 6.0).
For Unifi to adopt other devices, e.g. an Access Point, it is required to change the inform IP address. Because Unifi runs inside Docker by default it uses an IP address not accessible by other devices. To change this go to Settings > System > Advanced and set the Inform Host to a hostname or IP address accessible by your devices. Additionally the checkbox "Override" has to be checked, so that devices can connect to the controller during adoption (devices use the inform-endpoint during adoption).
Note that you must use 8080:8080. If you MUST change the port, it must be changed on both sides and manually changed in your system.properties file. Otherwise, devices will initially communicate and then break after.
Please note, Unifi change the location of this option every few releases so if it's not where it says, search for "Inform" or "Inform Host" in the settings.
In order to manually adopt a device take these steps:
ssh ubnt@$AP-IP
set-inform http://$address:8080/inform
The default device password is ubnt. $address is the IP address of the host you are running this container on and $AP-IP is the Access Point IP address.
When using a Security Gateway (router) it could be that network connected devices are unable to obtain an ip address. This can be fixed by setting "DHCP Gateway IP", under Settings > Networks > network_name, to a correct (and accessible) ip address.
Migration From Unifi-Controller
If you were using the mongoless tag for the Unifi Controller container, you can switch directly to the Unifi Network Application container without needing to perform any migration steps.
You cannot perform an in-place upgrade from an existing Unifi-Controller container, you must run a backup and then a restore.
The simplest migration approach is to take a full backup of your existing install, including history, from the Unifi-Controller web UI, then shut down the old container.
You can then start up the new container with a clean /config mount (and a database container configured), and perform a restore using the setup wizard.
This image uses a self-signed certificate by default. This naturally means the scheme is https.
If you are using a reverse proxy which validates certificates, you need to disable this check for the container.
To help you get started creating a container from this image you can either use docker-compose or the docker cli.
Note
Unless a parameter is flaged as 'optional', it is mandatory and a value must be provided.
docker-compose (recommended, click here for more info)
---
services:
  unifi-network-application:
    image: lscr.io/linuxserver/unifi-network-application:latest
    container_name: unifi-network-application
    environment:
      - PUID=1000
      - PGID=1000
      - TZ=Etc/UTC
      - MONGO_USER=unifi
      - MONGO_PASS=
      - MONGO_HOST=unifi-db
      - MONGO_PORT=27017
      - MONGO_DBNAME=unifi
      - MONGO_AUTHSOURCE=admin
      - MEM_LIMIT=1024 #optional
      - MEM_STARTUP=1024 #optional
      - MONGO_TLS= #optional
    volumes:
      - /path/to/unifi-network-application/data:/config
    ports:
      - 8443:8443
      - 3478:3478/udp
      - 10001:10001/udp
      - 8080:8080
      - 1900:1900/udp #optional
      - 8843:8843 #optional
      - 8880:8880 #optional
      - 6789:6789 #optional
      - 5514:5514/udp #optional
    restart: unless-stopped
docker cli (click here for more info)
docker run -d \
  --name=unifi-network-application \
  -e PUID=1000 \
  -e PGID=1000 \
  -e TZ=Etc/UTC \
  -e MONGO_USER=unifi \
  -e MONGO_PASS= \
  -e MONGO_HOST=unifi-db \
  -e MONGO_PORT=27017 \
  -e MONGO_DBNAME=unifi \
  -e MONGO_AUTHSOURCE=admin \
  -e MEM_LIMIT=1024 `#optional` \
  -e MEM_STARTUP=1024 `#optional` \
  -e MONGO_TLS= `#optional` \
  -p 8443:8443 \
  -p 3478:3478/udp \
  -p 10001:10001/udp \
  -p 8080:8080 \
  -p 1900:1900/udp `#optional` \
  -p 8843:8843 `#optional` \
  -p 8880:8880 `#optional` \
  -p 6789:6789 `#optional` \
  -p 5514:5514/udp `#optional` \
  -v /path/to/unifi-network-application/data:/config \
  --restart unless-stopped \
  lscr.io/linuxserver/unifi-network-application:latest

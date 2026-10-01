---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/deprecated-images-docker-unifi-controller-2b1a8131-1
title: "deprecated-images-docker-unifi-controller-2b1a8131"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: ["2023-09-06", "2024-01-01"]
keywords: ["latency", "memory", "parameters", "throughput"]
source: docs/RAG/collect-261001-unifi-ubiquiti/deprecated-images-docker-unifi-controller-2b1a8131.md
source_anchor: ""
source_lines: [1, 143]
sha256: d2573d937e545eff69a64850b5e5832100313c0bb4407611bb27ee7c8212d94b
---

# deprecated-images-docker-unifi-controller-2b1a8131

Warning
This image is deprecated. We will not offer support for this image and it will not be updated.
We recommend our unifi-network-application image instead: https://github.com/linuxserver/docker-unifi-network-application
linuxserver/unifi-controller¶
The Unifi-controller software is a powerful, enterprise wireless software engine ideal for high-density client deployments requiring low latency and high uptime performance.
Supported Architectures¶
We utilise the docker manifest for multi-platform awareness. More information is available from docker here and our announcement here.
Simply pulling lscr.io/linuxserver/unifi-controller:latest should retrieve the correct image for your arch, but you can also pull specific arch images via tags.
The architectures supported by this image are:
| Architecture | Available | Tag | 
|---|---|---|
| x86-64 | ✅ | amd64-<version tag> | 
| arm64 | ✅ | arm64v8-<version tag> | 
| armhf | ❌ |  | 
Version Tags¶
This image provides various versions that are available via tags. Please read the descriptions carefully and exercise caution when using unstable or development tags.
| Tag | Available | Description | 
|---|---|---|
| latest | ✅ | Stable Unifi Controller releases. | 
| mongoless | ✅ | Stable Unifi Controller releases without mongodb included. | 
Application Setup¶
From 2024-01-01 this image will be deprecated and it will no longer be updated. Please migrate to our Unifi Network Application image instead¶
See: https://info.linuxserver.io/issues/2023-09-06-unifi-controller for more information.
The webui is at https://ip:8443, setup with the first run wizard.
For Unifi to adopt other devices, e.g. an Access Point, it is required to change the inform IP address. Because Unifi runs inside Docker by default it uses an IP address not accessible by other devices. To change this go to Settings > System > Advanced and set the Inform Host to a hostname or IP address accessible by your devices. Additionally the checkbox "Override" has to be checked, so that devices can connect to the controller during adoption (devices use the inform-endpoint during adoption).
Please note, Unifi change the location of this option every few releases so if it's not where it says, search for "Inform" or "Inform Host" in the settings.
In order to manually adopt a device take these steps:
The default device password is ubnt. $address is the IP address of the host you are running this container on and $AP-IP is the Access Point IP address.
When using a Security Gateway (router) it could be that network connected devices are unable to obtain an ip address. This can be fixed by setting "DHCP Gateway IP", under Settings > Networks > network_name, to a correct (and accessable) ip address.
Strict reverse proxies¶
This image uses a self-signed certificate by default. This naturally means the scheme is https. If you are using a reverse proxy which validates certificates, you need to disable this check for the container.
Usage¶
To help you get started creating a container from this image you can either use docker-compose or the docker cli.
docker-compose (recommended, click here for more info)¶
---
version: "2.1"
services:
  unifi-controller:
    image: lscr.io/linuxserver/unifi-controller:latest
    container_name: unifi-controller
    environment:
      - PUID=1000
      - PGID=1000
      - TZ=Etc/UTC
      - MEM_LIMIT=1024 #optional
      - MEM_STARTUP=1024 #optional
    volumes:
      - /path/to/data:/config
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
docker cli (click here for more info)¶
docker run -d \
  --name=unifi-controller \
  -e PUID=1000 \
  -e PGID=1000 \
  -e TZ=Etc/UTC \
  -e MEM_LIMIT=1024 `#optional` \
  -e MEM_STARTUP=1024 `#optional` \
  -p 8443:8443 \
  -p 3478:3478/udp \
  -p 10001:10001/udp \
  -p 8080:8080 \
  -p 1900:1900/udp `#optional` \
  -p 8843:8843 `#optional` \
  -p 8880:8880 `#optional` \
  -p 6789:6789 `#optional` \
  -p 5514:5514/udp `#optional` \
  -v /path/to/data:/config \
  --restart unless-stopped \
  lscr.io/linuxserver/unifi-controller:latest
Parameters¶
Containers are configured using parameters passed at runtime (such as those above). These parameters are separated by a colon and indicate <external>:<internal> respectively. For example, -p 8080:80 would expose port 80 from inside the container to be accessible from the host's IP on port 8080 outside the container.
Ports (-p)¶
 | Parameter | Function | 
|---|---|
| 8443 | Unifi web admin port | 
| 3478/udp | Unifi STUN port | 
| 10001/udp | Required for AP discovery | 
| 8080 | Required for device communication | 
| 1900/udp | Required for Make controller discoverable on L2 network option | 
| 8843 | Unifi guest portal HTTPS redirect port | 
| 8880 | Unifi guest portal HTTP redirect port | 
| 6789 | For mobile throughput test | 
| 5514/udp | Remote syslog port | 
Environment Variables (-e)¶
 | Env | Function | 
|---|---|
| PUID=1000 | for UserID - see below for explanation | 
| PGID=1000 | for GroupID - see below for explanation | 
| TZ=Etc/UTC | specify a timezone to use, see this list. | 
| MEM_LIMIT=1024 | Optionally change the Java memory limit (in Megabytes). Set to default to reset to default | 
| MEM_STARTUP=1024 | Optionally change the Java initial/minimum memory (in Megabytes). Set to default to reset to default | 
Volume Mappings (-v)¶
 | Volume | Function | 
|---|---|
| /config | All Unifi data stored here | 
Miscellaneous Options¶
| Parameter | Function | 
|---|---|
Environment variables from files (Docker secrets)¶
You can set any environment variable from a file by using a special prepend FILE__.
As an example:
Will set the environment variable MYVAR based on the contents of the /run/secrets/mysecretvariable file.
Umask for running applications¶
For all of our images we provide the ability to override the default umask settings for services started within the containers using the optional -e UMASK=022 setting. Keep in mind umask is not chmod it subtracts from permissions based on it's value it does not add. Please read up here before asking for support.
User / Group Identifiers¶
When using volumes (-v flags), permissions issues can arise between the host OS and the container, we avoid this issue by allowing you to specify the user PUID and group PGID.
Ensure any volume directories on the host are owned by the same user you specify and any permissions issues will vanish like magic.
In this instance PUID=1000 and PGID=1000, to find yours use id your_user as below:
Example output:
Docker Mods¶
We publish various Docker Mods to enable additional functionality within the containers. The list of Mods available for this image (if any) as well as universal mods that can be applied to any one of our images can be accessed via the dynamic badges above.
Support Info¶
-  Shell access whilst the container is running:
-  To monitor the logs of the container in realtime:
-  Container version number:
-  Image version number:
Updating Info¶
Most of our images are static, versioned, and require an image update and container recreation to update the app inside. With some exceptions (ie. nextcloud, plex), we do not recommend or support updating apps inside the container. Please consult the Application Setup section above to see if it is recommended for the image.
Below are the instructions for updating containers:
Via Docker Compose¶
-  Update images: 
  -  All images:
  -  Single image:
-  
-  Update containers: 
  -  All containers:
  -  Single container:
-  
-  You can also remove the old dangling images:
Via Docker Run¶
-  Update the image:
-  Stop the running container:
-  Delete the container:

---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/leandrocardacci-docker-unifi-linuxserver-network-application-922d5e57-2
title: "leandrocardacci-docker-unifi-linuxserver-network-application-922d5e57"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["memory", "parameters", "throughput"]
source: docs/RAG/collect-261001-unifi-ubiquiti/leandrocardacci-docker-unifi-linuxserver-network-application-922d5e57.md
source_anchor: ""
source_lines: [142, 234]
sha256: c5118d9956d42fc30ad68e91ce9ec9d2192450bdddfd89523db2a0f5bc370db2
---

# leandrocardacci-docker-unifi-linuxserver-network-application-922d5e57

Containers are configured using parameters passed at runtime (such as those above). These parameters are separated by a colon and indicate <external>:<internal> respectively. For example, -p 8080:80 would expose port 80 from inside the container to be accessible from the host's IP on port 8080 outside the container.
| Parameter | Function | 
|---|---|
| -p 8443:8443 | Unifi web admin port | 
| -p 3478:3478/udp | Unifi STUN port | 
| -p 10001:10001/udp | Required for AP discovery | 
| -p 8080:8080 | Required for device communication | 
| -p 1900/udp | Required for Make controller discoverable on L2 network option | 
| -p 8843 | Unifi guest portal HTTPS redirect port | 
| -p 8880 | Unifi guest portal HTTP redirect port | 
| -p 6789 | For mobile throughput test | 
| -p 5514/udp | Remote syslog port | 
| -e PUID=1000 | for UserID - see below for explanation | 
| -e PGID=1000 | for GroupID - see below for explanation | 
| -e TZ=Etc/UTC | specify a timezone to use, see this list. | 
| -e MONGO_USER=unifi | Mongodb Username. Only evaluated on first run. Special characters must be url encoded. | 
| -e MONGO_PASS= | Mongodb Password. Only evaluated on first run. Special characters must be url encoded. | 
| -e MONGO_HOST=unifi-db | Mongodb Hostname. Only evaluated on first run. | 
| -e MONGO_PORT=27017 | Mongodb Port. Only evaluated on first run. | 
| -e MONGO_DBNAME=unifi | Mongodb Database Name (stats DB is automatically suffixed with _stat ). Only evaluated on first run. | 
| -e MONGO_AUTHSOURCE=admin | Mongodb authSource. For Atlas set to admin . Only evaluated on first run. | 
| -e MEM_LIMIT=1024 | Optionally change the Java memory limit (in Megabytes). Set to default to reset to default | 
| -e MEM_STARTUP=1024 | Optionally change the Java initial/minimum memory (in Megabytes). Set to default to reset to default | 
| -e MONGO_TLS= | Mongodb enable TLS. Only evaluated on first run. | 
| -v /config | Persistent config files | 
You can set any environment variable from a file by using a special prepend FILE__.
As an example:
-e FILE__MYVAR=/run/secrets/mysecretvariable
Will set the environment variable MYVAR based on the contents of the /run/secrets/mysecretvariable file.
For all of our images we provide the ability to override the default umask settings for services started within the containers using the optional -e UMASK=022 setting.
Keep in mind umask is not chmod it subtracts from permissions based on it's value it does not add. Please read up here before asking for support.
When using volumes (-v flags), permissions issues can arise between the host OS and the container, we avoid this issue by allowing you to specify the user PUID and group PGID.
Ensure any volume directories on the host are owned by the same user you specify and any permissions issues will vanish like magic.
In this instance PUID=1000 and PGID=1000, to find yours use id your_user as below:
id your_user
Example output:
uid=1000(your_user) gid=1000(your_user) groups=1000(your_user)
We publish various Docker Mods to enable additional functionality within the containers. The list of Mods available for this image (if any) as well as universal mods that can be applied to any one of our images can be accessed via the dynamic badges above.
- 
Shell access whilst the container is running: docker exec -it unifi-network-application /bin/bash
- 
To monitor the logs of the container in realtime: docker logs -f unifi-network-application
- 
Container version number: docker inspect -f '{{ index .Config.Labels "build_version" }}' unifi-network-application
- 
Image version number: docker inspect -f '{{ index .Config.Labels "build_version" }}' lscr.io/linuxserver/unifi-network-application:latest
Most of our images are static, versioned, and require an image update and container recreation to update the app inside. With some exceptions (noted in the relevant readme.md), we do not recommend or support updating apps inside the container. Please consult the Application Setup section above to see if it is recommended for the image.
Below are the instructions for updating containers:
- 
Update images: 
  - 
All images: docker-compose pull
  - 
Single image: docker-compose pull unifi-network-application
- 
- 
Update containers: 
  - 
All containers: docker-compose up -d
  - 
Single container: docker-compose up -d unifi-network-application
- 
- 
You can also remove the old dangling images: docker image prune
- 
Update the image: docker pull lscr.io/linuxserver/unifi-network-application:latest
- 
Stop the running container: docker stop unifi-network-application
- 
Delete the container: docker rm unifi-network-application
- 
Recreate a new container with the same docker run parameters as instructed above (if mapped correctly to a host folder, your /config folder and settings will be preserved)
- 
You can also remove the old dangling images: docker image prune
Tip
We recommend Diun for update notifications. Other tools that automatically update containers unattended are not recommended or supported.
If you want to make local modifications to these images for development purposes or just to customize the logic:
git clone https://github.com/linuxserver/docker-unifi-network-application.git
cd docker-unifi-network-application
docker build \
  --no-cache \
  --pull \
  -t lscr.io/linuxserver/unifi-network-application:latest .
The ARM variants can be built on x86_64 hardware and vice versa using lscr.io/linuxserver/qemu-static
docker run --rm --privileged lscr.io/linuxserver/qemu-static --reset
Once registered you can define the dockerfile to use with -f Dockerfile.aarch64.
- 13.02.25: - Revert JRE to 17.
- 12.02.25: - Bump JRE to 21.
- 11.08.24: - Important: The mongodb init instructions have been updated to enable auth (RBAC). We have been notified that if RBAC is not enabled, the official mongodb container allows remote access to the db contents over port 27017 without credentials. If you set up the mongodb container with the old instructions we provided, you should not map or expose port 27017. If you would like to enable auth, the easiest way is to create new instances of both unifi and mongodb with the new instructions and restore unifi from a backup.
- 11.08.24: - Rebase to Ubuntu Noble.
- 04.03.24: - Install from zip package instead of deb.
- 17.10.23: - Add environment variables for TLS and authSource to support Atlas and new MongoDB versions.
- 05.09.23: - Initial release.

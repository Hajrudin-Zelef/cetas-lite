---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hvelvigen-homelab-blob-head-systems-docker-networking-unifi-controller-md-3efba4c7
title: "hvelvigen-homelab-blob-head-systems-docker-networking-unifi-controller-md-3efba4c7"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hvelvigen-homelab-blob-head-systems-docker-networking-unifi-controller-md-3efba4c7.md
source_anchor: ""
source_lines: [1, 149]
sha256: 9e652d28697d554005c708378c8d4606da4e5696ad1bf926c344a41286ae8adb
---

# hvelvigen-homelab-blob-head-systems-docker-networking-unifi-controller-md-3efba4c7

A proper home for your UniFi Controller, not “that one random container that everything depends on”.
- Pattern name: UniFi Network Application
- Host Service Layer: Docker on Linux as a Service Host
- Category: Network Management
- Exposure: Internal LAN (optionally via reverse proxy)
- Criticality: Medium–High
- Intended users: Homelab operator / network administrator
This pattern assumes the Docker Host HSL is already in place:
- /srv/host-docker on dedicated disk
- /srv/docker and/srv/logs structured
- Docker Root Dir relocated from default
The UniFi controller becomes a defined service, not an ad-hoc container.
The UniFi Network Application manages:
- UniFi access points
- UniFi switches
- UniFi gateways
- Wireless configuration
- Guest networks and VLAN integration
- Firmware lifecycle
This pattern ensures:
- Predictable container deployment
- Persistent and backed-up configuration
- Controlled network exposure
- Safe upgrade workflow
- You want a self-hosted UniFi controller
- You standardise services on your Docker HSL
- You require portability and recoverability
- You manage multiple UniFi devices
- You rely on UniFi Cloud
- You use a CloudKey or Dream Machine as controller
- You only operate a single standalone AP
- Your Docker host is disposable or experimental
- Docker Host HSL complete
- /srv layout implemented
- Docker Root Dir relocated
- LAN routing functional
- Controller not exposed directly to the public internet
You are comfortable using:
- docker compose
- basic shell operations
- container log inspection
- Two-container deployment:
  - unifi-controller (UniFi Network Application)
  - unifi-db (MongoDB 4.4 instance)
- External MongoDB container required and managed as part of this pattern
- Both containers attached to a dedicated Docker bridge network (unifi-net )
- UniFi controller ports published on the Docker host for UI and device communication
Note: a host-network variant is possible if required for more complex L2 discovery scenarios, but this pattern standardises on a bridge network with explicit ports.
- /srv/docker/unifi-controller/config
- /srv/logs/unifi-controller
Create directories:
sudo mkdir -p /srv/docker/unifi-controller/config
sudo mkdir -p /srv/logs/unifi-controller
sudo chown -R <uid>:<gid> /srv/docker/unifi-controller /srv/logs/unifi-controller
config/ is the critical persistent data location.
logs/ is reserved for future file-based logging if required;
the initial pattern relies on docker logs for runtime inspection.
| Input | Example | Purpose | 
|---|---|---|
| TZ | Europe/London | Log and scheduler timezone | 
| PUID / PGID | 1000 / 1000 | Container runtime identity | 
| MEM_LIMIT | 1024M | JVM memory hint | 
| Network mode | host | Direct port binding | 
| Data path | ./config | Persistent controller data | 
- Navigate to service directory:
cd /srv/docker/unifi-controller
sudo nano docker-compose.yml
- Compose file:
services:
  unifi-db:
    image: mongo:4.4.25
    container_name: unifi-db
    restart: unless-stopped
    networks:
      - unifi-net
    environment:
      MONGO_INITDB_ROOT_USERNAME: root
      MONGO_INITDB_ROOT_PASSWORD: supersecretroot
    volumes:
      - /srv/docker/unifi-controller/db:/data/db
  unifi-controller:
    image: lscr.io/linuxserver/unifi-network-application:latest
    container_name: unifi-controller
    restart: unless-stopped
    depends_on:
      - unifi-db
    networks:
      - unifi-net
    environment:
      PUID: "1000"
      PGID: "1000"
      TZ: "Europe/London"
      MONGO_HOST: unifi-db
      MONGO_PORT: "27017"
      MONGO_DBNAME: unifi
      MONGO_USER: root
      MONGO_PASS: supersecretroot
      MONGO_AUTHSOURCE: admin
    volumes:
      - /srv/docker/unifi-controller/config:/config
    ports:
      - "8443:8443"
      - "8080:8080"
      - "3478:3478/udp"
      - "10001:10001/udp"
      - "1900:1900/udp"
      - "8843:8843"
      - "8880:8880"
      - "6789:6789"
      - "5514:5514/udp"
networks:
  unifi-net:
    driver: bridge
- Deploy:
docker compose pull
docker compose up -d
- Verify:
docker ps | grep unifi
docker logs unifi-controller --tail=80
- Direct LAN access:
https://<docker-host-ip>:8443/
- 
Accept self-signed certificate warning on first access.
- 
Run initial setup wizard and adopt devices.
docker compose up -d
docker compose down
docker compose restartdocker logs unifi-controller --tail=100
- Primary backup path:
/srv/docker/unifi-controller/config
Example archive:
sudo tar czf /srv/docker/unifi-controller-backup-$(date +%F).tar.gz \
  -C /srv/docker/unifi-controller configcd /srv/docker/unifi-controller
docker compose pull
docker compose up -d
Allow time for schema migrations after major version upgrades.
docker compose down
sudo tar czf /srv/docker/unifi-controller-archive-$(date +%F).tar.gz \
  -C /srv/docker/unifi-controller config
sudo rm -rf /srv/docker/unifi-controller
sudo rm -rf /srv/logs/unifi-controller
Remove any reverse proxy or DNS references as required.
- unifi-db andunifi-controller containers are both running
- Both services are attached to the unifi-net Docker network
- UniFi UI accessible via https://<docker-host-ip>:8443
- Devices adopted and reporting connected
- Backups tested (/srv/docker/unifi-controller/config )
- Upgrade workflow confirmed functional (docker compose pull && docker compose up -d )
When these conditions are met, the UniFi Network Application is structured, reproducible, and maintainable within the Docker Host Service Layer.

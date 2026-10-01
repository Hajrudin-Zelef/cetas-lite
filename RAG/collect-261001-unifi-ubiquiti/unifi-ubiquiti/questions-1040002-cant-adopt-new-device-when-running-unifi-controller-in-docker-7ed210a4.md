---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1040002-cant-adopt-new-device-when-running-unifi-controller-in-docker-7ed210a4
title: "network_mode: \"host\""
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1040002-cant-adopt-new-device-when-running-unifi-controller-in-docker--7ed210a4.md
source_anchor: ""
source_lines: [1, 49]
sha256: 878379acb10d08387268d5a999b6b07856adc792f573082d38d9daedf9fc3b47
---

# network_mode: "host"

I'm trying to host the Unifi Controller in a docker container but I can't adopt a new device (access point) unless I set networking mode to "host". I have followed several guides and opened all ports listed here. It's possible to manage already connected devices and to find new devices but adopting them results in a never ending "adopting" status.
The first configuration is based on jacobalberty/unifi:
version: "3.7"
services:
  controller:
    image: jacobalberty/unifi:arm32v7
    restart: unless-stopped
    init: true
    volumes:
      - ./data:/unifi
    environment:
      TZ: "Europe/Stockholm"
    ports:
      - "3478:3478/udp"
      - "5514:5514/udp"
      - "8080:8080"
      - "8443:8443"
      - "8843:8843"
      - "8880:8880"
      - "6789:6789"
      - "27117:27117"
      - "5656-5699:5656-5699/udp"
      - "1900:1900"
      - "10001:10001/udp"
#    network_mode: "host"
The other is a direct copy of linuxserver.io example:
version: "2.1"
services:
  unifi-controller:
    image: linuxserver/unifi-controller
    container_name: unifi-controller
    environment:
      - PUID=1000
      - PGID=1000
      - MEM_LIMIT=1024M #optional
    volumes:
      - ./data:/config
    ports:
      - 3478:3478/udp
      - 10001:10001/udp
      - 8080:8080
      - 8443:8443
      - 1900:1900/udp #optional
      - 8843:8843 #optional
      - 8880:8880 #optional
      - 6789:6789 #optional
      - 5514:5514 #optional
    restart: unless-stopped
The docker container is running on a Raspberry Pi connected to wired network. The host is running other containers but there are no conflicting ports. If there were port conflicts it shouldn't be possible to start the container anyway.

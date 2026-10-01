---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1074787-cant-access-webui-for-unifi-controller-locally-running-in-dock-dc60dbb9
title: "- 6789:6789 #mobile throughput test"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["throughput"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1074787-cant-access-webui-for-unifi-controller-locally-running-in-dock-dc60dbb9.md
source_anchor: ""
source_lines: [1, 60]
sha256: 88a3803bbc221b2224ea9e37c160ecdbbdef95c7b2e13ff1b9c7c378df3ab3ea
---

# - 6789:6789 #mobile throughput test

I am able to access other docker containers which are hosted on my server via their web interfaces on various ports, but Unifi Controller is giving me "refused to connect" type errors.
This is the output of docker ps
    CONTAINER ID   IMAGE                                  COMMAND         CREATED       STATUS                        PORTS                                                                                                                                                                                                                                                                                                                                           NAMES
175a359de1d2   ghcr.io/linuxserver/unifi-controller   "/init"         6 hours ago   Up About a minute             0.0.0.0:1900->1900/udp, :::1900->1900/udp, 0.0.0.0:3478->3478/udp, :::3478->3478/udp, 0.0.0.0:8080-8081->8080-8081/tcp, :::8080-8081->8080-8081/tcp, 0.0.0.0:8843->8843/tcp, :::8843->8843/tcp, 0.0.0.0:5514->5514/udp, :::5514->5514/udp, 0.0.0.0:10001->10001/udp, :::10001->10001/udp, 0.0.0.0:8880->8880/tcp, :::8880->8880/tcp, 8443/tcp   unifi-controller
This is my docker-compose.yml entry for the controller
  unifi-controller:
    container_name: unifi-controller
    image: ghcr.io/linuxserver/unifi-controller
    restart: unless-stopped
    environment:
      - PUID=1000
      - PGID=1000
      - TZ=Europe/London
    ports:
      - 3478:3478/udp
      - 10001:10001/udp
      - 8080:8080
      - 8081:8081
      - 8843:8843
      - 8880:8880
#      - 6789:6789 #mobile throughput test
      - 1900:1900/udp
      - 5514:5514/udp
    volumes:
      - "~/docker-services/unifi/config:/config"
There are some errors in the docker logs but I don't know what they mean or how to fix them. They seem relatively generic when I've searched for them - not entirely sure what the fix is.
:0 TCP candidates not supported yet
:0 Permanent error code on channel bind request: 403 - Forbidden IP
:0 All TURN channel bindings failed
:0 Invalid channel id: 8000f; SID: 0x000f; _maxSctpChannels: 0x0400
:0 Invalid channel id: c000f; SID: 0x000f; _maxSctpChannels: 0x0400
:0 TCP candidates not supported yet
:0 Permanent error code on channel bind request: 403 - Forbidden IP
:0 All TURN channel bindings failed
:0 STUN id:   4; fd: 217 172.18.0.5:42921 -> 192.168.1.222:50184 (eth0) DTLS id: 9 (192.168.1.235:42921) (PEER) timed out
:0 webRtcId 1 terminated with code: (-2147090409) WebRTC connection interrupted from far side
:0 SSL_read permanently failed: (6)
:0 Unable to do SSL I/O
:0 webRtcId 2 terminated with code: (-2147090409) WebRTC connection interrupted from far side
:0 TCP candidates not supported yet
:0 Permanent error code on channel bind request: 403 - Forbidden IP
:0 All TURN channel bindings failed
:0 SSL_read permanently failed: (6)
:0 Unable to do SSL I/O
:0 webRtcId 3 terminated with code: (-2147090409) WebRTC connection interrupted from far side
:0 TCP candidates not supported yet
:0 Permanent error code on channel bind request: 403 - Forbidden IP
:0 All TURN channel bindings failed
:0 SSL_read permanently failed: (6)
:0 Unable to do SSL I/O
:0 webRtcId 4 terminated with code: (-2147090409) WebRTC connection interrupted from far side
:0 TCP candidates not supported yet
:0 Permanent error code on channel bind request: 403 - Forbidden IP
:0 All TURN channel bindings failed
:0 SSL_read permanently failed: (6)
:0 Unable to do SSL I/O
:0 webRtcId 5 terminated with code: (-2147090409) WebRTC connection interrupted from far side
I've tried restarting the container, restarting docker and restarting the server. Also ruled out different browsers just in case (Chrome, Edge, Firefox). This was working previously but nothig has changed.
When connected to the same network on my phone I can access the unifi controller by mobile app. I can also access it from https://network.unifi.ui.com
For reference, other applications such as nginx are available on the ports specified in the docker-compose file. I can't work out why it's only unifi-controller that's being difficult.

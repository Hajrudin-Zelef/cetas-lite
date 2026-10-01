---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1085063-docker-swarm-over-ipsec-vti-unifi-site-to-site-aacf4c08
title: "questions-1085063-docker-swarm-over-ipsec-vti-unifi-site-to-site-aacf4c08"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1085063-docker-swarm-over-ipsec-vti-unifi-site-to-site-aacf4c08.md
source_anchor: ""
source_lines: [1, 19]
sha256: 1336f550cadf7aa7a1edf7121105a62bf5ceed8e9dc7c87359d801a5cd0ff18b
---

# questions-1085063-docker-swarm-over-ipsec-vti-unifi-site-to-site-aacf4c08

I have setup a Docker Swarm using 3 servers;
Master node
Home - local IP 192.168.178.250
Worker node 1
Remote - local IP 10.10.10.200
Worker node 2
Remote - local IP 10.10.10.250
Both sites have a Unifi gateway (USG Pro/USG) and I've setup a site-to-site connection between them (auto settings, IPSec VTI)
I was able to add the 2 nodes using join tokens from the master since both worker nodes can access 192.168.178.250 via the VPN tunnel, and I've been successful in setting up multiple containers. However, it seems that the connection keeps timing out within the ingress(?) network.
Here's a chart explaining my setup
For example, I have a UptimeRobot rule that checks the "NGINX (hello world)" availability. The NGINX container is available through the "traefik_default" network and via Traefik linked to a domain.
CURL attempt 1: https://nginx.domain.tld/ - 200 ok
CURL attempt 2: https://nginx.domain.tld/ - 200 ok
CURL attempt 3: https://nginx.domain.tld/ - 504 Gateway timeout
CURL attempt 4: https://nginx.domain.tld/ - 504 Gateway timeout
CURL attempt 5: https://nginx.domain.tld/ - 504 Gateway timeout
While curl on the public domain gives a 504 Gateway Timeout the direct IP (over the VPN) seems to be reachable http://10.10.10.200 - 200 Ok
At first I figured it might be Traefik, but when I explore the portainer interface than 2/3 times it can't load services or containers where the log is saying something like: Unable to connect to portainer agent on 10.10.10.200/250
I'm having a really hard time to figure out where it is going wrong, is there a different VPN type required to get a more stable setup, or is this even a viable solution at all?

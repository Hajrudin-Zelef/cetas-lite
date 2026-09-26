---
id: collect-260926-mikrotik/mikrotik/questions-1157942-container-in-ipvlan-l3-network-cannot-ping-its-dockers-host-co-2c62ce40
title: "questions-1157942-container-in-ipvlan-l3-network-cannot-ping-its-dockers-host-co-2c62ce40"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/vlan-bridge/questions-1157942-container-in-ipvlan-l3-network-cannot-ping-its-dockers-host-co-2c62ce40.md
source_anchor: ""
source_lines: [1, 22]
sha256: 77c0acb96807d757087b32b0a37b728cf7dae935191fd80e16993f7f09fb815f
---

# questions-1157942-container-in-ipvlan-l3-network-cannot-ping-its-dockers-host-co-2c62ce40

Container in ipvlan L3 network cannot ping its docker's host, "connection status" is "invalid" in Mikrotik firewall, but can ping everyone else - Server Fault
Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
Created, entered into container and did some pings:
docker run -it --rm --network ipvlan alpine /bin/sh
Ping results:
ok: google.com
ok: 10.9.0.1 (Mikrotik)
ok: 10.9.0.2 (another computer in network)
ok: 10.21.0.2 (another container in ipvlan)
not ok: 10.9.0.3 (Docker host)
On Mikrotik there is a firewall rule:
- chain: forward
- connection state: invalid
- action: drop
To make ping work I need either disable this rule or add another one allowing invalid connection state when source address is 10.9.0.3 and destination address is 10.21.0.0/16.
It's a security constraint of the IPVLAN and MACVLAN kernel drivers. There's a note in Docker documentation about the IPVLAN driver:
NOTE: the containers can NOT ping the underlying host interfaces as they are intentionally filtered by Linux for additional isolation.
In fact, it concerns not just pings, but the entire communication with a host.
Another discussion on GitHub points out the same thing but discloses more details:
Note: In both Macvlan and Ipvlan you are not able to ping or communicate with the default namespace IP address. For example, if you create a container and try to ping the Docker host's eth0 it will not work. That traffic is explicitly filtered by the kernel modules themselves to offer additional provider isolation and security.
The default namespace is not reachable per ipvlan design in order to isolate container namespaces from the underlying host.
l3withl3s?l3sdoesn't help.

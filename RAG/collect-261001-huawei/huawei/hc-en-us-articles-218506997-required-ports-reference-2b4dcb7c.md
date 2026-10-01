---
id: collect-261001-huawei/huawei/hc-en-us-articles-218506997-required-ports-reference-2b4dcb7c
title: "hc-en-us-articles-218506997-required-ports-reference-2b4dcb7c"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/hc-en-us-articles-218506997-required-ports-reference-2b4dcb7c.md
source_anchor: ""
source_lines: [1, 151]
sha256: f1d474bdca01785a9ccab3638e77e564bded26295f4cbc860b6e7129db4df008
---

# hc-en-us-articles-218506997-required-ports-reference-2b4dcb7c

The following tables list UDP and TCP ports used by UniFi applications and services. This reference is especially useful for deployments using self-hosted UniFi Network Servers, third-party gateways, or restrictive firewalls.
In a full UniFi deployment with native gateways, these ports are opened automatically.
Remote Management
Protocol
Port
Direction
Usage
TCP/UDP
53
Both
DNS lookups for remote access, updates, and Guest Portal redirection (also required for UniFi Network)
UDP
123
Egress
NTP (time sync). Required for establishing secure connections
UDP
3478
Both
STUN for remote access (also required for UniFi Network)
TCP
443
Both
Remote Access service, application GUI/API access via web browser (also required for UniFi Network)
TCP
8883
Egress
Remote Access service
TCP
5349
Ingress
Remote access support
UniFi Network
Protocol
Port
Direction
Usage
TCP/UDP
53
Both
DNS lookups for Guest Portal redirection and updates (also required for Remote Management)
UDP
3478
Both
STUN for device adoption and communication (also required for Remote Management)
TCP
5671
Ingress
Traffic Flow logging for UXGs adopted on L2 or L3 networks.
TCP
8080
Ingress
Device and application communication. Some U6 APs may have this port open on device for UP-Sense communication.
TCP
8443
Ingress
Application GUI/API (on UniFi Console)
TCP
8880 – 8882
Ingress
Hotspot portal redirection (HTTP)
TCP
8843
Ingress
Hotspot portal redirection (HTTPS)
TCP
8444
Ingress
Secure Portal for Hotspot
TCP
6789
Ingress
UniFi mobile speed test
TCP
27117
Ingress
Local database communication
TCP
28082
Ingress
Device packet capture and support file downloads
UDP
10001
Ingress
Device discovery during adoption
UDP
10101
Ingress
Client fingerprinting information
UDP
1900
Ingress
L2 discovery (“Make application discoverable on L2 network”)
UDP
5514
Ingress
Remote syslog capture
TCP/UDP
22
Both
SSH access (manual management, not used by default)
TCP
443
Both
Application GUI/API access via web browser (also required for Remote Management)
UDP
20100-22100
Both
Site Magic SD-WAN tunnel port range (if all ports are in use, it will then try another free port above 20000, then try above/below this range).
UniFi Protect
Protocol
Port
Direction
Usage
TCP
7441
Ingress
Outgoing RTSPS streams
TCP
7442
Both
WebSocket server for device communication
TCP
7443
Both
REST API (HTTPS)
TCP
7444
Both
WebSocket server for camera communication
TCP
7445
Ingress
Outgoing Protect streams
TCP
7447
Ingress
Outgoing RTSP streams
TCP
7550
Ingress
Camera streams
TCP
7552
Both
SSL camera connections
TCP
7888
Both
TCP Bridge
Stacked NVRs Only (MSR/MSP)
These ports are only required if using physically stacked NVRs. They should be added in addition to the ports above.

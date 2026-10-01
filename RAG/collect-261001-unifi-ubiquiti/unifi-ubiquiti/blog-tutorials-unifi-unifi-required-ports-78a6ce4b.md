---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/blog-tutorials-unifi-unifi-required-ports-78a6ce4b
title: "blog-tutorials-unifi-unifi-required-ports-78a6ce4b"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/blog-tutorials-unifi-unifi-required-ports-78a6ce4b.md
source_anchor: ""
source_lines: [1, 83]
sha256: 89351fca1351fbe3460ae3f27116ae463670906e97a9d5ad4597b02c18879129
---

# blog-tutorials-unifi-unifi-required-ports-78a6ce4b

Tutorial
Required Ports for UniFi
Full reference of UDP and TCP ports used by UniFi Network, Protect, Connect, Access, and Remote Management — plus how to change defaults.
 Summary This reference lists every default UDP and TCP port used by UniFi Remote Management, UniFi Network, UniFi Protect, UniFi Connect, and UniFi Access. It’s the document you want when self-hosting a UniFi Network server on Windows, macOS, or Linux, when running behind a third-party gateway, or when a hardened firewall is dropping traffic between the console and the rest of the deployment. The last section covers how to change default ports on a self-hosted Network server through system.properties.
Which ports does UniFi require?
UniFi is a family of applications — Remote Management, Network, Protect, Connect, and Access — and each one uses a specific set of UDP and TCP ports to talk to devices, to the cloud, and to client browsers. In typical deployments with UniFi gateways, the right ports open automatically through the integrated firewall. Administrators using third-party gateways, self-hosted Network servers on Windows/macOS/Linux, or hardened firewalls must allow these ports explicitly or the corresponding services break in subtle ways — failed device adoption, broken Guest Portal redirection, video streams that don’t render, or the console refusing to remote-access from outside.
The tables below group every port by UniFi application. Direction is given from the perspective of the UniFi Console or self-hosted server: Ingress means traffic arriving at the server, Egress means traffic leaving toward an external service, Both means the port handles bidirectional flows.
Remote Management ports
Remote management lets administrators reach UniFi devices over the internet without exposing the console directly. These ports support remote access, DNS resolution, and secure communication.
| Protocol & port | Direction | Usage | 
|---|---|---|
| TCP/UDP 53 | Both | DNS lookups for remote access, updates, and Guest Portal redirection (shared with Network) | 
| UDP 123 | Egress | NTP time sync — required for secure connections | 
| UDP 3478 | Both | STUN service for remote access (shared with Network) | 
| TCP 443 | Both | Remote Access service and web GUI/API (shared with Network) | 
| TCP 8883 | Egress | Remote Access service | 
| TCP 5349 | Ingress | Remote access support | 
UniFi Network ports
UniFi Network is the central platform for UniFi switches, routers, and Wi-Fi access points. These ports handle device adoption, controller communication, and management.
| Protocol & port | Direction | Usage | 
|---|---|---|
| TCP/UDP 53 | Both | DNS for Guest Portal redirection and updates | 
| UDP 3478 | Both | STUN for device adoption and communication | 
| TCP 8080 | Ingress | Device and application communication | 
| TCP 8443 | Ingress | Application GUI/API on the UniFi Console | 
| TCP 8880–8882 | Ingress | Hotspot portal redirection (HTTP) | 
| TCP 8843 | Ingress | Hotspot portal redirection (HTTPS) | 
| TCP 8444 | Ingress | Secure portal for Hotspot | 
| TCP 6789 | Ingress | UniFi mobile speed test | 
| TCP 27117 | Ingress | Local database communication | 
| UDP 10001 | Ingress | Device discovery during adoption | 
| UDP 1900 | Ingress | Layer-2 discovery (“Make application discoverable on L2 network”) | 
| UDP 5514 | Ingress | Remote syslog capture | 
| TCP/UDP 22 | Both | SSH access (disabled by default) | 
| TCP 443 | Both | Application GUI/API via web browser | 
UniFi Protect ports
UniFi Protect handles video streaming and device communication for cameras and network video recorders.
| Protocol & port | Direction | Usage | 
|---|---|---|
| TCP 7441 | Ingress | Outgoing RTSPS streams | 
| TCP 7442 | Both | WebSocket server for device communication | 
| TCP 7443 | Both | REST API (HTTPS) | 
| TCP 7444 | Both | WebSocket server for camera communication | 
| TCP 7445 | Ingress | Outgoing Protect streams | 
| TCP 7447 | Ingress | Outgoing RTSP streams | 
| TCP 7550 | Ingress | Camera streams | 
| TCP 7552 | Both | SSL camera connections | 
| TCP 7888 | Both | TCP bridge | 
Stacked NVRs (MSR/MSP)
These ports are required only when physically stacking network video recorders. Open them in addition to the base Protect ports above.
| Protocol & port | Direction | Usage | 
|---|---|---|
| TCP 7446 | Both | Protect streams between consoles | 
| TCP 7451 | Both | Protect streams between consoles | 
| TCP 7600 | Both | Protect application communications | 
UniFi Connect ports
UniFi Connect integrates with Lutron lighting processors and other automation devices.
| Protocol & port | Direction | Usage | 
|---|---|---|
| UDP 2647 | Ingress | Lutron Processor discovery (HomeWorks QSX, HomeWorks Wireless, RadioRA3) | 
| UDP 5353 | Ingress | Lutron Processor discovery (HomeWorks QS, RadioRA2) | 
| TCP 18080 | Ingress | Application GUI/API on the UniFi Console | 
| TCP 18443 | Both | WebSocket server for device communication (HTTPS) | 
| TCP 18884 | Both | MQTT server for device communication | 
| TCP 18888 | Both | Internal Lutron proxy between Lutron Processor and UniFi Connect | 
UniFi Access ports
UniFi Access manages door controllers and readers.
| Protocol & port | Direction | Usage | 
|---|---|---|
| TCP 12812 | Both | MQTT server for device communication | 
| TCP 12442 | Both | WebSocket server (UCP4) for device communication | 
| TCP 12443 | Both | HTTPS server for device communication | 
| TCP 12445 | Both | Open API | 
| TCP 12478 | Both | WebRTC TURN server for device communication | 
Changing default ports on a self-hosted UniFi Network server
Default port assignments can be modified on self-hosted UniFi Network servers running on Windows, macOS, or Linux. UniFi Consoles (Cloud Keys, Dream Machines, or other embedded controllers) do not permit port changes.
- Shut down the running UniFi Network application.
- Open <unifi_base>/data/system.properties in a text editor.
- Add or modify the line for the port you want to change. For example, to change the shutdown port from 8081 to8089 , addunifi.shutdown.port=8089 .
- Restart the UniFi Network application.
Make sure the modified line has no leading or trailing whitespace, no inline comments (the file ignores any line starting with #), and no stray characters — anything off-spec causes the change to be silently ignored and the default to persist. For broader management context across multi-vendor fleets that include UniFi alongside MikroTik, see our MikroTik NAT setup guide and the MikroTik DNS over HTTPS tutorial.
Take the next step
UniFi is excellent within its own ecosystem; in mixed fleets it sits alongside MikroTik routers, Intelbras OLTs, and other vendors that each carry their own port lists, their own management consoles, and their own remote-access models. MKController unifies that picture: one dashboard across UniFi, MikroTik, Intelbras, and other vendors, with discovery via SNMP, LLDP, and CDP and secure remote access through outbound tunnels — no per-site port forwarding required.

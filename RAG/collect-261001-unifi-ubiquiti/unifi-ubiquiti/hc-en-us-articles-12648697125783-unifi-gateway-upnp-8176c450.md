---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-12648697125783-unifi-gateway-upnp-8176c450
title: "hc-en-us-articles-12648697125783-unifi-gateway-upnp-8176c450"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-12648697125783-unifi-gateway-upnp-8176c450.md
source_anchor: ""
source_lines: [1, 16]
sha256: 060957dfb8c91b13afe260c9f50e52577a09e3bb47a1cbb780becc53f0aa78f8
---

# hc-en-us-articles-12648697125783-unifi-gateway-upnp-8176c450

UniFi Gateway - UPnP
UPnP is a feature found in Internet section of your Network application that allows you to dynamically open and forward ports.
Requirements
- A UniFi gateway or UniFi Cloud Gateway.
How does it work?
UPnP automatically creates port forwarding and firewall rules to allow traffic through the firewall. Certain types of traffic, for example clients connecting to online game servers, may use this feature to improve the connection.
We recommend to keep UPnP disabled unless it is required in your network.
Frequently Asked Questions
Will using UPnP affect the security of my network?
Yes, UPnP automatically opens and forward ports through the firewall. Forwarding ports allows devices on the Internet to connect to these services running on the client devices.
Should I enable UPnP when playing online games?
We recommend to keep this feature disabled even when playing online games. Please refer to the documentation of the game publisher to determine whether UPnP is required.
Should I enable UPnP when watching online videos?
No, this is not required.
I have a service on my LAN that I want to access remotely. Should I enable UPnP?
Instead of using UPnP, use the Wireguard or Teleport VPN to access the client device instead.

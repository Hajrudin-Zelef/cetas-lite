---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-12648701398807-unifi-gateway-multicast-dns-74b23f99
title: "hc-en-us-articles-12648701398807-unifi-gateway-multicast-dns-74b23f99"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-12648701398807-unifi-gateway-multicast-dns-74b23f99.md
source_anchor: ""
source_lines: [1, 17]
sha256: 2754639a0d33ee0ef399bdec8a202090513abb773b12fe06c4cea7e4f8223d37
---

# hc-en-us-articles-12648701398807-unifi-gateway-multicast-dns-74b23f99

UniFi Gateway - Multicast DNS (mDNS) Proxy
Multicast DNS (mDNS) Proxy is a UniFi Gateway feature that allows multicast device discovery traffic to be rebroadcast across different networks.
Requirements
- An independent UniFi Gateway or UniFi Cloud Gateway
How does it work?
A Multicast DNS Proxy is active on the UniFi Gateway and forwards multicast traffic from devices between different networks (VLANs). Enable this feature when features like AirPlay, AirPrint, or Chromecast should be shared across different networks/VLANs.
There are three modes: "Auto", "Off", and "Custom". "Auto" automatically retransmits common mDNS traffic across all VLANs, "Off" completely stops the retransmission of mDNS traffic, and "Custom" allows for individual services to be rebroadcast across individually selected networks while others are dropped. When "Custom" is selected, the VLAN Scope can be selected, with all or specific services selectable.
In addition to a full suite of built in mDNS protocols, new ones can be added to the list in the format of _service._protocol.local. The service is usually the name of the feature associated with the traffic, while protocol represents whether TCP or UDP is used for communication. For example, _airplay._tcp.local. 
Frequently Asked Questions
1. What are the default settings for Multicast DNS Proxy?
2. Is there a limit to the amount of Multicast DNS networks?
| Gateway | Maximum mDNS Networks | 
| Express / UXG Lite | 2 | 
| UDR / UDM / UDR7 / Express 7 | 30 | 
| UXG Max / UXG Pro | 50 | 
| UDM Pro / UDM SE / UDM Pro Max / UDW / UCG Ultra / UCG Max / UCG Fiber / UXG Fiber | 50 | 
| EFG | 60 |

---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-18125733726615-pro-av-traffic-optimization-on-unifi-switches-6c7233d1
title: "hc-en-us-articles-18125733726615-pro-av-traffic-optimization-on-unifi-switches-6c7233d1"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["cyber", "latency"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-18125733726615-pro-av-traffic-optimization-on-unifi-switches-6c7233d1.md
source_anchor: ""
source_lines: [1, 70]
sha256: 7763a73fd10f9ccaed726b1c2169221cac3dbf00f5f9398431cef494dadd6490
---

# hc-en-us-articles-18125733726615-pro-av-traffic-optimization-on-unifi-switches-6c7233d1

Pro AV Traffic Optimization on UniFi Switches
Pro AV on UniFi Switches uses Quality of Service (QoS) to automatically optimize traffic for specific audio and video environments by matching and prioritizing latency sensitive traffic. Configure Pro AV on individual ports using the Port Manager.
For a full overview of UniFi’s Traffic and Policy Management capabilities, see here.
For a full overview of UniFi's Network and Cyber Security capabilities, see here.
Requirements
- UniFi Network, version 8.4.59 or later
- UniFi Switch supporting Pro AV, version 7.1.26 or later
- ECS-24-PoE
- ECS-48-PoE
- EAV-XG-24-PoE
- EAV-Fiber
- USW-Pro-XG-8-PoE
- USW-Pro-XG-10-PoE
- USW-Pro-XG-24-PoE
- USW-Pro-XG-48-PoE
- USW-Pro-Max-16
- USW-Pro-Max-16-PoE
- USW-Pro-Max-24
- USW-Pro-Max-24-PoE
- USW-Pro-Max-48
- USW-Pro-Max-48-PoE
- USW-EnterpriseXG-24
- USW-Enterprise-48-PoE
- USW-Enterprise-24-PoE
- USW-Pro-24
- USW-Pro-24-PoE
- USW-Pro-48
- USW-Pro-48-PoE
- USW-Pro-Aggregation
- USW-Pro-HD-24
- USW-Pro-HD-24-PoE
Audio and Video Environments
When configuring Pro AV, select the option that matches your setup and the port(s) where audio and video devices are connected. Ports are configure individually for Pro AV and the options do not need to match on the same switch. The following options are available:
- Dante Audio - Matches Audinate Dante audio, PTP and control traffic.
- Q-SYS Audio and Video - Matches Q-SYS audio, video, PTP and control traffic.
- Q-SYS Video with Dante Audio - Matches Q-SYS video with Audinate Dante audio, PTP and control traffic.
- Shure Audio - Matches Shure audio, PTP and control traffic.
- AES67 Audio - Matches AES67 audio, PTP and control traffic.
- SDVoE with AES67 Audio - Matches SDVoE with AES67 audio, PTP and control traffic .
- SDVoE with Dante Audio - Matches SDVoE with Audinate Dante audio, PTP and control traffic.
- NDI with AES67 Audio - Matches NDI with AES67 audio, PTP and control traffic.
- NDI with Dante Audio - Matches NDI with Audinate Dante audio, PTP and control traffic.
- Crestron NVX/NAX - Matches Crestron NVX and NAX AV and control traffic.
How is Traffic Prioritized when using Pro AV?
Traffic is prioritized by matching on specific DSCP and CoS values, protocols and UDP/TCP ports. Traffic is assigned to a queue which determines priority over other traffic.
IGMP Settings - Snooping, Fast Leave and Querier
It is recommended to enable IGMP Snooping and define a specific switch as the querier. IGMP Snooping can be enabled for specific VLANs using the Network settings. Refer to the table below for the settings for each environment.
| Environment | IGMP Snooping | Fast Leave | Querier | Unknown Multicast Traffic | 
| Dante | Required | Recommended | Must be Present | Multicast Router Ports | 
| Q-SYS | Required | Required (NV-32-H) | Must be Present | Multicast Router Ports | 
| Shure | Required | Recommended | Must be Present | Multicast Router Ports | 
| AES67 | Required | Recommended | Must be Present | Multicast Router Ports | 
| SDVoE | Required | Required | Must be Present | Multicast Router Ports | 
| Crestron | Required | Required | Must be Present | Multicast Router Ports | 
| NDI | Required | Recommended | Must be Present | Multicast Router Ports | 
| OneIP (Pulse-Eight) | Required | Recommended | Must be Present | Multicast Router Ports | 
It is recommended to manually configure the Querier Switch for each VLAN. Either automatically set the address or manually define it. The switch with the lowest address will be elected as querier.
Note: Fast Leave can also be referred to as Immediate Leave and Unknown Multicast Traffic can be referred to as Unregistered Multicast Flooding.
Unknown Multicast Traffic and Known Protocols
Unknown multicast traffic is traffic that does not match any IGMP membership reports. It is recommended to forward this traffic to Multicast Router Ports or drop it using the Forward Unknown Multicast Traffic option. Traffic in the 224.0.0.0/24 range is used for local multicast communication and is always flooded regardless of which option is configured.
Known protocols using 224.0.1.x addresses (PTP, H.323, NTP, SLP and SSDP) are also flooded with the Flood Known Protocols option. Disabling this option may lead to issues with AV traffic as PTP is no longer flooded.
Note: UniFi Switch version 7.2 is required for Forward Unknown Multicast Traffic to Multicast Router Ports and Flood Known Protocols. If an older version is used, unknown traffic is dropped and known protocols are flooded by default.
AV Traffic and mDNS
If mDNS is enabled, it is recommended to exclude the VLANs (networks) that IGMP Snooping is enabled on. This stops protocols such as PTP from being forwarded to different VLANs.
Custom DSCP
Refer to the instructions in UniFi QoS and Traffic Shaping.
Multicast and Link Aggregation
- Source IP address
- Destination IP address
- Ingress physical port number

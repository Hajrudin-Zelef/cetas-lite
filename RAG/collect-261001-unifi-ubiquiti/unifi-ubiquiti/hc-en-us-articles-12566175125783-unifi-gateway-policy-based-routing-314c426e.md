---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-12566175125783-unifi-gateway-policy-based-routing-314c426e
title: "hc-en-us-articles-12566175125783-unifi-gateway-policy-based-routing-314c426e"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["cyber", "kill switch", "latency"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-12566175125783-unifi-gateway-policy-based-routing-314c426e.md
source_anchor: ""
source_lines: [1, 58]
sha256: 21eaceb7fae43dec24c80b2f95937c96b2841346366396f20983936d6d2fa167
---

# hc-en-us-articles-12566175125783-unifi-gateway-policy-based-routing-314c426e

UniFi Gateway - Policy-Based Routing
Policy Based Routes allow you to flexibly direct traffic through specific network interfaces—such as a particular WAN port or a VPN tunnel—based on custom rules and conditions.
For a full overview of UniFi’s Traffic and Policy Management capabilities, see here.
For a full overview of UniFi's Network and Cyber Security capabilities, see here.
Requirements
- A Next-Gen UniFi gateway or UniFi Cloud Gateway
When To Use Policy-Based Routing
There are many possible applications of policy-based routing, however some of the common use-cases include:
- 
Split Traffic Across WANs: Direct specific types of traffic through different WAN interfaces (e.g., WAN1 vs. WAN2).
  - Send video conferencing (Zoom, Teams) through a dedicated, low-latency fiber WAN, while all other traffic goes through the default WAN.
  - Force guest VLAN traffic to go through a filtered or throttled WAN.
- 
Proxy Traffic Through VPN: Force certain devices or destinations to route through a VPN tunnel instead of the default WAN.
  - Send outbound traffic for specific services (e.g., CRM, HR portals) through a secure tunnel to a Zero-Trust Exchange (e.g., Zscaler or Cloudflare).
  - Use a VPN Client (e.g., NordVPN) to overcome geo-based content restrictions.
Configuring Policy-Based Routes
- 
Navigate to Policy-Based Routing Rules: Follow the path depending on your UniFi Network version:
  - Network 9.4: Settings > Policy Table > Create New Policy > Route
  - Network 9.3: Settings > Policy Engine > Policy-Based Routes > Create Route
- Create Name: Assign a name to the rule.
- Select an Interface: This is VPN tunnel, WAN (Internet), or local network interface specified traffic will be sent through.
- 
(Optional) Configure the Kill Switch Toggle
  - Enabled: If the specified interface goes down, traffic will stop completely.
  - Disabled: If the specified interface goes down, traffic will instead flow through the default interface, or the interface with the next highest priority according to other routing rules.
- 
Configure Source and Destination: Traffic meeting all specified criteria will be routed through the interface configured in (2).
 
  - 
Source 
    - Any: Applies to any traffic originating from inside the local network.
    - Device / Network: Specify certain networks (VLANs) and/or devices.
  - 
Destination
    - Any: Applies to any, no matter what its destination is.
    - IP: Specify certain IP addresses or IP address ranges, and ports associated with the destination traffic.
    - 
Domain: Specify one or more domains associated with the destination traffic.
Note: Requires the client devices to use the UniFi gateway as the DNS server.
    - Region: Specify one or more countries associated with the destination traffic.
- 
Source 
Examples
Sending Streaming Traffic through a VPN
If you want to send specific streaming traffic from your Apple TV to a VPN Client tunnel, create a Policy Based Route with the following options:
- Type: Specific Traffic
- Category: Domain Name
- Domain Name: Add one or more domains used by the streaming service
- Target: Apple TV
- Interface: VPN Client
Routing All Traffic from a Virtual Network to a Secondary WAN
If you want to send all traffic from a Virtual Network to the secondary WAN port, create a Policy Based Route with the following options:
- Type: All Traffic
- Target: Select the Virtual Network name
- Interface: WAN2
See WAN Failover and Load Balancing for more information.

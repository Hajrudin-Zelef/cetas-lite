---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-nat-and-port-forwar-1c9df1c0-1
title: "sase-and-sd-wan-mx-design-and-configure-configuration-guides-nat-and-port-forwar-1c9df1c0"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-nat-and-port-forwar-1c9df1c0.md
source_anchor: ""
source_lines: [1, 49]
sha256: c185d84b281b06265737ae6d646f2676f0e4c7e11dea199c79c0a4933506e9d6
---

# sase-and-sd-wan-mx-design-and-configure-configuration-guides-nat-and-port-forwar-1c9df1c0

Port Forwarding and NAT Rules on the MX
Servers behind a firewall often need to be accessible from the internet. You can accomplish this by implementing Port Forwarding, 1:1 NAT (Network Address Translation), or 1:Many NAT on the MX security appliance. This article discusses when it is appropriate to configure each one, how to configure each one, and their corresponding limitations.
Port Forwarding
Overview
Port forwarding takes specific TCP or UDP ports destined to an internet interface of the MX security appliance and forwards them to specific internal IPs. This is best for users who do not own a pool of public IP addresses. This feature can forward different ports to different internal IP addresses, allowing multiple servers to be accessible from the same public IP address.
Figure 1. Example of port forwarding configuration
Figure 2. Illustration of port forwarding configuration
Configuration
- Navigate to Security & SD-WAN > Configure > Firewall.
- Click Add a port forwarding rule to create a new port forward.
- Configure the following:
- Description: Provide description of the rule
- Uplink: Listen on the public IP of internet 1, internet 2, or both
- Protocol: TCP or UDP
- Public port: Destination port of the traffic that is arriving on the WAN
- LAN IP: Local IP address to which traffic will be forwarded
- Local port: Destination port of the forwarded traffic that will be sent from the MX to the specified host on the LAN; if you simply wish to forward the traffic without translating the port, this should be the same as the Public port
- Allowed remote IPs: Remote IP addresses or ranges that are permitted to access the internal resource via this port forwarding rule
Note:
- Ports can be listed individually, or as a range
- Port ranges must be hyphenated; a comma-separated list is not accepted
- When mapping a range of public ports to a range of local ports, the ranges must be the same length
    
  - e.g. 8000-8500 public must be mapped to 8000-8500 local
- 
    It is not possible to forward a single TCP or UDP port to multiple LAN devices using port forwarding
Additional Considerations
Forwarding rules with Overlaps
When configuring 1:1 NAT, 1:Many NAT, and port forwarding rules on an MX, be aware that overlaps between 1:1 or 1:Many NAT and the MX uplink IP have the potential to conflict with any port forwarding rules or MX-hosted services on the same ports.
Forwarding L2TP/IPsec UDP Ports
If a port forward for ports UDP 500 or 4500 to a specific server is configured, the MX will reroute all non-Meraki site-to-site and L2TP/IPsec client VPN traffic to the LAN IP specified in the port forward. This is discussed with greater detail in IPSec VPN Port Overlap with Manual Port Forwarding Rules
Forwarding TCP 443/80
If a port forward for ports 443 or 80 is configured, you may be unable to reach the local status page via the MX's WAN IP address.
Note: This does not affect LAN or site-to-site client ability to reach the local status page.
1:1 NAT
Overview
1:1 NAT is for users with multiple public IP addresses available for use and for networks with multiple servers behind an firewall, such as two web servers and two mail servers. 1:1 NAT mapping can only be configured with IP addresses that do not belong to the MX security appliance. It can also translate public IP addresses in different subnets than the WAN interface address if the ISP routes traffic for the subnet towards the MX interface. Each translation added is a one-to-one rule, which means traffic destined to the public IP address can only go to one internal IP address. Within each translation, a user can specify which ports will be forwarded to the internal IP. When adding ports for NAT, a range or comma-separated list of ports are both acceptable.
Figure 1. Example of 1:1 NAT configuration
Figure 2. Illustration of 1:1 NAT configuration
Configuration
- Navigate to Security & SD-WAN > Configure > Firewall
- Click Add a 1:1 NAT mapping to create a mapping
- Configure the following:
- Name: A descriptive name for the rule
- Public IP: The IP address that will be used to access the internal resource from the WAN
- LAN IP: The IP address of the server or device that hosts the internal resource that you wish to make available on the WAN
- Uplink: The physical WAN interface on which the traffic will arrive
- Allowed inbound connections: The ports this mapping will provide access on and the remote IPs that will be allowed access to the resource. To enable an inbound connection, click Allow more connections and enter the following information:
    

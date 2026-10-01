---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-nat-and-port-forwar-1c9df1c0-2
title: "sase-and-sd-wan-mx-design-and-configure-configuration-guides-nat-and-port-forwar-1c9df1c0"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-nat-and-port-forwar-1c9df1c0.md
source_anchor: ""
source_lines: [50, 107]
sha256: 77fc9d3ad5f3504628a64ccec5e75dda1db27ac4400f172ea237adae139f4419
---

# sase-and-sd-wan-mx-design-and-configure-configuration-guides-nat-and-port-forwar-1c9df1c0

  - Protocol: Choose from TCP, UDP, ICMP ping, or any
  - Ports: Enter the port or port range that will be forwarded to the host on the LAN; you can specify multiple ports or ranges separated by commas
  - Remote IPs: Enter the range of WAN IP addresses that are allowed to make inbound connections on the specified port or port range; you can specify multiple WAN IP ranges separated by commas
Note:
- You can move a configured rule up or down in the list by dragging the ✥ symbol. Click the X to delete the rule entirely.
- Creating a 1:1 NAT rule does not automatically allow inbound traffic to the public IP listed in the NAT mapping. By default, all inbound connections are denied. You will have to configure Allowed inbound connections as described above in order to allow the inbound traffic.
Additional Considerations
1:1 NAT and Multiple MX Uplinks
If the MX primary uplink is not the same as the 1:1 NAT uplink, outbound traffic from the 1:1 NAT LAN device will, by default, egress out of the MX primary uplink. To prevent asynchronous routing, an uplink preference that points to the same uplink configured for the 1:1 NAT can be set. This configuration option can be found under Security & SD-WAN > Configure > SD-WAN & traffic shaping > Flow preferences.
Example:
- 
    MX primary uplink is WAN 1
- 
    1:1 NAT maps to WAN 2 Uplink/IP
- 
    You want all outbound internet traffic sourced from 1:1 NAT LAN device to use WAN 2
1:1 NAT and Load Balancing
If the MX is configured to load balance traffic across multiple WAN interfaces, outbound traffic from the 1:1 NAT LAN device will, by default, egress out of both WAN interfaces. To prevent asynchronous routing, an uplink preference configuration can be created, as shown in the example above.
1:1 NAT and Content Filtering
When a 1:1 NAT rule is configured for a given LAN IP, that device's outbound traffic will be mapped to the public IP configured in the 1:1 NAT rule rather than the primary WAN IP of the MX. Exceptions may occur when the MX is running some content filtering features that involve its web proxy. In this circumstance, outbound web traffic initiated by the 1:1 NAT LAN device will use the primary uplink as normal.
Port forwarding/NAT rules and Inbound firewall rules
If the manual inbound firewall is enabled, port forwarding and NAT rule behavior will be affected. Please refer to the NAT Exceptions with Manual Inbound Firewall KB article for details on how inbound firewall rules will change and what actions you need to take.
Hairpin Routing
Traffic sourced from the LAN of the MX that is destined for the public IP configured in the port forwarding/1:1 NAT/1:Many NAT section will be routed to the private IP address associated with the configured mapping.
In this process, the MX will accept the packet on the LAN and rewrite the IPv4 header. The rewritten header will be sourced from the MX's IP/MAC, or layer 3 interface in which the destination client resides, while also being destined for the private IP/MAC of the client mapped to the 1:1 NAT.
This practice does add complexities and may also be achieved with more ease via static DNS records where applicable.
In some cases, 1:1 NAT translation will not work properly immediately after installing a new MX or when using load balancing. Special considerations should be taken when configuring 1:1 NAT rules with Uplink preferences and multiple public IP addresses.
Example Configurations
Basic (Insecure) Configuration
A basic but insecure 1:1 NAT configuration can be set up to forward all traffic to the internal client. This should be configured when a 1:1 NAT needs to be made on a quick notice, but is not recommended due to security reasons. When all ports are forwarded to a client, attackers using a port scanner can target vulnerable services or gain access to the internal server.
Figure 1. Example of insecure 1:1 NAT configuration
Figure 2. Illustrating an insecure 1:1 NAT configuration
Detailed (Secure) Configuration
A more advanced configuration should include multiple rules and utilize a secondary uplink to provide redundancy for the web server. If one of the uplinks goes down, the secondary uplink is still in place to provide remote connectivity to the internal server. 1:1 NAT rules should also be configured to restrict specific remote IP addresses' access to specific services such as RDP.
Figure 1. Example of a secure 1:1 NAT configuration
Figure 2. Illustrating an example secure 1:1 NAT configuration
1:Many NAT
Overview
A 1:Many NAT configuration allows an MX to forward traffic from a configured public IP to internal servers. Like 1:1 NAT, 1:Many NAT can use a public IP address from the MX’s own WAN subnet or from a different routed subnet than the WAN interface IP, provided that the ISP routes traffic for that subnet to the MX interface. The public IP used for NAT must be different from the MX’s own WAN interface IP address. Unlike a 1:1 NAT rule, 1:Many NAT allows a single public IP to translate to multiple internal IPs on different ports. For each 1:Many IP definition, a single public IP must be specified, then multiple port forwarding rules can be configured to forward traffic to different devices on the LAN on a per-port basis. As with 1:1 NAT, a 1:Many NAT definition cannot use an IP address that belongs to the MX.
Figure 1. Example of 1:Many NAT configuration
Figure 2. Illustration of 1:Many NAT configuration
Configuration
- Navigate to Security & SD-WAN > Configure > Firewall
- Click Add a 1:1 NAT mapping to create a mapping
- Configure the following:
- Public IP: The IP address that will be used to access the internal resource from the WAN
- Uplink: The physical WAN interface on which the traffic will arrive
- Rules: A 1:Many NAT entry will be created with one associated forwarding rule. To add additional rules, click Add a port forwarding rule under the existing rule or rules for a particular 1:Many entry.
    
  - Description: A description of the rule
  - Protocol: TCP or UDP
  - Public port: Destination port of the traffic that is arriving on the WAN
  - LAN IP: Local IP address to which traffic will be forwarded
  - Local port: Destination port of the forwarded traffic that will be sent from the MX to the specified host on the LAN. If you simply wish to forward the traffic without translating the port, this should be the same as the Public port
  - Allowed remote IPs: Remote IP addresses or ranges that are permitted to access the internal resource via this port forwarding rule
Troubleshooting
For information on troubleshooting issues with port forwarding and NAT rules, please refer to this article.
In a 1:1 NAT, outbound traffic from a LAN device is expected to come from the public IP associated with that device, as configured in the 1:1 NAT rule. However, in a 1:Many NAT, outbound traffic initiated from the LAN device will be sourced from the MX's default public IP. Only the inbound, return communication for the 1:Many NAT traffic will use the rule's associated public IP.

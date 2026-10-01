---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-firewall-and-traffi-746f678a-3
title: "sase-and-sd-wan-mx-design-and-configure-configuration-guides-firewall-and-traffi-746f678a"
domain: general-networking
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-firewall-and-traffi-746f678a.md
source_anchor: ""
source_lines: [121, 167]
sha256: 90626cffee3abd4845ec26f642a86037daad759032b2dc0a829164768ad2a3a3
---

# sase-and-sd-wan-mx-design-and-configure-configuration-guides-firewall-and-traffi-746f678a

Use this option to forward traffic destined for the WAN IP of the MX on a specific port to any IP address within a local subnet or VLAN. Click Add a port forwarding rule to create a new port forward. You need to provide the following:
- 
    Description: A description of the rule.
- Uplink: Listen on the public IP of Internet 1, Internet 2, or both.
- Protocol: TCP or UDP.
- Public port: Destination port of the traffic that is arriving on the WAN.
- LAN IP: Local IP address to which traffic will be forwarded.
- Local port: Destination port of the forwarded traffic that will be sent from the MX to the specified host on the LAN. If you simply wish to forward the traffic without translating the port, this should be the same as the Public port.
- Allowed remote IPs: Remote IP addresses or ranges that are permitted to access the internal resource via this port forwarding rule.
You can also create a port forwarding rule to forward a range of ports. However, the range configured in the Public port field must be the same length as the range configured in the Local port field. The public ports will be forwarded to their corresponding local ports within the range. For instance, if you forward TCP 223-225 to TCP 628-630, port 223 would be translated to 628, port 224 would be translated to 629, and port 225 would be translated to 630.
1:1 NAT
Use this option to map an IP address on the WAN side of the MX (other than the WAN IP of the MX itself) to a local IP address on your network. Click Add a 1:1 NAT mapping to create a new mapping. You need to provide the following:
- Name: A descriptive name for the rule
- Public IP: The IP address that will be used to access the internal resource from the WAN.
- LAN IP: The IP address of the server or device that hosts the internal resource that you wish to make available on the WAN.
- Uplink: The physical WAN interface on which the traffic will arrive.
- Allowed inbound connections: The ports this mapping will provide access on, and the remote IPs that will be allowed access to the resource. To enable an inbound connection, click Allow more connections and enter the following information:
    
  - Protocol: Choose from TCP, UDP, ICMP ping, or any.
  - Ports: Enter the port or port range that will be forwarded to the host on the LAN. You can specify multiple ports or ranges separated by commas.
  - Remote IPs: Enter the range of WAN IP addresses that are allowed to make inbound connections on the specified port or port range. You can specify multiple WAN IP ranges separated by commas.
You can move a configured rule up or down in the list using the + icon. Click the X to remove it entirely.
Creating a 1:1 NAT rule does not automatically allow inbound traffic to the public IP listed in the NAT mapping. By default all inbound connections are denied. You will have to configure Allowed inbound connections as described above to allow the inbound traffic.
1:Many NAT
1:Many NAT, also known as Port Address Translation (PAT), is more flexible than 1:1 NAT. It allows you to specify one public IP that has multiple forwarding rules for different ports and LAN IPs. To add a 1:Many NAT listener IP, click Add 1:Many IP.
- Public IP: The IP address that will be used to access the internal resource from the WAN.
- Uplink: The physical WAN interface on which the traffic will arrive.
A 1:Many NAT entry will be created with one associated forwarding rule. To add additional rules, click Add a port forwarding rule under the existing rule or rules for a particular 1:Many entry.
- 
    Description: A description of the rule.
- Protocol: TCP or UDP.
- Public port: Destination port of the traffic that is arriving on the WAN.
- LAN IP: Local IP address to which traffic will be forwarded.
- Local port: Destination port of the forwarded traffic that will be sent from the MX to the specified host on the LAN. If you simply wish to forward the traffic without translating the port, this should be the same as the Public port.
- Allowed remote IPs: Remote IP addresses or ranges that are permitted to access the internal resource via this port forwarding rule.
Bonjour Forwarding
Use this feature to allow Bonjour to work between VLANs. Click Add a Bonjour forwarding rule to create a new forwarding rule.
- Description: Specify a name for the rule.
- Service VLANs: Select one or more VLANs where network services are running. Bonjour requests from the Client VLANs will be forwarded to these VLANs.
- Client VLANs:  Select one or more VLANs from which client Bonjour requests can originate. Requests on these VLANs will be forwarded to the Service VLANs. The list of services that can be forwarded include:
    
  - All services
  - AirPlay
  - Printers
  - AFP (Apple file sharing)
  - Scanners
  - iChat

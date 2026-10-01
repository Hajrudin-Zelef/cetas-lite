---
id: collect-261001-huawei/huawei/sase-and-sd-wan-mx-design-and-configure-deployment-guides-vpn-concentrator-deplo-9458fcc9-4
title: "sase-and-sd-wan-mx-design-and-configure-deployment-guides-vpn-concentrator-deplo-9458fcc9"
domain: huawei
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["datacenter"]
source: docs/RAG/collect-261001-huawei/sase-and-sd-wan-mx-design-and-configure-deployment-guides-vpn-concentrator-deplo-9458fcc9.md
source_anchor: ""
source_lines: [208, 278]
sha256: 46fe1bc386cf655fb86951771908bc52d71d456d7e45666496ffcc043412b888
---

# sase-and-sd-wan-mx-design-and-configure-deployment-guides-vpn-concentrator-deplo-9458fcc9

Cisco Meraki's AutoVPN technology leverages a cloud-based registry service to orchestrate VPN connectivity. In order for successful AutoVPN connections to establish, the upstream firewall must allow the VPN concentrator to communicate with the VPN registry service. The relevant destination ports and IP addresses can be found under the Help > Firewall info page in the Dashboard.
Uplink Health Monitoring
The MX also performs periodic uplink health checks by reaching out to well-known Internet destinations using common protocols. The full behavior is outlined here. In order to allow for proper uplink monitoring, the following communications must also be allowed:
- ICMP to 8.8.8.8 (Google's public DNS service)
- HTTP port 80
- DNS to the MX's configured DNS server(s)
Appendix
Appendix 1: One-armed concentrator operation
A WAN appliance operating in one-armed concentrator mode sends and receives traffic on a singular interface. This interface will always be the the first Internet or WAN port on the unit. A secondary port is not supported when deployed as a VPN concentrator.
It is important to understand the flow of traffic sent across an AutoVPN tunnel while the MX is acting as a one-armed concentrator. In the following scenario we have a host at a branch location trying to load a webpage located in the datacenter, over the site-to-site VPN.
- 
    The client sends traffic to the private address of the web server to its default gateway, the MX (in Routed mode) at the branch location.
- 
    The branch MX will look at its routing table and see that the destination IP address is contained within a subnet that is accessible over the Meraki AutoVPN.
- 
    The branch MX encrypts and encapsulates the data from the client and sends a packet source from its WAN interface, destined for the public IP address and port of the one-armed concentrator at the datacenter that was learned through the VPN registry.
- 
    This traffic is routed across the Internet to the edge of the datacenter.
- 
    The edge of the datacenter will NAT the traffic into a private address and send the traffic to the IP address of the one-armed concentrator.
- 
    The traffic will traverse the network internal to the datacenter and arrive at the one-armed concentrator. The MX will then decrypt and de-encapsulate the traffic and forward the original packet (sent by the client from the branch) upstream.
- 
    The upstream datacenter infrastructure routes traffic to the server.
- 
    The server receives the client traffic and sends a response to the client.
- 
    The response is then routed back through the internal datacenter network to the MX acting as a one-armed concentrator.
- 
    Upon receiving this response, the one-armed concentrator sees that the destination IP address is contained within a subnet that is accessible over the site-to-site VPN, looks up the contact information for the corresponding AutoVPN peer, encapsulates and encrypts the data, and sends the response on the wire.
- 
    The response, destined for the public IP and AutoVPN port of the branch MX, is then routed through the datacenter and NAT’ed out to the Internet.
- 
    The packet is then routed through the Internet to the branch MX.
- 
    The Branch MX receives the response, decrypts, de-encapsulates, and forwards the server's response downstream.
- 
    The response then traverses the internal branch network and is received by the client device.
Appendix 2: Routed mode concentrator operation
A WAN appliance operating as a Routed mode concentrator sends and receives encapsulated and encrypted traffic on its WAN interface and sends and receives de-encapsulated and decrypted traffic on its LAN interface.
It is important to understand the flow of traffic sent across an AutoVPN tunnel while the MX is acting as a Routed mode concentrator. In the following scenario we have a host at a branch location trying to load a webpage located in the datacenter, over the site-to-site VPN.
- 
    The client sends traffic to the private address of the web server to its default gateway, the MX (in Routed mode) at the branch location.
- 
    The branch MX will look at its routing table and see that the destination IP address is contained within a subnet that is accessible over the Meraki AutoVPN.
- 
    The branch MX encrypts and encapsulates the data from the client and sends a packet source from its WAN interface, destined for the public IP address and port of the Routed mode concentrator at the datacenter that was learned through the VPN registry.
- 
    This traffic is routed across the Internet to the edge of the datacenter.
- 
    The edge of the datacenter will NAT the traffic into a private address and send the traffic to the IP address of the Routed mode concentrator.
- 
    The traffic will traverse the network internal to the datacenter and arrive at the Routed mode concentrator's WAN interface. The MX will then decrypt and de-encapsulate the traffic.
- 
    The concentrator will look at its routing table and forward the original packet (sent by the client from the branch) downstream based on the most specific route to the destination address.
- 
    The downstream datacenter infrastructure routes traffic to the server.
- 
    The server receives the client traffic and sends a response to the client.
- 
    The response is then routed back through the internal datacenter network to the MX acting as a Routed mode concentrator.
- 
    Upon receiving this response, the Routed mode concentrator sees that the destination IP address is contained within a subnet that is accessible over the site-to-site VPN, looks up the contact information for the corresponding AutoVPN peer, encapsulates and encrypts the data, and sends the response on the wire out its WAN interface.
- 
    The response, destined for the public IP and AutoVPN port of the branch MX, is then routed through the datacenter and NAT’ed out to the Internet.
- 
    The packet is then routed through the Internet to the branch MX.
- 
    The Branch MX receives the response, decrypts, de-encapsulates, and forwards the server's response downstream.
- 
    The response then traverses the internal branch network and is received by the client device.

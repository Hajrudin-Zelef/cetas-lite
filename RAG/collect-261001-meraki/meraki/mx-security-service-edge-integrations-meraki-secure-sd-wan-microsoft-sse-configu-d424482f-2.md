---
id: collect-261001-meraki/meraki/mx-security-service-edge-integrations-meraki-secure-sd-wan-microsoft-sse-configu-d424482f-2
title: "mx-security-service-edge-integrations-meraki-secure-sd-wan-microsoft-sse-configu-d424482f"
domain: meraki
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/mx-security-service-edge-integrations-meraki-secure-sd-wan-microsoft-sse-configu-d424482f.md
source_anchor: ""
source_lines: [101, 175]
sha256: 97c2788e4fea8b17fe0211be12636f4c6310f5c7b73d22b3055980b52bae3b83
---

# mx-security-service-edge-integrations-meraki-secure-sd-wan-microsoft-sse-configu-d424482f

    To avoid the impact of a potential failure of Microsoft’s SSE solution DC failure, it is recommended to create one connection for each two different     Microsoft’s SSE solution DCs. Then manipulate BGP attributes to give preference to one Microsoft SSE DC over the other. The Weight, MED (Multiple Exit     Discriminator) or Path prepend fields under the routing section of the IPsec peer configuration can be used to prefer BGP routes from one peer over     another.  
- 
    For example, the default Weight for BGP peers is 0, setting the Weight of one Microsoft’s SSE DC 10, gives a higher priority to the routes advertised by the zone A over another zone with a default Weight of 0.
- 
    For symmetric return traffic, use AS path prepending ensure Microsoft SSE knows the primary tunnel to return traffic. For example, the path prepending field is empty by default, on the backup tunnel, add e.g. 64550 64550 to increase hop count seen by Microsoft SSE peer, this will ensure the tunnel with the least hops will be preferred. 
| Primary Microsoft SSE tunnel | Backup Microsoft SSE tunnel | 
- 
    WAN connectivity failure 
To avoid the impact of a potential WAN failure, it is recommended to have dual WAN links on your Secure Meraki SD-WAN device and configure an additional link under the connectivity tab. If WAN1 fails, connectivity to Microsoft’s SSE solution is established on WAN2. The tunnel over the WAN2 link will not establish until WAN1 fails on the MX Secure SD-WAN device.
- 
    Device failure 
To avoid the impact of device failure, it is recommended to deploy your Meraki Secure SD-WAN in a High Availability configuration. See MX Warm Spare guide to learn more about High availability.
This Topology and configuration highlights how to deploy a fully redundant connection to Microsoft’s SSE solution.
 
 
WAN 1 Active primary and secondary tunnels are independent tunnels with BGP attributes configured to prefer one peer over the other. The WAN 2 inactive tunnels become active when WAN1 connection fails.  
The Single device Dual WAN topology addresses DC failure and WAN connectivity failure, while the High Availability with Dual WAN virtual IP addresses DC failure, WAN connectivity failure and device failure.
Single device Dual WAN topology configuration
On MSFT Entra,
- 
    Configure two remote networks
- 
    Configure with two links per remote network pointing to each WAN IP 
On Meraki Dashboard,
- 
    Configure higher BGP Weight for preferred MSFT DC
- 
    Increase AS path for backup tunnel to MSFT DC to ensure symmetric routing
High Availability with Dual WAN virtual IP configuration
On MSFT Entra,
- 
    configure two remote networks
- 
    Configure with two links per remote peer pointing to each WAN IP
On Meraki Dashboard,
- Configure MX in Warm spare mode
- Configure virtual IP
- Configure higher BGP Weight for preferred MSFT DC
- Configure lower MED value for preferred MSFT DC to ensure symmetric routing
IPsec VPN Firewall
BGP uses TCP port 179, this means that TCP for 179 needs to be permitted through the VPN firewall to allow BGP communication between configured peers.
You can add firewall rules to control what traffic is allowed to pass through the VPN tunnel. These rules will apply to outbound VPN traffic to/from all MX-Z appliances in the Organization that participate in site-to-site VPN. These rules are configured in the same manner as the Layer 3 firewall rules described on this documentation. Note that VPN Firewall rules will not apply to inbound traffic or to traffic that is not passing through the VPN.
Serviceability
Event Logs
If you have any issues or would like to know more about the Cisco Secure Access peering details, navigate to Network-wide > Monitor > Event log
- 
    Select Event type Include - Non-Meraki VPN Negotiation
Packet Captures
The following options are available for a packet capture on MX/Z platforms:
- 
    Appliance: The appliance the capture will run on.
- 
    Interface: Select the interface to run the capture on; the interface names will vary depending on the appliance configuration. A few examples of interfaces you may see are:
- 
    Internet 1 or Internet 2 - Capture traffic on one active WAN uplink. Internet 2 will only appear if there is a second WAN link.
- 
    LAN - Captures traffic from all LAN ports
- 
    Cellular - Captures cellular traffic from the integrated cellular interface. This does not apply to USB modems.
- 
    Site-to-Site VPN - Captures AutoVPN traffic (MX/Z to MX/Z only). This does not apply to IPsec VPN peers.
- 
    Output: Select how the capture should be displayed; view output or download .pcap.
- 
    Verbosity: Select the level of the packet capture (only available when viewing the output directly to Dashboard).
- 
    Ignore: Optionally ignore capturing broadcast/multicast traffic.
- 
    Filter expressions: Apply a capture filter.
To capture packets, select the WAN interface and use the filter expressions for UDP 500 for Phase 1 or UDP 4500 for Phase 2.
API
The Meraki dashboard API is an interface for software to interact directly with the Meraki cloud platform and Meraki-managed devices. The API contains a set of tools known as endpoints for building software and applications that communicate with the Meraki dashboard for use cases such as provisioning, bulk configuration changes, monitoring, and role-based access controls. The dashboard API is a modern, RESTful API using HTTPS requests to a URL and JSON as a human-readable format. The dashboard API is an open-ended tool that can be used for many purposes.
For more information, read here.
24/7 Support
Cisco Meraki Support is available 24/7 to Enterprise customers for assistance with resolving network issues and providing answers to questions not covered by the documentation. For more information, read here.

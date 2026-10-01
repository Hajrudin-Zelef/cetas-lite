---
id: collect-261001-meraki/meraki/ms-meraki-campus-lan-5d88fe48-19
title: "ms-meraki-campus-lan-5d88fe48"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "license"]
source: docs/RAG/collect-261001-meraki/ms-meraki-campus-lan-5d88fe48.md
source_anchor: ""
source_lines: [790, 863]
sha256: 3bc94436cfc810498b03cd8c9dcd402155dbc68aaf73c23434e4ea4b6dfd0b17
---

# ms-meraki-campus-lan-5d88fe48

  - Navigate to Switch > Configure > Routing and DHCP
  - Delete any static routes other than the Default route for the desired switch
  - Delete any layer 3 interfaces other than the one which contains the next hop IP for the default route on the desired switch
  - Delete the last layer 3 interface to disable layer 3 routing
- For switch stacks performing L3 routing, it is possible that the management IP subnet can overlap with the subnet of any of it's own configured L3 interfaces
Please refer to the below table for scaling considerations when configuring SVI interfaces on Meraki Switches
Static Routes
General Guidance
- In order to route traffic elsewhere in the network, static routes must be configured for subnets that are not being routed by the switch or would not be using the default route already configured
- Static routes can be configured per switch or stack
- The Next hop IP is The IP address of the next layer 3 device along the path to this network. This address must exist in a subnet with a routed interface.
- You can edit an existing static route
- You can also delete an existing SVI but please note that A switch must retain at least one routed interface and the default route
- The default route cannot be manually deleted
- If OSPF is enabled, Dashboard provides the ability to pick and choose which static routes should be redistributed into the OSPF domain. You can also choose if you want to prefer the static route over OSPF or not
Routing Scaling Considerations for MS Platforms
| Model | Layer 3 Interfaces | Routes | Maximum Routable Clients | Features | 
| MS210 | 16 | 16 static routes | 8192 | Static Routing DHCP Relay | 
| MS225 | 16 | 16 static routes | 8192 |  | 
| MS2501 | 256 | 1024* (256 static routes) | 8192 | Static Routing OSPFv2 DHCP Relay DHCP Server Warm-spare (except MS390) Multicast Routing (PIM-SM) | 
| MS3502 | 256 | 16384* (256 static routes) | 24k |  | 
| MS350X | 256 | 8192 | 45k |  | 
| MS355 | 256 | 8192 (256 static routes) | 68k |  | 
| MS390 | 256 | 8192 (256 static routes) | 24k |  | 
| MS4102 | 256 | 16384* (256 static routes) | 24k |  | 
| MS425 | 256 | 8192 (256 static routes) | 212k |  | 
| MS450 | 256 | 8192 (256 static routes) | 68k |  | 
* The alert, "This switch is routing for too many hosts. Performance may be affected" will be displayed if the current number of routed clients exceeds the values listed in the table above
(1) The maximum number of learned OSPF routes is 900
(2) The maximum number of learned OSPF routes is 1500
L3 configuration changes on MS210, MS225, MS250, MS350, MS355, MS410, MS425, MS450 require the flushing and rebuilding of L3 hardware tables. As such, momentary service disruption may occur. We recommend making such changes only during scheduled downtime/maintenance window
MS390 Specific Guidance
- Please refer to the above guidance for MS390 platforms as well
Warm-Spare Switch Redundancy
It is recommended to use switch stacking to ensure reliability and high availability as opposed to warm-spare as it offers better redundancy and faster failover. If stacking is not available for any reason, warm-spare could be an option. Warm-spare with VRRP will also allow for the failure or removal of one of the distribution nodes without affecting endpoint connectivity to the default gateway.
General Guidance
- Both switches must be Layer 3 switches
- You will need to use two identical switches each with a valid license
- Have a direct connection between the two switches for the exchange of VRRP messages (Multicast address 224.0.0.18 every 300ms)
- Ensure that both the primary and spare have unique management IP addresses for communication with Dashboard (that does not conflict with the layer 3 interface IP addresses)
- Any changes made to L3 interfaces of MS Switches in Warm Spare may cause VRRP Transitions for a brief period of time. This might result in a temporary suspension in the routing functionality of the switch for a few seconds. We recommend making any changes to L3 interfaces during a change window to minimize the impact of potential downtime
When using Warm Spare on an MS switch it cannot be part of a switch stack or enable OSPF functionality as those features are mutually exclusive
All active L3 interfaces and routing functions on the "Spare" switch will be overwritten with the L3 configuration of the selected primary switch.
MS390 Specific Guidance
- MS390 series switches do not support warm spare/VRRP at this stage
DHCP Server
General Guidance
- MS switch platforms with layer 3 capabilities can be configured to run DHCP services (Please refer to datasheets for guidance on supported features)
- The MS switch can either disable DHCP (i.e. The MS will not process or forward DHCP messages on this subnet. This disables the DHCP service for this subnet), Run DHCP and respond to requests OR relay requests to another server
- If the Relay option is chosen, the MS will forward DHCP messages to a server in a different VLAN.
- If there are multiple DHCP relay server IPs configured for a single subnet, the MS will send the DHCP discover message to all servers. Whichever server responds back first is where the communication will continue
- You can proxy DNS requests to an upstream server in a different VLAN, to google DNS (8.8.8.8 and 8.8.4.4) or to Umbrella DNS servers
Please note that the Proxy to Umbrella features uses the OpenDNS server from Umbrella. If you require to use premium Umbrella services, please purchase the appropriate license(s) and instead choose the option "Proxy to Upstream DNS"
- DHCP options can also be specified
On MS, if an NTP server (option 42) is not configured, by default, the switch will use its SVI IP address as the NTP server option. This can cause problems for legacy devices that do not have hardcoded NTP servers since the MS does not respond to NTP requests
DHCP Snooping
Dashboard displays DHCP Servers seen by Meraki Switches on the LAN using DHCP snooping. Administrators can configure Email Alerts to be sent when a new DHCP server is detected on the network, block specific devices from being allowed to pass DHCP traffic through the switches, and see information about any currently active or allowed DHCP servers on the network.
Unlike DHCP, DHCP snooping does not require that the MS switches with layer 3 capabilities
- By default DHCP Servers can be explicitly blocked by entering the MAC address of the server in dashboard (This will prevent DHCP traffic sourced from that MAC from traversing the switches)
- Please note that DHVPv6 servers cannot be blocked using the MAC address
- You can also block or allow automatically detected DHCP servers from the DHCP Servers list
- Meraki switches detect a DHCP server if it detects a DHCP response from that server. (Dashboard will show further details such as MAC, VLANs and Subnets, Time last seen and a copy of the most recent DHCP packet)
- Meraki Switches configured as DHCP servers are automatically allowed
- It is recommended to check DHCP snooping on regular bases to track and action any Rogue DHCP servers
Blocking a DHCP server is done for its MAC address. Thus this server will be blocked for ALL VLANs and subnets.
If the policy is set to Deny DHCP Servers (i.e. Block DHCP Servers) then please when introducing a new DHCP server on your network (apart from the Meraki switches, e.g. Upstream your network) remember to unblock that DHCP server.
These features only apply to switches which are NOT bound to a configuration template
DHCPv6 is not logged on the DHCP servers and ARP page of the switch
MS390 Specific Guidance
- Please refer to the above for MS390 platforms as well
Dynamic ARP Inspection (DAI)
General Guidance
- Dynamic ARP Inspection (DAI) is a security feature in MS switches that protects networks against man-in-the-middle ARP spoofing attacks
- DAI inspects Address Resolution Protocol (ARP) packets on the LAN and uses the information in the DHCP snooping table on the switch to validate ARP packets.

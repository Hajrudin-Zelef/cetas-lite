---
id: collect-261001-meraki/meraki/ms-meraki-campus-lan-5d88fe48-8
title: "ms-meraki-campus-lan-5d88fe48"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "energy", "license", "voice"]
source: docs/RAG/collect-261001-meraki/ms-meraki-campus-lan-5d88fe48.md
source_anchor: ""
source_lines: [106, 174]
sha256: 73cc0963732409af098812235017329e65404727e22325c7278dcc52b05ca226
---

# ms-meraki-campus-lan-5d88fe48

The question of when a separate physical core is necessary depends on multiple factors. The ability of a distinct core to allow the campus to solve physical design challenges is important. However, it should be remembered that a key purpose of having a distinct campus core is to provide scalability and to minimize the risk from (and simplify) moves, adds, and changes in the campus. In general, a network that requires routine configuration changes to the core devices does not yet have the appropriate degree of design modularisation. As the network increases in size or complexity and changes begin to affect the core devices, it often points out design reasons for physically separating the core and distribution functions into different physical devices.
As a general rule of thumb, if your Distribution Layer is more than one stack (or two distribution units); it is recommended to introduce a dedicated Core Layer to interconnect the Distribution Layer and all the other network components
It's recommended to follow a collapsed core approach if you're distribution layer is:
- A single switch (No HSRP/VRRP/GLBP needed in this case)
- A stack of MS switches (No VRRP/warm-spare needed in this case)
- A pair of MS switches (Enable routing on access layer if possible, more guidance given in the below sections, otherwise enable VRRP/Warm-spare)
- Two stacks of Cisco catalyst switches (e.g. C9500)
Otherwise, it's recommended to follow a traditional three-tier approach to achieve a more scalable architecture
Planning, Design Guidelines and Best Practices
Planning for your Deployment
The following points summarize the design aspects for a typical Wired LAN that needs to be taken into consideration. Please refer to Meraki documentation for more information about each of the following items
- Hierarchical design traditional approach with 3 layers (access, aggregation, core) vs more common approach with 2 layers (access, collapsed core)
- Port density required on your access layer
- What port type/speeds are required on your MDF layer (GigE, mgigE, 10GigE, etc)
- Patching requirements between your IDF and MDF (Electrical, Multi-mode fibre, Single-mode fibre, etc)
- Number of stack members where applicable (this will influence your ether channels and thus number of ports on aggregation layer)
- Stackpower requirements where applicable (e.g MS390, C9300-M, etc)
- Port density required on your aggregation/collapsed-core layer
- Switching capacity is required on your aggregation/collapsed-core layer
- Consider using physical stacking on the access layer (typically useful if they are part of an IDF closet and cross-chassis port channeling is required)
- Layer 3 routing on Access layer (typically useful to reduce your broadcast domain and helps with fault isolation/downtime within your network)
- Calculate your PoE budget requirements (Which will influence your switch models, their power supplies and power supply mode)
- Check your Multicast requirements (IGMP snooping, Storm Control, Multicast routing, etc)
- If you require DHCP services on your access layer (distributed DHCP as opposed to centralized DHCP)
- End-to-end segmentation (Classification, Enforcement, etc.) using Security Group Tags requires MS390 Advanced License
- What uplink port speeds are required on the access layer (GigE, mgigE, 10GigE, etc)
- If any ports on your switch(s) that needs to be disabled
- On your access switches designate your upstream connecting ports (For non modular switches, e.g Ports 1-4)
- On your access switches, designate your Wireless LAN connecting ports (e.g Ports 5-10)
- On your access switches, designate your client-facing ports (e.g Ports 12-24)
- On your access switches, designate ports connecting to downstream switches (where applicable)
- On your access switches, designate ports that should provide PoE (e.g. Connecting downstream Access Points, etc)
- On your aggregation switches, designate ports connecting to upstream network (For non modular switches, e.g Ports 1-4)
- On your aggregation switches, designate ports connecting to downstream switches (e.g. Ports 10-16)
- On your switches, designate ports that should be isolated (e.g Restrict access between clients on the same VLAN)
- On your switches, designate ports that should be mirroring traffic (e.g Call recording software, WFM software, etc)
- On your switches, designate ports that should be in Trusted mode where applicable (For DAI inspection purposes)
- Choose your QoS mark and trust boundaries (i.e where to mark traffic, marking structure and values, trust or re-mark incoming traffic, etc)
- All client-facing ports should be configured as access ports
- Ports connecting downstream Access Points can either be configured as access (e.g NAT mode SSID, untagged bridge mode SSID, etc) or trunk (e.g tagged bridge mode SSID)
- Using Port Tags can be useful for administration and management purposes
- Using Port names can be useful for management purposes
- Do you require an access policy on your access ports (Meraki Authentication, external Radius, CoA, host mode, etc)
- What native VLAN is required on your access port(s) and will that be different per switch/stack?
- What management VLAN do you wish to use on your network? Will that be the same for all switches in the network or per switch/stack?
- Do you require a Voice VLAN on your access port(s)
- Size your STP domain based on your topology, no more than 9 hops in total.
- Designate your root switch based on your topology and designate STP priority values to your switches/stacks accordingly
- Do not disable STP unless absolutely required (e.g speed up DHCP process, entailed by network topology, etc)
- Use STP guards on switch ports to enhance network performance and stability
- Use STP BPDU guard on client-facing access ports
- Disable STP BPDU guard on ports connecting downstream switches
- Use STP Root guard on downstream ports connecting switches that are not supposed to become root
- Use STP Loop guard in a redundant topology on your blocking ports for further loop prevention (e.g Receiving data frames but not BPDUs)
- Always enable auto-negotiation unless the other end does not support that
- Enable UDLD when supported on the other end (Also please refer to Meraki firmware changelog for Meraki switches)
- It is recommended to enable UDLD in Alert-only mode on point to point links
- It is recommended to enable UDLD in Enforce mode on multi-point ports (e.g two or more UDLD-capable ports are connected through one or more switches that don't support UDLD)
- If using Layer 3 routing, plan your OSPF areas and routing flow from one area to the other (OSPF timers, interfaces, VLANs, etc)
- If enabling Adaptive Policy, choose the assignment method (Static vs Dynamic via Radius) and SGT per access port and whether trunk port peers are SGT capable or not
- Do you need to tag Radius and other traffic in a separate VLAN other than the management VLAN? (Refer to Alternate Management Interface)
- Check your MTU considerations taking into account all additional headers (e.g AnyConnet, Other VPNs, etc)
- Switch ACL requirements (e.g IPv4 ACLs, IPv6 ACLs, Group policy ACLs, etc). Click here for more information about Switch ACL operation
- Switch security requirements (e.g DHCP snooping behavior, Alerts, DAI, etc)
- For saving energy purposes, consider using Port schedules
Installation, Deployment & Maintenance
General Guidance
- Have your Campus LAN design finalized in terms of L2 and L3 nodes as well as the SVIs required where applicable (Please refer to the below sections for guidance on the design elements)
- Start with the network edge and have your firewalls and routers (e.g. Meraki MX SD-WAN & Security Appliance) connected to the public internet and able to access the Meraki cloud (Check firewall rules requirements for cloud connectivity)

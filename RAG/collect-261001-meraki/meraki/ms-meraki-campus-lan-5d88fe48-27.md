---
id: collect-261001-meraki/meraki/ms-meraki-campus-lan-5d88fe48-27
title: "ms-meraki-campus-lan-5d88fe48"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "license"]
source: docs/RAG/collect-261001-meraki/ms-meraki-campus-lan-5d88fe48.md
source_anchor: ""
source_lines: [1304, 1379]
sha256: cfa348419aaf851c6c1559a625e9c5136b861192fd744b02b5d3f8cf01995d20
---

# ms-meraki-campus-lan-5d88fe48

- Starting with MS 15.X, the MS390 will support Network Based Application Recognition (NBAR) Netflow v10 (IPFIX) for IPv4 and IPv6 traffic, as well as Encrypted Traffic Analytics (ETA) flow export for use with NetFlow analyzers like Cisco's Secure Network Analytics (formerly Stealthwatch Enterprise and Cloud).
- When the feature is enabled, every interface will collect flow records in both the input and output directions
- When configured it will be also enabled on every interface on every switch in the network that supports the feature and is configured correctly
- ETA requires MS390 Advanced License
- ETA requires a L3 SVI configured on exporting switch/stack
- ETA requires that the Collector that is reachable via L3 SVI
- The Netflow recorder configurations are very granular and support these fields.
Sample Topologies
The following section demonstrates some sample topologies encompassing both full Meraki architectures and hybrid architectures. The designs presented below take into consideration the design guidelines and best practices that have been presented in the previous sections of this article.
Topology 1 - Meraki Full Stack with Layer 3 Access
Logical architecture
Please refer to the following diagram for the logical architecture of Topology #1:
Assumptions
- It is assumed that Wireless roaming is confined within each zone area (No roaming between stacks)
- It is assumed that VLANs are local to each closet/zone and not spanning across multiple zones
- Corporate/BYOD SSID terminates in a single VLAN based on the AP zone
- Guest SSID only broadcasted in Zone 1
- IoT SSID only broadcasted in Zone 4
- Cisco ISE used for authentication and posturing
Considerations
- Putting all switches in the same Dashboard Network will help in providing a topology diagram for the entire Campus, however that also means that firmware upgrades will be performed for all switches within the same network which could be disruptive.
Please work with Meraki Support to assist in rolling firmware upgrades to different switches such that not all switches are scheduled for a firmware upgrade at the same time
- Access Stacks will offer DHCP services to SSID clients
- Either Edge MX OR Core Stack to offer DHCP services in Management VLANs, In either case make sure that the Static Routes on the MX pointing downstream are adjusted accordingly
- There is no use of VLAN 1 in this topology
- Transit VLANs are required to configure a default gateway per stack which needs to be separate from the Management VLAN range
- Only the SVI interfaces for Transit VLANs will be OSPF active interfaces. All other interfaces should be passive.
- If it is desired to use a quarantine VLAN returned from Cisco ISE (e.g. Guest VLAN for failed corp-auth) then it will be required to create dedicated SVIs per stack to host this traffic. Remember not to span VLANs across multiple stacks.
- OSPF cannot be turned on the Edge MX appliances since we are using multiple VLANs (VLAN 500 and 1925)
- During the in-life production of this topology, any change on a layer 3 interface will cause a brief interruption to packet forwarding. It is therefore recommended to do that during a maintenance window
- When adding new stacks, you must prepare the network for that by creating the required SVIs and Transit VLANs
Topology 2 - Meraki Full Stack with Layer 2 Access
Logical architecture
Please refer to the following diagram for the logical architecture of Topology #2:
Assumptions
- It is assumed that Wireless roaming is required everywhere in the Campus
- It is assumed that VLANs are spanning across multiple zones
- Corporate SSID (Broadcasted in all zones) users are assigned VLAN 10 on all APs. CoA VLAN is VLAN 30 (Via Cisco ISE)
- BYOD SSID (Broadcasted in all zones) users are assigned a VLAN 20 on all APs. CoA VLAN is VLAN 30 (Via Cisco ISE)
- IoT and Guest SSID broadcasted everywhere in Campus
- Access Switches will be running in Layer 2 mode (No SVIs or DHCP)
- Access Switch uplinks are in trunk mode with native VLAN = VLAN 100 (Management VLAN)
- STP root is at Distribution/Collapsed-core
- Distribution/Collapsed-core uplinks are in Trunk mode with Native VLAN = VLAN 1 (Management VLAN)
- All VLAN SVIs are hosted on the edge MX and not in Campus LAN
- Network devices will be assigned fixed IPs from the management VLAN DHCP pool. Default Gateway is 10.0.1.1
Considerations
- Putting all switches in the same Dashboard Network will help in providing a topology diagram for the entire Campus, However that also means that firmware upgrades will be performed for all switches within the same network which could be disruptive.
Please work with Meraki Support to assist in rolling firmware upgrades to different switches such that not all switches are scheduled for a firmware upgrade at the same time
- To enable Wireless roaming in Campus, SSIDs will be configured in Bridge mode which results in seamless Layer 2 roaming
- Layer 2 roaming requires that all APs are part of the same broadcast domain (i.e. Upstream VLAN consistency)
- All VLANs will be hosted on the MX appliance where DHCP will be running for both Network devices and end user clients. VRRP will be used to point devices to the "primary" default gateway across the 2 appliances
- Upstream VLAN consistency across all stacks results in a broadcast domain that spans across the whole campus. Consequently, STP must be tightly configured to protect the network from loops
- Based on the number of users, the VLAN size might need to be adjusted leading to a larger broadcast domain
Topology 3 - Hybrid Campus with Layer 3 MS390 Access
Logical architecture
Please refer to the following diagram for the logical architecture of Topology #3:
Assumptions
- It is assumed that Wireless roaming is confined within each zone area (No roaming between stacks)
- It is assumed that VLANs are local to each closet/zone and not spanning across multiple zones
- Corporate/BYOD SSID terminates in a single VLAN based on the AP zone
- Guest SSID only broadcasted in Zone 1
- IoT SSID only broadcasted in Zone 2
- Cisco ISE used for authentication and posturing
Considerations
- If you're using Virtual IP on the MX WAN uplinks, then the MXs must share the same broadcast domain on the WAN side.
- Access Stacks will offer DHCP services to SSID clients
- Core Stack or MX WAN Edge will offer DHCP services in Management VLANs
- Stacks in Native VLANs apart from VLAN 1 will need to be either pre-configured OR provisioned in VLAN 1 then changed to their respective native VLAN per the above diagram for ease of initial setup.
- Only the SVI interfaces for Management VLANs will be OSPF active interfaces. All other interfaces should be passive.
- If it's desired to use a quarantine VLAN returned from Cisco ISE (e.g. Guest VLAN for failed corp-auth) then it will be required to create dedicated SVIs per stack to host this traffic. Remember not to span VLANs across multiple stacks.
- During the in-life production of this topology, any change on a layer 3 interface will cause a brief interruption to packet forwarding. It is therefore recommended to do that during a maintenance window
- When adding new stacks, you must prepare the network for that by creating the required SVIs for Management VLAN(s)
- Consider configuring STP in this architecture as a failsafe option. However, it is not expected that you have any blocking links based on the design proposed
Wireless Roaming (Layer 3)
To enable wireless roaming for this architecture, a dedicated MX in concentrator mode is required. Please refer to the following diagram for more details:

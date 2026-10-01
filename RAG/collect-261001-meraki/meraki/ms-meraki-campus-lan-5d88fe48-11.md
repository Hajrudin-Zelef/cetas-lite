---
id: collect-261001-meraki/meraki/ms-meraki-campus-lan-5d88fe48-11
title: "ms-meraki-campus-lan-5d88fe48"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/ms-meraki-campus-lan-5d88fe48.md
source_anchor: ""
source_lines: [316, 387]
sha256: 444df7bb86d9f75ddf6a070744c25e2e1975c8125c399eec1c4cef3d6daddfcf
---

# ms-meraki-campus-lan-5d88fe48

All MS platforms (excluding MS390) use a separate routing table for management traffic. Configuring a Management IP within the range of a configured SVI interface can lead to undesired behavior. The Management VLAN must be separate from any configured SVI interface.
- Unrequired VLANs should be manually pruned from trunked interfaces to avoid broadcast propagation.
- If you require that your Radius, Syslog or SNMP traffic to be encapsulated in a separate VLAN (that is not necessarily exposed to the internet) then consider using the Alternate Management Interface on MS. Please refer to the table below for this feature compatibility:
| MS Switch Family | MS Switch Model | MS Firmware Support (first supported on) | 
| MS2xx | MS210 | MS14.5 | 
|  | MS225 | MS14.5 | 
|  | MS250 | MS14.5 | 
| MS3xx | MS350 | MS14.5 | 
|  | MS355 | MS14.5 | 
|  | MS390 | MS15 | 
| MS4xx | MS410 | MS14.5 | 
|  | MS425 | MS14.5 | 
|  | MS450 | MS14.5 | 
The Alternate Management Interface (AMI) functionality is enabled at a per-network level and, therefore, all switches within the Dashboard Network will use the same VLAN for the AMI. The AMI IP address can be configured per switch statically as shown below:
Please note that the subnet of the AMI (the subnet mask for the AMI IP address) is derived from Layer-3 interface for the AMI VLAN, if one has been configured on the switch. In the absence of a Layer-3 interface for the AMI VLAN, each switch will consider its AMI to be /32 network address
Layer 3 routing must be enabled on a switch for its AMI to be activated
MS390 Specific Guidance
- The default active VLANs on any MS390 port is 1-1000. This can be changed via local status page or in dashboard (See note below)
- Please ensure that the MS390 switch/stack has a maximum of 1000 VLANs
- The total number of VLANs supported on ANY MS390 switch port is 1000
For example, If you have an existing stack with each port set to Native VLAN 1, 1-1000 and the new member ports are set to native VLAN 1; allowed VLANs: 1,2001-2500 then your total number of VLAN in the stack will be 1000(1-1000)+500(2001-2500) = 1500. Dashboard will not allow the new member to be added to the stack and will show an error.
To utilize any VLANs outside of 1-1000 on an MS390, the switch or switch stack must have ALL of its trunk interfaces set to an allowed vlan list that contains a total that is less than or equal to 1000 VLANs, including any of the module interfaces that are not in use. Here's a quick way to do that.
MS390 Stacking Specific Guidance
- Please refer to the MS390 Stacking guidance provided below
DHCP
General Guidance
- DHCP is recommended for faster deployments and zero-touch
- It is recommended to fix the DHCP assignments on the DHCP server as this will ensure that other network applications (e.g. Radius) will always use the same source IP address range (i.e. the Management/AMI VLAN)
- Static IP addressing can also be used however to minimize initial provisioning it's recommended to use DHCP for initial setup, then change IP addressing from dashboard. Meraki MS switches will attempt to do DHCP discovery on all supported VLANs.
Please refer to the stacking section for further guidance on IP addressing when using switch stacks
MS390 Specific Guidance
- When installing an MS390, it is important to ensure that any DHCP services or IP address assignments used for management fall within the active VLAN range (1-1000 by default, unless changed via the local status page or dashboard)
- If you require using Static IP addressing (OR an IP Address outside of the default active VLANs 1-1000) please connect each MS390 switch with an uplink to the Meraki dashboard and upgrade firmware to latest stable prior to changing any configuration. Once the switch upgrades and reboots, you can now change the management IP as required (please ensure the upstream switch/device allows this VLAN in its port configuration).
MS390 Stacks Specific Guidance
- It is recommended to set the same IP address on all switches in dashboard once DHCP assigns IP addressing and the stack is online (e.g. dashboard shows that the management IP of the stack is 10.0.5.20, then please statically set this IP on all switch members of the stack)
- Thus, it is recommended to use Static IP address as opposed to DHCP. Please connect each MS390 switch with an uplink (do not connect any stacking cables at this stage) to dashboard and upgrade firmware to latest stable prior to changing any configuration. Once the switch upgrades and reboots, you can now change the management IP as required (please ensure the upstream switch/device allows this VLAN in its port configuration). Don't forget to assign the same IP address to all members of the stack. (Start with the primary switch, this should automatically assign the same IP to all members within the same stack)
MS390 Stack IP Address Provisioning Sequence for Best Results:
- Claim your MS390s into a dashboard network (do not create a stack)
- Set the firmware to 11.31+
- Connect an uplink to each switch (members un-stacked)
- Ensure that the stacking cables are not connected to any member
- Power on switches (members un-stacked)
- Have DHCP available on native VLAN 1
- Wait for firmware to be loaded and configuration to be synced
- Power off switches
- Disconnect all uplinks from all switches
- Connect stacking cables to all members to form a ring topology
- Connect one uplink to one member (only one link for the stack)
- Power on switches and wait for them to come online on dashboard
- Create a stack on dashboard by adding all members
- Wait for the stack ports to show online on all members in dashboard
- Observe the IP address used on the stack members (should be the same for all members)
- Click on the IP address of each switch and change settings from DHCP to Static. Configure the IP address that is used for the stack for each switch member
- Configure Link aggregation as needed and add more uplinks accordingly
- Make sure to abide to the maximum VLAN count as described in the below section when you provision your MS390 stack/switches
Please note that the Primary switch owns the Management IP and will resolve ARP requests to its own MAC address
Supported VLANs
General Guidance
- All MS platforms (except MS390): VLANs 1-4096 supported
- It is recommended to take some initiative in designing the campus to decrease the size of broadcast domains by limiting where VLANs traverse. This requires that your VLANs to be trunked to only certain floors of the building or even to only certain buildings depending on the physical environment. This reduces the flooding expanse of broadcast packets so that traffic doesn’t reach every corner of the network every time there’s a broadcast, reducing the potential impact of broadcast storms
Meraki MS platforms do not support the VTP Protocol
MS390 Specific Guidance
- MS390s support the following VLAN ranges: 1-1001 and 1006-4092 with a maximum VLAN count* of 1000
- MS390 has the following Default Active VLANs: VLAN ID 1-1000 (i.e. configured by default) However, the active VLANs can be changed via the local status page or dashboard (after the switch has come online)
The following VLAN ranges are reserved on MS390 switches: 1002-1005, 4093-4094
* MS390 switches support up to 1000 VLANs in total. It is recommended to configure the switch ports with the specific VLANs (or ranges) to stay within the 1000 VLAN count (e.g. 1-20, 100-300, 350, 900-1000, 1050-1250)
Please ensure that all trunk ports on MS390 switches are configured such that the maximum VLAN count is 1000. (i.e. Do not exceed the maximum VLAN count of 1000 on any switch port)
MS390 Stacks Specific Guidance
- Same guidance for MS390s
Spanning Tree Protocol & UDLD
General Guidance
- All Meraki MS switches (with the exception of MS390) support RSTP for loop prevention

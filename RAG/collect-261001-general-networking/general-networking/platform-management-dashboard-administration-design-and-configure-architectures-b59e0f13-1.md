---
id: collect-261001-general-networking/general-networking/platform-management-dashboard-administration-design-and-configure-architectures-b59e0f13-1
title: "platform-management-dashboard-administration-design-and-configure-architectures--b59e0f13"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/platform-management-dashboard-administration-design-and-configure-architectures--b59e0f13.md
source_anchor: ""
source_lines: [1, 128]
sha256: a08fb0f4dd1461798f7d26c3f1b5cfcdc55db427fb988d2ad541d45a19ec3ba6
---

# platform-management-dashboard-administration-design-and-configure-architectures--b59e0f13

General MS Best Practices
Layer 2 Features
- 
    STP 
  - 
        RSTP is enabled by default and should always be enabled. Disable only after careful consideration.
  - 
        PVST interoperability (Catalyst/Nexus) 
    - 
            VLAN 1 should be allowed on a trunk between Catalyst and MS. This is crucial for RSTP
    - 
            Make Catalyst the root switch
  - 
            
  - 
        Set root switch priority to “0 - likely root” 
    - 
            Higher end models such as the MS410, MS425 deployed at core or aggregation are suitable candidates for the role
    - 
            Ideally, the switch designated as the root should be one which sees minimal changes (config changes, link up/downs etc.) during daily operation
  - 
            
- 
        
- Keep the STP diameter under 7 hops, such that packets should not ever have to travel across more than 7 switches to travel from one point of the network to the other
- BPDU Guard should be enabled on all end-user/server access ports to avoid rogue switch introduction in network
- Loop Guard should be enabled on trunk ports that are connecting switches
- Root Guard should be enabled on ports connecting to switches outside of administrative control
- 
    MTU 
  - 
        Recommended to keep at default of 9578 unless intermediate devices don’t support jumbo frames. This is useful to optimize server-to-server and application performance. Avoid fragmentation when possible.
- 
        
- 
    Switchports 
  - 
        Trunk 
    - 
            Prune unnecessary VLANs off trunk ports using allowed VLAN list in order to reduce scope of flooding
    - 
            Ensure that the native VLAN and allowed VLAN lists on both ends of trunks are identical. Mismatched native VLANs on either end can result in bridged traffic
 Tagging 
    - 
            For ease of management, assign tags to switch ports. For example, switch<->switch links can be assigned “trunk”, switch<->AP can be “wireless” etc
  - 
            
  - 
        Aggregation 
    - 
            Only LACP is supported for link aggregation. Ensure the other end supports LACP
    - 
            It is recommended to configure aggregation on the dashboard before physically connecting to a partner device
  - 
            
- 
        
- UDLD (Unidirectional Link Detection)
    
  - This should be enabled on fiber trunks - in “Alert Only” mode
- Link Negotiation
    
  - This should be set to auto-negotiate for ports connecting Meraki devices
  - Use “forced” mode only if a device connected to the port does not support auto-negotiation
- Switchport count in a network
    
  - It is recommended to keep the total switch port count in a network to fewer than 8000 ports for reliable loading of the switch port page.
Layer 3 Features
- IP addressing and subnetting schema
    
  - Dedicate /24 or /23 subnets for end-user access
  - Avoid overlapping subnets as this may lead to inconsistent routing and forwarding
- 
    L3 Interfaces 
  - 
        Assign a dedicated management VLAN
  - 
        Avoid configuring a L3 interface for the management vlan. Use L3 interfaces only for data VLANs. This helps in separating management traffic from user data
- 
        
In case of switch stacks, ensure that the management IP subnet does not overlap with the subnet of any configured L3 interface. Overlapping subnets on the management IP and L3 interfaces can result in packet loss when pinging or polling (via SNMP) the management IP of stack members. NOTE: This limitation does not apply to the MS390 series switches.
L3 configuration changes on MS210, MS225, MS250, MS350, MS355, MS410, MS425, MS450 require the flushing and rebuilding of L3 hardware tables. As such, momentary service disruption may occur. We recommend making such changes only during scheduled downtime/maintenance window.
- 
    Access Control Lists (ACLs) 
  - 
        Summarize IP addresses as much as possible (before-after examples below).
- 
        
- Maximum ACL limit is 128 access control entries (ACEs) per network
Take control over your network traffic. Review user and application traffic profiles and other permissible network traffic to determine the protocols and applications that should be granted access to the network. Ensure traffic to the Meraki dashboard is permitted (Help > Firewall Info)
- OSPF
    
  - Found under Switching > Configure > OSPF Routing
  - All configured interfaces should use broadcast mode for hello messages
  - We recommend leaving the “hello” and “dead” timers to a default of 10s and 40s respectively. If more aggressive timers are required, ensure adequate testing is performed.
  - Ensure all areas are directly attached to the backbone Area 0. Virtual links are not supported
  - Configure a Router ID for ease of management
- With multiple VLANs on a trunk, OSPF attempts to form neighbor relationships over each VLAN, which may be unnecessary. To exchange routing information, OSPF doesn’t need to form neighbor relationships over every VLAN. Instead, a dedicated transit VLAN can be defined and allowed on trunks, typically between the core and aggregation layers with OSPF enabled and “Passive” set to “no.” For all other subnets that need to be advertised, enable OSPF and set “Passive” to “Yes.” This will reduce unnecessary load on the CPU. If you follow this design, ensure that the management VLAN is also allowed on the trunks.
- Configure MD5 authentication for added security
- 
    DHCP 
  - 
        Specify allowed DHCP servers to protect against rogue servers
  - 
        In a warm spare configuration, the load balancing mechanism for DHCP, in some case, may be inefficient and cause an issue where devices may try to get an address from a member with no leases remaining. This is addressed in a stacked configuration, where this issue will not occur.
- 
        
Topology
- We highly recommend having the total switch count in any dashboard network to be less than or equal to 400 switches. If switch count exceeds 400 switches, it is likely to slow down the loading of the network topology/ switch ports page or result in display of inconsistent output.
Multicast
- 
    The most important consideration before deploying a multicast configuration is to determine which VLAN the multicast source and receivers should be placed in. If there are no constraints, it is recommended to put the source and receiver in the same VLAN and leverage IGMP snooping for simplified configuration and operational management.
- 
    Multicast Routing 
  - 
        Meraki switches provide support for 30 multicast routing enabled L3 interfaces on a per switch level
  - 
        PIM SM requires the placement of a rendezvous point (RP) in the network to build the source and shared trees. It is recommended to place the RP as close to the multicast source as possible. Where feasible, connect the multicast source directly to the RP switch to avoid PIM’s source registration traffic which can be CPU intensive. Typically, core/aggregation switches are a good choice for RP placement
  - 
        Ensure every multicast group in the network has an RP address configured on Dashboard
  - 
        Ensure that the source IP address of the multicast sender is assigned an IP in the correct subnet. For example, if the sender is in VLAN 100 (192.168.100.0/24), the sender's IP address can be 192.168.100.10 but should not be 192.168.200.10.
  - 
        Make sure that all Multicast Routing enabled switches can ping the RP address from all L3 interfaces that have Multicast Routing enabled
  - 
        Configure an ACL to block non-critical groups such as 239.255.255.250/32 (SSDP). As of MS 12.12, Multicast Routing is no longer performed for the SSDP group of 239.255.255.250.
- 
        

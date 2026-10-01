---
id: collect-261001-meraki/meraki/ms-meraki-campus-lan-5d88fe48-20
title: "ms-meraki-campus-lan-5d88fe48"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "voice"]
source: docs/RAG/collect-261001-meraki/ms-meraki-campus-lan-5d88fe48.md
source_anchor: ""
source_lines: [864, 933]
sha256: 1213b2a7f44477c2e8872b531359e38c7eab463df5a6d373821227a1f8cee225
---

# ms-meraki-campus-lan-5d88fe48

- DAI performs validation by intercepting each ARP packet and comparing its MAC and IP address information against the MAC-IP bindings contained in the DHCP snooping table (i.e Any ARP packets that are inconsistent with the information contained in the DHCP snooping table are dropped)
- DAI associates a trust state with every port on the switch. Ports marked as trusted are excluded from DAI validation checks and all ARP traffic is permitted. Ports marked as untrusted are subject to DAI validation checks and the switch examines ARP requests and responses received on those ports.
- It is recommended to configure only ports facing end-hosts as untrusted (Trusted: disabled)
- It is recommended to configure ports connecting network devices (e.g switches, routers) as trusted to avoid connectivity issues
- Since DAI relies on the DHCP snooping tables, it is recommended to enable DAI only on subnets with DHCP enabled otherwise the ARP packet will be dropped
- DAI is disabled by default and needs to be enabled before configuring port settings
- DAI blocked events are logged, it is recommended to check those logs on regular bases
DAI is supported on the following platforms with MS10+:
MS210, MS225, MS250, MS350, MS355, MS390, MS410, MS425, MS450
MS390 Specific Guidance
- MS390 series switches do support DAI with firmware
Multicast
General Guidance
- The most important consideration before deploying a multicast configuration is to determine which VLAN the multicast source and receivers should be placed in.
- If there are no constraints, it is recommended to put the source and receiver in the same VLAN and leverage IGMP snooping for simplified configuration and operational management
- PIM SM requires the placement of a rendezvous point (RP) in the network to build the source and shared trees. It is recommended to place the RP as close to the multicast source as possible. Where feasible, connect the multicast source directly to the RP switch to avoid PIM’s source registration traffic which can be CPU intensive (Typically, core/aggregation switches are a good choice for RP placement)
- 
    Ensure every multicast group in the network has an RP address configured on Dashboard
- 
    Ensure that the source IP address of the multicast sender is assigned an IP in the correct subnet. For example, if the sender is in VLAN 100 (192.168.100.0/24), the sender's IP address can be 192.168.100.10 but should not be 192.168.200.10.
- 
    Make sure that all Multicast Routing enabled switches can ping the RP address from all L3 interfaces that have Multicast Routing enabled
- 
    Configure an ACL to block non-critical groups such as 239.255.255.250/32 (SSDP) (Please note that as of MS 12.12, Multicast Routing is no longer performed for the SSDP group of 239.255.255.250)
- 
    Disable IGMP Snooping if there are no layer 2 multicast requirements. IGMP Snooping is a CPU dependent feature, therefore it is recommended to utilize this feature only when required (For example, IPTV)
- 
    It is recommended to use 239.0.0.0/8 multicast address space for internal applications
- 
    Always configure an IGMP Querier if IGMP snooping is required and there are no Multicast routing enabled switches/routers in the network. A querier or PIM enabled switch/router is required for every VLAN that carries multicast traffic
- 
    Storm control is recommended to be set to 1%
- 
    Storm control expected behavior is that it will drop excessive packets if the limit has been exceeded
Storm control is not supported on the following MS platforms: MS120, MS220 and MS320
Multicast Scaling Considerations
Meraki switches provide support for 30 multicast routing enabled L3 interfaces on a per switch level
MS390 Specific Guidance
- All above guidance, plus:
- Without IGMP snooping, MS390 will flood all traffic
- With IGMP snooping, MS390 will not flood traffic
- Storm control expected behavior is that it will drop all packets if the limit has been exceeded (not just the excess traffic) until the monitored traffic drops below the defined limit (1 sec interval)
Link Aggregation
General Guidance
- It is very important to match Link Aggregation (aka Ether-Channel) settings between CatOS, Cisco IOS and Meraki MS Switches
- Please note that the defaults are different between the different platforms
- The supported protocols might also be different so please consult the configuration guides of each of your switches and ensure the configuration is consistent across your Ether-Channels.
- MS platforms support both 802.3ad and 802.1ax LACP
- Running any other state or protocol on the remote side (this includes pagp and just set to ‘ON’) will cause issues
- Up to 8 members in a single Ether-Channel
- Aggregates of ports spread over multiple members of a stack is supported
- LACP is set to active mode and LACPDUs will be sent out the ports trying to initiate a LACP negotiation
- In a Hybrid Campus LAN (includes Cisco IOS and/or CatOS devices) make sure that PAgP settings are the same on both sides. The defaults are different. CatOS devices should have PAgP set to off when connecting to a Cisco IOS software device if EtherChannels are not configured.
- It is recommended to configure aggregation on the dashboard before physically connecting to a partner device
- It is recommended to configure the downlink device first, wait for the config to state up to date, before configuring the aggregation uplink device (If the process is performed in the uplink side first, there may be an outage depending on the models of switches used)
- If you are setting up Ether-channel between Meraki switches, it is recommended to set it on auto-negotiate
- If you are setting up Ether-channel between Meraki and other switches (e.g. Cisco Catalyst), it is recommended to set it on forced IF you are unsure that the other switch(es) supports auto-negotiate mode.
- If you are setting up Ether-channel between Meraki MS and Cisco Catalyst, it may be advantageous on the Catalyst switch to disable the feature "spanning-tree etherchannel guard misconfig" if there are issues with getting the LACP aggregate established
In relation to SecureConnect, If an MR access-point that does not support LACP is plugged into a switchport which is part of an LACP aggregate group, the switchport will be disabled by LACP. MR access-points that do support LACP, when plugged into a switchport configured as a part of an LACP aggregate group will continue to function as they would if SecureConnect was disabled.
Link Aggregation is supported on ports sharing similar characteristics such as link speed and media-type (SFP/Copper).
MS390 Specific Guidance
- It is recommended to refresh your browser (Dashboard UI for switchports) before enabling link aggregation
- Please ensure that you enable link aggregation on dashboard before connecting multiple links to the other switch
- By default, prior to configuring LACP, the MS series runs an LACP Passive instance per port. This is to prevent loops when a bonded link is connected to a switch running the default configuration. Once LACP is configured, the MS will run an Active LACP instance with a 30-second update interval and will always send LACP frames along the configured links.
Oversubscription and QoS
General Guidance
- It is recommended for oversubscription on access-to-distribution uplinks to be below 20:1, and distribution-to-core uplinks to be 4:1 (This will be mostly dependent on the application requirements so should be considered as a rule of thumb)
- When congestion does occur, QoS is required to protect important traffic such as mission-critical data applications, voice, and video
- Meraki MS series switches support adding (i.e. Marking) and honoring of DSCP tags for incoming traffic (DSCP tags can be added, modified or trusted)
- QoS rules are processed top to bottom

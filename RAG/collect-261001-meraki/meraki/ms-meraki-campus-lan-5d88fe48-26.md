---
id: collect-261001-meraki/meraki/ms-meraki-campus-lan-5d88fe48-26
title: "ms-meraki-campus-lan-5d88fe48"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["full-duplex"]
source: docs/RAG/collect-261001-meraki/ms-meraki-campus-lan-5d88fe48.md
source_anchor: ""
source_lines: [1233, 1303]
sha256: 28d6931473156238cad37f7b7f425f71dacb1fc6bbc75b7cdd19866175d02d1b
---

# ms-meraki-campus-lan-5d88fe48

- Ensure that you configure a Bridge STP Priority that is higher than your access layer (e.g. 61441 or higher)
- Do not extend your STP domain by more than 5 hops
Please refer to the following diagrams for some topology guidelines:
Adaptive Policy
General Guidance
- MS390 is the only platform that supports Adaptive Policy
- Non MS390 platforms will not tag traffic and will not enforce Adaptive Policy
- Non MS390 platforms will drop traffic with a SGT tag
MS390 Specific Guidance
- MS390 switches support Adaptive Policy
- It is recommended to keep the default infrastructure group as it is (SGT value 2) which is used to tag all Meraki Cloud traffic.
- All other network devices (e.g. Access Points, Switches, etc) can be either part of this group or create a separate group for management traffic if required.
- Also please note that ACLs are processed from the top down, with the first rule taking precedence over any following rules.
- If the device connected to a MS390 trunk port does not support SGTs, please ensure that Peer SGT Capable is disabled in dashboard otherwise the device on the other end won't be able to communicate.
- If you are using a Radius server (e.g. Cisco ISE) to return SGT values then please ensure that it returns the value in HEX value (in this case do not use a static mapping on the port as Radius attribute: cisco-av-pair:cts:security-group-tag cannot override the static value)
Please note that you cannot have static group assignment AND an 802.1x access policy configured on a switchport. If 802.1x is used on the interface, you must configure the interface group tag to "Unspecified" for the configuration to work properly. In this case all Access-Accept messages for clients will require an SGT using the cisco-av-pair:cts:security-group-tag tag.
Adaptive Policy Scaling Considerations:
Maximum number of Adaptive Policy Groups: 60
Maximum number of policies configured: 3600
Maximum Custom ACLs per (Group > Group) policy: 10
Maximum number of ACE entries per Custom ACL: 16
Maximum number of ACES entries per source group to destination group policy: 160
Maximum IP to SGT mappings: 8000
Caution: If you DELETE a tag, it will be removed from mapping on every network device and every configuration including static port mappings and SSID configurations. DO NOT delete a tag unless that is the desired outcome. Also, Removing Adaptive Policy from a network will affect all Adaptive Policy capable devices in that network.
Operations, Administration and Maintenance
General Guidance
- Cable testing feature on dashboard can be used on all ports, however it can disrupt live traffic.
- Half-duplex mode is supported on all ports
- Syslog, SNMP are supported on all MS platforms (except MS390)
- L3 configuration changes on MS210, MS225, MS250, MS350, MS355, MS410, MS425, MS450 require the flushing and rebuilding of L3 hardware tables. As such, momentary service disruption may occur. it is recommended to make such changes only during scheduled downtime/maintenance window
- For MS220-8 and MS220-8P, it is recommended to keep the MAC entries below 8k. And for all other MS1xx and MS2xx platforms, it is recommended to keep the MAC entries below 16k. For any other MS platform, it is recommended to keep the MAC entries below 32k.
- Putting all switches in the same Dashboard Network will help in providing a topology diagram for the entire Campus, however that also means that firmware upgrades will be performed for all switches within the same network which could be disruptive.
Please work with Meraki Support to assist in rolling firmware upgrades to different switches such that not all switches are scheduled for a firmware upgrade at the same time
- It is recommended for an optimal experience with Dashboard Topology to keep the number of switches below 400 and number of devices below 1000 in a single dashboard network
- It is recommended to implement tag-based port permissions to restrict access to specific ports as required (Please refer to dashboard administration and management for more info)
- It is recommended for an optimal experience with Dashboard to keep the number of ACL entries per network below 128
- The Virtual stacking feature allows you to do port search with ease. Refer to this guide for guidance on how to search for ports and search filters.
- Port isolation allows a network administrator to prevent traffic from being sent between specific ports. This can be configured in addition to an existing VLAN configuration, so even client traffic within the same VLAN will be restricted
For MS210, MS225, and MS250 series switches, port isolation is only supported on the first 24 ports
- It may be necessary to configure a mirrored port or range of ports. This is often useful for network devices that require monitoring of network traffic, such as a VoIP recording solution or an IDS/IPS
MS switches support one-to-one or many-to-one mirror sessions. Cross-stack port mirroring is available on Meraki stackable switches. Only one active destination port can be configured per switch/stack
MS390 Specific Guidance
- Rebooting or power recycling a MS390 switch will reboot the entire stack. It is recommended to do that within a maintenance window
- Cable testing feature on dashboard can be used on all ports except the module ports (i.e uplink ports)
- The maximum jumbo frames supported on MS390 is 9198 bytes. Routed MTU is 1500 bytes
- Please note that half-duplex mode is not supported on mGig ports (only supported on GigE ports). As such, it is recommended to manually set MS390 mGig ports with full-duplex and the other switch port with full-duplex mode as well. (This will also speed up the process for links to establish)
- SNMP and Syslog are not yet supported on MS390s
- Netflow and Encrypted Traffic Analytics is supported on MS390s with MMS15.x and Advanced Licensing
- MS390s take considerably more time to boot as compared to other Meraki MS switching platforms.
- Please be patient until the switches complete the bootup process. (Do not power down or reset the device during a firmware upgrade. A device which has its power LED blinking white is indicating that it is going through a firmware upgrade).
- Also please note that stacks will take longer to boot.
- Refer to the below for indicative bootup times:
| Firmware | Stack Size | Boot Time | 
|---|---|---|
| MS 12.28 and older | 8 | 50m | 
| MS 12.28 and older | 1 | 11m30s | 
| MS 14.12+ | 8 | 22m | 
| MS 14.12+ | 1 | 7m18s | 
Management Port
- All Meraki MS platforms are equipped with a dedicated management port with the exception of the following models: MS120-8, MS120-8FP, MS120-8LP, MS220-8, MS220-8P
Local Status Page
- For ALL MS platforms (except MS390) without a dedicated Management Port: Please connect a wired client to one of the switchports and assign a static IP address 1.1.1.99 with Subnet mask 255.255.255.0 and browse to 1.1.1.100
- For MS390 platforms: Please connect a wired client to one of the switchports and assign a static IP address 10.128.128.132 with Subnet mask 255.0.0.0 and DNS 10.128.128.130 then browse to 10.128.128.130
- For ALL MS platforms (except MS390) with a dedicated Management Port: Please connect a wired client to the management port. No static IP address is needed. Simply browse to 1.1.1.100 to access the local status page
SM Sentry
- SM Sentry is a feature that is used to enroll end-user clients via Meraki devices (e.g. MR and MS)
- SM Sentry is supported on all MS platforms with the exception of: MS120, MS125, MS220, MS320
Netflow and Encrypted Traffic Analytics
General Guidance
- Encrypted Traffic Analytics is not supported on any MS platform at this stage except the MS390
MS390 Specific Guidance

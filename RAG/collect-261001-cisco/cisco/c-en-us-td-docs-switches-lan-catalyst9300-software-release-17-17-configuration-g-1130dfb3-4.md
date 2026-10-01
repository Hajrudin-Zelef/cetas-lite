---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-17-configuration-g-1130dfb3-4
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-17-configuration-g-1130dfb3"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-17-configuration-g-1130dfb3.md
source_anchor: ""
source_lines: [96, 138]
sha256: 6019c29d6b9c2751c24993545c06678ee598a08bdee6b227a5f303b1a5f17605
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-17-configuration-g-1130dfb3

| Step 12 | ipv6 ttl ttl-value Example: Device(config-mon-erspan-src-dst)# ipv6 ttl 4 | (Optional) Configures the IPv6 TTL value of packets in the ERSPAN traffic. | 
| Step 13 | ipv6 flow-label flow-label-value Example: Device(config-mon-erspan-src-dst)# ipv6 flow-label 2 | (Optional) Configures the IPv6 flow label in the ERSPAN traffic. flow-label-value: The range is from 0 to 1048575. | 
| Step 14 | mtu mtu-size Example: Device(config-mon-erspan-src-dst)# mtu 512 | Configures the MTU size for truncation. Any ERSPAN packet that is larger than the configured MTU size is truncated to the configured size. The MTU size range is 176 to 9000 bytes. The default value is 9000 bytes. | 
| Step 15 | origin ipv6-address ipv6-address Example: Device(config-mon-erspan-src-dst)# origin ipv6 address 2001:DB8:1::1 | Configures the IPv6 address used as the source of the ERSPAN traffic. | 
| Step 16 | vrf vrf-id Example: Device(config-mon-erspan-src-dst)# vrf 1 | (Optional) Configures the VRF name to use instead of the global routing table. | 
| Step 17 | exit Example: Device(config-mon-erspan-src-dst)# exit | Exits ERSPAN source session destination configuration mode, and returns to ERSPAN source session configuration mode. | 
| Step 18 | no shutdown Example: Device(config-mon-erspan-src)# no shutdown | Enables the configured sessions on an interface. | 
| Step 19 | end Example: Device(config-mon-erspan-src)# end | Exits ERSPAN source session configuration mode, and returns to privileged EXEC mode. | 
The ERSPAN destination session defines the session configuration parameters and the ports that receives the monitored traffic. To define an IPv6 ERSPAN destination session, complete the following procedure:
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable | Enables privileged EXEC mode. Enter your password if prompted. | 
| Step 2 | configure terminal Example: Device# configure terminal | Enters global configuration mode. | 
| Step 3 | monitor session session-number type erspan-destination Example: Device(config)# monitor session 3 type erspan-destination | Defines an ERSPAN destination session using the session ID and the session type, and enters ERSPAN monitor destination session configuration mode.    | 
| Step 4 | description string Example: Device(config-mon-erspan-dst)# description source 1 | (Optional) Describes the ERSPAN destination session.  | 
| Step 5 | destination interface interface-type interface-number Example: Device(config-mon-erspan-dst)# destination interface fortygigabitethernet 1/0/3 | Associates the ERSPAN destination session number with source ports, and selects the traffic direction to be monitored. | 
| Step 6 | source Example: Device(config-mon-erspan-dst)# source | Enters ERSPAN destination session source configuration mode. | 
| Step 7 | erspan-id erspan-flow-id Example: Device(config-mon-erspan-dst-src)# erspan-id 100 | Configures the ID used by source and destination sessions to identify the ERSPAN traffic, which must also be entered in the ERSPAN source session configuration. | 
| Step 8 | ipv6 address ipv6-address Example: Device(config-mon-erspan-dst-src)# ip address 2001:DB8::1 | Configures the IPv6 address that is used as the destination of the ERSPAN traffic. This IPv6 address must be an address on a local interface or loopback interface, and match the address on the destination switch. | 
| Step 9 | exit Example: Switch(config-mon-erspan-dst-src)#exit | Exits ERSPAN destination session source configuration mode, and returns to ERSPAN destination session configuration mode. | 
| Step 10 | no shutdown Example: Device(config-mon-erspan-dst)# no shutdown | Enables the configured sessions on an interface. | 
| Step 11 | end Example: Device(config-mon-erspan-dst)# end | Exits ERSPAN destination session source configuration mode, and returns to privileged EXEC mode. | 
The following sections provide configuration examples for ERSPAN.
The following example shows how to configure an ERSPAN source session:
Device> enable
Device# configure terminal
Device(config)# monitor session 1 type erspan-source
Device(config-mon-erspan-src)# description source1
Device(config-mon-erspan-src)# source interface GigabitEthernet 1/0/1 rx
Device(config-mon-erspan-src)# source interface GigabitEthernet 1/0/4 - 8 tx
Device(config-mon-erspan-src)# source interface GigabitEthernet 1/0/3
Device(config-mon-erspan-src)# destination
Device(config-mon-erspan-src-dst)# erspan-id 100
Device(config-mon-erspan-src-dst)# ip address 10.1.0.2
Device(config-mon-erspan-src-dst)# ip dscp 10
Device(config-mon-erspan-src-dst)# ip ttl 32
Device(config-mon-erspan-src-dst)# mtu 512
Device(config-mon-erspan-src-dst)# origin ip address 10.10.0.1
Device(config-mon-erspan-src-dst)# vrf monitoring
Device(config-mon-erspan-src-dst)# exit
Device(config-mon-erspan-src)# no shutdown
Device(config-mon-erspan-src)# end
  

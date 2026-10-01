---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-17-configuration-g-1130dfb3-3
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-17-configuration-g-1130dfb3"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-17-configuration-g-1130dfb3.md
source_anchor: ""
source_lines: [52, 95]
sha256: 6a4b5df9a52fb1c88801b80fe8eb29e2a54ec730fa547b9331e48aeafdb5bfd3
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-17-configuration-g-1130dfb3

| Step 5 | [no] header-type 3 Example: Device(config-mon-erspan-src)# header-type 3 | (Optional) Configures a switch to Type-III ERSPAN header. The default type is Type-II ERSPAN header. | 
| Step 6 | source {interface interface-type interface-number \| vlan vlan-id} [, \| - \| both \| rx \| tx] Example: Device(config-mon-erspan-src)# source interface fastethernet 0/1 rx | Configures the source interface or the VLAN, and the traffic direction to be monitored. | 
| Step 7 | filter {ip access-group {standard-access-list \| expanded-access-list \| acl-name } \| ipv6 access-group acl-name \| mac access-group acl-name \| sgt sgt-ID [, \| -] \| vlan vlan-ID [, \| -]} Example: Switch(config-mon-erspan-src)# filter vlan 3 | (Optional) Configures source VLAN filtering when the ERSPAN source is a trunk port. The filter sgt sgt-ID command configures SGT filtering in the ERSPAN source session. | 
| Step 8 | destination Example: Device(config-mon-erspan-src)# destination | Enters ERSPAN source session destination configuration mode. | 
| Step 9 | erspan-id erspan-flow-id Example: Device(config-mon-erspan-src-dst)# erspan-id 100 | Configures the ID used by source and destination sessions to identify the ERSPAN traffic, which must also be entered in the ERSPAN destination session configuration. | 
| Step 10 | ip address ip-address Example: Device(config-mon-erspan-src-dst)# ip address 10.1.0.2 | Configures the IP address that is used as the destination of the ERSPAN traffic. | 
| Step 11 | ip dscp dscp-value Example: Device(config-mon-erspan-src-dst)# ip dscp 10 | (Optional) Enables the use of IP differentiated services code point (DSCP) for packets that originate from a circuit emulation (CEM) channel. | 
| Step 12 | ip ttl ttl-value Example: Device(config-mon-erspan-src-dst)# ip ttl 32 | (Optional) Configures the IP TTL value of packets in the ERSPAN traffic. | 
| Step 13 | mtu mtu-size Example: Device(config-mon-erspan-src-dst)# mtu 512 | Configures the MTU size for truncation. Any ERSPAN packet that is larger than the configured MTU size is truncated to the configured size. The MTU size range is 176 to 9000 bytes. The default value is 9000 bytes. | 
| Step 14 | origin ip-address ip-address Example: Device(config-mon-erspan-src-dst)# origin ip address 10.10.0.1 | Configures the IP address used as the source of the ERSPAN traffic. | 
| Step 15 | vrf vrf-id Example: Device(config-mon-erspan-src-dst)# vrf 1 | (Optional) Configures the VRF name to use instead of the global routing table. | 
| Step 16 | exit Example: Device(config-mon-erspan-src-dst)# exit | Exits ERSPAN source session destination configuration mode, and returns to ERSPAN source session configuration mode. | 
| Step 17 | no shutdown Example: Device(config-mon-erspan-src)# no shutdown | Enables the configured sessions on an interface. | 
| Step 18 | end Example: Device(config-mon-erspan-src)# end | Exits ERSPAN source session configuration mode, and returns to privileged EXEC mode. | 
| Note | You cannot include source VLANs and filter VLANs in the same session. | 
The ERSPAN destination session defines the session configuration parameters and the ports that receive the monitored traffic.
To define an IPv4 ERSPAN destination session, complete the following procedure:
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable | Enables privileged EXEC mode. Enter your password if prompted. | 
| Step 2 | configure terminal Example: Device# configure terminal | Enters global configuration mode. | 
| Step 3 | monitor session session-number type erspan-destination Example: Device(config)# monitor session 1 type erspan-destination | Defines an ERSPAN destination session using the session ID and the session type, and enters ERSPAN monitor destination session configuration mode.    | 
| Step 4 | description string Example: Device(config-mon-erspan-dst)# description source1 | (Optional) Describes the ERSPAN destination session.  | 
| Step 5 | destination interface interface-type interface-number Example: Device(config-mon-erspan-dst)# destination interface GigabitEthernet1/0/1 | Associates the ERSPAN destination session number with source ports, and selects the traffic direction to be monitored. | 
| Step 6 | source Example: Device(config-mon-erspan-dst)# source | Enters ERSPAN destination session source configuration mode. | 
| Step 7 | erspan-id erspan-flow-id Example: Device(config-mon-erspan-dst-src)# erspan-id 100 | Configures the ID used by source and destination sessions to identify the ERSPAN traffic, which must also be entered in the ERSPAN source session configuration. | 
| Step 8 | ip address ip-address [force] Example: Device(config-mon-erspan-dst-src)# ip address 10.1.0.2 | Configures the IP address that is used as the destination of the ERSPAN traffic.  | 
| Step 9 | vrf vrf-id Example: Device(config-mon-erspan-dst-src)# vrf 1 | (Optional) Configures the VRF name to use instead of the global routing table. | 
| Step 10 | no shutdown Example: Device(config-mon-erspan-dst-src)# no shutdown | Enables the configured sessions on an interface. | 
| Step 11 | end Example: Device(config-mon-erspan-dst-src)# end | Exits ERSPAN destination session source configuration mode, and returns to privileged EXEC mode. | 
The ERSPAN source session defines the session configuration parameters and the ports or VLANs to be monitored. To define an IPv6 ERSPAN source session, complete the following procedure:
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable | Enables privileged EXEC mode. Enter your password if prompted. | 
| Step 2 | configure terminal Example: Device# configure terminal | Enters the global configuration mode. | 
| Step 3 | monitor session session-number type erspan-source Example: Device(config)# monitor session 1 type erspan-source | Defines an ERSPAN source session using the session ID and the session type, and enters ERSPAN monitor source session configuration mode.  | 
| Step 4 | description string Example: Device(config-mon-erspan-src)# description source1 | (Optional) Describes the ERSPAN source session.  | 
| Step 5 | [no] header-type 3 Example: Device(config-mon-erspan-src)# header-type 3 | (Optional) Configures a switch to Type-III ERSPAN header. The default type is Type-II ERSPAN header. | 
| Step 6 | source {interface interface-type interface-number \| vlan vlan-id} [, \| - \| both \| rx \| tx] Example: Device(config-mon-erspan-src)# source interface fortygigabitethernet 1/0/3 | Configures the source interface or the VLAN, and the traffic direction to be monitored. | 
| Step 7 | filter {ip access-group {standard-access-list \| expanded-access-list \| acl-name } \| ipv6 access-group acl-name \| mac access-group acl-name \| sgt sgt-ID [, \| -] \| vlan vlan-ID [, \| -]} Example: Switch(config-mon-erspan-src)# filter ipv6 access-group exampleacl | (Optional) Configures source VLAN filtering when the ERSPAN source is a trunk port. The filter sgt sgt-ID command configures SGT filtering in the ERSPAN source session. | 
| Step 8 | destination Example: Device(config-mon-erspan-src)# destination | Enters ERSPAN source session destination configuration mode. | 
| Step 9 | erspan-id erspan-flow-id Example: Device(config-mon-erspan-src-dst)# erspan-id 100 | Configures the ID used by source and destination sessions to identify the ERSPAN traffic, which must also be entered in the ERSPAN destination session configuration. | 
| Step 10 | ipv6 address ipv6-address Example: Device(config-mon-erspan-src-dst)# ipv6 address 2001:DB8::1 | Configures the IPv6 address that is used as the destination of the ERSPAN traffic. | 
| Step 11 | ipv6 dscp dscp-value Example: Device(config-mon-erspan-src-dst)# ipv6 dscp 2 | (Optional) Enables the use of IPv6 differentiated services code point (DSCP) for packets that originate from a circuit emulation (CEM) channel. | 

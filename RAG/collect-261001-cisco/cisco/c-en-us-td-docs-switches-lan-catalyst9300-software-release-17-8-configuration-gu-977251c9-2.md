---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-8-configuration-gu-977251c9-2
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-8-configuration-gu-977251c9"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-8-configuration-gu-977251c9.md
source_anchor: ""
source_lines: [35, 82]
sha256: 21bd941e8ff715fc60a31f46fbb59a197b63b16b1a348a10e60b39eef322d63e
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-8-configuration-gu-977251c9

                                    Universal algorithm—The universal load-balancing algorithm allows each device on the network to make a different load sharing decision for each source-destination address pair, which resolves load-sharing imbalances. The device is set to perform universal load sharing by default.
How to Configure Cisco Express Forwarding
CEF or distributed CEF is enabled globally by default. If for some reason it is disabled, you can re-enable it by using the ip cef or ip cef distributed global configuration command.
Procedure
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure terminal Example:  Device# configure terminal   | Enters global configuration mode. | 
| Step 2 | ip cef Example:  Device(config)# ip cef | Enables CEF operation on a non-stacking switch. Go to Step 4. | 
| Step 3 | ip cef distributed Example:  Device(config)# ip cef distributed | Enables CEF operation on a active switch. | 
| Step 4 | interface interface-id Example:  Device(config)# interface gigabitethernet 1/0/1 | Enters interface configuration mode, and specifies the Layer 3 interface to configure. | 
| Step 5 | ip route-cache cef Example:  Device(config-if)# ip route-cache cef | Enables CEF on the interface for software-forwarded traffic. | 
| Step 6 | end Example:  Device(config-if)# end | Returns to privileged EXEC mode. | 
| Step 7 | show ip cef Example:  Device# show ip cef | Displays the CEF status on all interfaces. | 
| Step 8 | show cef linecard [detail] Example:  Device# show cef linecard detail | (Optional) Displays CEF-related interface information on a non-stacking switch. | 
| Step 9 | show cef linecard [slot-number] [detail] Example:  Device# show cef linecard 5 detail | (Optional) Displays CEF-related interface information on a switch by stack member for all switches in the stack or for the specified switch. (Optional) For slot-number , enter the stack member switch number. | 
| Step 10 | show cef interface [interface-id] Example:  Device# show cef interface gigabitethernet 1/0/1 | Displays detailed CEF information for all interfaces or the specified interface. | 
| Step 11 | show adjacency Example:  Device# show adjacency | Displays CEF adjacency table information. | 
| Step 12 | copy running-config startup-config Example:  Device# copy running-config startup-config | (Optional) Saves your entries in the configuration file. | 
| Note | The ip route-cache cef command is enabled by default and it cannot be disabled. | 
How to Configure a Load-Balancing for CEF Traffic
The following sections provide information on configuring load-balancing for CEF traffic.
Enabling or Disabling CEF Per-Destination Load Balancing
To enable or disable CEF per-destination load balancing, perform the following procedure:
Procedure
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example:  Device# enable   | Enters global configuration mode. | 
| Step 2 | configure terminal Example:  Device# configure terminal   | Enters global configuration mode. | 
| Step 3 | interface interface-id Example:  Device(config-if)# interface gigabitethernet 1/0/1  | Enters interface configuration mode, and specifies the Layer 3 interface to configure. | 
| Step 4 | [no] ip load-sharing per-destination Example:  Device(config-if)# ip load-sharing per-destination  | Enables per-destination load balancing for CEF on the interface. The no ip load-sharing per-destination command disables per-destination load balancing for CEF on the interface. | 
| Step 5 | end Example:  Device(config-if)# end | Exits interface configuration mode and returns to privileged EXEC mode. | 
Selecting a Tunnel Load-Balancing Algorithm for CEF Traffic
Select the tunnel algorithm when your network environment contains only a few source and destination pairs. The device is set to perform universal load sharing by default.
To select a tunnel load-balancing algorithm for CEF traffic, perform the following procedure:
Procedure
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example:  Device# enable   | Enters global configuration mode. | 
| Step 2 | configure terminal Example:  Device# configure terminal   | Enters global configuration mode. | 
| Step 3 | ip cef load-sharing algorithm {original \| universal [id] } Example:  Device(config)# ip cef load-sharing algorithm universal | Selects a CEF load-balancing algorithm.  | 
| Step 4 | end Example:  Device(config)# end | Returns to privileged EXEC mode. | 
Example: Enabling or Disabling CEF Per-Destination Load Balancing
Per-destination load balancing is enabled by default when you enable CEF. The following example shows how to disable per-destination load balancing:
Device> enable
Device# configure terminal
Device(config)# interface Ethernet1/0/1
Device(config-if)# no ip load-sharing per-destination
Device(config-if)# end

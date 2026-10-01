---
id: collect-261001-cisco/cisco/cisco-ios-xr-vs-huawei-commands-cheat-sheet-a718c720
title: "cisco-ios-xr-vs-huawei-commands-cheat-sheet-a718c720"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-cisco/cisco-ios-xr-vs-huawei-commands-cheat-sheet-a718c720.md
source_anchor: ""
source_lines: [1, 85]
sha256: 1222548af596583e5f74e581eb271d0bee48ee9c1dc3ae6d3be184bd4a413dd0
---

# cisco-ios-xr-vs-huawei-commands-cheat-sheet-a718c720

Here’s the Cisco IOS-XR commands vs Huawei OS commands
| show | display | 
| \| include | \| include | 
| exit | quit | 
| end | return | 
| ..formal \| |  | 
| configure | system-view | 
| shutdown | shutdown | 
| no shutdown | undo shutdown | 
| no | undo | 
| clear | reset | 
| show running-config | display current-configuration | 
| show version | display version | 
| show tech-support | display diagnostic-information | 
| show configuration (after change) | display configuration candidate | 
| clear counters <interface> | reset counters <int> | 
| clear line <line><line number> |  | 
| terminal length <num> | screen-length <0-512> temporary *value 0 indicates none split screen | 
| terminal width <num> |  | 
| terminal monitor | terminal monitor | 
| terminal monitor disable | undo terminal monitor | 
| show clock | display clock | 
| show processes cpu | display cpu-usage | 
| show logging | display logbuffer | 
| ping | ping | 
| traceroute | tracert | 
| show ip interface brief | display ip interface | 
| show route | display ip routing-table | 
| show route longer-prefixes <X.X.X.X/length> | display ip routing-table <X.X.X.X> <length of IP >longer-match | 
| show route static | display ip routing-table protocol static | 
| show route summary |  | 
| show protocols |  | 
| show route ipv6 | display ipv6 routing-table | 
| show interface be<xx> | display interface Eth-trunkX | 
| monitor interface <interface> |  | 
| reload | reboot | 
| show users | display users | 
| show diag / show inventory |  | 
| show history | display history-command | 
| dir | dir | 
| show environment |  | 
| show memory summary | display memory-usage | 
| show platform | display device pic-status | 
| show controllers <interface> | display controller | 
| show access-lists | display acl <aclname> | 
 
| show osfp neighbor | display ospf nexthop | 
| show ospf database | display ospf lsdb | 
| show ospf interface | display ospf interface | 
| show route ospf | display ip routing-table protocol ospf | 
 
| show clns neighbor | display isis peer | 
| show isis database | display isis lsdb | 
| show isis route | display isis route | 
| show isis topology |  | 
| show isis interface | display isis interface | 
 
| show bgp | display bgp routing-table | 
| show bgp summary | display bgp peer | 
| show bgp neighbors | display bgp peer | 
| show bgp neighbor <ip> advertised-routes | display bgp routing-table peer <ip> advertised-routes | 
| show bgp neighbor <ip>routes | display bgp routing-table peer <ip> received-routes | 
| show bgp ipv6 unicast summary | display bgp  ipv6 peer | 
| show bgp all all summary | display bgp all summary | 
| show bgp ipv6 unicast neigh <ip> advertised-routes | display bgp ipv6 routing-table peer <ip> advertised-routes | 
| show bgp ipv6 unicast neighbor <ip> routes | display bgp ipv6 routing-table peer <ip> received-routes | 
| show bgp regexp | display bgp routing-table regular-expression | 
| show route bgp | display ip routing-table protocol bgp | 
 
| show mpls ldp summary | display mpls ldp all | 
| show mpls interface | display mpls interface | 
| show mpls ldp neighbor brief | display mpls ldp peer | 
| show rsvp interface | display mpls rsvp-te interface | 
| show rsvp neighbors | display mpls rsvp-te peer | 
 
| show mfib/mrib route | display multicast routing-table | 
| show pim interface | display pim interface | 
| show pim neighbor | display pim neighbor | 
| show igmp interface | display igmp interface | 
 
 
							
		 
		
Be First to Comment

---
id: collect-261001-cisco/cisco/wp-content-uploads-2023-06-command-line-cheat-sheets-cisco-juniper-alcatel-huawe-0e8f258a-1
title: "wp-content-uploads-2023-06-command-line-cheat-sheets-cisco-juniper-alcatel-huawe-0e8f258a"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["copyright", "memory", "optics"]
source: docs/RAG/collect-261001-cisco/wp-content-uploads-2023-06-command-line-cheat-sheets-cisco-juniper-alcatel-huawe-0e8f258a.md
source_anchor: ""
source_lines: [1, 129]
sha256: d2dbf2f7d3fc5f63e2d443425e8756b13983c92810f6eeb8fb3a9f22b3e4ae73
---

# wp-content-uploads-2023-06-command-line-cheat-sheets-cisco-juniper-alcatel-huawe-0e8f258a

COMMAND LINE CHEAT SHEET 
COMMAND LINE CHEAT SHEET (Cisco, Juniper, Nokia, Huawei)                                               Copyright © 2018-2019, By Gokhan Kosem www.ipcisco.com 
      
IOS XR HVRP JUNOS SROS 
BASIC 
show show show display 
exit exit / up exit quit 
run run - - 
end exit exit all return 
| include | match | match | include 
… formal | | display-set - - 
reload request system reboot admin reboot now reboot 
GENERAL CONFIGURATION 
show running-config show configuration admin display-config display current-configuration 
show startup-config - - display saved-configuration 
configure terminal configure / edit configure system view 
hostname hostname system host-name  hostname system name  systemname sysname  systemname 
show  (after conf change) show  | compare info (after conf change) - 
commit commit admin save save

COMMAND LINE CHEAT SHEET 
COMMAND LINE CHEAT SHEET (Cisco, Juniper, Nokia, Huawei)                                               Copyright © 2018-2019, By Gokhan Kosem www.ipcisco.com 
      
IOS XR HVRP JUNOS SROS 
shut down disable shut down shut down 
no shut down delete interfaces x disable no shutdown undo shut down 
no delete no undo 
SHOW 
show clock show system uptime show system time display clock 
show ntp status show ntp status show system ntp display ntp-service status 
show history show cli history history display history-command 
show platform show chassis fpc show card,  show mda display device pic-status 
admin show platform show chassis fpc detail show card detail,  show mda detail display device 
show environment show chassis environment - - 
show inventory show chassis hardware - - 
admin show environment | include 
PM show chassis hardware | match PSM show chassis  environment power-
supply display power 
show diags show chassis hardware show chassis environment - 
show memory summary show chassis routing engine show system memory-pools display memory-usage 
show processes cpu show system processes extensive show system cpu display cpu-usage

COMMAND LINE CHEAT SHEET 
COMMAND LINE CHEAT SHEET (Cisco, Juniper, Nokia, Huawei)                                               Copyright © 2018-2019, By Gokhan Kosem www.ipcisco.com 
      
IOS XR HVRP JUNOS SROS 
show users show system users show system users display users 
show version show version show version display version 
show licence - - display licence 
- show system alarms show system alarms display alarm all / active 
- show chassis alarms - - 
show arp show arp show router arp display arp all 
show interface show interfaces show router interface display ip interface 
show interface interface show interfaces interface show port port display ip interface interface 
show interface interface statistics  show port port statistics  
show interface brief show interface terse show router interface summary display ip interface brief 
show policy-map show class-of-service interface show router policy - 
show policy-map interface show interfaces queue - - 
show route show route show router route-table display ip routing-table 
show route summary show route summary show router route-table summary - 
show route ipv6 show route table inet6.0 show router route-table ipv6 display ipv6 routing-table

COMMAND LINE CHEAT SHEET 
COMMAND LINE CHEAT SHEET (Cisco, Juniper, Nokia, Huawei)                                               Copyright © 2018-2019, By Gokhan Kosem www.ipcisco.com 
      
IOS XR HVRP JUNOS SROS 
show route-map show policy show router policy display route-policy 
show snmp show snmp statistics show snmp counters display snmp statistics 
show tcp show system connections show system connections display tcp statistics 
show ipv4 traffic show system statistics - display ip statistics 
show protocols show route protocol - - 
show flash show flash file   ( + dir ) dir flash: 
show filesystem show system storage - dir 
show bfd session show bfd session show router bfd session display bfd session all 
show bfd interfaces location x - show router bfd interface display bfd interface 
show interfaces be x show interfaces aex show lag x display interface Eth-Trunk x 
show interfaces be x details show interfaces aex details show lag x detail - 
- - show lag x associations - 
TSHOOT 
ping ip_address ping  ip_address ping  ip_address ping  ip_address 
traceroute  ip_address traceroute  ip_address traceroute  ip_address tracert  ip_address

COMMAND LINE CHEAT SHEET 
COMMAND LINE CHEAT SHEET (Cisco, Juniper, Nokia, Huawei)                                               Copyright © 2018-2019, By Gokhan Kosem www.ipcisco.com 
      
IOS XR HVRP JUNOS SROS 
debug debug debug debugging 
no debug undebug all no debug undo debugging 
monitor interface interface monitor interface interface monitor port port - 
terminal monitor monitor start messages - terminal monitor /terminal trapping 
terminal monitor disable monitor stop messages - undo  terminal monitor 
show tech-support request support info admin tech-support display diagnostic-information 
show logging show log messages show log log-id 99 (all) display logbuffer 
show controllers interface show interfaces diagnostic optics 
interface - display controller 
show access-lists show firewall show filter ip x display acl x 
CLEAR 
clear clear clear reset 
clear counters interface clear interface statistics interface clear counter interface xx reset counters interface xx 
clear arp-cache clear arp clear router arp reset arp 
clear cef - - reset ip fast-forwarding 
clear route * clear ip route clear router route-adv reset ip forwarding-table statistis 
protocol all

COMMAND LINE CHEAT SHEET 
COMMAND LINE CHEAT SHEET (Cisco, Juniper, Nokia, Huawei)                                               Copyright © 2018-2019, By Gokhan Kosem www.ipcisco.com 
      
IOS XR HVRP JUNOS SROS 
clear access-list counters clear firewall clear filter - 
clear line line request system logout username - - 
OSPF 
show ospf (summary) show ospf overview show router ospf status display ospf brief 
show ospf database show ospf database show router ospf database display ospf lsdb 
show ospf interface show ospf interface show router ospf interface display ospf interface 
show ospf neighbor show ospf neighbor 
 show router ospf neighbor display ospf nexthop 
 
show route ospf show route protocol ospf show router ospf routes display ip routing-table protocol ospf 
show ospf virtual-links - show router ospf virtual-link display ospf vlink 
show ospf statistics show ospf statistics show router ospf statistics display ospf statistics 
ISIS 
display isis interface show isis interface show router isis interface display isis interface 
show clns neighbor show isis adjacency show router isis adjaceny display isis peer 
show isis database show isis database show router isis database display isis lsdb 
show isis topology show isis  topology show router isis topology -

COMMAND LINE CHEAT SHEET 
COMMAND LINE CHEAT SHEET (Cisco, Juniper, Nokia, Huawei)                                               Copyright © 2018-2019, By Gokhan Kosem www.ipcisco.com 
      

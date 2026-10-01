---
id: collect-261001-cisco/cisco/2017-01-tips-cisco-vs-huawei-vs-juniper-basic-cli-commands-53f33e33
title: "2017-01-tips-cisco-vs-huawei-vs-juniper-basic-cli-commands-53f33e33"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-cisco/2017-01-tips-cisco-vs-huawei-vs-juniper-basic-cli-commands-53f33e33.md
source_anchor: ""
source_lines: [1, 135]
sha256: b059927c1e4bd7345f5f2409d71731a484a6395d1c86ffcd4b1f1e4ffccc890b
---

# 2017-01-tips-cisco-vs-huawei-vs-juniper-basic-cli-commands-53f33e33

Are you familiar with the primary user interface used for configuring, monitoring, and maintaining your network devices? The networking leaders such as Cisco, Juniper, Huawei, they have their own basic command-line interface (CLI). What is the CLI Command?
Firstly we take the famous Cisco as an example.
Now in this article we listed some essential and basic Commands of Cisco, Huawei and Juniper, which can help you know the basic differences of commands among Cisco, Huawei and Juniper.
Cisco vs. Huawei Essential Command Mapping
| CISCO | HUAWEI | 
| ping | ping | 
| traceroute | tracert | 
| show | display | 
| show interfaces | display interface | 
| Show ip route | display routing-table | 
| Show ip interface | Display ip interface | 
| Show version | Display version | 
| Show ip bgp | Display bgp routing-table | 
| Show clock | Display clock | 
| Show port | Display port-mapping | 
| Show flash | dir flash: (on user view mode) | 
| Show logging | Display logbuffer | 
| Show snmp | Display snmp-agent statistics | 
| Show frame-relay pvc | Display fr pvc-info | 
| Show users | Display users | 
| Show terminal length | screen-length disable undo screen-length disable | 
| enable | Super | 
| disable | Super 0 (number is privilege level from 0 to 3, where 3 is default and equivalent to “enable” on Cisco) | 
| Conf t | System-view | 
| exit | quit | 
| end | return | 
| Show policy-map interface | Display qos policy interface | 
| send | send (on user view mode) | 
| write terminal (sh run) | display current-configuration | 
| Sh startup | Display saved-configuration | 
| [no equivalent: shows the files used for startup] | Display startup | 
| Write erase | Reset saved-configuration | 
| Write mem (or wr or copy run start) | save | 
| clear counters | reset (on user view mode) Reset counters interface | 
| ? | ? | 
| telnet | telnet | 
| Enable secret (conf mode) | Super pass cipher (system mode) | 
| Term mon | term debu | 
| clock | clock | 
| no | undo | 
| debug / no debug | debugging / undo debugging | 
| copy running-config | Save safely | 
| terminal monitor | terminal monitor | 
| terminal length | screen-length disable undo screen-length disable | 
| terminal no monitor | undo terminal monitor | 
| clear counters | reset counters interface | 
| clear interface | reset counters interface | 
| clear crypto | ipsec sa ike sa | 
| clear access-list counters | reset acl counter all | 
| reload | reboot | 
| shutdown | shutdown | 
| boot | boot bootrom | 
| Aaa | hwtacacs scheme | 
| terminal no monitor | undo terminal monitor | 
| tacacs-server | hwtacacs scheme (in conf command) | 
| snmp-server | tftp-server (in conf command) | 
| router bgp | bgp | 
| Router rip | rip | 
| ip tacacs | hwtacacs nas-ip (this command doesn’t exist !!!) | 
| mtu | Mtu (this command doesn’t exist !!!) | 
| clear ip cef | reset ip fast-forwarding | 
| clear ip route * | reset ip routing-table statistics protocol all | 
| Clear ip bgp | Reset bgp all | 
| Show tech | display diagnostic-information | 
| Sh ip nat translation | Display nat session | 
| Show Controller | display controller (but not relevant for non-modular chassis) | 
| show dsl int atm 0 | display dsl status interface Atm 2/0 | 
| sho atm pvc | Display atm pvc-info | 
| debug pvc nego | Debug atm all (very dangerous – might crash router) | 
| sho crypto isakmp sa | Display ike sa | 
| sho crypto isakmp key | Display ike peer | 
| sho crypto isakmp police | Display ike proposal | 
From https://lifeoflogs.blogspot.com/2011/04/cisco-vs-huawei-essential-command.html
CLI Commands Cisco vs. Juniper Router will Help in Troubleshooting
Basic CLI Commands
| Description | Cisco IOS | Juniper | 
| To Ping | ping | ping | 
|  | traceroute | traceroute | 
| To display date / time | show clock | show system uptime | 
| To display Chassis status | show environment | show chassis environment | 
| To display history of commands entered | show history | show cli history | 
|  | show ip traffic | show system statistics | 
|  | show logging | show log | 
|  | show processes | show system processes | 
|  | show running config | show configuration | 
|  | show tech-support | request support information | 
|  | show users | show system users | 
|  | show version | show version | 
|  | show arp | show arp | 
|  | show interface | show interfaces show interfaces detail show interfaces extensive | 
|  | show ip interface brief | show interfaces terse | 
|  | show ip route | show route | 
|  | show ip route summary | show route summary | 
|  | show route-map | show policy | 
|  | show tcp | show system connections | 
|  | clear counters | clear interface statistics | 
|  | clear arp-cache | clear arp | 
|  | clear line |  | 
|  | clear ip route |  | 
BGP Commands
| Description | Cisco IOS | Juniper | 
|  | show ip bgp | show route protocol bgp | 
|  | show ip bgp community | show route community | 
|  | show ip bgp dampened paths | show route damping decayed | 
|  | show ip bgp neighbors | show bgp neighbor | 
|  | show ip bgp neighbors address advertised-routes | show route advertising-protocol bgp address | 
|  | show ip bgp neighbors address received-routes | show route receive-protocol bgp address | 
|  | show ip bgp peer-group | show bgp group | 
|  | show ip bgp regexp | show route aspath-regex | 
|  | show ip bgp summary | show bgp summary | 
|  | clear ip bgp | clear bgp neighbor | 
|  | clear ip bgp dampening | clear bgp damping | 
OSPF Commands
| Description | Cisco IOS | Juniper | 
|  | show ip ospf database | show ospf database | 
|  | show ip ospf interface | show ospf interface | 
|  | show ip ospf neighbor | show ospf neighbor | 
IS-IS Commands
| Description | Cisco IOS | Juniper | 
|  | show clns neighbor | show isis adjacency | 
|  | show isis database | show isis database | 
|  | show isis route | show isis routes | 
|  | show isis topology | show isis routes | 
|  | show isis spf-log | show isis spf log | 
|  | clear clns neighbor | clear isis adjacency | 
|  | clear isis * | clear isis database | 
From https://forums.juniper.net/t5/Configuration-Library/CLI-commands-Cisco-VS-Juniper-router-will-help-in/td-p/68088
If you have more tips and references about the CLI Commands, you can share with us here. Welcome!
More Related…
Top 10 Commands Every Cisco IOS User Should Know
Top Commands for Verifying Cisco Switch Network Status and Operational State
Check Cisco Routers and Switches Using the IOS Environment Command
Expertise Builds Trust
20+ Years • 200+ Countries • 21500+ Customers/Projects
CCIE · JNCIE · NSE7 · ACDX · HPE Master ASE · Dell Server/AI Expert

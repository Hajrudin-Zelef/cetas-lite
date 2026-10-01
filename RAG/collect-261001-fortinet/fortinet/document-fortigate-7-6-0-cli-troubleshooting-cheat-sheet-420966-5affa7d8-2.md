---
id: collect-261001-fortinet/fortinet/document-fortigate-7-6-0-cli-troubleshooting-cheat-sheet-420966-5affa7d8-2
title: "document-fortigate-7-6-0-cli-troubleshooting-cheat-sheet-420966-5affa7d8"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["decode", "memory", "parameters"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-6-0-cli-troubleshooting-cheat-sheet-420966-5affa7d8.md
source_anchor: ""
source_lines: [105, 189]
sha256: d33fbc9cb0c901a8ae4fc6da396e131c5730c3333244f3997a47f344d9dece8c
---

# document-fortigate-7-6-0-cli-troubleshooting-cheat-sheet-420966-5affa7d8

| diagnose test application urlfilter <option> | Show the web filter debug output for the specified option. | 
| diagnose debug application dnsproxy -1 diagnose debug enable | Start real-time debugging for DNS proxy. DNS proxy is responsible for DNS filter, DNS translation, DNS resolution etc. | 
| diagnose debug enable diagnose test application dnsproxy | List the DNS proxy debug outputs. | 
| diagnose test application dnsproxy <option> | Show the DNS proxy debug output for the specified option. | 
| diagnose ips filter set "host <x.x.x.x> and port <port>" diagnose ips debug enable all diagnose debug enable | Start IPS engine debugs for Application Control and IPS Security profile | 
| diagnose ips debug enable av diagnose ips debug status show diagnose sys scanunit debug all enable diagnose sys scanunit debug level verbose diagnose sys scanunit debug show diagnose debug enable | Start real-time debugging for antivirus profile when antivirus profile is configured in flow mode. | 
| diagnose wad debug enable category scan diagnose wad stream-scan av-test "debug enable" diagnose wad stream-scan av-test "debug all:debug" diagnose sys scanunit debug all enable diagnose sys scanunit debug level verbose diagnose sys scanunit debug show diagnose debug enable | Start real time debugging for antivirus profile when antivirus profile is configured in proxy mode. | 
IPS engine
The IPS engine handles traffic related to flow-based processing.
|  | Real-time debugs are CPU intensive tasks. Running real-time IPS engine debugs with proper filters can result in high CPU usage. | 
| Command | Description | 
|---|---|
| diagnose test application ipsmonitor 1 | Show IPS engine information | 
| diagnose test application ipsmonitor 2 | Set the IPS engine enable/disable status. | 
| diagnose test application ipsmonitor 99 | Restart all IPS engines and monitor. | 
| diagnose test application ipsmonitor 97 | Start all IPS engines. | 
| diagnose test application ipsmonitor 98 | Stop all IPS engines. | 
| diagnose ips session list diagnose test application ipsmonitor 13 | Show the IPS sessions in each engine's memory space. | 
| diagnose ips filter set "host <x.x.x.x> and port <port>" diagnose ips debug enable all diagnose debug enable | Show IPS engine debugs for the traffic specified by the filter. | 
WAD
The WAD daemon handles proxy related processing.
|  | Real-time debugs are CPU intensive tasks. Running real-time WAD debugs with proper filters can result in high CPU usage. | 
| Command | Description | 
|---|---|
| diagnose test application wad 1000 | Show all WAD processes. | 
| diagnose test application wad 2 | Show total memory usage. | 
| diagnose test application wad 99 | Restart all WAD processes. | 
| diagnose wad debug display pid enable diagnose wad filter <filter> diagnose wad filter list diagnose wad debug enable level <level> diagnose wad debug enable category <category> diagnose debug enable | Start real-time debugging of the traffic processed by WAD daemon. | 
| diagnose wad filter <filter> | Set the filter for the WAD debugs. | 
| diagnose wad filter list | Show all the filters that have been set for debugging. | 
| diagnose wad filter clear | Clear the WAD filter settings. | 
| diagnose wad debug enable level <level> | Set the verbosity level of the debugs. | 
| diagnose wad debug enable category <category> | Set the traffic category. | 
| diagnose wad debug display pid enable | Show the WAS worker PID in debugs that handle the session request. | 
| diagnose debug enable | Start printing debugs in the console. | 
CPU profiling
| Command | Description | 
|---|---|
| diagnose sys profile cpumask <cpu_id> | Set the CPU core to profile. | 
| diagnose sys profile start | Start CPU profiling and wait for one to two minutes to stop. | 
| diagnose sys profile stop | Stop CPU profiling. | 
| diagnose sys profile module | Show the applied kernel modules. | 
| diagnose sys profile show detail diagnose sys profile show order | Show the CPU profiling result for the respective core. | 
Tree
| Command | Description | 
|---|---|
| tree | Show the entire command tree. | 
| tree execute | Show the execute command tree. | 
| tree diagnose | Show the diagnose command tree. | 
Routing
IPv4 and IPv6 routing
| Command | Description | 
|---|---|
| get router info routing-table all | Show routing table. | 
| get router info routing-table database get router info6 routing-table database | Show IPv4 and IPv6 routing database information. | 
| diagnose ip route list get router info kernel diagnose ipv6 route list get router info6 kernel | Show the IPv4 and IPv6 kernel routing table. | 
| get router info protocols get router info6 protocols | Show routing protocol information for IPv4 and IPv6. | 
| execute router restart | Restart the routing daemon | 
| get router info ospf status get router info6 ospf status | Show OSPF status for IPv4 and IPv6. | 
| get router info ospf neighbor get router info6 ospf neighbor | Show OSPF neighbors for IPv4 and IPv6. | 
| get router info ospf database brief | Show OSPF database in brief. | 
| get router info bfd neighbor get router info6 bfd neighbor | Show BFD neighbors for IPv4 and IPv6. | 
| diagnose test application bfd 1 diagnose test application bfd 2 diagnose test application bfd 3 | Show BFD statistics. | 
| diagnose debug application bfdd <debug level> diagnose debug enable | Start real-time BFD debugging . | 
| get router info bgp summary get router info6 bgp summary | Show BGP summary for IPv4 and IPv6. | 
| get router info bgp neighbors get router info6 bgp neighbors get router info bgp neighbors <x.x.x.x> advertised-routes get router info6 bgp neighbors <x:x::x:x/m> advertised-routes get router info bgp neighbors <x.x.x.x> received-routes get router info6 bgp neighbors <x:x::x:x/m> received-routes get router info bgp neighbors <x.x.x.x> routes get router info6 bgp neighbors <x:x::x:x/m> routes | Show BGP peer and the advertised and received routes from the BGP peer.  | 
| diagnose ip router bgp all enable diagnose ip router bgp level info diagnose debug enable | Start real-time BGP debugging. | 
| execute router clear bgp {all \| as <ASN> \| ip x.x.x.x \| ipv6 y:y:y:y:y:y:y:y} | Execute a hard reset based on the specified parameters:  | 
| execute router clear bgp {all \| ip x.x.x.x \| ipv6 y:y:y:y:y:y:y:y} soft {in\|out} | Executea soft reset based on the specified parameter:  | 
| get router info ospf status get router info6 ospf status | Show OSPF status for IPv4 and IPv6. | 
| get router info ospf interface get router info6 ospf interface | Show OSPF running on interface for IPv4 and IPv6. | 
| get router info ospf neighbor all get router info6 ospf neighbor all | Show OSFP neighbor information for IPv4 and IPv6. | 
| get router info ospf database brief get router info6 ospf database brief | Show OSPF database in brief for IPv4 and IPv6. | 
| diagnose ip router ospf all enable diagnose ip router ospf level info diagnose debug enable | Start real-time OSPF debugging. | 
Multicast routing
| Command | Description | 
|---|---|
| get router info multicast igmp interface | Show IGMP statistics for an interface. | 
| get router info multicast igmp groups | Show multicast groups subscribed to with IGMP. | 
| diagnose ip multicast get-igmp-limit | Show maximum IGMP states. | 
| diagnose ip router igmp decode enable diagnose ip router igmp level info diagnose debug console timestamp enable diagnose debug enable | Start real-time debugging of IGMP daemon. | 
| execute mrouter clear igmp-interface <interface> | Clear all IGMP entries from one interface. | 
| execute mrouter clear igmp-group <group-address> | Clear all IGMP entries for one or all groups. | 
| get router info multicast pim sparse-mode <interface>. | Show sparse-mode interface information. | 
| get router info multicast pim sparse-mode <neighbor> | Show sparse-mode neighbor information. | 

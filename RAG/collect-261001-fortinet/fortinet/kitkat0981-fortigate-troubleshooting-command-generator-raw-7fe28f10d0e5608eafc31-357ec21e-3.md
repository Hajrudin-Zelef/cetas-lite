---
id: collect-261001-fortinet/fortinet/kitkat0981-fortigate-troubleshooting-command-generator-raw-7fe28f10d0e5608eafc31-357ec21e-3
title: "kitkat0981-fortigate-troubleshooting-command-generator-raw-7fe28f10d0e5608eafc31-357ec21e"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["memory", "parameters"]
source: docs/RAG/collect-261001-fortinet/kitkat0981-fortigate-troubleshooting-command-generator-raw-7fe28f10d0e5608eafc31-357ec21e.md
source_anchor: ""
source_lines: [300, 472]
sha256: ab6a84c1c5d31507333ef8e08a3ee1e7c20839693965ee1d75cd2e2bf0e9a245
---

# kitkat0981-fortigate-troubleshooting-command-generator-raw-7fe28f10d0e5608eafc31-357ec21e

WAD
The WAD daemon handles proxy related processing.
Real-time debugs are CPU intensive tasks. Running real-
time WAD debugs with proper filters can result in high
CPU usage.
Command Description
diagnose test application wad 1000 Show all WAD processes.
diagnose test application wad 2 Show total memory usage.
diagnose test application wad 99 Restart all WAD processes.
diagnose wad debug display pid enable
diagnose wad filter <filter>
diagnose wad filter list
diagnose wad debug enable level
<level>
diagnose wad debug enable category
<category>
diagnose debug enable
Start real-time debugging of the
traffic processed by WAD daemon.
diagnose wad filter <filter> Set the filter for the WAD debugs.
diagnose wad filter list Show all the filters that have been
set for debugging.
diagnose wad filter clear Clear the WAD filter settings.
diagnose wad debug enable level
<level>
Set the verbosity level of the
debugs.
diagnose wad debug enable category
<category>
Set the traffic category.
diagnose wad debug display pid enable Show the WAS worker PID in debugs
that handle the session request.
diagnose debug enable Start printing debugs in the console.
CPU profiling
Command Description
diagnose sys profile cpumask <cpu_id> Set the CPU core to profile.
diagnose sys profile start Start CPU profiling and wait for one
to two minutes to stop.
diagnose sys profile stop Stop CPU profiling.
diagnose sys profile module Show the applied kernel modules.
diagnose sys profile show detail
diagnose sys profile show order
Show the CPU profiling result for the
respective core.
Tree
Command Description
tree Show the entire command tree.
tree execute Show the execute command tree.
tree diagnose Show the diagnose command tree.
IPv4 and IPv6 routing
Command Description
get router info routing-table all Show routing table.
get router info routing-table database
get router info6 routing-table
database
Show IPv4 and IPv6 routing
database information.
diagnose ip route list
get router info kernel
diagnose ipv6 route list
get router info6 kernel
Show the IPv4 and IPv6 kernel
routing table.
get router info protocols
get router info6 protocols
Show routing protocol information
for IPv4 and IPv6.
execute router restart Restart the routing daemon
get router info ospf status
get router info6 ospf status
Show OSPF status for IPv4 and
IPv6.
Command Description
get router info ospf neighbor
get router info6 ospf neighbor
Show OSPF neighbors for IPv4 and
IPv6.
get router info ospf database brief Show OSPF database in brief.
get router info bfd neighbor
get router info6 bfd neighbor
Show BFD neighbors for IPv4 and
IPv6.
diagnose test application bfd 1
diagnose test application bfd 2
diagnose test application bfd 3
Show BFD statistics.
diagnose debug application bfdd <debug
level>
diagnose debug enable
Start real-time BFD debugging .
get router info bgp summary
get router info6 bgp summary
Show BGP summary for IPv4 and
IPv6.
get router info bgp neighbors
get router info6 bgp neighbors
get router info bgp neighbors
<x.x.x.x> advertised-routes
get router info6 bgp neighbors
<x:x::x:x/m> advertised-routes
get router info bgp neighbors
<x.x.x.x> received-routes
get router info6 bgp neighbors
<x:x::x:x/m> received-routes
get router info bgp neighbors
<x.x.x.x> routes
get router info6 bgp neighbors
<x:x::x:x/m> routes
Show BGP peer and the advertised
and received routes from the BGP
peer.
l Substitute <x.x.x.x> with IPv4
address of the peer.
l Substitute <x:x::x:x/m> with
IPv6 address of the peer.
diagnose ip router bgp all enable
diagnose ip router bgp level info
diagnose debug enable
Start real-time BGP debugging.
execute router clear bgp {all | as
<ASN> | ip x.x.x.x | ipv6
y:y:y:y:y:y:y:y}
Execute a hard reset based on the
specified parameters:
l all: all BGP peers
l as <ASN>: BGP peers specified
by AS number
l ip x.x.x.x: BGP peer
specified by IPv4 address
(x.x.x.x)
l ipv6 y:y:y:y:y:y:y:y: BGP
peer specified by IPv6 address
(y:y:y:y:y:y:y:y)
execute router clear bgp {all | ip
x.x.x.x | ipv6 y:y:y:y:y:y:y:y} soft
{in|out}
Executea soft reset based on the
specified parameter:
l all: all BGP peers
l ip x.x.x.x: BGP peer
specified by IPv4 address
(x.x.x.x)
l ipv6 y:y:y:y:y:y:y:y: BGP
peer specified by IPv6 address
(y:y:y:y:y:y:y:y)
l in: received BGP routes only
l out: advertised BGP routes
only
A soft reset will occur in both
directions if neither in nor out is
specified.
get router info ospf status
get router info6 ospf status
Show OSPF status for IPv4 and
IPv6.
get router info ospf interface
get router info6 ospf interface
Show OSPF running on interface for
IPv4 and IPv6.
get router info ospf neighbor all
get router info6 ospf neighbor all
Show OSFP neighbor information for
IPv4 and IPv6.
get router info ospf database brief
get router info6 ospf database brief
Show OSPF database in brief for
IPv4 and IPv6.
diagnose ip router ospf all enable
diagnose ip router ospf level info
diagnose debug enable
Start real-time OSPF debugging.
FortiOS 7.6 Troubleshooting Cheat Sheet Fortinet Inc. 3


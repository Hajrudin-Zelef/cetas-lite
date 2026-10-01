---
id: collect-261001-fortinet/fortinet/kitkat0981-fortigate-troubleshooting-command-generator-raw-7fe28f10d0e5608eafc31-357ec21e-2
title: "kitkat0981-fortigate-troubleshooting-command-generator-raw-7fe28f10d0e5608eafc31-357ec21e"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-fortinet/kitkat0981-fortigate-troubleshooting-command-generator-raw-7fe28f10d0e5608eafc31-357ec21e.md
source_anchor: ""
source_lines: [134, 299]
sha256: a76a33bbe4a6a0cc46622d6b56336a6e259dee084c3ea6dd996746606f92c282
---

# kitkat0981-fortigate-troubleshooting-command-generator-raw-7fe28f10d0e5608eafc31-357ec21e

Command Description
diagnose sys session list Show session table after filtering.
diagnose sys session clear Clear the session table for the
specified filter.
diagnose firewall iprope list Show FortiGate’s internal firewall
table.
Network diagnostics
Command Description
execute ping-options {options}
execute ping <x.x.x.x>
Ping IP address <x.x.x.x> using the
specified options.
execute ssh-options {options}
execute ssh <x.x.x.x>
SSH to IP address <x.x.x.x> using
the specified options.
execute traceroute-options {options}
execute traceroute <x.x.x.x>
Traceroute IP address <x.x.x.x>
using the specified options.
get system arp
diagnose ip arp list
Show ARP entries.
diagnose netlink brctl list Show the names of all of the
switches on the FortiGate.
diagnose netlink brctl name host
<switch-name>
Show the switching table of the
specified switch.
get system interface
get sys interface physical
Show a summary of interface
details, including IP address
information.
diagnose ip address list Show IP address information.
diagnose hardware deviceinfo nic
<interface>
get hardware nic <interface>
Show detailed interface information.
get sys interface transceiver Show connected transceivers.
Packet sniffer
Command Description
diagnose sniffer packet <interface>
<'filter'> <verbose> <count> <a|l>
Execute the inbuilt packet sniffer,
filtered on a particular interface with
the specified filter. For more
information, see Performing a sniffer
trace or packet capture.
Debug flow
Command Description
diagnose debug reset Stop all the prior debugs that were
enabled and running in the
foreground or background.
diagnose debug flow filter clear Clear any IPv4 debug flow filters.
diagnose debug flow filter6 clear Clear any IPv6 debug flow filters.
diagnose debug flow filter <filter> Set a filter for running IPv4 traffic
debug flows.
diagnose debug flow filter6 <filter> Set a filter for running IPv6 traffic
debug flows.
diagnose debug flow show function-name
enable
Show the function name of the code
that the traffic accesses.
diagnose debug flow show iprope enable Show which internal firewall policy
that the traffic is going through.
diagnose debug console timestamp
enable
Start printing timestamps on
debugs.
diagnose debug flow trace start <n> Show n lines of IPv4 debugs.
diagnose debug flow trace start6 <n> Show n lines of IPv6 debugs.
diagnose debug enable Start printing debugs in the console.
For more detailed debug flow filter information, see
Technical Tip: Using filters to review traffic traversing the
FortiGate.
UTM
Command Description
diagnose debug urlfilter <filter>
diagnose debug application urlfilter -
1
diagnose debug enable
Start real-time debugging for web
filter traffic.
diagnose debug enable
diagnose test application urlfilter
List the web filter debug outputs.
diagnose test application urlfilter
<option>
Show the web filter debug output
for the specified option.
diagnose debug application dnsproxy -1
diagnose debug enable
Start real-time debugging for DNS
proxy. DNS proxy is responsible for
DNS filter, DNS translation, DNS
resolution etc.
diagnose debug enable
diagnose test application dnsproxy
List the DNS proxy debug outputs.
diagnose test application dnsproxy
<option>
Show the DNS proxy debug output
for the specified option.
diagnose ips filter set "host
<x.x.x.x> and port <port>"
diagnose ips debug enable all
diagnose debug enable
Start IPS engine debugs for
Application Control and IPS Security
profile
diagnose ips debug enable av
diagnose ips debug status show
diagnose sys scanunit debug all enable
diagnose sys scanunit debug level
verbose
diagnose sys scanunit debug show
diagnose debug enable
Start real-time debugging for
antivirus profile when antivirus
profile is configured in flow mode.
diagnose wad debug enable category
scan
diagnose wad stream-scan av-test
"debug enable"
diagnose wad stream-scan av-test
"debug all:debug"
diagnose sys scanunit debug all enable
diagnose sys scanunit debug level
verbose
diagnose sys scanunit debug show
diagnose debug enable
Start real time debugging for
antivirus profile when antivirus
profile is configured in proxy mode.
IPS engine
The IPS engine handles traffic related to flow-based processing.
Real-time debugs are CPU intensive tasks. Running real-
time IPS engine debugs with proper filters can result in
high CPU usage.
Command Description
diagnose test application ipsmonitor 1 Show IPS engine information
diagnose test application ipsmonitor 2 Set the IPS engine enable/disable
status.
diagnose test application ipsmonitor
99
Restart all IPS engines and monitor.
diagnose test application ipsmonitor
97
Start all IPS engines.
diagnose test application ipsmonitor
98
Stop all IPS engines.
diagnose ips session list
diagnose test application ipsmonitor
13
Show the IPS sessions in each
engine's memory space.
diagnose ips filter set "host
<x.x.x.x> and port <port>"
diagnose ips debug enable all
diagnose debug enable
Show IPS engine debugs for the
traffic specified by the filter.
FortiOS 7.6 Troubleshooting Cheat Sheet Fortinet Inc. 2


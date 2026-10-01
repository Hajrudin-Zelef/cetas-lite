---
id: collect-261001-fortinet/fortinet/kitkat0981-fortigate-troubleshooting-command-generator-raw-7fe28f10d0e5608eafc31-357ec21e-5
title: "kitkat0981-fortigate-troubleshooting-command-generator-raw-7fe28f10d0e5608eafc31-357ec21e"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-fortinet/kitkat0981-fortigate-troubleshooting-command-generator-raw-7fe28f10d0e5608eafc31-357ec21e.md
source_anchor: ""
source_lines: [665, 871]
sha256: 3e90944ddd648ae3f9a12609378e303d07e2f31f3e7efb7ee064438585c679ab
---

# kitkat0981-fortigate-troubleshooting-command-generator-raw-7fe28f10d0e5608eafc31-357ec21e

SSL VPN
SSL VPN web mode has become Agentless VPN, and SSL
VPN tunnel mode is no longer supported in 7.6.3 and
later. Therefore, SSL VPN related debug commands may
not work as expected.
Command Description
diagnose vpn ssl debug-filter list Show any filters that are set for SSL
VPN debug.
diagnose vpn ssl debug-filter clear Clear any filters that are set for SSL
VPN daemon debug.
diagnose vpn ssl debug-filter <filter> Set a filter for SSL VPN debugs.
diagnose debug application sslvpn -1
diagnose debug enable
Start SSL VPN debugs for traffic
that the filter is applied to.
diagnose vpn ssl list
get vpn ssl monitor
execute vpn sslvpn list
Show the current SSL VPN sessions
for both web and tunnel mode.
diagnose vpn ssl statistics
diagnose vpn ssl mux-stat
Show the SSL VPN statistics.
execute vpn sslvpn list Show all SSL VPN web and tunnel
mode connections.
execute vpn sslvpn del-tunnel Disconnect the users from tunnel
mode SSL VPN connection.
execute vpn sslvpn del-web Disconnect the users from web
mode SSL VPN connection.
Managed FortiSwitches
The successful execution of commands for managed
FortiSwitches requires that the feature is available on the
FortiSwitch device itself. See the FortiSwitchOS Feature
Matrix.
Enter ? to view additional options or parameters required
to obtain the required information in the diagnose
switch-controller switch-info commands.
Command Description
diagnose switch-controller switch-info
mac-table
Show managed FortiSwitch MAC
address list.
diagnose switch-controller switch-info
port-stats
Show managed FortiSwitch port
statistics.
diagnose switch-controller switch-info
trunk status
Show managed FortiSwitch trunk
information.
diagnose switch-controller switch-info
mclag
Show MCLAG related information
from FortiSwitch.
diagnose switch-controller switch-info
poe
Show POE-related information.
diagnose switch-controller switch-info
lldp
Show LLDP-related information.
diagnose switch-controller switch-info
port-properties
Show managed FortiSwitch port
properties.
diagnose switch-controller switch-info
acl-counters
Show managed FortiSwitch port
ACL counters information.
diagnose switch-controller switch-info
pdu-counters-list
Show managed FortiSwitch pdu-
counters information.
diagnose switch-controller switch-info
flapguard
Show managed FortiSwitch
flapguard information.
diagnose switch-controller switch-info
qos-stats
Show managed FortiSwitch QoS
statistics.
diagnose switch-controller switch-info
modules
Show modules related information
from FortiSwitch.
diagnose switch-controller switch-info
stp
Show managed FortiSwitch STP
instance status.
Command Description
diagnose switch-controller switch-info
bpdu-guard-status
Show managed FortiSwitch STP
BPDU guard status.
diagnose switch-controller switch-info
igmp-snooping
Show managed FortiSwitch IGMP
snooping information.
diagnose switch-controller switch-info
loop-guard
Show managed FortiSwitch loop-
guard status.
diagnose switch-controller switch-info
dhcp-snooping
Show managed FortiSwitch DHCP
snooping interface list.
diagnose switch-controller switch-info
arp-inspection
Show managed FortiSwitch ARP
inspection interface list.
diagnose switch-controller switch-info
option82-mapping
Show managed FortiSwitch DHCP
option 82 mapping information.
diagnose switch-controller switch-info
802.1X
Show managed FortiSwitch port
802.1X status.
diagnose switch-controller switch-info
802.1X-dacl
Show managed FortiSwitch port
802.1X dynamic ACL status.
diagnose switch-controller switch-info
mac-limit-violations
Show managed FortiSwitch violated
MACs information.
diagnose switch-controller switch-info
flow-tracking
Show managed FortiSwitch flow
information.
diagnose switch-controller switch-info
mirror
Show managed FortiSwitch mirror
information.
diagnose switch-controller switch-info
ip-source-guard
Show managed FortiSwitch source
guard information in hardware.
diagnose switch-controller switch-info
rpvst
Show managed FortiSwitch STP
port information when inter-
operating with rapid PVST network.
execute switch-controller get-conn-
status <FortiSwitch-SN>
Show FortiSwitch connection
status.
execute switch-controller get-
physical-conn standard <FortiSwitch-
SN>
Show FortiLink connectivity graph.
execute switch-controller diagnose-
connection <FortiSwitch-SN>
Show FortiSwitch connection
diagnostics.
Managed FortiAPs
Command Description
diagnose wireless-controller wlac -c
wtp
diagnose wireless-controller wlac -d
wtp
Show information about the FortiAP
devices.
diagnose wireless-controller wlac -c
sta
diagnose wireless-controller wlac -d
sta
Show information about the wireless
clients connected to the FortiAP
devices.
diagnose wireless-controller wlac help Show a list of debug options
available for the wireless controller.
diagnose wireless-controller wlac sta_
filter
diagnose wireless-controller wlac sta_
filter clear
diagnose wireless-controller wlac sta_
filter <aa:bb:cc:dd:ee:ff>255
diagnose debug enable
Start real-time debugging of a
wireless client/station that connects
to the FortiAP.
l <aa:bb:cc:dd:ee:ff>: MAC
address of endpoint/station
diagnose wireless-controller wlac -c
vap
Show virtual access point
information, including its MAC
address, BSSID, SSID, the interface
name, and the IP address of the APs
that are broadcasting it.
diagnose wireless-controller wlac wtp_
filter
diagnose wireless-controller wlac wtp_
filter clear
diagnose wireless-controller wlac wtp_
filter <FAP-SN> 0-<x.x.x.x>:5246 255
diagnose debug application cw_acd
0x7ff
Show the wireless termination point
(WTP), or FortiAP, debugging on the
wireless controller if FortiAP is
failing to connect to FortiGate.
l <FAP-SN>: FortiAP serial
number
l <x.x.x.x>: FortiAP IP address
FortiOS 7.6 Troubleshooting Cheat Sheet Fortinet Inc. 5


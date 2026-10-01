---
id: collect-261001-fortinet/fortinet/kitkat0981-fortigate-troubleshooting-command-generator-raw-7fe28f10d0e5608eafc31-357ec21e-4
title: "kitkat0981-fortigate-troubleshooting-command-generator-raw-7fe28f10d0e5608eafc31-357ec21e"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["agent", "decode"]
source: docs/RAG/collect-261001-fortinet/kitkat0981-fortigate-troubleshooting-command-generator-raw-7fe28f10d0e5608eafc31-357ec21e.md
source_anchor: ""
source_lines: [473, 664]
sha256: 7c69d4116cd5fff6da9724dc8679053ccdb96f1bd37d00d39d627ec339a02cbb
---

# kitkat0981-fortigate-troubleshooting-command-generator-raw-7fe28f10d0e5608eafc31-357ec21e

Multicast routing
Command Description
get router info multicast igmp
interface
Show IGMP statistics for an
interface.
get router info multicast igmp groups Show multicast groups subscribed
to with IGMP.
diagnose ip multicast get-igmp-limit Show maximum IGMP states.
diagnose ip router igmp decode enable
diagnose ip router igmp level info
diagnose debug console timestamp
enable
diagnose debug enable
Start real-time debugging of IGMP
daemon.
execute mrouter clear igmp-interface
<interface>
Clear all IGMP entries from one
interface.
execute mrouter clear igmp-group
<group-address>
Clear all IGMP entries for one or all
groups.
get router info multicast pim sparse-
mode <interface>.
Show sparse-mode interface
information.
get router info multicast pim sparse-
mode <neighbor>
Show sparse-mode neighbor
information.
get router info multicast pim sparse-
mode rp-mapping
Show RP to group mapping
information.
get router info multicast pim sparse-
mode table
Show sparse-mode routing table.
diagnose ip router pim-sm events
enable
diagnose ip router pim-sm all enable
diagnose ip router pim-sm level info
diagnose debug enable
Start real-time debugging of PIM
sparse mode.
SD-WAN
Command Description
diagnose sys sdwan health-check status Show SD-WAN health check
statistics.
diagnose sys sdwan service4
diagnose sys sdwan service6
Show SD-WAN rules in control
plane.
diagnose sys sdwan member Show SD-WAN members.
diagnose firewall proute list Show SDWAN rule and policy routes
in the data plane.
diagnose sys link-monitor status
diagnose sys link-monitor interface
<interface>
Show link monitoring statistics.
diagnose debug application link-
monitor -1
diagnose debug enable
Start real-time link monitor
debugging.
diagnose test application lnkmtd 1
diagnose test application lnkmtd 2
diagnose test application lnkmtd 3
Show link monitoring statistics.
Authentication
Command Description
diagnose firewall auth filter <filter> Set the filter used to list entries.
diagnose firewall auth list List filtered, authenticated IPv4
users.
diagnose wad user list List current users authenticated by
proxy (wad daemon).
diagnose debug application fnbamd -1
diagnose debug application authd -1
diagnose debug enable
Start real-time debugging for
remote and local authentication.
diagnose test authserver <auth_
protocol> <server_name> <user>
<password>
Test authentication directly from the
CLI.
Caution: The password is visible in
clear text; be careful when capture
this command to a log file.
Command Description
diagnose test authserver ldap <server_
name> <user> <password>
Test user authentication using an
LDAP server.
Caution: The password is visible in
clear text; be careful when capture
this command to a log file.
diagnose test authserver radius
<server_name> <auth_type> <user>
<password>
Test user authentication using a
Radius server.
Caution: The password is visible in
clear text; be careful when capture
this command to a log file.
diagnose debug fsso-polling detail
diagnose debug fsso-polling summary
Show information about the polls
from FortiGate to DC.
diagnose debug fsso-polling user
diagnose debug authd fsso list
Show FSSO logged on users when
Fortigate polls the DC.
diagnose debug application fssod -1
diagnose debug application smbcd -1
diagnose debug enable
Start real-time debugging when the
FortiGate is used for FSSO polling.
diagnose debug fsso-polling refresh-
user
execute fsso refresh
Refresh the current logged on FSSO
users and refresh the list.
Caution: This command can cause
an outage, use it carefully.
diagnose debug authd fsso server-
status
Show current status of connection
between FortiGate and the collector
agent.
diagnose debug application authd 8256
diagnose debug enable
Start real-time debugging for the
connection between FortiGate and
the collector agent.
diagnose debug authd fsso refresh-
logons
Resend the logged-on users list to
FortiGate from the collector agent.
diagnose debug application authd 8256
diagnose debug enable
Start real-time debugging for the
connection between FortiGate and
the collector agent.
diagnose debug application samld -1
diagnose debug enable
Start real-time SAML debugging.
IPsec
Command Description
diagnose vpn ike gateway list Show IPsec phase 1 information.
diagnose vpn tunnel list Show IPsec phase 2 information.
get vpn ipsec tunnel summary
get vpn ipsec tunnel details
Show summary and detailed
information about IPsec tunnels.
diagnose vpn tunnel flush Flush all Phase2 tunnel SAs
(Security Associations).
diagnose vpn tunnel flush <name>
[name]
Flush one or more specific Phase2
tunnels by name.
diagnose vpn ike gateway <clear |
flush>
Clear/flush IKE gateways (Phase1).
Apply diagnose vpn ike gateway
filter to filter on specific
gateways.
diagnose vpn ike gateway <clear |
flush> name <name>
Clear/flush a specific IKE gateway
(Phase1) by name.
diagnose vpn ike gateway filter Use various options to filter the IKE
gateways.
diagnose vpn ipsec status Show information about encryption
counters.
diagnose vpn ike log filter <filter> Set a filter for IKE daemon debugs.
diagnose debug application ike -1
diagnose debug enable
Start real-time debugging of IKE
daemon with the filter set.
diagnose vpn ike restart Restart the IKE process.
diagnose vpn ike counts
diagnose vpn ike routes
diagnose vpn ike errors
diagnose vpn ike stats
diagnose vpn ike status
diagnose vpn ike crypto
Show other information, such as IKE
counts, routes, errors, and statistics.
FortiOS 7.6 Troubleshooting Cheat Sheet Fortinet Inc. 4


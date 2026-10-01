---
id: collect-261001-fortinet/fortinet/kitkat0981-fortigate-troubleshooting-command-generator-raw-7fe28f10d0e5608eafc31-357ec21e-6
title: "kitkat0981-fortigate-troubleshooting-command-generator-raw-7fe28f10d0e5608eafc31-357ec21e"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/kitkat0981-fortigate-troubleshooting-command-generator-raw-7fe28f10d0e5608eafc31-357ec21e.md
source_anchor: ""
source_lines: [872, 992]
sha256: 4e13aa8a7f5527b92c5bbc1a3c871f3587f66b1b368da4be94423197866b5e54
---

# kitkat0981-fortigate-troubleshooting-command-generator-raw-7fe28f10d0e5608eafc31-357ec21e

High availability
Command Description
diagnose system ha status
get system ha status
Show HA status and information.
execute ha manage <index> <username> Log into and manage a specific HA
member.
diagnose sys ha checksum cluster Show checksum information of all
cluster members.
diagnose sys ha checksum show <vdom> Show detailed checksum
information for a VDOM.
diagnose sys ha checksum recalculate Recalculate HA checksums.
diagnose sys ha recalculate-extfile-
signature
Recalculate HA external files
signatures.
diagnose sys ha reset-uptime Reset the HA uptime. This is used to
test failover.
diagnose debug application hatalk -1
diagnose debug application hasync -1
diagnose debug application harelay -1
diagnose debug enable
Start real-time debugging of HA
daemons.
diagnose sys ha history read Show HA history.
execute ha synchronize stop
execute ha synchronize start
Manually start and stop HA
synchronization.
ZTNA
The WAD daemon handles proxy related processing.
The FortiClient NAC daemon (fcnacd) handles FortiGate
to EMS connectivity.
Command Description
diagnose endpoint fctems test-
connectivity <EMS>
Verify FortiGate to FortiClient EMS
connectivity.
execute fctems verify <EMS> Verify the FortiClient EMS’s
certificate.
diagnose test application fcnacd 2 Dump the EMS connectivity
information.
diagnose debug app fcnacd -1
diagnose debug enable
Run real-time FortiClient NAC
daemon debugs.
diagnose endpoint ec-shm list <ip>
<mac> <EMS_serial_number> <EMS_
tenant_id>
Show the endpoint record list.
Optionally, add filters.
diagnose endpoint lls-comm send ztna
find-uid <uid> <EMS_serial_number>
<EMS_tenant_id>
Query endpoints by client UID, EMS
serial number, and EMS tenant ID.
diagnose endpoint lls-comm send ztna
find-ip-vdom <ip> <vdom>
Query endpoints by the client IP-
VDOM pair.
diagnose wad dev query-by uid <uid>
<EMS_serial_number> <EMS_tenant_id>
Query from WAD diagnose
command by UID, EMS serial
number, and EMS tenant ID.
diagnose wad dev query-by ipv4 <ip> Query from WAD diagnose
command by IP address.
diagnose firewall dynamic list List EMS security posture tags and
all dynamic IP and MAC addresses.
diagnose test application fcnacd 7
diagnose test application fcnacd 8
Check the FortiClient NAC daemon
ZTNA and route cache.
diagnose wad worker policy list Display statistics associated with
application gateway rules.
diagnose wad debug enable category all
diagnose wad debug enable level
verbose
diagnose debug enable
Run real-time WAD debugs.
diagnose debug reset Reset debugs when completed
Logging
Command Description
diagnose log test Generate logs for testing.
execute log filter <filter> Set log filters.
execute log filter Show log filters.
exec log display Show filtered logs.
execute log delete Delete filtered logs.
diagnose debug application miglogd -1
diagnose debug enable
Start real-time debugging of logging
process miglogd.
execute log fortianalyzer test-
connectivity
Test connectivity between
FortiGate and FortiAnalyzer.
Traffic shaping
Command Description
diagnose firewall shaper traffic-
shaper list
Show configured traffic shapers.
diagnose firewall shaper traffic-
shaper stats list
Show traffic shaper statistics.
SIP session helper
Command Description
diagnose sys sip status Show SIP status.
diagnose sys sip mapping list Show SIP mapping list.
diagnose sys sip dialog list Show SIP dialogue list.
diagnose debug application sip -1
diagnose debug enable
Start real-time SIP debugging.
SIP ALG
Command Description
diagnose sys sip-proxy calls list Show list of active SIP proxy calls.
diagnose sys sip-proxy stats Show SIP proxy statistics.
diagnose sys sip-proxy session list Show SIP proxy session list.
diagnose debug application sip -1
diagnose debug enable
Start real-time SIP debugging.
FortiOS 7.6 Troubleshooting Cheat Sheet Fortinet Inc. 6

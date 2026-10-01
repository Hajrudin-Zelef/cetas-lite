---
id: collect-261001-fortinet/fortinet/fortigate-3-technical-tip-collecting-information-for-ha-issues-95845-554d3b3a
title: "fortigate-3-technical-tip-collecting-information-for-ha-issues-95845-554d3b3a"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-fortinet/fortigate-3-technical-tip-collecting-information-for-ha-issues-95845-554d3b3a.md
source_anchor: ""
source_lines: [1, 50]
sha256: 4946da887566001f721abfe197cb6eb0a165a56f456858b7f2a041eb1f87a909
---

# fortigate-3-technical-tip-collecting-information-for-ha-issues-95845-554d3b3a

Technical Tip: Collecting information for HA issues
Description
Solution
- The output of the following debug commands should be collected for any HA-related issues:
diagnose debug enable
diagnose debug console timestamp enable
diagnose debug application hatalk -1 <----- HA formation issues.
diagnose debug application hasync -1 <----- HA Sync issues.
execute ha synchronize start <----- Start HA Sync.
To disable debug:
diagnose debug disable <----- Command to disable the debug.
The HA framework consists of two daemons: hatalk and hasync. The 'hatalk' process monitors cluster management and failure monitoring. The 'hasync' process handles the synchronization of configuration files, the upgrade process, IKE notifications, external files, the ARP table, and the forwarding information base (FIB).
If no output is generated in the hasync or hatalk debugs, restarting the daemons may be necessary. This can be done by running the commands below on each unit.
To determine the process IDs running for hasync and hatalk:
diagnose sys process pidof hasync
diagnose sys process pidof hatalk
To restart the process:
diagnose system kill 11 <process_id>
Alternatively, to kill all simultaneously running HA processes, execute the following with super admin rights:
fnsysctl killall hatalk
fnsysctl killall hasync
- Run the following on both Primary/Secondary units and collect the info:
get system performance status
get system status
get system ha status
diagnose sys ha status
diagnose sys ha history read
diagnose debug crashlog read
diagnose sys ha checksum show
diagnose sys ha dump 5
diagnose sys ha dump-by group
execute ha synchronize stop <----- Stop synchronization.
execute ha synchronize start <----- Start on the backup unit first.
diagnose sys ha checksum recalculate
To access the secondary device in the CLI, run the following:
execute ha manage <Index-ID> <Admin-Username>
See this article: Technical Tip: How to access the secondary unit from the primary with the 'execute ha manage' command.
- Packet captures for seeing communication between HA ports:
diagnose hardware device nic <heartbeat interface>
diagnose sniffer packet port_ha "" 4 0 l    <----- port_ha should be the heartbeat interface.
To capture only the heartbeat packets:
diagnose sniffer packet any 'ether proto 0x8890' 4 0 l
There are three EtherType configurations in HA
- ha-eth-type: 8890 - Standard NAT/Route mode heartbeat.
- hc-eth-type: 8892 - Transparent mode heartbeat or session sync.
- l2ep-eth-type: 8893 - HA configuration synchronization (Layer 2 Endpoint).
Configuration parameters can be verified with the following commands:
get system ha
show full sys ha | grep -f eth
- Collect the FortiGate's HA and System Event logs for both units downloaded from the GUI/FortiAnalyzer or syslog (remote) server.

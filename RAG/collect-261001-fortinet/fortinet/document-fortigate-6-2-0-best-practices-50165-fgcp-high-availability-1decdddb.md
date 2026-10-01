---
id: collect-261001-fortinet/fortinet/document-fortigate-6-2-0-best-practices-50165-fgcp-high-availability-1decdddb
title: "document-fortigate-6-2-0-best-practices-50165-fgcp-high-availability-1decdddb"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["throughput"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-6-2-0-best-practices-50165-fgcp-high-availability-1decdddb.md
source_anchor: ""
source_lines: [1, 13]
sha256: c581aed83924b1483a48cb216292806326f8370d4a25469fb6abee78f7d597a8
---

# document-fortigate-6-2-0-best-practices-50165-fgcp-high-availability-1decdddb

FGCP high availability
FGCP high availability
Fortinet suggests the following practices related to high availability:
- Use Active-Active HA to distribute TCP and UTM sessions among multiple cluster units. An active-active cluster may have higher throughput than a standalone FortiGate unit or than an active-passive cluster.
- Use a different host name on each FortiGate unit when configuring an HA cluster. Fewer steps are required to add host names to each cluster unit before configuring HA and forming a cluster.
- Consider adding an Alias to the interfaces used for the HA heartbeat so that you always get a reminder about what these interfaces are being used for.
- Enabling load-balance-all can increase device and network load since more traffic is load-balanced. This may be appropriate for use in a deployment using the firewall capabilities of the FortiGate unit and IPS but no other content inspection.
- An advantage of using session pickup is that non-content inspection sessions will be picked up by the new primary unit after a failover. The disadvantage is that the cluster generates more heartbeat traffic to support session pickup as a larger portion of the session table must be synchronized. Session pickup should be configured only when required and is not recommended for use with SOHO FortiGate models. Session pickup should only be used if the primary heartbeat link is dedicated (otherwise the additional HA heartbeat traffic could affect network performance).
- If session pickup is not selected, after a device or link failover all sessions are briefly interrupted and must be re-established at the application level after the cluster renegotiates. For example, after a failover, users browsing the web can just refresh their browsers to resume browsing. Users downloading large files may have to restart their download after a failover. Other protocols may experience data loss and some protocols may require sessions to be manually restarted. For example, a user downloading files with FTP may have to either restart downloads or restart their FTP client.
- If you need to enable session pickup, consider enabling session-pickup-delay to improve performance by reducing the number of sessions that are synchronized.
- Consider using the session-sync-dev option to move session synchronization traffic off the HA heartbeat link to one or more dedicated session synchronization interfaces.
- To avoid unpredictable results, when you connect a switch to multiple redundant or aggregate interfaces in an active-passive cluster you should configure separate redundant or aggregate interfaces on the switch; one for each cluster unit.
- Use SNMP, syslog, or email alerts to monitor a cluster for failover messages. Alert messages about cluster failovers may help find and diagnose network problems quickly and efficiently.

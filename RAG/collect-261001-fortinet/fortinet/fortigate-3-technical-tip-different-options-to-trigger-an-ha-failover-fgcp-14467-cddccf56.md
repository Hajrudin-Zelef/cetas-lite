---
id: collect-261001-fortinet/fortinet/fortigate-3-technical-tip-different-options-to-trigger-an-ha-failover-fgcp-14467-cddccf56
title: "fortigate-3-technical-tip-different-options-to-trigger-an-ha-failover-fgcp-14467-cddccf56"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/fortigate-3-technical-tip-different-options-to-trigger-an-ha-failover-fgcp-14467-cddccf56.md
source_anchor: ""
source_lines: [1, 4]
sha256: 11b4be11f2b6865a59b71778d24d959966175a6592f643aaa393d486a20e23fc
---

# fortigate-3-technical-tip-different-options-to-trigger-an-ha-failover-fgcp-14467-cddccf56

Technical Tip: Different options to trigger an HA failover (FGCP)
| Description | This article describes different options available to trigger an HA failover in the FGCP HA environment to make sure the traffic fails over to the backup FortiGate. | 
| Scope | FortiGate. | 
| Solution | In Fortigate HA, the following options exist to test a fail-over:  Scenario 1: If override is enabled in the HA settings of both units (primary and secondary), set the HA priority of the primary FortiGate to a lower value than the secondary FortiGate's HA priority.  Scenario 2: If override is disabled in the HA settings of both units (primary and secondary), reset the uptime on the primary FortiGate with the following CLI command:  diagnose sys ha reset-uptime Scenario 3: Choose to reboot Primary-FortiGate to force failover regardless of the configuration:  execute reboot This operation will reboot the system ! Do you want to continue? (y/n) y  Scenario 4: Force failover regardless of the configuration and without rebooting the device:  execute ha failover set 1 Caution: This command will trigger an HA failover. It is intended for testing purposes. Do you want to continue? (y/n) y  Note: This command is intended for testing or troubleshooting purposes. Always ensure to unset the failover flag after testing to avoid unintended behavior.  To unset the failover flag:  execute ha failover unset 1 Caution: This command may trigger an HA failover. It is intended for testing purposes. Do you want to continue? (y/n)y  Scenario 5: Port monitoring is enabled in both FortiGates, and one of the ports monitored in the Primary FortiGate changes to down:  get sys ha status .......  MONDEV stats: FG180FTK22901111(updated 2 seconds ago): LACP-AAA: aggregate/00, down, rx-bytes/packets/dropped/errors=238365954599499/424560658809/0/0, tx=266947122763281/432678050210/0/0 LACP-XXX: aggregate/00, up, rx-bytes/packets/dropped/errors=204891583771678/317099587259/0/0, tx=190870014539576/305491834397/0/0  FG180FTK21902222(updated 4 seconds ago): LACP-AAA: aggregate/00, up, rx-bytes/packets/dropped/errors=681689930511/1693891118/0/0, tx=679685895624/1047846512/0/0 LACP-XXX: aggregate/00, up, rx-bytes/packets/dropped/errors=504580436600/810217132/0/0, tx=457930209521/749546395/0/0  Related article: Technical Tip: How to use failover flag to change Active unit |

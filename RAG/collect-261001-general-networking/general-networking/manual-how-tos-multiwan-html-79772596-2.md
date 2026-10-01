---
id: collect-261001-general-networking/general-networking/manual-how-tos-multiwan-html-79772596-2
title: "manual-how-tos-multiwan-html-79772596"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/manual-how-tos-multiwan-html-79772596.md
source_anchor: ""
source_lines: [114, 140]
sha256: ddb3f8d631d1963c417b1d645334614f15024dbf6dde71c19459e99d686dfe7c
---

# manual-how-tos-multiwan-html-79772596

This setup is configured globally via , there cannot be a distinction for different gateway groups.
Configuration
For a minimal working failover configuration, we need two gateways with different priorities.
Go to
Note
We assume both the main and metered gateways already exist due to DHCP configuration.
| Name | WAN_DHCP | 
| Upstream Gateway | X | 
| Failover States | X | 
| Priority | 253 | 
Note
The Priority must be a lower number than the metered ISP gateway. This will mark this gateway as preferred. Checking Failover States will kill all firewall states if a failover happens. This means you must enable gateway monitoring, otherwise there cannot be a failover.
| Name | LTE_DHCP | 
| Upstream Gateway | X | 
| Failback States | X | 
| Priority | 254 | 
Note
The Priority must be a higher number than the main ISP gateway. Checking Failback States will kill all firewall states if our main gateway comes back online.
Go to and enable the following:
| Gateway switching | X | 
This will allow the default gateway of this firewall to change when a failover happens. It is necessary for the failover and failback of states to trigger correctly.
Verification
To verify if the failover and failback kill firewall states as expected, the simplest test is unplugging the main ISP and wait for the gateway monitor to trigger the failover to the metered ISP.
Any client with a session to the internet will be forced to re-establish it. A good test would be a SSH or RDP session.
Afterwards, reconnect the main ISP and wait for the failback to happen. The same scenario with the sessions being forced to re-establish should repeat.
If there are issues, verify default gateway switching, gateway priorities, and if the correct failover and failback states options have been set.
For further diagnostics, use .

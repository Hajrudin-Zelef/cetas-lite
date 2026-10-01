---
id: collect-261001-fortinet/fortinet/fortigate-3-troubleshooting-tip-ha-cluster-not-establishing-when-one-firewall-ha-6d2ede6c
title: "fortigate-3-troubleshooting-tip-ha-cluster-not-establishing-when-one-firewall-ha-6d2ede6c"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-fortinet/fortigate-3-troubleshooting-tip-ha-cluster-not-establishing-when-one-firewall-ha-6d2ede6c.md
source_anchor: ""
source_lines: [1, 43]
sha256: ff53d892f6557f4f61001f3544c8fb5e70a23d4019d6923765f720da472f5dd1
---

# fortigate-3-troubleshooting-tip-ha-cluster-not-establishing-when-one-firewall-ha-6d2ede6c

Troubleshooting Tip: HA cluster not establishing when one firewall has more vdoms than the other
Description
This article describes how to fix an issue where, when one of the devices in an HA cluster has additional VDOMs other than the standard 10 VDOMs per firewall, the HA cluster does not form.
Scope
FortiGate-3000 or higher (FortiGate-1240B supports 25 VDOMs).v3.0 and higher.
Solution
Once the secondary is added in the cluster, if the cluster is not forming, run the commands below:
diagnose debug application hatalk -1
diagnose debug enable
To disable the debug, use the following:
diagnose debug disable
diagnose debug reset
The following message in the debug indicates that the number of VDOMs is different on each unit and indicates the number allowed. In this case, the primary has 3 VDOMs while the secondary is only allowed 2.
<hatalk> HA cannot be formed because this box has 3 vdoms. It exceeds the maximum number of vdoms allowed on the HA peer 'FGVMXXXXXXXXXXXX', which only allows maximum 2 vdoms.
<hatalk> parse options for ' FGVMXXXXXXXXXXXX', packet_version=5
<hatalk> HA cannot be formed because this box has 3 vdoms. It exceeds the maximum number of vdoms allowed on the HA peer 'FGVMXXXXXXXXXXXX', which only allows maximum 2 vdoms.
<hatalk> parse options for ' FGVMXXXXXXXXXXXX', packet_version=5
To find the number of VDOMs allowed on FortiGate, use the following CLI command:
get system status | grep “virtual domains”
The output should show the max allowed virtual domains:
Max number of virtual domains: 10
To check the allowed VDOMs on the secondary, use the following command to manage the secondary:
execute ha manage [ID] [admin_username] <---- The ID will usually be 0 or 1
Once on the secondary, use the 'get system status | grep “virtual domains”' command to check the maximum number of allowed virtual domains on the secondary.
By default, most FortiGate units support 10 VDOMs, and many FortiGate models support purchasing a license key to increase the maximum number. Some exceptions may apply. Each FortiGate in an HA cluster requires the same number of VDOMs for the cluster to form correctly.
Purchase a license key for the secondary firewall.
To obtain a VDOM license key:
Record the FortiGate serial number. It is possible to find the serial number in the web-based manager on the System Status page.
Login on the Fortinet Support website, and use the serial number of the registered FortiGate to purchase a license key for 25, 50, 100, 250, or 500 VDOMs.
- When the license key is received, go to System -> Maintenance -> License on the FortiGate.
- In the License Key field, enter the 32-character license key received from Fortinet.
- Select Apply.
From CLI:
The license can be applied using the following command:
config global
execute upd-vd-license <license key>
In an HA environment, the license needs to be applied to each unit. 
For more information about VDOMs, see Virtual Domains.
For more information about HA, see High Availability.
Related articles:
Troubleshooting Tip: How to troubleshoot HA synchronization issue using GUI
Technical Tip: Procedure for HA manual synchronization
Technical Tip: How to access secondary unit of HA cluster via CLI

---
id: collect-261001-fortinet/fortinet/fortigate-3-troubleshooting-tip-unable-to-upgrade-downgrade-firmware-due-to-inva-11e9c35c
title: "fortigate-3-troubleshooting-tip-unable-to-upgrade-downgrade-firmware-due-to-inva-11e9c35c"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["license", "licenses"]
source: docs/RAG/collect-261001-fortinet/fortigate-3-troubleshooting-tip-unable-to-upgrade-downgrade-firmware-due-to-inva-11e9c35c.md
source_anchor: ""
source_lines: [1, 4]
sha256: 400b3d49f100d920db212bd511f2aa3c7519e475e775fe5ec15c4f5e220f7f47
---

# fortigate-3-troubleshooting-tip-unable-to-upgrade-downgrade-firmware-due-to-inva-11e9c35c

Troubleshooting Tip: Unable to upgrade/downgrade firmware due to invalid license status even though FortiGate has a valid license
| Description | This article describes how to resolve an issue where it is not possible to upgrade/downgrade firmware, and the error message 'Unable to upgrade/downgrade firmware through FortiGuard because of invalid license status' is thrown despite the FortiGate having a valid license.    This issue happens when the FortiGate is in HA, and one of the HA members does not have a valid license. | 
| Scope | FortiGate. | 
| Solution | It is required that all HA members have valid licenses. If HA members have different license levels or expiration dates, all HA members will be downgraded to the lowest license.  To resolve this issue, make sure all HA members have valid licenses. To verify if all the FortiGate cluster members are properly licensed, the following debug command can be run: In an HA cluster with 2 members, 2 contracts are expected. If only 1 device is being licensed, the following error will be observed when the above command is enabled:  Another instance in which this issue may occur is if the HA cluster members are registered in two different accounts. Ensure that all the FortiGate HA cluster members are under the same FortiCloud account.  Starting from FortiOS v7.2.9, v7.4.6, v7.6.1 and above, FortiGate supports 'Single FortiGuard license for FortiGate A-P HA cluster'. It is possible to purchase a specific FortiGate SKU that can be applied to both HA members. For more information, see this article: Technical Tip: Additional Info regarding Single FortiGuard license for FortiGate A-P HA cluster feature. Note: If both FortiGates in an HA cluster are confirmed to have the same level of licensing but are still getting the same error, force communication to FortiGuard to have licenses re-checked by running the command execute update-now.  Related article: Technical Tip: The HA Cluster requirements |

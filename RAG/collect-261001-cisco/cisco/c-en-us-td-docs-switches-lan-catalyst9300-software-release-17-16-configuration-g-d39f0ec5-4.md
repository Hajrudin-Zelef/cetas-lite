---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-d39f0ec5-4
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-d39f0ec5"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-d39f0ec5.md
source_anchor: ""
source_lines: [369, 448]
sha256: 36d177b1c7a5aec2d9ce6f61fe6037aa5024a8ef7ae6510163de1319d18fbc24
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-d39f0ec5

If the maximum prefix limit at process level and neighbor level is set together then the max prefix limit at process level
will take precedence. When the max prefix limit at neighbor level is set greater than the max prefix limit set at process
level, the device displays this message:
Max prefix limit at neighbor level is set to a value (%d) greater than max prefix limit at process level (%d)
Example Configuring the Maximum-Prefix Limit for All Peers--Named Configuration
The following example, starting in global configuration mode, configures the maximum-prefix limit for all peers. The maximum
limit is set to 10,000 prefixes, the warning threshold is set to 90 percent, the restart timer is set to 4 minutes, a decay
penalty is configured for the restart timer with the dampenedkeyword, and all timers are configured to be reset to 0 every 60 minutes. When the maximum-prefix limit is exceeded, all peering
sessions will be torn down, all routes learned from all peers will be removed from the topology and routing tables, and all
peers will be placed in a penalty state for 4 minutes (user-defined penalty value). A dampening exponential decay penalty
will also be applied.
If the maximum prefix limit at process level and neighbor level is set together then the max prefix limit at process level
will take precedence. When the max prefix limit at neighbor level is set greater than the max prefix limit set at process
level, the device displays this message:
Max prefix limit at neighbor level is set to a value (%d) greater than max prefix limit at process level (%d)
Example Configuring the Maximum-Prefix Limit for Redistributed Routes--Autonomous System Configuration
The following example, starting in global configuration mode, configures the maximum-prefix limit for routes learned through
redistribution. The maximum limit is set to 5000 prefixes and the warning threshold is set to 95 percent. When the number
of prefixes learned through redistribution reaches 4750 (95 percent of 5000), warning messages will be displayed in the console.
Because the warning-only keyword is configured, the topology and routing tables will not be cleared and route redistribution will not be placed in
a penalty state.
When the maximum prefix limit is configured at both the process level and redistribution level, the limit set at the process
level will take precedence. In cases where the max prefix limit at redistribution level is set greater than the max prefix
limit set at process level, the device displays this message:
Max prefix limit at redistribute level is set to a value (%d) greater than max prefix limit at process level (%d)
Example Configuring the Maximum-Prefix Limit for Redistributed Routes--Named Configuration
The following example, starting in global configuration mode, configures the maximum-prefix limit for routes learned through
redistribution. The maximum limit is set to 5000 prefixes and the warning threshold is set to 95 percent. When the number
of prefixes learned through redistribution reaches 4750 (95 percent of 5000), warning messages will be displayed in the console.
Because the warning-only keyword is configured, the topology and routing tables will not be cleared and route redistribution will not be placed in
a penalty state.
When the maximum prefix limit is configured at both the process level and redistribution level, the limit set at the process
level will take precedence. In cases where the max prefix limit at redistribution level is set greater than the max prefix
limit set at process level, the device displays this message:
Max prefix limit at redistribute level is set to a value (%d) greater than max prefix limit at process level (%d)
Example Configuring the Maximum-Prefix Limit for an EIGRP Process--Autonomous System Configuration
The following example, starting in global configuration mode, configures the maximum-prefix limit for an EIGRP process, which
includes routes learned through redistribution and routes learned through EIGRP peering sessions. The maximum limit is set
to 50,000 prefixes. When the number of prefixes learned through redistribution reaches 37,500 (75 percent of 50,000), warning
messages will be displayed in the console.
When the maximum-prefix limit is exceeded, all peering sessions will be reset, the topology and routing tables will be cleared,
and redistributed routes and all peering sessions will be placed in a penalty state.
Device(config)# router eigrp 100
Device(config-router)# address-family ipv4 vrf RED
Device(config-router-af)# maximum-prefix 50000
Device(config-router-af)# end
Note
When the max prefix limit at process level is set lower than the max prefix limit set at redistribute level, the device displays
this message:
Max prefix limit at process level is set to a value (%d) lower than max prefix limit at redistribute level (%d)
When the max prefix limit at process level is set lower than the max prefix limit set at neighbor level, the device displays
this message:
(%d) lower than max prefix limit at neighbor level (%d)
Example Configuring the Maximum-Prefix Limit for an EIGRP Process--Named Configuration
The following example, starting in global configuration mode, configures the maximum-prefix limit for an EIGRP process, which
includes routes learned through redistribution and routes learned through EIGRP peering sessions. The maximum limit is set
to 50,000 prefixes. When the number of prefixes learned through redistribution reaches 37,500 (75 percent of 50,000), warning
messages will be displayed in the console.
When the maximum-prefix limit is exceeded, all peering sessions will be reset, the topology and routing tables will be cleared,
and redistributed routes and all peering sessions will be placed in a penalty state.
When the max prefix limit at process level is set lower than the max prefix limit set at redistribute level, the device displays
this message:
Max prefix limit at process level is set to a value (%d) lower than max prefix limit at redistribute level (%d)
When the max prefix limit at process level is set lower than the max prefix limit set at neighbor level, the device displays
this message:
(%d) lower than max prefix limit at neighbor level (%d)
Feature History for Configuring EIGRP Prefix Limit Support
This table provides release and related information for the features explained in this module. These features are available
in all the releases subsequent to the one they were introduced in, unless noted otherwise.
Table 1. Feature History for Configuring EIGRP Prefix Limit Support
Release
Feature
Feature Information
Cisco IOS XE Everest 16.5.1a
Configuring EIGRP Prefix Limit Support
The EIGRP Prefix Limit Support feature introduces the capability to limit the number of prefixes per VRF that are accepted
from a specific peer or to limit all prefixes that are accepted by an Enhanced Interior Gateway Routing Protocol (EIGRP) process
through peering and redistribution.
Use the Cisco Feature Navigator to find information about platform and software image support. To access Cisco Feature Navigator,
go to https://cfnng.cisco.com/.

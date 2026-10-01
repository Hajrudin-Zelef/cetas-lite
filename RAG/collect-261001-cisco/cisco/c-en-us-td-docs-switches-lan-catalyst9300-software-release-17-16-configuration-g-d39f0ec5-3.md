---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-d39f0ec5-3
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-d39f0ec5"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-d39f0ec5.md
source_anchor: ""
source_lines: [228, 368]
sha256: ee6650be1367e6357dfdf2147ad51d920ceeedb458014a3c54bee8c58ead6d6f
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-d39f0ec5

configureterminal
Example:
Device# configure terminal
Enters global configuration mode.
Step 3
routereigrpvirtual-instance-name
Example:
Device(config)# router eigrp virtual-name1
Enters router configuration mode and creates an EIGRP routing process.
A maximum of 30 EIGRP routing processes can be configured.
Configuring the Maximum-Prefix Limit for an EIGRP Process Autonomous System Configuration
The maximum-prefix limit can be configured for an EIGRP process to limit the number prefixes that are accepted from all sources.
This task is configured with the
maximum-prefixcommand. When the maximum-prefix limit is exceeded, sessions with the remote peers are brought down and all routes learned
from remote peers are removed from the topology and routing tables. Also, all routes learned from the RIB are discarded and
redistribution is suspended for the default or user-defined time period.
Default or user-defined restart, restart-count, and reset-time values for the process-level configuration of this feature,
configured with the
maximum-prefix command, are inherited by the
redistributemaximum-prefix and
neighbormaximum-prefix command configurations by default. If a single peer is configured with the
neighbormaximum-prefix command, a process-level configuration or a configuration that is applied to all neighbors will be inherited.
Note
VRFs have been created and configured.
EIGRP peering is established through the MPLS VPN.
This task can be configured only in IPv4 VRF address family configuration mode.
Procedure
Command or Action
Purpose
Step 1
enable
Example:
Device> enable
Enables privileged EXEC mode.
Enter your password if prompted.
Step 2
configureterminal
Example:
Device# configure terminal
Enters global configuration mode.
Step 3
routereigrpas-number
Example:
Device(config)# router eigrp 1
Enters router configuration mode and creates an EIGRP routing process.
A maximum of 30 EIGRP routing processes can be configured.
Step 4
address-familyipv4 [unicast]
vrfvrf-name
Example:
Device(config-router)# address-family ipv4 vrf RED
Enters address family configuration mode and creates a session for the VRF.
Limits the number of prefixes that are accepted under an address family by an EIGRP process.
The example configures a maximum-prefix limit of 10,000 prefixes, a reset time period of 10 minutes, a warning message to
be displayed at 80 percent of the maximum-prefix limit, and a restart time period of 2 minutes.
Step 6
end
Example:
Device(config-router-af)# end
Exits address-family configuration mode and enters privileged EXEC mode.
Configuring the Maximum-Prefix Limit for an EIGRP Process Named Configuration
The maximum-prefix limit can be configured for an EIGRP process to limit the number prefixes that are accepted from all sources.
This task is configured with the
maximum-prefixcommand. When the maximum-prefix limit is exceeded, sessions with the remote peers are brought down and all routes learned
from remote peers are removed from the topology and routing tables. Also, all routes learned from the RIB are discarded and
redistribution is suspended for the default or user-defined time period.
Default or user-defined restart, restart-count, and reset-time values for the process-level configuration of this feature,
configured with the
maximum-prefix command, are inherited by the
redistributemaximum-prefix and
neighbormaximum-prefix command configurations by default. If a single peer is configured with the
neighbormaximum-prefix command, a process-level configuration or a configuration that is applied to all neighbors will be inherited.
Note
VRFs have been created and configured.
EIGRP peering is established through the MPLS VPN.
This task can be configured only in IPv4 VRF address family configuration mode.
Procedure
Command or Action
Purpose
Step 1
enable
Example:
Device> enable
Enables privileged EXEC mode.
Enter your password if prompted.
Step 2
configureterminal
Example:
Device# configure terminal
Enters global configuration mode.
Step 3
routereigrpvirtual-instance-name
Example:
Device(config)# router eigrp virtual-name1
Creates an EIGRP routing process and enters router configuration mode.
A maximum of 30 EIGRP routing processes can be configured.
Limits the number of prefixes that are accepted under an address family by an EIGRP process.
The example configures a maximum-prefix limit of 10,000 prefixes, a reset time period of 10 minutes, a warning message to
be displayed at 80 percent of the maximum-prefix limit, and a restart time period of 2 minutes.
Device# show eigrp address-family ipv4 22 accounting
(Optional) Displays prefix accounting information for EIGRP processes.
Note
Connected and summary routes are not listed individually in the output from this
show command but are counted in the total aggregate count per process.
Example
The following is sample output from the
showeigrpaddress-familyaccounting command:
Device# show eigrp address-family ipv4 22 accounting
EIGRP-IPv4 VR(saf) Accounting for AS(22)/ID(10.0.0.1)
Total Prefix Count: 3 States: A-Adjacency, P-Pending, D-Down
State Address/Source Interface Prefix Restart Restart/
Count Count Reset(s)
A 10.0.0.2 Et0/0 2 0 0
P 10.0.2.4 Se2/0 0 2 114
D 10.0.1.3 Et0/0 0 3 0
Configuration Examples for Configuring the Maximum-Prefix Limit
Example Configuring the Maximum-Prefix Limit for a Single Peer--Autonomous System Configuration
The following example, starting in global configuration mode, configures the maximum-prefix limit for a single peer. The maximum
limit is set to 1000 prefixes, and the warning threshold is set to 80 percent. When the maximum-prefix limit is exceeded,
the session with this peer will be torn down, all routes learned from this peer will be removed from the topology and routing
tables, and this peer will be placed in a penalty state for 5 minutes (default penalty value).
If the maximum prefix limit at process level and neighbor level is set together then the max prefix limit at process level
will take precedence. When the max prefix limit at neighbor level is set greater than the max prefix limit set at process
level, the device displays this message:
Max prefix limit at neighbor level is set to a value (%d) greater than max prefix limit at process level (%d)
Example Configuring the Maximum-Prefix Limit for a Single Peer--Named Configuration
The following example, starting in global configuration mode, configures the maximum-prefix limit for a single peer. The maximum
limit is set to 1000 prefixes, and the warning threshold is set to 80 percent. When the maximum-prefix limit is exceeded,
the session with this peer will be torn down, all routes learned from this peer will be removed from the topology and routing
tables, and this peer will be placed in a penalty state for 5 minutes (default penalty value).
If the maximum prefix limit at process level and neighbor level is set together then the max prefix limit at process level
will take precedence.When the max prefix limit at neighbor level is set greater than the max prefix limit set at process level,
the device displays this message:
Max prefix limit at neighbor level is set to a value (%d) greater than max prefix limit at process level (%d)
Example Configuring the Maximum-Prefix Limit for All Peers--Autonomous System Configuration
The following example, starting in global configuration mode, configures the maximum-prefix limit for all peers. The maximum
limit is set to 10,000 prefixes, the warning threshold is set to 90 percent, the restart timer is set to 4 minutes, a decay
penalty is configured for the restart timer with the dampenedkeyword, and all timers are configured to be reset to 0 every 60 minutes. When the maximum-prefix limit is exceeded, all peering
sessions will be torn down, all routes learned from all peers will be removed from the topology and routing tables, and all
peers will be placed in a penalty state for 4 minutes (user-defined penalty value). A dampening exponential decay penalty
will also be applied.

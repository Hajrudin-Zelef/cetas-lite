---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-d39f0ec5-2
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-d39f0ec5"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-d39f0ec5.md
source_anchor: ""
source_lines: [77, 227]
sha256: e68ec95e7333f6f64716f00cf126b0b768af4847c36585f83961868d36f41913
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-d39f0ec5

If the EIGRP process enters into a suspended (pending or down) state, the device will not establish neighborships with new
peers and thus cease to transmit and stop processing hello packets.
Configuring the Maximum Number of Prefix Accepted from Peering Sessions Autonomous System Configuration
The maximum-prefix limit can be configured for all peering sessions or individual peering sessions with the neighbormaximum-prefix(EIGRP) command. When the maximum-prefix limit is exceeded, the session with the remote peer is torn down and all routes learned
from the remote peer are removed from the topology and routing tables. The maximum-prefix limit that can be configured is
limited only by the available system resources on the device.
Note
In EIGRP,
neighbor commands have been used traditionally to configure static neighbors. In the context of this feature, however, the
neighbormaximum-prefix command can be used to configure the maximum-prefix limit for both statically configured and dynamically discovered neighbors.
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
When you configure the neighbormaximum-prefix command to protect a single peering session, only the maximum-prefix limit, the percentage threshold, the warning-only configuration
options can be configured. Session dampening, restart, and reset timers are configured on a global basis.
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
Limits the number of prefixes that are accepted from all EIGRP neighbors.
Step 8
end
Example:
Device(config-router-af)# end
Exits address family configuration mode and enters privileged EXEC mode.
Configuring the Maximum Number of Prefixes Accepted from Peering Sessions Named Configuration
The maximum-prefix limit can be configured for all peering sessions or individual peering sessions with the neighbormaximum-prefix(EIGRP) command. When the maximum-prefix limit is exceeded, the session with the remote peer is torn down and all routes learned
from the remote peer are removed from the topology and routing tables. The maximum-prefix limit that can be configured is
limited only by the available system resources on the device.
Note
In EIGRP, neighbor commands have been used traditionally to configure static neighbors. In the context of this feature, however, the neighbormaximum-prefix command can be used to configure the maximum-prefix limit for both statically configured and dynamically discovered neighbors.
Default or user-defined restart, restart-count, and reset-time values for the process-level configuration of this feature,
configured with the maximum-prefix command, are inherited by the redistributemaximum-prefix and neighbormaximum-prefix command configurations by default. If a single peer is configured with the neighbormaximum-prefix command, a process-level configuration or a configuration that is applied to all neighbors will be inherited.
Note
VRFs have been created and configured.
EIGRP peering is established through the MPLS VPN.
This task can be configured only in IPv4 VRF address family configuration mode.
When you configure the neighbormaximum-prefix command to protect a single peering session, only the maximum-prefix limit, the percentage threshold, the warning-only configuration
options can be configured. Session dampening, restart, and reset timers are configured on a global basis.
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
Enters router configuration mode and creates an EIGRP routing process.
A maximum of 30 EIGRP routing processes can be configured.
Limits the number of prefixes that are accepted from all EIGRP neighbors.
Step 8
exit-address-family
Example:
Device(config-router-af)# exit-address-family
Exits address family configuration mode.
Configuring the Maximum Number of Prefixes Learned Through Redistribution Autonomous System Configuration
The maximum-prefix limit can be configured for prefixes learned through redistribution with the redistributemaximum-prefix (EIGRP) command. When the maximum-prefix limit is exceeded, all routes learned from the RIB will be discarded and redistribution
will be suspended for the default or user-defined time period. The maximum-prefix limit that can be configured for redistributed
prefixes is limited only by the available system resources on the device.
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
Limits the number of prefixes redistributed into an EIGRP process.
Step 6
end
Example:
Device(config-router-af)# end
Exits address family configuration mode and enters privileged EXEC mode.
Configuring the Maximum Number of Prefixes Learned Through Redistribution Named Configuration
The maximum-prefix limit can be configured for prefixes learned through redistribution with the redistributemaximum-prefix(EIGRP) command. When the maximum-prefix limit is exceeded, all routes learned from the RIB will be discarded and redistribution
will be suspended for the default or user-defined time period. The maximum-prefix limit that can be configured for redistributed
prefixes is limited only by the available system resources on the device.
Default or user-defined restart, restart-count, and reset-time values for the process-level configuration of this feature,
configured with the
maximum-prefix command, are inherited by the
redistributemaximum-prefix and
neighbormaximum-prefix command configurations by default. If a single peer is configured with the
neighbormaximum-prefix command, a process-level configuration or a configuration that is applied to all neighbors will be inherited.
Note
VRFs have been created and configured.
EIGRP peering is established through the MPLS VPN.
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

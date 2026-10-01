---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate-ba9ee001-10
title: "docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate--ba9ee001"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["copyright", "warrants"]
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate--ba9ee001.md
source_anchor: ""
source_lines: [1589, 1695]
sha256: 3d9ad00776b1c4356b9a5d2af5461b13ce4b21320675c756927f6d1a285ce7b3
---

# docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate--ba9ee001

Using FortiManager and FortiAnalyzer
FortiManager and FortiAnalyzer are supported similarly to NAT/Route mode. For more information about this
please consult the Fortinet documentation at http://docs.fortinet.com or the Knowledge base at
http://kb.fortinet.com .
Establishing a communication to a FortiAnalyzer is done as per the example hereafter (from global level if VDOM
is enabled). This setting is independent from being in Transparent mode. However, as stated earlier in this
section the management VDOM must have IP connectivity to the FortiAnalyzer.
FGT (global) # show system fortianalyzer
config system fortianalyzer
set status enable
set server 10.2.2.2
end
Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.
45

High Availability in Transparent Mode
This section contains information about configuring high availability in Transparent mode. It contains the
following topics:
l Virtual Clustering
l MAC Address Assignment
For complete information about HA, please refer to the Fortigate Administration Guide
or the HA Technical guides available at http://docs.fortinet.com or the Knowledge
Base at http://kb.fortinet.com .
Any other statement and feature description in this document apply to a FortiGate Cluster running in Active-
Passive mode.
Virtual Clustering
If VDOM (virtual domain) is enabled on a cluster operating Transparent Mode, HA Virtual Clustering can be
configured in active-passive mode.
This will provide:
l Failover protection between two instances of a VDOM operating on two different FortiGate in the cluster.
l Load balancing between the FortiGate units on a per-VDOM basis.
The roles have been defined such as, in normal operation:
l FortiGate1 is Master for Vdom1 and Slave for Vdom2
l FortiGate2 is Master for Vdom2 and Slave for Vdom1
In case of a failure or reboot of a FortiGate, the remaining unit will become Master for Vdom1 and Vdom2.
The VDOMs given in this example are showing physical ports but a VDOM can also
include VLAN interfaces.
The L2 connectivity between the FortiGate is showing 4 separate L2 switches, but it
could also be one single switch one each side configured with appropriate VLANs.
Configuration example
l FortiGate1:
FGT1 (global) # show system ha
config system ha
set mode a-p
46 Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.

High Availability in Transparent Mode MAC Address Assignment
set hbdev "port5" 0 "port6" 0
set vcluster2 enable
set override disable
set priority 200
config secondary-vcluster
set override enable
set priority 100
set vdom "Vdom2"
end
end
l FortiGate2:
FGT2 (global) # show system ha
config system ha
set mode a-p
set hbdev "port5" 0 "port6" 0
set vcluster2 enable
set override disable
set priority 200
config secondary-vcluster
set override enable
set priority 100
set vdom "Vdom2"
end
end
MAC Address Assignment
If a cluster is operating in Transparent mode, the FortiGate Clustering Protocol (FGCP) assigns a virtual MAC
address for the Master unit management IP address. Since you can connect to the management IP address from
any interface, all of the FortiGate interfaces appear to have the same virtual MAC address.
Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.
47

Best Practices
1. Create forwarding domains when VLANs are used and set vlanforward todisable on all relevant physical
interface.
2. The forward-domain ID can be different to the VLAN ID, but it is recommended for troubleshooting and readability
to keep them the same.
3. Only interfaces from the same forwarding domains can have firewall policies between each others.
4. In order to allow IVL (independent VLAN learning), the VLANs must be placed in separate forwarding domains.
5. If an out-of-band management is desired, use if possible a VDOM in NAT/Route mode as management VDOM
and create (an) other Transparent mode VDOM(s) for the user traffic.
6. As Spanning Tree BPDUs are not forwarded by default, insert the FortiGate with caution to avoid L2 loops.
7. Multicast packets are not forwarded by default; this might cause routing protocols (RIP2,OSPF) disruption.
8. When using HSRP or VRRP configure static MAC entries for the Virtual MAC addresses.
48 Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.

Copyright© 2018 Fortinet, Inc. All rights reserved. Fortinet®, FortiGate®, FortiCare® and FortiGuard®, and certain other marks are registered trademarks of Fortinet,
Inc., in the U.S. and other jurisdictions, and other Fortinet names herein may also be registered and/or common law trademarks of Fortinet. All other product or company
names may be trademarks of their respective owners. Performance and other metrics contained herein were attained in internal lab tests under ideal conditions, and
actual performance and other results may vary. Network variables, different network environments and other conditions may affect performance results. Nothing herein
represents any binding commitment by Fortinet, and Fortinet disclaims all warranties, whether express or implied, except to the extent Fortinet enters a binding written
contract, signed by Fortinet’s General Counsel, with a purchaser that expressly warrants that the identified product will perform according to certain expressly-identified
performance metrics and, in such event, only the specific performance metrics expressly identified in such binding written contract shall be binding on Fortinet. For
absolute clarity, any such warranty will be limited to performance in the same ideal conditions as in Fortinet’s internal lab tests. In no event does Fortinet make any
commitment related to future deliverables, features, or development, and circumstances may change such that any forward-looking statements herein are not accurate.
Fortinet disclaims in full any covenants, representations,and guarantees pursuant hereto, whether express or implied. Fortinet reserves the right to change, modify,
transfer, or otherwise revise this publication without notice, and the most current version of the publication shall be applicable.

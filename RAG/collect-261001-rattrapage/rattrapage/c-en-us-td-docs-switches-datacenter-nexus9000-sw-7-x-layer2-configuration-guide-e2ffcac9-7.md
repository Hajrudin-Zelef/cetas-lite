---
id: collect-261001-rattrapage/rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-layer2-configuration-guide-e2ffcac9-7
title: "c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-layer2-configuration-guide--e2ffcac9"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-layer2-configuration-guide--e2ffcac9.md
source_anchor: ""
source_lines: [455, 461]
sha256: 6624089092db6b32f7a9657ee8b096372f123d92012d81171f21d6ee4669ac93
---

# c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-layer2-configuration-guide--e2ffcac9

                                    Protocol migration—For backward compatibility with 802.1D devices, 802.1w selectively sends 802.1D configuration BPDUs and TCN BPDUs on a per-port basis. When a port is initialized, the migrate-delay timer is started (specifies the minimum time during which 802.1w BPDUs are sent), and 802.1w BPDUs are sent. While this timer is active, the device processes all BPDUs received on that port and ignores the protocol type. If the device receives an 802.1D BPDU after the port migration-delay timer has expired, it assumes that it is connected to an 802.1D device and starts using only 802.1D BPDUs. However, if the 802.1w device is using 802.1D BPDUs on a port and receives an 802.1w BPDU after the timer has expired, it restarts the timer and starts using 802.1w BPDUs on that port.
| Note | If you want all devices on the same LAN segment to reinitialize the protocol on each interface, you must reinitialize Rapid PVST+. | 
Rapid PVST+ Interoperation with 802.1s MST
Rapid PVST+ interoperates seamlessly with the IEEE 802.1s Multiple Spanning Tree (MST) standard. No user configuration is needed. To disable this seamless interoperation, you can use PVST Simulation.
High Availability for Rapid PVST+
The software supports high availability for Rapid PVST+. However, the statistics and timers are not restored when Rapid PVST+ restarts. The timers start again and the statistics begin from 0.
| Note | See the Cisco Nexus 9000 Series NX-OS High Availability and Redundancy Guide, for complete information on high-availability features. |

---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-83-7
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "latency", "memory"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-83.md
source_anchor: ""
source_lines: [444, 499]
sha256: 83e0ec4e10ab990b1443c1e932040ed03c034ed7fb1397f01a22f902b2bd57a2
---

# Overview

| Property | Description | 
|---|---|
| **mirror-buffers** (*percent**: 1..90* ; Default:**auto** ) | Maximum amount of packet buffers for mirrored traffic (% of the total buffer memory). | 
| **mirror-profile** (*name* ; Default:**default** ) | The name of the QoS profile to assign to the mirrored packets (see **`/in/eth/sw/qos/profile`** ). | 
| **multicast-buffers** (*percent**: 1..90* ; Default:**auto** ) | Maximum amount of packet buffers for multicast/broadcast traffic (% of the total buffer memory). | 
| **shared-buffers** (*percent**: 0..90* ; Default:**auto** ) | Maximum amount of packet buffers that are shared between ports (% of the total buffer memory). Setting it to 0 disables buffer sharing. The remaining buffer memory is split between the ports. All buffers are treated as shared on new generation switch chips that use **Dynamic Buffers** . The switch chip automatically adjusts port and queue buffer limits based on the current congestion level. Using the auto value allows the device to utilize 100% of the available buffer memory, which is the recommended setting for most scenarios. In specific use cases where latency is more important than avoiding packet drops, the buffer limit can be manually reduced. | 
| **lossless-buffers**  (*percent**: 0..100* ; Default:**auto** ) | If the device supports multiple shared buffer pools, this setting allows adjusting the size of the lossless pool (% of the *shared* buffer memory, where 100% means all shared buffers allocated by the**shared-buffers** setting). For example, if shared-buffers=50 and lossless-buffers=80, the lossless pool receives 40% of the total buffer memory (80% of 50% or "0.8 * 0.5 = 0.4"), and the lossy pool receives the remaining 10% of shared buffers. | 
| **lossless-traffic-class**  (*integer array: 0..7* ; Default:**auto** ) | The list of lossless traffic classes. | 
| **wred-threshold**  (*low \| medium \| high;*  Default:**medium** )  | A relative amount of packets above a shared queue cap (" `queueX-shared-packet-cap` " or "`queueX-shared-byte-cap` ") where random drops take place. This threshold is applied only to queues with enabled Weighed Random Early Detection (**wred=yes** ) that use shared buffers (**use-shared-buffers=yes)** . The higher the queue buffer fill level, the higher the packet drop chance. The*low* threshold means the random tail drop starts later; the*high* - sooner. | 

## QoS Monitor

**Command:** `/interface/ethernet/switch/qos/monitor`

Monitors hardware QoS resources.

| Property | Description | 
|---|---|
| **total-packet-cap***(integer)* | Total packet capacity. The maximum number of hardware packet descriptors that the device can store is all queues. | 
| **total-packet-use***(integer)* | Total packet usage. The current number of packet descriptors residing in the hardware memory. | 
| **total-byte-cap***(byte)* | Total tx memory capacity. | 
| **total-byte-use***(byte)* | Total tx memory usage. The current number of bytes occupied by the packets in all tx queues. | 
| **multicast-packet-cap** *(integer)* | Multicast packet capacity. The maximum number of hardware packet descriptors that can be used by multicast/broadcast traffic. Depends on the **multicast-buffers** setting. | 
| **multicast-packet-use** *(integer)* | Multicast packet usage. The hardware makes a copy of the packet descriptor for each multicast destination. | 
| **mirror-ingress-packet-cap** *(integer)* | Ingress mirror packet capacity. The maximum number of hardware packet descriptors that can be used by ingress mirrored traffic. Depends on the **mirror-buffers** setting. | 
| **mirror-ingress-packet-use** *(integer)* | Ingress mirror packet usage. | 
| **mirror-ingress-byte-cap***(byte)* | Ingress mirror byte capacity. Depends on the **mirror-buffers** setting. | 
| **mirror-ingress-byte-use***(byte)* | Ingress mirror byte usage. | 
| **mirror-egress-packet-cap** *(integer)* | Egress mirror packet capacity. The maximum number of hardware packet descriptors that can be used by egress mirrored traffic. Depends on the **mirror-buffers** setting. | 
| **mirror-egress-packet-use** *(integer)* | Egress mirror packet usage. | 
| **mirror-egress-byte-cap***(byte)* | Egress mirror byte capacity. Depends on the **mirror-buffers** setting. | 
| **mirror-egress-byte-use***(byte)* | Egress mirror byte usage. | 
| **shared-packet-cap***(integer)* | Shared packet capacity. The maximum number of hardware packet descriptors that can be shared between ports and tx queues. Depends on the **shared-buffers** setting. | 
| **shared-packet-use***(integer)* | Shared packet usage. The current number of shared packet descriptors used by all tx queues. | 
| **shared-byte-cap***(byte)* | Shared tx memory capacity. Depends on the **shared-buffers** setting. | 
| **shared-byte-use***(byte)* | Shared tx memory usage. The current number of shared buffers occupied by the packets in all tx queues. | 
| **lossy-pool-packet-cap**  (integer) | Shared packet capacity of the lossy pool. The field is omitted if the device does not support multiple shared pools. | 
| **lossy-pool-packet-use (integer)** | Shared packet usage of the lossy pool. The field is omitted if the device does not support multiple shared pools. | 
| **lossless-pool-packet-cap (integer)** | Shared packet capacity of the lossless pool. The field is omitted if the device does not support multiple shared pools. | 
| **lossless-pool-packet-use (integer)** | Shared packet usage of the lossless pool. The field is omitted if the device does not support multiple shared pools. | 
| **wred-packet-cap***(integer)* | The maximum packet count that a queue can use above the shared cap (" `queueX-shared-packet-cap` " in "`/in/eth/sw/qos/port print usage` ") to trigger a random tail drop. For example, if "`queue1-shared-packet-cap=3072` " and "`wred-packet-cap=512` ", WRED triggers when`queue1-packet-use` exceeds 3072, reaching 100% drop rate at 3072+512=3584 packets. | 
| **wred-byte-cap***(integer)* | The maximum byte count that a queue can use above the shared cap (" `queueX-shared-byte-cap` ") to trigger a random tail drop. For example, if "`queue1-shared-byte-cap=768KiB` " and "`wred-byte-cap=128KiB` ", WRED triggers when`queue1-packet-use` exceeds 768KiB, reaching 100% drop rate at 768+128=896KiB. | 

**Sub-menu:** `/interface/ethernet/switch/qos/profile`

QoS profiles determine priority field values (PCP, DSCP) for the forwarded/routed packets. Congestion avoidance/resolution is based on QoS profiles. Each packet gets a QoS profile assigned based on the ingress switch port QoS settings (see `/in/eth/sw/port`).

| Property | Description | 
|---|---|
| **color** (*green \| yellow \| red* ; Default:**green** ) | Traffic color for color-aware drop precedence management. Leave the default value (green) for color-blind drop precedence management. | 
| **dscp** (*integer: 0..63* ; Default:**0** ) | IPv4/IPv6 DSCP field value for the egress packets assigned to the QoS profile. | 
| **name** (*string* ; Default: ) | The user-defined name of the QoS profile. | 
| **pcp** (*integer: 0..7* ; Default:**0** ) | VLAN priority value (IEEE 802.1q PCP - Priority Code Point). Used only if the egress packets assigned to the QoS profile are VLAN-tagged (have the 802.1q header). The value can be further altered via the QoS Egress Map. | 
| **traffic-class** (*integer: 0..7* ; Default:**0** ) | The traffic class determines the packet priority and the egress queue (see **tx-manager** ). The queue number is usually the same as the traffic class (packets with tc0 go into queue0, tc1 - queue1, ... tc7 - queue7). Unlike pcp, where 0 means the default priority but 1 - the lowest one (and further customizable), traffic classes are strictly ordered. TC0 always selects the lowest priority, etc. | 
| **automap**  (yes *\| no* ; Default:**yes** ) | Automatically map packets with matching PCP or DSCP values to this QoS profile. Only applies to **trusted** ports. | 


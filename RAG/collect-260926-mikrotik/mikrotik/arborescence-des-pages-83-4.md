---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-83-4
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["latency", "memory", "throughput", "voice"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-83.md
source_anchor: ""
source_lines: [222, 286]
sha256: 4d3edaf68a857f6652d8fbcf6516a2ad90e1653519e03afd84c3d3eee8c1b174
---

# Overview

RouterOS allows enabling/disabling the shared pool for each queue individually - for example, to prevent low-priority traffic from consuming the entire hardware memory. In addition, port buffer limits may prevent a single low-speed port from consuming the entire shared pool. See QoS Settings and Transmission Manager for details.

The default, best-effort (PCP=0, DSCP=0) traffic class is 1, while the lowest priority (PCP=1) has traffic class 0.

## Hardware Resources

The hardware (switch chips) has limited resources (memory). There are two main hardware resources that are relevant to QoS:

- *Packet descriptors* - contain packet control information (target port, header alternation, etc).
- *Data buffers* - memory chunks containing the actual payload. Buffer size depends on the switch chip model. Usually - 256 bytes.

One packet descriptor may use multiple buffers (depending on the payload size); buffers may be shared by multiple descriptors - in cases of multicast/broadcast. If the hardware does not have enough free descriptors or buffers, the packet gets dropped (*tail-drop*).

Hardware resources can be limited per destination type (multicast/unicast), per port, and per each tx queue. If any limits are reached, no more packets can be enqueued for transmission, and further packets get dropped.

RouterOS obscures low-level hardware information, allowing to set resource limits either in terms of packets or a percentage of the total amount. RouterOS automatically calculates the required hardware descriptor and buffer count based on the user-specified packet limit and port's MTU. Moreover, RouterOS comes with preconfigured hardware resources, so there is no need to do a manual configuration in common QoS environments.

Changing any hardware resource allocation parameter in runtime results in a temporary device halt when no packets can be enqueued nor transmitted. Temporary packet loss is expected while the device is forwarding traffic.

## Resource Saving (Previous Generation Devices)

This section does not apply to devices that support **Dynamic Buffers**.

Since reallocating hardware resources in runtime is not an option, RouterOS cannot automatically free queue buffers reserved for inactive ports. Those buffers remain unused. However, if the user knows that the specific ports will *never* be used (e.g., stay physically disconnected), the respective queue resources can be manually freed by using the built-in "offline" tx-manager with minimum resources, for example:

Avoid configuring the "offline" tx-manager on **switch-cpu** port.

When configuring `tx-manager` setting to QSFP+ or QSFP28 interfaces, you must apply the same configuration to all four sub-interfaces of a port. For example, if the interface qsfp28-1-1 is active and linked at 100Gbps, while sub-interfaces (qsfp28-1-2, qsfp28-1-3, qsfp28-1-4) are showing a non-running flag, **do not assign the "offline" tx-manager to thouse non-running sub-interfaces**. Doing so will impact the 100Gbps link as well. However, if none of the four sub-interfaces are running, it is safe to assign the "offline" tx-manager setting.

## Traffic Prioritization

The hardware provides two types of traffic transmission prioritization:

- **Strict Priority** - traffic from higher queues is always transmitted first;
- **Enhanced Transmission Selection (ETS)** - multiple queues participate in packet transmission scheduling at the same time.

**Strict priority** queues are straightforward. If the highest priority queue (Q7) has packets, those are transmitted first. When Q7 is empty, packets from Q6 get transmitted, and so on. The packets from the lowest priority queue (Q0) are transmitted only if all other queues are empty.

The downside of strict prioritization is increased latency in lower queues while "overprioritizing" higher queues. Suppose the acceptable latency of TC5 is 20ms, TC3 - 50ms. Traffic appearing in Q5 gets immediately transmitted due to the strict priority of the queue, adding extra latency to every packet in the lower queues (Q4..Q0). A packet burst in Q5 (e.g., a start of a voice call) may temporarily "paralyze" Q3, increasing TC3 latencies over the acceptable 50ms (or even causing packet drops due to full queue) while TC5 packets get transmitted at <1ms (way below the 20ms limit). Slightly sacrificing TC5 latency by transmitting TC3 packets in between would make everybody happy. That **ETS** is for.

**Enhanced Transmission Selection (ETS)** schedule traffic for transmission from multiple queues (group members) in a weighted round-robin manner. A queue's weight sets the number of packets transmitted from the queue in each round. For example, if Q2, Q1, and Q0 are the group members, and their weights are 3, 2, and 1, respectively, the scheduler transmits 3 packets from Q2, 2 - from Q1, and 1 - from Q0. The actual Tx order is "Q2, Q1, Q0, Q2, Q1, Q2" - for even fairer scheduling.

There are two hardware groups: `low-priority-group` and `high-priority-group`. There is a strict priority ordering between the two groups: the low-priority-group is transmitting only when *all* queues in the high-priority-group are empty. However, it is possible to use only one group for all queues.

The default (built-in) RouterOS queue setup is listed below. Q3-Q5 share the bandwidth within the high-priority group, where packets are transmitted while Q6 and Q7 are empty. Q0-Q2 are the members of the low-priority-group, where packets are transmitted while Q3-Q7 are empty.

It is recommended that all group members are adjacent to each other.

# Active Queue Management (AQM)

WRED is a per-queue congestion control mechanism that signals congestion events to the end-points by dropping packets. WRED relies on the existence of rate throttling mechanisms in the end-points that react to packet loss, such as TCP/IP. WRED uses a randomized packet drop algorithm in an attempt to anticipate congestion events and respond to them by throttling traffic rates before the congestion actually happens. The randomness property of WRED prevents throughput collapse related to the global synchronization of TCP flows.

WRED can be enabled/disabled per each queue in each Tx Manager. Disable WRED for lossless traffic! Also, there is no reason to enable WRED on high-speed ports where congestion should not happen in the first place.

The behavior is controlled via *WRED threshold*. WRED threshold is the maximum number of packets/bytes that can exceed the queue shared buffer limit (cap). A random packet drop begins when queue usage exceeds their respective capacities:

- `queueX-packet-use > queueX-shared-packet-cap` or
- `queueX-byte-use > queueX-shared-byte-cap` .

The more usage exceeds capacity, the higher the packet drop chance, reaching 100% at `queueX-shared-packet-cap + wred-packet-threshold` (or byte).

RouterOS automatically chooses the actual WRED threshold values according to queue or shared pool capacities. The user may shift the thresholds in one way or another via QoS Settings.

WRED requires the respective Tx queues to use shared buffers (**use-shared-buffers=yes**).


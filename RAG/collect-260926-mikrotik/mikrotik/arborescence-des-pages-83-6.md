---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-83-6
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-83.md
source_anchor: ""
source_lines: [346, 443]
sha256: 1b10fea224a8ad5fd2881e12b7bcfb8534cfede06eca2ce6adf9ce4b3ca24b7c
---

# Overview

| Property | Description | 
|---|---|
| **egress-rate-queue0 .. egress-rate-queue7** (*integer: 0..18446744073709551615* ; Default**!egress-rate-queuex** ) | Sets egress traffic limitation (bits per second) for specific output queue. It is possible to specify the limit using suffixes like k, M, or G to represent kbps, Mbps, or Gbps. This setting can be combined with the overall per-port limit **egress-rate** (see ).**/in/eth/sw/port** | 
| **map**  (*name* ; Default:**default** ) | Allows user-defined QoS priority-to-profile mapping in the case of a trusted port or host (see **`/in/eth/sw/qos/map`** ). | 
| **pfc** (*name* ; Default:**disabled** )  | The name of the PFC profile to control ingress priority-based traffic flow (see **`/in/eth/sw/qos/priority-flow-control`**`).` | 
| **profile** (*name* ; Default:**default** ) | The name of the QoS profile to assign to the ingress packets by default (see **`/in/eth/sw/qos/profile`** ). | 
| **trust-l2** (*ignore \| trust \| keep* ; Default:**ignore** ) | Whenever to trust the Layer 2 headers of the incoming packets (802.1p PCP field):  | 
| **trust-l3** (*ignore \| trust \| keep* ; Default:**ignore** ) | Whenever to trust the Layer 3 headers of the incoming packets (IP DSCP field):  | 
| **tx-manager**  (*name* ; Default:**default** ) | The name of the Transmission Manager that is responsible for enqueuing and transmitting packets *from* the given port (see**`/in/eth/sw/qos/tx-manager`** ). | 

L3 trust mode has higher precedence than L2 unless *trust-l3=ignore* or the packet does not have an IP header.

Forwarded/routed packets obtain priority field values (PCP, DSCP) from the selected QoS profile, overwriting the original values unless the respective trust mode is set to **keep**.

Commands.

| Command | Description | 
|---|---|
| **print** | Print the above properties in a human-friendly format. | 
| **print stats** | Print port statistics: total and per-queue transmitted/dropped packets/bytes. | 
| **reset-counters** | Reset all counters in port statistics to zero. | 
| **print usage** | Print queue usage/resources. | 
| **print pfc**  | Pring Priority Flow Control stats | 
| **print rates**  | Print per-queue egress traffic limitation (set by **egress-rate-queueX** ) | 

### Port Stats

| Property | Description | 
|---|---|
| **name** | Port name. | 
| **tx-packet** | The total number of packets transmitted via this port. | 
| **tx-byte** | The total number of bytes transmitted via this port. | 
| **drop-packet** | The total number of packets should have been transmitted via this port but were dropped due to a lack of resources (e.g., queue buffers) or QoS Enforcement. | 
| **drop-byte** | The total number of bytes should have been transmitted via this port but were dropped. | 
| **tx-queue0-packet** .. **tx-queue7-packet** | The number of packets transmitted via this port from the respective queue. | 
| **tx-queue0-byte** ..**tx-queue7-byte** | The number of bytes transmitted via this port from the respective queue. | 
| **drop-queue0-packet** ..**drop-queue7-packet** | The number of packets dropped from the respective queue (or not enqueued at all due to lack of resources). | 
| **drop-queue0-byte** ..**drop-queue7-byte** | The number of bytes dropped from the respective queue. | 

### Port Resources/Usage

Due to hardware limitations, some switch chip models may break traffic flow while accessing QoS port `usage` data. Use port `usage` for diagnostics/troubleshooting only. For monitoring, use QoS `monitor` or Port `stats` instead.

| Property | Description | 
|---|---|
| **name** | Port name. | 
| **packet-cap** | Port's packet capacity. The maximum number of packets that can be enqueued for transmission via the port. | 
| **packet-use***<sup>1</sup>* | Port's packet usage. The number of packets that are currently enqueued in all port's queues. | 
| **byte-cap** | Port's byte capacity (buffer size). The maximum number of bytes that can be enqueued for transmission via the port. | 
| **byte-use***<sup>1</sup>* | Port's byte usage. The size of hardware buffers (in bytes) that are currently allocated for packets the enqueued packets. Since the buffers are allocated by blocks (usually - 256B each), the allocated buffer size can be bigger than the actual payload. | 
| **queue0-packet-cap** .. **queue7-packet-cap** <sup>*2*</sup> | Individual queue capacity. The maximum number of packets that can be enqueued in the respective queues (unless the **Shared Buffers** are enabled). | 
| **queue0-shared-packet-cap .. queue7-shared-packet-cap** <sup>*2*</sup> | Shared queue capacity (individual queue capacity + shared buffers). The maximum number of packets that can be enqueued in the respective queues. | 
| **queue0-packet-use** .. **queue7-packet-use** <sup>*2*</sup> | Queue packet usage. The number of enqueued packets in the respective queues. | 
| **queue0-byte-cap** .. **queue7-byte-cap** <sup>*2*</sup> | Individual queue capacity. The maximum number of bytes that can be enqueued in the respective queues (unless the **Shared Buffers** are enabled). | 
| **queue0-shared-byte-cap .. queue7-shared-byte-cap** <sup>*2*</sup> | Shared queue capacity (individual queue capacity + shared buffers). The maximum number of bytes that can be enqueued in the respective queues. | 
| **queue0-byte-use** .. **queue7-byte-use** <sup>*2*</sup> | Queue buffer usage (in bytes). The size of hardware buffers (in bytes) that are currently allocated for packets in the respective queues. | 
| **queue0-byte-max .. queue7-byte-max** <sup>*2*</sup> | Maximum queue buffer fill level (in bytes). Available only on devices that provide the queue statistics service. Use the **reset-counters** command to reset values. | 

<sup>1</sup> Port's packet/byte usage can exceed the capacity if **Shared Buffers** are enabled.

<sup>2</sup> Only the queues in use are printed.

### Port PFC Stats (Previous Generations)

| Property | Description | 
|---|---|
| **name** | Port name. | 
| **pfc** | PFC profile name. | 
| **pfc-rx**  | Received PFC frame count. | 
| **pfc-tx** | Transmitted PFC frame count. | 
| **pfc-paused-tc** | The list of traffic classes should be paused (from the sender's perspective). PFC pause frames (XOFF) are periodically sent with the listed timers set from this port. | 
| **pfc0-pause-threshold .. pfc7-pause-threshold**  | Pause thresholds of the respective traffic classes. Only PFC-enabled traffic classes are displayed. | 
| **pfc0-resume-threshold .. pfc7-resume-threshold** | Resume thresholds of the respective traffic classes. Only PFC-enabled traffic classes are displayed. | 
| **pfc0-use .. pfc7-use** | The current buffer usage of the respective traffic classes (in bytes). In other words, it is the total size of all queued packets on all ports that were received from this port. Only PFC-enabled traffic classes are displayed. | 

### Port PFC Stats (New Generations)

| Property | Description | 
|---|---|
| **name** | Port name. | 
| **pfc** | PFC profile name. | 
| **pfc0-use .. pfc7-use** | The current buffer usage of the respective traffic classes (in bytes). In other words, it is the total size of all queued packets on all ports that were received from this port. Only PFC-enabled traffic classes are displayed. | 
| **rx-pause** | Received pause frame count. | 
| **tx-pause** | Transmitted pause frame count. | 

## QoS Menu

**Sub-menu:** `/interface/ethernet/switch/qos`

Almost the entire QoS HW configuration is located under **`/in/eth/sw/qos`**. Such an approach allows storing all QoS-related configuration items in one place, easy monitoring and exporting (`/in/eth/sw/qos/export`).

QoS entries have two major flags:

- **H** - Hardware-offloaded.
- **I** - Inactive.

**Sub-menu:** `/interface/ethernet/switch/qos/settings`


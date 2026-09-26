---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-83-8
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-83.md
source_anchor: ""
source_lines: [500, 583]
sha256: 19177ae06c7a1909fd6e6d5e5519e284d7e470a9472dbea71e15508b9420040a
---

# Overview

**Sub-menu:** `/interface/ethernet/switch/qos/map`

Priority-to-profile mapping table(-s) for trusted packets. All switch chips have one built-in map - **default**. In addition, some models allow the user to define custom mapping tables and assign different maps to various switch ports via the **qos-map** property:

- devices based on **Marvell Prestera **98DX224S, 98DX226S**** , or**98DX3236** switch chip models support only one map - default.
- devices based on **Marvell Prestera 98DX8xxx** ,**98DX4xxx** switch chips, or**98DX325x** model devices support up to 12 maps (the default + 11 user-defined).

| Property | Description | 
|---|---|
| **name** (*string* ; Default: ) | The user-defined name of the mapping table. | 

### `VLAN Map`

**Sub-menu:** `/interface/ethernet/switch/qos/map/vlan`

Matches VLAN priorities (802.1p PCP/DEI fields) to QoS profiles. By default, all values are matched to the default QoS profile.

| Property | Description | 
|---|---|
| **dei-only** (yes *\| no* ; Default:**no** ) | Map only packets with DEI (formerly CFI) bit set in the VLAN header. | 
| **map** (*name* ; Default:**default** ) | The name of the mapping table. | 
| **profile** (*name* ; Default: ) | The name of the QoS profile to assign to the matched packets. | 
| **pcp** (*range: 0..7* ; Default:**0** ) | VLAN priority (PCP) value(-s) for the lookup. | 

### DSCP Map

**Sub-menu:** `/interface/ethernet/switch/qos/map/ip`

Matches DSCP values to QoS profiles.

| Property | Description | 
|---|---|
| **dscp** (*range: 0..63* ; Default:**0** ) | DSCP value(-s) for the lookup. | 
| **map** (*name* ; Default:**default** ) | The name of the mapping table. If not set, the standard (built-in) mapping table gets altered. | 
| **profile** (*name* ; Default: ) | The name of the QoS profile to assign to the matched packets. | 

**Sub-menu:** `/interface/ethernet/switch/qos/tx-manager`

Transmission (Tx) Manager controls packet enqueuing for transmission and packet tx order. Different switch ports can be assigned to different Tx managers. The maximum number of hardware Tx managers depends on the switch chip model.

| Property | Description | 
|---|---|
| **name** (*string* ; Default: ) | The user-defined name of the Tx Manager | 
| **queue-buffers** *(percent: 0%..100% \| bytes \| auto;*  Default: **auto** ) | The total amount of hardware Tx buffers allocated to all ports linked to this Tx Manager. Any value but **auto** is NOT scaled by the number of ports. For example, if queue-buffers=30%, and there are 3 ports using this Tx Manager, each respective port receives 10% of total available resources. Adding two more ports to the Tx Manager drops per-port buffers down to 6% (30/5). | 

Port status has not effect on the allocated resources. Running ports receive the same amount of queue buffers as disconnected or disabled ones if all of them are assigned to the same Tx Manager.

### Transmission Queue Scheduler

**Sub-menu:** `/interface/ethernet/switch/qos/tx-manager/queue`

Each port has eight Tx queues. The assigned Tx Manager controls packet enqueuing and schedules transmission orders. Each queue can have either strict priority (where packets with the highest traffic class are always transmitted first) or grouped together for a weighted round-robin tx schedule.

Creating a Tx Manager automatically creates all eight respective queue schedulers.

Changing any properties of Tx manager or queues completely halts traffic enqueueing and transmission during the offload process. Temporary packet loss is expected while the device is forwarding traffic.

| Property | Description | 
|---|---|
| **tx-manager** (*name* ;*read-only* ) | The linked Tx Manager | 
| **traffic-class** (*integer: 0..7* ; *read-only* ) | The traffic class (tc0..tc7) and the respective port queue (queue0..queue7) that the scheduler controls. | 
| **schedule**  (*strict-priority \| high-priority-group \| low-priority-group* ) |  | 
| **weight** *(integer: 0..255;* Default:**1** ) | The weight value for the traffic class if it is a member of a schedule group. The field is not used in the case of strict priority schedule. | 
| **queue-buffers** *(percent: 0%..100% \| bytes \| auto;*  Default: **auto** ) | The amount of hardware Tx buffers allocated to this queue. Any value but **auto** is NOT scaled by the number of ports, i.e., the value gets split on ports linked to the Tx Manager. When given in percent, it means percentage of the tx-manager's queue-buffers value. | 
| **use-shared-buffers** *(yes \| no)* | Allow the queue to use the shared buffer pool when **queue-buffers**  are full. If the queue is full and the shared buffers are disabled, the packet gets dropped. If the shared buffers are enabled, the queue may use up to**shared-packet-cap** or**shared-poolX-packet-cap** (see QoS Settings for details) packets from the shared pool. | 
| **wred** (*yes \| no* ; Default:**no)** | Enables/disables Weighted Random Early Detection for the given queue. | 
| ***ecn (****yes \| no; Default: **no**)* | Enables/disables ECN marking of the transmitted packets. | 
| **wred-actual**  (*yes \| no* ;  read-only)  | The actual WRED value. | 
| **ecn-actual** (*yes \| no* ;  read-only) | The actual ECN value. | 

On some device models, due to hardware limitations, enabling ECN on one queue turns on CE marking of ECN-capable packets on all queues. In such cases, `ecn-actual=yes` despite `ecn=no`.

**Sub-menu:** `/interface/ethernet/switch/qos/priority-flow-control`

PFC configuration is organized in profiles. Different switch ports can be assigned to different PFC profiles. The maximum number of hardware Tx managers depends on the switch chip model. The builtin profile named "**disabled**" cannot be changed.

| Property | Description | 
|---|---|
| **name** (*string* ; Default: ) | The user-defined name of the PFC profile | 
| **pause-threshold** (*percent: 0%..100% \| bytes \| auto;* Default:**auto)** | Transmits a pause frame (XOFF) when the total size of enqueued packets reaches this threshold. Enqueued packets are counted per ingress port. Applies only when **tx=yes** . The value can be given either explicitly in bytes or percent of the respective shared pool size (**shared-poolX-byte-cap** ). | 
| **resume-threshold** (*percent: 0%..100% \| bytes \| auto;* Default:**auto)** | Transmits a resume frame (XON) when the total size of enqueued packets drops down to this threshold. Enqueued packets are counted per ingress port. Applies only when **tx=yes** . The value can be given either explicitly in bytes or percent of the respective shared pool size (**shared-poolX-byte-cap** ). | 
| **rx** (*yes \| no* ; Default:**no)** | Enables receiving of PFC frames. The received PFC frame pauses the specific priority queues on the port that received the PFC frame for the duration specified by the PFC frame. Disabling rx disables queue pausing. | 
| **traffic-class** (*integer array: 0..7* )  | The list of PFC-enabled traffic classes. | 
| **tx** (*yes \| no* ; Default:**no)** | Enables transmition of PFC frames. |

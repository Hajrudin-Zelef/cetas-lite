---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-83-5
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "memory", "throughput"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-83.md
source_anchor: ""
source_lines: [287, 345]
sha256: 535e1663f1f8f826754305dab3a375702867643bc0b1695446ac05d68feff141
---

# Overview

Choosing a WRED threshold value is a tradeoff between congestion anticipation and burst absorption. Setting a higher WRED threshold may lead to earlier traffic rate throttling and, therefore, resolve congestion. On the other hand, a high threshold leads to packet drops in limited traffic bursts that could be absorbed by the queue buffers and transformed losslessly if WRED didn't kick in. For instance, initiating a remote database connection usually starts with heavier traffic ("packet burst") at the initialization phase; then, the traffic rate drops down to a "reasonable" level. Any packet drop during the initialization phase leads to nothing but a slower database connection due to the need for retransmission. Hence, lowering the WRED threshold or entirely disabling WRED on such traffic is advised. The opposite case is video streaming. Early congestion detection helps select a comfortable streaming rate without losing too much bandwidth on retransmission or/and "overshooting" by sacrificing the quality level by too much.

Use Switch Rules (ACL) or other QoS Marking techniques to differentiate traffic and put packets into queues with desired WRED settings.

The following script only applies WRED to TCP/IP traffic by redirecting it to queue2. UDP and other packets are left in queue1 - since their end-points usually cannot respond to early drops. Queue1 and queue2 are scheduled equally - without prioritizing one queue over another.

## (ECN)

Some switch chips can perform ECN marking of IP packets on the hardware level, according to RFC 3168. Hardware ECN marking is based on the WRED mechanism, but instead of dropping IP packets, they are marked with CE (Congestion Experienced, binary 11) in the ECN field (two least significant bits in IPv4/TOS or IPv6/TrafficClass octet). Only ECN-Capable IP packets may be marked - those with the ECN field value of ECT(1) or ECT(0) (binary 01 or 10, respectively). Not ECN-Capable Transport packets (ECN=00) never get marked. If a packet already has the CE mark (ECN=11), it never gets cleared, even if the device does not experience congestion.

Set **ecn=yes** on Tx Manager Queue to enable ECN marking.

ECN marking mechanism requires the respective Tx queues to use shared buffers (**use-shared-buffers****=yes)**.

The packet receives the CE mark if all conditions below are met:

1. The packet is either IPv4 or IPv6.
2. The ECN field value in IP header is either ECT(1) or ECT(0).
3. Egress port's Tx Queue has **ecn=yes** and uses shared buffers (**use-shared-buffers=yes** ).
4. `queueX-packet-use > queueX-shared-packet-cap` or`queueX-byte-use > queueX-shared-byte-cap` .

Since enabling ECN (ecn=yes) prevents ECN-capable packet drop, queue usage may exceed WRED thresholds if the traffic sender doesn't react to congestion notification in time.

Priority-Based Flow Control (PFC) provides lossless operation for up to eight traffic classes, so that congestion in one traffic class does not pause other traffic classes. In addition, PFC enables co-existence of loss-sensitive traffic types with loss tolerant traffic type in the same network.

PFC-capable switch chips are complaint with IEEE 802.1Qbb PFC, meaning that the respective devices are capable of generating and responding to PFC frames. On the triggering part, the PFC frame is sent by the source port and traffic class experiencing the congestion. The timer values of the generated PFC frames are 0xFFFF for pause (XOFF) and 0x0 for resume (XON), and the appropriate bit in the priority enable vector is set. On the response part, the received PFC frame pauses the specific priority queues on the port that received the PFC frame for the duration specified by the PFC frame.

In RouterOS, PFC configuration is organized in profiles, where each port can be assigned to a specific profile. A PFC profile defines the traffic classes to enable PFC on, pause/resume thresholds to send XOFF/XON PFC frames, respectively, and whenever the assigned ports should transmit or/and receive PFC frames.

While congestion occurs on egress ports, PFC is triggered on the ingress port. Shared buffers must be used to associate the amount of ingressed traffic with the respective packets waiting in Tx queues. For each PFC-enabled traffic class, set **use-shared-buffers=yes** to the respective Tx Queues.

RouterOS implements 1:1 mapping between traffic classes and Tx queues. Packets with assigned traffic class 0 get enqueued in queue0, TC1 - queue1, etc., up to TC7-Q7. Hence, the terms "traffic class" and "tx queue" are used interchangeably in this text.

When choosing pause and resume thresholds, consider a delay in transmitting a PFC frame and processing it by the other side. For example, device A experienced congestion at time T, transmitted a PFC pause frame to device B, and B processed the frame and halted transmission at time T+D. During the delta time D, device B still kept sending traffic. If device A has configured the pause threshold to 100%, it has no free buffers available, and, therefore, packets may drop, which is unacceptable for lossless traffic classes. Lowering the pause threshold, let's say, down to 80% issues a PFC pause frame while still having free memory to accumulate traffic during the delta time D. The same applies to resume threshold. Setting it to 0% keeps the device idle during the delta time, lowering the overall throughput.

PFC Rx requires setting the egress rate to all associated queues to calculate pause time, even if it matches the wire speed. For example, if PFC runs on traffic class 3, the assigned ports require the `egress-rate-queue3` setting.

# Property Reference

## Switch settings

**Sub-menu:** `/interface/ethernet/switch`

Switch QoS settings (in addition to the existing ones).

| Property | Description | 
|---|---|
| **qos-hw-offloading** (*yes \| no* ; Default:**no** ) | Allows enabling QoS for the given switch chip (if the latter supports QoS). New generation devices force **qos-hw-offloading=yes** at all times. | 

When you enable QoS, turning off the **qos-hw-offloading** setting will not completely revert to the previous functionality. It is recommended to reboot the device after disabling it.

## Port settings

**Sub-menu:** `/interface/ethernet/switch/qos/port`

Switch port QoS settings. Assigns a QoS profile to ingress packets on the given port. The assigned profile can be changed via match rules if the port is considered trusted.

By default, ports are untrusted and receive the default QoS profile (Best-Effort, PCP=0, DSCP=0), where priority fields are cleared from the egress packets.


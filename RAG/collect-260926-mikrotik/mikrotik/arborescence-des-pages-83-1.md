---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-83-1
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "latency", "memory", "voice"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-83.md
source_anchor: ""
source_lines: [1, 78]
sha256: 14d3be5ae12a3da17f0fa65fcbacc3140a728b882e060a9c7af687065545bfb1
---

# Overview

This document defines **Quality of Service (QoS)** usage in RouterOS based on **Marvell Prestera switch chips**.

QoS is a set of features in network switches that allow network administrators to prioritize traffic and allocate network resources to ensure that important data flows smoothly and with low latency.

The primary function of QoS in network switches is to manage network traffic in a way that meets the specific requirements of different types of network applications. For example, voice and video data require low latency and minimal packet loss to ensure high-quality communication, while file transfers and other data applications can tolerate higher levels of latency and packet loss.

QoS works by identifying the type of traffic flowing through the switch and assigning it a priority level based on its requirements. The switch can then use this information to alter packet headers and prioritize the flow of traffic, ensuring that higher-priority traffic is given preferential treatment over lower-priority traffic.

RouterOS v7.15+ is required to support all QoS features:

1. **QoS Marking.** QoS profile matching by ingress packet headers, then egress header alternation according to the assigned QoS profiles.
2. **QoS Enforcement** . Avoid or resolve congestion based on the assigned QoS profiles and traffic shaping.
3. **QoS Policy** . Assign QoS profiles via ACL rules.
4. Active Queue Management: **WRED** (Weighted Random Early Detection),**ECN** notification, and processing,**PFC** (Priority-based Flow Control).
5. Traffic shaping.

## QoS Changes in RouterOS v7.23

RouterOS version 7.23 introduced significant changes in Quality of Service offloading, simplifying the configuration process and introducing "**lossless**" traffic classes. While hardware resources are finite, an active queue management with the help of ECN and/or PFC prevents packet loss if all parties support the features and are configured properly. Prior RouterOS versions supported ECN and PFC as well, but in v7.23, the configuration process was streamlined.

Not all devices support lossless traffic. Check the QoS Device Support table.

### Configuration Change Summary

- QoS Settings: most changes were set to "auto" by default, allowing RouterOS to pick the best known settings for the setup.
- QoS Settings: added `lossless-traffic-class` and`lossless-buffers` , allowing explicit specification of lossless traffic and the reservation of queue resources for it. The user can leave those fields to "auto" as well.
- QoS Monitor: displays the shared pool usage for lossy and lossless separately.
- Tx Queue: removed shared pool index. Router OS automatically selects the shared pool based on the traffic type (lossy or lossless).
- QoS Profile: added automapping to the specified PCP and DSCP values (enabled by default). For example, adding a QoS profile with DSCP=46 automatically applies the profile to traffic received from trusted ports.

# QoS Terminology

These terms will be used throughout the article.

- **QoS** - Quality of Service.
- **ACL** - Access Control List, a set of switch rules used to filter network traffic based on specified criteria.
- **AQM** - Active Queue Management.
- **DSCP** - Differentiated Services Code Point, a 6-bit field in the IP header used to prioritize network traffic.
- **ECN** - Explicit Congestion Notification.
- **ETS** - Enhanced Transmission Selection.
- **PCP** - Priority Code Point, a 3-bit field in the VLAN header used to prioritize traffic within a VLAN.
- **PFC** - Priority-based Flow Control (IEEE 802.1Qbb).
- **RoCE** - RDMA over Converged Ethernet.
- **WRED** - Weighted Random Early Detection.
**/in/eth/sw/**
**/interface/ethernet/switch/**

# QoS Enhancements in new generation Marvell Prestera Switch Chips

MikroTik devices with new generation Marvell Prestera switch chips (e.g. CRS8xx series running the switch-marvell package) offer a new approach in QoS enforcement. Previous models required making a tradeoff between guaranteed and shared Tx queue buffers. By increasing the shared buffer percentage, the device can absorb larger traffic bursts, but at the same time, one congested port can occupy all shared buffers, limiting the QoS capabilities of other ports. Using the new generation switch chips enable the best of both worlds by introducing **Dynamic Buffers**. A switch chip dynamically adjusts port and queue buffer limits based on the current congestion level. If there is no congestion within the device, a single port can absorb a large burst of traffic, preventing packet loss. When the Tx queue size increases across multiple ports, the device reduces per-port and per-queue limits, enforcing fair use of shared resources while ensuring enough buffers for non-congested queues to continue forwarding traffic. The entire process is automated, requiring zero configuration from the user.

Another improvement in new generation Marvell Prestera switch chips is eliminating the shortage of enqueued packet descriptors for storing control data. Previous generation devices had Packet Cap and Use stats. When a device received a burst of small packets, it could start tail-dropping packets due to hitting Packet Cap, even though it had enough queue buffers to store the payload. New generation devices ensure there is always enough memory to store control data, leaving the Dynamic Buffer size as the only limiting factor. That's why there are no Packet Cap/Use stats in new generation devices.

New generation devices require QoS HW Offloading to be enabled at all times (**qos-hw-offloading=yes**). RouterOS ignores user requests to disable it while keeping the field for backward compatibility.

| Switch Chip | Models | QoS Profiles | QoS Maps | Tx Managers | WRED | ECN | PFC | Lossless Buffers | Dynamic Buffers | Port/Queue Usage Stats | 
|---|---|---|---|---|---|---|---|---|---|---|
| **98DX3236** | CRS305-1G-4S+IN CRS326-24G-2S+ (RM/IN) CRS328-24P-4S+RM CRS328-4C-20S-4S+RM | 128 | 1 | 8 |  |  |  |  |  | Current values | 
| **98DX226S** | CRS305-1G-4S+OUT (FiberBox Plus) CRS310-1G-5S-4S+ (netFiber 9/IN) CRS310-8G+2S+IN CRS318-16P-2S+OUT (netPower 16P) CRS320-8P-8B-4S+RM CRS418-8P-8G-2S+RM CRS418-8P-8G-2S+5axQ2axQ-RM | 128 | 1 | 8 |  |  |  |  |  | Current values | 
| **98DX224S** | CRS318-1Fi-15Fr-2S-OUT (netPower 15FR) | 128 | 1 | 8 |  |  |  |  |  | Current values | 
| **98DX2528** | CRS304-4XG-IN | 128 | 1 | 8 |  |  |  |  |  | Current values | 
| **98DX8525** | CCR2216-1G-12XS-2XQ CRS518-16XS-2XQ-RM | 1024 | 12 | 15 | ✔ | ✔ | ✔ | ✔ |  | Max fill <sup>1</sup> | 
| **98DX4310** | CRS504-4XQ (IN/OUT) CRS510-8XS-2XQ-IN RDS2216-2XG-4S+4XS-2XQ | 1024 | 12 | 15 | ✔ | ✔ | ✔ | ✔ |  | Max fill <sup>1</sup> | 
| **98DX8208** | CRS309-1G-8S+IN | 1024 | 12 | 15 | ✔ | ✔ | ✔ | ✔ |  | Current values <sup>2</sup> | 
| **98DX8212** | CRS312-4C+8XG-RM | 1024 | 12 | 15 | ✔ | ✔ | ✔ | ✔ |  | Current values <sup>2</sup> | 
| **98DX8216** | CRS317-1G-16S+RM | 1024 | 12 | 15 | ✔ | ✔ | ✔ | ✔ |  | Current values <sup>2</sup> | 
| **98DX8332** | CRS326-24S+2Q+RM CRS326-4C+20G+2Q+RM | 1024 | 12 | 15 | ✔ | ✔ | ✔ | ✔ |  | Current values <sup>2</sup> | 
| **98DX3257**  | CRS354-48G-4S+2Q+RM CRS354-48P-4S+2Q+RM | 1024 | 12 | 15 | ✔ | ✔ | ✔ | ✔ |  | Current values <sup>2</sup> | 
| **98DX3255** | CCR2116-12G-4S+ | 1024 | 12 | 15 | ✔ | ✔ | ✔ | ✔ |  | Current values <sup>2</sup> | 
| **98CX8410** | CRS520-4XS-16XQ-RM | 1024 | 12 | 15 | ✔ | ✔ | ✔ | ✔ |  | Unavailable <sup>3</sup> | 
| **98DX7335** | CRS812-8DS-2DQ-2DDQ-RM CRS804-4DDQ-hRM | 1024 | 12 | 15 | ✔ | ✔ | ✔ | ✔ | ✔ | Current values + Max fill | 

<sup>1</sup> The device gathers max queue fill statistics instead of displaying the current usage values. Use the **reset-counters** command to reset those stats.

<sup>2</sup> Due to hardware limitations, some switch chip models may break traffic flow while accessing QoS port/queue usage data.


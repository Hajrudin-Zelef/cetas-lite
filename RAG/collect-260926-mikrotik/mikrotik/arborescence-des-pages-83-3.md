---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-83-3
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "memory"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-83.md
source_anchor: ""
source_lines: [172, 221]
sha256: dba1437de887ba046bcf44bb25f1fef22a5ba452bf11659807f5f09e8bb65e4d
---

# Overview

Each switch port has Layer2 and Layer3 trust settings that will change how ingress packets are classified into QoS profiles and what PCP and DSCP values will be used. Below are tables that describe all possible options:

| **qos-trust-l2** | **qos-trust-l3** | Behavior | 
|---|---|---|
| **ignore** | **ignore** | The port is considered untrusted. Both headers are ignored, and the port's **profile** is forced to all ingress packets. This is the default setting. | 
| **ignore** | **trust** | Trust the Layer 3 header. Use the DSCP field from the IP header of ingress packets for QoS profile lookup (see `/in/eth/sw/qos/map/ip` ). If the lookup fails (no QoS profiles are mapped to the given DSCP value), the**default** QoS profile is used (not the switch port's QoS profile). The switch port's**profile** field is used only for non-IP traffic. | 
| **ignore** | **keep** | Trust the Layer 3 header. Use the DSCP field from the IP header of ingress packets for QoS profile lookup (see `/in/eth/sw/qos/map/ip` ). If the lookup fails, the**default** QoS profile is used. The switch port's**profile** field is used only for non-IP traffic. If the forwarded/routed packet is VLAN-tagged, its PCP value is set from the selected QoS profile. However, the original DSCP value of the packet is kept intact. | 
| **trust** | **ignore** | Trust the Layer 2 header, but ignore L3. If an ingress packet is VLAN-tagged, use the PCP field from the VLAN header for QoS profile lookup (see `/in/eth/sw/qos/map/vlan` ). If the lookup fails (no QoS profiles are mapped to the given PCP value), the**default** QoS profile is used. The switch port's**profile** field is used only for untagged traffic. | 
| **trust** | **trust** | Trust both headers, but Layer 3 has higher precedence. In the case of an IP packet, use the DSCP field for QoS profile lookup (see `/in/eth/sw/qos/map/ip` ). If the DSCP-to-QoS lookup fails, use the**default** profile. If the packet is not an IP packet but is VLAN-tagged, use the PCP field from the VLAN header for QoS profile lookup (see`/in/eth/sw/qos/map/vlan` ).  If the VLAN-to-QoS lookup fails, use the**default** QoS profile. Non-IP untagged packets use the switch port's**profile** . | 
| **trust** | **keep** | The same as **trust+trust** , but the original DSCP value is preserved in forwarded/routed packets. | 
| **keep** | **ignore** | Trust the Layer 2 header but ignore L3. If an ingress packet is VLAN-tagged, use the PCP field from the VLAN header for QoS profile lookup (see `/in/eth/sw/qos/map/vlan` ). If the lookup fails (no QoS profiles are mapped to the given PCP value), the**default** QoS profile is used. The switch port's**profile** field is used only for untagged traffic. If the packet is VLAN-tagged on both ingress and egress, the original PCP value is kept. | 
| **keep** | **trust** | Trust both headers, but Layer 3 has higher precedence. In the case of an IP packet, use the DSCP field for QoS profile lookup (see `/in/eth/sw/qos/map/ip` ). If the DSCP-to-QoS lookup fails, use the**default** profile. If the packet is not an IP packet but is VLAN-tagged, use the PCP field from the VLAN header for QoS profile lookup (see`/in/eth/sw/qos/map/vlan` ).  If the VLAN-to-QoS lookup fails, use the**default** QoS profile. Non-IP untagged packets use the switch port's**profile** . If the packet is VLAN-tagged on both ingress and egress, the original PCP value is kept. The DSCP value in forwarded/routed packets is set from the selected QoS profile. | 
| **keep** | **keep** | Trust both headers, but Layer 3 has higher precedence. In the case of an IP packet, use the DSCP field for QoS profile lookup (see `/in/eth/sw/qos/map/ip` ). If the DSCP-to-QoS lookup fails, use the**default** profile. If the packet is not an IP packet but is VLAN-tagged, use the PCP field from the VLAN header for QoS profile lookup (see`/in/eth/sw/qos/map/vlan` ).  If the VLAN-to-QoS lookup fails, use the**default** QoS profile. Non-IP untagged packets use the switch port's**profile** . Keep both the original PCP and/or DSCP values intact in cases of VLAN-tagged and/or IP packets, respectively. | 

| **Port settings**  |  | The selected QoS profile and the source for PCP / DSCP field values in forwarded/routed packets |  |  |  |  |  |  |  |  |  |  |  | 
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| **qos-trust-l2**  | **qos-trust-l3**  | VLAN-Tagged IP |  |  | Untagged IP |  |  | VLAN-Tagged Non-IP |  |  | Untagged Non-IP |  |  | 
|  |  | QoS Profile | PCP | DSCP | QoS Profile | PCP *<sup>1</sup>* | DSCP | QoS Profile | PCP | DSCP | QoS Profile | PCP *<sup>1</sup>* | DSCP | 
| ignore | ignore | profile | profile | profile | profile | profile | profile | profile | profile | - | profile | profile | - | 
| ignore | trust | map/ip | map/ip | map/ip | map/ip | map/ip | map/ip | profile | profile | - | profile | profile | - | 
| ignore | keep | map/ip | map/ip | original | map/ip | map/ip | original | profile | profile | - | profile | profile | - | 
| trust | ignore | map/vlan | map/vlan | map/vlan | profile | profile | profile | map/vlan | map/vlan | - | profile | profile | - | 
| trust | trust | map/ip | map/ip | map/ip | map/ip | map/ip | map/ip | map/vlan | map/vlan | - | profile | profile | - | 
| trust | keep | map/ip | map/ip | original | map/ip | map/ip | original | map/vlan | map/vlan | - | profile | profile | - | 
| keep | ignore | map/vlan | original | map/vlan | profile | profile | profile | map/vlan | original | - | profile | profile | - | 
| keep | trust | map/ip | original | map/ip | map/ip | profile | map/ip | map/vlan | original | - | profile | profile | - | 
| keep | keep | map/ip | original | original | map/ip | profile | original | map/vlan | original | - | profile | profile | - | 

<sup>1</sup> applies only when ingress traffic is untagged, but the egress needs to be VLAN-tagged.

Starting from **RouterOS v7.15**, it is possible to assign QoS profiles via Switch Rules (ACL).

**Sub-menu:** `/interface/ethernet/switch/rule`

| New/Changed Properties | Description | 
|---|---|
| **new-qos-profile** (*name* ) | The name of the QoS profile to assign to the matched packets. | 
| **keep-qos-fields** (*yes \| no* ; Default:**no** ) | Should the original values of QoS fields (PCP, DSCP) be kept ( *yes* ), or replace them with the ones from the assigned QoS profile (*no* )? Relevant only if**new-qos-profile** is set. | 
| **new-vlan-priority** (*0..7* ) | Deprecated and should be replaced with the respective **new-qos-profile** . Kept for backward compatibility. Relevant only if qos-hw-offloading=no. | 

The following example assigns a QoS profile based on the source MAC address.

# QoS Enforcement

## Hardware Queues

Each switch port has eight hardware transmission (tx) queues (queue0..queue7). Each queue corresponds to a traffic class (tc0..tc7) set by a QoS profile. Each ingress packet gets assigned to a QoS profile, which, in turn, determines the traffic class for tx queue selection on the egress port.

Hardware queues are of variable size - set by the Transmission Manager. Moreover, multiple ports and/or queues can share resources with each other (so-called *Shared Buffers*). For example, a device with 25 ports has memory (buffers) to queue 1200 packets in total. If we split the resources equally, each port gets 48 exclusive buffers with a maximum of 6 packets per queue (48/8) - which is usually insufficient to absorb even a short burst of traffic. However, choosing to share 50% of the buffers leaves each port with 24 exclusive buffers (3 per queue), but at the same time, a single queue can grow up to 603 buffers (3 exclusive + 600 shared).


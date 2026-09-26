---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-83-2
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "ethernet", "latency", "memory"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-83.md
source_anchor: ""
source_lines: [79, 171]
sha256: 54578b656e63dd918ad8f7fa0645488368c0e5432937867d54a2c7a15c204db2
---

# Overview

<sup>3</sup> Usage data for individual queues on a port are unavailable, only the total usage for the entire port can be accessed.

# Applications and Usage Examples

## Basic Configuration Example

In this example, we define just one QoS level - VoIP (IP Telephony) on top of the standard "Best Effort" class. Let's imagine that we have a CRS326-24G-2S+ device where:

- all ports are bridged and using **vlan-filtering** ;
- sfp-sfpplus1 is a VLAN trunk connected to another switch;
- ether1-ether9 are dedicated ports for IP phones;
- ether10-ether24 are standard ports for host connection;

First, we need to define QoS profiles. Defined `dscp` and `pcp` values that will be used in forwarded packets on egress:

Port-based QoS profile assignment on dedicated ports for IP phones applies to ingress traffic. Other Ethernet ports will use the default `profile` (where `dscp=0` and `pcp=0`):

The trunk port receives both types of QoS traffic, we need to enable `trust-l3` and `trust-l2` to differentiate them:

If you are running RouterOS version 7.23 or later, the VLAN priority and IP DSCP mapping is created automatically, so this step is not required. If you are using an earlier RouterOS version, you must manually create the VLAN priority and IP DSCP mapping with the QoS profile:

Finally, enable QoS hardware offloading for the above settings to start working:

It is possible to verify the port QoS settings with `print` command:

Now incoming packets on ports ether1-ether9 are marked with a Priority Code Point (PCP) value of 5 and a Differentiated Services Code Point (DSCP) value of 46, and incoming packets on ports ether10-ether24 are marked with PCP and DSCP values of 0. When packets are incoming to sfp-sfpplus1 port, any packets with a PCP value of 5 will retain their PCP value of 5 and DSCP value of 46, while all other packets will be marked with PCP and DSCP values of 0.

## Dante

Starting from RouterOS v7.15, all MikroTik QoS-Capable devices are compatible with Dante**.** 

Dante hardware use the following DSCP / Diffserv priority values for traffic prioritization.

| Dante Priority | Usage | DSCP Label | DSCP Value | 
|---|---|---|---|
| High | Time critical PTP events | CS7 | 56 | 
| Medium | Audio, PTP | EF | 46 | 
| None | Other traffic | BE | 0 | 

The example assumes that the switch is using its default configuration, which includes a default "bridge" interface and all Ethernet interfaces added as bridge ports, and any of these interfaces could be used for Dante.

First, create QoS Profiles to match Dante traffic classes, there is already a pre-existing "default" profile that corresponds to Dante's None priority.

If you are running RouterOS version 7.23 or later, the IP DSCP mapping is created automatically, so this step is not required. If you are using an earlier RouterOS version, create a QoS mapping to match QoS profiles based on DSCP values. 

Configure hardware queues to enforce QoS on Dante traffic.

Dante's High and Medium priority traffic is scheduled in strict order. The devices transmit time-critical PTP packets until queue7 gets empty, then proceed with audio (queue5). Other traffic gets transmitted only when PTP and audio queues are empty.

The next step is to enable trust mode for incoming Layer3 packets (IP DSCP field):

Finally, enable QoS hardware offloading for the above settings to start working:

When using Dante in multicast mode, it is beneficial to enable IGMP snooping on the switch. This feature directs traffic only to ports with subscribed devices, preventing unnecessary flooding. Additionally, enabling an IGMP querier (if not already enabled on another device in the same LAN), adjusting query intervals, and activating fast-leave can further optimize multicast performance.

## RDMA over Converged Ethernet (RoCE)

RoCE allows you to directly access memory on remote storage systems using Ethernet networks without involving the host CPU. This capability significantly reduces latency and CPU overhead, making RoCE ideal for high-performance computing and data center environments. RoCE also enables a converged network, where various services (such as data storage, networking, and multimedia) run over a single Ethernet infrastructure. This simplifies network management and reduces the cost and complexity of maintaining separate networks.

RoCE achieves this through the use of **ECN** and **PFC** mechanisms. These features help prevent network congestion and packet loss, ensuring reliable, lossless communication. See the device feature table for compatible switches. Although switches can support RoCE environments, the end hosts must also be compatible with the RoCE protocol and equipped with RDMA-capable network interface cards (NICs).

There are two main versions of RoCE. RoCEv1 operates as an Ethernet link layer protocol and uses Ethertype 0x8915. RoCEv2 works over standard IP networks, using UDP destination port number 4791. ECN bits in the IP header are marked to signal network congestion, and a Congestion Notification Packet (CNP) is used to acknowledge congestion to the sender. For traffic prioritization, DSCP 26 is used for RoCEv2 traffic, while DSCP 48 for CNPs.

The following example can be used for lossless RoCEv2 with PFC and ECN and it assumes that the switch is using its default configuration, which includes a default "bridge" interface and all Ethernet interfaces added as bridge ports. The minimal recommended RouterOS version is 7.17.

First, configure additional profiles. Non-RoCE traffic will be assigned to already existing "default" profile with traffic-class 1, RoCEv2 to traffic-class 3, and CNP to traffic-class 6.

If you are running RouterOS version 7.23 or later, the IP DSCP mapping is created automatically, so this step is not required. If you are using an earlier RouterOS version, create a QoS mapping to match QoS profiles based on DSCP values.

Configure hardware queues and scheduler. We are using ETS (`schedule=high-priority-group`) for traffic-class 1 and traffic-class 3 with 50% bandwith assigment each (`weight=1`), and strict priority scheduling for traffic-class 6. Additionally, enable ECN (`ecn=yes`) for traffic-class 3 to mark IP packets in the switch that experience congestion.

Although using `schedule=low-priority-group` allows you to create separate ETS scheduling and bandwidth allocation for a different set of traffic-classes, it is not recommended to use this setting together with `lldp-dcbx=yes`. The reason is that the ETS Configuration/Recommendation TLVs are designed to handle a single bandwidth allocation across traffic classes, thus `schedule=high-priority-group` should be used instead.

Configure PFC profile for traffic-class 3 to ensure a lossless environment for RoCEv2 traffic.

Set Layer3 trust mode (`trust-l3=keep`) on switch ports where RoCEv2 traffic is expected, set PFC (`pfc=pfc-tc3`) and egress-rate for queue3 to comply with PFC requirements (`egress-rate-queue3=10.0Gbps`). In this example, 10Gbps SFP+ interfaces are used, and the egress rate can be set to match the physical speed of the interface. Change this property depending on your interface speeds.

Enable QoS hardware offloading for the above settings to start working.

Enable the LLDP Data Center Bridging Capability Exchange Protocol (DCBX) to share QoS settings and capabilities with other neighboring devices.

As an optional step, increase the L2MTU to accommodate larger data packets.

# QoS Marking

## Understanding Map ranges

In order to avoid defining all possible PCP and DSCP mappings, RouterOS allows setting multiple values and ranges for PCP and DSCP values for QoS Profile mapping.

In the following example, PCP values 0 and 2 use the default QoS profile, 1, 3-4 - streaming, 5 - voip, and 6-7 - control.

## Understanding Port, Profile, and Map relation


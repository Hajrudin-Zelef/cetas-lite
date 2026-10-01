---
id: collect-261001-general-networking/general-networking/c-en-us-td-docs-switches-lan-c9000-lyr2-fwd-flowcontrol-stormcontrol-flow-contro-39252bf2-1
title: "c-en-us-td-docs-switches-lan-c9000-lyr2-fwd-flowcontrol-stormcontrol-flow-contro-39252bf2"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["asic"]
source: docs/RAG/collect-261001-general-networking/c-en-us-td-docs-switches-lan-c9000-lyr2-fwd-flowcontrol-stormcontrol-flow-contro-39252bf2.md
source_anchor: ""
source_lines: [1, 100]
sha256: 1553cb90505942131ed48c2a165467069dc0788674061259e49b5fa366c967f1
---

# c-en-us-td-docs-switches-lan-c9000-lyr2-fwd-flowcontrol-stormcontrol-flow-contro-39252bf2

Flow Control and Storm Control Configuration Guide
Bias-Free Language
Bias-Free Language
The documentation set for this product strives to use bias-free language. For the purposes of this documentation set, bias-free is defined as language that does not imply discrimination based on age, disability, gender, racial identity, ethnic identity, sexual orientation, socioeconomic status, and intersectionality. Exceptions may be present in the documentation due to language that is hardcoded in the user interfaces of the product software, language used based on RFP documentation, or language that is used by a referenced third-party product. Learn more about how Cisco is using Inclusive Language.
This table provides release and platform support information for the features explained in this module.
These features are available in all the releases subsequent to the one they were introduced in, unless noted otherwise.
Table 1.
Release
Feature Name and Description
Supported Platform
Cisco IOS XE 26.2.1ea
Storm Control: Storm Control feature support has been introduced.
Cisco C9550 Series Smart Switches
Cisco IOS XE 17.18.1
Storm Control is a feature in Cisco networking devices that helps prevent traffic storms caused by broadcast, multicast, or
unicast traffic overwhelming the network.
Cisco C9350 Series Smart Switches
Cisco C9610 Series Smart Switches
Storm Control
Storm control prevents traffic on a LAN from being disrupted by a broadcast, multicast, or unicast storm on one of the physical
interfaces. A LAN storm occurs when packets flood the LAN, creating excessive traffic and degrading network performance. Errors
in the protocol-stack implementation, mistakes in network configurations, or users issuing a denial-of-service attack can
cause a storm.
Storm control (or traffic suppression) monitors packets passing from an interface to the switching bus and determines if the
packet is unicast, multicast, or broadcast. The switch counts the number of packets of a specified type received within the
1-second time interval and compares the measurement with a predefined suppression-level threshold.
Measured Traffic Activity
Storm control uses one of these methods to measure traffic activity:
Bandwidth as a percentage of the total available bandwidth of the port that can be used by the broadcast, multicast, or unicast
traffic
Traffic rate in bits per second at which broadcast, multicast, or unicast packets are received
With each method, when the rising threshold is reached, the port blocks only the excessive traffic above the threshold. If
the falling suppression level is not specified, the device blocks all traffic until the traffic rate drops below the rising
suppression level. In general, the higher the level, the less effective the protection against broadcast storms.
Note
When the storm control threshold for multicast traffic is reached, all multicast traffic except control traffic, such as bridge
protocol data unit (BDPU) and Cisco Discovery Protocol frames, are blocked. However, the device does not differentiate between
routing updates, such as OSPF, and regular multicast data traffic, so both types of traffic are blocked.
Traffic Patterns
Broadcast traffic exceeded the configured threshold during the intervals from T1 to T2 and from T4 to T5. When the specified
traffic exceeds this threshold, storm control drops the excess traffic only during the interval when the storm occurs. In
the following interval (such as T3), if broadcast traffic is below the threshold, it is forwarded as usual.
The combination of the storm-control suppression level and the 1-second time interval controls the way the storm control algorithm
works. A higher threshold allows more packets to pass through. A threshold value of 100 percent means that no limit is placed
on the traffic. A value of 0.0 means that all broadcast, multicast, or unicast traffic on that port is blocked.
Note
Because packets do not arrive at uniform intervals, the 1-second time interval during which traffic activity is measured can
affect the behavior of storm control.
You use the storm-control interface configuration commands to set the threshold value for each traffic type.
Storm Control Using a Hardware Rate Limiter
Traffic storm control monitors incoming traffic levels over a configured interval. However, the reaction time taken by storm
control is slightly slower as it is based on statistics counters to identify a storm. With the hardware rate limiter, the
action is taken at the ASIC level, and as a result, the storm control action starts immediately; as soon as the traffic rate
reaches the set threshold level. The hardware rate limiter implements policers for broadcast, multicast, unicast, and unknown
unicast traffic.
Types of Traffic Managed by Storm Control
Broadcast Traffic: Broadcast packets are sent to all devices in the same broadcast domain (e.g., ARP requests, DHCP Discover
messages). Excessive broadcast traffic can lead to a broadcast storm, overwhelming network devices and causing degraded performance.
Multicast Traffic: Multicast packets are sent to a group of devices that have explicitly joined a multicast group (e.g., video
streaming or IP telephony). If misconfigured or excessively generated, multicast traffic can cause a multicast storm, affecting
network performance.
Unicast Traffic: Unicast packets are sent to a specific MAC address. When a switch receives these packets, it checks its MAC
address table to identify the corresponding port within the VLAN and forwards the packets accordingly. If the number of unicast
packets exceeds the configured storm control threshold, storm control automatically filters and drops the excess packets to
prevent network congestion.
Unknown Unicast Traffic: Unknown unicast packets are destined for a specific MAC address, but the switch does not have the
destination MAC in its MAC address table. These packets are flooded to all ports in the VLAN. Excessive unknown unicast traffic
can lead to an unknown unicast storm, consuming bandwidth and overwhelming devices.
Storm Control uses thresholds to monitor and limit the traffic rate for each traffic type. When traffic exceeds the configured
threshold, the switch takes action, such as dropping packets or shutting down the interface.
Broadcast Traffic: Storm Control monitors the rate of broadcast packets and drops packets exceeding the threshold.
Multicast Traffic: Storm Control applies similar monitoring as broadcast traffic and action for multicast packets.
Unicast Traffic: Storm Control can limit unicast flooding by dropping traffic that exceeds the configured threshold.
Unknown Unicast Traffic: Storm Control can limit unknown unicast flooding by dropping traffic that exceeds the configured
threshold.
Storm Control Percentage and Bits per second Options
Percentage: Enables administrators to specify a threshold as a percentage of the total available bandwidth of the interface,
for example, 10% of a 10Gbps link.
Bits per second (bps): Defines the threshold in bits per second, allowing granular control independent of the interface speed.
When traffic of a specified type exceeds the configured percentage or bps, Storm Control activates and drops excess packets,
preventing the storm from affecting the rest of the network.
Storm Control Behavior on Port Channel Interfaces and Counters
When Storm Control is applied to PortChannel interfaces (also known as EtherChannels), it behaves slightly differently compared
to individual physical interfaces. A PortChannel interface is a logical interface that aggregates multiple physical links.
When Storm Control is applied to an EtherChannel, the configuration propagates to the underlying physical interfaces (members).
Configuring Storm Control on Port Channels
Procedure
Command or Action
Purpose
Step 1
enable
Example:
Device> enable
Enables privileged EXEC mode.
Step 2
configure terminal
Example:
Device# configure terminal
Enters global configuration mode.
Step 3

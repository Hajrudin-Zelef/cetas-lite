---
id: collect-261001-cisco/cisco/enterprise-en-set-802-1p-for-vlan-thread-592702-861-0884f696
title: "enterprise-en-set-802-1p-for-vlan-thread-592702-861-0884f696"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-set-802-1p-for-vlan-thread-592702-861-0884f696.md
source_anchor: ""
source_lines: [1, 67]
sha256: 65910ba5110952538b4bfbd2213608ff9729e47e5da8f605b35f7bec14f8f6b5
---

# enterprise-en-set-802-1p-for-vlan-thread-592702-861-0884f696

Hi all,
This post describes how to set the 802.1p value of a new VLAN for packets in Layer 2 scenarios.
Problem Description
An S5300 switch is replaced with the NE40E-M2K at a site. The S5300 switch adds another VLAN tag to incoming packets with VLAN tags and sets the 802.1p value of the new VLAN tag.
A switch has the following two situations:
(1) For packets with VLAN tags, the 802.1p value is set to 5. An outer VLAN is added to the inbound interface, and the 802.1p values of the inner and outer VLANs are set to the same value.
The switch configuration is as follows:
acl number 4111
rule 5 permit vlan-id 111
rule 10 permit vlan-id 3111
#
interface Eth-Trunk11
qinq vlan-translation enable
undo port hybrid vlan 1
port hybrid tagged vlan 912 4004
port hybrid untagged vlan 3100 3111
port vlan-stacking vlan 100 stack-vlan 3100
port vlan-stacking vlan 111 stack-vlan 3111
port vlan-mapping vlan 912 map-vlan 912
port vlan-mapping vlan 4004 map-vlan 4004
stp disable
traffic-remark inbound acl 4111 rule 5 8021p 5
traffic-remark inbound acl 4111 rule 10 8021p 5
mode lacp
undo local-preference enable
lacp preempt enable
(2) For packets with VLAN tags, set the 802.1p value to 5, add an outer VLAN tag to the packets, set the 802.1p value of the new outer VLAN tag to 1, and set the 802.1p values of the inner and outer VLAN tags to different values.
interface GigabitEthernet0/0/12
undo negotiation auto
port hybrid tagged vlan 3000 4001
port vlan-stacking vlan 100 stack-vlan 3100 remark-8021p 1
port vlan-stacking vlan 111 stack-vlan 3111 remark-8021p 1
port vlan-mapping vlan 3000 map-vlan 3000
port vlan-mapping vlan 4001 map-vlan 4001
Problem Analysis
1. For the first case, you can set the 802.1p value of the current VLAN by complex traffic classification. The 802.1p value of the new VLAN inherits the configured 802.1p value. The following is a configuration example:
Configure complex traffic classification and re-mark the 802.1p value of the original packet.
traffic classifier vlan150 operator or
if-match any
traffic behavior vlan150
remark 8021p 5
traffic policy vlan150
share-mode
classifier vlan150 behavior vlan150 precedence 1
Add an outer VLAN 250 and configure the complex traffic classification matching the original VLAN 150.
interface GigabitEthernet0/3/8
portswitch
port vlan-stacking vlan 150 stack-vlan 250
traffic-policy vlan150 inbound vlan 150
8021p demapping does not need to be configured on the egress.
interface GigabitEthernet0/3/9
port trunk allow-pass vlan 250 650
Test result:
No traffic policy is configured on the inbound interface. Original packets enter BE queues, and the 802.1p value is 0.
Traffic policy is configured on the inbound interface. Remark the 802.1p value in the original VLAN tag to 5 and the 802.1p value in the new VLAN tag is also 5.
2. In the second scenario, VLANs are added in the inbound direction of the router after complex traffic classification is performed. To modify the 802.1p value of the new VLANs, complex traffic classification must be performed in the outbound direction.
However, there are two types of outbound interfaces. There is only one outbound interface, which transparently transmits the new VLAN tag. Whether the inner and outer VLAN IDs are the same cannot be determined on the outbound interface. The 802.1p values of the inner VLAN tag are 5, and the 802.1p values of the outer VLAN tag must be different.
Based on the product manual, a method is perform complex traffic classification to re-mark the 802.1p value of the original VLAN tag on the inbound interface, and then the packets are forced to enter the queue. The 802.1p priority demapping is configured on the outbound interface. The configuration is as follows:
Configure complex traffic classification, re-mark the 802.1p value of the original packet, and forcibly place the packet into the AF1 queue.
service-class af1 color green
port trunk allow-pass vlan 250
trust upstream default vlan 250
trust 8021p vlan 250
Configure a traffic policy on the inbound interface, re-mark the 802.1p value in the original VLAN tag to 5, and forcibly place packets into the AF1 queue. The 802.1p value in the new VLAN tag is 1.
That's all, welcome to study and discuss together!
For details about how to set 802.1p to take effect, see the following link.
MQC & where DSCP or 802.1p remarking take effect

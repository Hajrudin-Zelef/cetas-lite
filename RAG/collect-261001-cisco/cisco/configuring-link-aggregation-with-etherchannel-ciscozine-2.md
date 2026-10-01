---
id: collect-261001-cisco/cisco/configuring-link-aggregation-with-etherchannel-ciscozine-2
title: "configuring-link-aggregation-with-etherchannel-ciscozine"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-cisco/configuring-link-aggregation-with-etherchannel-ciscozine.md
source_anchor: ""
source_lines: [125, 143]
sha256: 956adf601cd22b9761ab45b128a5a949157f109ff7e8f3b7d72943c5446e7add
---

# configuring-link-aggregation-with-etherchannel-ciscozine

Which? 1600 Mbps (Fast EtherChannel, full duplex) or 16 Gbps (Gigabit EtherChannel)

1600 Mbps is 1.6 Gbps! Probably 16000 Mbps.

Is it possible to have one end configured as l2 while the other end as l3 port-channel? Does this kind of configuration (mix of l2&l3) work?

Many thanx

No, the interface on each sides must be the same: L2 – L2 or L3 – L3

Nice

I think Fabio means when using fast ethernet, you can have up to 1.6Gbps, but when using gigabit ethernet, you can have 16Gbps. Correct me if I am wrong, I am learninbg also :)

When PagP or LACP are ‘on’, no packet will be sent to make a channel. Indeed: “The link aggregation is forced to be formed without any LACP negotiation .In other words, the switch will neither send the LACP packet nor process any incoming LACP packet. This is similar to the on state for PAgP.”

Ugh! :) Please excuse my avatar. I must have been having a bad day!

So, if I have 2 cables linking 2 switches together, and the purpose I am aiming for is to double the bandwidth between the 2, should I be using channel-group [#] mode on, as opposed to mode active or passive, or does it make any difference? I am trying to get double the bandwidth – I am not aiming for fault tolerance. So I am trying to understand if the mode I am using has any effect on that. Like, let’s say I have 2 buildings connected by fiber. Let’s say each building has 2 stacked switches, so each stack is one logical switch. My understanding is, if I connect one cable between switch A in the stack (in building A), and the other end of this same cable to switch A in building B, and then connect the second cable to switch B in building A and the other end of this cable to switch B in building B, and then create a port-channel with mode active, this will give me fault-tolerance. But I don’t know if this will give me more bandwidth. So let’s say I want more bandwidth and don’t care about fault-tolerance. So if I connect both cables to switch A in both buildings, and use mode active, will that give me twice the bandwidth of what only one cable between the building would give me? Or do I need to use mode ON in the port-channel? Does the physical connection determine the outcome, or the mode used in the port-channel, or both? Many thanks!

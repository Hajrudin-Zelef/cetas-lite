---
id: collect-261001-fortinet/fortinet/questions-555536-fortigate-vpn-routing-issue-6311ba91
title: "questions-555536-fortigate-vpn-routing-issue-6311ba91"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["reasoning"]
source: docs/RAG/collect-261001-fortinet/questions-555536-fortigate-vpn-routing-issue-6311ba91.md
source_anchor: ""
source_lines: [1, 8]
sha256: 30b610f8c420920ca317866ad91789e7986074d2ac2e3c48f0f273bcf6fea043
---

# questions-555536-fortigate-vpn-routing-issue-6311ba91

You need to configure two phase 1s (and two phase 2s), one for each WAN interface on your 200B.  On the secondary/backup tunnel, configure monitor, as described in the Fortigate cookbook.  Reasoning is also there... to summarize, this allows a tunnel to monitor another tunnel and bring itself up when the other tunnel goes down (dead peer detection must also be enabled).  You might want to set the monitor-hold-delay to something fairly high, to allow you to follow up with your primary ISP and make sure that primary connection isn't flapping.  You can be alerted of the change by configuring email alerting, snmptrap monitoring, or use something like Gateway IP Monitor (I actually have all three configured).
Also, consider your routing needs.  Are both interfaces configured via DHCP or PPPoE (but with static addresses)?  Do you have ECMP and static routes?
I want to also make the suggestion of creating DNS failover, if you have an internal DNS server.  I've covered this in a blog post.
If you ever need to NAT your IPsec packets themselves (to an address other than that bound to the egress interface):
- use the Local Gateway Address for the NAT source address.
- enable the ability for two IPs in the same subnet to be bound to interfaces (overlapping).
- bind the additional IP to the interface.
Sorry to include this extra bit of info, but I had a hell of a time figuring it out.

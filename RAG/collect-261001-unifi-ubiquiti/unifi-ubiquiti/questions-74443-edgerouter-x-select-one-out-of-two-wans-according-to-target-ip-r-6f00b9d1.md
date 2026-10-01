---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-74443-edgerouter-x-select-one-out-of-two-wans-according-to-target-ip-r-6f00b9d1
title: "questions-74443-edgerouter-x-select-one-out-of-two-wans-according-to-target-ip-r-6f00b9d1"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-74443-edgerouter-x-select-one-out-of-two-wans-according-to-target-ip-r-6f00b9d1.md
source_anchor: ""
source_lines: [1, 19]
sha256: bb68ba4fc99facb6c58b7712d9bf217afbcbb4e27e10e6a510034ab82b0103b0
---

# questions-74443-edgerouter-x-select-one-out-of-two-wans-according-to-target-ip-r-6f00b9d1

I am using IPSET to define the interested IP range.
I want to direct packets to one out of two WAN connections according to an IP range. Here are my conditions.
- A specific IP range. It is defined by using IPSET command to modify the underlying object for the network-group function.
- WAN connection, using DHCP, on eth4.
- Another WAN connection, using DHCP, on eth3.
- ER-X router, v2.09
By default, IP packets should use eth3 wan connection unless the destination IP address lied in the interested IP range. If so, packets should use eth4.
I have tried the following way, but it does not work. All packets always use eth3 unless I pulled out the eth3 connector.
First, create a network-group named "specialroute" and then use IPSET to import the interested IP range (for packets destination addresses to match).
Then create a static route table, representing where matched packets should go.
set protocols static table 13 interface-route 0.0.0.0/0 next-hop-interface eth4
Then mark eth4 with the failover-only flag to direct packets to eth3 by default.
However, some packets with the destination in the pre-defined IP range should go to eth4. So, I tried using Policy-Based-Routing to achieve the goal.
set firewall modify PBR rule 90 action modify
set firewall modify PBR rule 90 description specialroute
set firewall modify PBR rule 90 destination group network-group specialroute
set firewall modify PBR rule 90 modify table 13
It did not work as the failover-only flag might prevail PBR rules, and the router never sent packets from eth4.
What should I do? I want to let packets use one WAN connection by default unless the destination address belongs to the pre-defined IP ranges.

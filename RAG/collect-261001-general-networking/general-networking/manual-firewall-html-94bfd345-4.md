---
id: collect-261001-general-networking/general-networking/manual-firewall-html-94bfd345-4
title: "manual-firewall-html-94bfd345"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/manual-firewall-html-94bfd345.md
source_anchor: ""
source_lines: [248, 296]
sha256: c2a52ab54ef0a8f923e9317fddaa828d100bf664e52b71c4497197dacce3eb1c
---

# manual-firewall-html-94bfd345

| Direction | Direction of the traffic, see also Direction. | 
The following options are specifically used for HA setups.
| Option | Description | 
|---|---|
| No XMLRPC Sync | Disable configuration sync for this rule, when Firewall Rules sync is enabled in | 
| State Type / NO pfsync | Prevent states created by this rule to be synced to the other node | 
Rules can also be scheduled to be active at specific days or time ranges, you can create schedules in and select one in the rule.
This feature can be used to forward traffic to another gateway based on more fine grained filters than static routes could (OSI layer 4 verses OSI layer 3) and can be used to build multi-wan scenario’s using gateway groups.
More information about Multi-Wan can be found in the “Multi WAN” chapter.
| Option | Description | 
|---|---|
| Gateway | When a gateway is specified, packets will use policy based routing using the specified gateway or gateway group. Usually this option is set on the receiving interface (LAN for example), which then chooses the gateway specified here. (This ignores default routing rules). Only packets flowing in the same direction of the rule are affected by this parameter, the opposite direction (replies) are not affected by this option. | 
| reply-to | By default traffic is always send to the connected gateway on the interface. If for some reason you don’t want to force traffic to that gateway, you can disable this behaviour or enforce an alternative target here. | 
Note
When using policy based routing, don’t forget to exclude local traffic which shouldn’t be forwarded.
You can do so by creating a rule with a higher priority, using a default gateway.
Tip
In our experience the packet capture function () can be a valuable tool to inspect if traffic is really heading the direction you would expect it to go, just choose a host to monitor and try to exchange some packets. When selecting all interfaces, it’s easy to see where traffic headed.
The advanced options contains some settings to limit the use of a rule or specify specific timeouts for the it. Most generic (default) settings for these options can be found under
| Option | Description | 
|---|---|
| Max states | Limits the number of concurrent states the rule may create. When this limit is reached, further packets that would create state will not match this rule until existing states time out. | 
| Max source nodes | Limits the maximum number of source addresses which can simultaneously have state table entries. | 
| Max established | Limits the maximum number of simultaneous TCP connections which have completed the 3-way handshake that a single host can make. | 
| Max source states | Limits the maximum number of simultaneous state entries that a single source address can create with this rule. | 
| Max new connections | Limit the rate of new connections over a time interval. The connection rate is an approximation calculated as a moving average. (number of connections / seconds) Only applies on TCP connections | 
| State timeout | State Timeout in seconds (applies to TCP only) | 
Some less common used options are defined below.
| Option | Description | 
|---|---|
| Source OS | Operating systems can be fingerprinted based on some tcp fields from the originating connection. These fingerprints can be used as well to match traffic on. (more detailed information can be found in the pf.os man page) | 
| allow options | By default the firewall blocks IPv4 packets with IP options or IPv6 packets with routing extension headers set. If you have an application that requires such packets (such as multicast or IGMP) you can enable this option. | 
| TCP flags | If specific TCP flags need to be set or unset, you can specify those here. | 
| Set priority | Packets matching this rule will be assigned a specific queueing priority. If the packet is transmitted on a VLAN interface, the queueing priority will be written as the priority code point in the 802.1Q VLAN header. If two priorities are given, packets which have a TOS of lowdelay and TCP ACKs with no data payload will be assigned to the second one. | 
| Match priority | Only match packets which have the given queueing priority assigned. | 
| Set local tag | Packets matching this rule will be tagged with the specified string. The tag acts as an internal marker that can be used to identify these packets later on. This can be used, for example, to provide trust between interfaces and to determine if packets have been processed by translation rules. Tags are “sticky”, meaning that the packet will be tagged even if the rule is not the last matching rule. Further matching rules can replace the tag with a new one but will not remove a previously applied tag. A packet is only ever assigned one tag at a time. | 
| Match local tag | Match packets that are tagged earlier (using set local tag) | 
| State Type | Influence the state tracking mechanism used, the following options are available. When in doubt, it’s usually best to preserve the default keep state  | 
Troubleshooting
While building your ruleset things can go wrong, it’s always good to know where to look for signs of an issue. One of the most common mistakes is traffic doesn’t match the rule and/or the order of the rule doesn’t make sense for whatever reason.
With the use of the “inspect” button, one can easily see if a rule is being evaluated and traffic did pass using this rule. It’s also possible to jump directly into the attached states to see if your host is in the list as expected.
Another valuable tool is the live log viewer, in order to use it, make sure to provide your rule with an easy to read description and enable the “log” option.
If your using source routing (policy based routing), debugging can sometimes get a bit more complicated. Since the normal system routing table may not apply, it helps to know which flow the traffic actually followed. The packet capture is a useful tool in that case.
Common issues in this area include return traffic using a different interface than the one it came into, since traffic follows the normal routing table on it’s way out (reply-to issue), or traffic leaving the wrong interface due to overselection (matching internal traffic and forcing a gateway).
Inspecting used netmasks is also a good idea, intending to match a host but providing a subnet is a mistake easily made
(e.g. 192.168.1.1/32 vs 192.168.1.1/24 is in reality all of 192.168.1.x).
Last but not least, remember rules are matched in order and the default (inbound) policy is block if nothing else
is specified, since we match traffic on inbound, make sure to add rules where traffic originates from
(e.g. lan for traffic leaving your network, the return should normally be allowed by state).

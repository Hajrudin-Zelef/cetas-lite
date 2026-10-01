---
id: collect-261001-general-networking/general-networking/manual-firewall-html-94bfd345-2
title: "manual-firewall-html-94bfd345"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/manual-firewall-html-94bfd345.md
source_anchor: ""
source_lines: [90, 165]
sha256: 1667d42811c403ffc1079b846a2bbd0aa56b509846b5e30283252edba9b4dc2e
---

# manual-firewall-html-94bfd345

API access is described in more detail in the firewall API reference manual.
Interface filter
Choose an interface to filter the current view. Floating, group and single interfaces can be selected.
If you choose the “LAN” interface, you will be presented with all floating, group and single interface rules that influence packet decisions of the LAN interface.
If you create a new rule while having an interface selected, it will be automatically added to dialog.
Categories filter
Choose one or multiple categories to filter the current view. This combines with the selection of the interface filter.
Categories can be created in and can enable grouping different logic constructs.
If you create a category for mailservers and tag rules with it, you can simply filter for this tag and only see your mailservers. As with the interface filter, selecting one or multiple tags will add them automatically to a new rule.
Settings
| Option | Description | 
|---|---|
| Enabled | Enable this rule | 
| Sort order | The order in which rules are being processed. | 
| Sequence | The order in which rules are being processed. Please note that this is not a unique identifier, the system will automatically recalculate the ruleset when rule positions are changed with the available “Move rule before this rule” button. | 
| Categories | For grouping purposes you may select multiple groups here to organize items. | 
| No XMLRPC Sync | Exclude this item from the HA synchronization process. An already existing item with the same UUID on the synchronization target will not be altered or deleted as long as this is active. | 
| Description | You may enter a description here for your reference (not parsed). | 
| Option | Description | 
|---|---|
| Invert interface (rule) | Match packets on all but the selected rule interface. | 
| Interface (rule) | Only match packets on the selected rule interfaces or groups. | 
| Invert interface (origin) | Match packets that were not received on the selected interface. | 
| Interface (origin) | Only match packets that were initially received on the selected interfaces or groups. This is mostly relevant for out direction rules when creating a security zone ruleset. | 
| Option | Description | 
|---|---|
| Quick | If a packet matches a rule specifying quick, then that rule is considered the last matching rule and the specified action is taken. When a rule does not have quick enabled, the last matching rule wins. | 
| Action | Choose what to do with packets that match the criteria specified below. Hint: the difference between block and reject is that with reject, a packet (TCP RST or ICMP port unreachable for UDP) is returned to the sender, whereas with block the packet is dropped silently. In either case, the original packet is discarded. | 
| Allow options | This allows packets with IP options to pass. Otherwise they are blocked by default. | 
| Direction | Direction of the traffic. The default policy is to filter inbound traffic, which sets the policy to the interface originally receiving the traffic. | 
| Version | The IP protocol version that should match, e.g., IPv4 or IPv6. | 
| Protocol | The transport protocol that should match, e.g., TCP or UDP. | 
| ICMP type | If the transport protocol is ICMP, this option allows you to specify the ICMP types. | 
| ICMPv6 type | If the transport protocol is IPV6-ICMP, this option allows you to specify the ICMP types. | 
| Invert Source | Use this option to invert the sense of the match. | 
| Source | The source IP address or alias that should match. | 
| Source Port | Source port number or well known name (imap, imaps, http, https, …), for ranges use a dash | 
| Invert Destination | Use this option to invert the sense of the match. | 
| Destination | The destination IP address or alias that should match. | 
| Destination Port | Destination port number or well known name (imap, imaps, http, https, …), for ranges use a dash | 
| Log | Log packets that are handled by this rule | 
| TCP flags | Use this to choose TCP flags that must be set this rule to match. | 
| TCP flags [out of] | Use this to choose TCP flags that must be cleared for this rule to match. | 
| TCP flags any | Match any combination of TCP flags. | 
| Schedule | Rules can also be scheduled to be active at specific days or time ranges, you can create schedules in and select one in the rule. If the rule times out the states will be removed and the rule will be skipped. This means, if there is still a matching rule after the scheduled rule that allows the traffic, it will be used instead. Keep this in mind when using scheduled rules, and carefully build the ruleset around them, e.g., with additional block rules. | 
| Divert-to | Send packets matching this rule to the service specified, when the service is not running, packets will be dropped. | 
| Option | Description | 
|---|---|
| State type | State tracking mechanism to use, default is full stateful tracking, sloppy ignores sequence numbers, use none for stateless rules. | 
| State policy | Choose how states created by this rule are treated, default (as defined in advanced), floating in which case states are valid on all interfaces or interface bound. Interface bound states are more secure, floating more flexible. | 
| NO pfsync | This prevents states created by this rule to be synced with pfsync. | 
| TCP established | State Timeout in seconds (TCP only) | 
| UDP first | The state timeout in seconds after the first UDP packet. | 
| UDP single | The state timeout in seconds if both hosts have sent UDP packets. | 
| UDP multiple | The state timeout in seconds if the source host sends more than one UDP packet but the destination host has never sent one back. | 
| Adaptive Timeouts [start] | When the number of state entries exceeds this value, adaptive scaling begins. All timeout values are scaled linearly with factor (adaptive.end - number of states) / (adaptive.end - adaptive.start). | 
| Adaptive Timeouts [end] | When reaching this number of state entries, all timeout values become zero, effectively purging all state entries immediately. This value is used to define the scale factor, it should not actually be reached (set a lower state limit). | 
| Max states | Limits the number of concurrent states the rule may create. When this limit is reached, further packets that would create state are dropped until existing states time out. | 
| Max source nodes | Limits the maximum number of source addresses which can simultaneously have state table entries. | 
| Max source states | Limits the maximum number of simultaneous state entries that a single source address can create with this rule. | 
| Max source connections | Limit the maximum number of simultaneous TCP connections which have completed the 3-way handshake that a single host can make. | 
| Max new connections [c] | Maximum new connections per host, measured over time. | 
| Max new connections [s] | Time interval (seconds) to measure the number of connections | 
| Overload table | Overload table used when max new connections per time interval has been reached. The default virusprot table comes with a default block rule in floating rules, alternatively specify your own table here. | 
| Option | Description | 
|---|---|
| Max packet rate [packets] | Maximum number of packets allowed during the configured time interval. | 
| Max packet rate [seconds] | Time interval in seconds over which the maximum packet rate is measured. | 
| Option | Description | 
|---|---|
| Traffic shaper | Shape packets using the selected pipe or queue in the rule direction. | 
| Traffic shaper [reverse] | Shape packets using the selected pipe or queue in the reverse rule direction. | 
| Option | Description | 
|---|---|
| Gateway | Leave as ‘default’ to use the system routing table. Or choose a gateway to utilize policy based routing. | 
| Disable reply-to | Explicit disable reply-to for this rule | 

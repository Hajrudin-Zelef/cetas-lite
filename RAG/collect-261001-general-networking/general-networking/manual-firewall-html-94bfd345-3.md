---
id: collect-261001-general-networking/general-networking/manual-firewall-html-94bfd345-3
title: "manual-firewall-html-94bfd345"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-general-networking/manual-firewall-html-94bfd345.md
source_anchor: ""
source_lines: [166, 247]
sha256: 77e7801f014a285bd5daa097ef82ad6280722fd73b2a7caab085ccf23259dc56
---

# manual-firewall-html-94bfd345

| Reply-to | Determines how packets route back in the opposite direction (replies), when set to default, packets on WAN type interfaces reply to their connected gateway on the interface (unless globally disabled). A specific gateway may be chosen as well here. This setting is only relevant in the context of a state, for stateless rules there is no defined opposite direction. | 
| Option | Description | 
|---|---|
| Match priority | Only match packets which have the given queueing priority assigned. | 
| Set priority | Packets matching this rule will be assigned a specific queueing priority. If the packet is transmitted on a vlan(4) interface, the queueing priority will be written as the priority code point in the 802.1Q VLAN header | 
| Set priority [low-delay] | Used in combination with set priority, packets which have a TOS of lowdelay and TCP ACKs with no data payload will be assigned this priority when offered. | 
| Match TOS / DSCP | Only match packets which have the given TOS/DSCP marker. | 
| Option | Description | 
|---|---|
| Set local tag | Packets matching this rule will be tagged with the specified string. The tag acts as an internal marker that can be used to identify these packets later on. This can be used, for example, to provide trust between interfaces and to determine if packets have been processed by translation rules. Tags are “sticky”, meaning that the packet will be tagged even if the rule is not the last matching rule. Further matching rules can replace the tag with a new one but will not remove a previously applied tag. A packet is only ever assigned one tag at a time. | 
| Match local tag | Used to specify that packets must already be tagged with the given tag in order to match the rule. | 
| Option | Description | 
|---|---|
| Created: Username | The user who created the rule. | 
| Created: Time | The time at which the rule was created. | 
| Created: Description | Additional information about how the rule was created. | 
| Updated: Username | The user who last updated the rule. | 
| Updated: Time | The time at which the rule was last updated. | 
| Updated: Description | Additional information about how the rule was last updated. | 
| Note | Add a note to describe why this rule was created or changed. | 
Divert-to
With divert we can integrate additional policy constructs into firewall rules. A rule with “divert-to” can send all packets to a divert socket, on which a service like Intrusion Prevention System can listen.
This can provide more fine grained control than intercepting all traffic for inspection. It can improve performance because large flows that do not necessarily need inspection can be excluded.
As an example, go to and set the “Capture mode” to “Divert (IPS)”.
Afterwards, create a firewall rule that matches the traffic that should be inspected and set the “Divert-to” settings in that rule. Matching packets will be diverted to the socket the intrusion detection listens on. After making a decision, the packet will then be forwarded or dropped.
Attention
Keep in mind that when the service that listens on the divert socket is stopped, the firewall rule will drop all matching packets.
Rules
Rules is the implementation that has been around since day one. Since it consists of static php pages, there is no API support, Over time, it will be replaced by Rules [new] and a Migration assistant can be found in .
User Interface
Our overview shows all the rules that apply to the selected interface (group) or floating section. For every rule some details are provided and when applicable you can perform actions, such as move, edit, copy, delete.
Below you will find some highlights about this screen.
  - Interface name
  - The name of the interface is part of the normal menu breadcrumb
  - Category
  - If categories are used in the rules, you can select which one you will show here.
  - Toggle inspection
  - You can toggle between inspection and rule view here, when in inspection mode, statistics of the rule are shown. (such as packet counters, number of active states, …)
  - Show / hide automatic rules
  - Some rules are automatically generated, you can toggle here to show the details. If a magnifying glass is shown you can also browse to its origin (The setting controlling this rule).
  - Automatic rules
  - The contents of the automatic rules
  - User rules
  - All user defined rules
Settings
Traffic that is flowing through your firewall can be allowed or denied using rules, which define policies. This section of the documentation describe the different settings, grouped by usage.
Some settings help to identify rules, without influencing traffic flow.
| Option | Description | 
|---|---|
| Category | The category this rule belongs to, can be used as a filter in the overview | 
| Description | Descriptive text | 
Below are the settings most commonly used:
| Option | Description | 
|---|---|
| Action | The action to perform. | 
| Disabled | Disable a rule without removing it, can be practical for testing purposes and to support easy enablement of less frequently used policies. | 
| Interface | Interface[s] this rule applies on. You can easily copy rules between interfaces and change this field to the new target interface. (remember to check the order before applying) | 
| TCP/IP Version | Does this rule apply on IPv4, IPv6 or both. | 
| Protocol | Protocol to use, most common are TCP and UDP | 
| Source | Source network or address, when combining IPv4 and IPv6 in one rule, you can use aliases which contain both address families. You can select multiple sources per rule. | 
| Source / Invert | Invert source selection (for example not 192.168.0.0/24) You can only invert single sources. | 
| Destination | Destination network or address, like source you can use aliases here as well. You can select multiple destinations per rule. | 
| Destination / Invert | When the filter should be inverted, you can mark this checkbox. You can only invert single destinations. | 
| Destination port range | For TCP and/or UDP you can select a service by name (http, https) or number (range), you can also use aliases here to simplify management. | 
| Log | Create a log entry when this rule applies, you can use to monitor if your rule applies. | 
Tip
The use of descriptive names help identify traffic in the live log view easily.
Tip
You can select multiple sources or destinations per rule, yet keep in mind that a nested alias might be the better choice. This feature is most useful if you plan to create security zones.
When a firewall rule needs to be constrained in terms of the number of packets it may process over time, it’s possible to combine the rule with the traffic shaper.
The process of shaping is explained in the Traffic Shaping section of our documentation. Below you will find the relevant properties for the firewall rule.
| Option | Description | 
|---|---|
| Traffic shaping/rule direction | Force packets being matched by this rule into the configured queue or pipe | 
| Traffic shaping/reverse direction | Force packets being matched in the opposite direction into the configured queue or pipe | 
Tip
Filter rules are more flexible than the ones specified in the shaper section itself as these can be combined with aliases as well. Although this feature is quite new, it’s certainly worth looking at when in need of a traffic shaper.
Some settings are usually best left default, but can also be set in the normal rule configuration.
| Option | Description | 
|---|---|
| Source port range | In case of TCP and/or UDP, you can also filter on the source port (range) that is used by the client. Since in most cases you can’t influence the source port, this setting is usually kept default (any ). | 
| Quick | If a packet matches a rule specifying quick, the first matching rule wins. When not set to quick the last matching rule wins. When not sure, best use quick rules and interpret the ruleset from top to bottom. | 

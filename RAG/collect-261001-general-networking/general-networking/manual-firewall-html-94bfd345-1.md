---
id: collect-261001-general-networking/general-networking/manual-firewall-html-94bfd345-1
title: "manual-firewall-html-94bfd345"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-general-networking/manual-firewall-html-94bfd345.md
source_anchor: ""
source_lines: [1, 89]
sha256: eb5daedd0f857262e87dcd50cddd12358c9bfc06d1b6435334cc360a6b2eab59
---

# manual-firewall-html-94bfd345

Rules
Index
Overview
OPNsense contains a stateful packet filter, which can be used to restrict or allow traffic from and/or to specific networks as well as influence how traffic should be forwarded (see also policy based routing in “Multi WAN”).
The rules section shows all policies that apply on your network, grouped by interface.
There are two implementations to choose from:
- Rules [new]: a modern MVC implementation with API support and improved rule management
- Rules: a static PHP page with no API support
Tip
Rules [new] will replace Rules over time, you can already migrate your existing rules with a helper in .
The basics
Before creating rules, it’s good to know about some basics which apply to all rules.
States
By default rules are set to stateful (you can change this, but it has consequences), which means that the state of a connection is saved into a local dictionary which will be resolved when the next packet comes in. The consequence of this is that when a state exists, the firewall doesn’t need to process all its rules again to determine the action to apply, which has huge performance advantages.
Another advantage of stateful packet filtering is that you only need to allow traffic in one direction to automatically allow related packets for the same flow back in. Below diagram shows a tcp connection from a client to a server for https traffic, when not using stateful rules, both the client should be permitted to send traffic to the server at port 443 as the server back to the client (usually a port >=1024).
The use of states can also improve security particularly in case of tcp type traffic, since packet sequence numbers and timestamps are also checked in order to pass traffic, it’s much harder to spoof traffic.
Note
When changing rules, sometimes its necessary to reset states to assure the new policies are used for existing traffic. You can do this in .
Note
In order to keep states, the system need to reserve memory. By default 10% of the system memory is reserved for states, this can be configured in . (The help text shows the default number of states on your platform)
States can also be quite convenient to find the active top users on your firewall at any time, we added an easy to use “session” browser for this purpose. You can find it under .
Tip
States also play an important rule into protecting services against (distributed) denial of service attacks (DDOS). Relevant topics available in our documentation are “synproxy” states, connection limits and syncookies
Action
Rules can be set to three different action types:
- Pass –> allow traffic
- Block –> deny traffic and don’t let the client know it has been dropped (which is usually advisable for untrusted networks)
- Reject –> deny traffic and let the client know about it. (only tcp and udp support rejecting packets, which in case of TCP means a RST is returned, for UDPICMP UNREACHABLE is returned).
For internal networks it can be practical to use reject, so the client does not have to wait for a time-out when access is not allowed. When receiving packets from untrusted networks, you usually don’t want to communicate back if traffic is not allowed.
Processing order
Firewall rules are processed in sequence per section, first evaluating the Floating rules section followed by all rules which belong to interface groups and finally all interface rules.
Internal (automatic) rules are usually registered first.
Rules can either be set to quick or not set to quick, the default is to use quick. When set to quick, the rule is
handled on “first match” basis, which means that the first rule matching the packet will take precedence over rules following in sequence.
When quick is not set, last match wins. This can be useful for rules which define standard behaviour.
Our default deny rule uses this property for example (if no rule applies, drop traffic).
Note
Internally rules are registered using a priority, floating uses 200000,
groups use 300000 and interface rules land on 400000 combined with the order in which they appear.
Automatic rules are usually registered at a higher priority (lower number).
Warning
NAT rules are always processed before filter rules! So for example, if you define a NAT : Destination NAT (Port Forwarding) rules without a associated rule, i.e. Filter rule association set to Pass, this has the consequence, that no other rules will apply!
Tip
The interface should show all rules that are used, when in doubt, you can always inspect the raw output of the ruleset in /tmp/rules.debug
Since and implementations exist side by side, there are some additional considerations regarding the processing order of rules.
If a filter rule has:
a single interface defined, it is an Interface Rule
a group interface defined, it is a Group Rule
any number of interfaces or one inverted interface defined, it is a Floating Rule
Processing order:
System defined rules at the beginning of the ruleset
and floating rules
and group rules
single interface rules
single interface rules
System defined rules at the end of the ruleset
Rule sequence
The sequence in which the rules are displayed and processed can be customized per section:
- Select a rule with the checkbox on the left side of the rule.
- Use the arrow button in the action menu on the right side of a rule in order to move selected rules before the rule where the action button is pressed.
While rules in are processed implicitly by the order they appear in the configuration file, implement a more explicit Sort order.
A Sort order will have this structure: 200000.0000250:
200000 The Priority group, defines the hierarchy the rule belongs to. It can be influenced by the interface in a rule.
.0000250 The Sequence, it defines the exact spot of the rule in context of the interface hierarchy. It can be influenced with the Sequence field inside a rule, or with the Move rule before this rule button.
As example, we create a few different rules:
200000.0000100 (Floating rule)
200000.0000200 (Floating rule)
300000.0000050 (Group rule)
300000.0000060 (Group rule)
400000.0002000 (Interface rule)
400000.0003100 (Interface rule)
When applying the filter, the rules will be processed in this Sort order. A lower sequence number inside the priority group means the priority is higher.
Note
The Sequence does not have to be unique, multiple rules can share the same number. Though what this means is that the filter is not populated as strictly in order as if all rules have a unique sequence.
Tip
When adding a new rule, the Sequence will be automatically populated with a unique number. You can either change this number manually to adjust the exact position of the rule, or use the Move rule before this rule button.
Direction
Traffic can be matched on in[coming] or out[going]  direction, our default is to filter on incoming direction.
In which case you would set the policy on the interface where the traffic originates from.
For example, if you want to allow https traffic coming from any host on the internet,
you would usually set a policy on the WAN interface allowing port 443 to the host in question.
Note
Traffic leaving the firewall is accepted by default (using a non-quick rule), when Disable force gateway in is not checked, the connected gateway would be enforced as well.
Implementations
Rules [new]
Rules [new] started out as Firewall Automation with a limited GUI to offer API support, since the static php implementation of Rules lacked automation requirements. Over time, more features have been added, and the GUI has been completely reworked. Long requested features like advanced search functionality, category folders improved performance and better visibility are the result.
You can find it in or .
User Interface
Our overview shows all the rules that apply to the selected interface, group or floating section. For every rule, some details are provided and when applicable you can perform actions such as move, edit, copy, delete.

---
id: collect-261001-general-networking/general-networking/manual-nat-html-52be0bd7-2
title: "manual-nat-html-52be0bd7"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/manual-nat-html-52be0bd7.md
source_anchor: ""
source_lines: [102, 167]
sha256: 956e659ef3acc8e3126c3ae0b59cceff5cac3221e1355390b98c4d91c19f05fb
---

# manual-nat-html-52be0bd7

If you only have one external IP, the automatic mode is suitable in most cases. With multiple external addresses, you may want to use hybrid or manual mode and add custom rules.
| Automatic Source NAT rule generation | The default. Follows the behaviour described above, and is suitable for most scenarios. | 
| Manual Source NAT rule generation | No automatic rules are generated. They can be added manually. | 
| Hybrid Source NAT rule generation | Automatic rules are generated, and additional manual rules can be added. | 
| Disable Source NAT rule generation | Disables Source NAT. This is used for transparent bridges, for example. | 
When adding a rule, the following fields are available:
| Option | Description | 
|---|---|
| Enable | Enable this rule | 
| Sequence | Rules are evaluated in sequence order. | 
| Categories | Assign categories for rule organization. | 
| No XMLRPC Sync | Exclude this rule from HA synchronization. | 
| Description | Enter a description to identify this rule. | 
| Option | Description | 
|---|---|
| Interface | Choose the interface(s) on which the traffic originates. | 
| Version | Select IPv4, IPv6 or both. | 
| Protocol | Assign categories for rule organization. | 
| Description | Select the protocol this rule should match. | 
| Option | Description | 
|---|---|
| Invert Source | Match everything except the specified source. | 
| Source Address | Specify the source network or alias to match. | 
| Source Port | Source port or port range. | 
| Option | Description | 
|---|---|
| Invert Destination | Match everything except the specified destination. | 
| Destination Address | Destination address or alias to match. | 
| Destination Port | Destination port or port range. | 
Tip
This translates the original source (e.g., an internal host) to a new source (e.g., the external IP address of the firewall).
| Option | Description | 
|---|---|
| Translate Source IP | Packets matching this rule will be mapped to the IP address given here. | 
| Translate Source Port | Source port number or well-known name. | 
| Pool Options | Choose how traffic is distributed between multiple translation addresses. | 
| Source Hash Key | Keep source-hash mappings stable across ruleset reloads. | 
| Static-port | Prevent changes to the source port of TCP and UDP packets. | 
| Endpoint Independent | Enable endpoint-independent mapping for UDP traffic. | 
| Option | Description | 
|---|---|
| Do not NAT | Enabling this option will disable NAT for traffic matching this rule and stop processing Source NAT rules. | 
| Log | Log packets that are handled by this rule. | 
| Set local tag | Assign a tag that other NAT or filter rules can match. | 
| Match local tag | Used to specify that packets must already be tagged with the given tag in order to match the rule. | 
NPTv6
Network Prefix Translation, shortened to NPTv6, is used to translate IPv6 addresses. A common usage for this is to translate global (“WAN”) IPs to local ones. In this regard, it is similar to NAT, although NPTv6 can only be used to map addresses one-to-one, unlike NAT which typically translates one external IP to several internal ones.
NPTv6 routes are listed at .
When adding a rule, the following fields are available:
| Option | Description | 
|---|---|
| Enable | Enable this rule | 
| Sequence | Rules are evaluated in sequence order. | 
| Categories | Assign categories for rule organization. | 
| Description | Enter a description to identify this rule. | 
| Option | Description | 
|---|---|
| Interface | Choose which interface this rule applies to. | 
| Option | Description | 
|---|---|
| Internal IPv6 Prefix | Enter the internal IPv6 prefix (source) for this network prefix translation. | 
| External IPv6 Prefix | Enter the external IPv6 prefix (target) for this network prefix translation. Leave empty to auto-detect the prefix address using the specified tracking interface instead. The prefix size specified for the internal prefix will also be applied to the external prefix. | 
| Track interface | Use prefix defined on the selected interface instead of the interface this rule applies to when target prefix is not provided. | 
| Option | Description | 
|---|---|
| Log | Log packets that are handled by this rule. |

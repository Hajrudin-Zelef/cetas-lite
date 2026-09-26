---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-47-1
title: "Common Actions and Associated properties"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-47.md
source_anchor: ""
source_lines: [1, 128]
sha256: fe7ab70a9f7e98c77cf7a08f9ac856f978c1a1d2403554e2bfe0f895f561bdc4
---

# Common Actions and Associated properties

| Property | Description | 
|---|---|
| **action** (*action name* ; Default:**accept** ) | Action to take if a packet is matched by the rule:  | 
| **address-list** (*name* ; Default: ) | Name of the address list to be used. Applicable if action is add-dst- to-address-list or add-src-to-address-list | 
| **address-list-timeout** (*none-dynamic \| none-static \| time* ; Default:**none-dynamic** ) | Time interval after which the address will be removed from the address list specified by `address-list` parameter. Used in conjunction with`add-dst-to-address-list` or`add-src-to-address-list` actions | 
| **jump-target** (*name* ; Default: ) | Name of the target chain to jump to. Applicable only if `action=jump` | 
| **log**  (*yes \| no; Default:* **no** ) | Add a message to the system log containing the following data: in-interface, out-interface, src-mac, protocol, src-ip:port->dst-ip:port, and length of the packet. Allows to log packets even if action is not " **log** ", useful for debugging firewall. | 
| **log-prefix** (*string* ; Default: ) | Adds specified text at the beginning of every log message. Applicable if or*action=log* configured.*log=yes* | 

## Stats

To view matching statistics by firewall rules, run `/ip firewall filter print stats` command or `/ipv6 firewall filter print stats` for IPv6 firewall.

| Property | Description | 
|---|---|
| **bytes** (*integer* ) | The total amount of bytes matched by the rule | 
| **packets** (*integer* ) | The total amount of packets matched by the rule | 

Statistics parameters can be reset by following commands:

| Command | Description | 
|---|---|
| **reset-counters** (*id* ) | Reset statistics counters for specific firewall rule or list of rules. | 
| **reset-counters-all** | Reset statistics counters for all firewall rules in the table. | 

## Other Useful Commands

By default print is equivalent to `print static` and shows only static rules. 

To print also dynamic rules use `print all`.

Or to print only dynamic rules use `print dynamic`.

# Matchers

Tables below shows all the properties that can be used as a matchers in the firewall rules.

Matchers are executed in a specific order.

For IPv4:

- Source MAC Address
- In/Out interfaces
- In/Out interface lists
- IP Range
- Address type
- Address list
- TTL
- DSCP
- Length
- TLS
- IPv4 Options
- Dst Port
- Src Port
- Any Port
- TCP Options
- TCP MSS
- ICMP Codes
- Ingress Priority
- Priority
- Packet Mark
- Realm (routing table)
- Hotsopot
- Connection Mark
- Connection State
- Connection NAT State
- Connection Bytes
- Connection Limit
- Connection Rate
- Ipsec Policy
- Helper
- String (content)
- PSD
- Layer7
- Random
- Nth
- PCC
- Limit
- Dst Limit
- Log

For IPv6:

- Address type
- Address list
- Source MAC Address
- In/Out interfaces
- In/Out interface lists
- Hop Limit
- DSCP
- Length
- TLS
- IPv6 Header
- Dst Port
- Src Port
- Any Port
- TCP Options
- TCP MSS
- ICMPv6 Codes
- Ingress Priority
- Priority
- Packet Mark
- Connection Mark
- Connection State
- Connection NAT State
- Connection Bytes
- Connection Limit
- Connection Rate
- Ipsec Policy
- Helper
- Match String (content)
- Random
- Nth
- PCC
- Limit
- Dst Limit
- Log


Properties are split in two parts:

- **stateless** - properties do not require connection tracking to function and can be used in stateless RAW firewall matching.
- **stateful** - properties either require connection tracking to function or is available only in stateful firewall config.

## Stateless Properties


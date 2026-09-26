---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-45
title: "Properties"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-45.md
source_anchor: ""
source_lines: [1, 23]
sha256: 8768fedb496acaa805bdbd11f17a349a65e25eee843bfedc57081d21c1584712
---

# Properties

Firewall address lists allow a user to create lists of IP addresses grouped together under a common name. Firewall filter, mangle, and NAT facilities can then use those address lists to match packets against them.

The address list records can also be updated dynamically via the `action=add-src-to-address-list` or `action=add-dst-to-address-list` items found in NAT, Mangle, and Filter facilities.

Firewall rules with action `add-src-to-address-list` or `add-dst-to-address-list` work in passthrough mode, which means that the matched packets will be passed to the next firewall rules.

# Properties

| Property | Description | 
|---|---|
| **address** (*DNS Name \| IP address/netmask \| IP-IP* ; Default: ) | A single IP address or range of IPs to add to the address list or DNS name. You can input for example, '192.168.0.0-192.168.1.255' and it will auto modify the typed entry to 192.168.0.0/23 on saving. IP-IP ranges are supported only for IPv4 addresses. | 
| **dynamic** (*yes,* no) | Allows creating data entry with dynamic form. | 
| **list** (*string* ; Default: ) | Name for the address list of the added IP address. | 
| **timeout** (*time* ; Default: ) | Time after address will be removed from the address list. If the timeout is not specified, the address will be stored in the address list permanently. | 
| **creation-time**  (*time* ; Default: ) | The time when the entry was created. | 

If the timeout parameter is not specified, then the address will be saved to the list permanently on the disk. If a timeout is specified, the address will be stored on the RAM and will be removed after a system's reboot.

Example

The following example creates a dynamic address list of people who are connecting to port 23 (telnet) on the router and drops all further traffic from them for 5 minutes. Additionally, the address list will also contain one static address list entry of 192.0.34.166/32 (www.example.com):

As seen in the output of the last print command, two new dynamic entries appeared in the address list (marked with a status of 'D'). Hosts with these IP addresses tried to initialize a telnet session to the router and were then subsequently dropped by the filter rule.

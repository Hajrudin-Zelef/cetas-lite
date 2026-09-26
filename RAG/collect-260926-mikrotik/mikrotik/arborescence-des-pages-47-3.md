---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-47-3
title: "Common Actions and Associated properties"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-47.md
source_anchor: ""
source_lines: [179, 191]
sha256: e9e608a6494bb19f519257b4f8d16be226f5f973dc1afa0c7d54b2f52b5ed7d7
---

# Common Actions and Associated properties

| Property | Description | 
|---|---|
| **connection-bytes** (*integer-integer* ; Default: ) | Matches packets only if a given amount of bytes has been transferred through the particular connection. 0 - means infinity, for example `connection-bytes=2000000-0` means that the rule matches if more than 2MB has been transferred through the relevant connection | 
| **connection-limit** (*integer,netmask* ; Default: ) | Matches connections per address or address block after a given value is reached. Should be used together with `connection-state=new` and/or with`tcp-flags=syn` because matcher is very resource-intensive | 
| **connection-mark** (*no-mark \| string* ; Default: ) | Matches packets marked via mangle facility with particular connection mark. If **no-mark** is set, the rule will match any unmarked connection | 
| **connection-nat-state** (*srcnat \| dstnat* ; Default: ) | Can match connections that are srcnatted, distracted, or both. Note that `connection-state=related` connections connection-nat-state is determined by the direction of the first packet. and if connection tracking needs to use dst-nat to deliver this connection to the same hosts as the main connection it will be in connection-nat-state=dstnat even if there are no dst-nat rules at all | 
| **connection-rate** (*Integer 0..4294967295* ; Default: ) | Connection Rate is a firewall matcher that allows capturing traffic based on the present speed of the connection | 
| **connection-state** (*established \| invalid \| new \| related \| untracked* ; Default: ) | Interprets the connection tracking analytics data for a particular packet:  | 
| **connection-type** (*ftp \| h323 \| irc \| pptp \| quake3 \| sip \| tftp* ; Default: ) | Matches packets from related connections based on information from their connection tracking helpers. A relevant connection helper must be enabled under the: `/ip firewall service-port` | 
| **layer7-protocol** (*name* ; Default: ) | Layer7 filter name defined in layer7 protocol menu. Read more>>. | 
| **p2p** () | Matches some unencrypted P2P protocols. Deprecated in modern days since mostly everything is encrypted and requires deep packet inspection to identify. **IPv4** only. | 
| **realm** (*integer: 0..4294967295* ; Default: ) | **IPv4** only. | 
| **routing-mark** (*string* ; Default: ) | Matches packets marked by mangle facility with particular routing mark |

---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-81-2
title: "Introduction"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-81.md
source_anchor: ""
source_lines: [119, 160]
sha256: e7196960e4f570f282cb17e847228d868b00f33f00d6a8ad24dc437856667996
---

# Introduction

| Property | Description | 
|---|---|
| **assured** (*yes \| no* ) | Indicates that this connection is assured and that it will not be erased if the maximum possible tracked connection count is reached. | 
| **confirmed** (*yes \| no* ) | Connection is confirmed and a packet is sent out from the device | 
| **connection-mark** (*string* ) | Connection mark that was set by the mangle rule. | 
| **connection-type** (*pptp \| ftp* ) | Type of connection, the property is empty if connection tracking is unable to determine a predefined connection type. | 
| **dst-address** (*ip* ) | Destination address. | 
| **dst-port**  (*integer* ) | Destination port. | 
| **dstnat** (*yes \| no* ) | A connection has gone through DST-NAT (for example, port forwarding). | 
| **dying** (*yes \| no* ) | The connection is dying due to a connection timeout. | 
| **expected** (*yes \| no* ) | Connection is set up using connection helpers (pre-defined service rules). | 
| **fasttrack** (*yes \| no* ) | Whether the connection is FastTracked. | 
| **gre-key** (*integer* ) | Contents of the GRE Key field. | 
| **gre-protocol** (*string* ) | Protocol of the encapsulated payload. | 
| **gre-version** (*string* ) | A version of the GRE protocol was used in the connection. | 
| **connection-mark**  (*string* ) | Connection mark assigned for the connection from firewall. | 
| **hw-offload**  (*yes \| no* ) | Hardware offloaded connection. | 
| **icmp-code** (*string* ) | ICMP Code Field | 
| **icmp-id** (*integer* ) | Contains the ICMP ID | 
| **icmp-type** (*integer* ) | ICMP Type Number | 
| **orig-bytes** (*integer* ) | Amount of bytes sent out from the source address using the specific connection. | 
| **orig-fasttrack-bytes** (*integer* ) | Amount of FastTracked bytes sent out from the source address using the specific connection. | 
| **orig-fasttrack-packets** (*integer* ) | Amount of FastTracked packets sent out from the source address using the specific connection. | 
| **orig-packets** (*integer* ) | Amount of packets sent out from the source address using the specific connection. | 
| **orig-rate** (*integer* ) | The data rate at which packets are sent out from the source address using the specific connection. | 
| **protocol** (*string* ) | IP protocol type | 
| **repl-bytes** (*integer* ) | Amount of bytes received from the destination address using the specific connection. | 
| **repl-fasttrack-bytes** (*string* ) | Amount of FastTracked bytes received from the destination address using the specific connection. | 
| **repl-fasttrack-packets** (*integer* ) | Amount of FastTracked packets received from the destination address using the specific connection. | 
| **repl-packets** (*integer* ) | Amount of packets received from the destination address using the specific connection. | 
| **repl-rate** (*string* ) | The data rate at which packets are received from the destination address using the specific connection. | 
| **reply-dst-address** (*ip* ) | Destination address expected of return packets. | 
| **reply-dst-port**  (*integer* ) | Destination port expected of return packets. | 
| **reply-src-address** (*ip* ) | Source address expected of return packets. | 
| **reply-src-port** (*integer* ) | Source port expected of return packets. | 
| **seen-reply** (*yes \| no* ) | The destination address has replied to the source address. | 
| **src-address** (*ip* ) | The source address. | 
| **src-port**  (*integer* ) | The source port. | 
| **srcnat** (*yes \| no* ) | Connection is going through SRC-NAT, including packets that were masqueraded through NAT. | 
| **tcp-state** (*string* ) | The current state of TCP connection :  | 
| **timeout** (*time* ) | Time after connection will be removed from the connection list. | 
| **uses-helper**  (*yes \| no* ) | "IP/Firewall/Service Port" helper has been applied to the particular connection. |

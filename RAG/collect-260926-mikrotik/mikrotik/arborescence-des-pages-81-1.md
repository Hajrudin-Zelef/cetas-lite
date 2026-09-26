---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-81-1
title: "Introduction"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-81.md
source_anchor: ""
source_lines: [1, 118]
sha256: 0ad55abb14e4ea519df728562d5e1f66d51ef437f5edf465b1defcb78f311aee
---

# Introduction

Connection tracking allows the kernel to keep track of all logical network connections or sessions, and thereby relate all of the packets which may make up that connection.

NAT relies on this information to translate all related packets in the same way.

Because of connection tracking you can use stateful firewall functionality even with stateless protocols such as UDP.

Firewall features affected by connection tracking:

- NAT
- firewall:
  - connection-bytes
  - connection-mark
  - connection-type
  - connection-state
  - connection-limit
  - connection-rate
  - layer7-protocol
  - new-connection-mark
  - tarpit

List of tracked connections can be seen in /ip firewall connection for IPv4 and /ipv6 firewall connection for IPv6.

# Connection states

Based on connection table entries arrived packet can get assigned one of the connection states: **new, invalid, established, related,** or **untracked**.

There are two different methods when the packet is considered **new**. The first one is in the case of stateless connections (like UDP) when there is no connection entry in the connection table. The other one is in the case of a stateful protocol (TCP). In this case, a new packet that starts a new connection is always a TCP packet with an *SYN* flag.

If a packet is not new it can belong to either an ***established*** or ***related*** connection or not belong to any connection making it ***invalid***. A packet with an ***established*** state, as most of you already guessed, belongs to an existing connection from the connection tracking table. A ***related*** state is very similar, except that the packet belongs to a connection that is related to one of the existing connections, for example, ICMP error packets or FTP data connection packets.

Connection state **notrack** is a special case when **RAW** firewall rules are used to exclude connection from connection tracking. This rule would make all forwarded traffic bypass the connection tracking, improving packet processing speed through the device.

Any other packet is considered ***invalid*** and in most cases should be dropped.

Based on this information we can set a basic set of filter rules to speed up packet filtering and reduce the load on the CPU by accepting *established/related* packets, dropping *invalid* packets, and working on more detailed filtering only for *new* packets.

Such a rule set must not be applied on routers with asymmetric routing, because asymmetrically routed packets may be considered invalid and dropped.

# FastTrack

IPv4 FastTrack is a special handler that bypasses Linux facilities allowing for faster packet forwarding. The handler is used for **TCP** and **UDP** connections marked with "`fasttrack-connection`" action. IPv4 FastTrack handler supports NAT (SNAT, DNAT, or both).

Note that not all packets of the connection can be FastTracked, so it is likely to see some packets going through a slow path even though the connection is marked for FastTrack. This is the reason why **fasttrack-connection** is usually followed by an identical "*action=accept*

FastTrack-ed packets are bypassing:

- firewall,
- connection tracking,
- simple queues,
- queue tree with *parent=global* ,
- IP accounting,
- IPSec,
- hotspot universal client,
- VRF assignment

It is up to the administrator to make sure FastTrack does not interfere with other configuration.

## Requirements

IPv4 FastTrack is active if the following conditions are met:

- no mesh, metarouter interface configuration;
- sniffer, torch, or traffic generator is not running;
- */tool mac-scan* is not actively used;
- */tool ip-scan* is not actively used;
- FastPath and Route cache is enabled under *IP/Settings*

## Example

For example, for SOHO routers with factory default configuration, you could FastTrack all LAN traffic with this one rule placed at the top of the Firewall Filter. The same configuration accept rule is required:

- Connection is FastTracked until the connection is closed, timed-out, or router is rebooted.
- Dummy rules will disappear only after FastTrack firewall rules will be deleted/disabled and the router rebooted.
- While FastPath and FastTrack both are enabled on the device only one can be active at a time.

Queues (except Queue Trees parented to interfaces), firewall filter, and mangle rules will not be applied for FastTracked traffic.

# Connection tracking settings

Connection tracking settings are managed from `/ip firewall connection tracking` menu.

### Properties

| Property | Description | 
|---|---|
| **enabled** (*yes \| no \| auto* ; Default:**auto** ) | Allows to disable or enable connection tracking. With disabled connection tracking firewall features listed above will stop working. If set to "auto" connection tracking is disabled until at least one firewall rule is added. | 
| **liberal-tcp-tracking** (*yes \| no;* Default:**no** ) | Enables or disables liberal TCP connection tracking by toggling the kernel parameter `nf_conntrack_tcp_be_liberal` . When set to**yes** , the system mark only out of window RST segments as INVALID. Enabling this setting may allow malformed packets that would otherwise be considered `invalid` by the firewall's`connection-state` matcher. This can increase exposure to certain evasion techniques. This property should be enabled only when troubleshooting or working around known issues. | 
| **loose-tcp-tracking** (*yes* ; Default:**yes** ) |  | 
| **tcp-syn-sent-timeout** (*time* ; Default:**5s** ) | TCP SYN timeout. | 
| **tcp-syn-received-timeout** (*time* ; Default:**5s** ) | TCP SYN timeout. | 
| **tcp-established-timeout** (*time* ; Default:**1d** ) | Time after which established TCP connection times out. | 
| **tcp-fin-wait-timeout** (*time* ; Default:**10s** ) |  | 
| **tcp-close-wait-timeout** (*time* ; Default:**10s** ) |  | 
| **tcp-last-ack-timeout** (*time* ; Default:**10s** ) |  | 
| **tcp-time-wait-timeout** (*time* ; Default:**10s** ) |  | 
| **tcp-close-timeout** (*time* ; Default:**10s** ) |  | 
| **udp-timeout** (*time* ; Default:**3****0s** ) | Specifies the timeout for UDP connections that have seen packets in one direction | 
| **udp-stream-timeout** (*time* ; Default:**3m** ) | Specifies the timeout of UDP connections that have seen packets in both directions | 
| **icmp-timeout** (*time* ; Default:**10s** ) | ICMP connection timeout | 
| **generic-timeout** (*time* ; Default:**10m** ) | Timeout for all other connection entries | 

**Read-only properties**

| Property | Description | 
|---|---|
| **max-entries** (*integer* ) | Max amount of entries that the connection tracking table can hold. This value depends on the installed amount of RAM. Note that the system does not create a maximum-size connection tracking table when it starts, it may increase if the situation demands it and the system still has free RAM, but the size will not exceed 1048576 | 
| **total-entries** (*integer* ) | Amount of connections that the connection table currently holds | 

# Connection List

List of tracked connections can be seen in /ip firewall connection for ipv4 and /ipv6 firewall connection for IPv6.

## Properties

All properties in the connection list are read-only


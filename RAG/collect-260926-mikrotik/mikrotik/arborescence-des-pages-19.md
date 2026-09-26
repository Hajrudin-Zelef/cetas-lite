---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-19
title: "Summary"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-19.md
source_anchor: ""
source_lines: [1, 38]
sha256: 0d494f9c4b9f537b0ebef9a2793ad33c0f7c924ef43e6d1cffc1ac8ec2199c0f
---

# Summary

Ping uses the Internet Control Message Protocol (ICMP) Echo messages to determine if a remote host is active or inactive and to determine the round-trip delay when communicating with it. Ping tool sends ICMP (type 8) message to the host and waits for the ICMP echo-reply (type 0). The interval between these events is called a round trip. If the response (that is called pong) has not come until the end of the interval, we assume it has timed out. The second significant parameter reported is TTL (Time to Live). Is decremented at each machine in which the packet is processed. The packet will reach its destination only when the TTL is greater than the number of routers between the source and the destination.

## Quick Example

RouterOS Ping tool allows you to configure various additional parameters like:

- arp-ping;
- address;
- src-address;
- count;
- dscp;
- interface;
- interval;
- routing-table;
- size;
- ttl;

Let's take a look ar very simple example:

The same we can achieve with more shorter CLI command:

It is also possible to ping multicast address to discover all hosts belonging to multicast group:

Ping by DNS name:

When you use the domain name and CLI for ping, router DNS will be used to resolve the address. When you use the Winbox Tools/Ping, your computer's DNS will be used to resolve the given address.

## MAC Ping

This submenu allows enabling the mac ping server.

When mac ping is enabled, other hosts on the same broadcast domain can use the ping tool to ping mac address:

Ping MAC address:

By default, a MAC ping attempts to reach the destination through all active interfaces. This can generate unwanted traffic and duplicate replies if the destination is reachable via multiple interfaces. To restrict a MAC ping to a specific interface, use the interface selector (append `%` followed by the interface name to the MAC address). For example: `/ping 00:11:22:33:44:55%ether1`

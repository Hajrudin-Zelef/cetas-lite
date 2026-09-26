---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-49-8
title: "DHCP Client"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-49.md
source_anchor: ""
source_lines: [696, 770]
sha256: 1830c13c3c0a7dafa11c38fa19924c50db4718ca02e2eeb19724ec0fbe164c52
---

# DHCP Client

The purpose of the DHCP relay is to act as a proxy between DHCP clients and the DHCP server. It is useful in networks where the DHCP server is not on the same broadcast domain as the DHCP client.

DHCP relay does not choose the particular DHCP server in the DHCP-server list, it just sends the incoming request to all the listed servers.

## Properties

| Property | Description | 
|---|---|
| **add-relay-info** (*yes \| no* ; Default:**no** ) | Adds DHCP relay agent information if enabled according to RFC 3046. Agent Circuit ID Sub-option contains mac address of an interface, Agent Remote ID Sub-option contains MAC address of the client from which request was received. | 
| **delay-threshold** (*time \| none* ; Default:**none** ) | If secs field in DHCP packet is smaller than delay-threshold, then this packet is ignored | 
| **dhcp-server** (IPv4 a*ddress [IPv4]* ; Default: ) | List of DHCP servers' IP addresses which should the DHCP requests be forwarded to | 
| **interface** (interface; Default: ) | Interface name the DHCP relay will be working on. | 
| **local-address** (*IP* ; Default:**0.0.0.0** ) | The unique IP address of this DHCP relay needed for DHCP server to distinguish relays. If set to **0.0.0.0** - the IP address will be chosen automatically from addresses that are assigned to an interface a relay is running | 
| **relay-info-remote-id** (*string* ; Default: ) | specified string will be used to construct Option 82 instead of client's MAC address. Option 82 consist of: interface from which packets was received + client mac address or **relay-info-remote-id** | 
| **name** (*string* ; Default: ) | Descriptive name for the relay | 
| **local-address-as-src-ip(** yes \| no; Default: **no** ) | Use local address as source address for Discover/Request packets sent to the DHCP server | 
| **disabled** (yes \| no; Default: no) | Whether relay is disabled or not. By default, it is not disabled | 
| **dhcp-server-vrf**  (vrf; Default: main) | Specifies the VRF on which the DHCP relay should operate. | 

## Configuration Example


Let us consider that you have several IP networks 'behind' other routers, but you want to keep all DHCP servers on a single router. To do this, you need a DHCP relay on your network which will relay DHCP requests from clients to the DHCP server.

This example will show you how to configure a DHCP server and a DHCP relay that serves 2 IP networks - 192.168.1.0/24 and 192.168.2.0/24 that are behind a router DHCP-Relay.

**IP Address Configuration**

IP addresses of DHCP-Server:

IP addresses of DHCP-Relay:

**DHCP Server Setup**

To setup 2 DHCP Servers on the DHCP-Server router add 2 pools. For networks 192.168.1.0/24 and 192.168.2.0:

Create DHCP Servers:

Configure respective networks:

**DHCP Relay Config**

Configuration of DHCP-Server is done. Now let's configure DHCP-Relay:

## DHCP Relay with VRF (introduced in 7.15)

Let's take the previous setup but we'll consider that the interface to the DHCP server and interfaces to DHCP clients are added in VRF:

In the DHCP-relay configuration dhcp-server-vrf should be added:

Due to VRF configuration there are several routing-tables - we should add additional routes:

To achieve successful DHCP-server - DHCP-relay communication we should add NAT rules:

# DHCPv6 Relay

## Summary

**Sub-menu:** `/ipv6 dhcp-relay`

DHCPv6 Relay in RouterOS acts as an intermediary that forwards DHCPv6 client solicitations received on a local interface to a remote DHCPv6 server. The relay adds a Hop‑By‑Hop option containing its own link‑local address, enabling the server to know the client’s network location. Responses from the server are returned to the relay, which removes the Hop‑By‑Hop option and delivers the reply to the originating client. This allows a single DHCPv6 server to service multiple subnets without being directly connected to each, while preserving proper address assignment and prefix delegation. The relay requires IPv6 forwarding to be enabled and appropriate firewall allowances for UDP ports 546 (client) and 547 (server).

## Properties

| Property | Description | 
|---|---|
| **comment** (*string* ; Default: ) | Descriptive name of an item. | 
| **delay-threshold** (*time \| none* ; Default:**none** ) | If secs field in DHCP packet is smaller than delay-threshold, then this packet is ignored. | 
| **dhcp-server** (*IPv6 address [IPv6]%interface* ; Default: ) | A list of DHCP server IP addresses to which DHCP requests should be forwarded (optionally, the interface can also be specified together with the IPv6 address). | 
| **interface** (*interface* ; Default: ) | Interface name the DHCP relay will be working on. | 
| **name** (*string* ; Default: ) | Descriptive name for the relay | 
| **dhcp-options**  (*DHCPv6 option* ; Default:**client_mac** ) | A list of DHCPv6 options to be inserted by the relay into forwarded DHCPv6 packets. By default, relay inserts option 79. | 
| **disabled** (*yes \| no* ; Default: no) | Whether relay is disabled or not. By default, it is not disabled. | 
| **link-address**  (*IPv6 address [IPv6]* ; Defaul: :: ) | An IPv6 address that may be used by the server to identify the link on which the client is located. | 
| **store-relayed-bindings**  (*yes \| no* ; Default:**no** ) | Inspects relayed DHCP advertisements and stores assigned prefixes. Should be used, for example, to avoid loss of routing information on relay reboot. By default, it is disabled. |

---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-43
title: "Introduction"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["throughput"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-43.md
source_anchor: ""
source_lines: [1, 132]
sha256: 7bd6ddd67a99ec93f51d86f151879f2584309599c42f15a787161d5dbc8cbfd5
---

# Introduction

Firewall filters are used to allow or block specific packets forwarded to your local network, originating from your router, or destined to the router.

There are two methods on how to set up filtering:

- allow specific traffic and drop everything else
- drop only malicious traffic, everything else is allowed.

Both methods have pros and cons, for example, from a security point of view first method is much more secure, but requires administrator input whenever traffic for a new service needs to be accepted. This strategy provides good control over the traffic and reduces the possibility of a breach because of service misconfiguration.

On the other hand, when securing a customer network it would be an administrative nightmare to accept all possible services that users may use. Therefore careful planning of the firewall is essential in advanced setups.

A firewall filter consists of three predefined chains that cannot be deleted:

- **input** - used to process packets**__entering the router__** through one of the interfaces with the destination IP address which is one of the router's addresses. Packets passing through the router are not processed against the rules of the input chain
- **forward** - used to process packets**__passing through the router__**
- **output** - used to process packets**__originating from the router__** and leaving it through one of the interfaces. Packets passing through the router are not processed against the rules of the output chain

Firewall filter configuration is accessible from `ip/firewall/filter` menu for IPv4 and `ipv6/firewall/filter` menu for IPv6.

# Firewall Example

Lets look at basic firewall example to protect router itself and clients behind the router, for both IPv4 and IPv6 protocols.

## IPv4 firewall

### Protect the router itself

Rules of thumb followed to set up the firewall:

- work with `new` connections to decrease the load on a router;
- accept what you need
- `drop` everything else,`log=yes` could be set to log some attackers, but keep in mind that it may add some load to he CPU on heavy attacks.

We always start by accepting already known and accepted connections, so the first rule should be to accept "established" and "related" connections.

Now we can proceed by accepting some new connections, in our example we want to allow access ICMP protocol from any address and everything else only from 192.168.88.2-192.168.88.254 address range. For that we create an address list and two firewall rules.

And lastly we drop everything else:

Complete set of just created rules:

### Protect the LAN devices

Concept in protecting the users is very similar, except that in this case we are blocking unwanted traffic and accepting everythign else.

At first we will create `address-list` with the name "not_in_internet" which we will use for the firewall filter rules:

Brief firewall filter rule explanation:

- packets with *connection-state=established,related* added to FastTrack for faster data throughput, the firewall will work with new connections only;
- drop *invalid* connection and log them with prefix "invalid";
- drop attempts to reach not public addresses from your local network, apply *address-list=not_in_internet* before, "bridge" is local network interface, log=yes attempts with prefix "!public_from_LAN";
- drop incoming packets that are not NAT`ed, ether1 is public interface, log attempts with "!NAT" prefix;
- jump to ICMP chain to drop unwanted ICMP messages
- drop incoming packets from the Internet, which are not public IP addresses, ether1 is a public interface, log attempts with prefix "!public";
- drop packets from LAN that does not have LAN IP, 192.168.88.0/24 is local network used subnet;

Allow only needed ICMP codes in "icmp" chain:

## 

IPv6 firewall 

### Protect the router itself

Very similar to IPv4 setup, except that we have to deal with more protocols required for IPv6 to function properly.

At first we create an `address-list` from which you allow access to the device:

Brief IPv6 firewall filter rule explanation:

- work with *new* packets, accept*established/related* packets;
- drop *link-local* addresses from Internet(public) interface/interface-list;
- accept access to a router from *link-local* addresses, accept*multicast* addresses for management purposes, accept your source*address-list* for router access;
- drop anything else;

In certain setups where the DHCPv6 relay is used, the src address of the packets may not be from the link-local range. In that case, the src-address parameter of rule #4 must be removed or adjusted to accept the relay address.

### Protect the LAN devices

This step is more important than it is for IPv4. In IPv4 setups clients mostly have addresses from local address range and are NATed to public IP, that way they are not directly reachable from the public networks.

IPv6 is a different story. In most common setups, enabled IPv6 makes your clients available from the public networks, so proper firewall filter rules to protect your customers are mandatory.

In brief we will very basic LAN protection should:

- accept *established/related* and work with*new* packets;
- drop *invalid* packets;
- accept ICMPv6 packets;
- accept *new* connections originated only from your clients to the public network;
- drop everything else.

# Matchers

All matcher properties are common and listed here.

# Actions

Tables below shows list of filter specific actions and associated properties.

| Property | Description | 
|---|---|
| **action** (*action name* ; Default:**accept** ) |   | 
| **hw-offload** (*no \| yes* ; Default:**yes** ) | Enables or disables FastTrack hardware offloading. Supported only on switches with FastTrack offloading and when `action=fasttrack-connection` is set. | 
| **reject-with** (*icmp-no-route \| icmp-admin-prohibited \| icmp-not-neighbour* *\| icmp-address-unreachable \| icmp-port-unreachable \| tcp-reset \| icmp-err-src-routing-header \| icmp-headers-too-long* ; Default:**icmp-no-route** ) | Specifies ICMP error to be sent back if the packet is rejected. Applicable if `action=reject`  | 

# RAW Filtering

The firewall RAW table allows to selectively bypass or drop packets before connection tracking, that way significantly reducing the load on the CPU. The tool is very useful for DoS/DDoS attack mitigation.

RAW filter configuration is accessible from `ip/firewall/raw` menu for IPv4 and `ipv6/firewall/raw` menu for IPv6.

The RAW table does not have matchers that depend on connection tracking ( like connection-state, layer7, etc.).

If a packet is marked to bypass the connection tracking packet de-fragmentation will not occur.

Also RAW firewall can have rules only in two chains:

- **prerouting** - used to process any packet entering the router
- **output** - used to process packets originated from the router and leaving it through one of the interfaces. Packets passing through the router are not processed against the rules of the output chain

And has one specific action:

| Property | Description | 
|---|---|
| **action** (*action name* ; Default:**accept** ) |  | 

## Basic RAW Example

Let's assume that we have OSPF configuration, but due to connection tracking OSPF have adjacency problems. We can use RAW rules to fix this, by not sending OSPF packets to connection tracking.

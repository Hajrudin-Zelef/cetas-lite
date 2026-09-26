---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-32-1
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-32.md
source_anchor: ""
source_lines: [1, 113]
sha256: c6d7a7d9c3322aaeeed610c95d9acb0a27592fe94a657ff361fb580c5356af5b
---

# Overview

MikroTik RouterOS implements Label Distribution Protocol (RFC 3036, RFC 5036, and RFC 7552) for IPv4 and IPv6 address families. LDP is a protocol that performs the set of procedures and exchange messages by which Label Switched Routers (LSRs) establish Label Switched Paths (LSPs) through a network by mapping network-layer routing information directly to data-link layer switched paths.

# Prerequisites for MPLS

## "Loopback" IP address

Although not a strict requirement, it is advisable to configure routers participating in the MPLS network with "loopback" IP addresses (not attached to any real network interface) to be used by LDP to establish sessions.

This serves 2 purposes:

- as there is only one LDP session between any 2 routers, no matter how many links connect them, the loopback IP address ensures that the LDP session is not affected by interface state or address changes
- use of loopback address as LDP transport address ensures proper penultimate hop popping behavior when multiple labels are attached to the packet as in the case of VPLS

In RouterOS "loopback" IP address can be configured by creating a dummy bridge interface without any ports and adding the address to it. For example:

## IP connectivity

As LDP distributes labels for active routes, the essential requirement is properly configured IP routing. LDP by default distributes labels for active IGP routes (that is - connected, static, and routing protocol learned routes, except BGP).

For instructions on how to set up properly IGP refer to appropriate documentation sections:

LDP supports ECMP routes.

You should be able to reach any loopback address from any location of your network before continuing with the LDP configuration. Connectivity can be verified with the ping tool running from loopback address to loopback address.

# Example Setup

Let's consider that we have already existing four routers setup, with working IP connectivity.

## Ip Reachability

Not going deep into routing setup here is the quit export of the IP and OSPF configurations:

Verify that IP connectivity and routing are working properly

## LDP Setup

In order to start distributing labels, LDP is enabled on interfaces that connect other LDP routers and not enabled on interfaces that connect customer networks.

On R1 it will look like this:

Note that the transport address gets set to 111.111.111.1. This makes the router originate LDP session connections with this address and also advertise this address as a transport address to LDP neighbors.

Other routers are set up similarly.

R2:

On R3:

On R4:

After LDP sessions are established, R2 should have two LDP neighbors:

The local mappings table shows what label is assigned to what route and peers the router have distributed labels to.

Remote mappings table on the other hand shows labels that are allocated for routes by neighboring LDP routers and advertised to this router:

We can observe that router has received label bindings for all routes from both its neighbors - R1 and R3.

The remote mapping table will have active mappings only for the destinations that have direct next-hop, for example, let's take a closer look at 111.111.111.4 mappings. The routing table indicates that the network 111.111.111.4 is reachable via 111.12.0.2 (R3):

And if we look again at the remote mapping table, the only active mapping is the one received from R3 with assigned label 19. This implies that when R2 when routing traffic to this network, will impose label 19.

Label switching rules can be seen in the forwarding table:

If we take a look at rule number 2, the rule says that when R2 received the packet with label 19, it will change the label to new label 19 (assigned by the R3).

As you can see from this example it is not mandatory that labels along the path should be unique.

Now if we look at the forwarding table of R3:

Rule number 0, shows that the out label is "**impl-null**". The reason for this is that R3 is the last hop before 111.111.111.4 will be reachable and there is no need to swap to any real label. It is known that R4 is the egress point for the 111.111.111.4 network (router is the egress point for directly connected networks because the next hop for traffic is not MPLS router), therefore it advertises the "implicit null" label for this route. This tells R3 to forward traffic for the destination 111.111.111.4/32 to R4 unlabelled, which is exactly what R3 forwarding table entry tells.

Action, when the label is not swapped to any real label, is called **Penultimate hop popping,** it ensures that routers do not have to do unnecessary label lookup when it is known in advance that the router will have to route the packet.

# Using traceroute in MPLS networks

RFC4950 introduces extensions to the ICMP protocol for MPLS. The basic idea is that some ICMP messages may carry an MPLS "label stack object" (a list of labels that were on the packet when it caused a particular ICMP message). ICMP messages of interest for MPLS are Time Exceeded and Need Fragment.

MPLS label carries not only label value, but also TTL field. When imposing a label on an IP packet, MPLS TTL is set to value in the IP header, when the last label is removed from the IP packet, IP TTL is set to value in MPLS TTL. Therefore MPLS switching network can be diagnosed by means of a traceroute tool that supports MPLS extension.

For example, the traceroute from R4 to R1 looks like this:

Traceroute results show MPLS labels on the packet when it produced ICMP Time Exceeded. The above means: that when R3 received a packet with MPLS TTL 1, it had label 18 on it. This match advertised label by R3 for 111.111.111.4. In the same way, R2 observed label 17 on the packet on the next traceroute iteration - R3 switched label 17 to label 17, as explained above. R1 received packet without labels - R2 did penultimate hop popping as explained above.

## Drawbacks of using traceroute in MPLS network

### Label switching ICMP errors

One of the drawbacks of using traceroute in MPLS networks is the way MPLS handles produced ICMP errors. In IP networks ICMP errors are simply routed back to the source of the packet that caused the error. In an MPLS network, it is possible that a router that produces an error message does not even have a route to the source of the IP packet (for example in the case of asymmetric label switching paths or some kind of MPLS tunneling, e.g. to transport MPLS VPN traffic).

Due to this produced ICMP errors are not routed to the source of the packet that caused the error but switched further along the label switching path, assuming that when the label switching path endpoint will receive an ICMP error, it will know how to properly route it back to the source.

This causes the situation, that traceroute in MPLS network can not be used the same way as in IP network - to determine failure point in the network. If the label switched path is broken anywhere in the middle, no ICMP replies will come back, because they will not make it to the far endpoint of the label switching path.

### Penultimate hop popping and traceroute source address

A thorough understanding of penultimate hop behavior and routing is necessary to understand and avoid problems that penultimate hop popping causes to traceroute.

In the example setup, a regular traceroute from R5 to R1 would yield the following results:

```
[admin@R5] > /tool traceroute 9.9.9.1
     ADDRESS                                    STATUS
   1         0.0.0.0 timeout timeout timeout
   2         2.2.2.2 37ms 4ms 4ms
                      mpls-label=17
   3         9.9.9.1 4ms 2ms 11ms
```
compared to:


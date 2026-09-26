---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-32-2
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "ethernet"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-32.md
source_anchor: ""
source_lines: [114, 228]
sha256: eb3d0d1924853712476d2f954f58637e0add5adb7555b2e3bfc4bfaca345e398
---

# Overview

```
[admin@R5] > /tool traceroute 9.9.9.1 src-address=9.9.9.5
     ADDRESS                                    STATUS
   1         4.4.4.3 15ms 5ms 5ms
                      mpls-label=17
   2         2.2.2.2 5ms 3ms 6ms
                      mpls-label=17
   3         9.9.9.1 6ms 3ms 3ms
```
The reason why the first traceroute does not get a response from R3 is that by default traceroute on R5 uses source address 4.4.4.5 for its probes because it is the preferred source for a route over which next-hop to 9.9.9.1/32 is reachable:

[admin@R5] > /ip route print
Flags: X - disabled, A - active, D - dynamic,
C - connect, S - static, r - rip, b - bgp, o - ospf, m - mme,
B - blackhole, U - unreachable, P - prohibit
 #      DST-ADDRESS        PREF-SRC        G GATEWAY         DISTANCE             INTERFACE
 ...
 3 ADC  4.4.4.0/24         4.4.4.5                           0                    ether1
 ...
 5 ADo  9.9.9.1/32                         r 4.4.4.3         110                  ether1
 ...

When the first traceroute probe is transmitted (source: 4.4.4.5, destination 9.9.9.1), R3 drops it and produces an ICMP error message (source 4.4.4.3 destination 4.4.4.5) that is switched all the way to R1. R1 then sends ICMP error back - it gets switched along the label switching path to 4.4.4.5.

R2 is the penultimate hop popping router for network 4.4.4.0/24 because 4.4.4.0/24 is directly connected to R3. Therefore R2 removes the last label and sends ICMP error to R3 unlabelled:

[admin@R2] > /mpls forwarding-table print
 # IN-LABEL             OUT-LABELS           DESTINATION        INTERFACE            NEXTHOP
 ...
 3 19                                        4.4.4.0/24         ether2               2.2.2.3
 ...

R3 drops the received IP packet because it receives a packet with its own address as a source address. ICMP errors produced by following probes come back correctly because R3 receives unlabelled packets with source addresses 2.2.2.2 and 9.9.9.1, which are acceptable to a route.

Command:

[admin@R5] > /tool traceroute 9.9.9.1 src-address=9.9.9.5
 ...

produces expected results, because the source address of traceroute probes is 9.9.9.5. When ICMP errors are traveling back from R1 to R5, the penultimate hop popping for the 9.9.9.5/32 network happens at R3, therefore it never gets to route packet with its own address as a source address.

# Optimizing label distribution

## Label binding filtering

During the implementation of the given example setup, it has become clear that not all label bindings are necessary. For example, there is no need to exchange IP route label bindings between R1 and R3 or R2 and R4, as there is no chance they will ever be used. Also, if the given network core is providing connectivity only for mentioned customer ethernet segments, there is no real use to distribute labels for networks that connect routers between themselves, the only routes that matter are /32 routes to endpoints or attached customer networks.

Label binding filtering can be used to distribute only specified sets of labels to reduce resource usage and network load.

There are 2 types of label binding filters:

- which label bindings should be advertised to LDP neighbors, configured in the `/mpls ldp advertise-filter` menu
- which label bindings should be accepted from LDP neighbors, configured in `/mpls ldp accept-filter` menu

Filters are organized in the ordered list, specifying prefixes that must include the prefix that is tested against the filter and neighbor (or wildcard).

In the given example setup all routers can be configured so that they advertise labels only for routes that allow reaching the endpoints of tunnels. For this 2 advertise filters need to be configured on all routers:

This filter causes routers to advertise only bindings for routes that are included by the 111.111.111.0/24 prefix which covers loopbacks (111.111.111.1/32, 111.111.111.2/32, etc). The second rule is necessary because the default filter results when no rule matches are to allow the action in question.

In the given setup there is no need to set up accept filter because by convention introduced by 2 abovementioned rules no LDP router will distribute unnecessary bindings.

Note that filter changes do not affect existing mappings, so to take the filter into effect, connections between neighbors need to be reset. either by removing neighbors from the LDP neighbor table or by restarting the LDP instance.

So on R2, for example, we get:

# LDP on Ipv6 and Dual-Stack links

RouterOS implements RFC 7552 to support LDP on dual-stack links.

Supported AFIs can be selected by LDP instance, as well as explicitly configured per LDP interface.

The example above enables LDP instance to use IPv4 and IPv6 address families and sets the preference to IPv6 with `preferred-afi` parameter. LDP interface configuration on the other hand explicitly sets that **ether2** supports only IPv4 and **ether3** supports only IPv6.

The main question occurs how AFI is selected when there are a mix of different AFIs and what if one of the supported AFIs flaps.

The logic behind sending hellos is as follows:

- if an interface has only one AFI:
  - dual-stack element is not sent
  - sends hello only if there is an IP address on the interface from the corresponding AFI.
- If an interface has both AFIs:
  - dual-stack element is always sent and contains the value from preferred-afi
  - sends hellos on each AFI if a corresponding address is present on the interface.

From all received hellos peer determines which AFI to use for connection and for which AFIs to bind and send labels. For LDP to be able to use a specific AFI, receiving hello for that specific AFI is mandatory. Hello packet contains the transport address necessary for proper LDP operation. By comparing received AFI addresses, is determined active/passive role.

The logic behind receiving and processing hellos is as follows:

- if the LDP instance has only one AFI (it means that all interfaces can have only that specific AFI operational):
  - drop hellos from not supported AFI
  - ignore/forget the dual-stack element for the hello packet
  - the role is determined only for this one specific AFI
  - labels are sent only for this one specific AFI
- if the LDP instance has both AFIs (interfaces can have different combinations of supported AFIs):
  - drop hellos from AFI that are not configured as supported on the interface.
  - ignore/forget the dual-stack element (preference is not taken into account) for hello packets, if an interface has only one supported AFI.
  - drop hello if received preference in dual-stack element does not match configured `preferred-afi` .

If there are changes in hello packets, the existing session is terminated only in case if address family used by labels is changed, otherwise, the session is preserved.

Dual-stack element in hello packets is set only if an interface is determined to be dual-stack compatible:

- Normally such an interface should be able to receive hellos from both AFIs,
  - Before proceeding LDP should wait for hello from the preferred AFI.
  - if hello is received only from one AFI:
    - if hello from preferred AFI is not received then it is considered an error.
    - otherwise, wait for missing hello for x seconds (x = 3 * hello-interval)
      - if missing hello appears within a time interval consider peer to be dual-stack
      - if missing hello did not appear, then consider peer to be single-stack
      - if missing hello appeared after the time interval then restart the session.
- the dual-stack element indicates that LDP wants to distribute labels for both AFIs.

In summary, the following combinations of AFIs and dual-stack element (ds6) are possible assuming that preferred-afi=ipv6:


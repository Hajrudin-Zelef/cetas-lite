---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-31-2
title: "Summary"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2026-02-12"]
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-31.md
source_anchor: ""
source_lines: [129, 259]
sha256: 4835d606a38bc0fafebbe71f4fb596aebec9f92a6b9996ae7b903d580f683c9c
---

# Summary

To see the actual active sessions with selected template parameters and negotiated capabilities refer to the BGP sessions menu:

This menu shows read-only cached BGP session information. It will show the current status of the session, flags, last received notification, and negotiated session parameters.

Even if the BGP session is not active anymore, the cache can still be stored for some time. Routes received from a particular session are removed only if the cache expires, this allows mitigating extensive routing table recalculations if the BGP session is flapping.

## Template Menu

The template contains all BGP protocol-related configuration options. It can be used as a template for dynamic peers and to apply a similar configuration to a group of peers. Note that this is not the same as peer groups on Cisco devices, where the group is more than just a common configuration.

# Best-Path Selection

BGP routers can receive multiple copies of the global routing table from multiple providers.

There should be some way to compare those multiple BGP routing tables and select the best route to the destination, the solution is the BGP Best Path Selection Algorithm.

The route is evaluated by the algorithm only if it is valid. In general, the route is considered valid if:

- NEXT_HOP of the route is valid and reachable
- AS_PATH received from external peers does not contain the local AS
- the route is not rejected by routing filters

For more information read nexthop selection and validation.

The best path algorithm also compares routes received only by a **single BGP instance**. Routes installed by different BGP instances are compared by the general algorithm, i.e. route distances are compared and the route with a lower distance is preferred.

If all the criteria are met, then the following actions take place:

1. The first path received is automatically considered the 'best path'. Any further received paths are compared to the first received to determine if the new path is better.
2. Prefer the path with the highest **WEIGHT** .
This parameter is not a part of the BGP standard, it is invented to quickly locally select the best route. A parameter is local to the router (assigned with routing filters in the BGP input) and cannot be advertised. A route without assigned WEIGHT has a default value of 0.
3. Prefer the path with the highest **LOCAL_PREF** .
This attribute is used only within an AS. A path without the LOCAL_PREF attribute has a value of 100 by default.
4. Prefer the path with the shortest **AS_PATH** . (skipped if`input.ignore-as-path-len` set to**yes** ).
Each AS_SET counts as 1, regardless of the set size. The AS_CONFED_SEQUENCE and AS_CONFED_SET are not included in the AS_PATH length.
5. ~~Prefer the path that was locally originated via aggregate or BGP network~~
6. Prefer the path with the lowest **ORIGIN** type.Interior Gateway Protocol (IGP) is lower than Exterior Gateway Protocol (EGP), and EGP is lower than INCOMPLETE in other words**IGP < EGP < INCOMPLETE**
7. Prefer the path with the lowest **multi-exit discriminator (MED).**The router compares the MED attribute only for paths that have the same neighboring (leftmost) AS. 
Paths without explicit MED value are treated with MED of 0.
8. Prefer **eBGP** over**iBGP** paths
9. Prefer lowest **IGP metric** .
10. If `multipath` is enabled, break here and and consider paths equal. This will install BGP ECMP routes.
11. Prefer the route that comes from the BGP router with the lowest **router ID** . If a route carries the**ORIGINATOR_ID** attribute, then the**ORIGINATOR_ID** is used instead of the router ID.
12. Prefer the route with the shortest **route reflection cluster list** . Routes without a cluster list are considered to have a cluster list of length 0.
13. Prefer the path that comes from the lowest neighbor address

# Routing Filter Notes

On BGP output routing filters are executed before BGP itself is modifying attributes, for example, if `nexthop-choice` is set to `force-self`, then the gateway set in the routing filters will be overridden.

On BGP input routing filters are applied to the received attributes, which means that, for example, setting the gateway will work no matter what `nexhop-choice` value is set.

# Running More than One Instance

As we already know for best path selection to work properly, BGP routes must be received from the same instance. But in certain scenarios it is necessary to run multiple BGP instances with their own separate tables. 

BGP determines whether sessions belongs to the same instance by comparing configured local router IDs. 

For example config below will run each peer in its own BGP instance


When `router-id` is not specified BGP will pick the "default" ID from `/routing id`.

Starting from v7.20 instance is no longer determined by router-id.

# Route Distinguisher

Route Distinguisher is a 64-bit integer, which is divided into three parts:

- type (always 2 bytes),
- administrator subfield,
- value or service provider subfield.

Currently, there are three format types defined.

| **2bytes** | 2bytes | 2bytes | 2bytes | 
|---|---|---|---|
| **Type1** | ASN | 4byte value |  | 
| **Type2** | 4-byte IP |  | value | 
| Type3 | 4-byte ASN |  | value | 

# BGP Unnumbered

RouterOS is capable of establishing BGP session dynamically without global IPv6 address on an interface and without specifying remote link-local address (BGP Unnumbered connection). Mechanism relies on a IPv6 neighbor discovery (RFC 4861) to get neighbor's link-local address and RFC5549 to advertise IPv4 NLRIs with IPv6 nexthops.

Unnumbered connection is configured if `remote.address` is empty and `local.address` is an interface. In this mode only one connection per interface is accepted.

For router without configured global addresses to be able to send replies for RAs, following ND config is needed:

```
/ipv6/nd/prefix/add prefix=none interface=sfp-sfpplus1
```

 

```
[admin@CCR2004_2XS_111] /ipv6/neighbor> print 
Flags: D - DYNAMIC; R - ROUTER
Columns: ADDRESS, MAC-ADDRESS, INTERFACE, VRF
#    ADDRESS                    MAC-ADDRESS        INTERFACE     VRF 
0 DR fe80::de2c:6eff:fec5:a7ff  DC:2C:6E:C5:A7:FF  sfp-sfpplus1  main
```

 

Now we can add unnumbered connection config:

```
/routing bgp 
add instance=myInstance local.address=sfp-sfpplus1 .role=ibgp name=unnumbered_2
```

 

```
[admin@CCR2004_2XS_111] /routing/bgp/connection> print 
Flags: D - DYNAMIC, X - DISABLED, I - INACTIVE 
 0   name="unnumbered_2" instance=v6_test 
     local.address=sfp-sfpplus1 .default-address=fe80::de2c:6eff:fea4:b42f%sfp-sfpplus1 .role=ibgp 
     routing-table=main as=333 
[admin@CCR2004_2XS_111] /routing/bgp/session> print 
Flags: E - ESTABLISHED 
 0 E name="unnumbered_2-1" instance=v6_test 
     remote.address=fe80::de2c:6eff:fec5:a7ff%sfp-sfpplus1 .as=333 .id=203.0.113.2 .capabilities=mp,rr,enhe,gr,as4 .afi= 
     .messages=5181 .bytes=98439 .eor="" 
     local.role=ibgp .address=fe80::de2c:6eff:fea4:b42f%sfp-sfpplus1 .as=333 .id=203.0.113.1 .cluster-id=203.0.113.1 
     .capabilities=mp,rr,enhe,gr,as4 .afi= .messages=5181 .bytes=98439 .eor="" 
     output.procid=20 
     input.procid=20 ibgp 
     multihop=yes hold-time=3m keepalive-time=1m uptime=3d14h20m55s640ms last-started=2026-02-12 18:26:59 prefix-count=0 
```

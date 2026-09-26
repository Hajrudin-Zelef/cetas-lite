---
id: collect-260926-mikrotik/mikrotik/mikrotik-failover-routing-description-living-document-1
title: "Static routes"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/mikrotik-failover-routing-description-living-document.md
source_anchor: ""
source_lines: [1, 146]
sha256: 426f2e9bb7c205f0584716768a0dd66d842c40bddf66aae4427d229c1872c0af
---

# Static routes

This is meant to be a somewhat-easier-to-digest recap of the discussion
that can be found on the MikroTik forum at this URL:
https://forum.mikrotik.com/viewtopic.php?f=23&t=157048&p=836497&hilit=failover#p836497
Note that the forum discussion not only addresses *failover*, but also
*load balancing* at the same time (without explicitly saying so in the
beginning).

For the sake of this document, we'll make some assumptions:

1. There are two uplinks, each bringing its own gateway: Gateway 1 and
gateway 2 .  (For the sake of this discussion, let the IP address of
gateway 1 be `1.2.3.4` and that of gateway 2 be`2.3.4.5` ).
2. The uplink via gateway 1 is to be the primary route to take, while
the uplink via gateway 2 is to be the backup route only.  In other
words, traffic should move via `1.2.3.4` unless this uplink breaks.
3. The primary uplink is attached via the `ether1` interface, while the
secondary uplink is attached via the`ether2` interface.

The obvious approach is to define two default routes, with the primary route having a smaller metric than the secondary route:

```
/ip route add gateway=1.2.3.4 distance=1 comment="Primary route"
/ip route add gateway=2.3.4.5 distance=2 comment="Secondary route"
```
This will have everything routed via `1.2.3.4` unless the `ether1`
interface goes down.  If `ether1` does go down, then the primary route
will become invalidated, and traffic will be routed through the secondary
route. The general picture is this:

```
+-----------+
| Internet  | <----------------------------------------------------------+
+-----------+                                                            |
  ^                                                                      |
  |                                                                      |
  |                                                                      |
+-----------+  Uplink 1 (primary)   +--------+  Uplink 2 (secondary)   +-----------+
| Gateway 1 | <-------------------- | Router | ----------------------> | Gateway 2 |
+-----------+                       +--------+                         +-----------+
```
However, this also means that if `ether1` stays up, then traffic will
happily be routed through the primary route.  In particular, this is the
case if there is, for instance, a DSL modem or an ONT connected to an
ethernet port of the MikroTik router and the DSL connection or the fiber
connection fails.  So, assuming that uplink 1 uses, say, a DSL modem, the
above network graph would be a little more accurate like this:

```
+-----------+
| Internet  | <------------------+
+-----------+                    |
  ^                              |
  |                              |
  |                              |
+-----------+                  +-----------+
| Gateway 1 |                  | Gateway 2 |
+-----------+                  +-----------+
  ^                              ^
  | Uplink 1                     | Uplink 2
  |                              |
+-----------+  Ethernet link   +-----------+
| DSL Modem | <--------------- |  Router   |
+-----------+                  +-----------+
```
This illustrates the basic problem in such a setup: If the DSL link
(labeled `Uplink 1` in the graph) fails, then the ethernet link from the
router to the DSL modem is still up.  This means that the router will not
detect the failed uplink, the primary route will still hold, and there will
not be a failover.  In other words, there will be no internet connectivity
for the router (or any networks it routes).

To counter this, MikroTik routers have the concept of gateway checking. To use gateway checking, the primary default route can be defined like this:

```
/ip route add gateway=1.2.3.4 distance=1 check-gateway=ping comment="Primary route"
```
This makes the router ping the gateway IP every 10 seconds, and if two
consecutive pings get lost, then the gateway is marked as `unreachable`.
This means that if the DSL link fails, then after a little bit of time, the
router will still notice that the gateway IP (which is located on the far
side of the DSL link) does not respond, and the primary route will not be
used any more.  Furthermore, as soon as the gateway IP is reachable again,
the primary default route will become active again.  So far, so good.

However, there is still a problem in practice these days.  A number of link
failures occur further upstream, somewhere between the gateway and the
internet.  This means that in the above example, both the ethernet and the
DSL links are up and running, and gateway 1 is pingable at `1.2.3.4`. This
means that, again, the router will not detect the broken internet
connection.

To address this issue, recursive routes can be used. (Alternatively, a script could be written that regularly pings an arbitrary host on the internet that is expected to be up, and if that host is not pingable, assumes that the primary uplink has failed.)

To use recursive routes, a watchdog host is needed that is located
somewhere on the internet and is expected to be up and, crucially,
pingable.  As an example, we will use CloudFlare's DNS server, found at
`1.1.1.1`, but really any IP address that can reliably be expected to be
pingable will do.

First, and counterintuitively, we define the primary default route to use the watchdog host as gateway:

```
/ip route add gateway=1.1.1.1 distance=1 check-gateway=ping comment="Primary virtual route"
```
Of course, `1.1.1.1` is not a valid gateway, and so this would never work:
Traffic going through the default route would be sent to `1.1.1.1`, which
is not directly attached to the router, so it would not be reachable except
through the default route, which would result in a loop.  I am not sure if
this loop would finally be detected by the router, but even if it did, the
packets would end up being routed through the secondary default route to
1.1.1.1, which would obviously drop them, so it would not work either way.

An additional route is necessary, telling the router to use `1.2.3.4` (the
actual gateway 1) as gateway for all traffic for `1.1.1.1`:

```
/ip route add dst-address=1.1.1.1/32 gateway=1.2.3.4 scope=10 comment="Primary route"
```
This way, traffic going to the primary default route will be shipped to
`1.1.1.1`, and since `1.1.1.1` is not directly attached, the routing table
is consulted again.  However, this time around, there is a specific route
for `1.1.1.1` which pushes traffic through to `1.2.3.4`, so there is no
loop.  Instead, traffic eventually goes through gateway 1 by default.

If, however, any part of uplink 1 breaks, then `1.1.1.1` is not reachable
any more, which is detected by the router, and the first default route will
be invalidated, making the secondary default route active.

Of course, there is still another problem with this setup: What if the
watchdog host, `1.1.1.1`, breaks or is not pingable for another reason, but
"our" internet uplink 1 still works fine?  In that case, the primary
default route would also be invalidated, which is not desired.  To counter
this, multiple watchdog hosts can be used.  As an example, we will use
`1.1.1.1`, `8.8.8.8` and `9.9.9.9`:

```
/ip route
add gateway=1.1.1.1 distance=1 check-gateway=ping comment="Primary virtual route A"
add gateway=8.8.8.8 distance=1 check-gateway=ping comment="Primary virtual route B"
add gateway=9.9.9.9 distance=1 check-gateway=ping comment="Primary virtual route C"
add dst-address=1.1.1.1/32 gateway=1.2.3.4 scope=10 comment="Primary route A"
add dst-address=8.8.8.8/32 gateway=1.2.3.4 scope=10 comment="Primary route B"
add dst-address=9.9.9.9/32 gateway=1.2.3.4 scope=10 comment="Primary route C"
```
This way, all three of the defined watchdog/virtual gateway hosts must fail at the same time in order for the primary uplink to be considered invalid.


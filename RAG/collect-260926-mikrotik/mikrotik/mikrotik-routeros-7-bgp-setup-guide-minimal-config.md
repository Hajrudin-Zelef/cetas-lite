---
id: collect-260926-mikrotik/mikrotik/mikrotik-routeros-7-bgp-setup-guide-minimal-config
title: "mikrotik-routeros-7-bgp-setup-guide-minimal-config"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/mikrotik-routeros-7-bgp-setup-guide-minimal-config.md
source_anchor: ""
source_lines: [1, 156]
sha256: 164141f6543ad146b6afef0d8b0490e28f646529b7d40d1664703f0b707edcb8
---

# mikrotik-routeros-7-bgp-setup-guide-minimal-config

This guide explains how to bgp peer a router and announce a prefix to an upstream with Mikrotik Router OS 7. This guide doesn’t cover the “IP” part and firewall part. We highly recommend that you have those two setup before continuing to setup BGP. Those guide is made to work with Servperso service only, no support / help can be provided for third-party services.

## Sample info used over this guide.

Those information has to be replaced by your own during the process.

### VM ip

IPv4: 45.154.99.193/27

IPv6: 2a0c:b640:10::4:193/112

GWv4: 45.154.99.222 AS34872

GWv6: 2a0c:b640:10::4:ffff AS34872

### RIPE allocated ressource

ASN: AS208210

IPv4: 192.0.2.0/24

IPv6: 2001:db8:1234::/48

## Configuration

On ROS7, all route have to be inserted on the local routing table. Before with ROS6 a tab called advertisement permit to announce without having to insert them in the main routing table.

### Filtering

The first step is to build a filter to ensure you won’t leak any prefix learned from someone else up to the upstreams / ixp. This is called a “route leak” and create a few troubles. The simplest way we can recommend for a small network is to fully list all prefix allowed to go over. For more advanced routing policy, we recommend that you read the full documentation on mikrotik wiki.

We build a multi-protocol route filter on our guide. ROS7 now handle multiple affinity on single session (both V6 + V4 on a single session).

Go to routing > filters > tab rule then add a new rule with + icon to add a new filters.

BGP-OUT: Filter used to export your prefix to internet. BGP-IN-DEFAULT: Import default route only (useful for single homed asn). BGP-IN-FULL: accept a full view from your upstream provider. Servperso can also export a full view or a default route up to you on single request.

Here some sample filtering useful for your asn

And export from routers itself:

```
/routing filter rule
add chain=BGP-OUT disabled=no rule="if ( afi ipv4 and dst == 192.0.2.0/24 ) {\r\
    \n accept;\r\
    \n}\r\
    \nif ( afi ipv6 and dst == 2001:db8:1234::/48 ) {\r\
    \n accept;\r\
    \n}\r\
    \n\r\
    \nreject;"
add chain=BGP-IN-DEFAULT disabled=no rule="if( afi ipv4 and dst == 0.0.0.0/0 ) {\r\
    \n accept;\r\
    \n}\r\
    \nif( afi ipv6 and dst == ::/0 ) {\r\
    \n accept;\r\
    \n}"
add chain=BGP-IN-FULL disabled=no rule="accept;"
```
### Insert your resources on the local routing table

As explained earlier, Mikrotik require to have all route inserted on the routing table; It was part of the ROS7 routing engine rewrite. So, on next step, we insert your resource on local routing table as “nullrouted” prefix. It also prevents route loop for unused resources. Used part of your block got “more specific” and override the null route. (More specific is always used). If you use your /24 directly on a single interface without subdividing, you can skip this step.

Go to /ip/routes (or ipv6 routes) then add a new route.

Or the CLI version

```
/ip route
add blackhole disabled=no dst-address=192.0.2.0/24 gateway="" routing-table=main suppress-hw-offload=no
/ipv6 route
add blackhole disabled=no dst-address=2001:db8:1234::/48 gateway="" routing-table=main suppress-hw-offload=no
```
### Configure BGP session

Go to /routing/bgp tab connection then use add button. Create one session per protocol. You need to do configuration in general, extra and on filter tab. Some parameters can be defined as a template. For that you can use the template tabs. This guide doesn’t talk about that to keep things simple. But definitely something to investigate if you manage a lot of bgp sessions.

#### General tab

#### Extra tab

On extra tab, we need to redistribute static and connected (no worries, route is filtered by filters). That part permit to reuse the static route we define earlier on the routing table and propagate them over BGP.

#### And filter tab

This part is used to apply filters, it reuses the filter we define earlier. It’s an important step to avoid route leaks.

#### Command line version

```
/routing bgp connection
add address-families=ip as=208210 disabled=no hold-time=infinity input.filter=BGP-IN-DEFAULT local.address=45.154.99.193 .role=ebgp name=SERVPERSO_v4 output.filter-chain=BGP-OUT .redistribute=connected,static remote.address=45.154.99.222/32 .as=34872 \
    routing-table=main
add address-families=ipv6 as=208210 disabled=no input.filter=BGP-IN-DEFAULT local.address=2a0c:b640:10::4:193 .role=ebgp name=SERVPERSO_v6 output.filter-chain=BGP-OUT .redistribute=connected,static remote.address=2a0c:b640:10::4:ffff/128 .as=34872 routing-table=\
    main
```
#### BGP affinity – Router OS 7 BGP Multithreading

Ros7 now handle multithreading features for BGP. Based on your hardware, we highly recommend using them if you receive or send a full view. We also recommend that you use them only if your router hardware has multiple CPU. If you want to create a separate process on bgp session, you can use input and output affinity and set them to “alone”. For a default route, we don’t recommend changing the affinity. This guide cover a subset about bgp multi threading. Consult mikrotik wiki for more

Mikrotik also provide a few diagnostic command.

/routing/stats/process/print

And here a few examples from one of our routers.

### Full-command export

```
/ip address
add address=45.154.99.193/27 interface=ether1 network=45.154.99.192
/ip route
add blackhole disabled=no dst-address=192.0.2.0/24 gateway="" routing-table=main suppress-hw-offload=no
/ipv6 route
add blackhole disabled=no dst-address=2001:db8:1234::/48 gateway="" routing-table=main suppress-hw-offload=no
/ipv6 address
add address=2a0c:b640:10::4:193/112 advertise=no interface=ether1
/routing bgp connection
add address-families=ip as=208210 disabled=no hold-time=infinity input.affinity=alone .filter=BGP-IN-DEFAULT local.address=45.154.99.193 .role=ebgp name=SERVPERSO_v4 output.affinity=alone .filter-chain=BGP-OUT .redistribute=connected,static remote.address=\
    45.154.99.222/32 .as=34872 routing-table=main
add address-families=ipv6 as=208210 disabled=no input.filter=BGP-IN-DEFAULT local.address=2a0c:b640:10::4:193 .role=ebgp name=SERVPERSO_v6 output.filter-chain=BGP-OUT .redistribute=connected,static remote.address=2a0c:b640:10::4:ffff/128 .as=34872 routing-table=\
    main
	
/routing filter rule
add chain=BGP-OUT disabled=no rule="if ( afi ipv4 and dst == 192.0.2.0/24 ) {\r\
    \n accept;\r\
    \n}\r\
    \nif ( afi ipv6 and dst == 2001:db8:1234::/48 ) {\r\
    \n accept;\r\
    \n}\r\
    \n\r\
    \nreject;"
add chain=BGP-IN-DEFAULT disabled=no rule="if( afi ipv4 and dst == 0.0.0.0/0 ) {\r\
    \n accept;\r\
    \n}\r\
    \nif( afi ipv6 and dst == ::/0 ) {\r\
    \n accept;\r\
    \n}"
add chain=BGP-IN-FULL disabled=no rule="accept;"
/system identity
set name=router-lab-servperso
```
## How to troubleshot

To check what’s your announcing, you can use a few diagnostic command and tools.

### Show route advertised on specific session

/routing/bgp/advertisements/print where peer=SERVPERSO_V4-1

This example come from another router, but normally you retrieve only your prefix filtered earlier on that list.

### Troubleshot over servperso tools

Servperso have a few useful tools to directly troubleshot (lookinglass, bgp session statistics, …). You can access our guide to troubleshoot over the link below.

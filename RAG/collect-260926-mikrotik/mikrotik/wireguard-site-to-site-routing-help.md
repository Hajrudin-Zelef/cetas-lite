---
id: collect-260926-mikrotik/mikrotik/wireguard-site-to-site-routing-help
title: "wireguard-site-to-site-routing-help"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/wireguard/wireguard-site-to-site-routing-help.md
source_anchor: ""
source_lines: [1, 62]
sha256: 021b04826b780a9e08a17f049009d432b1fc14310f377d85184e5f283049f519
---

# wireguard-site-to-site-routing-help

Yikes, I use no wizards, just do it manually LOL, and normally for keys one just puts  “++++++” or something never the real keys.

Okay that gives me a bit of a sense of what you are doing.

Interesting, in 4 router scenario, its rare to have each  one have a public IP.

Normally its one, so what is done is

ONE server Router at handshake

THREE client Routers  at handshake

Now if two of the four have public IPs what I recommend is a backup separate wireguard network

In this case, assuming Primary router goes down,  then make the second one with public IP the primary

ALTERNATE  server Router at handshake

Two client Routers at handshake.

It would appear you have all four on  same network, with all the same listening port in settings, all are listening on the same port in the input chain as well.

On the peer settings, each router  has three  peers.

I think I understand what you want to throw into the mix, you as admin via a remote laptop, for example, have connected to one of the routers and want the added ability to reach any subnet LAN or any other router for config purposes.

Not sure why you want to reach  one through the other as you  can simply DIRECTLY connect to the router you desire.

For example on wireguard app on my iphone, I would simply make four wireguard configs,  one for each router…  but its possible so why not think about it.

Now if one wants to wireguard and connect to ONE router and then reach the other three,  I know how to do that in the ONE server scenario, but this is different and will have to think about it.

By the way this is apparently  called FULL MESH or ( full mess ) TOPOLOGY.

As long as there is no requirement for any local subnet users  to use the internet of a different router, this should work just fine.  The admiin on the laptop coming into any router depending on firewall rules can access internet if desired…

LOGIC:

1. One connects to R1 via wireguard and am now at the LAN side of R1  ( used the ios wireguard app )
2. I want to reach R4, which has a specific wireguard IP ( via winbox likely ),  or I use the gateway of a subnet on R4,  to attempt to reach config of R4 ( via browser likely )
 **Routes**

- R1 knows that the wireguard address of R4, so nothing needs to be made
- R1 knows nothing about remote subnet on R4, so one needs a route made to that subnet
 **FW Rules**
One has to allow traffic that left the tunnel go back into the tunnel,   which is what I call a relay rule…

**CONCLUSION/SOLUTION.**

On each router add all the possible routes to other subnets.

ex. **R1**

*add dst-address=lansubnetR2 gateway=wireguard1
add dst-address=lansubnetR3 gateway=wireguard1
add dst-address=lansubnetR4 gateway=wireguard1*

On each router add a relay rule that allows traffic to exit the tunnel and then renter the tunnel.

*add action=accept chain=forward comment=“relay wg”  **in-interface=wireguard1  out-interface=wireguard1***

Do this  on all your routers and thus you as remote admin should be able to wireguard into any of the four specific routers and reach any other subnet or router to config.

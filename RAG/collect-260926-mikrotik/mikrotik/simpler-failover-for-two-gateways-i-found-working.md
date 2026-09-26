---
id: collect-260926-mikrotik/mikrotik/simpler-failover-for-two-gateways-i-found-working
title: "simpler-failover-for-two-gateways-i-found-working"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/simpler-failover-for-two-gateways-i-found-working.md
source_anchor: ""
source_lines: [1, 104]
sha256: 18225caf30bc980e31c901cb7cad5271e2c1665d0161a99574867f2d26410726
---

# simpler-failover-for-two-gateways-i-found-working

Hey,

like many others I was wondering how to accomplish a simple failover with two Gateways (here: DSL and LTE) with MikroTik involved.

Searching the Internet and this Board, all I was able to find was “Recursive Routes” with checking e.g. 8.8.8.8 as a “Gateway”.

This was not working at first and I wasn’t happy with recursion in the routes so I managed to get the task done with another way I was not able to find anywhere while searching, so I’m sharing this:

Done this on RB5009 yesterday - in Winbox:

**1. Prerequirements:**

- Network with DHCP done by MicroTik (in this case: 192.188.1.0/24)
- Standard Gateway in DHCP will be the MikroTik (here: 192.168.1.2)
- Internet available at (for Example) 192.168.1.1 (in this case DSL)
- Internet available at (for Example) 192.168.1.250 (LTE-Modem)

**2. Routing:**

- Standard Route 0.0.0.0/0 set to 192.168.1.250 with Distance 1 comment=LTE-Failover → (keep it DEACTIVATED)
- Standard Route 0.0.0.0/0 set to 192.168.1.1 with Distance 2

**3. Go to ROUTING → TABLES**

- Create a Routing Table named (for Example) “DSL” - check FIB

**4. Go To IP → ROUTES → Click +**

- Dst,Address: 0.0.0.0/0
- Gateway: 192.168.1.1 (your Primary Gateway)
- Routing Table: Select above created ROUTING TABLE (here: “DSL”)

**5. Go to IP → FIREWALL → Tab MANGLE**

Create a MANGLE-Rule:

- Tab → GENERAL
 – Chain: output
 – Dst.Address: 8.8.8.8
– Protocol: 1 (icmp)
- Tab → ACTION
 – Action: mark routing
– New Routing Mark: Select above created ROUTING TABLE (here: “DSL”)

**6. Go to TOOLS → NETWATCH**

-Tab → HOST

– Create a Netwatch Host:

— Host: 8.8.8.8

— Type: icmp

— Interval: 00:00:30

— Timeout: 5.00

-Tab → Down

/ip route enable [find comment=LTE-Failover]

-Tab → Up

/ip route disable [find comment=LTE-Failover]

What’s this doing?

We were creating TWO STANDARD ROUTES for Traffic leaving the local network to the internet.

The secondary route (in this case LTE) has a higher priority (say: “lower distance”) but is kept disabled.

By creating a second Routing Table and a firewall mangle-rule we will force the ICMP-Request to 8.8.8.8 through the primary gateway (in this case: DSL).

Netwatch is able to perform scripts if the host becomes unavailable through the primary route.

The DOWN-script will enable the secondary route which will become active immediately due to the higher priority (say: “lower distance”)

All traffic to Internet will go through the secondary route now.

Netwatch will still check every 30 seconds pinging 8.8.8.8 forced to the primary gateway as of our mangle-rule.

If 8.8.8.8 will be available again through the primary gateway the UP-script will deactivate the secondary route again.

All traffic will go through the primary route again.

Please note that you will not be able to use the host used ( in this case 8.8.8.8 ) as an upstream DNS-Server, since it won’t work when LTE kicks in.

I’m not an MikroTik-Expert by far, still learning, but I found this way a bit more straight-forward and understandable than the “recursive routes” many tutorials show up with. Also you can extend the scripts by sending EMails out (configure TOOLS → EMAIL first) by adding for example:

*:delay 10
/tool e-mail send to=youremail@host.com subject=“DSL is DOWN!!” body=“DSL inactive - LTE active”*

at the end of the script.

Still, I was wondering, if this is already documented somewhere, that’s why I posted it here. Please disregard or close if this is “too obviuous” or “already well documented” 

Have a great day, everyone, many greetings,

Martin!

*EDIT*: I was choosing this variant for failover over the “recursive Routes” because I’d like to maintain more control about failover.

The script can be extended, and getting an EMail, WHEN failover happens is quite nice. Also we could add even MORE Netwatch-hosts. For example: The FIRST netwatch checks 8.8.8.8 and if this fails a script may ENABLE the SECOND Netwatch-Host to check, just to verify, and only after BOTH would fail, the secondary route may kick in. I think this has more opportunities at all

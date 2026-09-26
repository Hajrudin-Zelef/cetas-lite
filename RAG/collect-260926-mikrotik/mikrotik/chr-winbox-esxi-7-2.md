---
id: collect-260926-mikrotik/mikrotik/chr-winbox-esxi-7-2
title: "chr-winbox-esxi-7"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution"]
source: docs/RAG/lot-mikrotik/tools/chr-winbox-esxi-7.md
source_anchor: ""
source_lines: [238, 252]
sha256: 1e201ac934999070e52f4d406d16bf41d86dc083b51f9bd42cc65c4f8e2de02e
---

# chr-winbox-esxi-7

- Some are only CHR routers for distribution to customers.
- Some are bandwidth control to/from my customers.
- Some are DHCP servers
- Some are used for re-directing deliquent customers to a web server ( web site - pay your bill - call us )
- Some are NAT444 ( not normal NAT or NAT44 ) ( where I can use 8-Live-IP-Addresses to CGN-NAT a 100.64.x.y/21 network ( NAT444 on a CHR is blazing fast and reliable )
- Some are btest servers ( one used to test/verify my customers get the speed they are paying for - and another btest server to test how fast a customer can go without any bandwidth rate-limiting )
- Some are dedicated for EoIP tunnels ( bridge remote networks - example customer with 3 sites wants a common network )
- Some are for BGP
- …

All of my CHRs run average 1-percent to a maximum of 15-percent of CPU use ( winbox CPU shown on the web gui )  and I average about 8 to 9 Gig of traffic to my customers.

FYI - I assume you know about the CHR btest server I manage.  One CPU is allocated to this public btest server.  When I btest using udp send-or-receive to 127.0.0.1 ( itself ) I get faster than 200-Gig.  And that CHR is on a busy VmWare ESXi server with 20+ other Windows/Linux/PfSense servers running on it also.

North Idaho Tom Jones

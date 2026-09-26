---
id: collect-260926-mikrotik/mikrotik/questions-176118-mikrotik-os-for-load-balancing-41325af5
title: "questions-176118-mikrotik-os-for-load-balancing-41325af5"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/questions-176118-mikrotik-os-for-load-balancing-41325af5.md
source_anchor: ""
source_lines: [1, 16]
sha256: f086414fac3c5e272d5450d3930d51ed4a5206f88abc12a6587432835926499e
---

# questions-176118-mikrotik-os-for-load-balancing-41325af5

Is it possible to load balance with a mikrotik routerOS between 2 ISPs and the mikrotik is in a LAN?
- 
        Load balancing is a bit too wide definition. What do you want to do exactly?Istvan– Istvan2010-08-30 15:39:21 +00:00Commented Aug 30, 2010 at 15:39
2 Answers 2
From my own workbench: http://www.mikrotik-routeros.com/?p=12
This config shows you how to use the RouterOS PCC mangle rules setup a basic load balancing system that holds static connections between IP addresses once established and also forces inbound traffic on any WAN interface to pass back out the same (stopping any triangular routing issues).
- 
        +1. Bvest yo ucan get - norte that this is not doubling bandwidth for one download or a set of downlaods that result on beign routed to the same external iterface.TomTom– TomTom2011-10-12 12:53:53 +00:00Commented Oct 12, 2011 at 12:53
There are various forms of load balancing and the answer depends very much on the topology of your network, but the short answer your question yes it is possible to load balance between two ISPs if the MikroTik is on the LAN.
In the following example, I am assuming a RouterBOARD or a PC with at least three network card interfaces.
LAN Interface = 192.168.0.1/24
ADSL router to ISP1 = 192.168.1.1/24
ADSL router to ISP2 = 192.168.2.1/24
Gateway for all PCs is the LAN router, 192.168.0.1. When packets hit the router, depending on your chosen load balancing methodology (ECMP, NTH, or PCC) you would forward the packets to either ISP1 via the IP for ISP1 or forward the packets to ISP2 via the IP for ISP2.
As an alternative to forwarding the load balanced packets to the IP itself for example when you are using PPPoE interfaces configured on the MikroTik router, you could forward packets to the interface instead.
If you give us some hints about your topology we can answer the question better.

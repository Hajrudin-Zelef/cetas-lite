---
id: collect-260926-mikrotik/mikrotik/questions-5657-ebgp-redistribution-on-mikrotik-ccr-a073a086
title: "questions-5657-ebgp-redistribution-on-mikrotik-ccr-a073a086"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-5657-ebgp-redistribution-on-mikrotik-ccr-a073a086.md
source_anchor: ""
source_lines: [1, 8]
sha256: 507cfe84a93cb60a737506c25080a45ad543a9671b715efab5facaa7042ee15e
---

# questions-5657-ebgp-redistribution-on-mikrotik-ccr-a073a086

I assume you already have eBGP up and running between you and your transit providers, and that you have only a default route, not the full internet BGP table. I assume also that you and your customer each have a public AS and a public subnet.
So, I will mention only minimum needed config to enable you advertise your customer AS and subnet, which can apply also for your customer, noting the router model they are using.
Take a look at  http://wiki.mikrotik.com/wiki/Manual:Routing/BGP and set the following :
remote-address a.b.c.d <== IP address of customer peer router running BGP with you
remote-as xxxxx <== customer public AS number
network x.y.z.0/24 <== any specific subnets from your AS that need to be advertised to your customer
default-originate always <== provide default route to your customer via BGP (your customer should not use this command on their end)
Again, note that this is a minimum config that will allow you to initiate BGP session with your customer and advertise their AS and subnets to your upstream providers, there are other elements that should be added, e.g. I prefer creating a VRF for the customer on your router, also, you could adjust local-preference or AS prepending in order to select which of your 2 transit providers would be primary preferred by BGP and which would be secondary, among other BGP features.

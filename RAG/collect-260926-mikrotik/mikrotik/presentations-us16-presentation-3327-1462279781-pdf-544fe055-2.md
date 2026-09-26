---
id: collect-260926-mikrotik/mikrotik/presentations-us16-presentation-3327-1462279781-pdf-544fe055-2
title: "presentations-us16-presentation-3327-1462279781-pdf-544fe055"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "revenue"]
source: docs/RAG/lot-mikrotik/RouterOS/presentations-us16-presentation-3327-1462279781-pdf-544fe055.md
source_anchor: ""
source_lines: [182, 256]
sha256: 79e9881ecd600e5b8732ac09179a0b7ddbb67dae39d00da2af043b0be0354204
---

# presentations-us16-presentation-3327-1462279781-pdf-544fe055

• Implement on a low priority segment –  The very worst thing you can do is 
implement a complex protocol like MPLS for the first time using your highest 
priority subscribers as Guinea pigs. Pick a low impact segment of the network 
to attempt to take your lab into production. For a WISP this can be as simple as 
one tower router and one core router.  
 
• Don’t be afraid to use new hardware – We often see WISPs go to great lengths 
to try and convert an existing production router into an MPLS capable router. 
Sometimes this is feasible and sometimes it’s actually less expensive to build 
MPLS for a small segment with new gear. Once everything is running and 
stable, a migration can be planned for the rest of the network. Adding a new 
tower can be a great opportunity to bring an MPLS trial online since new 
hardware is already required.

Business case for MPLS – part 1  
• Increase revenue using existing infrastructure – MPLS enabled connections 
will sell for much more than an basic Internet pipe due to the provider offering 
managed routing and QoS. It is not uncommon for an MPLS circuit to sell for 4 
times the cost of a non MPLS circuit at the same speeds.

Business case for MPLS – part 2  
• Lower OPEX cost by increasing agility – MPLS enabled networks are more agile and can 
solve complex problems that a customer may require more efficiently than a non MPLS 
network.  
 
• Build private customer networks outside your service area – If a business in your 
footprint has several locations in your service area, but one or more locations outside of 
the service area, you can still capture that sale and manage the entire private WAN 
infrastructure for that company. MPLS can be used to facilitate this. 
• Use MPLS over tunnels to extend the VRF over the Internet. 
• Purchase transit from the last mile provider and use MPLS to connect all the 
locations.  
• Manage the entire solution for a premium fee.  
• Customer gets a Tier 1 business class product with more personalized service. 
 
• Resell your network to Tier 1 and 2 providers – with MPLS enabled, you can quickly 
hand off last mile Layer 2 or 3 circuits to other providers for redundancy or primary 
transit. 
 
• Transport legacy technologies – MPLS can encapsulate legacy technologies like ATM, 
Frame Relay and PPP and deliver them with a much lower cost than traditional legacy 
end to end service.

MPLS WISP Use case #1 – PPPoE aggregation with 
VPLS  
• Problem – Need multiple L2 domains aggregated over a routed network for PPPoE 
• Solution – Use VPLS to extend L2 from the BRAS to the last mile.

MPLS WISP Use case #2 – BGP Peering with a full 
table  
• Problem – Need to sell public BGP transit to multiple customers without impacting BGP edge  
performance 
• Solution – Use VPLS to hand off a /30 that directly connects the CE and PE routers without the need  
for a full table in the core

MPLS WISP Use case #3 – Private L3 Transit (L3VPN)  
• Problem – Need to sell private transit to multiple customers and keep them isolated 
• Solution – Use L3VPN to build private VRF segment for each customer.

MPLS WISP Use case #4 – Private L2 Transit (L2VPN)  
  
• Problem – Need to sell private L2 transit to multiple customers over an L3 networks 
• Solution – Use L2VPN to build private L2 pseudowire for each customer.

Questions? 
The content of this presentation will be available at 
mum.mikrotik.com 
 
Please come see us at the IP ArchiTechs booth in the Exhibitor Hall 
 
Email: kevin.myers@iparchitechs.com 
Office: (303) 590-9943 
Web: www.iparchitechs.com 
 
Thank you for your time and enjoy the MUM!!

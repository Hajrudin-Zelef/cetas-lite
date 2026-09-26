---
id: collect-260926-mikrotik/mikrotik/mpls-bgp-and-ospf-design-for-wisp-3
title: "mpls-bgp-and-ospf-design-for-wisp"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/mpls-bgp-and-ospf-design-for-wisp.md
source_anchor: ""
source_lines: [369, 396]
sha256: a23ad48a79dd362fc4d608228bb15b416ec70153b6b88981320b53728ba48559
---

# mpls-bgp-and-ospf-design-for-wisp

Before we enable MPLS at a site, we set up a test VPLS tunnel across the link with temporary IP addresses on both ends (in the same /30 or /24), and we try pinging with size 1500 do-not-fragment from one address to the other. If we want to provide more than 1500 MTU (ex. 1508 for RFC 4638 PPPoE) then we temporarily change the IP MTU on the VPLS tunnel to 1508, and test the ping with 1508. That way, we can verify that all of the devices in between are properly passing the jumbo frames before we put customers on the service.

             
            
           
          
            
            
              
This will help you some on MTU sizing.

https://mum.mikrotik.com//presentations/US16/presentation_3327_1462279781.pdf

             
            
           
          
            
            
              How much horsepower does it take to run MPLS/VPLS.  We set up a test bed of 4 CRS112-BP-4S-IN switches which is what we might normally use at a smaller micro pop site where we have small 14-16" boxes and just need POE to the radio and UPS and they failed terribly.  They could only do 50-60 Mbps before the CPU got maxed out which is terrible when they can do handle a ton more otherwise in a normal configuration.  At normal tower sites where we have rack space we typically run CCR routers.  But even with these as I don’t have a test bed I would be very concerned of the ability to handle it, even more so at your core where you might have 40-100 VPLS tunnels terminating on one CCR

             
            
           
          
            
            
              You definitely don’t want to do it with a small CRS. Look at using a 3011 at smaller sites and CCR at others.

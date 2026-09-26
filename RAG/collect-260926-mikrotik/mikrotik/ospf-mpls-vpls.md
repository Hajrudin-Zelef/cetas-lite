---
id: collect-260926-mikrotik/mikrotik/ospf-mpls-vpls
title: "ospf-mpls-vpls"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/ospf-mpls-vpls.md
source_anchor: ""
source_lines: [1, 41]
sha256: 5796edce7f45e8ac4c909e7c7451b746637beda24bacb6147baa25e162192d6d
---

# ospf-mpls-vpls

vpls one side is up and other end is down help please, i have checked LDP and MTU on both routers but still its the same and i m stuck

             
            
           
          
            
            
              We need more info  so we can help.

Please state config, cenario, etc…

             
            
           
          
            
            
              Are you sure you have LDP enabled on every LSR between LER’s?

             
            
           
          
            
            
              
Check IP address (/32) of both ends appears on routing table

             
            
           
          
            
            
              
Be sure that the transport address is set on the LDP interface or global LDP options. VPLS requires dynamic targeted LDP sessions which must have a consistent transport address set.

Here is an example of working config for VPLS.

https://stubarea51.net/2018/04/23/wisp-design-building-highly-available-vpls-for-public-subnets/

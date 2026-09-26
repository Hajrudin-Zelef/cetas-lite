---
id: collect-260926-mikrotik/mikrotik/what-are-the-best-practices-for-securing-a-mikrotik-router-from-external-threats-2
title: "what-are-the-best-practices-for-securing-a-mikrotik-router-from-external-threats"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["exploit"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/what-are-the-best-practices-for-securing-a-mikrotik-router-from-external-threats.md
source_anchor: ""
source_lines: [193, 243]
sha256: d446c49d74f633e088df9c4390b55d32668592ff8da58d63273953f08a38cea2
---

# what-are-the-best-practices-for-securing-a-mikrotik-router-from-external-threats

Again if you want to use the mac-backdoor for whatever reason, good for you.

Personally I prefer (and recommend) to disable ALL “features” and ONLY enable those I actually need (which unfortunately none of the NOS out there (yes Im looking at you Cisco, Arista, Juniper, VyOS, Mikrotik, HPE, Aruba etc) do these days where most junk is enabled by default so the box will happily inform its surroundings (and even the internet) which version it runs etc).

             
            
           
          
            
            
              That was a REMOTE exploit possible only on routers that had “wrong” firewall settings and thus allowed access from the outside (WAN).

By setting:

/tool mac-server mac-winbox set allowed-interface-list=none

you are disabling it completely, also from LAN.

It creates a big inconvenience in normal management.

It makes much more sense in any normal setup to have a dedicated interface for management and allow winbox on that interface.

             
            
           
          
            
            
              Again, disabling not needed features is NOT an “inconvenience” rather the opposite.

             
            
           
          
            
            
              Actually, security always comes with some level of inconvenience.

It’s up to the admin to decide what he/she values most: convenience or security.

Accept inconvenience then or get rid of humans … 

we all should be quite aware the human factor (and it’s accompanying need for convenience) is the biggest thread for anything related to security

Access your router remotely directly

or

first fire up your VPN, then access your router

Most will take the first option. Simply for convenience.

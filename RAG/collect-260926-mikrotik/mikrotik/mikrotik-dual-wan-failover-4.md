---
id: collect-260926-mikrotik/mikrotik/mikrotik-dual-wan-failover-4
title: "mikrotik-dual-wan-failover"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/mikrotik-dual-wan-failover.md
source_anchor: ""
source_lines: [321, 333]
sha256: 88038982b83ee9a372c9c201917e8af7415518e8c59d971fac990baf21c0c3e8
---

# mikrotik-dual-wan-failover

I’ll stay polite so I won’t translate any of the Czech sayings related to this kind of statements, but it would at least take long to happen (if at all possible because I’m not deep into the recursive next hop search algorithm, so maybe there is some reason which excludes using the interface name as gateway). So if you want it now, use the workaround suggested or use a script. In fact, the next-hop search mechanism was also not originally intended for the failover use. Which BTW also means that the failover may happen up to about 10 seconds after the active WAN path breaks because this is how often the check-gateway pings are sent.

             
            
           
          
            
            
              
For PPPoE (used at your WAN1), there is a script-less way which @Sob has described: you create a copy of /ppp profile named default, give it a name like my-pppoe-profile, and set the remote-address item in that new profile to some private address which isn’t in conflict with any private subnet you use anywhere in your network - say, 10.22.33.44. In /interface pppoe-client configuration, you set the profile item to my-pppoe-profile. And in the individual route(s) to the anchor IP(s) used to monitor PPPoE availability, you use the 10.22.33.44 as a gateway address. This way, the remote-address setting from the /ppp profile my-pppoe-profile overrides the setting which came from the PPPoE server, and so it remains stable even though the PPPoE server sends you a different one each time.


So I’ve tried that  but sadly while trying to connect it says the connection is terminated and it isn’t able to make a connection when I’m using a random remote address.

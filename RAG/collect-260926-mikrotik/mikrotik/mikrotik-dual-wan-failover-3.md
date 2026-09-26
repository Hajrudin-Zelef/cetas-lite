---
id: collect-260926-mikrotik/mikrotik/mikrotik-dual-wan-failover-3
title: "mikrotik-dual-wan-failover"
domain: mikrotik
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/mikrotik-dual-wan-failover.md
source_anchor: ""
source_lines: [230, 320]
sha256: f860729a8b19052850391a259939722dfb7445f62a8557bcb23884793abd0ea3
---

# mikrotik-dual-wan-failover

Regarding the gateway IP provided by DHCP on the second WAN, there is again no point in monitoring that address itself but you need to use it as a gateway to the monitored destination in the recursive next-hop search scheme, and you cannot easily assign an alias to it (well, you can in some cases, but in exactly those cases it is pointless to do that). So if the DHCP server runs on the modem+router combo you’ve got from the ISP, there is a 0.001% chance that the gateway address will ever change; if that box acts as a bridge and the DHCP server is physically located at the other end of the WAN link, the chance that the gateway IP will change is much higher. So in the latter case, you would have to permit the dhcp client to install a default gateway, but you would tell it to set a high *distance* = low priority value to it and copy the address of that gateway to the individual route to the monitored destination using the *script* parameter. So each time a new DHCP assignment arrives, you’d check whether the gateway IP has changed as compared to the previous one and if yes, you’d modify the individual route(s) to the monitored IPs.

             
            
           
          
            
            
              Hello guys,

Thank you all for the precious help. Tonight I had some time to try the things up and everything seemed to work good with one exception. The remote address of the PPPoE  is changing. It seems to be either 5 or 12 but it changes.

So What I’ve done till now:

/ip route

add check-gateway=ping distance=1 gateway=8.8.8.8

add check-gateway=ping distance=2 gateway=8.8.4.4

add distance=1 dst-address=8.8.4.4/32 gateway=192.168.x.x scope=10

add distance=1 dst-address=8.8.8.8/32 gateway=x.x.x.x scope=10


So when the Line 1 fail and then reconnect it may take a different Remote address so the recurse fails as the gateway in line 3 is different.  You have mentioned some kind of a script but isn’t there an easier way to always take the current remote address and to put it as the GW without scripting? It’s really sad that this isn’t working with the pppoe interface. I was able to get the address of the GW from the status bar of the pppoe interface. I wasn’t completely able to understand the method proposed by @Sob’s.

Also I found in an article that the recursive method has the following limitation:

Whatever IP you use as your target is only reachable via the primary route. If the primary route is down, that IP address will be unreachable. If you use 8.8.8.8 to resolve DNS, the DNS service will be down when the primary route is down. Therefore if you use Google for DNS and use 8.8.8.8 as the routing target, you should use a different Google DNS server such as 8.8.4.4 for DNS instead.


The pppoe is using 1.1.1.1 and 8.8.8.8 as DNS and for the ADSL I have to check it because currently I’m using the ADSL address as DNS.

So except the problem with the changing remote address everything seems to work. If I’m able to reslove this too I’d be able to make it with multiple host checks and leave it this way.

             
            
           
          
            
            
              
You’re mixing together the DHCP case with the PPPoE case.

For DHCP (used at your WAN2), there is no other way than a script to get the assigned IP address of default gateway and set it as a gateway in the individual routes to the monitored anchor addresses, but you obviously don’t need it because the gateway IP provided by DHCP on WAN2 does not change.

For PPPoE (used at your WAN1), there is a script-less way which @Sob has described: you create a copy of */ppp profile* named *default*, give it a name like *my-pppoe-profile*, and set the *remote-address* item in that new profile to some private address which isn’t in conflict with any private subnet you use anywhere in your network - say, *10.22.33.44*. In */interface pppoe-client* configuration, you set the *profile* item to *my-pppoe-profile*. And in the individual route(s) to the anchor IP(s) used to monitor PPPoE availability, you use the 10.22.33.44 as a gateway address. This way, the *remote-address* setting from the */ppp profile my-pppoe-profile* overrides the setting which came from the PPPoE server, and so it remains stable even though the PPPoE server sends you a different one each time.




This is normal - for any destination address the routes with the longest, i.e. most exactly matching, *dst-address* prefix are chosen. So if at least one route with *dst-address=8.8.8.8/32* exists and is active, routes whose *dst-address* prefixes also match 8.8.8.8 but are shorter (wider), such as *8.8.8.0/24* or *0.0.0.0/0*, are never chosen for delivery of packets to 8.8.8.8. This has two consequences:

- you must not set *check-gateway=ping* for the individual routes to monitored anchor addresses, because if you do and the gateway becomes unreachable, the route becomes inactive and the check-gateway pings of the routes one level higher in the recursion start taking another route, ruining the idea of using inaccessibility of the anchor address as indication of network path failure
- you cannot use the anchor IP for any other purpose than network path monitoring because the anchor IP has to be inaccessible if the path whose availability it monitors is broken, so you cannot set up an alternative route to the anchor IP.

 
            
           
          
            
            
              Yes I don’t have problem with the WAN2 a its gateway is constant. I’m using the ADSL modem as GW and it won’t change. The route to WAN2 is static. The only thing that is changing is the remote address of the PPPoE which I’m using as WAN1 (main link).

The current set is:

WAN 1 - Optic → media convertor → Mikrotik at eth1

WAN 2  - phone line->ADSL modem ->Mikrotik at eth2

I have full access to the ADSL modem. I’m only not sure which DNS it was using but now on the Mikrotik  I’m using the modem as DNS.That’s why I believe the only problem is the PPPoE with it’s changing GW. I’ll try the proposed workaround for it and I’ll write if there is any success as it’s not still completely clear for me. If it was possible to use the pppoe interface instead of exact GW it would be way easier…

I’m really interested how in fact the current TP-link failover is in fact realized behind the wizard.

             
            
           
          
            
            
              
PPPoE creates a Point-to-Point interface. For all interfaces of this type, there is no actual need to use any address of the remote device because “the remote end of the tunnel” is the only address you need - whatever you send out that interface will end up on the single remote device. This is a difference to Point-to-Multipoint interface where you need an address of a particular device in addition to the name of the interface. For practical reasons, the gateway addresses are configured as IP addresses, which allows to quickly choose the interface by its associated “network” address, and there the IP address of the gateway device is translated into its MAC address.

So a common habit is to use IP address as a gateway even for PPP interfaces although in these cases it actually acts only as an alias to the interface name. The recursive next-hop search needs IP addresses of gateways, that’s a fact you have to merely accept 

But as the “remote” address of a PPP interface plays no other role in the process than the alias of the interface name, it is only meaningful in the local context of the sending device. Thus you may label the PPP interface with any “remote” address you like. And whilst */interface pppoe-client* doesn’t have a direct parameter *remote-address*, it does accept that parameter if provided by means of the *profile* and uses it to override the value provided by the server.





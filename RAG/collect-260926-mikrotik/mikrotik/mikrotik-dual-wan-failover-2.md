---
id: collect-260926-mikrotik/mikrotik/mikrotik-dual-wan-failover-2
title: "mikrotik-dual-wan-failover"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/mikrotik-dual-wan-failover.md
source_anchor: ""
source_lines: [107, 229]
sha256: c70c311cf274d89696ad99aac64518d5f5893ad2b000fb16486dfed5ea781fe4
---

# mikrotik-dual-wan-failover

First of all I want to say that I’m continuing with the article after I have set my settings for both ISPs. Here are the things I’ve done before continuing with the first method from the article.

1. Port 1&2 are set as WAN ports.
2. Port 3-5 are in a LAN bridge.
3. DHCP server for the bridge is set with the needed pool and the router used as DNS.
4. PPPoE-out for the main link is set with the needed credentials (Use peer DNS = true Add default route = false), Ethernet port 1.
5. Created route  Dest. Add. 0.0.0.0/0, GW PPPoE-out.

At this point I’m having internet through the ppoe  and everything works fine.

1. Add new address to the address list for the second ISP 192.169.1.2/24, network 192.168.1.0, Ethernet port 2.
2. Created route  Dest. Add. 0.0.0.0/0, GW 192.168.1.1 via Ethernet port 2.
3. DNS server set to 192.168.1.1 (the ADSL modem), Allow remote request = true.

WAN1 and WAN2 alongside with pppoe are added to the WAN list so the NAT and firewall rules can apply to all of them.

At this point I want to continue with the monitoring of the gateways explained in the article (its method one)

/ip route

add dst-address=8.8.8.8 gateway=PPPoE-out scope=10

add dst-address=8.8.4.4 gateway=192.168.1.1 scope=10

At this point both addresses are reachable

/ip route

add distance=1 gateway=8.8.8.8 check-gateway=ping

add distance=2 gateway=8.8.4.4 check-gateway=ping

And here comes the problem, once I set these two routes they are both unreachable. I know that I’m missing something and probably I don’t need the dest in point 5 and 7 but I’m not able to figure it out. Probably there is some kind of a conflict but I’ve tried everything I could imagine for now and it seems not to work. Could you please give me an advice about where I’m in fact messing the things up.

             
            
           
          
            
            
              First of all, the recursive routing on which the scriptless failover is based does not work if a route’s *gateway* is set to anything else than an IP number anywhere in the recursive chain. So you cannot use the interface name (*PPPoE-out*) as a *gateway* for *dst-address=8.8.8.8*, you have to use the IP address provided by the PPPoE server.

             
            
           
          
            
            
              
Fair enough but it isn’t working even with the static address of the second ISP. I mean the 0.0.0.0/0 with GW 8.8.8.8 is still unreachable.

             
            
           
          
            
            
              When I say “you must use as gateway the IP address provided by the PPPoE server”, I have in mind the address which that PPPoE server provides as a gateway, not the one it assigns to you. Is it what you mean by “static address of the second ISP”?

Normally, where you are a PPPoE client, the server assigns you your own address and indicates its own IP address which you may use as a gateway for anything you want to send via that server. But in most cases, you can use the interface name as well; recursive routing on Mikrotik is one of the exceptions where you can’t. I have seen you have set *add-default-route* in */interface pppoe-client* to *no*, but when you do that, you won’t learn the gateway address. So you have to set it to *yes* for a while to learn the address “manually”, or keep it on *yes* and set *default-route-distance* to e.g. *10* and add a blackhole route with a lower distance, so the final set of default routes would be

```
dst-address=0.0.0.0/0 gateway=8.8.8.8 distance=1
dst-address=0.0.0.0/0 gateway=8.8.4.4 distance=2
dst-address=0.0.0.0/0 type=blackhole distance=9
dst-address=0.0.0.0/0 gateway=the.ip.from.isp distance=10
```

Then, you would use an *on-up* script from a */ppp profile* attached to the */interface pppoe-client* to copy the gateway IP from the route with *distance=10* to the route with *dst-address=8.8.8.8/32*. But it only makes sense to do it this complex way if the PPPoE server doesn’t provide the same gateway IP address each time.

Where you are a DHCP client, you must use the IP address provided by the DHCP server as a default gateway (or use the routing table provided by the DHCP server as Option 121 but that’s out of scope of this).

             
            
           
          
            
            
              There’s also the trick with locally set remote address. Simply put one in PPP profile used by PPPoE client and then use it as gateway. I found it some time ago in this forum and although it looks completely wrong at first (how can I set remote address when I don’t control remote side, right?) it works. The used address is not actually used by anything by default, no packets are sent to it, so it don’t matter what you put there. Importatnt part is that it’s static. And I think it was possible to go even one step further and use 8.8.8.8 as this remote address/gateway and check-gateway=ping with it. I don’t remember the details, it probably had to be done with routing filter to add check-gateway option.

             
            
           
          
            
            
              
So what you are saying is that

- you don’t need to retrieve the real “remote” address from the PPPoE client, so *add-default-gateway* may stay at*no*
- you can assign different “remote” addresses to different PPPoE clients, which makes it possible to use the recursive next-hop search and thus scriptless failover even on several PPPoE connections even if the servers assign the same remote addresses to them

?

             
            
           
          
            
            
              Yes. In other words, if ISP would be giving you random 10.x.y.z every time you connect, you can set static 10.1.1.1 in PPP profile and use that. And it will work, because it’s PPP, a tunnel where you just feed everything into. On ethernet, gateway IP address is used by ARP, but with PPP it’s just a local hint where it is.

             
            
           
          
            
            
              
Ah I was afraid it won’t  be so straight forword with the PPPoE…

I meant that when I’m using the second ISP settings, everything is static i.e.  the adress is 192.168.1.2 and the GW is the ADS modem at 192.168.1.1. Thus with these settings while I set the route with dest 8.8.8.8 throught GW 192.168.1.1 it is reachable and when afterwords I set the dest 0.0.0.0/0 with GW 8.8.8.8 its unreachable.

Sadly I’m afraid that the remote ip is not the same and it may variate (I’ll double check it) which seems to make the things even more complicated as it’s obvoius that if the remote adress is changing the set up won’t work if this adress is not monitored. However I have static IP adresses from both ISPs.

             
            
           
          
            
            
              There is no point in **monitoring** the remote IP, it may even not be up at all on the remote end. For the purpose of identifying a local PPPoE tunnel to use by a gateway IP address, you may assign the local alias to the tunnel’s *remote-address* as per @Sob’s suggestion. For the purpose of monitoring the WAN link transparency, the monitored addresses should be some immortal addresses further in the internet, so instead of checking just the hop between your router and ISP’s PPPoE server, you check the whole path through the ISP up to the internet.


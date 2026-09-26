---
id: collect-260926-mikrotik/mikrotik/ospf-works-than-stops-by-itself-1
title: "ospf-works-than-stops-by-itself"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/ospf-works-than-stops-by-itself.md
source_anchor: ""
source_lines: [1, 245]
sha256: fb2a2ddb5e0fe42b644cf8b34e8e3b2d41b40a5aa2eee6ce00ac1e9cfeb92f1d
---

# ospf-works-than-stops-by-itself

ok here is some backround.

I have 2 routers connected by a simple switch.

10.26.1.1 is backbone router. R1

10.26.1.3 is a wireless AP and PPPoE Server




R1 configuration v4.5

```
/ip address
add address=10.4.1.1/24 broadcast=10.4.1.255 comment="" disabled=no interface=HalkenSector2 network=10.4.1.0
add address=10.26.1.1/24 broadcast=10.26.1.255 comment="" disabled=no interface=HalkenSector1 network=10.26.1.0
network=95.0.133.252
```




```
/routing ospf instance
set default comment="" disabled=no distribute-default=always-as-type-2 in-filter=ospf-in metric-bgp=20 metric-connected=20 metric-default=1 metric-other-ospf=auto \
    metric-rip=20 metric-static=20 name=default out-filter=ospf-out redistribute-bgp=no redistribute-connected=as-type-2 redistribute-other-ospf=as-type-2 \
    redistribute-rip=no redistribute-static=as-type-2 router-id=0.0.0.0
/routing ospf area
set backbone area-id=0.0.0.0 comment="" disabled=no instance=default name=backbone type=default
/routing ospf area range
add advertise=yes area=backbone comment="" cost=calculated disabled=yes range=0.0.0.0/0
/routing ospf network
add area=backbone comment="" disabled=no network=10.26.1.0/24
add area=backbone comment="" disabled=no network=10.4.1.0/24
```

R2 conf v3.30 Routing-TEST

```
/ip address
add address=10.26.1.3/24 broadcast=10.26.1.255 comment="" disabled=no \
    interface=vlan1 network=10.26.1.0
```




```
/routing ospf instance
set default comment="" disabled=no distribute-default=always-as-type-2 \
    in-filter=ospf-in metric-bgp=20 metric-connected=20 metric-default=1 \
    metric-other-ospf=auto metric-rip=20 metric-static=20 name=default \
    out-filter=ospf-out redistribute-bgp=no redistribute-connected=as-type-2 \
    redistribute-other-ospf=as-type-2 redistribute-rip=no \
    redistribute-static=as-type-2 router-id=0.0.0.0
/routing ospf area
set backbone area-id=0.0.0.0 disabled=no instance=default name=backbone type=\
    default
/routing ospf interface
add authentication=none authentication-key="" authentication-key-id=1 cost=10 \
    dead-interval=40s disabled=yes hello-interval=10s instance-id=0 \
    interface=all network-type=ptmp passive=no priority=1 \
    retransmit-interval=5s transmit-delay=1s
add authentication=none authentication-key="" authentication-key-id=1 cost=10 \
    dead-interval=40s disabled=yes hello-interval=10s instance-id=0 \
    interface=wlan1 network-type=ptmp passive=no priority=1 \
    retransmit-interval=5s transmit-delay=1s
/routing ospf network
add area=backbone disabled=no network=10.26.1.0/24
add area=backbone disabled=yes network=0.0.0.0/0
add area=backbone disabled=yes network=95.0.133.0/24
```

OSPF stops working after a few hours. Neigbors are still active but no routes are received or imported.

I tried every possible scenario to solve this problem.

Any ideas ?

             
                
            
           
          
            
            
              Anyone out here ? MT team ?

             
            
           
          
            
            
              An OSPF loopback interface solved the problem. Thanks to MT support but they need to further investigate the issue if this is by design or is a serious bug.

             
            
           
          
            
            
              Did you try both routers with the same version?  If so were you able to duplicate the problem?

             
            
           
          
            
            
              I have created a loopback /32 bridge interface and assigned an IP address to it. Than assigned PPPoE Servers Local address the same IP address as loopback. Now it works and routes doesnt dissapear after some time like before.

As I said, I dont really know why it works this way because it is suggested to me by Mikrotik support after some months of struggling about this problem.

my previous PPPoE configuration was correct network wise either. But was loosing all the routes after 30mins or so.

             
            
           
          
            
            
              i’ve the same issue. Using 4.10, 3.30 mixed. Could you give me more specific configuration?

Regards

NoXy

             
            
           
          
            
            
              Bring your OS versions to the same level and things will improve.  Add in loopbacks and they will get even better.

             
            
           
          
            
            
              Sorry guys, but how do I do this loopbacks. someone can write me an example?

Many thanks

             
            
           
          
            
            
              To fake out a loopback interface just create a bridge interface and don’t add physical interfaces to it. Assign an IP to the bridge and it acts like a loopback.

             
            
           
          
            
            
              Thank you. I’ll try that too.

Sorry for my bad english.


             
            
           
          
            
            
              
Noxy;

If you are loosing routes in a PPPoE environment, you probably need to use the loopback method. I have described it above in detail.

             
            
           
          
            
            
              
I would ask others for their opinion on that statement, is it just a fix for a issue or a possible tweak?

             
            
           
          
            
            
              I note that the configuration at the top of this thread shows both router ids are 0.0.0.0. Its possible that 0.0.0.0 for the router id really means “pick one of my existing IP addresses, and use that as the router-id”, but i’m not sure if thats what happens with routeros or not; I always explicitly set a router-id.

Otherwise, assuming Halkensector1 and vlan1 are connected together, and I didn’t miss anything, I would say that configuration should work (and it was stated that it did for several hours); but best practice would be to add a loopback, and in ospfv2, use the loopback’s IP as the router id.

Loopbacks never go down, and give the router an always reachable and up IP address regardless of the state of its physical interfaces (reachable as long as you have some sort of connectivity to the router that is).

So, I guess my opinion on that statement is that it is a fix for an issue (it implies the ospf implementation is not bug free), and using the same routeros versions and loopbacks are less likely to trigger or manifest the bug(s). I’m not quite sure what you mean by tweak in this case.

             
            
           
          
            
            
              Whenever you have the local IP address of PPPoE Server Profile same with the local LAN, OSPF stops working after sometime.

So If you have OSPF in your network, the example in the wiki :

http://wiki.mikrotik.com/wiki/PPPoE

is WRONG.

**You need to set the local address of the PPPoE Profile to the loopback address.**

I dont know if this is by design or a bug but I hope they will mention or give a warning about it in the related wiki page.

             
            
           
          
            
            
              Are you saying that adding a bridge, with no physical interfaces linked to it, and adding an ip address on that bridge, solves this issue ?

What would you suggest that the router id might be: the ip on the recently created bridge ?

And also, would the subnet on the bridge added to the ospf/networks list ?

I would appreciate an answer from someone who did it !

Thank you in advance ! 

             
            
           
          
            
            
              
The IP you add on the loopback should be a /32 (single IP, not a range). You then use thes the router-id is OSPF. Because the IP is on a loopback/bridge, it never “goes down” and so the OSPF process is more stable.

             
            

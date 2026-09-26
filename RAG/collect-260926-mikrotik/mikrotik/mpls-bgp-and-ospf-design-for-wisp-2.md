---
id: collect-260926-mikrotik/mikrotik/mpls-bgp-and-ospf-design-for-wisp-2
title: "mpls-bgp-and-ospf-design-for-wisp"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/mpls-bgp-and-ospf-design-for-wisp.md
source_anchor: ""
source_lines: [196, 368]
sha256: d5395b83820af8321ac052403a5dc2180678176379777919a5a046cbb7427a1e
---

# mpls-bgp-and-ospf-design-for-wisp

Also…you need to add a source interface to the BGP peering when using Loopbacks and it looks like you don’t have that set…otherwise, it will inherit the transit subnet address as the source for the peering.

             
            
           
          
            
            
              
Not true…in an iBGP design you want the peering address to use loopbacks that are advertised by OSPF so that if an interface goes down and another path is available, the peering will stay online.

             
            
           
          
            
            
              I set the source interface to the loopback and that fixed the peering issue. Setting syncronize=no fixed the router advertisement. I went through the second presentation and will watch the first one tomorrow. Thanks for the help.

             
            
           
          
            
            
              Hi Kevin.

Can you explain us more about, you know i’ve watched your presentation about

“Using eBGP and OSPF transit fabric for traffic engineering”.

Nice one but would you like to point us about BGP Communities,and how can we build up this kind of configuration, any example or maybe part of that configuration will be appreciated

Thanks

             
            
           
          
            
            
              BGP communities allow you to change the route processing by tagging prefixes that get flooded throughout your BGP network. What you want to use them for is up to you, but I use them for setting local pref and MED values.

For example, I set a particular community for a customer route depending on which gateway I want it to use. Both gateways are configured to look for the communities and set the MED to the core network accordingly. By simply changing the BGP community, I can dictate the path that traffic takes back from the core, without having to update routing filters everywhere.  Within the core, I set different community values based on whether the source of the update is from transit or peering. Then the core BGP routers can set the local pref based on this community string.

There are also well known communities such as no-export, which stops update outside of the local AS, and prevents you from leaking routes outside your network.

             
            
           
          
            
            
              if i have OSPF cost on one way and BGP communities on other way, which way will prefer?

and also thanks for explanation but do you mind to show us how do you configure?

Thanks

             
            
           
          
            
            
              OSPF is an IGP, BGP is an EGP, they have different use cases.

By default, eBGP > OSPF > iBGP in terms of administrative distance, but BGP communities are not attributes that directly affect the best path algorithm. Communities are used simply to tag a prefix that another peer can check to see if it needs to do any special processing of that prefix. For example, most Tier 1 ISPs have a list of communities that you can set to tell them what you want to do with those prefixes. This one is from NTT https://us.ntt.net/support/policy/routing.cfm#communities

Maybe you are dual-homed to a provider and want to tell them which link to prefer. By using the NTT example, you would configure community 2914:490 on your primary and 2914:480 on your backup link. When the NTT routers receive these, they then set the local pref of those prefixes accordingly.

```
/routing filter
add chain=NTT-primary-out action=accept set-bgp-communities=2914:490
add chain=NTT-backup-out action=accept set-bgp-communities=2914:480
```

The NTT side might be

```
/routing filter
add chain=BGP-in bgp-communities=2914:490 action=accept set-bgp-local-pref=120
add chain=BGP-in bgp-communities=2914:480 action=accept set-bgp-local-pref=110
```

             
            
           
          
            
            
              i true it doesn’t go i do not understand two thing.

1. 
Why have you done on each side diferent AS?
2. 
see the picture:

 
            
           
          
            
            
              The goal of this design is to use OSPF only for EQMP load balancing between the sites, but BGP as the overall routing protocol.

EBGP allows routing policy to be modified and advertised at each tower site. iBGP basically considers the entire AS with a more or less single unified routing policy for egress, but internally it will use the IGP for final routing decisions.

             
            
           
          
            
            
              i have a question as we are trying to do the same with mpls and vpls. if you are using vpls tunnels back to your core everywhere, how is bgp doing anything with your traffic ip addresses as everything is supposed to be traveling over the tunnel. I also noticed there were no traffic networks entered for bgp. Are they supposed to be and do the get entered on overy router in your network.

             
            
           
          
            
            
              This depends on your use case.

Some ISPs may use LDP signalled VPLS for private transport circuits. Other ISPs may use BGP signalled VPLS.

In most cases, it’s helpful to have iBGP to advertise public subjects and /32 loopbacks even if the majority of traffic is in VPLS.

In short, having BGP on the interior of the network provides deployment flexiblity.

Add to that MikroTik’s issues with large OSPF routing tables and BGP becomes a solid choice for internal routing while relying on OSPF for reachability.

             
            
           
          
            
            
              I can get everything connected but not carrying traffic properly.  Would you by chance have an example config of say a tower router set up with a couple of AP’s and backhaul, and one of your core router that you are routing it to using private subnets that we could see as an example to compare and learn from.  We can find bits an pieces out there on different parts but not a complete working config of both ends of a real world use case to see.  I think it would be most beneficial to a lot of WISPs as we hear more and more about moving to an MPLS/VPLS/BGP/OSPF network.

             
            
           
          
            
            
              Here’s an example of OSPF/MPLS/VPLS for a WISP with HA DCs and with configs…I’ll see what I can dig up for BGP 

https://www.stubarea51.net/2018/04/23/wisp-design-building-highly-available-vpls-for-public-subnets/

             
            
           
          
            
            
              Thanks.

On the MTU size, I see some people set it to 1530 for MPLS, some 1580, 1600, and 2000.  Is there any downside to setting it to 2000 across the board?

Also I am having issues getting the MPLS working out in the field through the various wireless links even though on my lab it works fine.  I have all the MTU on all equipment between both ends set to 2000.  If I try to ping anything larger than 1500 with do-not-fragment including the router I’m working on it fails.  So why can’t I ping larger than 1500 to all the equipment in between to find if there is an MTU issue somewhere along the path.

             
            
           
          
            
            
              
Hi,

Regarding the MTU setting, some devices or ISPs may have stricter limits on MTU, and it is best to use the same MPLS MTU on all devices if possible. We went with 1550 because it allows us several nested VLAN tags and also allows the customer several nested VLAN tags within their VPLS tunnel, while allowing us also to use RFC4638 with PPPoE.


---
id: collect-261001-general-networking/general-networking/r-networking-comments-e8rjrh-ospf-troubleshooting-a0510ec1
title: "r-networking-comments-e8rjrh-ospf-troubleshooting-a0510ec1"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "latency"]
source: docs/RAG/collect-261001-general-networking/r-networking-comments-e8rjrh-ospf-troubleshooting-a0510ec1.md
source_anchor: ""
source_lines: [1, 167]
sha256: 2c5de8c655023f42fa3d36c2ebe7249c71165721721a720614ef6aa5ac5e86fb
---

# r-networking-comments-e8rjrh-ospf-troubleshooting-a0510ec1

[supprimé]
      [deleted by user] 
        
    Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
You're going to need to provide a lot more topology and config information.
MTU too small on something somewhere? Maybe the LSDB is reaching a full packet and ends up being dropped when flooded.
I don't believe you'll get to FULL with MTU mismatches unless they are configured to disregard MTU. I believe you get stuck in 2WAY at that point.
You would see behavior like this in EIGRP with MTU mismatches -- adjacency forms, then as soon as it starts sending hte routing table, it drops out because packets are getting fragmented along...but OSPF won't let you get this far.
This is the type of behavior I'd see when an invalid route is learned and it's mucking up the adjacency -- but I don't think I'd see that in a basic L2/L3 topology. If VXLAN, MPLS, or GRE are involved, then maybe.
Not different MTU on the ends, something in the middle with a lower MTU. An intermediary switch etc.
sounds to me the route to each other is affected as the peering session is coming up; almost like its being handed a route that is invalid to get to the peer, so after it goes FULL, suddenly they can't reach each other any more. You clear the route, they fall back to (default? static?) and they can re-establish.
https://imgur.com/a/JbE12XU
Here's the very basic diagram.
When I turn up the link to 'Colocation switch', All switches remain adjacent in ospf (and new one discovered), but they all stop receiving keep-alives from the MX80.
If I force a 'clear ospf neighbor' on the MX80, the 3 switches rediscover the MX80 until the timer once again expires.
During everything the MX is surprisingly unaware this is happening and all neighbors stay in a 'full' state.
Can you turn on debugging for all or part of OSPF? For example, for the entire process or for LSAs?
Are all router IDs unique? In particular, is the third switch that joins using the same RID as one of the other devices?
The way you explain it, it sounds like your MX is directly connected to a single switch, which is connected to another switch, which is connected to the other switch. But your post also makes it sound like the third switch that causes issues is connecting directly to the MX80 ("my MX receives the hello..."). Can you draw a topology using draw.io?
Sounds like you have a conflicting Router ID problem to me. Have you mapped the RID to every router through "set routing-options router-id xxx" then set priority to the MX to be the DR, that way the other QFX won't fight to who becomes the ADJ router of the area 0. I would personally split the cluster into areas to minimize OSPF traffic and cut down on this sort of problem.
Router ID's appear correct. I have not forced the priority of the MX but I probably should.
If this was a larger footprint I'd have setup multiple areas, but I only have a few 'Core' routers.
So just to clarify, you have each 5100 vc hung off the other and all devices are serving as routers?
Just to clarify:
Is this all one PTMP network?
Which relationships stay established when you plug SW3 in, and which stop working?
What IGMP settings do you have on the QFX5100s?
It sounds like you really need to find out what's happening to these advertisements once the adjacency is made. Seeing that you need to minimize downtime as much as possible, how about setting up some port mirroring on the relevant interfaces on the switches so you can capture packets without impacting service?
Are your hello and dead timers consistent throughout all your devices?
Yes. There's no latency between either location either so they mostly always re-up at the same exact interval.
Yeah, answers already good.
This could be OSPF preferring it's own route for neighbour; once established the initial neighbour route goes away + then keepalives drop.
Or MTU issues, something like the hellos get through but not the keepalives.
Would need to see the OSPF and interface configs honestly.
Sounds like a duplicate router ID somewhere. Maybe on the new QFX you added.
You're not adding the new QFX to the VC I take it?
I'm living in an entirely Cisco world and so without looking over all your config this is really just thoughts. I feel like this might be some sort of DR/BDR thing. All this OSPF is forming over a broadcast network (Vlan 1). Once all these form can you see who the DR and BDRs are? And that the network type is broadcast? Can you debug ospf on any of these interfaces?
Normally I'd do everything with separate layer 3 p2p links for OSPF, or if on a trunk I'd designate a seperate vlan just to form the l3 connectivity via SVIs. However you might not be able to do this in your circumstance.
Make sure you don't have a duplicate router-id that is being learned when the adjacency comes up.
I'm wondering if you've got an IGMP/Multicast problem that gets triggered with a DR election by the last qfx coming online.
you have "LAG" is that MC-LAG?
The accept-data attribute is needed for VRRP over IRB interfaces in MC-LAG to enable OSPF or other Layer 3 protocols and applications to work properly over multichassis aggregated Ethernet (mc-aeX) interfaces.
are you using BFD ?
NOTE
On a QFX5100 switch, the minimum transmit interval must be 1000 milliseconds or greater. Sub-second timers are not supported in Junos OS 13.2X51-D10 and later. If you configure the minimum transmit interval timer lower than 1000 milliseconds, the state of the MC-LAG can be affected.
also double check you protect- RE rules
Are the switch really daisy chained like that?
you can also do TCP dump to files to see at what point the keep alive packets are dying.
could also be a MTU mismatch or smaller mtu in the middle fraggin the packets. a trace options on the devices for OSPF is needed along with some packet captures
Double check all the subnet masks match. Also check that there are no MAC overlaps.
It could be a problem with the Dr/her election I. A nbma network.
Honestly, it is a design problem at first. Why do you want to establish routers adjacencies in the same vlan? They are routers, they should establish adjacencies bin dedicated , point to point vlan.
It would be a lot more efficient for redundancy. For instance you would not have to troubleshoot this. Your redundancy will be based on STP plus ospf for instance (because of root STP and Dr election)
Keep things simple, do some segmentation
you can always use draw.io for a quick fast diagram..
I’m not an OSPF expert by any means but who is the DR before you add QFX 3? Who is DR after? Looking at the diagram you posted QFX 3 has the highest IP so I would expect it to be elected as DR if you haven’t manually assigned a priority DR. If there is a routing issue between that device and QFX 1/2 then they would be unable to build adjacencies to QFX 3.

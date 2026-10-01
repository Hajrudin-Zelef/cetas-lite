---
id: collect-261001-general-networking/general-networking/r-networking-comments-10wcy0o-help-with-sudden-ospf-failure-seems-to-be-ccafecfb
title: "r-networking-comments-10wcy0o-help-with-sudden-ospf-failure-seems-to-be-ccafecfb"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["latency"]
source: docs/RAG/collect-261001-general-networking/r-networking-comments-10wcy0o-help-with-sudden-ospf-failure-seems-to-be-ccafecfb.md
source_anchor: ""
source_lines: [1, 33]
sha256: fde757d89177146409dd7ca40ba05b31b7273d0c4ed6561e0f9c115091ab3fd0
---

# r-networking-comments-10wcy0o-help-with-sudden-ospf-failure-seems-to-be-ccafecfb

Help with sudden OSPF failure - seems to be retransmission issue - debugging inside 
        
        
        
    
    
    So recently had OSPF go down between an edge Palo Alto and a Cisco core switch.
No changes were made to either device as far as I can tell. Logs show no commits on the PA and no config changes to the Cisco switch (HA Pair)
after some investigation I get:
      ECHANGE to DOWN, Neighbor Down : Too many retransmission
found a few articles on both Cisco and PA side, both pointing to either MTU issue or Packet loss
      Packet loss is unlikely: 1000 pings from OSPF interfaces 0% loss, <2ms average latency. They are directly connected
Some further debugging:
      I can see the Adjacency process begin and the DBD packets transmit
I can see the VlXX OUT to the PA
I can see the VlXX IN from the PA
OUT packet len:1412
MTU on the router interface is 1500, verified by ping sweep
I can see the packets retransmit until they hit 25 where it is killed due to excessive retransmission
IN packets still come but are ignored for a timer until process repeats
So that size 1412 shouldn't get fragmented correct?
One other thing I noticed is that on the HA Pair, all the Vlan interfaces are setup HSRP, except the one facing the PAs. For whatever reason, those vlan interfaces are setup separate i.e. VlanXX x.x.x.1 and x.x.x.2 with no standby...
oddly after the Neighbor exchange goes down, I see a new election between the two devices but using other HSRP Vlan interfaces, but over the OSPF VlanXX already in process. This setup seems really odd to me? Could it cause issues?
It's been a while since I went really deep into OSPF troubleshooting so I'm at a little bit of a loss of where to go next. I can't find any other errors on the PA or Cisco devices other than the constant retransmission/kill loop.
Section des commentaires
I had a similar issue before that presented itself when the LSDB got big enough to fill an entire full-sized packet. The L2 MTU was too small between the Cisco and Palo, yet the L3 MTU was configured the same so the neighbors were never stuck in Exstart prior to the LSDB being too big. (OSPF has an MTU check built into the protocol, where each neighbor lists its interface MTU in the DBD packet and neighbors will stay stuck in Exstart if they detect that the MTU doesn't match). My issue was fixed by fixing the L2 MTU on the switch connecting the Cisco and Palos. The issue went completely undetected until the LSDB got big enough.
I know you verified MTU but that really seems like the issue. Can you try pinging between the Cisco and two Palo firewalls with size 1500 and df-bit? Maybe check in the Palos that the system MTU hasn't been set somewhere which is causing the problem? Even though the Cisco MTU is 1500, are you sure the Palo MTU isn't even higher and causing the issue?
Edit: Thinking about this again, it's interesting to note that IS-IS actually pads the Hellos to the MTU to verify that max size packets actually work. You could argue that OSPF is less robust because it simply "believes" its neighbor that the MTU listed is actually viable from a L2 MTU standpoint. OSPF never tests that full size packets will actually work, and when L2 MTU is a problem you run into situations like this.
It was an MTU issue. They were both at 1500, I checked that with DF bit pings. But I enabled jumbo frames on the Palo and set it to 1600 and did the same on the Cisco switch on both the trunk interface and the Vlan interface and it came up immediately. I can also ping 1600 sized packets with DF from the Vlan interface.
I could see it immediately on the first HSRP switch, while the other was still retransmitting until I added it to the trunk interface between them so the full path had the same MTU. Very interesting.
Also if using a keychain on the cisco side for ospf authentication make sure it did not expire or shift to a key not on the palo.
Commentaire supprimé par un membre de l’équipe de modération
There wasn't, it was MTU issue

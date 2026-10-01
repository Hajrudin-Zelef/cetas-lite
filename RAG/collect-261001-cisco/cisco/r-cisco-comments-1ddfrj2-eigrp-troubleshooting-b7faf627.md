---
id: collect-261001-cisco/cisco/r-cisco-comments-1ddfrj2-eigrp-troubleshooting-b7faf627
title: "r-cisco-comments-1ddfrj2-eigrp-troubleshooting-b7faf627"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/r-cisco-comments-1ddfrj2-eigrp-troubleshooting-b7faf627.md
source_anchor: ""
source_lines: [1, 25]
sha256: 3d9d3407532d21204bf6989e45f770950ecde813b4bb196b2c75e3bf3ccab267
---

# r-cisco-comments-1ddfrj2-eigrp-troubleshooting-b7faf627

EIGRP troubleshooting 
        
    I'm having some trouble with an EIGRP neigborship forming with a newly deployed 2960XR to our core. The 2960XR is on the other end of a service provider L2 link there is an additional 2960X (L2 only) in between the service provider's switch and our core 9500.
2960XR -->{Service Provider L2} -->2960X -->Core_C9500 On the 2960X using RSPAN to mirror the port where the L2 SP link lands I can see the hello multicast packets from the 2960XR. The port on the 2960X is configured as a trunk port as is the uplink to the 9500.
The 9500 does not appear to be seeing any EIGRP packets from the 2960XR. My initial thought is that this is a multicast problem and that the multicast from the 2960X is not making it to the 9500. Does this sound like it could reasonably be the issue? Do you have any recommendations for troubleshooting?
I've generally had no problem with configuring EIGRP I have multiple sites over service provider links, the only wild card that makes this different is the 2960X in between the SP switch and the core.
Update:
Turns out that this is a service provider issue, I am not able to properly pass tagged traffic over this connection. I converted the port on the far end to a routed port instead of a trunk port with a transit vlan and the port on the near end as an access port on the transit vlan. The EIGRP neighborship formed immediately and I had the L3 connectivity that I expected. This is not a long term solution but at least I know what the problem is and have opened a ticket with the service provider.
Section des commentaires
Turns out that this is a service provider issue, I am not able to properly pass tagged traffic over this connection. I converted the port on the far end to a routed port instead of a trunk port with a transit vlan and the port on the near end as an access port on the transit vlan. The EIGRP neighborship formed immediately and I had the L3 connectivity that I expected. This is not a long term solution but at least I know what the problem is and have opened a ticket with the service provider.
I've seen from your other posts you are unable to ping and cannot see MAC / ARP entries at each end of the link. Have you got your provider involved in this?
You may need to adjust your MTU on the interfaces to allow for the overheads if they are running a L2VPN etc
Thanks, MTU is on my list to check. I did confirm that the provider MTU is 9216, MTU on the 9500 is 9216 the system MTU settings on the 2960x and 2960XR is:
Have you confirmed L2 connectivity from 2960xr to 9500? Can you ping between them, do you have ARP entries in both directions, etc?
I have confirmed L2 connectivity between the 2960x and the 2960XR I can see CDP info, I'm not quite sure how to check L2 connectivity between the 2960XR and the 9500.
I have not been able to pass any L3 traffic I see no arp entries and I can't ping the EIGRP transport Vlan interface. At the moment I do not have physical access to the 2960XR to debug on that end.
You need to resolve that part first before worrying about EIGRP/multicast.
Are you using a routed port on the 9500 or an SVI? Do you have the necessary VLANs built all the way across? You're probably missing a VLAN somewhere, or you have a trunk vs. access mismatch, etc. Once you get basic L2+L3 working from 9500 to 2960xr, EIGRP will probably come right up
I would try to create an SVI on the 2960x (shut down the interface on 9500) and see if you can ping from the 2960x across the SP network.
Is the transit vlan present in the database?
Did you create your EIGRP transit vlan on 2960?
The transit VLAN/SVI has been created on the 2960XR the VLAN exists on the 2960X and the VLAN/SVI interface exist on the 9500. I see the EIGRP hello packets from the 2960XR on the port on the the 2960X with a source address of the Transit SVI and a multicast Destination 224.0.0.10.
The port on the 2960XR and the 2960X and the port between the 2960X and the 9500 is configured with only:
Is ip routing or correct sdm template enabled on 2960x ?
IP routing is not enabled on the 2960X functionally it is purely configured as a layer 2 switch. IP routing is configured on the 2960XR and the 9500.

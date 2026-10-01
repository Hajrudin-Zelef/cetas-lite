---
id: collect-261001-cisco/cisco/r-cisco-comments-e713t5-eigrp-neighbor-flapping-every-80-seconds-0dbdc0ed
title: "r-cisco-comments-e713t5-eigrp-neighbor-flapping-every-80-seconds-0dbdc0ed"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["asic"]
source: docs/RAG/collect-261001-cisco/r-cisco-comments-e713t5-eigrp-neighbor-flapping-every-80-seconds-0dbdc0ed.md
source_anchor: ""
source_lines: [1, 39]
sha256: 16ab6a9d5b5c0e1861a4c0e4722a4e9dd15524231ba5cbf77ba51cf9694a4777
---

# r-cisco-comments-e713t5-eigrp-neighbor-flapping-every-80-seconds-0dbdc0ed

EIGRP neighbor flapping every 80 seconds 
        
    Just inherited a data center build out where the main pieces of Cisco equipment are a pair of Nexus 93180YC-EXs and 4451 ISRs. The Nexus pair is running vPC and acting performing core switching/routing, while the 4451s terminate DMVPN tunnels. L3 Routing between the four is via EIGRP on a /29 SVI. All vPCs Port-Channel interfaces are a single VLAN:
interface Port-Channel26
 switchport
 switchport mode access
 switchport access vlan 101
 spanning-tree port type edge
 vpc 26
The EIGRP neighbors are all stable and working with one exception: the neighbors between the secondary Nexus and both 4451s bounces every 80 seconds. Logs look like this:
2019 Dec  1 23:47:09 Nexus93180-02 %EIGRP-5-NBRCHANGE_DUAL:  eigrp-101 [30032] (MyVRF-base) IP-EIGRP(0) 254: Neighbor 192.168.1.66 (Vlan101) is down: retry limit exceeded
2019 Dec  1 23:47:09 Nexus93180-02 %EIGRP-5-NBRCHANGE_DUAL:  eigrp-101 [30032] (MyVRF-base) IP-EIGRP(0) 254: Neighbor 192.168.1.66 (Vlan101) is down: Peer goodbye received
2019 Dec  1 23:47:09 Nexus93180-02 %EIGRP-5-NBRCHANGE_DUAL:  eigrp-101 [30032] (MyVRF-base) IP-EIGRP(0) 254: Neighbor 192.168.1.67 (Vlan101) is down: retry limit exceeded
2019 Dec  1 23:47:09 Nexus93180-02 %EIGRP-5-NBRCHANGE_DUAL:  eigrp-101 [30032] (MyVRF-base) IP-EIGRP(0) 254: Neighbor 192.168.1.66 (Vlan101) is up: new adjacency
2019 Dec  1 23:47:09 Nexus93180-02 %EIGRP-5-NBRCHANGE_DUAL:  eigrp-101 [30032] (MyVRF-base) IP-EIGRP(0) 254: Neighbor 192.168.1.67 (Vlan101) is up: new adjacency
2019 Dec  1 23:48:28 Nexus93180-02 %EIGRP-5-NBRCHANGE_DUAL:  eigrp-101 [30032] (MyVRF-base) IP-EIGRP(0) 254: Neighbor 192.168.1.66 (Vlan101) is down: retry limit exceeded
2019 Dec  1 23:48:28 Nexus93180-02 %EIGRP-5-NBRCHANGE_DUAL:  eigrp-101 [30032] (MyVRF-base) IP-EIGRP(0) 254: Neighbor 192.168.1.66 (Vlan101) is down: Peer goodbye received
2019 Dec  1 23:48:28 Nexus93180-02 %EIGRP-5-NBRCHANGE_DUAL:  eigrp-101 [30032] (MyVRF-base) IP-EIGRP(0) 254: Neighbor 192.168.1.67 (Vlan101) is down: retry limit exceeded
2019 Dec  1 23:48:28 Nexus93180-02 %EIGRP-5-NBRCHANGE_DUAL:  eigrp-101 [30032] (MyVRF-base) IP-EIGRP(0) 254: Neighbor 192.168.1.67 (Vlan101) is down: Peer goodbye received
2019 Dec  1 23:48:28 Nexus93180-02 %EIGRP-5-NBRCHANGE_DUAL:  eigrp-101 [30032] (MyVRF-base) IP-EIGRP(0) 254: Neighbor 192.168.1.66 (Vlan101) is up: new adjacency
2019 Dec  1 23:48:28 Nexus93180-02 %EIGRP-5-NBRCHANGE_DUAL:  eigrp-101 [30032] (MyVRF-base) IP-EIGRP(0) 254: Neighbor 192.168.1.67 (Vlan101) is up: new adjacency
I've checked general configuration, hello timers, subnet masks, and MTUs - everything's good. Was doing some Googling and found this may be expected behavior when trying to run EIGRP on an SVI where L2 is vPC. But why would it occur only on the secondary switch?
Section des commentaires
I have run into similar issues. Your topology would be fine with catalyst switches, but vpc on nexus have a few rules. I would recommend having routed links to each nexus
https://www.cisco.com/c/en/us/support/docs/ip/ip-routing/118997-technote-nexus-00.html
Got it. The 93180YC-EX is basically a re-branded 5k / 6k so looks like I'd have to upgrade to 7.3(0)N1(1) or later
We're currently running 7.0(3)I7(6
No, it’s not (different ASIC and different code base). I7(6) is more than enough to run peer-gateway together with layer3 peer-router.
I need more of your topology to help you..
What is the connection between the N9k's and the 4451's? From what I see I'm assuming it a Port-channel?
If that is the case, create some /31's to make the connection and not use a port-channel and let you IGP do the work for you.
The only place I am using vlans/vpc to connect to other devices are to switches and my FTD's. Everything else is L3.
> What is the connection between the N9k's and the 4451's? From what I see I'm assuming it a Port-channel?
Correct. Layer 2 Portchannel (switchport mode, single vlan, STP edge port)
> If that is the case, create some /31's to make the connection and not use a port-channel and let you IGP do the work for you.
Yes, I tried this a few hours ago and it worked fine. So that proves it's not a Layer 1/2 problem. It makes the routing more complicated since I'm 100% reliant to EIGRP redistribution in to our DMVPN, but that also lessens the chance of a blackhole so I'm probably OK with that.
I just had to clean up one of these yesterday. Go back and try it again and make sure spanning tree isn't blocking the vlan.
That is what my issue was.
Do you have layer 3 peer router enabled on the vpc?

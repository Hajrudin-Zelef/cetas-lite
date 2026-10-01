---
id: collect-261001-cisco/cisco/t5-switching-l3-routing-problem-with-nexus-vpc-pair-in-typical-core-position-td-aa260c21-2
title: "t5-switching-l3-routing-problem-with-nexus-vpc-pair-in-typical-core-position-td--aa260c21"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/t5-switching-l3-routing-problem-with-nexus-vpc-pair-in-typical-core-position-td--aa260c21.md
source_anchor: ""
source_lines: [47, 202]
sha256: 4ddba5f21e775af0b55a60c94b4bc2a36e543d105f0351b221f27d84e36d3a26
---

# t5-switching-l3-routing-problem-with-nexus-vpc-pair-in-typical-core-position-td--aa260c21

			Catalyst 3000
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-21-2022 10:08 AM
I think I found where you are misunderstanding, 
when the traffic receive for V100 it can not L3 routing using SVI of V200 it must route use V100 in other nexus Peer.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-21-2022 12:04 PM
hi mhm, i think your answer is inline with what i described in the latest post, showing routing behavior. so thank you for reaching out!
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-18-2022 03:02 PM - edited 03-18-2022 04:24 PM
peer-gateway & HSRP 
HSRP
HSRP have it VIP and this VIP is answer from both vPC this is how vPC work,
then what is the function of active HSRP, it the only one that reply to ARP request.
Peer-gateway
when config the SVI in each Peer this peer will the only peer answer to any traffic destination to this SVI IP, BUT with Peer-gateway this will enhancement with make both peer answer to SVI IP of one Peer. 
do show mac address vlan 
you will see "G" beside the SVI for each Peer.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-18-2022 05:14 PM
hi, thx for quick reply. i have peer-gateway configured, as seen in the vpn config snippet. still doesn't work.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-18-2022 06:50 PM
May be below example clarifies your setup, why vPC HSRP working active/Active :
https://www.ciscozine.com/nexus-vpc-hsrp-vrrp-active-active/
=====️ Preenayamo Vasudevam ️=====
***** Rate All Helpful Responses *****
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-19-2022 11:31 AM
"weird on/off connectivity between the subnets"
can you more elaborate ?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-19-2022 04:13 PM - edited 03-19-2022 04:14 PM
thx for all the replies, in first post, I tried to simplify the design a bit so its easy to understand. after reading the ciscozine behavior, i realize, it might be worth elaborating a bit further, as i didn't think the backup peer would send traffic out one of its active interfaces, if the the other interface was down. Here is a more accurate diagram:
|------N9K-1------|
DMZ----Firewall (Vlan200) || (VPC) CAT3k-1 (vlan100)
|------N9K-2------|
When vlan100 is shut on the N9K-2 backup VPC peer, the traffic from v100 destined to stuff in the DMZ off the firewall on vlan200 doesn't make it, or doesn't get a reply if it does make it. I need to troubleshoot this more on Monday, but essentially it looks like when the vlan 100 interfaces is down on N9K-2, the traffic should go N9K-1, which I would expect anyway when both interfaces are up, but there is something odd about how nexus peer process traffic, so even when both vlan100 are up, some hosts go through the backup peer v100 interface.
Anyway, shut v100 down on N9K-2 and traffic fails to make it to or receive from the DMZ hosts hanging off the firewall. Note, firewall is also port-channeled to the Nexus peers and running on v200, and both interfaces are up on the peers for v200. in other words, N9K-1 and N9K-2 have both v200 up and forwarding traffic. Supposedly, N9K-1 is HSRP primary. I shut only one of the V100 interfaces and something breaks on the workstation switch side (Cat3K) and they cannot talk to the DMZ hosts.
In thinking through this, i think perhaps the firewall may be seeing traffic from 2 x v100 interface native IP's on the nexus side, rather than the HSRP address. In this case, the firewall may be doing some sort of IPS since its seeing traffic from essentially two v100 nexus IP's, rather than their HSRP ip address.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-19-2022 05:07 PM
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-19-2022 05:18 PM
thx for diagram. its starting to look clearer. In step 8 though, why does 7k2 not just send directly to S over Po2? is it because its a PO and hashing tells it to use 7k1? If so, why is there not logic in the NXOS code to handle this situation? There is no reason why 7k2 could say, I have "peer-gateway" configured so I will just send to S, or for 7k1 to say, why not just send the traffic i just got from 7k2 over peer-link to the S via my 7k1 po2 uplink? It looks like this might be why im seeing traffic death? How do i handle this situation the best way in configuration then?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-19-2022 05:32 PM
in example not show why 7K2 forward traffic directly BUT, 
one of case is 7K2 shutdown SVI so the 7K2 don't have L3 capability to forward traffic "notice since traffic pass through different VLAN so it need L3 to forward traffic".
7K2 need L3 to forward traffic BUT 7K1 have this capability So it send traffic through Peer-link and then the it drop.
so in FW which is router mode how the route is config for VLAN 100?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-19-2022 09:34 PM
hi mhm, on my firewall, the route to v100 is set to the v200 HSRP address of the vpc peers. In the failure scenario, im only shutting the vlan100 interface on _ONE_ of the vpc peers. I would expect L2 traffic coming to the down vpc peer v100 to be forwarded over the vpc peer link for L3 processing. I need to analyze if this is happening or not, and then see if firewall is actually getting this traffic or the return traffic is failing. ill be able to do that next week. thgx again for everyone's replies.
I still think the logic in step 10 of the diagram is confusing, there is no reason for the VPC L3 peer to not forward the traffic, in my opinion. if L2 is working on 7k2, why not just shoot it over peer link and then out the 7k1 svi? Also, my traffic flow is a bit different that the diagram in that the PO1 and PO2 segments are different L3 vlans. and the nexus are doing routing for the traffic as well and L3 port-channeling.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-20-2022 07:04 PM
Match hash in both v100 and v200,
Make sure asa use mac address of vip of hsrp not real address.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-19-2022 12:31 PM
Hello
If i understand you OP, then what you see is normal, as stated by others , even if you run a FHRP such as HSRP in a vpc environment, data plane traffic is treated by the vpc a active/active so either switch can process , Only control plane traffic such as an arp query would you then obtain a reply from the HSRP primary 
Please rate and mark as an accepted solution if you have found any of the information provided useful.
This then could assist others on these forums to find a valuable answer and broadens the community’s global network.
Kind Regards
Paul
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-21-2022 05:00 AM
Hi
As other has stated, this is the normal behavior. Why, this is due to the loop prevention in the vpc environment. The loop prevention is accomplished by marking frames that is received from a "vpc port-channel" and then sent out on the vpc-link. When the vpc-peer receives the frame it will not send this frame out on a "vpc port-channel".

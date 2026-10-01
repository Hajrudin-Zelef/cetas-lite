---
id: collect-261001-general-networking/general-networking/r-networking-comments-spvkzq-eigrp-issue-absolutely-no-clue-what-the-problem-is-3764ee29
title: "r-networking-comments-spvkzq-eigrp-issue-absolutely-no-clue-what-the-problem-is-3764ee29"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["optics"]
source: docs/RAG/collect-261001-general-networking/r-networking-comments-spvkzq-eigrp-issue-absolutely-no-clue-what-the-problem-is-3764ee29.md
source_anchor: ""
source_lines: [1, 55]
sha256: a422b1c71c26be4e76e3c7f7dd568bc158272922bf2d5a1fb29477cd113b8093
---

# r-networking-comments-spvkzq-eigrp-issue-absolutely-no-clue-what-the-problem-is-3764ee29

eigrp issue... absolutely no clue what the problem is...
Diagram: https://imgur.com/a/zZFeOAx
So I've got a pair of N7K's full mesh connected to a pair of Cat9K's with each link peered with eigrp. Everything works except ONE link from the N7K-A to Cat9K-B. The link will initially peer, steadily climb to 5000 RTO and then fail and drop after about 20 seconds, rinse and repeat.
N7K-A:
interface Ethernet1/20
 mpls ip
 mtu 9150
 bfd interval 300 min_rx 300 multiplier 3
 no ip redirects
 ip address 192.168.0.90/31
 ip router eigrp 5
 ip pim sparse-mode
C9K-B:
interface HundredGigE1/0/51
 no switchport
 ip address 192.168.0.91 255.255.255.254
 no ip redirects
 ip pim sparse-mode
 mpls ip
 mpls label protocol ldp
 bfd interval 300 min_rx 300 multiplier 3
end
Am I missing something super obvious here? All of the other links peer without issue. I also have this exact same setup elsewhere and I have no problems with peering. The only difference is that this particular N7K pair is configured as a VPC pair while the other one isn't, but I don't think that's the issue here.
The Cat9K's are not stacked, but are running GLBP with redundant L2 trunks to downstream switches.
Logs don't really indicate anything outside of the BFD session teardown after the routes drop.
Edit: Pack it up folks. Layer-1, bad optic issue. Nothin' to see here.
Section des commentaires
What's the MTU on the C9K? The interface config for the N7K says it's 9150 but on the C9K there's no MTU line.
System mtu 9150
That's the L2 MTU, you also need to verify L3 MTU which is what the MTU config on the port is changing.
I don't have a nexus in front of me but look at the output of a routed port on an Arista
Max Frame size != Max Packet size. Packets can fragment, frames can not.
Start with layer 1. What type of link? Errors?
Uhhh... haha, I think you win the prize. Not sure how I didn't notice it before, seeing as how I think I've done show int a million times. Guess I was too busy freaking out when the messed up routing kept breaking stuff. Threw both interface as passive and ran ping. Lookin' like abstract ascii art.
Time to start swapping everything out.
If I had a nickel for every time I made the same mistake, I'd have at least a quarter.
I had a weird issue like this on a SAN (1 fiber path was reaaaally slow). After 2 weeks of troubleshooting with VMware and Dell, I went through to push all the cables in just for the hell of it and one "clicked" in place -- that was the path that was flaking out. It was plugged in just enough to show link but whenever we tried to push data over the link it slowed to a crawl.
Is this a “layer3 peer-router” scenario?
https://www.cisco.com/c/en/us/support/docs/ip/ip-routing/118997-technote-nexus-00.html
The cat9ks are not participating in any pc/vpc shenanigans. Vpc is purely for some firewall stuff connected to the n7ks.
Are you getting any packet loss between the links? If you do a constant ping between those links do you see any drops?
Wow, EIGRP. Havent heard of anyone still using that for awhile.
Haven’t worked on EIGRP for a decade, but, from those days you need two things to form an adjacency. Same AS and same K values. I have to assume your K values and AS are the same as all other links are working.
This leaves physical as the most likely option. Some type of unidirectional comms. Have you looked at the interface statistics to see if you’re taking errors or packet counters are increasing?
I don’t remember which debug you can use, something like ‘debug ip eigrp packets’ that would allow you to see packets as they arrive, if they arrive, and what process is being done upon them. That might glean some info.
You’ve got an interesting one for sure. I’d focus on L1 at this point. Even going so far as to change all optics, cables, and interfaces to rule out through escalating triage.
Are you running NX-OS 6.x by any chance? Saw exactly the same thing and I think it's to do with L3 VPC. Seem to remember we would've needed to upgrade to v7 train so we redesigned our solution to get around it.
have you even tried troubleshooting?
show eigrp address-family ipv4 <AS> events
show eigrp address-family interfaces
show eigrp address-family topology
show eigrp address-family neighbors
show eigrp address-family ipv4 <AS> interfaces detail HundredGigE1/0/51/ and Ethernet1/20
Also, show cdp neigh
Normal, show int HundredGigE1/0/51/ and Ethernet1/20 might be helpful. I would also remove the bfd while troubleshooting.

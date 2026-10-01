---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-151872-unifi-ap-how-to-issolate-the-guest-network-d71e4cf0
title: "questions-151872-unifi-ap-how-to-issolate-the-guest-network-d71e4cf0"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-151872-unifi-ap-how-to-issolate-the-guest-network-d71e4cf0.md
source_anchor: ""
source_lines: [1, 15]
sha256: 2840429408b07bd0bc19235fa1d18e9c4b46db6a54be0489a7dcf80f2cf56c82
---

# questions-151872-unifi-ap-how-to-issolate-the-guest-network-d71e4cf0

Information Security is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
0
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
How do I isolate my "Guest" network on my Unifi Access Point(AP) so they can't see all the computers on the network and thereby prohibiting access to our computers/data on other networks on same AP.
On my old Router, the "Guest" network was going through the WAN so they had an IP address 192.168.1.xxx and there they couldn't see each other. But on this AP, I can't connect them to the WAN, but only the LAN. Hence they have an IP address in segment 10.5.25.xxx
Is that something I have to setup through the VLAN/LAN settings?
First of all let you know that applying restrictions to your guests will let them still 'see' other hosts because these restrictions don't block ICMP protocol, but still they can't even ping other hosts neither access.
If you want to isolate completely your guest network I would recommend you VLANs. You can easily set them up if you have a Unifi Security Gateway.
Set up a new network with a complete different pool of IPs and subnet (10.0.1.0/24 for example) and mark it as VLAN '123' for example.
Now in Wireless Network settings you can edit your guest network and in the advanced settings select 'Use VLAN with VLAN ID' and put in again '123'.
From now on your guests should been put in VLAN '123' with an IP in range 10.0.1.1-10.0.1.254 and will be completely isolated from the main network.
Hope I helped you or gave you at least some ideas!

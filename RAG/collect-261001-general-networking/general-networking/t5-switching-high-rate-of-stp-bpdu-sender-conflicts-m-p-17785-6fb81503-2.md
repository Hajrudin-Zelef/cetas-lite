---
id: collect-261001-general-networking/general-networking/t5-switching-high-rate-of-stp-bpdu-sender-conflicts-m-p-17785-6fb81503-2
title: "t5-switching-high-rate-of-stp-bpdu-sender-conflicts-m-p-17785-6fb81503"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-switching-high-rate-of-stp-bpdu-sender-conflicts-m-p-17785-6fb81503.md
source_anchor: ""
source_lines: [44, 187]
sha256: f5db86231398980f7a72d12f0e0e01b4ac7e70e59dd1cb4860af5742a00dbef6
---

# t5-switching-high-rate-of-stp-bpdu-sender-conflicts-m-p-17785-6fb81503

			Meraki
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-24-2018 02:56 AM
Is the grape office switch a Meraki switch?
Next thought; if you can be confident the network is loop free, it may be worthwhile simply disabling spanning tree.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-24-2018 01:34 AM
The 10.x firmware had a lot of spanning tree improvements. Try going to 10.x if you don't make progress.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-24-2018 02:08 AM
@Philip D'Ath wrote:
Have you set the spanning tree priority on your "core" switch to force it to be the root? If not do this first.
The 10.x firmware had a lot of spanning tree improvements. Try going to 10.x if you don't make progress.
Yes, the core switch's bridge priority is set to 0 to force it to be the root. It is the "Root Bridge" switch I refer to throughout the question, which connects to the Rocket AP.
All of the switches are currently already on the 10.6 firmware.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-24-2018 02:41 AM
10.6! I would go to 10.22 promptly.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-24-2018 02:56 AM
Is the grape office switch a Meraki switch?
Next thought; if you can be confident the network is loop free, it may be worthwhile simply disabling spanning tree.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-24-2018 10:13 PM
@Philip D'Ath wrote:
Is the grape office switch a Meraki switch?
Next thought; if you can be confident the network is loop free, it may be worthwhile simply disabling spanning tree.
Yes, it's an 8-port MS220.
I will disable it for the time being and see how it goes. I may begin using redundant links sometime in the future as the network grows, so I would eventually have to enable STP again.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-25-2018 03:33 AM
You can do redundant links with port aggregation. Then you don´t need STP.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-10-2019 03:23 PM
Hi redsector.
STP is just not used for redundancy, it is for find and stop loop in network as well .
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-13-2020 06:40 AM
I realize that this is an old thread but thought I would update with information that would be relevant in today's environment. I had similar issue and the solution was to enable VLAN 1 on the core switch and walk it down to the other switches. After doing so the switch status will update to point to the Core Switch as being the Root Switch in the topology. Here is an article that was provided to me for further reading.
Identifying Root Switch:
Configuring STP on Meraki Switch:
Advance Setting of STP on Meraki Switches:
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-25-2018 09:36 AM
Hey Piet,
These warnings are actually ones you don't want to ignore, particularly if you are planning to introduce redundant links down the road.
Your network must be on the MS 10.x firmware release, as this is the version where BPDU conflict logging was introduced, as part of overall enhancements to anomaly detection.
If I understand correctly your topology is something like this:
You've got a switch, ROOT BRIDGE, to which one port has a connection with a Ubiquity AP.
Then you have three Meraki switches:
Workshop, Admin and Grape Office
Each of these three switches has another Ubiquity wireless bridge connected, and all three of them wirelessly bridge back to the AP connected to your root switch. The outcome is that from the perspective of your root switch, all three of these Meraki switches are downstream from the single port with the wireless bridge.
This is where your problem is stemming from, and what the logging has identified. The wireless bridge solution is for all intents and purposes acting as a dumb L2 switch. This dumb/unmanaged switch effectively interconnects four of your switches in a star topology:
(A) ROOT BRIDGE
(B) Workshop
(C) Admin
(D) Grape Office
In addition, you have a 5th "pseudo-switch", the unmanaged L2 switch formed by the wireless bridging, I'll refer to as:
(E) Unmanaged Switch
For the remainder of this post I will refer to these switches by A, B, C, D and E per the above list. The switches A through D must be running RSTP, not legacy STP, as the warnings you see logged are only produced from segments operating in RSTP mode.
Now here's the problem:
When any of the four switches A through D transmits a BPDU, the BPDU will be received by switch E (unmanaged switch / wireless bridge). BPDUs are sent to a special destination of 01:80:C2:00:00:00. This is the well-known address used by the IEEE STP/RSTP protocols.
If a switch supports STP, when it recevies a BPDU with this special destination address, it does not forward the BPDU out other ports. It just uses the data from the BPDU for its own STP calculations, and then may generate its own BPDUs to send out other ports.
But in the case of switch E, (R)STP is not supported, it's just an unmanaged L2 switch. So when E recevies a BPDU from any of the switches A through D, it will just flood the BPDU out all its other "ports" (wireless links in this case).
Example:
(1) Switch A transmits a BPDU.
(2) Switches B, C, and D all receive this same BPDU.
This is the root of the problem. It would actually be fine using legacy STP, in which a port won't become forwarding until the expiration of a long timer, which can be around 30 seconds. But with RSTP, ports become forwarding rapidly (hence the name RSTP!).
The mechanism RSTP uses to rapidly transition a port into the forwarding state and bypass the 30 second delay of the old standard is through the use of a proposal/agreement negotiation between ports. Instead of wating for timers to expire, a port will send out an initial BPDU with a "proposal" flag set. The other port that receives this "proposal" BPDU will send out a responding BPDU with an "agreement" flag set.
This affirmative proposal/agreement process is the primary mechanism used by RSTP to converge faster than the legacy standard which exclusively relies on waiting for timer expirations.
The problem:
This proposal/agreement mechanism is exlcusively point-to-point, it only works between explicit pairs of ports. It does not work in your scenario where a proposal BPDU from A is received by B, C, and D, with all of them potentially sending their own agreement responses (and then these agreements would again get flooded to all switches!).
In this scenario, RSTP convergence becomes undefined and may be unstable and/or introduce temporary loops that could come and go. Now, with your current topology you don't actually have any physical loop, so even with the RSTP convergence imstability, there is no risk of an actual loop forming, as it's physically impossible.
However, if you add a redundant link down the road such that there is a real physical loop that relies on spanning tree to be handled, now the door will be open to encounter more impactful problems.
Personally, I would recommend not using the single AP on the root switch to wirelessly bridge down to your three other switches. If you add two additional APs and have your three switches wirelessly link up to dedicated APs on the root bridge, then you will avoid this problem scenario.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-29-2018 10:23 PM

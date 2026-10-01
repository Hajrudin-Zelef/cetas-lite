---
id: collect-261001-general-networking/general-networking/t5-switching-high-rate-of-stp-bpdu-sender-conflicts-m-p-17785-6fb81503-3
title: "t5-switching-high-rate-of-stp-bpdu-sender-conflicts-m-p-17785-6fb81503"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-switching-high-rate-of-stp-bpdu-sender-conflicts-m-p-17785-6fb81503.md
source_anchor: ""
source_lines: [188, 222]
sha256: 2b2a91f5c88772886912242f94fe8b2d27b70660a52418cb81032fae72f502c0
---

# t5-switching-high-rate-of-stp-bpdu-sender-conflicts-m-p-17785-6fb81503

@test108 wrote:
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
Hi Andrew,
Thank you for all of the info!
I will take what you have said into consideration when deciding on adding redundant links in the future. Perhaps I will then, as you have suggested, rather replace the Rocket AP with three separate PtP links. For now disabling RSTP seems to have solved the problem, as I am no longer getting any errors.

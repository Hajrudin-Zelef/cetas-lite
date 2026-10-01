---
id: collect-261001-cisco/cisco/questions-195-how-can-i-generate-traffic-on-a-cisco-asa-b593c6a0
title: "questions-195-how-can-i-generate-traffic-on-a-cisco-asa-b593c6a0"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["research", "throughput"]
source: docs/RAG/collect-261001-cisco/questions-195-how-can-i-generate-traffic-on-a-cisco-asa-b593c6a0.md
source_anchor: ""
source_lines: [1, 21]
sha256: 9f7040dcb040af9f3974ff573fd34cf4efd7004ea7b8f173d25c40010034a026
---

# questions-195-how-can-i-generate-traffic-on-a-cisco-asa-b593c6a0

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
13
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
In a lot of locations, we have only one Cisco ASA 5505 each and one or more WiFi APs, besides the provider router. No servers or PCs, just WiFi service for occasional visitors. I want to remotely check if the provider gives us the contracted bandwidth. I could see used bandwidth on the ASA monitoring, and test one direction by sending traffic to the ASA using Iperf on my side.
Is there any way to let the ASA generate substantial traffic?
(Just noticed your comment about Raspberrys, and I want to follow up)
We have a large MPLS network where we need to measure bandwith and jitter across sites to make sure that our most sensitive applications isn't hurting.
One great way of doing this is, like you mentioned, deploying Raspberrys. They are small enough, so you can install them nearly anywhere. One great toolf for testing bandwith/jitter is iperf, which you then can set up to run as a (filtered!) daemon on each PI. Then just initiate tests from a central management station.
Downstream you might be able to download something to null: on the ASA but upstream you would be limited of how much the flash disk can transfer which is usually not that much.
Your best bet would be to connect a PC and run the tests but that requires someone to be on site of course.
Perhaps passive monitoring is a better approach. There are plenty of systems out there that will track interface state/errors/discards/bandwidth. Does a bandwidth graph fit your needs?
I had a similar requirement many moons ago with a remote router running IOS. Ended up using tcp-small-servers chargen service on router - runs on tcp/19. It will generate a stream of random characters etc. and can be scaled out using multiple telnet sessions from the remote router to other routers running chargen. Nowhere near the rich control/functionality of host based iperf tests. Also as we know ASA does not support telnet locally so this won't work here...
ASA does have that packet tracer utility since 7.0 to generate interesting traffic for tunnel configs but doesn't allow rate tuning. Such functionality would also be an interesting attack vector if exploited remotely and at scale... Any such native functionality on ASA would likely peg the control plane and any variety of inherent control plane policing would need to restrict the generated traffic rate.
I agree with above that Raspberry Pi is a good solution here. We just dropped some in to drive NOC style screens. They're awesome.
I actually was stuck with a similar request last week (of course the remote site was on the other side of the country). What I ended up doing was using mgen to send unidirectional udp and just verifying the counters on the remote side device. If you aren't comfortable with mgen, you could do the same with nping or one of the nmap tools.
As you pointed out, the issue with iperf, IxChariot, and so many other tools is that they require a remote responder. mgen, nping, and a few others do not.
If you have Solar Winds Engineering Tool Kit, you can use "WAN Killer" between the two locations. Make sure your ACL's allow both ends to talk to each other this way, and flood with the appropriate level of echo's or discard's from Point A to B and visa versa.
We use this to test the networks ability to handle traffic across VPN and MPLS connections, whether it's measuring the raw throughput from A to B and the PPS from A to B.

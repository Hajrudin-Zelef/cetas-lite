---
id: collect-261001-cisco/cisco/questions-5305-cisco-nexus-fabric-extenders-with-10g-sfp-in-path-d9587993
title: "questions-5305-cisco-nexus-fabric-extenders-with-10g-sfp-in-path-d9587993"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "optics", "research"]
source: docs/RAG/collect-261001-cisco/questions-5305-cisco-nexus-fabric-extenders-with-10g-sfp-in-path-d9587993.md
source_anchor: ""
source_lines: [1, 9]
sha256: 20a58aafb443b3d2c281f688fed82487d9092dbd03d1c1eeaa8adff127a486ab
---

# questions-5305-cisco-nexus-fabric-extenders-with-10g-sfp-in-path-d9587993

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
8
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
We are looking to deploy a new data center build-out with Nexus 7ks and Nexus 2ks(FEX). I know the Nexus 2ks come bundled with the Fabric Extenders to be used to uplink to the Nexus 7ks. Our plan however is to put Gigamon fiber taps in-line of all of our uplinks back to our Nexus 7ks. Based off what I understand the FET needs to connect on one end to a FET on the other end. Would the SFP+ used on the in-line fiber taps cause a problem? Does anyone have experience with this? I'm considering just going with 10G SFP+ optics everywhere and not using the FETs.
This shouldn't be an issue, as I've seen customers do it all the time and I've done it in the lab. There is some limitation on distance with the FET-10G vs 10-SR. Also the FET can ONLY be used to connect a N2K to either a N5K or N7K
FETs are not required to connect 2Ks, they are far cheaper than SFPs and for some odd reason (to me) you can only order them when you order 2ks. As Ricky said you are going to see a non standard Ethernet encapsulation but that is a matter of does wireshark support it. Cisco will be happy to take the extra $$s for SFP+

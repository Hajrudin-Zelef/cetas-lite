---
id: collect-261001-cisco/cisco/questions-88734-can-cisco-nexus-3232cs-sfp-1-33-and-1-34-ports-use-1g-optics-8770590c
title: "questions-88734-can-cisco-nexus-3232cs-sfp-1-33-and-1-34-ports-use-1g-optics-8770590c"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["optics", "research"]
source: docs/RAG/collect-261001-cisco/questions-88734-can-cisco-nexus-3232cs-sfp-1-33-and-1-34-ports-use-1g-optics-8770590c.md
source_anchor: ""
source_lines: [1, 15]
sha256: 263f94e078584bd8f5020c159af527d8773733d712b18eb024c462d360e32a5d
---

# questions-88734-can-cisco-nexus-3232cs-sfp-1-33-and-1-34-ports-use-1g-optics-8770590c

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
7
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have searched all over the place and I cannot find a definitive answer to this:
Can Cisco Nexus 3232C's SFP+ 1/33 and 1/34 ports use 1G optics?
For "reasons," it would be really convenient to plug in a 1G from a legacy connection and tie to a 100GbE VLAN.
You can see the ports on the left hand side, and the Cisco documentation says "The Cisco Nexus 3232C (N3K-C3232C) is a 1 rack unit (RU) switch with 32 10- or 100-Gigabit QSFP28-100 and 2 10G SPF+ ports." As far as I can tell, they are connected to the normal switching backplane and are not management ports (mgmt is on the back and has SFP mgmt, separately).
Does anyone have a source that can answer this question definitively?
"Can Cisco Nexus 3232C's SFP+ 1/33 and 1/34 ports use 1G optics?"
No, that particular model does not support 1 Gbps SFPs. The only interface on the switch that does support 1 Gbps is the management interface that cannot be use for switching.
There are other Nexus 3000 series that do support 1 Gbps, either directly on the interface, SFP+ slot, or QSFP slot using a QSFP adapter. You also need to be careful in that some Nexus switches will not support 1 Gbps in certain SFP+ or QSFP slots. I saw that problem bite a consultant on a Nexus switch where the first few and last few QSFP slots would not use 1 Gbps, but all in between would.
Hard to provide a source for something that is not possible. Checking the data sheet, there is no mention of 1G connectivity, just 10G, 40G and 100G, and 25G via breakouts. The transceiver compatibility matrix provides the same information. I'd take that as face value that there's no 1G connectivity.

---
id: collect-261001-cisco/cisco/questions-4939-what-path-does-a-packet-follow-through-a-cisco-nexus-5k-2k-infras-cad10ad9
title: "questions-4939-what-path-does-a-packet-follow-through-a-cisco-nexus-5k-2k-infras-cad10ad9"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/questions-4939-what-path-does-a-packet-follow-through-a-cisco-nexus-5k-2k-infras-cad10ad9.md
source_anchor: ""
source_lines: [1, 6]
sha256: 320074dcb1ccd88ccdb49156a1ff8347bbdbe12c24fe613512bb4dcd97321550
---

# questions-4939-what-path-does-a-packet-follow-through-a-cisco-nexus-5k-2k-infras-cad10ad9

I was told by someone that a packet going between two hosts connected to the same Nexus 2k would end up being passed all the way up to the 5k that is managing said 2k for logic to be applied and then passed back down to the 2k and out the port the destination host is connected to. That seems extremely inefficient to me, so I'm looking for someone to either confirm or deny the information I've been given.
1 Answer 1
The Nexus 2k isn't an actual switch, but a "fabric extender" - it functions as a separate, self-contained linecard for the Nexus 5k. The following CiscoLive PDF has great details on the Nexus 2k/5k architecture
It's clearly pointed out in page 56 that the 2k does not perform any local switching.
The Nexus 5k is what performs all of the forwarding, integrating the Nexus 2k fabric extenders remotely into its backplane via IEEE 802.1BR.
As a future point of advice, googling for "a day in the life of a packet on [x]" where [x] is your favorite piece of gear will often produce good documentation on the architecture of that gear. YMMV with this however, as not all vendors make this information public.

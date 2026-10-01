---
id: collect-261001-general-networking/general-networking/questions-62463-will-om4-fiber-work-with-my-1gbe-fiber-transceiver-7397284e
title: "questions-62463-will-om4-fiber-work-with-my-1gbe-fiber-transceiver-7397284e"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-62463-will-om4-fiber-work-with-my-1gbe-fiber-transceiver-7397284e.md
source_anchor: ""
source_lines: [1, 9]
sha256: 970f56969fce03d4b63ba3d6b6f43a6ef49b30251717ffbc3b35b2cee9ecfb79
---

# questions-62463-will-om4-fiber-work-with-my-1gbe-fiber-transceiver-7397284e

What is the fastest cabling that can be used to connect the transceiver that Meraki calls the 1GbE SFP SX Fiber Transceiver?
Source: meraki.cisco.com/products/switches/accessories
I'm trying to purchase fiber cabling between two switches that are 459 feet (140 meters) apart.
Even though the switches themselves are limited to 1GbE, I want to purchase the fastest cabling that will function with these switches' transceivers. My priority is on flexibility (not price): When the day comes, that these switches are upgraded to ones capable of much faster speeds, the cable I purchase now should accommodate those speeds.
Currently, the two switches are Meraki MS120-8, and they are equipped with 4 of the transceivers that are pointed out in the top screenshot of this post.
Ultimately, I'll be running two separate fiber cables (both 150 meters long) to hopefully get closer to 2GbE speed (via link aggregation). These cables will take different physical paths between these two switches, to reduce the likelihood of both cables being damaged at the same time.
Meraki won't advise me on the cabling at all. For example, I couldn't even get my Meraki rep to even speculate as to whether or not this [OM4 cable][5] would work. I need someone with experience to advise me. Otherwise, I'm just going to have order short strands of each prospective type I'd want, and just see which ones work before buying longer strands.
What is the fastest cabling that can be used to connect the transceiver that Meraki calls the 1GbE SFP SX Fiber Transceiver?
UPDATE: OM4 fiber cable was purchased and is now installed. It is working perfectly. Thanks for everyone's confirmation!

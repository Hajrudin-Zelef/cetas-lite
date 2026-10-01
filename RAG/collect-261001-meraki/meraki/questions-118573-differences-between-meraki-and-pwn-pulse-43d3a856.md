---
id: collect-261001-meraki/meraki/questions-118573-differences-between-meraki-and-pwn-pulse-43d3a856
title: "questions-118573-differences-between-meraki-and-pwn-pulse-43d3a856"
domain: meraki
role: reference
task: reference
actors: ["Apple", "Microsoft"]
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-meraki/questions-118573-differences-between-meraki-and-pwn-pulse-43d3a856.md
source_anchor: ""
source_lines: [1, 16]
sha256: 9d594e30bd3467f17ed9406592b58a1e742c64b0c8ff0fe2d749b7c18b648afb
---

# questions-118573-differences-between-meraki-and-pwn-pulse-43d3a856

Information Security is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
0
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I am working on the beginning of a writing assignment where I will document an emerging technology in information security. I have found this to be difficult because as the new technology emerged with regards to new attack vectors - they are soon patched, replaced, etc.
Anyways - I am considering writing about products such as pwn pulse. My question is, is there any immediate difference between the pwn pulse product and a product such as Cisco's Meraki?
Secondly, they seem to refer to these solutions simply as 'BYOD' solutions. Is this the accurate and appropriate terminology for these types of products? I ask because I will be forced to search through peer journals (and online databases) to research and document anything I can find that I can relate to such a paper and I don't anticipate finding much on 'BYOD'.
Regarding your specific question, the two capabilities share some overlap when it comes to BYOD, but there are some key differences:
- Meraki "provides device based security policies, built-in NAC, and built-in mobile device management."
- Pwn Pulse "detects, fingerprints, and analyzes rogue, misconfigured, and unauthorized wireless and wired devices including BYOx/mobile, Wireless, Bluetooth, wired and other network-enabled devices."
In short, the two systems are apples and oranges compared to each other--one (Meraki) is a cloud-enabled Wi-Fi access solution with Mobile Device Management capabilities, kind of like Microsoft Intune. The other (Pwn Pulse) is a cloud-enabled device visibility + threat detection solution. As you've already alluded, BYOD is a bit of a vague term as it can mean different things depending upon where you're coming from. I would argue that both of these offer BYOD solutions in that they both resolve one aspect of understanding what devices are in/around airspace demanding network resources, but this is but one aspect of the larger product offering.
Because your first question is asking for the difference between two products, it is considered a product recommendation kind of question and is unfortunately off topic here. But you are as capable of reading and comparing their marketing blurbs as anyone else, so that shouldn't prove a problem for your homework.
To answer your second question, these security products attempt to address the risks posed by user-owned devices in a Bring-Your-Own-Device (BYOD) environment. A BYOD organization explicitly grants permission for users to bring their own iPhones, tablets, laptops, etc., and all these devices could be infected with malware, Trojan horses, viruses, worms, or worse.
Because that's a mouthful to say, simply calling them "BYOD" seems to be in vogue, even though it's not a factually correct term for these devices. Try searching for "BYOD security" or "mobile device management" and see what else you come up with.

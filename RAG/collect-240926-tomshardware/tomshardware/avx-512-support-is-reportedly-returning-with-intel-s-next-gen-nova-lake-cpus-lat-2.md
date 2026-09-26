---
id: collect-240926-tomshardware/tomshardware/avx-512-support-is-reportedly-returning-with-intel-s-next-gen-nova-lake-cpus-lat-2
title: "avx-512-support-is-reportedly-returning-with-intel-s-next-gen-nova-lake-cpus-lat"
domain: tomshardware
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["intel", "consumer", "cost", "research"]
source: docs/RAG/clean_en/tomshardware/avx-512-support-is-reportedly-returning-with-intel-s-next-gen-nova-lake-cpus-lat.md
source_anchor: ""
source_lines: [63, 103]
sha256: c723f8676199a1838138ee5ef962a99e864b245baad1767a9f6dd9b0c2a8bc2f
---

# avx-512-support-is-reportedly-returning-with-intel-s-next-gen-nova-lake-cpus-lat

What had happened is that Intel simply disabled it in a way that some motherboard makers figured out how to hack around, in the BIOS. Intel quickly forced those motherboard makers back in line (newer BIOSes no longer had the hack) and the updated microcode slammed the door shut on any more of these shenanigans.
- 
Reply
 Initially, but then they stated something else, you're like a year and a half behind...bit_user said:That's not accurate. Intel had stated that hybrid CPUs would support only 256-bit. That means both core types in such a CPU would be restricted.
 Initially, Intel planned to remedy this with the Intel AVX10 specification, which was intended to strictly limit consumer processors to 256-bit instructions (AVX10/256) so it could run uniformly across both Performance-cores (P-cores) and E-cores. 1, 2, 3, 4, 5]However, this 256-bit limitation is no longer the case. Intel ultimately removed the 256-bit cap from the AVX10 specification and officially confirmed the return of full-width 512-bit vector execution in their consumer lineup. Moving forward, architectures like the upcoming Nova Lake will natively support 512-bit execution across both P-cores and E-cores. 1, 2, 3, 4, 5]
- 
Reply
 Yes, the precise mechanism the BIOSes of those boards used to enable it was (I think accurately) characterized as a hack.helper800 said:It was not really a hack either though.
 
 Here's how Anandtech (R.I.P.) described it:
 Dr. Ian Cutress (Anandtech) said:We’ve done some extensive research on what Intel has done in order to ‘disable’ AVX-512. It looks like that in the base firmware that Intel creates, there is an option to enable/disable the unit, as there probably is for a lot of other features. Intel then hands this base firmware to the vendors and they adjust it how they wish. As far as we understand, when the decision to drop AVX-512 from the POR was made, the option to enable/disable AVX-512 was obfuscated in the base firmware. The idea is that the motherboard vendors wouldn’t be able to change the option unless they specifically knew how to – the standard hook to change that option was gone.
 
 However, some motherboard vendors have figured it out. In our discoveries, we have learned that this works on ASUS, GIGABYTE, and ASRock motherboards, however MSI motherboards do not have this option. It’s worth noting that all the motherboard vendors likely designed all of their boards on the premise that AVX-512 and its high current draw needs would be there, so when Intel cut it, it meant perhaps that some boards were over-engineered with a higher cost than needed. I bet a few weren’t happy.
 **Source:** https://web.archive.org/web/20240906190403/https://www.anandtech.com/show/17047/the-intel-12th-gen-core-i912900k-review-hybrid-performance-brings-hybrid-complexity/2
 The article has other background info about the subject (including history behind it and the decision to remove it), for those who are interested.
 
 
 Once the feature was enabled, it was enabled. Yes, the P-cores really did have it.helper800 said:You could perform those workloads if you wanted so clearly the hardware physically supported the task.
 
 
I guess the word "hack" confused you. Otherwise, it seems unwise to argue about something you don't know very well. Much better to ask questions to test what you think you know or that would fill in the gaps in your knowledge (that you're aware of).helper800 said:I do not know much more on this topic though, so if there is nuance here I certainly missed it back when all this came out.
- 
Reply
 It was 100% the word 'hack.' I interpreted that as the motherboard venders enabling a feature set Intel deemed dangerous to the chip, not that they circumvented an artificial limitation. What was Intel's purpose in disabling AVX 512 at that point? Did they just not want it in their consumer chips for product segmentation reasons?bit_user said:Yes, the precise mechanism the BIOSes of those boards used to enable it was (I think accurately) characterized as a hack.
 
 Here's how Anandtech (R.I.P.) described it:
 **Source:** https://web.archive.org/web/20240906190403/https://www.anandtech.com/show/17047/the-intel-12th-gen-core-i912900k-review-hybrid-performance-brings-hybrid-complexity/2
 The article has other background info about the subject (including history behind it and the decision to remove it), for those who are interested.
 
 
 Once the feature was enabled, it was enabled. Yes, the P-cores really did have it.
 
 
I guess the word "hack" confused you. Otherwise, it seems unwise to argue about something you don't know very well. Much better to ask questions to test what you think you know or that would fill in the gaps in your knowledge (that you're aware of).
- 
Reply
 The Anandtech article (readable via the Archive.org link I included) explores the history and rationale of it being disabled. So, I'd just refer you to that page.helper800 said:What was Intel's purpose in disabling AVX 512 at that point? Did they just not want it in their consumer chips for product segmentation reasons?
 
 I can actually see the logic behind it, since the potential workaround of trying to fault AVX-512 threads over to the P-cores has some nasty performance pitfalls for naive code.
 
And, even if they limited AVX-512 to just when the E-cores are disabled, they didn't want to draw more fire to their hybrid strategy and give people yet more reasons to disable them. Of course, as we now know, those criticisms came anyway.

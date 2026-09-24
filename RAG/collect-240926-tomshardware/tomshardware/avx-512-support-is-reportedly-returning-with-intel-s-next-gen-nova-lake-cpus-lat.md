---
id: collect-240926-tomshardware/tomshardware/avx-512-support-is-reportedly-returning-with-intel-s-next-gen-nova-lake-cpus-lat
title: "avx-512-support-is-reportedly-returning-with-intel-s-next-gen-nova-lake-cpus-lat"
domain: tomshardware
role: reference
task: reference
actors: ["AMD", "Google", "Intel"]
dates: []
keywords: ["intel", "amd", "compute", "consumer", "cost", "research"]
source: docs/RAG/clean_en/tomshardware/avx-512-support-is-reportedly-returning-with-intel-s-next-gen-nova-lake-cpus-lat.md
source_anchor: ""
source_lines: [1, 103]
sha256: ef1a60c53c078772eddf69a1411d83e4d44c9ad7542a214655f7df42c8072719
---

# avx-512-support-is-reportedly-returning-with-intel-s-next-gen-nova-lake-cpus-lat

<!-- source: https://www.tomshardware.com/pc-components/cpus/avx-512-support-is-reportedly-returning-with-intels-next-gen-nova-lake-cpus-latest-linux-kernel-patches-reveal-p-cores-and-e-cores-will-gain-native-512-bit-execution -->

When Intel switched to a hybrid architecture with its 12th-Gen Alder Lake PUs, it removed AVX-512 support from the lineup entirely because the E-cores didn't support it. Since then, every subsequent generation has shipped without it... until now. Just today, a new Linux patch pushed in the RAID optimized path has revealed that AVX-512 is finally returning to Intel CPUs with Nova Lake, present on both P-cores and E-cores.

Intel has been working toward a unified AVX solution for the past few years, as it was originally a champion of the SMID extensions before running into the hybrid hurdle with Alder Lake. Getting past that hurdle is AVX10, which Intel first detailed a few years back. With AVX10.2, 512-bit instructions will run on the P-cores, while either core type can handle converged 256-bit instructions.

As such, the E-cores would have their processing width capped at 256-bit, while the P-cores would be open to the full 512-bit wide pipelines. Any thread could swiftly move between either core type with AVX10 implemented. Previously, if the scheduler shifted a 512-bit task running on a P-core to an E-core, the application would crash instantly because those E-cores couldn't process the instruction.

However, the new patches suggest that Intel has now mandated native 512-bit execution across both P-cores and E-cores, no longer requiring the latter to step down and process the data a bit slower. This is a major development over the standard we originally expected Intel to adopt; the E-cores are apparently becoming just as performant as the P-cores when it comes to SIMD instructions with Nova Lake and later.

Intel's original announcement showed that it had already uncoupled the software improvements of AVX-512 from the physical width of the register, so the new instruction features could remain present for both 512-bit and 256-bit execution. This includes things like masking, embedded broadcast for rounding math operations, and doubling the number of the registers themselves from 16 to 32.

It remains unclear if we will ever see this version of AVX10 on client CPUs, as it seems Nova Lake is going purely for 512-bit execution across both core types. AMD's current-gen Zen 5 processors also have full 512-bit wide registers, while the previous Zen 4 architecture divided a single 512-bit task across two 256-bit execution units over two clock cycles. This ensured execution remained disruption-free even if it took longer.

The last time we saw native AVX-512 support on an Intel client family was Rocket Lake (11th Gen), right before the hybrid era ushered in by Alder Lake. For modern AI workloads and other compute-heavy tasks such as encoding or simulations, AVX-512 instructions bring a huge performance benefit that's foolish to be left on the table. Keep in mind that this is just a Linux patch at the moment and that Intel hasn't officially announced native AVX-512 support for Nova Lake yet.

Get Tom's Hardware's best news and in-depth reviews, straight to your inbox.

*Follow* *Tom's Hardware on Google News**, or* *add us as a preferred source**, to get our latest news, analysis, & reviews in your feeds.*

- 
The article is partially wrong when it says this:Reply
 
 12th gen shipped with AVX 512, but was later disabled through BIOS updates, and then a silicon revision. The article makes it seem like it was never available on 12th gen...Admin said:When Intel switched to a hybrid architecture with its 12th-Gen Alder Lake PUs, it removed AVX-512 support from the lineup entirely because the E-cores didn't support it.
 
The last time we saw native AVX-512 support on an Intel client family was Rocket Lake (11th Gen), right before the hybrid era ushered in by Alder Lake.
- 
Reply
 That's not accurate. Intel had stated that hybrid CPUs would support only 256-bit. That means both core types in such a CPU would be restricted.The article said:With AVX10.2, 512-bit instructions will run on the P-cores, while either core type can handle converged 256-bit instructions.
 
 As such, the E-cores would have their processing width capped at 256-bit, while the P-cores would be open to the full 512-bit wide pipelines.
 
 However, Intel then got rid of the 256-bit limit, later saying that both core types would support full 512-bit mode (**source:** https://www.phoronix.com/news/Intel-AVX10-Drops-256-Bit ).
 
 The one thing they never said was that you'd have different core types, in a hybrid CPU, supporting different modes!
 
 
 Ugh, no. This is a misunderstanding of how AVX-512 works.The article said:Intel's original announcement showed that it had already uncoupled the software improvements of AVX-512 from the physical width of the register, so the new instruction features could remain present for both 512-bit and 256-bit execution.
 
 Since its introduction, AVX-512 consisted of a set of vector instructions that could take either 128-bit, 256-bit, or 512-bit operands. What they proposed in AVX10 was implementations which limited the maximum operand size to either 128-bit, 256-bit, or 512-bit. An architectural register was added, so that you could check the capabilities of the CPU and select a code path that didn't use any operands which were too wide.
 
 In other words, AVX10 didn't introduce the multiple operand sizes facility, it just permitted the CPU to place a cap on what sizes it supported. And if you can't execute any 512-bit instructions, then the registers also don't need to be 512 bits wide. So, that's how you can think about the relationship between the instruction set and register sizes.
 
 
 AVX-512 is when they increased from 16 to 32. Using regular AVX/AVX2 instructions (or SSE, in 64-bit mode), you can access only 16 vector registers. AVX-512 provides access to 32 of them, and AVX10 didn't change that.The article said:This includes things like masking, embedded broadcast for rounding math operations, and doubling the number of the registers themselves from 16 to 32.
 
 
 No, they killed it. AVX10 will always allow 512-bit operands.The article said:It remains unclear if we will ever see this version of AVX10 on client CPUs
 
 
Laptop Zen 5 cores work the same way as Zen 4, where most 512-bit operations are actually implemented as a pair of 256-bit ones. I assume Zen 5C cores are also like this, but haven't heard one way or the other.The article said:AMD's current-gen Zen 5 processors also have full 512-bit wide registers, while the previous Zen 4 architecture divided a single 512-bit task across two 256-bit execution units over two clock cycles.
- 
Reply
 It wasn'thelper800 said:The article is partially wrong when it says this:
 
 12th gen shipped with AVX 512, but was later disabled through BIOS updates, and then a silicon revision. The article makes it seem like it was never available on 12th gen...*formally* enabled. End users were never expected to see it.
 
 What had happened is that Intel simply disabled it in a way that*some* motherboard makers figured out how to hack around, in the BIOS. Intel quickly forced those motherboard makers back in line (newer BIOSes no longer had the hack) and the updated microcode slammed the door shut on any more of these shenanigans.
- 
Reply
 It was not really a hack either though. You could perform those workloads if you wanted so clearly the hardware physically supported the task. I do not know much more on this topic though, so if there is nuance here I certainly missed it back when all this came out.bit_user said:It wasn't*really* enabled. End users were never expected to see it.
 
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

---
id: collect-240926-tomshardware/tomshardware/amd-makes-fsr-4-upscaling-official-for-radeon-rx-7000-and-6000-series-cards-rdna
title: "amd-makes-fsr-4-upscaling-official-for-radeon-rx-7000-and-6000-series-cards-rdna"
domain: tomshardware
role: reference
task: reference
actors: ["AMD", "Google", "Intel", "Nvidia"]
dates: ["2025-08"]
keywords: ["amd", "compute", "cost", "fp8", "gpu", "intel", "nvidia"]
source: docs/RAG/clean_en/tomshardware/amd-makes-fsr-4-upscaling-official-for-radeon-rx-7000-and-6000-series-cards-rdna.md
source_anchor: ""
source_lines: [1, 57]
sha256: 1a1cb74f94f5fb00043116c02a410f2b0e3a233c2077ab1e2845c7d11a12ff35
---

# amd-makes-fsr-4-upscaling-official-for-radeon-rx-7000-and-6000-series-cards-rdna

<!-- source: https://www.tomshardware.com/pc-components/gpu-drivers/amd-makes-fsr-4-upscaling-official-for-radeon-rx-7000-and-6000-series-cards-rdna-3-and-rdna-2-chips-will-soon-enjoy-improved-visuals -->

Owners of Radeon RX 9000-series cards have been enjoying the benefits of the company's FSR 4 upscaling tech for some time now, and that feature has officially been exclusive to those products since launch. Gamers with older Radeons were left out in the cold, sparking community outrage—until today. AMD VP Jack Huynh has revealed that FSR 4.1 upscaling will be made available for RDNA 3 cards (RX 7000-series products) in July, and for RDNA 2 cards (RX 6000-series products) in "early 2027."

Community testing of AMD's INT8 FSR 4 code puts the penalty at around 10-20% versus FSR 3 on 6000-series cards, and a lower cost on 7000-series Radeons, but only testing with the full official version will tell the full story. Having said that, even with lower performance scaling, the quality-to-speed ratio of FSR 4.1 is almost certainly worth it.

The company has hinted at open-sourcing FSR4 in the wake of the aforementioned leak, a move that would probably be a good idea for the future of the technology, given how AMD's GPU division seems focused on catching up to Nvidia in the far more lucrative market AI accelerators. A more recent SDK update also suggested that FSR frame generation might get 4-6x multipliers, too, which would give it feature parity with Nvidia's MFG. In any event, making FSR 4 available to owners of older Radeons is a welcome step that will extend the useful life of those products for some time.

The news is certainly welcome for gamers with older Radeons, though technically minded folks already had access to FSR 4's improvements by way of community tools like Optiscaler. AMD initially restricted FSR 4 to RDNA 4's accelerated FP8 hardware, but an FSR 4 source code leak in August 2025 revealed that the company had also created an INT8 version of the AI upscaling model that was compatible with older cards. The community used that source code to enable support for FSR 4 on older Radeons through unofficial tools, creating an ongoing outcry for official support that's now been answered.

It's notable that the upcoming official release will incorporate the latest FSR 4.1 release. That upscaler improves on the original FSR 4.0 on most every front, with less blurring and smearing, better detail retention on thin lines and distant retention, and finer particle effects. There's also significantly less shimmer on object edges (aka improved temporal stability).

*Follow* *Tom's Hardware on Google News**, or* *add us as a preferred source**, to get our latest news, analysis, & reviews in your feeds.*

Get Tom's Hardware's best news and in-depth reviews, straight to your inbox.

Bruno Ferreira is a contributing writer for Tom's Hardware. He has decades of experience with PC hardware and assorted sundries, alongside a career as a developer. He's obsessed with detail and has a tendency to ramble on the topics he loves. When not doing that, he's usually playing games, or at live music shows and festivals.

- 
While the immediacy of RDNA 3 getting it is probably driven by the integrated proliferation (and lack of RDNA 4 here) I'm glad they're moving just the same. The wait for RDNA 2 seems pretty dumb, but at least they've officially committed to it.Reply
- 
Reply
 I don't know that it'll have that effect, lol.ManDaddio said:Well it's about time. But that means the Radeon used market prices are going way up now.
 Steam box might use fsr4 now?
Yes, we're assuming that Steam Deck ("early 2027") and Steam Machine will get FSR4 support now.
- 
Nice! As a 7900XTX user I'm quite pleased with this even though I'm not playing any games where I can't get 60fps@4K yet (the most my primary screen can do).Reply
 
Good to know this'll be available when I get round to playing games that need it for a decent experience.
- 
As a 9070XT owner I think this is great. Wider install base means more dev participation.Reply
 
 But 2 corrections: 1. FSR4.1 has more shimmering than FSR 4.0, probably revealed by the increased sharpness and clarity, and 2. The performance hit on RDNA 3,2 relative to RDNA4 will be noticeable as the 9060XT has more significantly more 8 bit than the 7900XTX , much less the 6900XT.
 
It should still be completely worth it for iGPUs though as AMD has fallen pretty far behind Intel in the iGPU upscaling category and this helps them catch up. It would be interesting to see a comparison of these iGPUs with upscaling enabled now that AMD iGPUs are getting something comparable to XeSS.
- 
Reply
 Is FSR4 support in enough games to justify prices going way up?ManDaddio said:Well it's about time. But that means the Radeon used market prices are going way up now.
 Steam box might use fsr4 now?
 
 Yes, the Steam Machine should have it. Possibly at launch if it comes out in July or later, or updated to have it very soon after launch.
 
 
People have already been testing FSR4 on RDNA 2/3 for months. I think the performance hit was less than expected given the major TOPS advantage RDNA4 cards have. But gamers may have to use a lower FSR4 quality level, or accept losing some FPS in exchange for better quality.rluker5 said:The performance hit on RDNA 3,2 relative to RDNA4 will be noticeable as the 9060XT has more significantly more 8 bit than the 7900XTX , much less the 6900XT.
- 
Reply
 The main problem with RDNA2's implementation wasn't speed but quality, because while RDNA3 could use INT8 in the matrix computations that RDNA4 does in FP8 for FSR4 (which can be done with little to no rounding errors), RDNA2 can't do wave matrix multiply accumulate (WMMA) at all and must rely on this being emulated on top - resulting in far more rounding errors and thus much lower quality and a much higher compute cost.usertests said:Is FSR4 support in enough games to justify prices going way up?
 
 Yes, the Steam Machine should have it. Possibly at launch if it comes out in July or later, or updated to have it very soon after launch.
 
 
People have already been testing FSR4 on RDNA 2/3 for months. I think the performance hit was less than expected given the major TOPS advantage RDNA4 cards have. But gamers may have to use a lower FSR4 quality level, or accept losing some FPS in exchange for better quality.
- 
Reply
 Didn't hear about that but OK.mitch074 said:The main problem with RDNA2's implementation wasn't speed but quality, because while RDNA3 could use INT8 in the matrix computations that RDNA4 does in FP8 for FSR4 (which can be done with little to no rounding errors), RDNA2 can't do wave matrix multiply accumulate (WMMA) at all and must rely on this being emulated on top - resulting in far more rounding errors and thus much lower quality and a much higher compute cost.
 
If AMD is giving themselves that much more time, maybe what they'll come up with will be better than the leaked version.

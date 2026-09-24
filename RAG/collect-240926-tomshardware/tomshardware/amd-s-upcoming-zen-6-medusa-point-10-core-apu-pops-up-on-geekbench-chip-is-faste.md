---
id: collect-240926-tomshardware/tomshardware/amd-s-upcoming-zen-6-medusa-point-10-core-apu-pops-up-on-geekbench-chip-is-faste
title: "amd-s-upcoming-zen-6-medusa-point-10-core-apu-pops-up-on-geekbench-chip-is-faste"
domain: tomshardware
role: reference
task: reference
actors: ["AMD", "Google", "Samsung"]
dates: []
keywords: ["amd", "benchmark", "benchmarks", "fp8", "memory"]
source: docs/RAG/clean_en/tomshardware/amd-s-upcoming-zen-6-medusa-point-10-core-apu-pops-up-on-geekbench-chip-is-faste.md
source_anchor: ""
source_lines: [1, 39]
sha256: 5f399d45e03b746cc53d614af3e8e1cbf220015fb6d9f66d9d5d7a5d325ee4f1
---

# amd-s-upcoming-zen-6-medusa-point-10-core-apu-pops-up-on-geekbench-chip-is-faste

<!-- source: https://www.tomshardware.com/pc-components/cpus/amds-upcoming-zen-6-medusa-point-10-core-apu-pops-up-on-geekbench-chip-is-faster-than-ryzen-ai-9-hx-370-and-even-ryzen-ai-max-395 -->

AMD is expected to announce its next-gen mobile CPUs at CES 2027, but leaks have already started to pour in, giving us a decent idea of the performance we can expect. Codenamed "Medusa Point," the Red Team's upcoming lineup will likely be based on the Zen 6 microarchitecture, and one of the SKUs has just popped up on Geekbench. It scored much better than the previous leak, beating most of its contemporaries.

The part showed up as "AMD Eng Sample 100-000001713-33_N" and was marked under "AMD Plum-MDS1," which we know is the platform associated with Medusa Point. It's a 10-core (4+6) chip, with 20 threads, clocked at roughly 2.0 GHz, carrying 10MB of L2 cache and 32MB of L3 cache. The L3 cache and clock speeds might be misreported. Currently, AMD only makes two other 10-core mobile parts — Ryzen AI 9 365 and Ryzen AI 9 465, so we're most likely looking at a purported Ryzen AI 9 565 here.

Coming to the scores, the chip netted 3,174 points in the single-core test and 15,092 points in the multi-core test. Both of those numbers are higher than the Strix Point flagship APU, the Ryzen AI 9 HX 370. On average, that SKU sits around 2,600 single-core points, so the Medusa Point score is 22% higher. In multi-core, the AI 9 HX 370 gets 13,400 points, making our main contender 13% faster on average.

It even beats the Strix Halo flagship, the Ryzen AI 9 Max+ 395, by over 400 points in the single-core benchmark, but loses in the multi-core test. Of course, the onboard graphics is no comparison between the two. Compared to a prior leak also showcasing a 10-core Medusa Point APU, this new listing is significantly better. The previous one came in at only 2,300 single-core and 13,002 multi-core points.

It seems like Zen 6 offers a noticeable leap in performance based on architectural improvements, since the core count between the chips we compared is identical. It's too early to judge anything, though, since Medusa Point is months away at this point, and this is just one SKU from the lineup. The top-end parts almost carry a mandate to be faster than their direct predecessor to be even worth releasing; it's the midrange where the real value proposition lies.

*Follow* *Tom's Hardware on Google News**, or* *add us as a preferred source**, to get our latest news, analysis, & reviews in your feeds.*

Get Tom's Hardware's best news and in-depth reviews, straight to your inbox.

- 
Medusa Point (MDS1) is expected to be the successor to 8-core Krackan Point. The true successor to Strix Point would be "Medusa Halo Mini" (with ~14 cores and 24 CUs of RDNA5).Reply
 
 Don't know about the L3 cache, but the 2 GHz anomaly could be caused by new "LP" cores.
 
It probably has "RDNA4m" graphics for full speed FSR4 FP8/WMMA.
- 
Most of the rumors are pointing at Medusa Point having 4+4c Zen 6 cores, and then 2 lp cores off on their own for idle/background stuff type tasks. Which makes those multi-thread benches look even better; the lp cores would not really be used in a bench, so its 4+4c Zen6 handily outperforming 4+6c Zen 5.Reply
- 
Reply
 That might have been true of Meteor Lake where it was hard to even get the LPE cores to be used (not sure if that behavior changed with updates), but not true of later chips. Benchmarks could be using these LP cores.GenericUser2001 said:the lp cores would not really be used in a bench, so its 4+4c Zen6 handily outperforming 4+6c Zen 5.
 
However, it's possible that two LP cores together don't even equate to 1 classic core's worth of performance when under heavy load, if their max turbo is somewhere around 2-3 GHz. And Medusa Point outperforming Strix Point 370 is definitely a nice result, when it is only the Krackan Point replacement.
- 
The 2GHz clock speed is obviously not correct because the single threaded couldReply*never* be that high if it was. The performance level does seem to indicate they're close to whatever real world targets are likely to be. That being said I'm curious what the release schedule is going to look like. Typically AMD has released mobile parts later, but Zen 6 is certainly looking very atypical release wise. The other issue being the 400 series hasn't been on the market for very long so something that completely outclasses it might not be welcome so soon to OEMs given that sales are likely very bad.
- 
Reply
If anyone knew what memory prices will look like when this releases, well they would be either buying Samsung/Sk Hynix/Micron stock, or buying shorts in those companies.thisisaname said:Will memory be affordable by the time this comes out?
- 
Reply
 It's faster than Strix Point (365/370), and supposed to have faster ST than Halo. It's still 0.4% faster ST than your score, but I assume the number they were comparing it to was the average of all 395 tests.Puchu said:So how is this faster?
 
It's obvious for Zen 6 to have higher single-threaded performance than Zen 5. The Ryzen AI MAX+ 395 only has a stock 5.1 GHz turbo clock which will not be hard for even budget Zen 6 APUs to match. Since this is an early result with bugged reporting, we can't tell what clocks this is really running at.

---
id: collect-240926-tomshardware/tomshardware/apple-launches-new-m6-and-m5-ultra-apple-silicon-chips-debuting-in-new-mac-mini-2
title: "apple-launches-new-m6-and-m5-ultra-apple-silicon-chips-debuting-in-new-mac-mini-"
domain: tomshardware
role: reference
task: reference
actors: ["Apple", "Intel"]
dates: []
keywords: ["benchmarks", "gpu", "intel", "lpddr5x", "memory"]
source: docs/RAG/clean_en/tomshardware/apple-launches-new-m6-and-m5-ultra-apple-silicon-chips-debuting-in-new-mac-mini-.md
source_anchor: ""
source_lines: [68, 94]
sha256: d03820d45595b4f9999996c1a9fc3829be739485b0e5659d46c65e1b8caa5c69
---

# apple-launches-new-m6-and-m5-ultra-apple-silicon-chips-debuting-in-new-mac-mini-

These days, I guess the main point of genlock is probably for video walls.
- 
Reply
 At least you can cluster them with Thunderbolt 5.Doomsday7 said:I am not sure about it. Even 64GB unified RAM do mean safety operating LLMs with 36B only. So as long as Apple does not find itself a way out of RAM crisis… Stay with my old Macbook Pro is driving me nuts. Where is that new RAM reduction as Purdue Univeristy is proposing to half LLM memory need and increase speed significantly? This is what Apple formerly was known for: Technology lead. Maybe M8?
 
 Got a link for Purdue proposal?
 
 If the middle cores are not significantly slower than the best/super cores, there may be no real loss to single-threading across 4 cores (amount of "super" cores M5 has) accounting for any small improvements they made. I don't know enough about Apple's core types to say.User of Computers said:They make no claims about M6's single-core performance compared to its predecessor, leading me to believe that there are only minor architectural improvements to be had. In that case, it strikes me as odd that Apple would reduce the number of S-cores from M5 to make room for their new middle P-cores.
 If I had to guess, it probably has something to do with multi-threaded effiency.
 
 Either way, I'm guessing it would be rare for the M6 to underperform vs. the M5. The MT boost will be noticeable, while a small ST deficit on two cores would be hidden.
 
My theory based on the rumored fast M6 and M7 release schedule, and that there will be no M6 Max/Ultra, is that this is a quick refinement of the M5 with LPDDR5X for "budget" products. M7-based products will be released shortly with substantial boosts to AI performance, but using more expensive LPDDR6 memory.
- 
Reply
 Eh, multi-GPU has definite bottlenecking issues and limitations. You'd at least want to use PCIe 5.0 x16, which is way faster than TB5.usertests said:At least you can cluster them with Thunderbolt 5.
 
Benchmarks of Intel's BattleMatrix solution showed it could deliver gains on batch processing workloads, but that would primarily interest service providers and not most end users.
- 
Reply
 I've seen 6% quoted by a lot of people and they reference https://nanoreview.net/en/cpu-compare/apple-m6-12-core-vs-apple-m5. However, I don't know if that is an official run or "estimated", I've so far taken it with a grain of salt.User of Computers said:They make no claims about M6's single-core performance compared to its predecessor, leading me to believe that there are only minor architectural improvements to be had. In that case, it strikes me as odd that Apple would reduce the number of S-cores from M5 to make room for their new middle P-cores.
 If I had to guess, it probably has something to do with multi-threaded effiency.
 
Like you, I'm guessing the IPC improvements are marginal to nonexistent. It could largely die shrink with small tweaks to accommodate two additional cores and the change of two LP cores to E/Middle cores (Apple calls them "performance cores")
- 
Reply
I definitely agree. There shouldn't be anything noticible from consumers' or reviewers' perspectives.usertests said:My theory based on the rumored fast M6 and M7 release schedule, and that there will be no M6 Max/Ultra, is that this is a quick refinement of the M5 with LPDDR5X for "budget" products. M7-based products will be released shortly with substantial boosts to AI performance, but using more expensive LPDDR6 memory.

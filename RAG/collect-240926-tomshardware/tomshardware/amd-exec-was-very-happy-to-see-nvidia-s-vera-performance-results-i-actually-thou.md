---
id: collect-240926-tomshardware/tomshardware/amd-exec-was-very-happy-to-see-nvidia-s-vera-performance-results-i-actually-thou
title: "amd-exec-was-very-happy-to-see-nvidia-s-vera-performance-results-i-actually-thou"
domain: tomshardware
role: reference
task: reference
actors: ["AMD", "Google", "Nvidia", "TSMC"]
dates: []
keywords: ["amd", "nvidia", "benchmark", "benchmarks", "chatgpt", "claude", "compute", "copilot", "throughput"]
source: docs/RAG/clean_en/tomshardware/amd-exec-was-very-happy-to-see-nvidia-s-vera-performance-results-i-actually-thou.md
source_anchor: ""
source_lines: [1, 47]
sha256: 566ae3b67105a94b6cd8c8060cbc79428354a7f2ccc204f279de193c2ce806b2
---

# amd-exec-was-very-happy-to-see-nvidia-s-vera-performance-results-i-actually-thou

<!-- source: https://www.tomshardware.com/pc-components/cpus/amd-exec-was-very-happy-to-see-nvidias-vera-performance-results-i-actually-thought-we-were-beating-them-by-smaller-numbers -->

Although the industry has largely learned to move out of the way when the big green giant that is Nvidia steps into the room, one AMD executive said he was “very happy” to see Nvidia publish __SPEC CPU 2026 benchmarks for its Vera CPU__ ahead of the Advancing AI 2026 event. AMD used the configuration in Nvidia’s white paper as a basis to run SPEC on its new Zen 6 ‘Venice’ CPUs, offering what it calls an “apples-to-apples” comparison between the two chips.

“We are very happy that Nvidia published their Vera performance [numbers],” said AMD’s Ravi Kuppuswany, corporate VP of compute and enterprise solutions. “We were actually being a little conservative. I actually thought we were beating them by smaller numbers than what I have here… we thought we’d have at least a 10% advantage. What we’re finding is… we have 20% advantage, and we have not even finished completely tuning.”

Kuppuswany’s comments overlaid the slide you can see above, claiming 2.2x higher throughput with Venice compared to Vera, and 1.2x faster per-core performance. AMD has certainly stacked the deck in its favor here (as did Nvidia when it first published its Vera results), so let’s go through what this chart actually says.

As a quick aside, you’ll see “estimated” in the images above and below. These numbers aren’t guesses (they’re based on real runs), but SPEC maintains strict reporting guidelines for “official” runs. And because these runs aren’t reported to SPEC and therefore haven’t gained official status, they must come with the “estimated” disclaimer.

The throughput number is the easiest one to clarify. AMD compared its __256-core, 600W Epyc 9996__ against the 88-core Vera, both in a 2P configuration. Yes, the CPU with more than twice the number of threads and an extra 150W stacked on top of its TDP has significantly higher throughput, as it should.

AMD would argue that it’s a fair comparison given that Nvidia is only offering Vera as a single, 88-core SKU. But the reality remains that the throughput comparison is one Vera could never win, and Nvidia would (__and has__) argued that it’s not trying to win in a race against the 256-core Venice. It’s building a single CPU for a single purpose.

The more interesting and consequential number here is the per-core performance. SPECrate_int is a throughput benchmark. A typical run loads all threads with a copy of an application and measures how much work gets done within a set amount of time. SPECspeed is the inverse of that, looking at a single application and how quickly it can run. AMD arrived at the numbers above by taking the overall SPECrate_int score and dividing it by the number of cores.

Get Tom's Hardware's best news and in-depth reviews, straight to your inbox.

Nvidia’s results show Vera achieving a score of 925 overall. AMD says a 96-core High Frequency Epyc CPU achieved a score of 1,210. This processor, however, doesn’t seem to exist. AMD has the 96-core Epyc 9686F that boosts up to 5 GHz (much higher than the standard max frequency across the stack), but it has a rated TDP of 500W, not 600W.

Regardless, AMD divided that score by the number of cores; about 6.3 for AMD, and about 5.3 for Nvidia (remember these are 2P configurations). AMD says 1.2x, which actually translates to about an 18.8% lead. That’s not far off enough to say AMD was maliciously juicing its own numbers, but it’s important to note.

Given that AMD is using two different SKUs here (or maybe just one, considering the 9686F discrepancy), we can do the same per-core napkin math on the Epyc 9996 against Vera. Once again, Nvidia shared an overall score of 925, while the Epyc 9996 achieved a score of 2,070. That gives AMD a per-core score of 4.04.

That number isn’t important as a comparison point to Nvidia – again, we’re comparing a 256-core CPU to an 88-core one – but rather as a metric to see how Venice scales when normalized for per-core performance on a fully-loaded chip. That’s about 35% lower than the per-core score from the souped-up 9686F.

Although dissecting these numbers and sidelining the back-and-forth of two of the most powerful companies in the world is interesting, it’s not all that informative. SPEC maintains its strict reporting requirements for a reason. We won’t be able to say, with certainty, how these chips match up until we have official, reported runs. And even then, there’s an additional layer of compiler optimization (AMD and Nvidia both used GCC 15.2) and the broader context of the servers and workloads that these chips will serve.

The battle lines have been drawn, though only when looking at integer-based workloads. Vectorized performance is important, as well, and that’s an area where AMD holds a strong position in the current server CPU market. It doesn’t look like that will change with Vera, though we don’t have floating-point results to draw any conclusions from yet.

*Follow* *Tom's Hardware on Google News**, or* *add us as a preferred source**, to get our latest news, analysis, & reviews in your feeds.*

- 
the ai bubble will stop, when these companies will have no cash any more and will want to really start to make money, instead of making billions of expenses.Reply
 
 at this time, chatgpt, claude, copilot and co will be really costly and expensive to use, for normal user.
 
so, normal users/companies will drastically reduce or just stop using them.
- 
I don't understand how we're supposed to be happy and excited that AMD's unreleased 96-core 600w chip offers only 20% more performance while consuming 33% more power than Nvidia's also-unreleased 88-core 450w chip.Reply
 
 Does "we thought we'd be 10% ahead but we're actually 20%" translate into "we thought we were going to be terrible on power efficiency, but we're actually only slightly bad."?
 
Is AMD's executive really bad at math, or does he just think we are?
- 
Apples-to-apples? Not quite. AMD is using its 256-core SKU for the overall throughput claim and switching to a 96-core high-frequency SKU for the per-core claim.Reply
When both vendors rely on TSMC, pushing higher clocks forces you way up the exponential power curve. Disclosing the operating frequencies and actual Perf/Watt for these SPEC runs would give us a much clearer picture than selective SKU metrics.

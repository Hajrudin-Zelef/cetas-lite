---
id: collect-261001-ia-llm/ia-llm/designing-ai-resistant-technical-evaluations-2
title: "designing-ai-resistant-technical-evaluations"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["claude", "compute", "distribution", "exploit", "memory", "opus 4", "reasoning", "tpu", "training"]
source: docs/RAG/collect-261001-ia-llm/designing-ai-resistant-technical-evaluations.md
source_anchor: ""
source_lines: [63, 108]
sha256: eb39fab7d3da1ed817fd296f28566d051750bd0c544a8bb549f1eb6ee8251ddd
---

# designing-ai-resistant-technical-evaluations

This wasn't my first interview defeated by a Claude model. I'd designed a live interview question in 2023 specifically because our questions at the time were based around common tasks that early Claude models had lots of knowledge of and so could solve easily. I tried to design a question that required more problem solving skill than knowledge, still based on a real (but niche) problem I'd solved at work. Claude 3 Opus beat part 1 of that question; Claude 3.5 Sonnet beat part 2. We still use it because our other live questions aren't AI-resistant either.

For the take-home, there was a straightforward fix. The problem had far more depth than anyone could explore in 4 hours, so I used Claude Opus 4 to identify where it started struggling. That became the new starting point for version 2. I wrote cleaner starter code, added new machine features for more depth, and removed multicore (which Claude had already solved, and which only slowed down development loops without adding signal).

I also shortened the time limit from 4 hours to 2 hours. I'd originally chosen 4 hours based on candidate feedback preferring less risk of getting sunk if they got stuck for a bit on a bug or confusion, but the scheduling overhead was causing multi-week delays in our pipeline. Two hours is much easier to fit into a weekend.

Version 2 emphasized clever optimization insights over debugging and code volume. It served us well—for several months.

When I tested a pre-release Claude Opus 4.5 checkpoint, I watched Claude Code work on the problem for 2 hours, gradually improving its solution. It solved the initial bottlenecks, implemented all the common micro-optimizations, and met our passing threshold in under an hour.

Then it stopped, convinced it had hit an insurmountable memory bandwidth bottleneck. Most humans reach the same conclusion. But there are clever tricks that exploit the problem structure to work around that bottleneck. When I told Claude the cycle count it was possible to achieve, it thought for a while and found the trick. It then debugged, tuned, and implemented further optimizations. By the 2-hour mark, its score matched the best human performance within that time limit—and that human had made heavy use of Claude 4 with steering.

We tried it out in our internal test-time compute harness for more rigor and confirmed it could both beat humans in 2 hours and continue climbing with time. Post-launch we even improved our harness in a generic way and got a higher score.

I had a problem. We were about to release a model where the best strategy on our take-home would be delegating to Claude Code.

Some colleagues suggested banning AI assistance. I didn't want to do this. Beyond the enforcement challenges, I had a sense that given people continue to play a vital role in our work, I should be able to figure out *some* way for them to distinguish themselves in a setting *with AI—*like they'd have on the job. I didn't want to give in yet to the idea that humans only have an advantage on tasks longer than a few hours.

Others suggested raising the bar to "substantially outperform what Claude Code achieves alone." The concern here was that Claude works fast. Humans typically spend half the 2 hours reading and understanding the problem before they start optimizing. A human trying to steer Claude would likely be constantly behind, understanding what Claude did only after the fact. The dominant strategy might become sitting back and watching.

Nowadays performance engineers at Anthropic still have lots of work to do, but it looks more like tough debugging, systems design, performance analysis, figuring out how to verify the correctness of our systems, and figuring out how to make Claude's code simpler and more elegant. Unfortunately these things are tough to test in an objective way without a lot of time or common context. It's always been hard to design interviews that represent the job, but now it's harder than ever.

But I also worried if I invested in designing a new take-home, either Claude Opus 4.5 would solve that too, or it would become so challenging that it would be impossible for humans to complete in two hours.

I realized Claude could help me implement whatever I designed quickly, which motivated me to try developing a harder take-home. I chose a problem based on one of the trickier kernel optimizations I'd done at Anthropic: an efficient data transposition on 2D TPU registers while avoiding bank conflicts. I distilled it into a simpler problem on a simulated machine and had Claude implement the changes in under a day.

Claude Opus 4.5 found a great optimization I hadn't even thought of. Through careful analysis, it realized it could transpose the entire computation rather than figuring out how to transpose the data, and it rewrote the whole program accordingly.

In my real case, this wouldn't have worked, so I patched the problem to remove that approach. Claude then made progress but couldn't find the most efficient solution. It seemed like I had my new problem, now I just had to hope human candidates could get it fast enough. But I had some nagging doubt, so I double-checked using Claude Code's "ultrathink" feature with longer thinking budgets ... and it solved it. It even knew the tricks for fixing bank conflicts.

In hindsight, this wasn't the right problem to try. Engineers across many platforms have struggled with data transposition and bank conflicts, so Claude has substantial training data to draw on. While I'd found my solution from first principles, Claude could draw on a larger toolbox of experience.

I needed a problem where human reasoning could win over Claude's larger experience base: something sufficiently out of distribution. Unfortunately, this conflicted with my goal of being recognizably like the job.

I thought about the most unusual optimization problems I'd enjoyed and landed on Zachtronics games. These programming puzzle games use unusual, highly constrained instruction sets that force you to program in unconventional ways. For example, in Shenzhen I/O, programs are split across multiple communicating chips that each hold only about 10 instructions with one or two state registers. Clever optimization often involves encoding state into the instruction pointer or branch flags.

I designed a new take-home consisting of puzzles using a tiny, heavily constrained instruction set, optimizing solutions for minimal instruction count. I implemented one medium-hard puzzle and tested it on Claude Opus 4.5. It failed. I filled out more puzzles and had colleagues verify that people less steeped in the problem than me could still outperform Claude.

Unlike Zachtronics games, I intentionally provided no visualization or debugging tools. The starter code only checks whether solutions are valid. Building debugging tools is part of what's being tested: you can either insert well-crafted print statements or ask a coding model to generate an interactive debugger in a few minutes. Judgment about how to invest in tooling is part of the signal.

I'm reasonably happy with the new take-home. It might have lower variance than the original because it comprises more independent sub-problems. Early results are promising: scores correlate well with the caliber of candidates' past work, and one of my most capable colleagues scored higher than any candidate so far.

I'm still sad to have given up the realism and varied depth of the original. But realism may be a luxury we no longer have. The original worked because it resembled real work. The replacement works because it simulates novel work.

We're releasing the original take-home for anyone to try with unlimited time. Human experts retain an advantage over current models at sufficiently long time horizons. The fastest human solution ever submitted substantially exceeds what Claude has achieved even with extensive test-time compute.


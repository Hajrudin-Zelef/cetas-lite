---
id: collect-261001-ia-llm/ia-llm/designing-ai-resistant-technical-evaluations-1
title: "designing-ai-resistant-technical-evaluations"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Anthropic"]
dates: ["2023-11", "2025-05"]
keywords: ["accelerator", "claude", "distribution", "exploit", "gpu", "inference", "memory", "opus 4", "tpu", "trainium"]
source: docs/RAG/collect-261001-ia-llm/designing-ai-resistant-technical-evaluations.md
source_anchor: ""
source_lines: [1, 62]
sha256: 8add60519d1f661012dbbad77a2190f85e47186e0a4f58ca8858baa70ed0e3b7
---

# designing-ai-resistant-technical-evaluations

## Get the developer newsletter

Product updates, how-tos, community spotlights, and more. Delivered monthly to your inbox.

*Written by Tristan Hume, a lead on Anthropic's performance optimization team. Tristan designed—and redesigned—the take-home test that's helped Anthropic hire dozens of performance engineers.*

Evaluating technical candidates becomes harder as AI capabilities improve. A take-home that distinguishes well between human skill levels today may be trivially solved by models tomorrow—rendering it useless for evaluation.

Since early 2024, our performance engineering team has used a take-home test where candidates optimize code for a simulated accelerator. Over 1,000 candidates have completed it, and dozens now work here, including engineers who brought up our Trainium cluster and shipped every model since Claude 3 Opus.

But each new Claude model has forced us to redesign the test. When given the same time limit, Claude Opus 4 outperformed most human applicants. That still allowed us to distinguish the strongest candidates—but then Claude Opus 4.5 matched even those. Humans can still outperform models when given unlimited time, but under the constraints of the take-home test, we no longer had a way to distinguish between the output of our top candidates and our most capable model.

I've now iterated through three versions of our take-home in an attempt to ensure it still carries signal. Each time, I’ve learned something new about what makes evaluations robust to AI assistance and what doesn't.

This post describes the original take-home design, how each Claude model defeated it, and the increasingly unusual approaches I've had to take to ensure our test stays ahead of our top model’s capabilities. While the work we do has evolved alongside our models, we still need more strong engineers—just increasingly creative ways to find them.

To that end, we're releasing the original take-home as an open challenge, since with unlimited time the best human performance still exceeds what Claude can achieve. If you can best Opus 4.5, we’d love to hear from you—details are at the bottom of this post.

In November 2023, we were preparing to train and launch Claude Opus 3. We’d secured new TPU and GPU clusters, our large Trainium cluster was coming, and we were spending considerably more than we had in the past on accelerators, but we didn't have enough performance engineers for our new scale. I posted on Twitter asking people to email us, which brought in more promising candidates than we could evaluate through our standard interview pipeline, a process that consumes significant time for staff and candidates

We needed a way to evaluate candidates more efficiently. So, I took two weeks to design a take-home test that could adequately capture the demands of the role and identify the most capable applicants.

Take-homes have a bad reputation. Usually they’re filled with generic problems which engineers find boring, and which make for poor filters. My goal was different: create something genuinely engaging that would make candidates excited to participate and allow us to capture their technical skills at a high-level of resolution.

The format also offers advantages over live interviews for evaluating performance engineering skills:

**Longer time horizon:** Engineers rarely face deadlines of less than an hour when coding. A 4-hour window (later reduced to 2 hours) better reflects the actual nature of the job. It's still shorter than most real tasks, but we need to balance that with how onerous it is.

**Realistic environment:** No one watching or expecting narration. Candidates work in their own editor without distraction.

**Time for comprehension and tooling:** Performance optimization requires understanding existing systems and sometimes building debugging tools. Both are hard to realistically evaluate in a normal 50 minute interview.

**Compatibility with AI assistance:** Anthropic's general candidate guidance asks candidates to complete take-homes without AI unless indicated otherwise. For this take-home, we explicitly indicate otherwise.

Longer-horizon problems are harder for AI to solve completely, so candidates can use AI tools (as they would on the job) while still needing to demonstrate their own skills.

Beyond these format-specific goals, I applied the same principles I use when designing any interview to make the take-home:

**Representative of real work:** The problem should give candidates a taste of what the job actually involves.

**High signal:** The take-home should avoid problems that hinge on a single insight and ensure candidates have many chances to show their full abilities — leaving as little as possible to chance. It should also have a wide scoring distribution,and ensure enough depth that even strong candidates don't finish everything.

**No specific domain knowledge:** People with good fundamentals can learn specifics on the job. Requiring narrow expertise unnecessarily limits the candidate pool.

**Fun:** Fast development loops, interesting problems with depth, and room for creativity.

I built a Python simulator for a fake accelerator with characteristics that resemble TPUs. Candidates optimize code running on this machine, using a hot-reloading Perfetto trace that shows every instruction, similar to the tooling we have on Trainium.

The machine includes features that make accelerator optimization interesting: manually managed scratchpad memory (unlike CPUs, accelerators often require explicit memory management), VLIW (multiple execution units running in parallel each cycle, requiring efficient instruction packing), SIMD (vector operations on many elements per instruction), and multicore (distributing work across cores).

The task is a parallel tree traversal, deliberately not deep learning flavored, since most performance engineers hadn't worked on deep learning yet and could learn domain specifics on the job. The problem was inspired by branchless SIMD decision tree inference, a classical ML optimization challenge as a nod to the past, which only a few candidates had encountered before.

Candidates start with a fully serial implementation and progressively exploit the machine's parallelism. The warmup is multicore parallelism, then candidates choose whether to tackle SIMD vectorization or VLIW instruction packing. The original version also included a bug that candidates needed to debug first, exercising their ability to build tooling.

The initial take-home worked well. One person from the Twitter batch scored substantially higher than everyone else. He started in early February, two weeks after our first hires through the standard pipeline. The test proved predictive: He immediately began optimizing kernels and found a workaround for a launch-blocking compiler bug involving tensor indexing math overflowing 32 bits.

Over the next year and a half, about 1,000 candidates completed the take-home, and it helped us hire most of our current performance engineering team. It proved especially valuable for candidates with limited experience on paper: several of our highest-performing engineers came directly from undergrad but showed enough skill on the take-home for us to hire confidently.

Feedback was positive. Many candidates worked past the 4-hour limit because they were enjoying themselves. The strongest unlimited-time submissions included full optimizing mini-compilers and several clever optimizations I hadn't anticipated.

By May 2025, Claude 3.7 Sonnet had already crept up to the point where over 50% of candidates would have been better off delegating to Claude Code entirely. I then tested a pre-release version of Claude Opus 4 on the take-home. It came up with a more optimized solution than almost all humans did within the 4-hour limit.


---
id: vague2-nerdykings/nerdykings/claude-riemann-hypothese-650-echecs
title: "650 Échecs Plus Tard, Claude Fait Une Découverte Sur Riemann"
domain: nerdykings
role: reference
task: article
actors: ["Anthropic"]
dates: ["2026-09-23"]
keywords: ["claude", "agent", "agents", "agi", "distribution", "formalization", "lean", "reasoning", "research"]
source: docs/RAG/Collect RAG Vague 2/03_nerdykings/claude-riemann-hypothese-650-echecs.md
source_anchor: ""
source_lines: [1, 70]
sha256: 624b363f399d22a86a6385384e637e75a56cd20dfaa5adf92e6465f4f6e7fbec
---

# 650 Échecs Plus Tard, Claude Fait Une Découverte Sur Riemann

## Metadata

- **Source** : https://www.nerdykings.com/blog/claude-riemann-hypothese-650-echecs.html
- **Site** : NerdyKings
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Anthropic revealed a remarkable experiment: they took a research version of Claude, not yet publicly available, and asked it to seriously tackle the **Riemann hypothesis** — one of the most famous unsolved problems in mathematics, open for over 160 years, with a $1M prize. Spoiler: Claude did not solve it, but what happened during the experiment is more interesting than the result.

What is the Riemann hypothesis? Dating from 1859, it indirectly concerns the distribution of prime numbers (2, 3, 5, 7, 11...). Mathematicians study them using the **Riemann zeta function**, which has "zeros." The hypothesis states that all non-trivial zeros of this function lie exactly on a particular line — the **critical line**. Nobody has proven this for all zeros, but mathematicians have proven it for a certain proportion. That's where Claude intervenes.

Did Claude solve Riemann? No. Before this experiment, the best results guaranteed that about **41.6%** of zeros lie on the critical line. Claude's new method raises this proportion to about **67.2%** — a huge progression. But don't over-interpret: this doesn't mean Claude "solved 67%" of the hypothesis. Even proving 99.999% of zeros lie on the line wouldn't prove the hypothesis — a single counterexample could falsify it. Anthropic itself acknowledges this new method doesn't seem able to lead directly to a complete proof. Concretely, Claude found a novel way to combine recent mathematical results (work by Aryan, Baluyot, Goldston, Suriajaya, and Turnage-Butterbaugh) with an older result by Bombieri from 2000 — treating the entire space with positivity and negativity defined together, in a quadratic-form formulation.

650 failed ideas, then a real little research lab. Claude didn't find the answer on the first try. In a first phase, the system explored about **650 different ideas**, none of which worked. Anthropic then let Claude search much longer with much more resources: two sessions, about **31 million output tokens**, and notably **~60 sub-agents working in parallel for about a day and a half**. Concretely: agents developing the main mathematical leads; others exploring alternative approaches; a final group dedicated to verifying results and hunting for errors. Far from writing a perfect prompt and getting the answer immediately, it looked like a real automated research lab with role distribution among agents — similar to emerging AI agent-team tools where agents specialize and critique each other, applied here to fundamental research.

The strange moment: 37 minutes of silence. One sub-agent was exploring an approach based on an existing mathematical idea but concluded it couldn't work as planned. Instead of stopping, it flipped the problem and tried viewing the same structure differently. Then something strange happened in the logs: for about **37 minutes** the sub-agent stayed nearly silent. Then a new idea appeared at once, yielding an important result around 50%. Even Claude immediately found this result suspicious, writing literally: **"too strong to be new"** — the result seemed so strong it must either already exist or contain an error. The main orchestrator shared the skepticism and launched several other agents to actively dismantle the reasoning. They found a few things to correct but no error that could eliminate the main result. A new research phase then reinforced the method up to the ~67% result.

How do we know Claude didn't hallucinate a proof? This is the most important question. Anthropic set up several verification layers: other agents reproduced the reasoning and searched for counterexamples; Claude downloaded **54 scientific papers** to check whether the idea was truly new; human mathematicians (**Levent Alpöge** and **Ralph Furman** at Anthropic, then external experts **Brian Conrey** and **Dan Goldston**) examined the result and performed an independent re-derivation; and the proof was formalized with **Lean**, a proof assistant that rigorously verifies each step. The result is very recent and must still be studied in depth by the mathematical community, but at this stage no fatal error has been publicly identified.

The funniest detail: the person who launched the experiment, **Jared Somner**, isn't even a mathematician. Anthropic explains that for much of the process, his interventions were mainly messages like "keep going" or "believe in yourself." Claude doesn't need emotional support per se, but these messages mainly encouraged it to continue rather than give up too quickly after 650 consecutive failures.

The author's view: Claude didn't solve the Riemann hypothesis, and this method may never get there. It's not AGI or a spontaneous "spark of intelligence" — it's the result of 60 agents, 31 million tokens, and a day and a half of raw computation on an extremely bounded, formally verifiable problem, similar to the internal geometry Claude uses to count characters: an impressive jury-rigged mechanism, not human-like understanding. But what strikes the author is the trajectory: until now LLMs were essentially good at understanding and summarizing discoveries already made by humans. Here a system explores, fails 650 times, doubts its own result, gets verified by dozens of other instances of itself, and ends up producing something recognized mathematicians validate — a real change in nature, even if the problem was tailor-made (formally verifiable, bounded, with a huge literature to digest). The real question: does this work only because Riemann is a perfect mathematical playground for brute-force-plus-verification, or have we just seen the prototype of a real autonomous research lab applicable to other fields?

## Key points

- Anthropic used an unreleased research version of Claude on the **Riemann hypothesis** (open since 1859, $1M prize).
- Claude did not solve it; it raised the proven proportion of zeros on the critical line from ~**41.6%** to ~**67.2%**.
- This does not constitute a proof — a single counterexample could falsify the hypothesis.
- Method combined recent results (Aryan, Baluyot, Goldston, Suriajaya, Turnage-Butterbaugh) with Bombieri's 2000 result via a quadratic-form formulation.
- The system explored ~**650 failed ideas** before a working approach emerged.
- Scale: 2 sessions, ~**31M output tokens**, ~**60 parallel sub-agents** over ~1.5 days.
- A sub-agent went silent for ~**37 minutes**, then produced a sudden result (~50%); Claude itself called it "too strong to be new."
- Verification: agent reproduction/counterexamples, 54 papers checked, human mathematicians, and **Lean** formalization.
- Experiment launched by **Jared Somner**, a non-mathematician, whose main input was "keep going."
- Author frames it as an automated research lab prototype, not AGI.

## Technical data / figures

| Item | Value |
|---|---|
| Problem | Riemann hypothesis (1859) |
| Prior proven proportion | ~41.6% |
| New proportion | ~67.2% (67.25%) |
| Failed ideas explored | ~650 |
| Output tokens | ~31 million |
| Sub-agents | ~60 parallel |
| Duration | ~1.5 days (two sessions) |
| Silent period before breakthrough | ~37 minutes |
| Intermediate result | ~50% |
| Papers checked for novelty | 54 |
| Human verifiers | Levent Alpöge, Ralph Furman, Brian Conrey, Dan Goldston |
| Formalization | Lean proof assistant |
| Experiment lead | Jared Somner |

- Key concepts: **Riemann zeta function**, **non-trivial zeros**, **critical line**, **quadratic form**, **Lean**
- Reference results: Aryan, Baluyot, Goldston, Suriajaya, Turnage-Butterbaugh; Bombieri (2000)

## Why this source matters for the RAG

This article provides a detailed, quantified account of an AI multi-agent system attacking a major open mathematical problem, including the exact verification pipeline (Lean, human experts) and honest limitations. It is essential for a RAG knowledge base on AI-for-mathematics, autonomous research agents, and formal verification.

## Source URL

https://www.nerdykings.com/blog/claude-riemann-hypothese-650-echecs.html

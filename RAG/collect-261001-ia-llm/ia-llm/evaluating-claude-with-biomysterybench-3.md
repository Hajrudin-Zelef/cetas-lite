---
id: collect-261001-ia-llm/ia-llm/evaluating-claude-with-biomysterybench-3
title: "evaluating-claude-with-biomysterybench"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "United States", "Z.ai"]
dates: []
keywords: ["claude", "benchmark", "benchmarks", "cost", "cyber", "glm", "opus 4", "reasoning", "research", "safeguards", "tool use"]
source: docs/RAG/collect-261001-ia-llm/evaluating-claude-with-biomysterybench.md
source_anchor: ""
source_lines: [73, 92]
sha256: 73cdd2d1568383cd1333a345919135848a42f7f66c1ec5f1c3f0ee93e9dc8906
---

# evaluating-claude-with-biomysterybench

`The headline accuracy numbers tell you how often each model gets the right answer, but not how it gets there. I wanted to know whether a correct answer on a hard problem means the same thing as a correct answer on a solvable one. Since every problem was attempted five times, I could look at per-problem solve counts: if a model solves something 5/5 it has a reliable method; if it solves it 1/5 it probably got lucky on a reasoning path it can't consistently find again. So I broke each model's solved problems down by solve count (0/5 through 5/5) on the two sets side by side.``The texture of "solved" changes sharply between the two sets. On human-solvable problems, Opus 4.6 is strongly bimodal — 86% of the problems it solves at all, it solves at least 4 out of 5 times. It either has the answer or it doesn't. On the human-difficult set that collapses to 44%, and the share of brittle wins (solved only 1–2 of 5 attempts) jumps from 9% to 44%. Sonnet 4.6 shows the same shift, and more sharply (75% reliable → 22%; 9% brittle → 56%). So the 77.4%→23.5% headline drop actually understates what's happening: on solvable problems the model is retrieving something it reliably knows, while on hard problems nearly half of its wins are paths it stumbles onto rather than reproduces. The accuracy gap is real, but the reliability gap underneath it is the more interesting story about where the capability frontier actually sits. Opus 4.7 and Mythos move the frontier a little (Mythos gets 94% of its solvable wins at ≥4/5) but the same bimodal-vs-brittle split holds on the difficult set for every model.`
We thought Claude Mythos Preview’s analysis held up and dove deeper into reliability, which is an important metric to measure model performance on. However, it also felt a little…boring? It added some nuance to the performance analysis we showed above, but did not fundamentally tackle a new question. Despite this, it seems like the models are starting to develop the seeds of research taste (even if they have a ways to go before producing deep insight).

BioMysteryBench is an encouraging measure of scientific capability. The most recent generations of Claude solve the majority of human-solvable problems reliably, and on a meaningful fraction of human-difficult tasks, it outperforms panels of five domain experts. Models are improving across generations, and are no longer merely keeping up with trained scientists on bioinformatics problems; on some tasks, they’re ahead.

We’re also delighted to see convergent work in this space: While finalizing this post, Genentech and Roche released CompBioBench. Their benchmark consists of 100 computational biology tasks “based on synthetic/augmented data and metadata scrambling/scrubbing of real datasets to create challenging problems with a single ground-truth answer that require multi-step reasoning, tool use, bespoke code, and interaction with real-world external resources.” Sound familiar? Their results echo those of BioMysteryBench, too: Claude Opus 4.6 reaches 81% overall and 69% on their hardest questions, reinforcing that frontier models are now genuinely useful collaborators for bioinformatics research.

We’re eager to build even longer-horizon, real-world tasks that push model research capabilities, and to hear creative ideas from others. Send us your interesting benchmarks, innovative uses of AI for science, and interactions with AI that prompted you to rethink what could be possible in your field at scienceblog@anthropic.com.

If you are interested in understanding how models perform on difficult verifiable computational biology tasks, you can access BioMysteryBench here and visit claude.com/lifesciences to learn more.

We built an index of how well today’s robots can perform US job tasks. Robots can already do three-quarters of physical tasks, mostly in limited settings, but are cost-competitive for just 0.3% of them.

Read more
We’re launching a new study using Anthropic Interviewer to learn from your experiences with AI, and we invite you to participate.

Read more
Like Claude Mythos Preview, GLM-5.3 has strong capabilities for autonomously building end-to-end cyber exploits. But GLM-5.3 is unlike other frontier models in that it has been released without meaningful safeguards to limit misuse.

Read more

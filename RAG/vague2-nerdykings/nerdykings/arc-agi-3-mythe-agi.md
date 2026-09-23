---
id: vague2-nerdykings/nerdykings/arc-agi-3-mythe-agi
title: "ARC-AGI-3 : Le Benchmark Qui Casse Le Mythe De L'AGI"
domain: nerdykings
role: reference
task: article
actors: ["Google"]
dates: ["2026-09-23"]
keywords: ["agi", "benchmark", "agent", "agentic", "agents", "benchmarks", "gemini", "memory", "training"]
source: docs/RAG/Collect RAG Vague 2/03_nerdykings/arc-agi-3-mythe-agi.md
source_anchor: ""
source_lines: [1, 43]
sha256: 6651d68560ac005c33ce4d0a8ac7eb8a9e748a9d8c8717ce1e5178b01d139b88
---

# ARC-AGI-3 : Le Benchmark Qui Casse Le Mythe De L'AGI

## Metadata

- **Source** : https://www.nerdykings.com/blog/arc-agi-3-mythe-agi.html
- **Site** : NerdyKings
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

The article examines ARC-AGI-3, a benchmark that challenges the narrative that AI is about to become general. While models pass law exams and solve doctorate-level physics, ARC-AGI-3 delivers a brutal result: untrained humans succeed on 100% of tasks, while the most advanced models score under 1%. The article first explains why current benchmarks are biased. Most are static tests (input → output): math, code, multiple choice, text analysis. These tests almost always end up integrated directly or indirectly into training data, so models become good not because they understand but because they recognize previously seen patterns — "crystallized intelligence." Real intelligence is the ability to face something totally new, without instructions or examples, and understand what to do.

ARC-AGI-3's ingenious idea: instead of giving a problem to solve, it gives an environment in which to evolve. You land in a kind of mini-game, you don't know the rules, the goal, or any instructions. You must interact and progressively understand what is happening — like being dropped into a video game you've never seen, with no tutorial. You test, observe, form hypotheses, correct errors, and gradually build understanding. The agent must do two things simultaneously: understand what to do and understand how to do it — discovering goal and strategy in parallel, purely from interactions. That is what current models cannot do.

Why humans excel: we have cognitive priors — basic mental structures developed from childhood. We intuitively understand what an object is, how elements interact, how space, symmetries, and repetitions work, letting us make sense of a new situation very quickly. AI models don't have this; they can simulate these concepts in static contexts but fail completely in dynamic interactive environments. The benchmark also measures the "how": not just whether you solve the problem but how many actions you take, how you explore, how fast you understand. A human tests a few actions, understands the system, and goes straight to the optimal solution; an AI often tests enormous numbers of possibilities, loops, repeats useless actions, and advances without real understanding. ARC-AGI-3 heavily penalizes brute force. The fundamental limitation of current models: GPT, Gemini, and company are incredible in well-defined contexts but are not made to act in an environment, not designed to explore, have no persistent experience memory, don't build a world model through interactions, and can't learn in real time from errors. The author concludes the future likely won't come from more data or bigger models, but from systems capable of acting, exploring, and learning autonomously — a shift from chatbot logic to agent logic.

## Key points

- ARC-AGI-3: humans 100% vs advanced AI models under 1%.
- Current static benchmarks suffer from data contamination → crystallized intelligence.
- ARC-AGI-3 uses interactive environments with unknown rules, goals, and no instructions.
- Agents must discover both the goal and the strategy from interactions alone.
- Humans rely on cognitive priors; models lack persistent world models and real-time learning.
- The benchmark measures exploration efficiency, penalizing brute force.
- Conclusion: the future is agentic systems that act, explore, and learn autonomously.

## Technical data / figures

| Item | Value |
|---|---|
| Human success rate | 100% |
| Advanced AI models | Under 1% |
| Benchmark type | Interactive environment (not static) |
| Instructions given | None |
| Measured dimensions | Success, action count, exploration efficiency |
| Prior paradigm | Static input → output tests |

## Why this source matters for the RAG

This source provides a sharp critique of AGI claims and explains the gap between pattern recognition and adaptable intelligence. It is valuable for RAG corpora on AGI evaluation, world models, and the shift from chatbots to autonomous agents.

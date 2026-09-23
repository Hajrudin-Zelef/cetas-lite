---
id: vague2-nerdykings/nerdykings/loop-transformers-claude-raisonnement-tokens
title: "Loop Transformers : La Fin Du Raisonnement En Tokens ?"
domain: nerdykings
role: reference
task: article
actors: ["Anthropic", "DeepSeek", "Google", "United States"]
dates: ["2026-04", "2026-09-23"]
keywords: ["attention", "benchmarks", "claude", "compute", "deepseek", "gemini", "inference", "reasoning", "research", "safeguards", "training"]
source: docs/RAG/Collect RAG Vague 2/03_nerdykings/loop-transformers-claude-raisonnement-tokens.md
source_anchor: ""
source_lines: [1, 43]
sha256: b60319af2e2c26759b5aee53f71329fda32dece7ea594258b65507e9df6759fb
---

# Loop Transformers : La Fin Du Raisonnement En Tokens ?

## Metadata

- **Source** : https://www.nerdykings.com/blog/loop-transformers-claude-raisonnement-tokens.html
- **Site** : NerdyKings
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

The article explores loop transformers, a research concept that gained popularity in April 2026, as an alternative to chain-of-thought reasoning. Today, all reasoning models (o3, Gemini Thinking, DeepSeek R2) generate large amounts of text to "think out loud." The author argues this is somewhat hacky: for a multi-hop question like "who is the wife of the 44th US president?", the model must chain two dependent steps, and chain-of-thought solves this by writing intermediate steps. But under the hood, each reasoning word must be converted to text, appended, then reconverted to internal representation — constant compression/decompression just to keep the thread of thought.

Loop transformers take a small block of layers and loop it on itself: the output of one loop becomes the input of the next, letting the model refine its internal state without ever exiting to text. Some researchers suspect this mechanism could explain why Claude performs particularly well on pure logic benchmarks (graph traversal, BFS) — tasks requiring step-by-step path following. The author stresses this is a researcher theory based on observed behavior; Anthropic has never officially confirmed it.

An April 2026 paper, "Loop, Think and Generalize," tested this setup and observed three distinct training stages: memorization (performance plateaus on training data), generalization (works only on similar problems), and systematic generalization (combining knowledge in ways never seen during training — a reusable reasoning procedure). Crucially, increasing the number of loops at inference time, even beyond what was seen during training, lets the model solve more complex problems — the loop count becomes a literal "thinking effort" slider. Looping creates instability (like audio feedback), addressed by a paper called "Parsy" using mathematical safeguards.

A mechanistic analysis (also April 2026) tracked the internal state per loop using dimensionality-reduction visualization, finding stable trajectories. Attention analysis revealed three distinct roles per loop: early loops roughly understand the problem, middle loops connect information, final loops stabilize and converge. Another paper, "Mixture of Recursions," proposes a router deciding how many loops each word deserves to save compute. The author concludes loop transformers won't replace chain-of-thought soon because CoT is text and thus supervisable/trainable, while hidden-space reasoning is harder to control — but loop transformers are exciting for small on-device models.

## Key points

- Chain-of-thought forces reasoning through text, causing constant compression/decompression overhead.
- Loop transformers recycle a small block of layers, refining internal state without text.
- Some researchers theorize this underlies Claude's strength on pure logic tasks (unconfirmed by Anthropic).
- "Loop, Think and Generalize" (April 2026) identifies three stages: memorization, generalization, systematic generalization.
- More inference-time loops act as a "thinking effort" slider, solving harder problems.
- "Parsy" paper adds mathematical safeguards against loop instability.
- Mechanistic analysis found three loop roles: understand, connect, stabilize/converge.
- CoT retains an advantage because it is text-based and therefore supervisable and trainable.

## Technical data / figures

- Concept popularity spike: April 2026
- Key papers: "Loop, Think and Generalize" (April 2026), "Parsy", "Mixture of Recursions"
- Training stages observed: 3 (memorization → generalization → systematic generalization)
- Loop functional roles: 3 (understanding → connecting → stabilizing)
- Target use case: small local/on-device models

## Why this source matters for the RAG

This source documents an emerging architecture (loop transformers) that challenges the dominant chain-of-thought paradigm, with concrete paper references and a mechanistic interpretability angle. It is valuable for RAG corpora on reasoning architectures, latent reasoning, and the future of small models.

---
id: ai-industry-kb-2026/05-chinese-ai-labs-deep-file-deepseek/what-did-not-happen-dated-negatives-as-important-as-the-posi
title: "What did not happen (dated negatives — as important as the positives)"
domain: chinese-ai-labs-deep-file-deepseek
role: deep-dive
task: actor-profile
actors: ["CISA", "DeepSeek"]
dates: ["2026-02-11", "2026-02-15", "2026-03-09", "2026-04-24", "2026-06-01", "2026-07-24", "2026-07-31", "2026-08-12", "2026-08-13", "2026-08-21", "2026-09-10", "2026-09-22"]
keywords: ["apache", "attention", "benchmarks", "deepseek", "license", "pricing", "reasoning", "training"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [2079, 2109]
section: "5. Chinese AI Labs — Deep File: DeepSeek"
sha256: c32d08b016210f5b0461af978ef51616329c1023a689612ff075deb321a9fe49
---

# What did not happen (dated negatives — as important as the positives)

- 2026-02-11 → 2026-03-09 (26 days): silent 1M upgrade to V4-Lite surfacing.
- 2026-03-09 → 2026-04-24 (46 days): leak phase to official V4 Preview.
- 2026-04-24 → 2026-06-01 (38 days): preview to permanent price cut.
- 2026-04-24 → 2026-07-31 (98 days): preview to Flash-0731 checkpoint.
- 2026-07-31 → 2026-08-13 (13 days): Flash-0731 to Pro-0813 GA.
- 2026-08-13 → 2026-08-21 (8 days): Pro-0813 GA to Flash-Vision-Exp preview.
- 2026-08-21 → 2026-09-10 (20 days): Vision-Exp to V4.1-Flash GA.
- 2026-04-24 → 2026-09-10 (139 days): official preview to new architecture family — the fastest structural update in DeepSeek's history (cellcog).
- The interval compression itself is the signal: checkpoints move from ~98 days apart to 8–20 days apart as training/post-training iteration velocity accelerates.

### What did not happen (dated negatives — as important as the positives)

- 2026-02-15: no model, no weights, no license, no pricing — rumor phase only.
- 2026-03-09: no GA release — V4-Lite is a leak/soft-preview under NDA framing.
- 2026-06-01: no V4.1 announcement — the event is a permanent 75% price cut.
- 2026-07-24: no engine change — the deepseek-chat/reasoner retirement is alias naming hygiene; V4-Flash answers before and after.
- 2026-08-13: no V4.1 — V4-Pro-0813 is a V4-line checkpoint with no technical blog post.
- 2026-09-10: no "Compressed Expert Dispatch" anywhere in DeepSeek material — the string is a rumor-cycle fabrication; CED is Causal Encoder-Decoder.
- 2026-08-12: no +10.6 on the Vals Index — Vals.ai records +9.48; no source for 10.6 was found.
- 2026-04-24: no Apache 2.0 — the shipped V4 license is MIT; March Apache-2.0 expectations are rumor residue.
- 2026-04-24: no "Engram"/"O(1)" in the official 58-page V4 report — zero mentions (buzzgrewal).
- 2026-09-22 (cutoff): no V4.1-Pro shipped yet; no independent reproduction of the 890 B/token claim; no independent reproduction of V4.1-Flash launch benchmarks; the Sept-14 routing conflict unresolved.
- 2026-08-12: no temperature parameter on V4-Pro-0813 — evaluations ran at max reasoning effort, a constraint on interpreting its scores.

### Cross-phase context

- The official V4 April technical report (58 pages) uses "Engram" and "O(1)" **zero times**; what V4 actually ships is a disciplined hybrid attention design — CSA (KV ~4×, top-k over compressed blocks, NSA/DSA lineage), HCA (KV ~128×, dense attention over compressed blocks) — per buzzgrewal's analysis. The "Compressed Expert Dispatch" string appears nowhere in DeepSeek material or credible secondary coverage (quoted-phrase search: zero relevant hits).
- Wave 1's V4-family benchmarks/pricing and wave3/02's V4 Preview verification are the priors this file extends — nothing is duplicated here; the additions are V4-Lite (Mar 9), the alias retirement (Jul 24), the June-1 price cut, the 0813 checkpoint detail, the beta-to-GA mechanics (Sept 8→10), and the rumor-phase documentation.

## Implications


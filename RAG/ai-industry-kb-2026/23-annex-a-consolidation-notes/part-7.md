---
id: ai-industry-kb-2026/23-annex-a-consolidation-notes/part-7
title: "Annex A — Consolidation notes (part 7)"
domain: appendix
role: appendix
task: reference
actors: ["Alibaba", "Anthropic", "CoreWeave", "Crusoe", "DeepSeek", "Google", "Meta", "MiniMax", "Moonshot", "Nebius", "Nvidia", "OpenAI", "Unsloth", "Z.ai", "vLLM"]
dates: ["2025-03", "2025-09-30", "2025-12", "2026-02", "2026-03", "2026-03-16", "2026-03-21", "2026-05-11", "2026-06", "2026-06-24", "2026-07-09", "2026-07-28", "2026-08-16", "2026-08-19", "2026-08-27", "2026-09-02", "2026-09-03"]
keywords: ["apache", "attention", "attribution", "benchmark", "compute", "deepseek", "diffusion", "disclosure", "foundry", "fp4", "fp8", "gemini"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [11215, 11261]
section: "Annex A — Consolidation notes"
sha256: 125dc42f909473e1014183628c9971b6cf5ad2b3e3b17cffced419f6e9560252
---

# Annex A — Consolidation notes (part 7)

1. **Wave 2's July 30 GPT-5.6 preview date** — superseded by the government-reviewed preview (June 25–26) and GA (July 9, 2026) — §1.
2. **Wave 2.1 §7's any-to-any Omni Flash framing** — nuanced: audio-reference inputs and scene extension were unsupported at preview; synchronized audio is the roadmap, not shipped — §1.
3. **The entire "April 11 wave" brief from the Wave 3.1 draft** — contradicted in full (ten-element debunk table) — §1.
4. **The February brief's "end of February" Qwen3.5 dating** — mid-February in the open section; not carried at all — §1, §2.
5. **Unsloth Dynamic Quantization 2.0 as "current"** — superseded by **Dynamic v3.0 (2026-08-19)**. v2.0's accuracy data (+0.12 vs +0.18 PPL delta, GLM-5.2 showcase) is retained only as the prior generation's figures, explicitly labeled — §2, §8, §10.
6. **Wave 2's "vLLM-Omni v0.28.0"** — retired; v0.28.0 is core vLLM — §6.
7. **"PagedAttention as vLLM's current core mechanism"** — legacy path since v0.25.0; mark as historical in inherited index/glossary entries — §6, §7.
8. **"Best open-weight model" as a durable title** — dated sequence, not a contradiction: GLM-5.2 (AA v4.1 51 at release) → Kimi K3 (AA 57, July 30) → GLM-5.3 (AA 60 v4.1-scale / ≈45 v4.3-scale). Any "best" statement must carry its benchmark suite and date; v4.1-era scores are not comparable to v4.3-era scores (rebase Sept 4, 2026) — §2, §3.
9. **Brief "GLM-5.2 is 744B"** → 753B total / ~40B active (744B-A40B footnoted as FP8-build shorthand) — §3.
10. **Brief "Kimi K2.6 close to Opus 4.6"** → ahead on SWE-Bench Pro (58.6% vs 53.4%) — §3.
11. **Brief listing "MiniMax H3" and "Hailuo 3.0" as separate items** → single merged entry — §3.
12. **Brief implication that Seed 2.1 Turbo's August tracker dates are its release** → 2026-06-24 FORCE pin with single-source caveat — §3.
13. **"Beats Opus 4.8" (Kimi K3) as blanket verdict** → benchmark-dependent only — §3.
14. **Wave 1's "2026 = the year MoE becomes the default architecture"** → restated with dense counterexamples — §9.
15. **Wave 1's "routing decisions finer and compute more efficient per flop … without hurting inter-GPU bandwidth"** — the strong "bandwidth-free" formulation marked [UNVERIFIED / contradicted as stated] per MoE Parallel Folding; Kimi K3 is the co-designed counterexample, not a refutation — §9.
16. **Wave 2 §7.1 MCP adoption claims** — the "78% enterprise" figure RETRACTED (replaced by Stacklok 41/30/29); "110M monthly downloads" re-dated to June 2026; the ~500M/month aggregator claim excluded; the "TCP/IP moment" quotation dropped; config portability downgraded from roadmap item to roadmap-aspiration — §13.
17. **Wave 2.1 §4.1 MCP spec-status RC caveat** — OBSOLETE; 2026-07-28 is Final; replaced by the adoption-lag caveat — §13.
18. **"No signed contract" caution (The IT Guys, 2026-08-27) on the NVIDIA–HF deal** — superseded by the Form 8-K definitive agreement (2026-09-02) and confirmation (2026-09-03); preserved only as a dated note — §16.
19. **Any brief framing presenting TGI or Candle as central HF orchestration** — superseded: TGI archived 2026-03-21 (maintenance December 2025), Candle active-but-niche; Inference Providers verified as the live routing layer — §16.
20. **Rumor-phase "Apache 2.0" for the DeepSeek V4 family** — residue; the family shipped **MIT** (V4, V4.1-Flash) — §5.
21. **"Compressed Expert Dispatch"** — deleted everywhere; CED = "Causal Encoder-Decoder" only — §5.
22. **"Engram as established DeepSeek technology"** — rumor-cycle material only (entered discourse March 2026, 36kr, pre-release V4 rumors); the official V4 April report never mentions it; the V4.1-Flash 196B-module attribution carries [VENDOR] — §5.
23. **Unsloth diffusion "2x"** — vendor-issued downgrade to 1.2–1.7x; use the later v0.1.808-beta figure — §9, §10.
24. **DeepSeek V4 pre-2026-08-16 promotional pricing ($0.07/$0.28)** — now the off-peak Flash rate; post-2026-08-16 figures are peak/off-peak — §14.
25. **Wave 2 [UNVERIFIED] flag on Lium** — retired: VERIFIED (datura-ai/lium, Bittensor SN51, active Sep 2026); MI355X "contact sales only" superseded except Crusoe's page (Vultr ~Aug 8, TensorWave $2.95, DO spot $4.50 — VERIFIED bookable) — §14.
26. **"CoreWeave–Meta $21B = largest AI cloud contract in history"** — superlative removed; parallel Nebius–Meta up-to-$27B (2026-03-16) prevents the claim — §14.
27. **"Sora launched 2026"** — 2025-09-30 (VentureBeat) — §12, §19.
28. **The February 2026 "all three live" timeline row (Sora + Kling + Seedance)** — an interpretive synthesis, not a reported event; kept labeled as context if reused — §12.
29. **September 11–12, 2026 Kling communications as a launch** — "AI Director" repositioning of the February model; no second launch, no "3.1" version downstream — §12.
30. **The humai.blog "Sora→Spud resource redirection" claim** — single-source [UNVERIFIED]; must not be merged with the confirmed Simo/Altman compute-reallocation reporting — §12.
31. **"TurboQuant invented by Red Hat"** — Google Research invention; Red Hat = independent evaluation (2026-05-11) — §7.
32. **"MLA on all engines"** — 3/4 (TGI archived 2026-03-21) — §7.
33. **"DSA reduces KV cache"** — attention compute; KV-cache reduction needs CSA/HCA-class compression — §7.
34. **"FP4/NVFP4 in production"** — experimental / SM100-native only; FP8 E4M3 is the 2026 production dtype — §7, §8.
35. **"O(N²)→constant KV memory"** — O(n²) = compute, KV memory = O(n), recurrent state = O(1) — §7.
36. **"Dynamic drift correction" as a method name** — [UNVERIFIED]; map to documented mechanisms — §7.
37. **Gemini Omni Flash "video + synchronized audio in one pass"** — preview nuance: audio-reference inputs and scene extension unsupported at preview; synchronized audio is the roadmap — §1.
38. **Wave 2.1 "vLLM-Omni" naming (Wave 2.1's own correction to Wave 2)** — carries over: v0.26.0 is core vLLM — §6, §7.
39. **Wave-2 730K+ HF datasets figure (via Wave 1)** — reconciled with the September 3 press 500K figure as methodological (public vs total repos), not averaged — §16.
40. **Wave-4 brief's OpenAI–HF breach dating ("Jul 16–21" compression)** — corrected to attack Jul 11–13 / disclosure Jul 16 / joint attribution Jul 21; the "answer grid stolen" phrasing and "nuclear war test" label likewise resolved (Scoop 2 = Payne's Project Kahn preprint, not a government report) — §17.
41. **NVFP4 ≠ MXFP4 conflation** — block-16 vs block-32, FP8-E4M3+FP32 scales vs E8M0, proprietary vs open spec; NVFP4 is an inference format, not a training default; MXFP4 is Hopper+ only (gpt-oss-120b does not run on A100) — §8.
42. **Marlin "SM80 floor"** — SM75 (Turing) — §8.
43. **NVL144 naming** — mapped to NVL72 (dies vs packages) — §15.
44. **Rubin "accelerated"/"pulled-in"** — on schedule within H2 2026 — §15.
45. **CES 2026 as the Rubin announcement** — production confirmation; announced GTC 2025 (March 2025) — §15.
46. **"1M is the new norm" (context)** — must carry the tier qualifier: **"flagship-tier norm"** (Qwen3.6 at 256K, Foundry 200K caps on Opus 4.8, Gemini 3.5 [UNVERIFIED], vendor-run NIH-at-1M) — §7.
47. **Independent 1M-context NIH validation gap** — carried as a field-wide gap, not a contradiction: most NIH-at-1M numbers are vendor-run (M3 9×/15×, V4 CSA/HCA figures); the Hopper FP8 bug (91%→13% NIAH) is the concrete reason vendor [VENDOR] labels and the long-context validation gate stay as a recurring theme — §7.

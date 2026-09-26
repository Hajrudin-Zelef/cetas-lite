---
id: ai-industry-kb-2026/23-annex-a-consolidation-notes/part-2
title: "Annex A — Consolidation notes (part 2)"
domain: appendix
role: appendix
task: reference
actors: ["AMD", "Alibaba", "Anthropic", "CoreWeave", "DeepSeek", "Google", "Groq", "Hugging Face", "Meta", "Mistral", "Moonshot", "Nebius", "Nvidia", "OpenAI", "OpenRouter", "SGLang", "SpaceX", "TensorRT-LLM", "Z.ai", "vLLM", "xAI"]
dates: ["2025-03", "2025-04", "2025-04-28", "2025-04-29", "2025-06", "2025-12", "2026-02", "2026-02-10", "2026-03", "2026-03-16", "2026-03-21", "2026-03-25", "2026-04-11", "2026-04-20", "2026-04-23", "2026-04-28", "2026-05-11", "2026-06-16", "2026-06-17", "2026-07", "2026-07-09", "2026-07-30", "2026-08", "2026-08-26", "2026-08-27", "2026-09-02", "2026-09-03", "2026-09-22"]
keywords: ["sol", "acquisition", "alignment", "attention", "attribution", "benchmark", "blackwell", "compute", "deepseek", "fp4", "fp8", "glm"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [11009, 11035]
section: "Annex A — Consolidation notes"
sha256: b73f0ad989fd69adb417c7c46d72b8d41387c0c1ac43150f514caabd19409e24
---

# Annex A — Consolidation notes (part 2)

1. **GLM-5.2 total parameters (744B vs 753B vs ~743B)** — §2, §3. Z.ai's "744B" is the FP8-build/VRAM shorthand; NVIDIA card shows 753B; vLLM config derives 743.4B backbone. Resolved as **three consistent numbers under different counting conventions**: 743.4B backbone (config-derived) + 9.95B MTP block = 753.3B full-weight total, 39.3B active. Consolidated headline: **753B total / ~40B active** with the "744B-A40B" FP8-shorthand footnote. Never "744B" bare, never "743B" as the headline total.
2. **Qwen3-30B-A3B "2026-02-10 release" vs April 2025 SKU** — §1, §2. The 2026-02-10 claim is dropped entirely; the SKU is **2025-04-29** (OpenRouter catalog April 28, 2025; CSDN April 29, 2025). The real mid-February 2026 Alibaba open-weight event was Qwen3.5-397B-A17B (Feb 16/17). No "Qwen3.5-30B-A3B" SKU exists.
3. **The "April 11, 2026 grouped release wave" (Mistral Small 4 + Medium 3.5 + Large 3; Grok 4.1/4.2; DeepSeek V3.2)** — §1, §2. Rejected in full; the components span December 2025 → August 2026, not a single day. A ten-element debunk table was produced in §1. Component releases (e.g. Mistral Medium 3.5, announced Apr 28, released Apr 29–30, 2026) retain their individual dates.
4. **GPT-5.6 Sol preview date (July 30, 2026 vs June 25–26 preview / July 9 GA)** — §1. Wave 2's July 30 date superseded: government-reviewed preview June 25–26, GA July 9, 2026 (bitrouter PR #674, pureai.com, superagent registry).
5. **Opus 4.6 pricing ($15/$75 vs $5/$25)** — §1. The businessworld.in $15/$75 figure was dropped as an extended-tier/editorial error against the $5/$25 consensus.
6. **Grok 4.7 TB4.0 (37.58% vs 38.0%)** — §1, §13. Both kept with provenance: 37.58% (xAI Grok Build harness) vs 38.0% (company announcements); the gap is harness/rounding noise, not a discrepancy.
7. **Muse Spark 1.2 AA (54 vs 57)** — §1. Both kept with methodology labels rather than forcing a single number.
8. **"SpaceXAI" vs xAI** — §1. The "SpaceXAI" naming (press-level drift in September Grok 4.7 coverage) was rejected in favor of xAI; no corporate confirmation of the merged name.
9. **"TurboQuant by Red Hat" vs Google Research** — §6, §7. Corrected: Google Research invented TurboQuant (arXiv:2504.19874, ICLR 2026, blog 2026-03-25); Red Hat published an independent vLLM evaluation on 2026-05-11 (llmkube#308). Attribution is invention-vs-evaluation clean.
10. **"MLA exactly 4× batch/GPU" vs measured figures** — §7. PARTIALLY VERIFIED: no primary measurement pins "exactly 4×"; it circulates as shorthand for the 93.3%/5.76× family. Consolidated wording uses 93.3% (DeepSeek-V2 paper), 5.76× throughput, 14.2× toy arithmetic [DIRECTIONAL]; "4× batch/GPU" is flagged [UNVERIFIED] shorthand.
11. **"MLA ~90% reduction"** — §7. VERIFIED-conservative: use 93.3% with the DeepSeek-67B baseline cited; controlled ablation ~96%; raw element ratio ~98.2%.
12. **"MLA supported by all engines" vs 3/4** — §7. Corrected to **3 of 4**: vLLM (late 2024, v0.6.x window, exact minor [UNVERIFIED]), SGLang (v0.3, Sep 2024), TensorRT-LLM (supported, version unpinned [UNVERIFIED]). TGI had Gaudi-branch MLA only and was archived 2026-03-21 — never mainline CUDA.
13. **"DSA reduces KV cache" vs attention compute** — §5, §7. Corrected: DSA (and MSA/IndexShare) reduce attention *compute*; KV-cache reduction requires separate compression (V4's CSA/HCA: 27%/10% and 10%/7% vs V3.2 at 1M [VENDOR]; V4.1's CED + FP4-E2M1 + SWA Bounded Replay).
14. **"FP4/NVFP4 in production"** — §7, §8. Restrained: FP8 E4M3 is the 2026 production serving dtype; FP4/NVFP4 are Blackwell/SM100-native or experimental (SGLang v0.5.6+), INT4 method-dependent. NVFP4 ≠ MXFP4 (block-16 vs block-32, FP8-E4M3+FP32 scales vs E8M0, proprietary vs open spec) — never conflated.
15. **"Dynamic drift correction" as a named technique** — §7. [UNVERIFIED] as a named technique across all waves; consolidated wording must map the phrase to documented mechanisms (MiKV's Dynamic Outlier Awareness, TurboQuant's QJL, GEAR's X ≈ D̂ + L + S, ResQ's PCA splits, WKVQuant's 2D alignment, Kitty, FoveatedKV).
16. **O(N²)→constant KV memory phrasing** — §7. Corrected: O(n²) is attention *compute*; KV *memory* is O(n); hybrid/recurrent layers make per-layer *state* O(1). Consolidated wording must not say KV memory was O(n²).
17. **vLLM v0.28.0 mislabeled "vLLM-Omni"** — §6, §7. vLLM v0.28.0 (2026-08-26) is a *core* vLLM release; Wave 2's "vLLM-Omni v0.28.0" mislabel retired, with Omni treated as a separate repo/product line. The V1-only milestone (V0 removal) is dated **v0.16.0 (March 2026)**, not v0.28.0. PagedAttention is a legacy path since **v0.25.0 (July 2026)** and must never be presented as the current core mechanism.
18. **The "+29% SGLang vs vLLM" figure counted twice** — §6. One benchmark (H100/Llama 3.1 8B community result: 16,200 vs 12,500 tok/s; AIMultiple, 1K ShareGPT, bfloat16), cited across Wave 1 and Wave 2 — consolidated to a single citation with the same-kernel (FlashInfer) control noted (canonical in §6b).
19. **"Inference ≈70% of AI cloud spend" vs Gartner 55%** — §14. CONTRADICTED at global level: Gartner's audited 55% figure (2026; 59% in 2027) wins; the 70–80% range is retained only for the narrower scope of GPU-cloud spend among production AI teams (secondary compilations). The two scopes must never be collapsed.
20. **"Largest AI cloud contract in history" (CoreWeave–Meta ~$21B)** — §14. Superlative NOT demonstrated industry-wide: the parallel Nebius–Meta up-to-$27B deal (2026-03-16) prevents any categorical claim. Canonical one-record appearance of the CoreWeave–Meta deal, without the superlative.
21. **NVL144 vs NVL72 naming** — §15. The GTC 2025 name NVL144 counted dies (72 × 2); the 2026 name NVL72 counts packages. Same rack. Every "Vera Rubin NVL144" entry maps to NVL72; do not merge with Rubin Ultra NVL576 (2027, HBM4e, NVLink 7).
22. **Rubin "accelerated"/"pulled-in" framing** — §15. Q3 2026 shipments sit inside the original H2 2026 window; canonical status is **on schedule**, not pulled-in. CES 2026 was production *confirmation*, not the announcement (announced GTC 2025, March 2025; Helios first teased Advancing AI, June 2025).
23. **NVIDIA–Groq "acquisition" vs license** — §15, §20. Wave 1's acquisition framing corrected to the verified structure: **non-exclusive license + talent** (Dec 2025). "Sohgo" → **Sohu**; source typo "Nevus" → **Nebius**.
24. **"agreed to acquire" (NVIDIA–Hugging Face) vs completed purchase** — §16. Primary sources (Form 8-K definitive agreement signed 2026-09-02; confirmation 2026-09-03; first report 2026-08-26 The Information) resolve toward agreed-to-acquire, never completed purchase. The IT Guys' 2026-08-27 "no signed contract" caution is superseded but preserved as a dated note.
25. **"927K+ datasets" (Hugging Face)** — §16. Dropped as [UNVERIFIED, no supporting source]; replaced with 500K (2026-09-03 deal coverage) plus the 730K–1M methodology range. The Wave-2-via-Wave-1 730K+ vs September-3-press 500K conflict is recorded as methodological (public vs total repos), not averaged.
26. **CVE-2026-25874 dual CVSS (9.8 vs 9.3)** — §16. Not a contradiction: a scoring-version difference (v3.1 vs v4.0). Dates fixed: NVD 2026-04-23, PoC 2026-04-28; status "unpatched as of 2026-09-22"; companion CVE-2026-0599 kept as footnote only.
27. **Kimi K2.6 release dates (April 13 vs April 20 vs April 21)** — §3. All three are real phases: preview opened ~April 13, GA April 20–21. Consolidated: GA **2026-04-20/21**, preview noted. Context 256K vs 262,144 is the same figure (shorthand vs exact tokens). GLM-5.2 weights June 16 vs June 17: pin 2026-06-16, note the HF card dated 2026-06-17 (±1 day).

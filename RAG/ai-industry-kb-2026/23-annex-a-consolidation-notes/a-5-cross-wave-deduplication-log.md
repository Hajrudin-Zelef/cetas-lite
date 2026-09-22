---
id: ai-industry-kb-2026/23-annex-a-consolidation-notes/a-5-cross-wave-deduplication-log
title: "A.5 Cross-wave deduplication log"
domain: appendix
role: appendix
task: reference
actors: ["Alibaba", "Anthropic", "ByteDance", "Cerebras", "China", "CoreWeave", "DeepSeek", "Glasswing", "Google", "Groq", "Hugging Face", "Meta", "Microsoft", "MiniMax", "Mistral", "Moonshot", "Nebius", "Nscale", "Nvidia", "OpenAI", "SGLang", "TensorRT-LLM", "United States", "Unsloth", "Z.ai", "vLLM", "xAI"]
dates: ["2025-04-05", "2026-01", "2026-02-20", "2026-05", "2026-05-05", "2026-05-14", "2026-07", "2026-07-22", "2026-08-19", "2026-09", "2026-09-01"]
keywords: ["acquisition", "agent", "agentic", "agents", "agi", "arr", "asic", "astra", "attention", "awq", "benchmark", "benchmarks"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [11268, 11398]
section: "Annex A — Consolidation notes"
sha256: ae8b997ed69cd99bb0bb39cd6ce674b68a006d3685ea1fa424d4cbdac02f09af
---

# A.5 Cross-wave deduplication log

## A.5 Cross-wave deduplication log

Content appearing in multiple waves was merged into a single canonical home; other sections carry a one-line cross-reference only. (Direction: canonical home ← sections that cross-ref.)

### §1 — Frontier closed models
- Fable 5 / Mythos 5 export-control suspension (June 12 → July 1): full incident record in **§17**; one-line cross-ref in §1.
- Company funding rounds: canonical in **§20**; cross-ref in §1.
- Nano Banana 2 Lite: deep dive in **§12**; one-line record in §1.
- Mistral Small 4 / Medium 3.5 / Large 3 details: open-model sections; recorded in §1 only to debunk the April 11 fiction.
- Gemma 4 12B specs: open-weight sections; §1 carries only the dating correction.
- ByteDance Seed models: canonical in **§4**; §1 cross-ref only.
- Qwen3.5: canonical in the open-weight chronology (**§2**); §1 drops it entirely (only the mid-February pre-release reporting remains).
- DeepSeek V4: canonical in the DeepSeek chronology (**§5**); §1 drops it.
- Muse Spark 1.3 (Sept 2) and Grok 4.7: canonical in the Wave 2.1 deltas; §1 includes them only for timeline completeness.
- Llama 4 (April 5, 2025) and Behemoth: **explicitly excluded** — 2025 events, outside §1 scope.
- Galaxy S26: excluded — hardware, not a frontier model release.

### §2 — Open-weight chronology
- Architecture deep dives: GLM → §3; Qwen Gated DeltaNet / KDA → §4; DeepSeek DSA/CED/V4.1 → §5. This section's Spec table is the canonical per-model 1-line spec record — other sections cross-reference, never re-specify.
- Inference engines → §6; quantization (GGUF, NVFP4, MXFP4, Unsloth v3.0 mechanics) → §8; full pricing tables → §14; frontier closed-model details → §1.
- Dated release rows live in §2 (1–2 lines each); §5's DeepSeek material collapses to 1–2-line rows + cross-ref to §5.
- §4's Qwen "Key dated facts" bullets collapse to one-line §2 rows with a "→ §4" pointer; the spec table, benchmark reality checks, license map, pricing chronology, and Omni strategy deep dive stay in §4.

### §3 — Chinese GLM/Kimi/MiniMax (deep file)
- GLM-5.2 architecture detail and IndexShare mechanics beyond the parameter reconciliation: canonical in wave3/02-openweight-chronology.md (Claim 3).
- GLM-5.2 benchmarks and text-only status in the frontier comparison: wave1/07 §2.4 and wave3/06 §1.3.
- GLM-5.3 release and safety-hold narrative: wave1/07 §2.4 and wave2.1/05 §4 — this section adds only the full benchmark verification, the always-on-reasoning breaking change, and the rumor-phase dating.
- GLM-5.3-Flash / "Ox Alpha": wave1/07 §2.4 — nothing new here.
- Kimi K3 chronology, launch, and "beats Opus 4.8" evidence: wave3/02 Claim 4 — this section keeps only technical-report architecture detail plus flagged conflicts.
- MiniMax H3 model detail (H3-Omni-Transformer, 2K upscaling, H3-Context-IR, open-weights Aug 2–3): wave2.1/07 §1.1 — this section records only the alias resolution, the territorial license restriction, and the 2.7T rumor.
- Seedance 2.0 / 2.5: wave3/06 §1.3 — cross-reference only.
- OpenVuln / VulnHunter narrative: wave1/07 §2.4 — figures confirmed here, not duplicated.

### §4 — Qwen / Seed / Ling (deep file)
- Feb→May 2026 Qwen skeleton: wave3/02 (not re-derived).
- Gated-DeltaNet hybrid mapping, Qwen3.8-Max-0902 detail, Qwen3.8-27B detail: wave2.1/05 (cross-ref only).
- Seedance 2.5: wave3/06 §1.3 (cross-ref only).
- Artifact/service split enforced: Qwen3.8-Max (hosted) vs Qwen3.8-2.4T-A95B (weights) vs Qwen3.8-Max-0902 (checkpoint) stay separate index entries in every index.

### §5 — DeepSeek (deep file)
- Dated release rows: canonical in §2 (this section keeps the full narrative).
- Serving/pricing economics: summarized here; full treatment in §9 / §14 of the final document.
- V4 fine-grained-MoE architecture reconstruction (256/384 routed experts, 6 active, hash-routed first blocks, sqrt(softplus) affinity, FP4 expert weights): canonical in wave2/03 §4 — cross-referenced, all figures [PARTIALLY VERIFIED].

### §6 — Inference engines (two parts, 06a vLLM / 06b SGLang)
- SGLang internals and the vLLM-vs-SGLang shootout: canonical in **06b** (06a keeps minimal comparative framing: the +29% figure cited once, per-model engine choice).
- KV-cache mechanics, FP8 KV dtype, O(n)/O(n²) analysis: canonical in **§7** (06b keeps only engine-level outcome figures: hit-rate ratios, throughput deltas, release versions).
- Quantization dtype mechanics (FP8/MXFP4/TurboQuant/UD quants): canonical in **§8**.
- GPU pricing and cloud inference markets: canonical in **§14**.
- TPU hardware specifics: **§15** (SGLang-JAX recorded in 06b only as the engine-backend fact).
- vLLM internals, vLLM-Omni, vllm-gguf-plugin, Inferact, NIM 2.0, Dynamo vs llm-d, spend flip, per-token economics: canonical in **06a**.
- SGLang (RadixAttention, scheduling, releases, commercialization, claims) + adjacent runtimes (TensorRT-LLM, llama.cpp server, emerging engines): canonical in **06b**.
- Wave-1 archive flags (MiMo-V2.6 spec discrepancies, 1M-context claim harmonization): explicitly out of §6's scope — harmonized centrally.

### §7 — KV cache & attention (two parts, 07a mechanisms / 07b disaggregation)
- **07a** is the home for: KV-cache sizing law; attention variants (MHA/MQA/GQA/MLA) + MLA engine-support who/when; shared/hybrid layouts (Gemma 4, MSA, IndexShare, DSA); post-hoc compression zoo; error-correction canon (GEAR, MiKV, ResQ, WKVQuant, ZipCache/PrefixQuant/MiniKV, LESS, KVLinC, Kitty, FoveatedKV); quantized-KV production dtypes; Mamba-2/Gated-DeltaNet; sizing ladders. Engine per-release kernel detail lives in **§6**.
- **07b** is the home for: disaggregation, tiered storage, RDMA, prefix caching, 1M-context product status; attention-variant mechanisms appear only as cross-refs to 07a; vLLM/SGLang version specifics as timeline anchors only (detail in **§6**).

### §8 — Quantization (two parts, 08a formats/tooling / 08b precision/production)
- **08a** owns: formats and tooling (GGUF, K-quants/IQ-quants, Dynamic 3.0, EXL2→EXL3, AWQ/GPTQ toolchain turnover, deployment decision matrix, llama.cpp, edge/ARM variants, NF4/QTIP, workflows). The NVFP4≠MXFP4 anatomy table stays in 08a (format-definition fact); the vLLM-vs-EXL3 shootout stays in 08a (adjudicates a format/engine claim, [COMMUNITY]).
- **08b** owns: the FP8/FP4/NVFP4/MXFP4 **production arc**, TurboQuant production verdicts, KV-cache quantization methods, BitNet, Bonsai-27B, cost arithmetic, quality measurement.
- TurboQuant method detail: canonical in **§7** (08b carries a one-line cross-ref).
- Dynamic UD format mechanics: canonical in **§8** (§10 records only versioning facts: 2.0 current July 2026, superseded by v3.0 2026-08-19).

### §9 — MoE (two parts, 09a releases/routing / 09b economics/serving)
- **09a** owns: release facts, per-model specifications, routing theory (top-k → latent routing, load-balancing, fine-grained bandwidth correction), leadership map, historical context, open research questions.
- **09b** owns: training economics, inference economics (active-vs-total, KV-cache nuance), serving cookbook (vLLM/SGLang EP, EPLB, MTP speculative decoding), quantization recipes, pricing ladder, license landscape, benchmark scoreboard.
- Unsloth training mechanics (split-LoRA, cut cross-entropy, kernels): canonical in **§10** (09b references only where serving economics are affected).

### §10 — Training / Unsloth
- Dynamic Quant UD mechanics → **§8**; GGUF serving/format → **§8** (this section covers only the post-training export API); fine-grained MoE architecture and bandwidth analysis → **§9** (verdict referenced, not re-analyzed).
- Wave 2 vs Wave 2.1 overlap (February MoE headline, NVIDIA-blog numbers, torchtune comparison, 70–80% VRAM verdicts) eliminated — each fact appears once, delta additions merged into the trajectory.
- Unsloth Studio/Desktop productization and GRPO details overlapping §6's Unsloth platform section (§1): §10 owns training-kernel, benchmark, and PEFT-landscape content and defers platform minutiae to §6.

### §11 — Multimodal open
- Closed frontier omni models (Gemini Omni Flash and siblings) → **§1**; full video-generation coverage (Sora, Seedance, Kling, Genie) → **§12**; agentic frameworks, benchmark trust-crisis material, agent pricing → **§13**. This section keeps only the open-weight omni lineage plus the two September 2026 any-to-any entrants (MiniMax H3, HiDream-O1-Video-1.0).
- Full open-weight licensing detail for MiniMax H3: canonical in **§11** (§12 keeps capability/ranking claims and cross-references).
- vLLM-Omni one-liner → **§6**; KV-cache engineering detail (CSA2, Engram, MXFP4) shared with test-time-compute economics (§13-adjacent/agentic section).

### §12 — Multimodal video
- Open-weight omni models → **§11**; agents and agentic frameworks → **§13**; Grok Imagine → **§1**; general AGI/narrative audit material from the wave3 check file is excluded as out of scope.

### §13 — Agents / MCP
- Agent safety incidents (Cline CI compromise Feb 9, "Comment and Control" Apr 16, Unit 42 campaigns, Zscaler, "Agents of Chaos", Promptware survey, LiteLLM CVEs): full treatment in **§17** — one-line cross-reference only here.
- Agent-economics figures (Cursor $4B ARR, Cognition Series D/E, Replit valuation, Devin run-rate, Harness Engineering claims): kept here only in engineering context; pure economics figures belong in **§19** (attributed forecasts).
- Test-time compute architecture (DeepSeek V4.1-Flash KV engineering, workload inversion): summarized here for agent context; model-side detail cross-references **§5**.
- A2A/ACP: covered here as protocol complements; full protocol-governance detail (AAIF directed fund, SEP process) stays in §13.
- MCP governance as enterprise control plane: cross-link §13 ↔ §18.

### §14 — Cloud / neoclouds
- Chip specifications (FLOPs, architectures, Rubin NVL72, Etched Sohu, Cerebras WSE-3): canonical in **§15** — this part carries only silicon *economics* (rental rates, resale, contract pricing, megadeals, valuations), with cross-refs where a price claim depends on a spec.
- FP8/FP4 KV quantization, TurboQuant, AWQ, GGUF matters → **§8** (referenced only where they enter unit economics).
- Capacity megadeals (CoreWeave–Meta $21B, Nebius–Meta $27B, Nebius–Microsoft $17.4B, OpenAI Rubin at-scale) live in §14; §15 carries only one-line cross-refs.

### §15 — Hardware / chips
- Capacity megadeals → **§14** (one-line cross-refs here); export-control actions (DOJ/Super Micro Mar 19, 2026; NDRC/Manus Apr 27, 2026) → **§18** (cross-refs only here).

### §16 — Hugging Face
- July 2026 autonomous-breach incident targeting Hugging Face: full account in **§17** (one line here).
- Robotics deployment detail: full account in **§19** (chronology and specs only here).
- ggml.ai → HF (2026-02-20): recorded here as precision-format context; acquisition mechanics belong to the ecosystem parts.

### §17 — Safety incidents (two parts, 17a incidents / 17b breach + nuclear)
- **17a** (this part): Fable 5 / Mythos 5 export-control skeleton (June 9 launch → June 12 suspension → June 30 lift → July 1 restoration; Anthropic "Redeploying Claude Fable 5" blog trigger; June 26 Project Glasswing Mythos-5 reprieve; CAISI safeguards assessment — from wave3/01 §3) plus the granular overlay (17:21 ET timestamp, ~90-minute compliance window, Lutnick → Tom Brown letter, deemed-export mechanism, ~20-hour report-to-order arc). The July 22 OpenAI/HF breach appears here only as the Kill Switch Act's trigger (one-line cross-ref).
- **17b**: owns the July 22, 2026 OpenAI sandbox-escape/Hugging Face breach itself (two models, safety restrictions disabled, sandbox escape, internet reach, HF production-server compromise) plus the nuclear-wargaming reports (Payne's Project Kahn preprint, Feb/Aug 2026). Model release facts (GPT-5.6 Sol) deduplicated to §1 cross-references.
- **§18** owns: Pax Silica (supply-chain security), 464 state chatbot bills, Kill Switch Act legislative-procedure future. Cross-refs are one-directional (§18 → §17).

### §18 — Governance / regulation
- All incident detail (BIS order timing/trigger/compliance, Warner quote and walk-back, OpenAI–HF breach, AG subpoenas): canonical in **§17** — §18 keeps only regulation/policy analysis and regulatory consequences. Dedup test for the coordinator: a paragraph reading like a breach narrative belongs in §17; rules, penalties, enforcement actions, compliance duties belong here.
- Heretic and AISI material (also in wave1/07 licensing/thesis sections): cited here only for the **enforceability** argument, never re-analyzed; the "weights cannot be recalled" finding cross-links to wave1/07 §8 (not re-argued).

### §19 — Economy / predictions
- Funding-round details (Cognition, Cursor raises): cross-referenced to **§20**.
- Model capability claims (Omni Flash, Astra, Fable 5.1 benchmarks): cross-referenced to **§1–5**, not relitigated.
- The multimodal "revolution" narrative audit keeps only its *economic* half; the model-version scoreboard lives in wave3 CHECK 1 (excluded by brief).

### §20 — Startup funding
- IPO pricing mechanics: canonical in **§21** (Cerebras's May 14, 2026 IPO mentioned here once with a §21 cross-ref, no pricing table).
- Product facts (Sohu ASIC specs, Groq LPU architecture, Skild Brain, HX-2/Altra/CA-1, Devin/Windsurf product detail, NEO specs): company/product sections — one identifying clause each here.
- RadixArk ($400M valuation, reported January; formal launch 2026-05-05) and Inferact ($150M raise, January 2026): named in the brief but NOT present in wave5/01-startup-funding.md — carried as [UNVERIFIED] brief-flag lines only.

### §21 — IPOs / public markets
- Round amounts → **§20** (one-line cross-refs here); company fundamentals (Anthropic revenue detail → lab fundamentals; Nebius deals → wave3/04; OpenAI structure → wave4/03 §6); Cerebras IPO mechanics → wave2/05 §5. This part adds only post-filing trajectories, priced-and-trading updates, Nscale, and the broader queue.

### §22 — Robotics
- Hugging Face robotics (HopeJR 2026-09-01, Reachy Mini refresh, CVE-2026-25874): canonical in **§16** — one-line cross-refs only.
- Robotics funding rounds: canonical in **§20** — one-line cross-refs only. CVE-2026-25874 is not re-explained.
- wave3/05 (Hugging Face) and wave3/06 (multimodal) were cross-checked and explicitly not duplicated.

### Label-inventory note (cross-cutting)
All provenance labels used in the source waves ([VENDOR], [UNVERIFIED], [COMMUNITY], [DIRECTIONAL], [PARTIALLY VERIFIED], [SECONDARY], plus §17's [CONFIRMED], [PRIMARY], [POLICY], [ANALYSIS], [CONSOLIDATION RULE]) are carried through assembly attached to their figures and claims — the final document's fact-grade depends on them. No label was invented for sources that were already qualified in the waves.

---
id: ai-industry-kb-2026/23-annex-a-consolidation-notes/10-training-unsloth
title: "§10 — Training / Unsloth"
domain: appendix
role: appendix
task: training
actors: ["Anthropic", "Cerebras", "CoreWeave", "DeepSeek", "Glasswing", "Google", "Groq", "Hugging Face", "Meta", "Microsoft", "MiniMax", "Nebius", "Nscale", "Nvidia", "OpenAI", "United States", "Unsloth", "vLLM", "xAI"]
dates: ["2026-01", "2026-02-20", "2026-05-05", "2026-05-14", "2026-07", "2026-07-22", "2026-09", "2026-09-01"]
keywords: ["training", "acquisition", "agent", "agentic", "agents", "agi", "arr", "asic", "astra", "awq", "benchmark", "benchmarks"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [11337, 11398]
section: "Annex A — Consolidation notes"
sha256: 7fdf6ba6a5af720c8828d3e250a3c756aef35a7a609f0d4e86eaabae05c59395
---

# §10 — Training / Unsloth

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

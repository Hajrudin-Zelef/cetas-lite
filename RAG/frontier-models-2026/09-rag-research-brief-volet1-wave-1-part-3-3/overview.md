---
id: frontier-models-2026/09-rag-research-brief-volet1-wave-1-part-3-3/overview
title: "RAG Research Brief — VOLET1 / Wave 1 (Part 3/3)"
domain: rag-research-brief-volet1-wave-1-part-3-3
role: deep-dive
task: reference
actors: ["AWS", "Alibaba", "Anthropic", "China", "DeepSeek", "Google", "Hugging Face", "Moonshot", "OpenAI", "OpenRouter", "Xiaomi", "Z.ai", "xAI"]
dates: ["2025-07-11", "2026-04", "2026-06-12", "2026-07", "2026-07-16", "2026-09-18", "2026-09-22"]
keywords: ["research", "agent", "agentic", "agents", "astra", "attention", "aws", "bedrock", "benchmarks", "claude", "context window", "cost"]
source: docs/RAG/Grands titres IA modèlesEN.md
source_anchor: ""
source_lines: [395, 454]
section: "RAG Research Brief — VOLET1 / Wave 1 (Part 3/3)"
sha256: f3d6d7c74f8338d096f4484c3d09e9272ee11049a60d2972006a720ea9e520dd
---

# RAG Research Brief — VOLET1 / Wave 1 (Part 3/3)
## Frontier AI Models Feb–Sept 2026: Chinese Open-Weight Dominance
**Compiled: 2026-09-22 — Research coverage for RAG ingestion**

---

## 1. XIAOMI MiMo-V2.6 (Pro / Flash / Pro-UltraSpeed)

**Release date:** Weights and announcement September 21–22, 2026 (Artificial Analysis lists Sept 21; Xiaomi announcement post dated Sept 22). Preceded by an unusual public "Live RL" training dashboard (mimo.xiaomi.com/rl) streaming the final RL run.

**Variants:**
- **MiMo-V2.6-Pro** — flagship, native omnimodal (text + image + audio + video in), 1M-token context window, sparse MoE: ~1.02T total parameters / ~42B active per token.
- **MiMo-V2.6-Flash** — efficiency tier: 309B total / ~15B active per token.
- **MiMo-V2.6-Pro-UltraSpeed** — inference tier delivering up to **20× faster output at the same quality** for latency-sensitive workloads; priced 10× the Pro tier.

**Architecture & training (key RAG facts):**
- Natively omnimodal, trained with scaled reinforcement learning on verifiable complex tasks (coding, general agents, visual, cybersecurity) in a single run.
- Production RL run: <6 days, 30 RL steps each for Flash and Pro, ~750,000 trajectories, 1,568 samples per update, 16 async rollouts per prompt, 3.5–3.7B tokens per step, contexts up to 1M tokens.
- Cost: ~$850K (Flash) and ~$2.62M (Pro) — comparatively modest frontier-class post-training bill.
- Anti-reward-hacking measures: frozen router, adversarial evaluation, anomaly detection, verifier cross-checks; graders that rank successful trajectories by quality, not just pass/fail.
- DeepSWE v1.1 (held-out) rose during the run: Flash 48.8 → 65.68; Pro 58.4 → 72.57 (final vendor table reports Pro 71.9).
- Predecessor MiMo-V2.5 (April 2026) scored only 19.0 on DeepSWE v1.1 — a ~50-point generational jump.

**Benchmarks (vendor-reported, Sept 2026):**
- **Artificial Analysis Intelligence Index v4.3: 46.32 → 46** — highest open-weight score on the leaderboard; ahead of GLM-5.3 (45), Kimi K3 (44), Qwen3.8 Max (45 on 0902 snapshot); tied with closed Grok 4.7 (46). Ahead of Grok 4.6 (44) and Gemini 3.8 Flash (41) on the same index.
- DeepSWE v1.1: Pro 71.9 vs Claude Opus 5 74.0, GPT-5.6 Sol ~73, Kimi K3 69.0, DeepSeek V4.1 Flash 74.2.
- AutomationBench v1.0.6: Pro 53.1 — beats Claude Opus 5 (50.3), GPT-5.6 Sol (45.8).
- Terminal-Bench 2.1: 89.9 (best in Xiaomi's table); Terminal-Bench 4.0: 34.9 (trails Opus 5 at 49.0 — hardest terminal agentic work still closed-lab territory).
- OSWorld-Verified: 82.0; Toolathlon-Verified: 76.9; MiMo Code Bench: 63.2; GDPval-AA 2.1 Elo: 1673 (vs Opus 5 1708); Agents' Last Exam: 31.6 (ties Opus 5).
- Visual: MiMo VisualCoding 72.3 — ahead of Opus 5 (70.0).
- Cyber: CyberGym 94.0; but ExploitBench 47.9 vs GPT-5.6 Sol 78.5; SEC Bench Pro 66.3 vs Sol 79.1 — offensive security remains a gap.
- Competitive coding gap: ProgramBench 26.5 vs Opus 5 37.0.

**License & price:** MIT license on Hugging Face + ModelScope. API: Flash $0.14/M input / $0.28/M output; Pro $0.435/M / $0.87/M (unchanged from V2.5); UltraSpeed 10× Pro. Artificial Analysis: $0.13 per Intelligence Index task; ~134 tok/s output. Xiaomi claims 1/20–1/60 the cost of overseas models at the same intelligence level. Available via OpenRouter.

**Positioning:** "Strongest open-source model to date" per Xiaomi; on par with Claude Opus 5 and GPT-5.6 Sol on most agent benchmarks per the launch post; still trails the top closed models (Claude Fable 5.1, GPT-6 Astra) on the hardest evals. Controversy note: some commentary alleges training-data borrowing from Anthropic's Claude — unverified, but worth flagging in RAG as disputed.

---

## 2. KIMI K3 / K2.x (Moonshot AI)

**Release timeline:**
- Kimi K2 (1.04T / 32.6B active, 128K ctx): July 11, 2025. K2-Instruct-0905 (256K): Sept 2025. K2 Thinking: Nov 2025. K2.5 (multimodal MoE, 1T/32B, 256K): Jan 27, 2026. K2.6: Apr 20, 2026. K2.7 Code: June 12, 2026.
- **Kimi K3: hosted service July 16, 2026; downloadable weights late July 2026 (~July 27); Amazon Bedrock GA September 18, 2026.**

**Architecture:** 2.8T total parameters — "first open model to reach 2.8T" per Moonshot. MoE "Stable LatentMoE": 16 of 896 experts active (~50B active). Innovations: Kimi Delta Attention + Attention Residuals (long-sequence/depth information flow); quantization-aware training from SFT onward (MXFP4 weights, MXFP8 activations). ~2.5× scaling-efficiency gain over K2. Native vision; 1M-token context.

**Benchmarks:**
- Artificial Analysis Intelligence Index: K3 (max) ~44 (v4.3.2) in Sept 2026 snapshots; earlier v4.1.1 readings ~59.7–60. Index-version changes make cross-version comparison invalid — note methodology version when citing.
- BenchAlign ~74.4 (per user brief; Chinese coding/agent composite where K3 frequently leads Chinese open models).
- Vendor claim: frontier-level across Moonshot's suite, consistently beating other tested open models; still trails proprietary Claude Fable 5 and GPT-5.6 Sol overall.
- K2.6: 80.2% SWE-bench Verified, 58.6% SWE-bench Pro (vendor).
- Known limitations (Moonshot): trained in "preserved thinking history" mode — unstable if harness drops thinking history; long-horizon bias can cause unexpected autonomous decisions on ambiguous instructions.

**License & price:** K2 series: modified MIT. K3: new "open weight" license — large MaaS businesses (>$20M/yr revenue) must sign a separate agreement; NOT traditional open source. Weights 1.56 TB on Hugging Face. Commercial access via OpenRouter, Moonshot API, and now AWS Bedrock (first open-weight model on Bedrock with explicit prompt caching; runs in Bedrock security/compliance boundary).

**Positioning:** Moonshot is the most open-weight-forward Chinese frontier lab and the only one repeatedly setting the open-model size ceiling (K2 → K2.5 → K3). K3 competes as the top Chinese open model on several composites but was overtaken by MiMo-V2.6-Pro (46) and Qwen3.8-Max-0902 (45) on AA Index v4.3 in Sept 2026.

---


---
id: collect-261001-general-networking/general-networking/grands-titres-ia-modelesen-5
title: "VOLET 1 — Vague 1 (EN)"
domain: general-networking
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "China", "DeepSeek", "Google", "Hugging Face", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "Xiaomi", "Z.ai", "xAI"]
dates: ["2026-04", "2026-09-22"]
keywords: ["agent", "agentic", "agents", "astra", "attention", "benchmark", "benchmarks", "claude", "context window", "cost", "cyber", "cybersecurity"]
source: docs/RAG/collect-261001-general-networking/Grands titres IA modèlesEN.md
source_anchor: ""
source_lines: [351, 435]
sha256: b2e41aa603bf9a2715af841388061135da6183ba6aedcbfc89a4c5c8f32a48fb
---

# VOLET 1 — Vague 1 (EN)

## 3.10 GLM-5.3-Flash — benchmarks (vendor-reported)
| Benchmark | GLM-5.3-Flash | Reference |
|---|---|---|
| DeepSWE v1.1 | **63.4** | vs 46.2 for GLM-5.2 |
| AA Intelligence Index | **57** | — |
| GDPVal-AA v2 | **1773** | Claude Opus 4.8: 1582 |
| AutomationBench | **48.8** | Opus 4.8: 41.0 |
| Terminal-Bench 2.1 | 84.3 | Opus 4.8: 85.0 |
| HLE w/ Tools | 55.3 | Opus 4.8: 57.9 |
| Z.ai Code Bench v1.0 (max effort) | 29.0 | Opus 4.8: 29.5 — near parity |

## 3.11 GLM-5.3-Flash — license, price, availability
- **MIT license, open weights day one** (contrast with GLM-5.3 non-Flash, staged behind safety review).
- **API pricing (per 1M):** $0.15 input / $0.50 output / $0.03 cached input — ~33× cheaper input and ~50× cheaper output than Anthropic Opus 5 ($5/$25); ~$0.045/task on AA Intelligence Index v4.1.1 with discounts (~10× cheaper than prior frontier-tier intelligence per Z.ai).
- **OpenRouter:** live (formerly Ox Alpha); Z.ai API; GLM Coding Plan; ZCode.
- Gap flagged for wave 2: the brief's note (SWE-bench Verified 70–75%+ "per snapshots" for the GLM-5.3-FlashX/MiniMax class) could not be confirmed in web sources — no SWE-bench Verified figure published for GLM-5.3-Flash to date.

---

# 4. Cross-cutting comparison (vague 1 models)

| Dimension | Grok 4.20 | DeepSeek V4-Pro | DeepSeek V4.1 Flash | GLM-5.2 | GLM-5.3 | GLM-5.3-Flash |
|---|---|---|---|---|---|---|
| Release | 17 Feb 2026 (beta) | ~17 Feb 2026 | 10 Sep 2026 | 13–16 Jun 2026 | 14 Aug 2026 | 26 Aug 2026 |
| Weights | Closed | Open (MIT) | Open (MIT) | Open (MIT) | Open (MIT, staged) | Open (MIT) |
| Params (total/active) | n.d. | 1.6T / 49B | 552B / 8B–16B (+196B Engram) | 744B / 40B | ~743B / ~40B | 320B / 18B |
| Context | 256K (→2M) | 1M | 1M (384K out) | 1M | 1M | 1M |
| Multimodal | text+image+video | vision (V4.1) | native vision | text only | text (coding) | native image+video |
| Signature tech | 4-agent council, adversarial consensus | Engram memory, mHC, Muon | Causal Encoder-Decoder, FP4 KV, CSA2 | IndexShare attention | post-training scaling, cyber training | sparse+linear attention, CN-chip serving |
| SWE-bench Verified | n.d. | 80.6% | — | — (Pro: 62.1%) | — | — |
| Terminal-Bench | n.d. | 67.9% (2.0) | 90.6 (2.1) | 81.0 (2.1) | 28.3 (3.0) | 84.3 (2.1) |
| API in/out ($/1M) | 1.25 / 2.50 | 1.74 / 3.48 | 0.15 / 0.60 (off-peak) | 1.40 / 4.40 | n.d. (Coding Plan) | 0.15 / 0.50 |

# 5. Global trend (Feb–Sep 2026)
- **Chinese open-weight models (DeepSeek, Qwen, Kimi, GLM/Z.ai, MiniMax, MiMo) dominate the open-source top tier** in coding, agentic work, and price/performance. The pattern repeats: frontier-class scores at 5–10% of Western API prices, MIT-licensed weights, 1M-token contexts as standard.
- **Closed models (Claude Opus/Fable 5.x, GPT-5.x/6, Gemini 3.x, Grok 4.x) retain leads** on the hardest closed benchmarks and enterprise agentic tooling — but the gap on *long-horizon coding* has narrowed to single-digit points (GLM-5.2 vs Opus 4.8), and DeepSeek V4.1 Flash outright retired its own premium tier.
- **Architectural convergence on inference efficiency:** sparse MoE everywhere, KV-cache compression (FP4, CSA2, IndexShare), conditional memory (Engram), reasoning-effort dials, agent-task synthetic post-training.
- **Hardware sovereignty subplot:** GLM-5.3-Flash serving production traffic on ~100K domestic Chinese accelerators — first credible demonstration that SOTA open models can be trained/served without Nvidia.
- **Benchmark fragmentation:** vendors increasingly report on *different* benchmarks (SWE-bench Verified vs Pro vs DeepSWE vs Terminal-Bench), making like-for-like comparison unreliable — a structural caveat for the RAG.

# 6. Caveats & methodology notes
- All benchmark numbers above are **vendor-published** unless noted; independent evaluations (Artificial Analysis, community harnesses) exist for some models and generally show slightly lower figures.
- An agentic benchmark scores a *system* (model + harness + tools + retries), not a model alone — cross-harness comparisons are not apples-to-apples.
- GLM-5.3 (non-Flash) weights were still pending (~2 weeks post-launch) at last check — verify before citing as "open."
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


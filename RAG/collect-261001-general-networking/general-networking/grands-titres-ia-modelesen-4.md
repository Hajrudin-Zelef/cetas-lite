---
id: collect-261001-general-networking/general-networking/grands-titres-ia-modelesen-4
title: "VOLET 1 — Vague 1 (EN)"
domain: general-networking
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "China", "DeepSeek", "Google", "Hugging Face", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "SGLang", "United States", "Unsloth", "Z.ai", "vLLM"]
dates: ["2026-06", "2026-08"]
keywords: ["agent", "agents", "attention", "attribution", "benchmark", "benchmarks", "claude", "compute", "cyber", "cybersecurity", "deepseek", "distribution"]
source: docs/RAG/collect-261001-general-networking/Grands titres IA modèlesEN.md
source_anchor: ""
source_lines: [267, 350]
sha256: 079d942717e0548b6091a1396606d049316a33dd4d788c3cd25e7fefa27db8fe
---

# VOLET 1 — Vague 1 (EN)

## 2.8 DeepSeek V4.1 Flash — license, price, availability
- **MIT license, weights + repo on Hugging Face and ModelScope** — permissive end of the open-weights spectrum (no revenue threshold, no field-of-use clause); self-host, fine-tune, serve commercially, royalty-free. Inference folder + separate encoding module in repo; community building inference support from day one.
- **Official API pricing (per 1M, off-peak):** $0.15 input / $0.60 output / $0.003 cached input. **Peak** (Mon–Fri 01:00–04:00 & 06:00–10:00 UTC): doubled — $0.30 / $1.20 / $0.006. Weekends always off-peak.
- **Speed (observed):** ~190–427 tok/s community-reported; 215 tok/s p50 in one routing analysis; p50 TTFT 842 ms, p95 2.40 s, ~1.2% error rate.
- **Distribution:** DeepSeek app (rolling out), official API, **OpenRouter** (`deepseek/deepseek-v4.1-flash` — 1,048,576-token context, reasoning default `high`, tools, JSON, streaming; upstreams include DeepInfra, Novita, Venice; 99.92% 3-day availability); partners WorkBuddy, CodeBuddy, OpenCode Go, ClinePass.
- **Positioning:** killed its own premium tier — "a lab deciding the premium tier was no longer worth charging for." Best value in its class on price/perf/agent metrics per multiple independent analyses.

---

# 3. GLM-5.2 / GLM-5.3 / GLM-5.3-Flash (Zhipu AI / Z.ai)

## 3.1 GLM-5.2 — release
- **Released 13 June 2026** on the GLM Coding Plan; **API, chatbot, and MIT open weights followed 16–17 June 2026** (weights ~2–4 days after launch; Hugging Face + ModelScope; official FP8 build day one; community GGUF quantisations — Unsloth — shortly after).
- Unusual launch: **zero benchmark numbers published at release** (no SWE-bench, no Terminal-Bench) — full benchmark suite published with the open weights on June 17.
- Succeeds GLM-5.1 in the GLM-5 family; Z.ai (formerly Zhipu AI) monetises via hosted API + flat-rate GLM Coding Plan, not by withholding weights.
- **"New DeepSeek moment"**: as of 16 June 2026, Artificial Analysis ranked GLM-5.2 the **leading open-weights model on its Intelligence Index v4.1 (score 51)**, ahead of DeepSeek V4-Pro and Kimi K2.6 — a reversal of the GLM-5.1 mid-pack picture.

## 3.2 GLM-5.2 — architecture
- **MoE: 744B total / ~40B active per token** (Z.ai's own spec: "744B-A40B"; some write-ups say 753B — the HF safetensors count includes embedding/output-head weights outside the MoE budget; same model).
- **Context: 1,048,576 tokens (1M)** — 5× GLM-5.1's ~200K; max output 128K–262K (sources vary: 131,072 in model card; 262,144 per AgentRiot).
- **IndexShare sparse-attention technique:** groups of four transformer layers share one lightweight indexer — keeps 1M-context inference costs under control.
- **Dual reasoning modes: High (fast) and Max (deep reasoning)**; also non-thinking mode.
- **Text only — no vision.**
- Agent-first design; **Anthropic-compatible endpoint** → works with Claude Code, Cline, OpenCode, OpenClaw (one base-URL change); exposed in Claude Code as model ID `glm-5.2[1m]`.
- **ZCode:** official Claude-Code-style agent harness shipped alongside, built to drive GLM-5.2.

## 3.3 GLM-5.2 — benchmarks
| Benchmark | GLM-5.2 | Reference |
|---|---|---|
| AA Intelligence Index v4.1 | **51 — #1 open-weights** | ahead of DeepSeek V4-Pro, Kimi K2.6 |
| Terminal-Bench 2.1 | **81.0** | Kimi K3: 88.3; Claude Opus 4.8: 85.0; DeepSeek V4-Flash-0731: 82.7 |
| SWE-bench Pro | **62.1** | — |
| DeepSWE v1.1 | 46.2 | — |
| Agents' Last Exam | 23.8 | — |
| AutomationBench Public | 12.9 | — |
| DSBench-FullStack/Hard (internal) | 68.7 / 59.6 | — |
| (Z.ai vendor table, via VentureBeat) | edges past OpenAI GPT-5.5 on several multi-step engineering benchmarks | vendor's own table, not independent |

- **Positioning:** first open-weights model to close the gap with closed frontier models on *long-horizon* coding — within single-digit points of Claude Opus 4.8 on long-horizon coding tasks; hardest general-reasoning crown still held by closed frontier.

## 3.4 GLM-5.2 — license, price, availability
- **MIT license — fully permissive open weights** (no regional restrictions).
- **API pricing (per 1M, unchanged from GLM-5.1):** $1.40 input / $4.40 output / $0.26 cached input.
- **Speed:** ~141 tok/s.
- Available: GLM Coding Plan, Z.ai API/chatbot, Claude Code, third-party platforms; self-host via SGLang/vLLM-class + GGUF.

---

## 3.5 GLM-5.3 — release (14 August 2026)
- **Released 14 August 2026** (announcement thesis: "Scaling post-training is all we did for GLM-5.3").
- **Same ~743–744B-parameter MoE base as GLM-5.2** — no new base model, no bigger architecture; **all claimed gains from post-training alone** (statement about where the frontier is moving: next jump from training smarter, not bigger).
- Available at launch via **GLM Coding Plan**, working with **ZCode, Claude Code, OpenCode**.
- **Weights staged:** not downloadable at launch — open release planned ~2 weeks later, pending security reviews. Unlike GLM-5.2, API access + weights staged behind safety review, not shipped day one.
- Crowded launch week: Google shipped Gemini 3.7 Flash (Aug 13); competing coding updates from DeepSeek (V4-Pro) and Alibaba (Qwen).

## 3.6 GLM-5.3 — architecture & the cyber angle
- Same MoE base as 5.2 (~743B total / ~40B active class).
- **Post-training focused on coding + cybersecurity:** trained with data and environments built specifically to find software vulnerabilities → **"emergent cyber capability"** that arrived faster than the team expected.
- **2,436 vulnerabilities uncovered across 269 projects during testing** (some flaws up to 40 years old), per Zhipu.

## 3.7 GLM-5.3 — benchmarks (all vendor-published — treat as claims, not independent verification)
- **+50% coding capability vs GLM-5.2** (Zhipu's in-house coding benchmark).
- **Terminal-Bench 3.0: 28.3**; **SWE-Marathon: 42.5** — claimed open-source SOTA.
- **Agents' Last Exam (CLI): open-source SOTA** (claimed).
- **CyberGym: 84.5** — tops the vulnerability-discovery benchmark, edging Anthropic's Mythos 5 in Zhipu's headline comparison; **>2× GLM-5.2 on exploitation benchmarks**.
- Zhipu calls it **"the strongest open-weights coding model available today"** (own benchmarks).
- Direct competition: DeepSeek V4-Pro, Alibaba Qwen coding updates.

---

## 3.8 GLM-5.3-Flash ("FlashX") — release (26 August 2026)
- **Launched 26–27 August 2026** — first **natively multimodal** model in the GLM-5 series (image + video inputs from the architecture level up, not bolted on).
- **Stealth preview as "Ox Alpha":** anonymised model briefly **topped OpenRouter's usage charts in mid-August** before attribution; community forensics (tokenizer match, Z.AI error codes, Java stack trace, Aug 21–22) pointed to Zhipu before Z.ai/Bloomberg/Caixin/CNBC confirmed.
- **Runs entirely on ~100,000 domestically manufactured Chinese AI accelerators** — no Nvidia hardware; architected for deployment across domestic Chinese GPU families. Pointed rebuttal to the narrative that SOTA Chinese AI depends on Nvidia; major signal on semiconductor self-sufficiency under US export controls.
- **Z.ai shares gained ~8% on the announcement.**

## 3.9 GLM-5.3-Flash — architecture
- **MoE: 320B total / 18B active per token.**
- **Hybrid sparse + linear attention** — design choice cutting compute significantly.
- **Context: 1M tokens.**
- **Natively multimodal:** native image + video input; text output.
- **MIT open-weights license** — weights at `huggingface.co/zai-org/GLM-5.3-Flash`; local inference via SGLang, vLLM, TokenSpeed.
- Post-trained on the 743B-parameter base (per explainx summary of launch materials).


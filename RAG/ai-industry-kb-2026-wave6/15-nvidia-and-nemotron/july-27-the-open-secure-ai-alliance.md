---
id: ai-industry-kb-2026-wave6/15-nvidia-and-nemotron/july-27-the-open-secure-ai-alliance
title: "July 27 — the Open Secure AI Alliance"
domain: nvidia-and-nemotron
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Anthropic", "Google", "Hugging Face", "Microsoft", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "Unsloth"]
dates: ["2026-07-27", "2026-08-11"]
keywords: ["agent", "agentic", "attention", "benchmarks", "blackwell", "claude", "cost", "disclosure", "distillation", "distribution", "embedding", "gguf"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [7417, 7459]
section: "§15. NVIDIA and Nemotron"
delta_of: ai-industry-kb-2026
sha256: 8f07e15990ac5515d9cbf96062669ebbcafad8e919d6ee4e1706a9621ace4352
---

# July 27 — the Open Secure AI Alliance

### July 27 — the Open Secure AI Alliance
- **2026-07-27** — NVIDIA launched the **Open Secure AI Alliance** with **30+ founding tech companies**, triggered by the July 21 breach disclosure [SECONDARY].
- Goals: develop and share **open-source security tools, evaluation frameworks, benchmarks, and best practices** for increasingly autonomous AI systems — inspectable, modifiable, runnable on teams' own infrastructure [SECONDARY].
- NVIDIA's stated position: blanket restrictions on open frontier AI would weaken defensive capacity and concentrate power/dependence/vulnerability in a few closed providers [SECONDARY].
- NVIDIA contributes open models, weights, data, and agent-harness research, including the open-source **Nvidia Labs Object-Oriented Agent** project on GitHub [SECONDARY].
- Per-outlet member lists vary (Reuters: Adobe, CrowdStrike, HF, Dell; techxplore: Microsoft, IBM, Palantir, CrowdStrike, Cisco, Dell, HF; InsideAI: Microsoft, IBM, Cisco, Cloudflare, HF) — **treat any single list as partial**; notably absent: Google, OpenAI, Anthropic, with Anthropic described as "the most visible holdout" [SECONDARY].
- (The July 24 letter and July 27 alliance are §21-owned coalition events; the NVIDIA-actor framing is the delta here.)


### New verified facts — expansion

### Nemotron 3.5 Lightning — release, sizing, license
- NVIDIA released Nemotron 3.5 Lightning on approximately August 11, 2026, announcing it jointly with the NeMo Switchyard router; press coverage dates the launch to August 11–12, 2026. [SECONDARY] [S1][S2]
- The model is described by secondary reporting as a 30B-total / 3B-active-parameter open MoE with 1M-token context, released under the OpenMDW-1.1 license with commercial use permitted. [SECONDARY] [S1][S49]
- CONTRADICTION — vendor-adjacent sizing: other NVIDIA material describes the same Lightning model at 31.6B total / ~3.6B active parameters. Preserve both: 30B/3B per secondary reporting vs 31.6B/~3.6B per NVIDIA material. Do not average them. [SECONDARY] [S1][S6]
- The architecture is a hybrid of Mamba-2 state-space blocks, sparse MoE feed-forward layers and attention, with a Multi-Token Prediction (MTP) head. [VENDOR] [S1][S6]
- Speed is the headline claim: up to 4x output speed, attributed to multi-token prediction plus "DSpark" and "DFlash" drafters, and an NVFP4 checkpoint. [VENDOR] [S1][S49]
- On the vendor-cited PinchBench (10,000 tasks), Lightning completed tasks 30% faster than Qwen3.6 35B. [VENDOR] [S1][S49]
- Recommended sampling is temperature 1.0 and top_p 0.95. (single secondary coverage) [SECONDARY] [S1]
- Deployment targets: a single DGX Spark or a single H100; community routes include Ollama, LM Studio, llama.cpp and Unsloth; distribution via build.nvidia.com, OpenRouter, Hugging Face and ModelScope. [SECONDARY] [S1][S49]
- The ggml-org community build names it NVIDIA-Nemotron-3.5-Lightning-30B-A3B-GGUF, targeting CPU and low-VRAM local inference. (single community coverage) [COMMUNITY] [S7]
- The text-only Nemotron chat checkpoints (Ultra, Super, Nano, 3.5 Lightning) are text-in/text-out only; multimodality lives in the separate Omni variant. (single secondary coverage) [SECONDARY] [S8]

### NVFP4 — the 4-bit production path
- NVFP4 is NVIDIA's 4-bit floating-point format using E2M1 encoding (2 exponent bits, 1 mantissa bit), built for Blackwell-era 4-bit tensor-core inference. [VENDOR] [S9][S50]
- Quantization uses blocks of 16 elements with dual scaling: a per-block float32 scale plus a coarser per-tensor scale. [VENDOR] [S9][S50]
- The reference Nemotron-3-Ultra-NVFP4 checkpoint is produced by quantization-aware distillation (QAD): a higher-precision teacher distills into the NVFP4 student, rather than pure post-training quantization. (single vendor coverage) [VENDOR] [S9]
- NVFP4 documents a precision-exception path for values unrepresentable in E2M1, so out-of-range activations do not silently clip. (single vendor coverage) [VENDOR] [S9]
- Lightning's speed story explicitly includes its NVFP4 checkpoint — the 4-bit format is part of the serving recipe, not an afterthought. (single secondary coverage) [SECONDARY] [S1]
- The Nemotron-3-Embed-8B-BF16 Hugging Face card is a separate 8B-class embedding checkpoint in the Nemotron 3 family, not the Lightning LLM. (single vendor coverage) [VENDOR] [S10]

### NeMo Switchyard — what it actually is (CORRECTED)
- NeMo Switchyard is an open-source LLM router library announced alongside Nemotron 3.5 Lightning; it routes each LLM call to the cheapest model that can still do the job. [VENDOR] [S2][S3]
- CORRECTION to earlier drafts: Switchyard is NOT label-only routing. It ships tuning-free routers including an LLM classifier with session affinity, a stage router that reads recent tool activity, and an escalation router that starts cheap and promotes on sustained difficulty. [VENDOR] [S1][S51]
- A tunable prefill router learns from the model's residual stream to predict which candidate model will succeed. (single secondary coverage) [SECONDARY] [S1]
- The reference server accepts OpenAI, Anthropic and Responses API requests, preserving native API compatibility. (single vendor coverage) [VENDOR] [S3]
- Implementation: written in Rust with Python bindings. Core crates are `libsy` (routing algorithms, entry point `Algorithm::run_stream`), `libsy-llm-client` (HTTP client driving the stream), `switchyard-runner` (TOML config parsing), `switchyard-server` (thin HTTP demo server), `switchyard-py` (Python bindings); `switchyard-translation` converts between OpenAI Chat Completions, OpenAI Responses and Anthropic Messages via a vendor-neutral IR. (single vendor coverage) [VENDOR] [S4]
- Install path: `pip install nemo-switchyard`, or embed via the Rust crate `switchyard-libsy`; it can also run as a NeMo Relay plugin (loading a `routes.toml`) or a LiteLLM Router/proxy plugin. [VENDOR] [S3][S5]
- Published result 1 (LangChain, 145 multi-turn agentic tasks): routing between Lightning and Claude Opus 4.8 with the escalation router cut cost 74% versus a frontier-only baseline, sending only 7% of calls to the frontier model, at roughly a 6-point accuracy tradeoff. [VENDOR] [S1][S2]
- Published result 2 (Cognition, staged routing in Devin Desktop, FrontierCode Main): routing between Opus 5 and Kimi K2.7 reached 50.6% at $3.11 mean cost — within 2.8 points of Opus 5 accuracy at approximately 28% lower mean cost. (single vendor coverage) [VENDOR] [S1]
- Published result 3 (Switchyard's own chart, Terminal-Bench 2.1): staged, escalation and classifier routes reach 71–76% accuracy for 13–30% less than the Opus 4.8 baseline, while single fixed models stay below 56%. Cost basis is average ISP token cost. [VENDOR] [S3][S5]
- Limitation to note: Switchyard picks the model but the harness makes the call — transport, retries and credentials stay with the operator's gateway. It does not replace the serving layer. (single vendor coverage) [VENDOR] [S3]


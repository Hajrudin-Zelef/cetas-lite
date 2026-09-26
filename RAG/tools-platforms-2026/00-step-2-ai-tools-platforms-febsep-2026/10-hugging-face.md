---
id: tools-platforms-2026/00-step-2-ai-tools-platforms-febsep-2026/10-hugging-face
title: "10. Hugging Face"
domain: step-2-ai-tools-platforms-febsep-2026
role: deep-dive
task: platform
actors: ["AMD", "AWS", "Alibaba", "Anthropic", "Apple", "Baseten", "Cerebras", "China", "DeepSeek", "Hugging Face", "Meta", "Microsoft", "Moonshot", "Nvidia", "OpenAI", "United States", "Z.ai", "vLLM"]
dates: ["2026-06-04", "2026-06-17", "2026-06-30", "2026-07", "2026-09", "2026-09-03"]
keywords: ["accelerator", "acquisition", "agent", "agentic", "agents", "apache", "benchmark", "benchmarks", "blackwell", "claude", "compute", "deepseek"]
source: docs/RAG/Outils & plateformes IAEN.md
source_anchor: ""
source_lines: [477, 526]
section: "Step 2 — AI Tools & Platforms (Feb–Sep 2026)"
sha256: 00fa148a0c56519c30483f224d4bb6592158f09304a9289efe4575671b0da577
---

# 10. Hugging Face

## 10. Hugging Face

### 10.1 Platform scale (2026)
- 18M+ developers, researchers, and creators; 200,000+ companies; 3M+ models; 500,000 datasets; 1M+ applications (Spaces). **[official — NVIDIA blog citing HF figures, Sept 3, 2026]**
- Growth Jan–Aug 2026: public model repos 2.43M → 2.96M (+21.5%); datasets 711K → 1M; Spaces 1.00M → 1.44M. Distribution highly skewed: ~85.6% of models have <200 lifetime downloads; 1.5% of repos account for 99.2% of downloads. **[official — HF "State of Open Models: Summer 2026", Aug 14, 2026]**
- Runtime layer growing fastest: repos declaring `gguf` +464%, `lerobot` +194%, `mlx` +148% vs +16% for `transformers`/`peft`. **[official — same report]**
- GGUF local-inference traffic: Qwen 39.6M GGUF downloads/month (vs Gemma 20.8M, Llama 7.5M); Qwen-based derivatives 151,448 (2.6× Meta's footprint). **[official — same report]**
- ~$150M annualized revenue (2026) — press-reported, not a company disclosure. **[secondary — cryptobriefing/Bloomberg Opinion]** ⚠️ estimate only.

### 10.2 NVIDIA acquisition — announced September 3, 2026 (pending close)
- NVIDIA signed a **definitive agreement on Sept 2, 2026** to acquire Hugging Face for **$12,930,300,000**; announced by CEO Jensen Huang in an NVIDIA blog post on Sept 3, 2026. Deal expected to close in **H1 2027**, subject to regulatory approvals and customary conditions. **[official — NVIDIA blog + SEC 8-K filing]**
- Structure: ~$11.9B payable to HF stockholders + up to ~$1.0B equity-based retention program for HF employees joining NVIDIA. **[official — 8-K]**
- Platform commitment: HF "will remain an open platform" — developers keep choosing their own models, frameworks, clouds, and compute providers; "NVIDIA compute will not be required"; multi-cloud and multi-accelerator support continues. **[official — NVIDIA blog]**
- HF's three co-founders (Clément Delangue, Julien Chaumond, Thomas Wolf) join NVIDIA with the whole team; per Bloomberg Billionaires Index each founder worth ~$1.8B. **[independent — Bloomberg]**
- Context: The Information, Bloomberg, and TechCrunch reported the deal under negotiation Aug 26–27, 2026 before confirmation. HF previously rejected a ~$500M NVIDIA investment (late 2025) that would have valued it at ~$7B. Multiple bidders were reported; NVIDIA moved after Delangue approached Huang directly. **[independent — TechCrunch/Bloomberg, via secondary]**
- This is NVIDIA's largest outright company acquisition (Mellanox: $6.9B, 2020). **[secondary — temperature2 editorial]**

### 10.3 Platform features & launches (June–September 2026)

**Hub & inference**
- **Baseten joins Inference Providers** (Aug 6, 2026): conversational/text-gen tasks at launch (Kimi K3, DeepSeek V4 Flash, GLM-5.2, etc.); more tasks to follow. Integration via `huggingface_hub` ≥1.26.1, `@huggingface/inference` JS, OpenAI-compatible router at `router.huggingface.co/v1` with `model:provider` syntax. Two billing modes: custom key (billed by provider) or routed-by-HF (billed on HF account, provider rates passed through with **no HF markup**). PRO users get $2/month Inference credits. **[official — https://huggingface.co/blog/baseten]**
- **Inference Providers in agent harnesses**: integrated with Pi, OpenCode, Hermes Agents, OpenClaw, etc. **[official — Baseten blog post]**
- **Agentic Resource Discovery launch** (June 17, 2026): let agents search Hub resources. **[official — https://huggingface.co/blog/agentic-resource-discovery-launch]**
- **Model pages show Every Eval Ever results** (June 30, 2026): community eval results featured on model pages. **[official — https://huggingface.co/blog/eee-community-evals]**
- **hf CLI redesigned as agent-optimized Hub interface** (June 4, 2026). **[official — https://huggingface.co/blog/hf-cli-for-agents]**
- **agent-usage dataset** (published July 2026): tracks `agent/<name>` tokens from coding agents calling the Hub — Claude Code led July with 44.4%, Codex rose 10.4%→20.8% Apr–Jul. **[official — State of Open Models report]**
- **Foundry Managed Compute** (Microsoft, July 7), **one-click to Amazon SageMaker Studio** (July 7), **zero-egress storage with SkyPilot** (July 7) — compute/storage integrations. **[official blog]**
- **Cerebras + Gemma 4 for real-time voice AI** (July 1). **[official blog]**
- **Funes** (Sept 3, 2026): "Give Your Coding Agents a Memory You Own". **[official blog]**

**Spaces & compute**
- **Spaces Dev Mode** (PRO): SSH/VS Code access into Spaces for fast iteration. **[official pricing/docs — via secondary guides; verify on huggingface.co/pricing]**
- **ZeroGPU**: shared GPU pool (RTX Pro 6000 Blackwell) giving free GPU-backed Spaces; PRO raises quota (third-party guides cite 25–40 min/day with over-quota credits). **[secondary]** ⚠️ exact quota varies across third-party guides; verify on official pricing page.
- **HF Jobs**: run training/jobs on HF infra — "Run a vLLM Server on HF Jobs in One Command" (June 26), migrate GitHub CI to HF Jobs (June 9), Async GRPO with LoRA across HF Jobs (Sept 10). **[official blog]**
- **Storage Buckets**: powering Papers with Code search + LeRobot/Strands Agents data loops (Aug 2026 posts). **[official blog]**
- **Gradio Workflow**: new workflow system posts Aug 25, Sept 10. **[official blog]**

**Browser / local AI**
- **@huggingface/kernels** (Sept 1, 2026): 207 WebGPU kernels as versioned Hub repos (Apache-2.0), JS loader, plus **Fleet** — in-browser crowdsourced GPU benchmarking. 2.57× geomean faster than ORT WebGPU on Apple M4 (vendor benchmark). Kernels page also covers CUDA/ROCm/Metal. **[official — https://huggingface.co/blog/webgpu-kernels]**
- **llama.cpp/ggml team joined Hugging Face** (Feb 2026); project stays open-source and community-governed. **Transformers now runs llama.cpp quants** (Sept 22, 2026). **[official — State of Open Models report; blog]**
- **tokenizers v1** (Sept 21, 2026). **Kernels major updates** (July 6). **[official blog]**
- **Jun Kim (oMLX creator) joins HF** to support the MLX community (Sept 22, 2026). **[official blog]**
- **oMLX, Nunchaku 4-bit diffusion in Diffusers** (July 23). **[official blog]**

**Ecosystem / research posts**
- **State of Open Models: Summer 2026** (Aug 14, 2026): flagship data report — Chinese labs lead the open frontier ceiling (754B–2.78T params/month vs US mostly <130B); Qwen became the community base model; licenses stay permissive (59% Apache-2.0 among Chinese >20B releases) though Kimi K3/Qwen 3.8 2.4T recently added non-commercial/revenue-share terms; small models take 83% of all-time downloads. **[official — https://huggingface.co/blog/state-of-open-models-summer-2026]**
- **UK AISI × EvalEval reproducible benchmarks** (Sept 22), **BenchMIRT** (AllenAI, Sept 1), **Open ASR Leaderboard global-south language** (Aug 28), **Real World VoiceEQ** (July 15), **FFASR leaderboard** (June 24). **[official blog]**
- **Meta Muse Glimmer** spotlight (Aug 10): local, agentic, multimodal, open source. **Thinking Machines Inkling** (July 15), **GLM-5.2** (Z.ai, June 17). **[official blog]**
- **LeRobot v0.6.0** (July 7); Strands Agents + LeRobot robotics posts (June 17, Aug 13). **Grabette** robot data (July 21). **[official blog]**


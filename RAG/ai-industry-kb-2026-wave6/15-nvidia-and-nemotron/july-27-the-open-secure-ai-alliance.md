---
id: ai-industry-kb-2026-wave6/15-nvidia-and-nemotron/july-27-the-open-secure-ai-alliance
title: "July 27 — the Open Secure AI Alliance"
domain: nvidia-and-nemotron
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Anthropic", "Google", "Hugging Face", "Microsoft", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "Unsloth", "vLLM"]
dates: ["2026-05-31", "2026-06", "2026-06-01", "2026-07-15", "2026-07-20", "2026-07-27", "2026-08-11", "2026-08-13"]
keywords: ["agent", "agentic", "attention", "attribution", "benchmarks", "blackwell", "chatgpt", "claude", "compute", "cost", "datacenter", "diffusion"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [7417, 7501]
section: "§15. NVIDIA and Nemotron"
delta_of: ai-industry-kb-2026
sha256: df3b946e9df0b9346cf0fde8843adbfb872f766d636bb8c5822378c68b690e92
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

### Cosmos 3 — two-tower Mixture-of-Transformers world models
- NVIDIA announced Cosmos 3 at its GTC Taipei event on June 1, 2026 (one research note dates the announcement ~May 31, 2026 — preserve the one-day discrepancy). [SECONDARY] [S11][S12][S13]
- Cosmos 3 is a family of omnimodal world foundation models for physical AI, unifying three capabilities that earlier Cosmos generations split across separate models: physical reasoning, world generation and action generation. [VENDOR] [S11][S13]
- Architecture: a two-tower Mixture-of-Transformers (MoT). The Reasoner tower is an autoregressive vision-language model (initialized from Qwen3-VL) that interprets text/image/video/audio and builds a world-state representation; the Generator tower is diffusion-based and produces physics-aware video, audio and action sequences conditioned on the reasoner's understanding. [VENDOR] [S11][S13]
- The towers have separate parameters per layer but interact via joint attention; information flows from reasoner to generator, and the reasoner can run independently while the generator always activates both towers. (single vendor coverage) [VENDOR] [S13]
- Two launch sizes: Cosmos 3 Nano = 16B total (8B reasoner + 8B generator), aimed at workstation inference (e.g., NVIDIA RTX PRO 6000); Cosmos 3 Super = 64B total (32B + 32B), aimed at datacenter deployment on Hopper/Blackwell for large-scale synthetic data and advanced physical reasoning. [VENDOR] [S11][S13]
- Cosmos 3 Edge = 4B on-device world model, launched July 15, 2026 with weights on Hugging Face July 20, 2026 (`nvidia/Cosmos3-Edge`); it targets Jetson Thor for real-time robot control at 15 Hz (640×360 resolution, 32 actions per inference). [SECONDARY] [S14][S15]
- NVIDIA claims Cosmos 3 Edge ranks #1 on VANTAGE-Bench for vision analytics among 4B-parameter models and state-of-the-art robot policy learning in its size class; specific scores were not disclosed. (single vendor coverage) [VENDOR] [S15]
- The Edge reasoner uses a 2B dense transformer and Qwen3-VL-compatible message conventions; reasoner context reaches up to 256K tokens. [SECONDARY] [S14][S16]
- Robotics operating modes: forward dynamics (action+image+text → future video, i.e., a learned simulator), inverse dynamics (text+video → action), and policy (image+text → video+action); the model emits numerical actions such as joint angles, gripper positions and trajectories. (single secondary coverage) [SECONDARY] [S16]
- Licensing: open weights under OpenMDW-1.1 (commercial use permitted, "Built on NVIDIA Cosmos" attribution); training scripts, deployment tools and datasets also open-sourced; code at github.com/nvidia/cosmos. [SECONDARY] [S13][S16]
- Access routes: Hugging Face (`nvidia/Cosmos3-Nano`, `nvidia/Cosmos3-Super`, `nvidia/Cosmos3-Edge`), build.nvidia.com (preview-gated), Cosmos NIM microservices (Reasoner NIM GA, Generator NIM v1.0.0), self-serve via vLLM-Omni, and third-party hosting (e.g., DeepInfra lists Nano at $0.0108/second). [SECONDARY] [S16][S53]
- A ready VLA-style policy fine-tune, Cosmos3-Nano-Policy-DROID (16B), was released alongside, trained on the DROID dataset. [SECONDARY] [S16][S54]
- CONTRADICTION (minor): announcement dated May 31, 2026 in one research note vs June 1, 2026 (GTC Taipei) in event records. Preserve both. [SECONDARY] [S12][S16]

### GR00T N1.7 and N2 — humanoid robot foundation models
- NVIDIA announced GR00T N1.7 in early access with commercial licensing, bringing generalized robot skills including advanced dexterous control to production-ready deployments. [VENDOR] [S17][S55]
- Industrial adopters named for the Isaac GR00T N line: AGIBOT, Humanoid, LG Electronics, NEURA Robotics and Noble Machines. [SECONDARY] [S17][S56]
- At the GTC keynote, Jensen Huang previewed GR00T N2, a next-generation robot foundation model based on DreamZero research with a new world-action-model architecture. [VENDOR] [S17][S56]
- NVIDIA's claim: N2 helps robots succeed at new tasks in new environments more than twice as often as leading vision-language-action models; it ranks No. 1 on MolmoSpaces and RoboArena for generalist robot policies. [VENDOR] [S17][S56]
- N2 was slated to be available by the end of the year (2026) at preview time — it remains a previewed, not generally available, model. [SECONDARY] [S17][S56]
- The GR00T stack sits on the Newton physics engine 1.0 plus the PhysX SDK (multiphysics simulation, dexterous manipulation) and the Jetson Thor robotic computing platform for sim-to-real deployment. (single secondary coverage) [SECONDARY] [S17]
- Separately, NVIDIA announced the Isaac GR00T Reference Humanoid Robot (June 1, 2026): a Unitree H2/H2 Plus chassis (~6 ft, 150 lb, 31 DOF) with Sharpa Wave tactile five-finger hands (22 DOF, 75 DOF total), head-mounted stereo camera (140° horizontal / 102° vertical FOV), wrist cameras, whole-body control with up to 120 Nm arm torque and 360 Nm leg torque, 7 kg rated / 15 kg peak arm payload, and Jetson Thor onboard compute. (single secondary coverage) [SECONDARY] [S18]
- LG Group and NVIDIA signed an MoU on August 13, 2026 (Santa Clara) to jointly develop a next-generation bipedal humanoid robot, targeting public unveiling in Q1 2027: Jetson Thor onboard compute, Isaac GR00T foundation model, and NVIDIA Halos for Robotics safety stack. LG Electronics supplies actuators, LG Innotek sensors, LG Energy Solution batteries — while LG also develops its own in-house robot foundation model (RFM) fed by manufacturing validation data. (single secondary coverage) [SECONDARY] [S19]
- The LG–NVIDIA scope extends beyond the robot: an AI factory pilot site powered by NVIDIA's Vera Rubin platform in H1 2027 (paving the way for an 80MW AI factory in Cheonan, South Korea, by H1 2028), CLOiD wheel-based robots on LG's Tennessee washing-machine line for real-world validation, a robot data platform with LG CNS, and an AI-defined vehicle platform on NVIDIA DRIVE Hyperion. (single secondary coverage) [SECONDARY] [S19]

### Alpamayo — driving world models (CORRECTED LINEAGE)
- CORRECTION: there is no verified "Alpamayo 1.5" release in the collected material. The verified lineage is Alpamayo 1 (CES 2026) → Alpamayo 2 Super (June 2026). Do not cite a 1.5 version. [SECONDARY] [S20][S21]
- NVIDIA launched the Alpamayo family at CES 2026 (January 5–6, 2026): open AI models, simulation tools and datasets for autonomous-vehicle reasoning, pitched by Jensen Huang as "the ChatGPT moment for physical AI." [SECONDARY] [S20][S22]
- Alpamayo 1 is a 10B-parameter chain-of-thought reasoning vision-language-action (VLA) model; it breaks driving problems into steps, reasons through possibilities and selects the safest path, emitting trajectories plus human-readable explanations. [SECONDARY] [S20][S22]
- The CES launch bundle: Alpamayo 1 (10B, open weights + inference scripts on Hugging Face), AlpaSim (open-source simulation framework on GitHub), and 1,700+ hours of real-world driving datasets on Hugging Face. (single secondary coverage) [SECONDARY] [S22]
- Named early adopters: Lucid, JLR and Uber, using it to accelerate Level 4 deployment; Alpamayo serves as a teacher model distilled into smaller in-vehicle runtimes. (single secondary coverage) [SECONDARY] [S22]
- Alpamayo 2 Super: a 34B-parameter VLA for robotaxis/autonomous driving, pairing a 32B VLM backbone (built on the Cosmos 3 Super Reasoner, post-trained with RL) with a 2.3B diffusion-based action decoder. [SECONDARY] [S21][S57]
- CONTRADICTION: one aggregator describes Alpamayo 2 Super as a "32-billion-parameter" model while NVIDIA-focused reporting says 34B (32B backbone + 2.3B decoder ≈ 34B). Preserve both; the 34B figure is better sourced. [SECONDARY] [S21][S23]
- From one pass over full-surround camera video it emits a planned trajectory, a causal explanation of that trajectory, and a meta-action; the trajectory API returns 64 waypoints spanning 0.1–6.4 seconds at 0.1 s intervals. [SECONDARY] [S21][S57]
- Inputs: multi-camera RGB video, text, and egomotion history with timestamps (validated notebook: six cameras, four historical frames per camera; egomotion as 3D translation + 3×3 rotation, multi-timestep); some reporting says up to seven cameras. [SECONDARY] [S21][S24]
- Training data: roughly 115,000 hours of multi-camera driving video with egomotion/trajectory annotations, about 3.7M Chain-of-Causation (CoC) traces, and over 1B images. [SECONDARY] [S21][S57]
- NVIDIA reports a LingoQA score of 79.2, ranking Alpamayo 2 Super first among 37 evaluated models. (single vendor coverage) [VENDOR] [S24]
- License: OpenMDW-1.1 (Linux Foundation permissive license), covering fine-tuning, derivatives and commercial redistribution; NVIDIA is extending OpenMDW across the whole Alpamayo family so earlier R&D-positioned releases become commercially deployable. [SECONDARY] [S21][S24]
- Announced June 1, 2026, with availability expected summer 2026 on GitHub and Hugging Face; an open-source CoC Auto-Labeling Pipeline was also announced for annotation-free causal label generation. [SECONDARY] [S21][S23]
- Prior Alpamayo versions accumulated 400,000+ downloads (per NVIDIA via secondary). (single vendor coverage) [VENDOR] [S23]


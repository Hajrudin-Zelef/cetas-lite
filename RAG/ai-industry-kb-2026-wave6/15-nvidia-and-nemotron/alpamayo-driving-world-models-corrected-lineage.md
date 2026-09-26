---
id: ai-industry-kb-2026-wave6/15-nvidia-and-nemotron/alpamayo-driving-world-models-corrected-lineage
title: "Alpamayo — driving world models (CORRECTED LINEAGE)"
domain: nvidia-and-nemotron
role: deep-dive
task: actor-profile
actors: ["Hugging Face", "Nvidia", "OpenAI", "SGLang", "vLLM"]
dates: ["2025-12-19", "2026-06", "2026-06-01", "2026-08-03"]
keywords: ["chatgpt", "cost", "diffusion", "fine-tuning", "fp8", "full-duplex", "gpu", "inference", "latency", "license", "llama", "multimodal"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [7486, 7519]
section: "§15. NVIDIA and Nemotron"
delta_of: ai-industry-kb-2026
sha256: 9c51a2c1172fe0a5f324041c3f5ada83f6c685c2a98f010d59fa47cbe903a8cf
---

# Alpamayo — driving world models (CORRECTED LINEAGE)

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

### Nemotron VoiceChat 11B and Nemotron-3-Nano-Omni — speech and omni models
- On August 3, 2026, NVIDIA published NemotronLabs VoiceChat 11B on Hugging Face under OpenMDW 1.1: an 11B-parameter, end-to-end, real-time full-duplex speech-to-speech model for conversational AI (checkpoint `nvidia/NVIDIA-NemotronLabs-VoiceChat-11B`). [SECONDARY] [S25][S26]
- It is a single unified model — raw audio in, speech out — listening and speaking simultaneously, handling interruptions and calling tools mid-conversation in one forward pass with no ASR→LLM→TTS cascade. [SECONDARY] [S25]
- Claimed performance: ~450 ms turn-taking latency, full-duplex operation, barge-in support, and the first open full-duplex model with tool calling. (single vendor coverage) [VENDOR] [S25]
- Architecture (based on the SALM-Duplex paper, Interspeech 2025), frame-locked at 12.5 Hz (one 80 ms frame per step): a Fast Conformer perception encoder + Nemotron-H hybrid-Mamba thinker emitting one text token and one function token per frame; a Gemma3-1B EAR-TTS talker converting each text token into 31-quantizer RVQ code stacks; an RVQ-VAE codec decoding to 22.05 kHz audio. The model decides when to speak — no external VAD. [SECONDARY] [S26][S27]
- The checkpoint is a single 44 GB fp32 `model.safetensors` with a NeMo-style config; its tokenizer/config come from the sibling repo `nvidia/NVIDIA-Nemotron-Nano-9B-v2`. [SECONDARY] [S27]
- Serving: offline single-turn inference via vLLM-Omni and SGLang-Omni recipes (3-stage pipeline: thinker → talker → code2wav); experimental native duplex realtime serving is tracked but not yet the default path. [SECONDARY] [S26][S27][S28]
- Measured footprint: on one H200 the offline recipe peaks around 129 GB (talker engine reserves ~47 GB KV pool, thinker ~73 GB including 17.7 GB bf16 weights); one GPU hosts all four stages. (single secondary coverage) [SECONDARY] [S27]
- Nemotron-3-Nano-Omni-30B-A3B is the multimodal chat model of the line: video, audio, image and text in, text out. It is the only multimodal chat model in the Nemotron line; Ultra/Super/Nano/3.5 Lightning are text-only. [SECONDARY] [S8][S29]
- A community serving recipe runs the Omni model in FP8 at ~33 GB, fitting one L40S 48 GB GPU, with `--max-model-len 32768` and `enable_thinking=false` for low-latency voice chat. (single community coverage) [COMMUNITY] [S29]
- The broader Nemotron catalog separates modalities into distinct models: `nemotron-asr-streaming`, `nemotron-voicechat`, `nemotron-ocr-v1`/`v2`, `nemotron-parse`, `nemotron-page-elements-v3`, `nemotron-table-structure-v1`, `nemotron-graphic-elements-v1`, `llama-nemotron-embed-vl-1b-v2`, `llama-nemotron-rerank-vl-1b-v2`, plus safety models (`nemotron-3.5-content-safety`, `llama-3_1-nemotron-safety-guard-8b-v3`, etc.). A multimodal Nemotron system is a pipeline of models, each hop its own latency/cost/failure surface. (single secondary coverage) [SECONDARY] [S8]

### Download-metric correction
- The "50M+ downloads" figure sometimes attached to the Nemotron family is UNVERIFIED and conflates two metrics. (single unverified coverage) [UNVERIFIED] [S30]
- The verified NVIDIA developer-forum figure is 5M Hugging Face downloads, posted 2025-12-19. (single secondary coverage) [SECONDARY] [S30]
- Where "50M+" appears in NVIDIA-adjacent material it refers to training samples, not downloads — the two must not be merged. (single secondary coverage) [SECONDARY] [S30]



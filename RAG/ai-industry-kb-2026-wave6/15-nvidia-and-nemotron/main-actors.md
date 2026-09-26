---
id: ai-industry-kb-2026-wave6/15-nvidia-and-nemotron/main-actors
title: "Main actors"
domain: nvidia-and-nemotron
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Anthropic", "China", "Groq", "Nvidia", "Poolside", "United States", "Z.ai"]
dates: ["2025-08-25", "2025-12-19", "2026-03-11"]
keywords: ["acquisition", "agent", "compute", "cost", "distillation", "energy", "fp4", "fp8", "full-duplex", "glm", "inference", "latency"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [7571, 7658]
section: "§15. NVIDIA and Nemotron"
delta_of: ai-industry-kb-2026
sha256: 06680b540227500ea656a13b3971b5fbb59b9d1e79eba933e19a29ab7b7d30b4
---

# Main actors

| Model / artifact | Metric | Value | Label | Sources |
|---|---|---|---|---|
| Nemotron 3.5 Lightning | Total params (secondary) | 30B | [SECONDARY] | S1 |
| Nemotron 3.5 Lightning | Active params (secondary) | 3B | [SECONDARY] | S1 |
| Nemotron 3.5 Lightning | Total params (vendor-adjacent) | 31.6B | [VENDOR] | S6 |
| Nemotron 3.5 Lightning | Active params (vendor-adjacent) | ~3.6B | [VENDOR] | S6 |
| Nemotron 3.5 Lightning | Context | 1M tokens | [SECONDARY] | S1 |
| Nemotron 3.5 Lightning | License | OpenMDW-1.1, commercial OK | [SECONDARY] | S1 |
| Nemotron 3.5 Lightning | Output speed claim | Up to 4x | [SECONDARY] | S1 |
| Nemotron 3.5 Lightning | PinchBench vs Qwen3.6 35B | 30% faster (10k tasks) | [VENDOR] | S1 |
| Nemotron 3.5 Lightning | Recommended sampling | temp 1.0, top_p 0.95 | [SECONDARY] | S1 |
| Nemotron 3.5 Lightning | Deployment floor | 1× DGX Spark or 1× H100 | [SECONDARY] | S1 |
| Switchyard | LangChain 145-task cost cut vs Opus 4.8 | −74% | [VENDOR] | S1, S2 |
| Switchyard | Frontier-call share (LangChain) | 7% to Opus 4.8 | [VENDOR] | S1 |
| Switchyard | Accuracy tradeoff (LangChain) | ~6 points | [VENDOR] | S1, S2 |
| Switchyard | FrontierCode Main (Cognition/Devin) | 50.6% @ $3.11 mean cost | [VENDOR] | S1 |
| Switchyard | Gap to Opus 5 (FrontierCode) | Within 2.8 pts, ~28% cheaper | [VENDOR] | S1 |
| Switchyard | Terminal-Bench 2.1 (own chart) | 71–76% @ 13–30% less than Opus 4.8 | [VENDOR] | S3, S5 |
| NVFP4 | Format | E2M1 (2 exp / 1 mantissa bits) | [VENDOR] | S9 |
| NVFP4 | Block size | 16 elements | [VENDOR] | S9 |
| NVFP4 | Scaling | Per-block FP32 + per-tensor dual scale | [VENDOR] | S9 |
| NVFP4 | Ultra checkpoint method | Quantization-aware distillation | [VENDOR] | S9 |
| Cosmos 3 Nano | Total params | 16B (8B reasoner + 8B generator) | [VENDOR] | S11, S13 |
| Cosmos 3 Super | Total params | 64B (32B + 32B) | [VENDOR] | S11, S13 |
| Cosmos 3 Edge | Total params | 4B | [SECONDARY] | S15, S16 |
| Cosmos 3 Edge | Control rate | 15 Hz, 32 actions/inference @ 640×360 | [SECONDARY] | S16 |
| Cosmos 3 | Reasoner context | Up to 256K tokens | [SECONDARY] | S15 |
| Cosmos 3 | License | OpenMDW-1.1, commercial OK | [SECONDARY] | S13, S17 |
| Cosmos 3 Nano | Third-party hosting (DeepInfra) | $0.0108/second | [SECONDARY] | S17 |
| Cosmos3-Nano-Policy-DROID | Size | 16B VLA fine-tune | [SECONDARY] | S17 |
| GR00T N1.7 | Availability | Early access + commercial license | [SECONDARY] | S18 |
| GR00T N2 | Availability | Previewed; slated end of 2026 | [SECONDARY] | S18 |
| GR00T N2 | Claimed success lift vs VLA models | >2x on new tasks/environments | [VENDOR] | S18 |
| Isaac GR00T Reference | Total DOF | 75 (31 body + hands) | [SECONDARY] | S19 |
| Isaac GR00T Reference | Arm / leg torque | 120 Nm / 360 Nm | [SECONDARY] | S19 |
| Isaac GR00T Reference | Arm payload | 7 kg rated / 15 kg peak | [SECONDARY] | S19 |
| Alpamayo 1 | Size | 10B reasoning VLA | [SECONDARY] | S21, S23 |
| Alpamayo 2 Super | Size | 34B (32B VLM + 2.3B action decoder) | [SECONDARY] | S22 |
| Alpamayo 2 Super | Training video | ~115,000 hours multi-camera | [SECONDARY] | S22 |
| Alpamayo 2 Super | CoC traces | ~3.7M | [SECONDARY] | S22 |
| Alpamayo 2 Super | Training images | >1B | [SECONDARY] | S22 |
| Alpamayo 2 Super | Trajectory output | 64 waypoints, 0.1–6.4 s @ 0.1 s | [SECONDARY] | S22 |
| Alpamayo 2 Super | LingoQA | 79.2, 1st of 37 models | [VENDOR] | S25 |
| Alpamayo family | Prior-version downloads | 400,000+ | [VENDOR] | S24 |
| NemotronLabs VoiceChat 11B | Params | 11B full-duplex S2S (single source) | [SECONDARY] | S26 |
| VoiceChat 11B | Frame lock | 12.5 Hz (80 ms frames) | [SECONDARY] | S27, S28 |
| VoiceChat 11B | Claimed turn-taking latency | ~450 ms | [VENDOR] | S26 |
| VoiceChat 11B | RVQ quantizers | 31 per frame | [SECONDARY] | S27 |
| VoiceChat 11B | Output sample rate | 22.05 kHz | [SECONDARY] | S27, S28 |
| VoiceChat 11B | Checkpoint size | 44 GB fp32 safetensors | [SECONDARY] | S28 |
| VoiceChat 11B | Offline peak VRAM (H200) | ~129 GB | [SECONDARY] | S28 |
| Nemotron-3-Nano-Omni | Size | 30B-A3B, audio/video/image/text → text | [SECONDARY] | S8, S30 |
| Nemotron-3-Nano-Omni | FP8 serving footprint | ~33 GB (fits L40S 48 GB) | [COMMUNITY] | S30 |
| Nemotron family (HF) | Verified downloads | 5M (2025-12-19) | [SECONDARY] | S31 |
| "50M+" figure | Correct referent | Training samples, NOT downloads | [SECONDARY] | S31 |


| Nemotron 3 Super | Release | 2026-03-11 (catalogs) / GTC 2026 | [SECONDARY] | S31, S32 |
| Nemotron 3 Super | Total / active params | 120B / ~12B (12.7B in some catalogs) | [SECONDARY] | S31, S32 |
| Nemotron 3 Super | Architecture | Mamba-2 + Transformer + LatentMoE + MTP | [SECONDARY] | S31, S32 |
| Nemotron 3 Super | Context | 1M tokens (91.75% RULER@1M; 256K default config) | [SECONDARY] | S32, S36 |
| Nemotron 3 Super | Speed (Artificial Analysis) | 449 output tok/s, #1 of 51 in tier, 2.2x GPT-OSS-120B | [SECONDARY] | S31 |
| Nemotron 3 Super | SWE-Bench Verified | 60.47% (vs GPT-OSS-120B 41.90%, vs Qwen3.5-122B 66.40) | [VENDOR] | S33, S36 |
| Nemotron 3 Super | MMLU-Pro / AIME 2025 / HMMT Feb25 / LiveCodeBench | 83.73% / 90.21% / 93.67% / 81.19% | [VENDOR] | S36 |
| Nemotron 3 Super | Training data | 15.6T tokens, 153 datasets, 20 langs, 43 prog langs | [SECONDARY] | S32 |
| Nemotron 3 Super | Hosted pricing (median) | $0.30 / $0.80 per 1M in/out; DeepInfra from $0.10/$0.50 | [SECONDARY] | S31 |
| Nemotron 3 Super | Min self-host hardware | 8x H100-80GB; NIM, 15+ providers | [SECONDARY] | S32 |
| Nemotron 3 Super | License (reported) | NVIDIA Nemotron Open Model License | [SECONDARY] | S32 |
| Nemotron 3 Super | Technical report | arXiv 2604.12374; research.nvidia.com PDF | [COMMUNITY] | S35 |
| Jetson Thor T5000 | AI compute / memory / power | 2,070 FP4 TFLOPS / 128GB / 130W | [SECONDARY] | S37, S38 |
| Jetson Thor T5000 | GA / module price / dev kit | 2025-08-25 / $2,999 / ~$3,499 | [SECONDARY] | S37, S39 |
| Jetson Thor T4000 | AI compute / memory | 1,200 FP4 TFLOPS / 64GB, shipping, $1,999 | [COMMUNITY] | S39 |
| Jetson Thor T3000 / T2000 | AI compute / memory / ship | 865 / 400 FP4 TFLOPS; 32/16GB; Q1 2027 | [COMMUNITY] | S39 |
| Jetson Thor vs AGX Orin | Compute / efficiency lift | 7.5x AI compute, 3.5x energy efficiency | [SECONDARY] | S37 |
| NVIDIA Dynamo | Claimed throughput lift | Up to 7x (multi-node) | [VENDOR] | S41 |
| NVIDIA Dynamo | Multimodal E/P/D TTFT improvement | 30% faster on image workloads | [COMMUNITY] | S42 |
| NVIDIA Dynamo | Video generation | Real-time 1080p on single B200 | [COMMUNITY] | S42 |
| Cosmos 1 (first gen) | Training video | 20M hours | [SECONDARY] | S46 |
| Cosmos-Predict2.5 | Scales | 2B and 14B; 200M curated clips + RL post-training | [COMMUNITY] | S45 |

## Main actors
- **NVIDIA** — author of both July coalition moves (letter co-host, alliance founder); the Nemotron open-model line; Nemotron Coalition; Object-Oriented Agent; the $6B Poolside license [SECONDARY/VENDOR].
- **Jensen Huang** — first-ever X post amplified the July 24 letter [SECONDARY].
- **Nemotron 3 Ultra** — the 550B flagship checkpoint: 90% sparsity, Mamba-2 + LatentMoE + NVFP4 [SECONDARY].
- **Groq** — Groq 3 LPX full production (Aug 24); DOJ "shadow acquisition" probe (Sept 11) [VENDOR/SECONDARY].
- **The Open Secure AI Alliance** — 30+ companies; open-source security tooling for autonomous AI [SECONDARY].
- **GLM 5.2 (Z.ai)** — the Chinese open-weight model that did HF's forensic work in the breach [SECONDARY].


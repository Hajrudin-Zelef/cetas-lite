---
id: ai-industry-kb-2026-wave6/15-nvidia-and-nemotron/nemotron-voicechat-11b-and-nemotron-3-nano-omni-speech-and-o
title: "Nemotron VoiceChat 11B and Nemotron-3-Nano-Omni — speech and omni models"
domain: nvidia-and-nemotron
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Anthropic", "China", "Groq", "Hugging Face", "Nvidia", "OpenRouter", "Poolside", "SGLang", "Samsung", "TensorRT-LLM", "United States", "Z.ai", "vLLM"]
dates: ["2025-06", "2025-08-25", "2025-12-15", "2025-12-19", "2026-02", "2026-03", "2026-03-11", "2026-03-16", "2026-06-01", "2026-06-04", "2026-07-15", "2026-08-03", "2026-08-11", "2026-08-24"]
keywords: ["omni", "voice", "accelerator", "acquisition", "agent", "agentic", "attention", "benchmarks", "blackwell", "compute", "context window", "cost"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [7502, 7650]
section: "§15. NVIDIA and Nemotron"
delta_of: ai-industry-kb-2026
sha256: b6a011ba19f35f346db2920d3f9d56b8e6ce6e949b5a6d3f1f3e69bcfebad460
---

# Nemotron VoiceChat 11B and Nemotron-3-Nano-Omni — speech and omni models

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


### Nemotron 3 Super — the March 2026 mid-tier flagship
- NVIDIA released Nemotron 3 Super (`nemotron-3-super-120b-a12b`) at GTC in March 2026 — reported as March 11, 2026 in secondary catalogs, while the base corpus dates GTC 2026 announcements to March 16, 2026; preserve both dates. [SECONDARY] [S31][S32]
- Architecture: hybrid Mamba-2 + Transformer sparse-attention + LatentMoE with a Multi-Token Prediction (MTP) head; 120B total parameters, ~12B active per token (reported as 12.7B in some catalogs — preserve both). 1M-token context window. [SECONDARY] [S31][S32][S33]
- CONTRADICTION: active parameters reported as 12B (TokenCost/awesomeagents/dev.to) versus 12.7B (other catalogs). Preserve both; they describe the same checkpoint. [SECONDARY] [S31][S32]
- Speed: 449 output tokens/sec per Artificial Analysis — ranked #1 of 51 models in its intelligence tier, 2.2x faster than GPT-OSS-120B. (single secondary coverage) [SECONDARY] [S31]
- Long-context: 91.75% on RULER at full 1M length; 96.30% at 256K, 95.67% at 512K. The 1M window defaults to 256K in the shipped configuration due to VRAM constraints; extending to the full million requires `VLLM_ALLOW_LONG_MAX_MODEL_LEN=1`. [SECONDARY] [S32][S36]
- Benchmarks (vendor-reported, via community): SWE-Bench Verified 60.47% (vs GPT-OSS-120B 41.90%), MMLU-Pro 83.73%, AIME 2025 90.21% (no tools), HMMT Feb 2025 93.67% (no tools), LiveCodeBench 81.19%, GPQA 79.23 (82.70% with tools), SciCode subtask 42.05%. Artificial Analysis: Intelligence Index 35.97, Coding 31.19, Agentic 40.18; GPQA Diamond 80.0%, HLE 19.2%, SciCode 36.0%, TAU-2 67.8%, IFBench 71.5%, LCR 60.0%. (single vendor coverage) [VENDOR] [S36]
- Competitive position is mixed: it beats GPT-OSS-120B on most tasks but trades wins with Qwen3.5-122B-A10B — leading on long-context retrieval (RULER@1M 91.75%), HMMT math and LiveCodeBench, trailing on general knowledge (MMLU-Pro 83.73 vs 86.70) and science reasoning (GPQA 79.23 vs 86.60), and on agentic coding (SWE-Bench 60.47 vs 66.40). (single secondary coverage) [SECONDARY] [S32]
- Training: 15.6T tokens, 153 datasets, 20 languages, 43 programming languages, 10+ RL environments; pre-training cutoff June 2025, post-training cutoff February 2026. NVFP4 quantization-aware training from day one; BF16/FP8/NVFP4 weights published. (single secondary coverage) [SECONDARY] [S32]
- Full open release: weights on Hugging Face, 153 datasets and 15 RL environments; technical report at research.nvidia.com and arXiv 2604.12374. [COMMUNITY] [S35][S32]
- Deployment: minimum 8x H100-80GB self-host; NIM containers; 15+ cloud providers; free on OpenRouter and build.nvidia.com. Median hosted pricing $0.30/$0.80 per 1M input/output tokens; as low as $0.10/$0.50 on DeepInfra. [SECONDARY] [S31][S32]
- License: reported as the NVIDIA Nemotron Open Model License for Super [S32] — while Lightning/Cosmos 3/Alpamayo 2 Super are reported under OpenMDW-1.1 [S1][S13][S21]. Preserve both attributions; they may reflect NVIDIA's license renaming over 2026. [SECONDARY]

### Jetson Thor — the on-device compute ladder for Cosmos/GR00T
- Jetson AGX Thor (T5000 module) reached general availability August 25, 2025: NVIDIA Blackwell GPU, 128GB memory, up to 2,070 FP4 TFLOPS at 130W; a lighter T4000 version at ~70W and 1,200 FP4 TFLOPS. 7.5x the AI compute and 3.5x the energy efficiency of Jetson AGX Orin. [SECONDARY] [S37][S38]
- On July 15, 2026 NVIDIA expanded the family to four tiers: T5000 (2,070 FP4 TFLOPS, 128GB, shipping, module $2,999), T4000 (1,200 FP4 TFLOPS, 64GB, shipping, module $1,999), T3000 (865 FP4 TFLOPS, 32GB, Q1 2027) and T2000 (400 FP4 TFLOPS, 16GB, Q1 2027) — positioned as successors to Orin AGX and Orin NX. All four share Blackwell GPUs, Arm Neoverse V3AE CPUs and JetPack 7.2.1. (single community coverage) [COMMUNITY] [S39]
- Named use: running Isaac GR00T N1.5 and VLMs/VLA models on-device; early adopters include Boston Dynamics and Figure for humanoid robots. Developer kit ~$3,499. (single secondary coverage) [SECONDARY] [S37]
- T5000 hardware detail: 14-core Arm Neoverse V3AE CPU, 2,650-core Blackwell GPU, 96 fifth-generation Tensor Cores, PCIe Gen 5, third-generation Programmable Vision Accelerator, optical flow accelerators and dual encoders/decoders for low-latency visual tasks. (single secondary coverage) [SECONDARY] [S38]
- Carrier ecosystem: Auvidea's X242 carrier (same form factor as AGX Orin, >98% power efficiency via GaN MOSFETs), Lanner EAI-I351 (8x GMSL2 deserializers, -25°C to 70°C operating range) and Vecow EAC-7000 series — all shipping production systems around the Thor modules. [SECONDARY] [S38][S48]

### NVIDIA Dynamo — disaggregated inference at GTC 2026
- At GTC 2026 NVIDIA introduced "disaggregated inference" through Dynamo, described as an AI factory operating system: the inference pipeline splits into prefill/attention (compute- and KV-cache-heavy, run on Vera Rubin systems) and decode/token generation (latency-sensitive, offloaded to Groq LPU racks with expanded memory). (single secondary coverage) [SECONDARY] [S40]
- NVIDIA's technical blog (via secondary) claims Dynamo boosts inference throughput up to 7x in real-world multi-node benchmarks. (single vendor coverage) [VENDOR] [S41]
- Core components: SLO Planner (GPU scheduling against latency/throughput targets), KV-aware router (routes to reuse existing KV cache), NIXL low-latency communication library, KV Block Manager (offloads cache to CPU/SSD/object storage), and Grove (Kubernetes-native topology-aware scheduling). (single secondary coverage) [SECONDARY] [S41]
- Production shape: LangChain and NeMo Agent Toolkit integrations; multimodal disaggregated encode/prefill/decode with embedding cache (30% faster time-to-first-token on image workloads); native video-generation support (real-time 1080p on a single B200); Kubernetes Inference Gateway plugin; S3/Azure blob storage-tier KV offload. Prebuilt runtimes for SGLang, TensorRT-LLM and vLLM at 1.0.1. [COMMUNITY] [S42]
- TrendForce notes the third-generation Groq LP30 chip (Samsung-fabricated) entered full-scale production for 2H26 shipment, with a higher-performance LP40 planned for the Feynman architecture — the hardware roadmap behind the decode-offload story. (single secondary coverage) [SECONDARY] [S40]

### Cosmos lineage — from Cosmos 1 to Cosmos 3
- The first Cosmos generation was unveiled by Jensen Huang at CES: world foundation models trained on 20M hours of video (driving, hand manipulation, human motion, navigation, first-person POV, nature dynamics), in Nano/Super/Ultra sizes, under an open license on GitHub — pitched as doing "for the world of robotics and industrial AI what Llama 3 has done for enterprise AI." (single secondary coverage) [SECONDARY] [S46]
- The GTC-era major Cosmos release added three named components: Cosmos Transfer (structured video inputs — segmentation/depth/lidar/pose/trajectory maps — to controllable photoreal video for synthetic data), Cosmos Predict (multi-frame generation to predict intermediate actions), and Cosmos Reason (open chain-of-thought reasoning over video). Plus two Omniverse-powered blueprints for large-scale synthetic data generation. Early adopters: 1X, Agility Robotics, Figure AI, Foretellix, Skild AI, Uber. (single vendor coverage) [VENDOR] [S43]
- Cosmos 3 unified what the earlier generation split across four families — Predict (world generation), Transfer (controlled generation), Reason (scene understanding), Policy (action generation) — into the single two-tower MoT model. [COMMUNITY] [S16][S53]
- Community literature notes track the 2.5 generation: Cosmos-Predict2.5 (flow-based, Text2World/Image2World/Video2World unified, grounded by Cosmos-Reason1, trained on 200M curated clips with RL post-training, 2B and 14B scales) and Cosmos-Transfer2.5 (control-net-style Sim2Real/Real2Real, 3.5x smaller than Transfer1 at higher fidelity) — under the NVIDIA Open Model License at github.com/nvidia-cosmos. (single community coverage) [COMMUNITY] [S45]

## Figures and metrics
| Model / item | Size (total/active) | Date | Note |
|---|---|---|---|
| Nemotron 3 Nano | 31.6B / 3.2B | 2025-12-15 | Efficiency entry [VENDOR] |
| Nemotron 3 Super | ~120B / 12.7B | 2026 | Mid tier [SECONDARY] |
| Nemotron 3 Ultra | 550B / 55B | 2026-06-01 launch; weights 2026-06-04 | 90% sparsity; Mamba-2 + LatentMoE + NVFP4 [SECONDARY] |
| Nemotron 3.5 Lightning | 30B / 3B | 2026-08-11 | + NeMo Switchyard [SECONDARY] |
| Nemotron-Cascade 2 | 30B / 3B | 2026-03 | IMO 2025 gold; Open Model License [SECONDARY] |
| Nemotron 4 | ≥1T | — | In development, NOT announced [SECONDARY] |
| Groq 3 LPX | — | 2026-08-24 full production | Agentic-AI inference speed [VENDOR] |

- Poolside deal: **$6B Model Factory license + $1B investment** — not an acquisition, not a Laguna license [SECONDARY].
- Free access: full Nemotron 3 family (Ultra/Super/Nano/3.5) on OpenRouter :free routes [COMMUNITY].
- Breach response stat: GLM 5.2 (Chinese open weights) did the forensic work US closed models refused [SECONDARY].


### New verified metrics — expansion

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


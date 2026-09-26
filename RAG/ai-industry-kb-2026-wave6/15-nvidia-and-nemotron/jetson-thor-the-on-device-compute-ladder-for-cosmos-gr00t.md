---
id: ai-industry-kb-2026-wave6/15-nvidia-and-nemotron/jetson-thor-the-on-device-compute-ladder-for-cosmos-gr00t
title: "Jetson Thor — the on-device compute ladder for Cosmos/GR00T"
domain: nvidia-and-nemotron
role: deep-dive
task: funding-deals
actors: ["China", "Groq", "Nvidia", "OpenRouter", "Poolside", "SGLang", "Samsung", "TensorRT-LLM", "United States", "Z.ai", "vLLM"]
dates: ["2025-08-25", "2025-12-15", "2026-06-01", "2026-06-04", "2026-07-15", "2026-08-11", "2026-08-24"]
keywords: ["compute", "accelerator", "acquisition", "agent", "agentic", "attention", "benchmarks", "blackwell", "decode", "disaggregated", "embedding", "energy"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [7533, 7570]
section: "§15. NVIDIA and Nemotron"
delta_of: ai-industry-kb-2026
sha256: 3f5628c14b4c3684c2d519deea7518beea74afcff9e385d2adac83b114e38707
---

# Jetson Thor — the on-device compute ladder for Cosmos/GR00T

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


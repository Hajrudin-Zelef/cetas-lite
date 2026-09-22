---
id: etape4-trackb-local-inference/00-local-inference/2-7-jetson-edge-2026
title: "2.7 Jetson / edge (2026)"
domain: step-4-track-b-local-inference-stack-llama-cpp-ollama-lm-stu
role: deep-dive
task: reference
actors: ["Alibaba", "Meta", "Mistral", "Nvidia", "TensorRT-LLM", "vLLM"]
dates: ["2026-07"]
keywords: ["benchmark", "benchmarks", "blackwell", "compute", "energy", "fp4", "gguf", "gpu", "inference", "llama", "llama.cpp", "lpddr5x"]
source: docs/RAG/etape4_trackB_local_inference.md
source_anchor: ""
source_lines: [850, 922]
section: "Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)"
sha256: be73b6916da8b69d77fc6272f504b245decbc69abbcfd91a29148bebcb96259a
---

# 2.7 Jetson / edge (2026)

### 2.7 Jetson / edge (2026)

**Jetson AGX Thor (T5000)** — NVIDIA's "physical AI" flagship (robotics/humanoids), not a desktop LLM box:
- Specs: Blackwell GPU, 2,070 TFLOPS FP4 sparse, 128 GB LPDDR5X, 40–130 W configurable, 14-core Arm Neoverse-V3AE CPU, Multi-Instance GPU; ~7.5× AI compute and ~3.5× energy efficiency claims vs AGX Orin (note: mixes FP4-sparse vs INT8-sparse in the comparison). **Dev kit $3,499**; bulk module $2,999. [independent] (http://aiwiki.ai/wiki/jetson_thor), (https://www.techradar.com/pro/nvidia-quietly-unveiled-its-fastest-mini-pc-ever-capable-of-topping-2070-tflops-and-if-you-squint-enough-you-might-even-think-it-looks-like-an-rtx-5090)
- Developer kit released Aug 25, 2025 (shipments from Nov 20, 2025 per pre-order coverage); runs JetPack 7 (Ubuntu 24.04, kernel 6.8); JetPack 7.1 current production release. [independent] (https://github.com/liunix61/awesome-embedded-ai-stack/blob/HEAD/hardware-landscape/jetson.md), (https://twowintech.com/nvidia-jetson-agx-thor-full-analysis/)
- July 2026: NVIDIA announced lower **T3000 / T2000** modules (smaller memory/form factor, Q1 2027 availability; pricing unannounced; CNX/ServeTheHome flagged specs as preliminary). [independent] (http://aiwiki.ai/wiki/jetson_thor)
- LLM angle: 128 GB lets Thor run Llama-70B-class models locally, but ~273 GB/s bandwidth (like DGX Spark) limits token throughput; community consensus: relevant for 70B+ edge inference and VLA/robotics workloads, overkill for simple robotics (Orin Nano class suffices). [independent] (https://github.com/alpininsight/capi-provider-ssh/blob/HEAD/docs/roadmap/nvidia-jetson-edge-devices.md)
- Hands-on: HotHardware/ServeTheHome demos of GR00T N1 on Thor; wayin.ai tested local LLMs/VLMs on the dev kit (128 GB, 273 GB/s bandwidth noted as "slow for GPU memory, fast for CPU/RAM"). [independent] (http://aiwiki.ai/wiki/jetson_thor), (https://wayin.ai/wayinvideo/s/chshare0uMzeE0gKLh/)

---

## PART 3 — GGUF QUANTIZATION BENCHMARKS (2026)

### 3.1 Quality vs size: the standard numbers

**Measured perplexity/KL data (ikawrakow / Artefact2 measurements on Mistral-7B-class models; bits-per-weight, KL-divergence median, ln(PPL(Q)/PPL(base)))** [independent — community measurements, widely cited as the reference] (https://gist.github.com/Artefact2/b5f810600771265fc1e39442288e8ec9):

| Quant | Bits/weight | KL-div median | ln(PPL(Q)/PPL(base)) | Notes |
|---|---|---|---|---|
| Q8_0 | ~8.5 | 0.0043 | 0.0005 | ~lossless |
| Q6_K | 6.57 | 0.0032 | −0.0008 | ~lossless |
| Q5_K_M | 5.67 | 0.0043 | 0.0005 | near-lossless |
| Q5_K_S | 5.52 | 0.0045 | 0.0005 | — |
| **Q4_K_M** | 4.83 | 0.0075 | 0.0060 | community default |
| Q4_K_S | 4.57 | 0.0083 | 0.0081 | — |
| **IQ4_NL** | 4.56 | 0.0085 | 0.0074 | — |
| **IQ4_XS** | 4.32 | 0.0088 | 0.0079 | ≈Q4_K_M quality, fewer bits |
| Q3_K_L | 4.22 | 0.0152 | 0.0205 | — |
| IQ3_M | 3.63 | 0.0186 | 0.0268 | — |
| Q3_K_M | 3.89 | 0.0171 | 0.0258 | — |
| IQ3_XS | 3.32 | 0.0296 | 0.0458 | — |
| Q3_K_S | 3.50 | 0.0304 | 0.0511 | cliff: avoid |
| Q2_K | 3.00 | 0.0588 | 0.1103 | — |
| IQ2_M | 2.76 | 0.0702 | 0.1223 | — |

**PPL-increase summary table (gguf-switchboard, based on Llama-3.1-8B; measured except i-quants which are extrapolated)** [independent] (https://github.com/pradeepgudipati/gguf-switchboard/blob/HEAD/docs/QUANT_SCORING.md):

| Quant | PPL increase vs FP16 | Quality score | Confidence |
|---|---|---|---|
| Q8_0 | 0.1% | 99.9 | Measured |
| Q6_K | 0.4% | 99.6 | Measured |
| Q5_K_M | 1.1% | 98.9 | Measured |
| Q4_K_M | 3.3% | 96.7 | Measured |
| Q4_K_S | 4.1% | 95.9 | Measured |
| Q3_K_L | 6.7% | 93.3 | Measured |
| Q3_K_M | 8.7% | 91.3 | Measured |
| Q3_K_S | 22.4% | 77.6 | Measured |
| Q2_K | ~35% | ~65 | Extrapolated |
| IQ4_NL/IQ4_XS, IQ3_*, IQ2_* | various | various | Extrapolated — treat as order-of-magnitude only |

Caveat (stated in the source itself): this is a generic table; quantization loss varies by architecture and training; per-model measured data (imatrix runs, quantizer model cards) beats the table for that model. [independent]

### 3.2 2026 consensus: K-quants vs i-quants, imatrix

- **Q4_K_M remains the default/recommended 4-bit quant** ("fast, recommended", "community sweet spot: minimal quality loss, major VRAM savings"). Q4_K_S is the size/speed/quality optimum one step down. [independent] (https://github.com/casteldazur/awesome-local-ai/blob/HEAD/guides/vram-requirements.md), (https://huggingface.co/mradermacher/UNAversal-8x7B-v1beta-GGUF)
- **i-quants (imatrix-calibrated) are the 2026 standard for squeezing more quality per bit.** Quantizer guidance (mradermacher, a major GGUF publisher) repeated across model cards: "IQ-quants are often preferable over similar sized non-IQ quants" — specifically: IQ3_S beats Q3_K*; IQ3_M beats Q3_K_L; IQ4_XS approaches Q4_K_M quality at 4.32 vs 4.83 bits/weight. [independent] (https://huggingface.co/mradermacher/Llama3-8B-ScaleQuest-i1-GGUF), (https://huggingface.co/mradermacher/Qwen2.5-14B-Instruct-abliterated-i1-GGUF)
- **imatrix calibration** = importance-matrix computed from a calibration corpus; weights that matter more get more bits. The "i1-" prefix on mradermacher repos denotes imatrix quants. Quality varies with calibration data and the compute used to produce it (mradermacher credits access to a private supercomputer for higher-quality imatrix runs). [independent] (same)
- Practical ladder at cutoff (per major quantizer model cards):
  - Desperate/low VRAM: IQ2_M / IQ3_XS ("IQ3_XXS probably better than Q2_K")
  - Mid: IQ3_S/M, IQ4_XS
  - Default: Q4_K_S / Q4_K_M ("fast, recommended")
  - High quality: Q5_K_M, Q6_K ("very good quality"), Q8_0 ("best quality", ~lossless) [independent] (https://huggingface.co/mradermacher/UNAversal-8x7B-v1beta-GGUF)
- **ikawrakow's quant-PPL graph** (nethype.de) is the community's canonical visual for lower-quality quant comparison (lower = better), linked from most quantizer model cards. [independent] (https://www.nethype.de/huggingface_embed/quantpplgraph.png)
- FP4 (NVFP4) is emerging as a hardware format on Blackwell (DGX Spark, RTX 50) via TensorRT-LLM / vLLM and incoming llama.cpp CUDA support — separate from GGUF integer quants; expect format convergence work through 2026. [secondary] (https://www.compute-market.com/blog/rx-9070-xt-vs-rtx-5060-ti-local-ai-2026)

### 3.3 Notes on sources & limitations

- Most quantization PPL data is community-measured (ikawrakow, Artefact2) on a handful of architectures; i-quant entries in summary tables are often extrapolated, not measured — flagged above.
- r/LocalLLaMA was not directly mined for this report; community claims above come from benchmark repos, quantizer model cards, and review sites.
- No new GGUF quant *types* beyond the K/i-quant families surfaced in 2026 coverage at cutoff; the movement is in imatrix quality, speculative-decoding drafters (MTP/DFlash/DSpark in LM Studio 0.4.22+), and hardware FP4.

---


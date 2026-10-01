---
id: collect-261001-ia-llm/ia-llm/etape4-trackb-local-inference-16
title: "Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "Alibaba", "Intel", "Meta", "Mistral", "Nvidia", "TensorRT-LLM", "vLLM"]
dates: []
keywords: ["llama", "llama.cpp", "amd", "benchmark", "blackwell", "compute", "fp4", "gguf", "intel", "memory", "mistral", "mxfp4"]
source: docs/RAG/collect-261001-ia-llm/etape4_trackB_local_inference.md
source_anchor: ""
source_lines: [863, 937]
sha256: a98652a5e1546ef100028428cd4947996bbe2eaafcc5fe32bf427e8a9624a7ee
---

# Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)

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

## OPEN QUESTIONS / UNCERTAINTIES (flagged for follow-up)

1. Exact LM Studio download/user counts — no vendor-published figure found; "millions of downloads" is secondary only. [unverified]
2. LM Studio Enterprise tier and Secure Cloud pricing — not published at cutoff. [unverified]
3. DFlash / DSpark assistant drafters in LM Studio 0.4.22 — changelog-confirmed names only; architectural details not verified. [unverified]
4. IPEX-LLM archival (Jan 28, 2026) — from a community research doc; not independently confirmed against Intel's repo at cutoff. [unverified]
5. Snapdragon X2 Elite / X2 Elite Extreme specs (80 TOPS claim) — single secondary source. [unverified]
6. DGX Spark "2.5×" post-launch software speedup and GPT-OSS-120B 58.8 tok/s (vLLM/MXFP4) figures — vendor/community-reported, not independently reproduced here. [vendor-reported/secondary]
7. Street prices (RTX 5090 ~$2,500–$3,799; RX 9070 XT ~$740) move with GDDR7/memory supply; figures are dated snapshots (Aug–Sep 2026). [secondary]
8. AMD's listed price for the Ryzen AI Halo developer platform ($3,999) and ASRock AI BOX-A395 pricing (unannounced) — from secondary coverage. [secondary]

---

## APPENDIX A — 2026 TIMELINE (LM Studio / local hardware / GGUF-relevant)


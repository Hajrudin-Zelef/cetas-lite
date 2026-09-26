---
id: ai-industry-kb-2026/08-quantization-model-formats/overview
title: "8. Quantization & Model Formats"
domain: quantization-model-formats
role: deep-dive
task: quantization
actors: ["AMD", "Alibaba", "Apple", "Hugging Face", "Intel", "Nvidia", "SGLang", "Unsloth", "vLLM"]
dates: ["2023-03", "2023-08-22", "2026-02-20", "2026-08-19", "2026-08-27", "2026-09", "2026-09-22"]
keywords: ["quantization", "acquisition", "amd", "attention", "awq", "blackwell", "datacenter", "decode", "fp4", "fp8", "gguf", "gptq"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [3767, 3808]
section: "8. Quantization & Model Formats"
sha256: 0e90150c235d99900a1abce16d933ce1a1c4bff1053f2111b11493e770564bf6
---

# 8. Quantization & Model Formats
Keywords: GGUF, K-quants, IQ-quants, imatrix, ggml.ai, Hugging Face, Unsloth Dynamic 3.0, EXL2, ExLlamaV3, EXL3, QTIP, AWQ, GPTQ, AutoRound, HQQ, GPTQModel, llm-compressor, ModelOpt, bitsandbytes, torchAO, Marlin, Machete, NVFP4, MXFP4, NF4, vLLM, llama.cpp, ARM edge

## Summary

As of 2026-09-22, the quantization ecosystem has completed its split into two production stacks with GGUF as the undisputed local-inference standard. Hugging Face's absorption of **ggml.ai** (announced 2026-02-20, team and MIT license preserved) cemented GGUF — supported by llama.cpp, Ollama, LM Studio, GPT4All, Jan and koboldcpp, with tens of thousands of checkpoints on Hugging Face — as the reference container for quantized models, while datacenter serving standardized on safetensors-based FP8 / NVFP4 / MXFP4 under vLLM and SGLang. (The FP8/FP4/NVFP4/MXFP4 production arc, TurboQuant and KV-cache quantization detail live in sibling part 08b; this part covers formats and tooling.)

Measured per-type accuracy data (Artefact2's KL-divergence/perplexity table; mradermacher's imatrix cards) now pins every GGUF K-quant and I-quant to numbers: IQ3_S beats Q3_K_M at equal size, IQ4_XS matches Q4_K_M quality at ~4.3 bits/weight, Q4_K_S is the community "optimal size/speed/quality" pick, and Q4_K_M remains the default. Unsloth's selective-layer quantization generation moved to **Dynamic v3.0** (launched 2026-08-19; SUPERSEDES Dynamic 2.0 — Wave 1's v2.0 section is deprecated as "current"): >10% better top-1 accuracy than every other provider at the same size [VENDOR], with 1-bit `UD-TQ1_0` and 2-bit `UD-Q2_K_XL` variants; the first independent replication test (2026-08-27) confirmed the direction on KL divergence but measured a 9.5% decode-speed regression and one UD-Q3_K_XL file that wedges llama-server [COMMUNITY].

The 2026 tooling shakeout is decisive: **AutoAWQ is officially deprecated** (AWQ quantization now via llm-compressor), **GPTQModel** replaced AutoGPTQ, **ExLlamaV2 is archived** with development continuing on **ExLlamaV3** (v1.4.6–v1.4.9, September 2026 patch cadence) and the **EXL3** format (a streamlined QTIP variant, coherent at 1.6 bpw), **optimum-quanto is in maintenance mode** (Hugging Face redirects users to bitsandbytes or torchAO), and **Intel AutoRound** is the strongest measured 4-bit PTQ method (91.4% logprob top-1 agreement vs BF16, beating GPTQ 90.5% and AWQ on 27/32 lm-eval configs). A reproducible vLLM (NVFP4/FP8+MTP) vs ExLlamaV3 (EXL3 4.05bpw) shootout on Qwen3.8-Flash-Next on RTX PRO 6000 (Sept 2026) [COMMUNITY] sharpened the verdict: EXL3 wins only short single-stream decode (169–173 tok/s); vLLM prefills ~2.7× faster and scales where TabbyAPI/ExLlamaV3 fails (21/60 completions at 32k×4). The brief's "AWQ as production standard" claim VERIFIES only in nuance: AWQ is the dominant compat INT4 path on both vLLM and SGLang ("more commonly used of the two" vs GPTQ), but AWQ tooling lags newest architectures, and a Blackwell decode shootout had AWQ beating both NVFP4 recipes — the standard holds as a compat format, not as an always-fastest one.

## Key dated facts

### GGUF: the de facto local-inference standard

- GGUF (GGML Universal File) is a binary format storing tensors plus metadata in a single file, designed for fast save/load. Developed by Georgi Gerganov and the community; initial release **2023-08-22**; current **v3**. It superseded the older GGML format to keep backward compatibility as new architectures were added.
- As of 2026 it is **the standard format for distributing quantized LLMs for local inference**, natively supported by **llama.cpp, Ollama, LM Studio, GPT4All, Jan, and koboldcpp**.
- **Hugging Face hosts tens of thousands of GGUF checkpoints** with first-class integration: a metadata viewer, an inference-endpoint service, and a JavaScript parser library.
- llama.cpp exposes backends for **CUDA, Metal, HIP (AMD), Vulkan, SYCL, and CPU** — the portability basis for GGUF across NVIDIA, AMD, Apple Silicon and CPU-only machines.
- Conceptual note: **GGUF is a container, not a compression method**. A GGUF file may hold unquantized tensors or any of the K-quant / I-quant families (Q2_K … Q8_0, IQ1_S … IQ4_XS). Quality comes from the quant inside, not the container.
- GGUF filenames follow a published naming convention: base name, size label, fine-tune, version, encoding, type, shard (5-digit `00003-of-00009` counters), with optional `mmproj-` (vision projector) and `mtp-` (multi-token-prediction draft) prefixes.

### Hugging Face acquires ggml.ai (2026-02-20)

- **ggml.ai joins Hugging Face — announced 2026-02-20.** Primary evidence: the official announcement posted as GitHub Discussion ggml-org/llama.cpp #19759 ("ggml.ai joins Hugging Face to ensure the long-term progress of Local AI"), covered by Simon Willison (2026-02-20) and secondary press.
- Background: ggml.ai was founded in 2023 by **Georgi Gerganov** (llama.cpp creator; originally backed by Nat Friedman and Daniel Gross).
- This is a **genuine acquisition, not a sponsorship or loose partnership** — the team becomes Hugging Face employees. No financial terms disclosed.
- Commitments: projects stay **100% MIT-licensed** and community-driven; Gerganov's team keeps full technical leadership and works on llama.cpp full-time; HF provides long-term sustainable resources.
- Joint goal: **"single-click" integration between ggml-based engines and the Transformers library** (Transformers as the "source of truth" for model definitions).
- Two HF engineers — **ngxson (Son) and allozaur (Alek)** — were already core llama.cpp contributors, making the absorption, in HF's words, "a very natural process."
- llama.cpp (March 2023) made 4-bit inference possible on a MacBook and ignited the local-AI movement; the repository passed **126,000+ GitHub stars** by 2026.
- 2026 practitioner guides list **GGUF Q4_K_M via Ollama/LM Studio as the default for local CPU/Mac deployment**.
- [UNVERIFIED] A single secondary outlet (pasqualepillitteri.it, 2026) claims NVIDIA offered $12.9B to acquire Hugging Face — no Reuters/Bloomberg corroboration as of 2026-09-22. The MIT license on llama.cpp/GGML is in any case irrevocable.

### GGUF K-quants: measured per-type accuracy data

- Hugging Face's reference table for a Llama-2-7B-class model: FP16 5.9565 baseline; Q8_0 5.9584 (+0.03%); Q6_K 5.9642 (+0.13%); Q5_K_M 5.9796 (+0.39%); Q4_K_M 6.0565 (+1.68%). Newer models react differently; use as illustrative only.
- Community quality-vs-FP16 rules of thumb (2026): Q8_0 ~99.5% (0.55x size, near-lossless, the "safe" choice); Q6_K ~99% (0.42x, diminishing returns above); Q5_K_M ~98% (0.36x, recommended ≤14B); **Q4_K_M ~96% (0.30x, best quality/size trade-off, the default)**; Q3_K_M ~90% (0.25x, very large models only); Q2_K ~80% (0.18x, aggressive).
- **Q4_K_M** is a mixed 4/6-bit scheme (attention plus first/last blocks at 6-bit), measuring ~0.1–0.15 perplexity above FP16 on WikiText with 1–3% MMLU loss.
- **Importance matrices (imatrix)** materially improve 1–3-bit quants: `llama-imatrix` computes per-weight importance from calibration text, and `llama-quantize --imatrix` uses it during quantization. For IQ1/IQ2 mixes the tool warns if no imatrix is supplied.
- Artefact2's GGUF quantization overview (measured; bits per weight, median KL divergence vs FP16, q99 KL, fraction of top tokens differing, ln(PPL(Q)/PPL(base))):


---
id: collect-240926-huggingface/huggingface/hauhaucs-qwen3-8-27b-uncensored-hauhaucs-aggressive-mtp-gguf-hugging-face-1
title: "hauhaucs-qwen3-8-27b-uncensored-hauhaucs-aggressive-mtp-gguf-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["AMD", "Alibaba", "Hugging Face", "Nvidia"]
dates: []
keywords: ["gguf", "qwen", "agentic", "attention", "benchmark", "blackwell", "consumer", "deepseek", "gpu", "llama", "llama.cpp", "quantization"]
source: docs/RAG/clean_en/huggingface/hauhaucs-qwen3-8-27b-uncensored-hauhaucs-aggressive-mtp-gguf-hugging-face.md
source_anchor: ""
source_lines: [1, 128]
sha256: 0d8b1e2909cd88656da160df14a372e35de1f2b940762359f032422be8cf1ca9
---

# hauhaucs-qwen3-8-27b-uncensored-hauhaucs-aggressive-mtp-gguf-hugging-face

<!-- source: https://huggingface.co/HauhauCS/Qwen3.8-27B-Uncensored-HauhauCS-Aggressive-MTP-GGUF -->

**HauhauCS FastMTP: up to 3.02x document TG and 1.93x reasoning TG versus non-MTP — plus up to 35.2% more document TG and 21.1% more reasoning TG than standard embedded MTP.**


**Join the Discord** for updates, roadmaps, projects, or just to chat.


Qwen3.8-27B uncensored by HauhauCS **0/465 Refusals*** . 

This is the **Aggressive variant**: direct answers, no refusal behavior, and minimal preamble on hard prompts.

**Every text GGUF preserves Qwen3.8's native NextN head, and this release adds HauhauCS FastMTP: a specific acceleration sidecar qualified across the complete quant lineup at maximum native context.** Vision is included through the separate BF16 projector.

**Hugging Face's Hardware Compatibility widget may not recognize K_P quants.** If files appear to be missing, click **View variants** or open **Files and versions**.


No changes to datasets or intended capabilities. This release preserves Qwen3.8-27B's text, reasoning, agentic, image, and video capabilities while applying the HauhauCS Aggressive uncensoring profile.

Pick Aggressive when you specifically want the model to get to the answer without first talking itself into compliance. For reliability-critical, specifically long-context agentic work, a Balanced release is normally the safer default when/if one is available.

BPW is the encoded tensor-payload average across the complete text model, including its embedded MTP tensors, rounded to two decimals. The projector and FastMTP sidecar work with every text quant; download the projector only for image or video input.

K_P ("Perfect") quants are HauhauCS custom quantizations that use model-specific analysis to selectively preserve quality where it matters most. Every model gets its own optimized quantization profile.

A K_P quant effectively bumps quality up by one or two quant levels at only around 5–15% more size than the base quant. The files remain standard GGUFs and work with llama.cpp, LM Studio, and other GGUF-compatible runtimes with no special build or plugin.

**Note:** K_P quants may show as `?` in LM Studio's quant column. This is a display issue only—the model loads and runs normally.

- Dense 27B causal language model with a vision encoder
- 64 language-model layers
- Hidden size 5,120; FFN size 17,408
- 248,320-token padded vocabulary
- 48 Gated DeltaNet layers and 16 gated-attention layers
- Native embedded MTP/NextN preserved, plus the HauhauCS FastMTP 32K acceleration profile
- 262,144-token native context; extensible up to 1,000,000 with framework-specific configuration
- Native text, image, and video understanding
- Based on Qwen/Qwen3.8-27B

HauhauCS FastMTP is the custom, variant-specific acceleration profile built for this exact Aggressive release: a compact 32K draft sidecar and per-quant serving profiles qualified for TG, acceptance, maximum native context, and VRAM.

It delivers **up to 3.02x document TG and 1.93x reasoning TG versus non-MTP, plus up to 35.2% more document TG and 21.1% more reasoning TG than the standard embedded-MTP profile.** The unchanged full target verifies every drafted token, so FastMTP accelerates generation without replacing the target model or changing its answers. The construction and selection methodology is exclusive to HauhauCS releases.

The benchmark ladder:

| Comparison | Document TG | Reasoning TG | Scope | 
|---|---|---|---|
| Standard embedded MTP vs MTP disabled | **2.23x** (`+123.4%` ) | **1.60x** (`+59.6%` ) | Final Q8_K_P, depth 2 | 
| HauhauCS FastMTP profile vs standard embedded MTP | **+35.2%** | **+21.1%** | Final Q8_K_P, depth 3 vs depth 2 | 
| HauhauCS FastMTP vs embedded MTP at identical depth | **+11.1%** | **+18.2%** | Final Q8_K_P, depth 3 | 
| HauhauCS FastMTP vs MTP disabled | **3.02x** (`+202.0%` ) | **1.93x** (`+93.3%` ) | Final Q8_K_P service | 

These results were measured on one RTX PRO 6000 Blackwell 96 GB per isolated lane at `204800` configured context, full CUDA offload, `--no-mmap`, and the official reasoning sampler. FastMTP accelerates TG; PP is reported alongside it for a complete serving comparison.

There are two acceleration paths:

- **Embedded MTP:** use any target GGUF by itself with`--spec-type draft-mtp` in a current upstream llama.cpp build.
- **HauhauCS FastMTP:** pair that same target with`Qwen3.8-27B-Uncensored-HauhauCS-Aggressive-FastMTP-32K.gguf` and the HauhauCS runtime patch below.

The compact draft uses a standard GGUF `d2t` token map plus a minimal Qwen3.8 runtime consumer. Build it once. The example below uses CUDA; for ROCm/HIP or Vulkan, replace `-DGGML_CUDA=ON` with `-DGGML_HIP=ON` or `-DGGML_VULKAN=ON`. For CPU-only, omit the backend flag.

```
git clone https://github.com/ggerganov/llama.cpp
cd llama.cpp
git checkout 4df29be4f4c3673f428170fda944a5b19f743bb8
curl -L -o HauhauCS-FastMTP-llama.cpp.patch \
  https://huggingface.co/HauhauCS/Qwen3.8-27B-Uncensored-HauhauCS-Aggressive-MTP-GGUF/resolve/main/HauhauCS-FastMTP-llama.cpp.patch
git apply --check HauhauCS-FastMTP-llama.cpp.patch
git apply HauhauCS-FastMTP-llama.cpp.patch
cmake -S . -B build -DGGML_CUDA=ON -DCMAKE_BUILD_TYPE=Release
cmake --build build --config Release -j"$(nproc)"
```
If draft loading reports `expected 5120, 248320, got 5120, 32768`, the FastMTP sidecar is correct but the executable is unpatched. Launch the freshly built `./build/bin/llama-server` from this checkout.

Then serve any target quant with the one shared FastMTP sidecar:

```
MODEL=Qwen3.8-27B-Uncensored-HauhauCS-Aggressive-Q4_K_P.gguf
DRAFT=Qwen3.8-27B-Uncensored-HauhauCS-Aggressive-FastMTP-32K.gguf
DEPTH=3
CUDA_VISIBLE_DEVICES=0 ./build/bin/llama-server \
  --model "$MODEL" \
  --spec-draft-model "$DRAFT" \
  --spec-draft-ngl all \
  --spec-type draft-mtp \
  --spec-draft-n-max "$DEPTH" \
  --spec-draft-p-min 0 \
  --ctx-size 204800 \
  --parallel 1 \
  --batch-size 2048 \
  --ubatch-size 512 \
  --n-gpu-layers all \
  --split-mode none \
  --flash-attn on \
  --no-mmap \
  --temp 1.0 \
  --top-k 20 \
  --top-p 0.95 \
  --min-p 0 \
  --presence-penalty 0 \
  --repeat-penalty 1.0 \
  --jinja \
  --reasoning on \
  --reasoning-effort xhigh \
  --reasoning-preserve \
  --reasoning-format deepseek \
  --host 127.0.0.1 \
  --port 8080
```
Three-run medians for the uncached 9.8K-token document fixture and three-case means for reasoning. Every FastMTP result reproduced the corresponding embedded-MTP output hashes.

| Quant | Depth | PP tok/s | Document TG | Reasoning TG | vs embedded n2, Doc / Reason | vs MTP-off, Doc / Reason | 
|---|---|---|---|---|---|---|
| Q2_K_P | 3 | 3351.29 | 213.95 | 145.09 | +11.6% / +1.7% | 2.27x / 1.48x | 
| Q3_K_P | 3 | 3317.16 | 216.15 | 137.99 | +20.5% / +8.8% | 2.54x / 1.56x | 
| Q4_K_P | 3 | 3204.98 | 187.26 | 123.52 | +18.0% / +2.3% | 2.67x / 1.71x | 
| Q5_K_P | 3 | 2842.05 | 168.29 | 110.50 | +17.8% / +4.6% | 2.61x / 1.66x | 
| Q6_K_P | 3 | 3081.90 | 156.57 | 103.51 | +26.5% / +13.8% | 2.95x / 1.91x | 
| Q8_K_P | 3 | 3285.86 | 138.18 | 90.07 | +35.2% / +21.1% | 3.02x / 1.93x | 
| IQ2_M | 3 | 3050.94 | 219.19 | 135.00 | +13.8% / +0.4% | 2.33x / 1.39x | 
| IQ3_M | 3 | 3269.27 | 204.98 | 128.45 | +21.5% / +7.9% | 2.40x / 1.45x | 
| IQ3_XS | 3 | 3165.75 | 210.64 | 138.34 | +19.1% / +5.9% | 2.38x / 1.51x | 
| IQ4_XS | 3 | 3445.30 | 211.09 | 135.77 | +21.7% / +9.3% | 2.68x / 1.66x | 

The full-window gate used the final scrubbed Q3_K_P and FastMTP files: **190,000 uncached prompt tokens plus 64 generated tokens completed at 1613.81 PP tok/s and 131.81 TG tok/s, with 92.0% draft acceptance and no truncation inside the configured maximum native context.**

Single-run reference results from the final public files at a configured max token context, full CUDA offload, `--no-mmap`, the official thinking sampler, and embedded MTP. The workload used an uncached 9.8K-token document-continuation prompt followed by 512 generated tokens.


---
id: collect-mindstudio/mindstudio/dots3-note-local-deployment
title: "How to Deploy dots3-note Preview Locally with vLLM or SGLang"
domain: mindstudio
role: reference
task: article
actors: ["Hugging Face", "Nvidia", "OpenAI", "SGLang", "vLLM"]
dates: ["2026-09-23"]
keywords: ["sglang", "vllm", "agent", "attention", "decode", "fp8", "gpu", "gpus", "latency", "memory", "moe", "multimodal"]
source: docs/RAG/Collect RAG/02_mindstudio/dots3-note-local-deployment.md
source_anchor: ""
source_lines: [1, 59]
sha256: d5bdf30477746aad4651f1231910d3e72f08ed43674486f34c55888ca5f72510
---

# How to Deploy dots3-note Preview Locally with vLLM or SGLang

## Metadata

- **Source** : https://www.mindstudio.ai/blog/dots3-note-local-deployment
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This is a hardware and command guide to self-hosting **dots3-note preview**'s **FP8** checkpoint on 8-GPU nodes using **vLLM** or **SGLang**, with speculative decoding. dots3-note preview is a **Mixture-of-Experts (MoE) multimodal model** from **dots studio (Xiaohongshu)**, the first open-weight release in the dots3 family. It has **280B total parameters** with **16B activated per token**, handles text, image, video, and audio input, and supports context lengths up to **512K tokens**. Running it locally means serving the FP8 checkpoint on a single **8-GPU node**.

The model card recommends the FP8 checkpoint (**dots-studio/dots3-note-prev-fp8**) since **BF16** needs significantly more GPU memory for the same setup. vLLM's own example targets **eight NVIDIA H100 GPUs** with tensor-parallel size 8 and expert-parallel size 8. Context length is configurable and should be tuned down from the full 512K ceiling depending on memory, concurrency, and which modalities (image, audio, video) are used, since multimodal inputs add memory overhead beyond raw token count. The architecture explains the GPU demand: **256 routed experts plus one shared expert with top-8 routing**, **45 MoE layers plus one dense layer**, and a mix of **Dense Sparse Attention (DSA)** and **sliding window attention (SWA)** layers in roughly a **1:3 ratio**. Such a model benefits from expert parallelism (EP=8) spreading the 256 experts across all eight GPUs, which is why both official paths default to EP=8 alongside TP=8.

**SGLang deployment** uses the prebuilt Docker image **lmsysorg/sglang:dev-dots3-note**, which handles dependency versions and pulls the checkpoint from Hugging Face on first run. Key flags: `--attention-backend fa3` (backend for prefill, decode, and draft attention), `--moe-a2a-backend deepep` (all-to-all communication for MoE layers spread across GPUs), `--enable-dp-attention`, `--dp-size 8`, `--tp-size 8`, `--ep-size 8`, `--page-size 64`, `--trust-remote-code`, `--enable-multimodal`, and the NEXTN speculative flags. Prefill CUDA graph capture isn't supported yet on this release. Native source support lives in SGLang **PR #33829**; until merged, the Docker image or building from the PR revision are the supported routes. The project maintains a dedicated cookbook page. Optional flags: `--language-only` (skip vision/audio encoders) and `--tool-call-parser dots` (OpenAI-compatible tool calling).

**vLLM deployment** has native support on main branch, so a recent **nightly build** is the practical requirement rather than waiting on a PR. The reference command targets eight H100s: `vllm serve dots-studio/dots3-note-prev-fp8 --tensor-parallel-size 8 --enable-expert-parallel --moe-backend deep_gemm --max-model-len 262144`. The context length is set to **262144** (half the model's 512K max), a reasonable middle ground for memory versus context on an 8-GPU node; adjust `--max-model-len` to workload. Optional flags mirror SGLang: `--language-model-only` for text-only serving, and `--enable-auto-tool-choice --tool-call-parser dots` for automatic tool calling.

**MTP/NEXTN speculative decoding** is supported on both frameworks. The architecture includes a dedicated MTP component (**one shared layer, 1.13B parameters**) built to support this. A small draft mechanism proposes several tokens ahead, and the full model verifies them in a single forward pass. According to the model's deployment docs, enabling MTP/NEXTN can **reduce time-per-output-token (TPOT) by more than 50%**, a meaningful latency win for chat or agent loops. In SGLang this uses `--speculative-algorithm NEXTN` plus `--speculative-num-steps`, `--speculative-eagle-topk`, `--speculative-num-draft-tokens`, and `--speculative-draft-model-path` (pointing back at the same FP8 checkpoint); in vLLM it's `--speculative-config '{"method":"mtp","num_speculative_tokens":3}'`. It's optional, so it's reasonable to get a baseline working first.

**Transformers** support (Hugging Face **PR #47844**) is good for quick single-GPU testing, not production multi-GPU serving. It requires compatible PyTorch/torchvision for your NVIDIA driver, plus **torchcodec** and **FFmpeg** for audio/video input, and installing Transformers from the PR branch. A minimal example loads the checkpoint with **AutoModelForMultimodalLM** and runs `generate()`. For real traffic or concurrent requests, the model card points to SGLang or vLLM for OpenAI-compatible serving.

## Key points

- dots3-note preview is a 280B MoE multimodal model with 16B active, up to 512K context, text/image/video/audio.
- The FP8 checkpoint is the recommended self-hosting path; BF16 needs much more memory.
- Both SGLang and vLLM deploy on one 8-GPU node with TP=8 and EP=8.
- MTP/NEXTN speculative decoding can cut time-per-output-token by more than 50%.
- Tool calling is supported via a dedicated `dots` parser flag on both frameworks.
- Transformers is for quick testing only; use SGLang/vLLM for production.
- SGLang support is in an open PR (#33829); vLLM has native main-branch support (nightly).

## Technical data / figures

| Item | Value |
|---|---|
| Total parameters | 280B (MoE) |
| Active parameters | 16B per token |
| Architecture | 256 routed experts + 1 shared, top-8 routing |
| Layers | 45 MoE + 1 dense |
| Attention | DSA + SWA (~1:3 ratio) |
| Context | Up to 512K tokens |
| Modalities | Text, image, video, audio |
| Recommended checkpoint | dots3-note-prev-fp8 |
| Baseline deployment | 8 GPUs (H100 in vLLM example) |
| Parallelism | TP=8, EP=8 |
| SGLang image | lmsysorg/sglang:dev-dots3-note |
| SGLang PR | #33829 |
| vLLM context example | 262144 |
| MTP component | 1 shared layer, 1.13B params |
| MTP TPOT reduction | >50% |
| HF Transformers PR | #47844 |

## Why this source matters for the RAG

It provides exact deployment commands and architectural details for a large multimodal MoE model, including FP8 quantization and speculative decoding. It is a strong reference for self-hosting frontier-scale multimodal models with vLLM/SGLang.

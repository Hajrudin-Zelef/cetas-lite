---
id: collect-240926-mindstudio/mindstudio/how-to-deploy-dots3-note-preview-locally-with-vllm-or-sglang
title: "how-to-deploy-dots3-note-preview-locally-with-vllm-or-sglang"
domain: mindstudio
role: reference
task: reference
actors: ["Hugging Face", "Nvidia", "OpenAI", "SGLang", "vLLM"]
dates: []
keywords: ["sglang", "vllm", "agent", "agents", "attention", "decode", "fp8", "gpu", "gpus", "latency", "memory", "moe"]
source: docs/RAG/clean_en/mindstudio/how-to-deploy-dots3-note-preview-locally-with-vllm-or-sglang.md
source_anchor: ""
source_lines: [1, 127]
sha256: 2e9a1540dcb4c810bb1823d011d72c83de032260dbb4a358408571cc41840c03
---

# how-to-deploy-dots3-note-preview-locally-with-vllm-or-sglang

<!-- source: https://www.mindstudio.ai/blog/dots3-note-local-deployment -->

## What is dots3-note preview and what does it take to run it?

dots3-note preview is a Mixture-of-Experts multimodal model from dots studio (Xiaohongshu), the first open-weight release in the dots3 family. It has 280B total parameters with 16B activated per token, handles text, image, video, and audio input, and supports context lengths up to 512K tokens. Running it locally means serving the FP8 checkpoint on a single 8-GPU node with either vLLM or SGLang, both of which now carry native or near-native support for the architecture.

## TL;DR

- **dots3-note preview** is a 280B-parameter MoE model with 16B active parameters, built for text, image, video, and audio understanding with up to 512K context.
- The **FP8 checkpoint** (`dots-studio/dots3-note-prev-fp8` ) is the recommended way to self-host it, since BF16 needs significantly more GPU memory for the same setup.
- Both **SGLang** and**vLLM** support one-node deployment on 8 GPUs using tensor parallelism and expert parallelism (TP=8, EP=8).
- **Speculative decoding** via MTP/NEXTN is optional but can cut time-per-output-token by more than 50%, based on figures from the model’s own deployment notes.
- **Tool calling** is available out of the box on both serving frameworks using a dedicated`dots` parser flag.
- Transformers support exists for quick local testing but isn’t meant for production-scale multi-GPU serving. SGLang and vLLM are the frameworks to use for that.
- As of this writing, SGLang and Transformers support ships via open pull requests rather than merged mainline code, so exact install commands matter.

## Other agents ship a demo. Remy ships an app.

Real backend. Real database. Real auth. Real plumbing. Remy has it all.

## What hardware do you need to run dots3-note preview?

The model card recommends serving the FP8 checkpoint on a single 8-GPU node. That’s the baseline configuration for both SGLang and vLLM in the official examples, and vLLM’s own example specifically targets eight NVIDIA H100 GPUs with tensor-parallel size 8 and expert-parallel size 8.

BF16 is also supported but requires substantially more memory to hold the same 280B-parameter, 16B-active MoE architecture, so FP8 is the practical default for anyone without a much larger cluster. Context length is configurable and should be tuned down from the full 512K ceiling depending on available memory, expected concurrency, and which modalities (image, audio, video) you’re actually using, since multimodal inputs add memory overhead beyond raw token count.

The architecture itself explains some of the GPU demand: 256 routed experts plus one shared expert with top-8 routing, 45 MoE layers plus one dense layer, and a mix of Dense Sparse Attention (DSA) and sliding window attention (SWA) layers in roughly a 1:3 ratio. A model this shape benefits from expert parallelism (EP=8) that spreads the 256 experts across all eight GPUs, which is why both official deployment paths default to EP=8 alongside TP=8.

## How do you deploy dots3-note preview with SGLang?

The most direct path is the prebuilt Docker image `lmsysorg/sglang:dev-dots3-note`, which handles dependency versions for you and pulls the checkpoint from Hugging Face automatically on first run:

```
docker run --gpus all --ipc=host -p 8000:8000 \
  lmsysorg/sglang:dev-dots3-note \
  sglang serve \
    --model-path dots-studio/dots3-note-prev-fp8 \
    --served-model-name dots3-note-prev \
    --host 0.0.0.0 \
    --port 8000 \
    --context-length 524288 \
    --enable-dp-attention \
    --dp-size 8 \
    --tp-size 8 \
    --ep-size 8 \
    --moe-dense-tp-size 1 \
    --page-size 64 \
    --trust-remote-code \
    --attention-backend fa3 \
    --moe-a2a-backend deepep \
    --enable-multimodal \
    --speculative-algorithm NEXTN \
    --speculative-num-steps 3 \
    --speculative-eagle-topk 1 \
    --speculative-num-draft-tokens 4 \
    --speculative-draft-model-path dots-studio/dots3-note-prev-fp8
```
A few flags are worth understanding rather than copy-pasting blind. `--attention-backend fa3` sets the backend for prefill, decode, and (when speculative decoding is active) the draft model’s attention as well. `--moe-a2a-backend deepep` controls the all-to-all communication pattern MoE layers need when experts are spread across GPUs. Prefill CUDA graph capture isn’t supported yet, so don’t expect that optimization on this release.

Native source support for dots3-note lives in SGLang pull request #33829. Until that PR merges into mainline, the Docker image or building from the PR revision are the two supported routes. The SGLang project also maintains a dedicated cookbook page with one-node recipes and tuning notes for this model specifically, which is the better reference than generic SGLang docs if you hit configuration edge cases.

Two optional flags are useful in practice: `--language-only` loads just the text model and skips the vision/audio encoders if you don’t need multimodal input, and `--tool-call-parser dots` turns on OpenAI-compatible function/tool calling.

## How do you deploy dots3-note preview with vLLM?

## Remy is new. The platform isn't.

Remy is the latest expression of years of platform work. Not a hastily wrapped LLM.

vLLM has native support for dots3-note preview on its `main` branch, so the practical requirement is a recent nightly build rather than waiting on a PR to merge, as is the case with SGLang. The reference deployment targets eight H100 GPUs:

```
vllm serve dots-studio/dots3-note-prev-fp8 \
  --served-model-name dots3-note-prev \
  --host 0.0.0.0 \
  --tensor-parallel-size 8 \
  --enable-expert-parallel \
  --moe-backend deep_gemm \
  --max-model-len 262144
```
Note the context length here is set to 262144 (half of the model’s max 512K), which is a reasonable middle ground for memory versus context tradeoffs on an 8-GPU node. Increase or decrease `--max-model-len` based on your actual workload.

Optional flags mirror SGLang’s: `--language-model-only` skips the vision and audio encoders for text-only serving, and `--enable-auto-tool-choice --tool-call-parser dots` turns on automatic tool calling in an OpenAI-compatible format.

## What is MTP/NEXTN speculative decoding and is it worth enabling?

Both SGLang and vLLM support speculative decoding for dots3-note preview using a Multi-Token Prediction (MTP) approach, sometimes labeled NEXTN. The model architecture includes a dedicated MTP component (one shared layer, 1.13B parameters) built specifically to support this.

The idea behind speculative decoding is straightforward: a small draft mechanism proposes several tokens ahead, and the full model verifies them in a single forward pass instead of generating one token at a time. When the draft tokens are accepted, this cuts the number of expensive full-model passes needed per output token.

According to the model’s deployment documentation, enabling MTP/NEXTN can reduce time-per-output-token (TPOT) by more than 50%. That’s a meaningful latency win for interactive use cases like chat or agent loops where per-token speed matters.

In SGLang, this is controlled with `--speculative-algorithm NEXTN` plus related tuning flags (`--speculative-num-steps`, `--speculative-eagle-topk`, `--speculative-num-draft-tokens`, and a `--speculative-draft-model-path` that in this case points back at the same FP8 checkpoint). In vLLM, it’s a single JSON config flag: `--speculative-config '{"method":"mtp","num_speculative_tokens":3}'`.

It’s optional in both frameworks, so if you’re debugging a fresh deployment, it’s reasonable to get a baseline working first and add speculative decoding once the rest of the stack is stable.

## Is Transformers a viable way to deploy dots3-note preview?

Transformers support (tracked in Hugging Face pull request #47844) is good for quick local testing on a single GPU or machine, not for production multi-GPU serving. The setup requires installing compatible PyTorch and torchvision builds for your NVIDIA driver, plus `torchcodec` and FFmpeg if you want audio or video input support, then installing Transformers directly from the PR branch since it hasn’t merged yet.

A minimal example loads the FP8 checkpoint with `AutoModelForMultimodalLM` and runs `generate()` directly, which is useful for verifying the model loads and produces sensible output before committing to a full SGLang or vLLM deployment. For anything resembling real traffic or concurrent requests, though, the model card explicitly points to SGLang or vLLM for OpenAI-compatible serving.

## Frequently Asked Questions

### What GPU count does dots3-note preview need?

The documented configurations target a single 8-GPU node, specifically NVIDIA H100s in vLLM’s reference example, running the FP8 checkpoint with tensor parallelism and expert parallelism both set to 8.

### Does dots3-note preview support tool calling?

Yes. Both SGLang and vLLM support OpenAI-compatible tool/function calling through a dedicated `dots` tool-call parser flag (`--tool-call-parser dots` in SGLang, and `--enable-auto-tool-choice --tool-call-parser dots` in vLLM).

## Seven tools to build an app. Or just Remy.

Editor, preview, AI agents, deploy — all in one tab. Nothing to install.

### What’s the difference between the BF16 and FP8 checkpoints?

The FP8 checkpoint (`dots-studio/dots3-note-prev-fp8`) is quantized and needs less GPU memory to serve, which is why it’s the recommended default for one-node deployment. The BF16 checkpoint offers full precision but needs considerably more memory for the same 280B-parameter architecture.

### Can I run dots3-note preview without a GPU cluster?

The officially documented deployment paths (SGLang and vLLM) assume an 8-GPU node. Transformers can load the model for basic single-machine testing, but it’s not positioned as a substitute for multi-GPU serving in the model’s own deployment guidance.

### What is MTP/NEXTN speculative decoding used for here?

It’s an optional feature that uses the model’s built-in Multi-Token Prediction layer to draft multiple tokens ahead and verify them together, which the documentation states can reduce time-per-output-token by more than 50% when enabled on either SGLang or vLLM.

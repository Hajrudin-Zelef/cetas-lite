---
id: collect-240926-mindstudio/mindstudio/how-to-run-tencent-s-hy4-preview-locally-with-vllm-or-sglang-1
title: "how-to-run-tencent-s-hy4-preview-locally-with-vllm-or-sglang"
domain: mindstudio
role: reference
task: reference
actors: ["DeepSeek", "Hugging Face", "Moonshot", "OpenAI", "SGLang", "Z.ai", "vLLM"]
dates: []
keywords: ["sglang", "vllm", "apache", "attention", "consumer", "context window", "deepseek", "exploit", "fp8", "glm", "gpu", "gpus"]
source: docs/RAG/clean_en/mindstudio/how-to-run-tencent-s-hy4-preview-locally-with-vllm-or-sglang.md
source_anchor: ""
source_lines: [1, 101]
sha256: cccd4419c525361ca7aa3676d6f67d4ab8c801f5181cc2c5fbc0d7b3bee467bf
---

# how-to-run-tencent-s-hy4-preview-locally-with-vllm-or-sglang

<!-- source: https://www.mindstudio.ai/blog/run-tencent-hy4-preview-locally -->

## What is Tencent’s Hy4 preview model?

Hy4 preview is Tencent’s newest open-weight flagship model, a Mixture-of-Experts (MoE) architecture with 770 billion total parameters and 49 billion activated per token. It ships under Apache 2.0, with both a full-precision release and an FP8 quantized version, and it’s designed to be served through standard inference engines like vLLM and SGLang rather than run as a raw checkpoint. Running it locally means self-hosting, not spinning up a lightweight chatbot on a laptop. This is a multi-GPU deployment exercise.

## TL;DR

- Hy4 preview is a **770B-parameter MoE model** with only 49B parameters activated per token, spread across 78 layers where the first layer is dense and the remaining 77 use MoE with 256 routed experts plus 1 shared expert.
- The model uses **Gated DeepSeek Sparse Attention (Gated DSA)** with IndexCache for cross-layer index reuse, plus identity Hyper-Connections (iHC) on the residual pathway, architectural choices borrowed and adapted from DeepSeek and GLM.
- Tencent ships an **FP8 quantized checkpoint** (Hy4-preview-FP8) specifically for production serving, alongside the full-weight release, both distributed via Hugging Face, ModelScope, GitCode, and CNB.
- Official deployment paths run through **prebuilt Docker images** :`vllm/vllm-openai:hy4-preview` and`lmsysorg/sglang:hy4-preview` , both configured for 8-way tensor parallelism.
- The model supports **speculative decoding** out of the box via a built-in MTP (multi-token prediction) layer, roughly 10B total parameters with 0.7B activated, cutting latency on long generations.
- Context length is rated at **1 million tokens** , and the model defaults to a “high” reasoning effort mode that can be toggled off for faster, more direct responses.
- Tencent explicitly labels this a **preview with known limitations** , including a tendency to over-reason and over-verify on simpler tasks, similar to issues seen in the earlier Hy3 preview.

## Remy doesn't build the plumbing. It inherits it.

Remy ships with all of it from MindStudio — so every cycle goes into the app you actually want.

## What hardware do you need to run Hy4 preview?

Both official deployment recipes assume `--tensor-parallel-size 8` or `--tp-size 8`, meaning an 8-GPU node is the baseline configuration Tencent tested and documented. This isn’t a model you run on a single consumer GPU, or even a single high-end workstation card. With 770B total parameters, even the FP8 checkpoint requires substantial aggregate VRAM across multiple accelerators, which is why Tencent’s own examples default straight to 8-way tensor parallelism.

Practically, this puts Hy4 preview in the same deployment category as other frontier-scale open-weight MoE models: it’s built for multi-GPU servers or cloud instances with high-bandwidth interconnects, not local desktop inference. Anyone planning a deployment should budget for an 8-GPU cluster with enough interconnect bandwidth to keep tensor-parallel communication from becoming the bottleneck, since MoE models move a lot of data between experts during inference.

## How do you deploy Hy4 preview with vLLM?

Tencent provides a dedicated Docker image tagged `vllm/vllm-openai:hy4-preview`, built specifically to support this model’s architecture (including its custom attention backend and tool-calling format). The documented launch command looks like this:

```
docker run --gpus all \
  -p 8000:8000 \
  --ipc=host \
  -v ~/.cache/huggingface:/root/.cache/huggingface \
  vllm/vllm-openai:hy4-preview tencent/Hy4-preview-FP8 \
    --tensor-parallel-size 8 \
    --speculative-config '{"num_speculative_tokens":3,"method":"mtp"}' \
    --attention-backend FLASHMLA_SPARSE \
    --tool-call-parser hy_v4 \
    --reasoning-parser hy_v4 \
    --enable-auto-tool-choice \
    --port 8000 \
    --served-model-name hy4-preview
```
A few details matter here. The `--speculative-config` flag activates the model’s built-in MTP layer for speculative decoding, using 3 speculative tokens per step, which should meaningfully speed up generation without changing output quality. The `FLASHMLA_SPARSE` attention backend is required because Hy4 preview uses Gated DSA rather than standard dense attention, so a generic vLLM config won’t correctly exploit the model’s sparse attention pattern. The `hy_v4` parsers handle the model’s specific tool-calling and reasoning-trace formats, which differ from more common templates like ChatML.

Once the container is running, the model is exposed through an OpenAI-compatible API on port 8000, so any existing OpenAI SDK client can talk to it by just changing the base URL.

## How do you deploy Hy4 preview with SGLang?

SGLang offers a parallel deployment path via `lmsysorg/sglang:hy4-preview`, described as multi-arch (supporting both x86 and Arm). The launch command is structurally similar to the vLLM one but uses SGLang’s own flags for speculative decoding:

```
docker pull lmsysorg/sglang:hy4-preview
docker run --gpus all --ipc=host -p 8000:8000 lmsysorg/sglang:hy4-preview \
  python3 -m sglang.launch_server \
    --model tencent/Hy4-preview-FP8 \
    --tp-size 8 \
    --reasoning-parser auto \
    --tool-call-parser auto \
    --speculative-algorithm NEXTN \
    --speculative-num-steps 3 \
    --speculative-eagle-topk 1 \
    --speculative-num-draft-tokens 4 \
    --port 8000 \
    --served-model-name hy4-preview
```
The main functional difference from vLLM here is that SGLang’s reasoning and tool-call parsers can be set to `auto`, letting the server infer format handling rather than requiring an explicit `hy_v4` parser name. The speculative decoding setup (`NEXTN` algorithm with eagle-style top-k drafting) is SGLang’s mechanism for using the same underlying MTP layer for faster generation, tuned with different step and draft-token counts than the vLLM config.

### Everyone else built a construction worker.

We built the contractor.

    One file at a time.

UI, API, database, deploy.

Either path lands you at the same place: a local OpenAI-compatible endpoint on port 8000, ready to accept standard chat completion requests.

## What can you actually do once it’s running?

After deployment, interacting with Hy4 preview is identical to using any other OpenAI-compatible endpoint:

```
from openai import OpenAI
client = OpenAI(base_url="http://127.0.0.1:8000/v1", api_key="EMPTY")
response = client.chat.completions.create(
    model="hy4-preview",
    messages=[{"role": "user", "content": "Hello! Can you briefly introduce yourself?"}],
    temperature=0.9,
    top_p=1.0,
)
print(response.choices[0].message.content)
```
Tencent recommends `temperature=0.9` and `top_p=1.0` as defaults. By default the model runs in a “high” reasoning effort mode, meaning it produces deep chain-of-thought before answering, which suits math, coding, and multi-step reasoning tasks but adds latency for simple queries. For faster, more direct answers, you can pass `extra_body={"chat_template_kwargs": {"reasoning_effort": "no_think"}}` to skip the extended reasoning trace.

## Is Hy4 preview worth deploying right now?

That depends on what you’re optimizing for. On the capability side, Tencent’s internal blind evaluation, 163 experts rating outputs across 203 real engineering tasks, put Hy4 preview slightly ahead of GLM 5.3 (2.99 vs. 2.92 average score, winning 46.8% of head-to-head comparisons) and slightly ahead of Kimi K3 (2.99 vs. 2.94, winning 51.2%). Tencent also describes it as the largest generation-over-generation improvement they’ve measured for this model line, built on expanded pre-training, a larger post-training run, and a 1 million token context window.


---
id: collect-240926-huggingface/huggingface/unsloth-nvidia-nemotron-3-5-lightning-30b-a3b-gguf-hugging-face-2
title: "unsloth-nvidia-nemotron-3-5-lightning-30b-a3b-gguf-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["Nvidia", "OpenAI", "SGLang", "TensorRT-LLM", "vLLM"]
dates: ["2026-11-08"]
keywords: ["nvidia", "agents", "blackwell", "data centre", "diffusion", "distribution", "fine-tuning", "gpu", "gpus", "inference", "memory", "moe"]
source: docs/RAG/clean_en/huggingface/unsloth-nvidia-nemotron-3-5-lightning-30b-a3b-gguf-hugging-face.md
source_anchor: ""
source_lines: [117, 278]
sha256: 92d7a355eb7f984675e916477ac9c50cb37d60a2060bf05cd4682553c36352bc
---

# unsloth-nvidia-nemotron-3-5-lightning-30b-a3b-gguf-hugging-face

- The model underwent a continued pre-training phase to train its **Multi-Token Prediction (MTP)** layers. In this stage, MTP heads learn to predict multiple future tokens, providing richer training signals to the base model. This phase aligns the MTP layers with the base model's distribution.

Stage 3: Supervised Fine-Tuning

- The model was further fine-tuned on synthetic code, math, science, tool calling, instruction following, structured outputs, and general knowledge data. This stage incorporated data designed to support long-range retrieval and multi-document aggregation.

Stage 4: Reinforcement Learning

- The model underwent multi-environment reinforcement learning using GRPO (Group Relative Policy Optimization) across math, code, science, instruction following, multi-step tool use, multi-turn conversations, and structured output environments. It utilized an asynchronous RL architecture that decouples training from inference and leverages MTP to accelerate rollout generation.
- Software used for reinforcement learning: NeMo RL, NeMo Gym

NVIDIA-Nemotron-3.5-Lightning-30B-A3B-BF16 is a result of the above work.

- **Input Type(s):** Text
- **Input Format(s):** String
- **Input Parameters:** One-Dimensional (1D): Sequences
- **Other Properties Related to Input:** Maximum context length up to 1M tokens. Supported languages include English, Spanish, French, German, Italian, and Japanese.

- **Output Type(s):** Text
- **Output Format:** String
- **Output Parameters:** One-Dimensional (1D): Sequences
- **Other Properties Related to Output:** Maximum context length up to 1M tokens

Our AI models are designed and/or optimized to run on NVIDIA GPU-accelerated systems. By leveraging NVIDIA's hardware (e.g. GPU cores) and software frameworks (e.g., CUDA libraries), the model achieves faster training and inference times compared to CPU-only solutions.

- **Runtime Engine(s):** PyTorch
- **Supported Hardware Microarchitecture Compatibility:** NVIDIA Ampere - A100; NVIDIA Blackwell; NVIDIA Hopper - H100-80GB
- **Preferred/Supported Operating System(s):** Linux

The integration of foundation and fine-tuned models into AI systems requires additional testing using use-case-specific data to ensure safe and effective deployment. Following the V-model methodology, iterative testing and validation at both unit and system levels are essential to mitigate risks, meet technical and functional requirements, and ensure compliance with safety and ethical standards before deployment.

- GA (08/11/2026)

All deployment snippets below assume:

```
export MODEL_CKPT=nvidia/NVIDIA-Nemotron-3.5-Lightning-30B-A3B-BF16
```
And for DSpark:

```
export DSPARK_CKPT=nvidia/NVIDIA-Nemotron-3.5-Lightning-30B-A3B-NVFP4-DSpark
```
The BF16 recipes below cover vLLM. For TensorRT-LLM, SGLang, W4A16 (Blackwell / Hopper / Ampere from a single checkpoint), and DGX Spark (DSpark) recipes, see the NVFP4 card.


Lightning 3.5 ships with two external draft models for speculative decoding as well as MTP (Multi-Token Prediction). While we currently recommend DSpark for all cases - your usecase may align with DFlash and MTP:

- **DSpark:** A semi-autoregressive speculative-decoding drafter that proposes a whole block of candidate tokens in a single forward pass from a parallel backbone. This is recommended for DGX Spark, as well as low-concurrency data centre deployments.
- **DFlash:** A speculative-decoding drafter that uses a lightweight block-diffusion model to generate an entire draft block in one forward pass.
- **MTP:** A modeling technique that trains the network to predict several future tokens at each position instead of only the next one.

- vLLM Nightly: `vllm/vllm-openai:v0.27.1`

For max throughput deployments, use the following configuration, no speculative decoding strategy is best for this serving configuration, and due to memory constraints the Mamba cache `dtype` is set as FP16:

```
vllm serve --model $MODEL_CKPT \
    --max-num-seqs 128 \
    --enable-prefix-caching \
    --async-scheduling \
    --mamba-backend flashinfer \
    --mamba-ssm-cache-dtype float16 \
    --enable-mamba-cache-stochastic-rounding \
    --mamba-cache-philox-rounds 5 \
    --reasoning-parser nemotron_v3 \
    --tool-call-parser qwen3_coder \
    --enable-auto-tool-choice
```
For long-context, multi-GPU serving (TP8 with expert parallelism):

```
VLLM_ALLOW_LONG_MAX_MODEL_LEN=1 vllm serve --model $MODEL_CKPT \
    --moe-backend flashinfer_cutlass \
    --mamba-backend flashinfer \
    --enable-prefix-caching \
    --mamba-cache-mode align \
    --max-model-len 1048576 \
    --enable-expert-parallel \
    --tensor-parallel-size 8 \
    --reasoning-parser nemotron_v3 \
    --tool-call-parser qwen3_coder \
    --enable-auto-tool-choice
```
```
VLLM_ALLOW_LONG_MAX_MODEL_LEN=1 vllm serve --model $MODEL_CKPT \
    --max-num-seqs 128 \
    --max-model-len 1048576 \
    --max-num-batched-tokens 10240 \
    --no-enable-prefix-caching \
    --async-scheduling \
    --speculative_config.model $DSPARK_CKPT \
    --speculative_config.num_speculative_tokens 5 \
    --mamba-backend flashinfer \
    --reasoning-parser nemotron_v3 \
    --tool-call-parser qwen3_coder \
    --enable-auto-tool-choice
```
- **Context Length:** If you're memory-constrained — or want more KV-cache headroom at high concurrency — lower`--max-model-len` to match your workload and drop`VLLM_ALLOW_LONG_MAX_MODEL_LEN=1` .

The examples below use the OpenAI-compatible client and work with the serving backend above. Recommended sampling settings are **Temperature 1.0** and **Top_P 0.95**.

The vLLM snippets above register the model as `nvidia/NVIDIA-Nemotron-3.5-Lightning-30B-A3B-BF16` via `--served-model-name`.

```
from openai import OpenAI
client = OpenAI(base_url="http://localhost:8000/v1", api_key="EMPTY")
MODEL = "nvidia/NVIDIA-Nemotron-3.5-Lightning-30B-A3B-BF16"
```
Lightning 3.5 exposes reasoning control through chat-template kwargs: thinking enabled (the default), thinking disabled for direct answers, and a runtime thinking budget.

## **Reasoning ON / OFF and streaming examples: Click to expand!**

**Reasoning ON (default)**

```
response = client.chat.completions.create(
    model=MODEL,
    messages=[{"role": "user", "content": "Write a haiku about GPUs"}],
    max_tokens=16000,
    temperature=1.0,
    top_p=0.95,
    extra_body={"chat_template_kwargs": {"enable_thinking": True}}
)
print(response.choices[0].message.content)
```
**Reasoning OFF**

```
response = client.chat.completions.create(
    model=MODEL,
    messages=[{"role": "user", "content": "What is the capital of Japan?"}],
    max_tokens=16000,
    temperature=1.0,
    top_p=0.95,
    extra_body={"chat_template_kwargs": {"enable_thinking": False}}
)
print(response.choices[0].message.content)
```
**Streaming**

```
stream = client.chat.completions.create(
    model=MODEL,
    messages=[{"role": "user", "content": "Explain speculative decoding in two sentences"}],
    max_tokens=16000,
    temperature=1.0,
    top_p=0.95,
    stream=True,
)
for chunk in stream:
    print(chunk.choices[0].delta.content or "", end="", flush=True)
```
For vLLM, add the following to any serve command above:

```
    --enable-auto-tool-choice \
    --tool-call-parser qwen3_coder \
    --reasoning-parser nemotron_v3
```
**NOTE:** For coding agents, add `extra_body={"chat_template_kwargs": {"force_nonempty_content": True}}` to the API call, as shown below.


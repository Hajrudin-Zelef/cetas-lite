---
id: collect-240926-huggingface/huggingface/nvidia-nvidia-nemotron-3-ultra-550b-a55b-bf16-hugging-face-2
title: "Set the IP for the head node in RAY_HEAD_IP"
domain: huggingface
role: reference
task: reference
actors: ["China", "Nvidia", "SGLang", "vLLM"]
dates: []
keywords: ["agentic", "attention", "blackwell", "distillation", "embeddings", "fine-tuning", "gpu", "gpus", "hbm", "inference", "kv cache", "memory"]
source: docs/RAG/clean_en/huggingface/nvidia-nvidia-nemotron-3-ultra-550b-a55b-bf16-hugging-face.md
source_anchor: ""
source_lines: [96, 213]
sha256: 9a1d0d115840b72b43baab17104d80f1d7b03b44f86f91cd7cc1f0ae6f2505e0
---

# Set the IP for the head node in RAY_HEAD_IP

The model utilizes the **LatentMoE** architecture, where tokens are projected into a smaller latent dimension for expert routing and computation, improving accuracy per byte. The Ultra model is pre-trained using an NVFP4 recipe — sharing the quantization-aware pre-training approach pioneered in the Nemotron 3 family. The majority of linear layers use NVFP4 for weights, activations, and gradients, while select layers (including latent projections, MTP layers, QKV/attention projections, and embeddings) are maintained in BF16 or MXFP8 for training stability. The model includes **Multi-Token Prediction (MTP)** layers using a shared-weight design across prediction heads. This improves training signal quality, enables faster inference via native speculative decoding, and supports more stable autoregressive drafting at longer draft lengths compared to independently trained offset heads.

Stage 1: Pre-Training

- NVIDIA-Nemotron-3-Ultra-550B-A55B-Base-BF16 model was pre-trained for approximately 20T tokens using crawled and synthetic code, math, science, and general knowledge data. Training leveraged an NVFP4 recipe for efficiency. All datasets are disclosed in the Training and Evaluation Datasets section of this document. Major portions of the pre-training corpus are released in the Nemotron-Pre-Training-Datasets collection.
- Software used for pre-training: Megatron-LM

Stage 2: Supervised Fine-Tuning

- The model was further fine-tuned on synthetic code, math, science, tool calling, instruction following, structured outputs, and general knowledge data. This stage incorporated data designed to support long-range retrieval and multi-document aggregation. All datasets are disclosed in the Training and Evaluation Datasets section of this document. Major portions of the fine-tuning corpus are released in the Nemotron-Post-Training-v3 collection. Data Designer is one of the libraries used to prepare these corpora.

Stage 3: Reinforcement Learning

- The model underwent multi-environment reinforcement learning using asynchronous GRPO (Group Relative Policy Optimization) across math, code, science, instruction following, multi-step tool use, multi-turn conversations, and structured output environments. It utilized an asynchronous RL architecture that fully decouples training from inference across separate GPU devices, leveraging in-flight weight updates and MTP to accelerate rollout generation. Conversational quality was further refined through RLHF. All datasets are disclosed in the Training and Evaluation Datasets section of this document. The RL environments and datasets are released as part of NeMo Gym.
- Software used for reinforcement learning: NeMo RL, NeMo Gym

Stage 4: Multi-Domain On-Policy Distillation (MOPD)

- The model underwent **Multi-Domain On-Policy Distillation (MOPD)** to improve reasoning across many task types while staying efficient. This technique uses strong teacher models to guide training on the model's own generated attempts (on-policy rollouts), helping recover accuracy and improve performance across coding, math, instruction following, tool use, and agentic workflows. By distilling teacher signal onto the student's own trajectories rather than offline traces, MOPD better aligns the student's behavior with what it would actually produce at inference time, yielding stronger gains than purely off-policy distillation.

NVIDIA-Nemotron-3-Ultra-550B-A55B-BF16 model is a result of the above work.

The end-to-end training recipe is available in the NVIDIA Nemotron Developer Repository. Evaluation results can be replicated using the NeMo Evaluator SDK. Data Designer is one of the libraries used to prepare the pre and post training datasets. More details on the datasets and synthetic data generation methods can be found in the technical report NVIDIA Nemotron 3 Ultra Technical Report.

- **Input Type(s):** Text
- **Input Format(s):** String
- **Input Parameters:** One-Dimensional (1D): Sequences
- **Other Properties Related to Input:** Maximum context length up to 1M tokens. Supported languages include: English, French, Spanish, Italian, German, Japanese, Hindi, Korean, Brazilian Portuguese, and Chinese

- **Output Type(s):** Text
- **Output Format:** String
- **Output Parameters:** One-Dimensional (1D): Sequences
- **Other Properties Related to Output:** Maximum context length up to 1M tokens

Our AI models are designed and optimized to run on NVIDIA GPU-accelerated systems. By leveraging NVIDIA's hardware (e.g. GPU cores) and software frameworks (e.g., CUDA libraries), the model achieves faster training and inference times compared to CPU-only solutions.

- Runtime Engine(s): NeMo 26.04.01
- Supported Hardware Microarchitecture Compatibility: NVIDIA Ampere - A100; NVIDIA Blackwell; NVIDIA Hopper - H100-80GB
- Operating System(s): Linux

The integration of foundation and fine-tuned models into AI systems requires additional testing using use-case-specific data to ensure safe and effective deployment. Following the V-model methodology, iterative testing and validation at both unit and system levels are essential to mitigate risks, meet technical and functional requirements, and ensure compliance with safety and ethical standards before deployment.

- v1.0 - GA

The Ultra BF16 checkpoint is a frontier-scale model. The minimum recommended hardware is:

- **Single-node:** 8× B200 (≈1.5 TB aggregate HBM — fits BF16 weights plus KV cache with headroom)
- **Multi-node:** ≥8 GPUs across H100 / H200 / GB200 / GB300, orchestrated with Ray v2

All deployment snippets below default to **port 8000**, with **chunked prefill** and **MTP (5 speculative tokens)** enabled.

The recommended multi-processing backend for multi-node BF16 deployments is Ray v2. Below is a template for launching a Ray cluster:

```
# Set the IP for the head node in RAY_HEAD_IP
export RAY_HEAD_IP=<head_node_ip>
export RAY_PORT=6379
export RAY_ADDRESS=${RAY_HEAD_IP}:${RAY_PORT}
# Start Ray head node (vLLM/SGLang will run on this node)
ray start --head --node-ip-address=${RAY_HEAD_IP} --port=${RAY_PORT}
# Start Ray worker node(s)
ray start --address=${RAY_HEAD_IP}:${RAY_PORT} --block
# Verify Ray cluster is ready
ray status --address=${RAY_HEAD_IP}:${RAY_PORT}
```
`ray[cgraph]` is required: `uv pip install "ray[cgraph]"`

Recommended container: `vllm/vllm-openai:v0.22.0`.

For more detailed information, please see this cookbook.

```
export MODEL_CKPT=PATH/TO/MODEL/CHECKPOINT
```
**8× B200 single-node deployment:**

```
docker run -d --name nemotron-ultra-vllm \
  --gpus all \
  --ipc=host \
  --network=host \
  --shm-size=16g \
  --ulimit memlock=-1 \
  --ulimit stack=67108864 \
  -v $MODEL_CKPT:/model:ro \
  -e VLLM_WORKER_MULTIPROC_METHOD=spawn \
  -e SAFETENSORS_FAST_GPU=1 \
  -e NVIDIA_TF32_OVERRIDE=1 \
  -e VLLM_LOGGING_LEVEL=INFO \
  vllm/vllm-openai:v0.22.0 \
  /model \
  --host 0.0.0.0 \
  --port 8000 \
  --served-model-name nvidia/nemotron-3-ultra \
  --trust-remote-code \
  --tensor-parallel-size 8 \
  --enable-expert-parallel \
  --dtype bfloat16 \
  --max-model-len 262144 \
  --gpu-memory-utilization 0.90 \
  --max-num-seqs 16 \
  --max-num-batched-tokens 32768 \
  --enable-chunked-prefill \
  --enable-prefix-caching \
  --reasoning-parser nemotron_v3 \
  --enable-auto-tool-choice \
  --tool-call-parser qwen3_coder \
  --mamba-ssm-cache-dtype float16 \   
  --mamba-backend flashinfer \     
  --enable-mamba-cache-stochastic-rounding \   
  --mamba-cache-philox-rounds 5 \
  --speculative-config '{"method": "nemotron_h_mtp", "num_speculative_tokens": 5}' \
  --model-loader-extra-config '{"enable_multithread_load": true, "num_threads": 96}'
```
**Multi-node deployment (e.g. 2× 4×GB300 with Ray):**

After launching the Ray head and worker per the multi-node setup above:


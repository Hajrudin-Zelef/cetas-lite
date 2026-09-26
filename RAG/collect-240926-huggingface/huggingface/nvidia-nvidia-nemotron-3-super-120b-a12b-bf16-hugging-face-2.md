---
id: collect-240926-huggingface/huggingface/nvidia-nvidia-nemotron-3-super-120b-a12b-bf16-hugging-face-2
title: "with uv: uv pip install vllm==0.18.1 --torch-backend=auto"
domain: huggingface
role: reference
task: reference
actors: ["China", "Nvidia", "OpenAI", "SGLang", "vLLM"]
dates: []
keywords: ["vllm", "agent", "agents", "attention", "blackwell", "fine-tuning", "fp8", "gpu", "gpus", "hbm", "inference", "memory"]
source: docs/RAG/clean_en/huggingface/nvidia-nvidia-nemotron-3-super-120b-a12b-bf16-hugging-face.md
source_anchor: ""
source_lines: [107, 296]
sha256: f65b8c1f7d3ef633dbab0b7b2bc74fc2ef4cce4552388fa6c9dc91a157adad3c
---

# with uv: uv pip install vllm==0.18.1 --torch-backend=auto

- The model was further fine-tuned on synthetic code, math, science, tool calling, instruction following, structured outputs, and general knowledge data. This stage incorporated data designed to support long-range retrieval and multi-document aggregation. All datasets are disclosed in the Training and Evaluation Datasets section of this document. Major portions of the fine-tuning corpus are released in the Nemotron-Post-Training-v3 collection. Data Designer is one of the libraries used to prepare these corpora.

Stage 3: Reinforcement Learning

- The model underwent multi-environment reinforcement learning using asynchronous GRPO (Group Relative Policy Optimization) across math, code, science, instruction following, multi-step tool use, multi-turn conversations, and structured output environments. It utilized an asynchronous RL architecture that fully decouples training from inference across separate GPU devices, leveraging in-flight weight updates and MTP to accelerate rollout generation. Conversational quality was further refined through RLHF. All datasets are disclosed in the *Training and Evaluation Datasets* section of this document. The RL environments and datasets are released as part of NeMo Gym.
- Software used for reinforcement learning: NeMo RL, NeMo Gym

NVIDIA-Nemotron-3-Super-120B-A12B-BF16 model is a result of the above work.

The end-to-end training recipe is available in the NVIDIA Nemotron Developer Repository. Evaluation results can be replicated using the NeMo Evaluator SDK. Data Designer is one of the libraries used to prepare the pre and post training datasets. More details on the datasets and synthetic data generation methods can be found in the technical report NVIDIA Nemotron 3 Super Technical Report.

- **Input Type(s):** Text
- **Input Format(s):** String
- **Input Parameters:** One-Dimensional (1D): Sequences
- **Other Properties Related to Input:** Maximum context length up to 1M tokens. Supported languages include: English, French, German, Italian, Japanese, Spanish, and Chinese

- **Output Type(s):** Text
- **Output Format:** String
- **Output Parameters:** One-Dimensional (1D): Sequences
- **Other Properties Related to Output:** Maximum context length up to 1M tokens

Our AI models are designed and optimized to run on NVIDIA GPU-accelerated systems. By leveraging NVIDIA's hardware (e.g. GPU cores) and software frameworks (e.g., CUDA libraries), the model achieves faster training and inference times compared to CPU-only solutions.

- Runtime Engine(s): NeMo 25.11.01
- Supported Hardware Microarchitecture Compatibility: NVIDIA Ampere - A100; NVIDIA Blackwell; NVIDIA Hopper - H100-80GB
- Operating System(s): Linux

- v1.0 - GA

For each inference backend, you'll need the custom `super_v3` reasoning parser. Download it with:

```
wget https://huggingface.co/nvidia/NVIDIA-Nemotron-3-Super-120B-A12B-BF16/raw/main/super_v3_reasoning_parser.py
```
OR

```
curl -O https://huggingface.co/nvidia/NVIDIA-Nemotron-3-Super-120B-A12B-BF16/raw/main/super_v3_reasoning_parser.py
```
For advanced deployment configurations, visit this resource.

For more detailed information, please see this cookbook.

```
pip install vllm==0.18.1
# with uv: uv pip install vllm==0.18.1 --torch-backend=auto
export MODEL_CKPT=PATH/TO/MODEL/CHECKPOINT
```
```
# Optional: --enable-expert-parallel
vllm serve $MODEL_CKPT \
  --served-model-name nvidia/nemotron-3-super \
  --async-scheduling \
  --dtype auto \
  --kv-cache-dtype fp8 \
  --tensor-parallel-size 8 \
  --max-model-len 262144 \
  --enable-expert-parallel \
  --swap-space 0 \
  --trust-remote-code \
  --gpu-memory-utilization 0.9 \
  --max-cudagraph-capture-size 128 \
  --enable-chunked-prefill \
  --mamba-ssm-cache-dtype float32 \
  --reasoning-parser nemotron_v3 \
  --enable-auto-tool-choice \
  --tool-call-parser qwen3_coder
```
Context length defaults to 256k above. To use up to 1M, set `VLLM_ALLOW_LONG_MAX_MODEL_LEN=1` and `--max-model-len 1048576`.


**B200/B300 (BF16)**: The larger HBM capacity per device means the BF16 checkpoint fits on 2 GPUs. Set `--tensor-parallel-size 2` and remove `--enable-expert-parallel`. All other flags remain the same.


Container:

```
docker pull lmsysorg/sglang:nightly-dev-cu13-20260316-d852f26c
```
For more detailed information, please see this cookbook.

```
sglang serve \
  --model-path PATH/TO/CHECKPOINT \
  --served-model-name nvidia/nemotron-3-super \
  --trust-remote-code \
  --tp 8 \
  --ep 8 \
  --tool-call-parser qwen3_coder \
  --reasoning-parser nemotron_3
```
Context length defaults to 256k above. To use up to 1M, set `SGLANG_ALLOW_OVERWRITE_LONGER_CONTEXT_LEN=1` and `--context-length 1048576`.


**B200/B300 (BF16)**: The larger HBM capacity per device means the BF16 checkpoint fits on 2 GPUs. Set `--tp 2 --ep 2`. All other flags remain the same.


Container:

```
docker pull nvcr.io/nvidia/tensorrt-llm/release:1.3.0rc8
```
For more detailed information, please see this cookbook.

```
cat > ./extra-llm-api-config.yml << EOF
kv_cache_config:
  enable_block_reuse: false
  free_gpu_memory_fraction: 0.8
  mamba_ssm_cache_dtype: float16
  dtype: fp8
moe_config:
  backend: TRTLLM
cuda_graph_config:
  enable_padding: true
  max_batch_size: 256
enable_attention_dp: true
enable_chunked_prefill: true
num_postprocess_workers: 4
stream_interval: 10
EOF
trtllm-serve PATH/TO/BF16/CHECKPOINT \
  --host 0.0.0.0 \
  --port 8123 \
  --max_batch_size 256 \
  --tp_size 8 --ep_size 8 \
  --max_num_tokens 8192 \
  --trust_remote_code \
  --reasoning_parser nano-v3 \
  --tool_parser qwen3_coder \
  --extra_llm_api_options extra-llm-api-config.yml
```
**B200/B300 (BF16)**: The larger HBM capacity per device means the BF16 checkpoint fits on 2 GPUs. Set `--tp_size 2 --ep_size 2` and reduce `max_batch_size` to `128` in both the config file and the serve command. All other flags remain the same.


The examples below use the OpenAI-compatible client and work with any of the serving backends above.

NOTE: For coding agents add the following to the API call - `extra_body={“chat_template_kwargs”: {“force_nonempty_content”: True}`


```
from openai import OpenAI
client = OpenAI(base_url="http://localhost:8000/v1", api_key="EMPTY")
MODEL = "nvidia/nemotron-3-super"
```
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
**Low-effort reasoning**

Uses significantly fewer reasoning tokens than full thinking mode. Recommended as a starting point before tuning explicit token budgets.

```
response = client.chat.completions.create(
    model=MODEL,
    messages=[{"role": "user", "content": "What is the capital of Japan?"}],
    max_tokens=16000,
    temperature=1.0,
    top_p=0.95,
    extra_body={"chat_template_kwargs": {"enable_thinking": True, "low_effort": True}}
)
print(response.choices[0].message.content)
```
OpenCode is an AI coding agent that runs in your terminal. It connects to any OpenAI-compatible endpoint, making it compatible with all three serving backends above (vLLM, SGLang, and TRT-LLM).

Create or update your `~/.config/opencode/opencode.json`:


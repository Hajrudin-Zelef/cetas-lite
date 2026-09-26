---
id: collect-240926-huggingface/huggingface/nvidia-nvidia-nemotron-3-ultra-550b-a55b-bf16-hugging-face-3
title: "Set the IP for the head node in RAY_HEAD_IP"
domain: huggingface
role: reference
task: reference
actors: ["Nvidia", "OpenAI", "SGLang", "TensorRT-LLM", "vLLM"]
dates: []
keywords: ["agent", "agents", "attention", "blackwell", "fp8", "gpu", "gpus", "latency", "memory", "moe", "nvidia", "parameters"]
source: docs/RAG/clean_en/huggingface/nvidia-nvidia-nemotron-3-ultra-550b-a55b-bf16-hugging-face.md
source_anchor: ""
source_lines: [214, 451]
sha256: b33c31d4e2d6ae3c3b38c5daed6aca8cce7a2bffadf0c81e82c358a2efba965d
---

# Set the IP for the head node in RAY_HEAD_IP

```
# Run on Ray head node
vllm serve $MODEL_CKPT \
  --host 0.0.0.0 \
  --port 8000 \
  --served-model-name nvidia/nemotron-3-ultra \
  --tensor-parallel-size 8 \
  --distributed-executor-backend ray \
  --trust-remote-code \
  --dtype bfloat16 \
  --gpu-memory-utilization 0.90 \
  --max-model-len 262144 \
  --max-num-seqs 256 \
  --max-num-batched-tokens 32768 \
  --enable-chunked-prefill \
  --enable-prefix-caching \
  --reasoning-parser nemotron_v3 \
  --mamba-ssm-cache-dtype float16 \   
  --mamba-backend flashinfer \     
  --enable-mamba-cache-stochastic-rounding \   
  --mamba-cache-philox-rounds 5 \
  --enable-auto-tool-choice \
  --tool-call-parser qwen3_coder \
  --speculative-config '{"method": "nemotron_h_mtp", "num_speculative_tokens": 5}' \
  --kv-cache-dtype fp8 \
  --model-loader-extra-config '{"enable_multithread_load": true, "num_threads": 96}' \
  --compilation-config '{"pass_config": {"fuse_allreduce_rms": false}}' \
  --distributed-timeout-seconds 3600
```
Context length defaults to 256k above. To use up to 1M, set `VLLM_ALLOW_LONG_MAX_MODEL_LEN=1` and `--max-model-len 1048576`.

**Useful environment variables:** `VLLM_FLASHINFER_ALLREDUCE_BACKEND=trtllm`, `VLLM_FLASHINFER_MOE_BACKEND=latency` (TRTLLM-Gen) or `VLLM_FLASHINFER_MOE_BACKEND=throughput` (CUTLASS).

Container (tested on 8× B200):

```
docker pull lmsysorg/sglang:v0.5.12.post1
```
For more detailed information, please see this cookbook.

**8× B200 single-node deployment (BF16, chunked prefill + MTP on by default):**

```
docker run -d --name nemotron-ultra-sglang \
  --gpus all \
  --cap-add SYS_NICE \
  --ipc=host \
  --network=host \
  --shm-size=16g \
  --ulimit memlock=-1 \
  --ulimit stack=67108864 \
  -v $MODEL_CKPT:/model:ro \
  -e SAFETENSORS_FAST_GPU=1 \
  -e NVIDIA_TF32_OVERRIDE=1 \
  -e SGLANG_DISABLE_DEEP_GEMM=1 \
  lmsysorg/sglang:v0.5.12.post1 \
  python3 -m sglang.launch_server \
  --model-path /model \
  --host 0.0.0.0 \
  --port 8000 \
  --served-model-name nvidia/nemotron-3-ultra \
  --tp-size 8 \
  --ep-size 8 \
  --context-length 262144 \
  --mem-fraction-static 0.85 \
  --chunked-prefill-size 32768 \
  --fp8-gemm-backend triton \
  --moe-runner-backend triton \
  --mamba-scheduler-strategy no_buffer \
  --disable-piecewise-cuda-graph \
  --reasoning-parser nemotron_v3 \
  --tool-call-parser qwen3_coder \
  --speculative-algorithm EAGLE \
  --speculative-num-steps 5 \
  --speculative-eagle-topk 1 \
  --speculative-num-draft-tokens 5 \
  --trust-remote-code \
  --log-level info
```
Context length defaults to 256k above. To use up to 1M, set `SGLANG_ALLOW_OVERWRITE_LONGER_CONTEXT_LEN=1` and `--context-length 1048576`.

**Tool calls + reasoning parsing:** when calling the chat completions endpoint with `tools`, you **must** set `"chat_template_kwargs": {"enable_thinking": true, "force_nonempty_content": true}` in the request body to parse both reasoning and tool calls correctly.

**Important Note:** Current support for Nemotron 3 Ultra is limited to NVIDIA Blackwell architecture (including B200/B300 and GB200/GB300). While support for NVIDIA Hopper systems is planned, it is not currently available.

Container:

```
docker pull nvcr.io/nvidia/tensorrt-llm/release:1.3.0rc17
```
For more detailed information, please see this cookbook.

**8xB200:**

```
cat > ./extra-llm-api-config.yml << EOF
backend: pytorch
trust_remote_code: true
tensor_parallel_size: 8
pipeline_parallel_size: 1
context_parallel_size: 1
gpus_per_node: 8
moe_expert_parallel_size: 1
disable_overlap_scheduler: false
cuda_graph_config:
  enable_padding: true
  max_batch_size: 256
enable_chunked_prefill: true
enable_attention_dp: false
max_batch_size: 256
max_seq_len: null
max_num_tokens: 32768
num_postprocess_workers: 4
kv_cache_config:
  enable_block_reuse: false
  max_tokens: null
  max_attention_window: null
  sink_token_length: null
  free_gpu_memory_fraction: 0.75
  host_cache_size: null
  cross_kv_cache_fraction: null
  secondary_offload_min_priority: null
  event_buffer_max_size: 0
  attention_dp_events_gather_period_ms: 5
  enable_partial_reuse: true
  copy_on_partial_reuse: true
  use_uvm: false
  max_gpu_total_bytes: 0
  iteration_stats_interval: 1
  dtype: fp8
  tokens_per_block: 32
  mamba_state_cache_interval: 256
  use_kv_cache_manager_v2: false
  max_util_for_resume: 0.95
moe_config:
  backend: TRTLLM
  max_num_tokens: null
  load_balancer: null
  disable_finalize_fusion: false
  use_low_precision_moe_combine: false
EOF
TLLM_ALLOW_LONG_MAX_MODEL_LEN=1 trtllm-serve  \
<bf16_ckpt> \
--max_batch_size 256 \
--tp_size 8 --ep_size 1 \
--max_num_tokens 32768 \
--trust_remote_code \
--reasoning_parser nano-v3 \
--tool_parser qwen3_coder \
--chat_template $MODEL_DIR/chat_template.jinja \
--extra_llm_api_options extra-llm-api-config.yml
```
**Long-context configuration:**

For long-context benchmarking, set `TLLM_ALLOW_LONG_MAX_MODEL_LEN=1` as an environment variable and add `--max_seq_len <seq_len>` as the desired maximum context length. The MTP `speculative_config` block above carries over unchanged — on rc16, `max_draft_len` is the authoritative field and `num_nextn_predict_layers` is treated as deprecated.

The examples below use the OpenAI-compatible client and work with any of the serving backends above.

NOTE: For coding agents add the following to the API call - `extra_body={"chat_template_kwargs": {"force_nonempty_content": True}}`

```
from openai import OpenAI
client = OpenAI(base_url="http://localhost:8000/v1", api_key="EMPTY")
MODEL = "nvidia/nemotron-3-ultra"
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
**Medium-effort reasoning**

Uses significantly fewer reasoning tokens than full thinking mode. Recommended as a starting point before tuning explicit token budgets.

```
response = client.chat.completions.create(
    model=MODEL,
    messages=[{"role": "user", "content": "What is the capital of Japan?"}],
    max_tokens=16000,
    temperature=1.0,
    top_p=0.95,
    extra_body={"chat_template_kwargs": {"enable_thinking": True, "medium_effort": True}}
)
print(response.choices[0].message.content)
```
**Tool calling with reasoning (SGLang requires explicit chat template kwargs)**

```
response = client.chat.completions.create(
    model=MODEL,
    messages=[{"role": "user", "content": "What's the weather in New York?"}],
    tools=[{
        "type": "function",
        "function": {
            "name": "get_weather",
            "description": "Get the current weather for a city.",
            "parameters": {
                "type": "object",
                "properties": {
                    "city": {"type": "string"},
                    "unit": {"type": "string", "enum": ["celsius", "fahrenheit"]}
                },
                "required": ["city"]
            }
        }
    }],
    tool_choice="required",
    max_tokens=256,
    temperature=1.0,
    top_p=0.95,
    extra_body={"chat_template_kwargs": {"enable_thinking": True, "force_nonempty_content": True}}
)
```
OpenCode is an AI coding agent that runs in your terminal. It connects to any OpenAI-compatible endpoint, making it compatible with all three serving backends above (vLLM, SGLang, and TensorRT-LLM).

Create or update your `~/.config/opencode/opencode.json`:


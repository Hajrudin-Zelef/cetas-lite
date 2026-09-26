---
id: collect-240926-huggingface/huggingface/nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-nvfp4-hugging-face-3
title: "nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-nvfp4-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["Nvidia", "OpenAI", "SGLang", "vLLM"]
dates: []
keywords: ["nvfp4", "nvidia", "agents", "attention", "context window", "decode", "fp8", "gguf", "gpu", "gpus", "llama", "llama.cpp"]
source: docs/RAG/clean_en/huggingface/nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-nvfp4-hugging-face.md
source_anchor: ""
source_lines: [308, 518]
sha256: 68a6f3273ae5406c67126be4d82beee68e3475c5cf7ee62a62a25bf028ee7191
---

# nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-nvfp4-hugging-face

```
cat > nemotron-35-lightning-nvfp4-mtp.yaml << EOF
kv_cache_config:
  dtype: fp8
  enable_block_reuse: false
  mamba_state_config:
     periodic_snapshot_interval: 8192
  free_gpu_memory_fraction: 0.8
  mamba_ssm_cache_dtype: float16
  mamba_ssm_stochastic_rounding: true
  mamba_ssm_philox_rounds: 5
moe_config:
   backend: MARLIN
nvfp4_gemm_config:
  allowed_backends: [marlin, cutlass, cublaslt, cuda_core]
cuda_graph_config:
    enable_padding: true
    max_batch_size: 8
speculative_config:
  decoding_type: MTP
  max_draft_len: 3
  allow_advanced_sampling: true
enable_chunked_prefill: true
num_postprocess_workers: 4
print_iter_log: true
stream_interval: 10
disable_overlap_scheduler: false
EOF
trtllm-serve \
$MODEL_CKPT \
--max_batch_size 8 \
--max_num_tokens 8192 \
--reasoning_parser nemotron-v3 \
--tool_parser qwen3_coder \
--config nemotron-35-lightning-nvfp4-mtp.yaml
```
**Validated context:** 1M tokens by default; lower `--max_seq_len` if memory-constrained.

For more indepth instructions on how to deploy through SGLang, head here or the SGLang cookbook


- Container: `lmsysorg/sglang:dev-nemotron3-5-lightning`

**Balanced — no speculative decoding:**

```
sglang serve \
    --model-path $MODEL_CKPT \
    --mamba-backend flashinfer \
    --mamba-ssm-dtype float16 \
    --enable-mamba-cache-stochastic-rounding \
    --mamba-cache-philox-rounds 5 \
    --mem-fraction-static 0.85 \
    --cuda-graph-max-bs-decode 16 \
    --reasoning-parser nemotron_3 \
    --tool-call-parser qwen3_coder \
    --port 8000
```
On Hopper the NVFP4 weights run through W4A16 kernels; the `flashinfer` Mamba backend is not required (FA3 target attention is selected by default):

```
sglang serve \
    --model-path $MODEL_CKPT \
    --mamba-ssm-dtype float16 \
    --mem-fraction-static 0.85 \
    --cuda-graph-max-bs-decode 16 \
    --reasoning-parser nemotron_3 \
    --tool-call-parser qwen3_coder \
    --port 8000
```
Smaller memory budget — `--mem-fraction-static` and `--cuda-graph-max-bs-decode` are lowered accordingly:

```
sglang serve \
    --model-path $MODEL_CKPT \
    --mamba-ssm-dtype float16 \
    --mem-fraction-static 0.78 \
    --cuda-graph-max-bs-decode 4 \
    --reasoning-parser nemotron_3 \
    --tool-call-parser qwen3_coder \
    --port 8000
```
The blocks above run the balanced baseline. To enable one of the speculative decoding strategies described above, append the matching flags to any command:

**MTP** — the draft head is embedded in the target checkpoint (no separate download):

```
    --speculative-algorithm EAGLE \
    --speculative-draft-model-path $MODEL_CKPT \
    --speculative-num-steps 5 \
    --speculative-eagle-topk 1 \
    --speculative-num-draft-tokens 6
```
**DFlash** — separate draft model (`--speculative-dflash-block-size` is `6` on B200, `4` on H100 / DGX Spark):

```
    --speculative-algorithm DFLASH \
    --speculative-draft-model-path nvidia/NVIDIA-Nemotron-3.5-Lightning-30B-A3B-NVFP4-DFlash \
    --speculative-dflash-block-size 6
```
**DSpark** — separate draft model:

```
    --speculative-algorithm DSPARK \
    --speculative-draft-model-path $DSPARK_CKPT \
    --speculative-dspark-block-size 3
```
- **Context length:** The commands above serve the model's full context window by default. Set`--context-length` to a smaller value if you're memory-constrained or want more KV-cache headroom at higher concurrency.

The following recipes are validated and provided by our fantastic partners.

- ollama version: 0.32.9

For local, single-command use, Nemotron 3.5 Lightning is on Ollama, with tool-calling and thinking enabled.

```
ollama run nemotron-3.5-lightning
```
**Validated context:** Ollama will dynamically set context based on available VRAM (under 24GB -> 4K, 24-48GB -> 32K, 48GB+ -> 256K). You can use `/set parameter num_ctx DESIRED_CONTEXT_LENGTH` which may result in CPU offloading for higher context limits.

Official GGUF weights: `ggml-org/NVIDIA-Nemotron-3.5-Lightning-30B-A3B-GGUF`. 

- llama.cpp version: https://github.com/ggml-org/llama.cpp/releases/latest

Simplest start:

```
llama-server -hf ggml-org/NVIDIA-Nemotron-3.5-Lightning-30B-A3B-GGUF \
  --port 8000
```
Full server recipe with recommended sampling:

```
llama-server \
  -hf ggml-org/NVIDIA-Nemotron-3.5-Lightning-30B-A3B-GGUF:Q4_K_M \
  --temp 1.0 --top-p 0.95 \
  -np 1 \
  -c 40960 \
  --port 8000 \
  -ngl 99 \
  -fa on \
  --jinja \
  --no-webui \
  --fit off
```
**Validated context:** the examples set `-c 40960` (~40K); raise it as VRAM allows.

You can also run Nemotron 3.5 Lightning 30B A3B through LM Studio from this resource!

The examples below use the OpenAI-compatible client and work with any of the serving backends above. All backends serve on port `8000` (vLLM and TRT-LLM by default; SGLang via `--port 8000`), so the `base_url` works as-is. Recommended sampling settings are **Temperature 1.0** and **Top_P 0.95**.

The vLLM snippets above pass the checkpoint to `--model $MODEL_CKPT`, and vLLM registers the model under that same identifier (`nvidia/NVIDIA-Nemotron-3.5-Lightning-30B-A3B-NVFP4`) — none of the commands set `--served-model-name`. For the other backends — or if you add a `--served-model-name` of your own — copy the identifier returned by `GET /v1/models` into MODEL below.

```
from openai import OpenAI
client = OpenAI(base_url="http://localhost:8000/v1", api_key="EMPTY")
MODEL = "nvidia/NVIDIA-Nemotron-3.5-Lightning-30B-A3B-NVFP4"
```
Lightning 3.5 exposes reasoning control through chat-template kwargs: thinking enabled (the default), and thinking disabled for direct answers.

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
The TRT-LLM snippet above already launches with the required parsers (`--reasoning_parser nemotron-v3 --tool_parser qwen3_coder`). For vLLM, add the following to any serve command above: 

```
    --enable-auto-tool-choice \
    --tool-call-parser qwen3_coder \
    --reasoning-parser nemotron_v3
```
**NOTE:** For coding agents, add `extra_body={"chat_template_kwargs": {"force_nonempty_content": True}}` to the API call, as shown below.


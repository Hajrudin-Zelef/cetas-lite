---
id: collect-240926-huggingface/huggingface/nvidia-nvidia-nemotron-3-super-120b-a12b-nvfp4-hugging-face-2
title: "with uv: uv pip install vllm==0.20.0 --torch-backend=auto"
domain: huggingface
role: reference
task: reference
actors: ["China", "Nvidia", "OpenAI"]
dates: []
keywords: ["vllm", "agents", "attention", "blackwell", "fp4", "fp8", "gpu", "gpus", "inference", "memory", "moe", "nvfp4"]
source: docs/RAG/clean_en/huggingface/nvidia-nvidia-nemotron-3-super-120b-a12b-nvfp4-hugging-face.md
source_anchor: ""
source_lines: [95, 294]
sha256: 316401d0a6f48619afc8e97b6de48070f9f25489fdf5c2c52dae1bb6e556cc46
---

# with uv: uv pip install vllm==0.20.0 --torch-backend=auto

- The model underwent multi-environment reinforcement learning using asynchronous GRPO (Group Relative Policy Optimization) across math, code, science, instruction following, multi-step tool use, multi-turn conversations, and structured output environments. It utilized an asynchronous RL architecture that fully decouples training from inference across separate GPU devices, leveraging in-flight weight updates and MTP to accelerate rollout generation. Conversational quality was further refined through RLHF. All datasets are disclosed in the *Training and Evaluation Datasets* section of this document. The RL environments and datasets are released as part of NeMo Gym.
- Software used for reinforcement learning: NeMo RL, NeMo Gym

NVIDIA-Nemotron-3-Super-120B-A12B-NVFP4 model is a result of the above work.

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
- Supported Hardware Microarchitecture Compatibility: NVIDIA Blackwell
- Operating System(s): Linux

- v1.0 - GA

For each inference backend, you'll need the custom `super_v3` reasoning parser. Download it with:

```
wget https://huggingface.co/nvidia/NVIDIA-Nemotron-3-Super-120B-A12B-NVFP4/raw/main/super_v3_reasoning_parser.py
```
OR

```
curl -O https://huggingface.co/nvidia/NVIDIA-Nemotron-3-Super-120B-A12B-NVFP4/raw/main/super_v3_reasoning_parser.py
```
For advanced deployment configurations, visit this resource.

NOTE: For running on Spark - please use the following instructions


For more detailed information, please see this cookbook.

```
pip install vllm==0.20.0
# with uv: uv pip install vllm==0.20.0 --torch-backend=auto
export MODEL_CKPT=nvidia/NVIDIA-Nemotron-3-Super-120B-A12B-NVFP4
```
```
vllm serve $MODEL_CKPT \
  --served-model-name nvidia/nemotron-3-super \
  --async-scheduling \
  --dtype auto \
  --max-model-len 262144 \
  --swap-space 0 \
  --trust-remote-code \
  --kv-cache-dtype fp8 \
  --gpu-memory-utilization 0.9 \
  --max-cudagraph-capture-size 128 \
  --enable-chunked-prefill \
  --mamba-ssm-cache-dtype float16 \
  --reasoning-parser-plugin /app/super_v3_reasoning_parser.py \
  --reasoning-parser super_v3 \
  --enable-auto-tool-choice \
  --tool-call-parser qwen3_coder
```
Context length defaults to 256k above. To use up to 1M, set `VLLM_ALLOW_LONG_MAX_MODEL_LEN=1` and `--max-model-len 1048576`.


To deploy the NVFP4 chekpoint on NVIDIA DGX Spark, make sure that you are using the `vllm/vllm-openai:v0.27.1` container image and use the following command:

```
docker run --rm -it --gpus all \
  -e VLLM_NVFP4_GEMM_BACKEND=marlin \
  -e VLLM_ALLOW_LONG_MAX_MODEL_LEN=1 \
  -e VLLM_FLASHINFER_ALLREDUCE_BACKEND=trtllm \
  -e VLLM_USE_FLASHINFER_MOE_FP4=0 \
  -e HF_TOKEN=$HF_TOKEN \
  -v ~/.cache/huggingface:/root/.cache/huggingface \
  -v $(pwd)/super_v3_reasoning_parser.py:/app/super_v3_reasoning_parser.py \
  -p 8000:8000 \
  vllm/vllm-openai:v0.27.1 \
    --model nvidia/NVIDIA-Nemotron-3-Super-120B-A12B-NVFP4 \
    --served-model-name nvidia/nemotron-3-super \
    --host 0.0.0.0 \
    --port 8000 \
    --async-scheduling \
    --dtype auto \
    --kv-cache-dtype fp8 \
    --tensor-parallel-size 1 \
    --pipeline-parallel-size 1 \
    --data-parallel-size 1 \
    --trust-remote-code \
    --gpu-memory-utilization 0.90 \
    --enable-chunked-prefill \
    --max-num-seqs 4 \
    --max-model-len 1000000 \
    --moe-backend marlin \
    --mamba_ssm_cache_dtype float16 \
    --quantization fp4 \
    --speculative_config '{"method":"mtp","num_speculative_tokens":3,"model":"nvidia/Nemotron-3-Super-120B-A12B-BF16-MTPv2","moe_backend":"triton"}' \
    --reasoning-parser-plugin /app/super_v3_reasoning_parser.py \
    --reasoning-parser super_v3 \
    --enable-auto-tool-choice \
    --tool-call-parser qwen3_coder
```
Container:

```
docker pull lmsysorg/sglang:dev-cu13-nemotronh-nano-omni-reasoning-v3
```
For more detailed information, please see this cookbook.

```
docker run --gpus all -it --rm \
  -p 30000:30000 \
  -v ~/.cache/huggingface:/root/.cache/huggingface \
  -e HF_TOKEN=$HF_TOKEN \
  --shm-size 16g \
  lmsysorg/sglang:dev-cu13-nemotronh-nano-omni-reasoning-v3 \
  python3 -m sglang.launch_server \
    --model-path nvidia/NVIDIA-Nemotron-3-Super-120B-A12B-NVFP4 \
    --served-model-name nvidia/nemotron-3-super \
    --host 0.0.0.0 \
    --port 30000 \
    --trust-remote-code \
    --quantization modelopt_fp4 \
    --mem-fraction-static 0.8 \
    --max-running-requests 8 \
    --tool-call-parser qwen3_coder \
    --reasoning-parser nemotron_3 \
    --disable-piecewise-cuda-graph
```
Context length defaults to 256k above. To use up to 1M, set `SGLANG_ALLOW_OVERWRITE_LONGER_CONTEXT_LEN=1` and `--context-length 1048576`.


Container:

```
docker pull nvcr.io/nvidia/tensorrt-llm/release:1.3.0rc12
```
For more detailed information, please see this cookbook.

```
cat > extra-llm-api-config.yml << 'EOF'
kv_cache_config:
  dtype: fp8
  enable_block_reuse: false
  free_gpu_memory_fraction: 0.9
  mamba_ssm_cache_dtype: float16
  mamba_ssm_stochastic_rounding: true
  mamba_ssm_philox_rounds: 5
moe_config:
   backend: CUTLASS
cuda_graph_config:
    enable_padding: true
    max_batch_size: 8
enable_attention_dp: false
enable_chunked_prefill: true
stream_interval: 1
print_iter_log: true
speculative_config:
  decoding_type: MTP
  num_nextn_predict_layers: 3
  allow_advanced_sampling: true
EOF
docker run --gpus all -it --rm \
  -p 8123:8123 \
  -v ~/.cache/huggingface:/root/.cache/huggingface \
  -v "$(pwd)/extra-llm-api-config.yml:/workspace/extra-llm-api-config.yml:ro" \
  -e HF_TOKEN=$HF_TOKEN \
  -e TLLM_ALLOW_LONG_MAX_MODEL_LEN=1 \
  --shm-size 16g \
  --ulimit memlock=-1 --ulimit stack=67108864 \
  -w /workspace \
  nvcr.io/nvidia/tensorrt-llm/release:1.3.0rc12 \
  trtllm-serve nvidia/NVIDIA-Nemotron-3-Super-120B-A12B-NVFP4 \
    --host 0.0.0.0 \
    --port 8123 \
    --served_model_name nvidia/nemotron-3-super \
    --max_batch_size 8 \
    --tp_size 1 --ep_size 1 \
    --max_num_tokens 8192 \
    --trust_remote_code \
    --reasoning_parser nano-v3 \
    --tool_parser qwen3_coder \
    --extra_llm_api_options /workspace/extra-llm-api-config.yml \
    --max_seq_len 1048576
```
The examples below use the OpenAI-compatible client and work with any of the serving backends above.

NOTE: For coding agents add the following to the API call - `extra_body={“chat_template_kwargs”: {“force_nonempty_content”: True}`


```
from openai import OpenAI
client = OpenAI(base_url="http://localhost:8000/v1", api_key="EMPTY")
MODEL = "nvidia/nemotron-3-super"
```
**Reasoning ON (default)**


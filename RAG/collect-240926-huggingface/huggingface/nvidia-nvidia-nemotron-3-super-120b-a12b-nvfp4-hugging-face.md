---
id: collect-240926-huggingface/huggingface/nvidia-nvidia-nemotron-3-super-120b-a12b-nvfp4-hugging-face
title: "with uv: uv pip install vllm==0.20.0 --torch-backend=auto"
domain: huggingface
role: reference
task: reference
actors: ["AWS", "Alibaba", "China", "DeepSeek", "Google", "Hugging Face", "Meta", "MiniMax", "Mistral", "Moonshot", "Nvidia", "OpenAI", "SGLang", "Z.ai", "vLLM"]
dates: ["2025-01-05", "2025-02-10", "2025-06", "2025-08-04", "2025-12", "2025-29-04", "2026-02", "2026-02-24", "2026-03", "2026-03-11", "2026-11-03"]
keywords: ["vllm", "agent", "agentic", "agents", "alignment", "attention", "benchmark", "benchmarks", "blackwell", "compute", "deepseek", "distribution"]
source: docs/RAG/clean_en/huggingface/nvidia-nvidia-nemotron-3-super-120b-a12b-nvfp4-hugging-face.md
source_anchor: ""
source_lines: [1, 665]
sha256: 6d3deb7699cce63b086aa209545987d136b9bef4a48bfb19ae8051d6478c74ad
---

# with uv: uv pip install vllm==0.20.0 --torch-backend=auto

<!-- source: https://huggingface.co/nvidia/NVIDIA-Nemotron-3-Super-120B-A12B-NVFP4 -->

| **Total Parameters** | 120B (12B active) | 
| **Architecture** | LatentMoE - Mamba-2 + MoE + Attention hybrid with Multi-Token Prediction (MTP) | 
| **Context Length** | Up to 1M tokens | 
| **Minimum GPU Requirement** | 1× B200 OR 1× DGX Spark | 
| **Supported Languages** | English, French, German, Italian, Japanese, Spanish, Chinese | 
| **Best For** | Agentic workflows, long-context reasoning, high-volume workloads (e.g. IT ticket automation), tool use, RAG | 
| **Reasoning Mode** | Configurable on/off via chat template ( `enable_thinking=True/False` ) | 
| **Speculative Decoding** | Includes a built-in MTP head, with an updated MTPv2 head available as a separate checkpoint | 
| **License** | NVIDIA Nemotron Open Model License | 
| **Release Date** | March 11, 2026 | 

Use `temperature=1.0` and `top_p=0.95` across **all tasks and serving backends** — reasoning, tool calling, and general chat alike.


For more details on how to deploy and use the model - see the Quick Start Guide below!

**Model Developer:** NVIDIA Corporation

**Model Dates:** December 2025 - March 2026

**Data Freshness:**

- The post-training data has a cutoff date of February 2026.
- The pre-training data has a cutoff date of June 2025.

**Nemotron-3-Super-120B-A12B-NVFP4** is a large language model (LLM) trained by NVIDIA, designed to deliver strong agentic, reasoning, and conversational capabilities. It is optimized for collaborative agents and high-volume workloads such as IT ticket automation. Like other models in the family, it responds to user queries and tasks by first generating a reasoning trace and then concluding with a final response. The model's reasoning capabilities can be configured through a flag in the chat template.

The model employs a hybrid **Latent Mixture-of-Experts (LatentMoE)** architecture, utilizing interleaved Mamba-2 and MoE layers, along with select Attention layers. Distinct from the Nano model, the Super model incorporates **Multi-Token Prediction (MTP)** layers for faster text generation and improved quality, and it is trained using **NVFP4** quantization to maximize compute efficiency. The model has **12B active parameters** and **120B parameters in total**.

The supported languages include: English, French, German, Italian, Japanese, Spanish, and Chinese

This model is ready for commercial use.

**Governing Download Terms:** Use of this model is governed by the NVIDIA Nemotron Open Model License.

**Governing Download Terms with NIM:** The NIM container is governed by the NVIDIA Software License Agreement and Product-Specific Terms for AI Products. Use of this model is governed by the NVIDIA Nemotron Open Model License.

| Benchmark | Nemotron-3-Super | Nemotron-3-Super FP8 | Nemotron-3-Super NVFP4 | 
|---|---|---|---|
| **General Knowledge** |  |  |  | 
| MMLU-Pro | 83.73 | 83.63 | 83.33 | 
| **Reasoning** |  |  |  | 
| HMMT Feb25 (with tools) | 94.73 | 94.38 | 95.36 | 
| GPQA (no tools) | 79.23 | 79.36 | 79.42 | 
| LiveCodeBench (v6 2024-08↔2025-05) | 78.69 | 78.44 | 78.44 | 
| LiveCodeBench (v5 2024-07↔2024-12) | 81.19 | 80.99 | 80.56 | 
| SciCode (subtask) | 42.05 | 41.38 | 40.83 | 
| HLE (no tools) | 18.26 | 17.42 | 17.42 | 
| **Agentic** |  |  |  | 
| Terminal Bench (hard subset) | 25.78 | 26.04 | 24.48 | 
| **TauBench V2** |  |  |  | 
| Airline | 56.25 | 56.25 | 54.75 | 
| Retail | 62.83 | 63.05 | 63.38 | 
| Telecom | 64.36 | 63.93 | 63.27 | 
| Average | 61.15 | 61.07 | 60.46 | 
| **Chat & Instruction Following** |  |  |  | 
| IFBench (prompt) | 72.58 | 72.32 | 73.30 | 
| Scale AI Multi-Challenge | 55.23 | 54.35 | 52.8 | 
| Arena-Hard-V2 (Hard Prompt) | 73.88 | 76.06 | 76.00 | 
| **Long Context** |  |  |  | 
| AA-LCR | 58.31 | 57.69 | 58.06 | 
| RULER-500 @ 128k (500 samples per task) | 96.79 | 96.85 | 95.99 | 
| RULER-500 @ 256k (500 samples per task) | 96.60 | 96.33 | 96.52 | 
| RULER-500 @ 512k (500 samples per task) | 96.09 | 95.66 | 96.23 | 
| **Multilingual** |  |  |  | 
| MMLU-ProX (avg over languages) | 79.35 | 79.21 | 79.37 | 

All evaluation results were collected via Nemo Evaluator SDK and for most benchmarks, the Nemo Skills Harness. For reproducibility purposes, more details on the evaluation settings can be found in the Nemo Evaluator SDK configs folder and the reproducibility tutorial for Nemotron 3 Super. The open source container on Nemo Skills packaged via NVIDIA's Nemo Evaluator SDK used for evaluations can be found here. In addition to Nemo Skills, the evaluations also used dedicated open-source packaged containers for Tau-2 Bench (default prompt), Terminal Bench Hard (48 tasks), ScaleAI Multi Challenge Multi-turn Instruction Following, and Ruler.

The following benchmarks are not onboarded yet in our open source tools and for these we used either their official open source implementation or otherwise an internal scaffolding that we plan to open source in the future: SWE Bench Verified (OpenHands), SWE Bench Multilingual (OpenHands), BrowseComp with Search (internal implementation with Serp API), Terminal Bench Core 2.0 (Harbor).

NVIDIA-Nemotron-3-Super-120B-A12B-NVFP4 is a general purpose reasoning and chat model intended to be used in English, Code, and supported multilingual contexts. This model is optimized for collaborative agents and high-volume workloads. It is intended to be used by developers designing AI Agent systems, chatbots, RAG systems, and other AI-powered applications. This model is also suitable for complex instruction-following tasks and long-context reasoning.

Hugging Face - 03/11/2026 via Hugging Face

- **Architecture Type:** Mamba2-Transformer Hybrid Latent Mixture of Experts (LatentMoE) with Multi-Token Prediction (MTP)
- **Network Architecture:** Nemotron Hybrid LatentMoE
- **Number of model parameters:** 120B Total / 12B Active

The model utilizes the **LatentMoE** architecture, where tokens are projected into a smaller latent dimension for expert routing and computation, improving accuracy per byte. The Super model is pre-trained using NVFP4 quantization — the first model in the Nemotron 3 family trained at this precision. The majority of linear layers use NVFP4 for weights, activations, and gradients, while select layers (including latent projections, MTP layers, QKV/attention projections, and embeddings) are maintained in BF16 or MXFP8 for training stability. The model includes **Multi-Token Prediction (MTP)** layers using a shared-weight design across prediction heads. This improves training signal quality, enables faster inference via native speculative decoding, and supports more stable autoregressive drafting at longer draft lengths compared to independently trained offset heads.

Stage 1: Pre-Training

- NVIDIA-Nemotron-3-Super-120B-A12B-Base-BF16 model was pre-trained for over 25T tokens using crawled and synthetic code, math, science, and general knowledge data. Training leveraged NVFP4 quantization for efficiency. All datasets are disclosed in the Training and Evaluation Datasets section of this document. Major portions of the pre-training corpus are released in the Nemotron-Pre-Training-Datasets collection.
- Software used for pre-training: Megatron-LM

Stage 2: Supervised Fine-Tuning

- The model was further fine-tuned on synthetic code, math, science, tool calling, instruction following, structured outputs, and general knowledge data. This stage incorporated data designed to support long-range retrieval and multi-document aggregation. All datasets are disclosed in the Training and Evaluation Datasets section of this document. Major portions of the fine-tuning corpus are released in the Nemotron-Post-Training-v3 collection. Data Designer is one of the libraries used to prepare these corpora.

Stage 3: Reinforcement Learning

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

```
{
    "$schema": "https://opencode.ai/config.json",
    "model": "local/nvidia-nemotron-3-super",
    "provider": {
        "local": {
            "npm": "@ai-sdk/openai-compatible",
            "name": "local_backend",
            "options": {
                "baseURL": "http://localhost:8000/v1",
                "apiKey": "EMPTY"
            },
            "models": {
                "nvidia-nemotron-3-super": {
                    "name": "nvidia/nemotron-3-super",
                    "limit": {
                        "context": 1000000,
                        "output": 32768
                    }
                }
            }
        }
    },
    "agent": {
        "build": {
            "temperature": 1.0,
            "top_p": 0.95,
            "max_tokens": 32000
        },
        "plan": {
            "temperature": 1.0,
            "top_p": 0.95,
            "max_tokens": 32000
        }
    }
}
```
Update `baseURL` to match whichever backend you are running. The default port above (`8000`) matches the vLLM example; SGLang and TRT-LLM use `30000` and `8123` respectively.


To learn more about other supported agent scaffolds - check out this resource

##  **Advanced: Budget-Controlled Reasoning**

Set a hard token ceiling on the reasoning trace using `reasoning_budget`. The model will attempt to close the trace at the next newline before the budget is hit; if none is found within 500 tokens it closes abruptly at `reasoning_budget + 500`.

```
from typing import Any, Dict, List
import openai
from transformers import AutoTokenizer
class ThinkingBudgetClient:
    def __init__(self, base_url: str, api_key: str, tokenizer_name_or_path: str):
        self.tokenizer = AutoTokenizer.from_pretrained(tokenizer_name_or_path)
        self.client = openai.OpenAI(base_url=base_url, api_key=api_key)
    def chat_completion(        self,
        model: str,
        messages: List[Dict[str, Any]],
        reasoning_budget: int = 512,
        max_tokens: int = 1024,
        **kwargs,
    ) -> Dict[str, Any]:
        assert max_tokens > reasoning_budget, (
            f"reasoning_budget must be less than max_tokens. "
            f"Got {max_tokens=} and {reasoning_budget=}"
        )
        # Step 1: generate the reasoning trace up to the budget
        response = self.client.chat.completions.create(
            model=model, messages=messages, max_tokens=reasoning_budget, **kwargs
        )
        reasoning_content = response.choices[0].message.content
        if "" not in reasoning_content:
            reasoning_content = f"{reasoning_content}.\n\n\n"
        reasoning_tokens_len = len(
            self.tokenizer.encode(reasoning_content, add_special_tokens=False)
        )
        remaining_tokens = max_tokens - reasoning_tokens_len
        assert remaining_tokens > 0, (
            f"No tokens remaining for response ({remaining_tokens=}). "
            "Increase max_tokens or lower reasoning_budget."
        )
        # Step 2: continue from the closed reasoning trace
        messages.append({"role": "assistant", "content": reasoning_content})
        prompt = self.tokenizer.apply_chat_template(
            messages, tokenize=False, continue_final_message=True
        )
        response = self.client.completions.create(
            model=model, prompt=prompt, max_tokens=remaining_tokens, **kwargs
        )
        return {
            "reasoning_content": reasoning_content.strip().strip("").strip(),
            "content": response.choices[0].text,
            "finish_reason": response.choices[0].finish_reason,
        }
```
**Example usage** (32-token reasoning budget):

```
client = ThinkingBudgetClient(
    base_url="http://localhost:8000/v1",
    api_key="EMPTY",
    tokenizer_name_or_path="nvidia/NVIDIA-Nemotron-3-Super-120B-A12B-NVFP4",
)
result = client.chat_completion(
    model="nvidia/NVIDIA-Nemotron-3-Super-120B-A12B-NVFP4",
    messages=[
        {"role": "system", "content": "You are a helpful assistant. /think"},
        {"role": "user", "content": "What is 2+2?"},
    ],
    reasoning_budget=32,
    max_tokens=512,
    temperature=1.0,
    top_p=0.95,
)
print(result)
```
**Data Modality:** Text
**The total size:** 15,573,172,908,990 Tokens
**Total number of datasets:** 153
**Dataset partition:** *Training [100%], testing [0%], validation [0%]*
**Time period for training data collection:** 2013 to February 24, 2026
**Time period for testing data collection:** 2013 to February 24, 2026
**Time period for validation data collection:** 2013 to February 24, 2026
**Data Collection Method by dataset:** Hybrid: Automated, Human, Synthetic
**Labeling Method by dataset:** Hybrid: Automated, Human, Synthetic

NVIDIA-Nemotron-3-Super-120B-A12B-NVFP4 is pre-trained on a large corpus of high-quality curated and synthetically-generated data. It is trained in the English language, as well as 19 other languages and 43 programming languages. Our sources cover a variety of document types such as: webpages, dialogue, articles, and other written materials. The corpus spans domains including legal, math, science, finance, and more. We also include a small portion of question-answering, and alignment style data to improve model accuracy. The model was trained for approximately 25 trillion tokens.

The post-training corpus for NVIDIA-Nemotron-3-Super-120B-A12B-NVFP4 of high-quality curated and synthetically-generated data. Primary languages used for post-training include English, French, German, Italian, Japanese, Spanish, and Chinese.

During post-training, we generate synthetic data by distilling trajectories, solutions, and translations from strong teacher models and agent systems, often grounded in real tasks or documents and aggressively filtered for quality. For math, code, and science, we start from curated problem sets and use open source permissive models such as GPT-OSS-120B to produce step-by-step reasoning traces, candidate solutions, best-of-n selection traces, and verified CUDA kernels. For long-context and science, we build synthetic QA and reasoning data by retrieving passages from long documents, generating MCQ/OpenQA questions and answers, and paraphrasing them into multiple prompt/response formats to ensure diversity. Across all pipelines we stack automated verification—compilers, numerical checks, language identification—to ensure our data is high quality.

More details on the datasets and synthetic data generation methods can be found in the technical report ***NVIDIA Nemotron 3 Super***.

## **Click to explore the full dataset catalogue used for training**

The foundation of the model is trained on the **Nemotron-3-Nano** corpus, comprising the following collections:

| Dataset Collection | Token Counts | Description | 
|---|---|---|
| **Nemotron-CC-v2** &**v2.1** | 9.13T | A massive collection of English web data filtered from Common Crawl, including 2.5T+ tokens of new organic, translated, and synthetically rephrased content. | 
| **Nemotron-CC-Code-v1** | 427.9B | High-quality code tokens extracted from Common Crawl using the Lynx + LLM pipeline to preserve structure and equations. | 
| **Nemotron-Pretraining-Code-v1** &**v2** | 1.09T | Curated GitHub code references with multi-stage filtering, deduplication, and large-scale synthetic code data. | 
| **Nemotron-CC-Math-v1** | 133.3B | High-quality math pre-training dataset preserving LaTeX formatting and mathematical structures. | 
| **Nemotron-Pretraining-Specialized-v1** | 336.4B | Synthetic datasets targeting specialized domains such as STEM reasoning and scientific coding. | 

The GitHub Crawl was collected using the GitHub REST API and the Amazon S3 API. Each crawl was operated in accordance with the rate limits set by its respective source, either GitHub or S3. We collect raw source code and subsequently remove any having a license which does not exist in our permissive-license set (for additional details, refer to the technical report).

| Dataset | Modality | Dataset Size | Collection Period | Collecting Organisation | 
|---|---|---|---|---|
| English Common Crawl | Text | 3.36T | 4/8/2025 | NVIDIA Advanced Deep Learning Research | 
| English Common Crawl 1.1 | Text | Not disclosed | 10/2/2025 | NVIDIA Advanced Deep Learning Research | 
| Multilingual Common Crawl | Text | 812.7B | 5/1/2025 | NVIDIA Advanced Deep Learning Research | 
| GitHub Crawl | Text | 747.4B | 4/29/2025 | NVIDIA Advanced Deep Learning Research | 

| Dataset | Model(s) used | 
|---|---|
| Global Regulation | Unknown | 
| TAUS Translation Memory | Unknown | 
| Scale HLE | Unknown | 
| HackerRank Coding | Unknown | 
| RL data for Search | Gemini 3; GPT-5 * | 

- Models used for prompt generation only

| Dataset | Model(s) used | 
|---|---|
| Simple Minesweeper | - | 
| Simple Sudoku | - | 
| Multitool Typewriter Hard | - | 
| Machine Translation of News Commentary and TAUS Translation Memory | - | 
| Machine Translation of STEM - | Qwen2.5-14B-Instruct | 
| Competitive Coding RL data from Nemotron Cascade | - | 
| Long context RL | - | 
| Single-step SWE RL for patch generation | - | 
| OpenHands SWE | - | 

| Dataset | Modality | Dataset Size | Seed Dataset | Model(s) used for generation | 
|---|---|---|---|---|
| Nemotron-Pretraining-Formal-Logic | Text | 128,022,285 | Nemotron Personas | Qwen3-235B-A22B-Thinking-2507 | 
| Nemotron-Pretraining-Economics | Text | 73,374,154 | - | Qwen3-235B-A22B-Thinking-2507 | 
| Nemotron-Pretraining-Multiple-Choice | Text | 1,609,214,470 | MMLU Auxiliary Train | DeepSeek-V3; Qwen3-235B-A22B | 
| Nemotron-Pretraining-Code-Concepts | Text | 7,294,510,156 | - | gpt-oss-20b; gpt-oss-120b | 
| Nemotron-Pretraining-Unconditional-Algorithmic | Text | 196,492,899 | - | gpt-oss-120b; Qwen3-235B-A22B | 
| Synthetic Tasks from DeepSeek-V3 and Qwen3-235B-A22B | Text | 6.7B | train splits of Into the Unknown; AI2 ARC (AI2 Reasoning Challenge); BLiMP (Benchmark of Linguistic Minimal Pairs); CommonSenseQA; GLUE; HeadQA; Hendrycks Ethics; Memo Trap; modus-tollens; NeQA; pattern-matching-suppression; mastermind_24_mcq_random; mastermind_24_mcq_close; quote-repetition; redefine-math; Repetitive Algebra; sig-figs; MMLU-Pro; MC-TACO; MedConceptsQA; MMLU_dataset; OpenbooksQA; PIQA (Physical Interaction Question Answering); SocialIQA; SuperGLUE; tinyAI2_arc; tinyMMLU; tinyWinogrande; TruthfulQA; WebQuestions; Winogrande; GPQA; MBPP | DeepSeek v3; Qwen3-235B-A22B | 
| Synthetic Art of Problem Solving from DeepSeek-R1 | Text | 40B | Art of Problem Solving; American Mathematics Competitions 8; American Mathematics Competitions 10; | DeepSeek-R1 | 
| Synthetic Moral Stories and Social Chemistry from Mixtral-8x22B-v0.1 | Text | 327M | social-chemestry-101; Moral Stories | Mixtral-8x22B-v0.1 | 
| Synthetic Social Sciences seeded with OpenStax from DeepSeek-V3, Mixtral-8x22B-v0.1, and Qwen2.5-72B | Text | 83.6M | OpenStax - CC BY-SA subset | DeepSeek-V3; Mixtral-8x22B-v0.1; Qwen2.5-72B | 
| Synthetic Health Sciences seeded with OpenStax from DeepSeek-V3, Mixtral-8x22B-v0.1, and Qwen2.5-72B | Text | 9.7M | OpenStax - CC BY-SA subset | DeepSeek-V3; Mixtral-8x22B-v0.1; Qwen2.5-72B | 
|  | Text | 175M | OpenStax - CC BY-SA subset; GSM8K; Open Textbook Library - CC BY-SA & GNU subset | DeepSeek-R1, DeepSeek-V3; DeepSeek-V3-0324; Qwen2.5-72B | 
| Nemotron-PrismMath | Text | 4.6B | Big-Math-RL-Verified; OpenR1-Math-220k | Qwen2.5-0.5B-instruct, Qwen2.5-72B-Instruct; DeepSeek-R1-Distill-Qwen-32B | 
| Synthetic Question Answering Data from Papers and Permissible Books from Qwen2.5-72B-Instruct | Text | 350M |  | Qwen2.5-72B-Instruct | 
| Refreshed Nemotron-MIND from phi-4 | Text | 73B | Common Crawl | phi-4 | 
| Nemotron-CC-Math-4plus | Text | 52.3B | Common Crawl | phi-4 | 
| Nemotron-CC-Math-3 | Text | 80.9B | Common Crawl | phi-4 | 
| Synthetic AGIEval seeded with AQUA-RAT, LogiQA, and AR-LSAT from DeepSeek-V3 and DeepSeek-V3-0324 | Text | 4.0B | AQUA-RAT; LogiQA; AR-LSAT | DeepSeek-V3; DeepSeek-V3-0324 | 
| Synthetic AGIEval seeded with AQUA-RAT, LogiQA, and AR-LSAT from Qwen3-30B-A3B | Text | 4.2B | AQUA-RAT; LogiQA; AR-LSAT | Qwen3-30B-A3B | 
|  | Text |  |  | Qwen2.5-32B-Instruct; Qwen2.5-Math-72B; Qwen2.5-Math-7B; Qwen2.5-72B-Instruct | 
| Synthetic MMLU Auxiliary Train from DeepSeek-R1 | Text | 0.5B | MMLU Auxiliary Train | DeepSeek-R1 | 
|  | Text |  |  | Qwen2.5-72B-Instruct | 
| Synthetic Common Crawl from Qwen3-30B-A3B and Mistral-Nemo-12B-Instruct | Text | 415.8B | Common Crawl | Qwen3-30B-A3B; Mistral-NeMo-12B-Instruct | 
| Synthetic Multilingual Data from Common Crawl from Qwen3-30B-A3B | Text |  | Common Crawl | Qwen3-30B-A3B | 
| Synthetic Multilingual Data from Wikimedia from Qwen3-30B-A3B | Text |  | Wikimedia | Qwen3-30B-A3B | 
| Synthetic Math Data from Wikimedia from Nemotron-4-340B-Instruct | Text |  | - | Nemotron-4-340B-Instruct | 
| Synthetic Common Crawl Code from phi-4 | Text | 427.9B | Common Crawl | phi-4 | 
| Synthetic Scientific Coding from Qwen3-235B-A22B | Text | 1.2B | Wikimedia | Qwen3-235B-A22B | 
| Tool Calling Data | Text | 26.2B |  | Qwen3-235B-A22B-2507; gpt-oss-120b | 
| Synthetic Essential-Web from QwQ-32B | Text | 28.1B | Essential-Web | QwQ-32B | 
| Translated Synthetic Crawl | Text | 389.9B | Common Crawl | Qwen3-30B-A3B | 
| Translated Synthetic Wikipedia | Text | 7.9B | Wikimedia | Qwen3-30B-A3B | 
| Synthetic Art of Problem Solving from gpt-oss-120b and Qwen2.5-32B-Instruct | Text | Undisclosed | Art of Problem Solving; American Mathematics Competitions 8; American Mathematics Competitions 10 | gpt-oss-120b; Qwen2.5-32B-Instruct | 
| Synthetic Stack Exchange from gpt-oss-120b and Qwen2.5-32B-Instruct | Text | Undisclosed | Stack Exchange | gpt-oss-120b; Qwen2.5-32B-Instruct | 
| Synthetic OpenCodeReasoning from DeepSeek-R1-0528 | Text | Undisclosed | OpenCodeReasoning | DeepSeek-R1-0528 | 
| Synthetic HackerRank Coding from DeepSeek-R1-0528 | Text | Undisclosed | HackerRank Coding Dataset | DeepSeek-R1-0528 | 
| Synthetic SWE-Gym from Qwen3-Coder-480B-A35B-Instruct | Text | Undisclosed | SWE-Gym | Qwen3-Coder-480B-A35B-Instruct | 
|  | Text | Undisclosed | Art of Problem Solving; American Mathematics Competitions 8; American Mathematics Competitions 10; Stack Exchange | gpt-oss-120b; Qwen2.5-32B-Instruct; Goedel-Prover-V2-32B | 
|  | Text | Undisclosed | Stack Exchange; SCP-116K; LIMO; TACO; Code Contest; Codeforces | DeepSeek-R1; DeepSeek-R1-0528; Qwen2.5-32B-Instruct; Qwen3-235B-A22B; | 
| Synthetic Safety from DeepSeek-R1-0528, gpt-oss-120b and Mixtral-8x7B-v0.1 | Text | Undisclosed | Nemotron Content Safety Dataset V2; Gretel Synthetic Safety Alignment Dataset; RedTeam-2K; Malicious Tasks; Nemotron-Personas-USA | DeepSeek-R1-0528; gpt-oss-120b; Mixtral-8x7B-v0.1 | 
| Synthetic STEM from Qwen3-235B-A22B-Instruct-2507 and gpt-oss-120b | Text | Undisclosed |  | Qwen3-235B-A22B-Instruct-2507; gpt-oss-120b | 
| Synthetic KernelBook from DeepSeek-R1-0528 | Text | Undisclosed | KernelBook | DeepSeek-R1-0528 | 
| Synthetic Tool Calling from Qwen3-235B-A22B-Thinking-2507 and Qwen3-Next-80B-A3B-Thinking | Text | Undisclosed | ToolBench; glaive-function-calling-v2; APIGen Function-Calling; Nemotron-Personas-USA | Qwen3-235B-A22B-Thinking-2507; Qwen3-Next-80B-A3B-Thinking | 
| Synthetic Chat from gpt-oss-120b, Mixtral-8x22B-Instruct-v0.1, Qwen3-235B-A22B-Instruct-2507 , and Qwen3-235B-A22B-Thinking-2507 | Text | Undisclosed | C4; LMSYS-Chat-1M; ShareGPT; GSM8K; PRM800K; FinQA; WikiTableQuestions; Riddles; glaive-function-calling-v2; SciBench; tigerbot-kaggle-leetcodesolutions-en-2k; OpenBookQA; Advanced Reasoning Benchmark; Software Heritage; Khan Academy Math Keywords; WildChat-1M; Nemotron-Personas-USA | gpt-oss-120b; Mixtral-8x22B-Instruct-v0.1; Qwen3-235B-A22B-Instruct-2507; Qwen3-235B-A22B-Thinking-2507 | 
| Synthetic Long Context from Qwen3-235B-A22B-Instruct-2507 | Text | Undisclosed | CORE; PG-19; DOAB CC BY & CC BY-SA subset; NDLTD | Qwen3-235B-A22B-Instruct-2507 | 
|  | Text | Undisclosed | NVIDIA Internal | gpt-oss-120b; DeepSeek-R1-0528; Qwen3-32B; and Qwen3-235B-A22B-Thinking-2507 | 
| Synthetic STEM from Qwen3-235B-A22B-Thinking-2507 | Text | Undisclosed | ICHO-IPH0; Physics Big; Scale HLE; OpenMathReasoning; OpenCodeReasoning | Qwen3-235B-A22B-Thinking-2507 | 
| Synthetic DocFinQA and SWE-smith from Qwen3-Coder-480B-A35B-Instruct and Kimi-K2-Thinking | Text | Undisclosed | DocFinQA; SWE-smith | Qwen3-Coder-480B-A35B-Instruct; Kimi-K2-Thinking | 
| Synthetic Math from gpt-oss-120b and Qwen2.5-32B-Instruct | Text | Undisclosed | - | gpt-oss-120b; Qwen2.5-32B-Instruct | 
| Synthetic Essential-Web from gpt-oss-120b | Text | Undisclosed | Essential-Web | gpt-oss-120b | 
| Synthetic Scale HLE from gpt-oss-120b | Text | Undisclosed | Scale HLE | gpt-oss-120b | 
| Synthetic CDQuestions from gpt-oss-120b | Text | Undisclosed | CDQuestions | gpt-oss-120b | 
| Synthetic Stack Exchange from gpt-oss-120b | Text | Undisclosed | Stack Exchange | gpt-oss-120b | 
| Synthetic GPQA from gpt-oss-120b and Qwen2.5-32B-Instruct | Text | Undisclosed | Stack Exchange | gpt-oss-120b; Qwen2.5-32B-Instruct | 
| Synthetic Vedantu from gpt-oss-120b | Text | Undisclosed | Vedantu | gpt-oss-120b | 
| Synthetic SWE-Gym and R2E-Gym-Subset from Qwen3-Coder-480B-A35B-Instruct | Text | Undisclosed | SWE-Gym; R2E-Gym-Subset | Qwen3-Coder-480B-A35B-Instruct | 
| Synthetic SWE-Gym from Qwen3-Coder-480B-A35B-Instruct | Text | Undisclosed | SWE-Gym | Qwen3-Coder-480B-A35B-Instruct | 
| Synthetic SWE-Gym and R2E-Gym-Subset from DeepSeek-R1-0528 | Text | Undisclosed | SWE-Gym; R2E-Gym-Subset | DeepSeek-R1-0528 | 
|  | Text | Undisclosed | HelpSteer2; HelpSteer3; LMSYS-Chat-1M; Nemotron-Personas-USA | gpt-oss-120b; Qwen3-235B-A22B-Instruct-2507; Qwen3-235B-A22B-Thinking-2507 | 
|  | Text | Undisclosed | - | Qwen3-30B-A3B-Instruct-2507; Qwen3-30B-A3B-Thinking-2507; Qwen3-235B-A22B-Instruct-2507; Qwen3-235B-A22B-Thinking-2507 | 
| Synthetic Search STEM MCQ from Qwen3-235B-A22B and DeepSeek-R1-0528 | Text | Undisclosed | - | Qwen3-235B-A22B; DeepSeek-R1-0528 | 
| Synthetic Search STEM OPENQ from DeepSeek-R1-0528 | Text | Undisclosed | - | DeepSeek-R1-0528 | 
| Synthetic OpenSTEM from Qwen2.5-32B-Instruct and DeepSeek-R1-0528 | Text | Undisclosed | - | Qwen2.5-32B-Instruct; DeepSeek-R1-0528 | 
| Synthetic MCQ from Qwen2.5-32B-Instruct and DeepSeek-R1-0528 | Text | Undisclosed | - | Qwen2.5-32B-Instruct; DeepSeek-R1-0528 | 
| Synthetic MCQ10 from DeepSeek-R1-0528 | Text | Undisclosed | - | DeepSeek-R1-0528 | 
| Synthetic MCQ4 from Qwen3-235B-A22B, DeepSeek-R1-0528, and Qwen3-235B-A22B-Instruct-2507 | Text | Undisclosed | - | Qwen3-235B-A22B; DeepSeek-R1-0528; Qwen3-235B-A22B-Instruct-2507 | 
| Synthetic OpenMathReasoning from gpt-oss-120b and Qwen2.5-32B-Instruct | Text | Undisclosed | OpenMathReasoning | gpt-oss-120b; Qwen2.5-32B-Instruct | 
| Synthetic Offline Search MCQA HLE from DeepSeek-R1-0528 | Text | Undisclosed | - | DeepSeek-R1-0528 | 
| Synthetic Offline Search MCQA GPQA from Qwen3-235B-A22B and DeepSeek-R1-0528 | Text | Undisclosed | - | Qwen3-235B-A22B; DeepSeek-R1-0528 | 
|  | Text | Undisclosed | - |  | 
|  | Text | Undisclosed | WildChat-1M; arena-human-preference-140k |  | 
|  | Text | Undisclosed | Nemotron Content Safety Dataset V2; Gretel Synthetic Safety Alignment Dataset; RedTeam-2K; Malicious Tasks; | DeepSeek-R1-0528; gpt-oss-120b; DeepSeek-R1-Distill-Qwen-7B; Qwen3-30B-A3B-Thinking-2507; Qwen3-235B-A22B-Instruct-2507; Mixtral-8x7B-v0.1 | 
| Synthetic Code from Qwen3-32B | Text | Undisclosed | English Common Crawl; English Common Crawl 1.1 | Qwen3-32B | 
| Synthetic OpenCodeReasoning from DeepSeek-R1 | Text | Undisclosed | OpenCodeReasoning | DeepSeek-R1 | 
| Synthetic LIMO from DeepSeek-R1-0528 | Text | Undisclosed | LIMO | DeepSeek-R1-0528 | 
| Synthetic SCP from DeepSeek-R1-0528 | Text | Undisclosed | SCP-116K | DeepSeek-R1-0528 | 
| Synthetic Stack Exchange from DeepSeek-R1-0528 | Text | Undisclosed | Stack Exchange | DeepSeek-R1-0528 | 
| Synthetic Common Crawl from Qwen3-30B-A3B | Text | Undisclosed | Common Crawl | Qwen3-30B-A3B | 
| Synthetic Wikipedia from Qwen3-30B-A3B | Text | Undisclosed | Wikimedia | Qwen3-30B-A3B | 
| Synthetic Essential-Web from Qwen3-30B-A3B and Qwen3-235B-A22B-Thinking-2507 | Text | Undisclosed | Essential-Web | Qwen3-30B-A3B; Qwen3-235B-A22B-Thinking-2507 | 
| Synthetic Textbook Math from Qwen3-30B-A3B, Qwen3-235B-A22B, phi-4 | Text | Undisclosed | Common Crawl; FineMath | Qwen3-30B-A3B; Qwen3-235B-A22B; phi-4 | 
| Synthetic Math and Code from DeepSeek-R1 and DeepSeek-R1-0528 | Text | Undisclosed |  | DeepSeek-R1; DeepSeek-R1-0528 | 
| Synthetic Nemotron-Personas-USA from gpt-oss-120b and Qwen3-8B | Text | Undisclosed | Nemotron-Personas-USA | gpt-oss-120b; Qwen3-8B | 
| Synthetic Text-To-SQL | Text | Undisclosed | - | gpt-oss-120b | 
| Synthetic Agentless SWE | Text | Undisclosed | SWE-Bench-Train; SWE-Fixer-Train; SWE-reBench; SWE-smith | DeepSeek-R1-0528 | 
| Synthetic Search Graph Walk | Text | Undisclosed | - | MiniMax-M2 | 
| Synthetic CUDA 100k | Text | Undisclosed | KernelBook; HuggingFace Transformers; FlashInfer | DeepSeek-R1-0528; gpt-oss-120b | 
| Synthetic Safety | Text | Undisclosed | Nemotron Content Safety Dataset V2; Gretel Synthetic Safety Alignment Dataset; RedTeam-2K; HarmfulTasks | gpt-oss-120b; NVIDIA-Nemotron-Nano-9B-v2; gemma-3-4b-it | 
| Synthetic Agentic Diverse Domains | Text | Undisclosed | - | DeepSeek-R1-0528; Qwen3-235B-A22B-Thinking-2507; Qwen3-235B-A22B-Instruct-2507; Qwen3-32B; gpt-oss-120b; DeepSeek-V3.2 | 
| Synthetic SWE Unverified | Text | Undisclosed | - | gpt-oss-120b; Qwen3-Coder-480B-A35B-Instruct; GLM-4.7-Flash | 
| Synthetic Scale HLE from Deepseek-V3 | Text | Undisclosed | Scale HLE | DeepSeek-V3-0324 | 
| Synthetic CDQuestions from Deepseek-V3 | Text | Undisclosed | CDQuestions | DeepSeek-V3-0324 | 
| Synthetic Stack Exchange from Deepseek-V3 | Text | Undisclosed | Stack Exchange | DeepSeek-V3-0324 | 
| Synthetic GPQA from Deepseek-V3 | Text | Undisclosed | Stack Exchange | DeepSeek-V3-0324 | 
| Synthetic Vedantu from Deepseek-V3 | Text | Undisclosed | Vedantu | DeepSeek-V3-0324 | 
| Synthetic Tool Call Schema for RL | Text | Undisclosed | ToolBench; glaive-function-calling-v2; APIGen Function-Calling; Nemotron-Personas-USA | Qwen3-235B-A22B-Thinking-2507; Qwen3-Next-80B-A3B-Thinking | 
| Synthetic Data for Search | Text | Undisclosed | Wikimedia | MiniMax-M2 | 
| Synthetic Instruction Following for RL | Text | Undisclosed | - | NVIDIA-Nemotron-Nano-9B-v2; Qwen3-235B-A22B-Thinking-2507 | 
| Synthetic Conversational Agentic Tool-Use RL | Text | Undisclosed | - | DeepSeek-V3.2; DeepSeek-R1-0528; Qwen3-235B-A22B-Thinking-2507; Qwen3-32B; gpt-oss-120b; Qwen3-235B-A22B-Instruct-2507 | 
| Synthetic Terminal Pivot RL | Text | Undisclosed | SWE-smith; Nemotron-Cascade-RL-SWE; Vendor supplied | DeepSeek-V3.2; Qwen3-Coder-480B-A35B-Instruct; Kimi-K2.5; Qwen3-235B-A22B-Instruct-2507 | 

For our post-training recipe, we focused on 9 main languages in addition to English: French, German, Italian, Japanese, Spanish, and Chinese

Those languages were represented in the form of multilingual reasoning and translation tasks.

The following table depicts our sample distribution for the 6 languages and 5 translation pairs.

| Language | Size | 
|---|---|
| English | 13.48M | 
| Italian | 53k | 
| German | 53k | 
| Spanish | 53k | 
| French | 53k | 
| Japanese | 53k | 
| Chinese | 53k | 
| English <-> Italian | 43.2k | 
| English <-> German | 43.2k | 
| English <-> Spanish | 43.2k | 
| English <-> French | 43.2k | 
| English <-> Japanese | 43.2k | 

- **Data Collection Method by dataset** : Hybrid: Human, Synthetic
- **Labeling Method by dataset** : Hybrid: Automated, Human, Synthetic

- **Acceleration Engine:** PyTorch
- **Test Hardware:**
  - NVIDIA Hopper
    - 1-8x H100
    - 1-8x H200
  - NVIDIA Grace Blackwell
    - GB200
- NVIDIA Hopper

NVIDIA believes Trustworthy AI is a shared responsibility and we have established policies and practices to enable development for a wide array of AI applications. When downloaded or used in accordance with our terms of service, developers should work with their internal model team to ensure this model meets requirements for the relevant industry and use case and addresses unforeseen product misuse.

We advise against circumvention of any provided safety guardrails contained in the Model without a substantially similar guardrail appropriate for your use case. For more details: Safety and Explainability Subcards.

Please report model quality, risk, security vulnerabilities or NVIDIA AI Concerns here.

```
@misc{nvidia_nemotron_3_2025,
  title  = {NVIDIA Nemotron 3: Efficient and Open Intelligence},
  author = {{NVIDIA}},
  year   = {2025},
  url    = {https://arxiv.org/abs/2512.20856},
  note   = {White Paper}
}
```
- Downloads last month
- 836,387

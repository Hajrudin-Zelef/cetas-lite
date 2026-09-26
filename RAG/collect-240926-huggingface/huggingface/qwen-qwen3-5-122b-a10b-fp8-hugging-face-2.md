---
id: collect-240926-huggingface/huggingface/qwen-qwen3-5-122b-a10b-fp8-hugging-face-2
title: "Set the following accordingly"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "OpenAI", "SGLang", "vLLM"]
dates: ["2025-08-07"]
keywords: ["agent", "claude", "context window", "fp8", "gpus", "inference", "kv cache", "memory", "multimodal", "qwen", "reasoning", "sglang"]
source: docs/RAG/clean_en/huggingface/qwen-qwen3-5-122b-a10b-fp8-hugging-face.md
source_anchor: ""
source_lines: [109, 229]
sha256: bae902e2782aba7742c8940bc7eb5c67ea23998484ca0878c8f0c9edbbb7ac02
---

# Set the following accordingly

|  | GPT-5-mini 2025-08-07 | Claude-Sonnet-4.5 | Qwen3-VL-235B-A22B | Qwen3.5-122B-A10B | Qwen3.5-27B | Qwen3.5-35B-A3B | 
|---|---|---|---|---|---|---|
| STEM and Puzzle |  |  |  |  |  |  | 
| MMMU | 79.0 | 79.6 | 80.6 | 83.9 | 82.3 | 81.4 | 
| MMMU-Pro | 67.3 | 68.4 | 69.3 | 76.9 | 75.0 | 75.1 | 
| MathVision | 71.9 | 71.1 | 74.6 | 86.2 | 86.0 | 83.9 | 
| Mathvista(mini) | 79.1 | 79.8 | 85.8 | 87.4 | 87.8 | 86.2 | 
| DynaMath | 81.4 | 78.8 | 82.8 | 85.9 | 87.7 | 85.0 | 
| ZEROBench | 3 | 4 | 4 | 9 | 10 | 8 | 
| ZEROBench_sub | 27.3 | 26.3 | 28.4 | 36.2 | 36.2 | 34.1 | 
| VlmsAreBlind | 75.8 | 85.5 | 79.5 | 96.7 | 96.9 | 97.0 | 
| BabyVision | 20.9 | 18.6 | 22.2 | 40.2 / 34.5 | 44.6 / 34.8 | 38.4 / 29.6 | 
| General VQA |  |  |  |  |  |  | 
| RealWorldQA | 79.0 | 70.3 | 81.3 | 85.1 | 83.7 | 84.1 | 
| MMStar | 74.1 | 73.8 | 78.7 | 82.9 | 81.0 | 81.9 | 
| MMBench <sub>EN-DEV-v1.1</sub> | 86.8 | 88.3 | 89.7 | 92.8 | 92.6 | 91.5 | 
| SimpleVQA | 56.8 | 57.6 | 61.3 | 61.7 | 56.0 | 58.3 | 
| HallusionBench | 63.2 | 59.9 | 66.7 | 67.6 | 70.0 | 67.9 | 
| Text Recognition and Document Understanding |  |  |  |  |  |  | 
| OmniDocBench1.5 | 77.0 | 85.8 | 84.5 | 89.8 | 88.9 | 89.3 | 
| CharXiv(RQ) | 68.6 | 67.2 | 66.1 | 77.2 | 79.5 | 77.5 | 
| MMLongBench-Doc | 50.3 | -- | 56.2 | 59.0 | 60.2 | 59.5 | 
| CC-OCR | 70.8 | 68.1 | 81.5 | 81.8 | 81.0 | 80.7 | 
| AI2D_TEST | 88.2 | 87.0 | 89.2 | 93.3 | 92.9 | 92.6 | 
| OCRBench | 82.1 | 76.6 | 87.5 | 92.1 | 89.4 | 91.0 | 
| Spatial Intelligence |  |  |  |  |  |  | 
| ERQA | 54.0 | 45.0 | 52.5 | 62.0 | 60.5 | 64.8 | 
| CountBench | 91.0 | 90.0 | 93.7 | 97.0 | 97.8 | 97.8 | 
| RefCOCO(avg) | -- | -- | 91.1 | 91.3 | 90.9 | 89.2 | 
| ODInW13 | -- | -- | 43.2 | 44.5 | 41.1 | 42.6 | 
| EmbSpatialBench | 80.7 | 71.8 | 84.3 | 83.9 | 84.5 | 83.1 | 
| RefSpatialBench | 9.0 | 2.2 | 69.9 | 69.3 | 67.7 | 63.5 | 
| LingoQA | 62.4 | 12.8 | 66.8 | 80.8 | 82.0 | 79.2 | 
| Hypersim | -- | -- | 11.0 | 12.7 | 13.0 | 13.1 | 
| SUNRGBD | -- | -- | 34.9 | 36.2 | 35.4 | 33.4 | 
| Nuscene | -- | -- | 13.9 | 15.4 | 15.2 | 14.6 | 
| Video Understanding |  |  |  |  |  |  | 
| VideoMME <sub>(w sub.)</sub> | 83.5 | 81.1 | 83.8 | 87.3 | 87.0 | 86.6 | 
| VideoMME <sub>(w/o sub.)</sub> | 78.9 | 75.3 | 79.0 | 83.9 | 82.8 | 82.5 | 
| VideoMMMU | 82.5 | 77.6 | 80.0 | 82.0 | 82.3 | 80.4 | 
| MLVU | 83.3 | 72.8 | 83.8 | 87.3 | 85.9 | 85.6 | 
| MVBench | -- | -- | 75.2 | 76.6 | 74.6 | 74.8 | 
| LVBench | -- | -- | 63.6 | 74.4 | 73.6 | 71.4 | 
| MMVU | 69.8 | 70.6 | 71.1 | 74.7 | 73.3 | 72.3 | 
| Visual Agent |  |  |  |  |  |  | 
| ScreenSpot Pro | -- | 36.2 | 62.0 | 70.4 | 70.3 | 68.6 | 
| OSWorld-Verified | -- | 61.4 | 38.1 | 58.0 | 56.2 | 54.5 | 
| AndroidWorld | -- | -- | 63.7 | 66.4 | 64.2 | 71.1 | 
| Tool Calling |  |  |  |  |  |  | 
| TIR-Bench | 24.6 | 27.6 | 29.8 | 53.2 / 42.5 | 59.8 / 42.3 | 55.5 / 38.0 | 
| V* | 71.7 | 58.6 | 85.9 | 93.2 / 90.1 | 93.7 / 89.0 | 92.7 / 89.5 | 
| Medical VQA |  |  |  |  |  |  | 
| SLAKE | 70.5 | 73.6 | 54.7 | 81.6 | 80.0 | 78.7 | 
| PMC-VQA | 36.3 | 55.9 | 41.2 | 63.3 | 62.4 | 62.0 | 
| MedXpertQA-MM | 34.4 | 54.0 | 47.6 | 67.3 | 62.4 | 61.4 | 

* MathVision: our model’s score is evaluated using a fixed prompt, e.g., “Please reason step by step, and put your final answer within \boxed{}.” For other models, we report the higher score between runs with and without the \boxed{} formatting.

* BabyVision: scores reported as "with CI / without CI".

* TIR-Bench and V*: scores reported as "with CI / without CI".

* Empty cells (--) indicate scores not yet available or not applicable.

Qwen3.5 models operate in thinking mode by default, generating thinking content signified by `<think>\n...</think>\n\n` before producing the final responses.
To disable thinking content and obtain direct response, refer to the examples here.


For streamlined integration, we recommend using Qwen3.5 via APIs. Below is a guide to use Qwen3.5 via OpenAI-compatible API.

Qwen3.5 can be served via APIs with popular inference frameworks. In the following, we show example commands to launch OpenAI-Compatible API servers for Qwen3.5 models.

Inference efficiency and throughput vary significantly across frameworks. We recommend using the latest framework versions to ensure optimal performance and compatibility. For production workloads or high-throughput scenarios, dedicated serving engines such as SGLang, KTransformers or vLLM are strongly recommended.


The model has a default context length of 262,144 tokens. If you encounter out-of-memory (OOM) errors, consider reducing the context window. However, because Qwen3.5 leverages extended context for complex tasks, we advise maintaining a context length of at least 128K tokens to preserve thinking capabilities.


SGLang is a fast serving framework for large language models and vision language models. SGLang from the main branch of the open-source repository is required for Qwen3.5, which can be installed using the following command in a fresh environment:

```
uv pip install 'git+https://github.com/sgl-project/sglang.git#subdirectory=python&egg=sglang[all]'
```
See its documentation for more details.

The following will create API endpoints at `http://localhost:8000/v1`:

- **Standard Version** : The following command can be used to create an API endpoint with maximum context length 262,144 tokens using tensor parallel on 8 GPUs.```
python -m sglang.launch_server --model-path Qwen/Qwen3.5-122B-A10B-FP8 --port 8000 --tp-size 8 --mem-fraction-static 0.8 --context-length 262144 --reasoning-parser qwen3
```
- **Tool Use** : To support tool use, you can use the following command.```
python -m sglang.launch_server --model-path Qwen/Qwen3.5-122B-A10B-FP8 --port 8000 --tp-size 8 --mem-fraction-static 0.8 --context-length 262144 --reasoning-parser qwen3 --tool-call-parser qwen3_coder
```
- **Multi-Token Prediction (MTP)** : The following command is recommended for MTP:```
python -m sglang.launch_server --model-path Qwen/Qwen3.5-122B-A10B-FP8 --port 8000 --tp-size 8 --mem-fraction-static 0.8 --context-length 262144 --reasoning-parser qwen3 --speculative-algo NEXTN --speculative-num-steps 3 --speculative-eagle-topk 1 --speculative-num-draft-tokens 4
```

vLLM is a high-throughput and memory-efficient inference and serving engine for LLMs. vLLM from the main branch of the open-source repository is required for Qwen3.5, which can be installed using the following command in a fresh environment:

```
uv pip install vllm --torch-backend=auto --extra-index-url https://wheels.vllm.ai/nightly
```
See its documentation for more details.

For detailed Qwen3.5 usage guide, see the vLLM Qwen3.5 recipe.

The following will create API endpoints at `http://localhost:8000/v1`:

- **Standard Version** : The following command can be used to create an API endpoint with maximum context length 262,144 tokens using tensor parallel on 8 GPUs.```
vllm serve Qwen/Qwen3.5-122B-A10B-FP8 --port 8000 --tensor-parallel-size 8 --max-model-len 262144 --reasoning-parser qwen3 
```
- **Tool Call** : To support tool use, you can use the following command.```
vllm serve Qwen/Qwen3.5-122B-A10B-FP8 --port 8000 --tensor-parallel-size 8 --max-model-len 262144 --reasoning-parser qwen3 --enable-auto-tool-choice --tool-call-parser qwen3_coder 
```
- **Multi-Token Prediction (MTP)** : The following command is recommended for MTP:```
vllm serve Qwen/Qwen3.5-122B-A10B-FP8 --port 8000 --tensor-parallel-size 8 --max-model-len 262144 --reasoning-parser qwen3 --speculative-config '{"method":"qwen3_next_mtp","num_speculative_tokens":2}'
```
- **Text-Only** : The following command skips the vision encoder and multimodal profiling to free up memory for additional KV cache:```
vllm serve Qwen/Qwen3.5-122B-A10B-FP8 --port 8000 --tensor-parallel-size 8 --max-model-len 262144 --reasoning-parser qwen3 --language-model-only
```


---
id: collect-240926-huggingface/huggingface/qwen-qwen3-8-flash-next-hugging-face-4
title: "Set the following accordingly"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "SGLang", "vLLM"]
dates: []
keywords: ["agentic", "cost", "embeddings", "inference", "leaderboard", "parameters", "qwen", "reasoning", "rotary", "sglang", "training", "vllm"]
source: docs/RAG/clean_en/huggingface/qwen-qwen3-8-flash-next-hugging-face.md
source_anchor: ""
source_lines: [337, 409]
sha256: ae8731a0b12025a11309f3d599a1c9622db0ede839a89807aa2d80b593c0040d
---

# Set the following accordingly

1. **Sampling Parameters** : We suggest using the following sets of sampling parameters:
  - Thinking Mode: `temperature=1.0` ,`top_p=0.95` ,`top_k=20` ,`min_p=0.0` ,`presence_penalty=0.0` ,`repetition_penalty=1.0`
  - Instruct (or non-thinking) mode: `temperature=0.7` ,`top_p=0.80` ,`top_k=20` ,`min_p=0.0` ,`presence_penalty=1.5` ,`repetition_penalty=1.0`
 For supported frameworks, you can adjust the `presence_penalty` parameter between 0 and 2 to reduce endless repetition. However, using a higher value may occasionally result in language mixing and a slight decrease in model performance.
2. Thinking Mode: 
3. **Adequate Output Length** : To optimize performance on agentic tasks, we recommend allocating sufficient output length to allow the model to generate detailed and comprehensive responses. For frameworks that support separate token limits for internal reasoning and final outputs, we suggest the following configuration within the 1M context length:
  - Reasoning Content: Set the maximum output length to 262,144 tokens.
  - Final Response: Set the maximum output length to 131,072 tokens.
 These settings provide the necessary capacity for complex reasoning while ensuring ample space for high-quality final deliverables.
4. **Processing Ultra-Long Texts** : Qwen3.8-Flash-Next natively supports context lengths of up to 262,144 tokens. For long-horizon tasks where the total length (including both input and output) exceeds this limit, we recommend using RoPE scaling techniques to handle long texts effectively, e.g., YaRN.YaRN is currently supported by several inference frameworks, e.g., vLLM, SGLang, and TokenSpeed. In general, there are two approaches to enabling YaRN for supported frameworks: 
  - Modifying the model configuration file: In the `config.json` file, change the`rope_parameters` fields in`text_config` to:```
{
    "mrope_interleaved": true,
    "mrope_section": [
        11,
        11,
        10
    ],
    "rope_type": "yarn",
    "rope_theta": 10000000,
    "partial_rotary_factor": 0.25,
    "factor": 4.0,
    "original_max_position_embeddings": 262144
}
```
  - Passing command line arguments: For vLLM, you can use ```
VLLM_ALLOW_LONG_MAX_MODEL_LEN=1 vllm serve ... --hf-overrides '{"text_config": {"rope_parameters": {"mrope_interleaved": true, "mrope_section": [11, 11, 10], "rope_type": "yarn", "rope_theta": 10000000, "partial_rotary_factor": 0.25, "factor": 4.0, "original_max_position_embeddings": 262144}}}' --max-model-len 1000000  
```
For SGLang, you can use ```
SGLANG_ALLOW_OVERWRITE_LONGER_CONTEXT_LEN=1 python -m sglang.launch_server ... --json-model-override-args '{"text_config": {"rope_parameters": {"mrope_interleaved": true, "mrope_section": [11, 11, 10], "rope_type": "yarn", "rope_theta": 10000000, "partial_rotary_factor": 0.25, "factor": 4.0, "original_max_position_embeddings": 262144}}}' --context-length 1000000
```
For TokenSpeed, you can use ```
TOKENSPEED_ALLOW_OVERWRITE_LONGER_CONTEXT_LEN=1 tokenspeed serve ... --hf-overrides '{"text_config": {"rope_parameters": {"mrope_interleaved": true, "mrope_section": [11, 11, 10], "rope_type": "yarn", "rope_theta": 10000000, "partial_rotary_factor": 0.25, "factor": 4.0, "original_max_position_embeddings": 262144}}}' --max-model-len 1000000  
```
 All the notable open-source frameworks implement static YaRN, which means the scaling factor remains constant regardless of input length, **potentially impacting performance on shorter texts.** We advise modifying the`rope_parameters` configuration only when processing long contexts is required. 
It is also recommended to modify the`factor` as needed. For example, if the typical context length for your application is 524,288 tokens, it would be better to set`factor` as 2.0.
5. **Long Video Understanding** : To optimize inference efficiency for plain text and images, the`size` parameter in the released`video_preprocessor_config.json` is conservatively configured. It is recommended to set the`longest_edge` parameter in the video_preprocessor_config file to 469,762,048 (corresponding to 224k video tokens) to enable higher frame-rate sampling for hour-scale videos and thereby achieve superior performance. For example,```
{"longest_edge": 469762048, "shortest_edge": 4096}
```
Alternatively, override the default values via engine startup parameters. For implementation details, refer to: vLLM / SGLang.

If you find our work helpful, feel free to give us a cite.

```
@techreport{qwen2026design,
    title       = {On the Design of {Qwen3.8-Next} Architecture: Evaluation, Efficiency, and Training Stability},
    author      = {{Qwen Team}},
    institution = {Alibaba Group},
    month       = {August},
    year        = {2026}
}
@misc{qwen3.8flashnext,
    title  = {{Qwen3.8-Flash-Next}: A New Architecture, Towards Ultimate Cost-Efficiency},
    author = {{Qwen Team}},
    month  = {August},
    year   = {2026},
    url    = {https://qwen.ai/blog?id=qwen3.8-flash-next}
}
```
- Downloads last month
- 830,208

## Spaces using Qwen/Qwen3.8-Flash-Next 13

## Collection including Qwen/Qwen3.8-Flash-Next

- Idavidrein/gpqa · Diamond View evaluation results leaderboard
- datacurve/deep-swe · Deep Swe View evaluation results    leaderboard  58.7<sup>*</sup>
- ScaleAI/SWE-bench_Pro · SWE Bench Pro View evaluation results leaderboard
- llamaindex/ExtractBench leaderboard
- Mean View evaluation resultssource
- Short View evaluation resultssource
- Medium View evaluation resultssource

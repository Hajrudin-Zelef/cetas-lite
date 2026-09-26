---
id: collect-240926-huggingface/huggingface/deepseek-ai-deepseek-v4-1-flash-hugging-face-2
title: "deepseek-ai-deepseek-v4-1-flash-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["Anthropic", "DeepSeek"]
dates: []
keywords: ["deepseek", "agent", "benchmark", "claude", "inference", "kv cache", "leaderboard", "license", "mit license", "parameters", "reasoning", "tool calling"]
source: docs/RAG/clean_en/huggingface/deepseek-ai-deepseek-v4-1-flash-hugging-face.md
source_anchor: ""
source_lines: [82, 129]
sha256: 7a9e3ac0a98dfb50590b45b5dda5877c4ab41224e3124f303b4c983338832d60
---

# deepseek-ai-deepseek-v4-1-flash-hugging-face

All scaffolds use N=8 samples per task on DeepSWE v1.1 and N=3 on Terminal-Bench 2.1, with Linux containers, `temperature=1.0`, `top_p=0.95`, a 1M-token context limit, and max_steps=500 per agent. Terminal-Bench 2.1 is evaluated without network access.

| Benchmark (Metric) | Claude Code | Codex | OpenCode | Pi | mini-SWE | DSH Minimal | DSH Standard | DSH PTC | 
|---|---|---|---|---|---|---|---|---|
| DeepSWE v1.1 (Resolved) | 69.8 | 65.6 | 65.5 | 66.2 | 74.2 | 72.6 | 70.5 | 67.6 | 
| Terminal-Bench 2.1 (Pass@1) | 88.0 | 84.1 | 85.0 | 86.1 | 90.3 | 90.6 | 85.8 | 85.8 | 

This release does not include a Jinja-format chat template. The `encoding` folder contains a self-contained Python reference implementation (`encoding.py`) with test cases for multi-turn conversations, tool calling, thinking mode, numeric reasoning effort, mid-conversation system messages, and interleaved image content.

For production use, we additionally release deepseek-recipe, a set of Rust libraries with Python bindings that provides the same prompt format as a maintained, protocol-aware toolkit. It converts Messages, Chat Completions, and Responses API requests into the Conversation format, encodes them into DeepSeek V4 and V4.1 prompts or token IDs, and parses model output back into complete or streamed responses — covering thinking, tool calls, images, and generation settings. Model inference, tool execution, and HTTP transport are left to the caller.

Please refer to the `inference` folder for instructions on weight conversion and running inference locally.

**Recommended sampling parameters:**

| Parameter | Value | 
|---|---|
| `temperature` | 1.0 | 
| `top_p` | 0.95 or 1.0 | 
| `context_window` | 1M tokens | 
| `max_tokens` | ≥ 256K | 

The `evaluation` folder contains step-by-step instructions for reproducing the DeepSWE v1.1 benchmark results, covering both the `dsh-minimal` agent and the official `mini-swe-agent`. The patch required to integrate `dsh-minimal` with Pier is also included there.

This repository and the model weights are licensed under the MIT License.

```
@misc{deepseekai2026deepseekv41flash,
      title={DeepSeek-V4.1-Flash: Pushing the Limits of KV Cache Compression},
      author={DeepSeek-AI},
      year={2026},
}
```
If you have any questions, please raise an issue or contact us at service@deepseek.com.

- Downloads last month
- 606,028

## Spaces using deepseek-ai/DeepSeek-V4.1-Flash 29

## Collection including deepseek-ai/DeepSeek-V4.1-Flash

- Idavidrein/gpqa · Diamond View evaluation results    leaderboard  90.9
- datacurve/deep-swe · Deep Swe View evaluation results leaderboard
- llamaindex/ExtractBench leaderboard
- Mean View evaluation resultssourcePipeline name: deepseek_v4_1_flash_extract_oneshot_structured_output_file (served via the DeepSeek API, thinking disabled)87.11<sup>*</sup>
- Short View evaluation resultssourcePipeline name: deepseek_v4_1_flash_extract_oneshot_structured_output_file (served via the DeepSeek API, thinking disabled)94.44<sup>*</sup>
- Medium View evaluation resultssourcePipeline name: deepseek_v4_1_flash_extract_oneshot_structured_output_file (served via the DeepSeek API, thinking disabled)81.49<sup>*</sup>

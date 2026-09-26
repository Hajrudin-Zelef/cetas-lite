---
id: collect-240926-huggingface/huggingface/deepseek-ai-deepseek-v4-pro-hugging-face-2
title: "messages -> string"
domain: huggingface
role: reference
task: reference
actors: ["China", "DeepSeek", "OpenAI"]
dates: []
keywords: ["agentic", "benchmark", "context window", "deepseek", "inference", "leaderboard", "license", "mit license", "parameters", "reasoning"]
source: docs/RAG/clean_en/huggingface/deepseek-ai-deepseek-v4-pro-hugging-face.md
source_anchor: ""
source_lines: [94, 170]
sha256: 5594a6ea56e2c70649c5457056ebf495becf1fe7d5e339820a0492e7d56e0ed3
---

# messages -> string

| Benchmark (Metric) | V4-Flash Non-Think | V4-Flash High | V4-Flash Max | V4-Pro Non-Think | V4-Pro High | V4-Pro Max | 
|---|---|---|---|---|---|---|
| **Knowledge & Reasoning** |  |  |  |  |  |  | 
| MMLU-Pro (EM) | 83.0 | 86.4 | 86.2 | 82.9 | 87.1 | **87.5** | 
| SimpleQA-Verified (Pass@1) | 23.1 | 28.9 | 34.1 | 45.0 | 46.2 | **57.9** | 
| Chinese-SimpleQA (Pass@1) | 71.5 | 73.2 | 78.9 | 75.8 | 77.7 | **84.4** | 
| GPQA Diamond (Pass@1) | 71.2 | 87.4 | 88.1 | 72.9 | 89.1 | **90.1** | 
| HLE (Pass@1) | 8.1 | 29.4 | 34.8 | 7.7 | 34.5 | **37.7** | 
| LiveCodeBench (Pass@1) | 55.2 | 88.4 | 91.6 | 56.8 | 89.8 | **93.5** | 
| Codeforces (Rating) | - | 2816 | 3052 | - | 2919 | **3206** | 
| HMMT 2026 Feb (Pass@1) | 40.8 | 91.9 | 94.8 | 31.7 | 94.0 | **95.2** | 
| IMOAnswerBench (Pass@1) | 41.9 | 85.1 | 88.4 | 35.3 | 88.0 | **89.8** | 
| Apex (Pass@1) | 1.0 | 19.1 | 33.0 | 0.4 | 27.4 | **38.3** | 
| Apex Shortlist (Pass@1) | 9.3 | 72.1 | 85.7 | 9.2 | 85.5 | **90.2** | 
| **Long Context** |  |  |  |  |  |  | 
| MRCR 1M (MMR) | 37.5 | 76.9 | 78.7 | 44.7 | 83.3 | **83.5** | 
| CorpusQA 1M (ACC) | 15.5 | 59.3 | 60.5 | 35.6 | 56.5 | **62.0** | 
| **Agentic** |  |  |  |  |  |  | 
| Terminal Bench 2.0 (Acc) | 49.1 | 56.6 | 56.9 | 59.1 | 63.3 | **67.9** | 
| SWE Verified (Resolved) | 73.7 | 78.6 | 79.0 | 73.6 | 79.4 | **80.6** | 
| SWE Pro (Resolved) | 49.1 | 52.3 | 52.6 | 52.1 | 54.4 | **55.4** | 
| SWE Multilingual (Resolved) | 69.7 | 70.2 | 73.3 | 69.8 | 74.1 | **76.2** | 
| BrowseComp (Pass@1) | - | 53.5 | 73.2 | - | 80.4 | **83.4** | 
| HLE w/ tools (Pass@1) | - | 40.3 | 45.1 | - | 44.7 | **48.2** | 
| MCPAtlas (Pass@1) | 64.0 | 67.4 | 69.0 | 69.4 | **74.2** | 73.6 | 
| GDPval-AA (Elo) | - | - | 1395 | - | - | **1554** | 
| Toolathlon (Pass@1) | 40.7 | 43.5 | 47.8 | 46.3 | 49.0 | **51.8** | 

This release does not include a Jinja-format chat template. Instead, we provide a dedicated `encoding` folder with Python scripts and test cases demonstrating how to encode messages in OpenAI-compatible format into input strings for the model, and how to parse the model's text output. Please refer to the `encoding` folder for full documentation.

A brief example:

```
from encoding_dsv4 import encode_messages, parse_message_from_completion_text
messages = [
    {"role": "user", "content": "hello"},
    {"role": "assistant", "content": "Hello! I am DeepSeek.", "reasoning_content": "thinking..."},
    {"role": "user", "content": "1+1=?"}
]
# messages -> string
prompt = encode_messages(messages, thinking_mode="thinking")
# string -> tokens
import transformers
tokenizer = transformers.AutoTokenizer.from_pretrained("deepseek-ai/DeepSeek-V4-Pro")
tokens = tokenizer.encode(prompt)
```
Please refer to the inference folder for detailed instructions on running DeepSeek-V4 locally, including model weight conversion and interactive chat demos.

For local deployment, we recommend setting the sampling parameters to `temperature = 1.0, top_p = 1.0`. For the Think Max reasoning mode, we recommend setting the context window to at least **384K** tokens.

This repository and the model weights are licensed under the MIT License.

```
@misc{deepseekai2026deepseekv4,
      title={DeepSeek-V4: Towards Highly Efficient Million-Token Context Intelligence},
      author={DeepSeek-AI},
      year={2026},
}
```
If you have any questions, please raise an issue or contact us at service@deepseek.com.

- Downloads last month
- 514,445

## Spaces using deepseek-ai/DeepSeek-V4-Pro 100

## Collection including deepseek-ai/DeepSeek-V4-Pro

## Paper for deepseek-ai/DeepSeek-V4-Pro

- openai/gsm8k · Gsm8k View evaluation results    leaderboard  92.6
- Idavidrein/gpqa · Diamond View evaluation results    leaderboard  90.1
- SWE-bench/SWE-bench_Verified · Swe Bench Resolved View evaluation results    leaderboard  80.6
- ScaleAI/SWE-bench_Pro · SWE Bench Pro View evaluation results    leaderboard  55.4
- IntelligenceLab/Long-Horizon-Terminal-Bench · Lhtb Solved View evaluation results source leaderboard
- TIGER-Lab/MMLU-Pro · Mmlu Pro View evaluation results    leaderboard  87.5
- actava/chi-bench leaderboard

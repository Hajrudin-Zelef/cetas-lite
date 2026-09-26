---
id: collect-240926-huggingface/huggingface/qwen-qwen3-6-35b-a3b-hugging-face-4
title: "Set the following accordingly"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba"]
dates: []
keywords: ["agentic", "inference", "leaderboard", "parameters", "qwen"]
source: docs/RAG/clean_en/huggingface/qwen-qwen3-6-35b-a3b-hugging-face.md
source_anchor: ""
source_lines: [495, 534]
sha256: 489a2549b74e4e62e7943b8c92c9ee9e0dabaf2c35cfda2dccd7414943f2fd4a
---

# Set the following accordingly

1. **Sampling Parameters** :
  - We suggest using the following sets of sampling parameters depending on the mode and task type:  
    - **Thinking mode for general tasks** :`temperature=1.0` ,`top_p=0.95` ,`top_k=20` ,`min_p=0.0` ,`presence_penalty=1.5` ,`repetition_penalty=1.0`
    - **Thinking mode for precise coding tasks (e.g., WebDev)** :`temperature=0.6` ,`top_p=0.95` ,`top_k=20` ,`min_p=0.0` ,`presence_penalty=0.0` ,`repetition_penalty=1.0`
    - **Instruct (or non-thinking) mode** :`temperature=0.7` ,`top_p=0.80` ,`top_k=20` ,`min_p=0.0` ,`presence_penalty=1.5` ,`repetition_penalty=1.0`
  - For supported frameworks, you can adjust the `presence_penalty` parameter between 0 and 2 to reduce endless repetitions. However, using a higher value may occasionally result in language mixing and a slight decrease in model performance.
2. We suggest using the following sets of sampling parameters depending on the mode and task type:  
3. **Adequate Output Length** : We recommend using an output length of 32,768 tokens for most queries. For benchmarking on highly complex problems, such as those found in math and programming competitions, we suggest setting the max output length to 81,920 tokens. This provides the model with sufficient space to generate detailed and comprehensive responses, thereby enhancing its overall performance.
4. **Standardize Output Format** : We recommend using prompts to standardize model outputs when benchmarking.
  - **Math Problems** : Include "Please reason step by step, and put your final answer within \boxed{}." in the prompt.
  - **Multiple-Choice Questions** : Add the following JSON structure to the prompt to standardize responses: "Please show your choice in the`answer` field with only the choice letter, e.g.,`"answer": "C"` ."
5. **Long Video Understanding** : To optimize inference efficiency for plain text and images, the`size` parameter in the released`video_preprocessor_config.json` is conservatively configured. It is recommended to set the`longest_edge` parameter in the video_preprocessor_config file to 469,762,048 (corresponding to 224k video tokens) to enable higher frame-rate sampling for hour-scale videos and thereby achieve superior performance. For example,```
{"longest_edge": 469762048, "shortest_edge": 4096}
```

If you find our work helpful, feel free to give us a cite.

```
@misc{qwen36_35b_a3b,
    title = {{Qwen3.6-35B-A3B}: Agentic Coding Power, Now Open to All},
    url = {https://qwen.ai/blog?id=qwen3.6-35b-a3b},
    author = {{Qwen Team}},
    month = {April},
    year = {2026}
}
```
- Downloads last month
- 3,138,204

## Spaces using Qwen/Qwen3.6-35B-A3B 73

## Collection including Qwen/Qwen3.6-35B-A3B

- Idavidrein/gpqa · Diamond View evaluation results    leaderboard  86
- SWE-bench/SWE-bench_Verified · Swe Bench Resolved View evaluation results    leaderboard  73.4
- ScaleAI/SWE-bench_Pro · SWE Bench Pro View evaluation results    leaderboard  49.5
- MathArena/aime_2026 · MathArena Aime 2026 View evaluation results    leaderboard  92.7
- TIGER-Lab/MMLU-Pro · Mmlu Pro View evaluation results    leaderboard  85.2
- llamaindex/ExtractBench leaderboard
- Mean View evaluation resultssource

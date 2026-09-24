---
id: collect-240926-huggingface/huggingface/unsloth-deepseek-v4-flash-0731-gguf-hugging-face
title: "unsloth-deepseek-v4-flash-0731-gguf-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["DeepSeek", "Unsloth", "Z.ai"]
dates: []
keywords: ["deepseek", "gguf", "agent", "agentic", "agents", "benchmark", "benchmarks", "glm", "license", "mit license", "quantization", "reasoning"]
source: docs/RAG/clean_en/huggingface/unsloth-deepseek-v4-flash-0731-gguf-hugging-face.md
source_anchor: ""
source_lines: [1, 44]
sha256: 0eebffac32d101eca56d07ca52f3988db9951adcb4eaee58246b2cab9455ad46
---

# unsloth-deepseek-v4-flash-0731-gguf-hugging-face

<!-- source: https://huggingface.co/unsloth/DeepSeek-V4-Flash-0731-GGUF -->

## Read our How to Run DeepSeek-V4-0731 Guide!

    *Unsloth Dynamic 2.0 achieves superior accuracy & outperforms other leading quants.*
  

- To run DeepSeek-V4-Flash-0731 in full precision lossless, run Q8 (UD-Q8_K_XL), which is 162GB and only 7GB bigger than Q4 (UD-Q4_K_XL).
- See our DeepSeek-V4 guide for quantization analysis and instructions.
- You can now run DeepSeek-V4-Flash-0731 in Unsloth Studio with toggles for High and Max thinking.
- **New DSpark support allowing up to 2x faster decoding!** Docs for DSpark

**DeepSeek-V4-Flash-0731** is the official release of **DeepSeek-V4-Flash**, superseding the preview version, with substantially enhanced agentic capabilities. It has the same model structure as DeepSeek-V4-Flash-DSpark, i.e. it comes with a speculative decoding module attached.

DeepSeek-V4-Flash-0731 outperforms DeepSeek-V4-Pro (Preview) on benchmarks listed below despite its far smaller activated parameter count, and is broadly competitive with the strongest proprietary models available.

| Benchmark | DeepSeek-V4-Flash-0731 | DeepSeek-V4-Flash (Preview) | DeepSeek-V4-Pro (Preview) | GLM-5.2 | Opus-4.8 | 
|---|---|---|---|---|---|
| Terminal Bench 2.1 | 82.7 | 61.8 | 72.1 | 81.0 | 85.0 | 
| NL2Repo | 54.2 | 39.4 | 38.5 | 48.9 | 69.7 | 
| Cybergym | 76.7 | 38.7 | 52.7 | - | 83.1 | 
| DeepSWE | 54.4 | 7.3 | 12.8 | 46.2 | 58.0 | 
| Toolathlon-Verified | 70.3 | 49.7 | 55.9 | 59.9 | 76.2 | 
| Agents' Last Exam | 25.2 | 15.8 | 16.5 | 23.8 | 25.7 | 
| AutomationBench Public | 25.1 | 10.8 | 12.8 | 12.9 | 27.2 | 
| DSBench-FullStack † | 68.7 | 37.0 | 41.8 | 61.8 | 71.6 | 
| DSBench-Hard † | 59.6 | 25.8 | 31.1 | 54.5 | 71.7 | 

Notes:

1. For the Code Agent tasks among the public benchmarks above, DeepSeek-V4-Flash-0731 is evaluated with the minimal mode of DeepSeek Harness (to be released) as the agent framework, using the `max` reasoning effort level with`temperature = 1.0, top_p = 0.95` .
2. † DSBench-FullStack is an internal full-stack development test set; DSBench-Hard is an internal test set of difficult coding-agent problems.

This repository and the model weights are licensed under the MIT License.

```
@misc{deepseekai2026deepseekv4,
      title={DeepSeek-V4: Towards Highly Efficient Million-Token Context Intelligence},
      author={DeepSeek-AI},
      year={2026},
}
```
- Downloads last month
- 205,279

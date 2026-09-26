---
id: collect-260926-rattrapage/rattrapage/nvidia-nemotron-3-ultra-how-to-run-locally-2
title: "NVIDIA Nemotron 3 Ultra - How To Run Locally"
domain: rattrapage
role: reference
task: reference
actors: ["MiniMax", "Moonshot", "Nvidia", "OpenAI", "Unsloth", "Z.ai"]
dates: []
keywords: ["nvidia", "agent", "agentic", "benchmark", "benchmarks", "cost", "gguf", "glm", "kimi", "reasoning", "research", "throughput"]
source: docs/RAG/lot-rattrapage/servers-reviews/NVIDIA Nemotron 3 Ultra - How To Run Locally.md
source_anchor: ""
source_lines: [145, 218]
sha256: 8de558e18923cc02add849d7ab829409a3ce74bae5272b8155237151284978b7
---

# NVIDIA Nemotron 3 Ultra - How To Run Locally

    --model unsloth/NVIDIA-Nemotron-3-Ultra-550B-A55B-GGUF/UD-IQ3_XXS/NVIDIA-Nemotron-3-Ultra-550B-A55B-UD-IQ3_XXS-00001-of-00006.gguf \
    --alias "unsloth/NVIDIA-Nemotron-3-Ultra-550B-A55B" \
    --temp 1.0 \
    --top-p 0.95 \
    --port 8001
```
{% endcode %}
Then in a new terminal, after installing the OpenAI client with `pip install openai`:
```python
from openai import OpenAI
openai_client = OpenAI(
    base_url = "http://127.0.0.1:8001/v1",
    api_key = "sk-no-key-required",
)
completion = openai_client.chat.completions.create(
    model = "unsloth/NVIDIA-Nemotron-3-Ultra-550B-A55B",
    messages = [
        {"role": "user", "content": "What is 2+2?"},
    ],
)
print(completion.choices[0].message.reasoning_content)
print(completion.choices[0].message.content)
```
And on 4 B200s, around 40 tokens / s is seen for generation!
### Unsloth GGUF Benchmarks
We also did KLD analysis for our GGUF quants - on a log mean KLD scale, the model loses very little accuracy when quantized down to even 1bit due to our [dynamic methodology](/docs/basics/dynamic-3.0-ggufs.md) where more important layers are left in higher precision and the rest in lower bits.
For a linear scale:
### Official Benchmarks
Nemotron 3 Ultra is NVIDIA's largest Nemotron 3 reasoning model and is positioned for leading accuracy on frontier reasoning, coding and agentic tasks while optimizing time to task completion through high throughput.
Ultra is especially suited for workloads where task success depends on sustained reasoning rather than short single-turn responses:
* Autonomous coding sessions across large repositories
* Deep research across many sources with conflicting evidence
* Enterprise workflows with persistent tool-using loops
* EDA / chip design verification and failure analysis
As shown in Figure 1 and Figure 2 Nemotron 3 Ultra leads on accuracy on agent productivity, instruction following, and long context tasks and provides leading throughout, saving 30% on costs compared to other leading open models. 
Figure 1: Nemotron 3 Ultra leads among open models on agentic benchmarks for agent productivity, coding, and instruction following.
Figure 2: Nemotron 3 Ultra saves up to 30% in costs and leads on the cost efficiency frontier
More benchmarks from NVIDIA:
| Benchmark                                     | N-3-Ultra 550B-A55B | MiniMax-2.7 230B-A10B | GLM-5.1 744B-A40B | Kimi-K2.6 1T-A32B |       |       |      |
| --------------------------------------------- | :-----------------: | :-------------------: | :---------------: | :---------------: | :---: | :---: | :--: |
| **Agentic**                                   |                     |                       |                   |                   |       |       |      |
| Terminal Bench 2.1                            |         56.4        |          55.5         |        59.3       |        67.2       |  49.9 |  49.2 | 54.2 |
| GDPVal                                        |         46.7        |          47.6         |        54.7       |        50.4       |  34.6 |  54.6 | 50.2 |
| SWE-Bench Verified                            |         71.9        |          72.2         |        73.8       |        69.5       |  69.9 |  74.0 | 72.4 |
| SWE-Bench Multilingual                        |         67.7        |          69.2         |        73.8       |        65.9       |  67.7 |  71.9 | 72.1 |
| ProfBench (Search)                            |         56.0        |          52.0         |        46.0       |        56.0       |  53.0 |  59.9 | 57.0 |
| PinchBench                                    |         90.0        |          77.6         |        81.2       |        90.2       |  86.6 |  88.6 | 91.3 |
| TauBench V3                                   |                     |                       |                   |                   |       |       |      |
| Airline                                       |         81.5        |          75.3         |        85.0       |        85.8       |  76.5 |  80.8 | 80.8 |
| Retail                                        |         86.4        |          84.9         |        84.1       |        82.9       |  88.5 |  88.9 | 89.1 |
| Telecom                                       |         92.9        |          89.6         |        96.9       |        97.8       |  98.0 |  96.3 | 98.3 |
| Banking                                       |         22.6        |          14.6         |        12.8       |        23.1       |  20.9 |  25.9 | 26.7 |
| Average                                       |         70.9        |          66.1         |        69.7       |        72.4       |  71.0 |  73.2 | 73.7 |
| BrowseComp                                    |         44.4        |          54.1         |        59.4       |        61.3       |  40.5 |  59.4 | 46.9 |
| Vals.ai Financial Agent 1.1                   |                     |                       |                   |                   |       |       |      |
| without web search                            |         60.1        |          51.3         |        60.2       |        54.0       |  61.3 |  58.9 | 58.4 |
| with web search                               |         53.7        |          50.5         |        60.7       |        58.8       |  59.0 |  62.3 | 60.1 |
| **Reasoning and Knowledge**                   |                     |                       |                   |                   |       |       |      |
| IOI 2025                                      |        570.0        |           --          |       456.5       |       585.0       | 441.3 | 580.1 |  --  |
| LiveCodeBench (v6)                            |         89.0        |          77.2         |        85.7       |        90.2       |  79.3 |  92.5 | 90.9 |
| IMOAnswerBench (no tools)                     |         88.6        |          68.3         |        86.8       |        91.1       |  83.1 |  93.0 | 91.1 |
| IMOAnswerBench (with tools)                   |         92.3        |          75.1         |        91.1       |       93.71       | 84.51 |  85.4 | 89.6 |
| Apex-Shortlist (no tools)                     |         74.9        |          28.9         |        71.1       |        77.4       |  61.4 |  85.8 | 82.4 |
| Apex-Shortlist (with tools)                   |         84.8        |          51.9         |        79.0       |        73.2       |  60.4 |  86.5 | 82.0 |
| GPQA (no tools)                               |         87.0        |          86.6         |        86.1       |        91.0       |  87.1 |  87.8 | 88.5 |
| SciCode (subtask)                             |         44.6        |          38.3         |        47.7       |        52.0       |  48.0 |  50.5 | 48.2 |
| HLE (no tools)                                |         26.7        |          23.1         |        27.2       |        34.8       |  28.5 |  37.7 | 32.2 |
| HLE (with tools)                              |         37.4        |           --          |        50.4       |        54.0       |  48.3 |  48.2 | 45.1 |
| CritPt (no tools)                             |         3.1         |          0.6          |        3.7        |        9.1        |  2.4  |  14.0 | 10.6 |
| MMLU-Pro                                      |         86.8        |          81.9         |        85.9       |        88.1       |  88.3 |  87.5 | 86.4 |
| OmniScience Accuracy                          |         24.1        |          20.5         |        31.3       |        35.5       |  35.9 |  46.8 | 39.9 |
| OmniScience Non-Hallucination                 |         78.7        |          74.4         |        66.8       |        67.1       |  7.4  |  5.7  |  2.8 |
| **Chat & Instruction Following**              |                     |                       |                   |                   |       |       |      |
| IFBench (prompt loose)                        |         81.7        |          74.6         |        76.6       |        73.7       |  78.2 |  79.1 | 82.0 |

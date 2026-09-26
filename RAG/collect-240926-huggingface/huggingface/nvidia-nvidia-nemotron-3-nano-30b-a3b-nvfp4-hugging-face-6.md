---
id: collect-240926-huggingface/huggingface/nvidia-nvidia-nemotron-3-nano-30b-a3b-nvfp4-hugging-face-6
title: "Load tokenizer and model"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "China", "DeepSeek", "MiniMax", "Mistral", "Moonshot", "Nvidia"]
dates: []
keywords: ["alignment", "deepseek", "kimi", "mistral", "nvidia", "qwen"]
source: docs/RAG/clean_en/huggingface/nvidia-nvidia-nemotron-3-nano-30b-a3b-nvfp4-hugging-face.md
source_anchor: ""
source_lines: [446, 520]
sha256: be98aafe7c3f2b25fcfa68477d267939af194b94b20527979f8983a9ef79a3e6
---

# Load tokenizer and model

| Synthetic SWE-Gym from Qwen3-Coder-480B-A35B-Instruct | Text | Undisclosed | SWE-Gym | Qwen3-Coder-480B-A35B-Instruct | 
| Synthetic SWE-Gym and R2E-Gym-Subset from DeepSeek-R1-0528 | Text | Undisclosed | SWE-Gym; R2E-Gym-Subset | DeepSeek-R1-0528 | 
| Synthetic HelpSteer, LMSYS-Chat-1M, and Nemotron-Personas-USA from gpt-oss-120b, Qwen3-235B-A22B-Instruct-2507, and Qwen3-235B-A22B-Thinking-2507 | Text | Undisclosed | HelpSteer2; HelpSteer3; LMSYS-Chat-1M; Nemotron-Personas-USA | gpt-oss-120b; Qwen3-235B-A22B-Instruct-2507; Qwen3-235B-A22B-Thinking-2507 | 
| Synthetic Structured Outputs from Qwen3-30B-A3B-Instruct-2507, Qwen3-30B-A3B-Thinking-2507, Qwen3-235B-A22B-Instruct-2507, and Qwen3-235B-A22B-Thinking-2507 | Text | Undisclosed | - | Qwen3-30B-A3B-Instruct-2507; Qwen3-30B-A3B-Thinking-2507; Qwen3-235B-A22B-Instruct-2507; Qwen3-235B-A22B-Thinking-2507 | 
| Synthetic Search STEM MCQ from Qwen3-235B-A22B and DeepSeek-R1-0528 | Text | Undisclosed | - | Qwen3-235B-A22B; DeepSeek-R1-0528 | 
| Synthetic Search STEM OPENQ from DeepSeek-R1-0528 | Text | Undisclosed | - | DeepSeek-R1-0528 | 
| Synthetic OpenSTEM from Qwen2.5-32B-Instruct and DeepSeek-R1-0528 | Text | Undisclosed | - | Qwen2.5-32B-Instruct; DeepSeek-R1-0528 | 
| Synthetic MCQ from Qwen2.5-32B-Instruct and DeepSeek-R1-0528 | Text | Undisclosed | - | Qwen2.5-32B-Instruct; DeepSeek-R1-0528 | 
| Synthetic MCQ10 from DeepSeek-R1-0528 | Text | Undisclosed | - | DeepSeek-R1-0528 | 
| Synthetic MCQ4 from Qwen3-235B-A22B, DeepSeek-R1-0528, and Qwen3-235B-A22B-Instruct-2507 | Text | Undisclosed | - | Qwen3-235B-A22B; DeepSeek-R1-0528; Qwen3-235B-A22B-Instruct-2507 | 
| Synthetic OpenMathReasoning from gpt-oss-120b and Qwen2.5-32B-Instruct | Text | Undisclosed | OpenMathReasoning | gpt-oss-120b; Qwen2.5-32B-Instruct | 
| Synthetic Offline Search MCQA HLE from DeepSeek-R1-0528 | Text | Undisclosed | - | DeepSeek-R1-0528 | 
| Synthetic Offline Search MCQA GPQA from Qwen3-235B-A22B and DeepSeek-R1-0528 | Text | Undisclosed | - | Qwen3-235B-A22B; DeepSeek-R1-0528 | 
| Synthetic Human Preference from QwQ-32B, Qwen3-30B-A3B, Qwen3-235B-A22B, Qwen3-235B-A22B-Instruct-2507, Mistral-Small-3.1-24B-Instruct-2503, Mistral-Small-3.2-24B-Instruct-2506, MiniMax-M1-80k, MiniMax-M1-40k, Kimi-K2-Instruct, DeepSeek-V3-0324, DeepSeek-R1-0528 | Text | Undisclosed | - | QwQ-32B; Qwen3-30B-A3B; Qwen3-235B-A22B; Qwen3-235B-A22B-Instruct-2507; Mistral-Small-3.1-24B-Instruct-2503; Mistral-Small-3.2-24B-Instruct-2506; MiniMax-M1-80k; MiniMax-M1-40k; Kimi-K2-Instruct; DeepSeek-V3-0324; DeepSeek-R1-0528 | 
| Synthetic WildChat-1M and arena-human-preference-140k from DeepSeek-R1, gemma-2-2b-it, gemma-3-27b-it, gpt-oss-20b, gpt-oss-120b, Mistral-7B-Instruct-v0.3, Mixtral-8x22B-Instruct-v0.1, Nemotron-4-340B-Instruct, NVIDIA-Nemotron-Nano-9B-v2, Phi-4-mini-instruct, Phi-3-small-8k-instruct, Phi-3-medium-4k-instruct, Qwen3-235B-A22B, QwQ-32B | Text | Undisclosed | WildChat-1M; arena-human-preference-140k | DeepSeek-R1; gemma-2-2b-it; gemma-3-27b-it; gpt-oss-20b; gpt-oss-120b; Mistral-7B-Instruct-v0.3; Mixtral-8x22B-Instruct-v0.1; Nemotron-4-340B-Instruct; NVIDIA-Nemotron-Nano-9B-v2; Phi-4-mini-instruct; Phi-3-small-8k-instruct; Phi-3-medium-4k-instruct; Qwen3-235B-A22B; QwQ-32B | 
| Synthetic Safety from DeepSeek-R1-0528, gpt-oss-120b, DeepSeek-R1-Distill-Qwen-7B, and Mixtral-8x7B-v0.1 | Text | Undisclosed | Nemotron Content Safety Dataset V2; Gretel Synthetic Safety Alignment Dataset; RedTeam-2K; Malicious Tasks; | DeepSeek-R1-0528; gpt-oss-120b; DeepSeek-R1-Distill-Qwen-7B; Qwen3-30B-A3B-Thinking-2507; Qwen3-235B-A22B-Instruct-2507; Mixtral-8x7B-v0.1 | 
| Synthetic Code from Qwen3-32B | Text | Undisclosed | English Common Crawl; English Common Crawl 1.1 | Qwen3-32B | 
| Synthetic OpenCodeReasoning from DeepSeek-R1 | Text | Undisclosed | OpenCodeReasoning | DeepSeek-R1 | 
| Synthetic LIMO from DeepSeek-R1-0528 | Text | Undisclosed | LIMO | DeepSeek-R1-0528 | 
| Synthetic SCP from DeepSeek-R1-0528 | Text | Undisclosed | SCP-116K | DeepSeek-R1-0528 | 
| Synthetic Stack Exchange from DeepSeek-R1-0528 | Text | Undisclosed | Stack Exchange | DeepSeek-R1-0528 | 
| Synthetic Common Crawl from Qwen3-30B-A3B | Text | Undisclosed | Common Crawl | Qwen3-30B-A3B | 
| Synthetic Wikipedia from Qwen3-30B-A3B | Text | Undisclosed | Wikimedia | Qwen3-30B-A3B | 
| Synthetic Essential-Web from Qwen3-30B-A3B and Qwen3-235B-A22B-Thinking-2507 | Text | Undisclosed | Essential-Web | Qwen3-30B-A3B; Qwen3-235B-A22B-Thinking-2507 | 
| Synthetic Textbook Math from Qwen3-30B-A3B, Qwen3-235B-A22B, phi-4 | Text | Undisclosed | Common Crawl; FineMath | Qwen3-30B-A3B; Qwen3-235B-A22B; phi-4 | 
| Synthetic Math and Code from DeepSeek-R1 and DeepSeek-R1-0528 | Text | Undisclosed | Magicoder-Evol-Instruct-110K; opc-sft-stage2; TACO; OpenCodeReasoning; OpenMathReasoning; NuminaMath CoT | DeepSeek-R1; DeepSeek-R1-0528 | 
| Synthetic Nemotron-Personas-USA from gpt-oss-120b and Qwen3-8B | Text | Undisclosed | Nemotron-Personas-USA | gpt-oss-120b; Qwen3-8B | 

| Dataset | # of Tokens in Nemotron Nano 2 | # of Tokens in Nemotron 3 Nano | 
|---|---|---|
| English Common Crawl | 3,360,110,334,818 | 3,456,523,212,210 | 
| English Synthetic CC | 1,949,464,641,123 | 4,340,740,677,920 | 
| Crawl++ | 360,389,153,262 | 360,389,153,262 | 
| Math | 124,606,230,663 | 154,217,502,165 | 
| Synthetic Math | 73,007,767,155 | 73,007,767,155 | 
| Code | 747,409,228,724 | 1,043,856,922,136 | 
| Synthetic Code | 175,067,553,293 | 453,117,917,176 | 
| Common Crawl Code | 0 | 263,072,374,097 | 
| English Wiki | 17,349,266,926 | 17,349,266,926 | 
| Synthetic Wiki | 0 | 7,850,648,552 | 
| Books | 0 | 0 | 
| Papers | 191,586,493,365 | 191,586,493,365 | 
| PDF-to-text | 141,096,578,533 | 141,096,578,533 | 
| Code SFT | 60,025,726,817 | 102,863,752,325 | 
| STEM SFT | 272,680,426,295 | 359,826,214,274 | 
| General SFT | 6,057,478,645 | 6,057,478,645 | 
| Tool-Calling SFT | 0 | 26,244,716,867 | 
| Multilingual | 2,172,261,909,350 | 1,743,892,490,859 | 
| Synthetic multilingual | 997,710,364,950 | 595,140,661,135 | 
| **Total** | **10,648,823,153,919** | **13,336,833,827,602** | 

We use a considerable amount of synthetic data. Out of 10.6 trillion tokens, 3,534,013,958,278 tokens are synthetically generated.

We extracted data for fifteen languages from the following three Common Crawl snapshots: CC-MAIN-2024-51, CC-MAIN-2025-08, CC-MAIN-2025-18. The fifteen languages included were Arabic, Chinese, Danish, Dutch, French, German, Italian, Japanese, Korean, Polish, Portuguese, Russian, Spanish, Swedish, and Thai. As we did not have reliable multilingual model-based quality classifiers available, we applied just heuristic filtering instead—similar to what we did for lower quality English data in the Nemotron-CC pipeline, but selectively removing some filters for some languages that did not work well. Deduplication was done in the same way as for Nemotron-CC. Additionally, we used data from Wikipedia and FineWeb-2 (Penedo et al., 2025) for these fifteen languages as well as four additional languages: Czech, Finnish, Hebrew, and Hindi.

| Language | Total Tokens | 
|---|---|
| Arabic | 118,056,362,726 | 
| Danish | 117,747,321,618 | 
| German | 146,613,691,781 | 
| Spanish | 469,156,575,409 | 
| French | 139,982,002,289 | 
| Italian | 298,858,370,174 | 
| Japanese | 682,755,693,336 | 
| Korean | 127,099,747,538 | 
| Dutch | 89,041,592,681 | 
| Polish | 105,356,493,147 | 
| Portuguese | 243,249,275,089 | 
| Russian | 185,314,014,057 | 
| Swedish | 74,954,953,299 | 
| Thai | 160,778,944,467 | 
| Chinese | 211,007,236,689 | 

We collect a total of 922,476,782,017 tokens of code in 43 different languages.


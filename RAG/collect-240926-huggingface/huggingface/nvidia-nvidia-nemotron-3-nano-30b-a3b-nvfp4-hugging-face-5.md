---
id: collect-240926-huggingface/huggingface/nvidia-nvidia-nemotron-3-nano-30b-a3b-nvfp4-hugging-face-5
title: "Load tokenizer and model"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "Mistral", "Moonshot", "Nvidia"]
dates: []
keywords: ["agent", "alignment", "benchmark", "deepseek", "kimi", "mistral", "nvidia", "qwen", "reasoning", "tool calling", "tool use", "training"]
source: docs/RAG/clean_en/huggingface/nvidia-nvidia-nemotron-3-nano-30b-a3b-nvfp4-hugging-face.md
source_anchor: ""
source_lines: [395, 445]
sha256: 8de4ca9975fc35660b809ec7b9ad360eba97f682804f68480feb3610c151a0d6
---

# Load tokenizer and model

| Dataset | Modality | Dataset Size | Seed Dataset | Model(s) used for generation | 
|---|---|---|---|---|
| Synthetic Art of Problem Solving from DeepSeek-R1 | Text | 40B | Art of Problem Solving; American Mathematics Competitions 8; American Mathematics Competitions 10; | DeepSeek-R1 | 
| Synthetic Moral Stories and Social Chemistry from Mixtral-8x22B-v0.1 | Text | 327M | social-chemestry-101; Moral Stories | Mixtral-8x22B-v0.1 | 
| Synthetic Social Sciences seeded with OpenStax from DeepSeek-V3, Mixtral-8x22B-v0.1, and Qwen2.5-72B | Text | 83.6M | OpenStax - CC BY-SA subset | DeepSeek-V3; Mixtral-8x22B-v0.1; Qwen2.5-72B | 
| Synthetic Health Sciences seeded with OpenStax from DeepSeek-V3, Mixtral-8x22B-v0.1, and Qwen2.5-72B | Text | 9.7M | OpenStax - CC BY-SA subset | DeepSeek-V3; Mixtral-8x22B-v0.1; Qwen2.5-72B | 
| Synthetic STEM seeded with OpenStax, Open Textbook Library, and GSM8K from DeepSeek-R1, DeepSeek-V3, DeepSeek-V3-0324, and Qwen2.5-72B | Text | 175M | OpenStax - CC BY-SA subset; GSM8K; Open Textbook Library - CC BY-SA & GNU subset | DeepSeek-R1, DeepSeek-V3; DeepSeek-V3-0324; Qwen2.5-72B | 
| Nemotron-PrismMath | Text | 4.6B | Big-Math-RL-Verified; OpenR1-Math-220k | Qwen2.5-0.5B-instruct, Qwen2.5-72B-Instruct; DeepSeek-R1-Distill-Qwen-32B | 
| Synthetic Question Answering Data from Papers and Permissible Books from Qwen2.5-72B-Instruct | Text | 350M |  | Qwen2.5-72B-Instruct | 
| Refreshed Nemotron-MIND from phi-4 | Text | 73B | Common Crawl | phi-4 | 
| Nemotron-CC-Math-4plus | Text | 52.3B | Common Crawl | phi-4 | 
| Nemotron-CC-Math-3 | Text | 80.9B | Common Crawl | phi-4 | 
| Synthetic AGIEval seeded with AQUA-RAT, LogiQA, and AR-LSAT from DeepSeek-V3 and DeepSeek-V3-0324 | Text | 4.0B | AQUA-RAT; LogiQA; AR-LSAT | DeepSeek-V3; DeepSeek-V3-0324 | 
| Synthetic AGIEval seeded with AQUA-RAT, LogiQA, and AR-LSAT from Qwen3-30B-A3B | Text | 4.2B | AQUA-RAT; LogiQA; AR-LSAT | Qwen3-30B-A3B | 
| Synthetic Art of Problem Solving from Qwen2.5-32B-Instruct, Qwen2.5-Math-72B, Qwen2.5-Math-7B, and Qwen2.5-72B-Instruct | Text |  | Art of Problem Solving; American Mathematics Competitions 8; American Mathematics Competitions 10; GSM8K; PRM800K | Qwen2.5-32B-Instruct; Qwen2.5-Math-72B; Qwen2.5-Math-7B; Qwen2.5-72B-Instruct | 
| Synthetic MMLU Auxiliary Train from DeepSeek-R1 | Text | 0.5B | MMLU Auxiliary Train | DeepSeek-R1 | 
| Synthetic Long Context Continued Post-Training Data from Papers and Permissible Books from Qwen2.5-72B-Instruct | Text |  |  | Qwen2.5-72B-Instruct | 
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
| Synthetic Art of Problem Solving and Stack Exchange from gpt-oss-120b, Qwen2.5-32B-Instruct, and Goedel-Prover-V2-32B | Text | Undisclosed | Art of Problem Solving; American Mathematics Competitions 8; American Mathematics Competitions 10; Stack Exchange | gpt-oss-120b; Qwen2.5-32B-Instruct; Goedel-Prover-V2-32B | 
| Synthetic Multilingual Science and Code data from DeepSeek-R1, DeepSeek-R1-0528, Qwen2.5-32B-Instruct, and Qwen3-235B-A22B, translated with Qwen2.5-32B-Instruct and Qwen2.5-14B-Instruct | Text | Undisclosed | Stack Exchange; SCP-116K; LIMO; TACO; Code Contest; Codeforces | DeepSeek-R1; DeepSeek-R1-0528; Qwen2.5-32B-Instruct; Qwen3-235B-A22B; | 
| Synthetic Safety from DeepSeek-R1-0528, gpt-oss-120b and Mixtral-8x7B-v0.1 | Text | Undisclosed | Nemotron Content Safety Dataset V2; Gretel Synthetic Safety Alignment Dataset; RedTeam-2K; Malicious Tasks; Nemotron-Personas-USA | DeepSeek-R1-0528; gpt-oss-120b; Mixtral-8x7B-v0.1 | 
| Synthetic STEM from Qwen3-235B-A22B-Instruct-2507 and gpt-oss-120b | Text | Undisclosed |  | Qwen3-235B-A22B-Instruct-2507; gpt-oss-120b | 
| Synthetic KernelBook from DeepSeek-R1-0528 | Text | Undisclosed | KernelBook | DeepSeek-R1-0528 | 
| Synthetic Tool Calling from Qwen3-235B-A22B-Thinking-2507 and Qwen3-Next-80B-A3B-Thinking | Text | Undisclosed | ToolBench; glaive-function-calling-v2; APIGen Function-Calling; Nemotron-Personas-USA | Qwen3-235B-A22B-Thinking-2507; Qwen3-Next-80B-A3B-Thinking | 
| Synthetic Chat from gpt-oss-120b, Mixtral-8x22B-Instruct-v0.1, Qwen3-235B-A22B-Instruct-2507 , and Qwen3-235B-A22B-Thinking-2507 | Text | Undisclosed | C4; LMSYS-Chat-1M; ShareGPT; GSM8K; PRM800K; FinQA; WikiTableQuestions; Riddles; glaive-function-calling-v2; SciBench; tigerbot-kaggle-leetcodesolutions-en-2k; OpenBookQA; Advanced Reasoning Benchmark; Software Heritage; Khan Academy Math Keywords; WildChat-1M; Nemotron-Personas-USA | gpt-oss-120b; Mixtral-8x22B-Instruct-v0.1; Qwen3-235B-A22B-Instruct-2507; Qwen3-235B-A22B-Thinking-2507 | 
| Synthetic Long Context from Qwen3-235B-A22B-Instruct-2507 | Text | Undisclosed | CORE; PG-19; DOAB CC BY & CC BY-SA subset; NDLTD | Qwen3-235B-A22B-Instruct-2507 | 
| Synthetic Tool Use Interactive Agent from gpt-oss-120b, DeepSeek-R1-0528, Qwen3-32B, and Qwen3-235B-A22B-Thinking-2507 | Text | Undisclosed | NVIDIA Internal | gpt-oss-120b; DeepSeek-R1-0528; Qwen3-32B; and Qwen3-235B-A22B-Thinking-2507 | 
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

---
id: collect-240926-huggingface/huggingface/unsloth-nvidia-nemotron-3-5-lightning-30b-a3b-gguf-hugging-face-5
title: "unsloth-nvidia-nemotron-3-5-lightning-30b-a3b-gguf-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "California", "DeepSeek", "MiniMax", "Mistral", "Moonshot"]
dates: []
keywords: ["benchmark", "chatgpt", "deepseek", "kimi", "mistral", "pretraining", "qwen", "reasoning", "regulation", "tool calling", "training"]
source: docs/RAG/clean_en/huggingface/unsloth-nvidia-nemotron-3-5-lightning-30b-a3b-gguf-hugging-face.md
source_anchor: ""
source_lines: [370, 412]
sha256: 428d67be15ef4f6daf0aebd66ca13933aaaf6d0b1852aed5bbecde0e3599aeec
---

# unsloth-nvidia-nemotron-3-5-lightning-30b-a3b-gguf-hugging-face

| Dataset | Modality | Dataset Size | Seed Dataset | Model(s) used for generation | 
|---|---|---|---|---|
| Nemotron-Pretraining-Fact-Seeking | Text | 35.0B | FineWiki | Qwen3-30B-A3B-Instruct-2507 | 
| Nemotron-Pretraining-Legal | Text | 4.3B | CommonPile (caselaw_access_project_filtered); California Code of Regulations; Judicial Ethics Opinions; GLOBALCIT; CUAD; Nemotron Personas; ToSDR Terms of Service Corpus; CodeHima/TOS_Dataset; ContractNLI; CaseHOLD; Code of Federal Regulations; Canadian Case Law (subsets that allow commercial use) | Qwen3-235B-A22B-Thinking-2507 | 
| Nemotron-Pretraining-Formal-Logic | Text | 128M | Nemotron Personas | Qwen3-235B-A22B-Thinking-2507 | 
| Nemotron-Pretraining-Economics | Text | 73.4M | - | Qwen3-235B-A22B-Thinking-2507 | 
| Nemotron-Pretraining-Multiple-Choice | Text | 1.6B | MMLU Auxiliary Train | DeepSeek-V3; Qwen3-235B-A22B | 
| Nemotron-Pretraining-Code-Concepts | Text | 7.3B | - | gpt-oss-20b; gpt-oss-120b | 
| Nemotron-Pretraining-Unconditional-Algorithmic | Text | 196.5M | - | gpt-oss-120b; Qwen3-235B-A22B | 
| More Synthetic Tasks from DeepSeek-V3 and Qwen3-235B-A22B | Text | 1.1B | train splits of acp_bench; ai2_arc; babi; gsm8k; hendrycks_math; IFEval; MedText; mediqa_qa; mlqa; MMLU-Pro; mmlu-pro-plus; MMLU-ProX; nq_open; tinyGSM8k; truthful_qa; truthfulqa-multi; MATH-lighteval; mmlu; awesome-chatgpt-prompts; super_glue | DeepSeek v3; Qwen3-235B-A22B | 
| Synthetic Tasks from DeepSeek-V3 and Qwen3-235B-A22B | Text | 6.7B | train splits of Into the Unknown; AI2 ARC (AI2 Reasoning Challenge); BLiMP (Benchmark of Linguistic Minimal Pairs); CommonSenseQA; GLUE; HeadQA; Hendrycks Ethics; Memo Trap; modus-tollens; NeQA; pattern-matching-suppression; mastermind_24_mcq_random; mastermind_24_mcq_close; quote-repetition; redefine-math; Repetitive Algebra; sig-figs; MMLU-Pro; MC-TACO; MedConceptsQA; MMLU_dataset; OpenbooksQA; PIQA (Physical Interaction Question Answering); SocialIQA; SuperGLUE; tinyAI2_arc; tinyMMLU; tinyWinogrande; TruthfulQA; WebQuestions; Winogrande; GPQA; MBPP | DeepSeek v3; Qwen3-235B-A22B | 
| Synthetic Art of Problem Solving from DeepSeek-R1 | Text | 40B | Art of Problem Solving; American Mathematics Competitions 8; American Mathematics Competitions 10 | DeepSeek-R1 | 
| Synthetic Moral Stories and Social Chemistry from Qwen3-235B-A22B-Thinking-2507 and Mixtral-8x22B-v0.1 | Text | 15.2M | social-chemestry-101; Moral Stories | Qwen3-235B-A22B-Thinking-2507; Mixtral-8x22B-v0.1 | 
| Synthetic Moral Stories and Social Chemistry from Mixtral-8x22B-v0.1 | Text | 327M | social-chemestry-101; Moral Stories | Mixtral-8x22B-v0.1 | 
| Synthetic Social Sciences seeded with OpenStax from DeepSeek-V3, Mixtral-8x22B-v0.1, and Qwen2.5-72B | Text | 83.6M | OpenStax - CC BY-SA subset | DeepSeek-V3; Mixtral-8x22B-v0.1; Qwen2.5-72B | 
| Synthetic Health Sciences seeded with OpenStax from DeepSeek-V3, Mixtral-8x22B-v0.1, and Qwen2.5-72B | Text | 9.7M | OpenStax - CC BY-SA subset | DeepSeek-V3; Mixtral-8x22B-v0.1; Qwen2.5-72B | 
| Synthetic STEM seeded with OpenStax, Open Textbook Library, and GSM8K from DeepSeek-R1, DeepSeek-V3, DeepSeek-V3-0324, and Qwen2.5-72B | Text | 175M | OpenStax - CC BY-SA subset; GSM8K; Open Textbook Library - CC BY-SA & GNU subset | DeepSeek-R1, DeepSeek-V3; DeepSeek-V3-0324; Qwen2.5-72B | 
| Nemotron-PrismMath | Text | 4.6B | Big-Math-RL-Verified; OpenR1-Math-220k | Qwen2.5-0.5B-instruct, Qwen2.5-72B-Instruct; DeepSeek-R1-Distill-Qwen-32B | 
| Synthetic Question Answering Data from Papers and Permissible Books from Qwen2.5-72B-Instruct | Text | 350M | arXiv; National Institutes of Health ExPorter; BioRxiv; PMC Article; USPTO Backgrounds; peS2o; Global Regulation; CORE; PG-19; DOAB CC BY & CC BY-SA subset; NDLTD | Qwen2.5-72B-Instruct | 
| Synthetic Rephrased Math Data from Common Crawl from phi-4 | Text | 73B | Common Crawl | phi-4 | 
| Synthetic Math Data from Common Crawl 4plus | Text | 52.3B | Common Crawl | phi-4 | 
| Synthetic Math Data from Common Crawl 3 | Text | 80.9B | Common Crawl | phi-4 | 
| Synthetic AGIEval seeded with AQUA-RAT, LogiQA, and AR-LSAT from DeepSeek-V3 and DeepSeek-V3-0324 | Text | 4.0B | AQUA-RAT; LogiQA; AR-LSAT | DeepSeek-V3; DeepSeek-V3-0324 | 
| Synthetic AGIEval seeded with AQUA-RAT, LogiQA, and AR-LSAT from Qwen3-30B-A3B | Text | 4.2B | AQUA-RAT; LogiQA; AR-LSAT | Qwen3-30B-A3B | 
| Synthetic Art of Problem Solving from Qwen2.5-32B-Instruct, Qwen2.5-Math-72B, Qwen2.5-Math-7B, and Qwen2.5-72B-Instruct | Text | Undisclosed | Art of Problem Solving; American Mathematics Competitions 8; American Mathematics Competitions 10; GSM8K; PRM800K | Qwen2.5-32B-Instruct; Qwen2.5-Math-72B; Qwen2.5-Math-7B; Qwen2.5-72B-Instruct | 
| Synthetic MMLU Auxiliary Train from DeepSeek-R1 | Text | 0.5B | MMLU Auxiliary Train | DeepSeek-R1 | 
| Synthetic Long Context Continued Post-Training Data from Papers and Permissible Books from Qwen2.5-72B-Instruct | Text | Undisclosed | arXiv; National Institutes of Health ExPorter; BioRxiv; PMC Article; USPTO Backgrounds; peS2o; Global Regulation; CORE; PG-19; DOAB CC BY & CC BY-SA subset; NDLTD | Qwen2.5-72B-Instruct | 
| Synthetic Common Crawl from Qwen3-30B-A3B and Mistral-Nemo-12B-Instruct | Text | 415.8B | Common Crawl | Qwen3-30B-A3B; Mistral-NeMo-12B-Instruct | 
| Synthetic Multilingual Data from Common Crawl from Qwen3-30B-A3B | Text | Undisclosed | Common Crawl | Qwen3-30B-A3B | 
| Synthetic Multilingual Data from Wikimedia from Qwen3-30B-A3B | Text | Undisclosed | Wikimedia | Qwen3-30B-A3B | 
| Synthetic Math Data from Wikimedia from Nemotron-4-340B-Instruct | Text | Undisclosed | - | Nemotron-4-340B-Instruct | 
| Synthetic Common Crawl Code from phi-4 | Text | 427.9B | Common Crawl | phi-4 | 
| Synthetic Scientific Coding from Qwen3-235B-A22B | Text | 1.2B | Wikimedia | Qwen3-235B-A22B | 
| Tool Calling Data | Text | 26.2B | - | Qwen3-235B-A22B-2507; gpt-oss-120b | 
| Synthetic Essential-Web from QwQ-32B | Text | 28.1B | Essential-Web | QwQ-32B | 
| Translated Synthetic Crawl | Text | 389.9B | Common Crawl | Qwen3-30B-A3B | 
| Translated Synthetic Wikipedia | Text | 7.9B | Wikimedia | Qwen3-30B-A3B | 
| Synthetic Long Context from Qwen3-235B-A22B-Instruct-2507 | Text | Undisclosed | CORE; PG-19; DOAB CC BY & CC BY-SA subset; NDLTD | Qwen3-235B-A22B-Instruct-2507 | 
| Synthetic Search STEM OPENQ from DeepSeek-R1-0528 | Text | Undisclosed | - | DeepSeek-R1-0528 | 
| Synthetic MCQ from Qwen2.5-32B-Instruct and DeepSeek-R1-0528 | Text | Undisclosed | - | Qwen2.5-32B-Instruct; DeepSeek-R1-0528 | 
| Synthetic Offline Search MCQA HLE from DeepSeek-R1-0528 | Text | Undisclosed | - | DeepSeek-R1-0528 | 
| Synthetic Offline Search MCQA GPQA from Qwen3-235B-A22B and DeepSeek-R1-0528 | Text | Undisclosed | - | Qwen3-235B-A22B; DeepSeek-R1-0528 | 
| Synthetic Human Preference from QwQ-32B, Qwen3-30B-A3B, Qwen3-235B-A22B, Qwen3-235B-A22B-Instruct-2507, Mistral-Small-3.1-24B-Instruct-2503, Mistral-Small-3.2-24B-Instruct-2506, MiniMax-M1-80k, MiniMax-M1-40k, Kimi-K2-Instruct, DeepSeek-V3-0324, DeepSeek-R1-0528 | Text | Undisclosed | - | QwQ-32B; Qwen3-30B-A3B; Qwen3-235B-A22B; Qwen3-235B-A22B-Instruct-2507; Mistral-Small-3.1-24B-Instruct-2503; Mistral-Small-3.2-24B-Instruct-2506; MiniMax-M1-80k; MiniMax-M1-40k; Kimi-K2-Instruct; DeepSeek-V3-0324; DeepSeek-R1-0528 | 

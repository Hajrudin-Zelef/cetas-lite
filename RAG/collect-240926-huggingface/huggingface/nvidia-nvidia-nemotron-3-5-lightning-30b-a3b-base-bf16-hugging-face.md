---
id: collect-240926-huggingface/huggingface/nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-base-bf16-hugging-face
title: "nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-base-bf16-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["AWS", "Alibaba", "California", "China", "DeepSeek", "Google", "Hugging Face", "MiniMax", "Mistral", "Moonshot", "Nvidia", "OpenAI"]
dates: ["2025-01-05", "2025-02-10", "2025-08-04", "2025-09", "2025-12", "2025-29-04", "2025-30-09", "2026-05", "2026-11-08"]
keywords: ["nvidia", "agents", "alignment", "attention", "benchmark", "benchmarks", "blackwell", "chatgpt", "compute", "deepseek", "distillation", "fine-tuning"]
source: docs/RAG/clean_en/huggingface/nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-base-bf16-hugging-face.md
source_anchor: ""
source_lines: [1, 271]
sha256: af052d85a01d9b76e7fcb81f076cd75d461eec7cc383f2604469425b258ee200
---

# nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-base-bf16-hugging-face

<!-- source: https://huggingface.co/nvidia/NVIDIA-Nemotron-3.5-Lightning-30B-A3B-Base-BF16 -->

*Looking for the post-trained model? See NVIDIA-Nemotron-3.5-Lightning-30B-A3B-BF16 for the full-precision release, or NVIDIA-Nemotron-3.5-Lightning-30B-A3B-NVFP4 for optimized inference.*


**Model Developer:** NVIDIA Corporation

**Model Dates:** December 2025 - May 2026

**Data Freshness:**

- The pre-training data has a cutoff date of September 2025.

**NVIDIA-Nemotron-3.5-Lightning-30B-A3B-Base-BF16** is a large language model (LLM) trained by NVIDIA. This is the base (pre-trained) checkpoint of the Nemotron 3.5 Lightning family — no supervised fine-tuning, reinforcement learning, or distillation has been applied — making it the natural starting point for developers and researchers building their own post-trained models.

The model employs a hybrid **Mixture-of-Experts (MoE)** architecture, utilizing interleaved Mamba-2 and MoE layers, along with select Attention layers. The model incorporates **Multi-Token Prediction (MTP)** layers — trained via a dedicated continued pre-training phase — for richer training signals and native speculative decoding, and it is pre-trained using an **NVFP4** recipe to maximize compute efficiency. The model has **3B active parameters** and **30B parameters in total**.

The pre-training corpus spans English, 19 other spoken languages, and 43 programming languages.

This model is ready for commercial and non-commercial use.

NVIDIA Nemotron™ is a family of open models with open weights, training data, and recipes, delivering leading efficiency and accuracy for building specialized AI agents.

Use of this model is governed by the OpenMDW License Agreement, version 1.1 (OpenMDW-1.1).

| Benchmark | Qwen3.5-35B-A3B | Gemma-4-26B-A4B | Nemotron-3 Nano 30B-A3B | Nemotron-3.5 Lightning | Nemotron-3 Super 120B-A12B | 
|---|---|---|---|---|---|
| **General** |  |  |  |  |  | 
| MMLU | 81.07 | 77.81 | 78.48 | 78.59 | 86.01 | 
| MMLU-Pro (5-shot) | 64.49 | 50.02 | 64.20 | 67.94 | 74.43 | 
| AGIEval-EN (CoT) | 70.51 | 55.28 | 68.45 | 70.02 | 77.92 | 
| **Math** |  |  |  |  |  | 
| GSM8K (8-shot, CoT) | 90.07 | 77.03 | 91.43 | 91.28 | 90.67 | 
| Minerva Math (4-shot) | 59.66 | 43.74 | 82.64 | 82.78 | 84.84 | 
| **Code** |  |  |  |  |  | 
| MBPP (3-shot) | 70.76 | 68.39 | 73.82 | 78.59 | 81.71 | 
| HumanEval | 66.46 | 50.00 | 75.00 | 77.44 | 80.49 | 
| **Commonsense understanding** |  |  |  |  |  | 
| ARC-Challenge (25-shot) | 95.39 | 92.83 | 91.98 | 92.66 | 96.08 | 
| HellaSwag | 85.61 | 85.26 | 85.55 | 85.55 | 88.97 | 
| OpenBookQA | 44.20 | 48.80 | 46.80 | 47.60 | 48.60 | 
| PIQA (acc) | 82.32 | 82.21 | 82.64 | 83.35 | 83.90 | 
| PIQA (acc-norm) | 82.54 | 83.84 | 84.33 | 85.20 | 85.47 | 
| WinoGrande (5-shot) | 79.24 | 79.08 | 79.16 | 79.95 | 78.93 | 
| **Global-MMLU-Lite (5-shot)** |  |  |  |  |  | 
| Average | 80.94 | 74.78 | 74.62 | 75.53 | 85.72 | 
| German (de) | 81.25 | 75.50 | 75.75 | 76.00 | 87.25 | 
| Spanish (es) | 83.25 | 76.50 | 79.25 | 78.00 | 87.25 | 
| French (fr) | 82.00 | 74.50 | 74.75 | 76.25 | 85.75 | 
| Italian (it) | 85.25 | 76.50 | 77.00 | 77.75 | 86.75 | 
| Japanese (ja) | 77.25 | 73.75 | 70.75 | 73.25 | 84.25 | 
| Korean (ko) | 78.25 | 73.50 | 70.50 | 71.25 | 82.50 | 
| Portuguese (pt) | 82.25 | 76.25 | 75.00 | 77.00 | 87.50 | 
| Chinese (zh) | 78.00 | 71.75 | 74.00 | 74.75 | 84.50 | 
| **Multilingual Math — MGSM (8-shot)** |  |  |  |  |  | 
| German (de) | 84.80 | 69.60 | 85.60 | 84.80 | 90.40 | 
| Spanish (es) | 89.20 | 78.80 | 84.40 | 88.00 | 88.00 | 
| French (fr) | 84.40 | 67.20 | 81.60 | 82.40 | 85.60 | 
| Japanese (ja) | 72.80 | 54.80 | 70.40 | 69.60 | 81.60 | 
| Russian (ru) | 90.00 | 76.40 | 86.40 | 87.60 | 91.20 | 
| Chinese (zh) | 85.60 | 68.00 | 82.80 | 78.40 | 85.60 | 
| **Long context — RULER** |  |  |  |  |  | 
| RULER 256K | 82.36 | 85.73 | 71.71 | 76.88 | 83.03 | 
| RULER 1M | 56.43 | 72.93 | 51.23 | 69.62 | 66.98 | 

Accuracy numbers measured by NVIDIA under a consistent harness (NeMo Gym / Nemo Evaluator SDK); they may differ from vendors' self-reported numbers.

All evaluation results were collected via Nemo Evaluator SDK and NVIDIA's open source container of LM Evaluation Harness except for RULER which uses the NeMo Skills harness. The open source container on LM Evaluation Harness packaged via NVIDIA's Nemo Evaluator SDK used for evaluations can be found here. OSS Evaluation Recipes for Nemotron-3.5 Lightning are available in NeMo Gym.

This model is intended for developers and researchers building LLMs. As the base checkpoint of the Nemotron 3.5 Lightning family, it is the recommended starting point for pre-training research, continued pre-training on domain corpora, and building custom post-trained variants (SFT, RL, and distillation) via NeMo RL, NeMo Gym, and Megatron-LM.

Hugging Face - 08/11/2026

- **Architecture Type:** Mamba2-Transformer Hybrid Mixture of Experts (MoE) with Multi-Token Prediction (MTP)
- **Network Architecture:** Nemotron Hybrid MoE
- **Number of model parameters:** 30B Total / 3B Active

The model was pre-trained with over 20T tokens and supports up to 1M context length. The pre-training phase used an NVFP4 recipe. The model includes **Multi-Token Prediction (MTP)** layers, which predict multiple future tokens to provide richer training signals and enable faster inference via speculative decoding.

Stage 1: Pre-Training

- NVIDIA-Nemotron-3.5-Lightning-30B-A3B-Base-BF16 model was pre-trained using an NVFP4 recipe with crawled and synthetic code, math, science, and general knowledge data.
- Software used for pre-training: Megatron-LM

Stage 2: Continued Pre-Training for Multi-Token Prediction (MTP)

- The model underwent a continued pre-training phase to train its **Multi-Token Prediction (MTP)** layers. In this stage, MTP heads learn to predict multiple future tokens, providing richer training signals to the base model and enabling native speculative decoding at inference time.

NVIDIA-Nemotron-3.5-Lightning-30B-A3B-Base-BF16 model is a result of the above work.

- **Input Type(s):** Text
- **Input Format(s):** String
- **Input Parameters:** One-Dimensional (1D): Sequences
- **Other Properties Related to Input:** Maximum context length up to 1M tokens. The pre-training corpus spans English, 19 other spoken languages, and 43 programming languages.

- **Output Type(s):** Text
- **Output Format:** String
- **Output Parameters:** One-Dimensional (1D): Sequences
- **Other Properties Related to Output:** Maximum context length up to 1M tokens

Our AI models are designed and optimized to run on NVIDIA GPU-accelerated systems. By leveraging NVIDIA's hardware (e.g. GPU cores) and software frameworks (e.g., CUDA libraries), the model achieves faster training and inference times compared to CPU-only solutions.

- **Runtime Engine(s):** PyTorch; Megatron-LM; NeMo
- **Supported Hardware Microarchitecture Compatibility:** NVIDIA Ampere - A100; NVIDIA Blackwell; NVIDIA Hopper - H100-80GB
- **Operating System(s):** Linux

The integration of foundation and fine-tuned models into AI systems requires additional testing using use-case-specific data to ensure safe and effective deployment. Following the V-model methodology, iterative testing and validation at both unit and system levels are essential to mitigate risks, meet technical and functional requirements, and ensure compliance with safety and ethical standards before deployment.

- v1.0 - GA (08/11/2026)

**Data Modality:** Text
**Training Data Size:** More than 20 Trillion Tokens
**Dataset partition:** *Training [100%], testing [0%], validation [0%]*
**Time period for training data collection:** 2013 to December 2025
**Time period for testing data collection:** 2013 to December 2025
**Time period for validation data collection:** 2013 to December 2025
**Data Collection Method by dataset:** Hybrid: Automated, Manually-Collected, Synthetic
**Labeling Method by dataset:** Hybrid: Automated, Manually-Labeled, Synthetic

**Properties:** NVIDIA-Nemotron-3.5-Lightning-30B-A3B-Base-BF16 is pre-trained on a large corpus of high-quality curated and synthetically-generated data. It is trained in the English language, as well as 19 other spoken languages and 43 programming languages. Our sources cover a variety of document types such as: webpages, dialogue, articles, and other written materials. The corpus spans domains including legal, math, science, finance, and more. We also include a small portion of question-answering, and alignment style data to improve model accuracy. The model was pre-trained for more than 20 trillion tokens.

## For Detailed Dataset Information: Click here!

The foundation of the model is trained on the Nemotron 3 corpus, comprising the following datasets from the Nemotron Pretraining Datasets collection:

| Dataset Collection | Token Counts | Description | 
|---|---|---|
| **Nemotron-CC-v2** &**v2.1** | 9.1T | A massive collection of English web data filtered from Common Crawl, including 2.5T+ tokens of new organic, translated, and synthetically rephrased content. | 
| **Nemotron-CC-Code-v1** | 427.9B | High-quality code tokens extracted from Common Crawl using the Lynx + LLM pipeline to preserve structure and equations. | 
| **Nemotron-Pretraining-Code-v1** &**v2** &**v3** | 1.7T | Curated GitHub code references with multi-stage filtering, deduplication, and large-scale synthetic code data. | 
| **Nemotron-CC-Math-v1** | 133.3B | High-quality math pre-training dataset preserving LaTeX formatting and mathematical structures. | 
| **Nemotron-Pretraining-Specialized-v1** &**v1.1** &**v1.2** &**Nemotron-Pretraining-SFT-v1** | 660.0B | Synthetic datasets targeting specialized domains such as STEM reasoning and scientific coding. | 
| **Nemotron-Pretraining-Legal-v1** | 4.3B | Synthetic datasets targeting the legal domain. | 

The English Common Crawl data was downloaded from the Common Crawl Foundation (see their FAQ for details on their crawling) and includes the snapshots CC-MAIN-2013-20 through CC-MAIN-2025-13. The data was subsequently deduplicated and filtered in various ways described in the Nemotron-CC paper. Additionally, we extracted data for fifteen languages from the following three Common Crawl snapshots: CC-MAIN-2024-51, CC-MAIN-2025-08, CC-MAIN-2025-18. The fifteen languages included were Arabic, Chinese, Danish, Dutch, French, German, Italian, Japanese, Korean, Polish, Portuguese, Russian, Spanish, Swedish, and Thai. As we did not have reliable multilingual model-based quality classifiers available, we applied just heuristic filtering instead—similar to what we did for lower quality English data in the Nemotron-CC pipeline, but selectively removing some filters for some languages that did not work well. Deduplication was done in the same way as for Nemotron-CC.

The GitHub Crawl was collected using the GitHub REST API and the Amazon S3 API. Each crawl was operated in accordance with the rate limits set by its respective source, either GitHub or S3. We collect raw source code and subsequently remove any having a license which does not exist in our permissive-license set.

| Dataset | Modality | Dataset Size | Collection Period | Collecting Organisation | 
|---|---|---|---|---|
| English Common Crawl | Text | 3.36T | 4/8/2025 | NVIDIA Advanced Deep Learning Research | 
| English Common Crawl 1.1 | Text | Not disclosed | 10/2/2025 | NVIDIA Advanced Deep Learning Research | 
| Multilingual Common Crawl | Text | 812.7B | 5/1/2025 | NVIDIA Advanced Deep Learning Research | 
| GitHub Crawl | Text | 747.4B | 4/29/2025 | NVIDIA Advanced Deep Learning Research | 
| GitHub Crawl 1.1 | Text | 172.7B | 9/30/2025 | NVIDIA Advanced Deep Learning Research | 

| Dataset | Model(s) used | 
|---|---|
| Global Regulation | Unknown | 
| TAUS Translation Memory | Unknown | 
| Scale HLE | Unknown | 
| HackerRank Coding | Unknown | 
| RL data for Search | Gemini 3; GPT-5 | 

| Dataset | Model(s) used | 
|---|---|
| Simple Minesweeper | Undisclosed | 
| Simple Sudoku | Undisclosed | 
| Multitool Typewriter Hard | Undisclosed | 
| Machine Translation of News Commentary and TAUS Translation Memory | Undisclosed | 
| Machine Translation of STEM - | Qwen2.5-14B-Instruct | 
| Competitive Coding RL data from Nemotron Cascade | Undisclosed | 
| Long context RL | Undisclosed | 
| Single-step SWE RL for patch generation | Undisclosed | 
| OpenHands SWE | Undisclosed | 

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
| Synthetic WildChat-1M and arena-human-preference-140k from DeepSeek-R1, gemma-2-2b-it, gemma-3-27b-it, gpt-oss-20b, gpt-oss-120b, Mistral-7B-Instruct-v0.3, Mixtral-8x22B-Instruct-v0.1, Nemotron-4-340B-Instruct, NVIDIA-Nemotron-Nano-9B-v2, Phi-4-mini-instruct, Phi-3-small-8k-instruct, Phi-3-medium-4k-instruct, Qwen3-235B-A22B, QwQ-32B | Text | Undisclosed | WildChat-1M; arena-human-preference-140k | DeepSeek-R1; gemma-2-2b-it; gemma-3-27b-it; gpt-oss-20b; gpt-oss-120b; Mistral-7B-Instruct-v0.3; Mixtral-8x22B-Instruct-v0.1; Nemotron-4-340B-Instruct; NVIDIA-Nemotron-Nano-9B-v2; Phi-4-mini-instruct; Phi-3-small-8k-instruct; Phi-3-medium-4k-instruct; Qwen3-235B-A22B; QwQ-32B | 
| Synthetic Code from Qwen3-32B | Text | Undisclosed | English Common Crawl; English Common Crawl 1.1 | Qwen3-32B | 
| Synthetic OpenCodeReasoning from DeepSeek-R1 | Text | Undisclosed | OpenCodeReasoning | DeepSeek-R1 | 
| Synthetic OpenCodeReasoning from DeepSeek-R1-0528 | Text | Undisclosed | OpenCodeReasoning | DeepSeek-R1-0528 | 
| Synthetic HackerRank Coding from DeepSeek-R1-0528 | Text | Undisclosed | HackerRank Coding Dataset | DeepSeek-R1-0528 | 
| Synthetic LIMO from DeepSeek-R1-0528 | Text | Undisclosed | LIMO | DeepSeek-R1-0528 | 
| Synthetic SCP from DeepSeek-R1-0528 | Text | Undisclosed | SCP-116K | DeepSeek-R1-0528 | 
| Synthetic Stack Exchange from DeepSeek-R1-0528 | Text | Undisclosed | Stack Exchange | DeepSeek-R1-0528 | 
| Synthetic Stack Exchange from gpt-oss-120b and Qwen2.5-32B-Instruct | Text | Undisclosed | Stack Exchange | gpt-oss-120b; Qwen2.5-32B-Instruct | 
| Synthetic Stack Exchange from gpt-oss-120b | Text | Undisclosed | Stack Exchange | gpt-oss-120b | 
| Synthetic Art of Problem Solving from gpt-oss-120b and Qwen2.5-32B-Instruct | Text | Undisclosed | Art of Problem Solving; American Mathematics Competitions 8; American Mathematics Competitions 10 | gpt-oss-120b; Qwen2.5-32B-Instruct | 
| Synthetic Common Crawl from Qwen3-30B-A3B | Text | Undisclosed | Common Crawl | Qwen3-30B-A3B | 
| Synthetic Wikipedia from Qwen3-30B-A3B | Text | Undisclosed | Wikimedia | Qwen3-30B-A3B | 
| Synthetic Essential-Web from Qwen3-30B-A3B and Qwen3-235B-A22B-Thinking-2507 | Text | Undisclosed | Essential-Web | Qwen3-30B-A3B; Qwen3-235B-A22B-Thinking-2507 | 
| Synthetic Essential-Web from gpt-oss-120b | Text | Undisclosed | Essential-Web | gpt-oss-120b | 
| Synthetic Textbook Math from Qwen3-30B-A3B, Qwen3-235B-A22B, phi-4 | Text | Undisclosed | Common Crawl; FineMath | Qwen3-30B-A3B; Qwen3-235B-A22B; phi-4 | 
| Synthetic Math and Code from DeepSeek-R1 and DeepSeek-R1-0528 | Text | Undisclosed | Magicoder-Evol-Instruct-110K; opc-sft-stage2; TACO; OpenCodeReasoning; OpenMathReasoning; NuminaMath CoT | DeepSeek-R1; DeepSeek-R1-0528 | 
| Synthetic Math from gpt-oss-120b and Qwen2.5-32B-Instruct | Text | Undisclosed | - | gpt-oss-120b; Qwen2.5-32B-Instruct | 
| Synthetic OpenMathReasoning from gpt-oss-120b and Qwen2.5-32B-Instruct | Text | Undisclosed | OpenMathReasoning | gpt-oss-120b; Qwen2.5-32B-Instruct | 
| Synthetic KernelBook from DeepSeek-R1-0528 | Text | Undisclosed | KernelBook | DeepSeek-R1-0528 | 
| Synthetic Scale HLE from gpt-oss-120b | Text | Undisclosed | Scale HLE | gpt-oss-120b | 
| Synthetic CDQuestions from gpt-oss-120b | Text | Undisclosed | CDQuestions | gpt-oss-120b | 
| Synthetic GPQA from gpt-oss-120b and Qwen2.5-32B-Instruct | Text | Undisclosed | Stack Exchange | gpt-oss-120b; Qwen2.5-32B-Instruct | 
| Synthetic Vedantu from gpt-oss-120b | Text | Undisclosed | Vedantu | gpt-oss-120b | 
| Synthetic Search STEM MCQ from Qwen3-235B-A22B and DeepSeek-R1-0528 | Text | Undisclosed | - | Qwen3-235B-A22B; DeepSeek-R1-0528 | 
| Synthetic OpenSTEM from Qwen2.5-32B-Instruct and DeepSeek-R1-0528 | Text | Undisclosed | - | Qwen2.5-32B-Instruct; DeepSeek-R1-0528 | 
| Synthetic MCQ10 from DeepSeek-R1-0528 | Text | Undisclosed | - | DeepSeek-R1-0528 | 
| Synthetic MCQ4 from Qwen3-235B-A22B, DeepSeek-R1-0528, and Qwen3-235B-A22B-Instruct-2507 | Text | Undisclosed | - | Qwen3-235B-A22B; DeepSeek-R1-0528; Qwen3-235B-A22B-Instruct-2507 | 

**Data Collection Method by dataset** 

- Hybrid: Automated, Human, Synthetic
**Labeling Method by dataset**
- Hybrid: Automated, Human, Synthetic
**Properties:** This corpus comprises a mix of high-quality standard benchmarks for base language models, covering general knowledge, math, code, commonsense understanding, reading comprehension, multilingual understanding, and long context.

**Data Collection Method by dataset** 

- Hybrid: Automated, Human, Synthetic
**Labeling Method by dataset**
- Hybrid: Automated, Human, Synthetic
**Properties:** This corpus comprises a mix of high-quality standard benchmarks for base language models, covering general knowledge, math, code, commonsense understanding, reading comprehension, multilingual understanding, and long context.

- **Acceleration Engine:** PyTorch
- **Test Hardware:**
  - NVIDIA Hopper
    - 1-8x H100
    - 1-8x H200
  - NVIDIA Blackwell
    - GB200
- NVIDIA Hopper

NVIDIA believes Trustworthy AI is a shared responsibility and we have established policies and practices to enable development for a wide array of AI applications. Developers should work with their internal model team to ensure this model meets requirements for the relevant industry and use case and addresses unforeseen product misuse.

We advise against circumvention of any provided safety guardrails contained in the Model without a substantially similar guardrail appropriate for your use case. For more details: Safety and Explainability Subcards.

For more detailed information on ethical considerations for this model, please see the Model Card++ Bias, and Privacy Subcards.

Please report model quality, risk, security vulnerabilities or NVIDIA AI Concerns here.

- Downloads last month
- 188,031

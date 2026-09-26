---
id: collect-240926-huggingface/huggingface/nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-base-bf16-hugging-face-2
title: "nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-base-bf16-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["AWS", "Alibaba", "China", "Google", "Nvidia", "OpenAI"]
dates: ["2025-01-05", "2025-02-10", "2025-08-04", "2025-12", "2025-29-04", "2025-30-09"]
keywords: ["nvidia", "alignment", "gemini", "license", "memory", "pretraining", "reasoning", "regulation", "research", "training"]
source: docs/RAG/clean_en/huggingface/nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-base-bf16-hugging-face.md
source_anchor: ""
source_lines: [111, 166]
sha256: 4f6ede362ec2ef6216db84d1c82de253e2ab2f43d430b9fb329f460cdc4486f3
---

# nvidia-nvidia-nemotron-3-5-lightning-30b-a3b-base-bf16-hugging-face

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


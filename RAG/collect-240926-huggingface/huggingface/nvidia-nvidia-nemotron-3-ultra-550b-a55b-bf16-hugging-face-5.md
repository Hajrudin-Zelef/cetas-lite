---
id: collect-240926-huggingface/huggingface/nvidia-nvidia-nemotron-3-ultra-550b-a55b-bf16-hugging-face-5
title: "Set the IP for the head node in RAY_HEAD_IP"
domain: huggingface
role: reference
task: reference
actors: ["AWS", "Alibaba", "China", "Google", "Nvidia", "OpenAI"]
dates: ["2025-01-05", "2025-02-10", "2025-08-04", "2025-29-04", "2025-30-09"]
keywords: ["gemini", "license", "memory", "nvidia", "pretraining", "reasoning", "regulation", "research", "training"]
source: docs/RAG/clean_en/huggingface/nvidia-nvidia-nemotron-3-ultra-550b-a55b-bf16-hugging-face.md
source_anchor: ""
source_lines: [576, 628]
sha256: b9291a0272c03b12c03f129ecfa4bbfbda7b7ef933fb28c7aa2af48d0e74d393
---

# Set the IP for the head node in RAY_HEAD_IP

For all domains, we apply a unified data filtering pipeline to ensure that only high-quality, license-compliant, and verifiable samples are used for post-training. We first discard malformed examples using structural checks (e.g., missing tool definitions when tool calls are present). We then aggressively filter reasoning traces exhibiting pathological repetition, such as repeated n-grams within a sliding window or across the entire trajectory, which we found to be a strong indicator of malformed or low-quality reasoning. Finally, based on internal audits of synthetically generated datasets, we observed that some teacher models occasionally produce reasoning traces and final responses that implicitly align with specific political entities or promote nationalistic narratives. To mitigate this, we apply targeted keyword- and regex-based filters and remove all trajectories matching such behavior.

Alongside the model, we release our final pre-training and post-training data, as outlined in this section. For ease of analysis, there is a sample set that is ungated. For all remaining code, math and multilingual data, gating and approval is required, and the dataset is permissively licensed for model training purposes.

More details on the datasets and synthetic data generation methods can be found in the technical report ***NVIDIA Nemotron 3 Ultra***.

For more information about the datasets used to train this model, please see the Public Summary of Training Content

## For Detailed Dataset Information: Click here!

The foundation of the model is trained on the **Nemotron-3-Ultra** corpus, comprising the following datasets from the Nemotron Pretraining Datasets collection:

| Dataset Collection | Token Counts | Description | 
|---|---|---|
| **Nemotron-CC-v2** &**v2.1** | 9.1T | A massive collection of English web data filtered from Common Crawl, including 2.5T+ tokens of new organic, translated, and synthetically rephrased content. | 
| **Nemotron-CC-Code-v1** | 427.9B | High-quality code tokens extracted from Common Crawl using the Lynx + LLM pipeline to preserve structure and equations. | 
| **Nemotron-Pretraining-Code-v1** &**v2** &**v3** | 1.7T | Curated GitHub code references with multi-stage filtering, deduplication, and large-scale synthetic code data. | 
| **Nemotron-CC-Math-v1** | 133.3B | High-quality math pre-training dataset preserving LaTeX formatting and mathematical structures. | 
| **Nemotron-Pretraining-Specialized-v1** &**v1.1** &**v1.2** &**Nemotron-Pretraining-SFT-v1** | 660.0B | Synthetic datasets targeting specialized domains such as STEM reasoning and scientific coding. | 
| **Nemotron-Pretraining-Legal-v1** | 4.3B | Synthetic datasets targeting the legal domain. | 

The English Common Crawl data was downloaded from the Common Crawl Foundation (see their FAQ for details on their crawling) and includes the snapshots CC-MAIN-2013-20 through CC-MAIN-2025-13. The data was subsequently deduplicated and filtered in various ways described in the Nemotron-CC paper. Additionally, we extracted data for fifteen languages from the following three Common Crawl snapshots: CC-MAIN-2024-51, CC-MAIN-2025-08, CC-MAIN-2025-18. The fifteen languages included were Arabic, Chinese, Danish, Dutch, French, German, Italian, Japanese, Korean, Polish, Portuguese, Russian, Spanish, Swedish, and Thai. As we did not have reliable multilingual model-based quality classifiers available, we applied just heuristic filtering instead—similar to what we did for lower quality English data in the Nemotron-CC pipeline, but selectively removing some filters for some languages that did not work well. Deduplication was done in the same way as for Nemotron-CC.

The GitHub Crawl was collected using the GitHub REST API and the Amazon S3 API. Each crawl was operated in accordance with the rate limits set by its respective source, either GitHub or S3. We collect raw source code and subsequently remove any having a license which does not exist in our permissive-license set (for additional details, refer to the technical report).

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


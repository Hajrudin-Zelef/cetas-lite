---
id: collect-240926-huggingface/huggingface/nvidia-nvidia-nemotron-3-nano-30b-a3b-nvfp4-hugging-face-4
title: "Load tokenizer and model"
domain: huggingface
role: reference
task: reference
actors: ["AWS", "Alibaba", "China", "Nvidia"]
dates: ["2025-01-05", "2025-02-10", "2025-08-04", "2025-29-04"]
keywords: ["license", "memory", "nvidia", "regulation", "research"]
source: docs/RAG/clean_en/huggingface/nvidia-nvidia-nemotron-3-nano-30b-a3b-nvfp4-hugging-face.md
source_anchor: ""
source_lines: [369, 394]
sha256: 37c5ca01027b47520ed4761f1f4d07811c6abc67299c3da52cf082b86b0ccaaa
---

# Load tokenizer and model

| Dataset | 
|---|
| Global Regulation | 
| TAUS Translation Memory | 
| Scale HLE | 
| HackerRank Coding | 

| Dataset | 
|---|
| Simple Minesweeper | 
| Simple Sudoku | 
| Multitool Typewriter Hard | 
| Machine Translation of News Commentary and TAUS Translation Memory | 
| Machine Translation of STEM data using Qwen2.5-14B-Instruct | 

The English Common Crawl data was downloaded from the Common Crawl Foundation (see their FAQ for details on their crawling) and includes the snapshots CC-MAIN-2013-20 through CC-MAIN-2025-13. The data was subsequently deduplicated and filtered in various ways described in the Nemotron-CC paper. Additionally, we extracted data for fifteen languages from the following three Common Crawl snapshots: CC-MAIN-2024-51, CC-MAIN-2025-08, CC-MAIN-2025-18. The fifteen languages included were Arabic, Chinese, Danish, Dutch, French, German, Italian, Japanese, Korean, Polish, Portuguese, Russian, Spanish, Swedish, and Thai. As we did not have reliable multilingual model-based quality classifiers available, we applied just heuristic filtering instead—similar to what we did for lower quality English data in the Nemotron-CC pipeline, but selectively removing some filters for some languages that did not work well. Deduplication was done in the same way as for Nemotron-CC.

The GitHub Crawl was collected using the GitHub REST API and the Amazon S3 API. Each crawl was operated in accordance with the rate limits set by its respective source, either GitHub or S3. We collect raw source code and subsequently remove any having a license which does not exist in our permissive-license set (for additional details, refer to the technical report).

| Dataset | Modality | Dataset Size | Collection Period | Collecting Organisation | 
|---|---|---|---|---|
| English Common Crawl | Text | 3.36T | 4/8/2025 | NVIDIA Advanced Deep Learning Research | 
| English Common Crawl 1.1 | Text | Not disclosed | 10/2/2025 | NVIDIA Advanced Deep Learning Research | 
| Multilingual Common Crawl | Text | 812.7B | 5/1/2025 | NVIDIA Advanced Deep Learning Research | 
| GitHub Crawl | Text | 747.4B | 4/29/2025 | NVIDIA Advanced Deep Learning Research | 


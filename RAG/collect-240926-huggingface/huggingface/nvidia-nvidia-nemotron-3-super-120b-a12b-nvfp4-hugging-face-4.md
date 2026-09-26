---
id: collect-240926-huggingface/huggingface/nvidia-nvidia-nemotron-3-super-120b-a12b-nvfp4-hugging-face-4
title: "with uv: uv pip install vllm==0.20.0 --torch-backend=auto"
domain: huggingface
role: reference
task: reference
actors: ["AWS", "Alibaba", "Google", "Nvidia", "OpenAI"]
dates: ["2025-01-05", "2025-02-10", "2025-08-04", "2025-29-04"]
keywords: ["gemini", "license", "memory", "nvidia", "pretraining", "reasoning", "regulation", "research", "training"]
source: docs/RAG/clean_en/huggingface/nvidia-nvidia-nemotron-3-super-120b-a12b-nvfp4-hugging-face.md
source_anchor: ""
source_lines: [473, 513]
sha256: 9fd03b11c4e6e1837bc90ddc374b4d0b5fd957802c405139c7267b8bfe669b07
---

# with uv: uv pip install vllm==0.20.0 --torch-backend=auto

The foundation of the model is trained on the **Nemotron-3-Nano** corpus, comprising the following collections:

| Dataset Collection | Token Counts | Description | 
|---|---|---|
| **Nemotron-CC-v2** &**v2.1** | 9.13T | A massive collection of English web data filtered from Common Crawl, including 2.5T+ tokens of new organic, translated, and synthetically rephrased content. | 
| **Nemotron-CC-Code-v1** | 427.9B | High-quality code tokens extracted from Common Crawl using the Lynx + LLM pipeline to preserve structure and equations. | 
| **Nemotron-Pretraining-Code-v1** &**v2** | 1.09T | Curated GitHub code references with multi-stage filtering, deduplication, and large-scale synthetic code data. | 
| **Nemotron-CC-Math-v1** | 133.3B | High-quality math pre-training dataset preserving LaTeX formatting and mathematical structures. | 
| **Nemotron-Pretraining-Specialized-v1** | 336.4B | Synthetic datasets targeting specialized domains such as STEM reasoning and scientific coding. | 

The GitHub Crawl was collected using the GitHub REST API and the Amazon S3 API. Each crawl was operated in accordance with the rate limits set by its respective source, either GitHub or S3. We collect raw source code and subsequently remove any having a license which does not exist in our permissive-license set (for additional details, refer to the technical report).

| Dataset | Modality | Dataset Size | Collection Period | Collecting Organisation | 
|---|---|---|---|---|
| English Common Crawl | Text | 3.36T | 4/8/2025 | NVIDIA Advanced Deep Learning Research | 
| English Common Crawl 1.1 | Text | Not disclosed | 10/2/2025 | NVIDIA Advanced Deep Learning Research | 
| Multilingual Common Crawl | Text | 812.7B | 5/1/2025 | NVIDIA Advanced Deep Learning Research | 
| GitHub Crawl | Text | 747.4B | 4/29/2025 | NVIDIA Advanced Deep Learning Research | 

| Dataset | Model(s) used | 
|---|---|
| Global Regulation | Unknown | 
| TAUS Translation Memory | Unknown | 
| Scale HLE | Unknown | 
| HackerRank Coding | Unknown | 
| RL data for Search | Gemini 3; GPT-5 * | 

- Models used for prompt generation only

| Dataset | Model(s) used | 
|---|---|
| Simple Minesweeper | - | 
| Simple Sudoku | - | 
| Multitool Typewriter Hard | - | 
| Machine Translation of News Commentary and TAUS Translation Memory | - | 
| Machine Translation of STEM - | Qwen2.5-14B-Instruct | 
| Competitive Coding RL data from Nemotron Cascade | - | 
| Long context RL | - | 
| Single-step SWE RL for patch generation | - | 
| OpenHands SWE | - | 


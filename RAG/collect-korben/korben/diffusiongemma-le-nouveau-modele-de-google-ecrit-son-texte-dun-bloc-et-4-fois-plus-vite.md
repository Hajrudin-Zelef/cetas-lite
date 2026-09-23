---
id: collect-korben/korben/diffusiongemma-le-nouveau-modele-de-google-ecrit-son-texte-dun-bloc-et-4-fois-plus-vite
title: "DiffusionGemma : le nouveau modèle de Google écrit son texte d'un bloc, et 4 fois plus vite"
domain: korben
role: reference
task: article
actors: ["Apple", "Google", "Hugging Face", "Nvidia"]
dates: ["2026-09-23"]
keywords: ["diffusion", "accelerator", "apache", "consumer", "gemini", "gpu", "inference", "license", "memory", "mixture of experts", "moe", "nvidia"]
source: docs/RAG/Collect RAG/01_korben/diffusiongemma-le-nouveau-modele-de-google-ecrit-son-texte-dun-bloc-et-4-fois-plus-vite.md
source_anchor: ""
source_lines: [1, 55]
sha256: f610a47830f08da23d4e936c8ea66fe0c61000db5020cec9c1c5b32b15ee954b
---

# DiffusionGemma : le nouveau modèle de Google écrit son texte d'un bloc, et 4 fois plus vite

## Metadata

- **Source** : https://korben.info/diffusiongemma-le-nouveau-modele-de-google-ecrit-son-texte-dun-bloc-et-4-fois-plus-vite.html
- **Site** : Korben.info
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

The article reports on DiffusionGemma, a new open AI model from Google DeepMind that reaches more than 1,000 tokens per second on a single H100 (Nvidia's data-center accelerator) and about 700 on an RTX 5090 gaming card. That is roughly four times what classic Gemma models of comparable size produce.

The difference lies in how the text is generated. Usual language models are autoregressive: they write left to right, one token at a time. DiffusionGemma does it differently. It works like image generators, which start from a cloud of noise and denoise it gradually until the requested picture. The model lays down a canvas of 256 dummy tokens, goes over it several times to refine its estimates, then finalizes the whole block at once.

Under the hood it is a Mixture of Experts with 26 billion parameters, an architecture where only a small part of the model wakes up at each calculation, 3.8 billion here. The whole thing fits in 18 GB of compressed video memory, i.e., a large consumer graphics card. Locally, this approach shifts the bottleneck from memory bandwidth (the speed at which the card reads its own data) to pure computation. In the cloud, servers pool the requests of thousands of users and their chips run constantly, whereas your home GPU spends most of its time waiting for data. Diffusion occupies those lost cycles.

Then there are non-linear tasks, where the writing order does not follow the reading order. Google even refined a version on Sudoku, a puzzle deemed impossible for classic models since each cell depends on cells not yet written. DiffusionGemma, which corrects its canvas continuously, reaches 80% success while dropping the computation steps from 48 to 12.

Everything is not rosy, however. In an image, a failed pixel goes unnoticed; a badly predicted token can render an entire paragraph incoherent and force a full restart. And for a five-word answer, refining a complete canvas wastes computation. That is why the big cloud Gemini models do not use it. The model is experimental, but it is released under the Apache 2.0 license, the same as the rest of the Gemma 4 family, so it is commercially usable without restriction. The weights are downloadable now on Hugging Face, with optimization carried out hand in hand with Nvidia. Apple's MLX tool is also involved, so Macs are served. The author concludes that Google is most interesting right now on these local models, far more than on Gemini.

## Key points

- DiffusionGemma generates text in blocks rather than token by token, like image-diffusion models.
- Speed: >1,000 tokens/s on an H100, ~700 on an RTX 5090, about 4x classic Gemma models of comparable size.
- Architecture: Mixture of Experts, 26B parameters, 3.8B active; fits in 18 GB of compressed VRAM.
- Shifts the bottleneck from memory bandwidth to pure computation, exploiting idle GPU cycles locally.
- Non-linear tasks: a Sudoku-refined version reaches 80% success, dropping computation steps from 48 to 12.
- Limitations: one bad token can corrupt a whole paragraph; short answers waste computation on a full canvas.
- License: Apache 2.0 (like Gemma 4), weights on Hugging Face, optimization with Nvidia, MLX support for Mac.
- Author finds Google most interesting on these local open models rather than on Gemini.

## Technical data / figures

| Item | Value |
|---|---|
| Model | DiffusionGemma (Google DeepMind) |
| Generation method | Diffusion (block-wise), 256-token canvas refined in passes |
| Parameters | 26 billion (MoE) |
| Active parameters | 3.8 billion |
| Throughput | >1,000 tokens/s (H100); ~700 tokens/s (RTX 5090) |
| Speed vs classic Gemma | ~4x |
| VRAM (compressed) | 18 GB |
| License | Apache 2.0 |
| Weights | Hugging Face |
| Optimization | With Nvidia |
| Mac support | MLX |
| Sudoku success | 80% (steps 48 → 12) |
| Cloud Gemini | Not used (wasteful for short outputs) |

## Why this source matters for the RAG

This article explains a novel LLM architecture (diffusion-based text generation) with precise performance and hardware figures, plus concrete trade-offs and local-vs-cloud implications. It is valuable for questions about open models, inference speed, model architectures and local AI deployment.

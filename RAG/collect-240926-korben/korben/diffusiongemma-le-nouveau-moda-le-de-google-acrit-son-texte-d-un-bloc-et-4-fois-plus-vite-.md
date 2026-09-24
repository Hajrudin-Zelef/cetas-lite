---
id: collect-240926-korben/korben/diffusiongemma-le-nouveau-moda-le-de-google-acrit-son-texte-d-un-bloc-et-4-fois-plus-vite-
title: "DiffusionGemma: Google's new model writes its text in one block, and 4 times faster"
domain: korben
role: reference
task: reference
actors: ["Anthropic", "Apple", "Google", "Hugging Face", "Nvidia"]
dates: []
keywords: ["diffusion", "accelerator", "apache", "claude", "consumer", "gemini", "gpu", "license", "memory", "mixture of experts", "nvidia", "parameters"]
source: docs/RAG/clean_en/korben/diffusiongemma-le-nouveau-moda-le-de-google-acrit-son-texte-d-un-bloc-et-4-fois-plus-vite-korben.md
source_anchor: ""
source_lines: [1, 33]
sha256: eb4bbbc720166e3b6551739b4792361d64ae9f4806cfd2ab661ffa1c837618ab
---

# DiffusionGemma: Google's new model writes its text in one block, and 4 times faster

<!-- source: https://korben.info/diffusiongemma-le-nouveau-modele-de-google-ecrit-son-texte-dun-bloc-et-4-fois-plus-vite.html -->

# DiffusionGemma: Google's new model writes its text in one block, and 4 times faster

## Key takeaways AI-generated summary

1. DiffusionGemma generates text in blocks rather than token by token, reaching over 1,000 tokens/second on H100 and about 700 on RTX 5090, four times faster than comparable-sized classic Gemma models.
2. The model works like an image generator: it lays down a canvas of 256 fictitious tokens, refines it over several passes, then finalizes the entire block at once, which allows it to solve non-linear tasks like Sudoku with 80% success.
3. DiffusionGemma (26 billion parameters, 3.8 billion active) fits in 18 GB of compressed video memory, ships under Apache 2.0 license with weights downloadable on Hugging Face, and works on Mac via MLX.

Over 1,000 tokens per second on a single H100 card, the accelerator Nvidia sells to data centers, and about 700 on an RTX 5090, its high-end gaming card. That's the throughput Google DeepMind is announcing for DiffusionGemma, its new open AI model, roughly four times what classic Gemma models of comparable size produce.

The whole difference lies in how text is generated. Usual language models are autoregressive: they write from left to right, one token at a time, the token being the small piece of a word that an AI manipulates. DiffusionGemma does everything differently.

It works like image generators, which start from a cloud of noise and denoise it little by little until the requested photo. The model lays down a canvas of 256 fictitious tokens, goes over it several times to refine its estimates, then finalizes the entire block at once.

Under the hood, there's a Mixture of Experts with 26 billion parameters, an architecture where only a small part of the model wakes up at each calculation, 3.8 billion here. As a result the whole thing fits in 18 GB of video memory in compressed version, that is a large consumer graphics card.

The interest locally is that this approach shifts the bottleneck from memory bandwidth, the speed at which the card reads its own data, to pure computation. In the cloud, servers pool requests from thousands of users and their chips run continuously, whereas your GPU at home spends most of its time waiting for data. Diffusion occupies these lost cycles.

And then there are non-linear tasks, where the writing order doesn't follow the reading order. Google even fine-tuned a version on Sudoku, a puzzle reputed impossible for classic models since each cell depends on cells not yet written. DiffusionGemma, which corrects its canvas continuously, reaches 80% success while cutting computation steps from 48 to 12.

Not everything is rosy though. In an image, a missed pixel goes unnoticed. A mispredicted token, on the other hand, can make an entire paragraph incoherent and force starting over. And for a five-word answer, trimming down a complete canvas wastes computation. That's actually why the big cloud Gemini models don't go for it.

The model is experimental, but it ships under Apache 2.0 license, the same as the rest of the Gemma 4 family, so usable commercially without restriction. The weights can be downloaded right now on Hugging Face, the reference platform for open models, with an optimization carried out hand in hand with Nvidia. MLX, Apple's tool for running AI locally, is also part of it, so Macs are served.

If you want my opinion, it's on these local models that Google is most interesting right now, much more than on Gemini.

Source: ARS Technica

## Comments

starfix!in Surfshark doesn't make you invMorganein Discord guesses your age sansts3rv1in The Ray-Ban Display arrive eponponin Openpilot - The NHTSA passes lesfabiengin Claude Code makes you choose

---
id: vague2-nerdykings/nerdykings/sakana-diffusionblocks-vram-entrainement-ia-1
title: "DiffusionBlocks : Sakana AI Divise La Mémoire D'Entraînement Par 4"
domain: nerdykings
role: reference
task: article
actors: ["Anthropic", "DeepSeek", "Nvidia", "OpenAI", "Sakana"]
dates: ["2026-09-23"]
keywords: ["diffusion", "sakana", "chatgpt", "claude", "deepseek", "gpu", "gpus", "llama", "memory", "nvlink", "research", "training"]
source: docs/RAG/Collect RAG Vague 2/03_nerdykings/sakana-diffusionblocks-vram-entrainement-ia.md
source_anchor: ""
source_lines: [1, 54]
sha256: 6c379f06105c10ca6cc309c563d8be405afebd84e46efc1ef2ba97f2ef15b659
---

# DiffusionBlocks : Sakana AI Divise La Mémoire D'Entraînement Par 4

## Metadata

- **Source** : https://www.nerdykings.com/blog/sakana-diffusionblocks-vram-entrainement-ia.html
- **Site** : NerdyKings
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

Training a powerful AI today requires dozens or even hundreds of interconnected GPUs, and above all enormous memory. A team from the Japanese lab **Sakana AI** has published a method called **DiffusionBlocks**, able to reduce this required memory by 3x, 4x, or more — without sacrificing performance. The basic idea is almost disarming in its simplicity: split the model into several blocks and train each block independently.

Why AI training consumes so much memory: using a model like ChatGPT is mostly a forward pass — text traverses the network layers and the model produces an answer. Training is much heavier. The model first does the forward pass, compares the result with the expected answer, then runs backward to determine which parts of the network to correct — that's backpropagation. To go backward, the system must keep in memory the intermediate results of each layer (activations), the model weights, the gradients used to correct them, and the optimizer states. Concretely, for a 10-billion-parameter model, the weights alone already represent ~40 GB in FP32. Adding gradients, optimizer states, and activations, the required memory far exceeds 100 GB. When several GPUs share the load, they must constantly exchange information, adding a communication bottleneck.

DiffusionBlocks: train the model block by block. This observation led Sakana AI to a radical question: must all layers really be trained at the same time? With DiffusionBlocks, researchers split the network into several large blocks — e.g., a 12-layer transformer can be divided into three blocks of four layers. Instead of loading all 12 layers and computing gradients for the whole model at once, the system trains one block at a time, then moves to the next once finished. Since memory now depends mainly on one block's size rather than the whole model, splitting into three blocks theoretically reduces memory by ~3x, and four blocks by ~4x. But for this to work, each block must know precisely what it must learn — and that's where the diffusion principle giving the method its name comes in.

The link with diffusion image generation: in an image generation model, one starts from pure noise, which the model progressively denoises — first finding large shapes, then building structure, then adding fine details. Sakana AI's researchers noticed that the passage through a transformer's layers could be seen the same way: each group of layers receives imperfect information, makes a small correction, then passes a slightly cleaner state to the next group. With DiffusionBlocks, each block becomes a specialist of one precise denoising step: the first learns to work with very noisy information, the second handles a more structured version, the last focuses on details and produces the final result. Above all, each block can learn its task without waiting for the others' results — like a production line where each team has a well-defined mission without keeping the whole line in memory permanently.

Results surprisingly close to — sometimes better than — classical training. On an image classification test with three blocks, classical training got ~60.5% accuracy vs 59.3% for DiffusionBlocks, with a theoretical memory footprint ~3x smaller. On image generation, the result is even more surprising: on ImageNet, the normally trained model got an FID score of 12.09 (lower is better for quality and diversity) vs 10.63 for DiffusionBlocks — the lower-memory method slightly beat classical training. Researchers also tested a small Llama 2-inspired language model split into four blocks, with performance remaining close to classical training. So the method isn't limited to image generation.

The limit: too many blocks degrade quality. One might think splitting into 20 or 30 blocks would save maximum memory, but experiments show a clear limit: the more blocks, the fewer layers each contains, and the more it loses its ability to understand the transformation it must perform. For image generation, two or three blocks generally give the best results, and from six blocks onward quality degrades markedly. There's a real trade-off between memory saved and each block's learning capacity.

What it could change at scale: if DiffusionBlocks works on multi-billion-parameter models, the consequences could be significant. Researchers and companies with fewer resources could train or customize much larger models on more affordable GPUs. It could also simplify distributed training: since blocks train independently, they no longer need to constantly exchange activations and gradients with other GPUs. Each machine could process a different block with far less communication — potentially reducing dependence on ultra-fast connections like **NVLink** or **InfiniBand**, similar in spirit to DeepSeek's recent Dual Path. Caution: the experiments presented are still very far from the scale of models like GPT, Claude, or DeepSeek, and researchers acknowledge scaling to much larger models remains to be demonstrated. The method also concerns mainly training — to quickly run an autoregressive LLM, all its layers must still be loaded into memory. So DiffusionBlocks, as is, doesn't let you run a giant model on a small GPU.

The author's view: despite this limit, DiffusionBlocks questions a rule almost taken for granted in modern AI — that all layers of a network must necessarily be trained together. Sakana AI shows that by turning each group of layers into a specialist of a denoising step, a network can be trained block by block while keeping competitive performance. It's the same kind of bet found in recent attempts to rethink transformer architecture: at some point it's no longer enough to have more GPUs; you must also question how models are trained. For now it's promising research, not yet an industrial revolution, but if it scales, it could make AI training much more accessible and shake up part of the current GPU economy.

## Key points

- **DiffusionBlocks** from Sakana AI reduces training memory by 3x–4x+ without sacrificing performance.
- Idea: split the network into large blocks and train one block at a time instead of all layers at once.
- Memory depends on one block's size, not the whole model; more blocks = more savings (up to a limit).
- Uses a diffusion-inspired view: each block specializes in one denoising step (noise → shapes → structure → details).
- Blocks learn independently without waiting for others, easing distributed training and communication.
- Classification test: ~60.5% (classical) vs 59.3% (DiffusionBlocks) with ~3x less memory.
- ImageNet generation: FID 12.09 (classical) vs 10.63 (DiffusionBlocks) — lower-memory method slightly better.
- Limit: too many blocks degrade quality; 2–3 blocks best, degradation from 6 blocks.
- Concerns mainly training; running autoregressive LLMs still needs all layers in memory.
- Scaling to GPT/Claude/DeepSeek-scale models remains unproven.

## Technical data / figures

| Metric | Classical training | DiffusionBlocks |
|---|---|---|
| Classification accuracy (3 blocks) | ~60.5% | 59.3% |
| ImageNet FID (lower is better) | 12.09 | 10.63 |
| Theoretical memory reduction | — | ~3x (3 blocks), ~4x (4 blocks) |
| Optimal block count (images) | — | 2–3 blocks |
| Degradation threshold | — | from 6 blocks |
| 10B model weights (FP32) | ~40 GB | — |
| Total memory (10B, classical) | >100 GB | — |


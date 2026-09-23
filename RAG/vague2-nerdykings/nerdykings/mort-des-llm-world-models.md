---
id: vague2-nerdykings/nerdykings/mort-des-llm-world-models
title: "La Mort Des LLM : World Models Vont Tout Changer En 2026"
domain: nerdykings
role: reference
task: article
actors: ["Meta", "Nvidia", "OpenAI"]
dates: ["2026-09-23"]
keywords: ["blackwell", "embedding", "gpus", "nvidia", "robotics"]
source: docs/RAG/Collect RAG Vague 2/03_nerdykings/mort-des-llm-world-models.md
source_anchor: ""
source_lines: [1, 47]
sha256: 68d40b6a7ece42ff3bd6ed4f8a11ff6c3558984308efb12d367fa7f08cc975f8
---

# La Mort Des LLM : World Models Vont Tout Changer En 2026

## Metadata

- **Source** : https://www.nerdykings.com/blog/mort-des-llm-world-models.html
- **Site** : NerdyKings
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

The article argues that after two years of increasingly fluent LLMs, a persistent malaise remains: these systems can describe the world but not live in it. This limit could change everything in 2026 through World Models. It cites Yann LeCun's long-standing diagnosis: language models are brilliant but fundamentally blind to reality; they learn to talk about the world by reading texts, as if one could understand physics by flipping through a dictionary. For LeCun, language is not the origin of intelligence but a late layer built on something more fundamental — an intuitive understanding of the world, movement, objects, and the physical consequences of actions.

The "all-language" paradigm has hit a ceiling, and a more radical approach emerges: teach AI to predict the world itself, not words. A World Model is not a pair of eyes but an internal multisensory simulation engine. It learns not from PDF libraries but by ingesting massive video streams coupled with force, pressure, and movement data, seeking to build an understanding of Newtonian laws. At the heart is the JEPA architecture (Joint Embedding Predictive Architecture), whose break is the difference between generating and predicting. Sora tries to reconstruct every pixel — a titanic task that collapses under noise — while JEPA works exclusively in latent space, turning the world into abstract mathematical vectors. It doesn't ask "what exact color will the pixel be" but "what will the semantic and physical state of the object be," ignoring superfluous details to focus on physical invariants and causality. Its key module, the controller, develops intuitive physics: it doesn't see a cup fall but infers its mass, friction, and inertia — like your brain instantly projecting "object broken on the floor" when you push a cup.

The article describes a "war of doctrines." Yann LeCun and AMI Labs (founded after leaving Meta) bet on radical efficiency with V-JEPA: predicting the world in latent space, aiming for learning efficiency close to living beings (where an LLM reads the whole internet to learn a rule, V-JEPA learns causality from a few hours of raw video). OpenAI with Sora 2 Pro stays faithful to "bigger is smarter," making physics emerge through brute-force generation (video autoregressive modeling), costly in GPUs with no guarantee of causal understanding. Nvidia with Cosmos impressed at CES 2026 with World Foundation Models designed as the operating system for factories and autonomous cars, tokenizing physical reality for its Blackwell chips.

Why there is no effective domestic robot yet: unpredictability. Classic robotics relied on rigid programming; the robot fails as soon as reality deviates by a millimeter. With World Models as a motor brain, robots like Tesla Optimus Gen 3 or Figure 02 simulate the action in latent space before executing it — a robot grabbing a slippery object senses the coming instability and adjusts pressure before the physical accident. The future is not the death of LLMs but a fusion into Large World Models (LWM), inspired by human neurobiology: the LLM as prefrontal cortex (language, long-term planning, abstraction, concepts and intentions) and the World Model as motor cortex (sensory perception, spatial geometry, real-time dynamics, thinking in force vectors and trajectories). The LLM decomposes a sentence and sends a high-level command; the World Model simulates the path, anticipates obstacles, and adjusts each servomotor. The LLM gives the order; the World Model ensures it survives physical reality.

## Key points

- LLMs describe the world but cannot inhabit it; World Models aim to fix this in 2026.
- Yann LeCun: language is a late layer on top of intuitive physical understanding.
- World Models are internal multisensory simulation engines trained on video + force/pressure/movement.
- JEPA predicts in latent space (semantic/physical state) rather than reconstructing pixels like Sora.
- Doctrine war: LeCun/AMI Labs (V-JEPA), OpenAI (Sora 2 Pro), Nvidia (Cosmos).
- World Models enable predictive robotics (Optimus Gen 3, Figure 02) instead of reactive programming.
- Future is fusion into Large World Models (LWM): LLM = prefrontal cortex, World Model = motor cortex.

## Technical data / figures

| Item | Detail |
|---|---|
| Key architecture | JEPA (Joint Embedding Predictive Architecture) |
| JEPA approach | Latent-space prediction of physical state |
| Sora approach | Pixel reconstruction / video autoregressive modeling |
| Actors | LeCun/AMI Labs (V-JEPA), OpenAI (Sora 2 Pro), Nvidia (Cosmos) |
| Nvidia hardware | Blackwell |
| Robots cited | Tesla Optimus Gen 3, Figure 02 |
| Future paradigm | LWM (Large World Models) |
| Nvidia event | CES 2026 |

## Why this source matters for the RAG

This source explains the World Models paradigm and the debate over language-centric vs simulation-based AI, including the JEPA architecture and robotics implications. It is valuable for RAG corpora on AI architectures, embodied AI, and the future of LLMs.

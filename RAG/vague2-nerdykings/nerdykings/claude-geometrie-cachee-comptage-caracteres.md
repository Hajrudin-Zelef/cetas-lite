---
id: vague2-nerdykings/nerdykings/claude-geometrie-cachee-comptage-caracteres
title: "L'inexplicable découverte cachée dans Claude"
domain: nerdykings
role: reference
task: article
actors: ["Anthropic"]
dates: ["2026-09-23"]
keywords: ["claude", "attention", "reasoning", "research"]
source: docs/RAG/Collect RAG Vague 2/03_nerdykings/claude-geometrie-cachee-comptage-caracteres.md
source_anchor: ""
source_lines: [1, 65]
sha256: f3f750642d019c11645b92a905aa0d234e90aa1a9424bba855c0a5166194d0e0
---

# L'inexplicable découverte cachée dans Claude

## Metadata

- **Source** : https://www.nerdykings.com/blog/claude-geometrie-cachee-comptage-caracteres.html
- **Site** : NerdyKings
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

How does Claude know where to wrap a line of text at the right width when it never "sees" text, only a sequence of tokens? An interpretability team at Anthropic answered this, and the answer is far stranger than a simple character counter: Claude builds, internally, a real geometry — a mechanism reminiscent of what your brain uses to know where you are in a room.

A simple question never really asked: when you ask an LLM to format text at fixed width — a table, a commented code block, justified text in a terminal — it must know how many characters remain before the margin. But an LLM doesn't "see" text as we understand it; it manipulates tokens (word pieces converted into numbers). It has no ruler, no cursor, no native notion of "80 characters." Yet it works most of the time. How?

This is exactly the question posed by seven Anthropic researchers — **Wes Gurnee, Emmanuel Ameisen, Isaac Kauvar, Julius Tarng, Adam Pearce, Chris Olah, and Joshua Batson** — in a paper with a striking title: **"When Models Manipulate Manifolds."** Their subject: **Claude 3.5 Haiku**, mechanically dissected on a single task — line wrapping in fixed-width text.

Place cells... in a language model: the model doesn't represent the number of characters as a simple digit stored somewhere. It represents it as points on curved, low-dimensional geometric surfaces — "manifolds." These manifolds are themselves cut into small units the researchers call "feature families": directions in the model's internal space that activate together to encode a precise position. The twist is the comparison the researchers themselves make: this organization resembles that of **place cells** discovered in the mammalian hippocampus — the mechanism that earned John O'Keefe and the couple May-Britt and Edvard Moser the 2014 Nobel Prize in Medicine. In a biological brain, these cells activate according to the animal's physical position in its environment. In Claude, a comparable structure seems to encode... a position in a sentence.

How the model "knows" where the margin is: once this map of the character count is set, attention heads directly manipulate it to estimate the remaining distance before the line's end. Researchers show these manifolds are organized orthogonally to one another, creating perfectly linear decision boundaries — a way for the model to "read" at a geometric glance whether it must still write characters or wrap now. Another finding: from the earliest layers, before even "reasoning" in the proper sense, the model performs what the authors call rich sensory processing — a form of early perception of text structure, even though it only receives a token sequence and never "sees" anything in the visual sense.

The real twist: the illusions that fool Claude. The most interesting part of the paper — the through-line of the author's video — are the illusions. Researchers managed to build specific character sequences able to hijack this counting mechanism — to make the model believe it's closer to or farther from the margin than it really is. The result: deliberately caused layout errors, exactly like an optical illusion tricks the human eye into perceiving a curved line where it is perfectly straight. And that's the paper's strongest point for the author: an illusion, by definition, only fools systems using a shortcut — a learned heuristic, not a universal, robust rule. The fact that Claude can be "fooled" this way is the best proof that this mechanism isn't an exact calculator-style module but a learned geometric construction, with its flaws.

So, a spark of intelligence or not? It's tempting to conclude "Claude has a brain," "it perceives space like us." Honestly, that would be going much too fast. This paper proves neither consciousness nor human-like understanding. It only proves that a model trained solely to predict the next token can, without being explicitly asked, build a surprisingly sophisticated internal machinery to solve a very concrete sub-problem. It isn't magic; it's optimization finding the most efficient solution available in parameter space — and that solution looks like geometry because geometry is probably the most efficient tool for this kind of task, whether in a biological brain or an artificial neural network.

The author's view: what's craziest isn't the existence of this internal geometry — we already knew neural networks encode concepts in a distributed way. What's impressive is the precision of mechanistic interpretability today: a team can open the black box, isolate an ultra-specific task — wrapping a line of text — and extract a complete mathematical structure, to the point of knowing exactly which character sequences will derail it. Does it change anything for you directly? No. But it shows interpretability research — understanding why a model does what it does, not just observing that it does — is advancing very fast at Anthropic. And in a world where more and more decisions are delegated to these models, knowing precisely how they "think" internally isn't an academic detail; it's probably one of the most important projects of 2026.

## Key points

- Anthropic researchers studied how Claude 3.5 Haiku wraps text at fixed width without ever "seeing" characters.
- Paper: **"When Models Manipulate Manifolds"** by Gurnee, Ameisen, Kauvar, Tarng, Pearce, Olah, Batson.
- The model encodes character counts as points on curved low-dimensional **manifolds**, split into "feature families."
- This resembles mammalian hippocampal **place cells** (O'Keefe, Moser & Moser, Nobel 2014), but encoding position in a sentence.
- Attention heads manipulate the manifold to estimate remaining distance to the margin.
- Manifolds are orthogonal, producing linear decision boundaries for wrapping.
- Early layers perform "rich sensory processing" — early structural perception before reasoning.
- Researchers built character sequences that create **illusions**, causing layout errors — evidence of a learned heuristic, not an exact calculator.
- The paper proves neither consciousness nor human-like understanding; it shows optimization discovering geometry.
- Interpretability is advancing fast and matters as models take on more decisions.

## Technical data / figures

| Item | Value |
|---|---|
| Paper title | When Models Manipulate Manifolds |
| Institution | Anthropic (interpretability team) |
| Subject model | Claude 3.5 Haiku |
| Task studied | Fixed-width text line wrapping |
| Researchers | Wes Gurnee, Emmanuel Ameisen, Isaac Kauvar, Julius Tarng, Adam Pearce, Chris Olah, Joshua Batson |
| Core finding | Character count encoded as low-dimensional manifolds / feature families |
| Biological analogy | Hippocampal place cells (Nobel 2014) |
| Mechanism | Orthogonal manifolds → linear decision boundaries |
| Key demonstration | Adversarial "illusions" causing wrapping errors |

- Key concepts: **manifolds**, **feature families**, **place cells**, **attention heads**, **mechanistic interpretability**, **rich sensory processing**

## Why this source matters for the RAG

This article summarizes a landmark mechanistic-interpretability result showing a language model spontaneously learning geometric representations for a concrete sub-task, with a striking neuroscience analogy. It is valuable for a RAG knowledge base on interpretability, emergent representations, and AI cognition.

## Source URL

https://www.nerdykings.com/blog/claude-geometrie-cachee-comptage-caracteres.html

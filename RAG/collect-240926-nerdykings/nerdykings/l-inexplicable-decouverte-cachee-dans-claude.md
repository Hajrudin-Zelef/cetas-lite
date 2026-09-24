---
id: collect-240926-nerdykings/nerdykings/l-inexplicable-decouverte-cachee-dans-claude
title: "The inexplicable discovery hidden in Claude"
domain: nerdykings
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["claude", "attention", "research"]
source: docs/RAG/clean_en/nerdykings/l-inexplicable-decouverte-cachee-dans-claude.md
source_anchor: ""
source_lines: [1, 43]
sha256: fcc6eae55140202802f40a1b6e6c55fbfc5912419d44d019b11ed688c6b13b60
---

# The inexplicable discovery hidden in Claude

<!-- source: https://www.nerdykings.com/blog/claude-geometrie-cachee-comptage-caracteres.html -->

# The inexplicable discovery hidden in Claude

How does Claude know where to wrap a line of text at just the right width… when it never sees text, only a sequence of tokens? A team of interpretability researchers at Anthropic has just answered this question, and the answer is far stranger than a simple character counter: Claude internally builds a genuine geometry — a mechanism reminiscent of the one your brain uses to know where you are in a room.

## A simple question, never really asked

When you ask an LLM to format text at a fixed width — a table, a commented code block, justified text in a terminal — it needs to know how many characters remain before it hits the margin. Except that an LLM doesn't "see" text in the way we understand it. It manipulates tokens, chunks of words converted into numbers. It has no ruler, no cursor, no native notion of "80 characters." And yet, it works, most of the time. How?

This is exactly the question seven Anthropic researchers — Wes Gurnee, Emmanuel Ameisen, Isaac Kauvar, Julius Tarng, Adam Pearce, Chris Olah, and Joshua Batson — asked themselves in a paper with a punchy title: *"When Models Manipulate Manifolds."* Their subject: **Claude 3.5 Haiku**, mechanically dissected on a single task — line wrapping in fixed-width text.

## Place cells... in a language model

What they find is that the model doesn't represent the number of characters as a simple number stored somewhere. It represents it as points on curved, low-dimensional geometric surfaces — **"manifolds."** These manifolds are themselves divided into small units that the researchers call "feature families": directions in the model's internal space that activate together to encode a precise position.

The twist is the comparison the researchers themselves make: this organization resembles that of the **place cells** discovered in the hippocampus of mammals — the mechanism that earned John O'Keefe and the couple May-Britt and Edvard Moser the Nobel Prize in Medicine in 2014. In a biological brain, these cells activate according to the animal's physical position in its environment. In Claude, a comparable structure seems to encode... a position in a sentence.

## How the model "knows" where the margin is

Once this map of character counts is established, **attention heads** come to manipulate it directly to estimate the distance remaining before the end of the line. The researchers show that these manifolds are organized orthogonally to one another, which creates perfectly linear decision boundaries — a way, for the model, to "read" at a geometric glance whether it should still write characters or wrap the line now.

Another discovery: from the very first layers, even before "thinking" in the proper sense, the model performs what the authors call **rich sensory processing** — a form of early perception of the text's structure, even though it only receives a sequence of tokens as input, never "seeing" anything in the visual sense of the term.

## The real twist: the illusions that fool Claude

The most interesting part of the paper, the one that served as the throughline for my video, is the **illusions**. The researchers managed to construct specific character sequences capable of hijacking this counting mechanism — of making the model believe it is closer to or farther from the margin than it actually is. The result: layout errors deliberately provoked, exactly as an optical illusion tricks the human eye into perceiving a curved line where it is perfectly straight.

And that, for me, is the strongest point of the paper: an illusion, by definition, only fools systems that use a shortcut — a learned heuristic, not a universal and robust rule. The fact that Claude can be "tricked" in this way is the best proof that this mechanism is not an exact calculator-style computation module, but rather a *learned* geometric construction, with its flaws.

## So, a spark of intelligence or not?

It's tempting to conclude "Claude has a brain," "it perceives space like we do"... Let's be honest: that would be going much too fast. This paper proves neither consciousness nor understanding in the human sense of the term. It merely proves that a model trained solely to predict the next token can, without being explicitly asked, build an astonishingly sophisticated internal machinery to solve a very concrete sub-problem. It's not magic, it's optimization finding the most efficient solution available in parameter space — and that solution looks like geometry here because geometry is probably the most efficient tool for this kind of task, whether in a biological brain or in an artificial neural network.

## My take

What I find craziest is not the existence of this internal geometry in itself — we already knew that neural networks encode concepts in a distributed way. What blows me away is the level of precision of **mechanistic interpretability** today: a team can open the black box, isolate an ultra-specific task — wrapping a line of text — and extract a complete mathematical structure from it, down to knowing exactly which character sequences will make it derail.

Concretely, does it change anything for you? Not directly. But it shows that interpretability research — understanding *why* a model does what it does, not just observing that it does it — is advancing really fast at Anthropic. And in a world where we're increasingly going to entrust decisions to these models, knowing precisely how they "think" internally is not an academic detail. It's probably one of the most important undertakings of 2026.

### 🛠️ Tools you can test related to this article

A selection of my tested tools, relevant for going further.

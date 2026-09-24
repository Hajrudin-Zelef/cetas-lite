---
id: collect-240926-nerdykings/nerdykings/loop-transformers-la-fin-du-raisonnement-en-tokens
title: "Loop Transformers: The End of Reasoning in Tokens?"
domain: nerdykings
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "Google", "OpenAI"]
dates: ["2026-04"]
keywords: ["reasoning", "agents", "attention", "benchmarks", "chatgpt", "claude", "deepseek", "gemini", "guardrails", "inference", "memory", "parameters"]
source: docs/RAG/clean_en/nerdykings/loop-transformers-la-fin-du-raisonnement-en-tokens.md
source_anchor: ""
source_lines: [1, 45]
sha256: b7e6df319d7abfa86224476e6a259f86b596e4649d7432a201d902518602ed5b
---

# Loop Transformers: The End of Reasoning in Tokens?

<!-- source: https://www.nerdykings.com/blog/loop-transformers-claude-raisonnement-tokens.html -->

# Loop Transformers: The End of Reasoning in Tokens?

Does Claude think differently from ChatGPT? Not at the system prompt level — at the architecture level, underneath. A theory is currently circulating in AI research: instead of making a model "talk" so it can think (the classic chain of thought, full of tokens generated before the answer), **loop transformers** loop a few layers of the network back onto themselves, letting the internal state refine directly, without ever going through text. And some researchers suspect this is one of the mechanisms that might be running behind Claude.

## Why chain of thought is a bit of a hack

Today, all reasoning models — o3, Gemini Thinking, DeepSeek R2, whatever — all do the same thing: before answering, they generate a ton of text to "think out loud." It works very well, it's become the absolute standard. But let's be honest for two seconds: it's a somewhat bizarre method. An AI that literally has to write words to itself in order to manage to think.

Take a simple question: "who is the wife of the 44th American president?" To answer, the model has to do two steps: retrieve that the 44th president is Obama, then remember that his wife is Michelle. Two steps that depend on each other — **multi-hop** reasoning. A single-step question, an AI solves it by pure pattern matching. But as soon as it chains several dependent steps, it has to follow a thread, keep track of where it is — and a classic model has to do all that in a single pass, without ever going back.

This is exactly the problem that chain of thought came to solve: instead of forcing the model to solve everything in one shot, we let it write an intermediate step, reread it, adjust its reasoning, write the next step. The problem is what happens under the hood: every word of the reasoning has to be transformed into text, appended to everything already written, then retransformed into an internal representation for the next step. The model compresses, decompresses, recompresses constantly, just to keep track of its own thought. A bit like if, to solve a math problem, you were forced to write each step on a post-it, read it aloud, then decipher it again before you could think about the next step.

## What if we looped the brain instead of making the mouth talk

The logical question is: what if the model could refine its thinking directly, on the inside, without ever going through words? That's the idea behind loop transformers, a concept that exploded in popularity in April 2026. The principle is simple: usually, a language model is a stack of dozens of different layers, and information passes through this stack only once. With a loop transformer, you take a small block of a few layers and, instead of using it once, you loop it back onto itself. The output of the first loop becomes the input of the second, which becomes the input of the third — the model refines its internal state without ever needing to produce text to "reread itself."

This is where the connection to Claude comes in. Some researchers suspect this mechanism would explain why Claude does particularly well on pure logic benchmarks, like traversing a graph or doing a BFS — tasks where you have to follow a path step by step, exactly the kind of problem where looping an internal state makes more sense than writing tokens. Note, this is a theory from researchers observing a behavior: **Anthropic has never officially confirmed this**. But the fact that serious researchers are starting to dig into this track precisely because of Claude's behavior is worth paying attention to.

A paper published in April 2026, "Loop, Think and Generalize," tested exactly this setup. By observing training, the researchers see the model go through three distinct stages. First **memorization**: the model learns the examples by heart, its performance on training data climbs and plateaus. Then **generalization**: it starts to generalize, but only on problems that resemble what it has already seen. Finally, the most interesting stage, **systematic generalization**: the model becomes capable of combining bits of knowledge in a way it has never seen during training. It has learned a real reusable reasoning procedure, not just answers by heart — something a classic transformer struggles to do.

## The number of loops as a reflection slider

The craziest detail: if you increase the number of loops at inference time, even beyond what the model saw during training, it becomes capable of solving even more complex problems. The number of loops literally becomes a slider — you want the model to think longer about a hard problem? You increase the number of loops.

Obviously, as soon as you loop an architecture back onto itself, it creates a new problem: instability. A bit like feedback in a microphone too close to an amp — as soon as the signal is reinjected in a loop onto itself, the slightest error can amplify with each turn and end up blowing everything up. A paper called "Parsy" looked into this, with a solution that consists of putting mathematical guardrails in place to prevent the internal state from running away with each loop.

## Does the model really think

Better performance on benchmarks doesn't prove anything by itself — it could be a lucky statistical accident. That's where a mechanistic analysis of looped models, also published in April 2026, becomes fascinating. The researchers tracked the model's internal state at each loop. Since this state has thousands of dimensions, they use a technique a bit like a satellite map zooming out: instead of displaying every pixel, it keeps only the broad trends, to project all that chaos onto a 2D or 3D map and see whether the model goes off in all directions or follows a coherent trajectory.

Result: the model does indeed follow stable trajectories, and looking more closely at how it uses its attention at each loop, the researchers notice that the loops fall into three distinct roles, almost like a real thinking process. The first loops serve to **understand** the problem roughly. The middle loops serve to **connect** information to each other. The last loops serve to **stabilize**, to converge toward the final answer. It's the same block of layers doing all of this each time — it's just that what it's given as input changes at each loop, which forces it to behave differently at each step. Exactly what you'd expect if you imagined real human reflection, only more mechanical. This kind of work also connects with interpretability efforts at Anthropic, which are also trying to understand what happens inside a model when it doesn't say it — and it echoes other recent avenues like agents that think without words or MIT's RLMs, which all explore this same idea: getting reasoning out of text.

A quick aside: another paper, "Mixture of Recursions," noticed that we waste computation by forcing every word to go through the same number of loops, even the easy words. Their solution is a router that decides how many loops each word really deserves, with tricks to save the memory that consumes. Nothing paradigm-shifting, but a good engineering optimization.

## My take

Are loop transformers going to replace chain of thought tomorrow? No, clearly not, for a simple reason: chain of thought has an unfair advantage, it's written in text. That means we can supervise it, correct it, train directly on it. A looped model thinks in a hidden, invisible space that we can't easily reread or correct. Architecturally, it's cleaner — but to train it, it's much harder to control. And concretely, if you're not limited by memory, simply adding more parameters to a classic model almost always wins against looping a small block.

Where I find it really exciting is for small models, the ones you run on a phone or locally, where you can't afford to have billions of parameters but where you still want real iterative reasoning. On the Claude question, I remain cautious: it's a theory from researchers observing a behavior, not a confirmation from Anthropic. But it's exactly the kind of research avenue that, in six months, could well become the norm for small models. One to watch very closely.

### 🛠️ Tools you can try related to this article

A selection of my tested tools, relevant for going further.

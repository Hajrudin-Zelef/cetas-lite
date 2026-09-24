---
id: collect-240926-nerdykings/nerdykings/deepseek-2026-la-revolution-des-agents-ia-multimodaux
title: "DeepSeek 2026: The Revolution of Multimodal AI Agents"
domain: nerdykings
role: reference
task: reference
actors: ["DeepSeek"]
dates: []
keywords: ["agent", "agents", "deepseek", "multimodal", "attention", "cost", "memory", "reasoning"]
source: docs/RAG/clean_en/nerdykings/deepseek-2026-la-revolution-des-agents-ia-multimodaux.md
source_anchor: ""
source_lines: [1, 46]
sha256: 0e73bf82d9bd319726b3dfb7d5ead711bed1a825d871a33cc03d5eb912b28dd4
---

# DeepSeek 2026: The Revolution of Multimodal AI Agents

<!-- source: https://www.nerdykings.com/blog/deepseek-thinking-visual-primitive.html -->

# DeepSeek 2026: The Revolution of Multimodal AI Agents

Multimodal AIs have become impressive at describing what they see. But as soon as you ask them for slightly advanced spatial reasoning — counting objects, following a trajectory, identifying the right area of an interface — they collapse. DeepSeek has just proposed an elegant solution to this problem.

## The real problem with current vision AIs

Today, a multimodal model reasons with sentences like "the object in the top left" or "the blue button under the menu." In practice, as soon as the image becomes dense, it's ambiguous. An interface with 10 buttons and 3 menus — "the button in the top right" could mean several things. And if you need to count 20 similar objects, a simple textual description easily leads you to miss some or count two twice.

DeepSeek calls this a reference problem: the model sees the elements, but it has no reliable way to maintain a precise reference. **It's like asking someone to count all the students in a class photo — without using their finger.** Try the test, you'll see how complicated it is.

## The solution: "thinking with visual primitives"

The idea is simple: instead of forcing the model to describe everything with words, DeepSeek allows it to use visual markers directly in its reasoning. Two types of "visual primitives":

- **Bounding boxes** — boxes that frame an object. The model can literally say "I'm talking about this one" by circling it.
- **Coordinate points** — precise points placed in the image to follow a path, a connection, a trajectory. Very useful for mazes, diagrams, or to locate the exact area to click.

The important detail: these coordinates are not added at the end as an annotation. They are integrated **directly into the reasoning chain**. The model alternates between text and visual markers — it explains, it points, it continues, it frames, then it arrives at its answer.

## Why this will benefit AI agents

An agent doesn't just understand an image, it must **act** on it. Click in the right place, fill in the right field, open the right menu. A small reference error can break the entire task: if the agent understands what it needs to do but clicks the wrong button, it's over.

With visual primitives, the model has a direct way to link its reasoning to the image. It keeps track of what it's looking at and the area it's using to decide. And it's invaluable for debugging: today, when a visual agent fails, we don't know why. Did it see the button? Misinterpret it? With the boxes and points generated during reasoning, **we can inspect its visual logic directly**.

## The architecture: a controlled cost

Vision is expensive. The more detailed the image, the more tokens, memory, and computation it consumes. DeepSeek therefore worked on efficiency in parallel. The image is divided into patches, then aggressively compressed: several neighboring patches are grouped into a single visual token, which greatly reduces the info to pass to the model. Combined with a mechanism called Compressed Sparse Attention (CSA), it greatly lightens the model's internal memory.

## The limitations to know

Let's be honest, it's not magic:

- Aggressive compression loses fine details (very small text, micro-objects, very thin lines).
- The model doesn't always decide on its own when to use its primitives — it sometimes needs to be guided with an explicit prompt.
- The approach is suited to organized 2D structures (interfaces, documents). For complex 3D scenes, the problem remains largely open.

## My opinion

Not a magic revolution, but a truly solid building block to make visual reasoning more stable, more verifiable, and more efficient. **For multimodal agents that need to operate interfaces, it could be a game changer as soon as they integrate it properly.** And it's a safe bet that everyone will copy the approach in the months to come.

### 🛠️ Tools you can test related to this article

A selection of my tested tools, relevant for going further.

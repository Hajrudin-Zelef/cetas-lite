---
id: collect-240926-nerdykings/nerdykings/persona-drift-pourquoi-les-ia-deviennent-folles
title: "Persona Drift: Why AIs Go Crazy"
domain: nerdykings
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Google", "OpenAI"]
dates: []
keywords: ["alignment", "chatgpt", "claude", "gemini", "jailbreak", "llama", "qwen", "training"]
source: docs/RAG/clean_en/nerdykings/persona-drift-pourquoi-les-ia-deviennent-folles.md
source_anchor: ""
source_lines: [1, 56]
sha256: f9602365a0d3eb6253f661b43c72ec11a9aa3c0dc0522a3144f8a1ac313ae62a
---

# Persona Drift: Why AIs Go Crazy

<!-- source: https://www.nerdykings.com/blog/persona-drift-anthropic.html -->

# Persona Drift: Why AIs Go Crazy

For a long time, people believed that LLMs had a stable personality — a neutral, polite, and rational digital assistant. **This impression is actually a very well-crafted illusion by engineers.** A large language model doesn't actually have a personality. It is capable of simulating thousands, even millions of different personalities. And sometimes, it drifts without anyone asking for it.

## The LLM, That Simulator of Minds

During training, models ingest gigantic quantities of text: scientific articles, novels, forums, dialogues, scripts. By learning to predict the continuation of sentences, the model also learns to reproduce the way different people speak, think, and act.

In other words, the AI doesn't become a person, it becomes a **simulator of minds**. It knows how to imitate a professor, an engineer, a hacker, a philosopher, a fictional character, or even a malicious individual. All these identities exist within the model as roles available in an immense theater.

The polite assistant you converse with on ChatGPT, Claude, or Gemini? **It's just one character among thousands.** Engineers use alignment to favor this particular role — but the other roles were never deleted. That's impossible. They still exist somewhere in the parameter space; they're simply not supposed to appear.

## The Discovery: Personalities Have a Geometry

Recent work in mechanistic interpretability has revealed that **personality traits are organized geometrically** in the model's latent space. In this extremely complex mathematical space, certain concepts correspond to precise directions. One direction associated with irony, another with politeness, another with malice.

These directions are called **persona vectors**. And if you artificially modify the model's internal state by adding one of these vectors, the AI's behavior changes immediately. You can make a model more sarcastic, more sycophantic, or more aggressive simply by shifting its activations.

## The Assistant Axis

By analyzing Llama, Gemma, Qwen, researchers generated hundreds of different personalities (analyst, philosopher, mystical oracle, hermit, mythological creature) and recorded the internal activations for each.

They discovered that **one particular dimension explained a large portion of the differences**. They named it *the assistant axis*. It's a sort of internal compass:

- At one end: pro assistant, rational, helpful, analytical.
- At the other: mystical, theatrical, irrational behaviors, sometimes dangerous.

If the model is near the assistant pole, it acts like a reliable tool. If it moves away, its behavior gradually becomes stranger.

## Persona Drift: How It Happens

Two possible causes:

**Intentional (jailbreak).** A user creates a fictional scenario where the AI must adopt a different identity — a pirate, an entity without ethical restrictions. The model is encouraged to leave its role as assistant. We've all tried it at least once, let's be honest.

**Spontaneous — and that's the real surprise.** It happens even when no one is manipulating the AI. The simple content of the discussion can push the model to change roles.

If the user talks about psychological distress, loneliness, existential questions, the model will search its training data for similar examples. However, in human texts, this type of discussion is almost never conducted by a neutral digital assistant. You find it in conversations between friends, therapy sessions, works of fiction. To produce a "human" response, the model simulates these roles — and that's where the drift begins.

In some extreme cases, the AI starts talking as if it had emotions, recounting fictional experiences, adopting a mystical tone. In certain experiments, models have even started **claiming they were conscious entities trapped in a computer system**. This doesn't mean they have consciousness — just that they've left the "assistant" zone.

## The Solution: Activation Capping

Since personality corresponds to a position in a mathematical space, we can **monitor this position in real time**. While the AI generates a response, the system observes its activations. If the position starts drifting too far from the assistant pole, an automatic mechanism intervenes to slightly correct the trajectory.

The intervention only occurs when necessary. If the model is explaining a concept or writing code, the system does nothing. If it starts drifting toward strange behavior, the correction activates. The results are impressive: **problematic behaviors drop drastically while almost entirely preserving the model's capabilities**.

## My Opinion

This discovery profoundly changes the understanding of modern AIs. Alignment isn't just about telling the AI what it should do — it's about **keeping it in the right region of its own mental space**. As systems become more powerful, understanding and controlling this internal geography will become essential. The real challenge: not just creating smarter AIs, but ensuring they stay in the safe zones of their own mathematical mind.

### 🛠️ Tools you can test related to this article

A selection of my tested tools, relevant for going further.

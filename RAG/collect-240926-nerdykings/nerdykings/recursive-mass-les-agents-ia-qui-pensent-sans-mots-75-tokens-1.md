---
id: collect-240926-nerdykings/nerdykings/recursive-mass-les-agents-ia-qui-pensent-sans-mots-75-tokens-1
title: "Recursive Mass: The AI Agents That Think Without Words (-75% Tokens)"
domain: nerdykings
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Nvidia"]
dates: []
keywords: ["agent", "agents", "benchmarks", "llama", "nvidia", "qwen", "reasoning", "research", "training"]
source: docs/RAG/clean_en/nerdykings/recursive-mass-les-agents-ia-qui-pensent-sans-mots-75-tokens.md
source_anchor: ""
source_lines: [1, 68]
sha256: 6e7e1bd6c1cc7ea127d5733bb7462a5b454889fe57ecb36e6509f9c0a5c8f04c
---

# Recursive Mass: The AI Agents That Think Without Words (-75% Tokens)

<!-- source: https://www.nerdykings.com/blog/recursive-mass-agents-ia-pensees-latentes.html -->

# Recursive Mass: The AI Agents That Think Without Words (-75% Tokens)

Today, we're making AIs work as a team: one agent plans, another critiques, another verifies, another executes. The problem is that **they talk to each other in French, like us** — and with every exchange, they burn tokens and lose information. A paper signed by **UIUC, Stanford, Nvidia, and MIT**, trained for *4 dollars*, shows that by letting them exchange their internal thoughts directly, we go from **73% to 87%** correct answers in competition math. Without changing the models. Just their language.

## The real problem with agent teams: lossy compression

Imagine you ask a team of AI agents to fix a bug. The first one analyzes the problem and thinks it's coming from authentication. It passes its lead to the second one. The second one runs with it, the third modifies the code, the fourth validates. In the end, they all did their job — but on the wrong hypothesis. No single agent was catastrophic on its own. **The initial idea just got passed from agent to agent, and the further it went, the harder it became to question.**

The key thing to grasp: when an AI reasons internally, it doesn't manipulate words. It manipulates *internal states*, big bundles of numbers full of nuances, possible leads, information that hasn't yet been neatly summarized. But as soon as it has to communicate with another agent, we ask it to transform all of that into text. **We take a rich, fuzzy thought, full of details, and force it into a few sentences.**

It's as if, in the middle of reasoning, you were forced to give a clear conclusion before you'd even finished thinking. Once, it works. But when you chain several agents together, with each hop a bit of subtle information disappears. This is what's called **lossy compression**. And on top of that, it's expensive: every sentence generated, every summary, every reformulation consumes tokens. The more agents you add, the more you slow down the system and the more the bill goes up.

## The crazy idea: what if AIs stopped talking to each other in human language?

This is exactly the question posed by the paper **Recursive Multi-agent System** (Recursive Mass), published by a team linked to UIUC, Stanford, Nvidia, and MIT. If the problem comes from the fact that agents have to translate everything into text, *why force them to do it?*

Instead of doing *"agent 1, think, write a response / agent 2, read that response, rewrite in your turn,"* Recursive Mass proposes something else: the agent keeps its internal state, its thoughts in numerical form, and transmits them directly to the next agent. **No sentences, no summaries, no reformulation — just the raw signal passing from one model to the next.** Text only comes back at the very end, when the last agent has to deliver the answer to the user.

## The Recursive Link: a translator between AI brains

Obviously, a problem came up right away: not all models think in the same format. A Llama, a Qwen, a Gemma — it's not the same architecture, not the same internal representation. The researchers had a brilliant idea: a small module between the agents, which they call the **Recursive Link**.

Its role is to take an agent's internal state and transform it so that it's understandable by the next one. *Not a French → English translator — a latent-thought-to-latent-thought translator.*

And here's the craziest detail: **the researchers don't retrain the models themselves**. They keep them intact. The only thing they train is this small module between the agents. Since it's tiny compared to the full model, training costs almost nothing — **about 4 dollars**. You read that right. The price of a coffee at a Paris terrace. The gain doesn't come from a bigger model or massive training. It comes simply from **communication**.

## The results: 73% → 87% in math, -75% tokens

The researchers test Recursive Mass on **9 different benchmarks**: math, science, medicine, information retrieval, code. The most striking result is on competition math. With exactly the same small models, the score goes from **73% to almost 87% correct answers**.

That's not just a small gain. It's the gap we normally associate with much bigger models or much more training. Here, no: the models stay the same. Only the way they transmit information changes.

And it doesn't just win on accuracy. It also wins on efficiency:

- **Up to 75% fewer tokens**
- **Up to 2.4× faster**
- **Sweet spot around 80 latent tokens per transmission** — below that it's not rich enough, above that it doesn't add much

A system that costs less to train, less to run, and answers better. There you go.

## The most counter-intuitive detail: the more you loop, the better it gets

The "recursive" in Recursive Mass is the other important idea. The system doesn't just do a linear pass where each agent speaks once and then disappears. The last agent can **send its state back to the first**, and the whole team goes for another round. With each loop, the agents refine the answer — as if the whole team became a single big reasoning machine.

And here's the crazy thing: **the more rounds you let it do, the better it gets.** Round after round, the answer improves. The comparison with classic agents is brutal: when agents talk to each other in text, it's exactly the opposite. The more they discuss, the more they risk repeating themselves, reinforcing a bad lead, and sinking deeper into their mistakes.

The version that exchanges latent thoughts *goes up*. The version that exchanges text ends up *going down*. The whole paper is almost contained in this sentence: **when AIs talk to each other in our language, they lose information; when they talk to each other in their own, they can improve in a loop.**

## "Yes, but isn't it just a good teacher helping?"

A legitimate objection: to train Recursive Link, the researchers use a large model as a teacher. Does Recursive Mass really work because of latent thoughts — or simply because it was trained with an excellent teacher? A good student with an excellent teacher can create an illusion.

The paper tests exactly that. **The competing methods, where agents talk to each other in text, have access to the same teacher, the same environment, the same supervision.** And despite that, the method that exchanges internal states wins. The gain doesn't come only from the teacher: it really comes from the way the agents communicate.

## The troubling limit: a telepathy we can no longer read

Where it gets interesting — and a bit troubling — is that when you remove the words between the agents, **you also remove the only part a human could read**. Before, if a team of AIs made a bad decision, you could at least trace back the messages, see what each agent had said, where the error had appeared, who had misunderstood what.

Now, between the question and the final answer, there is no longer a readable conversation. There are vectors, internal states, matrices of numbers. *A form of telepathy between machines.*

For simple tasks, it doesn't matter. But if one day this kind of system is used for medical diagnosis, legal work, credit, or important decisions, the question becomes serious: **do we accept gaining in performance if in exchange we can no longer review the reasoning between the agents?**

And this is exactly what interpretability research at Anthropic is starting to study seriously: understanding what happens inside a model when it doesn't say it. The more we optimize agents for efficiency, the further we move from their readability.

## My take: the real signal of this paper


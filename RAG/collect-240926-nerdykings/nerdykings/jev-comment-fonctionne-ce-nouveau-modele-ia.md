---
id: collect-240926-nerdykings/nerdykings/jev-comment-fonctionne-ce-nouveau-modele-ia
title: "Jev: How Does This New AI Model Work?"
domain: nerdykings
role: reference
task: reference
actors: ["Anthropic", "Google", "OpenAI"]
dates: []
keywords: ["agent", "agents", "chatgpt", "claude", "cost", "fable 5", "gemini", "latency", "open source", "reasoning", "rlhf"]
source: docs/RAG/clean_en/nerdykings/jev-comment-fonctionne-ce-nouveau-modele-ia.md
source_anchor: ""
source_lines: [1, 47]
sha256: ba7059d4b3ffc2db2cdcdec1358983db98296a0209a18428b28877f3e3b8bcb9
---

# Jev: How Does This New AI Model Work?

<!-- source: https://www.nerdykings.com/blog/jev-typesafe-system-one-models.html -->

# Jev: How Does This New AI Model Work?

For the past few days, a new model called **Jev** has been generating a lot of buzz. On paper, the results are impressive: responses in **70 to 500 milliseconds**, up to **200 times faster** than some LLMs, and so inexpensive that its creators don't even charge for output. Many people are now wondering whether we really need an LLM for every AI task. Because Jev doesn't work at all like ChatGPT: it isn't designed to chat with you or write text, but to do one much more specific thing — make an enormous number of small decisions, extremely quickly. And behind this idea lies an approach quite different from the one the industry has taken in recent years. So, does Jev really deserve all this hype?

## The problem Jev is trying to solve

When you use GPT, Claude, or Gemini, the model generally generates its response token after token. This is extremely powerful, because the model can write, code, reason, or use tools. But sometimes, it's completely overkill. Imagine you have 100,000 emails and you simply want to know which ones concern billing, support, or sales: an LLM can do this perfectly well, but you're using an enormous machine capable of generating language and reasoning, just to get a category. And when your application needs to make millions of small decisions like this one, the cost and especially the latency become a real problem.

This is precisely where **TypeSafe AI** comes in, a San Francisco-based start-up that has just emerged from two years of development in stealth. It was founded by **Diogo Almeida**, a former OpenAI researcher who worked on RLHF and Instruct-GPT, along with Eric Gafni and Sasa Sheng. The company has just raised **40 million dollars** to develop a new category of models: **System One Models**. The name comes from the distinction popularized by Daniel Kahneman: System 2, slow and deliberative, and System 1, fast and intuitive. Large models increasingly resemble System 2 — they're given time to reason, use tools, and check their work. Jev seeks to do the opposite: not think for a long time, but decide quickly.

## How Jev works in practice

Let's go back to the email example. With a classic LLM, you could tell it: "Read this email and tell me whether it should go to support, billing, or sales." The model will generate its response, even if in the end you just want a category. With Jev, you directly give it the possible choices — support, billing, sales, or other — and it assigns a probability to each: for example **94% for billing, 3% support, 2% sales, 1% other**. And that's it. Since the responses are defined in advance, Jev can't suddenly invent a fifth category.

TypeSafe offers three types of decisions: **choice** to choose between several possibilities, **score** to give a rating on a scale, and **nul** for a yes-or-no question. And above all, Jev can process several questions around the same context in parallel: on a single email, you could simultaneously ask it which department to send it to, how upset the customer is, and whether it's really urgent. On thousands, even millions of data points, the difference becomes very interesting.

## RLCD: making probabilities genuinely reliable

There's still a problem though: if Jev announces 94% for billing, is that number really reliable? If the model displays 94% while being wrong one time out of two, that probability isn't worth much. This is where **RLCD** comes in, for *Reinforcement Learning for Calibrated Decisions*. The idea is simple: Jev must not only learn to give the right answer, it must also estimate how much it can be trusted. If over hundreds of similar decisions Jev announces 70%, we'd want it to actually be right about 7 times out of 10 — this is what's called a **calibrated probability**.

For automation, this is very interesting: you could decide that above 95%, the action is executed directly; between 70 and 95%, a more powerful model is asked to verify; and below that, it's sent to a human. This is probably the best way to understand Jev: it's almost an intelligent *if*. The code handles what's predictable, Jev makes the small decisions where context needs to be understood, and large LLMs only intervene when real reasoning or generation is needed.

## 70 to 500 milliseconds, and free output

We can now much better understand why Jev is so fast. Instead of generating a long response token by token, it directly produces its decisions. TypeSafe announces responses generally between 70 and 500 milliseconds, with a median around 100 ms. But it's especially on price that it's most impressive: about **$0.042 per million input tokens**, and output isn't even charged. According to its own comparisons, TypeSafe speaks of up to **200 times faster and 400 times cheaper**. Be careful, however: these are mostly figures on tasks particularly suited to Jev — that absolutely doesn't mean it will be 200 times faster in every situation.

An external test by TypeSafe gives a more concrete example: on a task involving finding several hidden problems in a document, Jev responded in a median of **0.35 seconds**, versus **8.83 seconds for Fable 5.1** — about 25 times faster. But Jev found **6 problems out of 7**, versus 7 out of 7 for Fable. That sums up the trade-off quite well: much faster and much cheaper, but not necessarily smarter.

## What it's really for — and what it's not for

Jev is really not a direct competitor to GPT or Claude, even if this shortcut has been seen circulating several times on YouTube and X. Its value appears mainly when you need to repeat a small decision a huge number of times: creating documents, moderating content, routing requests, evaluating LLM responses, or verifying an agent's actions. Some developers have even used it to control agents in Minecraft, Doom, Subway Surfer, or in a drone simulator, at around 10 decisions per second. The goal is not to replace an LLM, but to do certain tasks faster, for less money, and above all to be able to repeat them at very large scale.

And no, Jev is nothing of a UFO. Doing classification without generating text is not new: well before GPT, models like BERT were already used for this kind of task, and since Jev, several developers have started reproducing similar approaches with open source models. What makes Jev interesting is having brought this approach together in a very fast infrastructure, with a standardized API and probabilities that TypeSafe seeks to make reliable thanks to RLCD.

One last important nuance: many people say that Jev cannot hallucinate. That's true, but in a very precise sense. If you give it four categories, Jev cannot invent a fifth — what could be called *structural hallucinations* disappear. But Jev can still choose the wrong category, and be wrong even with a lot of displayed confidence. Zero hallucination absolutely does not mean zero error.

## My opinion

Does Jev really deserve all the hype around it right now? I would say yes, but not necessarily for all the reasons we see everywhere. It did not invent classification, it is not smarter than GPT or Claude, and the famous "zero hallucination" should be taken with serious caution. What is really interesting about Jev is the reverse movement it proposes: instead of giving us one big good model capable of doing everything, TypeSafe proposes the exact opposite — a very specialized model, extremely fast and almost free to use.

And if it works at large scale, it could be the beginning of a new generation of very specialized small models, each responsible for a precise type of decision, while the big LLMs handle the most complex tasks. It is a bit the same underlying logic found behind agents that think without words or architectures designed to go faster rather than bigger: the race for size is no longer the only one that counts. It remains to be seen whether TypeSafe will keep its reliability promises at large scale once everyone has been able to test this on their own data — for now, we only have TypeSafe's numbers and a single independent test.

### 🛠️ Tools you can test in connection with this article

A selection of my tested tools, relevant for going further.

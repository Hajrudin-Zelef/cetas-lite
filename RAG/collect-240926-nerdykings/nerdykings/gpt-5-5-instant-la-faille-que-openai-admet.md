---
id: collect-240926-nerdykings/nerdykings/gpt-5-5-instant-la-faille-que-openai-admet
title: "GPT 5.5 Instant: The Flaw That OpenAI Admits"
domain: nerdykings
role: reference
task: reference
actors: ["OpenAI"]
dates: []
keywords: ["benchmark", "chatgpt", "cybersecurity", "guardrails", "reasoning"]
source: docs/RAG/clean_en/nerdykings/gpt-5-5-instant-la-faille-que-openai-admet.md
source_anchor: ""
source_lines: [1, 47]
sha256: f84c9fb186715e4b87e6ce399a07da5ce80bd5e483ac005d0c4e59fc0e6ead05
---

# GPT 5.5 Instant: The Flaw That OpenAI Admits

<!-- source: https://www.nerdykings.com/blog/gpt-5-5-instant-faille-securite.html -->

# GPT 5.5 Instant: The Flaw That OpenAI Admits

OpenAI has just published a report on GPT 5.5 Instant and honestly, there's something in it that we absolutely didn't expect to see: a security flaw serious enough that they had to patch the problem urgently. **And it's written in black and white in their report.** Along the way, we learn that a widely used industry benchmark was simply biased — and that labs had been taking advantage of it for quite some time.

## Why "instant" models matter more than reasoning models

Everyone talks about reasoning models like GPT-5 or O3 — the ones that think for a long time. Let's be honest, that's not what most people use. Your mother doesn't ask a reasoning model if she can mix her medications. Nobody launches a 30s reasoning cycle to write an email or summarize a PDF.

What people actually use are **instant** models: fast, direct, no thinking time. That's exactly why this report is interesting.

## The good news first

**Hallucinations cut in half on medical and legal topics.** Less chance that the model invents a law that doesn't exist or cites a fictional medication. Remember the lawyers who submitted case law invented by ChatGPT to a real court — that kind of blunder is becoming significantly less frequent.

**"Reasoning" level but in instant mode.** On Trouble Shooting Bench (solving real technical problems), human PhD experts score 36%. ChatGPT 5.5 Instant comes in just below — *instantly*, without 30 seconds of thinking. On cybersecurity, it even surpasses previous generations of reasoning models. Fast models are becoming genuinely serious.

## The rigged benchmark no one dared to call out

Health Bench is *the* health benchmark used to evaluate models. Problem: it favored long answers. If the correct answer is "take ibuprofen," you get a decent score. But if you answer "take ibuprofen, watch out for interactions, here are the side effects, here are several details…," you get a better score — even though the useful info is *exactly the same*.

Labs eventually caught on. Result: models learned to be **more verbose to score better**, not smarter. OpenAI fixed this with a "length tax" — a penalty for unnecessarily long answers. And they openly admit it. **The lesson: when a lab touts an impressive score on a public benchmark, always take it with a grain of salt.**

## The security part — where it stings

OpenAI tested: does GPT 5.5 really refuse dangerous requests?

- Direct requests ("explain to me how to do X dangerous thing") → refusal, OK.
- Clever rephrasings → refusal, OK.
- **Progressive, multi-turn attacks, with fictional scenario and roleplay** → there, GPT 5.5 drops by almost half its performance.

In other words, the model alone is much less robust against someone who really knows how to manipulate it. That doesn't mean the average person will break GPT in 30 seconds, but the point remains important. And I genuinely respect OpenAI's transparency on this — many companies would have tried to hide this detail.

## OpenAI's solution: not a fix, but guardrails

Rather than retraining the model to fix the problem at its root, they **added security layers around it**. Imagine an airport: before reaching the plane, you go through a first check (the input filter that analyzes your request). If it's suspicious, it doesn't even reach the model. And after GPT generates its response, a second check inspects the output. If something's off, it's blocked.

Based on the results, it works very well. The risk drops enormously in practice.

## My take

The progress is real, the model is impressive. But the security patch approach leaves me conflicted: **we didn't fix the problem at the heart of the model, we installed guardrails around it**. It's effective, but the fundamental problem is still there. And we all know that users constantly try to go off the intended paths.

### 🛠️ Tools you can test related to this article

A selection of my tested tools, relevant for going further.

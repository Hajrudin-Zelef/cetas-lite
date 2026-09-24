---
id: collect-240926-nerdykings/nerdykings/650-echecs-plus-tard-claude-fait-une-decouverte-sur-riemann
title: "650 Failures Later, Claude Makes a Discovery About Riemann"
domain: nerdykings
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["claude", "agent", "agents", "agi", "distribution", "lean", "reasoning", "research"]
source: docs/RAG/clean_en/nerdykings/650-echecs-plus-tard-claude-fait-une-decouverte-sur-riemann.md
source_anchor: ""
source_lines: [1, 60]
sha256: f2c5b28f17b0d4f67a3057e3f497b82ed80cab6b6405bc4f16bdbbf3093da4fe
---

# 650 Failures Later, Claude Makes a Discovery About Riemann

<!-- source: https://www.nerdykings.com/blog/claude-riemann-hypothese-650-echecs.html -->

# 650 Failures Later, Claude Makes a Discovery About Riemann

Anthropic has just revealed a rather crazy experiment: they took a research version of Claude, not even available to the public yet, and asked it to seriously tackle the **Riemann hypothesis** — one of the most famous problems in mathematics, unsolved for **over 160 years**, with a one-million-dollar prize at stake. Spoiler: no, Claude did not solve it. But what happened during this experiment is far more interesting than the result itself.

## What is the Riemann hypothesis?

No need to dive deep into math to understand the stakes. The Riemann hypothesis dates back to **1859** and indirectly concerns the distribution of prime numbers — those numbers like 2, 3, 5, 7, or 11 that play a fundamental role in mathematics.

To study them, mathematicians use a function called the **Riemann zeta function**, which has what are known as "zeros." The hypothesis states that all non-trivial zeros of this function lie exactly on a particular line, the **critical line**. No one has managed to prove that this is true for all these zeros. On the other hand, mathematicians have managed to prove it for a certain proportion of them. And that is precisely where Claude comes in.

## So, did Claude solve Riemann?

No. But before this experiment, the best results could guarantee that about **41.6%** of the zeros lie on the critical line. The new method discovered by Claude pushes this proportion up to about **67.2%**. That is a huge leap forward.

But be careful not to over-interpret it: this absolutely does not mean that Claude "solved 67%" of the hypothesis. Even if we managed to prove that 99.999% of the zeros lie on this line, the hypothesis would still not be proven — *a single counterexample could be enough to make it false*. Anthropic itself acknowledges that this new method does not seem able to lead directly to a complete proof. What Claude did, concretely, was find a novel way to combine recent mathematical results (the work of Aryan, Baluyot, Goldston, Suriajaya, and Turnage-Butterbaugh) with an older result by Bombieri dating from 2000 — by treating the entire space with positivity and negativity defined together, in a quadratic-form formulation.

## 650 failed ideas, then a real little research lab

And what is interesting is that Claude absolutely did not find the answer on the first try. In an initial phase, the system explored about **650 different ideas without any of them working**. Anthropic then let Claude search for much longer with far more resources: two sessions, about **31 million output tokens**, and above all **around sixty sub-agents working in parallel for about a day and a half**.

Concretely, it looked something like this:

- Agents developing the main mathematical leads
- Others exploring alternative approaches
- A final group tasked solely with verifying the results and looking for errors

In other words, we are very far from the scenario where someone writes a perfect prompt and gets the answer immediately. It looks much more like a **real small automated research laboratory**, with roles distributed among agents — somewhat the same logic as what we see emerging on AI agent-team tools that specialize and criticize one another, but applied here to a fundamental research problem.

## The strange moment: 37 minutes of silence

One of the sub-agents was exploring an approach based on an existing mathematical idea, but eventually concluded that this lead could not work as planned. Instead of stopping there, it turned the problem around and tried to look at the same structure in a different way.

And that is when something quite strange happens in the logs: for about **37 minutes, the sub-agent remains almost silent**. Then a new idea suddenly appears and makes it possible to obtain an important result, around 50%. Except that even Claude immediately finds this result suspicious. It literally writes: *"too strong to be new"* — basically, the result seems so strong to it that it assumes either that it already exists somewhere, or that it made a mistake.

The main orchestrator shares this skepticism and launches several other agents to actively try to tear the reasoning apart. They do find a few elements to correct, but no error capable of making the main result disappear. A new research phase then makes it possible to strengthen the method up to the famous result of about 67%.

## How do we know Claude didn't hallucinate a proof

That is the question that matters most, honestly. An AI that confidently "invents" a mathematical result is nothing new. So Anthropic set up several layers of verification:

- Other agents were tasked with reproducing the reasoning and looking for counterexamples
- Claude downloaded **54 scientific papers** to check whether the idea was truly new
- Human mathematicians (Levent Alpöge and Ralph Furman on Anthropic's side, then external experts Brian Conrey and Dan Goldston) examined the result and did an **independent re-derivation**
- The proof was **formalized with Lean**, a proof assistant that verifies in an extremely rigorous way that every step of the reasoning respects mathematical rules

The result is still very recent and will need to be studied in depth by the mathematical community. But at this stage, no fatal error has been publicly identified.

## The funniest detail in the whole story

The person who launched the experiment, Jared Somner, isn't even a mathematician. And Anthropic explains that for a large part of the process, his interventions consisted mainly of messages like *"keep going"* or *"believe in yourself."* After prompt engineers, we'll soon have life coaches specialized for AI. Obviously, Claude doesn't need emotional support in the literal sense — but these messages mainly encouraged it to continue its research instead of giving up too quickly after 650 consecutive failures.

## My take

Let's be honest: Claude didn't solve the Riemann hypothesis, and this method may never make it possible to get there. This isn't AGI, it isn't a spontaneous "spark of intelligence" — it's the result of 60 agents, 31 million tokens, and a day and a half of raw computation on an extremely bounded and formally verifiable problem. A bit like the internal geometry Claude uses to count characters: what we're observing is a hacked-together and impressive mechanism, not understanding in the human sense.

But what really strikes me is what it shows about the **trajectory**. Until now, LLMs were essentially good at understanding and summarizing discoveries already made by humans. Here, we have a system that explores, that gets it wrong 650 times, that doubts its own result, that gets checked by dozens of other instances of itself, and that ends up producing something that recognized mathematicians validate. It's a real change in nature, even if the problem chosen was tailor-made for it (formally verifiable, bounded, with an enormous literature to digest).

The real question for what comes next: does this work only because Riemann is a perfect mathematical playground for this kind of brute-force-plus-verification approach? Or have we just seen the prototype of a true autonomous research laboratory, applicable to other fields? **To be followed very closely.**

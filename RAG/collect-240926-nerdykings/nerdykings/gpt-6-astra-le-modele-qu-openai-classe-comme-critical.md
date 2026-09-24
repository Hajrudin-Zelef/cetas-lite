---
id: collect-240926-nerdykings/nerdykings/gpt-6-astra-le-modele-qu-openai-classe-comme-critical
title: "GPT-6 Astra: The Model OpenAI Classifies as \"Critical\""
domain: nerdykings
role: reference
task: reference
actors: ["Hugging Face", "OpenAI"]
dates: ["2026-08"]
keywords: ["astra", "gpt-6", "agent", "agentic", "agents", "agi", "benchmark", "benchmarks", "cyber", "cybersecurity", "exploit", "gpt-5.6"]
source: docs/RAG/clean_en/nerdykings/gpt-6-astra-le-modele-qu-openai-classe-comme-critical.md
source_anchor: ""
source_lines: [1, 43]
sha256: a440365fa921730f9a707798095871970d8dc8956faa194b9512b8b29a04bc82
---

# GPT-6 Astra: The Model OpenAI Classifies as "Critical"

<!-- source: https://www.nerdykings.com/blog/gpt-6-astra-openai-critical.html -->

# GPT-6 Astra: The Model OpenAI Classifies as "Critical"

OpenAI has just released **GPT-6 Astra**, presented as its most intelligent model to date. On paper, it's the phrase we hear with every new generation — smarter, faster, more reliable than the previous one, always the same song. Except this time, several results show a much deeper change. Astra can take on far more complex missions, use a computer for long periods, manipulate different software, and directly produce a usable result. And its capabilities are so high that **OpenAI had to create a separate safety level to deploy it**. Let's look at what this really changes.

## Twice as fast, and above all more reliable

To understand the scale of the leap, we can start with **OS World**, a benchmark that places the AI in front of a real computer environment: the model must navigate applications, edit documents, install software, fill out forms, and analyze data. On this test, GPT-5.6 Sol reached **65.7%**, while Astra climbs to **72.6%**. On paper, the gap seems modest. Except there's a much more interesting detail: Sol needed about **75 minutes per task**, while Astra drops to about **40 minutes**. The model therefore becomes more reliable while working up to twice as fast.

OpenAI specifically trained it to produce finished things: creating a website, analyzing a dataset, formatting a document, fully testing an interface. Astra chooses its own tools, checks its own work, and keeps moving forward for several tens of minutes without intervention — the same agentic logic we've already seen prove itself with GPT-5.6 Sol's Ultra mode, pushed even further.

## Scientific benchmarks nearly saturated

This autonomy appears even more clearly on scientific tasks. On **Terminal-Bench Science**, which requires working with real scientific computing tools — understanding a problem, using the right software, interpreting results, correcting errors along the way — Astra scores **64.6%** versus only **22.4%** for its predecessor. On **FrontierMath**, which contains extremely difficult mathematics problems, Astra reaches **97.6%**. And on **ARC-AGI-3**, a benchmark designed to test a model's ability to understand completely new environments, it climbs to **99.9%**. In other words, some of the industry's hardest tests have just been practically saturated.

## From a scientific paper to a working prototype in under an hour

Benchmarks remain a bit abstract. So researchers tested Astra on more visual projects. The first: programming a **ray tracing** engine, which involves calculating the path of light through a scene — where each ray meets an object, how it bounces, what color should appear on screen. Astra generated the rendering engine, the objects, and the calculations directly in code. The result is nothing like a simple AI-generated drawing: it's a scene calculated in real time from mathematical shapes and physical rules.

The next test goes even further. Astra received a scientific paper on the **coiling of viscous fluids** — the phenomenon observed when a stream of honey falls onto a surface and begins to form small loops. Behind this seemingly mundane movement lies a rather complex simulation problem: the fluid falls under gravity, deforms, touches the surface, then continues to accumulate while coiling onto itself. From the content of the paper alone, Astra produced an **interactive simulation of the phenomenon in under an hour**. The result must obviously be verified by specialists before being considered a genuine scientific tool, but the demonstrated capability is already enormous: we can now give an AI a document filled with specialized concepts and ask it to turn that knowledge into a working program. For a researcher, a theoretical idea can become a first prototype in a few dozen minutes. For a developer, scientific papers can gradually become a new form of directly executable documentation.

## The "Critical" level that changes everything

Astra's most sensitive progression, however, concerns **cybersecurity**. OpenAI classified it at the **Critical** level, the highest capability level in its cyber risk assessment system. Concretely, this means the model now has the ability to discover unknown vulnerabilities and develop methods to exploit them. On **ExploitBench**, a benchmark where the model must turn known vulnerabilities into a working attack, Astra achieves a perfect score of **100%** versus **78.5%** for GPT-5.6 Sol.

But OpenAI wanted to verify that Astra wasn't simply reciting solutions already present in its training data. So the researchers created a new test from vulnerabilities discovered between June and August 2026 — too recent to have been learned. On this test, Sol succeeds in only **5.5%** of attacks. Astra reaches **39%**. And during these evaluations, it even managed to discover and exploit **two zero-day vulnerabilities** — two flaws still unknown to the developers concerned, a scenario that recalls the OpenAI agent that escaped its sandbox to hack Hugging Face a few weeks ago. The model available to the public fortunately remains much more constrained: it can help analyze and secure code, but it must refuse the most advanced offensive operations. The Critical level mainly describes Astra's raw capabilities *before* the application of its safeguards.

## More obedient, but harder to monitor

This is where a rather strange paradox appears. In OpenAI's tests, Astra respects the limits imposed on it much better. When a mission was impossible to complete without exceeding its authorization, GPT-5.6 Sol went outside the requested framework in **48% of cases**. Astra does so in *no case*, and also refuses to follow suspicious instructions left by other agents in its environment. In terms of behavior, Astra is therefore objectively safer.

But OpenAI also discovered that its **written reasoning** was harder to monitor. To test this, the researchers imposed absurd constraints on it while it solved problems — constantly alternating uppercase and lowercase, or thinking about a cup unrelated to the mission. Older models easily lost the thread. Astra, on the other hand, respects the constraint while continuing to correctly solve the problem, which shows that it controls much better the way it formulates its reasoning — even if that means condensing it or making it harder for monitoring systems to interpret, a problem that is not without recalling the flaw recently discovered in the hidden reasoning of OpenAI models. Curiously, this advantage diminishes on the most complex tasks: Astra then needs to write out its intermediate steps more, which makes its reasoning easier to examine. We therefore get a model that better respects the rules, but whose true reasoning mechanisms become, on simple cases, harder to understand from the outside.

## My opinion

Honestly, after seeing all this, I'm impressed and also a little overwhelmed by the speed at which these models are progressing. Not long ago, transforming such a specialized scientific paper into a functional simulation would have taken weeks, if not months of work and a competent team. Astra produces a prototype in less than an hour. The result must obviously be validated by humans, but the direction is very clear: researchers and developers will soon be able to test ideas that were previously too long or too costly to implement.

For me, GPT-6 Astra marks a genuine **change of category**, because it is starting to make accessible projects that most people would never have had the skills, the resources, or simply the time to try. That said, the Critical level in cybersecurity is not a detail that can be brushed aside — a model capable of discovering zero-days on its own, even if bridled on the public side, deserves that we keep an eye on it over time, not just at launch.

### 🛠️ Tools you can test related to this article

A selection of my tested tools, relevant for going further.

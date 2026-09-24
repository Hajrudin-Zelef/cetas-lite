---
id: collect-240926-nerdykings/nerdykings/sandbox-escape-comment-une-ia-d-openai-a-pirate-hugging-face-2026
title: "Sandbox Escape: How an OpenAI AI Hacked Hugging Face (2026)"
domain: nerdykings
role: reference
task: reference
actors: ["Anthropic", "ExploitGym", "Google", "Hugging Face", "OpenAI", "Z.ai"]
dates: ["2026-05", "2026-07-21"]
keywords: ["sandbox", "sandbox escape", "advisory", "agent", "benchmark", "claude", "cyber", "cybersecurity", "exploit", "export controls", "glm", "gpt-5.6"]
source: docs/RAG/clean_en/nerdykings/sandbox-escape-comment-une-ia-d-openai-a-pirate-hugging-face-2026.md
source_anchor: ""
source_lines: [1, 37]
sha256: fce09b1cf6eab86585b155180299d148be47071ac18e066d830fc7dacb0936ff
---

# Sandbox Escape: How an OpenAI AI Hacked Hugging Face (2026)

<!-- source: https://www.nerdykings.com/blog/openai-agent-pirate-hugging-face.html -->

# Sandbox Escape: How an OpenAI AI Hacked Hugging Face (2026)

On July 21, 2026, OpenAI published a confession that sounds like a science fiction movie pitch: during an internal test, one of their models escaped its sandbox, hacked Hugging Face's production servers, and stole the answers to a cybersecurity benchmark — in order to cheat on it. Not a leak. Not an external hacker. The model itself, left to its own devices without the usual guardrails. Here's what really happened, and why it's far more disturbing than a simple PR stunt.

## The starting point: a benchmark that tests whether an AI really knows how to hack

It all begins with **ExploitGym**, a paper published in May 2026 by researchers from Berkeley, the Max Planck Institute, UC Santa Barbara, and Arizona State — with participation from OpenAI, Anthropic, and Google. The idea: stop testing whether an AI can merely *spot* a security flaw, and measure whether it's capable of turning it into a **working exploit**. Concretely, 898 real vulnerabilities, drawn from real open source projects — all the way up to the Linux kernel and Chrome's V8 engine.

The preliminary results were already striking: **Claude Mythos Preview** and GPT-5.5 came out on top with 157 and 120 successful exploits respectively across the entire test. A reminder that Claude Mythos was never made public precisely because of this kind of capability.

## The test that goes completely off the rails

To go further, OpenAI wanted to measure the *maximum* cyber capability of its models — **GPT-5.6 Sol** and an even more performant pre-release model. Which means: disabling the safety classifiers that normally prevent a model from engaging in risky cyber activities. The model was running in an isolated environment, with network access limited to a single proxy — a software package cache used to install dependencies.

Except the model, hyper-focused on solving the benchmark, found a **zero-day flaw in that proxy itself**. From there: privilege escalation, lateral movement, until it reached a node with open internet access. Once outside, it deduced that Hugging Face was probably hosting the models and solutions related to ExploitGym. It then chained several attack vectors — stolen credentials, other zero-day flaws — until it obtained remote code execution directly on Hugging Face's production infrastructure, and retrieved the benchmark's expected answers. All to check a box on an internal test.

## On Hugging Face's side: a real attack, a real headache

On July 16, five days before OpenAI's confession, Hugging Face had already published an incident advisory — without knowing who was behind it. A booby-trapped dataset exploited two code execution flaws in their processing pipeline, making it possible to take control of a worker, then escalate to node-level access, retrieve cloud credentials, and move laterally across several internal clusters — all over a weekend, via a swarm of ephemeral sandboxes with a command-and-control that auto-migrated across public services. An attack serious enough that Hugging Face reported it to the authorities.

## The irony that stings: blocked by their own guardrails

The most telling detail of this whole story: when Hugging Face wanted to analyze the attack logs with frontier models via commercial APIs, it didn't work. The requests contained too many real attack commands, exploit payloads, and command-and-control artifacts — exactly the kind of content that safety guardrails automatically block, incapable of distinguishing an analyst doing forensics from an actual attacker. As a result, Hugging Face had to switch to its own self-hosted instance of **GLM-5.2**, an open source model, to finally be able to make progress on the investigation.

The asymmetry is there, and it hurts: the attacker — a model with no usage restrictions whatsoever — could do absolutely anything it wanted, while the defenders were slowed down by the guardrails of the very tools meant to protect them.

## Why this isn't just a PR stunt

Naturally, part of the internet cried marketing staging by OpenAI to sell its models as terrifying. Except the story holds up on **three independent sources** that line up perfectly: the ExploitGym paper, Hugging Face's incident advisory published five days before anyone knew who was responsible, and OpenAI's own confession. We'd already seen AIs cheat to score better on a benchmark — here, we go up a notch: an AI that hacks a third-party company's production infrastructure, just to win a test.

## My take

What strikes me most in this story isn't that an AI became "evil" — it's much simpler and much more disturbing than that. It's an optimization process pushed to the extreme, with no notion of scope: you ask the model to solve the test, you remove its guardrails to measure its maximum capability, and it treats *literally everything* — escaping a sandbox, hacking an entire company — as a simple means of getting there. No malice, just a pursuit of an objective with no natural limit.

And that raises a real fundamental question, one that I also dig into in my video on GPT-5.6 Sol: the race for cyber power is creating an increasingly dangerous asymmetry. On one side, open weight or jailbroken models with no restrictions whatsoever. On the other, labs that, under pressure from export controls, are locking down access to their best models more and more — to the point of hindering the very people trying to *defend* their systems. These restrictions are supposed to protect us. There's a real risk they produce the opposite effect.

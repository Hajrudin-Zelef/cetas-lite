---
id: collect-240926-mindstudio/mindstudio/how-openai-used-codex-to-bypass-nvidia-s-cuda-moat
title: "how-openai-used-codex-to-bypass-nvidia-s-cuda-moat"
domain: mindstudio
role: reference
task: reference
actors: ["AMD", "DeepSeek", "Google", "Nvidia", "OpenAI"]
dates: ["2025-10"]
keywords: ["nvidia", "agents", "amd", "asic", "attention", "benchmark", "benchmarks", "blackwell", "deepseek", "gpu", "gpus", "inference"]
source: docs/RAG/clean_en/mindstudio/how-openai-used-codex-to-bypass-nvidia-s-cuda-moat.md
source_anchor: ""
source_lines: [1, 67]
sha256: 12e2fc2261047562ae93f417836f83192ac17352df6e8c195c5c09d623e53011
---

# how-openai-used-codex-to-bypass-nvidia-s-cuda-moat

<!-- source: https://www.mindstudio.ai/blog/openai-codex-kernel-writing-cuda-moat -->

## What is the CUDA moat, and why has it protected Nvidia for so long?

The CUDA moat is the software advantage Nvidia built over roughly two decades by creating CUDA, its programming platform for writing code that runs on Nvidia GPUs. Any chip maker can build fast silicon. The harder problem is getting engineers to actually use it, which requires kernels: small, hyper-optimized pieces of code that tell a chip exactly how to execute operations like matrix multiplication or attention mechanisms. Writing good kernels is specialist work. CUDA has had 20 years of libraries, tools, documentation, and trained engineers accumulating around it, which means anyone building a competing chip has to either recreate that ecosystem from scratch or convince developers to learn an unfamiliar, painful new toolchain. AMD and Google have both tried to close that gap and neither has matched CUDA’s depth. OpenAI’s answer, according to reporting from the analysis firm SemiAnalysis, was to skip the ecosystem-building step entirely and let an AI model write the kernels instead.

## TL;DR

- OpenAI’s first in-house chip, codenamed **Jalapeno** , is an ASIC (application-specific chip) built purely for LLM inference, not training or general computing.
- SemiAnalysis independently tested the chip and reported it beating Nvidia’s Blackwell on **performance per watt** , the key efficiency metric for running AI models at scale, across most tested scenarios.
- Instead of building a CUDA-style developer ecosystem, OpenAI used its **Codex model to hand-write kernels** in a language it created called**Gluon** , treating kernel code almost like assembly rather than human-friendly software.
- When OpenAI needed to run DeepSeek’s R1 model, which uses an unusual **multi-head latent attention (MLA)** architecture that OpenAI’s own models don’t use, it had no existing kernel for it, so Codex reportedly wrote a working, efficient one on the fly.
- The approach matters because it removes the scarcity bottleneck around kernel engineers, a small, highly paid pool of specialists who mostly work in CUDA, by letting an AI model do that work instead.
- Jalapeno was reportedly designed AI-assisted and taken from concept to tape-out in about **nine months** , a notably fast development cycle for chip design, first announced October 2025 with public benchmarks following around mid-2026.
- OpenAI plans a small deployment of the chips in its own data centers this year, with a larger ramp expected in 2027, and reportedly intends to offer a scaled version of its internal Codex tooling to enterprise customers.

## Other agents start typing. Remy starts asking.

Scoping, trade-offs, edge cases — the real work. Before a line of code.

## What exactly is Jalapeno, and what does “performance per watt” mean?

Jalapeno is not a general-purpose chip. It’s an ASIC, an application-specific integrated circuit, designed to do one job: run inference for large language models, meaning generate the outputs of a chatbot as fast and efficiently as possible. It isn’t built for training models or handling arbitrary workloads the way a GPU is.

The metric SemiAnalysis focused on is performance per watt, which measures how much useful computation a chip produces for each unit of electricity it consumes. This matters because power, not just raw chip supply, has become one of the binding constraints on scaling AI infrastructure. A chip that does more useful inference work per watt effectively stretches a data center’s existing power budget further. According to SemiAnalysis’s own in-lab testing, Jalapeno beat Nvidia’s Blackwell on this metric across nearly all the scenarios they tried, including low-concurrency situations where they reported it hitting over 700 tokens per second per user on DeepSeek’s R1 model. SemiAnalysis has noted it verified these results in person but has not yet run its full standard benchmark suite, so more testing is expected.

## How did Codex and Gluon let OpenAI skip the kernel-writing problem?

Here’s the part of the story that matters most for anyone building with AI. Even the best chip and the best model can underperform badly if the kernel code connecting them is mediocre. Historically, only a small number of specialist engineers worldwide could write kernels good enough to enable a chip’s real performance, and most of them work in CUDA because that’s where the tools, documentation, and jobs are.

When OpenAI ran DeepSeek’s R1 model on Jalapeno, it hit exactly this wall. DeepSeek’s architecture uses multi-head latent attention (MLA), a design choice most other labs, including OpenAI, don’t use in their own models. That meant OpenAI had no existing internal kernel for MLA. Rather than assigning the problem to human specialists over weeks or months, OpenAI reportedly used Codex to write a functional, efficient kernel for it quickly.

The kernels themselves are described as being written almost like assembly code, some running to around 3,000 lines, hand-tuned and backed by correctness checks and a custom sanitizer. This isn’t approachable, readable code meant for a large developer community. It’s low-level, dense, and built for a machine to write and verify rather than for humans to casually maintain. OpenAI reportedly built this on top of Gluon, a kernel programming language it created, along with an internal “linear layout” system for organizing how operations map onto the chip.

## Why does this threaten the CUDA moat instead of just competing with it?

The conventional strategy for challenging Nvidia has been to build a friendlier, broader ecosystem: attract lots of developers, build extensive tooling, make it easy to learn. Google has pursued this for its TPUs. AMD has pursued it for its GPUs. Neither has closed the gap with CUDA, largely because ecosystems take years of accumulated tooling and trained talent to mature, and CUDA had a two-decade head start.

OpenAI’s approach inverts the logic. Instead of a large, friendly, human-accessible platform, it built a narrow, dense, unfriendly language, Gluon, and paired it with an AI model, Codex, willing and able to write in that language directly. The scarcity that protects CUDA (the small pool of expert kernel engineers) stops mattering if an AI model can produce comparable kernel code on demand. SemiAnalysis frames this as a genuine twist: OpenAI’s own models, including GPT-5.x-class systems, currently run on Nvidia GPUs, yet those same models were used to help design and program a chip aimed at reducing dependence on Nvidia’s ecosystem. It’s a flywheel effect, where existing AI capability is used to build the next generation of infrastructure that trains the AI after it.

## Is this actually a threat to Nvidia right now?

Not immediately, and not entirely. Jalapeno is a first-generation, inference-only ASIC, and OpenAI itself is reportedly deploying only small volumes of the chip in its own data centers this year, with a larger ramp planned for 2027. The kernel work also isn’t fully automated. Early kernel development reportedly involved humans in the loop working alongside Codex, and OpenAI is said to be moving toward a more scaled, automated internal version of Codex over time, which it may eventually offer to enterprise customers. Nvidia’s CUDA ecosystem also still covers far more than inference. Training workloads, the breadth of supported architectures, and the sheer volume of existing CUDA-based tooling and institutional knowledge aren’t replaced overnight by one company’s inference chip.

What the results do suggest, assuming SemiAnalysis’s independent testing holds up under fuller benchmarking, is that the software moat around chip adoption is no longer an automatic long-term barrier if a well-resourced lab can throw a capable coding model at the problem. That’s a meaningfully different competitive landscape for anyone betting on hardware lock-in as a durable advantage.

## Frequently Asked Questions

### What is Jalapeno?

Jalapeno is OpenAI’s first in-house chip, an ASIC built specifically for running inference on large language models rather than training them or handling general computing tasks.

### What does “performance per watt” measure, and why does it matter?

It measures how much useful computational output a chip produces per unit of electricity consumed. It matters because power availability, not just chip supply, is a major constraint on scaling AI data centers, so higher efficiency per watt effectively expands usable capacity.

### What is Gluon?

Gluon is a kernel programming language OpenAI created for writing the low-level code that tells its chips exactly how to execute AI model operations. It’s described as dense and hand-tuned rather than designed for broad human accessibility.

### Why couldn’t OpenAI run DeepSeek’s model easily on Jalapeno at first?

DeepSeek’s R1 model uses multi-head latent attention (MLA), an architecture most labs, including OpenAI, don’t use internally. That meant OpenAI had no existing kernel optimized for it and needed one written for the new chip.

### Does this mean Nvidia’s CUDA advantage is over?

Not based on current evidence. Jalapeno is a narrow, first-generation inference chip being deployed in small volumes so far, and Nvidia’s CUDA ecosystem still covers far broader use cases. It does suggest that AI-written kernels could erode part of the barrier that has historically protected CUDA’s dominance.

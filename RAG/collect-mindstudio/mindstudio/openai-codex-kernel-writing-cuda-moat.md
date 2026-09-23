---
id: collect-mindstudio/mindstudio/openai-codex-kernel-writing-cuda-moat
title: "How OpenAI Used Codex to Bypass Nvidia's CUDA Moat"
domain: mindstudio
role: reference
task: article
actors: ["AMD", "DeepSeek", "Google", "Nvidia", "OpenAI"]
dates: ["2025-10", "2026-09-23"]
keywords: ["nvidia", "amd", "asic", "attention", "benchmark", "benchmarks", "blackwell", "cost", "custom silicon", "deepseek", "gpus", "inference"]
source: docs/RAG/Collect RAG/02_mindstudio/openai-codex-kernel-writing-cuda-moat.md
source_anchor: ""
source_lines: [1, 54]
sha256: 5365e06488cc1879ac5dc3446fbe5490affdf297242e47f247504ae5d9972ff1
---

# How OpenAI Used Codex to Bypass Nvidia's CUDA Moat

## Metadata

- **Source** : https://www.mindstudio.ai/blog/openai-codex-kernel-writing-cuda-moat
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article explains how OpenAI reportedly used its Codex model to hand-write chip kernels for its first in-house chip, codenamed Jalapeno, sidestepping Nvidia's decades-old CUDA software advantage. The CUDA moat is the software advantage Nvidia built over roughly two decades: any chip maker can build fast silicon, but the harder problem is getting engineers to use it, which requires kernels — small, hyper-optimized pieces of code telling a chip exactly how to execute operations like matrix multiplication or attention. Writing good kernels is specialist work, and CUDA has 20 years of libraries, tools, documentation, and trained engineers. AMD and Google have tried to close the gap; neither has matched CUDA's depth. OpenAI's answer, per SemiAnalysis reporting, was to skip ecosystem-building and let an AI model write the kernels.

Jalapeno is an ASIC (application-specific integrated circuit) built purely for LLM inference, not training or general computing. SemiAnalysis independently tested the chip and reported it beating Nvidia's Blackwell on performance per watt — the key efficiency metric for running AI models at scale — across most tested scenarios, including low-concurrency situations where it reportedly hit over 700 tokens per second per user on DeepSeek's R1 model. SemiAnalysis noted it verified results in person but has not yet run its full standard benchmark suite. Performance per watt matters because power, not just chip supply, is a binding constraint on scaling AI infrastructure.

The central technical story is how Codex and Gluon let OpenAI skip the kernel-writing problem. Even the best chip and model can underperform if the kernel code is mediocre. When OpenAI ran DeepSeek's R1 on Jalapeno, it hit a wall: DeepSeek's architecture uses multi-head latent attention (MLA), a design most labs including OpenAI don't use, so OpenAI had no existing internal kernel for it. Rather than assigning the problem to human specialists over weeks or months, OpenAI reportedly used Codex to write a functional, efficient kernel quickly. The kernels are described as written almost like assembly code, some around 3,000 lines, hand-tuned and backed by correctness checks and a custom sanitizer. OpenAI built this on top of Gluon, a kernel programming language it created, plus an internal "linear layout" system for mapping operations onto the chip.

This inverts the conventional strategy for challenging Nvidia. Instead of a large, friendly, human-accessible ecosystem (as Google did with TPUs and AMD with its GPUs), OpenAI built a narrow, dense, unfriendly language, Gluon, paired with an AI model willing and able to write in it directly. The scarcity protecting CUDA — the small pool of expert kernel engineers — stops mattering if an AI model can produce comparable kernel code on demand. SemiAnalysis frames it as a flywheel: OpenAI's own models currently run on Nvidia GPUs, yet those same models helped design and program a chip aimed at reducing dependence on Nvidia's ecosystem. The article cautions this isn't an immediate threat: Jalapeno is a first-generation, inference-only ASIC deployed in small volumes this year with a larger ramp in 2027; kernel work still involved humans in the loop early on; and CUDA still covers far more than inference (training, breadth of architectures, tooling). Jalapeno was reportedly designed AI-assisted and taken from concept to tape-out in about nine months, first announced October 2025 with public benchmarks around mid-2026. OpenAI reportedly intends to offer scaled internal Codex tooling to enterprise customers.

## Key points

- OpenAI's first in-house chip, codenamed Jalapeno, is an inference-only ASIC.
- SemiAnalysis testing reportedly showed Jalapeno beating Nvidia Blackwell on performance per watt in most scenarios, including >700 tokens/sec/user on DeepSeek R1 at low concurrency.
- OpenAI used Codex to write kernels in its own language, Gluon, instead of building a CUDA-style developer ecosystem.
- Codex reportedly wrote a working MLA kernel for DeepSeek R1 on the fly; kernels run ~3,000 lines, assembly-like, with correctness checks and a custom sanitizer.
- The approach removes the scarcity bottleneck of expert kernel engineers, weakening the CUDA moat's long-term protection.
- Jalapeno went concept-to-tape-out in ~9 months, announced October 2025, benchmarks ~mid-2026.
- Small deployments in OpenAI data centers this year; larger ramp in 2027.
- Not an immediate threat: inference-only, small volumes, humans in the loop, and CUDA still dominates training and breadth.

## Technical data / figures

| Item | Value |
|---|---|
| Chip name | Jalapeno |
| Type | ASIC (LLM inference only) |
| Key metric | Performance per watt |
| vs Nvidia Blackwell | Reported faster per watt in most scenarios |
| R1 low-concurrency throughput | >700 tokens/sec/user |
| Kernel language | Gluon (OpenAI-created) |
| Kernel size | ~3,000 lines (assembly-like) |
| Target model with MLA | DeepSeek R1 |
| Design cycle | ~9 months concept to tape-out |
| Announced | October 2025 |
| Public benchmarks | ~mid-2026 |
| Deployment | Small volumes 2026; larger ramp 2027 |
| Testing source | SemiAnalysis (in-lab, in-person) |

## Why this source matters for the RAG

It documents a concrete case of AI writing low-level hardware kernels, which reframes the durability of CUDA lock-in and the economics of custom silicon — highly relevant to inference cost and hardware strategy. The Gluon/Codex approach is a key data point for anyone assessing whether software moats remain a lasting barrier in AI infrastructure.


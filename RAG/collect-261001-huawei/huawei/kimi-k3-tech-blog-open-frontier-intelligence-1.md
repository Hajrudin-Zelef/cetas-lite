---
id: collect-261001-huawei/huawei/kimi-k3-tech-blog-open-frontier-intelligence-1
title: "kimi-k3-tech-blog-open-frontier-intelligence"
domain: huawei
role: reference
task: reference
actors: ["Anthropic", "Moonshot", "Nvidia"]
dates: ["2026-07-27"]
keywords: ["kimi", "agent", "agentic", "asic", "attention", "benchmark", "benchmarks", "claude", "compute", "context window", "decode", "fable 5"]
source: docs/RAG/collect-261001-huawei/kimi-k3-tech-blog-open-frontier-intelligence.md
source_anchor: ""
source_lines: [1, 74]
sha256: 609d08bc47923e9e06ad18af9ffae3dce6d393d62c2ef999abf0ef67523b3104
---

# kimi-k3-tech-blog-open-frontier-intelligence

Today, we are introducing Kimi K3 — our most capable model. Kimi K3 is a 2.8T-parameter model built on our Kimi Delta Attention and Attention Residuals, with native vision capabilities and a 1-million-token context window. It is the world's first open 3T-class model, designed for frontier intelligence across long-horizon coding, knowledge work, and reasoning.

While its overall performance still trails the most powerful proprietary models, Claude Fable 5 and GPT 5.6 Sol, Kimi K3 demonstrated frontier-level performance across our evaluation suite, consistently outperforming other tested models.

Kimi K3 is available today on Kimi.ai, Kimi Work, Kimi Code, and the Kimi API. At launch, Kimi K3 will use max thinking effort by default, with low- and high-effort modes to be introduced in subsequent updates. We are currently working closely with inference partners and open-source maintainers to align technical details and ensure a reliable rollout across the ecosystem. The full model weights will be released by July 27, 2026. Further details on the architecture, training, and evaluations will be released alongside the Kimi K3 technical report.

## An Open 3T-Class Model

Kimi K3 is the first open model to reach 2.8 trillion parameters. It marks the latest step in Kimi's sustained push at the scaling frontier: for nine of the past twelve months, Kimi models have set the upper bound of open-model sizes.

Kimi K3 is built on Kimi Delta Attention (KDA) and Attention Residuals (AttnRes), two architectural updates designed to improve how information flows across sequence length and model depth. We have also scaled up Mixture of Experts (MoE) sparsity, effectively activating 16 out of 896 experts when paired with a Stable LatentMoE framework. Together with refined training and data recipes, these structural changes yield an approximate 2.5× improvement in overall scaling efficiency compared to Kimi K2, allowing the model to convert compute into intelligence more effectively.

## Coding

Kimi K3 has strong long-horizon coding performance. Operating with minimal human oversight, it can sustain long engineering sessions, navigate massive repositories, and orchestrate terminal tools.

Kimi K3 also excels in tasks blending software engineering with visual reasoning — it leverages screenshots and visuals to optimize game dev, frontend, and CAD.

The case studies below show how Kimi K3's coding capability translates into open-ended software creation and scientific research.

### Kernel Optimization

We tested the models' capability to optimize GPU kernels. Each model works independently in an identical sandbox, with up to 24 hours to profile, rewrite, and benchmark four tasks spanning AttnRes, KDA, and a 512-head-dimension MLA kernel across NVIDIA Hopper GPUs and GPGPU from an alternative vendor. Kimi K3 performed competitively with Fable 5 (with fallback) and substantially outperformed Opus 4.8, GPT 5.6 Sol, and GPT 5.5.

Claude Fable 5 was evaluated by a third party, and its results may include fallback behavior. Across most models, some trajectories include small, acceptable precision shortcuts that remain within our numerical tolerance. GPGPU denotes general-purpose GPUs used for computation beyond graphics rendering.

In the late stages of Kimi K3 development, an early version of Kimi K3 handled the majority of the team's kernel optimization works.

### GPU Compiler Development

We further tested whether Kimi K3 could build a GPU programming system from scratch. Kimi K3 developed MiniTriton, a compact Triton-like compiler with its own tile-level IR layer over MLIR, optimization passes, and a PTX code-generation pipeline. Across supported roofline benchmarks, MiniTriton delivers performance on par with or better than Triton and torch.compile — beating Triton on certain workloads. Beyond microbenchmarks, MiniTriton sustains end-to-end nanoGPT training with stable convergence, the loss curve closely tracking the reference with only minor divergence — validating the full pipeline on a realistic workload. These results demonstrate that Kimi K3 can build a coherent end-to-end compiler — from DSL frontend and IR passes to PTX codegen and runtime — rather than isolated kernels; its from-scratch Tensor Core path already rivals Triton’s extensively optimized stack.

### Game Dev and Digital Creation

Kimi K3 combines strong 3D reasoning, coding, and vision capabilities to turn concepts, images, and videos into fully playable interactive experiences. Kimi K3 achieves true "vision in the loop" by seamlessly iterating between code and live screenshots—instantly seeing and refining outputs.

### Chip Design

As an early proof of concept, Kimi K3 designed a chip to serve a nano model built on its own architecture. In a single 48-hour autonomous run, K3 built, optimized, and verified the chip using open-source EDA tools on the Nangate 45nm library. Within 4 mm², the chip closes timing at 100 MHz and sustains over 8,700 tokens/s decode throughput in simulation, packing 1.46M standard cells, 0.277 MB of SRAM, and an INT4 MAC array with fused dequantization. A chip built by a model, for a model, reflects K3's long-horizon agentic capabilities.

### Coding for Research

Kimi K3 bridges scientific literature and executable code, autonomously implementing, validating, and analyzing complex computational research workflows.

In one case, Kimi K3 completed in about two hours what would typically require one to two weeks of work by an experienced researcher. To reproduce the I–Love–Q universal relations in computational astrophysics, it reviewed and cross-validated 20+ papers, implemented the full numerical pipeline, evaluated 300+ equations of state, identified inconsistencies in published formulas, generated 3,000+ lines of Python code, and produced an interactive HTML dashboard for exploring the results.

## Knowledge Work

Kimi K3 advances end-to-end knowledge work. Beyond public benchmarks, Kimi K3 (max) demonstrates consistent gains across our internal evaluations, which are derived from recurring patterns and challenges observed in real-world user-agent workflows. These consistent advantages across distinct production-oriented workflows reflect a broad improvement in Kimi K3's agentic knowledge work capabilities.

### Research with Interactive Visualization

Below are a few examples of what Kimi K3 in Kimi Work can produce across financial consulting and scientific research:

### Case 1: Interactive 42 years of AI ASIC industry research website

An interactive research report you can drill into: 42 years of the ASIC industry, created through 120+ rounds of recursive self-improvement. Kimi K3 transforms evidence into bespoke charts, animated diagrams, and interactive visual narratives. It pulled data via 2.8k+ web searches/fetches and 1.1k+ terminal data pulls, across 11k+ pages spanning 87 quarterly reports and 99 original PDFs.

### Case 2: Fusion Industry Research

A consulting-style industry report with interactive visualizations—including timelines, Funnel Chart, Range Bar Chart, Gantt Charts, and publication-quality slides.

### Case 3: GWTC-5 Gravitational-wave Analysis

An analysis of 391 gravitational-wave events using 20+ concurrent subagents, producing 7 scientific visualizations, 2 tables, and a literature synthesis from 10+ papers.

Kimi K3 is also particularly effective at producing infographic-style presentations, such as the fully editable heatmap and annual report shown below:

### Widgets and Dashboard

In Kimi Work, we introduce two new features - Widgets and Dashboard - which make interactions with Kimi K3 more visual and persistent. Widgets let you generate interactive components directly within a chat, with connections to local data or external plugins for continuous updates. Dashboard brings the widgets you care about most into one persistent, personalized view organized around a topic, project, or goal.

### Video Editing


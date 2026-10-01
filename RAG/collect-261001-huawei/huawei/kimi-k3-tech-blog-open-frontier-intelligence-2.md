---
id: collect-261001-huawei/huawei/kimi-k3-tech-blog-open-frontier-intelligence-2
title: "kimi-k3-tech-blog-open-frontier-intelligence"
domain: huawei
role: reference
task: reference
actors: ["Anthropic", "Apple", "Moonshot", "OpenAI", "Z.ai", "vLLM"]
dates: ["2026-07-16"]
keywords: ["kimi", "agent", "agentic", "agents", "attention", "benchmark", "benchmarks", "claude", "cyber", "disaggregated", "fable 5", "glm"]
source: docs/RAG/collect-261001-huawei/kimi-k3-tech-blog-open-frontier-intelligence.md
source_anchor: ""
source_lines: [75, 117]
sha256: 664f6d394a5d9b40b7aabddc49eef29a719d1147272e081c8b1b3b91d2dbaba3
---

# kimi-k3-tech-blog-open-frontier-intelligence

Kimi K3 excels at motion design, animation, and video editing because its native multimodal architecture understands text, images, and video within the same model.

In one example, K3 created a 3Blue1Brown-style motion-graphics explainer of its own architecture, translating technical ideas into animated diagrams and transitions.

In another, Kimi K3 edited its own teaser video from 56 source clips, handling clip selection, motion-matched cuts, frame-accurate beat synchronization, audio processing, and multiple rounds of revision. A high-density short video like this would typically take an experienced editor one to two working days, or a beginner three to five.

### Architecture and Infrastructure

Kimi K3 is built on Kimi Delta Attention (KDA) and Attention Residuals (AttnRes). KDA provides an efficient foundation for scaling attention, while AttnRes selectively retrieves representations across depth rather than accumulating them uniformly. Together, they form the architectural backbone of a model designed to scale well beyond the trillion-parameter regime.

Kimi K3 uses Stable LatentMoE, effectively activating 16 of 896 experts. At this level of sparsity, routing and optimization become first-order challenges. Quantile Balancing derives expert allocation directly from router-score quantiles, eliminating heuristic updates and a sensitive balancing hyperparameter, while Per-Head Muon extends Muon by optimizing attention heads independently for more adaptive learning at scale. Sigmoid Tanh Unit (SiTU) and Gated MLA improve activation control and attention selectivity respectively. Together, these advances enable stable and efficient training at the 2.8-trillion-parameter scale.

Kimi K3 applies quantization-aware training from the SFT stage onward, using MXFP4 weights with MXFP8 activations for broad hardware compatibility. To prevent expert imbalance from degrading throughput at large expert-parallel scales, we introduce a fully balanced expert-parallel training method with static shapes and no host synchronization on the critical path. Since inference efficiency likewise benefits from larger high-bandwidth communication domains, we recommend deploying Kimi K3 on supernode configurations with 64 or more accelerators. Finally, as KDA poses new challenges for conventional prefix caching, we have contributed a corresponding implementation to the vLLM community, to be released alongside the model. KDA with prefill cache allows us to serve Kimi K3 at a highly competitive token price despite its scale and long context.

More technical details will be available in our coming report.

## Availability

- Kimi K3 Agents: Download or update to the latest Kimi app from your mobile app store, available on iOS, Android, and HarmonyOS, or visit kimi.ai.
- Work with Kimi K3: Download the latest Kimi Work desktop app, version 3.1.0 or later, available for Windows and Apple silicon Macs.
- Code with Kimi K3: Run Kimi Code in your terminal and select Kimi K3 using the `/model` command.
- Build with the Kimi API: Visit the Kimi API Platform and select `kimi-k3` . Pricing is $0.30/MTok for cache-hit input, $3.00/MTok for cache-miss input, and $15.00/MTok for output. Powered by Mooncake's disaggregated inference architecture, the official Kimi API achieves a cache hit rate above 90% in coding workloads.
- Bring Kimi to your organization: Kimi Enterprise provides enterprise-grade data privacy and member management, with complete separation between personal and organization accounts. Visit the pricing page and select “Get Kimi Enterprise” to subscribe for your team.

### Full Benchmark Table

## Footnotes

All Kimi K3 results reported below are obtained with the reasoning effort set to 'max', setting temperature = 1.0 and top-p = 1.0. Depending on the benchmark, each model is evaluated under one of three agentic harnesses — Kimi Code, Claude Code, or Codex — as specified in the notes below.

### Coding benchmarks

1. **DeepSWE.** Kimi K3 is evaluated with the Kimi Code harness. The GLM-5.2 score is taken from the GLM-5.2 release blog (https://z.ai/blog/glm-5.2); all remaining scores are from the official DeepSWE leaderboard (https://deepswe.datacurve.ai/), under which Kimi K3 attains 67.3 with the mini-SWE-agent harness. We report the DeepSWE v1.1 tasks.
2. **Terminal-Bench 2.1.** Kimi K3 is evaluated with the Kimi Code harness. For all other models, we report the best score across harnesses: GLM-5.2 with Claude Code (https://z.ai/blog/glm-5.2); Claude Opus 4.8 and Claude Fable 5 with Terminus 2 (https://artificialanalysis.ai/evaluations/terminalbench-v2-1); GPT 5.5 and GPT 5.6 Sol with Codex (https://openai.com/index/previewing-gpt-5-6-sol/).
3. **Program Bench.** Kimi K3 is evaluated with the Kimi Code harness. The GLM-5.2 score is from https://z.ai/blog/glm-5.2; all other scores are from https://www.vals.ai/benchmarks/programbench.
4. **SWE Marathon.** Kimi K3, Claude Opus 4.8, and Claude Fable 5 are evaluated with the Claude Code harness; GPT-5.6 Sol is evaluated with the Codex harness. The GLM-5.2 score is from https://z.ai/blog/glm-5.2. Our evaluation is based on an H20-calibrated branch of the official v1.1 tasks (https://www.swe-marathon.org/): the Docker images, performance gates, and reference oracles for the GPU tasks have been recalibrated for H20, while the correctness and anti-cheat validators remain unchanged. Additionally, Claude Fable 5 hit fallbacks on 35% of the tasks in our evaluation, which may have negatively impacted its measured performance.
5. **FrontierSWE.** Kimi K3 is evaluated with the Kimi Code harness and GPT-5.6 Sol with the Codex harness; all other results are from https://www.frontierswe.com/. Dominance scores are recomputed from the raw scores using the official evaluation script and are current as of July 16, 2026.
6. **PostTrain Bench.** Scores for GLM-5.2, GPT-5.5, and Claude Opus 4.8 are adopted from the official PostTrainBench (https://posttrainbench.com/) results. Kimi K3, Claude Fable 5, and GPT-5.6 Sol are evaluated with the official Harbor implementation at maximum reasoning effort, averaged over three runs on H20 GPU (instead of H100 in the official setting) — Kimi K3 and Claude Fable 5 with the Claude Code harness, and GPT-5.6 Sol with the Codex harness.
7. **MLS Bench Lite.** Kimi K3 is evaluated with the Kimi Code harness; GLM-5.2 and the Claude models with the Claude Code harness; GPT-5.5 and GPT-5.6 Sol with the Codex harness.
8. **KCB 2.0.** Kimi K3 is evaluated with both the Kimi Code and Claude Code harnesses; GLM-5.2, Claude Opus 4.8, and Claude Fable 5 with the Claude Code harness; GPT-5.5 and GPT-5.6 Sol with the Codex harness. All models are evaluated at maximum reasoning effort, except GPT-5.5, which uses the "xhigh" setting. We also note that on this in-house benchmark, 10% of the tasks entered GPT-5.6 Sol's cyber guard.

### Productivity and agentic benchmarks


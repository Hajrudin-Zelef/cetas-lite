---
id: collect-240926-huggingface/huggingface/unsloth-qwen3-8-flash-next-gguf-hugging-face-1
title: "Read our How to Run Qwen3.8-Flash-Next Guide!"
domain: huggingface
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "OpenAI", "Unsloth"]
dates: []
keywords: ["qwen", "agent", "agentic", "agents", "agi", "attention", "benchmark", "claude", "context window", "deepseek", "embedding", "embeddings"]
source: docs/RAG/clean_en/huggingface/unsloth-qwen3-8-flash-next-gguf-hugging-face.md
source_anchor: ""
source_lines: [1, 108]
sha256: 7f5eecf8e4e9cf0a41e5dd18167d9ba4f5405d9d30bb6a15ed9cc8e2bda47546
---

# Read our How to Run Qwen3.8-Flash-Next Guide!

<!-- source: https://huggingface.co/unsloth/Qwen3.8-Flash-Next-GGUF -->

# Read our How to Run Qwen3.8-Flash-Next Guide!

    *Unsloth Dynamic 3.0 achieves superior accuracy & outperforms other leading quants.*
  

- MTP is now available for 1.3-1.7x Faster inference in Unsloth. Read Guide
- To run, please use llama.cpp or use our Unsloth Desktop app.
- See below for Qwen3.8-Flash-Next run in Unsloth Desktop with thinking controls:

As the frontier of foundation models pushes toward ever-larger parameter counts and ever-longer context windows, the question is no longer just how much we can scale, but how efficiently we can do so. Sustainable progress toward artificial general intelligence (AGI) that benefits everyone demands architectural innovation. Today, we are sharing a concrete step in that direction: Qwen3.8-Flash-Next.

This experimental preview of the architecture that will underpin Qwen4 is built around a fundamental rethinking of how the core components of modern large language models (LLMs) interact at scale.

The first open-weight release under this architecture is Qwen3.8-Flash-Next, which introduces:

- **Hybrid Attention with QSA** : The Gated DeltaNet and Gated Attention pairing has been reworked into Gated DeltaNet and Qwen Sparse Attention (QSA). Rather than selecting individual tokens for processing, QSA operates at the micro-block level. This cuts long-context latency significantly, a critical gain as agentic workloads increasingly dominate real-world usage.
- **Gated Residual** : Residual streams with normalization are what make deep LLM training manageable. Gated Residual modulates information flowing through widened residual streams via an element-wise, data-dependent read gate and a per-branch scalar write gate. This brings finer-grained expressiveness across layers while preserving training stability and keeping inference overhead low.
- **N-gram Embedding** : Embeddings provide a unique axis for parameter scaling that requires less computation and is more amenable to offloading than Mixture-of-Experts (MoE). By indexing with short n-grams, this approach makes parameter scaling highly efficient for memory-constrained accelerators without sacrificing quality.
- **Tailored Training Recipe** : The Muon and AdamW optimizers are applied to specific weight categories to maximize efficiency. Guided by refitted scaling laws, we eliminate traditional batch-size warmups and start directly at the target batch size, substantially reducing total optimizer steps while safely supporting larger learning rates for robust convergence.

For more details, please refer to our blog post Qwen3.8-Flash-Next and the technical report.

We are excited to embark on this next chapter with you and welcome your feedback as we build what comes next.

- Type: Causal Language Model with Vision Encoder
- Training Stage: Pre-training & Post-training
- Language Model
  - Number of Parameters: 125B with 6B activated, plus 51B n-gram embedding and 4B MTP
  - Hidden Dimension: 2560
  - Token Embedding: 248320 (Padded)
  - N-gram Embedding: 20,000,000 (bigrams/trigrams at layer 2)
  - Number of Layers: 48
  - Hidden Layout: 12 × (3 × (Gated DeltaNet → MoE) → 1 × (Qwen Sparse Attention → MoE))
  - Gated DeltaNet:
    - Number of Linear Attention Heads: 48 for V and 16 for QK
    - Head Dimension: 128
  - Qwen Sparse Attention:
    - Number of Attention Heads: 24 for Q and 2 for KV
    - Head Dimension: 256
    - Rotary Position Embedding Dimension: 64
    - Indexer Structure: MQA with 4 Query Heads and 1 Shared Key Head
    - Indexer Head Dimension: 128
    - Budget: 512 blocks or 2048 tokens
  - Mixture Of Experts
    - Number of Experts: 512
    - Number of Activated Experts: 10 Routed + 1 Shared
    - Expert Intermediate Dimension: 640
  - Gated Residual:
    - Number of Branches: 4
    - Bottleneck Rank: 320
  - LM Output: 248320 (Padded)
  - MTP: 1 layer, trained with multi-steps
- Context Length: 262,144 natively and extensible up to 1,000,000 tokens.

|  | Qwen3.8-Flash-Next | Qwen3.8-27B | Qwen3.7-Plus | DeepSeek-V4-Flash-0731 | Claude-Opus-4.6 (Max) | 
|---|---|---|---|---|---|
| # Params | 125B | 27B | 397B | 284B | -- | 
| # Activated params | 6B | 27B | 17B | 13B | -- | 
| # N-gram embedding params | 51B | -- | -- | -- | -- | 
| Coding |  |  |  |  |  | 
| Agentic coding DeepSWE 1.1 | **58.7** | 42.2 | 16.5 | 54.4 | -- | 
| Agentic coding SWE-bench Pro | **62.5** | 61.7 | 55.8 | 56.0 | 53.4 | 
| Multilingual software engineering SWE-bench Multilingual | **81.0** | 73.8 | 75.8 | -- | 77.5 | 
| Repo-level code generation NL2Repo-Bench | 48.1 | 42.3 | 41.1 | **54.2** | 47.6 | 
| Agent |  |  |  |  |  | 
| Long-horizon office work CoWorkBench | **73.9** | 70.7 | 65.1 | 45.1 | 68.2 | 
| Professional job tasks JobBench | **55.7** | 33.4 | 27.6 | 41.3 | 36.6 | 
| Frontier agentic tasks Agents' Last Exam | Pass@1 24.3 Score **51.2** | Pass@1 20.4 Score 42.9 | Pass@1 13.2 Score 33.6 | Pass@1 **25.2**Score -- | -- | 
| Real-world tool use Toolathlon Verified (Pass@1) | **73.5** | 67.1 | 50.6 | 70.3 | -- | 
| General |  |  |  |  |  | 
| Instruction following IFBench | **81.3** | 79.5 | 79.1 | 79.2 | 62.5 | 
| Scientific reasoning GPQA Diamond | **91.7** | 89.2 | 90.3 | 90.8 | 91.3 | 
| Multidisciplinary reasoning HLE | 35.9 | 30.8 | 34.7 | 33.8 | **40.0** | 
| Competitive coding LiveCodeBench v6 | **91.9** | 90.3 | 89.6 | 90.6 | 88.8 | 

1. DeepSWE 1.1: evaluated with the Claude Code and mini-SWE-agent harnesses, temp=1.0, top_p=0.95, 256K context window. We report the highest score across the two harnesses; notably, Qwen3.8-Flash-Next performs best on mini-SWE-agent.

2. SWE-bench Pro: except for Claude-Opus-4.6 (Max), for which we report the officially published score, all models are evaluated with the Claude Code harness, temp=1.0, top_p=0.95, 256K context window. Problematic tasks were corrected and all baseline models were re-evaluated on the refined benchmark.

3. SWE-bench Multilingual: evaluated with the mini-SWE-agent harness, temp=1.0, top_p=0.95, 256K context window.

4. NL2Repo-Bench: evaluated with the Claude Code harness. To prevent reward hacking, we disable Bash commands that attempt to access the specific repository, such as pip download, pip install and git clone.

5. CoWorkBench: an in-house cowork benchmark for evaluating long-horizon office and productivity agent tasks across computer science, finance, law, medical and other productivity domains.

6. HLE: judged by GPT-4o.

7. The best result in each row is shown in bold.

8. Empty cells (--): scores are not yet available or are not applicable.

|  | Qwen3.8-Flash-Next | Qwen3.8-27B | Qwen3.7-Plus | Claude-Opus-4.6 (Max) | 
|---|---|---|---|---|
| Agentic Multimodal Intelligence |  |  |  |  | 
| Multimodal tool use ClawEval-MM | Pass@3 **64.4**Average **60.4** | Pass@3 57.4 Average 56.9 | Pass@3 57.4 Average 60.1 | Pass@3 52.5 Average 54.7 | 
| Application recreation RecreationBench | **49.9** | 47.1 | 30.2 | -- | 
| Mobile use AndroidWorld | **84.5** | 81.9 | 81.0 | 62.0 | 
| Computer use OSWorld 2.0 | Binary **19.4**Partial **52.3** | Binary 19.4 Partial 48.0 | Binary 2.8 Partial 21.5 | -- | 
| Visual web development Vision2Web | **64.0** | 62.9 | 42.1 | -- | 
| General Multimodal Intelligence |  |  |  |  | 
| Embodied intelligence ERQA | **72.3** | 65.5 | 69.8 | 40.8 | 
| Long video understanding LVBench | **76.6** | 72.4 | 76.2 | 63.0 | 
| Real-world perception RealWorldQA | **88.5** | 85.9 | 86.9 | 73.9 | 
| Visual math problem solving MathVision | Without CI **90.6**With CI **95.7** | Without CI 90.0 With CI 94.6 | Without CI 90.3 With CI 88.7 | Without CI 65.5 | 
| Scientific chart analysis CharXiv (RQ) | Without CI 84.6 With CI **90.6** | Without CI 83.7 With CI 90.2 | Without CI **85.8**With CI 85.9 | Without CI 66.0 | 


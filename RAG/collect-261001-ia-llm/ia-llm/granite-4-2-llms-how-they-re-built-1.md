---
id: collect-261001-ia-llm/ia-llm/granite-4-2-llms-how-they-re-built-1
title: "Default: previous thinking is stripped to save context"
domain: ia-llm
role: reference
task: reference
actors: ["Google", "OpenAI", "SGLang", "vLLM"]
dates: []
keywords: ["agent", "agentic", "agents", "apache", "attention", "compute", "context window", "distribution", "embedding", "embeddings", "fine-tuning", "gemini"]
source: docs/RAG/collect-261001-ia-llm/granite-4-2-llms-how-they-re-built.md
source_anchor: ""
source_lines: [1, 70]
sha256: f939d44f1b09b5a3f1274a80e858229cb22bc0da91442d27288952717525bc1f
---

# Default: previous thinking is stripped to save context

*A technical walkthrough of how we built the Granite 4.2 reasoning model family.*

**Authors:** Granite Team, IBM

**TL;DR:** Granite 4.2 is our first family of dense, decoder-only reasoning LLMs, released in three sizes: **3B, 8B, and 30B**. These  models are post-trained from Granite-4.1 base models. Granite-4.1 base models were pre-trained from scratch on roughly 15T tokens with a five-phase strategy that extends the context window to 512K tokens, supervised fine-tuned on chain-of-thought, reasoning, and agentic-trajectory data, then post-trained with a **multi-stage reinforcement learning pipeline**. That pipeline includes agentic RL, where the 8B and 30B models learn to act with tools inside real sandboxed environments. Every model has a **thinking / non-thinking** switch, a **low-effort** thinking mode that spends a short reasoning budget on easy questions, and native tool calling. All Granite 4.2 models are released under the Apache 2.0 license.

**Links:**

Granite 4.2 is the reasoning-focused release of the Granite language-model family. Earlier Granite releases were strong instruction-following assistants; Granite 4.2 adds explicit reasoning. Every model can produce a chain of thought before its answer and can run in **thinking** or **non-thinking** mode depending on how much deliberation a task needs. A **low-effort** mode falls between the two, spending a short reasoning budget on easy questions.

The three sizes (**3B, 8B, and 30B**) share the same architectural design and follow the same training pipeline (base model, SFT, then multi-stage RL), each at its own scale. All three are strong reasoners and instruction followers. The clearest capability split shows up in post-training. The **8B and 30B** models additionally go through an **agentic RL** block that teaches them to operate as agents: calling tools, editing and running code, driving a terminal, and searching the web inside real environments. Every model supports native tool calling. Served through an OpenAI-compatible endpoint (for example, with vLLM), it emits tool calls in the OpenAI function-calling format and plugs into agentic harnesses without extra glue. Granite 4.2 is also supported in SGLang, see the SGLang cookbook for a ready-to-serve recipe.

The rest of this post walks through the build: architecture, pre-training, supervised fine-tuning, the multi-stage RL pipeline, and results.

Granite 4.2 models are built on a decoder-only dense transformer architecture with the following core components:

- **Attention:** Grouped Query Attention (GQA) with 40 attention heads and 8 KV heads
- **Position Embedding:** Rotary Position Embedding (RoPE) with θ = 10,000,000
- **Feed-Forward:** MLP with SwiGLU activation
- **Normalization:** RMSNorm (ε = 1e-5)
- **Embeddings:** Separate input/output embeddings (not tied)
- **Precision:** bfloat16

| Component | 3B Dense | 8B Dense | 30B Dense | 
|---|---|---|---|
| Embedding size | 2560 | 4096 | 4096 | 
| Number of layers | 40 | 40 | 64 | 
| Attention head size | 64 | 128 | 128 | 
| Number of attention heads | 40 | 32 | 32 | 
| Number of KV heads | 8 | 8 | 8 | 
| MLP hidden size | 8192 | 12800 | 32768 | 
| MLP activation | SwiGLU | SwiGLU | SwiGLU | 
| Sequence length | 131072 | 131072 | 131072 | 
| Position embedding | RoPE | RoPE | RoPE | 
| # Parameters | 3B | 8B | 30B | 

Granite 4.2 models are post-trained from Granite-4.1 base models. Granite-4.1 base models were trained from scratch on approximately **15 trillion tokens** using a five-phase training strategy. Phases 1–2 focus on foundational pre-training, phases 3–4 perform mid-training with progressively higher-quality data annealing, and phase 5 introduces long-context training, extending the context window to **512K tokens**. Each phase uses a distinct data mixture and learning-rate schedule, gradually shifting from broad web-scale data toward more curated, high-quality sources.

The pre-training recipe closely follows the previous generation; for a detailed treatment of the data blend, phase schedule, and long-context extension, see the Granite 4.1 blog.

Supervised fine-tuning (SFT) turns the base model into a reliable instruction-following, reasoning, and tool-using assistant. The SFT data mixture combines agentic (31.6%) and non-agentic (68.4%) data, totaling approximately 7.2 million samples, or roughly 100B tokens, of which about 65B are trainable.

The **agentic corpus** covers a broad range of domains, including software engineering (SWE, 69%), tool calling (12.1%), terminal use (8.0%), math (3.5%), search (0.8%), and action (0.2%). These samples and trajectories are generated using a diverse set of agent scaffolds and harnesses, including OpenHands, OpenCode, Terminus-2, SWE-agent, OpenResearcher, MiniSWE, OpenSeeker, EnvScaler, Gemini CLI, Hermes, Codex, and Goose. The agentic data combines samples from both open-source datasets and our own synthetically generated RL environments, spanning a variety of agent–harness combinations.

The **non-agentic corpus** consists of several major categories: instruction following (18.8%), coding (18.8%), math (14.6%), multilingual (7.0%), science (5.4%), reasoning (3.0%), and safety (0.8%).

We apply multiple stages of quality control before a sample enters the final SFT mixture. First, data from different sources is normalized and reformatted into a consistent OpenAI Chat format, making the conversation structure and tool interactions uniform across datasets and scaffolds.

We then use GPT-OSS-120B and Gemma 4 as LLM-based judges to assess sample quality. Low-scoring samples are removed, as are samples containing hallucinated or fabricated information, invalid tool interactions, or tool calls to functions that are not defined in the corresponding tool list. Several targeted, dataset-specific heuristic rules are also applied where appropriate to further improve quality and remove known sources of noise.

Finally, we perform both local and global deduplication. Deduplication is based on SHA-256 hashes computed over the combination of the `tools` and `messages` fields, removing duplicate samples both within individual data sources and across the overall SFT mixture.

The complete corpus is first globally shuffled to reduce ordering effects and ensure that samples from different domains are well mixed during training. The shuffled corpus is then partitioned into equally sized `.parquet` shards, which are tokenized using the model's tokenizer and chat template and prepared for large-scale distributed training.

Before launching the final large-scale runs, we tune hyperparameters on representative configurations, sweeping learning-rate schedules, initial learning rates, and warm-up ratios to find settings that train stably across model sizes. The final training configuration is summarized below:

| Parameter | Value | 
|---|---|
| Compute | 32–128 nodes (by model size), 4× Grace/GB200 per node | 
| Sequence length (packed) | 131,072 (128K) | 
| Global batch size | 128 | 
| Learning rate | 1.0e-5, constant after warm-up; 3.0e-6 for Phase 2 | 
| LR warm-up | 2.5% of `train_iters` | 
| Training duration | ~2 epochs | 
| Parallelism | TP=2, PP=1, CP=4 or CP=2 | 

For the 30B model, we additionally perform a second phase of SFT focused specifically on agentic coding. In this phase, agentic, SWE, and coding data are upsampled to increase their effective contribution to the training distribution, while approximately 16% of the mixture is retained as replay data from the original SFT corpus.

The 30B model is then fine-tuned for roughly one additional epoch at a lower learning rate of 3.0e-6. This targeted second phase increases the model's exposure to agentic coding trajectories without discarding the capabilities acquired during the initial SFT phase.


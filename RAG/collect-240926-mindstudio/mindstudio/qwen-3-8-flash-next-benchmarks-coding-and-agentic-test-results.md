---
id: collect-240926-mindstudio/mindstudio/qwen-3-8-flash-next-benchmarks-coding-and-agentic-test-results
title: "qwen-3-8-flash-next-benchmarks-coding-and-agentic-test-results"
domain: mindstudio
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Hugging Face", "OpenAI", "SGLang", "Z.ai", "vLLM", "xAI"]
dates: []
keywords: ["agentic", "benchmark", "benchmarks", "qwen", "agents", "attention", "claude", "compute", "consumer", "cost", "deepseek", "embedding"]
source: docs/RAG/clean_en/mindstudio/qwen-3-8-flash-next-benchmarks-coding-and-agentic-test-results.md
source_anchor: ""
source_lines: [1, 104]
sha256: 64aaf46baa2fc0e6b6b1b3871cb281b878259ee4d4ef6493fc4393dcb60a4d78
---

# qwen-3-8-flash-next-benchmarks-coding-and-agentic-test-results

<!-- source: https://www.mindstudio.ai/blog/qwen-3-8-flash-next-benchmarks -->

## Qwen 3.8 Flash Next scored 56 out of 80 on independent KingBench testing, landing behind GLM 5.3 Flash’s 63 out of 80. The gap comes almost entirely from front-end and 3D rendering tasks, while the two models tied on math and agentic pipeline tests, both hitting perfect scores. That split matters more than the headline number: it tells you what this model is actually good for.

Qwen 3.8 Flash Next is Alibaba’s preview of the architecture it plans to use for Qwen 4. It’s not a flagship release. It’s a small, cheap, fast model built to test a new set of design choices before they get scaled up, and the independent benchmark numbers give a clearer read on it than the vendor’s own marketing slides.

## TL;DR

- **Qwen 3.8 Flash Next** is a 125 billion parameter mixture-of-experts model that activates only 6 billion parameters per token, making it extremely cheap to run despite its large total size.
- On **KingBench’s eight-question coding and agentic test suite** , it scored 56/80 (70%), compared to GLM 5.3 Flash’s 63/80 (78.75%).
- The model **tied GLM 5.3 Flash on math and agentic tasks** , both scoring perfect 10s, but lost significant ground on 3D rendering and front-end polish, including a near-total failure on a folding table animation test.
- Architecture-wise, it uses a **hybrid attention system** (gated delta blocks plus Qwen Sparse Attention) and a novel**n-gram embedding layer** that stores roughly 51 billion parameters as phrase lookups that can live in regular RAM instead of GPU memory.
- API pricing is aggressive: **$0.16 per million input tokens and $0.47 per million output tokens** , about 12 times cheaper than Qwen 3.8 Max.
- The model is **open weight** under the Qwen community license, with GGUF quantizations already available for local use, requiring roughly 96 to 128GB of RAM for comfortable operation.
- On the broader KingBench leaderboard, it **outperforms GPT-5.6, Sonnet 5, and Grok 4.5** , which is a strong result for a model activating just 6 billion parameters per token.

## One coffee. One working app.

You bring the idea. Remy manages the project.

## What is Qwen 3.8 Flash Next?

Qwen 3.8 Flash Next is Alibaba’s experimental preview model, released as a sneak peek at the architecture that will underpin the next generation of Qwen models. It’s a 125 billion parameter mixture-of-experts (MoE) model, but only 6 billion of those parameters activate for any given token, which is what earns it the “Flash” name.

The model ships as a thinking model by default, meaning it generates reasoning steps before producing a final answer, though that behavior can be turned off. Context length runs natively to 262,000 tokens and can be stretched to 1 million tokens using YaRN extension. It’s fully open weight and available on Hugging Face under the Qwen community license.

Alibaba’s own benchmarks claim the model beats Claude Opus 4.5 on SWE-bench Pro (62.5 versus 53.4) and scores 91.7 on GPQA Diamond, while also outperforming DeepSeek V4 Flash, a considerably larger model, on coding tasks. Vendor benchmarks are worth treating skeptically, which is exactly why independent testing on a fixed, repeatable question set is useful.

## How does the architecture actually work?

Three design choices define this model, and they’re worth understanding because they explain both its speed and its blind spots.

First, the attention system is hybrid rather than uniform. Across 48 layers, the model follows a repeating pattern: three gated delta blocks (a form of linear attention) followed by one block of Qwen Sparse Attention (QSA). Unlike sparse attention methods that pick individual tokens to focus on, QSA operates at the microblock level, which cuts latency significantly for long-context and agentic workloads where the model is processing large volumes of tokens.

Second, there’s an n-gram embedding layer, and this is the more unusual choice. Of the 125 billion total parameters, around 51 billion are dedicated to bigram and trigram embeddings, essentially a built-in phrase dictionary. Common word groupings get stored as fixed entries rather than computed fresh each time. Critically, this layer doesn’t need to sit in GPU memory. It can live in ordinary system RAM, which is how the model scales up its total parameter count without demanding proportionally more compute.

Third, the MoE routing uses 512 experts with only 11 activated per token, combined with multi-token prediction for faster inference and gated residuals for training stability.

The upshot: this is a model that’s memory-hungry (because of the sheer parameter count) but compute-light (because so few parameters activate per token). That combination is unusual and it’s the whole bet behind the architecture.

## How did it perform on KingBench’s coding tests?

KingBench runs eight fixed tasks, each scored out of 10, for a maximum of 80 points. Testing pit Qwen 3.8 Flash Next directly against GLM 5.3 Flash, its closest competitor in the fast-and-cheap category.

Results broke down as follows:

- **Elevator simulation** (three elevators, floor-based logic, animation): Flash Next scored 5, GLM 5.3 Flash scored 6.
- **Three.js contact lens case** (clickable L/R caps, geometry): Flash Next scored 8, GLM 5.3 Flash scored 7. A win for Flash Next.
- **Three.js folding table** (slider-controlled fold animation): Flash Next scored 2, a near-total failure, versus GLM 5.3 Flash’s 7. This was the single biggest gap in the entire test set.
- **SVG panda eating a burger** : Flash Next scored 6, GLM 5.3 Flash scored 8.
- **Bow and arrow simulator game** (targets, leaderboard): both models scored 8, a tie.
- **Hard math problem** (permutation counting, correct answer 2460): both models scored a perfect 10.
- **Agentic pipeline task** (train a Pandemma 2B model, build a local web UI that generates outputs on refresh, unsupervised): both models scored a perfect 10.
- **3D wristwatch** (real-time hands, date, dual time zone): both scored 7.

## Remy doesn't write the code. It manages the agents who do.

Remy runs the project. The specialists do the work. You work with the PM, not the implementers.

Final tally: Qwen 3.8 Flash Next finished at 56/80 (70%), GLM 5.3 Flash finished at 63/80 (78.75%). The pattern is clear. Where the task is pure logic, math, or multi-step agentic execution, the two models are indistinguishable. Where the task demands visual polish, correct 3D geometry, or animation fidelity, Flash Next falls behind, and the folding table test shows it can fail outright.

## Is Qwen 3.8 Flash Next worth using?

It depends entirely on the workload. On the broader KingBench leaderboard, Flash Next lands in the middle of the pack, ahead of GPT-5.6, Sonnet 5, and Grok 4.5, which is a genuinely strong showing for a model activating only 6 billion parameters per token. For comparison, the full-size GLM 5.3 sits at the top of that leaderboard with 91.25%, and Qwen’s own larger 3.8 Max scores 81.25%.

So this isn’t a replacement for a primary coding model, and it’s not trying to be. Where it makes sense is as a cheap, fast backend for agentic pipelines, tool-calling workflows, and long-context processing, exactly the use cases its architecture was built around. The perfect scores on both the math and agentic tests back this up: those are the tasks that map most directly to real automation work, and Flash Next handles them as well as a competing model that costs more to run.

If your use case leans heavily on generating polished front-end code, 3D scenes, or animations, the benchmark results suggest you’ll hit rough edges more often than with GLM 5.3 Flash.

## Can you run Qwen 3.8 Flash Next locally?

Yes, and the architecture is specifically designed to make this feasible despite the large total parameter count. GGUF quantizations are already available, with file sizes ranging from roughly 72 to 74GB for 1-bit dynamic quants up to 94 to 111GB for 4-bit versions. An official FP8 version exists as well.

Realistically, comfortable local operation requires 96 to 128GB of RAM. A Mac with 128GB of unified memory is a good fit, as is a PC with 128GB of system RAM paired with a GPU that offloads the MoE experts to system memory. Because only 6 billion parameters activate per token, and because the large n-gram embedding layer is designed to sit in regular RAM rather than GPU memory, the model can generate at usable speeds without needing multiple high-end GPUs, unlike a dense model of similar total size that would grind to a halt on CPU-heavy setups.

For local deployment, Ollama offers the simplest path via a direct pull command referencing the Hugging Face repository and desired quantization. Llama.cpp works for an OpenAI-compatible local server. For production workloads, vLLM and SGLang are the recommended serving frameworks.

Given the API pricing, $0.16 per million input tokens and $0.47 per million output tokens, running it locally mostly makes sense for privacy requirements or offline use rather than cost savings, since the hosted API is cheap enough that hardware investment is hard to justify for most users.

## Frequently Asked Questions

### What is Qwen 3.8 Flash Next used for?

It’s best suited to agentic workflows, tool-calling pipelines, and long-context tasks where cost and speed matter more than front-end visual polish. Benchmark results show it matches larger, pricier models on math and multi-step automation tasks.

### How does Qwen 3.8 Flash Next compare to GLM 5.3 Flash?

## Other agents start typing. Remy starts asking.

Scoping, trade-offs, edge cases — the real work. Before a line of code.

On independent KingBench testing, GLM 5.3 Flash scored higher overall (63/80 versus 56/80), mainly due to stronger performance on 3D rendering and front-end animation tasks. The two models tied on math and agentic tasks, both scoring perfect marks.

### Why does Qwen 3.8 Flash Next only activate 6 billion parameters?

It uses a mixture-of-experts architecture with 512 experts, only 11 of which activate per token, plus a large n-gram embedding layer that stores common phrase patterns without needing to be computed on the fly. This keeps per-token compute low even though the total parameter count is 125 billion.

### Can Qwen 3.8 Flash Next be run on consumer hardware?

It can, though it needs significant RAM, roughly 96 to 128GB, because of its total parameter size. Because compute per token is low, it can run at usable speeds on a well-specced Mac with unified memory or a PC with a large RAM pool and GPU offloading, unlike a dense model of comparable size.

### Is Qwen 3.8 Flash Next better than Qwen 3.8 Max?

No. It’s a smaller, cheaper preview model meant to showcase a new architecture ahead of Qwen 4, not to replace the flagship. Qwen 3.8 Max scores considerably higher on independent benchmarks, and Flash Next is positioned as a low-cost option for specific workloads rather than a general-purpose upgrade.

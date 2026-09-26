---
id: collect-240926-misc/misc/qwen3-8-27b-review-the-open-weight-27b-worth-running-1
title: "qwen3-8-27b-review-the-open-weight-27b-worth-running"
domain: orcarouter
role: reference
task: reference
actors: ["AMD", "Alibaba", "Anthropic", "Hugging Face", "OpenAI"]
dates: ["2026-08-14", "2026-08-15"]
keywords: ["qwen", "agent", "agentic", "amd", "apache", "attention", "benchmark", "benchmarks", "claude", "cost", "datacenter", "fp8"]
source: docs/RAG/clean_en/misc/qwen3-8-27b-review-the-open-weight-27b-worth-running.md
source_anchor: ""
source_lines: [1, 72]
sha256: e8e1d7eae2eab8fd7ba7066b67b429d6a50ee11f2e36abdfc96a38b2117538ea
---

# qwen3-8-27b-review-the-open-weight-27b-worth-running

<!-- source: https://www.orcarouter.ai/blog/qwen-3-8-27b-review -->

**Qwen3.8 27B** is the best open-weights 27B you can self-host right now, and it is not particularly close. Weights dropped August 14, 2026 under Apache 2.0: a dense 27B (28B counting the vision encoder) with 262K native context, native image and video input, and vendor-reported agentic-coding scores that sit at or above Claude Opus 4.6 Max — SWE-bench Pro 61.7, DeepSWE 1.1 42.2, LiveCodeBench v6 90.3. The headline price is zero: download 55.6 GB and run it. The caveats are real too — independent testing finds it roughly three times slower and more token-hungry than its predecessor, every benchmark on the card is still Alibaba-reported, and the 1M context is a hosted-only feature. Verdict: if you have a 24 GB+ GPU and want frontier-adjacent capability without a metered bill, self-host it — but go in knowing exactly where the numbers are soft.

## Verdict first: should you use Qwen3.8 27B?

Short answer — **yes for local and private workloads; wait for third-party numbers before a procurement decision.** **Qwen3.8 27B** is the rare model where the open weights are the product and the API is the convenience. Because the license is Apache 2.0, the per-token price is permanently zero once you have the hardware, and nothing you build on it can be re-licensed or billed retroactively. What you give up is speed, verification, and convenience.

If you already run Qwen3.6-27B and are happy with it, the upgrade is real but not free in wall-clock terms. If you came here from the Qwen3.8-Max launch, this is the smaller, self-hostable member of the same generation — a different tool for a different job.

## The fast facts, checked August 15, 2026

• **Released** — weights went live August 14, 2026 at **Qwen/Qwen3.8-27B** on Hugging Face, mirrored on ModelScope; timezone-bracketed coverage also cites August 13, and the release landed about a week after the Qwen3.8 generation was announced.

• **License** — Apache 2.0: download, modify, redistribute, commercial use, with an explicit patent grant. Permanent.

• **Parameters** — 27B dense (28B counting the vision encoder), 64 layers, hidden size 5,120, vocabulary 248,320.

• **Architecture** — hybrid attention: 48 Gated DeltaNet linear-attention layers against 16 full Gated Attention layers (a 3:1 split). This is why a 27B dense model can carry 262K tokens of native context.

• **Context** — 262,144 tokens natively; extendable to 1,000,000 via YaRN on the hosted version.

• **Input** — native image and video alongside text; returns text. The vision encoder is why the parameter count reads 28B.

• **Thinking** — reasoning mode on by default and disableable per request; **reasoning_effort** (low/medium/high) and **preserve_thinking** for long agent runs.

• **Weights** — 55.6 GB of BF16 safetensors in 18 shards; FP8 and community GGUFs also ship.

The specs above come from the Hugging Face model card and config for **Qwen3.8 27B**, read today. The benchmark claims are Alibaba-reported; as of today no independent lab has reproduced them.

## What Qwen3.8 27B is genuinely good at

The strongest case is agentic software engineering. Alibaba reports a 3x jump on DeepSWE 1.1 over the previous 27B (42.2 vs 13.3) and a score above Claude Opus 4.6 Max on SWE-bench Pro (61.7 vs 53.4). Those are the numbers that earned it the "Opus at home" nickname — a dense 27B doing frontier-adjacent coding work on hardware a person can actually own.

Three things besides the raw numbers make it a defensible buy:

• **Native multimodal.** Image and video in, text out — a document with diagrams or a video walkthrough is not a separate model. The visual-reasoning scores (85.6 with the chain-of-thought feature enabled, 94.6 visual math, both vendor-reported) are where the vision encoder earns its keep.

• **262K native context.** Long-horizon agent tasks, big codebases, and multi-hour transcripts fit in one window. The hybrid linear-attention design keeps the KV cost of that context lower than a pure full-attention model of the same size.

• **Free, permanent weights.** This is the whole point of the category: a closed API model is rented; **Qwen3.8 27B** is owned. At 4-bit it runs on a 24 GB card (RTX 3090/4090-class), and AMD shipped Day-0 support (up to 24.5 tokens/s on a Ryzen AI Max+ 395 and 51.8 tokens/s on a Radeon AI PRO R9700, per AMD's own blog).

## Where the review gets ugly: speed, tokens, and verification

Independent testing so far is one tester's harness, not a lab — but it is the best signal we have. Running **Qwen3.8 27B** against **Qwen3.6-27B** on ten real-world tasks (judged by GPT-5.5 via the llmcompare harness), the 3.8 won nine of ten and scored 8.838 vs 6.862 on average — "one of the more impressive intelligence lifts" the tester had seen. It also used almost three times the tokens and was considerably slower; one task took roughly 600% longer. So the quality jump is real, and so is the price in wall-clock time.

Four more things to hold against it:

• **No third-party benchmarks yet.** Every headline figure in this article is Alibaba-reported as of August 15. The vendor card is detailed, but reproduction by an independent lab has not happened. If your decision depends on verified numbers, wait for them — the weights will still be there.

• **VRAM floor is real.** BF16 needs an 80 GB-class GPU. To fit a 24 GB card you must run 4-bit, and community reports (dev.to comment thread, days after release) say quantized builds "lose focus after long context." Official weights if you have the VRAM; quantized if you do not.

• **1M context is hosted-only.** The open weights cap at 262K. The extendable-to-1M figure is a Qwen Cloud hosted feature, not something a local copy does out of the box.

• **"Beats Opus" needs qualifiers.** Agentic benchmarks reward harness-specific behaviors, and a top comment on the dev.to launch post put it bluntly: "they do not beat opus on real-world usage." Treat the scoreboard as an upper bound on a good day, not a promise.

## How it sits in the Qwen 3.8 family

• **vs Qwen3.6-27B** — the same hardware class and 262K context, but a major capability jump at the cost of speed and token efficiency. If you already run 3.6 and are throughput-bound, staying is defensible.

• **vs Qwen3-Coder-30B-A3B** — the Coder is a MoE with roughly 3.3B active parameters that runs much faster (~90–110 tokens/s). The dense 27B is slower per token but inherits the generation's agentic gains; if the 27B's software-engineering scores hold up under independent testing, the Coder-30B's role shrinks.

• **vs Qwen3.8-Max** — the 2.4T MoE flagship is a datacenter model (API-only, around $2/$6 per million tokens in published coverage) with 1M context and video. The 27B is the deployable one. They answer different questions.

## Cost: the local-versus-API question, settled

For a closed model you compare per-token prices. For **Qwen3.8 27B** the marginal token costs nothing on your own hardware — the real price is the upfront GPU and your time. At 4-bit that is a 24 GB card; at BF16 it is an 80 GB-class machine. If you only need to evaluate the model, or you lack the hardware, the hosted route exists: OrcaRouter now lists **Qwen3.8 27B** with a free, rate-limited tier that bills $0 and returns HTTP 429 past its cap, plus a paid tier at $0.33 per million input and $2.40 per million output tokens (checked August 15). Qwen Cloud's hosted version, with the 1M context, is marked "coming soon."

The honest framing: for an open-weights model, a provider is a convenience, not a dependency. The free tier is the cheapest possible way to decide whether a 55.6 GB download is worth it. But once you have decided, the weights are the thing — and they are free.

## How to actually try it


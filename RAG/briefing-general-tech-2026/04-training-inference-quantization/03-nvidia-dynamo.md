---
id: briefing-general-tech-2026/04-training-inference-quantization/03-nvidia-dynamo
title: "Nvidia Dynamo and disaggregated serving"
domain: training-inference-quantization
role: deep-dive
task: architecture
actors: ["Amazon", "Cerebras", "Groq", "Nokia", "Nvidia"]
dates: ["2026-03", "2026-08", "2026-09"]
keywords: ["disaggregated", "disaggregated serving", "agentic", "decode", "gpus", "hyperscaler", "inference", "kv cache", "lpddr", "prefill", "quantization", "training"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g05-3"
source_lines: [4502, 4590]
canonical_for: ["disaggregated-serving"]
sha256: bae3436605e8c4cbd4d0148e261b25117da5d45fd0a1f40c404792d9adb76f71
---

# Nvidia Dynamo and disaggregated serving

<a id="g05-3"></a>
### 5.3 Nvidia Dynamo and disaggregated serving

If there is one software artifact that embodies the inference side of the 2026
split, it is Nvidia Dynamo. Generally available around GTC in March 2026,
Dynamo is Nvidia's open-source inference-serving framework, and its central
design decision is the one this chapter keeps returning to: a serving
deployment is not a fleet of identical GPUs but a managed fabric of
specialized resources, with prefill and decode explicitly separated and
scheduled independently. Version 1.4, released in August 2026, pushed the
disaggregated serving model further, making prefill/decode separation a first-
class, production-supported pattern rather than an advanced configuration.

The engineering logic is worth spelling out in full, because it explains the
entire inference-architecture half of this chapter — the LPDDR racks, the
telco deployments, the cost-per-token metric, and the quantization wave alike.

Prefill — ingesting the prompt, the system instructions, the retrieved
documents, the tool definitions, the conversation history — is compute-
intensive and bursty. It processes a large input tensor in one shot: high
arithmetic intensity, a premium on peak FLOPS and on high-bandwidth memory
that can feed the compute units. A prefill-heavy deployment (long documents,
retrieval-augmented generation, agentic tool loops that re-ingest context)
wants the densest compute it can get, and it wants it in bursts.

Decode — emitting tokens one by one — is the opposite workload. Each step
reads the entire weight set and the growing KV cache to produce a single
token: low arithmetic intensity, bound by memory bandwidth, latency-sensitive
because the user is waiting. A decode-heavy deployment (long generations,
chain-of-thought reasoning, agentic rollouts) wants large memory capacity for
the KV cache, high memory bandwidth per watt, and minimal inter-node chatter —
and it does not need peak FLOPS at all.

Running both phases on the same hardware configuration means compromising
both: prefill-optimized silicon idles its FLOPS during decode, decode-
optimized memory is overkill for prefill's bursty compute. Separating them —
disaggregation — means each phase gets the hardware it actually wants, sized
independently, scaled independently, and priced independently. An operator can
overprovision cheap decode memory while keeping prefill compute lean, or vice
versa, matching the fleet to the workload mix instead of the workload mix to
the fleet.

That is also why disaggregation and quantization are the same story told
twice. Dynamo optimizes *which* hardware runs each phase; quantization shrinks
*what* runs on it — fewer bits per weight means less memory traffic per decode
step, a smaller KV cache per unit of context, cheaper memory per token. The
two multiply: disaggregated placement × quantized weights × compressed KV
cache is the full 2026 inference-economics stack, and Dynamo is the software
layer that makes the first term deployable.

2026 produced several dated, heterogeneous examples of the disaggregation
pattern in production and in announcement, and their heterogeneity is the
point — the idea escaped any single vendor's stack:

- **AWS + Cerebras (summer 2026):** a hybrid serving split in which AWS Trainium handles the compute-heavy prefill phase while Cerebras — whose wafer-scale architecture keeps memory adjacent to compute, exactly the decode-phase bottleneck — handles token generation. This is disaggregation implemented across two vendors' silicon: prefill where FLOPS are cheap, decode where memory is fast. It also demonstrates that the disaggregated pattern is not an Nvidia-only story, even though Dynamo is Nvidia's framework.
- **Groq LPX (GTC 2026):** Groq's inference hardware takes the opposite route to disaggregation-by-cluster — instead of separating phases across machines, it collapses the decode bottleneck into a single chip. The LPU carries 512 MB of on-chip SRAM, keeping weights and KV state adjacent to compute and removing off-chip memory from the critical path for decode. Same economics as disaggregation, different architecture: inference-optimized memory hierarchy rather than phase-separated fleets. Both are answers to the same question — what does hardware look like when decode, not training, is the workload?
- **AI-RAN (GTC March 2026 → eight operators in field trials by September 2026):** the Nvidia/T-Mobile/Nokia initiative to run AI inference on telecom infrastructure. Once decode is power-efficient and phases are separable, inference capacity doesn't need a hyperscaler's purpose-built datacenter — it can live in the network, close to users, where latency and bandwidth economics favor it. Eight operators in field trials by September makes this a deployment program, not a concept.

Dynamo's significance, against this backdrop, is standardization.
Disaggregation as an architecture is something any competent infrastructure
team could sketch; Dynamo turned it into a deployable default — an open
framework, GA'd in March 2026, matured through v1.4 by August — that lets
operators treat prefill and decode as independently provisioned, independently
scaled, independently priced phases. Everything in the inference-rack section
below — LPDDR memory pools, telco deployment, the cost-per-token metric that
arbitrates it all — follows from that decomposition being available as
software rather than as a research result.

| Disaggregation datapoint (2026) | Detail |
|---|---|
| Nvidia Dynamo | Open-source serving framework; GA ~GTC March 2026; v1.4 August 2026 with prefill/decode disaggregation as first-class pattern |
| AWS + Cerebras | Summer 2026: Trainium for prefill, Cerebras for decode — cross-vendor disaggregation |
| Groq LPX | GTC 2026: LPU with 512 MB on-chip SRAM — decode bottleneck collapsed into the chip |
| AI-RAN | Nvidia/T-Mobile/Nokia at GTC March 2026 → 8 operators in field trials September 2026 — inference in the telco fabric |
| The logic | Prefill = compute-bursty (wants FLOPS); decode = memory-bound, latency-sensitive (wants capacity + bandwidth/watt); separate them, size them independently |

Disaggregation has second-order effects beyond the hardware sizing, and 2026
operators learned them in production. Independent scaling means the prefill
and decode fleets can have different refresh cycles — prefill tracking the
FLOPS-per-dollar curve of new accelerators, decode tracking the gigabytes-per-
dollar curve of memory — instead of the whole fleet turning over on the
training generation's cadence. Independent pricing means prefill-heavy
workloads (long documents, RAG, context re-ingestion) and decode-heavy
workloads (long generations, reasoning traces) can be billed against the
resources they actually consume rather than subsidizing each other inside a
blended per-token price. And separated failure domains mean a decode-fleet
incident doesn't take down prompt processing. None of these were the headline
reason anyone adopted Dynamo; all of them are reasons nobody reverted.


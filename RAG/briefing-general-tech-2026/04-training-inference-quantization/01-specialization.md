---
id: briefing-general-tech-2026/04-training-inference-quantization/01-specialization
title: "The great specialization: training vs inference"
domain: training-inference-quantization
role: deep-dive
task: architecture
actors: ["Amazon", "Cerebras", "Groq", "MLCommons", "Nokia", "Nvidia", "Positron", "Taalas", "UALink"]
dates: ["2026-02-20", "2026-03", "2026-09"]
keywords: ["inference", "training", "blackwell", "cost per token", "cowos", "decode", "disaggregated", "disaggregated serving", "funding", "gpu", "gpus", "hbm"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g05-1"
source_lines: [4302, 4409]
canonical_for: ["training-inference-split"]
sha256: 28e31caf5adbbf0536c0ff3aad07f18319c3cec9e886a988c4db336584cda185
---

# The great specialization: training vs inference

<a id="g05-1"></a>
### 5.1 The great specialization: training vs inference

For three years, "AI infrastructure" meant one thing: buy the densest rack of
Nvidia GPUs you could finance, stack it with high-bandwidth memory, and serve
whatever workload showed up. Training and inference shared the machine, the
memory hierarchy and the networking. A training cluster idled between runs
could serve inference; an inference fleet was just a training fleet with a
different job scheduler. The hardware didn't care what the software was doing,
because the hardware had been designed for the hardest workload — training —
and inference simply inherited it.

What 2026 confirmed — in vendor roadmaps, in operator purchases, in startup
funding rounds, and in the MLPerf submissions recorded in September — is that
this shared era is ending. The two workloads have opposite physics, and
infrastructure is finally being designed around that difference rather than
around the convenience of a single SKU.

Training is compute-saturating, memory-hungry and brutally synchronous. A
training run is one giant computation spread across thousands of accelerators
that must stay in lockstep: every step requires exchanging gradients and
activations across the entire cluster, so the bottleneck is inter-GPU
bandwidth and power delivery, not any single chip's arithmetic. The 2026
training rack is therefore defined by density above all else: high-bandwidth
HBM (HBM3e-class stacks across the 2026 fleet, HBM4-class designs arriving on
the newest accelerators), NVLink/UALink scale-up fabrics that treat a full
rack as a single memory domain, and power envelopes that would have been
unthinkable two years earlier. The Rubin NVL72 rack sits at roughly 190–230 kW
— more than double the Blackwell GB200 NVL72's 100–120 kW class, a figure
frequently misattributed to Rubin, and the correction matters precisely
because the gap between the two numbers is the measure of how far training
density moved in one generation. NVL144, at 600 kW, sits on the roadmap as a
planned 2027 deployment, not a shipping product — a rack that is no longer
datacenter equipment in any conventional sense but a purpose-built power plant
with a compute payload.

Inference has the opposite shape, and 2026 was the year the industry stopped
pretending otherwise. Inference is memory-bound rather than compute-bound:
generating tokens is limited by how fast weights and cached state can be
streamed to the compute units, not by how fast the units can multiply. It is
latency-sensitive rather than throughput-greedy: a user waiting for the first
token experiences delay, not batch efficiency. And — crucially — it is
decomposable: the prefill phase (processing the prompt, the system
instructions, the retrieved documents, the conversation history in one
compute-intensive pass) and the decode phase (emitting tokens one by one, each
step reading the full weight set and the growing KV cache) have different
resource profiles and no longer need to live on the same silicon.

That decomposability is the deep cause of everything in the inference half of
this chapter. It is what made 2026's disaggregated serving designs possible —
Nvidia's Dynamo, the Trainium-plus-Cerebras split, Groq's LPU — because once
prefill and decode are separate scheduling problems, each can be sized, priced
and placed independently. It is what lets inference hardware leave the HBM
straitjacket behind: decode cares about memory capacity and bandwidth-per-
watt, not peak floating-point throughput, so large, cheap memory pools built
on LPDDR5X (Positron's bet, funded with an $875M Series C on September 11,
2026) or dense on-chip SRAM (Groq's LPX, shown at GTC 2026 with 512 MB of SRAM
on the LPU) or even packaging-free silicon with no HBM and no CoWoS at all
(Taalas's HC1, February 20, 2026) are credible inference substrates. And it is
what lets inference leave the hyperscaler core: once the hardware is power-
efficient and the phases are separable, capacity can live at colocation
facilities, enterprise on-prem, and telco points of presence — the AI-RAN
initiative (Nvidia, T-Mobile, Nokia at GTC in March 2026, eight operators in
field trials by September 2026) being the dated proof that inference capacity
is being planned into the telecom fabric.

The business consequence is a split in the capital stack itself. Training
still concentrates where power and scale-up fabrics are: purpose-built AI
datacenters with hundreds of megawatts, financed against training contracts.
Inference disperses to where users and cheap power are, financed against per-
token margins. The two are no longer the same purchase decision, the same
depreciation schedule, or the same engineering discipline — and the vendors,
from Nvidia's Dynamo to Positron's LPDDR5X racks, have reorganized around that
fact.

|  | Training | Inference |
|---|---|---|
| Bound by | Inter-GPU bandwidth, power delivery | Memory bandwidth-per-watt, latency |
| Memory | High-bandwidth HBM (HBM3e → HBM4-class) | Large cheap pools: LPDDR5X, SRAM, HBM where justified |
| Interconnect | NVLink/UALink scale-up; rack as one memory domain | Disaggregated: prefill and decode separable, placeable independently |
| Power envelope | Rubin NVL72 ~190–230 kW; NVL144 600 kW planned 2027 | Power-efficient; edge/colo/telco-capable |
| Placement | Purpose-built AI datacenters | Hyperscaler core + colo + telco + edge (AI-RAN: 8 operators in trials, 09/2026) |
| Economic unit | The training run / the cluster contract | The token; cost per token |

The following sections walk through this split rack by rack, format by format:
first the training side (dense, HBM, high power), then the inference side
(disaggregated, LPDDR-friendly, edge-capable), then the metric that arbitrates
between them — cost per token — and the quantization techniques that moved
that metric more than anything else in 2026.

Why did the shared era last as long as it did? Part of the answer is that
through 2025, inference simply wasn't big enough to justify its own hardware.
Training spend dominated the capital conversation, inference was a rounding
error on the fleet, and buying one SKU for both workloads was operationally
simpler than maintaining two engineering disciplines, two vendor relationships
and two depreciation schedules. The split became economical only when the
composition of spend flipped — when inference grew from a side workload into
65–90% of operating cost. Specialization has a fixed cost (separate designs,
separate software, separate teams); it pays that cost only when the workload
is large enough. By 2026, inference was.

There is a second, subtler reason: the decomposability insight had to be
proven in production, not just in papers. Prefill/decode disaggregation was
discussed for years before Dynamo made it a GA'd, supported default in March
2026. Ideas don't split industries; deployable defaults do. The 2026 split is
therefore as much a software story as a silicon story — and the capital
followed the software.


---
id: briefing-general-tech-2026/08-market-analysis/04-agentic-infrastructure
title: "The agentic infrastructure shift"
domain: market-analysis
role: deep-dive
task: architecture
actors: ["Fujitsu", "Google", "Intel", "Nvidia", "Positron", "Qualcomm"]
dates: ["2026-09-11"]
keywords: ["agentic", "ai200", "crescent island", "decode", "disaggregated", "gpu", "hbm", "inference", "lpddr", "lpddr5x", "memory shortage", "monaka"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g09-4"
source_lines: [9504, 9599]
canonical_for: ["agentic-infrastructure"]
sha256: de9f06cb271c3d625d9afe05a6b31c675594432bdaf51703ded2ba6428a848aa
---

# The agentic infrastructure shift

<a id="g09-4"></a>
### 9.4 The agentic infrastructure shift

Somewhere in 2026, the center of gravity of AI infrastructure demand moved from training
to inference — and, more precisely, from inference as a serving afterthought to
inference as the design driver of the whole stack. The agentic workload is what moved
it: multi-step agents that plan, call tools, read long contexts, and iterate do not look
like chatbots at the infrastructure level. They are memory-hungry, latency-sensitive in
the prefill phase, throughput-hungry in the decode phase, and brutally cost-per-token
sensitive at scale. Every architectural trend verified in the February–September window
is legible as a response to that workload.

The most quotable data point of the year came from Intel's Lip-Bu Tan at Computex: the
CPU-to-GPU ratio in AI systems is moving from roughly 1:8 toward roughly 1:1. Read
carefully, that is not a statement about Intel versus Nvidia; it is a statement about
what the machines are *for*. Training clusters were GPU farms with token CPUs attached.
Agentic inference systems need serious CPU capacity — for orchestration, tool execution,
data movement, and the growing share of the pipeline that is not matrix multiplication.
A 1:1 ratio describes a balanced inference appliance, not a training rig. The datacenter
of 2026 is being rebalanced around the agent.

#### The 1:1 ratio, unpacked

Lip-Bu Tan's Computex figure — CPU-to-GPU ratios moving from roughly 1:8
toward roughly 1:1 — deserves unpacking because it is easily misread as a
vendor pitch. It is better read as a workload signature. A 1:8 machine is a
training machine: acres of matrix-multiply silicon with just enough CPU to
feed it. A 1:1 machine is an inference appliance: the CPUs are doing real
work — orchestrating agentic loops, executing tool calls, moving data between
the growing number of pipeline stages, running the non-attention parts of the
model and the increasingly heavy pre- and post-processing around it. The ratio
is the industry's way of saying, without saying it, that the unit of
computation has changed from the training run to the agentic task.

The economic logic follows the workload. Training is a batch process: you can
tolerate scheduling inefficiency because the job runs for weeks and the
objective is throughput per dollar over the whole run. Agentic inference is an
interactive process with strict phase-level latency requirements — prefill
must be fast because the agent is waiting on context, decode must be steady
because the user is watching tokens arrive — and the cost that matters is
cost per served token at quality, not cost per FLOP. Disaggregated
prefill/decode is the architectural expression of this: scale the
bandwidth-hungry phase and the capacity-hungry phase independently, schedule
them independently, price them independently. What was a research curiosity
for chatbots became production practice for agents because agents made both
phases worse at once — longer contexts in, more tokens out per task.

The decentralization of inference geography is the same logic at the map
level. Training centralized because gigawatt campuses are efficient and
talent is scarce; inference is decentralizing — toward edge, colocation, and
telco facilities — because inference revenue is latency- and locality-bound
in a way training never was, and because a megawatt-scale inference
deployment fits sites that could never host a training cluster. The carriers'
2026 awakening to their own assets — real estate, power feeds, last-mile
proximity — is the early signal of a build-out that will look less like the
hyperscale campuses of 2023–2025 and more like the CDN build-out of the
2010s: everywhere, smaller, closer.

The serving architecture followed. Disaggregated prefill and decode — splitting the
prompt-processing phase from the token-generation phase onto different hardware, each
scaled and scheduled independently — moved from research curiosity to production
practice in the window. The logic is agentic: prefill is compute- and
memory-bandwidth-intensive over long contexts; decode is memory-capacity-intensive and
latency-sensitive. Agents make both phases worse simultaneously (longer contexts, more
generated tokens per task), so the economic case for disaggregation, which was marginal
for chatbots, became compelling for agents. Expect every serious inference platform to
be disaggregated by the end of 2027; the technique is too well matched to the workload
to stay niche.

The geography of inference moved too. Training centralized in a handful of gigawatt
campuses; agentic inference is decentralizing toward edge, colocation and telco
facilities — closer to users, closer to data, and sized in megawatts rather than
gigawatts. The economics are straightforward: inference revenue is latency- and
locality-sensitive in a way training never was, and the power envelope of an inference
deployment fits sites that could never host a training cluster. The telco-inference
story, in particular, bears watching: carriers own the real estate, the power feeds and
the last mile, and 2026 was the year they started to realize it.

Underneath all of this sits the memory economics of section 9.2, and the hardware
response is the LPDDR inference-chip class: Crescent Island, Qualcomm's AI200,
Positron's LPDDR5X parts, Fujitsu's MONAKA. The shared thesis is that agentic inference
is memory-capacity-bound rather than compute-bound, and that LPDDR — cheaper per byte,
more available than HBM, adequate in bandwidth for batched inference — is the right
trade. This is the architectural bet that the memory shortage made rational: when HBM is
fully booked into 2028, you do not wait; you design around it. The $875 million Positron
Series C of 11/09/2026 is the market's vote that the bet has legs.

A coda from the research world: the DeepMind Institute debate, which ran through the
period, was in part about exactly this — what the infrastructure of the agentic era
should look like and who should build it. The debate's existence is itself a signal:
when the leading research institution's internal argument is about infrastructure rather
than algorithms, the bottleneck has moved. The algorithms of the agentic era are
substantially here; the infrastructure to run them at planetary scale is being built
now, under the memory constraint, around the 1:1 ratio, in disaggregated and
increasingly decentralized form.


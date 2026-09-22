---
id: briefing-general-tech-2026/04-training-inference-quantization/05-inference-racks
title: "Inference racks: disaggregated, LPDDR, edge"
domain: training-inference-quantization
role: deep-dive
task: architecture
actors: ["Amazon", "Cerebras", "Groq", "Nokia", "Nvidia", "Positron", "Taalas"]
dates: ["2026-02-20", "2026-03", "2026-09", "2026-09-11"]
keywords: ["disaggregated", "inference", "lpddr", "accelerator", "cost per token", "cowos", "decode", "dram", "gpus", "hbm", "hyperscaler", "kv cache"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g05-5"
source_lines: [4668, 4716]
sha256: 4fdca90030c2a7fd266be17a0a1f27635b2f494eb08951e7c03ddde65c30c6d4
---

# Inference racks: disaggregated, LPDDR, edge

<a id="g05-5"></a>
### 5.5 Inference racks: disaggregated, LPDDR, edge

The inference rack of 2026 is defined by the opposite choices from Section
5.4's training rack: cheap memory over fast memory, separated phases over
lockstep domains, placement flexibility over power density. Where training
asks "how do we keep 72 GPUs in lockstep," inference asks "how do we serve the
most tokens per dollar, per watt, per square meter" — and in 2026 the answers
arrived as four distinct hardware stories, each dated, each a different bet on
the same economics.

**Disaggregation as the organizing principle.** As Section 5.3 described, prefill and decode are different workloads with different hardware wants. The inference fleet of 2026 — and "fleet" is more accurate than "rack," because it need not be a rack — is increasingly two fleets: compute-dense hardware for prefill, where peak FLOPS still matters for bursty prompt ingestion, and memory-dense, power-efficient hardware for decode, where the KV cache lives and per-watt memory bandwidth is the currency. Dynamo v1.4 made this a supported deployment pattern in Nvidia's open framework; the AWS Trainium-plus-Cerebras split of summer 2026 made it a cross-vendor production reality. The practical consequence for operators is independent scaling: a workload mix that shifts toward long generations gets more decode memory without buying more prefill FLOPS, and the fleet's cost tracks the workload instead of the workload being contorted to fit a uniform fleet.

**LPDDR and the end of the HBM monopoly.** Decode is memory-capacity-bound, not memory-bandwidth-extreme. What the decode phase needs is gigabytes of affordable, power-sipping memory sitting near the compute — capacity for the KV cache, enough bandwidth per watt to stream weights once per token, and nothing more. That profile opened the door to LPDDR — mobile-derived DRAM that trades peak bandwidth for cost and power efficiency — as a serious inference substrate. For a decade, "serious AI memory" meant HBM, full stop; 2026 was the year credible, funded challengers said the decode phase doesn't need it:

- **Positron — Series C, $875M, September 11, 2026.** An inference-only hardware bet built on LPDDR5X memory, funded at a scale that makes it impossible to dismiss as an experiment. The proposition is pure decode economics: large, cheap memory pools beat HBM on cost per token for the memory-bound phase, and an inference-only company can optimize the entire stack — silicon, memory, networking, software — around that single phase without carrying training's constraints. The round's size is itself a data point about where 2026 capital believed the inference margin would be captured: not in general-purpose accelerators, but in memory-architecture specialization.
- **Groq LPX — GTC 2026.** Groq's answer attacks the same bottleneck from the opposite direction: instead of cheaper off-chip memory, no off-chip memory in the critical path at all. The LPU carries 512 MB of on-chip SRAM, keeping weights and KV state adjacent to compute. For decode — where every token generation is a full pass over weights plus cache — SRAM's bandwidth and latency characteristics are ideal, and the power cost of moving data millimeters instead of centimeters is transformative. It is the most literal possible interpretation of "memory-bound workload wants memory near compute."
- **Taalas HC1 — February 20, 2026.** The most radical statement of the trend: an AI accelerator with no HBM and no CoWoS advanced packaging at all. CoWoS — the advanced packaging that marries compute dies to HBM stacks — has been a supply-chain bottleneck and a cost center for the entire AI boom; designing it out entirely is a bet that inference silicon can live on conventional packaging with conventional memory and still win on cost per token. Dated to February, the HC1 is also a reminder that the inference-specialization thesis was being built in silicon well before the September benchmarks confirmed it.

**The edge and the telco as deployment targets.** Once inference hardware is power-efficient and phases are separable, it no longer needs to live in a hyperscaler's purpose-built datacenter next to a 200 kW training rack. The placement map of 2026 inference widened accordingly: colocation facilities with cheaper power, enterprise on-premises deployments, and — most structurally — telco points of presence. The AI-RAN initiative (Nvidia, T-Mobile, Nokia, presented at GTC in March 2026, with eight operators in field trials by September 2026) is the dated proof point: inference capacity planned into the telecom fabric, close to users, where latency and metro-bandwidth economics favor local serving over backhaul to a distant region. Training hardware could never support that map — you cannot put a 230 kW NVL72 in a central office. Inference hardware, disaggregated and LPDDR-efficient, can.

| Inference bet (2026) | Memory answer | The proposition |
|---|---|---|
| Positron (Series C $875M, 11/09/2026) | LPDDR5X | Large, cheap memory pools beat HBM on decode economics; inference-only full-stack optimization |
| Groq LPX (GTC 2026) | 512 MB on-chip SRAM | Keep weights/KV adjacent to compute; remove off-chip memory from the decode critical path |
| Taalas HC1 (20/02/2026) | No HBM, no CoWoS | The HBM tax — and the advanced-packaging bottleneck — is avoidable for inference silicon |
| AI-RAN (GTC 03/2026 → 8 operators in trials 09/2026) | Telco-sited capacity | Inference belongs in the network fabric, not just the cloud core |
| AWS Trainium + Cerebras (summer 2026) | Cross-vendor phase split | Prefill on Trainium, decode on Cerebras — disaggregation across vendors' silicon |

The common thread is that every one of these bets is downstream of
decomposability. None of them makes sense for training; all of them make sense
the moment prefill and decode are separate problems and decode is understood
as a memory-capacity workload. The inference rack of 2026 is not one design —
it is the space of designs that the training rack's constraints had been
suppressing, now released.

Read the five 2026 bets together and a pattern emerges: every one of them is
downstream of decomposability, and every one of them is capital voting on the
same thesis. Positron's $875M says decode memory should be cheap and abundant.
Groq's SRAM says decode memory should be adjacent. Taalas's no-HBM design says
the advanced-packaging bottleneck is optional. AI-RAN's eight operator trials
say inference capacity belongs in the network. Trainium-plus-Cerebras says
phases can split across vendors. None of these companies is building a
training rack; all of them are building answers to the question "what does
hardware look like when decode, not training, is the workload?" The diversity
of the answers — cheap DRAM, on-chip SRAM, no HBM at all, telco sites, cross-
vendor splits — is itself the signal: when a design space opens, capital
explores all of it at once, and 2026 was the exploration year.


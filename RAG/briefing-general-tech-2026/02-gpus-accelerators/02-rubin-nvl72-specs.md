---
id: briefing-general-tech-2026/02-gpus-accelerators/02-rubin-nvl72-specs
title: "Rubin NVL72: architecture and specifications"
domain: gpus-accelerators
role: deep-dive
task: hardware
actors: ["AMD", "CoreWeave", "Huawei", "Intel", "MLCommons", "Nvidia", "Qualcomm"]
dates: []
keywords: ["agentic", "benchmark", "blackwell", "cost per token", "disaggregated", "disaggregated serving", "fp4", "gpu", "gpus", "hbm4", "helios", "inference"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g03-2"
source_lines: [1509, 1581]
canonical_for: ["nvidia-vera-rubin"]
sha256: 777a33af9abeec2507160ac0ba4c0d1c1644e803202dcca99a833d03b2799ce4
---

# Rubin NVL72: architecture and specifications

<a id="g03-2"></a>
### 3.2 Rubin NVL72: architecture and specifications

The NVL72 is the atomic unit of the Rubin era: a single rack containing 72 Rubin GPUs and 36 Vera CPUs, tied together by NVLink 6 into one coherent compute domain. Nvidia's headline figure for the rack is **3.6 exaFLOPS of NVFP4 compute per rack** — the first time a single rack has crossed into the exaFLOPS regime at 4-bit precision.

The interconnect numbers deserve careful reading, because they are the most commonly misreported figures in the 2026 coverage. The NVL72's scale-up domain delivers **260 TB/s of aggregate bandwidth** — which works out to **3.6 TB/s per GPU**. That 260 TB/s figure is the aggregate across the domain. It is *not* a per-link figure, and presenting it as per-link bandwidth inflates the number by the count of GPUs in the domain. Any 2026 source quoting "260 TB/s per link" for NVL72 has misread the spec.

#### Specifications

| Specification | Rubin NVL72 | Notes |
|---|---|---|
| GPUs per rack | 72 Rubin | — |
| CPUs per rack | 36 Vera | 1 CPU per 2 GPUs |
| Interconnect | NVLink 6 | Scale-up domain |
| Compute per rack | 3.6 exaFLOPS NVFP4 | First exa-scale rack at FP4 |
| Scale-up bandwidth (aggregate) | 260 TB/s | = 3.6 TB/s per GPU; **not per-link** |
| Rack power | ≈ 190–230 kW | — |

#### NVLink 6 in one paragraph: what "scale-up" means

A GPU rack faces a fundamental problem: 72 GPUs are only as useful as the fabric connecting them. Without a high-bandwidth scale-up interconnect, each GPU is an island and multi-GPU training or serving pays a crushing communication tax. NVLink is Nvidia's answer — a proprietary GPU-to-GPU fabric that, in its sixth generation, binds all 72 Rubin GPUs (and the 36 Vera CPUs) into one coherent domain with 260 TB/s of aggregate bandwidth. The "6" is the generation; the architectural point is that the domain behaves as a single system for software purposes. AMD's Helios, Huawei's UnifiedBus, and Qualcomm's PCIe scale-up are all answers to the same problem — different fabrics, same necessity. In 2026, the scale-up fabric became as important a differentiator as the GPU itself, because the workloads (large-model training, disaggregated inference) all demand that dozens of accelerators act as one.

#### 36 Vera CPUs: the 1:2 ratio

The NVL72's 36 Vera CPUs for 72 Rubin GPUs — one CPU per two GPUs — is the physical expression of the control-plane thesis Intel's Tan articulates for the industry (§3.10). The Vera CPUs handle orchestration, data movement, and the host-side work of disaggregated serving; the ratio says how much control plane Nvidia thinks an exa-scale rack needs. It is also a notable density of CPU compute in its own right: 36 server-class CPUs per rack is a small datacenter's worth of general-purpose compute embedded in an AI system. As agentic workloads shift work toward the CPU side, that embedded capacity becomes a strategic reserve — the rack carries its own orchestration headroom rather than depending on external hosts.

#### The power figure, corrected

The power figure is the second commonly misreported number. Rubin NVL72 racks draw roughly **190 to 230 kW** — a step up that pushes datacenter power delivery and liquid cooling to the front of every deployment conversation. The frequently seen "100–120 kW" figure belongs to the previous generation: it is the Blackwell GB200 NVL72's envelope, not Rubin's. Conflating the two understates Rubin's power requirements by nearly half, and it misleads capacity planning for the H2 2026 deployments.

The implications are concrete. A 190–230 kW rack cannot be air-cooled in any conventional datacenter; it requires liquid cooling (direct-to-chip at minimum), reinforced power delivery, and in many facilities a retrofit. This is why the deployment story of Rubin in 2026 is inseparable from the neocloud story: CoreWeave and its peers were building greenfield AI datacenters designed around exactly these envelopes, while hyperscalers were adapting existing footprints. When Huang said "full production" in January, the unasked question was whether the *datacenters* were in full production too — and the H2 2026 delivery calendar was, in part, the answer.

#### What 4-bit doesn't do: the limits of the precision story

The chapter's FP4 enthusiasm needs one counterweight: 4-bit precision is not free. Quantizing to FP4 can degrade model quality on sensitive tasks, requires careful calibration (and sometimes quantization-aware retraining), and its benefits concentrate in memory-bound serving — compute-bound phases gain less. The industry's convergence on FP4 as the headline format reflects a judgment that the quality cost is acceptable for most serving workloads, not that it is zero. Buyers evaluating NVFP4/MXFP4 claims should ask the quality question explicitly: at what benchmark cost does the throughput gain come? None of the verified sources in this round published quality-degradation figures alongside the throughput multiples — a gap in the 2026 record worth noting, since throughput without quality is an incomplete metric.

#### Architecture: continuity with escalation

The architectural story of NVL72 is continuity with escalation. The 72-GPU NVL domain concept, the Grace-to-Vera CPU pairing philosophy, the NVLink scale-up fabric — all of it extends the Blackwell playbook. What changed is the density: exa-scale FP4 in one rack, Vera CPUs replacing Grace, NVLink 6 widening the domain. Nvidia's message at GTC was that the platform shift customers must absorb is not architectural but operational: 190–230 kW per rack means the datacenter, not the chip, is now the binding constraint on deployment speed.

That constraint is visible in the deployment details that did emerge. The first production rack at CoreWeave's Livingston site in June, the eight-partner H2 ramp, and the September confirmation of a multi-rack 504-GPU CoreWeave cluster (see §3.3) all point to the same reality: Rubin is shipping into a power-and-cooling-limited world, and the vendors winning 2026 are the ones whose customers can actually feed the racks.

#### Exa-scale in one rack: why the milestone matters

"3.6 exaFLOPS per rack" is the first exa-scale rack figure at 4-bit precision — a milestone worth pausing on. An exaFLOP is a quintillion operations per second; putting that in a single rack means the unit of supercomputing has shrunk from the room (the exa-scale systems of the early 2020s filled datacenter halls) to the cabinet row. The practical consequence is deployment granularity: buyers can now provision exa-scale inference capacity rack by rack, scaling linearly with floor space and megawatts rather than in datacenter-sized increments. It is also the figure that makes the power envelope (§3.2's 190–230 kW) legible — exa-scale compute was never going to be air-cooled. The milestone belongs to Nvidia first, but the industry's direction is set: AMD's 2.9 and Huawei's 60 EFLOPS (system-level) are all measured against the precedent that one rack can now hold an exaFLOP.

#### NVFP4: the precision story in one paragraph

NVFP4 — 4-bit floating point in Nvidia's implementation — is the precision format around which the Rubin efficiency claims orbit. The significance of 4-bit inference is straightforward: halving the bits per weight roughly halves the memory footprint and doubles the effective memory bandwidth for the same hardware, which is exactly the trade the memory-constrained inference market of 2026 rewards. Nvidia's MLPerf submission applied NVFP4 across weights, attention, and KV cache (see §3.3), and the GTC cost-per-token claims (10x throughput per watt at 1/10 the cost per token vs. Blackwell — vendor claims, unverified) are, at bottom, NVFP4 claims as much as they are silicon claims. The format is part of the platform, not a footnote to it.

---

#### The bandwidth arithmetic, worked through

Because the 260 TB/s figure was the most misreported number of the Rubin launch, here is the arithmetic stated once, completely. The NVL72 scale-up domain provides 260 TB/s of *aggregate* bandwidth — the sum of all GPU-to-GPU paths in the domain. Divided across 72 GPUs, that is approximately 3.6 TB/s per GPU of domain bandwidth. A *per-link* figure would describe a single GPU-to-GPU connection and would be a small fraction of the aggregate. Reporting "260 TB/s per link" therefore overstates the hardware by roughly the number of endpoints in the domain — an error of nearly two orders of magnitude, and one that appeared repeatedly in 2026 coverage.

The same discipline applies to AMD's Helios (§3.6), which quotes the identical 260 TB/s aggregate figure: same rack unit, same aggregate-domain convention, same misreading hazard. When comparing any two vendors' bandwidth claims, the checklist is: (1) aggregate or per-link? (2) scale-up (in-rack) or scale-out (between racks)? (3) measured or claimed? If a source cannot answer all three, the number is not comparable.

#### How to compare two racks: a checklist

The NVL72-vs-Helios comparison that dominated 2026 commentary is a case study in what a fair rack comparison requires — and how rarely it was done. The minimum viable comparison needs five rows, each labeled by evidence status:

| Dimension | Rubin NVL72 | Helios (MI455X) | Evidence status |
|---|---|---|---|
| Compute/rack (FP4) | 3.6 EFLOPS | 2.9 EFLOPS | Both vendor-claimed; Rubin additionally MLPerf-measured |
| Memory/rack | (not disclosed in verified sources) | 31 TB HBM4 | AMD-claimed |
| Scale-up bandwidth | 260 TB/s aggregate | 260 TB/s aggregate | Both vendor-claimed |
| Rack power | ~190–230 kW | (not disclosed in verified sources) | Nvidia-claimed |
| Availability | In production | Announced, unshipped | Verified vs. announced |

The table's real message is its gaps: memory capacity for Rubin and power for Helios were not in the verified sources for this round, which means half the "Rubin vs. Helios" commentary of 2026 was argued from incomplete spec sheets. A reference dossier should show the empty cells, not fill them.

---

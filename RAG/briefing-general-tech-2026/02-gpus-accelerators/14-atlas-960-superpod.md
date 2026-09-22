---
id: briefing-general-tech-2026/02-gpus-accelerators/14-atlas-960-superpod
title: "Atlas 960 SuperPoD: scale, NPO optics and the 2.3x/2.5x claims"
domain: gpus-accelerators
role: deep-dive
task: hardware
actors: ["AMD", "CoreWeave", "Huawei", "Nvidia"]
dates: ["2026-09", "2026-09-17"]
keywords: ["npo", "optics", "superpod", "ascend", "asic", "fp4", "fp8", "gpus", "hbm", "helios", "inference", "rack-scale"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g03-14"
source_lines: [2321, 2394]
canonical_for: ["huawei-ascend"]
sha256: 3f5c9963ea22d11ef5624baf589a2c04d2225674bc3703cedfe09e5373940c62
---

# Atlas 960 SuperPoD: scale, NPO optics and the 2.3x/2.5x claims

<a id="g03-14"></a>
### 3.14 Atlas 960 SuperPoD: scale, NPO optics and the 2.3x/2.5x claims

The Atlas 960 SuperPoD is Huawei's rack-scale — more precisely, *datacenter-scale* — answer to NVL72 and Helios: not a chip, not a rack, but a hundreds-of-cabinets system presented as a single product. Two configurations were detailed in the verification window, and they must not be confused.

#### 220 cabinets: the physical footprint

Two hundred and twenty cabinets is not a cluster but a datacenter hall — rows of liquid-cooled racks with dedicated power distribution, network spine, and service corridors. At 15,488 chips, the full Atlas 960 SuperPoD is closer in scale to a national supercomputer than to a commercial AI cluster: the comparison set is Frontier, El Capitan, and the other exa-scale national systems, not NVL72 deployments. That framing clarifies Huawei's buyer: the SuperPoD is built for state-scale and hyperscale-lab customers, not for enterprise datacenters. It also clarifies the delivery challenge: installing 220 cabinets is a construction project with power, cooling, and logistics phases — the Q4 2027 date is a *build* schedule as much as a *ship* schedule.

#### The two configurations

**The Atlas 960E SuperPoD** — announced **September 17, 2026** at Huawei Connect — is the nearer-term system: **4,096 NPUs, 8 EFLOPS FP8 / 16 EFLOPS FP4, 1 PB of HBM**, built on **NPO (near-packaged optics) Hi-ONE** optical interconnects at **7.2 Tbit/s** per module.

**The full Atlas 960 SuperPoD** — from **Eric Xu's 2025 roadmap, confirmed in September 2026** — is the flagship: **15,488 chips, 30 EFLOPS FP8 / 60 EFLOPS FP4, 4,460 TB of memory, 34 PB/s of interconnect bandwidth, 220 cabinets**, with delivery targeted for **Q4 2027** (announced/planned).

| Configuration | Chips/NPUs | Compute | Memory | Interconnect | Cabinets | Delivery |
|---|---|---|---|---|---|---|
| Atlas 960E SuperPoD | 4,096 NPUs | 8 EFLOPS FP8 / 16 EFLOPS FP4 | 1 PB HBM | NPO Hi-ONE 7.2 Tbit/s | — | Announced 17/09/2026 |
| Atlas 960 SuperPoD (full) | 15,488 chips | 30 EFLOPS FP8 / 60 EFLOPS FP4 | 4,460 TB | 34 PB/s | 220 | Q4 2027 (announced) |

One correction, stated plainly: a **"120 EFLOPS"** figure for the full system appears in **no Huawei announcement**. It does not exist in the verified record and must not be used. The confirmed FP4 figure for the full system is 60 EFLOPS. The persistence of the 120 EFLOPS number in secondary coverage is a case study in how unverified figures propagate: repeated often enough to feel sourced, sourced nowhere.

#### 1 PB of HBM: petabyte-scale memory, unpacked

The Atlas 960E's 1 PB (petabyte) of HBM — 1,000 terabytes — is difficult to intuit, so the decomposition helps: 4,096 NPUs × ~244 GB each ≈ 1 PB (consistent, within system overhead, with the 960DT's 288 GB per-chip figure from §3.13). A petabyte of high-bandwidth memory in one system means the entire working set of a 10-trillion-parameter model — weights, optimizer states, activations, KV caches — can reside in fast memory simultaneously, which is the precondition for Huawei's 2.3x/2.5x system claims: the gains come from never waiting on slower tiers. The number also reframes the per-chip gap (§3.15): Huawei compensates for weaker chips partly by simply having *more* fast memory in the system. Memory capacity at system scale is doing work that per-chip FLOPS cannot.

#### NPO optics: the verified figures and their scope

The optical numbers Huawei disclosed apply to the **960E's 4,096-NPU configuration only** — they must not be transposed onto the full 15,488-chip system:

| NPO figure (960E, 4,096 NPUs) | Value |
|---|---|
| Hi-ONE optical modules | ~5,500 (replacing ~48,000 classic 800G modules) |
| Power saved | >550 kW |
| MTBF | Doubled |
| Availability | 99.8% (<18h unplanned downtime/year) |
| Per-module rate | 7.2 Tbit/s |

NPO — moving the optical engine adjacent to the switch ASIC rather than in a pluggable module — attacks the power and density wall that copper and pluggable optics hit at SuperPoD scale. At 4,096 NPUs, the interconnect would otherwise require ~48,000 pluggable 800G modules, each drawing power, each a failure point, each consuming faceplate space. Replacing them with ~5,500 near-packaged engines cuts the module count by nearly 9x, saves over half a megawatt, and — by eliminating the pluggable connector, historically a reliability weak point — doubles MTBF to reach 99.8% availability.

The SCMP reported on September 17 that this is the **first introduction of NPO in a SuperPoD architecture**, with **"mass production of the world's first NPO product with an integrated light source."** The "integrated light source" detail is the manufacturing crux: co-packaged and near-packaged optics have been industry roadmaps for years, but the laser integration — historically done with external laser sources — is what made them hard to manufacture at volume. Huawei's claim to be mass-producing it first is, if it holds through delivery, a genuine interconnect milestone regardless of whose chips sit at the endpoints. It is also the kind of claim the industry will verify quickly: either the modules ship or they don't.

#### UnifiedBus ("Lingqu")

Alongside the optics, Huawei disclosed a unified interconnect fabric that merges **12+ interconnect protocols** and cuts latency from **7 µs to 2 µs**. The protocol count is the point: at SuperPoD scale, the tax is not just bandwidth but the translation overhead between the dozen-odd protocols a heterogeneous system speaks — intra-node, inter-node, storage, management, and the various generations of each. A single fabric collapses that overhead. The 7→2 µs latency reduction is Huawei's headline number for the effect; as with all latency claims for unshipped systems, it awaits measurement, but the architectural direction (fewer protocols, lower translation tax) is unambiguous.

#### 34 PB/s: what the interconnect number means

The full SuperPoD's 34 PB/s of interconnect bandwidth is the system's most load-bearing figure: at 15,488 chips, the fabric must circulate each training step's gradients and each inference step's activations across the entire machine without becoming the bottleneck. Thirty-four petabytes per second is roughly 130x the NVL72's 260 TB/s domain bandwidth — a scale jump that mirrors the chip-count jump (15,488 vs. 72). The ratio is the point Huawei wants noticed: the fabric scales with the system, so the system behaves as one computer rather than 215 racks that happen to share a building. Whether 34 PB/s is achieved, and at what power cost, is unmeasured — but as an architectural target it defines what "SuperPoD" means: the interconnect is not the system's plumbing, it is the system's processor.

#### The 2.3x / 2.5x claims: read the fine print

Huawei claims the Atlas 960 SuperPoD delivers **2.3x training and 2.5x inference** gains — but the comparison, confirmed in the verified sources, is **Atlas 960 SuperPoD versus Atlas 950 SuperPoD**, measured on **10-trillion-parameter models** (training at 16K sequence length; inference at 256K context). Three corrections to common misreadings:

1. The baseline is the **950 SuperPoD** — a *system*, itself announced/planned for Q4 2026 — **not** the Ascend 910C chip. Coverage that frames this as "960 vs 910C" is wrong on both the baseline and the unit of comparison.
2. It is a **system-to-system** comparison, **not chip-to-chip**. The gains compound chip improvements, scale improvements (more chips), interconnect improvements (NPO, UnifiedBus), and software improvements — they cannot be attributed to the 960DT die alone.
3. There are **no independent benchmarks** of either system. Both the claim and the baseline live inside Huawei's announcements. The 950 SuperPoD — the denominator of the fraction — is itself not yet delivered.

The 2.3x/2.5x figures are therefore Huawei's internal generational-improvement claims for its own systems on very large models — meaningful as a statement of the company's ambitions, unverifiable as performance facts until independent measurement exists. The workload choice is also telling: 10-trillion-parameter models at 16K training sequences and 256K inference contexts are frontier-scale workloads that only a SuperPoD-class system can run at all. Huawei is benchmarking its system on the workloads its system was built for — which is fair, and which also means the numbers say nothing about smaller-scale competitiveness.

---

#### The SuperPoD concept: why systems got this big

The SuperPoD is the logical endpoint of the rack-scale trend that defines this chapter. The reasoning runs: if one NVL72-class rack is a single compute domain, then the next unit of scaling is the multi-rack domain — and the unit after that is the datacenter-scale domain. Huawei's 220-cabinet, 15,488-chip system is that third step made product: not a cluster assembled by the customer from racks, but a *system* engineered, delivered, and (in principle) benchmarked as one. The 34 PB/s interconnect figure is the number that makes the concept work or fail — at SuperPoD scale, the fabric is the computer, and a 15,488-chip system with a weak fabric is just 15,488 slow chips.

The comparison with the Western approach is instructive. Nvidia's multi-rack story in 2026 was seven NVL72 racks (504 GPUs, CoreWeave, September) — large, but an order of magnitude smaller than the full Atlas 960 vision, and assembled from rack-scale building blocks rather than engineered as a single system. Huawei is betting that the next scaling step is a *system-design* problem favoring vertical integration, not a *cluster-assembly* problem favoring modularity. It is the same bet, at larger scale, that Huawei's NPO and UnifiedBus investments serve.

#### NPO in one paragraph: what "near-packaged" buys

Conventional datacenter optics are pluggable modules: a self-contained transceiver that plugs into the switch faceplate, converts electrical to optical, and draws its own power. At SuperPoD scale, tens of thousands of pluggables become a power, density, and reliability crisis — each module a failure point, each connector a signal-integrity compromise. Near-packaged optics moves the optical engine next to the switch ASIC, shortening the electrical path, cutting per-bit power, and eliminating the pluggable connector entirely. The trade-off is manufacturability: integrating the light source (the laser) with the engine is hard, which is why NPO has lived on industry roadmaps for years without mass production. Huawei's claim — "mass production of the world's first NPO product with an integrated light source" (SCMP, September 17) — is a claim to have solved the manufacturing crux first. If true, it is a lead in interconnect technology that is independent of the Ascend chips' competitiveness — an asset that retains value even if the per-chip gap (§3.15) never closes.

---


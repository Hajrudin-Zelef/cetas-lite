---
id: briefing-general-tech-2026/09-appendices/01-glossary
title: "Glossary"
domain: appendices
role: appendix
task: reference
actors: ["AMD", "Amazon", "Broadcom", "China", "Credo", "Google", "Huawei", "Intel", "Meta", "Microsoft", "Nvidia", "OCP", "OIF", "Qualcomm", "TSMC", "UALink"]
dates: ["2026-09-21"]
keywords: ["18a", "accelerator", "agi", "ai200", "ascend", "asic", "awq", "blackwell", "chiplet", "coherent optics", "cost per token", "cpo"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g10-1"
source_lines: [9675, 9819]
sha256: ad846b89dcdb743a9c925b1b97b301ac90e0bb701fbb29e5cba66761ca5479d6
---

# Glossary

<a id="g10-1"></a>
### 10.1 Glossary

**NPO (near-packaged optics)** — Optical transceiver modules placed on the board immediately next to the switch ASIC (typically within a few centimeters), keeping the electrical path short without integrating the optics into the ASIC package itself.
It is an intermediate step between pluggable optics and full co-packaging: shorter copper reach than pluggables, but without the manufacturing and serviceability burden of CPO.
In this dossier, NPO is the optics architecture behind Huawei's Atlas 900 SuperCluster A3 SuperPoD figures.
It preserves field-replaceable modules, unlike CPO, which matters to operators wary of servicing co-packaged boards.

**CPO (co-packaged optics)** — Optical engines integrated into the same package or substrate as the switch ASIC, eliminating long electrical traces and the power overhead of pluggable modules.
It is aimed at the bandwidth-density and power limits of next-generation AI fabrics, but remains a manufacturing and serviceability challenge; as of 2026 it is a roadmap item for most vendors rather than a shipping volume technology.
In this dossier, CPO appears chiefly as a roadmap technology invoked in optical narratives rather than as shipping volume hardware.

**LPO (linear pluggable optics)** — Pluggable optical modules built without a DSP chip, using linear driver and transimpedance-amplifier circuitry while relying on the host ASIC's equalization instead.
Lower power and latency than DSP-based modules, at the cost of tighter link-budget constraints and shorter supported reach.
Relevant in the dossier wherever module power dominates rack power budgets.
Its adoption hinges on host-side equalization capability, which is why LPO links are typically specified as a matched host-module pair.

**UALink (Ultra Accelerator Link)** — An open interconnect standard for scale-up AI accelerator fabrics, developed by the UALink Consortium whose founding members include AMD, Broadcom, Cisco, Google, HPE, Intel, Meta and Microsoft.
It is positioned as an open alternative to Nvidia's NVLink for linking accelerators inside a server or rack domain.
As of this dossier's snapshot, products were expected in 2027, not 2026 — a point the sensitive-notes section enforces strictly.
The consortium published its 1.0 specification in 2025, with silicon and systems from member companies expected in 2027.

**NVLink** — Nvidia's proprietary high-bandwidth interconnect for linking GPUs inside a server or rack domain, notably the NVL72 rack-scale domain whose aggregate bandwidth is quoted at 260 TB/s.
The performance reference against which open alternatives such as UALink are measured, and a key element of Nvidia's platform lock-in.
Each NVLink generation has roughly doubled per-link bandwidth, and the interconnect is a major differentiator of Nvidia's rack-scale systems.

**EFLOPS** — Exa floating-point operations per second: 10^18 FLOPS.
The unit used to express the aggregate compute of large AI clusters, including the Huawei SuperPoD figures (8/16 and 30/60 EFLOPS-class configurations) cited in this dossier.
Precision (FP16, FP8, INT8) should always accompany an EFLOPS figure; without it, the number is meaningless.
Dossier figures name the precision where the source provided it; dense FP8 EFLOPS and sparse figures are not interchangeable.

**KV cache** — The cached key and value tensors of a transformer's attention layers, stored during token generation so that previously processed tokens are not recomputed at each decoding step.
Its size grows with context length and batch size, making it one of the main memory consumers — and cost drivers — of long-context inference, and a central parameter of token economics.
Techniques such as paged attention, prefix caching and quantization exist specifically to shrink its footprint.

**Disaggregated inference** — An inference-serving architecture that runs the prefill and decode phases on separate pools of hardware, each sized to its phase's bottleneck (compute for prefill, memory bandwidth for decode).
It improves utilization and cost per token on large deployments and underpins much of the 2026 discussion about inference-cost reduction.
It pairs naturally with phase-aware scheduling and is a recurring theme in the 2026 inference-cost literature.

**Prefill / decode** — The two phases of LLM inference.
*Prefill* processes the whole input prompt at once: compute-bound and highly parallel.
*Decode* generates output tokens one at a time, each step depending on the previous one: memory-bandwidth-bound and latency-sensitive.
The distinction drives hardware sizing, batching strategies, KV-cache management and per-token pricing models.
Batching strategies such as continuous batching and chunked prefill exist to keep both phases utilized on shared hardware.

**1600ZR / 800ZR** — Coherent optical transceiver specifications from the OIF for 800 Gb/s and 1.6 Tb/s transmission over data-center-interconnect distances.
The "ZR" family defines interoperable coherent pluggables for DCI links, letting operators buy metro optics from multiple vendors.
A distinct market from the short-reach multimode optics inside the data center.
Multi-vendor interoperability testing is the practical hurdle these OIF specifications are designed to clear.

**DCI (data center interconnect)** — The fiber links connecting data centers to each other, typically metro to long-haul distances, including subsea segments.
Served by coherent optics (ZR/ZR+ class) and muxponder-based DWDM systems; a distinct market with its own vendors, analysts and operators from intra-data-center networking.
The ZR/ZR+ pluggable classes discussed in the dossier target exactly this segment.

**ASIC** — Application-specific integrated circuit: a chip designed for one workload rather than general-purpose computing.
In this dossier the term covers custom AI accelerators (Google TPU, AWS Trainium, Meta MTIA, Huawei Ascend, startup inference chips) as opposed to merchant GPUs — the "build vs. buy" axis of the 2026 silicon story.
The dossier's custom-silicon chapters treat ASICs as the strategic alternative to merchant GPUs for hyperscalers and sovereign programs.

**HBM (high bandwidth memory)** — 3D-stacked DRAM placed next to the compute die in the same package, delivering far higher memory bandwidth than GDDR or DDR at much higher cost per gigabyte.
Successive generations — HBM3, HBM3E, HBM4 — are a defining cost and supply item of AI accelerators; HBM allocation was one of the pressure points of the 2025–2026 memory crunch, and HBM stack counts are a standard way to compare accelerators.
Each HBM generation raises per-stack bandwidth and capacity; stack counts per accelerator are a standard comparison metric in the dossier.

**LPDDR (low-power DDR)** — Low-power variants of DDR memory, historically for mobile devices and increasingly used in AI inference hardware where power efficiency matters more than raw bandwidth.
LPDDR5X is the current high-end generation.
Some 2026 inference cards (e.g., Qualcomm's AI200) use LPDDR without disclosing the exact variant — the dossier names only "LPDDR" in such cases.
Its lower operating voltage trades peak bandwidth for power efficiency, which suits inference cards with strict thermal envelopes.

**SerDes** — Serializer/deserializer: the high-speed electrical signaling blocks that move data between chips, packages and boards.
SerDes lane rates (112G, 224G generations) set the ceiling for chip-to-chip and chip-to-optics interconnects, and SerDes IP is a core asset of connectivity vendors such as Credo and Broadcom.
The 224G-SerDes generation underpins both 1.6T optics and next-generation accelerator fabrics.

**DSP (in optical modules)** — The digital signal processor chip inside a coherent or high-rate optical module, handling equalization, forward error correction and impairment compensation.
It is typically the most power-hungry part of a DSP-based module, which is why DSP-free LPO designs exist — trading reach and margin for watts.
Removing it (LPO) shifts the equalization burden to the host ASIC, which must then be designed for the purpose.

**Muxponder** — A device combining multiplexing and transponder functions: it aggregates multiple lower-rate client signals (e.g., 8×100G) into one higher-rate DWDM wavelength for transport.
A common building block of DCI and metro optical networks, sitting between the router and the line system.
In DCI architectures it sits between the router's client ports and the DWDM line system.

**RoCE (RDMA over Converged Ethernet)** — A protocol enabling remote direct memory access over Ethernet, widely used for AI cluster networking as an alternative to InfiniBand.
Backend AI fabrics in 2026 are largely an Ethernet (RoCE, Ultra Ethernet) vs.
InfiniBand contest, with Broadcom silicon on the Ethernet side.
Its 2026 momentum is tied to industry efforts to harden Ethernet for AI-scale congestion control and telemetry.

**Chiplet** — A small, purpose-built die combined with other chiplets in a single package (via advanced packaging such as 2.5D interposers or 3D stacking) to build a large chip that would be uneconomical as one monolithic die.
The standard construction method for modern CPUs, GPUs and AI accelerators, and the enabler of mixing logic, HBM and I/O dies from different process nodes.
Advanced packaging (2.5D interposers, 3D stacking) is the enabling technology, and packaging capacity is itself a supply-chain chokepoint.

**E-cores / P-cores** — Intel's hybrid-CPU terminology: *efficiency cores* (throughput per watt, background and parallel work) vs.
*performance cores* (single-thread speed, latency-sensitive work), with a hardware thread director assigning work between them.
The vocabulary of Intel's client chips, including the Core Ultra Series 3 covered in this dossier.
The hybrid approach lets one die cover both burst responsiveness and sustained multi-threaded throughput.

**Node (process node)** — A generation of semiconductor manufacturing technology, e.g., Intel 18A or TSMC's 2 nm-class process.
Node names are commercial labels, not literal transistor dimensions; density, power and SRAM scaling vary between foundries at the "same" node, so cross-vendor node comparisons are approximate by construction.
Intel 18A and TSMC N2-class processes are the leading-edge references of the dossier's 2026 foundry coverage.

**Rack-scale** — Designing the full rack — compute, networking, power delivery and cooling — as a single engineered system rather than assembling individual servers.
Nvidia's NVL72 and AMD's Helios are rack-scale AI platforms.
Rack power is the binding facility constraint of the era: roughly 100–120 kW for Blackwell GB200-generation racks, moving toward ~190–230 kW class designs.
Facility power and liquid-cooling readiness, not chip supply alone, increasingly gate rack-scale deployments.

**Token economics** — The unit economics of AI services expressed per token: inference cost per million tokens versus price per million tokens charged to customers.
The central business-model question of 2026, as inference costs fell steeply (the McKinsey 85–95% figure) while memory and power costs did not — compressing margins for some providers and opening them for others.
Falling cost per token does not automatically mean falling prices: the spread funds model R&D, capacity and margin.

**Inference vs. training** — *Training* builds a model from data: massive, one-off, compute-intensive, tolerant of batching and checkpointing.
*Inference* runs a trained model to serve requests: latency-sensitive, memory-bandwidth-bound in the decode phase, and a continuous cost.
The two workloads stress different parts of the system, follow different scaling economics, and increasingly run on different hardware.
The dossier's hardware chapters are organized around this split, since it determines which bottlenecks matter.

**Quantization** — Reducing the numerical precision of model weights and activations (e.g., FP8, INT4) to cut memory footprint and speed up inference, with small accuracy trade-offs.
Microscaling formats such as NVFP4 (Nvidia) and MXFP4 (from the OCP Microscaling specification) define 4-bit block-scaled representations for efficient low-precision inference.
Quantization is one of the documented levers behind the 2026 inference-cost collapse.
The dossier treats it as one of the verified levers behind the 2026 collapse in inference costs.

**AWQ / GPTQ** — Post-training quantization methods that compress LLM weights to low precision without retraining.
AWQ (Activation-aware Weight Quantization) protects the small fraction of "salient" weights that matter most for accuracy; GPTQ uses approximate second-order (Hessian) information for layer-wise quantization.
Both are standard tooling for serving quantized open-weight models, and both appear in the dossier's treatment of efficient inference.
They are post-training methods — no retraining required — which is why they spread so fast through the open-weight ecosystem.

**MoE (mixture of experts)** — A sparse model architecture in which each token is routed to a small subset of "expert" sub-networks rather than the full model.
Total parameter counts can be very large while the active compute per token stays modest — the standard architecture of most frontier open-weight models in 2026, and the reason "parameter count" alone no longer describes serving cost.
Routing overhead and expert-parallel communication are the engineering costs of the architecture's efficiency gains.

**AGI (artificial general intelligence)** — AI systems matching or exceeding human capability across a broad range of cognitive tasks.
A contested, moving definition; in 2026 it functions as much as a strategic and regulatory concept (export controls, safety institutes, corporate charters, international declarations) as a technical milestone.
The dossier uses the term only where its sources use it, and never as a dated prediction.

**RSI (recursive self-improvement)** — A scenario in which an AI system improves its own design, potentially triggering rapid capability jumps.
A long-standing concern in AI-safety literature and one of the risk models tracked by safety institutes; relevant to the dossier's September governance chapters.
It remains a theoretical risk model in 2026, not an observed phenomenon.

**AI safety institute** — Government bodies created to evaluate frontier AI systems and coordinate mitigations for catastrophic risks (the UK AI Safety Institute, the US AI Safety Institute, and counterparts announced from 2023–2024 onward).
They commission capability evaluations, publish risk assessments and feed into international declarations such as the 22-country text adopted on 21/09/2026.
Their evaluation results feed directly into the governance chapters' September timeline.

**Export controls** — Government restrictions on exporting advanced semiconductors, manufacturing equipment and related technology — notably US controls limiting the sale of high-end AI accelerators and chipmaking tools to China.
A structural driver of China's domestic-accelerator programs (Huawei Ascend, domestic HBM efforts) throughout this dossier, and the backdrop to the September US–China AI-governance standoff.
They are also the reason the dossier tracks two parallel supply chains: Western merchant silicon and China's domestic stack.

**"RAMageddon"** — The dossier's shorthand for the 2025–2026 memory supply crunch: DRAM contract and spot prices surged (consumer DDR5 SKUs roughly ×3.5 from late 2025 into Q1 2026), HBM allocation tightened, and memory — not GPUs — became the binding constraint on AI buildouts, in the memory makers' own telling.
The term captures the moment the industry's bottleneck moved up the stack from compute to memory.
The term is the dossier's own coinage for the episode, not an industry-standard label.


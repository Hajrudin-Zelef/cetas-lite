---
id: briefing-general-tech-2026/00-front-matter/03-thematic-keywords
title: "Thematic keywords (source)"
domain: front-matter
role: reference
task: reference
actors: ["AMD", "China", "Huawei", "Nokia", "Nvidia", "OIF", "Qualcomm", "Telxius", "UALink"]
dates: []
keywords: ["accelerator", "agentic", "agi", "ai200", "asic", "capex", "chiplet", "coherent optics", "cost per token", "dci", "decode", "disaggregated"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g01-keywords"
source_lines: [322, 383]
sha256: af6c48a7cbbbf2d77cac465ee02a71e17a83326166b0f636981c701418487f18
---

# Thematic keywords (source)

<a id="g01-keywords"></a>
### Thematic keywords

- **AI accelerator** — a chip or system designed to execute neural-network mathematics (matrix multiplications and related kernels) far faster and more efficiently than a general-purpose CPU.
  In this dossier it covers datacenter GPUs, custom ASICs and dedicated inference processors alike (chapter 3).
  *See also: ASIC, HBM.*
- **Agentic infrastructure** — the compute, networking and software stack built to run autonomous AI agents at production scale, rather than single-turn chat workloads.
  It is the cross-cutting shift analyzed in §9.4.
  *See also: token economics.*
- **ASIC** — an application-specific integrated circuit: a chip designed for one workload instead of general-purpose programmability, trading flexibility for efficiency.
  Custom ASICs designed by hyperscalers are gaining ground against merchant GPUs (§4.6).
  *See also: AI accelerator, hyperscaler.*
- **Chiplet** — a small modular silicon die combined with others in a single package to build a large processor.
  Chiplets are the packaging approach behind modern high-end accelerators and server CPUs (chapters 3–4).
  *See also: HBM.*
- **Coherent optics** — optical transmission that encodes data in the phase, amplitude and polarization of light, enabling very high capacity over long distances on a single wavelength.
  Covered via OIF 1600ZR and 800G coherent pluggables (§6.7, §6.12).
  *See also: data center interconnect (DCI).*
- **Data center interconnect (DCI)** — the optical links that connect datacenters to each other across metro and long-haul distances.
  The D7070 muxponder and the Telxius–Nokia deployment sit in this segment (§6.3, §6.12).
  *See also: coherent optics.*
- **Disaggregated inference** — splitting LLM serving into separate prefill (prompt processing) and decode (token generation) stages, each running on hardware sized for its own profile.
  Nvidia Dynamo is the reference implementation discussed in §5.3.
  *See also: prefill/decode, rack-scale.*
- **Export controls** — government restrictions on selling advanced chips, tools and technology to certain destinations.
  They shape the US–China stack competition analyzed in §9.3.
- **FP8/INT4 quantization** — reducing the numerical precision of model weights and activations to 8-bit floating point or 4-bit integer, cutting memory footprint and cost per token.
  FP8 is the safe production standard (§5.8); INT4-class formats are the aggressive standard (§5.9).
  *See also: KV cache, token economics.*
- **HBM** — high-bandwidth memory: DRAM stacks placed on the same package as the accelerator for very high memory bandwidth.
  It defines dense training racks (§5.4) and is central to AMD's MI455X memory play (§3.6).
  *See also: chiplet, KV cache.*
- **Hyperscaler** — an operator of massive cloud datacenter fleets.
  Their ~$725–730 billion of 2026 capex sets the tempo of the whole hardware cycle (§4.5).
  *See also: ASIC, token economics.*
- **KV cache** — the stored key and value states of already-processed tokens, the dominant memory cost of long-context inference.
  TurboQuant compresses it to 3 bits (§5.13).
  *See also: disaggregated inference, FP8/INT4 quantization.*
- **Linear pluggable optics (LPO)** — pluggable optical modules that remove the digital signal processor, trading some link margin for lower power and latency.
  Covered in §6.9.
  *See also: near-packaged optics (NPO), coherent optics.*
- **Near-packaged optics (NPO)** — optical engines co-packaged close to the switch ASIC rather than in front-panel pluggables, shortening electrical reach.
  Huawei uses NPO in the Atlas 960 SuperPoD (§3.14).
  *See also: linear pluggable optics (LPO).*
- **NVLink/UALink** — high-bandwidth GPU-to-GPU interconnects: Nvidia's proprietary NVLink versus UALink, the open-industry coalition's alternative.
  The coalition against NVLink is covered in §4.10.
  *See also: rack-scale.*
- **Open weights** — the practice of publicly releasing a model's trained weights.
  The Hugging Face precedent matters to the recursive self-improvement debate (§8.6).
  *See also: recursive self-improvement (RSI).*
- **Prefill/decode** — the two phases of LLM inference: processing the input prompt (prefill) and generating output tokens one by one (decode).
  Their very different compute profiles motivate disaggregated serving (§5.3).
  *See also: disaggregated inference, KV cache.*
- **Rack-scale** — designing AI compute as a complete, integrated rack system (compute, networking, power, cooling) rather than as standalone servers.
  NVL72, Helios and the AI200 are rack-scale systems (chapter 3).
  *See also: NVLink/UALink, disaggregated inference.*
- **Recursive self-improvement (RSI)** — the prospect of AI systems improving their own capabilities without human direction.
  It is a central governance concern in the AGI debate (§8.6).
  *See also: open weights.*
- **Token economics** — the cost-and-revenue arithmetic of AI services expressed per generated token.
  Cost per token is treated as the decisive metric of the inference era (§5.6).
  *See also: FP8/INT4 quantization, hyperscaler.*

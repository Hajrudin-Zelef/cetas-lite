---
id: briefing-general-tech-2026/09-appendices/02-actors-index
title: "Actors index"
domain: appendices
role: appendix
task: reference
actors: ["AMD", "Amazon", "Anthropic", "Apple", "Broadcom", "Cerebras", "China", "CoreWeave", "Credo", "Crusoe", "FS.com", "Fujitsu", "Gartner", "Google", "Groq", "Huawei", "IDC", "Intel", "MACOM", "MLCommons", "Meta", "Micron", "Nebius", "Nokia", "Nvidia", "OCP", "OIF", "OpenAI", "Positron", "PrismML", "Qualcomm", "SK Hynix", "Samsung", "TSMC", "Taalas", "Telxius", "TrendForce", "UALink", "UN"]
dates: ["2025-12-22", "2026-05-12", "2026-06", "2026-07", "2026-07-14", "2026-09", "2026-09-03", "2026-09-16", "2026-09-21", "2026-09-23"]
keywords: ["18a", "accelerator", "ai200", "ascend", "asic", "benchmark", "blackwell", "capex", "claude", "cost per token", "custom silicon", "dci"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g10-2"
source_lines: [9820, 9948]
sha256: 525468639d0b1d769d9fc543fb7836a6113676115ada8909ad94ad780e55db61
---

# Actors index

<a id="g10-2"></a>
### 10.2 Actors index

This index covers every company, laboratory, consortium and research firm named in the dossier's February–September 2026 narrative.
Each entry states what the actor is, then its specific role in the dossier's events, ordered as they appear across the narrative.
The same scope discipline as the main text applies throughout: announced versus shipped, vendor claim versus verified fact, projection versus confirmed result.

**Nvidia** — The dominant AI-accelerator vendor and the gravitational center of the 2026 infrastructure narrative.
In this dossier: the Rubin platform's preview MLPerf submission (3.7× on Qwen3-VL, 2.5× on DeepSeek-R1 versus Blackwell — a "preview submission," not a final result), Nvidia's own cost-per-token claims for Rubin, the 260 TB/s aggregate-bandwidth figure of the NVL72 domain, and the 100–120 kW figure for Blackwell GB200 racks.
Nvidia is also the reference point for FS.com's "100% verified on NVIDIA" vendor claim and the proprietary incumbent that the UALink consortium's open interconnect is organized against.

**AMD** — The second merchant AI-accelerator vendor and Nvidia's principal challenger.
In this dossier: the MI400 accelerator and the Helios rack-scale platform, the alliance with Anthropic, and the crossing of a $1 trillion market capitalization on 21/09/2026 (close around $1,005B) during a sector-wide rally — a date the dossier fixes precisely, against looser "September" framings.

**Intel** — The x86 incumbent, pursuing a foundry business and AI accelerators in parallel with its client and server CPUs.
In this dossier: the Core Ultra Series 3 client lineup including the premium X9 tier (whose existence was initially denied in drafting, then corrected), the 18A process node as a foundry milestone, and the E-core/P-core hybrid terminology that structures its product story.

**Huawei** — The Chinese telecom and semiconductor group at the center of the dossier's sovereignty chapters.
In this dossier: the Ascend accelerator line and the SuperPoD scale-up systems — the 4,096-NPU supernode (Atlas 960E) versus the 15,488-NPU full cluster as *distinct products*, with the NPO optics figures (5,500 vs.
48,000 modules, >550 kW, 99.8%) tied to the 4,096-NPU configuration only.
Also: system-vs-system 2.3×/2.5× claims against the 950 SuperPoD; a per-chip gap of roughly two years behind Blackwell and ~10× under Rubin with no independent benchmarks; PyTorch support via the out-of-tree TorchNPU integration (in-tree native support planned, not delivered); and Nikkei's "more than 10" chipset portfolio figure, with no model-by-model list.

**Apple** — The consumer-hardware giant and custom-silicon designer.
In this dossier: the September event that launched the iPhone 18 Pro (not an "iPhone 17"); TrendForce's ranking tying Apple with Huawei at 24.8% each (not Apple alone in second place); and The Information's 16/09/2026 report of M8 Ultra-based enterprise AI servers targeted for 2029 — a rumor, not an Apple announcement, and the dossier never presents it as one.

**Samsung** — The memory leader and foundry competitor.
In this dossier: its commercial campaigns in the memory space, characterized as offensive rather than defensive, with the associated market-share figures treated as projections rather than confirmed results — part of the broader "RAMageddon" memory narrative.

**Google (incl.
DeepMind)** — The hyperscaler, TPU designer and frontier-model lab.
In this dossier: the DeepMind Institute for AI Safety, led by Shane Legg together with Demis Hassabis and James Manyika (three names, not one); and the "Googlebook" device announced on 12/05/2026, for which "Aluminium" is an internal codename, not the commercial name.

**Qualcomm** — The mobile-SoC leader moving into data-center AI inference.
In this dossier: the AI200 inference accelerator, described with "LPDDR" memory only — the exact variant (e.g., LPDDR5X) is undisclosed, per-card bandwidth is undisclosed, and 2026 shipping is unconfirmed.
The dossier attaches all three qualifiers to every mention.

**Fujitsu** — The Japanese IT and computing conglomerate (enterprise systems, HPC, custom ARM processors).
In this dossier it appears as part of the enterprise and sovereign-compute landscape against which the 2026 infrastructure buildout is measured — a reference incumbent rather than a protagonist of any single event.

**FS.com** — The optical-transceiver and network-hardware vendor.
In this dossier: its 22/12/2025 marketing claim of "100% verified on NVIDIA" — a vendor claim, not an Nvidia certification — and a 100G+ segment growth figure that could not be independently verified and is therefore excluded from the dossier's factual claims.

**Meta** — The hyperscaler and open-weight model publisher.
In this dossier it figures as a UALink Consortium founding member and as one of the large AI-infrastructure buyers whose capex decisions and model releases shape the demand side of the 2026 story.

**OpenAI** — The frontier-model company.
In this dossier: the GPT-6 Astra launch on 03/09/2026 (not July 2026); the 21/09/2026 institutional position issued through Global Affairs (not Sam Altman's personal words); and Altman's UN Security Council briefing, *planned* for 23/09/2026 — after the dossier's snapshot date, and therefore reported as planned, never as accomplished.

**Anthropic** — The frontier-model company and major AWS partner.
In this dossier: the Claude model line's trajectory through the spring "model war," and the alliance with AMD around MI400/Helios, which gives the AMD challenge a flagship software partner.

**Micron** — The US memory maker (DRAM, NAND, HBM).
In this dossier it stands on the memory-makers' side of the "memory is the dominant constraint" finding — the dossier confirms the substance as an industry-side statement while noting that the exact sentence is an editorial interpretation, not a direct quote.

**SK Hynix** — The Korean memory leader and key HBM supplier to Nvidia.
Same dossier role as Micron: a primary source for the memory-constraint narrative and the HBM allocation pressure of 2025–2026, at the heart of the "RAMageddon" chapters.

**TSMC** — The world's leading chip foundry, manufacturing for Nvidia, AMD, Apple, Qualcomm and Broadcom among others.
In this dossier it is the indispensable upstream of the accelerator supply chain and the reference for 2 nm-class process technology — the capacity allocator everyone else depends on.

**Broadcom** — The networking-silicon leader (Ethernet switching, AI-fabric silicon) and custom-AI-ASIC designer.
In this dossier: a UALink Consortium founding member and a central actor in the Ethernet-vs-InfiniBand fabric contest for AI clusters, where its merchant silicon sets the pace for the RoCE camp.

**Credo** — The connectivity-chip company specializing in SerDes, retimers and DSPs for data-center links.
In this dossier it represents the merchant-silicon layer of the 800G/1.6T optical and copper interconnect ramp — the unglamorous chips without which neither scale-up nor scale-out fabrics work.

**MACOM** — The analog and photonic semiconductor supplier (optical components, laser drivers, TIAs).
In this dossier it figures in the optical-module supply chain behind ZR-class and AI-fabric transceiver volumes.

**Nokia** — The telecom-equipment vendor with a large optical-networking business.
In this dossier it appears in the DCI and carrier-optics context, where coherent ZR-class pluggables are reshaping metro and long-haul architectures and displacing traditional transponder shelves.

**Telxius** — The telecom-infrastructure operator (subsea cables, terrestrial backbones).
In this dossier it illustrates the DCI layer: the physical fiber assets carrying inter-data-center traffic, and the capacity market underneath the AI buildout.

**CoreWeave** — The GPU-focused neocloud.
In this dossier it exemplifies the 2026 neocloud boom: specialized AI-cloud capacity sold against hyperscaler shortages and financed by the year's capital deluge.

**Crusoe** — The AI-cloud and data-center operator.
Same dossier role as CoreWeave: a neocloud scaling AI capacity outside the traditional hyperscalers, part of the diversified-supply story of 2026.

**Nebius** — The AI-infrastructure company built on the former Yandex N.V. assets.
Same dossier role: European-anchored neocloud capacity in the 2026 buildout, adding a geographic-diversification angle.

**Positron** — The startup building FPGA-based AI inference systems.
In this dossier it represents the alternative-silicon inference camp alongside Groq, Cerebras and Taalas — the "attack the cost per token from below the GPU" thesis.

**Taalas** — The AI-ASIC startup designing LLM-specific inference chips.
Same dossier role as Positron: custom silicon purpose-built for the decode phase, betting that inference economics reward specialization.

**Groq** — The inference-chip company (LPU architecture) selling low-latency token generation as a service.
In this dossier it is one of the reference points for the disaggregated-inference and token-economics discussion — speed as a product, not just FLOPS.

**Cerebras** — The wafer-scale AI-chip company.
In this dossier it represents the most radical alternative compute substrate in the inference-vs-training hardware debate: an entire wafer as one chip, trading generality for memory bandwidth.

**PrismML** — The ML startup that launched Bonsai on 14/07/2026.
Initially suspected to be a hallucination during verification; the re-verification pass confirmed the launch as real — the dossier's standing example of why the suspicion pass exists.

**MLCommons** — The industry consortium running the MLPerf benchmark suites.
In this dossier: the venue of Rubin's preview submission and the institutional reference for how inference-benchmark claims must be read — preview vs. official submissions, with all the caveats that distinction carries.

**OIF (Optical Internetworking Forum)** — The standards body behind the 800ZR/1600ZR coherent pluggable specifications.
In this dossier it is the normative reference for DCI optics claims: if a ZR figure is cited, OIF is the standard it is measured against.

**OCP (Open Compute Project)** — The open-hardware consortium standardizing data-center designs, including the MXFP4/MXFP8 microscaling formats.
In this dossier it anchors the open-hardware and quantization-format discussion, and the UALink-adjacent open-ecosystem narrative.

**IDC** — The technology market-research firm.
In this dossier it is one of the analyst sources for market sizing, cited with the standard caution: analyst figures are estimates and projections, not confirmed results.

**TrendForce** — The Taiwan-based market-research firm focused on memory and consumer electronics.
In this dossier: the source tying Apple and Huawei at 24.8% each in the relevant smartphone ranking — the correction of the "Apple solo second" misreading comes directly from its data.

**Gartner** — The technology advisory and research firm.
In this dossier it appears among the analyst sources framing enterprise AI-adoption and infrastructure-spending estimates.

**LightCounting** — The market-research firm specializing in optical communications.
In this dossier it is the reference analyst source for optical-transceiver volumes and the 800G/1.6T ramp — the numbers behind the optics chapters.

**Cignal AI** — The market-research firm covering optical hardware.
Same dossier role as LightCounting: an analyst source for coherent-optics and DCI market figures.

**Counterpoint** — The market-research firm covering smartphones and semiconductors.
In this dossier it appears among the analyst sources for device-market readings, including the iPhone 18 cycle.

**McKinsey** — The management consultancy.
In this dossier: its June 2026 report "The technology shifts reducing AI inference costs," the source of the 85–95% inference-cost-reduction figure — initially suspected to be a hallucination, confirmed real by the re-verification pass, and now one of the dossier's load-bearing statistics.


---
id: etape6-phasec-optics-cabling/00-front-matter/part-41
title: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure (part 41)"
domain: front-matter
role: reference
task: reference
actors: ["Broadcom", "Intel", "Nvidia"]
dates: ["2023-01", "2023-03", "2023-04", "2023-04-27", "2025-06-04", "2025-10-08", "2026-02", "2026-02-20", "2026-09", "2026-09-22"]
keywords: ["optics", "asic", "cpo", "ethernet", "gpus", "hbm", "intel", "latency", "nvidia", "rack-scale", "serdes"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [1492, 1509]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: 59a515a8af7fa176e6f4eac8ecd1571911e7133c12914a534492bc85619b24d2
---

# Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure (part 41)

**Broadcom Tomahawk 6 (102.4T generation).** 102.4 Tbps per ASIC with 100G or 200G PAM4 SerDes; silicon-level radix examples 64 × 1.6TbE / 128 × 800GbE / 256 × 400GbE / 512 × 200G / 1,024 × 100G [vendor-reported] (storagenewsletter.com, 2025-06-04). Shipping announced 2025-06-04; by September 2026 shipping with named OEM systems (Arista 7060XE7, HPE Juniper QFX5250) — no longer pre-announcement [independent]. Feature claims: scale-up and scale-out Ethernet, Cognitive Routing 2.0, arbitrary Clos/rail/torus topologies, long-reach passive copper support, optional co-packaged-optics (CPO) implementation [vendor-reported]. **Tomahawk 6 Davisson CPO** announced 2025-10-08, described as the industry's first 102.4 Tbps Ethernet switch with co-packaged optics [official] (globenewswire.com). Independent coverage of Arista's Tomahawk 6 system family corroborates the 1.6T-rack-scale positioning for AI infrastructure [independent] (networkworld.com).

**Broadcom Jericho3-AI (+ Ramon3 fabric element).** Purpose-built for AI/ML backend fabrics rather than shallow-buffer ToR switching; 28.8 Tbps network-facing capacity reported [independent] (networkworld.com). One account reports 144 × 106G PAM4 network-facing SerDes (18 × 800GbE / 36 × 400GbE / 72 × 200GbE); a deeper account reports 304 total SerDes: 144 network/downlink lanes plus 160 fabric-facing lanes toward Ramon3 [independent] (nextplatform.com, 2023-04-27). Positioned for fabrics connecting more than 32,000 GPUs, with traffic spraying/reordering, end-to-end scheduling, high radix, and vendor-claimed sub-10 ns automatic path convergence [vendor-reported] (telecomtv.com). Broadcom said Jericho3-AI was available to qualified customers in April 2023 [independent].

**Cisco Silicon One — G200, P200, G300.**
- **G200**: 51.2 Tbps, 5 nm, 512 × 112G-class SerDes; used in Cisco Nexus 9364E-SG2 systems for high-density 800G AI/ML and web-scale fabrics [official] (cisco.com N9364E-SG2-Q hardware installation guide).
- **P200**: powers Nexus N9364E-SP2R (64 × 800G) with 144 MB on-die shared buffer plus 16 GB HBM; Cisco's system datasheet reports 102.4 Tbps bidirectional = 51.2 Tbps one-way port bandwidth — must not be compared directly with unidirectional ASIC figures [official] (cisco.com N9364E-SP2R datasheet).
- **G300**: announced February 2026, 102.4 Tbps, with new system N9364F-SG3 reported as 64 ports supporting 800G/1.6T; syndicated announcement copy describes industry-leading on-chip buffering, but exact buffer size, SerDes, form factor, power, and commercial ship date still need official-document verification [vendor-reported] (storagenewsletter.com, 2026-02-20; aapnews syndication).
- Context: Cisco's UADP ASIC family powers the Catalyst 9000 enterprise/campus line; the 800G data-center Nexus 9000 systems above use Silicon One, not UADP — UADP is therefore not the current 800G data-center switching ASIC [independent].

**Marvell Teralynx 10.** Programmable 51.2 Tbps, 5 nm Ethernet switch ASIC integrating 512 × 112G LR SerDes; headline configurations 64 × 800GbE / 128 × 400GbE / 32 × 1.6T at silicon level [official] (marvell.com Teralynx 10 product brief). Previewed March 2023, initially sampling-guided Q2 2023; by mid-2026 Marvell reported production/customer deployment and volume production, plus SONiC support [vendor-reported] (convergedigest.com). Claimed latency as low as 500 ns, sub-600 ns across packet sizes in specified forwarding modes [independent] (servethehome.com). Feature claims: congestion-aware routing, permutable flex-forwarding, Flashlight telemetry, P4 INT, large buffers, zero-latency-loss programmability [independent] (networkworld.com roundup).

**NVIDIA Spectrum-4 (and Spectrum-X800 platform).** Custom NVIDIA 51.2 Tbps, 4 nm Ethernet ASIC using 100G PAM4 SerDes; configurations 64 × 800GbE / 128 × 400GbE / 256 × 200GbE [independent] (convergedigest.com). SN5600-class systems: 160 MB shared packet buffer, 12.8 Tbps crypto engine, MACsec and VXLANsec, adaptive routing, high-precision congestion control and telemetry, programmable parser/pipeline [vendor-reported] (NVIDIA Spectrum-X800 solution brief). **Spectrum-X800 is a platform, not just a chip**: combines Spectrum-4 switches with NVIDIA SuperNICs and software; full NVIDIA adaptive-routing behavior depends on the NVIDIA host-side stack and should not be attributed to a bare switch paired with commodity NICs [independent]. NADDOD N9570-128QC (128 × 400G, Spectrum-4, 51.2 Tbps, enterprise SONiC) presented February 2026 as a Spectrum-4 whitebox alternative pairing with ConnectX-8 SuperNICs for RoCE adaptive routing [secondary] (naddod.medium.com). A Dell datasheet mirror describes next-generation liquid-cooled CPO family — SN6800-LD (409.6 Tb/s across 4 × 102.4 Tb/s modules, 512 MMC-12 800G optical interfaces, 5RU), SN6810-LD (128 × 800G, 102.4 Tb/s, 2RU), SN6600-LD (128 × 800G via 64 OSFP, 102.4 Tb/s, 2RU) — 48 VDC/54 VDC busbar power, 160 MB packet buffer per 102.4T module; indicates 102.4T-class NVIDIA silicon in Dell's pipeline, though the ASIC name is not stated in the visible excerpt [vendor-reported]. No official evidence found for an "SN5800" system as of 2026-09-22 — do not treat as announced [unverified].

**Intel Tofino — status: exited.** Intel stopped future investment/development in its network-switching product line in January 2023, continuing to support existing products and customers [independent] (nasdaq.com). Tofino 2 topped out at 12.8 Tbps with 32 × 400G; not an 800G-generation contender; Intel's product page labels the 6.4 Tbps Tofino part retired/discontinued [official] (intel.com). No credible evidence of a revived 800G Tofino roadmap as of September 2026 [independent].

**Juniper/HPE silicon note; Arista silicon strategy.** Arista's high-end 800G/1.6T systems use merchant silicon: Tomahawk 5 in fixed 7060X6 class, Jericho3-AI in deep-buffer/modular AI fabrics, Tomahawk 6 in 7060XE7 — evidence-based product mapping, not an absolute claim about every Arista internal component [independent] (arista.com 7060X6 series). Juniper's current data-center ASIC story in the 800G systems below is merchant (Tomahawk 5/6) plus Trident 5 in QFX5140; a fresh primary source for a current-generation Juniper Express5/Trio 800G data-center ASIC was not captured — treat Juniper custom-silicon 800G claims as a gap [unverified].


---
id: etape6-phasec-optics-cabling/00-front-matter/10-3-nvidia-quantum-x800-infiniband-protocol-infiniband-xdr-
title: "10.3 NVIDIA Quantum-X800 InfiniBand (protocol: InfiniBand XDR — not Ethernet)"
domain: front-matter
role: reference
task: actor-profile
actors: ["Broadcom", "Huawei", "Nvidia", "Oracle"]
dates: ["2026-06"]
keywords: ["ethernet", "nvidia", "accelerator", "asic", "dci", "full-duplex", "gpu", "inference", "latency", "optics", "pricing", "serdes"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [1533, 1558]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: 8f249d46b51aafe529aae903a8de880bdd9117e0c1a50d3fd8671bd2b84ed579
---

# 10.3 NVIDIA Quantum-X800 InfiniBand (protocol: InfiniBand XDR — not Ethernet)

**FS.com (whitebox, published pricing).** **N9600-64OD**: 64 × 800G OSFP, Broadcom BCM78900 Tomahawk 5, 51.2 Tbps (102.4 Tbps full-duplex in FS's table), 165.2 MB packet buffer, sub-1 µs / 1 µs (64B) latency, RoCEv2/PFC/ECN/DLB/GLB, PicOS® and SONiC support, breakout to 128 × 400GbE or 256 × 200/100GbE [official] (FS.COM eu-en product 250955). FS.com list pricing (SKU 250955, excluding optics/cables/software/support and volume discounts): €40,333.86 (€33,894.00 VAT excl.) on FS.com Europe and £36,163.20 (£30,136.00 VAT excl.) on FS.com UK [official]. **N8650-32OD**: 32 × 800G OSFP, BCM78902 Tomahawk 5, 25.6 Tbps — half-width sibling in FS's AI switch portfolio [official] (FS.com N-series blog). FS's AI-switch portfolio announcement (Tomahawk 3/4/5-based, PicOS® AI Switch System, AmpCon-DC management) positions these for AI training/inference and HPC; the 800G datasheet carried an "expected availability Q3 2025" notice [official]. Caution: FS **N9550-64D** is a Tomahawk 4 (BCM56990) 400G switch at 25.6/51.2 Tbps uni/bidirectional — not an 800G-port system despite appearing in the same portfolio tables [official].

**UfiSpace, Ruijie, ZTE, Huawei, Celestica.** UfiSpace S9300 strategy scales from Trident 4-based 12.8T/400G to Tomahawk 5-based 51.2T/800G; S9700 DDC strategy scales to Jericho3/Ramon3 with 800G and claimed 921 Tbps multi-system capacity; the exact 800G S9300 model number and faceplate were not found — do not convert the 400G S9300-32D into an 800G model [official] (ufispace.com AI solution page). Ruijie RG-S6980-64QC (official datasheet): 51.2 Tbps but only 64 × 100/200/400GE QSFP-DD service ports — **not a native 800G-port system** despite the 51.2T aggregate [official] (ruijie.com datasheet, June 2026). A reseller lists RG-S6990-64OC2XS / -X as 64 × 800G OSFP with claimed 102.4 Tbps (full-duplex-implied) and AI-fabric positioning; with no official Ruijie document captured, treat this model/spec cluster as [unverified] [secondary] (network-switch.com). **ZTE**: no verified native-800G data-center switch model found — report the gap [unverified]. **Celestica**: no exact 800G model (e.g., verified DS-series 800G SKU) with official port/ASIC mapping captured — report the gap [unverified]. Huawei **CloudEngine 16800-X** chassis family (X4, X8, X16): official datasheet shows capacities of 179/387, 357/774, 714/1,548 Tbps — dual figures whose accounting (version/capacity basis) needs careful reading and may not be directly comparable; secondary reporting claims up to 288 × 800GE for the family, with RDMA/RoCE, PFC, AI ECN documented, but the exact 800G line-card model, per-slot density, and orderable configuration remain unverified [official] (Huawei 16800-X datasheet). A Huawei-central secondary piece claims a "world's first 800GE data center core switch" framing — treat the superlative as marketing, not independently verified [secondary].

### 10.3 NVIDIA Quantum-X800 InfiniBand (protocol: InfiniBand XDR — not Ethernet)

- Quantum-X800 is XDR InfiniBand at 800G per port; uses 200G-per-lane SerDes; not port-level backward-compatible with NDR optics [vendor-reported] (koicomputers.com datasheet mirror).
- Q3200-RA: two switch ASICs/groups of 36 × 800G each; aggregate descriptions vary between 57.6 Tbps and 2 × 28.8 Tbps across sources; 2U; 36 OSFP cages carrying 72 logical 800G ports [vendor-reported] (gpusmith.com datasheet mirror).
- Q3400-RA: 144 × 800G nonblocking logical ports, 72 OSFP cages, 115.2 Tbps, 4U, air-cooled [vendor-reported]. Q3401-RD: 144 × 800G, 72 OSFP cages, 115.2 Tbps, 4U, DC-power variant [vendor-reported]. Q3450-LD: 144 × 800G, 144 MPO connectors, predominantly liquid-cooled [secondary].
- Source-quality warning: several Quantum-X800 documents found are mirrors/reseller copies — tagged [vendor-reported]/[secondary], never [official]; the 57.6T vs 2×28.8T variance on Q3200-RA reflects inconsistent aggregation accounting across copies [independent].

### 10.4 Port configurations, form factors, and breakout

- Standard 51.2T lane math: 64 × 800G = 128 × 400G = 256 × 200G = 512 × 100G; half-width 25.6T: 32 × 800G = 64 × 400G; 102.4T: 64 × 1.6T = 128 × 800G = 256 × 400G = 512 × 200G [independent].
- **OSFP** dominates NVIDIA and many AI-backend products for thermal headroom (SN5600, 7060X6-64PE); **QSFP-DD800** is used where backward compatibility and established QSFP mechanics matter (Cisco N9364E-SG2-Q, Juniper QFX5241-64QD, Edgecore AIS800-64D); Cisco and Edgecore explicitly offer both cage types [independent].
- 800G **DR8** generally uses eight 100G optical lanes over MPO-16; 800G **2×DR4** exposes two independent 400G optical links, often on dual MPO-12 connectors [secondary] (FS.com 800G InfiniBand blog).
- **800G → 2 × 400G breakout** is operationally important while switch uplinks move to 800G but server/GPU NICs remain 400G; AI-cluster cabling guides document MPO-based 800G-to-2×400G fan-out practice [secondary] (Medium/aicplight888).
- Native 800G endpoints are increasing with ConnectX-8/9-class SuperNICs and newer accelerator generations, reducing but not eliminating the need for 2×400G breakout [independent].

### 10.5 AI backend vs. general data-center (frontend) positioning

- Explicit **AI backend / scale-out** products: NVIDIA Spectrum-X800 platform (SN5600 + SuperNICs), NVIDIA Quantum-X800 InfiniBand, Broadcom Jericho3-AI/Ramon3 fabrics, Arista Etherlink/7060X6/7060XE7/7800R4, Tomahawk 5/6 systems with RoCE and advanced load balancing, Cisco G200/G300 Nexus AI systems, Juniper QFX5240/QFX5250, Edgecore DCS560/AIS800, Micas M2-W6940-64OC, FS.com N9600 AI-switch line, Oracle Acceleron RoCE [independent].
- **General data-center plus AI-capable**: Dell Z9864F-ON, Marvell Teralynx 10-based designs, Juniper QFX5140 (inference/storage/leaf-spine), Huawei 16800-X [independent].
- Adoption pattern: 800G first concentrated in inter-switch spine/superspine and AI backend links; 2×400G breakout persists because GPU/server NIC endpoints often lag native switch port speed [independent].
- Do not describe all 800G switches as AI-only: the same 51.2T fixed systems serve cloud/Web 2.0, enterprise, and DCI roles (Dell, Edgecore, Micas positioning) [independent].


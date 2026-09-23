---
id: etape6-phased1-fabrics-spine-leaf/00-fabrics-spine-leaf/wave-13-switch-silicon-radix-update-tomahawk-6-systems-20252
title: "Wave 13 — Switch-silicon radix update: Tomahawk 6 systems (2025–2026)"
domain: step-6-phase-d1-data-center-fabric-architectures-spine-leaf-
role: deep-dive
task: reference
actors: ["Broadcom", "Google", "Meta", "Microsoft", "Nvidia", "TSMC", "xAI"]
dates: ["2025-06", "2026-03"]
keywords: ["3nm", "benchmark", "blackwell", "compute", "cost", "cpo", "disaggregated", "ethernet", "gpu", "gpus", "hyperscaler", "latency"]
source: docs/RAG/etape6_phaseD1_fabrics_spine_leaf.md
source_anchor: ""
source_lines: [371, 425]
section: "Step 6 Phase D1 — Data-Center Fabric Architectures: Spine/Leaf, Clos, Vendor Reference Designs"
sha256: 8b783254d0e6fd0cea9d12fdbf75eceb5c303a43b7d08eefa07efd2ebaf00bc7
---

# Wave 13 — Switch-silicon radix update: Tomahawk 6 systems (2025–2026)

## Wave 13 — Switch-silicon radix update: Tomahawk 6 systems (2025–2026)

- **Broadcom Tomahawk 6**: announced shipping June 2025 [official — https://www.globenewswire.com/news-release/2025/06/03/3092820/19933/en/Broadcom-Ships-Tomahawk-6-World-s-First-102-4-Tbps-Switch.html]; 3nm, 102.4 Tb/s, 224G SerDes; production-volume shipments reported from March 2026 [secondary]. **Davisson CPO variant** (BCM78919): 16×6.4T optical engines (TSMC COUPE), field-replaceable ELSFP lasers, 512-XPU scale-up / 100k+ XPU two-tier scale-out claims, ~70% optical-interconnect power reduction claim [vendor-reported — https://www.storagenewsletter.com/2025/10/10/broadcom-shipping-tomahawk-6-davisson-102-4-tb-s-ethernet-switch-with-co-packaged-optics/].
- **Edgecore AIS1600-64O / AIS800-128O** (Feb 2026): world's-first 102.4T open switches per vendor — 64×1.6T OSFP1600 (3RU) and 128×800G OSFP800 (4RU), Tomahawk 6-based [vendor-reported — https://markets.financialcontent.com/wral/article/bizwire-2026-2-23-edgecore-networks-sets-new-benchmark-for-ai-infrastructure-with-worlds-first-1024t-open-networking-switches].
- **DriveNets AI Fabric** on Tomahawk 6: systems announced, shipping targeted Q3 2026, aimed at hundreds of thousands of XPUs [secondary — https://www.ainvest.com/news/drivenets-gave-broadcom-investors-tomahawk-6-read-revenue-visibility-real-test-2607/].
- Fabric-math impact: a 128×800G super-spine tier supports 128 pods; 64×800G spine × 64-port leaves ≈ 4,096 server ports non-blocking per pod — single-pod 8k-GPU fabrics become radix-feasible [analysis from vendor specs].

---

## Wave 14 — Hyperscaler deployments deep-dive

### 14.1 Meta: RoCE AI Zones and Disaggregated Scheduled Fabric

From Meta engineering publications [official]:
- **AI Zone** (two-stage Clos for training): rack training switches (RTSW) as leaves with copper DAC intra-rack GPU scale-up; cluster training switches (CTSW, modular, deep buffers) as spines over single-mode fiber + 400G pluggables; non-blocking within a zone. **Aggregator training switches (ATSW)** extend the RoCE domain across zones in a building; cross-zone links are **oversubscribed by design**, mitigated by ECMP plus a topology-aware scheduler that finds a "minimum cut" for rank assignment [official — https://engineering.fb.com/2024/08/05/data-center-engineering/roce-network-distributed-ai-training-at-scale/].
- **Enhanced ECMP + QP scaling**: AI traffic's low entropy defeated standard ECMP hashing; hashing on destination queue-pair number gave up to 40% better AllReduce performance [official].
- **Disaggregated Scheduled Fabric (DSF)** (Oct 2024–2025): disaggregated line/fabric cards as separate devices; VOQ-based scheduled fabric on OCP-SAI + FBOSS for proactive congestion avoidance; open Ethernet RoCE to endpoints incl. MTIA and third-party XPUs/NICs. DSF platforms: **Arista 7700R4** — 7700R4C-38PE DSF leaf (Broadcom Jericho3-AI, 18×800G + 20×800G fabric ports, 14.4T, 16 GB buffers) and 7720R4-128PE DSF spine (Broadcom Ramon3, 128×800G, 102.4T) [official — https://engineering.fb.com/2024/10/15/data-infrastructure/open-future-networking-hardware-ai-ocp-2024-meta/ and https://engineering.fb.com/2025/10/20/data-center-engineering/disaggregated-scheduled-fabric-scaling-metas-ai-journey/].
- **RSC**: Research SuperCluster phase 2 — 16,000 GPUs on a three-level Clos fabric, used for LLaMA-era training [secondary].

### 14.2 xAI Colossus (Memphis)

NVIDIA's official account: 100,000 Hopper GPUs on **Spectrum-X Ethernet** for the RDMA network; built in **122 days** (19 days rack-to-training); **three-tier fabric**; 95% data throughput via Spectrum-X congestion control with zero application latency degradation or packet loss from flow collisions (vs ~60% on standard Ethernet per NVIDIA) [official — http://nvidianews.nvidia.com/news/spectrum-x-ethernet-networking-xai-colossus]. Expansion reporting: doubled toward 200,000 GPUs (adding H200), 2026 reports describe 230,000+ GPUs incl. 30,000+ GB200 Blackwell [secondary — https://markets.financialcontent.com/wral/article/tokenring-2026-1-1-colossus-rising-how-xais-memphis-supercomputer-redefined-the-global-compute-race]. Exact topology undisclosed — **[unverified]** on specifics.

### 14.3 Google and Microsoft (reference)

- Google's Jupiter/Aurora Clos generations: published papers describe 1+ Pb/s bisection fabrics; current-generation details proprietary — **[gap]** retained.
- Microsoft Azure remains the canonical hyperscale SONiC deployment (tens of thousands of switches); current counts not officially published — **[gap]** retained.

---

## Wave 15 — Fabric economics: whitebox TCO and market data

### 15.1 Edgecore's published TCO model (100 ToR switches, 5 years)

Edgecore's whitebox TCO analysis (32×100G ToR class, 5-year horizon, $0.10/kWh, PUE 1.4) [vendor-reported — https://www.edge-core.com/download/whitebox-switching-data-center-tco-analysis/?wpdmdl=6365&refresh=6aa0b13dd3e8b1788916029]:

| Component | Proprietary | Whitebox (SONiC) |
|---|---|---|
| Street price/switch | $22,000 | $12,000 |
| Annual NOS license | $2,500 | $0 (community) / $800 (enterprise support) |
| Annual HW support (next-day) | 15% of HW list | 10% of HW list |
| Typical power | 300 W | 260 W |
| Initial staff enablement | $150,000 | $90,000 |

Result: 5-year TCO $5.43M vs $2.49M — **~54% savings ($2.94M)**; payback ≈ 18 months [vendor-reported]. Treat as vendor-constructed model, not independent measurement.

### 15.2 Market context

- Whitebox DC Ethernet switch market: ~$3.0B in 2025 → forecast $6.6B by 2032 (11.8% CAGR); ~485k units in 2025 at ~$6k average price [secondary — https://pdf.marketpublishers.com/globalinfo/global-data-center-white-box-ethernet-switch-supply-gir.pdf].
- 650 Group (via OIF, Feb 2025): 400G/800G port shipment splits show Arista leading branded 400G/800G volumes with whitebox a major share of 400G [secondary — https://www.oiforum.com/wp-content/uploads/650-Group-February-2025-OIF-Presentation-rev2.pdf].
- Gray-market price snapshots (indicative only): Cisco N9K-C93600CD-GX (36×400G) $8,999–11,999; N9K-C93180YC-FX3 class $4,250–4,300 new via marketplace listings [secondary — treat as street/gray-market, not list].
- Practitioner sentiment (Sept 2026): Arista perceived as lower setup cost vs Cisco's higher initial cost but long-term ROI [secondary].

---


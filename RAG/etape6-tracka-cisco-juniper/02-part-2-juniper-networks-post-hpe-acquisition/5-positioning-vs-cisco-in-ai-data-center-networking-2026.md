---
id: etape6-tracka-cisco-juniper/02-part-2-juniper-networks-post-hpe-acquisition/5-positioning-vs-cisco-in-ai-data-center-networking-2026
title: "5. Positioning vs Cisco in AI data-center networking (2026)"
domain: part-2-juniper-networks-post-hpe-acquisition
role: deep-dive
task: reference
actors: ["AMD", "Broadcom", "Huawei", "Nvidia", "Oracle"]
dates: ["2026-03"]
keywords: ["accelerator", "acquisition", "amd", "blackwell", "coherent optics", "compute", "dci", "ethernet", "gpus", "helios", "nvidia", "optics"]
source: docs/RAG/etape6_trackA_cisco_juniper.md
source_anchor: ""
source_lines: [356, 384]
section: "PART 2 — JUNIPER NETWORKS (post-HPE acquisition)"
sha256: 2e2fdfff9697136fdf61059a65f15c20c26c46354f26d847af1583b8f0e72c0c
---

# 5. Positioning vs Cisco in AI data-center networking (2026)

- **Ethernet for AI / open standards:** HPE Juniper's strategy is explicitly standards-based Ethernet: Junos AI load balancing (GLB/DLB), multivendor fabric management (Apstra), and Mist AIOps layered on merchant silicon — Juniper's message is "secure, high-performance, multivendor, lower TCO" vs. closed stacks. [vendor-reported — Praveen Jain quote in Broadcom Tomahawk 6 PR]
- **Ultra Ethernet / UEC participation:** not confirmed. Juniper praised Broadcom's Tomahawk 6 as an "open, **UEC-compliant** architecture," but **no source in this pass confirms Juniper/UEC Consortium membership**. [flagged — unverified]
- **Congestion management (DCQCN / RoCEv2):** Junos EVO docs reference RoCE/rdma-opcode support and AI-ML load-balancing features, but **no 2026-dated Juniper announcement on DCQCN or RoCEv2 congestion features** was found. [gap flagged]
- **400G/800G portfolio:** QFX5240/5241 (64×800G / 32×800G, Tomahawk 5), PTX10002 (800G fixed), PTX12000/PTX10000 (800GbE, 1.6T-ready platforms); 1.6T readiness is a 2026 theme but no 1.6T ports ship yet. [official/secondary]
- **Partnerships:**
  - **Broadcom:** deepening. Dec 18, 2025 — HPE announced the **first AMD "Helios" AI rack-scale architecture** with integrated **scale-up Ethernet built with Broadcom**: purpose-built HPE Juniper Networking scale-up switch + software, using Broadcom's **Tomahawk 6** (102.4 Tbps), based on the open **Ultra Accelerator Link over Ethernet (UALoE)** standard, supporting trillion-parameter training. HPE claims **first OEM to productize a 100% liquid-cooled Tomahawk 6 switch** (from the Q2 FY26 earnings call). The Tomahawk 6 launch PR quotes Juniper SVP Praveen Jain (AI Clusters & Cloud Ready Data Center). Broadcom's industry-first **800G Thor Ultra NIC** is positioned alongside the HPE Juniper AI-optimized fabric. [official/secondary]
  - **AMD:** co-launch of the Helios rack-scale architecture (scale-up Ethernet over UALoE). [official]
  - **NVIDIA:** HPE "AI Factory" and **AI Grid with NVIDIA** (announced at NVIDIA GTC, March 2026): HPE Juniper PTX/MX routers + **800G ZR+ coherent optics**, MACsec, IP/MPLS for scale-across/DCI; HPE AI Factory with **Spectrum-X Ethernet**, BlueField-3, Blackwell GPUs, Vera Rubin NVL72 on the roadmap; AI Factory Lab in Grenoble (announced Dec 1, 2025, opened Q2 2026); a "3-2-1 integration strategy" (three silicon platforms from two companies into one fabric). HPE says the NVIDIA relationship is a two-way dependency: HPE adds integration, sovereignty, financing, GSI co-sell. [official/secondary — businesswire; Futurum]
- **Scale-out / scale-up / scale-across taxonomy:** HPE now frames AI networking as scale-out (QFX, e.g. Oracle), scale-up (Tomahawk 6-based switch for Helios), scale-across/DCI (PTX12000). [official]

## 5. Positioning vs Cisco in AI data-center networking (2026)

- **Market data (IDC, Q2 2026)** via Next Platform (Sept 20, 2026) and kad8:
  - Global Ethernet switch revenue **$18.9B** in Q2 2026 (+64.5% y/y for the data-center segment, $12.3B). [independent — IDC via Next Platform/kad8]
  - **Cisco: $5.43B, +36.2%** (28.7% global share); ~$2.2B from data-center switches; router business $1.6B (+31.6%). Largest vendor overall. [independent]
  - **NVIDIA: ~$2.5B (per kad8, +181% y/y; ~20% of DC Ethernet) — but Next Platform cites $3.86B growing 2.8×.** The two figures appear to use different scopes; **flagged as unresolved.** [independent — discrepancy noted]
  - **Arista: $2.53B, +37.6%** (13.4% global; 18.7% DC). Arista launched a modular router at >115 Tbps in early 2026 aimed at AI. [independent — IDC via Next Platform; Arista launch via Think Insights]
  - **HPE (incl. Juniper): $1.15B Ethernet, +6.1%** — "a lot of incremental business for HPE, whose Ethernet business was largely wireless under Aruba." Huawei $1.55B (+28.6%); ODMs ~$3.56B (+2.43×). [independent — Next Platform's IDC augmentation]
  - **Reading:** HPE-Juniper is #4 by Ethernet revenue (behind Cisco, NVIDIA/Arista, Huawei) but is the only vendor with full campus-to-DC-to-routing plus compute/storage; its growth lags the AI-cluster leaders (NVIDIA, Arista) in DC switching share.
- **Silicon positioning:** Juniper claims Express 5 (28.8 Tb/s) beats Cisco P100 (19.2 Tb/s) by ~33% throughput — a vendor claim Cisco declined to confirm or deny. [vendor-reported]
- **Analyst commentary:**
  - Futurum (Discover 2026 recap): "the networking story is the strongest, most defensible part of HPE's hand"; honest tension is the two coexisting management planes (HPE Mist vs Aruba Central); key tests = self-driving delivery in brownfield, security story above the network, NVIDIA dependency. [independent]
  - ACG's Ray Mota: "HPE delivers the architectural and operational foundation service providers need to fully participate in the AI value chain." [independent quote]
  - EMA's Shamus McGillicuddy: Juniper was showing more value than anyone else; HPE's acquisition centers on that. [independent]
  - Think Insights industry analysis: HPE-Juniper is a new full-stack competitor "explicitly positioned to challenge Cisco and Nvidia in AI-native networking"; Arista retains hyperscale/DC advantage; white-box/SONiC competes on price. [independent]
  - ETR AI survey cited by Futurum: ~21% of orgs building AI apps use HPE for hosting/deployment (and plan to stay), 14% considering — "most of the field is still open." [independent]
  - HPE's own claim (Feb 2026, MWC): "I don't think you could find one company who has a leadership in compute, storage, security and networking across the board" — the full-stack pitch vs Cisco/Arista. [vendor-reported]
- **Customer wins vs Cisco (2026):** Oracle's gigawatt-scale HPE Juniper win is the marquee 2026 reference; Wi-Fi 7 AP sales 7× and campus record orders suggest campus displacement momentum; no 2026-dated named hyperscale DC wins against Cisco surfaced. [official on Oracle; gaps flagged]


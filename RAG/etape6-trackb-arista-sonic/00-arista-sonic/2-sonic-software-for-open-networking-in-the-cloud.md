---
id: etape6-trackb-arista-sonic/00-arista-sonic/2-sonic-software-for-open-networking-in-the-cloud
title: "2. SONiC (SOFTWARE FOR OPEN NETWORKING IN THE CLOUD)"
domain: step-6-track-b-arista-sonic-cumulus-data-center-fabric-open-
role: deep-dive
task: reference
actors: ["AWS", "Broadcom", "Microsoft", "Nvidia"]
dates: ["2025-05", "2025-12-09", "2026-03-29", "2026-04-15", "2026-05", "2026-05-27", "2026-06", "2026-10"]
keywords: ["agent", "aws", "ethernet", "inference", "nvidia", "training"]
source: docs/RAG/etape6_trackB_arista_sonic.md
source_anchor: ""
source_lines: [96, 162]
section: "Step 6 — Track B: Arista + SONiC + Cumulus (Data-Center Fabric & Open Networking)"
sha256: 54f58c53dbf27af72443ca0ed20e8f6778612cdf09e2346341d007aeccd586e9
---

# 2. SONiC (SOFTWARE FOR OPEN NETWORKING IN THE CLOUD)

## 2. SONiC (SOFTWARE FOR OPEN NETWORKING IN THE CLOUD)

### 2.1 Community status and Microsoft's role

- SONiC is a **Linux Foundation project** (moved there in 2022) developing an open-source, Debian-based network operating system for switches, hardware-agnostic and modular **[independent — Network World, Jan 2026]**.
- Started life at **Microsoft**, adapted from its Debian-based Azure Cloud Switch; Microsoft remains a core contributor/user **[independent — The Register, May 27, 2026]**.
- SONiC 4.5 released May 2025, community-supported until at least October 2026 **[independent — Network World]**.
- Commercial vendors (Dell, Cisco, Juniper, Arista, NVIDIA, etc.) support SONiC on their hardware; Arista's 7060XE7 explicitly supports open network OS options (SONIC/OpenSwitch) **[independent/secondary]**.

### 2.2 2026 releases and distributions

**Enterprise SONiC 4.6 (Stordis) — GA June 2026**
- Operations-focused release: route-table **overflow detection with automatic recovery and retry**; improved route consistency checker (per-prefix inspection, transient-route ignore, CLI-triggered recovery); better telemetry; tighter access control; broader hardware list **[secondary — Stordis release breakdown]**.
- Base system moved to **Debian 12 Bookworm**, Linux kernel 6.1; Broadcom SAI Adapter 15.3.0; Broadcom SDK 6.5.35; FRR 8.2.2; Config DB version_4_6_1 **[secondary]**.
- Covers leaf-spine fabrics, Ethernet networks for AI clusters, and campus edge **[secondary]**.
- ⚠️ Release-notes documentation footnote labels Debian 12 as "Bullseye" in one component table; the feature section confirms Bookworm — minor doc inconsistency flagged by Stordis **[secondary]**.

**Netberg SONiC 202511.n0 — released April 15, 2026**
- Successor to the 202411 release; positioned for data centers, AI-driven infrastructure, and campus networks **[secondary — PRLog/Netberg]**.
- Key features: **AI workload optimization** (enhanced **SRv6** support for high-performance AI training/inference backend networks with efficient load balancing); advanced observability (high-frequency telemetry, advanced diagnostics); **Debian 13 (Trixie)** base with updated kernel, security patches, modern toolchain **[secondary]**.

**Aviz Networks turnkey enterprise SONiC — announced December 9, 2025**
- "Aviz Certified Community SONiC": full lifecycle management of community SONiC with OEM-NOS rigor — one continuously qualified code base; **50%+ savings** claimed from software-first operations; 24×7 global support with direct escalation and RMA handling; layered security with SLA-backed CVE patching **[secondary — BusinessWire, Dec 9, 2025]**.
- Certified image qualified across Accton/Edgecore, Celestica, **Cisco, NVIDIA**, Wistron (more coming); FTAS test reports across 100+ enterprise deployment scenarios at the OCP-certified **Aviz ONE Center**; validated runbooks for data center, edge, and **AI fabrics (IP CLOS, EVPN/VXLAN, MLAG, OpenStack, Spectrum-X)** **[secondary]**.
- CEO Vishal Shukla: "Since 2023, SONiC enterprise deployment has shifted from early adopters to large, mainstream enterprises." **[secondary]**.

### 2.3 Major adopter moves 2026

- **Cisco extends SONiC to Nexus 9000** (reported May 27, 2026): Cisco — which supported SONiC on routers for years — announced it will soon extend support to its **N9000 series data-center switches**. Quote: "The N9000 Series is expanding to include a foundation for SONiC, built on Cisco Cloud Scale and Silicon One — alongside platforms powered by NVIDIA Spectrum-X Ethernet switch silicon for AI-class fabrics." Analyst framing: hyperscalers use SONiC for the AI-service networks; as enterprises build on-prem AI (many won't run AI in the cloud), SONiC on everyday data-center hardware "is an idea whose time has come" **[independent — The Register]**.
- **Arista 7060XE7** supports SONiC/OpenSwitch as open-OS alternatives to EOS **[secondary — Zacks, Network World]**.
- **Dell** has long argued SONiC is ideal for enterprise switching and supported it while promoting NOS choice **[independent — The Register]**.
- Dell, Cisco, Juniper, Arista, NVIDIA all support SONiC on their hardware **[independent]**.

### 2.4 Key networking features (SONiC general, per Network World state-of-play)

- MCLAG (L2/L3), weighted ECMP, BGP/OSPF/BFD/IS-IS via FRRouting **[independent]**.

---

## 3. NVIDIA CUMULUS LINUX

### 3.1 NVIDIA stewardship and release status

- Cumulus Networks was acquired by **NVIDIA in 2020**; Cumulus Linux is NVIDIA's switch-optimized Linux-based NOS **[independent — Network World, Jan 2026]**.
- **Latest releases (as of early/mid-2026)**:
  - **Cumulus Linux 5.15** — most recent mainline version (not LTS) **[independent — Network World]**.
  - **Cumulus Linux 5.14** — Debian Bookworm-based **[official — NVIDIA cumulusnetworks/docs GitHub]**.
  - **Cumulus Linux 5.11** — LTS release (debuted 2024), supported **until 2027** **[independent — Network World]**.
  - Cumulus Linux 5.5/5.3 documentation marked "old" (Debian Buster-based) **[official — NVIDIA docs GitHub]**.
- **Support policy [official — NVIDIA docs]**: NVIDIA supports mainline and LTS branches as separate codebases; LTS prioritized for stability (security + critical bug fixes only, no new functionality); mainline gets 12 months of support after the LTS release to migrate. **End-of-life releases are not supported.**
- **NVIDIA's current recommendation [official]**: run the latest **Cumulus Linux 5.y.z** on Spectrum switches; the latest **4.3.z** on Broadcom switches.
- Cumulus Linux 5.5+ includes the **NVIDIA NetQ** agent and CLI for data-center network monitoring and operational health **[official — NVIDIA docs]**.

### 3.2 Cumulus VX / NVIDIA AIR (2025–2026 development)

- NVIDIA **no longer releases Cumulus VX as a standalone image** (per the Cumulus Linux 5.13 "What's New" document): to simulate a Cumulus Linux switch, use **NVIDIA AIR** — the cloud-hosted data-center simulation platform (digital twin of IT infrastructure) **[secondary — ipSpace.net blog, Jun 2025 post; noted updated Mar 2026]**.
- **Update (March 29, 2026)**: a reader report claims NVIDIA has resumed publishing Cumulus VX images, but the newer images are **only accessible with a support contract** — single-reader claim, **[unverified]**.
- Commentary (ipSpace): the forced-NVIDIA-AIR move mirrors past failed vendor attempts (Juniper Junosphere, Cisco before CML); the cloud-only requirement drew community criticism **[secondary — ipSpace.net]**.

### 3.3 Key features and ecosystem

- **Unnumbered interfaces** (simplified BGP/OSPF template for leaf-spine); **Redistribute Neighbor (RDNBR)** for VM/host mobility with L3 discovery; **Prescriptive Topology Manager (PTM)** for connection verification; **NVUE** (NVIDIA User Experience) full-CLI object model enabling advanced programmability **[independent — Network World]**.
- **FireMon** added NVIDIA Cumulus support to its Policy Manager (2025.2.6 feature release, ~May 2026): unified security policy management for open-networking fabrics built on Cumulus, alongside firewalls, AWS/Azure/GCP, Zscaler, Cisco ACI/NSX **[secondary — DataCenterNews Asia]**.
- **NetQ** bundled for telemetry and operational health monitoring **[official]**.

---


---
id: collect-261001-general-networking/general-networking/etape6-trackb-arista-sonic-3
title: "Step 6 — Track B: Arista + SONiC + Cumulus (Data-Center Fabric & Open Networking)"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "AWS", "Broadcom", "Intel", "Meta", "Microsoft", "Nvidia"]
dates: ["2023-07-19", "2025-06-11", "2025-12-09", "2026-03-29", "2026-05", "2026-05-27", "2026-07-16"]
keywords: ["agent", "alignment", "amd", "aws", "ethernet", "intel", "latency", "nvidia", "optics"]
source: docs/RAG/collect-261001-general-networking/etape6_trackB_arista_sonic.md
source_anchor: ""
source_lines: [117, 187]
sha256: 52360dc30a69954f7cf6edbf32e93dc53365dbe29bf0213be94300eaf6cbd7f1
---

# Step 6 — Track B: Arista + SONiC + Cumulus (Data-Center Fabric & Open Networking)

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

## 4. OPEN NETWORKING LANDSCAPE 2026

### 4.1 Ultra Ethernet Consortium (UEC) / Super Ethernet

**Specification status**
- UEC is a **Linux Foundation Joint Development Foundation** project founded **July 19, 2023** by AMD, Arista, Broadcom, Cisco, Eviden/Atos, HPE, Intel, Meta, Microsoft **[secondary — LLM-Systems-Wiki UEC page, updated Aug 26, 2026]**.
- **Specification v1.0**: 562-page document released **June 11, 2025**, defining a full Ethernet-based communication stack (NICs, switches, optics, cables) with the new **Ultra Ethernet Transport (UET)** **[secondary]**.
- Current release: **1.0.3 (July 16, 2026)**; 1.0.2 also released in 2026 **[secondary]**.
- Consortium scale: **100+ member companies, 1,500+ participants** **[secondary — Medium/Teradata Labs, Jun 2026]**.
- UET design goals: per-packet multipathing, lossy-or-lossless operation, fast loss recovery, libfabric compatibility; removes in-order-delivery requirement and lossless-fabric dependency of RoCEv2 **[secondary]**.

**2026 technical priorities** (per Chad Hintz, UEC marketing co-chair / AMD, via Network World):
1. **Programmable Congestion Management (PCM)** — implement congestion-control algorithms in a standard language, portable across UE NICs.
2. **Congestion Signaling (CSIG)** — packets carry high-fidelity congestion information for faster, more accurate transport reactions.
3. **Small-message performance** — UET 1.0's 104-byte headers are fine for 4KB packets (2.5% overhead) but costly for 256-byte transactions; UEC pursuing a reduced-size forwarding header to cut overhead roughly in half for optimized deployments (matters for HPC and local scale-up networks).
4. **In-Network Collectives (INC)** — move AI/HPC reduction operations from hosts into the network **[independent — Network World, ~Jan 2026]**.

**Vendor alignment**
- **Broadcom**: "first UE-compatible Tomahawk" switch with ~250ns latency (claimed); UEC-aligned features (Cognitive Routing 2.0, Global Load Balancing) in Tomahawk 6 **[secondary — Medium; financialcontent]**.
- **Arista**: founding/steering UEC member; Etherlink roadmap tracks UET (per-packet spraying, NSCC/RCCC, trimming, in-network collectives) **[secondary — LLM-Systems-Wiki]**.
- **NVIDIA**: UEC member (~2024) but pushes its proprietary **Spectrum-X** stack; Spectrum-X is the bundled alternative to UEC Ethernet **[secondary]**.
- Ethernet Alliance 2026 focus: higher bandwidth, AI demands, interoperability events, another 200G/lane plugfest **[independent — Network World]**.

### 4.2 Tomahawk 6 — shipments and deployments


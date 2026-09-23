---
id: etape6-phasea-vendors-dc/00-front-matter/2-nvidia-networking-mellanox-heritage
title: "§2 NVIDIA Networking (Mellanox heritage)"
domain: front-matter
role: reference
task: actor-profile
actors: ["Cohere", "CoreWeave", "Microsoft", "Nvidia", "Oracle", "TSMC"]
dates: ["2026-02", "2026-05", "2026-07-15"]
keywords: ["nvidia", "asic", "blackwell", "cpo", "ethernet", "fp8", "gpu", "gpus", "inference", "latency", "optics", "pricing"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [91, 144]
section: "Step 6 — Phase A: Enterprise Data-Center Switching Vendors"
sha256: 49096bc725aa3b35f79648ee99fa0699fc987055f35501fe43897ee63510684d
---

# §2 NVIDIA Networking (Mellanox heritage)

## §2 NVIDIA Networking (Mellanox heritage)

### 2.1 Strategic position: #1 in data-center Ethernet switching (Q1 2026)

- **IDC Q1 2026** (published Jun 18, 2026): NVIDIA became the **#1 vendor in data-center Ethernet switching by revenue** — **$2.1B, 21.5% share, +192.7% YoY**; Arista 20.7%, Cisco 17.8%; DC switching segment $10B (+61% YoY); total Ethernet switch market $15.4B (+39.8%) [independent — IDC via SDxCentral, Data Center Knowledge, Business Insider]. IDC's Paul Nicholson called it "one of the most significant vendor landscape shifts IDC has tracked in enterprise networking" [independent].
- Growth driver: **Spectrum-X platform** (tightly integrated GPU-plus-networking bundle with BlueField DPUs and LinkX interconnects); 800 Gb/s switches alone were 35.8% of DC segment revenue [independent].
- **Networking segment revenue reached ~$14.8B in Q1 FY2027 (+199% YoY)**, ~18% of NVIDIA's $81.6B quarterly revenue; prior quarters: $4.9B (Q2 FY26) → $11B (Q4 FY26) [vendor-reported; secondary — ainvest.com, SDxCentral].
- Jensen Huang's GTC claim: NVIDIA is "the largest networking company in the world" [vendor-reported].

### 2.2 Spectrum-X Ethernet (scale-out AI fabric)

- **Spectrum-X800 Ethernet** (announced GTC 2024, shipping 2025-2026): first end-to-end 800 Gb/s Ethernet platform for AI; initial adopters **Microsoft Azure, Oracle Cloud Infrastructure, CoreWeave** [official].
- **Spectrum-X Photonics** (CPO, GTC 2025): two configs — **SN6810**: 128× 800G (102.4 Tbps); **SN6800**: 512× 800G (409.6 Tbps); multi-chip design (central packet engine + 8 SerDes chiplets), 224 Gb/s lanes; Spectrum-X Photonics Ethernet switches "coming in 2026 from leading infrastructure and system vendors" [official]. H2-2026 ship status unconfirmed [gap].
- **Spectrum-6 ASIC**: 102.4 Tbps class; powers Dell PowerSwitch SN 6600 LD and Vera Rubin NVL72 deployments (May 2026) [vendor-reported].
- **Spectrum-XGS** (scale-across): announced for multi-site AI-factory scale-out; details in step 5 track D; 2026 commercial status not separately confirmed [gap].
- Positioning: Spectrum-X = open Ethernet with RDMA/AI congestion-control for "AI-optimized networking in every data center"; InfiniBand = highest-performance AI-dedicated fabric [official].

### 2.3 Quantum-X InfiniBand (XDR generation, 2026 current)

- **Quantum-X800 (XDR)** is the 2026 shipping generation (announced GTC 2024):
  - **Q3200-RA** (920-9B34F): 36 XDR ports over 18 OSFP cages, 2U, managed, production-release [official — NVIDIA networking-docs, 2026 Rel2A].
  - **Q3400-RA** (920-9B36F): 144 XDR ports over 72 OSFP cages, 4U, 8 PSUs, managed [official].
  - **Q3400-RD / Q3401-RD**: DGX switch variant, 48V DC bus bar [official].
  - XDR InfiniBand: 800 Gb/s per port (ConnectX-8 default IB mode), sub-microsecond latency, hardware-accelerated encryption, SHARP v3/v4 in-network computing; NVOS 25.03.1010 (2026 Rel2A), SHARP 3.16.3, UFM Enterprise 6.26.1, UFM Telemetry 1.24.2 [official].
- **Quantum-2 (NDR, 400G)** remains current/installed base: QM9700 (64× 400G, 51.2 Tbps, 1U), QM9790 unmanaged, CS9500 modular (up to 2,048× 400G = 1.6 Pbps); MLNX-OS **3.12.6300 released 2026-07-15** for QM9700/QM9701 [official].
- **Quantum-X Photonics (CPO InfiniBand)**: 144× 800G (115.2 Tbps), liquid-cooled, SHARPv4 with FP8 support, 14.4 TFLOPS in-network processing; announced GTC 2025, availability "later this year" (2025) per original plan [official]. 2026 ship confirmation not located [gap].
- Deployment evidence 2026: **LANL Mission/Vision supercomputers** (HPE-built) selected **Vera Rubin + Quantum-X800**; operational late 2027 [secondary — financialcontent.com]; **HPE Cray GX5000** offers Quantum-X800 (GTC 2026) [vendor-reported]; **DOE Argonne Solstice/Equinox** (100K + 10K Blackwell GPUs) use NVIDIA networking [secondary — stocktitan.net].
- InfiniBand retains HPC leadership; Ethernet (Spectrum-X) is "roughly on par" with InfiniBand demand inside NVIDIA's networking revenue [secondary — SDxCentral].

### 2.4 ConnectX SuperNICs (2026)

- **ConnectX-9 SuperNIC**: GA **February 2026** (firmware v82.48.1000); up to 800 Gb/s per port over InfiniBand and Ethernet; drives up to **1.6 Tb/s of throughput to Rubin GPUs**; programmable IO, intelligent congestion control; fully integrated with Spectrum-X and Quantum-X800 [official — NVIDIA firmware release notes].
- **ConnectX-8 SuperNIC**: mass production; 800 Gb/s XDR IB (default) / 2× 400 GbE, PCIe 6.0 x16, crypto + secure boot; C8180 HHHL and C8280Z mezzanine for GB300 [official — networking-docs].
- ConnectX-7/8 installed base persists; ConnectX-10 on roadmap [secondary — GitHub hardware-lineup memo].

### 2.5 BlueField DPUs and LinkX interconnects

- **BlueField-4** unveiled at GTC 2026 (with Rubin/STX/CMX generations) [official — GTC 2026 coverage]. BlueField-3 remains the shipping SuperNIC/DPU for Spectrum-X bundles [secondary].
- **LinkX**: end-to-end cable/transceiver portfolio — 800G twin-port OSFP transceivers (MMA4Z00-NS), 400G QSFP112 (MMA1Z00-NS400), passive fiber MPO cables, DAC/ACC copper; part of every Spectrum-X AI-factory bill of materials [official — NVIDIA Quantum-2 platform datasheet].
- Optics ecosystem for CPO: TSMC, Coherent, Lumentum, Fabrinet, Foxconn, Corning, SENKO [official].

### 2.6 Customer wins and competitive notes 2026

- Spectrum-X/X800 initial adopters: Microsoft Azure, OCI, CoreWeave [official]. Spectrum-X labeled by IDC as "the preferred network interconnect for large-scale AI training" [independent].
- The core competitive dynamic: buyers increasingly procure **networking as part of a bundled GPU-plus-switch stack** rather than standalone — favoring NVIDIA against Arista/Cisco on AI deals [independent — IDC via letsdatascience.com; remio.ai analysis].
- Arista's counter-positioning: open EOS, multi-vendor accelerators/NICs, UEC bet vs NVIDIA vertical integration [secondary — remio.ai].

### 2.7 Pricing

- NVIDIA does not publish switch/NIC list pricing; all pricing is quote/OEM-based. No verifiable 2026 street prices for SN5600/SN6810/Q3400/ConnectX-9 were located [gap].
- IDC ASP inference: DC Ethernet averaged ~$2.1B/quarter on NVIDIA's share driven by 800G mix (35.8% of DC segment revenue from 800G switches) [independent].

---


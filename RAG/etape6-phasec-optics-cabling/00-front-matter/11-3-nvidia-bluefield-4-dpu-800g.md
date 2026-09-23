---
id: etape6-phasec-optics-cabling/00-front-matter/11-3-nvidia-bluefield-4-dpu-800g
title: "11.3 NVIDIA BlueField-4 DPU (800G)"
domain: front-matter
role: reference
task: actor-profile
actors: ["AMD", "Broadcom", "CoreWeave", "Crusoe", "IREN", "Lambda", "Mistral", "Nebius", "Nvidia", "Oracle"]
dates: ["2024-05", "2025-10-14", "2025-10-28", "2026-03", "2026-09", "2026-09-22"]
keywords: ["nvidia", "amd", "benchmark", "compute", "ethernet", "gpu", "helios", "latency", "lpo", "memory", "mistral", "nvlink"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [1673, 1726]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: fbafe943bea2db8e27a0a5f6884ccf52b16467b0bfeb69a195213e1ea8e9baaf
---

# 11.3 NVIDIA BlueField-4 DPU (800G)

### 11.3 NVIDIA BlueField-4 DPU (800G)

- NVIDIA announced BlueField-4 at GTC DC in Washington, D.C. on October 28, 2025 [vendor-reported].
- BlueField-4 is rated at 800 Gb/s throughput, double BlueField-3 [vendor-reported].
- BlueField-4 is claimed to deliver six times more compute than BlueField-3 [vendor-reported].
- BlueField-4 combines an NVIDIA Grace CPU (64-core Arm Neoverse V2 class per secondary reporting) with the ConnectX-9 SuperNIC [vendor-reported].
- BlueField-4 targets AI storage, networking, security, multi-tenancy, cloud elasticity, and DOCA microservices [vendor-reported].
- Networking adoption: Arrcus harnesses BlueField-4 for gigascale AI factories (October 28, 2025) [vendor-reported] (businesswire.com).
- In March 2026 at GTC 2026, NVIDIA launched the BlueField-4 STX storage reference architecture with storage-optimized BF-4 combining Vera CPU and ConnectX-9 SuperNIC plus Spectrum-X Ethernet [vendor-reported].
- NVIDIA expects partners to start shipping BlueField-4 STX systems in the second half of 2026; Oracle, Mistral AI, and CoreWeave were named early adopters [vendor-reported].
- Storage partners building BlueField-4 platforms include AIC, Cloudian, DDN, Dell Technologies, HPE, Hitachi Vantara, IBM, Nutanix, Pure Storage, Supermicro, VAST Data, and WEKA [vendor-reported].
- STX-based platform availability was guided for H2 2026; no standalone BlueField-4 card retail price or firm GA date for general availability was found as of September 22, 2026 [independent].
- STX early adopters for context-memory storage include CoreWeave, Crusoe, IREN, Lambda, Mistral AI, Nebius, OCI, and Vultr [vendor-reported].

### 11.4 Broadcom Thor Ultra (800G, sampling)

- Broadcom launched the **Thor Ultra** 800GbE AI NIC on October 14, 2025 [independent].
- As of September 2026, the strongest availability statement found is **sampling to select customers**, not broad production shipment [secondary].
- Thor Ultra is specified as aggregate 800GbE with PCIe Gen6 x16 host interface [vendor-reported].
- Thor Ultra supports UEC-oriented packet-level multipathing, out-of-order placement, and selective retransmission [vendor-reported].
- Thor Ultra advertises programmable congestion control, packet trimming, and congestion signaling [vendor-reported].
- Thor Ultra includes line-rate encryption/decryption, secure boot, signed firmware, and attestation [vendor-reported].
- Reported configurations: 8×100G lanes (100G PAM4) or 4×200G lanes (200G PAM4) SerDes options [secondary].
- Reported form factors: CEM PCIe and OCP 3.0; single OSFP112 or dual QSFP112 port options in secondary Hot Chips 2026 reporting [secondary].
- Reported OSFP port modes: 1×800G, 2×400G, 4×200G, or 8×100G [secondary].
- Reported media support: DAC, AEC, optical transceivers, and linear-drive optics [vendor-reported].
- Hot Chips 2026 secondary architecture details: 5 nm process, ~2.4 billion transistors, 27×27 mm package, chip power 40–42 W, board power 50–55 W before optics [secondary].
- Broadcom's "15% faster task time" and low bit-error claims are vendor projections, not independent benchmark results [vendor-reported].
- Thor Ultra targets rack-to-rack scale-out, complementing scale-up fabrics such as NVLink (72–256 XPUs per rack-scale domain) [independent].
- No credible public unit price for Thor Ultra was found as of September 22, 2026 [unverified].

### 11.5 Broadcom Thor 2 / BCM57608 (400G — correction, not 800G)

- Broadcom's **Thor 2** (BCM57608) is a **400G** PCIe Gen5 AI Ethernet adapter family, unveiled May 2024 [independent].
- Characteristics: 400GbE, PCIe Gen5, 5 nm, eight SerDes lanes supporting 100/50G PAM4 and 10/25G NRZ [vendor-reported].
- Thor 2 includes a third-generation RoCE pipeline, low-latency congestion control, and telemetry [vendor-reported].
- Thor 2 claims passive copper reach up to 5 m and explicit LPO support [vendor-reported].
- Thor 2 was broadly available from Broadcom and multiple server vendors [independent].

### 11.6 Broadcom Stingray PS1100R (100G storage adapter — correction, not 800G)

- **PS1100R is a 100GbE storage adapter, not a 400G or 800G DPU** [official].
- Reported characteristics: BCM58804H Stingray SoC, eight Arm Cortex-A72 cores at 3.0 GHz, DDR4-2400 channels, 100 Gb/s crypto, RAID 5/6 and erasure-coding acceleration, integrated NetXtreme E-Series 100GbE, PCIe Gen3 [official].

### 11.7 AMD Pensando (Pollara 400G, Vulcano 800G)

- AMD Pensando **Pollara** is a **400GbE** AI NIC: PCIe Gen5 x16, 1×400G / 2×200G / 4×100G, RDMA with RCCL optimization [vendor-reported].
- AMD Pensando **Vulcano** is the relevant **800G** Pensando AI NIC, reported as PCIe Gen6 with UEC support, part of the AMD Helios rack design [vendor-reported].
- Oracle's planned 50,000-GPU AMD Instinct MI450 Helios supercluster (by 2026) is reported to connect each GPU with up to **three 800 Gbps AMD Pensando "Vulcano" AI-NICs**, with networking aligned to Ultra Ethernet Consortium standards [secondary].
- Timeline evidence conflicts: TechRadar described a 2026 Vulcano launch target, while a later secondary roadmap interpretation says Pollara persists through 2026 and an 800G Pensando transition comes in 2027 [secondary].
- AMD benchmark claims versus ConnectX-7 or Thor 2 are AMD marketing unless independently reproduced [independent].
- Salina is a separate Pensando DPU in secondary material and should not be conflated with Vulcano [secondary].
- No credible unit price for Vulcano was found as of September 22, 2026 [unverified].


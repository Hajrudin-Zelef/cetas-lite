---
id: etape6-phasef1-nic-dpu-smartnic/00-nic-dpu-smartnic/wave-12-master-comparison-tables
title: "Wave 12 — Master comparison tables"
domain: phase-f1-nic-dpu-and-smartnic-nvidia-amd-pensando-intel-ipu-
role: deep-dive
task: pricing
actors: ["AMD", "AWS", "Google", "Intel", "Nvidia", "Oracle"]
dates: []
keywords: ["amd", "apache", "asic", "aws", "ethernet", "helios", "intel", "kv cache", "lean", "license", "nvidia", "open source"]
source: docs/RAG/etape6_phaseF1_nic_dpu_smartnic.md
source_anchor: ""
source_lines: [427, 472]
section: "Phase F1 — NIC, DPU and SmartNIC (NVIDIA, AMD Pensando, Intel IPU, Marvell, offloads, pricing)"
sha256: 82d498f92ce178e3351edb35461cb9b94bb1d64a9b7cdc2b4e1271d506580878
---

# Wave 12 — Master comparison tables

## Wave 12 — Master comparison tables

### 12.1 Device-class comparison (2026 snapshot)

| Device class | Representative products | Implementation | Host isolation / OS | Programmability | Notes |
|---|---|---|---|---|---|
| DPU (full) | NVIDIA BlueField-3/4 | Arm SoC + ASIC datapath + accelerators | Own Linux (DPU OS), integrated BMC | DOCA SDK, P4, DPDK | Crypto enabled/disabled variants; 16–64 Arm cores |
| DPU (full) | AMD Pensando Salina 400 | P4-programmable MPU SoC + 16× N1 cores | PSM/AMD stack | P4, Pensando SDK | 232 P4 engines; 400G×2 |
| UEC NIC (AI) | AMD Pollara 400 | ASIC RDMA NIC | Host driver | Programmable RDMA (P4 pipeline elements) | First UEC 1.0 NIC; Oracle deployed |
| UEC NIC 800G | AMD Vulcano 800 | ASIC | Host driver | TBA | Helios-designed-in; 2027 transition |
| IPU (ASIC) | Intel Mount Evans (E2000) | ASIC + 16× N1 cores | Own OS (Google infra) | P4, IPDK | Google co-design; C3 VM |
| IPU (FPGA) | Intel Oak Springs Canyon | Xeon D + Agilex FPGA | Own OS | P4/RTL | 2×100G; OVS/NVMe-oF/RoCEv2 |
| DPU (ASIC) | Marvell OCTEON 10 | Arm N2 + ASIC datapath | Linux SDK | DPDK, VPP, SDK | 5 nm; 1T switch on-chip; DPU400 |
| SmartNIC (FPGA) | Napatech NT200A02, Silicom, Achronix | FPGA | Host DPDK | RTL/P4 | Line-rate capture, PTP |
| SuperNIC | NVIDIA ConnectX-8, BlueField-3 E-series | ASIC (+lean Arm) | Host driver | ASAP², DOCA (limited) | 400/800G AI scale-out |
| In-house | AWS Nitro, Azure SmartNIC, Google IPU | Custom | Proprietary | Internal | Not merchant |

### 12.2 Software-stack comparison

| Stack | Vendor | Interfaces | Key workloads | License posture (observed) |
|---|---|---|---|---|
| DOCA | NVIDIA | DOCA Flow, DPDK, P4, gRPC/REST, App Shield | OVS, IPSec/TLS, NVMe-oF, AI storage (KV cache) | NVIDIA-licensed, free-to-use with hardware |
| IPDK | Intel | P4, DPDK, SPDK, OVS-DPDK | Infra offload on IPU | Open-source components (Apache-2.0) |
| Pensando SDK/PSM | AMD | P4, policy (DSC) | Distributed services, firewall, load-balancer offload | Proprietary |
| OCTEON SDK | Marvell | DPDK, VPP, Linux | vRouter, 5G UPF, firewall | Proprietary SDK |
| SONiC/DASH | Open-source | SAI | Smart-switch / DPU networking | Open source |

### 12.3 RDMA transport comparison

| Transport | Fabric requirement | Ordering | Multipath | Congestion control | Typical NIC |
|---|---|---|---|---|---|
| InfiniBand | IB fabric (lossless) | Ordered | N/A (fat-tree) | IB CC (ECN-based) | ConnectX-7/8 IB mode |
| RoCEv2 | Lossless Ethernet (PFC/ECN/DCQCN) | Ordered | ECMP hash only | DCQCN | ConnectX-6/7/8, E2000, Pollara |
| UET (UEC 1.0) | Standard Ethernet, no PFC required | Unordered (first-class) | Native spraying | NSCC / RCCC (ACK-based) | Pollara 400, Vulcano |
| iWARP | TCP/IP Ethernet | TCP-ordered | TCP | TCP CC | Legacy Intel/Chelsio |

### 12.4 Buyer checklist (deployment)

- Confirm PSID/crypto variant (crypto-enabled vs disabled) matches regulatory domain.
- Verify PCIe Gen5 x16 slot availability and **auxiliary power cable kit** (BF-3 B3220/B3240 exceed slot power).
- Dual-slot clearance for B3240; HHHL vs FHHL mechanical fit.
- QSFP112 optics/cabling compatibility (400G-DR4/SR4, breakout options).
- DOCA version pinning vs host OFED/MOFED; firmware-PSID compatibility matrix.
- Secure boot / measured boot for DPU Arm complex (BlueField has secure boot + BMC; details per model — verify with vendor).
- Grey-market risk: third-party brands (e.g. ecer "Comelink") may lack NVIDIA warranty; prefer OEM/channel.


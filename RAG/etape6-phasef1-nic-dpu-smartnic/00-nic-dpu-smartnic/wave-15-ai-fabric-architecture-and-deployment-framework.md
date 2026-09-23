---
id: etape6-phasef1-nic-dpu-smartnic/00-nic-dpu-smartnic/wave-15-ai-fabric-architecture-and-deployment-framework
title: "Wave 15 — AI fabric architecture and deployment framework"
domain: phase-f1-nic-dpu-and-smartnic-nvidia-amd-pensando-intel-ipu-
role: deep-dive
task: pricing
actors: ["AMD", "Broadcom", "Google", "Intel", "Nvidia"]
dates: ["2024-05", "2025-12", "2026-06", "2026-09-22"]
keywords: ["amd", "asic", "benchmarks", "chiplet", "cost", "ethernet", "gpu", "gpus", "helios", "inference", "intel", "nvidia"]
source: docs/RAG/etape6_phaseF1_nic_dpu_smartnic.md
source_anchor: ""
source_lines: [537, 603]
section: "Phase F1 — NIC, DPU and SmartNIC (NVIDIA, AMD Pensando, Intel IPU, Marvell, offloads, pricing)"
sha256: c97512d28799133fff3c037cd77f388c797bb6ea252651d1ebbeb7bb62d87a04
---

# Wave 15 — AI fabric architecture and deployment framework

## Wave 15 — AI fabric architecture and deployment framework

### 15.1 Scale-up vs scale-out networking

- **Scale-up (GPU-to-GPU within node/rack)**: NVIDIA NVLink / NVSwitch; proprietary; BlueField-4 pairs with **multiple ConnectX-9 chips delivering 1.6 Tb/s per GPU** starting with Rubin [vendor-reported].
- **Scale-out (across racks, Ethernet)**: two philosophies — (a) NVIDIA Spectrum-X (lossless Ethernet tuned for AI, telemetry-driven congestion control) and (b) UEC/open Ethernet (Pollara/Vulcano + Tomahawk 6 + Arista Etherlink) [independent].
- BlueField-4 itself: **800 Gb/s total, typically 2×400 Gb/s bidirectional links**, Ethernet front-end for north–south traffic, security services on-DPU [vendor-reported] [source: https://www.sdxcentral.com/analysis/nvidias-bluefield-4-a-first-look-at-the-dpu-built-to-run-ai-factories/ ].

### 15.2 AI-native storage on DPU

- NVIDIA positions BlueField-4 as an **AI-native storage controller**: NVMe-oF/TCP, NVMe/TCP support carried over and expanded; KV-cache (NIXL/Dynamo context) served from DPU-attached storage for inference; partners (VAST Data et al.) run full storage stacks on BlueField; targeted partner availability H2 2026 [vendor-reported] [source: https://nvidianews.nvidia.com/news/nvidia-bluefield-4-powers-new-class-of-ai-native-storage-infrastructure-for-the-next-frontier-of-ai ].
- Precedent: BlueField-2/3 already used as NVMe-oF target offload (SNAP), virtio-blk; DOCA 2.9 adds alpha target-offload for BF-3 storage controller [official].

### 15.3 Decision framework — which NIC/DPU/IPU

| Workload | Leading option (2026) | Rationale |
|---|---|---|
| AI training scale-out, NVIDIA GPUs | ConnectX-8 SuperNIC / BlueField-3 E-series; BF-4 from 2026 | 800G, Multi-Host, Spectrum-X tuning, DOCA |
| AI on AMD GPUs (Helios) | Pollara 400 (2026), Vulcano 800 (2027) | UEC 1.0 compliance, RCCL integration |
| Cloud C3-class VMs (Google) | Intel E2000 (Mount Evans) | PSP encryption, P4 datapath, Google-qualified |
| Telco/5G UPF, vRouter | Marvell OCTEON 10 (CN103/DPU400) | Inline crypto, 1T switch, VPP |
| Enterprise virtualization/storage | BlueField-3 P-series | OVS, NVMe-oF, virtio 1K scale, DMS |
| PTP/capture appliances | FPGA SmartNICs (Napatech, Silicom) | Deterministic timestamping |
| Sovereign/AI-cloud builders | BlueField-3/4 or OCTEON | DPU isolation + secure boot + BMC |

### 15.4 Operational risks

- Firmware/driver pinning: BF-3 requires DOCA-Host (not MLNX_OFED); PSID-specific firmware; PLDM update path needs full power cycle in 2.9.x [official].
- Grey-market NICs may lack vendor warranty/support; verify PSID against NVIDIA firmware compatibility list [secondary].
- DOCA 2.9.0 does not support ConnectX-8 — mixing CX-8 with older DOCA branches breaks support [official].
- P-series dual-slot cards need chassis clearance + auxiliary power; validate against OEM HCL (Lenovo LP1809, HPE P71949-B21 cable kits) [official].
- UEC 1.0 is young: errata through 1.0.3 (Jul 2026); interoperability testing ongoing (VIAVI validation tools launched June 2026); production multi-vendor UEC fabrics **not yet demonstrated at scale** as of 2026-09-22 [unverified].
- TCO comparison claims (e.g. AMD's 33% vs Broadcom-based 800G) are vendor-internal; demand independent benchmarks before procurement [secondary].

---
*End of Phase F1. 15 waves, append-only. Research cutoff 2026-09-22. Single writer; no other workspace files modified.*

## Wave 16 — Broadcom: Thor 2, Thor Ultra, NetXtreme E-Series, Stingray status

### 16.1 Broadcom Thor 2 (400G AI NIC, launched May 2024)

- Broadcom unveiled what it claimed was the **first 5 nm Ethernet NIC built for AI data centers** (May 2024): low-power 400G PCIe Gen5 Ethernet NIC resolving XPU-bandwidth/cluster-size bottlenecks [secondary] [source: https://electronics360.globalspec.com/article/21074/ethernet-nic-unveiled-by-broadcom-for-ai-data-centers ].
- Foundation is the **Thor 2 ASIC SoC** doing all AI NIC processing — fully hardware-accelerated, high performance + low power; offered as **board product, chiplet, or IP** for integration flexibility [secondary] [source: https://www.sdxcentral.com/news/broadcoms-400g-ethernet-nic-engineered-for-unique-ai-requirements/ ].
- Third-generation RoCE pipeline; enhances RDMA with **UEC-derived techniques**: out-of-order packet placement with in-order completion, selective ACK/retransmission, packet-level multipathing across paths, configuration-free congestion-control algorithms [secondary].
- 8 SerDes lanes supporting 100/50G PAM4 and 25G NRZ; passive copper DAC reach up to 5 m; ultra-low-power linear pluggable optics; designed to be driven on Tomahawk 5 switches with extended DAC reach [secondary] [source: https://www.fs.com/uk/blog/fs-broadcom-400g-rocerdma-ethernet-nic-revolutionizes-ai-networking-landscape-24061.html ].
- ServTheHome context: 400G matches PCIe Gen5 x16 throughput; offloads/RDMA essential since CPUs can't service multiple 400G NICs without them [secondary] [source: https://www.servethehome.com/broadcom-400gbe-nics-launched-for-the-ai-era/ ].

### 16.2 Broadcom Thor Ultra (800G AI backend NIC, ~2025)

- **Thor Ultra: 800G Ethernet NIC purpose-built for AI backend networks** — clean-sheet design, not a Thor 2 evolution; focused exclusively on AI scale-out (Thor 2 also served enterprise) [secondary] [source: https://www.networkworld.com/article/4072253/broadcom-drops-the-hammer-on-ai-networking-with-thor-ultra.html ].
- Implements **UEC 1.0 specifications**; hardware-accelerated RDMA modernization. Broadcom frames it as completing a 3-year AI-networking portfolio build: **Tomahawk 6** switches (scale-out) + **Jericho 4** (inter-DC) + Thor Ultra NIC.
- Broadcom quote (Hasan Siraj, head of software products & ecosystem): "this NIC is fully compliant with Ultra Ethernet features right at 800 gig, and there is nothing else in the industry that can cater to this" — vendor claim, not independently verified [vendor-reported].
- Note: this fills the "Broadcom 2026 roadmap" gap flagged in Wave 6 — current-generation NIC is **Thor Ultra 800G**, not a Stingray refresh.

### 16.3 NetXtreme E-Series (merchant NIC portfolio)

- BCM957508-**P2200G**: dual-port 200GbE QSFP56, PCIe 4.0 x16, BCM57508 200GbE MAC + integrated dual-channel 200GbE SFI transceiver; port configs 2×100/50/25/10GbE and 2×200GbE; 400G connectivity for HA features [official/secondary] [source: https://www.primeline-solutions.com/media/datasheet/broadcom-netxtreme-gigabit-ethernet.pdf ].
- BCM957508-**P2100G**: dual-port 100GbE QSFP56, PCIe 4.0 x16 [official/secondary].
- Feature set: **SR-IOV up to 1,000 VFs**, multihost (dual-host x8), TruFlow flow processing engine, on-chip tunneling for Geneve/VXLAN/NVGRE, tunnel-aware stateless offloads, vSwitch acceleration, DCB (PFC/ETS/QCN), GPUDirect acceleration, RoCE v1/v2 [official/secondary].
- BCM957414A4140C **P150p**: single-port 50GbE QSFP28, PCIe 3.0 x8; BCM957504-**P425G**: 4×25GbE SFP28, PCIe 4.0 x16, multihost [official/secondary].

### 16.4 Stingray (SmartNIC lineage) status

- Broadcom **NetXtreme-S Stingray PS410T-H04** (4×10G 10GBase-T PCIe SmartNIC) still receiving Windows driver releases: v236.1.152.0 dated **17 December 2025** (bnxtnd.inf) [secondary] [source: https://treexy.com/products/driver-fusion/database/network-adapters/broadcom/ps410t-h04-netxtreme-s-stingray-4x10g-10gbase-t-pcie-smartnic/ ].
- Assessment: Stingray remains a maintenance-mode enterprise SmartNIC line (10G-class parts); Broadcom's 2026 AI-networking push rides on Thor 2/Thor Ultra, not Stingray [independent].
- AMD competitive claim to watch: Pollara/Vulcano pitched against "Broadcom Tomahawk-6-based 800G" switching with 33% lower cost — Broadcom disputes by positioning Thor Ultra as fully UEC-compliant at 800G [secondary] [source: https://www.networkworld.com/article/4072253/broadcom-drops-the-hammer-on-ai-networking-with-thor-ultra.html ].


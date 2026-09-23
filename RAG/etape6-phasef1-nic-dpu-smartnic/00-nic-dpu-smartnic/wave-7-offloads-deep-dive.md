---
id: etape6-phasef1-nic-dpu-smartnic/00-nic-dpu-smartnic/wave-7-offloads-deep-dive
title: "Wave 7 — Offloads deep-dive"
domain: phase-f1-nic-dpu-and-smartnic-nvidia-amd-pensando-intel-ipu-
role: deep-dive
task: pricing
actors: ["AMD", "AWS", "Baidu", "Broadcom", "CoreWeave", "Google", "Intel", "Microsoft", "Nvidia", "Oracle"]
dates: []
keywords: ["amd", "asic", "aws", "compute", "disaggregated", "ethernet", "gpu", "helios", "hyperscaler", "intel", "latency", "maia"]
source: docs/RAG/etape6_phaseF1_nic_dpu_smartnic.md
source_anchor: ""
source_lines: [204, 271]
section: "Phase F1 — NIC, DPU and SmartNIC (NVIDIA, AMD Pensando, Intel IPU, Marvell, offloads, pricing)"
sha256: fd63439bfec4192ba1807f8622948381a82179a1fc965a9d63c82f1ab6d62956
---

# Wave 7 — Offloads deep-dive

## Wave 7 — Offloads deep-dive

### 7.1 Taxonomy of offloads (SmartNIC → DPU → SuperNIC)

- **Traditional SmartNIC**: partial offload — datapath on the card, control plane on host CPU. **DPU SmartNIC**: complete offload — datapath and control plane run on the DPU's embedded CPU/OS; a separate compute domain for infrastructure [secondary] [source: https://www.fs.com/blog/traditional-smartnic-vs-dpu-smartnic-5408.html ].
- NVIDIA's framing: ASIC adapters (ConnectX) for "well-defined, price-power-performance efficient" offloads; BlueField DPUs combine Arm flexibility + ConnectX offloads for the software-defined (SDX) world [official] [source: https://developer.nvidia.com/blog/?p=23409 ].
- Why offload matters (2025 industry summary): east–west traffic dominance, encryption-everywhere (zero-trust), AI/ML GPU starvation waiting on data, NVMe-oF pushing storage I/O into the network; without offload CPUs burn cycles on packet handling/crypto/encapsulation [independent] [source: https://network-switch.com/blogs/networking/smart-nics-and-dpus-explained ].

### 7.2 Networking offloads

- **RDMA / RoCE**: ConnectX line's core; RoCEv2 acceleration on ConnectX-8; Zero-Touch RoCE on BlueField-3 [official]. RoCE vs InfiniBand: same Verbs API, Ethernet transport; congestion control (DCQCN) handled in NIC [independent].
- **Overlay offloads**: VXLAN/GENEVE/NVGRE encap/decap in hardware (ASAP² on NVIDIA; Pensando P4 pipeline; Intel IPDK) [vendor-reported].
- **vSwitch/OVS offload**: OVS datapath offloaded via ASAP² (NVIDIA), P4 (Pensando), FPGA (Intel Oak Springs Canyon / Altera); BlueField "bump-in-the-wire" architecture runs vSwitch/vRouter control+dataplane on the DPU [official].
- **SR-IOV**: supported across ConnectX-8, BlueField-2/3, OCTEON 10; VirtIO acceleration on BlueField-3 [vendor-reported].
- **Flow processing**: connection tracking (L4 firewall), NAT, hierarchical QoS, header rewrite, flow mirroring, flow aging/statistics [vendor-reported].
- **Programmable congestion control**: ConnectX-8 advanced congestion control; Pensando hardware-based congestion control with intelligent multipathing; DCQCN as the de-facto RoCE mechanism [vendor-reported/secondary].
- **MPI/GPU**: MPI tag matching (BlueField-3), GPUDirect RDMA/Storage, NCCL/RCCL offload assist, collectives offload (ConnectX-8 network compute) [vendor-reported].

### 7.3 Storage offloads

- **NVMe-oF**: initiator/target offload on BlueField (SNAP storage virtualization), Intel Mount Evans NVM interface emulating NVMe devices; iSER over RDMA [vendor-reported].
- **NVMe/TCP** vs NVMe/RoCE: NIC selection guidance — RDMA NICs for high-performance low-latency (NVMe/RoCE), TOE NICs as a mid-point for NVMe/TCP [independent] [source: https://intelligentvisibility.com/ethernet-storage-nic-guide ].
- **Erasure coding / compression / decompression**: hardware engines on BlueField-3 (storage offload), Pensando DSC line [vendor-reported].
- Fungible's NVMe-over-TCP disaggregated storage cluster is the canonical example of DPU-accelerated composable storage [secondary].

### 7.4 Security offloads

- **Inline crypto**: IPsec and TLS "data-in-motion" crypto on ConnectX-6 Dx/BlueField-2; MACsec/IPsec/TLS + PKA root-of-trust on BlueField-3; true inline crypto on OCTEON 10 (most of IPsec in hardware; 50 Gbps IPsec at 50% of a single N2 core per Marvell) [vendor-reported].
- **Zero-trust**: DPUs as policy enforcement points — network isolation between tenants at the NIC; Google C3's PSP-encrypted 200G networking via E2000 [vendor-reported/secondary].
- **Secure boot / RoT**: BlueField crypto-enabled SKUs include secure boot + hardware root of trust; OCTEON 10 eHSM/TrustZone [vendor-reported].

### 7.5 Operational reality check [independent]

- A 2026 practitioner note (dev.to): DPUs win where network/storage I/O genuinely steals CPU cycles — multi-tenant infra (vSwitch offload), storage-heavy clusters (NVMe-oF/erasure coding), high-throughput security (inline TLS/IPsec); below ~10 Gbps per host a DPU is "a Ferrari in traffic" [independent] [source: http://dev.to/renato_silva_71eef0fc385f/npu-dpu-qpu-which-one-actually-belongs-in-your-stack-3a07 ].
- Overhead: DPU = another programmable device with its own firmware, OS (Arm Linux), and failure modes; requires dedicated infra-team capacity [independent].

---
*End of Waves 6–7. Next: competitive landscape & adoption (Wave 8), open items/conflicts/glossary/source index (Wave 9).*

## Wave 8 — Competitive landscape and adoption (2026)

### 8.1 Hyperscaler and CSP adoption

- **Oracle Cloud Infrastructure (OCI)**: selected **NVIDIA BlueField-3 DPU** as the latest addition to its networking stack (announced at GTC); OCI frames it as part of its long-established approach of offloading infrastructure tasks from CPUs [official] [source: https://nvidianews.nvidia.com/news/oracle-cloud-infrastructure-chooses-nvidia-bluefield-data-center-acceleration-platform ].
- **Google Cloud**: co-designed Intel E2000 (Mount Evans) and deploys it in **C3 VMs** with 4th Gen Xeon; 200G PSP-encrypted networking [official/secondary].
- **Microsoft Azure**: in-house SmartNIC lineage (Azure Catapult, now FPGA+ASIC generations); absorbed Fungible's DPU team (2023); co-designs Maia AI ASICs with Marvell [secondary]. Threatens AMD Pensando business at Azure per Blocks&Files analysis [secondary].
- **AWS**: Nitro system (custom DPU lineage, v5/v6 generations) — offloads storage, networking, security; not sold merchant [secondary].
- **Oracle + AMD**: first CSP to deploy **Pensando Pollara 400** (UEC NIC) [vendor-reported].
- **Baidu, JD.com, UCloud**: BlueField DPU users per NVIDIA ecosystem list [official].
- **CoreWeave, Palo Alto Networks**: listed in NVIDIA's BlueField ecosystem at GTC DC 2025 (BlueField-4 briefing) [vendor-reported] [source: https://www.sdxcentral.com/news/nvidia-reveals-next-gen-dpu-to-help-offload-gigascale-ai-infrastructure/ ].
- **VAST Data**: runs its entire storage stack directly on BlueField hardware; storage partners (VAST, DDN, WEKA, etc.) validate the DPU-as-storage-controller pattern [vendor-reported/secondary] [source: https://www.sdxcentral.com/analysis/nvidias-bluefield-4-a-first-look-at-the-dpu-built-to-run-ai-factories/ ].

### 8.2 Pensando enterprise footprint

- Pensando's distributed services platform deployed "at scale" across cloud and enterprise clients including **Goldman Sachs, IBM Cloud, Microsoft Azure, Oracle Cloud** [secondary] [source: https://medium.com/@lixian_58397/the-most-comprehensive-dpu-smartnic-vendors-with-its-product-line-summary-db7899a725c8 ].
- **Pensando Elba SoC**: DPU for intelligent network switches (shipping) [secondary].
- **Pensando Capri DPU**: used in **Aruba CX 10000** switch series (distributed services switch) [secondary].
- Helios roadmap (AMD deep-dive, Sept 2026): Pensando Elba = shipping; **Pensando Vulcano = next-gen DPU, designed into Helios, H2 2026**; Helios rack = 72× MI455X + Venice EPYC + Vulcano NICs [secondary] [source: https://github.com/redhat-et/physical-ai-platform-intel/blob/HEAD/deliverables/intel/companies/amd-deep-dive.md ].

### 8.3 Competitive dynamics

- **AI networking bifurcation**: (a) NVIDIA vertical stack (ConnectX-8 SuperNIC / BlueField-4 + Spectrum-X + InfiniBand) vs (b) open Ethernet (AMD Pollara/Vulcano + UEC + Broadcom Tomahawk 6 switching + Arista 7800R4). Competition centers on software programmability, congestion control, transport, TCO — not raw bandwidth [secondary/analysis].
- **BlueField-4 topology detail** (SDxCentral, from NVIDIA director): 800 Gb/s total, typically **two 400 Gb/s bidirectional links**; Ethernet at the front end for north–south; GPU east–west handled by **multiple ConnectX-9 chips packaged together delivering 1.6 Tb/s per GPU** starting with Rubin; BlueField-4's 800G seen as sufficient headroom for the Rubin generation [vendor-reported/secondary].
- **BlueField-5** already on roadmap: slated **2028** with NVIDIA's **Feynman** architecture, per NVIDIA's "annual rhythm" cadence [vendor-reported].
- DPU market forecast: $554M (2021) → **$5.5B by 2031** at 27% CAGR (Allied Market Research via DataCenterKnowledge) [secondary]; 2026 realized revenue not located [gap].
- "Third socket" narrative: DPU as infrastructure root-of-trust, first line of defense, primary data orchestrator — freeing 100% of host CPU for applications [independent/analysis] [source: https://github.com/jorgeadev/the-architects-blueprint/blob/HEAD/web/src/content/blog/2026/04/28/beyond-the-cpu-architecting-hyperscale-analytics-.md ].
- Vendor landscape summary (merchant + in-house): NVIDIA BlueField, AMD Pensando, Intel IPU, Marvell OCTEON, Broadcom Stingray, Fungible (→Microsoft), Napatech (→Alphawave Semi), Xsight Labs, Kalray, Ethernity, Chelsio, Silicom, Achronix, Netronome, Asterfusion; in-house: AWS Nitro, Azure SmartNIC, Google IPU [secondary] [source: https://medium.com/@lixian_58397/the-most-comprehensive-dpu-smartnic-vendors-with-its-product-line-summary-db7899a725c8 ].


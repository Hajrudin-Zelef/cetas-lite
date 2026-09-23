---
id: etape6-phasef1-nic-dpu-smartnic/00-nic-dpu-smartnic/wave-9-open-items-conflicts-glossary-source-index
title: "Wave 9 — Open items, conflicts, glossary, source index"
domain: phase-f1-nic-dpu-and-smartnic-nvidia-amd-pensando-intel-ipu-
role: deep-dive
task: pricing
actors: ["AMD", "Broadcom", "Google", "Intel", "Nvidia"]
dates: ["2026-09-22"]
keywords: ["accelerator", "acquisition", "amd", "asic", "benchmark", "compute", "datacenter", "ethernet", "gpus", "helios", "intel", "ipo"]
source: docs/RAG/etape6_phaseF1_nic_dpu_smartnic.md
source_anchor: ""
source_lines: [272, 360]
section: "Phase F1 — NIC, DPU and SmartNIC (NVIDIA, AMD Pensando, Intel IPU, Marvell, offloads, pricing)"
sha256: 647b68d30fb84b5bf49b10bd57bc8f2e7baccfa301795ad28edf2768058dfe34
---

# Wave 9 — Open items, conflicts, glossary, source index

## Wave 9 — Open items, conflicts, glossary, source index

### 9.1 Open items / gaps

1. BlueField-4 retail/OEM pricing — not found (2026-09-22).
2. DOCA current major version (2.9.2 in docs archive; no 2026 major-version confirmation).
3. Intel Mount Morgan / Hot Springs Canyon shipment confirmation; 800G IPU status; F2000X-PL part number — unconfirmed.
4. Marvell OCTEON 12 — not confirmed (task scope mentioned it; search found nothing).
5. Broadcom Stingray 2026 roadmap — no current public material.
6. Post-acquisition Fungible product status — no standalone product line observed.
7. Napatech→Alphawave Semi acquisition — noted but not verified in this wave.
8. 2026 realized DPU market revenue — forecast only.
9. AMD "Pensando 800G" 2027 roadmap item — secondary source only.
10. ConnectX-8 ecer.com $2,500 listing — third-party brand ("Comelink"), grey-market risk.

### 9.2 Conflicts registered

- C1: Grabnpay BF-3 listing says 16 GB DDR5 while NVIDIA docs list 32 GB on P-series — sub-model confusion [unverified].
- C2: SHI commercial ($3,053) vs SHI public-sector ($4,425) for the same BF-3 B3210E SKU — channel pricing variance, not an error [secondary].
- C3: AMD's "33% lower switching costs" (Vulcano vs Tomahawk-6-based 800G) is an internal comparison, not an independent benchmark [secondary].
- C4: BlueField-4 launch date — "early availability 2026 as part of Vera Rubin" (NVIDIA) vs H2 2026 storage-platform availability; consistent but staged [vendor-reported].

### 9.3 Glossary

- **DPU** (Data Processing Unit): SoC combining Arm CPU cores, ASIC network/storage/security accelerators, and high-speed I/O; runs its own OS; offloads infrastructure workloads from the host CPU. NVIDIA's term.
- **IPU** (Infrastructure Processing Unit): Intel's term for the same class; ASIC (Mount Evans) and FPGA (Oak Springs Canyon) variants.
- **SmartNIC**: NIC with programmable offloads (FPGA or ASIC); partial offload (datapath only) in traditional definition.
- **SuperNIC**: NVIDIA's term for ConnectX-8-class NICs optimized for AI/HPC scale-out (800G, Multi-Host, network compute).
- **ASAP²**: NVIDIA Accelerated Switching and Packet Processing — OVS/overlay/flow offload framework.
- **DOCA**: NVIDIA's DPU software framework (SDK, runtime, orchestration).
- **IPDK**: Intel Infrastructure Programmer Development Kit — P4/DPDK/SPDK-based DPU programming model.
- **RoCE / RoCEv2**: RDMA over Converged Ethernet; v2 runs over UDP/IP.
- **NVMe-oF**: NVMe over Fabrics — transports: RDMA (RoCE/iWARP/IB), FC, TCP.
- **SR-IOV**: Single Root I/O Virtualization — PCIe device sharing across VMs.
- **PSP**: Google's open encryption protocol for line-rate network encryption (used with E2000).
- **UEC**: Ultra Ethernet Consortium — open Ethernet spec for AI/HPC; 1.0 published 2025.
- **NIXL / Dynamo**: NVIDIA libraries for KV-cache/context sharing (BlueField-4 storage platform).

### 9.4 Source index (verbatim URLs)

- https://nvidianews.nvidia.com/news/nvidia-extends-data-center-infrastructure-processing-roadmap-with-bluefield-3
- https://docs.nvidia.com/doca/archive/2-9-2/doca+installation+guide+for+linux/index.html
- https://github.com/nvidia/doca-platform/blob/HEAD/docs/public/platform-support.md
- https://www.servethehome.com/nvidia-bluefield-4-with-64-arm-cores-and-800g-networking-announced-for-2026/
- https://www.servethehome.com/nvidia-bluefield-4-processor-at-hot-chips-2026/
- https://nvidianews.nvidia.com/news/nvidia-bluefield-4-powers-new-class-of-ai-native-storage-infrastructure-for-the-next-frontier-of-ai
- https://nvidianews.nvidia.com/news/oracle-cloud-infrastructure-chooses-nvidia-bluefield-data-center-acceleration-platform
- https://www.sdxcentral.com/news/nvidia-reveals-next-gen-dpu-to-help-offload-gigascale-ai-infrastructure/
- https://www.sdxcentral.com/analysis/nvidias-bluefield-4-a-first-look-at-the-dpu-built-to-run-ai-factories/
- https://www.techpowerup.com/342340/nvidia-launches-bluefield-4-dpus-with-800-gb-s-throughput-for-ai-data-centers
- https://itdaily.com/news/datacenter/nvidia-bluefield-4-dpu/
- https://www.networkworld.com/article/4011095/amd-rolls-out-first-ultra-ethernet-compliant-nic.html
- https://www.techradar.com/pro/amd-debuts-a-400gbe-ai-network-card-with-an-800gbe-pcie-gen6-nic-coming-in-2026-but-will-the-industry-be-ready
- https://infotechlead.com/networking/amd-pensando-vulcano-800-challenges-nvidia-connectx-8-as-ai-networking-race-heats-up-broadcom-arista-and-marvell-intensify-competition-97488
- https://finance.biggo.com/news/1bad314f-c6c1-4284-b172-100bc3930dbf
- https://www.techtimes.com/articles/313781/20260106/ces-2026-amd-details-helios-ai-rack-next-gen-instinct-mi400-gpus.htm
- https://techly.relevanpress.com/amd-announces-pensando-salina-400-dpu-_-pollara-40--129573/
- https://github.com/redhat-et/physical-ai-platform-intel/blob/HEAD/deliverables/intel/companies/amd-deep-dive.md
- https://www.datacenterdynamics.com/en/news/intel-and-google-cloud-jointly-launch-data-center-accelerator-chip/
- https://www.hpcwire.com/2022/05/10/intel-extends-ipu-roadmap-through-2026/
- https://www.servethehome.com/intel-ipu-plans-revealed-for-800gbps-ipus-in-2025-dpu/
- https://www.blocksandfiles.com/ai-ml/2023/10/04/intel-semi-detaching-altera-business-for-future-ipo/1590209
- http://www.marvell.com/content/dam/marvell/en/public-collateral/embedded-processors/marvell-octeon-10-dpu-platform-product-brief.pdf
- https://www.techPowerUp.com/283905/marvell-extends-octeon-leadership-with-industrys-first-5nm-dpus
- https://fuse.wikichip.org/news/5732/marvell-launches-5nm-octeon-10-dpus-with-neoverse-n2-cores-ai-acceleration/
- https://www.semiaccurate.com/2021/06/28/marvell-announces-their-5nm-octeon-10-dpu/
- https://cloudswit.ch/blogs/marvell-octeon-10-cn103xx-open-router-gateway/
- https://www.datacenterknowledge.com/deals/microsoft-extends-data-center-ip-with-acquisition-of-dpu-startup-fungible
- https://www.networkworld.com/article/971816/microsoft-to-acquire-fungible-for-augmenting-azure-networking-storage.html
- https://www.FS.com/c/nvidia-ethernet-nics-4014
- https://www.fs.com/uk/c/nvidia-ethernet-nics-4014
- https://www.naddod.com/collections/nvidia-networking/infiniband-adapters
- https://www.shi.com/product/44101934/NVIDIA-BlueField-2-DPU-E-Series-SmartNIC-MBF2M516A-CECOT
- https://www.shi.com/product/43607275/NVIDIA-BlueField-2-Ethernet-DPU-MBF2M516A-CENOT
- https://www.shi.com/Product/46801187/NVIDIA-BLUEFIELD-3-B3210E-E-SERIES-FHHL-DPU
- https://cdn.pathfactory.com/assets/10412/contents/1195564/153be327-b398-454e-b6af-6a0aed28d960.pdf
- https://developer.nvidia.com/blog/?p=23409
- https://www.fs.com/blog/traditional-smartnic-vs-dpu-smartnic-5408.html
- https://network-switch.com/blogs/networking/smart-nics-and-dpus-explained
- https://intelligentvisibility.com/ethernet-storage-nic-guide
- https://medium.com/@lixian_58397/the-most-comprehensive-dpu-smartnic-vendors-with-its-product-line-summary-db7899a725c8
- https://www.ebay.com/itm/306620041473
- https://www.lambda-tek.com/Nvidia-900-9D3B6-00CV-AA0~sh/B48707817
- https://www.compsource.com/buy/9009D3B600CCEA0/Mellanox-4134/Nvidia-Bluefield3-B3210e-ESeries-Fhhl-Dpu-9009D3B600CCEA0-9009D3B600CCEA0/
- https://www.pc-canada.com/item/nvidia-bluefield-3-b3220-p-series-fhhl-dpu-200gbe-default-mode-ndr2-00-ib-dual/900-9d3b6-00sv-aa0-dell

---
*End of Phase F1. 9 waves, append-only. Research cutoff 2026-09-22. Single writer; no other workspace files modified.*


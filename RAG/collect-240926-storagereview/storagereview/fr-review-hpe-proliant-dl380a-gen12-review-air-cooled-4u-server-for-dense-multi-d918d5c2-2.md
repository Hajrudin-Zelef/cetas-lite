---
id: collect-240926-storagereview/storagereview/fr-review-hpe-proliant-dl380a-gen12-review-air-cooled-4u-server-for-dense-multi-d918d5c2-2
title: "fr-review-hpe-proliant-dl380a-gen12-review-air-cooled-4u-server-for-dense-multi--d918d5c2"
domain: storagereview
role: reference
task: reference
actors: ["Intel", "Nvidia"]
dates: []
keywords: ["blackwell", "compute", "distribution", "fine-tuning", "fp4", "governance", "gpu", "gpus", "inference", "intel", "licenses", "liquid cooling"]
source: docs/RAG/clean_en/storagereview/fr-review-hpe-proliant-dl380a-gen12-review-air-cooled-4u-server-for-dense-multi--d918d5c2.md
source_anchor: ""
source_lines: [3, 43]
sha256: 43d61ff2e9997eafdac066cad51e917cbb1f9b7dd7da50164bf1b2c725db9995
---

# fr-review-hpe-proliant-dl380a-gen12-review-air-cooled-4u-server-for-dense-multi--d918d5c2

The HPE ProLiant Compute DL380a Gen12 server is aimed at enterprise AI teams looking for high compute density without changing their rack configuration. This 4U air-cooled chassis integrates easily, supports up to eight double-width GPUs, and offers full PCIe Gen5 connectivity. It can be configured with two Intel Xeon 6 processors, each with 144 cores, 4 TB of DDR5 memory across 32 DIMM modules, and sixteen E3.S NVMe bays for high throughput and capacity. The goal is simple: achieve production-grade inference capacity and precise fine-tuning at scale without resorting to liquid cooling.
For accelerators, HPE offers a full range including NVIDIA H200 NVL, H100 NVL, L40S, L20, L4 cards and the RTX PRO 6000 Blackwell Server Edition, with power options compatible with high-consumption components. In this test, we focus on the RTX PRO 6000 Server, which offers an excellent compromise for enterprise AI. Each card carries 96 GB of ECC GDDR7 memory, a PCIe Gen5 x16 interface, FP4-compatible Tensor cores, and a 600 W thermal envelope, suited to air-cooled racks. Our configuration was equipped with four cards, a sensible starting point for high-throughput inference and targeted optimization, with room for growth.
HPE completes the platform with essential operational elements. iLO 7 handles out-of-band configuration, system status, and power management, thanks to the Silicon Root of Trust (a secure enclave ensuring firmware integrity), RSA 4096-bit encryption support, and a detachable iLO DC-MHS module that strengthens supply chain verification. The server also integrates with the HPE Private Cloud AI platform for multi-team governance and reproducible deployments at scale.
HPE ProLiant Compute DL380a Gen12 – Technical Specifications
| Categories | Specifications | 
|---|---|
| Processor type | HPE ProLiant Compute DL380a Gen12 | 
| Processor family | 6th Generation Intel® Xeon® Scalable Processors | 
| Available processor cores | 64 to 144 cores, depending on processor | 
| Number of processors | 2 | 
| Processor speed | Up to 2.4 GHz, depending on processor | 
| Maximum memory | RDIMM 4 TB (2 TB per processor) | 
| Memory slots | 32 DIMM slots | 
| Memory type | HPE DDR5 Smart Memory | 
| Memory protection | RAS: Advanced ECC, online spare memory, mirroring, combined channel functionality (locking), HPE Fast Fault-Tolerant Memory (ADDDC) | 
| Drive support | SFF NVMe and EDSFF | 
| Security | Optional locking bezel, intrusion detection, and integrated HPE TPM 2.0 module | 
| Infrastructure management | HPE iLO Standard with intelligent provisioning (embedded), HPE OneView Standard (requires download) • Optional: HPE iLO Advanced and HPE OneView Advanced (licenses required) | 
| Power supply | Up to 8 M-CRPS. Single 1+1 redundancy for the motherboard. Dual 2+1 redundancy for GPUs. | 
| Expansion connectors | 6 | 
| System fans | 4 dual-rotor fans and 8 hot-swappable single-rotor fans included | 
| Form factor | 4U rack | 
| Warranty | 3/3/3: Server warranty | 
Design and assembly of the HPE ProLiant DL380a Gen12 server
The HPE ProLiant Compute DL380a Gen12 rack server is a 4U dual-socket server designed for scalable, high-performance deployments. Measuring 6.88 x 17.63 x 31.60 cm, it combines dense CPU and GPU compute power with efficient air cooling for reliable operation even under heavy workloads.
Weighing between 82.7 and 137.8 kg depending on configuration, the chassis supports high-capacity components, redundant power, and offers easy front access for maintenance. Its design prioritizes performance, scalability, and efficient thermal management, making it ideal for enterprise and data center environments.
On the storage side, the HPE ProLiant DL380a Gen12 offers 4- or 8-bay configurations in SFF or EDSFF formats. Our test model was equipped with the HPE DL380a Gen12 NS204i-u front bay kit, supporting two hot-swappable NVMe M.2 boot devices. The chassis also included eight 2.5-inch bays, occupied by two HPE U.3 SSDs with a total capacity of 15.36 TB. HPE offers several front bay options, providing great flexibility to adapt to different deployment needs.
The unit can be transported using its two side handles, which means at least two people are required for safe racking and installation. It uses a 2U rail kit with telescopic rails, allowing for easy installation and simplified maintenance without fully dismantling the rack.
At the rear of the HPE ProLiant DL380a Gen12 server, the optimized layout promotes airflow, expandability, and ease of maintenance. The system supports up to eight MCRPS power supplies (1 to 8) thanks to an integrated ventilation panel, ensuring optimal cooling even at full load. Expandability is significant, with multiple PCIe Gen5 x16 slots (slots 1 to 6) compatible with integrated and optional expansion cards, as well as OCP A and B slots for flexible network adapter configuration.
Connectivity includes a dedicated iLO network port, multiple USB 3.2 Gen 1 ports, and a VGA port for local management. It is important to note that slot 1 is available only when the HPE DL380a Gen12 4EDSFF Direct Cable for NVD (P74716-B21) is installed and cannot be used with SFF NVMe drives. As for slot 4, it is not supported in configurations with 4 or 8 DW GPUs.
Power for the HPE ProLiant DL380a Gen12 server is provided by hot-swappable M-CRPS Titanium modular power supply kits. Compatible models include the 1,500 W (P67244-B21), 2,400 W (P67252-B21), and 3,200 W (P67248-B21) versions. The system supports up to eight power supplies, offering N+1 redundancy to ensure continuous operation even if a power module fails. Power requirements and distribution may vary depending on GPU configuration. Our test model was equipped with five 2,400 W M-CRPS power supplies, providing sufficient capacity to power the system's four GPUs (600 W TDP) while ensuring reliable redundancy.
Observing the interior of the HPE ProLiant DL380a Gen12 from the top, it is clear that HPE designed this chassis with a focus on cooling the GPUs, positioned at the front of the system to benefit from direct, unobstructed airflow. The cooling system includes four hot-swappable fan groups, each consisting of a 92 x 56 mm dual-rotor fan and two 40 x 28 mm single-rotor fans. The smaller fans concentrate airflow on the processor and memory modules, ensuring efficient thermal management of lower components. By comparison, the larger dual-rotor fans are specifically designed to generate significant airflow directly onto the GPU array. This balanced design ensures optimal cooling of compute and acceleration components, even under high and sustained workloads.
When examining the GPU configuration, our unit was pre-wired for four PCIe 5.0 GPUs, each cleanly installed in the front GPU cage. The system was configured with NVIDIA RTX PRO 6000 GPUs (Blackwell Server Edition, 96 GB), belonging to NVIDIA's new professional range optimized for AI, rendering, and compute workloads. Depending on configuration, the DL380a Gen12 can support 4 or 8 double-width GPUs or up to 16 single-width accelerators, offering great flexibility for a wide range of enterprise and AI deployments.
The GPUs compatible with this platform are as follows:
- NVIDIA RTX PRO 6000 Server Edition (96 GB)
- NVIDIA H200 NVL (141 GB)
- NVIDIA H100 NVL (94 GB)
- NVIDIA L40S (48 GB)
- NVIDIA L20 (48 GB)
- NVIDIA L4 (24 GB)
This flexible GPU configuration, combined with high-bandwidth PCIe Gen5 lanes, ensures that the DL380a Gen12 is ready for dense inference tasks and large-scale AI training environments. Inside the chassis, once the cooling shroud is installed, HPE's ingenuity in airflow management becomes apparent. This shroud features precisely molded deflectors that efficiently direct air toward the processors, memory modules, and VRMs, ensuring uniform system cooling.

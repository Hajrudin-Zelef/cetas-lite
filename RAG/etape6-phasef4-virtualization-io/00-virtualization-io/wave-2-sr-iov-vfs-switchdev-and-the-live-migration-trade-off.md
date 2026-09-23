---
id: etape6-phasef4-virtualization-io/00-virtualization-io/wave-2-sr-iov-vfs-switchdev-and-the-live-migration-trade-off
title: "Wave 2 — SR-IOV: VFs, switchdev, and the live-migration trade-off"
domain: phase-f4-i-o-virtualization-cpu-acceleration-extensions
role: deep-dive
task: reference
actors: ["AMD", "AWS", "Nvidia"]
dates: []
keywords: ["amd", "aws", "gpu", "latency", "memory", "neocloud", "nvidia", "research", "throughput"]
source: docs/RAG/etape6_phaseF4_virtualization_io.md
source_anchor: ""
source_lines: [76, 132]
section: "Phase F4 — I/O Virtualization & CPU Acceleration Extensions"
sha256: a0e71d1ba478ed4fcdf8f170762b30ed7272e6508bbc4de857bb1511982e54a7
---

# Wave 2 — SR-IOV: VFs, switchdev, and the live-migration trade-off

## Wave 2 — SR-IOV: VFs, switchdev, and the live-migration trade-off

### 2.1 SR-IOV fundamentals
- Single Root I/O Virtualization lets one PCIe physical device present itself as multiple PCIe functions: one full-featured **Physical Function (PF)** managed by the host, plus lightweight **Virtual Functions (VFs)** that guests use directly [official].
  Source: https://docs.nvidia.com/networking/display/mlnxenv24040660/single+root+io+virtualization+(sr-iov)
- NVIDIA ConnectX family adapters expose **up to 127 VFs per port**; each VF appears as an additional device sharing the PF's physical resources [vendor-reported].
  Source: https://docs.nvidia.com/networking/display/mlnxenv24040660/single+root+io+virtualization+(sr-iov)
- ConnectX-6 Dx advertises "up to 1K virtual functions per port" and 8 physical functions (product brief) — the highest VF density in the ConnectX line [vendor-reported].
  Source: https://solutions.asbis.com/api/uploads/files/40/pb-connectx-6-dx-en-card.pdf
- The payoff: guest drivers talk to hardware with **almost no VM exits** — community reference figures put SR-IOV at 95–99% of bare-metal throughput vs ~70–90% for virtio and 20–40% for full emulation [secondary].
  Source: https://github.com/denisergocmen924/cloud_fundamentals/blob/HEAD/Cloud%20Learning%20With%20Workbooks/Cloud_hardware.roadmap/English_Hardware/Phase_6_Virtualization_Hardware.md
- Hyperstack (neocloud, SR-IOV offering): iPerf inter-VM on GPU VMs reached **350 Gbps at 24 threads** out of a 400 Gbps theoretical max, vs 8.2 Gbps with legacy virtio-net on the same setup [vendor-reported].
  Source: https://www.hyperstack.cloud/blog/case-study/leveraging-high-speed-networking-for-high-performance-cloud-applications

### 2.2 PF/VF model and the embedded switch
- The PF owns configuration (VF creation, MAC/VLAN assignment, QoS); VFs get a deliberately limited feature set — advanced management stays in the PF/host driver [secondary].
  Source: https://www.quora.com/What-is-SR-IOV-passthrough
- Traffic between VFs on the same NIC is switched by the NIC's **embedded switch (eSwitch)** — hardware VEB/VEPA — without touching the host CPU or the external switch [vendor-reported].
  Source: https://github.com/dpdk/dpdk/blob/HEAD/doc/guides/nics/mlx5.rst
- DPDK's mlx5 PMD supports the full hierarchy: embedded switch, PFs, SR-IOV VFs, Linux auxiliary **Sub-Functions (SFs)**, and their port representors — SFs being the newer, more flexible alternative to VFs on mlx5 [official].
  Source: https://github.com/dpdk/dpdk/blob/HEAD/doc/guides/nics/mlx5.rst
- NVIDIA device-ID trail for VFs (Mellanox PCI vendor 15b3): ConnectX-6 VF = 101c, ConnectX-5 Ex VF = 101a, ConnectX-5 VF = 1018, ConnectX-4 Lx VF = 1016, ConnectX-4 VF = 1014; ConnectX-8 = 1023, ConnectX-9 = 1025, BlueField-3 = a2dc [secondary].
  Source: https://github.com/nvidia/deepops/blob/HEAD/docs/k8s-cluster/roce_backend.md
  Source: https://github.com/nvidia/k8s-launch-kit/blob/HEAD/skills/k8s-network-engineer/references/glossary.md

### 2.3 switchdev: getting the kernel back in control
- In **switchdev (SR-IOV hardware-offload) mode**, each VF gets a **port representor** netdev on the host; the host can program the eSwitch (via tc/flower offload) so traffic policy, switching, and visibility live in the kernel/OVS layer while the fast path stays in hardware [secondary].
  Source: https://docs.redhat.com/en/documentation/red_hat_openstack_platform/16.1/html/network_functions_virtualization_planning_and_configuration_guide/part-sriov-nfv-configuration
- OpenStack Neutron expresses this as vnic-type `direct` (plain VF passthrough) vs `direct` + `binding-profile '{"capabilities": ["switchdev"]}'` (hardware offload), plus `direct-physical` for whole-PF assignment to one instance [official].
  Source: https://docs.redhat.com/en/documentation/red_hat_openstack_platform/16.1/html/network_functions_virtualization_planning_and_configuration_guide/part-sriov-nfv-configuration
- Kubernetes SR-IOV device plugin + Multus patterns assign VFs to pods (DPUs handle north-south, NIC VFs east-west in NVIDIA's k8s-launch-kit reference) [secondary].
  Source: https://github.com/nvidia/k8s-launch-kit/blob/HEAD/skills/k8s-network-engineer/references/glossary.md
- VF LAG (HW bonding of VFs across physical ports) is supported from ConnectX-4 Lx upward in SwitchDev mode, but capped (e.g. max 32 VFs with LAG, 20 with full QoS on ConnectX-4 Lx firmware v14.27.1016) [vendor-reported].
  Source: https://forums.developer.nvidia.com/t/does-connectx4-lx-support-hw-lag-sr-iov-vf-lag/206427

### 2.4 The live-migration problem
- Fundamental tension: a VM bound directly to a VF cannot be live-migrated — the destination host has different hardware state, and device state is not part of the migration stream [official].
  Source: https://github.com/dpdk/dpdk/blob/HEAD/doc/guides/howto/lm_bond_virtio_sriov.rst
- DPDK's documented workaround: **bonding PMD** — a virtio + VF bond inside the VM with the VF as primary; before migration, fail traffic over to the virtio member, migrate, then re-attach the VF on the destination [official].
  Source: https://github.com/dpdk/dpdk/blob/HEAD/doc/guides/howto/lm_bond_virtio_sriov.rst
- Red Hat OpenStack docs state flatly: "VM migration with SR-IOV attached instances is not supported"; security groups also unavailable on SR-IOV ports [official].
  Source: https://docs.redhat.com/en/documentation/red_hat_enterprise_linux_openstack_platform/7/html/networking_guide/sr-iov-support-for-virtual-networking
- Research direction: "SRVM" hypervisor support for passthrough-device live migration claimed 9.6× throughput and 98% lower latency vs switching to paravirtual during migration, with no added downtime — academic prototype, not productized [secondary].
  Source: https://www.researchgate.net/publication/311491063_SRVM_Hypervisor_Support_for_Live_Migration_with_Passthrough_SR-IOV_Network_Devices
- Modern mitigation in the wild: **vDPA** (virtio data path acceleration) and virtio live-migration-friendly devices are the standards-track answer; AWS sidesteps the issue operationally via instance-retirement notices rather than live migration for SR-IOV (ENA) instances [secondary].
- Operational rule of thumb: choose SR-IOV/DPDK for throughput and determinism (NFV, trading, AI storage), choose virtio/vhost when live migration, security groups, and orchestration flexibility matter [secondary].

### 2.5 Prerequisites and caveats
- SR-IOV requires: SR-IOV-capable NIC firmware, BIOS/UEFI with SR-IOV enabled, IOMMU (VT-d/AMD-Vi) enabled, hypervisor support, and guest VF drivers [vendor-reported].
  Source: https://docs.nvidia.com/networking/display/mlnxenv24040660/single+root+io+virtualization+(sr-iov)
- Without an IOMMU, a VF's DMA is untranslated — a malicious/buggy guest could DMA into another VM's memory; the IOMMU is a precondition, not an option [secondary].
  Source: https://github.com/denisergocmen924/cloud_fundamentals/blob/HEAD/Cloud%20Learning%20With%20Workbooks/Cloud_hardware.roadmap/English_Hardware/Phase_6_Virtualization_Hardware.md
- If the host IOMMU lacks interrupt remapping, KVM needs `allow_unsafe_assigned_interrupts=1` — a security downgrade (see Wave 3) [official].
  Source: https://docs.redhat.com/en/documentation/red_hat_enterprise_linux_openstack_platform/7/html/networking_guide/sr-iov-support-for-virtual-networking
- [unverified] Exact max-VF counts for ConnectX-7/8/9 per port are not published in the sources collected; 127/port (ConnectX family doc) vs 1K/port (ConnectX-6 Dx brief) conflict is generational, not contradictory.

---

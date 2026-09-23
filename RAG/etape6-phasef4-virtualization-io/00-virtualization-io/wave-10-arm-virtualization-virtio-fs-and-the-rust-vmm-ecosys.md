---
id: etape6-phasef4-virtualization-io/00-virtualization-io/wave-10-arm-virtualization-virtio-fs-and-the-rust-vmm-ecosys
title: "Wave 10 — ARM virtualization, virtio-fs, and the rust-vmm ecosystem"
domain: phase-f4-i-o-virtualization-cpu-acceleration-extensions
role: deep-dive
task: reference
actors: ["AWS", "Intel"]
dates: []
keywords: ["aws", "datacenter", "gpu", "graviton", "intel", "memory", "nvidia", "throughput"]
source: docs/RAG/etape6_phaseF4_virtualization_io.md
source_anchor: ""
source_lines: [548, 621]
section: "Phase F4 — I/O Virtualization & CPU Acceleration Extensions"
sha256: fa6e6a266b58f45fdd9d9c9f6d8325484430b020e88836a6dce2e64edbd121fe
---

# Wave 10 — ARM virtualization, virtio-fs, and the rust-vmm ecosystem

## Wave 10 — ARM virtualization, virtio-fs, and the rust-vmm ecosystem

### 10.1 KVM on ARM: GIC, ITS, SMMUv3
- QEMU's `virt` machine for aarch64 guests: GIC versions selectable (`2` = max 8 vCPUs; `3` = up to 512; `4` needs nested `virtualization=on`, up to 317; `host`/`max` track the host) — MSI/MSI-X delivered via GICv2m or the **GICv3 ITS**, which is the default on modern machine types [secondary].
  Source: https://github.com/utmapp/qemu/blob/HEAD/docs/system/arm/virt.rst
- Guest IOMMU on ARM: `-machine virt,iommu=smmuv3` instantiates a machine-wide **SMMUv3** for the guest — the ARM counterpart to a virtual VT-d, used for nested DMA isolation and assigned-device experiments [secondary].
  Source: https://github.com/tuke-code/qemu/blob/HEAD/docs/system/arm/virt.rst
- The `virt` board exposes PCI/PCIe, hotpluggable DIMMs/NVDIMMs, 32 virtio-mmio transports, and large RAM/vCPU counts — the standard vehicle for ARM cloud VMs (AWS Graviton, Ampere) running virtio-net/blk [secondary].
  Source: https://github.com/vincent290587/qemu-nrf52/blob/HEAD/vincent290587-qemu-nrf52-f4c8d56/docs/system/arm/virt.rst
- Relevance to I/O virtualization: ARM hosts (BlueField DPU Arm cores, Graviton) run DPDK/SPDK natively (DPDK 25.11 ships aarch64 packages), and vDPA/SR-IOV on ARM NICs follow the same PF/VF model — the ecosystem is ISA-neutral at the virtio layer [secondary].
  Source: https://rpmfind.net/linux/rpm2html/search.php?query=dpdk

### 10.2 virtio-fs and DAX
- **virtio-fs** (device ID 26) shares host directories with guests via a FUSE-protocol virtqueue; cloud-hypervisor wires it as `--fs tag=<tag>,socket=<virtiofsd-sock>,num_queues=<n>,queue_size=<n>` and requires `--memory shared=on`; guest mounts with `mount -t virtiofs <tag> <dir>`; needs guest kernel ≥ 5.10 [secondary].
  Source: https://github.com/cloud-hypervisor/cloud-hypervisor/blob/HEAD/docs/fs.md
- **DAX (direct access)** lets the guest mmap host page cache directly for near-native file performance; cloud-hypervisor notes DAX support in the daemon as not yet stable — a known limitation when sizing virtio-fs for build caches or model weights [secondary].
  Source: https://github.com/cloud-hypervisor/cloud-hypervisor/blob/HEAD/docs/fs.md
- Ecosystem: `virtiofsd` (Rust, rust-vmm-based though developed outside the umbrella) and the **vhost-device** daemon collection (GPIO, I2C, RNG, SCMI, SCSI, vsock, sound, video) give hypervisor-agnostic virtio backends usable from QEMU, cloud-hypervisor, or any vhost-user frontend [secondary].
  Source: https://github.com/linaro/linaro-astro/blob/HEAD/src/content/blogs/rust-device-backends-for-every-hypervisor.mdx

### 10.3 rust-vmm / Firecracker / microVMs
- rust-vmm is the shared Rust component set (KVM API wrappers, virtio device models, vmmemory, vm-allocator, VFIO wrappers, linux-loader) behind **Firecracker** (AWS microVMs), **Cloud Hypervisor** (KVM + MSHV), **Kata Containers' Dragonball**, and **libkrun** — FOSDEM 2026 tracked the ecosystem's monorepo evolution [secondary].
  Source: https://thenewstack.io/intel-releases-cloud-hypervisor-based-on-same-components-as-amazons-firecracker/
  Source: https://ecd43db9x666f7364656dx6f7267.gateway.web.tr/https/2026/events/attachments/WEHLEY-rust-vmm_evolution_on_ecosystem_and_monorepo/slides/266719/rust-vmm_q4zaofh.pdf
- Minimal-VMM device model pattern (illustrated by fluxvm): virtio-net over TAP + `/dev/vhost-net` with `KVM_IRQFD`, virtio-blk on raw files, virtio-mmio discovery via cmdline — with the honest performance note that user-mode networking (slirp/passt) never beats TAP+vhost, and the knobs that matter are `vhost=on`, multiqueue, guest TSO/GSO, and TAP busy-poll [secondary].
  Source: https://github.com/zyvorai/fluxvm/blob/HEAD/crates/fluxvm-hypervisor/DESIGN.md

---

## Part H — Sourced performance figures compendium (do not mix eras)
- vhost-user + virtio-pmd (userspace both sides): **2–4×** vs kernel vhost-net/virtio-net (Red Hat; costs poll cores + usability) [vendor-reported].
- OVS-DPDK vs native kernel OVS: **~10×** (Red Hat blog) / **~2.5×** inter-VM iPerf3 (Intel Ubuntu guide) / **~1.45×** (Intel NFV guide) — topology-dependent, single numbers not portable [vendor-reported].
- SR-IOV vs virtio vs emulation (community reference): 95–99% / 70–90% / 20–40% of bare metal [secondary].
- SPDK NVMe/TCP on BlueField-3 (VPP plugin, 2026, experimental): 1.32× geomean throughput vs SPDK/XLIO; 1.96×/2.25× throughput-per-core vs POSIX/io_uring backends [secondary].
- AVX-512 on Emerald Rapids: doubled/tripled throughput in Phoronix tests, no major power/heat regression (vs Skylake-SP-era 33% all-core frequency drop under heavy AVX-512 — historical, not transferable) [independent].
- DPDK 18.02 PVP (historical): 6.12–6.93 Mpps @ 64 B, 2.10–2.64 Mpps @ 1518 B (1 vhost-user + 1 virtio core, old silicon) [independent — dated 2018].

---

## Part F (continued) — additional source URLs
- https://docs.oasis-open.org/virtio/virtio/v1.4/csprd01/virtio-v1.4-csprd01-diff-from-v1.3-wd01.pdf
- https://docs.oasis-open.org/virtio/virtio/v1.4/csprd01/virtio-v1.4-csprd01-diff-from-v1.2-cs01.pdf
- https://github.com/weltling/virtio-villain/commit/aa87bbd443b7a1876ba1699fd20b28dd2db8f108
- https://github.com/mirivlad/tos/blob/HEAD/docs/evidence/STAGE4D1_FIRST_VIRTQUEUE.md
- https://github.com/davidlin2k/rdma-book/blob/HEAD/src/part-5-deployment/ch15-cloud-and-virtualization/virtio-vdpa.md
- https://github.com/k8snetworkplumbingwg/sriov-network-device-plugin/blob/HEAD/docs/vdpa/README.md
- https://www.redhat.com/pt-br/blog/vdpa-kernel-framework-part-3-usage-vms-and-containers
- https://www.redhat.com/it/blog/breaking-cloud-native-network-performance-barriers?source=tag&term=121
- https://github.com/hairongchen/kernel-wiki/blob/HEAD/wiki/concepts/concept-virtio-data-plane.md
- https://www.redhat.com/de/blog/journey-vhost-users-realm
- https://www.intel.com/content/www/us/en/developer/articles/technical/set-up-open-vswitch-with-dpdk-on-ubuntu-server.html
- https://www.intel.com/content/www/us/en/developer/articles/technical/using-open-vswitch-with-dpdk-for-inter-vm-nfv-applications.html?page=1
- https://docs.nvidia.com/datacenter/tesla/pdf/MIG_User_Guide.pdf
- https://github.com/interloperok/ai.infracalculator/blob/HEAD/docs/mig-feasibility.md
- https://github.com/nvidia/deepops/blob/HEAD/workloads/examples/k8s/gpu-usage/README.md
- https://github.com/nvidia/cloud-native-docs/blob/HEAD/openshift/mig-ocp.rst
- https://github.com/utmapp/qemu/blob/HEAD/docs/system/arm/virt.rst
- https://github.com/tuke-code/qemu/blob/HEAD/docs/system/arm/virt.rst
- https://github.com/vincent290587/qemu-nrf52/blob/HEAD/vincent290587-qemu-nrf52-f4c8d56/docs/system/arm/virt.rst
- https://github.com/cloud-hypervisor/cloud-hypervisor/blob/HEAD/docs/fs.md
- https://github.com/linaro/linaro-astro/blob/HEAD/src/content/blogs/rust-device-backends-for-every-hypervisor.mdx
- https://thenewstack.io/intel-releases-cloud-hypervisor-based-on-same-components-as-amazons-firecracker/
- https://ecd43db9x666f7364656dx6f7267.gateway.web.tr/https/2026/events/attachments/WEHLEY-rust-vmm_evolution_on_ecosystem_and_monorepo/slides/266719/rust-vmm_q4zaofh.pdf
- https://github.com/zyvorai/fluxvm/blob/HEAD/crates/fluxvm-hypervisor/DESIGN.md
- https://forums.developer.nvidia.com/t/assigning-vlan-ids-to-virtual-nics-under-sr-iov-operation-not-supported/285457/2
- https://github.com/redhat-partner-solutions/rhel-sriov-test/blob/HEAD/sriov/tests/SR_IOV_Permutation/README.md
- https://github.com/redhat-partner-solutions/rhel-sriov-test/blob/HEAD/sriov/tests/SR_IOV_Permutation_DPDK/README.md
- https://github.com/asterfusion/helium_dpu/blob/HEAD/ET2500/dpdk-24.03/doc/guides/nics/bnxt.rst
- https://kib.kiev.ua/x86docs/Intel/VT-d/D51397-014.pdf
- https://www.slideshare.net/slideshow/vt-d-posted-interruptsfinal/38366857
- http://iommu.com/datasheets/mirrors/kib.kiev.ua-x86docs/Intel/VT-d/D51397-006.pdf

---


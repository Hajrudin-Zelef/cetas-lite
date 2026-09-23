---
id: etape6-phasef4-virtualization-io/00-virtualization-io/wave-3-cpu-virtualization-extensions-vt-x-vt-d-amd-v-amd-vi-
title: "Wave 3 — CPU virtualization extensions: VT-x/VT-d, AMD-V/AMD-Vi, ATS/PRI/PASID"
domain: phase-f4-i-o-virtualization-cpu-acceleration-extensions
role: deep-dive
task: reference
actors: ["AMD", "Intel", "Microsoft"]
dates: []
keywords: ["amd", "gpus", "intel", "latency", "mcp", "memory"]
source: docs/RAG/etape6_phaseF4_virtualization_io.md
source_anchor: ""
source_lines: [133, 185]
section: "Phase F4 — I/O Virtualization & CPU Acceleration Extensions"
sha256: 654b14bb62d6dc8161985189904c9e8bb0297892c14a9f0c930a5078d075511e
---

# Wave 3 — CPU virtualization extensions: VT-x/VT-d, AMD-V/AMD-Vi, ATS/PRI/PASID

## Wave 3 — CPU virtualization extensions: VT-x/VT-d, AMD-V/AMD-Vi, ATS/PRI/PASID

### 3.1 The two halves: CPU virtualization vs directed I/O
- **Intel VT-x / AMD-V (SVM)** virtualize the CPU itself: root/non-root operation, VM exits, and second-level address translation — **EPT** (Intel Extended Page Tables) / **NPT** aka RVI (AMD Nested Page Tables) — so guest-physical → host-physical translation happens in hardware [secondary].
  Source: https://github.com/yusuf-daglioglu/my_dotfiles/blob/HEAD/cross_platform/minimum_hardware_requirements.md
- **Intel VT-d / AMD-Vi** virtualize I/O: the IOMMU translates device DMA addresses and isolates interrupts, so a passed-through device can only touch its own VM's memory [official].
  Source: https://projectacrn.github.io/2.5/developer-guides/hld/hv-vt-d.html
- Linux detection fingerprints: Intel `DMAR-IR: Enabled IRQ remapping in x2apic mode`; AMD `AMD-Vi: Interrupt remapping enabled`; boot params `intel_iommu=on` / `amd_iommu=on` (AMD often needs only `iommu=pt`) [secondary].
  Source: https://github.com/adior-enigma/dusky/blob/HEAD/Documents/pensive/linux/Important%20Notes/KVM/KVM%20Setup/Verify%20VT-x%20and%20Kernel%20Modules%20and%20IOMMU.md
- **Intel VT-c** (connectivity): VMDq (per-VM network queues on the NIC, offloading traffic sorting from the hypervisor) + Virtual Machine Direct Connect (SR-IOV-based port virtualization) — the NIC-side complement to VT-d [secondary].
  Source: https://www.techtarget.com/searchitoperations/tip/Major-I-O-virtualization-vendors-and-technologies-explained

### 3.2 VT-d architecture (Intel VT for Directed I/O, rev up to 5.10)
- Three pillars per the architecture spec: **DMA remapping** (address translation for device DMA), **interrupt remapping** (isolation/routing of device interrupts to the right VM), **interrupt posting** (direct delivery of virtual interrupts to vCPUs, cutting VM exits) [official].
  Source: https://projectacrn.github.io/2.5/developer-guides/hld/hv-vt-d.html
- Latest published VT-d spec: Rev 5.10, Order D51397-018 (cdrdv2-public.intel.com); mirrors document Rev 4.1 (D51397-016) and Rev 3.2 (D51397-012) — capability registers advertise posted-interrupt support, 5-level paging for first-level translation, 1 GB pages, PASID support [official].
  Source: https://cdrdv2-public.intel.com/868911/D51397-018-vt-directed-io-spec.pdf
- Interrupt address range 0xFEEx_xxxx: DWORD writes without PASID are treated as interrupt requests and bypass DMA remapping; with PASID they are translated normally — a subtlety that matters for shared-virtual-memory devices [official].
  Source: http://iommu.com/datasheets/mirrors/kib.kiev.ua-x86docs/Intel/VT-d/D51397-016.pdf
- Scalable mode (modern Intel IOMMUs): two-level PASID tables per device → per-PASID page tables, enabling Shared Virtual Addressing (SVA) and nested translation; page sizes 4 KB / 2 MB / 1 GB [secondary].
  Source: https://github.com/hodgesds/kernel-docs/blob/HEAD/overview/iommu.md
- Linux drivers: `drivers/iommu/intel/iommu.c` (core), `pasid.c`, `svm.c` (SVA); AMD side: `drivers/iommu/amd/` with device table, IVRS ACPI init, v2 page tables, PPR log for I/O page faults [secondary].
  Source: https://github.com/hodgesds/kernel-docs/blob/HEAD/overview/iommu.md

### 3.3 ATS / PRI / PASID — the device-side protocol trio
- **PASID (Process Address Space ID):** tags DMA requests with a process context so one device can serve multiple address spaces (foundation of SVA/SVM) [official].
  Source: https://cdrdv2-public.intel.com/868911/D51397-018-vt-directed-io-spec.pdf
- **ATS (Address Translation Services):** device caches translations from the IOMMU and performs its own address translation (with ATC — Address Translation Cache) to cut IOMMU lookup latency [secondary].
  Source: https://github.com/hodgesds/kernel-docs/blob/HEAD/overview/iommu.md
- **PRI (Page Request Interface):** device can request the IOMMU/OS to page-in memory on I/O page faults instead of failing — enables overcommit and on-demand paging for accelerators/GPUs [secondary].
  Source: https://github.com/hodgesds/kernel-docs/blob/HEAD/overview/iommu.md
- AMD-Vi equivalents: PPR (Peripheral Page Request) log for I/O page faults, Guest Virtual APIC (GA) log for interrupt virtualization, device-table-driven PASID configuration [secondary].
  Source: https://github.com/hodgesds/kernel-docs/blob/HEAD/overview/iommu.md
- [unverified] Adoption depth of ATS/PRI/PASID in shipping NICs/DPUs (beyond GPUs/FPGAs) is not well documented in collected sources — a gap for AI-cluster DMA users.

### 3.4 Interrupt remapping: security rationale and quirks
- Without interrupt remapping, an assigned device can inject arbitrary MSI interrupts (the 0xFEEx_xxxx range) — a classic VM-escape primitive; remapping validates the interrupt source against IRTE entries [official].
  Source: http://iommu.com/datasheets/mirrors/kib.kiev.ua-x86docs/Intel/VT-d/D51397-016.pdf
- If the IOMMU lacks interrupt remapping, KVM needs `allow_unsafe_assigned_interrupts=1` — explicitly flagged as unsafe [official].
  Source: https://docs.redhat.com/en/documentation/red_hat_enterprise_linux_openstack_platform/7/html/networking_guide/sr-iov-support-for-virtual-networking
- Field quirk (2025): an AMD-Vi interrupt-remapping bug report describes virtio-serial devices on secondary-IOAPIC IRQs (24–30) never receiving interrupts with remapping enabled, while IRQ 5 devices work — attributed to IOMMU table setup for the secondary IOAPIC range [secondary].
  Source: https://github.com/josefbacik/kerneldev-mcp/blob/HEAD/docs/implementation/amd-vi-interrupt-remapping-bug.md
- Practical check: IOMMU groups must isolate the device (ACS on PCIe root ports); the `pcie_acs_override` patch fakes isolation and should not be used on hosts handling sensitive data [secondary].
  Source: https://github.com/adior-enigma/dusky/blob/HEAD/Documents/pensive/linux/Important%20Notes/KVM/KVM%20Setup/Verify%20VT-x%20and%20Kernel%20Modules%20and%20IOMMU.md

### 3.5 Virtualization extensions in modern VMMs
- Microsoft's OpenVMM (2026) emulates all three: ARM SMMU, AMD-Vi (`--amd-iommu`), and a full Intel VT-d emulator (PR #3697, rev 4.1 legacy mode — DMAR ACPI tables, second-level page-table walk, interrupt remapping, queued invalidation, 34 unit tests) — mutually exclusive per VM [secondary].
  Source: https://github.com/microsoft/openvmm/pull/3697
  Source: https://github.com/microsoft/openvmm/blob/HEAD/Guide/src/reference/openvmm/management/cli.md
- ACRN (Intel IoT/edge hypervisor) implements DMA remapping with second-level translation for requests-without-PASID; no first-level/nested translation; GVT-g graphics virtualization is incompatible with VT-d on the same unit [official].
  Source: https://projectacrn.github.io/2.5/developer-guides/hld/hv-vt-d.html

---

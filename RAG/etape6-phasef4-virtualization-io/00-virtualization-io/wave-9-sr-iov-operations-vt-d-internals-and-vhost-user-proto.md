---
id: etape6-phasef4-virtualization-io/00-virtualization-io/wave-9-sr-iov-operations-vt-d-internals-and-vhost-user-proto
title: "Wave 9 — SR-IOV operations, VT-d internals, and vhost-user protocol notes"
domain: phase-f4-i-o-virtualization-cpu-acceleration-extensions
role: deep-dive
task: reference
actors: ["AMD", "AWS", "Broadcom", "Intel", "Nvidia"]
dates: []
keywords: ["amd", "aws", "fp8", "gpu", "intel", "latency", "memory", "nvidia", "research"]
source: docs/RAG/etape6_phaseF4_virtualization_io.md
source_anchor: ""
source_lines: [492, 547]
section: "Phase F4 — I/O Virtualization & CPU Acceleration Extensions"
sha256: fd0bb8e921912d0a9fd5b5ce1f422b186b21b4c8287594084ba9792c778d46e1
---

# Wave 9 — SR-IOV operations, VT-d internals, and vhost-user protocol notes

## Wave 9 — SR-IOV operations, VT-d internals, and vhost-user protocol notes

### 9.1 SR-IOV operations playbook (Linux, vendor-neutral + mlx5/Broadcom notes)
- Creating VFs: `echo <n> > /sys/class/net/<PF>/device/sriov_numvfs` (or the PCI-device path `.../0000:82:00.0/sriov_numvfs`); `echo 0` tears them down. Non-persistent across driver reload — production uses udev rules, NetworkManager, or orchestration (OpenStack `sriov_numvfs`, K8s SR-IOV device plugin) [secondary].
  Source: https://github.com/asterfusion/helium_dpu/blob/HEAD/ET2500/dpdk-24.03/doc/guides/nics/bnxt.rst
- Per-VF tuning on the PF: `ip link set <PF> vf <id> mac <addr> vlan <id> [qos <n>] txrate|min_tx_rate|max_tx_rate <Mbps> trust on|off spoofchk on|off state auto|enable|disable` — RHEL partner test matrices permute spoofchk × trust × vlan × QoS × max_tx_rate and assert ping through every combination [secondary].
  Source: https://github.com/redhat-partner-solutions/rhel-sriov-test/blob/HEAD/sriov/tests/SR_IOV_Permutation/README.md
- **Ordering gotcha on mlx5 (ConnectX):** configure VF MAC/VLAN/trust/spoofchk *before* switching the eSwitch to `switchdev` mode — `ip link set ... vf 0 vlan` fails with "Operation not supported" after the switch (Red Hat bug 1856468). Sequence: legacy mode → create VFs → set VF attrs → `devlink dev eswitch set pci/<PCI> mode switchdev` [secondary].
  Source: https://forums.developer.nvidia.com/t/assigning-vlan-ids-to-virtual-nics-under-sr-iov-operation-not-supported/285457/2
- NVIDIA MLNX_EN SR-IOV knobs beyond `ip link`: per-VF sysfs `spoofchk`, per-VF rate limit, and **VF groups** for vSwitch (OVS) — a group's bandwidth is split evenly among member VFs, leftovers go to VFs that haven't hit their individual cap [vendor-reported].
  Source: https://docs.nvidia.com/networking/display/mlnxenv24040660/single+root+io+virtualization+(sr-iov)
- Broadcom bnxt (DPDK PMD): supports PF *and* VF operation under vfio-pci/uio_pci_generic/igb_uio; **flow bifurcation** splits traffic — data plane to DPDK, control plane to the kernel stack — so ethtool and kernel control keep working while DPDK owns the fast path [vendor-reported].
  Source: https://github.com/asterfusion/helium_dpu/blob/HEAD/ET2500/dpdk-24.03/doc/guides/nics/bnxt.rst
- Binding a VF to DPDK: unbind from the kernel driver, set `driver_override` to `vfio-pci`, bind; then run testpmd with `-a <VF-PCI>` (RHEL partner DPDK permutation tests do exactly this in containers) [secondary].
  Source: https://github.com/redhat-partner-solutions/rhel-sriov-test/blob/HEAD/sriov/tests/SR_IOV_Permutation_DPDK/README.md

### 9.2 VT-d internals worth knowing (from the architecture spec)
- Normative doc is the **Intel VT-d Architecture Specification Rev. 5.10 (D51397-018)**; mirrors of older revs (2.2/2.5/3.2/3.4/4.1) circulate, but 5.10 is the reference for current silicon [official].
  Source: https://cdrdv2-public.intel.com/868911/D51397-018-vt-directed-io-spec.pdf
- DMA remapping structure walk (scalable mode): Root Table → Context Table → **PASID Directory → PASID Table** → first-stage/second-stage page tables; the spec's Appendix A pins snoop/memory-type behavior for each structure (e.g. atomic Posted-Interrupt-Descriptor updates are snooped, WB) [official].
  Source: https://cdrdv2-public.intel.com/868911/D51397-018-vt-directed-io-spec.pdf
- **Queued invalidation** is the scalable TLB shootdown path: software posts invalidation descriptors to the Invalidation Queue; `QIES` (Global Status bit 26) reports enable status; IOTLB/Device-TLB invalidations complete asynchronously with `inv_wait_dsc` status writes [official].
  Source: http://iommu.com/datasheets/mirrors/kib.kiev.ua-x86docs/Intel/VT-d/D51397-006.pdf
- **Interrupt remapping** gates MSI/MSI-X through the Interrupt Remapping Table Address register + IRTEs; `IRES` (bit 25) reports enable status; in x2APIC mode compatibility-format interrupts are blocked unless explicitly handled [official].
  Source: http://iommu.com/datasheets/mirrors/kib.kiev.ua-x86docs/Intel/VT-d/D51397-006.pdf
- **Posted-interrupt IRTE format (§9.11):** the IRTE carries the **Posted Descriptor Address** (high/low), a Present bit, Fault-Processing-Disable (FPD), Urgent flag, vector, and SID/SQ/SVT qualifier fields — hardware atomically sets the PIR bit for the guest vector and raises the notification event, with **no VM exit** when the vCPU is running [official].
  Source: https://kib.kiev.ua/x86docs/Intel/VT-d/D51397-014.pdf
- Xen's posted-interrupt implementation (Feng Wu, Intel, XPDDS14): hypervisor updates the IRTE on guest MSI/MSI-X reconfiguration (via `XEN_DOMCTL_bind_pt_irq`), maintains per-vCPU Posted Descriptors with `SN` (suppress-notification), `NV` (notification vector) and `PIR[0-255]`, and migrates interrupt state on vCPU scheduling — "no VMM overhead at all" on the steady-state path [vendor-reported].
  Source: https://www.slideshare.net/slideshow/vt-d-posted-interruptsfinal/38366857
- Buyer implication: posted interrupts + interrupt remapping are table-stakes for dense virtualized networking/storage; confirm both in `dmesg`/capability bits (ECAP) before sizing vCPU-per-host for latency-sensitive VNFs [secondary].

### 9.3 vhost-user protocol essentials (for implementers)
- The vhost-user protocol is a message set letting QEMU (primary) hand virtqueue memory layout and configuration to an external handler (e.g. OVS-DPDK, SPDK vhost target) over a Unix socket; the handler configures the rings and does the actual packet processing [vendor-reported].
  Source: https://www.redhat.com/de/blog/journey-vhost-users-realm
- DPDK ships two complementary pieces: the **vhost-user library** (userspace vhost protocol implementation) and the **virtio-pmd** (DPDK PMD implementing the virtio spec, so a DPDK app can *be* a virtio device endpoint) [vendor-reported].
  Source: https://www.redhat.com/de/blog/journey-vhost-users-realm
- In the kernel's data-plane optimization taxonomy, the trajectory runs: virtio-net → vhost-net (kernel) → vhost-user/DPDK (userspace poll) → **vDPA** (hardware implements the virtqueue; hypervisor only does control-plane setup) — zero software involvement on the steady-state I/O path [secondary].
  Source: https://github.com/hairongchen/kernel-wiki/blob/HEAD/wiki/concepts/concept-virtio-data-plane.md

---

## Part G — Research log
- Wave 1 (virtio): OASIS specs 1.2/1.3/1.4 state, vhost-net vs vhost-user, AWS ENA, GCP gVNIC, crosvm.
- Wave 2 (SR-IOV): PF/VF, eSwitch/switchdev/representors, ConnectX + BlueField, OpenStack/K8s, migration limits.
- Wave 3 (CPU/IOMMU): VT-x/AMD-V, EPT/NPT, VT-d/AMD-Vi, DMA/interrupt remapping, posted interrupts, PASID/ATS/PRI, ACRN/OpenVMM.
- Wave 4 (SIMD): AVX-512 generations, AVX10.2/APX, AMX INT8→FP8, Emerald→Diamond Rapids, EPYC notes.
- Wave 5 (DPDK/SPDK): DPDK 25.11 LTS + 25.07, SPDK v26.01 LTS + v26.05, perf reports, CVE notes.
- Wave 6 (GPU virt): NVIDIA vGPU 19.x licensing, MIG, AMD MxGPU, Intel GVT-g EOL → SR-IOV, AI relevance.
- Wave 7 (vDPA/virtio-1.4): vDPA framework + K8s/CNF story, admin virtqueues, cs01 status correction.
- Wave 8 (OVS-DPDK/MIG): OVS-DPDK tuning and ops, MIG profile tables A100→B200, K8s exposure.
- Wave 9 (ops/internals): SR-IOV sysfs+ip-link playbook, switchdev ordering gotcha, bnxt flow bifurcation, VT-d Rev 5.10 internals, posted-interrupt IRTE, vhost-user protocol pieces.

*End of Phase F4 research file. No further waves planned for this phase.*

---


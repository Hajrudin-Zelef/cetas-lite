---
id: etape6-phasef4-virtualization-io/00-virtualization-io/wave-7-vdpa-and-the-virtio-1-4-admin-queue-track-deep-dive
title: "Wave 7 — vDPA and the virtio 1.4 admin-queue track (deep dive)"
domain: phase-f4-i-o-virtualization-cpu-acceleration-extensions
role: deep-dive
task: reference
actors: ["Intel", "Nvidia"]
dates: ["2026-04", "2026-07"]
keywords: ["benchmark", "intel", "latency", "nvidia"]
source: docs/RAG/etape6_phaseF4_virtualization_io.md
source_anchor: ""
source_lines: [429, 475]
section: "Phase F4 — I/O Virtualization & CPU Acceleration Extensions"
sha256: a653774e4bb1e775d57e86652f640c44550842b9c87f6b44c41c7b556317d767
---

# Wave 7 — vDPA and the virtio 1.4 admin-queue track (deep dive)

## Wave 7 — vDPA and the virtio 1.4 admin-queue track (deep dive)

### 7.1 vDPA: hardware virtqueues, vendor control plane
- **vDPA (virtio Data Path Acceleration)**: NIC hardware implements the virtio data path (virtqueues, descriptor rings, notifications) in its DMA engine; the guest runs an **unmodified standard virtio driver** while the data path bypasses the hypervisor entirely [secondary].
  Source: https://github.com/davidlin2k/rdma-book/blob/HEAD/src/part-5-deployment/ch15-cloud-and-virtualization/virtio-vdpa.md
- Kernel framework (`drivers/vdpa/`): `vdpa_bus` abstraction with two main drivers — **virtio_vdpa** (exposes device as a virtio-net netdev) and **vhost_vdpa** (exposes as a vhost-vdpa device using a vhost-net protocol extension so userspace apps can access rings directly) [secondary].
  Source: https://github.com/k8snetworkplumbingwg/sriov-network-device-plugin/blob/HEAD/docs/vdpa/README.md
- The full notification chain (Red Hat): guest MMIO kick → KVM ioeventfd → vhost-vdpa driver → vDPA bus op rings the real hardware doorbell; RX DMA completion → vDPA VF driver → virtqueue callback → irqfd → KVM injects MSI-X into the guest [vendor-reported].
  Source: https://www.redhat.com/pt-br/blog/vdpa-kernel-framework-part-3-usage-vms-and-containers
- Why it matters vs SR-IOV/VFIO: **live migration works without per-vendor VFIO migration drivers** — because DMA goes through virtqueues, the standard virtio migration story applies; VFIO passthrough needs each vendor to implement the VFIO migration framework [secondary].
  Source: https://github.com/cocoonstack/cloud-hypervisor/blob/HEAD/docs/vdpa.md
- Cloud-hypervisor exposes vDPA via `--vdpa path=<dev>,num_queues=<n>,iommu=on|off` — a concrete hypervisor integration [secondary].
  Source: https://github.com/cloud-hypervisor/cloud-hypervisor/blob/main/docs/vdpa.md
- Kubernetes story: vDPA CNI + vDPA device plugin give pods a virtio mdev interface bound to a VF, with a **single vendor-neutral virtio-net DPDK PMD** on the CNF side — the vendor's control plane stays hidden behind the vDPA kernel framework [vendor-reported].
  Source: https://www.redhat.com/it/blog/breaking-cloud-native-network-performance-barriers?source=tag&term=121
- [unverified] vDPA hardware support remains limited to newer NIC generations (the NIC must implement the virtio descriptor layout in its DMA engine); adoption breadth in 2026 is vendor-dependent.

### 7.2 virtio 1.4: the admin-virtqueue change (cs01, 8 April 2026)
- The standout 1.4 feature is **admin virtqueues** (feature bit `VIRTIO_F_ADMIN_VQ`): a device-agnostic command interface, "the virtqueue analog to a transport" — existing control virtqueues are device-type-specific (net, scsi) and hard to extend; admin queues carry cross-device admin commands instead, with support for **multiple** admin queues for QoS/scalability [official].
  Source: https://docs.oasis-open.org/virtio/virtio/v1.4/virtio-v1.4.pdf
- PCI transport changes: the common configuration gains `admin_queue_index` / `admin_queue_num` registers (placed after `queue_reset`); when `VIRTIO_F_ADMIN_VQ` is negotiated, `num_queues` **excludes** the admin queues [official] [independent — test-suite confirms layout].
  Source: https://docs.oasis-open.org/virtio/virtio/v1.4/virtio-v1.4.pdf
  Source: https://github.com/weltling/virtio-villain/commit/aa87bbd443b7a1876ba1699fd20b28dd2db8f108
- Spec hygiene in 1.4: common config also gained `queue_notif_config_data` (0x38) and `queue_reset` (0x3A), each live only under negotiated features [secondary].
  Source: https://github.com/mirivlad/tos/blob/HEAD/docs/evidence/STAGE4D1_FIRST_VIRTQUEUE.md
- Version-status correction: OASIS published **virtio-v1.4-csprd01 (9 Dec 2025)** and **virtio-v1.4-cs01 (8 April 2026)** — so a formal 1.4 Committee Specification *does* exist as of April 2026; earlier Wave 1 wording that "no formal publication found" referred to the missing 1.3, and the 1.4 cs01 PDF is the current normative draft [official].
  Source: https://docs.oasis-open.org/virtio/virtio/v1.4/virtio-v1.4.pdf
  Source: https://docs.oasis-open.org/virtio/virtio/v1.4/csprd01/virtio-v1.4-csprd01-diff-from-v1.2-cs01.pdf
- Open-source conformance tooling: `virtio-villain` test suite (July 2026) validates admin-queue common-config fields per spec §4.1.4.3 — implementers can check `admin_queue_num ≥ 1` and queue-range non-overlap [independent].
  Source: https://github.com/weltling/virtio-villain/commit/aa87bbd443b7a1876ba1699fd20b28dd2db8f108

---
## Wave 8 — DPDK PMD ecosystem / OVS-DPDK and NVIDIA MIG profiles

### 8.1 OVS-DPDK: the vSwitch workhorse
- OVS gained a DPDK datapath in OVS 2.2 and a DPDK-backed `vhost-user` virtual interface in OVS 2.4 — the DPDK datapath gives lower latency and higher performance than the kernel datapath while `vhost-user` ports attach guests to it [secondary].
  Source: https://github.com/sapcc/neutron/blob/HEAD/doc/source/admin/config-ovs-dpdk.rst
- Red Hat cites roughly **~10×** performance of OVS-DPDK over native kernel OVS; Intel's Ubuntu setup guide measured ~2.5× (inter-VM iPerf3) and ~1.45× in another test — the spread shows how topology-dependent these numbers are; do not quote a single figure as a 2026 benchmark [vendor-reported].
  Source: https://www.redhat.com/de/blog/journey-vhost-users-realm
  Source: https://www.intel.com/content/www/us/en/developer/articles/technical/set-up-open-vswitch-with-dpdk-on-ubuntu-server.html
- Multiqueue matters: an old KVM/DPDK benchmark showed 13.02 Mpps single-queue vs 21.13 Mpps dual-queue on 10 GbE (historical, RHEL 7 era) — multiqueue vhost-user (`queues=N` in libvirt) remains a required tuning knob [secondary — dated].
  Source: http://www.linux-kvm.org/images/c/c8/DPDK.pdf
- Operations notes (OpenNebula docs, current): OVS/DPDK version compatibility must match (`ovs-vswitchd --version`), hugepages per NUMA node, `dpdk-socket-mem` aligned to NUMA topology, PCI device bound to the PMD's driver, and PMD threads poll at 100% CPU by design — `pmd-sleep-max` trades wake latency for idle CPU [secondary].
  Source: https://github.com/opennebula/one-docs/blob/HEAD/content/product/cluster_configuration/networking_system/openvswitch_dpdk.md
- libvirt/QEMU wiring: `<interface type='vhostuser'>` with unix socket path + `<model type='virtio'/>`, OVS side `type: dpdkvhostuserclient` with `vhost-server-path` — client/server modes must be complementary [secondary].
  Source: https://github.com/opennebula/one-docs/blob/HEAD/content/product/cluster_configuration/networking_system/openvswitch_dpdk.md


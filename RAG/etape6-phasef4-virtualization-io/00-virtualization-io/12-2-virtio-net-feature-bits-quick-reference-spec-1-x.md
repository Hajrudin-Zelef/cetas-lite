---
id: etape6-phasef4-virtualization-io/00-virtualization-io/12-2-virtio-net-feature-bits-quick-reference-spec-1-x
title: "12.2 virtio-net feature bits quick reference (spec 1.x)"
domain: phase-f4-i-o-virtualization-cpu-acceleration-extensions
role: deep-dive
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws", "hyperscaler", "research"]
source: docs/RAG/etape6_phaseF4_virtualization_io.md
source_anchor: ""
source_lines: [682, 689]
section: "Phase F4 — I/O Virtualization & CPU Acceleration Extensions"
sha256: 6146c7503caf6b0c8c83d66e467f6de766dc241f2124b79f1bff4121b309a111
---

# 12.2 virtio-net feature bits quick reference (spec 1.x)

### 12.2 virtio-net feature bits quick reference (spec 1.x)
- Core negotiated features: `VIRTIO_NET_F_CSUM` (partial checksum offload), `VIRTIO_NET_F_GUEST_CSUM/TSO/ECN/UFO`, `VIRTIO_NET_F_MAC`, `VIRTIO_NET_F_GUEST_ANNOUNCE`, `VIRTIO_NET_F_MQ` (multiqueue, queue pairs = 2N), `VIRTIO_NET_F_CTRL_VQ` (the third control queue for MAC/VLAN/RX-mode programming), `VIRTIO_NET_F_CTRL_RX/VLAN/MAC_ADDR`, `VIRTIO_NET_F_MRG_RXBUF` (mergeable buffers), `VIRTIO_NET_F_STATUS`, `VIRTIO_NET_F_SPEED_DUPLEX`, `VIRTIO_NET_F_RSS`, `VIRTIO_NET_F_HASH_REPORT`, `VIRTIO_NET_F_NOTIF_COAL` [official].
  Source: https://docs.oasis-open.org/virtio/virtio/v1.4/virtio-v1.4.pdf
- vDPA/virtio-pmd implication: a DPDK virtio PMD or hardware vDPA NIC negotiates the same feature bits — feature parity between the paravirtual guest driver and the backend is what keeps live migration safe across hypervisor upgrades [secondary].
- [unverified] Azure MANA specifics were not captured in this research round — treat the hyperscaler-NIC comparison as AWS+GCP-weighted; a MANA pass belongs in a future refresh.

---


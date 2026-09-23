---
id: etape7-phasea-os/00-os/a40-minimal-and-immutable-variants-fedora-coreos-opensuse-mi
title: "A40 — Minimal and immutable variants: Fedora CoreOS, openSUSE MicroOS, bootc"
domain: step-7-phase-a-server-operating-systems-linux-os-layer
role: deep-dive
task: reference
actors: ["Oracle"]
dates: []
keywords: ["distribution", "packaging", "pricing"]
source: docs/RAG/etape7_phaseA_os.md
source_anchor: ""
source_lines: [617, 669]
section: "Step 7 — Phase A: Server Operating Systems (Linux OS Layer)"
sha256: 1a12d9c9d6aa0c7beedac42983f1873176a25999d338d556ecf16e1ac1999c16
---

# A40 — Minimal and immutable variants: Fedora CoreOS, openSUSE MicroOS, bootc

## A40 — Minimal and immutable variants: Fedora CoreOS, openSUSE MicroOS, bootc

**Why they exist.** Container-optimized hosts don't need a general-purpose mutable OS — they need atomic updates, rollback, and minimal attack surface. Three options bracket this space:

**Fedora CoreOS** `[secondary]`:

- Image-based, immutable host OS; provisioned via **Ignition**; updates via **rpm-ostree** (atomic, with rollback).
- Upstream of **Red Hat CoreOS** (the OpenShift host OS) — the skills transfer directly to OpenShift estates.
- 6-month-ish release cadence tracking Fedora; not an LTS product — hosts are meant to be reprovisioned, not long-lived.

**openSUSE MicroOS** `[secondary]`:

- **Transactional updates** with btrfs snapshots and automatic rollback on failed health checks; supports Ignition or cloud-init depending on image.
- Upstream of **SUSE Linux Enterprise Micro** — the SLES analog of the CoreOS/RHCOS relationship.

**bootc (bootable containers)** `[secondary]`:

- A newer cross-distribution project: the OS itself is distributed and updated as an **OCI container image**, installed to disk with `bootc install`. Works on Fedora/CentOS Stream derivatives today; the model (image-based Linux) is what RHEL's "image mode" productizes.

**Positioning vs this file's distributions.** Immutable hosts are the right answer for Kubernetes node fleets and appliance-style edge nodes (see 7B). For Proxmox hypervisors, general VMs and stateful services (databases, storage), the mutable distributions documented above remain correct — immutability would fight the workload.

---

## A41 — Commercial pricing comparison matrix (date-labeled)

All prices USD, per year, list/typical — **re-verify at purchase; packaging changes frequently**.

| Product | Metric | Typical price | Provenance |
|---|---|---|---|
| Ubuntu Pro (server) | per server | ~$500 | `[vendor-reported]` |
| Ubuntu Pro (workstation) | per workstation | ~$25 | `[vendor-reported]` |
| Ubuntu Pro (personal) | up to 5 machines | $0 | `[vendor-reported]` |
| Ubuntu Pro + Support | per node | Pro + support uplift (quote) | `[vendor-reported]` |
| Ubuntu Legacy add-on | per node | quote (extends to 12/15-yr window) | `[vendor-reported]` |
| RHEL Server Standard | per system (socket-pair class) | ~$799–$1,299 class | `[secondary]` |
| RHEL Server Premium | per system | ~$1,299 (1-yr quote observed) | `[secondary]` |
| RHEL Developer (individual) | up to 16 nodes | $0 (dev/test, non-production) | `[vendor-reported]` |
| RHEL Business Developers | up to 25 instances/user | $0 (dev/test, non-production) | `[official]` |
| CIQ RLC Pro | per node | $350 self / $600 std / $825 prem | `[vendor-reported]` |
| CIQ RLC Pro Hardened | per node | $775 std / $1,000 prem | `[vendor-reported]` |
| CIQ RLC Pro AI | per node | $775 std / $1,000 prem | `[vendor-reported]` |
| Oracle Linux Basic | per physical CPU pair | $699 | `[official]` |
| Oracle Linux Premier (incl. Ksplice) | per physical CPU pair | $1,399 | `[official]` |
| Oracle Linux Premier Plus | per physical CPU pair | $2,499 | `[official]` |
| TuxCare ELS (CentOS 7) | per system | ~$42.50/yr | `[secondary]` |
| Debian (all) | — | $0 (Freexian ELTS quote-based) | `[official]` |
| Rocky/AlmaLinux (community) | — | $0 | `[official]` |
| Canonical Landscape | per host | **unverified** (~$15/host/mo per third-party aggregator — do not budget on this) | `[unverified]` |

**Reading guide.** Node-based pricing (CIQ, Canonical) is simplest to forecast; socket-pair pricing (Red Hat, Oracle) rewards dense virtualization (fewer sockets, many VMs) and punishes small physical fleets. The free tiers (Debian, community rebuilds, Ubuntu Pro personal, RHEL Developer) legitimately cover homelab/small-estate use at $0.

---


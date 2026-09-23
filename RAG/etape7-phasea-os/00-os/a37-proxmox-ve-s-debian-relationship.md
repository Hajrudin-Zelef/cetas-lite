---
id: etape7-phasea-os/00-os/a37-proxmox-ve-s-debian-relationship
title: "A37 — Proxmox VE's Debian relationship"
domain: step-7-phase-a-server-operating-systems-linux-os-layer
role: deep-dive
task: reference
actors: ["AWS", "Oracle"]
dates: ["2025-08-05", "2026-08"]
keywords: ["aws", "cost", "distribution", "license"]
source: docs/RAG/etape7_phaseA_os.md
source_anchor: ""
source_lines: [572, 616]
section: "Step 7 — Phase A: Server Operating Systems (Linux OS Layer)"
sha256: 1d0d26b37f661ed828978a5c4b4231e85d6e977ac6d9bac0ecf4c1629efdb005
---

# A37 — Proxmox VE's Debian relationship

## A37 — Proxmox VE's Debian relationship

**Proxmox VE is built on Debian.** Proxmox VE 9.0 (released **2025-08-05**) is based on **Debian 13 "Trixie"** with Proxmox's own kernel **6.14.8-2** (built from the Ubuntu 25.04 kernel source — Proxmox has long shipped Ubuntu-based kernels on its Debian base for newer hardware enablement), QEMU 10.0.2, LXC 6.0.4, ZFS 2.3.3 and Ceph Squid 19.2.3 `[secondary]`.

**Proxmox VE 8.4** is based on **Debian 12 "Bookworm"** with kernel 6.8; its support window runs to approximately August 2026 — i.e., effectively ended around this writing, making VE 9 the current target `[secondary]`.

**Practical consequences:**

1. **Proxmox hosts are Debian systems.** Standard Debian administration (APT, systemd, networking) applies; Proxmox adds its own repositories (`pve-enterprise` with subscription, `pve-no-subscription` without) and replaces the stock kernel.
2. **Debian upgrades under Proxmox must follow Proxmox's documented major-upgrade path** (VE 8 → 9 includes the Bookworm → Trixie base upgrade with Proxmox-specific steps) — a bare Debian `full-upgrade` procedure is not sufficient and can break the cluster stack `[secondary]`.
3. **Guest OS choice is independent** — Proxmox runs any of the distributions in this file as KVM guests or LXC containers. Debian/Ubuntu LXC templates are first-class citizens via `pveam`.
4. **Kernel pinning matters:** Proxmox's kernel is separate from Debian's; `apt-mark hold` strategies must account for `proxmox-kernel-*` and `proxmox-default-kernel` metapackages.

---

## A38 — Cloud and virtualization images matrix

| Distribution | AWS | Azure | GCP | Proxmox LXC template | Notes |
|---|---|---|---|---|---|
| Debian 12/13 | Official AMIs (Debian AWS account) | Official publisher images | Official images | Yes (`debian-13-standard`) | cloud.debian.org builds; per-point-release IDs |
| Ubuntu LTS | Official AMIs (Canonical AWS account, incl. Pro/Pro FIPS variants) | Official publisher images | Official images | Yes | **Minimal** images available for containers/VMs |
| RHEL 9/10 | Official AMIs (Red Hat account, PAYG or BYOS) | Official publisher images | Official images | Community templates | PAYG images bundle the subscription hourly |
| Rocky 9/10 | Official AMIs (RESF account) | Marketplace images | Official images | Community templates | Also official **cloud, container and live** image variants |
| AlmaLinux 9/10 | Official AMIs | Marketplace images | Official images | Community templates | Official cloud images per minor |
| Oracle Linux 9/10 | OCI-native; AWS/Azure images published | Marketplace images | Published images | — | OCI platform images documented per release |

**PAYG vs BYOS.** On hyperscalers, RHEL and Ubuntu Pro are available as **pay-as-you-go images** (subscription cost folded into the hourly rate) or **bring-your-own-subscription** images for estates with existing contracts. Debian, Rocky and AlmaLinux cloud images carry no OS license cost — the standard choice for cost-sensitive fleets.

---

## A39 — Server vs desktop relevance

| Distribution | Server fit | Desktop fit | Notes |
|---|---|---|---|
| Debian | Excellent (stable, minimal netinst) | Good (large desktop community, conservative packages) | The "boring in the best way" choice |
| Ubuntu Server | Excellent (largest ecosystem, cloud-init/MAAS) | — (separate Ubuntu Desktop product) | Server and Desktop are distinct install profiles |
| Ubuntu Desktop 26.04/24.04 | Poor (not a server product) | Excellent | Relevant only as admin workstations |
| RHEL | Excellent (certified HW/SW ecosystem) | Workstation SKU exists, niche | Server-first product line |
| Rocky/Alma | Excellent (server-identical to RHEL) | Adequate (same packages as RHEL) | No distinct desktop product; fine as admin desktops |
| Oracle Linux | Excellent for Oracle workloads | Niche | Server-first |

**Admin-workstation note.** For the human side of the estate (admin laptops/desktops), Ubuntu Desktop LTS and Debian are the natural companions to an Ubuntu/Debian server fleet; RHEL Workstation exists for RHEL shops. This does not affect guest/host OS selection.

---


---
id: etape7-phasea-os/00-os/a4-debian-12-bookworm-current-status-and-lts-transition
title: "A4 — Debian 12 \"Bookworm\": current status and LTS transition"
domain: step-7-phase-a-server-operating-systems-linux-os-layer
role: deep-dive
task: reference
actors: ["AWS", "Google"]
dates: ["2023-06-10", "2026-04", "2026-04-23", "2026-06-11", "2026-07-11", "2026-08", "2026-08-31", "2026-09", "2028-06-30", "2031-04", "2031-06-30", "2033-06-30", "2036-04"]
keywords: ["aws", "memory", "research"]
source: docs/RAG/etape7_phaseA_os.md
source_anchor: ""
source_lines: [98, 146]
section: "Step 7 — Phase A: Server Operating Systems (Linux OS Layer)"
sha256: 05366799a4ae86c22879b9fc20de836f7e702469b161815c844a3ea90467bb34
---

# A4 — Debian 12 "Bookworm": current status and LTS transition

## A4 — Debian 12 "Bookworm": current status and LTS transition

**Release and kernel.** Bookworm was released 2023-06-10 with **Linux 6.1 LTS** `[secondary]`. At observation time the latest point release is **12.15, dated 2026-07-11** `[secondary]`.

**Lifecycle transition (unresolved conflict).** Debian's security team handed Bookworm to the Debian LTS team in mid-2026. Two sources disagree on the date: one reports the handoff on **2026-06-11** `[secondary]`, while endoflife.date's machine-readable data lists **2026-07-11** as the end of regular support `[secondary]`. Both are recorded here without adjudication. **LTS support runs until 2028-06-30; Freexian ELTS until 2033-06-30** `[secondary]`.

**What LTS means operationally.** Bookworm under Debian LTS receives security fixes for the main archive from the LTS team (a smaller volunteer + Freexian-sponsored team), with reduced architecture coverage and without point-release ISOs after the transition `[secondary]`. It is fully supported for server use, but new hardware enablement is limited compared to Trixie (kernel 6.1 vs 6.12). For Proxmox VE 8 hosts — which are Debian 12 based — this is the relevant timeline: Proxmox VE 8.x support tracks its own lifecycle (see A37), not Debian's directly.

**Freexian ELTS for Bookworm.** Freexian documents extended support for Debian 12 through 2033-06-30 at `https://www.freexian.com/lts/extended/docs/debian-12-support/` `[secondary]`. This is the commercial bridge for estates that cannot leave Bookworm.

---

## A5 — Debian 11 "Bullseye": LTS just ended — action required

Bullseye's Debian LTS window **ended 2026-08-31** `[secondary]` — roughly three weeks before this writing. Any Bullseye system still in production is now receiving security updates **only** if covered by Freexian ELTS (commercial, until 2031-06-30) `[secondary]`. This is the single most time-sensitive lifecycle fact in the Debian section: unpatched Bullseye hosts/VMs should be migrated to Bookworm or Trixie, or enrolled in ELTS.

**Migration path.** Debian supports in-place major upgrades (`apt full-upgrade` across releases) and documents Bullseye → Bookworm → Trixie as sequential upgrades; skipping a release (Bullseye → Trixie directly) is not supported by the release notes `[secondary]`. Proxmox VE hosts have their own documented upgrade path per major version (see A37).

---

## A6 — Debian cloud images, installer and provisioning

**Official cloud images.** Debian publishes official cloud images (qcow2, vhd, vmdk, raw, and cloud-provider-specific builds) via `cloud.debian.org` for AWS, Azure, Google Cloud, OpenStack and generic/no-cloud environments `[secondary]`. For AWS, Debian maintains official AMIs under the Debian AWS account; Azure and GCE images are published through the marketplaces with Debian as publisher `[secondary]`. Exact AMI IDs rotate per point release and region and are **not** recorded here (they change every point release; look them up in the Debian cloud wiki at research time).

**Installer.** Trixie ships the Debian Installer "Bookworm/Trixie" generation with **non-free firmware included by default** since Bookworm — a major practical change for server installs on hardware needing proprietary NIC/storage firmware `[secondary]`. Preseeding (`preseed.cfg`) and fully automated installs remain supported; cloud-init is available as a package and in cloud images `[secondary]`.

**Provisioning relevance.** For Proxmox estates, the practical Debian path is: Proxmox's own ISO (Debian-based) for hosts, and Debian cloud images or LXC templates for guests. Debian provides official LXC images used by Proxmox's template downloader (`pveam`) `[secondary]`.

---

## A7 — Ubuntu 26.04 LTS "Resolute Raccoon": the new LTS

> **Provenance warning:** Ubuntu 26.04 details below are drawn from secondary press coverage of the April 2026 release. First-party Canonical release notes were not retrievable in this research pass; treat kernel and feature specifics as `[secondary]` until confirmed.

**Release.** Ubuntu 26.04 LTS "Resolute Raccoon" was released **2026-04-23** `[secondary]`, succeeding 24.04 as the current LTS. Standard support runs through **April 2031**; Ubuntu Pro/ESM extends to **April 2036** `[secondary]`. The first point release, **26.04.1, shipped in August 2026**, enabling upgrade prompts from 24.04 LTS `[secondary]`.

**Kernel and desktop.** 26.04 ships **Linux 7.0** as the GA kernel `[secondary]`, GNOME 50 on the desktop, and continues Ubuntu's Wayland-by-default posture `[secondary]`.

**Notable server-relevant changes** `[secondary]`:

- **sudo-rs (memory-safe sudo in Rust) replaces classic sudo** as the default privilege-escalation tool; uutils coreutils continue to be evaluated as GNU replacements.
- **TPM-backed full-disk encryption** available in the installer (building on the 24.04 feature).
- NVMe/TCP host support improvements and updated Ceph client tooling relevant to storage nodes.
- Python 3.14, GCC 15, systemd 258 in the base `[secondary]`.

**Upgrade posture (September 2026).** With 26.04.1 released, `do-release-upgrade` from 24.04 LTS is officially offered `[secondary]`. For production server fleets, the conservative play remains 24.04 until the .1 point release has had wider exposure — that condition is now met, but fleet upgrades should still be staged (see 7C automation guidance).

---


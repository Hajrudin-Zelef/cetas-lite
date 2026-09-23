---
id: etape7-phasea-os/00-os/a12-ubuntu-core-the-immutable-iot-edge-variant
title: "A12 — Ubuntu Core: the immutable IoT/edge variant"
domain: step-7-phase-a-server-operating-systems-linux-os-layer
role: deep-dive
task: reference
actors: ["Oracle"]
dates: ["2019-05-07", "2022-05-18", "2024-05-31", "2024-06", "2024-06-30", "2025-05", "2025-05-13", "2025-05-20", "2025-07", "2025-11", "2026-06", "2027-05-31", "2029-05-31", "2029-06-30", "2030-05-31", "2032-05-31", "2033-05-31", "2035-05-31", "2036-05-31", "2039-05-31"]
keywords: ["cost", "gpu", "hyperscaler", "open source", "packaging", "pricing"]
source: docs/RAG/etape7_phaseA_os.md
source_anchor: ""
source_lines: [196, 255]
section: "Step 7 — Phase A: Server Operating Systems (Linux OS Layer)"
sha256: 0f11997f68fbae91ab83e293b09e1d01e684ccf1c3e9cc44d4b4d71366727be7
---

# A12 — Ubuntu Core: the immutable IoT/edge variant

## A12 — Ubuntu Core: the immutable IoT/edge variant

**What it is.** Ubuntu Core is Canonical's **fully snap-based, immutable** OS: the kernel, base system and applications are all snaps, updated transactionally with automatic rollback on failure `[official]`. Current generation is **Ubuntu Core 24** (released June 2024, based on Ubuntu 24.04) `[secondary]`, with 10-year support matching the LTS `[secondary]`.

**Server relevance.** Ubuntu Core is designed for IoT, edge gateways and appliances — not general servers. It matters to this stack in two places: (1) as the OS inside Canonical's appliance-style products, and (2) as a conceptual reference for immutable infrastructure, alongside Fedora CoreOS and openSUSE MicroOS (A40). For Proxmox VMs and general servers, Ubuntu Server (deb-based) remains the correct choice.

**Ubuntu Core vs Ubuntu Server — do not confuse.** Core has no `apt`, no traditional userspace mutability, and a different security model (snap confinement). Skills and tooling built for Ubuntu Server do not transfer directly.

---

## A13 — RHEL 10: current enterprise major

> **Date caveat:** RHEL 10's GA day is reported as **2025-05-20** in most press and **2025-05-13** in at least one Red Hat regional page `[secondary]`. The month (May 2025) is firm; the day is flagged `[unverified]`.

**Platform.** RHEL 10 ships **Linux 6.12** (kernel-6.12.0 family, heavily backported as per Red Hat practice) `[secondary]`, **DNF 5** as the default package manager frontend, **Wayland-only** GNOME (X11 server removed), **Ptyxis** as the default terminal, **PipeWire** for audio, and the removal of several legacy components: the X.org server (Wayland-only GNOME), `iptables` backend remnants (nftables only), SHA-1 signatures disabled by default in crypto policies, and deprecated drivers `[secondary]`.

**Developer and AI positioning.** Red Hat positions RHEL 10 as the base for its AI strategy (RHEL AI, InstructLab) with tuned math libraries and GPU enablement out of the box `[vendor-reported]`.

**Current minor.** At observation time the current minor is **RHEL 10.2** (June 2026), following 10.1 (November 2025) `[secondary]`. RHEL minors arrive roughly every 6 months (May/November cadence).

**Lifecycle.** Full support until **2030-05-31**, maintenance (Production Phase 3) until **2035-05-31**, optional Extended Life Cycle Support (ELS) add-on to **2039-05-31** `[secondary]`. **Extended Update Support (EUS)** — the ability to stay pinned on a specific minor release with backported security fixes — is available on x86-64 server subscriptions `[vendor-reported]`.

---

## A14 — RHEL 9: the incumbent enterprise major

**Platform.** RHEL 9 (GA 2022-05-18) ships **Linux 5.14** (Red Hat's backport-heavy 5.14.0 tree, not comparable to upstream 5.14) `[secondary]`, DNF 4, Python 3.9, GCC 11, OpenSSL 3.0, and GNOME 40 (Wayland default, X11 still available) `[secondary]`.

**Current minor.** **RHEL 9.8** (June 2026) is current at observation time `[secondary]`. Minor releases continue on the ~6-month cadence.

**Lifecycle.** Full support until **2027-05-31**, maintenance until **2032-05-31**, ELS add-on until **2036-05-31** `[official]`/`[secondary]`. **Action date:** full support ends in under 4 years — fleets standardizing on RHEL 9 should have their RHEL 10 migration/upgrade plan (Leapp, A17) drafted by 2026-2027.

**RHEL 9 EUS.** Extended Update Support lets subscribers pin to selected minors (e.g., 9.4, 9.6) for 24 months of critical-impact security backports `[vendor-reported]` — the standard mechanism for change-controlled estates that cannot track every minor.

---

## A15 — RHEL 8 and 7: legacy tail

**RHEL 8** (GA 2019-05-07, kernel 4.18): **8.10 was the final minor release**; full support ended 2024-05-31; maintenance runs to **2029-05-31**; ELS add-on to **2033-05-31** `[official]`/`[secondary]`. No new 8.x minors will ship — only security/bug backports.

**RHEL 7** (GA 2014, kernel 3.10): maintenance ended 2024-06-30; ELS add-on runs to **2029-05-31** `[secondary]`. TuxCare offers an independent **ELS4** program for RHEL 7/CentOS 7 extending vendor-grade patching to **2029-06-30** `[vendor-reported]` — a bridge for estates physically unable to migrate (see A30).

---

## A16 — RHEL licensing and pricing: the subscription model

**What a subscription buys.** RHEL is open source; the subscription buys **access to certified binaries, security errata, the Customer Portal knowledge base, support SLAs, and legal indemnification** `[official]`. The OS itself (sources) is freely available — the rebuilds (Rocky/Alma/Oracle) exist precisely because of this.

**Subscription metric.** Physical servers: **socket-pair** (one subscription covers 2 sockets); virtualized: **virtual-node pair** or per-VM models depending on SKU `[vendor-reported]`. Standard vs Premium differ principally in **SLA** (business-hours vs 24/7, response-time targets); Premium on x86-64 server includes EUS rights `[vendor-reported]`.

**Price points (date-labeled, verify at purchase):**

- Third-party reporting of Red Hat list pricing cites **RHEL Server Standard (physical or virtual nodes) around $799–$1,299/system/year** depending on SKU and term, with Premium higher `[secondary]`. A Red Hat discussion page quotes Premium (1 year) at **$1,299** `[secondary]`. These figures move with Red Hat's packaging changes — treat as order-of-magnitude, `[secondary]`.
- **Red Hat Developer Subscription for Individuals: no cost, up to 16 nodes**, for personal/development use (not production) `[vendor-reported]`.
- **Red Hat Enterprise Linux for Business Developers: no cost, up to 25 instances per registered user** for development/test (not production), announced July 2025 `[official]`.

**Price comparison warning.** RHEL pricing is SKU-dense (Workstation, Server, Virtual Datacenters, Edge, academic, hyperscaler PAYG images). Any single number quoted without its SKU is meaningless — always compare SKU-to-SKU (A43 pricing matrix lists SKUs where known).

---


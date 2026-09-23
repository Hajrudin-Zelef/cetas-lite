---
id: etape7-phasea-os/00-os/a8-ubuntu-24-04-lts-noble-numbat-the-workhorse
title: "A8 — Ubuntu 24.04 LTS \"Noble Numbat\": the workhorse"
domain: step-7-phase-a-server-operating-systems-linux-os-layer
role: deep-dive
task: reference
actors: []
dates: ["2020-04", "2022-04", "2024-04", "2025-05-31", "2026-09", "2027-04", "2029-05", "2030-04", "2032-04", "2034-04", "2036-04"]
keywords: ["cost", "gpus", "pricing"]
source: docs/RAG/etape7_phaseA_os.md
source_anchor: ""
source_lines: [147, 195]
section: "Step 7 — Phase A: Server Operating Systems (Linux OS Layer)"
sha256: cadad1ac50789b8b566e85499d380a6b0375915962b51ed31cff35e2078b263c
---

# A8 — Ubuntu 24.04 LTS "Noble Numbat": the workhorse

## A8 — Ubuntu 24.04 LTS "Noble Numbat": the workhorse

**Release and kernel.** Ubuntu 24.04 LTS was released April 2024 with **Linux 6.8** as the GA kernel `[secondary]`. The **24.04.5 point release (September 2026)** rolls the HWE stack forward to **Linux 7.0** `[secondary]` — meaning a fresh 24.04.5 server install runs a newer kernel than the original 24.04 GA.

**Base platform.** systemd 255, Netplan 1.0 (now the default network renderer on server, replacing ifupdown), Python 3.12, OpenSSL 3.0, GCC 13 `[secondary]`. Netplan 1.0 is operationally significant: all server network configuration is YAML under `/etc/netplan`, and cloud-init writes netplan configuration on first boot.

**Kernel cadence — GA vs HWE.** Ubuntu LTS releases ship a **GA kernel** (fixed major version for the life of the LTS on the original installer) and a **Hardware Enablement (HWE)** stack that rolls newer kernels in at point releases `.2` through `.5` `[official]`. Servers installed from the original 24.04 ISO stay on 6.8 unless the HWE stack is opted into; servers installed from 24.04.2+ ISOs get the newer HWE kernel by default `[official]`. For Proxmox guests and bare-metal nodes with recent NICs/GPUs, the HWE kernel is usually the right choice; for maximum stability on certified hardware, pin the GA kernel.

**Lifecycle.** Standard (free) security maintenance until **April/May 2029**; Ubuntu Pro/ESM until **April 2034**; paid Legacy add-on until **April 2036** `[official]`/`[vendor-reported]`.

---

## A9 — Older Ubuntu LTS: 22.04 and 20.04 status

**Ubuntu 22.04 LTS "Jammy Jellyfish"** (April 2022, kernel 5.15 GA): standard support until **April 2027**, ESM until **April 2032**, Legacy add-on until **April 2034** `[official]`. Still the most common Ubuntu version in existing fleets per secondary adoption data; 22.04 → 24.04 upgrades are mature.

**Ubuntu 20.04 LTS "Focal Fossa"** (April 2020, kernel 5.4 GA): standard support **ended 2025-05-31** `[official]`. Systems still on 20.04 are now dependent on Ubuntu Pro ESM (until April 2030) or the Legacy add-on (until April 2032) `[vendor-reported]`. Unsubscribed 20.04 systems are unpatched — same urgency class as Debian 11 (A5).

**Interim releases.** Ubuntu's non-LTS releases (24.10, 25.04, 25.10) have 9-month support windows and are irrelevant for server infrastructure; they exist as technology previews for the next LTS `[official]`.

---

## A10 — Ubuntu Pro, ESM and the Legacy add-on: pricing and structure

**What Ubuntu Pro is.** A subscription that extends security maintenance to 10 years (ESM), adds FIPS-validated crypto modules, CIS/USG hardening automation, Livepatch, and Landscape eligibility `[vendor-reported]`. The base OS remains free; Pro pays for the extended security and compliance artifacts.

**Tiers (Canonical's published structure)** `[vendor-reported]`:

- **Ubuntu Pro (personal): free for up to 5 machines** — full ESM access, the reason many homelabs run Pro on personal servers at zero cost.
- **Ubuntu Pro (enterprise): ~$25/workstation/year, ~$500/server/year** — typical list pricing reported in Canonical's own documentation repository; volume and term discounts apply.
- **Ubuntu Pro + Support:** adds 24/7 or 24/5 phone/ticket support with defined SLAs on top of Pro.
- **Legacy add-on:** extends coverage from year 10 to year 12 (15 years total on newer LTS like 20.04/22.04 structure) — quote-based, aimed at estates that cannot migrate.

**ESM coverage detail.** ESM covers the **main** and **universe** repositories (universe coverage was the major Pro improvement over the old ESM model, which covered main only) `[official]`. Without Pro, universe packages on an EOL LTS receive no security updates — a frequently missed exposure on servers running universe software.

**Pricing caveat.** The $25/$500 figures are Canonical-published typical prices; actual enterprise quotes vary by volume, term and region. They are `[vendor-reported]`, not independently audited, and Canonical can change list pricing — re-verify at purchase time.

---

## A11 — Canonical Landscape: fleet management

**What it is.** Landscape is Canonical's systems-management platform: package/update orchestration across fleets, inventory, compliance reporting, repository mirroring, and role-based administration, delivered as SaaS or self-hosted (on-premises) `[secondary]`.

**Licensing position.** Historically Landscape was available to Ubuntu Advantage/Pro support subscribers; current Canonical documentation positions Landscape as included with Ubuntu Pro subscriptions and also available for the free personal tier at small scale `[unverified]` — the exact current bundling could not be confirmed against Canonical's docs in this pass. Third-party software directories quote **~$15/host/month** for Landscape `[secondary]`, but this figure is **not** from Canonical and is marked `[unverified]` for pricing decisions.

**Operational relevance.** Landscape is the canonical (lowercase-c) answer to "how do I patch 200 Ubuntu servers without Ansible" — it complements rather than replaces configuration management (7C): Landscape handles package/update/compliance state; Ansible/Terraform handle configuration and provisioning.

---


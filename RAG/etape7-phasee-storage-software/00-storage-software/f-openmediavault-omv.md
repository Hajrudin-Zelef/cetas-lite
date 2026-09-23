---
id: etape7-phasee-storage-software/00-storage-software/f-openmediavault-omv
title: "F. OpenMediaVault (OMV)"
domain: step-7-phase-e-software-defined-nas-storage-ceph-rclone-true
role: deep-dive
task: reference
actors: ["AMD", "CISA", "Intel", "United States"]
dates: ["2025-04", "2025-04-16", "2025-05", "2025-12", "2026-07"]
keywords: ["advisory", "amd", "cost", "gpu", "intel", "license", "memory", "pricing", "research"]
source: docs/RAG/etape7_phaseE_storage_software.md
source_anchor: ""
source_lines: [189, 268]
section: "Step 7 — Phase E: Software-Defined & NAS Storage (Ceph, rclone, TrueNAS, OpenMediaVault, Unraid, QNAP/Synology/Asustor, ZFS)"
sha256: 38cf2788091754c4dbc9d78a113858f5a61c91b363e7b3f798031f9acd346665
---

# F. OpenMediaVault (OMV)

## F. OpenMediaVault (OMV)

### F1. Project status 2026

- OpenMediaVault is a GPLv3 Debian-based NAS OS with a web UI; **OMV 8 "Synchrony"** was released **24 December 2025**, rebased on **Debian 13 "Trixie"** [secondary].
- Defaults: mdraid + ext4/xfs/btrfs; **ZFS available via plugin** (not native default) [secondary].
- Docker apps via **OMV-Extras** plugin repository; strengths are old PCs, Raspberry Pi and low-power boxes [secondary].
- Community comparison (2026): OMV = "a Debian box with a web UI" — best for users comfortable with Debian who want GUI-assisted storage management [secondary].

### F2. OMV vs TrueNAS vs Unraid (2026 independent framing)

- **TrueNAS**: ZFS integrity/silent-corruption protection; free, no feature gates; larger learning curve [secondary].
- **Unraid**: mismatched-drive expansion; easiest app templates; paid license [secondary].
- **OMV**: free and open-source, smallest hardware floor (~1 GB RAM, 4 GB storage boot); plugin-based ZFS; no single-pool mixed-drive magic [secondary].

---

## G. Unraid

### G1. Product and licensing 2026

- Unraid OS is a paid, Slackware-based NAS OS (USB-flash boot, license tied to the USB drive); **current stable observed: 7.3.2 (July 2026)** with 7.3.3-rc.1 in testing (Sept 2026) [secondary][official].
- **Pricing (2024 pricing change, current through 2026)**: **Starter $49**, **Unleashed $109**, **Lifetime $249**; optional **$36/year** updates subscription — skipping it keeps the purchased version working [secondary].
- Older documentation still shows legacy tiers ($59 Basic/6 devices, $89 Plus/12, $129 Pro/unlimited, lifetime) from before the 2024 change; treat pre-2024 figures as superseded [secondary].
- **30-day free trial**, no credit card required [official].
- Release cadence: 7.0 (Jan 2025), 7.1 (May 2025), 7.2 (Oct 2025), 7.3-rc (Apr 2026) [secondary].

### G2. Architecture and features

- **Array model**: per-disk XFS/btrfs/ZFS filesystems + 1–2 dedicated parity disk(s); unique selling point is **mixed drive sizes** with one-at-a-time expansion and no wasted capacity — nothing else handles a shoebox of mismatched drives as well [secondary].
- **ZFS pools** supported as cache/pool devices (not the default array model); **foreign ZFS pool import** added in 7.1 [secondary].
- Docker management via **Community Apps**: 1000+ one-click templates; best-in-class GPU passthrough to containers (checkboxes in UI); KVM VMs supported [secondary].
- 7.3.x notes mention Intel microcode updates, lshw, kernel-firmware refreshes, Dynamix File Manager and GUI Search plugins, Outgoing Proxy Manager [official].
- Weaknesses vs TrueNAS: array write performance slower than RAID (parity penalty), not enterprise/HA focused, USB boot drive is a single point of failure (mitigated by USB backup tooling) [secondary].
- Typical verdict (2026): media-server shoebox with mixed drives → Unraid pays for itself in drives not replaced; ZFS-integrity-first → TrueNAS [secondary].

---

## H. Appliance vendors: Synology, QNAP, Asustor

### H1. Synology — DSM and 2026 lineup

- Synology DiskStation Manager (**DSM**) is the appliance OS; the 2026 desktop refresh is the **25 Plus series**: DS425+, DS725+, DS925+, DS1525+, DS1825+ [secondary].
- **DS1825+** (8-bay, successor to DS1821+): AMD Ryzen V1500B quad-core, 8 GB DDR4 ECC (expandable to 32 GB), 2× 2.5GbE, 2× M.2 NVMe (cache or pool), PCIe 3.0 x8 slot for 10/25 GbE NICs, expandable to **18 bays / 360 TB** via 2× DX525 expansion units (USB-C), SHR/JBOD/RAID 0/1/5/6/10, 3-year warranty (extendable to 5). US MSRP **$1,150** diskless; UK retail observed at **£1,229.99**; Qatar retail observed at **QAR 6,040** (Sept 2026) [secondary].
- **DS725+** (2-bay): US MSRP **$520** diskless; with DX525 expansion up to 140 TB raw [secondary].
- **neo+ series** (launched ~Aug 2026): DS725neo+, DS925neo+, DS1525neo+, DS1825neo+ — same platform as Plus series, 4 GB non-ECC DDR4 default (upgradeable to 32 GB ECC), built-in M.2, 2.5GbE, DX525 expansion; positioned as budget-accessible as memory market dynamics shift. India launch pricing observed: DS725neo+ Rs 61,123; DS925neo+ Rs 67,632; DS1525neo+ Rs 98,304; DS1825neo+ Rs 122,687 [secondary].
- Key DSM apps: **Synology Drive** (private cloud, cross-platform sync), **Active Backup Suite** (Windows/Linux/macOS/VM/cloud 3-2-1 backups), **Surveillance Station** (on-prem NVR), **Container Manager** (Docker) [secondary].

### H2. Synology drive-compatibility policy (controversy)

- Effective **April 2025 (announced 2025-04-16)**: on Plus-series models released from 2025 onward, **only Synology-branded or Synology-certified third-party drives offer full functionality and support**; non-certified drives technically work but lose software/firmware-related features [secondary].
- Synology's stated rationale: higher performance, reliability, support efficiency ("with our proprietary hard drive solution, we have already seen significant benefits") [vendor-reported].
- Community impact: Synology HDDs are competitively priced, but **Synology SSDs are markedly more expensive** than competitors' — a real cost penalty for all-flash builds [secondary].
- This is the single most-cited reason in 2026 for homelab users to consider TrueNAS/Unraid/UGREEN instead of new Synology hardware [secondary].

### H3. QNAP — QTS/QuTS hero and security history

- QNAP's NAS OS is **QTS** (and **QuTS hero** with ZFS on higher-end models); 2026 lineups continue the TVS/TS-x64/x73A families — specific new-model pricing was not captured in this research pass [unverified].
- **DeadBolt ransomware (2022)**: internet-exposed QNAP devices running outdated QTS/Photo Station were exploited; files got `.deadbolt` extension; ransom notes on the NAS login page demanded Bitcoin. Singapore Police/CSA issued a joint advisory urging QTS/app updates and disabling port-forwarding [official][secondary].
- **QSnatch** (2020): data-stealing malware infected **62,000+ QNAP devices** per a joint US CISA/UK NCSC advisory; payload included CGI password logger, credential scraper, SSH backdoor, and config/log exfiltration over HTTPS [secondary].
- Mitigation posture repeated by vendors and CERTs: keep QTS and apps patched, disable port forwarding/UPnP, change default ports, disable EZ-Connect-style cloud relay if unused, keep offline/offsite backups [secondary].

### H4. Asustor — ADM and DeadBolt

- Asustor (ASUS subsidiary) NAS run **ADM (Asustor Data Master)**; affected models in the 2022 DeadBolt wave included AS5104T, AS5304T, AS6404T, AS7004T, AS5202T, AS6302T, AS1104T per the NZ CERT advisory [secondary].
- Attack vector was believed to be the **EZ Connect** remote-access service (and possibly the Plex app); Asustor **disabled the myasustor.com DDNS service** during investigation and issued recovery firmware — encrypted data was unrecoverable without backups [secondary].
- Ransom was 0.03 BTC (~$1,100–1,200 at the time) per victim with a unique Bitcoin address; unlike the QNAP wave, no vendor-level master-key offer was reported for Asustor [secondary].
- Asustor's hardening guidance: change default ports (8000/8001, 80/443), disable EZ Connect, close Plex ports/disable Plex, immediate backup, disable Terminal/SSH and SFTP [secondary].

### H5. Appliance pricing signals (Sept 2026)

- Synology DS1825+: $1,150 (US MSRP, diskless) / £1,229.99 (UK) / QAR 6,040 (Qatar) [secondary].
- Synology DS725+: $520 (US MSRP, diskless) [secondary].
- Synology neo+ series: Rs 61,123–122,687 (India launch) [secondary].
- UniFi UNAS Pro (rackmount 7-bay, no license): **$499** [secondary].
- UGREEN NASync (desktop, Docker-capable): **from $374**, no license [secondary].
- QNAP/Asustor 2026 street pricing was not systematically captured; treat as a gap.

---


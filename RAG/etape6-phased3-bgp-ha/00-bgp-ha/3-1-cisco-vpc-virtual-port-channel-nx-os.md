---
id: etape6-phased3-bgp-ha/00-bgp-ha/3-1-cisco-vpc-virtual-port-channel-nx-os
title: "3.1 Cisco vPC (virtual Port Channel) — NX-OS"
domain: phase-d3-bgp-underlay-and-high-availability-in-the-data-cent
role: deep-dive
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["nvidia"]
source: docs/RAG/etape6_phaseD3_bgp_ha.md
source_anchor: ""
source_lines: [145, 190]
section: "Phase D3 — BGP underlay and high availability in the data center"
sha256: 245ca5ef5cbd622fbbd14bb00e88fe2f8d13c3372c7ed113f4a5b27e1b0e0936
---

# 3.1 Cisco vPC (virtual Port Channel) — NX-OS

### 3.1 Cisco vPC (virtual Port Channel) — NX-OS

- **Architecture:** two Nexus switches form a vPC domain; downstream devices see one logical port-channel via LACP with a shared virtual system ID. A dedicated **vPC peer link** (port-channel) carries control traffic and flooded/unknown traffic; a separate **peer-keepalive** link (routed, recommended in its own VRF — never over the peer link itself) provides liveness detection [official — Cisco Nexus 9000 NX-OS Interfaces Configuration Guides, 6.x through 10.1(x)].
- **Timers:** peer-keepalive interval default **1 second**, configurable **400 ms–10 s**; hold-timeout **3–10 s** (default 3 s) during which the secondary ignores keepalives to let convergence settle; timeout **3–20 s** (default 5 s) after which the secondary takes over if no hello arrived [official — NX-OS 6.x vPC guide].
- **Failure behavior** [official]:
  - **Peer-link failure, peer alive:** secondary detects via keepalive that the peer is still up and **disables all its vPC member ports** to prevent loops; traffic continues through the primary. Orphan devices attached only to the secondary are blackholed in this state.
  - **Keepalive-link failure only:** no traffic impact; no role change; roles unchanged when the link recovers.
  - **Whole peer failure:** surviving peer stops receiving keepalives and takes over; all traffic uses remaining links until the failed switch returns.
- **vPC peer-gateway:** lets the peer link carry HSRP/VRRP-routed traffic so the local switch can route packets destined to the peer's gateway MAC — standard in DC designs with FHRP [official].
- **Auto-recovery:** `auto-recovery` / `auto-recovery reload restore` lets the secondary bring up vPC legs when the primary never comes back (e.g., after reload); **not enabled by default** — must be configured on the secondary [official — NX-OS 9.2(x)/10.1(x) guides].
- **Object tracking:** `track` objects/lists on core-facing links and peer links can force a vPC role switchover on module failure to avoid blackholing when peer link and uplinks share a module [official — NX-OS 10.1(x)].
- **Version notes:** documented consistently across NX-OS 6.x, 7.x, 9.2(x), 10.1(x) for Nexus 9000; vPC is a data-center (NX-OS) feature — not available on IOS-XE Catalyst in the same form [official][secondary].
- Sources: https://www.Cisco.com/c/en/us/td/docs/dcn/nx-os/nexus9000/101x/configuration/interfaces/cisco-nexus-9000-nx-os-interfaces-configuration-guide-101x/b-cisco-nexus-9000-nx-os-interfaces-configuration-guide-93x_chapter_01000.html

### 3.2 Arista MLAG — EOS

- **Architecture:** MLAG domain with `domain-id`; peers communicate over a **peer-link** port-channel and a **local-interface** SVI (conventionally VLAN 4094) carrying the peer IP; both switches present the **same LACP system ID** to downstream devices, so a host's regular LACP bundle spreads across both switches. Each MLAG is bound by an **MLAG ID** (`mlag 3` on the port-channel), which need not match the port-channel number [official/independent — Arista community docs, EOS 4.35/4.36 guides].
- **Heartbeat:** default interval **2 s** (older docs) / **4 s** (EOS 4.36.1F), configurable **1–30 s**; if heartbeats stop for **2.5× the interval** (older) / **30 s** (4.36.1F), the MLAG association dissolves and both switches revert to independent state [official — EOS 4.36.1F MLAG guide; Arista community].
- **Dual-primary detection:** `peer-address heartbeat <ip> [vrf <name>]` sends UDP heartbeats over a separate path (e.g., VRF mgmt) so a peer-link failure doesn't cause dual-primary; **must be configured on both peers** [official — EOS 4.35/4.36]. Community lab testing on vEOS 4.23.0F confirmed the option exists and noted heartbeat-over-peer-link can leave both peers in Primary after reconnect — prefer the separate-path heartbeat [independent].
- **Recovery tuning:** `dual-primary recovery delay mlag <0-86400> non-mlag <0-86400>` staggers port bring-up after a split resolves (suggested: mlag 60, non-mlag 0) [official — EOS 4.36.1F].
- **MLAG ISSU:** in-service software upgrade of one peer with minimal disruption on active MLAG interfaces [secondary].
- **Version notes:** heartbeat and dual-primary-detection knobs documented in EOS 4.23–4.36 trains [official].
- Sources: https://arista.com/en/um-eos/eos-multi-chassis-link-aggregation?searchword=eos%20section%2013%204%20configuring%20mlag, https://arista.my.site.com/AristaCommunity/s/article/mlag-basic-configuration

### 3.3 NVIDIA Cumulus Linux CLAG (MLAG)

- **Architecture:** Cumulus calls it MLAG (docs) / CLAG (CLI heritage). Requires a dedicated bond named **`peerlink`** between peers plus an automatically created **`peerlink.4094`** L3 subinterface; the subinterface's IPv6 link-local provides L3 connectivity between peers [official — Cumulus docs, ESF generic config guide].
- **Peer IP:** `nv set mlag peer-ip linklocal`; do **not** use 169.254.0.1 (reserved for BGP unnumbered) [official].
- **Backup IP:** `nv set mlag backup <peer-address> vrf mgmt` keeps the MLAG domain alive when the peer link goes down — use the loopback or management IP, reachable and different from peerlink.4094 [official].
- **Peer-link sizing:** NVIDIA recommends sizing the peer link at **half the bandwidth of the single-homed interfaces** it must back up, since it serves as the failure backup path [official].
- **Interop note:** Cumulus documents FRR interop quirks for BGP unnumbered (§1.3); the same FRR-based control plane underlies CLAG state sync [official].
- Source: https://github.com/cumulusnetworks/docs/blob/HEAD/content/guides/esf-generic-config-guide.md

### 3.4 Dell VLT (Virtual Link Trunking) — OS10 / OS9

- **Architecture:** VLT domain with a **VLTi** interconnect (port-channel between peers) and a **VLT backup link** (separate path; Dell recommends the OOB management network) for peer liveliness [official — Dell SmartFabric OS10 VLT Reference Architecture Guide; Dell OS9 config guides 9.14.2.x].
- **Primary/secondary election:** default lowest MAC wins; configurable via `role priority` [official — Dell OS9 S3100/S4048T-ON guides].
- **Peer liveliness:** backup-link keepalives default **30 s** interval, configurable **1–30 s**; on VLTi failure the system checks the heartbeat immediately without waiting for the timer, for faster convergence [official — OS10 VLT RA guide].
- **Failure behavior** [official]:
  - **VLTi down, peer alive (heartbeat OK):** secondary shuts down its VLT ports (and associated L3 interfaces on the VLT VLAN); traffic flows via the primary only.
  - **VLTi down AND heartbeat down (full isolation):** secondary assumes primary role; previously learned MAC/ARP entries are retained but new ones are not synchronized; orphan (non-VLT) ports keep working unchanged.
  - **Primary node failure:** both VLTi and heartbeat fail; secondary takes over the primary role. Role changes are reported as SNMP traps. When the original primary returns, the promoted peer **retains** the primary role until reassigned.
- **Graceful LACP:** enabled by default (not disableable); covers node reload and secondary-detects-VLTi-down-with-heartbeat-up scenarios [official — OS10 VLT RA guide].
- **STP:** RSTP/RPVST+ supported with VLT (MSTP not supported on the documented OS10EE train); BPDUs processed on the primary; enable STP before VLT to avoid loops [official].
- Sources: https://infohub.delltechnologies.com/en-us/l/dell-emc-smartfabric-os10-virtual-link-trunking-reference-architecture-guide-1/peer-liveliness-check-2/


---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/tappaas-community-blob-head-src-larsrossen-network-unifi-os-readme-md-45b34268
title: "tappaas-community-blob-head-src-larsrossen-network-unifi-os-readme-md-45b34268"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-unifi-ubiquiti/tappaas-community-blob-head-src-larsrossen-network-unifi-os-readme-md-45b34268.md
source_anchor: ""
source_lines: [1, 38]
sha256: 70c3befaa80303a7c486e4f8d96e2678e93733e2d42949973613c1b7b1a6388a
---

# tappaas-community-blob-head-src-larsrossen-network-unifi-os-readme-md-45b34268

Primary audience: TAPPaaS operator / network administrator.
Self-hosted UniFi OS Server — Ubiquiti's full UniFi OS (UniFi Network and more)
running on your own hardware, no Cloud Key / Dream Machine appliance required. It is the
management plane that adopts and configures your UniFi switches and access points, and
within TAPPaaS it is the controller that ADR-008's switch-manager / ap-manager drive
through the UniFi OS API.
| Capability | Access from | How | 
|---|---|---|
| UniFi OS console + UniFi Network UI | mgmt /home /work zones | https://unifi-os.<domain> (internal) | 
| Direct console | mgmt zone | https://<vm-ip>:11443 | 
| UniFi OS API (for automation / ADR-008) | mgmt /home /work | https://unifi-os.<domain>/proxy/network/api | 
| Device adoption (switches, APs) | LAN | standard UniFi inform/discovery | 
- Not the standalone UniFi Network Controller (that is network/unifi , NixOSservices.unifi , legacy:8443/api ). This module is UniFi OS Server — the full UniFi
OS shell with the/proxy/network API.
- No internet exposure. The friendly name is reachable only from internal zones; it is not published to the internet (no inbound NAT/port-forward).
- No UniFi Protect / cameras validated here (Network-only). Protect works on UniFi OS Server but needs more disk/RAM — size up if you add it.
- No automatic admin account. UniFi OS Server requires a one-time interactive owner setup; see INSTALL.md.
- A cluster:vm -capable node with ~50 GB on the target storage pool and themgmt zone.
- Internet egress from the VM (to download the UniFi OS Server installer and adopt/update).
- UniFi switches/APs on the LAN to adopt (the reason to run it).
- Runs on a Debian 13 (trixie) guest, not NixOS: UniFi OS Server ships only as a
Debian/Ubuntu ELF installer running rootless podman under auosserver user, which does
not fit NixOS. (Debian 12 was tried and kernel-panicked under this Proxmox/QEMU.)
- UI/API listen on :11443 (UniFi OS console), with the Network app on 8443/8444 — there is nothing on :443.
- First-run owner setup is manual (interactive). Self-hosted UniFi OS has no API-key /
Integration feature, so automation authenticates as a local admin (stored once by
setup-credentials.sh ) — neither can be auto-provisioned at install.
network:proxy publishes the friendly name https://unifi-os.<domain> (split-horizon
DNS → Caddy → the VM's :11443, HTTPS upstream), with an access-list permitting only the
mgmt, home, and work zones — not reachable from the internet.
| Depends on | Purpose | 
|---|---|
| cluster:vm | Creates the Debian 13 VM from the cloud image | 
| templates:debian | OS prep (apt update/upgrade + qemu-guest-agent) | 
| backup:vm | Scheduled VM backup to PBS | 
| network:proxy | Internal-only reverse proxy for unifi-os.<domain> | 
4 vCPU / 6 GB RAM / 50 GB disk. Measured Network-only idle: ~1.4 GB RAM used (MongoDB/Java auto-tune to available RAM). Bump RAM to 8 GB for large fleets or UniFi Protect. Disk is thin-provisioned on ZFS (≈5.5 GB actually used); ~30 GB suffices for Network-only.
For installation steps see INSTALL.md.

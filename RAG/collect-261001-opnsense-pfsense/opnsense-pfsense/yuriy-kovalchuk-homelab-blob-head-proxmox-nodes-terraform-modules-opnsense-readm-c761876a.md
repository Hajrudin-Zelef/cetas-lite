---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/yuriy-kovalchuk-homelab-blob-head-proxmox-nodes-terraform-modules-opnsense-readm-c761876a
title: "yuriy-kovalchuk-homelab-blob-head-proxmox-nodes-terraform-modules-opnsense-readm-c761876a"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/yuriy-kovalchuk-homelab-blob-head-proxmox-nodes-terraform-modules-opnsense-readm-c761876a.md
source_anchor: ""
source_lines: [1, 66]
sha256: bd13833cefe31783fd571a562b83ac775c65fe336a1ac826f3cf199f91c5c39f
---

# yuriy-kovalchuk-homelab-blob-head-proxmox-nodes-terraform-modules-opnsense-readm-c761876a

Terraform module that provisions the OPNsense firewall/router VM on the gateway Proxmox node. It also creates the `vmbr1` VLAN-aware LAN bridge and attaches it to a physical NIC.

- Proxmox installed with `vmbr0` connected to the home router
- Devbox shell active (`devbox shell` from repo root)
- OPNsense ISO manually uploaded to Proxmox (see below)

Proxmox cannot decompress `.bz2` ISOs automatically. Download and decompress locally first:

```
curl -L https://mirror.ams1.nl.leaseweb.net/opnsense/releases/26.1/OPNsense-26.1.6-dvd-amd64.iso.bz2 \
  | bunzip2 > OPNsense-26.1.6-dvd-amd64.iso
```
Then upload via Proxmox web UI: **local → ISO Images → Upload**.

Fill in the values in `terraform/prd/.env.example` → `.env` at repo root:

- `GATEWAY_PROXMOX_ENDPOINT` — Proxmox API URL (e.g.`https://192.168.0.254:8006/` )
- `GATEWAY_PROXMOX_PASSWORD` — Proxmox root password
- `GATEWAY_LAN_BRIDGE_PORT` — physical NIC to attach to`vmbr1` (run`ip link show` on Proxmox to find it)

```
cd proxmox-nodes/terraform/prd/gateway
terragrunt plan
terragrunt apply
```
This creates:

- `vmbr1` — VLAN-aware Linux bridge attached to the physical LAN NIC
- OPNsense VM (ID 1001) with WAN on `vmbr0` and LAN on`vmbr1`

After apply, the VM boots from the ISO by default. In Proxmox web UI:

- VM → Options → Boot Order → move `virtio0` (disk) to the top

Open the Proxmox console for the VM. Once the ISO boots and shows a login prompt:

1. Login with `installer` /`opnsense`
2. The installer TUI launches automatically:
  - Select keymap
  - Choose **Install (ZFS)** (recommended)
  - Select target disk → `vtbd0` (the virtio disk)
  - Confirm and wait for installation to complete
3. Set a root password when prompted
4. **Before rebooting** — fix the boot order (step 4 above) so it boots from disk not ISO
5. Reboot

From the OPNsense console menu:

- **Option 1** — Assign interfaces
- WAN → `vtnet0` (connected to`vmbr0` )
- LAN → `vtnet1` (connected to`vmbr1` )

From the OPNsense console:

- **Option 2** — Set interface IP address → WAN → set a static IP in your home network range (e.g.`192.168.0.200/24` , gateway`192.168.0.1` )
- **Option 8** — Shell → run`pfctl -d` to temporarily disable the firewall

Access the GUI at `https://192.168.0.200` from your machine.
Default credentials: `root` / `opnsense`.

`pfctl -d` is reset on reboot — the firewall comes back automatically.


See `NETWORK.md` for the full VLAN design and firewall rules.

VLANs are configured via Ansible — see `ansible/playbooks/opnsense/README.md`.

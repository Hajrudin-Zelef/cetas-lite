---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/stoicturk182-blog-blob-head-content-posts-unifi-ssh-guide-md-e3f88ff1-1
title: "Newer firmware - APs and switches"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: ["2026-03-28"]
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/stoicturk182-blog-blob-head-content-posts-unifi-ssh-guide-md-e3f88ff1.md
source_anchor: ""
source_lines: [1, 83]
sha256: 642788231e2bb9c0ed6f2fc5ebfea9c1f32f25f0201840d8dee16fd4faf6a608
---

# Newer firmware - APs and switches

| title | Unifi SSH Guide | 
|---|---|
| date | 2026-03-28 | 
| lastmod | 2026-03-28 | 
| draft | false | 
| author | Andrew Jones | 
| description | SSH deployment methods involving Ubiquity network devices. | 
| categories |  | 
| tags |  | 
| featuredImage | /images/posts/hiclipart.com.png | 
| toc | true | 
| SSH | 
| Unifi | SSH | 
SSH behaviour across the UniFi product family is not consistent. The UNAS Pro runs a full Debian-based Linux environment. UniFi switches run a proprietary firmware with a limited shell. UniFi gateways (UDM, UDM Pro, UXG) run a locked-down Linux environment with their own persistence model. Each requires a different approach and carries different caveats around authentication, key storage, and what survives a reboot or firmware update.
This guide documents the known working methods for each device type as of UniFi Network 9.x.
UniFi devices fall into two distinct tiers when it comes to SSH configuration ownership. Understanding which tier a device belongs to determines how SSH is configured, where credentials are managed, and what survives a reboot or firmware update.
These devices run full Linux environments and own their SSH stack independently of the controller. SSH can be configured directly on the device via sshd_config, authorized_keys, and standard Linux tooling. The controller can push credentials to these devices, but the device retains its own SSH daemon and can be accessed regardless of controller availability.
| Device | OS / Environment | Default SSH Port | Key Auth Support | Persistent Key Storage | 
|---|---|---|---|---|
| UNAS Pro | Debian Linux | 22 | Yes | Yes (rwfs overlay) | 
| UDM Pro / UDM SE | UniFi OS (Debian-based) | 22 | Yes | Partial (lost on firmware update) | 
| UDR / UXG Pro | UniFi OS (Debian-based) | 22 | Yes | Partial (lost on firmware update) | 
These devices have no meaningful independent SSH configuration. SSH credentials are pushed down from the controller at adoption and updated whenever the controller's Device SSH Authentication settings change. There is nothing to configure on the device directly — SSH access is entirely dependent on what the controller has provisioned.
| Device | OS / Environment | Default SSH Port | Key Auth Support | Persistent Key Storage | 
|---|---|---|---|---|
| UniFi Managed Switch (USW) | Proprietary firmware | 22 | No | No | 
| UniFi Access Point (UAP) | BusyBox / OpenWrt-based | 22 | No | No | 
Note: The USW-Ultra is an exception even within Tier 2 — it has no SSH access at all. See the Unadopted Device section for details.
Unadopted devices use factory default credentials. These are well known and should be treated as temporary — they are overwritten the moment a device is adopted by a controller.
| Device Type | Newer Firmware | Older Firmware | 
|---|---|---|
| APs, Switches, Cameras | ui /ui | ubnt /ubnt | 
| Gateways, Consoles (UDM, UXG, CloudKey) | root /ui | root /ubnt | 
If you are unsure which applies, try the newer credentials first. If those fail, try the older set. If both fail, the device has most likely been adopted by another controller and retains that controller's SSH password — a factory reset is the only recovery path in that scenario.
If the device has not received a DHCP lease (e.g. connected directly to a laptop with no DHCP server), it will fall back to a static IP:
| Device Type | Default Fallback IP | 
|---|---|
| Gateways (UDM, UXG, USG) | 192.168.1.1 | 
| APs, Switches, Cameras | 192.168.1.20 | 
Reference: UniFi Default SSH Credentials - UniHosted
# Newer firmware - APs and switches
ssh ui@<DEVICE-IP>
# Newer firmware - gateways and consoles
ssh root@<DEVICE-IP>
# Older firmware - APs and switches
ssh ubnt@<DEVICE-IP>
# Older firmware - gateways and consoles
ssh root@<DEVICE-IP>
# password: ubnt
If a device is on the network but not appearing in the controller, or is pointing at the wrong controller, SSH in using the default credentials and run the set-inform command to point it at your controller manually:
set-inform http://<CONTROLLER-IP>:8080/inform
For a locally hosted controller at 10.2.1.x for example:
set-inform http://10.2.1.10:8080/inform
The device will then appear as pending adoption in the UniFi Network controller. The controller must be reachable from the device on port 8080 for this to work.
An orphaned device is one that was previously adopted by a controller you no longer have access to. It will not respond to default credentials and cannot be pointed at a new controller via SSH without the original controller password. The only recovery path is a physical factory reset:
- Locate the reset pinhole on the device (underside or rear panel)
- Press and hold with a paperclip for approximately 10 seconds until the LED flashes
- Allow the device to reboot to a steady LED state
- Default credentials are restored and the device is ready for fresh adoption
The USW-Ultra series switches do not support SSH at all — neither before nor after adoption. The traditional set-inform adoption workflow is not available on these devices. If DHCP Option 43 is not configured to push the controller inform URL automatically, alternatives include using the UniFi mobile app on the same network segment, or temporarily spinning up a local controller on a laptop connected to the same LAN to adopt the device and then redirect it to the primary controller.
Reference: USW-Ultra and the adoption challenge - Cody Deluisio
The set-inform command is not limited to local network adoption. Devices at remote sites can be pointed at a controller over the internet using either a static WAN IP or a DDNS hostname. This is the standard method for managing remote site devices from a centralised controller without requiring a VPN.
Before attempting remote adoption the following must be in place:
| Requirement | Detail | 
|---|---|
| Port 8080 open inbound | The controller must accept TCP 8080 from the internet or the remote site's WAN IP | 
| Controller reachable by hostname or IP | Either a static WAN IP or a DDNS hostname resolving to the controller's WAN IP | 
| Device on the remote network | Device must have an active internet connection | 
| SSH access to the remote device | Either physically present or via an existing management path | 
If the controller site has a static WAN IP, point the device directly at it:
set-inform http://<CONTROLLER-WAN-IP>:8080/inform
For sites without a static WAN IP, Cloudflare DDNS provides a stable hostname that updates automatically as the WAN IP changes. This is the recommended approach for homelab and SMB deployments.
Assuming Cloudflare DDNS is configured and a hostname such as <DDNS.DOmain> resolves to the controller's current WAN IP:
set-inform http://<FQDN or IP>:8080/inform
The device resolves the hostname via public DNS, connects to the WAN IP on port 8080, and registers with the controller. It will then appear as pending adoption in the UniFi Network controller.
Cloudflare DDNS works by running a client on the controller host (or gateway) that periodically checks the current WAN IP and updates the Cloudflare DNS A record if it has changed. The controller's inform URL therefore remains stable even when the ISP changes the WAN IP.
The update client can be a lightweight script running on a cron job or a systemd timer, calling the Cloudflare API:
curl -s -X PUT "https://api.cloudflare.com/client/v4/zones/<ZONE-ID>/dns_records/<RECORD-ID>" \
  -H "Authorization: Bearer <API-TOKEN>" \
  -H "Content-Type: application/json" \
  --data "{\"type\":\"A\",\"name\":\"<DDNS-HOSTNAME>\",\"content\":\"$(curl -s https://checkip.amazonaws.com)\",\"ttl\":60,\"proxied\":false}"
Set TTL to 60 seconds on the Cloudflare DNS record so IP changes propagate quickly. Do not proxy the record through Cloudflare (orange cloud) — the device must reach the controller directly on port 8080, and Cloudflare's proxy only passes HTTP/HTTPS traffic on standard ports.
Reference: Cloudflare API - Update DNS Record

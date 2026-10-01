---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/reginleif-lab-docs-project-1-installing-opnsense-on-proxmox-md-at-8743a328fdd2d3964181b5ab
title: "[OPNsenseHQ]"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-opnsense-pfsense/reginleif-lab-docs-project-1-installing-opnsense-on-proxmox-md-at-8743a328fdd2d3964181b5ab2a1ee416dc.md
source_anchor: ""
source_lines: [1, 152]
sha256: 6d711276d0640f9badabe0f8ec007b694caa585a26e3daa0b73f93987cab016d
---

# [OPNsenseHQ]

| title | Project 1: OPNsense Gateway Deployment (HQ) | 
|---|---|
| tags |  | 
| sites |  | 
| status | completed | 

| opnsense | deployment | proxmox | 

| hq | 

Deploy OPNsense (`OPNsenseHQ`) on Proxmox to serve as the secure gateway and firewall for the `reginleif.io` domain (`172.16.0.0/24`).

For educational context about firewalls, NAT, virtualization networking, and DNS resolution, see the dedicated concepts guide.

Note

**Pre-VLAN Addressing:** This project uses flat network addressing (`172.16.0.0/24`). IP addresses and network configuration will change when VLAN segmentation is implemented in Project 11. The gateway will move from `172.16.0.1` to VLAN-specific addresses (e.g., `172.16.5.1` for Infrastructure VLAN).

Configure the virtual machine with these specifications. Settings are optimized for OPNsense's FreeBSD base.

- **OS Type:** Other / Unknown (FreeBSD-based)
- **Machine:** q35 (Native PCIe)
- **BIOS:** OVMF (UEFI)
- **CPU:** Type Host (Passthrough AES-NI) | 2 Cores
- **RAM:** 4096 MB (4 GB)
- **Disk:** 40 GB
- **Controller:** VirtIO SCSI Single (IO Thread enabled)
- **Network:** VirtIO (Paravirtualized) x2
  - *Net0:* Bridge`vmbr0` (WAN)
  - *Net1:* Bridge`vmbr1` (LAN - Isolated)
- **QEMU Guest Agent:** Enabled

**Downloads:** OPNsense: https://opnsense.org/download/

Before booting, ensure the Proxmox host has the correct bridges:

- **vmbr0:** Connected to Physical LAN (WAN for the Lab).
- **vmbr1:** Linux Bridge (No Physical Port) -> Acts as the "Switch" for`172.16.0.0/24` .

**Creating vmbr1 (if not exists):**

1. Navigate to Proxmox Node → Network.
2. Create → Linux Bridge.
3. Name: `vmbr1` , No physical port (isolated bridge).
4. Apply Configuration.

Perform the initial OPNsense installation from the ISO image.

1. Boot `OPNsense-dvd-amd64.iso` .
2. Login as `installer` /`opnsense` .
3. Select **ZFS** (or UFS) and install to the VirtIO disk.
4. Reboot and remove ISO.

Since the web UI isn't reachable yet, configure via the VM Console:

1. **Assign Interfaces:**  - **WAN:**`vtnet0` (Upstream/Home Network)
  - **LAN:**`vtnet1` (Lab Network)
2. **Set Interface IP Address (WAN):**  - Select **WAN** .
  - IPv4 Configuration Type: `Static`
  - IPv4: `192.168.1.240`
  - Mask: `24`
  - Upstream Gateway: `192.168.1.254` (Home Router in our case)
  - IPv6: *None*
3. Select 
4. **Set Interface IP Address (LAN):**  - Select **LAN** .
  - IPv4: `172.16.0.1`
  - Mask: `24`
  - Gateway: *None* (This*is* the gateway)
  - IPv6: *None*
  - DHCP Server: **Do not enable** - this can also be disabled during the first-time wizard in the GUI. (DHCP will be handled by Domain Controllers for AD integration - DNS/DHCP options, dynamic DNS updates, and centralized management).
5. Select 

Once the Web UI is accessible at `https://172.16.0.1` (access it with another VM or tunnel for example):

1. 
**First-Time Wizard Configuration:**During the first-time wizard in the GUI: 
  - On the **General** step, set**Hostname** to`OPNsenseHQ` .
  - On the **Interfaces** step, uncheck**Block private networks** and**Block bogon networks** for the WAN interface (since the lab's WAN connects to a private home LAN).
2. On the 
3. 
**Disable Hardware Offloading** if not by default (Interfaces > Settings):
  - Disable Hardware Checksum Offload
  - Disable Hardware TCP Segmentation Offload
  - Disable Hardware Large Receive Offload
  - **Reboot** the VM after saving.

**Why disable hardware offloading?** VirtIO virtual NICs can cause packet corruption with offloading enabled. Disabling ensures stability with minimal performance impact in virtualized environments.

**Why allow private/bogon networks on WAN?** OPNsense treats RFC1918 addresses (10.x.x.x, 172.16-31.x.x, 192.168.x.x) as potentially spoofed when arriving on WAN. In production, this is a security feature. In a lab where your "internet" is actually a home network, you must disable these blocks to allow upstream connectivity.


OPNsense enables NAT by default, but automatic rules may not generate in some configurations. Verify and configure manually if needed:

1. Navigate to **Firewall → NAT → Outbound** .
2. Check if **Hybrid outbound NAT rule generation** shows rules for`172.16.0.0/24` .
3. If automatic rules exist, you're done. If not, create a manual rule:

**Manual NAT Rule (Required if no automatic rules exist):**

1. Set mode to **Manual outbound NAT rule generation** and click**Save** .
2. Click **Add** (+ button).
3. Configure:
  - **Interface:** WAN
  - **Address Family:** IPv4
  - **Protocol:** any
  - **Source address:**`172.16.0.0/24`
  - **Source port:** any
  - **Destination address:** any
  - **Destination port:** any
  - **Translation / target:** Interface address
  - **Description:**`LAN to WAN NAT`
4. Click **Save** then**Apply Changes** .

**Why might automatic rules not generate?** Common causes include: WAN interface missing a gateway definition, interface configuration order during setup, or OPNsense not recognizing the LAN interface type. Manual rules provide explicit control and are easier to troubleshoot.


OPNsense runs Unbound as its DNS resolver. Configure it with reliable upstream DNS servers and prevent DHCP from overriding these settings.

If WAN uses DHCP, the upstream router may push its own DNS servers, overriding your configuration.

1. Navigate to **Interfaces → WAN** .
2. Scroll to **DHCP client configuration** .
3. Uncheck **Use DNS servers provided by the DHCP server** (or check "Reject" options).
4. Click **Save** and**Apply Changes** .

1. Navigate to **System → Settings → General** .
2. Under **DNS servers** , add:
  - `1.1.1.1` (Cloudflare - primary)
  - `8.8.8.8` (Google - secondary)
3. For each DNS server, set **Use gateway:** to the WAN gateway (ensures DNS queries go out the correct interface).
4. Uncheck **Allow DNS server list to be overridden by DHCP/PPP on WAN** .
5. Click **Save** .

1. Navigate to **Services → Unbound DNS → General** .
2. Ensure these settings:
  - **Enable Unbound:** Checked
  - **Listen Port:** 53
  - **Network Interfaces:** All
  - **DNSSEC:** Checked
3. Click **Save** and**Apply Changes** .

**Why DNSSEC?** DNSSEC validates that DNS responses haven't been tampered with, protecting against cache poisoning and spoofing attacks.


From the OPNsense shell (console or SSH, or use GUI):

```
# [OPNsenseHQ]
ping google.com
ping 1.1.1.1
```
**Why configure DNS explicitly?** In lab environments where OPNsense's WAN connects to a home router via DHCP, the router often pushes its own DNS (e.g., ISP DNS or router IP). These may be slow, unreliable, or block certain queries. Using Cloudflare (1.1.1.1) and Google (8.8.8.8) ensures fast, reliable resolution for your lab VMs.

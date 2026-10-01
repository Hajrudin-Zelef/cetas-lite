---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/ideafieldpro-opnsense-install-config-vpn-blob-head-readme-md-c179bd92-1
title: "ideafieldpro-opnsense-install-config-vpn-blob-head-readme-md-c179bd92"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: []
keywords: ["amd", "intel"]
source: docs/RAG/collect-261001-opnsense-pfsense/ideafieldpro-opnsense-install-config-vpn-blob-head-readme-md-c179bd92.md
source_anchor: ""
source_lines: [1, 191]
sha256: 9c93d663b23a99f3777b0b451a8948a8e0f6f72302ea909e4c41e9fb64daba62
---

# ideafieldpro-opnsense-install-config-vpn-blob-head-readme-md-c179bd92

This case study documents a production-style deployment of the OPNsense firewall on bare-metal hardware, acting as the secure entry point to a private infrastructure. The appliance terminates WAN traffic, provides DHCP services for the LAN, and exposes a WireGuard-based VPN for remote administration of internal servers. The focus is on designing a clean network baseline, making opinionated configuration choices, and validating that the environment is maintainable and observable over time. OPNsense Documentation

- **Scenario** : Deploy a dedicated OPNsense firewall to front a small production network, consolidating perimeter security, DHCP, and VPN access into a single hardened appliance.
- **Constraints** : Bare-metal hardware with dual NICs, low idle power usage, and a simple /24 LAN segment to support servers, management devices, and end-user clients.
- **Approach** : Use OPNsense with ZFS, Intel NICs, Dnsmasq for integrated DHCP, and WireGuard for modern, low-overhead VPN access. Document all steps and decisions as an operator runbook.
- **Outcome** : A reproducible configuration that delivers stable LAN connectivity, controlled static mappings for critical hosts, and secure remote access into the LAN over VPN, backed by regular backups and monitoring.

At a high level, the deployment places OPNsense between the ISP-provided WAN and an internal LAN, with DHCP-managed clients on the LAN and WireGuard VPN clients joining via a dedicated tunnel subnet.

```
flowchart LR
  wan["WAN / Internet"]
  opnsense["OPNsense Firewall"]
  lan["Server/Host & Clients"]
  vpnClients["WireGuard VPN Clients"]
  wan --> opnsense
  vpnClients --- opnsense
  opnsense --> lan
```
    Key design elements:

- **WAN** : DHCP or static from ISP, terminated on OPNsense WAN interface.
- **LAN** :`192.168.1.0/24` with OPNsense at`192.168.1.1` , serving as default gateway and DNS for clients.
- **DHCP** : Dynamic range`192.168.1.100–192.168.1.200` , with static mappings reserved below`.100` for servers and infrastructure.
- **VPN** : WireGuard tunnel network`10.10.10.0/24` , with clients allowed to reach select LAN subnets.

✓ Network infrastructure design and implementation (WAN/LAN segmentation, DHCP ranges, static mappings)

✓ Firewall configuration and security hardening for a production-style perimeter device

✓ DHCP server management and IP allocation using Dnsmasq on OPNsense 25.x

✓ VPN deployment for secure remote access (WireGuard road-warrior setup into LAN resources)

✓ Production environment best practices (backups, updates, monitoring, logging)

✓ Technical documentation and knowledge transfer through a step-by-step operator runbook

✓ Troubleshooting and problem resolution for connectivity, DHCP, and VPN issues

- Install OPNsense on dedicated hardware for optimal network security
- Configure dual-NIC setup for WAN/LAN traffic routing
- Implement DHCP services for automated IP management
- Establish VPN access for secure remote server administration
- Document production-ready firewall deployment

- **CPU** : 1.5 GHz dual-core processor (Intel/AMD x86-64)
- **RAM** : 4 GB minimum (8 GB recommended for production)
- **Storage** : 40 GB SSD (for optimal performance)
- **Network Interfaces** : Minimum 2 NICs (WAN + LAN)
  - Intel NICs strongly recommended for stability
  - Realtek NICs may have driver limitations

For production environments with VPN and DHCP services: reddit

- **CPU** : Intel i5 or equivalent (T-series for lower power consumption)
- **RAM** : 16 GB DDR4
- **Storage** : 120 GB SSD (enterprise-grade)
- **NICs** : Dual Intel Gigabit or 2.5GbE interfaces
- **Power** : ~12-50W idle consumption depending on hardware

Visit official download page: https://opnsense.org/download/

Verify checksum after download

`sha256sum OPNsense-XX.X-OpenSSL-dvd-amd64.iso`
Compare with checksum listed on download page

**Linux/macOS:**

```
dd if=OPNsense-XX.X-OpenSSL-dvd-amd64.iso of=/dev/sdX bs=4M status=progress
sync
```
**Windows:**

- Use Rufus or balenaEtcher
- Select ISO file and target USB drive
- Write in DD Image mode

Document your network topology before installation:

| Interface | Purpose | Network Range | Gateway | 
|---|---|---|---|
| WAN | Internet Connection | DHCP/Static from ISP | ISP Gateway | 
| LAN | Internal Network | 192.168.1.0/24 | 192.168.1.1 | 

This section covers installing OPNsense onto bare-metal hardware using bootable media. The goal is to end up with a resilient base system (preferably using ZFS) that can support firewalling, DHCP, and VPN services without manual intervention after reboots.

1. Insert USB drive into target hardware
2. Access BIOS/UEFI (typically F2, F12, or DEL)
3. Set boot priority to USB device
4. Save and reboot

1. 
**Login Credentials** (at boot screen): docs.opnsense
  - Username: `installer`
  - Password: `opnsense`
2. Username: 
3. 
**Select Installation Options** :
  - Choose **"Install (UFS)"** or**"Install (ZFS)"**
  - ZFS recommended for data integrity and snapshots
4. Choose 
5. 
**Disk Configuration** :
  - Select target installation disk
  - For ZFS: Choose **"stripe"** for single disk setups
6. 
**Set Root Password** :
  - Create strong password for root account
  - **Document this securely** - required for all administrative access
7. 
**Complete Installation** :
  - Wait for installation to complete (~5-10 minutes)
  - Remove installation media when prompted
  - System will reboot

Upon first boot, you'll see the console menu: docs.opnsense

```
0) Logout                    7) Ping host
1) Assign interfaces         8) Shell
2) Set interface(s) IP       9) pfTop
3) Reset root password      10) Filter logs
4) Reset to factory         11) Restart web interface
5) Reboot system            12) Upgrade from console
6) Halt system              13) Restore configuration
```
**If interfaces aren't auto-detected:**

1. Select **Option 1** - Assign interfaces
2. Enter interface names when prompted:
  - **WAN** : Typically`vnet0` ,`em0` ,`igb0` , or`re0`
  - **LAN** : Typically`vnet1` ,`em1` ,`igb1` , or`re1`
3. Confirm assignments with `y`

**Tip**: Identify NICs by MAC addresses or temporarily connect to known networks

Here you bring the newly installed firewall onto the network by assigning IPs and enabling basic services. The objective is to establish a secure management plane on the LAN, with HTTPS access to the web GUI and a small initial DHCP scope for clients.

1. 
Select **Option 2** - Set interface(s) IP address
2. 
Choose **LAN** interface (typically option 2)
3. 
Configure IPv4: IPv4 Address: 192.168.1.1 Subnet Mask: 24 (255.255.255.0)
4. 
Skip IPv6 if not needed
5. 
**Enable DHCP server** :`y`
  - Start address: `192.168.1.100`
  - End address: `192.168.1.200`
6. Start address: 
7. 
**Do NOT revert to HTTP** :`n` (keep HTTPS)

1. Connect client computer to LAN interface
2. Open browser and navigate to: `https://192.168.1.1`
3. Accept self-signed certificate warning
4. **Login credentials** :
  - Username: `root`
  - Password: [your configured password]
5. Username: 

The web interface will launch a configuration wizard:

1. 
**General Information** :
  - Hostname: `firewall`
  - Domain: `yourdomain.local`
  - Primary DNS: `8.8.8.8` or`1.1.1.1`
  - Secondary DNS: `8.8.4.4` or`1.0.0.1`
2. Hostname: 
3. 
**Time Server** :
  - NTP Server: `pool.ntp.org`
  - Timezone: Select your timezone
4. NTP Server: 
5. 
**WAN Configuration** :
  - Type: DHCP (for most ISP connections) or Static IP
  - If static, enter ISP-provided details
6. 
**Set Admin Password** :
  - Change default root password if not done during install
7. 
**Reload Configuration**

This section configures DHCP using the modern Dnsmasq service in OPNsense 25.x, replacing the legacy ISC DHCP daemon. The intent is to centralize IP address management, keep static infrastructure addresses predictable, and integrate DHCP with local DNS so hosts are easily discoverable.


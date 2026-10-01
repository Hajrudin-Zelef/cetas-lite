---
id: collect-261001-fortinet/fortinet/github-isarmadfarooq-fortigate-ngfw-security-lab-a-complete-hands-on-lab-for-deploying-and-1
title: "Required"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["cybersecurity"]
source: docs/RAG/collect-261001-fortinet/github-isarmadfarooq-fortigate-ngfw-security-lab-a-complete-hands-on-lab-for-deploying-and-configuri.md
source_anchor: ""
source_lines: [1, 133]
sha256: f2d5fe2e88686987512aa409a2e00b2ff2fdc34303516cac3e632f50b2b3ceaa
---

# Required

```
███████╗ ██████╗ ██████╗ ████████╗██╗ ██████╗  █████╗ ████████╗███████╗
██╔════╝██╔═══██╗██╔══██╗╚══██╔══╝██║██╔════╝ ██╔══██╗╚══██╔══╝██╔════╝
█████╗  ██║   ██║██████╔╝   ██║   ██║██║  ███╗███████║   ██║   █████╗  
██╔══╝  ██║   ██║██╔══██╗   ██║   ██║██║   ██║██╔══██║   ██║   ██╔══╝  
██║     ╚██████╔╝██║  ██║   ██║   ██║╚██████╔╝██║  ██║   ██║   ███████╗
╚═╝      ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚══════╝
                      SECURITY LAB — ASSIGNMENT 02
```
**A complete, hands-on lab for deploying and configuring a FortiGate Next-Generation Firewall in a virtualized enterprise environment covering VM deployment, interface hardening, multi-layer policy implementation, and live threat validation.**

This repository documents the complete deployment and configuration of a **FortiGate VM Next-Generation Firewall** within a VMware Workstation virtualized environment as part of the **Advanced Network Security** course (MS Cybersecurity, NUCES Islamabad).

| # | Task | Description | Difficulty | 
|---|---|---|---|
| 1 | **VM Deployment** | FortiGate VM provisioning, NIC configuration, CLI setup | 🟡 Intermediate | 
| 2 | **GUI Hardening** | HTTPS management, admin credential security | 🟢 Beginner | 
| 3 | **Policy Implementation** | IPv4 firewall, DoS protection, web filtering | 🔴 Advanced | 
| 4 | **Validation** | Live website blocking and DoS attack testing | 🟡 Intermediate | 

```
FortiGate VM (FortiOS)  ·  VMware Workstation  ·  hping3  ·  Kali Linux
IPv4 Firewall Policies  ·  DoS Anomaly Detection  ·  URL Web Filtering  ·  NAT
```
```
                        ┌─────────────────────────────────────────┐
                        │           LAB TOPOLOGY                  │
                        └─────────────────────────────────────────┘
     ┌──────────┐           ┌─────────────────────────────────────────────┐
     │ INTERNET │ ◄────────►│              FortiGate VM                   │
     └──────────┘  WAN/ISP  │         (FortiOS — VMware Workstation)      │
                   port2    │                                             │
                            │  ┌──────────┐  ┌──────────┐  ┌──────────┐  │
                            │  │  port1   │  │  port2   │  │  port3   │  │
                            │  │  (Mgmt)  │  │  (WAN)   │  │  (LAN)   │  │
                            │  └────┬─────┘  └────┬─────┘  └────┬─────┘  │
                            └───────┼─────────────┼─────────────┼────────┘
                                    │             │             │
                         192.168.139.131     ISP Public    Internal
                         (Host-Only)          IP Addr    192.168.x.x/24
                                    │                          │
                             ┌──────┴──────┐          ┌───────┴────────┐
                             │  Admin/Mgmt  │          │  LAN Clients   │
                             │  Workstation │          │  (DHCP Pool)   │
                             └─────────────┘          └────────────────┘
                                                               │
                                                    ┌──────────┴──────────┐
                                                    │   Kali Linux VM     │
                                                    │  (Attacker / Test)  │
                                                    └─────────────────────┘
```
| VMware NIC | Network Type | Subnet | FortiGate Port | Role | 
|---|---|---|---|---|
| VMNet1 | Host-Only | 192.168.40.0/24 | port2 | WAN Simulation | 
| VMNet2 | Host-Only | 192.168.50.0/24 | port3 | LAN | 
| VMNet3 | Host-Only | Custom | port4 | DMZ (optional) | 
| VMNet8 | NAT | VMware NAT | port1 | **Management** | 
| VMNet11 | Host-Only | Custom | port5 | Additional Zone | 
| VMNet12 | Host-Only | Custom | port6 | Additional Zone | 

```
# Required
VMware Workstation Pro (v16 or later)     # Hypervisor
FortiGate VM image (FortiGate-VM64.ovf)   # Obtain from Fortinet support portal
Kali Linux VM                             # For DoS testing (Task 4)
# Optional but recommended
7-Zip or WinRAR                           # For extracting FortiGate VM archive
Any modern browser (Chrome, Firefox)     # For FortiGate GUI access
```
| Resource | Minimum | Recommended | 
|---|---|---|
| Host RAM | 8 GB | 16 GB | 
| Host CPU | 4 cores | 8 cores | 
| Free Disk | 40 GB | 60 GB | 
| FortiGate VM RAM | 1 GB | **2 GB** | 
| FortiGate VM vCPU | 1 | 1 | 
| FortiGate VM HDD | 30 GB | 30 GB | 

**A Fortinet support portal account is required** to download the FortiGate VM image.

Register at: https://support.fortinet.com

⚠️ 

```
# 1. Clone this repository
git clone https://github.com/yourusername/fortigate-security-lab.git
cd fortigate-security-lab
# 2. Review the lab documentation
cat README.md
# 3. Follow tasks in order:
#    Task 1 → Task 2 → Task 3 → Task 4
# 4. Default FortiGate credentials (change immediately!)
#    Username: admin
#    Password: (blank — press Enter on first login)
```
Navigate to the official Fortinet support portal and download the VMware-compatible FortiGate VM image.

**Download URL:** https://support.fortinet.com/support/#/downloads/vm

```
Select: FortiGate → VM Images → VMware ESXi/Workstation → FortiGate-VM64.ovf
```
After downloading, extract the `.zip` archive. You should see:

```
FortiGate-VM64/
├── FortiGate-VM64.ovf       ← OVF descriptor (import this)
├── FortiGate-VM64.mf        ← Manifest (checksums)
├── fortios.vmdk             ← Primary disk image
└── datadrive.vmdk           ← Data disk
```
💡 **Tip:** Verify the SHA256 checksum of the downloaded archive against the value published on the Fortinet portal before proceeding.


Open VMware Workstation's **Virtual Network Editor** to create the required host-only network adapters.

**Path:** `Edit → Virtual Network Editor` *(requires Administrator privileges)*

Or search **"Virtual Network Editor"** from the Windows Start Menu.

```
Default networks available:
  VMNet1  →  Host-Only
  VMNet8  →  NAT
Add the following networks (Host-Only type):
  VMNet2  →  192.168.50.0/24    (LAN)
  VMNet3  →  [your choice]      (Additional zone)
  VMNet11 →  [your choice]      (Additional zone)
  VMNet12 →  [your choice]      (Additional zone)
```
**Configuration rules:**


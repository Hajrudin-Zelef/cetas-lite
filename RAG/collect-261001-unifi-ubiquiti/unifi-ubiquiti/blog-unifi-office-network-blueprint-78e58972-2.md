---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/blog-unifi-office-network-blueprint-78e58972-2
title: "blog-unifi-office-network-blueprint-78e58972"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "safeguards", "voice"]
source: docs/RAG/collect-261001-unifi-ubiquiti/blog-unifi-office-network-blueprint-78e58972.md
source_anchor: ""
source_lines: [57, 113]
sha256: 0ae924ad0980893d03ed6e08c3b6ab9e6152ce96b4a782387f5dec6c72c14b49
---

# blog-unifi-office-network-blueprint-78e58972

- Cameras: UniFi G6 Series Cameras should be installed at entrances, hallways, and shared spaces. Choose Turret for ceiling mounts with flexible 3-axis aiming, Bullet for wall-mount perimeter coverage, Dome for vandal-resistant high-traffic areas, or Instant for quick Wi-Fi deployments without Ethernet cabling. All wired G6 models (Bullet, Turret, Dome) are IP66 all-weather rated and connect via PoE; the G6 Instant connects over Wi-Fi and uses USB-C power.
Pre-Wiring for Flexibility
An often-overlooked part of network setup is the pre-wiring of network cables. For new builds, Cat6A is the recommended choice — it supports 10GBASE-T up to 100 m and gives the most headroom. Existing Cat5e or Cat6 can support 2.5GbE and sometimes 5GbE over shorter runs, so a complete re-cable may not be necessary if your existing infrastructure qualifies. Whether setting up a new office or remodeling, having network cabling verified and in place before device installation saves time and reduces costly changes.
Wi-Fi Coverage: Leveraging 2.4 GHz, 5 GHz, and 6 GHz Bands
Wi-Fi 7 APs with tri-band support (U7 Pro and above) can use the 6 GHz band — first introduced in Wi-Fi 6E — alongside the traditional 2.4 GHz and 5 GHz bands. Wi-Fi 7 adds features like 320 MHz channels and Multi-Link Operation (MLO), subject to client and regional support:
Wi-Fi Band Selection
2.4 GHz – This band provides the most comprehensive coverage but operates at slower speeds, making it suitable for devices farther from access points or less demanding tasks.
5 GHz – This frequency offers higher speeds but a shorter range, making it ideal for bandwidth-heavy tasks like video conferencing or file sharing in close proximity to access points.
6 GHz (Wi-Fi 6E/7) – Provides additional spectrum and significantly less legacy congestion than 5 GHz. Range is broadly similar to or slightly below 5 GHz depending on power and obstructions — the advantage is capacity, not coverage distance. Requires WPA3 and compatible client devices; 6 GHz availability varies by region. Only available on the U7 Pro and higher — the U7 Lite does not have a 6 GHz radio.
Plan AP placement to achieve at least -65 dBm RSSI and 25 dB SNR at workstation height across all coverage areas. Validate with a post-install survey — real-world obstructions, interference sources, and client mix will affect actual performance.
Security and Monitoring Integration
Security is an integral part of your network design. UniFi G6 Series Cameras connect through the same infrastructure and are managed via the Protect application on your UniFi console. Network devices use the Network application. Both applications run on the same console hardware and are accessible through Site Manager, but they are distinct interfaces — not a single unified dashboard for all functions.
Network Configuration Blueprint
A proper office network requires logical segmentation to separate traffic types, improve security, and simplify management. This VLAN structure is our recommended starting point for most office deployments.
Recommended VLAN Structure
| VLAN ID | Name | Purpose | Subnet | DHCP Range | 
|---|---|---|---|---|
| 99 | Management | Switches, APs, gateway admin | 10.0.99.0/24 | 10.0.99.100-199 (static or reserved IPs for infrastructure devices) | 
| 10 | Corporate | Employee workstations | 10.0.10.0/24 | 10.0.10.100-254 | 
| 20 | Voice | VoIP phones | 10.0.20.0/24 | 10.0.20.100-254 | 
| 30 | Cameras | Protect cameras | 10.0.30.0/24 | 10.0.30.100-254 | 
| 40 | IoT | Printers, displays, smart devices | 10.0.40.0/24 | 10.0.40.100-254 | 
| 50 | Guest | Visitor WiFi (isolated) | 10.0.50.0/24 | 10.0.50.100-254 | 
Why Segment Your Network?
VLAN segmentation limits lateral movement when paired with restrictive firewall policies. VLANs create logical separation, but firewall rules and isolation policies enforce the boundary — UniFi's default inter-VLAN behavior can still permit traffic if policies are not explicitly configured. PCI DSS strongly recommends segmentation to reduce scope but does not universally require it. HIPAA requires reasonable safeguards based on risk assessment — it does not prescribe VLANs specifically. This design uses VLAN 99 as a dedicated management VLAN rather than the default VLAN 1, keeping infrastructure management traffic isolated from all user and device networks.
Port Profile Configuration
The Switch Pro Max 48 PoE supports port profiles to simplify VLAN assignment. Create reusable Ethernet port profiles in the Network application:
| Profile Name | VLAN Mode | Native VLAN | Tagged VLANs | Use Case | 
|---|---|---|---|---|
| AP Trunk | Trunk | 99 (Mgmt) | 10, 20, 40, 50 | Access Points | 
| Camera | Access | 30 | None | Protect cameras | 
| Workstation | Access | 10 | None | Employee PCs | 
| VoIP Phone | Access | 20 | None | Desk phones | 
Assign profiles under Devices → Switch → Ports → Port Manager in the current Network UI. Reusable Ethernet port profiles are also available under Settings → Profiles. See the current switch settings guide for the latest navigation.
Zone-Based Firewall Policies
UniFi Network 9.0+ uses zone-based firewalling. Group your VLANs into zones and apply policies between them. The general approach is default-deny between zones with explicit exceptions for required services:
| Source Zone | Destination Zone | Policy | Notes | 
|---|---|---|---|
| Guest | All Internal | Block | Use the Hotspot portal or Guest zone; enable client isolation | 
| IoT | All Internal | Block | Allow only necessary services (DHCP, DNS, NTP via gateway) | 
| Cameras | Gateway + NVR IP | Allow | Cameras need DHCP, DNS, NTP, adoption/update paths, and NVR access | 
| Cameras | All Other Internal | Block | Prevent camera-to-workstation traffic | 
| Voice | Internet + Gateway | Allow | VoIP needs DNS, NTP, SIP/RTP to provider; allow local PBX IP if applicable | 
| Voice | All Other Internal | Block | Prevent phone-to-workstation traffic | 
| Authorized Admin Devices | Management | Allow | Only designated admin workstations/IPs can reach VLAN 99 | 
| All Other Internal | Management | Block | Prevent user/device VLANs from reaching management interfaces | 
| Management | Infrastructure Services | Allow | Management devices need access to DHCP, DNS, NTP, firmware updates | 
This is a starting framework — adapt it to your specific services. A "Cameras → NVR only" rule that forgets gateway services (DHCP, DNS) will break camera adoption. A "Voice → Internet only" rule may break a local PBX or provisioning server. Test each policy thoroughly after applying.
For IPv6, apply equivalent zone policies or disable IPv6 on internal networks if your deployment does not require it.
Installation Steps
Step 1: Mounting and Pre-Wiring
Begin by mounting all devices — switches, access points, and cameras — in the locations determined during your planning phase. For new cable runs, use Cat6A for maximum headroom. Verify existing runs against the multigig speed requirements of each device (2.5GbE for most APs; the U7 Pro XG and XGS have 10GbE RJ45 interfaces but will negotiate at 2.5GbE on the Pro Max 48's multigig PoE ports). Include a UPS for the network rack and budget for patch panels, an SFP+ DAC cable for the gateway-to-switch uplink, and cable management.
Step 2: Set Up the Dream Machine Pro Max
Once the devices are mounted, start by turning on the Dream Machine Pro Max and connecting it to the Internet via the WAN port. If you're unsure whether the Pro Max is the right gateway for your needs, see our UniFi Gateway Comparison Guide for alternatives.
First-Boot Wizard:
- Connect via Mobile App or Web: Access the setup wizard through the UniFi Network mobile app or navigate to unifi.ui.com in a browser.
- Sign In: Enter your UI.com account credentials. If upgrading from an existing setup, select Restore from Cloud Backup to migrate your configuration.

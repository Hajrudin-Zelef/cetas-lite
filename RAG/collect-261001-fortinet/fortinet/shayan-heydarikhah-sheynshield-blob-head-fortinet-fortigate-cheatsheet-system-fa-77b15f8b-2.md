---
id: collect-261001-fortinet/fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-system-fa-77b15f8b-2
title: "Test TFTP server reachability"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["license", "parameters"]
source: docs/RAG/collect-261001-fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-system-fa-77b15f8b.md
source_anchor: ""
source_lines: [295, 676]
sha256: a9ac04eb860376aadd433c71665cd95f5040c419222acd848d6e61dfe6abba0e
---

# Test TFTP server reachability

In FortiOS 7.2.0, FortiAP and FortiSwitch devices cannot follow the same multi-build upgrade path as FortiGate devices; they upgrade directly to the target version. ([Fortinet Documentation][3])

- TFTP server installed.
- Firmware copied to TFTP root.
- Correct firmware filename confirmed.
- Correct hardware image confirmed.
- TFTP server IP confirmed.
- Network connectivity confirmed.
- Firewall rules verified.

Example:

```
TFTP Root
└── a.pkg
```
`execute ping 192.168.20.200`
- Ping successful.
- TFTP server reachable.
- Firmware file available.

`execute restore image tftp a.pkg 192.168.20.200````
This operation will replace the current firmware version!
Do you want to continue? (y/n)
```
- Upgrade warning reviewed.
- Firmware source verified.
- Correct image confirmed.
- Upgrade approved.
- Confirmation entered.

FortiGate performs operations including:

- Firmware download.
- Signature verification.
- Firmware integrity verification.
- Firmware installation.
- Reboot.

```
Ping
 ↓
TFTP
 ↓
Download
 ↓
Verify Signature
 ↓
Verify Integrity
 ↓
Install
 ↓
Reboot
 ↓
New FortiOS
```
Controlled upgrades allow a firmware image to be staged in a separate partition before activation. ([Fortinet Documentation][1])

`execute restore secondary-image tftp a.pkg`
Possible sources:

```
FTP
TFTP
USB
```
`execute set-next-reboot secondary`
Important

`secondary` must be used when the newly staged firmware is located in the secondary partition. The `primary`/`secondary` value refers to the partition containing the firmware that should be loaded on the next reboot. ([Fortinet Documentation][1])

- Firmware image prepared.
- Firmware image verified.
- Secondary partition available.
- Firmware staged.
- Staging completed successfully.
- Correct partition identified.
-  `set-next-reboot` configured.
- Maintenance window approved.
- Reboot scheduled.
- Post-reboot validation planned.

```
Current Firmware
      │
      ▼
Secondary Partition
      │
      ├── Load New Firmware
      │
      ▼
Verify
      │
      ▼
set-next-reboot secondary
      │
      ▼
Maintenance Window
      │
      ▼
Reboot
      │
      ▼
New Firmware
```
- Staged upgrades.
- Maintenance-window optimization.
- Multiple FortiGate deployment.
- Script-based deployment.
- FortiManager-controlled workflow.

- Firmware downloaded from trusted Fortinet source.
- Correct hardware model verified.
- Firmware build verified.
- Signature verification performed.
- Integrity verification performed.
- Unexpected signature warning investigated.

```
Firmware Image
      │
      ▼
Digital Signature
      │
      ▼
FortiGate Verification
      │
   ┌──┴──┐
  PASS   FAIL
   │      │
   ▼      ▼
Continue STOP
          Investigate
```
Caution

**Never ignore firmware authenticity warnings in production.**

- Stop upgrade if unexpected.
- Verify firmware source.
- Verify file.
- Verify model compatibility.
- Verify build.
- Re-download firmware if necessary.
- Investigate before proceeding.

Use console-based boot testing when you need to verify a firmware image without immediately making it the permanent default firmware.

- Console cable.
- TFTP server.
- Firmware image.
- Local management information.
- TFTP server IP.
- Firmware filename.

`execute reboot`
Confirm:

```
y
```
Watch for:

```
Press any key to display configuration menu ..
```
- Press key during boot window.
- Enter boot menu.
- Avoid missing the short interruption window.

```
Boot
 │
 ▼
Press Any Key
 │
 ├── YES → Boot Menu
 │
 └── NO  → Normal Boot
```
Typical options:

| Key | Function | 
|---|---|
| `g` | Get firmware image from TFTP | 
| `f` | Format boot device | 
| `b` | Boot backup firmware and set as default | 
| `c` | Configuration/information | 
| `q` | Continue normal boot | 
| `h` | Help | 

- Enter boot menu.
-  Select `g` .
- Enter TFTP server address.
- Enter local FortiGate address.
- Enter firmware filename.
- Download image.
- Verify image.
- Select run-without-saving option where supported.
- Test firmware.
- Validate basic functionality.

[!DANGER]
`f` can format the boot device. **Never select it casually in production.**


Select:

```
c
```
Potential configuration options:

| Key | Function | 
|---|---|
| `p` | Firmware download port | 
| `d` | DHCP mode | 
| `i` | Local IP address | 
| `s` | Local subnet mask | 
| `g` | Local gateway | 
| `v` | VLAN ID | 
| `t` | TFTP server | 
| `f` | Firmware filename | 
| `e` | Reset TFTP parameters | 
| `r` | Review parameters | 
| `n` | Network diagnostic/ping | 
| `q` | Quit | 
| `h` | Help | 

- Download port verified.
- Local IP configured.
- Subnet mask configured.
- Gateway configured if required.
- VLAN configured if required.
- TFTP server configured.
- Firmware filename configured.
- Parameters reviewed.
- Network test completed.

- USB drive available.
- Firmware image copied to USB root.
- Correct filename verified.
- Correct hardware image verified.
- USB connected to FortiGate.

`execute restore image usb <filename>`
Example:

`execute restore image usb a.pkg````
This operation will replace the current firmware version!
Do you want to continue? (y/n)
```
- Warning reviewed.
- Correct firmware confirmed.
- Upgrade approved.
- Reboot expected.

`get system status``execute update-now`
- FortiOS version.
- Build number.
- Device model.
- System status.
- FortiGuard connectivity.
- AV definitions.
- Attack definitions.
- Critical services.

GUI location:

```
System
└── Settings
    └── Startup Settings
```
- Detect Firmware requirement identified.
- Firmware filename configured.
- Firmware copied to USB root.
- USB connected.
- Startup behavior understood.
- Upgrade performed if required.
- Auto-detection disabled after temporary use if no longer required.

Warning

Avoid leaving automatic firmware detection enabled unnecessarily after maintenance or testing.

Caution

**Firmware downgrade is a high-risk operation.**

- Operational reason documented.
- Downgrade path verified.
- Target firmware verified.
- Hardware compatibility verified.
- Configuration backup created.
- Recovery plan prepared.
- Previous firmware available.
- Console access available.
- Maintenance window approved.

`execute ping 192.168.20.200`
Then:

`execute restore image tftp a.pkg 192.168.20.200`
FortiGate may display:

```
This operation will downgrade the current firmware version!
```
- Downgrade warning reviewed.
- Configuration impact understood.
- Confirmation approved.

`get system status``execute update-now`
- Firmware version.
- Build.
- Configuration.
- Interfaces.
- Routing.
- Firewall policies.
- VPN.
- HA.
- Security Fabric.
- FortiGuard.
- Critical services.

`get system status`
- Correct FortiOS version.
- Correct build.
- Correct hardware model.
- System healthy.

- Interfaces operational.
- VLANs operational.
- Routing verified.
- Default route verified.
- DNS verified.
- Internet connectivity verified.

- Firewall policies verified.
- Security profiles verified.
- IPS operational.
- Antivirus operational.
- Web filtering operational.
- Application control operational.

- IPsec tunnels verified.
- SSL VPN verified if used.
- Authentication verified.
- VPN traffic tested.

- HA status verified.
- Primary/secondary roles verified.
- Synchronization verified.
- Session synchronization verified where applicable.
- HA monitoring verified.

- Fabric root healthy.
- Fabric members connected.
- FortiAP status verified.
- FortiSwitch status verified.
- Fabric topology healthy.

- FortiGuard connectivity verified.
- AV definitions updated.
- Attack definitions updated.
- License status verified.

- Critical services verified.
- Logs reviewed.
- Authentication tested.
- Monitoring restored.
- Alerts reviewed.

- Internal traffic tested.
- Internet traffic tested.
- Critical applications tested.
- NAT verified.
- VIPs verified.
- Routing verified.
- VPN traffic verified.


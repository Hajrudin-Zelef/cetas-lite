---
id: collect-261001-fortinet/fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-system-fa-77b15f8b-4
title: "Test TFTP server reachability"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-system-fa-77b15f8b.md
source_anchor: ""
source_lines: [980, 1270]
sha256: 1578c37a9278aaf11702953199f6ac6392bd1df52461d3368eeed127b31c368d
---

# Test TFTP server reachability

The boot-menu `f` option can format the boot device. Treat it as destructive.

```
NO BACKUP
   ↓
NO PRODUCTION UPGRADE
```
```
Current
   ↓
Supported Path
   ↓
Target
```
Never assume:

```
Current → Latest
```
is supported.

```
Source
 ↓
Model
 ↓
Build
 ↓
Signature
 ↓
Integrity
```
Controlled upgrades allow:

```
Stage
 ↓
Validate
 ↓
Schedule
 ↓
Reboot
 ↓
Activate
```
This can reduce maintenance-window pressure.

```
Reboot
 ↓
System Status
 ↓
Definitions
 ↓
Services
 ↓
HA
 ↓
Fabric
 ↓
Traffic
```
```
Firmware
+
Config Backup
+
Console
+
TFTP
+
USB
=
Recovery Capability
```
```
LAB
 ↓
PILOT
 ↓
LOW-RISK
 ↓
PRODUCTION
 ↓
CRITICAL
```
Firmware images and configuration backups are operationally sensitive assets.

- Restrict access.
- Use trusted sources.
- Protect backups.
- Protect recovery media.
- Document version provenance.

```
Fabric Management
↓
Central firmware/device management
```
```
Feature
↓
New features
↓
Gray
```
```
Mature
↓
Maintenance-oriented release
↓
Green
```
```
Upgrade Path
↓
Current
↓
Intermediate
↓
Target
```
```
TFTP Upgrade
↓
execute ping
↓
execute restore image tftp
↓
Signature Verification
↓
Integrity Verification
↓
Install
↓
Reboot
```
```
Controlled Upgrade
↓
restore secondary-image
↓
Stage Firmware
↓
set-next-reboot secondary
↓
Maintenance Window
↓
Reboot
```
```
Federated Upgrade
↓
Fabric Management
↓
Follow Upgrade Path
↓
Intermediate Builds
↓
Multiple Reboots
↓
Target Firmware
```
```
Recovery
↓
Console
↓
Boot Menu
↓
TFTP / Backup / USB
↓
Known-Good Firmware
```
```
Post Upgrade
↓
get system status
↓
execute update-now
↓
Check Services
↓
Check HA
↓
Check Fabric
↓
Check Traffic
```
```
# System
get system status
# Connectivity
execute ping <IP>
# Security definitions
execute update-now
# Reboot
execute reboot
# TFTP firmware
execute restore image tftp <file> <server>
# USB firmware
execute restore image usb <file>
# Controlled firmware staging
execute restore secondary-image tftp <file>
# Boot secondary partition next
execute set-next-reboot secondary
# Boot primary partition next
execute set-next-reboot primary
# Federated upgrade
execute federated-upgrade initialize
# Federated status
execute federated-upgrade status
# Cancel federated upgrade
execute federated-upgrade cancel
```
Before clicking **Upgrade**:

- I know the current firmware.
- I know the target firmware.
- I verified the upgrade path.
- I verified hardware compatibility.
- I verified firmware authenticity.
- I created a configuration backup.
- I have console/OOB access.
- I have a recovery firmware image.
- I have a recovery method.
- I verified HA/Fabric state.
- I verified FortiGuard.
- I verified NTP.
- I verified licensing.
- I have an approved maintenance window.
- I have a rollback plan.
- I know what I will validate after reboot.

-  `get system status`
-  `execute update-now`
- Verify security definitions.
- Verify interfaces.
- Verify routing.
- Verify policies.
- Verify NAT.
- Verify VPN.
- Verify HA.
- Verify Security Fabric.
- Verify FortiGuard.
- Verify critical applications.
- Review logs.
- Confirm production traffic.
- Document final firmware/build.

`FortiGate Firmware Upgrade` ·
`FortiOS Upgrade` ·
`FortiOS 7.2.0 Firmware Upgrade` ·
`FortiGate Upgrade Path` ·
`FortiGate Fabric Management` ·
`Fortinet Security Fabric Upgrade` ·
`FortiGate Federated Upgrade` ·
`FortiGate TFTP Upgrade` ·
`FortiGate USB Firmware Upgrade` ·
`FortiGate Controlled Upgrade` ·
`FortiOS Firmware Verification` ·
`FortiGate Firmware Signature` ·
`FortiGate Firmware Recovery` ·
`FortiGate Boot Menu` ·
`FortiGate Downgrade` ·
`FortiGuard Update` ·
`execute update-now` ·
`execute restore image tftp` ·
`execute restore secondary-image` ·
`execute set-next-reboot` ·
`FortiAP Firmware Upgrade` ·
`FortiSwitch Firmware Upgrade` ·
`FortiOS Upgrade Path` ·
`FortiGate Maintenance` ·
`FortiGate Firmware Troubleshooting`

- 
YouTube — SheynShield 
  - Fortinet NSE content
  - FortiGate troubleshooting
  - Network Security Engineering

**Never treat a firmware upgrade as a single reboot operation.**

**A professional FortiGate upgrade is a controlled lifecycle:**

`PLAN → BACKUP → VERIFY → STAGE → UPGRADE → REBOOT → UPDATE → VALIDATE → DOCUMENT`

**Engineering Secure Networks**

**SheynShield | Security & Design Knowledge Base**

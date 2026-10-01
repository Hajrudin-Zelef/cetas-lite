---
id: collect-261001-fortinet/fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-system-fa-77b15f8b-3
title: "Test TFTP server reachability"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["licenses"]
source: docs/RAG/collect-261001-fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-system-fa-77b15f8b.md
source_anchor: ""
source_lines: [677, 979]
sha256: 602c94462f6e2f1a122ef3935797c8851455712a8e91476abd89782998570c8e
---

# Test TFTP server reachability

- Backup configuration.
- Verify upgrade path.
- Verify target firmware.
- Verify hardware compatibility.
- Verify firmware authenticity.
- Verify firmware integrity.
- Verify FortiGuard.
- Verify licenses.
- Verify NTP.
- Verify HA.
- Verify Fabric.
- Verify VPN.
- Verify routing.
- Verify critical services.
- Prepare rollback plan.
- Prepare console access.
- Prepare TFTP/USB recovery.
- Approve maintenance window.

- Monitor firmware transfer.
- Monitor signature verification.
- Monitor integrity verification.
- Monitor reboot.
- Do not interrupt power.
- Do not interrupt firmware transfer.
- Monitor HA state.
- Monitor Fabric state.
- Monitor management connectivity.

-  `get system status`
-  `execute update-now`
- Verify AV definitions.
- Verify attack definitions.
- Verify interfaces.
- Verify routing.
- Verify firewall policies.
- Verify NAT.
- Verify VPNs.
- Verify HA.
- Verify Fabric.
- Verify FortiGuard.
- Verify logs.
- Verify monitoring.
- Test critical applications.
- Confirm production traffic.
- Document final state.

```
Firmware Upgrade Failed
        │
        ▼
Can FortiGate Boot?
   ┌────┴────┐
  YES       NO
   │         │
   ▼         ▼
Check OS   Console
   │         │
   │         ▼
   │      Boot Menu
   │         │
   │         ▼
   │      TFTP / USB
   │         │
   └────┬────┘
        ▼
Verify Firmware
        │
        ▼
Verify Signature
        │
        ▼
Verify Integrity
        │
        ▼
Verify Upgrade Path
        │
        ▼
Recover / Retry
```
-  Run `get system status` .
- Verify firmware version.
- Verify build.
- Verify interfaces.
- Verify routing.
- Verify services.
- Verify FortiGuard.
- Verify security definitions.
- Review logs.

- Connect console.
- Interrupt boot.
- Enter boot menu.
- Verify TFTP connectivity.
- Verify firmware image.
- Verify image filename.
- Verify boot partition.
- Use recovery workflow.
- Avoid destructive formatting unless explicitly required.

| Purpose | Command | 
|---|---|
| System status | `get system status` | 
| Test connectivity | `execute ping <IP>` | 
| Update definitions | `execute update-now` | 
| Reboot | `execute reboot` | 
| TFTP upgrade | `execute restore image tftp <file> <server>` | 
| USB upgrade | `execute restore image usb <file>` | 
| Stage firmware | `execute restore secondary-image tftp <file>` | 
| Select secondary boot | `execute set-next-reboot secondary` | 
| Select primary boot | `execute set-next-reboot primary` | 
| Initialize federated upgrade | `execute federated-upgrade initialize` | 
| Federated status | `execute federated-upgrade status` | 
| Cancel federated upgrade | `execute federated-upgrade cancel` | 
| Restart federated upgrade | `execute federated-upgrade restart` | 

Important

Always verify command availability and exact behavior against the specific FortiOS build and hardware platform before production execution.

```
# Test TFTP server reachability
execute ping 192.168.20.200
# Start firmware upgrade
execute restore image tftp a.pkg 192.168.20.200
# Confirm the firmware replacement
y
# Confirm additional warning if displayed
y
# After reboot
get system status
# Update security definitions
execute update-now
```
- Ping successful.
- Firmware file exists.
- Correct model verified.
- Upgrade path verified.
- Firmware signature verified.
- Upgrade confirmed.
- FortiGate rebooted.
- New firmware verified.
- Security definitions updated.
- Production traffic tested.

```
# Stage new firmware into secondary partition
execute restore secondary-image tftp a.pkg
# Configure the partition containing the staged firmware
# to be loaded at the next reboot
execute set-next-reboot secondary
# Reboot during the maintenance window
execute reboot
```
- Firmware staged successfully.
- Secondary partition contains target firmware.
-  `set-next-reboot secondary` configured.
- Maintenance window reached.
- Reboot performed.
- New firmware loaded.
- System status verified.
- Production traffic validated.

Warning

Do not blindly use `primary` after staging an image to `secondary`. `set-next-reboot` selects the partition that will be booted next. ([Fortinet Documentation][1])

| Method | Source | Activation | Primary Use | 
|---|---|---|---|
| GUI Upgrade | FortiGuard / uploaded image | Immediate or scheduled | Normal upgrade | 
| Fabric Upgrade | FortiGuard | Immediate or scheduled | Fabric-wide upgrade | 
| Federated Upgrade | FortiGuard | Coordinated | Multi-build/Fabric upgrade | 
| TFTP Restore | TFTP | Immediate | Manual/recovery | 
| USB Restore | USB | Immediate | Offline/recovery | 
| Controlled Upgrade | Secondary partition | Next reboot | Staged activation | 
| Boot Menu TFTP | TFTP | Boot-time | Recovery/testing | 

- Console cable.
- Laptop.
- TFTP server.
- Known-good firmware.
- Previous firmware.
- USB recovery media.
- Configuration backup.
- Management IP information.
- TFTP server IP.
- Firmware filename.

```
Failed Upgrade
      │
      ▼
Console Access
      │
      ▼
Boot Menu
      │
      ├── TFTP
      │
      ├── Backup Firmware
      │
      └── USB / Alternative Recovery
      │
      ▼
Boot Known-Good Image
      │
      ▼
Verify Configuration
      │
      ▼
Verify Network
      │
      ▼
Verify Services
      │
      ▼
Restore Production State
```
- Firmware downloaded from trusted source.
- Firmware authenticity verified.
- Signature warnings investigated.
- Firmware model compatibility verified.
- Firmware file protected.
- TFTP access restricted.
- USB media protected.
- Configuration backups protected.
- Recovery images protected.

- Management access restricted.
- Console access controlled.
- Firmware upgrade permissions restricted.
- Maintenance window approved.
- Upgrade activity logged.
- Recovery procedure documented.

- OOB management available.
- Console access available.
- HA state verified.
- Fabric topology verified.
- NTP synchronized.
- FortiGuard operational.
- Licenses valid.
- Backup available.
- Previous firmware available.
- TFTP server ready.
- USB recovery ready.
- Maintenance window available.
- Rollback plan documented.

Important

**Fabric Management** is the central GUI area used for firmware management in FortiOS 7.2.x.

Important

**FortiOS 7.2.0 introduced Feature and Mature firmware maturity tags.** ([Fortinet Documentation][2])

Tip

Feature firmware is associated with newer functionality; Mature firmware does not introduce new major features and focuses on maintenance, fixes, and applicable vulnerability patches. ([Fortinet Documentation][5])

Important

`get system status` can be used to identify the FortiOS version, build, and maturity notation.

Important

A **federated update** can automatically follow multiple firmware builds in an upgrade path for supported FortiGate devices. ([Fortinet Documentation][3])

Warning

FortiAP and FortiSwitch do not follow the same multi-build upgrade path behavior as FortiGate in the FortiOS 7.2.0 federated-update workflow. ([Fortinet Documentation][3])

Important

`execute restore image tftp` replaces the current firmware and causes a reboot after installation.

Important

Firmware signature verification is a security control. Do not ignore unexpected authenticity warnings.

Important

`execute restore secondary-image` stages firmware in another partition for later activation. ([Fortinet Documentation][1])

Important

`execute set-next-reboot secondary` selects the secondary partition for the next boot. ([Fortinet Documentation][1])

Warning

`execute set-next-reboot primary` does **not** mean "boot the newly uploaded firmware." It means boot the firmware stored in the primary partition.

Important

`execute update-now` should be considered during post-upgrade security-definition validation.

Warning

A firmware downgrade can introduce configuration migration/reversion problems.

Caution


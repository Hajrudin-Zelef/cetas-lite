---
id: collect-261001-fortinet/fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-system-fa-77b15f8b-1
title: "Test TFTP server reachability"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["licenses"]
source: docs/RAG/collect-261001-fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-system-fa-77b15f8b.md
source_anchor: ""
source_lines: [1, 294]
sha256: cd8bfd349a70fc89b3b46fb499c907291dee688085f9a0adf00028cb52341efc
---

# Test TFTP server reachability

**FortiOS 7.2.0 | Security Fabric | Fabric Management | Federated Upgrade | Firmware Upgrade | TFTP | USB | Controlled Upgrade | Firmware Verification | Recovery**

**SheynShield | Engineering Secure Networks**


FortiGate **Fabric Management** provides centralized firmware and device management for devices participating in the Security Fabric.

- FortiGate devices identified.
- FortiAP devices identified.
- FortiSwitch devices identified.
- Fabric topology verified.
- Fabric root FortiGate identified.
- Fabric members authorized.
- Device registration status verified.
- Device operational status verified.
- Current firmware versions documented.
- Target firmware version documented.
- Upgrade status reviewed.
- Firmware maturity reviewed.

```
System
└── Fabric Management
```
- Device authorization.
- Device registration.
- Device monitoring.
- Firmware version visibility.
- Firmware maturity visibility.
- Upgrade management.
- Upgrade scheduling.
- Multi-device upgrade.
- Fabric upgrade status.

Caution

**Never start a production firmware upgrade without a rollback/recovery plan.**

- Current FortiOS version documented.
- Current build number documented.
- Hardware model verified.
- HA topology documented.
- Security Fabric topology documented.
- Managed FortiAP devices documented.
- Managed FortiSwitch devices documented.
- Critical services documented.
- Critical VPNs documented.
- Routing dependencies documented.
- External dependencies documented.

- FortiGuard connectivity verified.
- DNS resolution verified.
- Internet connectivity verified where required.
- Routing verified.
- Management connectivity verified.
- Out-of-band management available where possible.
- Console access available.

- NTP configured.
- NTP synchronization verified.
- System time verified.
- Time zone verified.

- FortiCare status verified.
- FortiGuard services verified.
- Required licenses verified.
- Security service subscriptions verified.

- Configuration backup created.
- Backup stored outside the device.
- Backup tested/reviewed.
- Recovery configuration available.
- Previous firmware image available.

- Target firmware downloaded.
- Correct hardware image selected.
- Firmware source verified.
- Firmware authenticity verified.
- Firmware integrity verified.
- Upgrade path verified.
- Firmware maturity reviewed.
- Feature/Mature classification reviewed.

- Firmware tested in lab where possible.
- Pilot device identified.
- Low-risk devices identified.
- Production devices grouped by risk.
- Critical devices scheduled last where possible.

- Maintenance window approved.
- Expected reboot time estimated.
- Stakeholders notified.
- Rollback procedure documented.
- Recovery equipment available.
- TFTP server prepared if required.
- USB recovery media prepared if required.

Starting with **FortiOS 7.2.0**, FortiOS firmware images use maturity tags. ([Fortinet Documentation][2])

- Feature tag identified.
- New functionality requirements reviewed.
- Production risk evaluated.
- Feature release warning reviewed.

- Mature tag identified.
- Release contains no new major features.
- Bug fixes reviewed.
- Vulnerability patches reviewed.
- Production suitability evaluated.

```
Feature → Gray
Mature  → Green
```
`get system status`
Look for:

```
GA.F → Feature
GA.M → Mature
```
- Current firmware maturity identified.
- Target firmware maturity identified.
- Mature → Feature transition evaluated.
- Feature firmware warning reviewed.

Important

**Feature vs Mature is a release maturity classification, not simply "new vs old".**

- Current FortiOS version identified.
- Target FortiOS version identified.
- Official upgrade path checked.
- Intermediate versions identified.
- Direct upgrade availability verified.
- Hardware compatibility verified.
- Configuration migration impact reviewed.
- Feature changes reviewed.
- HA dependencies reviewed.
- Fabric dependencies reviewed.

```
Current Version
      │
      ▼
Target Version
      │
      ▼
Supported Upgrade Path?
      │
   ┌──┴──┐
  YES    NO
   │      │
   ▼      ▼
Proceed  STOP
         Review
         Intermediate
         Versions
```
Warning

**Do not assume that every FortiOS release can be upgraded directly to every newer release.**

When the upgrade path contains multiple builds, FortiGate can automatically follow the path during a federated update. ([Fortinet Documentation][3])

- Follow upgrade path option reviewed.
- Intermediate versions reviewed.
- Multiple reboot requirement understood.
- Direct upgrade risk evaluated.

- Local configuration backup created.
- External backup created.
- FortiManager backup verified if applicable.
- Backup timestamp recorded.
- Backup location documented.
- Recovery configuration available.
- Configuration changes frozen before upgrade.

For custom/scheduled Fabric upgrades:

- Configuration backup timing understood.
- Pending configuration changes reviewed.
- Final configuration backup performed if necessary.
- No uncommitted production changes remain.

Caution

In a custom scheduled upgrade, the configuration backup is created when the upgrade is scheduled. Configuration changes made afterward may not be included in that scheduled backup. ([Fortinet Documentation][4])

- FortiGuard connectivity verified.
- DNS resolution verified.
- Routing verified.
- Licensing verified.
- Auto-update functionality verified.

After firmware installation:

`execute update-now`
- AV definitions updated.
- Attack definitions updated.
- FortiGuard connectivity verified.
- Security services operational.
- System processes healthy.

```
Firmware Upgrade
      │
      ▼
Reboot
      │
      ▼
get system status
      │
      ▼
execute update-now
      │
      ▼
Verify Definitions
      │
      ▼
Verify Services
      │
      ▼
Validate Traffic
```
Important

**A successful firmware installation does not automatically mean the complete system is healthy.**

```
System
└── Fabric Management
    └── Fabric Upgrade
```
- Logged into correct FortiGate.
- Root FortiGate identified.
- Fabric topology verified.
- Target firmware selected.
- Latest firmware reviewed.
- All available upgrades reviewed.
- Maturity level reviewed.
- Upgrade path reviewed.
- Immediate vs Custom schedule selected.
- Configuration backup confirmed.
- Upgrade schedule reviewed.
- Device list reviewed.
- Upgrade started during approved window.

Fortinet documents Fabric Upgrade as the centralized mechanism for upgrading the root/Fabric devices and managed devices. ([Fortinet Documentation][4])

Before Fabric firmware-management workflows:

- FortiGuard Auto-Update Tunnel requirements reviewed.
- FortiGuard connectivity verified.
- DNS verified.
- Routing verified.
- NTP verified.
- Licensing verified.
- Fabric connectivity verified.

```
Fabric Device
      │
      ▼
FortiGate / Fabric Root
      │
      ▼
FortiGuard
      │
      ▼
Firmware / Security Services
```
A federated upgrade can coordinate firmware upgrades across Fabric devices and can follow a multi-build upgrade path where supported. ([Fortinet Documentation][3])

`execute federated-upgrade initialize``execute federated-upgrade status``execute federated-upgrade cancel``execute federated-upgrade restart`
Note

`restart` is available in later 7.2.x documentation. Always verify the exact CLI options against the FortiOS build being used.

- Fabric topology verified.
- Root FortiGate identified.
- Target version selected.
- Upgrade path reviewed.
- Intermediate versions reviewed.
- Upgrade schedule selected.
- Configuration backup confirmed.
- Fabric devices reviewed.
- FortiGate upgrade sequence reviewed.
- FortiAP behavior reviewed.
- FortiSwitch behavior reviewed.
- Multiple reboot requirement understood.

Important


---
id: collect-261001-fortinet/fortinet/sheynshield-fortinet-fortigate-cheatsheet-foundation-md-at-172fb9490610bedf2470836990d0ff5-3
title: "sheynshield-fortinet-fortigate-cheatsheet-foundation-md-at-172fb9490610bedf2470836990d0ff5347fa74a0-"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["cost", "distribution", "memory", "sandbox"]
source: docs/RAG/collect-261001-fortinet/sheynshield-fortinet-fortigate-cheatsheet-foundation-md-at-172fb9490610bedf2470836990d0ff5347fa74a0-.md
source_anchor: ""
source_lines: [1006, 1540]
sha256: afc66d51f69cdfe77dc533b2a8eb96c1423dcb8f81c7a313c17e0f02f726b00b
---

# sheynshield-fortinet-fortigate-cheatsheet-foundation-md-at-172fb9490610bedf2470836990d0ff5347fa74a0-

```
FortiGate
    =
Network Security
FortiWeb
    =
Web Application Security
```
FortiMail provides email security.

Concept:

```
Internet
   ↓
FortiMail
   ↓
Mail Infrastructure
```
Typical security areas:

- Anti-spam
- Malware protection
- Email security
- Content inspection
- Email policy enforcement

FortiNAC provides Network Access Control.

Concept:

```
Device
   ↓
Identity / Device Profiling
   ↓
Policy
   ↓
Network Access
```
Potential use cases:

- Device visibility
- Network access control
- Segmentation
- Endpoint/device profiling
- Guest access

FortiAuthenticator provides identity and authentication services.

Concept:

```
User
  ↓
Identity
  ↓
Authentication
  ↓
Authorization
  ↓
Network Access
```
It can integrate with enterprise identity environments and authentication mechanisms.

FortiSandbox provides sandbox-based malware analysis.

Concept:

```
Suspicious File
       ↓
FortiSandbox
       ↓
Analysis
       ↓
Behavior
       ↓
Verdict
```
Sandboxing complements signature-based detection by providing deeper analysis of suspicious objects.

FortiADC provides application delivery capabilities.

It can be conceptually compared to an application delivery/load-balancing platform.

```
Clients
   │
   ▼
FortiADC
   │
   ├── Server 1
   ├── Server 2
   └── Server 3
```
Typical concepts:

- Load balancing
- Application delivery
- Server health monitoring
- Traffic distribution

Wireless Access Point.

```
Client
  ↓
FortiAP
  ↓
Network
```
Wireless LAN Controller functionality/platform.

Concept:

```
                    Controller
                         │
              ┌──────────┼──────────┐
              ▼          ▼          ▼
           FortiAP     FortiAP     FortiAP
```
Wireless management can use CAPWAP.

FortiToken provides authentication-token functionality.

Concept:

```
Username + Password
        +
OTP / Token
        ↓
     MFA
```
FortiToken Mobile provides mobile token functionality.

FortiOS is the operating system used by FortiGate.

Conceptual stack:

```
                    FortiOS
────────────────────────────────
Firewall
Routing
VPN
SD-WAN
Security Profiles
Authentication
HA
VDOM
Logging
Wireless
Management
Troubleshooting
────────────────────────────────
Hardware Acceleration
────────────────────────────────
FortiGate Hardware
```
Configuration can be managed through:

```
GUI
CLI
FortiManager
Automation / APIs
```
Useful for:

- Configuration
- Monitoring
- Dashboard
- Reports
- Wizards

Essential for:

- Advanced configuration
- Troubleshooting
- Debugging
- Automation
- Precise control

```
GUI
 =
Visibility + Convenience
CLI
 =
Precision + Troubleshooting
```
Before major changes:

```
Firmware Upgrade
HA Changes
Major Policy Changes
Routing Changes
Interface Changes
VDOM Changes
```
perform:

```
BACKUP
```
Recommended backup architecture:

```
FortiGate
   ↓
Encrypted Backup
   ↓
Secure Repository
   ↓
Versioned Storage
```
Never publish production configuration backups to a public GitHub repository.

General workflow:

```
Backup File
    ↓
Verify Hardware
    ↓
Verify FortiOS Version
    ↓
Verify VDOM Structure
    ↓
Restore
    ↓
Reboot if Required
    ↓
Validate
```
Before restoring, check:

```
Platform
Firmware
Interfaces
VDOMs
Certificates
Encrypted Secrets
HA Configuration
```
Recommended process:

```
Current Version
      ↓
Check Upgrade Path
      ↓
Check Hardware Compatibility
      ↓
Read Release Notes
      ↓
Backup Configuration
      ↓
Upgrade
      ↓
Reboot
      ↓
Validate
```
Never think:

```
Upload
  ↓
Reboot
  ↓
Done
```
A production firmware upgrade is a controlled change.

When normal boot/recovery procedures are insufficient, a low-level firmware recovery process may be required.

High-level workflow:

```
Firmware Image
      ↓
TFTP Server
      ↓
FortiGate Boot Menu
      ↓
Network Connection
      ↓
Firmware Download
      ↓
Flash / Run Image
      ↓
Boot
```
Typical procedure:

1. Prepare the correct firmware image.
2. Configure a TFTP server.
3. Connect the required FortiGate interface.
4. Connect to the FortiGate console.
5. Reboot the device.
6. Enter the boot menu.
7. Select the firmware download/recovery option.
8. Enter the TFTP server IP.
9. Configure the FortiGate local IP.
10. Enter the firmware filename.
11. Download the firmware.
12. Select the appropriate boot/save option.
13. Reboot.
14. Validate system operation.

Do not assume that a single interface such as `WAN1` is universally used for firmware recovery.

The required interface and exact boot-menu procedure depend on the FortiGate model.

```
get system status
```
Useful information:

- Hostname
- Serial number
- FortiOS version
- System information

```
show system interface
```
```
config system interface
    edit "port1"
        set ip 192.168.1.1 255.255.255.0
        set allowaccess https ping ssh
    next
end
```
```
diagnose sys top
```
Useful for observing:

```
CPU
Processes
Memory
System Activity
```
```
execute reboot
```
```
execute factoryreset
```
**DESTRUCTIVE COMMAND**

Use only when the operational impact is fully understood.

```
execute shutdown
```
On supported releases/platforms:

```
execute backup config flash
```
Always verify syntax for the installed FortiOS version.

At NSE 4 level, you must understand how FortiGate forwards traffic.

At NSE 7 level, you must be able to determine **why the expected behavior is not occurring**.

The core methodology is:

```
FOLLOW THE PACKET
```
NSE 4-level FortiGate engineering should cover:

```
Interfaces
VLAN
Routing
Static Routes
Dynamic Routing
VDOM
HA
SD-WAN
VPN
DHCP
DNS
```
```
Firewall Policies
NAT
Authentication
Antivirus
IPS
Web Filtering
Application Control
SSL Inspection
Security Profiles
```
```
Administrators
Admin Profiles
Management Access
Certificates
Backup
Restore
Firmware
Logging
```
At advanced level, the goal changes from:

```
"What command do I use?"
```
to:

```
"Why is this system behaving this way?"
```
Important domains:

```
Advanced Threat Protection
Enterprise Firewall
Advanced FortiGate Troubleshooting
Secure Access
Cloud Security
Performance
Architecture
Traffic Flow
Security Inspection
Hardware Acceleration
```
The advanced engineer should think in layers.

```
Physical
   ↓
Interface
   ↓
Layer 2
   ↓
Layer 3
   ↓
Routing
   ↓
Policy
   ↓
NAT
   ↓
Session
   ↓
Security Inspection
   ↓
Hardware Offload
   ↓
Return Traffic
```
The question is not:

```
"Which command should I type?"
```
The question is:

```
"At which stage did the expected behavior change?"
```
```
Routing
   =
WHERE?
```
```
Firewall Policy
   =
WHETHER?
```
```
NAT
   =
Address Translation
```
```
Firewall Policy
   =
Traffic Authorization
```
```
IDS
   =
Detection
```
```
IPS
   =
Detection + Prevention
```
```
TCP/443
   ≠
Guaranteed HTTPS Application Identity
```
Application identification requires deeper inspection/context.

```
More Inspection
      ↓
More Processing
      ↓
Potential Performance Impact
```
Hardware acceleration can reduce the performance cost for supported traffic/features.

```
Traffic
   ↓
Can It Be Offloaded?
   │
   ├── YES → Hardware Acceleration
   │
   └── NO  → CPU Processing
```
A packet going out successfully does not guarantee a successful session.

Always check:


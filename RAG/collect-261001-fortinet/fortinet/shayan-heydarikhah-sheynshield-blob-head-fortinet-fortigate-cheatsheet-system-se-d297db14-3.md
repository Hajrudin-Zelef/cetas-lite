---
id: collect-261001-fortinet/fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-system-se-d297db14-3
title: "Configure administrator timeout"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-system-se-d297db14.md
source_anchor: ""
source_lines: [574, 896]
sha256: 9d71a72cdfc2cd9b56c3c2cc5a5fa3734202b996f42fd7215fa9d07d5a8493c8
---

# Configure administrator timeout

                 port2
                   |
             +-----------+
             |  FGT-1    |
             +-----------+
Primary:
0.0.0.0/0
    via 192.168.254.2
    distance 10
Backup:
0.0.0.0/0
    via 12.12.12.2
    distance 250
Concept:
Incoming Interface:
    WAN
Destination:
    192.168.20.0/24
Outgoing Interface:
    port3
Gateway:
    12.12.12.2
Example:
config system link-monitor
    edit "WAN1"
        set addr-mode ipv4
        set srcintf "port1"
        set server-config default
        set server-type static
        set server 192.168.254.2
        set protocol ping
        set gateway-ip 192.168.254.2
        set source-ip 0.0.0.0
        set interval 500
        set probe-timeout 500
        set failtime 5
        set recoverytime 5
        set probe-count 30
        set ha-priority 1
        set update-cascade-interface enable
        set update-static-route enable
        set update-policy-route enable
        set status enable
    next
    edit "WAN2"
        set addr-mode ipv4
        set srcintf "port2"
        set server-config default
        set server-type static
        set server 192.168.200.1
        set protocol ping
        set gateway-ip 192.168.200.1
        set source-ip 0.0.0.0
        set interval 500
        set probe-timeout 500
        set failtime 5
        set recoverytime 5
        set probe-count 30
        set ha-priority 1
        set update-cascade-interface enable
        set update-static-route enable
        set update-policy-route enable
        set status enable
    next
end
| Parameter | Meaning | 
|---|---|
| srcintf | Interface used for the probe | 
| server | Health-check destination | 
| protocol | Probe protocol | 
| gateway-ip | Gateway associated with monitor | 
| interval | Probe interval | 
| probe-timeout | Probe timeout | 
| failtime | Failure threshold | 
| recoverytime | Recovery threshold | 
| probe-count | Probe count | 
| update-static-route | Update static route state | 
| update-policy-route | Update policy route state | 
| update-cascade-interface | Cascade interface state changes | 
For a routing/failover lab:
Firewall policies
    |
    +-- WAN access
    +-- Internal access
NAT should be enabled/disabled according to the actual traffic design.
Do not blindly use NAT disabled in production. The correct setting depends on whether the traffic requires source NAT.
On both FortiGates:
config system settings
    set auxiliary-session enable
    set asymroute enable
end
For multi-FortiGate designs:
Keep participating FortiGate devices on compatible and preferably identical FortiOS releases when reproducing/operating a feature-sensitive design.
Example conceptual Netwatch logic:
Netwatch
 |
 +-- UP
 |    |
 |    +-- Enable primary route
 |    +-- Disable backup route
 |
 +-- DOWN
      |
      +-- Disable primary route
      +-- Enable backup route
Example:
Host:
192.168.20.200
Route logic:
Primary:
0.0.0.0/0 -> 192.168.254.2
Backup:
0.0.0.0/0 -> 192.168.200.1
Masquerade:
IP
 └── Firewall
      └── NAT
           └── masquerade
FortiGate provides configuration save behavior through:
config system global
    set cfg-save automatic
end
Available modes include:
automatic
manual
revert
Fortinet documents these configuration-save modes under config system global.
config system global
    set cfg-save automatic
end
Changes are automatically saved.
Concept:
CLI/GUI Change
      |
      v
RAM
      |
      v
Flash
config system global
    set cfg-save manual
end
Changes are applied but must be explicitly saved.
execute cfg saveconfig system global
    set cfg-save revert
    set cfg-revert-timeout 600
end
Concept:
Change configuration
       |
       v
Temporary state
       |
       +---- Commit/save ---> Keep changes
       |
       +---- Timeout/reboot -> Revert
600 means 600 seconds.
Save:
execute cfg save
Reload/reboot:
execute cfg reloadSAVE   = write configuration
RELOAD = reboot/reload device
Fortinet documents execute cfg save and execute cfg reload in configuration save mode.
On supported FortiGate hardware, TPM can protect sensitive cryptographic material.
TPM can help protect:
- Passwords
- Keys
- Sensitive configuration data
Fortinet states that TPM is disabled by default on supported FortiGate hardware.
Enable:
config system global
    set private-data-encryption enable
end
FortiGate then prompts for a:
32 hexadecimal digit
master/private data encryption key
Simplified:
32-hex-digit
Master Encryption Password
          |
          v
       AES-128-CBC
          |
          v
Sensitive Data
          +
          |
          v
        TPM
          |
          v
     RSA-2048
          |
          v
Primary Key
Fortinet describes the master-encryption-password as protecting sensitive data using AES-128-CBC and the TPM-generated 2048-bit primary key as protecting that master password using RSA-2048.
Critical exam point:
TPM != Full Disk Encryption
Fortinet explicitly notes that the TPM module does not encrypt the disk drive.
TPM-protected configuration introduces an important dependency.
Backup
  |
  v
TPM-protected encryption information
  |
  v
Restore
Restore behavior depends on:
TPM state
+
Master Encryption Password
| Restore Condition | Result | 
|---|---|
| TPM disabled | Protected configuration cannot be restored | 
| TPM enabled + wrong master key | Cannot restore protected configuration | 
| TPM enabled + matching master key | Configuration can be restored | 
For an HA cluster:
FGT-A
  |
  +-- Master Encryption Key
  |
  v
FGT-B
HA members must use the same master encryption key for the cluster to form correctly and synchronize protected configuration.
diagnose hardware test infodiagnose hardware deviceinfo tpmdiagnose tpm
Examples include:
- Alert email credentials
- BGP/routing-related credentials
- External resource credentials
- FortiGuard proxy password
- FortiToken seeds
- HA password
- IPsec PSK
- Link Monitor server password
- Local certificate private keys
- LDAP/RADIUS/FSSO credentials
- Modem/PPPoE credentials
- NTP credentials
- SDN connector credentials
- SNMP credentials
- Wireless security credentials
Private Data Encryption
        |
        +-- Passwords
        +-- PSKs
        +-- Private Keys
        +-- Authentication Secrets
        +-- Service Credentials
FortiGate replacement messages can customize user-facing messages generated by security features.
Common groups/categories include:
UTM
 ├── admin
 ├── alertmail
 ├── custom-message
 ├── fortiguard-wf
 ├── ftp
 ├── http
 ├── icap
 ├── mail
 ├── nac-quar
 ├── spam
 ├── sslvpn
 ├── traffic-quota
 ├── utm
 └── webproxy
AUTH
 ├── auth
 └── webproxy
Enable GUI replacement-message groups:
config system settings
    set gui-replacement-message-groups enable
endconfig emailfilter profile
    edit "newmsgs"
        set replacemsg-group "newutm"
    next
end
Firewall policy:
config firewall policy
    edit 1
        set replacemsg-override-group "newauth"
        set inspection-mode proxy
        set emailfilter-profile "newmsgs"
    next
endFirewall Policy
       |
       +-- Proxy inspection
       |
       +-- Email Filter Profile
       |
       +-- Replacement Message
FortiGate GUI provides a CLI/script interface for deploying configuration commands.
Typical workflow:
GUI
 |
 v
CLI / Script
 |
 v
Paste commands
 |
 v
Execute
A CLI script comment begins with:
#
Example:
# Configure administrator timeout
config system global
    set admintimeout 10
end
Everything after # on that line is treated as a comment.
Workspace/transaction mode is useful when configuration changes need controlled deployment.
Concept:
Administrator A
       |
       v
Transaction
       |
       +---- Validate
       |
       +---- Commit
       |
       +---- Abort
Instead of immediately exposing every change to the rest of the system, configuration changes are held in a transaction until committed.

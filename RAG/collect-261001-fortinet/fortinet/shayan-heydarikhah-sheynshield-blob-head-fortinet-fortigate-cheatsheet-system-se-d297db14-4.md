---
id: collect-261001-fortinet/fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-system-se-d297db14-4
title: "Configure administrator timeout"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-system-se-d297db14.md
source_anchor: ""
source_lines: [897, 1186]
sha256: 6e4b7f51d9e8fd135b4071b4725629d707fd002e36d3e6bb015634c4be479d82
---

# Configure administrator timeout

Fortinet documents Workspace Mode as a change-control mechanism where changes are made in a local CLI process and become available to other processes after commit.
execute config-transaction start
Or specify a timeout:
execute config-transaction start 5
The timeout is specified in minutes.
FortiOS supports a transaction timeout from 1 to 60 minutes, with a default of 5 minutes.
execute config-transaction commit
Concept:
Transaction
    |
    v
Commit
    |
    v
Configuration becomes active globally
execute config-transaction abort
Concept:
Transaction
    |
    v
Abort
    |
    v
Discard transaction changes
Example:
config transaction id=1 will expire in 30 seconds
config transaction id=1 will expire in 20 seconds
config transaction id=1 will expire in 10 seconds
config transaction id=1 has expired
If the administrator does not commit the transaction before expiration, the transaction expires.
You may encounter:
Can not config the object since either the object or the referenced objects
are being configured by other transactions.
Command fail.
Return code 14
Meaning:
Another active transaction
        |
        v
Object/reference locked
        |
        v
Your configuration cannot modify it
diagnose sys config-transaction show txn-meta
Useful for:
Active transaction information
Concurrent transaction/session information
diagnose sys config-transaction show txn-info
Useful for:
Transaction changes
Modification information
diagnose sys config-transaction show txn-entity
Shows affected entities.
diagnose sys config-transaction show txn-lock
Shows locked objects.
Useful when another administrator is modifying configuration.
diagnose sys config-transaction show txn-cli-commands
Shows CLI commands associated with the transaction.
diagnose sys config-transaction show mctx
Useful for examining memory/context information related to the transaction.
If memory resources become exhausted, configuration transaction operations may be affected.
Some configuration areas cannot be changed through a workspace transaction.
Examples include:
config system console
config system resource-limits
config system elbc
config system global
set split-port
set vdom-admin
set management-vdom
set wireless-mode
set internal-switch-mode
end
config system settings
set opmode
end
config system npu
config system np6
config system wireless
set mode
end
config system vdom-property
config system storage
Fortinet documents these restrictions for Workspace Mode.
admin-lockout-threshold
        |
        v
Failed login count
admin-lockout-duration
        |
        v
Lockout time
reuse-password enable
        |
        v
Password reuse allowed
reuse-password disable
        |
        v
New password required
min-change-characters > 0
        |
        v
Overrides reuse-password
NTP
 |
 +-- Normal network time synchronization
PTP
 |
 +-- Precision/time-sensitive environments
ntpv3 enable
    =
NTPv3
ntpv3 disable
    =
NTPv4
E2E = End-to-End
P2P = Peer-to-Peer
FortiOS <= 6.0
    -> Not supported
6.2.0 - 6.2.2
    -> Permanently enabled
6.2.3+
    -> Disabled by default
ECMP
SD-WAN
Multiple paths
Asymmetric traffic
       |
       v
Auxiliary Session
       |
       v
Additional session handling
config system settings
    set asymroute enable
endFirst packet / session setup
            |
            v
           CPU
            |
            v
       Session created
            |
            v
       NPU eligibility
            |
            v
           NPU
            |
            v
    High-speed forwarding
Path/interface change
        |
        v
Session update required
        |
        v
DIRTY
automatic
    |
    +-- Automatically save
manual
    |
    +-- execute cfg save
revert
    |
    +-- Save point + timeout
start
  |
  v
modify
  |
  +---- abort ---> discard
  |
  +---- commit --> apply
Private Data Encryption
          |
          v
Master Encryption Password
          |
          +---- AES-128-CBC
          |
          v
Sensitive Data
TPM
 |
 +---- RSA-2048 Primary Key
 |
 v
Protects Master Encryption Password
| Goal | Command | 
|---|---|
| Current date | execute date | 
| Current time | execute time | 
| NTP status | diagnose sys ntp status | 
| PTP debug | diagnose debug application ptpd -1 | 
| Session table | diagnose sys session list | 
| TPM hardware info | diagnose hardware deviceinfo tpm | 
| TPM diagnostics | diagnose tpm | 
| Hardware test | diagnose hardware test info | 
| Transaction metadata | diagnose sys config-transaction show txn-meta | 
| Transaction info | diagnose sys config-transaction show txn-info | 
| Transaction entities | diagnose sys config-transaction show txn-entity | 
| Transaction locks | diagnose sys config-transaction show txn-lock | 
| Transaction CLI commands | diagnose sys config-transaction show txn-cli-commands | 
| Transaction memory context | diagnose sys config-transaction show mctx | 
- Change factory/default admin credential
- Create named administrator accounts
- Enable appropriate password policy
- Disable unnecessary password reuse
- Configure administrator idle timeout
- Configure login lockout
- Restrict management interfaces
- Prefer HTTPS/SSH
- Disable unnecessary Telnet/HTTP exposure
- Configure correct timezone
- Configure NTP
- Verify NTP synchronization
- Configure logging destination
- Verify certificate/time-dependent features
- Identify whether routing is symmetric
- Identify ECMP
- Identify SD-WAN
- Identify PBR
- Identify ADVPN
- Identify multiple ISP paths
- Determine whether return traffic can change interface
-  Evaluate auxiliary-session
-  Evaluate asymroute
- Verify NPU offload
-  Investigate dirty sessions
- Compare CPU/NPU behavior under load
- Use Workspace Mode for sensitive changes
- Understand transaction timeout
- Commit deliberately
- Abort failed changes
- Check transaction locks
-  Use cfg-save revert for risky remote changes
- Set a reasonable revert timeout
- Keep configuration backups
- Verify TPM support
- Understand Private Data Encryption
- Protect the master encryption key
- Document HA key requirements
- Test backup/restore procedures
- Understand TPM restore dependencies
- Remember: TPM is not disk encryption
                    FORTIGATE SYSTEM BASICS
                             |
       +---------------------+----------------------+
       |                     |                      |
    SECURITY                TIME                 SESSION
       |                     |                      |
 Password Policy            NTP                 Auxiliary
 Lockout                    PTP                 Asymroute
 Admin Timeout              TZ                  ECMP
 Management Ports           Logs                SD-WAN
       |                     |                      |
       +---------------------+----------------------+
                             |
                         PROCESSING
                             |
                     +-------+-------+
                     |               |
                    CPU             NPU
                     |               |
              Session Setup      Fast Path
              Policy/Routing     Offloading
                             |
                             v
                      CONFIGURATION
                             |
                +------------+------------+
                |                         |
          Save Modes                 Transactions
          automatic                 start/commit
          manual                    abort/timeout
          revert                    locks
                             |
                             v
                            TPM
                             |
                    Private Data Encryption
Auxiliary Session and Asymmetric Routing are the same thing.
❌ Wrong.
asymroute          = asymmetric routing behavior
auxiliary-session  = additional session handling
NPU handles the entire session from the first packet.
❌ Oversimplified.

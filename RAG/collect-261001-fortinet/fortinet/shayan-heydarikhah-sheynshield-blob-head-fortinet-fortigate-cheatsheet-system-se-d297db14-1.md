---
id: collect-261001-fortinet/fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-system-se-d297db14-1
title: "Configure administrator timeout"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: ["2026-08-29"]
keywords: []
source: docs/RAG/collect-261001-fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-system-se-d297db14.md
source_anchor: ""
source_lines: [1, 248]
sha256: 15e339f03d084300c0f8de80e12c4aa89916f6570f4375211a30b926a3d85154
---

# Configure administrator timeout

FortiOS 7.2.x | NSE4 / NSE7 Practical Reference
SheynShield | Engineering Secure Networks
A practical covering Administrator Security, NTP/PTP, Management Ports, Session Processing, Auxiliary Sessions, NPU Offloading, Link Monitor, Configuration Save Modes, TPM and Workspace Transactions.
On a factory-default FortiGate, the initial admin account behavior depends on the FortiOS/model generation.
Never leave the factory/default administrator credential unchanged.
Immediately after initial access:
- Change the administrator password.
- Create named administrator accounts.
- Avoid daily administration with the shared admin account.
- Restrict management access to trusted interfaces.
- Prefer HTTPS/SSH over insecure management protocols.
- Enable MFA where supported.
FortiOS password policy can control whether an administrator is allowed to reuse the previous password.
config system password-policy
    set status enable
    set reuse-password disable
end
| Setting | Result | 
|---|---|
| reuse-password enable | Password reuse is allowed | 
| reuse-password disable | New password must differ from old password | 
| min-change-characters > 0 | Requires a minimum number of characters to differ | 
| min-change-characters +reuse-password | min-change-characters takes precedence | 
Example:
config system password-policy
    set status enable
    set minimum-length 12
    set min-upper-case-letter 1
    set min-lower-case-letter 1
    set min-number 1
    set min-non-alphanumeric 1
    set reuse-password disable
end
Fortinet documents reuse-password as the control determining whether administrators may reuse the same password; min-change-characters overrides it when both are enabled.
Accurate system time is not optional on a production FortiGate.
Incorrect time can affect:
- Logs
- Event correlation
- Scheduled policies
- Certificates
- SSL-dependent features
- Authentication
- Security investigations
- Troubleshooting
- SIEM correlation
First-day rule: Configure the timezone and a reliable time source during initial deployment.
Fortinet specifically notes that accurate FortiOS system time is required for features such as scheduling, logging and SSL-dependent functionality.
execute date 2026-08-29
execute time 12:30:00
Check:
execute date
execute time
Example for Tehran:
config system global
    set timezone 41
end
FortiOS timezone 41 corresponds to GMT+03:30 Tehran.
config system global
    set timezone 41
    set dst disable
end
dst controls daylight-saving-time behavior.
Important: Do not blindly enable DST because an old configuration example contains set dst enable. Timezone/DST behavior should match the actual jurisdiction and current FortiOS/timezone rules.
FortiOS exposes dst as the CLI control for daylight-saving time.
FortiGate can synchronize its clock using:
- FortiGuard NTP
- Custom NTP servers
- NTP authentication
- NTP server mode
Basic configuration:
config system ntp
    set ntpsync enable
    set type custom
    set syncinterval 60
    config ntpserver
        edit 1
            set server ntp.example.com
            set ntpv3 disable
        next
    end
endset ntpv3 enable
means:
NTPv3
while:
set ntpv3 disable
means:
NTPv4
FortiOS documents ntpv3 as the switch for NTPv3; when disabled, NTPv4 is used.
FortiOS supports NTP authentication using:
- MD5
- SHA1
Example:
config system ntp
    set ntpsync enable
    set type custom
    config ntpserver
        edit 1
            set server ntp.example.com
            set authentication enable
            set key-id 123
            set key "SECRET"
            set ntpv3 enable
        next
    end
end
| NTP Version | Authentication | 
|---|---|
| NTPv3 | MD5 | 
| NTPv4 | SHA1 | 
Fortinet's CLI reference documents MD5/SHA1 authentication and the ntpv3 behavior.
Security note: NTP authentication is useful for trusted/internal NTP infrastructure. Do not expose unnecessary NTP services to the Internet.
config system ntp
    set syncinterval 1
end
The value is in minutes.
FortiOS supports an NTP synchronization interval from 1 to 1440 minutes.
FortiGate can also provide NTP service to downstream devices.
config system ntp
    set server-mode enable
    set interface "LAN"
end
Conceptually:
                    Internet / Internal NTP
                              |
                              v
                       +-------------+
                       |  FortiGate  |
                       | NTP Client  |
                       +-------------+
                              |
                       NTP Server Mode
                              |
             +----------------+----------------+
             |                |                |
             v                v                v
           PC-01            SW-01            Server
FortiGate relays NTP requests to its configured upstream NTP server when server mode is enabled.
PTP is designed for environments requiring much tighter time synchronization than ordinary NTP.
Typical use cases:
- Telecom
- Industrial systems
- Financial systems
- Time-sensitive applications
- Specialized data-center environments
FortiOS supports PTP configuration including:
- Interface
- Multicast/hybrid mode
- End-to-End delay
- Peer-to-Peer delay
Example:
config system ptp
    set status enable
    set interface "port1"
    set delay-mechanism E2E
end
The delay is measured across the complete path between the master and slave.
Simplified concept:
Master                         Slave
  |                              |
  | -------- Sync -------------->|
  |                              |
  | <------ Delay_Req -----------|
  |                              |
  | -------- Delay_Resp -------->|
The slave uses timestamps to estimate path delay.
Each link/path segment measures its own propagation delay through peer-delay exchanges.
Master ---- Switch ---- Switch ---- Slave
           P2P delay    P2P delay
E2E = End-to-End path delay
P2P = Peer-to-Peer link delay
Check system time:
execute date
execute time
Check NTP:
diagnose sys ntp status
Test UDP/123 reachability:
execute telnet 192.168.20.200 123
Note: TCP telnet to UDP/123 is not a definitive NTP test. It only proves whether something is listening/reachable using the tested transport. For real NTP validation, use FortiGate NTP diagnostics and packet capture.
PTP debugging:
diagnose debug application ptpd -1
On a Windows NTP server, verify:
netstat -nao
Check the Windows Time service:
services.msc
    Windows Time
Restart:
net stop w32time
net start w32time
Typical policy location:
Computer Configuration
 └── Administrative Templates
     └── System
         └── Windows Time Service
             ├── Global Configuration Settings
             └── Time Providers
Verify:
- Windows NTP Client enabled
- NTP Server configured
- Appropriate AnnounceFlags
- Windows Time service running
Example:
config system global
    set admin-port 80
    set admin-sport 443
    set admin-https-redirect enable
    set admin-ssh-port 2142
    set admin-telnet-port 2323
end
Prefer:
HTTPS
SSH
Avoid:
HTTP
TELNET
unless there is a specific operational requirement.
set admin-https-redirect enable
Concept:
HTTP :80
   |
   | redirect
   v
HTTPS :443
config system global
    set default-service-source-port 20-30
end
This affects the source-port range used for system-generated TCP/UDP service traffic.
Do not confuse this setting with the destination service port of firewall policies.
FortiOS controls administrator idle timeout through:
config system global
    set admintimeout 10
end
For administrative environments:
5–10 minutes
is generally preferable to leaving long idle sessions active.
FortiOS defines admintimeout as the number of minutes before an idle administrator session times out.
Example:
config system global
    set admin-lockout-threshold 2
    set admin-lockout-duration 60
end
Meaning:
2 failed attempts
        |
        v
Account locked
        |
        v
60 seconds

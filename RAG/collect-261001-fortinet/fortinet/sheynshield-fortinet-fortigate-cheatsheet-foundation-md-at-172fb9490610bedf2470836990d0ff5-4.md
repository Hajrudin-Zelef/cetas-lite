---
id: collect-261001-fortinet/fortinet/sheynshield-fortinet-fortigate-cheatsheet-foundation-md-at-172fb9490610bedf2470836990d0ff5-4
title: "sheynshield-fortinet-fortigate-cheatsheet-foundation-md-at-172fb9490610bedf2470836990d0ff5347fa74a0-"
domain: fortinet
role: reference
task: reference
actors: ["Apple", "Intel"]
dates: []
keywords: ["agent", "intel", "memory"]
source: docs/RAG/collect-261001-fortinet/sheynshield-fortinet-fortigate-cheatsheet-foundation-md-at-172fb9490610bedf2470836990d0ff5347fa74a0-.md
source_anchor: ""
source_lines: [1541, 1820]
sha256: 87eedb9af72eeaa9017f24708504c6def360fce8d8b963ea5bd2ab0d18b528e0
---

# sheynshield-fortinet-fortigate-cheatsheet-foundation-md-at-172fb9490610bedf2470836990d0ff5347fa74a0-

```
Forward Path
+
Return Path
```
```
☐ Dedicated Management VLAN
☐ OOB Management where possible
☐ HTTPS
☐ SSH
☐ Disable HTTP where unnecessary
☐ Disable Telnet
☐ Restrict Management Source IPs
☐ Trusted CA Certificate
☐ MFA
☐ Unique Administrator Accounts
☐ Least Privilege
☐ Appropriate Idle Timeout
☐ Logging
☐ Backup Configuration
☐ Secure Backup Storage
☐ Review Admin Accounts Regularly
```
For Internet-facing policies, evaluate:

```
☐ Correct Source Interface
☐ Correct Source Address
☐ Correct Destination Interface
☐ Correct Destination Address
☐ Correct Service
☐ Correct Schedule
☐ NAT
☐ Antivirus
☐ IPS
☐ Web Filter
☐ DNS Filter
☐ Application Control
☐ SSL/TLS Inspection
☐ File Filter
☐ DLP where required
☐ Logging
```
Use security controls according to:

```
Risk
+
Business Requirement
+
Performance
+
Compliance
```
Before upgrade:

```
☐ Confirm Current FortiOS Version
☐ Confirm Target Version
☐ Verify Upgrade Path
☐ Check Hardware Compatibility
☐ Read Release Notes
☐ Review Known Issues
☐ Backup Configuration
☐ Verify Backup
☐ Check HA Status
☐ Check FortiGuard Status
☐ Check Critical Services
☐ Define Maintenance Window
☐ Prepare Rollback Plan
```
After upgrade:

```
☐ Confirm Version
☐ Check Interfaces
☐ Check Routing
☐ Check Policies
☐ Check VPN
☐ Check HA
☐ Check SD-WAN
☐ Check Security Profiles
☐ Check Logs
☐ Check FortiGuard
☐ Test Critical Applications
```
```
☐ Link up?
☐ Interface errors?
☐ Speed/duplex?
☐ Cable/SFP?
```
```
☐ VLAN correct?
☐ Tagging correct?
☐ Switching correct?
☐ MAC learning correct?
```
```
☐ Correct IP?
☐ Correct subnet?
☐ Gateway reachable?
☐ Route exists?
```
```
☐ Correct policy?
☐ Policy order?
☐ Source address?
☐ Destination address?
☐ Service?
☐ Schedule?
```
```
☐ NAT required?
☐ Correct translated address?
☐ Return path?
```
```
☐ IPS blocking?
☐ AV blocking?
☐ Web Filter blocking?
☐ App Control blocking?
☐ SSL inspection issue?
☐ DNS filtering?
```
```
☐ Session created?
☐ Session state correct?
☐ Return traffic?
```
```
☐ CPU utilization?
☐ Memory?
☐ NP offload?
☐ CP offload?
☐ Resource exhaustion?
```
```
Route ≠ Permission
Policy = Authorization
NAT = Address Translation
IDS = Detection
IPS = Detection + Prevention
NGFW = Context-Aware Security
UTM = Multiple Integrated Security Functions
SPU = Security Processing Architecture
NP = Network Processing / Acceleration
CP = Content / Security Processing Acceleration
CPU = General-Purpose Processing
FortiManager = Centralized Management
FortiAnalyzer = Logging + Analytics + Reporting
FortiSIEM = SIEM + Security Operations
FortiWeb = WAF
FortiMail = Email Security
FortiNAC = Network Access Control
FortiAuthenticator = Identity + Authentication
FortiSandbox = Malware Analysis
FortiGuard = Security Intelligence + Security Services
FortiClient = Endpoint Agent
FortiClient EMS = Endpoint Management
FortiAP = Wireless Access Point
FortiWLC = Wireless LAN Controller
FortiADC = Application Delivery / Load Balancing
FortiToken = MFA / Authentication Token
```
```
                         FORTIGATE
                            │
       ┌────────────────────┼────────────────────┐
       │                    │                    │
    NETWORK              SECURITY            MANAGEMENT
       │                    │                    │
       ▼                    ▼                    ▼
   Routing              Firewall            FortiManager
   VLAN                  IPS                FortiAnalyzer
   VPN                   AV                 FortiSIEM
   SD-WAN                Web Filter          FortiClient EMS
   HA                    App Control
   VDOM                  SSL Inspection
                         Threat Intel
       │                    │
       └────────────┬───────┘
                    ▼
             Hardware Processing
                    │
             ┌──────┴──────┐
             ▼             ▼
            NP             CP
             │             │
       Network Path    Content/Security
       Acceleration     Acceleration
```
When troubleshooting any FortiGate problem, think:

```
                 TRAFFIC
                    │
                    ▼
              Physical Link
                    │
                    ▼
                Interface
                    │
                    ▼
               VLAN / L2
                    │
                    ▼
                 Routing
                    │
                    ▼
             Firewall Policy
                    │
                    ▼
                  NAT
                    │
                    ▼
           Security Inspection
                    │
          ┌─────────┼─────────┐
          ▼         ▼         ▼
         IPS        AV      App Control
          │         │         │
          └─────────┼─────────┘
                    ▼
              SSL Inspection
                    │
                    ▼
                 Session
                    │
                    ▼
              Hardware Path
                    │
                    ▼
                Forwarding
                    │
                    ▼
              Return Traffic
```
```
NSE 4
 │
 ├── Understand FortiGate
 ├── Configure FortiGate
 ├── Configure Security
 ├── Configure Networking
 ├── Configure VPN
 ├── Configure HA
 ├── Configure SD-WAN
 └── Perform Basic Troubleshooting
             │
             ▼
          NSE 7
             │
             ├── Understand Traffic Flow
             ├── Advanced Troubleshooting
             ├── Performance Analysis
             ├── Security Inspection
             ├── Hardware Acceleration
             ├── Enterprise Architecture
             ├── Advanced Threat Protection
             ├── Secure Access
             └── Cloud Security
```
At basic level:

```
"How do I configure it?"
```
At professional level:

```
"How does it work?"
```
At advanced level:

```
"Why does it behave differently from what I expect?"
```
At expert level:


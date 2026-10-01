---
id: collect-261001-fortinet/fortinet/sheynshield-fortinet-fortigate-cheatsheet-foundation-md-at-172fb9490610bedf2470836990d0ff5-1
title: "sheynshield-fortinet-fortigate-cheatsheet-foundation-md-at-172fb9490610bedf2470836990d0ff5347fa74a0-"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["throughput"]
source: docs/RAG/collect-261001-fortinet/sheynshield-fortinet-fortigate-cheatsheet-foundation-md-at-172fb9490610bedf2470836990d0ff5347fa74a0-.md
source_anchor: ""
source_lines: [1, 451]
sha256: 87d827c46bb8c95a4227e01054d90adcfe58569f62bb47d65e313a6ef95f309d
---

# sheynshield-fortinet-fortigate-cheatsheet-foundation-md-at-172fb9490610bedf2470836990d0ff5347fa74a0-

**SheynShield | Engineering Secure Networks**

**FortiGate • FortiOS • NGFW • UTM • Hardware Architecture • Management • Security • Networking • Troubleshooting • NSE 4 • NSE 7**


This is designed as a **Fortinet technical foundation and engineering reference**, with emphasis on the knowledge required to understand FortiGate at both **NSE 4 professional** and **NSE 7 advanced troubleshooting / solution** levels.

It is intentionally structured around four questions:

```
How does FortiGate work?
        ↓
How does FortiGate secure traffic?
        ↓
How does FortiGate process traffic?
        ↓
How do we troubleshoot and design it?
```
Traditional firewalls primarily make forwarding decisions using network and transport information:

```
Source IP
Destination IP
Protocol
Port
Session
```
Example:

```
Source      = 10.10.10.10
Destination = 8.8.8.8
Protocol    = TCP
Port        = 443
Action      = ACCEPT
```
The firewall knows that the traffic is using TCP port 443.

However:

```
TCP/443
   ≠
Guaranteed Application Identity
```
Port numbers alone do not provide complete application visibility.

UTM combines multiple security functions into a single security platform.

Typical capabilities include:

```
Firewall
+
Antivirus
+
IPS
+
Web Filtering
+
Application Control
+
Anti-Spam
+
VPN
```
```
Multiple Security Technologies
             ↓
      One Security Platform
```
The major advantage is integration between multiple security functions.

NGFW extends traditional firewall capabilities with deeper inspection and contextual security.

Typical capabilities:

```
Firewall
+
Application Control
+
IPS
+
Antivirus
+
Web Filtering
+
SSL/TLS Inspection
+
Identity Awareness
+
Threat Intelligence
+
Security Analytics
```
```
IP + Port
```
```
IP
+
Port
+
Application
+
User
+
Content
+
Threat
+
Context
```
IDS primarily detects suspicious activity.

```
Traffic
   ↓
Detection
   ↓
Alert
   ↓
Log
```
```
Detect
+
Alert
```
IPS adds active prevention.

```
Traffic
   ↓
Detection
   ↓
Analysis
   ↓
Decision
   ↓
Block / Drop / Reset
```
```
Detect
+
Prevent
```
```
IDS = Detect
IPS = Detect + Prevent
```
A simplified FortiGate security flow:

```
Incoming Traffic
       ↓
Interface
       ↓
Routing
       ↓
Firewall Policy
       ↓
Security Inspection
       ↓
Enforcement
       ↓
Logging
       ↓
Forwarding
```
Security inspection may involve:

```
Antivirus
IPS
Application Control
Web Filtering
DNS Filtering
SSL/TLS Inspection
File Filtering
DLP
Botnet Detection
Threat Intelligence
```
A simplified FortiGate architecture:

```
                         FortiGate
                            │
              ┌─────────────┴─────────────┐
              │                           │
             CPU                         SPU
                                          │
                              ┌───────────┴───────────┐
                              │                       │
                             NP                      CP
```
Where:

```
CPU
↓
General-Purpose Processing
NP
↓
Network Processing / Acceleration
CP
↓
Content / Security Processing Acceleration
SPU
↓
Fortinet Security Processing Architecture
```
Actual hardware architecture varies between FortiGate models and processor generations.

General-purpose processor.

Typical responsibilities may include:

- Management
- Control-plane processing
- Routing/control functions
- Processes that cannot be offloaded
- System services

**Security Processing Unit**

SPU is Fortinet's broader specialized hardware-processing architecture.

It can include:

```
SPU
 ├── NP
 └── CP
```
Optimized for network and packet processing.

Think:

```
Traffic
   ↓
NP
   ↓
High-Speed Network Processing
```
Potential acceleration areas depend on the platform and generation.

Designed to accelerate content/security-intensive operations.

Conceptually:

```
Traffic
   ↓
CP
   ↓
Security / Content Processing
```
Never assume:

```
NP = Everything Network Related
CP = Everything Security Related
```
Actual offloading depends on:

```
FortiGate Model
+
Processor Generation
+
FortiOS Version
+
Feature Configuration
+
Traffic Characteristics
+
Inspection Mode
```
FortiGate platforms are available across multiple performance classes.

A simplified historical/product-family model:

Examples include:

```
1000
2000
3000
5000
7000
```
Typical environments:

- Large enterprises
- Data centers
- Service providers
- High-throughput deployments
- Large security infrastructures

Examples include:

```
100
200
300
400
500
600
700
800
900
```
Typical environments:

- Enterprise branches
- Regional offices
- Campus networks
- Medium-sized organizations

Examples include:

```
30
40
50
60
80
```
Model suffixes such as:

```
F
G
```
represent different hardware generations and configurations.

**Important:** Model numbers should never be treated as a simple linear performance ranking. Always check the specific FortiGate data.


```
                         FORTINET
                            │
          ┌─────────────────┼─────────────────┐
          │                 │                 │
       NETWORK           SECURITY         OPERATIONS
          │                 │                 │
          ▼                 ▼                 ▼
      FortiGate          FortiWeb        FortiManager
      FortiSwitch        FortiMail       FortiAnalyzer
      FortiAP            FortiNAC        FortiSIEM
      FortiWLC           FortiDDoS       FortiClient EMS
      FortiADC            FortiAuth
```
FortiGuard provides Fortinet security intelligence and subscription services.

Conceptually:

```
                     FortiGuard
                         │
                         ▼
                Threat Intelligence
                         │
          ┌──────────────┼──────────────┐
          ▼              ▼              ▼
         AV             IPS          Web/DNS
       Updates        Updates       Categorization
          │              │              │
          └──────────────┼──────────────┘
                         ▼
                     FortiGate
```
Relevant intelligence may include:

- Antivirus updates
- IPS signatures
- Web categorization
- Application information
- Botnet intelligence
- DNS/security intelligence
- Other security databases

FortiGate can be understood using two major planes:

```
                 FortiGate
                    │
        ┌───────────┴───────────┐
        │                       │
        ▼                       ▼
    Data Plane              Management Plane
        │                       │
        ▼                       ▼
 Traffic Forwarding       Administration
 Firewall Policies        GUI / CLI
 Routing                  Monitoring
 Security Inspection      Configuration
```
```
Data Plane
    ≠
Management Plane
```
Many FortiGate platforms use:

```
192.168.1.99/24
```
as a factory-default management address.

Typical initial credentials:

```
Username: admin
Password: blank
```
Exact behavior varies by:

- Hardware model
- FortiOS release
- Factory-default configuration

```
Factory Default
      ↓
Change Password
      ↓
Enable MFA
      ↓
Restrict Management Access
```
Never leave a production FortiGate with default credentials.

Administrative access is configured per interface.

Example:

```
config system interface
    edit "port1"
        set ip 192.168.1.1 255.255.255.0
        set allowaccess https ping ssh
    next
end
```
Meaning:


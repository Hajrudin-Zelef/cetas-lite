---
id: collect-261001-fortinet/fortinet/sheynshield-fortinet-fortigate-cheatsheet-foundation-md-at-172fb9490610bedf2470836990d0ff5-2
title: "sheynshield-fortinet-fortigate-cheatsheet-foundation-md-at-172fb9490610bedf2470836990d0ff5347fa74a0-"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["agent", "incident"]
source: docs/RAG/collect-261001-fortinet/sheynshield-fortinet-fortigate-cheatsheet-foundation-md-at-172fb9490610bedf2470836990d0ff5347fa74a0-.md
source_anchor: ""
source_lines: [452, 1005]
sha256: 2282ea99a25909aee1c463800e9fcd269c32d98d242c0986c7090e17790c6456
---

# sheynshield-fortinet-fortigate-cheatsheet-foundation-md-at-172fb9490610bedf2470836990d0ff5347fa74a0-

```
HTTPS
  ↓
Web GUI
SSH
  ↓
CLI
PING
  ↓
ICMP
```
Possible management protocols depend on platform and configuration.

Recommended architecture:

```
                 Production Network
                        │
                        │
                     FortiGate
                        │
                        │
                 Management VLAN
                        │
                        ▼
                 Admin / Jump Host
```
Enterprise architecture:

```
Production Plane
       ≠
Management Plane
```
Preferred controls:

```
Management VLAN
+
OOB Management
+
Jump Server
+
Source IP Restriction
+
MFA
+
RBAC
+
Trusted Certificate
```
Administrative access is controlled through administrator accounts and administrator profiles.

Concept:

```
Administrator Account
          +
Administrator Profile
          ↓
Administrative Permissions
```
Provides broad administrative privileges.

Conceptually:

```
Global
+
VDOM
+
System
+
Network
+
Security
```
Can be limited according to:

- Permissions
- Features
- Read/write access
- VDOM scope
- Administrative role

Use:

```
Unique Admin Accounts
+
Least Privilege
+
MFA
```
instead of sharing one administrator account.

Concurrent administrative sessions determine whether the same administrative identity can maintain multiple active sessions.

Concept:

```
Administrator
    │
    ├── Session 1
    ├── Session 2
    └── Session 3
```
Prefer unique accounts for individual engineers.

Benefits:

```
Accountability
+
Auditability
+
Traceability
+
Incident Investigation
```
Production FortiGate management should use a trusted certificate.

Recommended architecture:

```
Enterprise CA
     │
     ▼
FortiGate Management Certificate
     │
     ▼
HTTPS
     │
     ▼
Administrator
```
Benefits:

- Trusted identity
- Reduced browser warnings
- Better management security
- Enterprise PKI integration

Changing the certificate or management port is not a substitute for access control.

VDOM means:

**Virtual Domain**

VDOMs allow a FortiGate to logically separate configurations and security domains.

Concept:

```
                 FortiGate
                     │
          ┌──────────┼──────────┐
          │          │          │
        VDOM-A     VDOM-B     VDOM-C
          │          │          │
       Network    Network    Network
       Policies   Policies   Policies
```
VDOMs can be used for:

- Multi-tenancy
- Organizational separation
- Administrative separation
- Network segmentation
- Resource isolation

Management scope can be:

```
Global
+
VDOM
```
depending on the administrator role.

One of the most important FortiGate concepts:

```
Routing
   ↓
WHERE?
```
versus:

```
Firewall Policy
   ↓
WHETHER?
```
Determines where the packet should go.

Determines whether the traffic is permitted.

Therefore:

```
Default Route
      ≠
Internet Permission
```
FortiGate evaluates policies according to policy lookup.

Conceptually:

```
Policy 1
   ↓
Policy 2
   ↓
Policy 3
   ↓
Policy 4
   ↓
...
   ↓
Implicit Deny
```
If no appropriate policy allows the traffic:

```
Traffic
   ↓
DENIED
```
```
No Matching Allow Policy
          ↓
        Deny
```
A policy can contain concepts such as:

```
Source Interface
Source Address
Destination Interface
Destination Address
Schedule
Service
Action
NAT
Authentication
Security Profiles
Logging
```
Mental model:

```
Traffic
   ↓
Interface
   ↓
Address
   ↓
Service
   ↓
Policy Match
   ↓
Security Controls
   ↓
Action
```
Typical LAN → Internet flow:

```
LAN Client
    ↓
FortiGate LAN
    ↓
Routing
    ↓
Firewall Policy
    ↓
NAT
    ↓
Security Inspection
    ↓
WAN
    ↓
Internet
```
Typical policy:

```
Source      = LAN
Destination = Internet
Service     = Required Services
Action      = ACCEPT
NAT         = Enabled where required
```
A static/default route does not automatically authorize traffic.

```
Route
 ↓
Path
Policy
 ↓
Permission
```
Security profiles provide additional inspection and enforcement.

Common security profiles/features include:

```
Antivirus
IPS
Web Filter
DNS Filter
Application Control
SSL/SSH Inspection
File Filter
DLP
Botnet Protection
```
A simplified policy architecture:

```
Firewall Policy
       │
       ├── Antivirus
       ├── IPS
       ├── Web Filter
       ├── Application Control
       ├── SSL Inspection
       └── Logging
```
For common Internet access scenarios:

```
Private IP
    ↓
FortiGate
    ↓
Source NAT
    ↓
Public IP
    ↓
Internet
```
Example concept:

```
192.168.10.10
      ↓
NAT
      ↓
203.0.113.10
```
NAT is a translation mechanism.

It is not itself a security policy.

```
NAT
 ≠
Firewall Authorization
```
FortiGate can support authentication workflows associated with firewall policies.

Conceptually:

```
Client
  ↓
Firewall Policy
  ↓
Authentication
  ↓
Redirect
  ↓
Authenticated Session
```
Example configuration structure:

```
config firewall policy
    edit <policy-id>
        set authentication-redirect-address <FQDN>
        set redirect-url <URL>
    next
end
```
Use the exact syntax and supported options for the installed FortiOS version.

FortiTelemetry provides communication between FortiClient and FortiGate in supported deployments.

Concept:

```
FortiClient
     │
     │ Telemetry
     ▼
FortiGate
```
A commonly associated port is:

```
TCP/8013
```
Exact communication behavior depends on the FortiOS and FortiClient versions and architecture.

CAPWAP is used for centralized wireless AP communication.

Concept:

```
FortiGate / Controller
          │
          │ CAPWAP
          ▼
       FortiAP
```
Common CAPWAP ports:

```
UDP/5246 → Control
UDP/5247 → Data
```
```
5246 = Control
5247 = Data
```
FortiManager provides centralized management.

```
                    FortiManager
                         │
          ┌──────────────┼──────────────┐
          ▼              ▼              ▼
      FortiGate       FortiGate       FortiGate
```
Typical functions:

- Centralized device management
- Policy management
- Configuration management
- Revision control
- Administrative workflows
- Large-scale deployment

```
FortiManager
     =
Centralized Management
```
FortiAnalyzer focuses on:

```
Logging
+
Analytics
+
Reporting
+
Investigation
```
Typical flow:

```
FortiGate
    │
    │ Logs
    ▼
FortiAnalyzer
    │
    ├── Search
    ├── Analytics
    ├── Reports
    └── Investigation
```
```
FortiAnalyzer
     =
Logs + Analytics + Reporting
```
Endpoint agent.

Capabilities can include:

- VPN
- Endpoint security
- ZTNA
- Telemetry
- Endpoint posture/security capabilities

Centralized management for FortiClient endpoints.

```
                    FortiClient EMS
                           │
             ┌─────────────┼─────────────┐
             ▼             ▼             ▼
        Endpoint 1    Endpoint 2    Endpoint 3
        FortiClient   FortiClient   FortiClient
```
```
FortiClient
    =
Endpoint Agent
FortiClient EMS
    =
Endpoint Management
```
FortiSIEM provides SIEM and security operations capabilities.

Concept:

```
Servers
Network Devices
Security Devices
Applications
Endpoints
      │
      ▼
  FortiSIEM
      │
      ├── Events
      ├── Correlation
      ├── Monitoring
      └── Security Operations
```
```
FortiSIEM
    =
SIEM / Security Operations
```
FortiWeb is a:

**Web Application Firewall (WAF)**

Concept:

```
Internet
   ↓
FortiWeb
   ↓
Web Application
```
Protection can focus on:

- Web application attacks
- HTTP/HTTPS traffic
- Application-layer threats
- OWASP-related attack patterns
- API/application protection


---
id: collect-261001-fortinet/fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-user-auth-e82ca855-4
title: "shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-user-auth-e82ca855"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-user-auth-e82ca855.md
source_anchor: ""
source_lines: [901, 1070]
sha256: c3bed88edecf73d899021b2688290814442de77a63f7168b93426751308ef150
---

# shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-user-auth-e82ca855

```
diagnose user-device-store user-count query \
"CN=callcenter,OU=information-tech,OU=publish-users,DC=test,DC=com"
```
`diagnose wireless-controller wlac-d sta online````
                         RADIUS
                           │
             ┌─────────────┼─────────────┐
             │             │             │
            1812          1813          VSA
             │             │             │
      Authentication   Accounting    Vendor Data
             │             │             │
             └─────────────┼─────────────┘
                           │
                           ▼
                       FortiGate
                           │
              ┌────────────┼────────────┐
              │            │            │
           Groups         RSSO         CoA
              │            │            │
              │       User ↔ IP     Authorization
              │
              ▼
       Firewall Policy
              │
              ▼
           Session
              │
              ▼
      Dynamic Shaping
```
```
                    RADIUS
                       │
       ┌───────────────┼────────────────┐
       │               │                │
    1812/AAA         1813             VSA
       │               │                │
 Authentication    Accounting      Vendor Data
       │               │                │
       └───────────────┼────────────────┘
                       │
                       ▼
                    NPS/AD
                       │
                       ▼
                 Group Mapping
                       │
                       ▼
                  FortiGate
                       │
          ┌────────────┼────────────┐
          │            │            │
        Policy        RSSO         CoA
          │            │            │
          ▼            ▼            ▼
       Access       User/IP      Dynamic Auth
          │
          ▼
       Session
          │
          ▼
   Dynamic Shaping
```
- RADIUS server reachable
- UDP/1812 reachable
- Shared secret correct
- NAS identity correct
- Authentication method correct
- NPS policy matches
- AD authentication succeeds

- Username format verified
- Case sensitivity understood
- Required AVP/VSA returned
- Group mapping verified
- FortiGate group matches returned identity

- UDP/1813 reachable
- Accounting enabled
- Start packet received
- Interim update received when required
- Stop packet received

- RSSO configured
- Accounting working
- User/IP association learned
- RSSO group configured
- Policy references RSSO identity

- CoA supported and configured
- Required attributes returned
- Session can receive updates
- Authorization change is verified

- Dynamic shaping enabled
- Correct user/session identified
- Required RADIUS attributes present
- Correct firewall policy matched
- Shaper statistics verified

**FortiGate RADIUS is much more than username/password authentication.**


The enterprise identity chain is:

```
Active Directory
      ↓
NPS
      ↓
RADIUS Authentication
      ↓
AVP / VSA
      ↓
FortiGate Group
      ↓
Firewall Policy
      ↓
Session
      ↓
Accounting / RSSO
      ↓
CoA / Dynamic Authorization
      ↓
Dynamic Shaping
```
```
1. 1812 = Authentication
2. 1813 = Accounting
3. AVP = Standard RADIUS Attribute
4. VSA = Vendor-Specific Attribute
5. RSSO = RADIUS-based user/session awareness
6. CoA = Change of Authorization
7. Authentication Success ≠ Policy Match Success
```
```
If RADIUS authentication succeeds
but the user receives the wrong access:
DO NOT STOP AT RADIUS.
Trace:
AD
 ↓
NPS Policy
 ↓
RADIUS Response
 ↓
AVP / VSA
 ↓
FortiGate Group
 ↓
Firewall Policy
 ↓
Session
 ↓
Accounting / RSSO
 ↓
Shaper / Authorization
```
- 
YouTube — SheynShield 
  - Fortinet NSE content
  - FortiGate troubleshooting
  - Network Security Engineering

**SheynShield | Engineering Secure Networks**

`FortiGate` · `RADIUS` · `NPS` · `Active Directory` · `RSSO` · `VSA` · `AVP` · `CoA` · `Dynamic Shaping` · `Fortinet` · `Network Security` · `NSE4` · `NSE7`

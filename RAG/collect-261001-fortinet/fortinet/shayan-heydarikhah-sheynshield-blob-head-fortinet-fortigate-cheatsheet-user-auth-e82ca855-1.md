---
id: collect-261001-fortinet/fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-user-auth-e82ca855-1
title: "shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-user-auth-e82ca855"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-user-auth-e82ca855.md
source_anchor: ""
source_lines: [1, 298]
sha256: ea288b5a71c536d03588f875f48ac927ba1dd9d43dc7053eb1f44032951da524
---

# shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-user-auth-e82ca855

**SheynShield | Security & Design Knowledge Base**
**FortiGate / FortiOS | NSE4 → NSE7**

Practical **RADIUS Server checklist** covering **RADIUS authentication, NPS, Active Directory, VSA, AVP, RSSO, RADIUS Accounting, CoA, group mapping, dynamic shaping, wireless attributes, troubleshooting, and FortiGate CLI**.


**RADIUS — Remote Authentication Dial-In User Service** provides centralized:

- Authentication
- Authorization
- Accounting

Typical FortiGate architecture:

```
                    ┌─────────────────────┐
                    │ Active Directory    │
                    │       / LDAP        │
                    └──────────┬──────────┘
                               │
                               ▼
                    ┌─────────────────────┐
                    │        NPS          │
                    │      RADIUS         │
                    └──────────┬──────────┘
                               │
                 ┌─────────────┴─────────────┐
                 │                           │
              UDP 1812                   UDP 1813
            Authentication               Accounting
                 │                           │
                 └─────────────┬─────────────┘
                               ▼
                    ┌─────────────────────┐
                    │     FortiGate       │
                    │                     │
                    │ Authentication      │
                    │ Group Mapping       │
                    │ RSSO                │
                    │ Dynamic Shaping      │
                    └──────────┬──────────┘
                               │
                               ▼
                       Firewall Policy
                               │
                               ▼
                            Access
```
- RADIUS/NPS server is available
- Active Directory integration is working
- FortiGate is defined as a RADIUS client/NAS
- Shared secret is configured consistently
- Authentication port is reachable
- Accounting port is reachable when required
- NPS policies are configured
- FortiGate user group is configured
- Firewall policy references the correct group

| Function | UDP Port | 
|---|---|
| Authentication / Authorization | `1812` | 
| Accounting | `1813` | 
| Legacy Authentication | `1645` | 
| Legacy Accounting | `1646` | 

- UDP/1812 allowed from FortiGate to RADIUS
- UDP/1813 allowed when accounting is required
- Legacy ports are used only when required
- Routing between FortiGate and RADIUS is correct
- Intermediate firewalls permit required traffic

```
1812 → Authentication
1813 → Accounting
1645 → Legacy Authentication
1646 → Legacy Accounting
```
**NSE Gotcha:** Authentication can work perfectly while accounting-dependent features fail because UDP/1813 is unavailable.

⚠️ 

```
User & Authentication
└── RADIUS Servers
    └── Create New
```
Example:

```
Name:              rad-1
Server IP:         192.168.20.200
Authentication:    MSCHAPv2
NAS IP:            192.168.20.254
Auth Port:         1812
Accounting Port:   1813
```
- RADIUS object created
- Correct server IP configured
- Correct authentication method selected
- Shared secret configured
- NAS IP configured correctly when required
- Authentication port verified
- Accounting port verified
- NPS recognizes FortiGate as a RADIUS client

```
config user radius
    edit "rad-1"
        set server "192.168.20.200"
        set secret <RADIUS_SECRET>
    next
end
```
🔐 Never commit a real RADIUS shared secret to GitHub.


Create a FortiGate remote authentication group:

```
User & Authentication
└── User Groups
    └── Create New
```
Example:

```
Group Name:      rad-g-1
Remote Server:   rad-1
```
- RADIUS server selected
- Correct remote group configuration
- Returned RADIUS attributes understood
- Group referenced by firewall policy
- NPS group membership verified

```
Active Directory
      ↓
AD Group
      ↓
NPS Policy
      ↓
RADIUS Response
      ↓
FortiGate Group
      ↓
Firewall Policy
      ↓
Access
```
💡 **Critical troubleshooting rule:**
`Authentication = SUCCESS` does not necessarily mean `Firewall Policy Match = SUCCESS`.


Example:

```
config user radius
    edit "rad-1"
        set password-encoding auto
        set username-case-sensitive enable
    next
end
```
- Password encoding matches the deployment
- Username format is consistent
- Username case behavior is understood
- AD username format is verified
- NPS username handling is verified

```
User01
user01
USER01
```
Depending on the FortiOS/RADIUS/backend behavior, these may not be treated identically.

```
FortiGate
   ↓
RADIUS
   ↓
NPS
   ↓
Active Directory
```
Keep username normalization and case handling consistent across the entire authentication chain.

Standard RADIUS attributes follow:

```
Attribute = Value
```
Common attributes:

| Attribute | Code | Purpose | 
|---|---|---|
| User-Name | `1` | Username | 
| NAS-IP-Address | `4` | NAS address | 
| Framed-IP-Address | `8` | Client IP | 
| Class | `25` | Session/accounting correlation | 
| Vendor-Specific | `26` | Vendor attribute container | 
| NAS-Identifier | `32` | NAS identifier | 
| Acct-Input-Octets | `42` | Received octets | 
| Acct-Output-Octets | `43` | Transmitted octets | 
| Acct-Session-Id | `44` | Session identifier | 
| Event-Timestamp | `55` | Event timestamp | 

- Attribute is supported
- Attribute code is correct
- Attribute value is correct
- FortiGate interprets the attribute as expected
- NPS policy returns the required attribute

**VSA = Vendor-Specific Attribute**

Fortinet-specific RADIUS information can be carried using the RADIUS Vendor-Specific Attribute mechanism.

```
RADIUS Attribute 26
        │
        ▼
Fortinet VSA
        │
        ├── Group
        ├── Client IP
        ├── VDOM
        ├── IPv6 information
        ├── Interface
        └── Access Profile
```
```
Vendor:
Fortinet
Vendor-ID:
12356
```
| VSA | Attribute | Purpose | 
|---|---|---|
| `1` | `fortinet-group-name` | Remote group mapping | 
| `2` | `fortinet-client-ip-address` | Client IP | 
| `3` | `fortinet-vdom-name` | VDOM information | 
| `4` | `fortinet-client-ipv6-address` | IPv6 client information | 
| `5` | `fortinet-interface-name` | Interface information | 
| `6` | `fortinet-access-profile` | Administrative access profile | 

- Fortinet vendor ID is correct
- Correct VSA is returned by RADIUS
- NPS vendor-specific attribute is configured correctly
- Attribute value matches FortiGate expectations
- Returned VSA is visible during troubleshooting
- Group mapping uses the correct returned value

A common enterprise design is:

```
Active Directory Group
          ↓
      NPS Policy
          ↓
    RADIUS Response
          ↓
    Fortinet VSA
          ↓
     FortiGate Group
          ↓
    Firewall Policy
```
Example:

```
AD User:
    alice
AD Group:
    IT
NPS:
    Match IT group
RADIUS:
    Fortinet-Group-Name = IT
FortiGate:
    Remote Group = IT
```
- User belongs to expected AD group
- NPS policy matches the user
- NPS policy returns expected attribute
- Fortinet VSA is correctly configured
- Returned group name is correct
- FortiGate remote group matches returned value
- Firewall policy references the expected group

If:

```
Authentication = SUCCESS
Policy Match    = FAILURE
```
check:


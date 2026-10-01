---
id: collect-261001-fortinet/fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-user-auth-e82ca855-2
title: "shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-user-auth-e82ca855"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-user-auth-e82ca855.md
source_anchor: ""
source_lines: [299, 673]
sha256: 7e1d00a1ce1a42dc8a0bfbc5f32362806d3d3b0f9a409b4588803086421f606e
---

# shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-user-auth-e82ca855

```
NPS Policy
   ↓
RADIUS Attributes
   ↓
VSA / Filter-Id
   ↓
FortiGate Group
   ↓
Firewall Policy
```
before assuming the RADIUS password authentication itself is broken.

RADIUS Accounting provides session information such as:

- Session start
- Interim updates
- Session stop
- User/session correlation
- Usage information
- RSSO information
- Dynamic policy context

```
FortiGate
   │
   │ Accounting-Request
   ▼
RADIUS Server
   │
   ├── Start
   ├── Interim-Update
   └── Stop
```
- UDP/1813 reachable
- Accounting server configured
- Accounting secret matches
- Start packets received
- Interim updates configured where required
- Stop packets received
- Session IDs are consistent
- Accounting data is visible on the RADIUS server

**RSSO = RADIUS Single Sign-On**

RSSO allows FortiGate to obtain user/session information associated with RADIUS authentication and use that identity for policy enforcement.

```
Security Fabric
└── External Connectors
    └── RADIUS Single Sign-On
```
Example:

```
Name:
    rsso-agent
```
Then create/use a corresponding FortiGate user group.

```
User & Authentication
└── User Groups
    └── rsso-f
```
```
User
  ↓
Network Access Device
  ↓
RADIUS Authentication
  ↓
RADIUS Server
  ↓
Accounting / Session Information
  ↓
FortiGate RSSO
  ↓
User ↔ IP Association
  ↓
FortiGate User Group
  ↓
Firewall Policy
```
- RSSO connector configured
- RADIUS authentication works
- RADIUS accounting works
- User/session information reaches FortiGate
- User-to-IP association is learned
- RSSO group is configured
- Firewall policy references RSSO identity

| Feature | RSSO | FSSO Polling | FSSO Collector | 
|---|---|---|---|
| Primary source | RADIUS | AD | AD/DC | 
| Authentication awareness | RADIUS-based | Polling | Logon events | 
| Main identity source | RADIUS session | AD | AD | 
| Best fit | RADIUS environments | Simple AD | Enterprise AD | 
| Session awareness | RADIUS-based | AD polling | AD logon events | 

```
RADIUS infrastructure
        ↓
       RSSO
Simple AD environment
        ↓
   FSSO Polling
Enterprise AD / multiple DCs
        ↓
 FSSO Collector / DC Agent
```
**CoA = Change of Authorization**

CoA allows the RADIUS infrastructure to request changes to an existing authenticated session.

```
RADIUS
   │
   │ CoA
   ▼
FortiGate
   │
   ├── Update authorization
   ├── Update session
   └── Apply supported attributes
```
- Dynamic authorization
- User policy changes
- Bandwidth changes
- Wireless environments
- Session updates
- Post-authentication authorization

- CoA supported by target FortiOS release
- RADIUS server supports required CoA behavior
- FortiGate CoA configuration verified
- Required network path is available
- Shared secret/configuration is correct
- Session is eligible for the requested update

RADIUS attributes can participate in dynamic traffic-shaping designs.

Example:

```
config user radius
    edit "rad-1"
        set radius-coa enable
        set radius-all-server enable
    next
end
```
⚠️ Exact CLI availability and behavior can vary by FortiOS release. Validate commands against the target FortiOS version.

Example:

```
config firewall policy
    edit 1
        set dynamic-shaping enable
    next
end
```
```
RADIUS User
     ↓
Identity
     ↓
RADIUS Attributes
     ↓
FortiGate Session
     ↓
Firewall Policy
     ↓
Dynamic Shaper
     ↓
User Traffic
```
Example:

```
User A → 20 Mbps
User B → 10 Mbps
User C →  5 Mbps
```
- User identity is correctly learned
- Required RADIUS attributes are returned
- Accounting is working when required
- Policy supports dynamic shaping
- Dynamic shaping is enabled
- User/session is associated correctly
- Shaper statistics confirm enforcement

Wireless environments may use additional attributes to carry information related to:

- SSID
- Wireless controller/device
- Wireless termination point
- Association time

Fortinet-specific wireless information may include:

| Code | Concept | 
|---|---|
| `7` | Fortinet SSID | 
| `23` | Wireless device/controller | 
| `24` | Wireless termination point | 
| `25` | Association time | 

- RADIUS authentication works
- Accounting is received
- Wireless client association is correct
- AP/controller information is correct
- Session state is synchronized
- Roaming behavior is validated
- Dynamic shaping is verified

Termination behavior affects how session state changes are handled.

Conceptual values:

| Value | Concept | 
|---|---|
| `0` | Default/normal behavior | 
| `1` | RADIUS request/session verification behavior | 
| `2` | Resource/session management and cleanup | 

**FortiOS-version and deployment dependent**. Validate the exact behavior against the target release and RADIUS integration.

⚠️ Treat these values as

- Required termination behavior is identified
- FortiOS version is verified
- RADIUS server behavior is verified
- Wireless session behavior is tested
- Stale-session scenarios are tested

RADIUS can return attributes used by FortiGate for identity/group decisions.

Example:

```
RADIUS
   │
   └── Filter-Id = IT
             │
             ▼
      FortiGate Group
             │
             ▼
      Firewall Policy
```
- Returned attribute is expected
- Attribute value is correct
- FortiGate group expects the same value
- NPS policy returns the attribute
- Firewall policy references the correct group

```
AD Group
   ↓
NPS Policy
   ↓
RADIUS Attribute
   ↓
FortiGate Group
   ↓
Firewall Policy
```
Any mismatch can cause policy selection failure.

RSA SecurID provides OTP-based authentication.

Conceptual architecture:

```
User
 │
 ├── Username
 └── OTP
      │
      ▼
FortiGate
      │
      ▼
RADIUS
      │
      ▼
RSA Authentication Infrastructure
```
`diagnose test authserver radius <server> <username> <password>`
- OTP integration is validated
- RADIUS authentication works
- Token/OTP policy is correct
- No OTP secrets are stored in documentation
- Production credentials are never committed to GitHub

🔐 **Never publish real passwords, OTP seeds, RADIUS secrets, certificates containing private keys, or production credentials.**


- FortiGate can reach RADIUS server
- Correct routing exists
- UDP/1812 is allowed
- UDP/1813 is allowed when required
- No intermediate firewall blocks RADIUS

- RADIUS server is online
- FortiGate is configured as a RADIUS client
- Shared secret matches
- NAS IP is correct
- Authentication method matches
- NPS policy is enabled

Run:

`diagnose test authserver radius <server> <username> <password>`
Validate:

- Authentication request reaches server
- Server returns Access-Accept
- No Access-Reject
- No unexpected Access-Challenge
- Authentication method is compatible

If authentication succeeds:

```
Authentication
      ↓
      ?
FortiGate Group
```
Check:

- NPS policy
- AD group membership
- VSA
- Filter-Id
- Returned attributes
- FortiGate remote group
- Firewall policy

If RSSO fails:

- Authentication works
- Accounting works
- UDP/1813 is reachable
- RSSO connector is enabled
- Session information is received
- User ↔ IP association exists
- RSSO group is correct

Check:

`diagnose firewall shaper dynamic-shaper stats`
Then:

`diagnose firewall shaper dynamic-shaper list`
Specific client:

`diagnose firewall shaper dynamic-shaper list ip 192.168.20.20`
Validate:

- Correct user identity
- Correct policy
- Required RADIUS attributes
- Correct shaper
- Active session
- Shaper statistics

Inspect:

`diagnose sys session list`
Then correlate:


---
id: collect-261001-fortinet/fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-user-auth-e82ca855-3
title: "shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-user-auth-e82ca855"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-user-auth-e82ca855.md
source_anchor: ""
source_lines: [674, 900]
sha256: d899b3f246a930f9c59c00b20d881e1fec4f5da4ca686da0af03734f22810211
---

# shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-user-auth-e82ca855

```
User
 ↓
Authentication
 ↓
Group
 ↓
Policy
 ↓
Session
 ↓
Shaper
```
| Symptom | Primary Checks | 
|---|---|
| Authentication fails | Server IP, secret, port, authentication method | 
| RADIUS timeout | Routing, firewall, UDP/1812 | 
| Accounting missing | UDP/1813, accounting configuration | 
| User authenticates but wrong policy matches | Group mapping, VSA, Filter-Id | 
| RSSO user not detected | Accounting, RSSO configuration, user/IP association | 
| Dynamic shaping fails | RADIUS attributes, accounting, policy, shaper | 
| Wrong group | NPS policy, VSA, Filter-Id | 
| Wireless session remains stale | Accounting, termination behavior | 
| CoA fails | CoA configuration, reachability, server support | 
| Username mismatch | Case sensitivity, username format | 
| MSCHAPv2 fails | NPS policy, AD integration, authentication method | 
| Policy does not match after successful login | Group mapping and returned attributes | 

```
                    Active Directory
                           │
                           ▼
                    ┌─────────────┐
                    │     NPS     │
                    │   RADIUS    │
                    └──────┬──────┘
                           │
              ┌────────────┴────────────┐
              │                         │
          UDP 1812                  UDP 1813
       Authentication             Accounting
              │                         │
              └────────────┬────────────┘
                           ▼
                    ┌─────────────┐
                    │  FortiGate  │
                    └──────┬──────┘
                           │
          ┌────────────────┼─────────────────┐
          │                │                 │
          ▼                ▼                 ▼
       Groups             RSSO        Dynamic Shaping
          │                │                 │
          └────────────────┼─────────────────┘
                           ▼
                    Firewall Policy
                           │
                           ▼
                          User
```
- AD user exists
- AD group membership is correct
- NPS receives request
- NPS policy matches
- Access-Accept is returned
- Required AVP/VSA is returned
- FortiGate recognizes the user
- Correct FortiGate group is selected
- Correct firewall policy matches
- Accounting starts
- RSSO learns the session when required
- Dynamic shaping applies when configured

```
Authentication
    ↓
Who are you?
Authorization
    ↓
What can you access?
```
Therefore:

```
RADIUS Authentication = SUCCESS
```
does **not** guarantee:

```
Firewall Policy Match = SUCCESS
```
```
1812 → Authentication
1813 → Accounting
```
If authentication works but RSSO/accounting-dependent behavior fails:

```
Check UDP/1813
```
```
AVP
 ↓
Standard RADIUS Attribute
VSA
 ↓
Vendor-Specific Attribute
```
Fortinet-specific information commonly uses the vendor-specific mechanism.

Always inspect:

```
AD
 ↓
NPS
 ↓
RADIUS Attribute
 ↓
FortiGate Group
 ↓
Firewall Policy
```
RSSO is not simply:

```
Username + Password
```
Think:

```
RADIUS Authentication
       +
Accounting / Session Information
       ↓
FortiGate User ↔ IP Awareness
```
If authentication works but shaping does not:

```
RADIUS Attributes
        ↓
Accounting
        ↓
User / Session
        ↓
Firewall Policy
        ↓
Dynamic Shaper
```
Trace the complete chain.

- RADIUS server IP is correct
- Shared secret is strong
- Shared secret matches on both sides
- UDP/1812 is restricted to trusted clients
- UDP/1813 is restricted to trusted clients
- NAS IP is correct
- NPS policies are least-privilege
- Authentication method is appropriate
- AD integration is functioning
- Authentication logging is enabled

- RADIUS server object configured
- Correct RADIUS group configured
- Firewall policy references correct group
- Required VSA/AVP mappings validated
- Accounting configured where required
- RSSO configured where required
- CoA configured where required
- Dynamic shaping configured where required
- Authentication troubleshooting logs available

- Never publish RADIUS shared secrets
- Never publish production passwords
- Never publish OTP seeds
- Never publish private keys
- Restrict RADIUS clients to trusted NAS devices
- Monitor authentication failures
- Monitor abnormal accounting traffic
- Review NPS policies regularly
- Use secure management access
- Validate supported secure RADIUS options for the target environment

```
config user radius
    edit "rad-1"
        set server "192.168.20.200"
        set secret <RADIUS_SECRET>
    next
end
```
```
config user radius
    edit "rad-1"
        set password-encoding auto
        set username-case-sensitive enable
    next
end
```
```
config user radius
    edit "rad-1"
        config accounting-server
            edit 1
                set status enable
                set server 192.168.20.200
                set port 1813
                set secret <RADIUS_SECRET>
            next
        end
    next
end
```
```
config user radius
    edit "rad-1"
        set radius-coa enable
        set radius-all-server enable
    next
end
```
⚠️ Verify exact syntax against the target FortiOS release.

`diagnose test authserver radius <server> <username> <password>``diagnose firewall auth list``diagnose sys session list``diagnose firewall shaper dynamic-shaper stats``diagnose firewall shaper dynamic-shaper list``diagnose firewall shaper dynamic-shaper list ip 192.168.20.20``diagnose user-device-store user-count list 1`
Example:


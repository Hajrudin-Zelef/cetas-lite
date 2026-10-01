---
id: collect-261001-fortinet/fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-network-d-7c4e7ff6-3
title: "shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-network-d-7c4e7ff6"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-network-d-7c4e7ff6.md
source_anchor: ""
source_lines: [463, 704]
sha256: 9688ad9016073e06e82afa7a9ab0fcd3082256e9ddd27a1ed0fdeef8edbae666
---

# shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-network-d-7c4e7ff6

| Guest account still works after expiration | Guest lifecycle/expiration configuration | 
| Wrong firewall policy is applied | User/group mapping, policy order | 
| Authentication server cannot be reached | Routing, DNS, firewall policy, server availability | 
- Guest network is isolated from corporate networks.
- Guest VLAN is dedicated.
- Management interfaces are inaccessible.
- Internal DNS is not unnecessarily exposed.
- Internal services are explicitly denied.
- Guest-to-guest communication is restricted where required.
- Strong credentials are used.
- Temporary accounts have expiration.
- Authentication timeout is configured.
- Unused accounts are disabled.
- Centralized authentication is used where appropriate.
- Authentication groups follow least privilege.
- HTTPS portal is enabled where appropriate.
- Trusted certificate is deployed.
- Certificate expiration is monitored.
- Certificate hostname/SAN is correct.
- Certificate chain is valid.
- Authentication events are logged.
- Failed logins are monitored.
- Guest access is auditable.
- Central logging is configured where required.
- Suspicious authentication behavior is investigated.
- Guest users receive only required access.
- Guest users cannot reach administrative interfaces.
- Authentication groups map only to intended policies.
- Temporary access automatically expires.
- Guest access is reviewed periodically.
Client
  ↓
Interface / VLAN
  ↓
Firewall Policy
  ↓
Authentication
  ↓
User / Group
  ↓
Access Decision
  ↓
Internet / Resource
- Understand where Captive Portal fits in the traffic flow.
- Understand the relationship between User → Group → Policy.
- Understand Local vs LDAP vs RADIUS authentication.
- Understand Guest Management.
- Understand authentication timeout behavior.
- Understand Guest VLAN isolation.
- Understand HTTPS/certificate requirements.
- Know the key authentication troubleshooting commands.
Captive Portal
       │
       ▼
Authentication Interface
       │
       ▼
Local / LDAP / RADIUS
       │
       ▼
Identity Verification
Captive Portal provides the access/authentication mechanism.
LDAP/RADIUS provide external identity/authentication services.
Authentication answers:
Who are you?
Authorization answers:
What are you allowed to access?
Authentication
      │
      ▼
Identity
      │
      ▼
User Group
      │
      ▼
Firewall Policy
      │
      ▼
Authorization
Guest
  │
  ├── Internet       ✅
  ├── Internal LAN   ❌
  ├── Servers        ❌
  └── Management     ❌
idle-timeout
    ↓
Inactivity-based expiration
hard-timeout
    ↓
Fixed authentication lifetime
session-timeout
    ↓
New sessions denied after timeout
When Captive Portal fails, do not immediately blame authentication.
Use:
1. Client
   ↓
2. DHCP
   ↓
3. DNS
   ↓
4. VLAN / Interface
   ↓
5. Firewall Policy
   ↓
6. Authentication
   ↓
7. LDAP / RADIUS
   ↓
8. NAT
   ↓
9. Routing
   ↓
10. Certificate
NSE troubleshooting principle: Start with the lower layers and traffic path before assuming the authentication backend is the problem.
diagnose firewall auth listdiagnose firewall auth mac listdiagnose test authserver ldap <server> <username> <password>config user setting
    set auth-timeout-type idle-timeout
    set auth-timeout 5
endconfig user ldap
    edit "LDAP"
        set server <LDAP-SERVER>
        set cnid "sAMAccountName"
        set dn "dc=example,dc=local"
    next
endconfig firewall policy
    edit 10
        set name "Guest-Captive-Portal"
        set srcintf "guest"
        set dstintf "wan"
        set srcaddr "all"
        set dstaddr "all"
        set schedule "always"
        set service "ALL"
        set action "accept"
        set nat enable
    next
end
⚠️ Version Note: Exact CLI options can differ between FortiOS releases, FortiGate models and deployment modes. Always validate the command tree on the target FortiOS version.
- Guest VLAN configured.
- DHCP working.
- DNS working.
- Routing verified.
- Internet connectivity verified.
- Guest-to-internal connectivity tested and denied.
- Captive Portal enabled.
- Authentication flow tested.
- HTTP/HTTPS behavior tested.
- HTTPS certificate validated.
- Mobile client behavior tested.
- Local/LDAP/RADIUS source tested.
- User group tested.
- Authentication success tested.
- Authentication failure tested.
- Timeout tested.
- Guest expiration tested.
- Re-authentication tested.
- Guest → Internal denied.
- Guest → Management denied.
- Guest → Servers denied.
- Required authentication services allowed.
- Least-privilege access implemented.
- Guest-to-guest access reviewed.
- Authentication logs available.
- Failed authentication monitored.
- Guest sessions visible.
- Central logging configured where required.
- Authentication anomalies can be investigated.
- Policy order reviewed.
- NAT verified.
- Certificate expiration monitored.
- Guest account lifecycle documented.
- Troubleshooting procedure documented.
- Configuration backup/change record maintained.
- FortiOS version documented.
                         GUEST CLIENT
                              │
                              ▼
                         GUEST VLAN
                              │
                              ▼
                          FORTIGATE
                              │
                              ▼
                     FIREWALL POLICY
                              │
                              ▼
                      CAPTIVE PORTAL
                              │
                 ┌────────────┼────────────┐
                 │            │            │
                 ▼            ▼            ▼
               LOCAL         LDAP        RADIUS
                 │            │            │
                 └────────────┼────────────┘
                              │
                              ▼
                      AUTHENTICATION
                              │
                     ┌────────┴────────┐
                     │                 │
                  REJECT             ACCEPT
                     │                 │
                     ▼                 ▼
                   DENY             AUTHORIZED
                                       │
                                       ▼
                                    INTERNET
Rule #1 — Captive Portal is not the identity database.
Rule #2 — Authentication tells you who the user is; authorization determines what the user can access.
Rule #3 — Guest networks should be isolated from internal and management networks by design.
Rule #4 — Never troubleshoot Captive Portal authentication before verifying VLAN, DHCP, DNS and firewall policy.
Rule #5 — HTTPS captive portals require a certificate strategy, not merely an HTTPS toggle.
Rule #6 — Temporary guest access should have a defined lifecycle: Create → Authenticate → Access → Expire → Disable.
Rule #7 — LDAP/RADIUS failures can originate from connectivity, identity, group mapping, policy or server-side configuration—not necessarily the Captive Portal itself.
Rule #8 — Treat the Guest network as untrusted and apply least-privilege access.
- FortiGate Firewall Policy
- FortiGate User Authentication
- FortiGate LDAP
- FortiGate RADIUS
- FortiGate Guest Management
- FortiGate VLAN
- FortiGate Wireless
- FortiAP
- Network Access Control
- 802.1X
- FortiGate Security Profiles
- FortiGate Logging
- FortiView
- FortiGate Troubleshooting
- FortiGate HA
FortiGate Captive Portal
FortiOS Captive Portal
Fortinet Captive Portal
FortiGate Guest WiFi
FortiGate Guest Network
FortiGate LDAP Authentication
FortiGate RADIUS Authentication
FortiGate User Authentication
FortiGate Guest Management
FortiAP Captive Portal

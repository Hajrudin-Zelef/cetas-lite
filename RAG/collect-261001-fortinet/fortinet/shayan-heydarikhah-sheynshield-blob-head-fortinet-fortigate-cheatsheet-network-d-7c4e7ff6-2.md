---
id: collect-261001-fortinet/fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-network-d-7c4e7ff6-2
title: "shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-network-d-7c4e7ff6"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-network-d-7c4e7ff6.md
source_anchor: ""
source_lines: [222, 462]
sha256: 0d409ec257407a400970bb8bc4f3e2649ee867b429e5c3793ba079671ef09df5
---

# shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-network-d-7c4e7ff6

- Authentication method is compatible.
- RADIUS policy permits the request.
- User exists.
- User/group attributes are correct.
- Authentication test succeeds.
- RADIUS logs have been reviewed when troubleshooting.
UDP/1812 → Authentication
UDP/1813 → Accounting
⚠️ Always verify the exact RADIUS implementation, ports and authentication method used by the deployment.
User groups connect authentication identity to firewall access control.
User
 │
 ▼
Authentication Server
 │
 ▼
User Group
 │
 ▼
Firewall Policy
 │
 ▼
Network Access
- User group has been created.
- Correct authentication server has been added.
- LDAP/RADIUS membership has been verified.
- Group is referenced by the intended firewall policy.
- Unauthenticated users cannot bypass the policy.
- Different guest groups have appropriate access levels.
- Group-based authorization has been tested.
guest-users
├── guest01
├── guest02
└── guest03
Guest Management is useful for temporary or controlled guest access.
- Guest account has been created.
- Username has been generated.
- Strong/random password has been configured.
- Expiration time has been configured.
- Maximum access duration has been defined.
- Guest account owner/operator is known.
- Expired accounts are removed or disabled.
- Guest credentials are distributed securely.
- Guest activity is logged.
Create
  │
  ▼
Activate
  │
  ▼
Authenticate
  │
  ▼
Access
  │
  ▼
Expire
  │
  ▼
Disable / Remove
HTTPS should be used for captive portal authentication where supported and appropriate.
- HTTPS captive portal is enabled.
- Appropriate certificate is configured.
- Certificate is valid.
- Certificate is not expired.
- Certificate hostname/SAN matches the intended name.
- Certificate trust chain is valid.
- Client devices trust the issuing CA where possible.
- HTTP-to-HTTPS behavior has been tested.
- Mobile clients have been tested.
- Browser certificate warnings have been investigated.
HTTP Request
     │
     ▼
FortiGate
     │
     ▼
HTTPS Redirect
     │
     ▼
Captive Portal
     │
     ▼
Authentication
- Validity period checked.
- Issuer checked.
- SAN checked.
- Trust chain checked.
- Client trust verified.
- Certificate expiration monitoring is configured.
🔐 Security Principle: HTTPS captive portal security depends on a correct certificate and trust strategy, not simply enabling HTTPS.
Authentication timeout controls how long authenticated access remains valid.
config user setting
    set auth-timeout-type idle-timeout
    set auth-timeout 5
end
- Timeout requirement has been identified.
- Idle timeout has been considered.
- Hard timeout has been considered.
- Session timeout behavior is understood.
- Guest experience has been tested after timeout.
- Re-authentication has been tested.
- Timeout values match the security requirement.
User Login
    │
    ▼
Authenticated
    │
    ├── Activity continues
    │
    └── Idle period reached
              │
              ▼
       Authentication Expires
| Type | Behavior | 
|---|---|
| idle-timeout | Authentication expires after inactivity | 
| hard-timeout | Authentication expires after a fixed period | 
| session-timeout | New sessions are denied after timeout while existing sessions can continue | 
🧠 NSE Memory: Understand the behavioral difference between idle, hard, and session timeout.
- Guest → Internal LAN = DENY
- Guest → Server VLAN = DENY
- Guest → Management VLAN = DENY
- Guest → Firewall management = DENY
- Guest → Security infrastructure = restricted
- Guest → Internet = allowed only as required
- DNS access is controlled.
- DHCP access is controlled.
- Administrative interfaces are unreachable from the Guest VLAN.
- Inter-client communication is restricted where required.
- East-west guest traffic has been evaluated.
                   Guest Network
                         │
              ┌──────────┼──────────┐
              │          │          │
              ▼          ▼          ▼
           Internet    Internal    Mgmt
             ✅          ❌          ❌
Golden Principle: A Guest VLAN should be treated as untrusted.
Email collection can be used in specific guest Wi-Fi and hotspot scenarios.
- Email collection requirement has been identified.
- Captive Portal VAP is configured where applicable.
- Portal type is configured appropriately.
- Firewall policy supports the required behavior.
- Guest consent/privacy requirements have been considered.
- Only required information is collected.
- Data handling follows organizational requirements.
config wireless-controller.vap
    edit "freewifi"
        set security captive-portal
        set portal-type email-collect
    next
end
Policy example:
config firewall policy
    edit 1
        set email-collect enable
    next
end
- Hotels
- Retail
- Shopping Centers
- Guest Wi-Fi
- Public Wi-Fi
- Marketing Wi-Fi
- Hotspot deployments
diagnose firewall auth listdiagnose firewall auth mac list
- Authenticated users are visible.
- Expected users appear.
- Authentication state is correct.
- Unexpected users are investigated.
- Authentication failures are logged.
- Guest sessions are monitored.
- Expired sessions disappear as expected.
- Authentication events are available for investigation.
- Client received an IP address.
- DHCP is working.
- Default gateway is correct.
- DNS is working.
- Client can reach FortiGate.
- Client is actually connected to the Guest VLAN.
- VLAN tagging is correct.
- Wireless SSID/VAP mapping is correct where applicable.
- Correct policy exists.
- Correct source interface.
- Correct source address.
- Correct destination interface.
- Correct destination address.
- Correct service.
- Correct user/group configuration.
- NAT is enabled where required.
- Policy order has been reviewed.
- Authentication is actually being enforced.
diagnose firewall auth list
- User authentication exists.
- User belongs to the expected group.
- Authentication has not expired.
- Credentials are correct.
- Authentication failure reason has been investigated.
- LDAP server is reachable.
- Base DN is correct.
- Username attribute is correct.
- Bind credentials are valid.
- LDAP group membership is correct.
- DNS resolution works if required.
Test:
diagnose test authserver ldap <server> <username> <password>
- RADIUS server is reachable.
- UDP/1812 is reachable.
- NAS IP is correct.
- Shared secret matches.
- Authentication method matches.
- RADIUS policy allows the request.
- Access-Accept is returned.
- RADIUS logs have been reviewed.
- Certificate is valid.
- Certificate is not expired.
- SAN is correct.
- Client trusts the CA.
- Complete certificate chain is trusted.
- HTTPS portal loads without unexpected warnings.
- Authentication succeeds.
- Firewall policy permits traffic.
- NAT works.
- Default route exists.
- DNS resolution works.
- Return traffic is allowed.
- Security profiles are not unintentionally blocking required traffic.
- Upstream connectivity is healthy.
| Symptom | Primary Investigation Areas | 
|---|---|
| Portal does not appear | VLAN, DHCP, policy, redirect, client connectivity | 
| Login page appears but authentication fails | Credentials, LDAP/RADIUS, group membership | 
| LDAP authentication fails | DN, server reachability, bind account, LDAP settings | 
| RADIUS authentication fails | NAS IP, shared secret, port, RADIUS policy | 
| User authenticates but Internet fails | Policy, NAT, routing, DNS | 
| User repeatedly sees login page | Timeout, session handling, authentication state | 
| HTTPS certificate warning | Certificate, SAN, trust chain, expiration | 
| Guest reaches internal network | Missing/incorrect isolation policy | 
| Some applications fail | DNS, policy, service restrictions, application behavior | 
| Authentication suddenly expires | Timeout configuration | 

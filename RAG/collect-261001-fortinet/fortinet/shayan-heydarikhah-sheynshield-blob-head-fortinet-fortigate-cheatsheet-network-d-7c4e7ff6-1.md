---
id: collect-261001-fortinet/fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-network-d-7c4e7ff6-1
title: "shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-network-d-7c4e7ff6"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-network-d-7c4e7ff6.md
source_anchor: ""
source_lines: [1, 221]
sha256: 99ecbc27ee7b65ed2d4579b6629257c4c04ca3a4f425a37221caa33869aa5aa9
---

# shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-network-d-7c4e7ff6

FortiGate Captive Portal | FortiOS | Guest Wi-Fi | Guest Network | LDAP | RADIUS | Authentication | Network Access Control
SheynShield — Engineering Secure Networks
                         Guest Client
                              │
                              ▼
                         Guest VLAN
                              │
                              ▼
                          FortiGate
                              │
                              ▼
                       Firewall Policy
                              │
                              ▼
                      Captive Portal
                              │
                ┌─────────────┼─────────────┐
                │             │             │
                ▼             ▼             ▼
              Local          LDAP         RADIUS
                │             │             │
                └─────────────┼─────────────┘
                              │
                              ▼
                       Authentication
                              │
                    ┌─────────┴─────────┐
                    │                   │
                  ACCEPT              REJECT
                    │                   │
                    ▼                   ▼
                 Internet              DENY
- Captive Portal requirement has been identified.
- Guest users are separated from internal users.
- Dedicated Guest VLAN/interface is available.
- Guest subnet has been defined.
- Firewall policy is responsible for enforcing access.
- Authentication source has been selected.
- User group has been defined.
- Guest-to-Internet access has been designed.
- Guest-to-internal access has been explicitly denied.
- Authentication timeout has been defined.
- Logging and monitoring requirements have been defined.
- Guest account lifecycle has been defined.
- Dedicated Guest VLAN has been created.
- Guest subnet has been assigned.
- DHCP scope is configured.
- DNS configuration is available.
- Default gateway points to FortiGate.
- Guest clients can reach the FortiGate gateway.
- Guest clients cannot directly access internal VLANs.
- Guest clients cannot access management networks.
- Guest clients can reach required authentication services.
- Routing has been verified.
- Internet connectivity has been verified.
                         FortiGate
                             │
              ┌──────────────┴──────────────┐
              │                             │
          Internal                      Guest VLAN
              │                             │
       ┌──────┼──────┐                 Guest Wi-Fi
       │      │      │                      │
    Users  Servers   Mgmt              Captive Portal
                                             │
                                             ▼
                                          Internet
- Guest → Internet = ALLOW
- Guest → Internal LAN = DENY
- Guest → Server VLAN = DENY
- Guest → Management VLAN = DENY
- Guest → Security infrastructure = DENY unless required
- Guest → Authentication services = ALLOW where required
- Guest → FortiGate management = DENY
- Guest → Other guest clients = RESTRICT where required
Security Principle: Treat the Guest network as an untrusted network. Design access from deny-by-default rather than assuming guest users are trusted.
Captive Portal authentication must be integrated with the appropriate firewall policy and traffic flow.
- Source interface = Guest interface/VLAN.
- Destination interface = WAN or required destination.
- Source address is correctly defined.
- Destination address is correctly defined.
- User/group requirements are correctly configured.
- Schedule is appropriate.
- Required services are permitted.
- NAT is enabled where Internet access requires it.
- Authentication is enforced through the intended policy.
- More-specific deny policies exist where required.
- Policy order has been reviewed.
- No earlier policy unintentionally bypasses authentication.
- Return traffic is permitted.
- Routing is correct.
config firewall policy
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
⚠️ Note: The exact policy requirements depend on the FortiOS release, authentication design and deployment mode. Validate the configuration against the target FortiOS version.
Guest Client
     │
     ▼
Ingress Interface
     │
     ▼
Firewall Policy Match
     │
     ▼
Authentication
     │
 ┌───┴────┐
 ▼        ▼
Reject   Accept
          │
          ▼
       Internet
Choose the authentication backend based on the deployment requirements.
| Authentication Source | Typical Use | 
|---|---|
| Local Users | Small environments, labs, testing | 
| LDAP | Enterprise directory authentication | 
| RADIUS | Centralized authentication, NPS | 
| Guest Management | Temporary guest accounts | 
| User Groups | Identity-based access control | 
- Authentication source selected.
- Authentication server is reachable.
- Credentials have been verified.
- User group has been created.
- Authentication server is associated with the correct group.
- Group is linked to the intended firewall policy.
- Authentication logging is enabled.
- Authentication failure behavior has been tested.
Client
  │
  ▼
Captive Portal
  │
  ▼
FortiGate Local Database
  │
 ┌┴──────────┐
 ▼           ▼
Accept      Reject
- Local user has been created.
- Username has been configured.
- Strong password has been configured.
- User has been assigned to the appropriate group.
- User is not unnecessarily privileged.
- Expiration has been configured where appropriate.
- Login has been tested from the Guest network.
- Logout has been tested.
- Re-authentication has been tested.
- Expired accounts are disabled or removed.
Guest Client
     │
     ▼
Captive Portal
     │
     ▼
FortiGate
     │
     ▼
LDAP Server
     │
     ▼
Directory / Active Directory
- LDAP server IP/FQDN is configured.
- LDAP connectivity is verified.
- Correct LDAP type is selected.
- Base DN is configured correctly.
- Username/CN attribute is correct.
- Bind account is configured where required.
- Bind password is valid.
- LDAP user exists.
- LDAP group mapping is correct.
- Firewall can reach the LDAP server.
- Authentication test succeeds.
config user ldap
    edit "LDAP"
        set server <LDAP-SERVER>
        set cnid "sAMAccountName"
        set dn "dc=example,dc=local"
        set type regular
        set username <BIND-USER>
        set password <PASSWORD>
    next
end
🔐 Security: Never commit real LDAP passwords, bind credentials, API keys or secrets to GitHub.
diagnose test authserver ldap <server> <username> <password>
- Authentication returns success.
- User exists in the directory.
- User/group membership is correct.
- LDAP server is reachable.
- DNS resolution works if FQDN is used.
- Firewall rules permit the required LDAP traffic.
Client
  │
  ▼
Captive Portal
  │
  ▼
FortiGate
  │
  │ RADIUS Authentication
  ▼
RADIUS Server
  │
 ┌┴───────────────┐
 ▼                ▼
Access-Accept   Access-Reject
- RADIUS server is configured.
- Correct server IP/FQDN is configured.
- Correct authentication port is configured.
- Shared secret is configured correctly.
- FortiGate/NAS IP is recognized by the RADIUS server.

---
id: collect-261001-cisco/cisco/virtualized-enterprise-lab-pfsense-opnsense-hardening-md-at-86697a2135698e4a483468811b1834
title: "Resultant sshd_config:"
domain: cisco
role: reference
task: reference
actors: []
dates: ["2026-03"]
keywords: []
source: docs/RAG/collect-261001-cisco/virtualized-enterprise-lab-pfsense-opnsense-hardening-md-at-86697a2135698e4a483468811b18349db44685c6.md
source_anchor: ""
source_lines: [1, 116]
sha256: db940111472e5c995a29311adeed4e4ff87ebf00256b2caa2cd39dccab15e78f
---

# Resultant sshd_config:

**Lab:** Virtualized Enterprise Security Lab (pfSense/OPNsense)

**Standard:** CIS Controls v8 · NIST SP 800-53 Rev 5 · ISO/IEC 27001:2022

**Architecture:** Zero-Trust (ZTNA) · SASE-aligned · Micro-segmented

**Last Reviewed:** March 2026


**CIS:** 4.1 | **NIST:** AC-17, IA-2 | **ISO 27001:** A.8.5

All SSH to pfSense/OPNsense restricted to public-key authentication. Password login is explicitly disabled, eliminating brute-force attack vectors against the management plane. Each admin's **Ed25519 public key** is imported under `System > User Manager > Admin User > Authorized Keys`.

```
# Resultant sshd_config:
PasswordAuthentication no
PubkeyAuthentication yes
PermitRootLogin without-password
MaxAuthTries 3
AllowUsers admin
```
**CIS:** 9.2 | **NIST:** SI-3 | **ISO 27001:** A.8.7

**pfBlockerNG** runs on the internal Unbound resolver, blocking malware C2, phishing, and cryptomining domains before any network connection is established. IOT/GUEST queries that resolve known C2 FQDNs receive `NXDOMAIN` — the connection is never attempted.

```
Active Feeds: hpHosts, Abuse.ch URLhaus, SANS ISC, Stevenblack Hosts
DNSBL Action: Unbound Redirect (NXDOMAIN)
Update Frequency: Every 4 hours
```
**CIS:** 8.2 | **NIST:** AU-2, SI-4 | **ISO 27001:** A.8.15

All firewall events (block/pass), Suricata alerts (EVE JSON), DHCP leases, and authentication events are shipped to the central syslog aggregator at `10.0.10.20` (Graylog/Elastic SIEM/Wazuh). Provides full audit trail for compliance and SOC investigation.

```
<syslog>
  <enable>1</enable>
  <logall>1</logall>
  <remoteserver>10.0.10.20</remoteserver>
  <ipproto>udp</ipproto>
</syslog>
```
**CIS:** 13.4 | **NIST:** SC-5 | **ISO 27001:** A.8.20

**Anti-spoofing** enabled on every VLAN interface — prevents traffic with a source IP outside the interface's subnet from being forwarded. **Bogon filtering** on WAN drops RFC 1918 and unallocated IP space arriving from the internet.

```
Interfaces > WAN: Block private networks ✅, Block bogon networks ✅
Firewall Advanced: IP Source Routing disabled, IPv6 disabled (unless required)
```
**CIS:** 13.7 | **NIST:** SI-3, SI-4 | **ISO 27001:** A.8.16

Suricata runs in **Inline IPS mode** (not passive tap), actively dropping malicious packets in real-time. Active rule sets: Emerging Threats Open, `et/trojan`, `et/scan`, `et/iot`, JA3 fingerprinting, Abuse.ch SSLBL. All alerts forwarded via EVE JSON to SIEM.

**CIS:** 12.8 | **NIST:** IA-3, SC-8 | **ISO 27001:** A.8.24

All service-to-service communication in PROD requires **mTLS** with **SPIFFE/SVID** identity certificates (issued by internal SPIFFE/Tornjak CA). Certificates rotate automatically every 24 hours. TLS 1.0/1.1/RC4/3DES blocked via firewall alias rule. Enforces ZTNA — no implicit trust from network location.

**CIS:** 4.6 | **NIST:** AC-3 | **ISO 27001:** A.8.3

The pfSense/OPNsense WebGUI is accessible **only from `10.0.10.5/32`** (admin jump-host). All other sources are explicitly blocked. The GUI runs on non-standard port **`8443`** to reduce automated scanner exposure.

```
<rule>
  <type>pass</type>
  <source><address>10.0.10.5</address></source>
  <destination><address>10.0.10.1</address><port>8443</port></destination>
  <descr>ALLOW: Admin jump-host only → WebGUI</descr>
</rule>
```
**CIS:** 13.4 | **NIST:** SC-7 | **ISO 27001:** A.8.22

Every inter-VLAN path requires an **explicit permit rule**. The final rule on every interface is `Block ALL — Source: Any, Destination: Any, Log: YES`. No zone has implicit access to any other zone. All deny hits are logged for SIEM correlation, enforcing true micro-segmentation.

|  | MGMT | PROD | IOT | GUEST | DMZ | 
|---|---|---|---|---|---|
| **MGMT** | N/A | ❌ | ❌ | ❌ | ❌ | 
| **PROD** | ❌ | Alias-limited | ❌ | ❌ | API-only | 
| **IOT** | ❌ | ❌ | ❌ | ❌ | ❌ | 
| **GUEST** | ❌ | ❌ | ❌ | ❌ | ❌ | 
| **DMZ** | ❌ | API-only | ❌ | ❌ | N/A | 

**CIS:** 12.6 | **NIST:** SC-8, SC-28 | **ISO 27001:** A.8.24

Site-to-site VPN uses **WireGuard + Pre-Shared Key (PSK)** per RFC 8784, providing quantum-resistance by adding a symmetric key layer over the Curve25519 key exchange. PSK is rotated every 30 days via automated script. Legacy VPN methods (IKEv1, PPTP, L2TP/PAP) are fully disabled. Cipher: ChaCha20-Poly1305 (AEAD).

**CIS:** 6.3 | **NIST:** IA-2 | **ISO 27001:** A.8.5

All admin access requires **TOTP-based MFA** via FreeRADIUS. Authentication flow:

```
Admin → Jump-Host SSH (Ed25519 key + TOTP)
       → pfSense WebGUI (Username + Password + RADIUS TOTP challenge)
```
```
RADIUS: FreeRADIUS @ 10.0.10.25:1812
TOTP Standard: RFC 6238 (30s time step, 6 digits)
Session Timeout: 8 hours
Concurrent Logins: 1 per admin account
```
Optional: YubiKey 5 NFC (OTP mode) for NIST AAL3 assurance level.

| Step | CIS v8 | NIST 800-53 | ISO 27001:2022 | 
|---|---|---|---|
| 1. SSH Key-Only | 4.1 | AC-17, IA-2 | A.8.5 | 
| 2. pfBlockerNG DNSBL | 9.2 | SI-3 | A.8.7 | 
| 3. Remote Syslog | 8.2 | AU-2, SI-4 | A.8.15 | 
| 4. Anti-Spoofing | 13.4 | SC-5 | A.8.20 | 
| 5. Suricata IPS | 13.7 | SI-3, SI-4 | A.8.16 | 
| 6. mTLS PROD | 12.8 | IA-3, SC-8 | A.8.24 | 
| 7. WebGUI MGMT-Only | 4.6 | AC-3 | A.8.3 | 
| 8. VLAN Micro-seg | 13.4 | SC-7 | A.8.22 | 
| 9. WireGuard PSK | 12.6 | SC-8, SC-28 | A.8.24 | 
| 10. TOTP MFA | 6.3 | IA-2 | A.8.5 | 

*All hardening measures align with 2026 enterprise security standards — ZTNA, SASE, Micro-segmentation, Zero-Trust.*

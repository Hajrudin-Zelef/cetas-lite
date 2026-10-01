---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/critical-unifi-os-auth-bypass-flaws-41b8965d
title: "critical-unifi-os-auth-bypass-flaws-41b8965d"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Google"]
dates: ["2026-05-21"]
keywords: ["advisory", "exploit"]
source: docs/RAG/collect-261001-unifi-ubiquiti/critical-unifi-os-auth-bypass-flaws-41b8965d.md
source_anchor: ""
source_lines: [1, 22]
sha256: bf4645c7204cc25001f30f15b84ad8fb424de750b30295269215a0a7377ce47d
---

# critical-unifi-os-auth-bypass-flaws-41b8965d

Ubiquiti has addressed three critical vulnerabilities within the UniFi OS Server that attackers can chain together to achieve unauthenticated remote code execution (RCE) with root privileges.
Disclosed on May 21, 2026, via Security Advisory Bulletin 064 (SAB-064), the flaws are tracked as CVE-2026-34908, CVE-2026-34909, and CVE-2026-34910.
Each vulnerability carries a maximum CVSS 3.1 severity score of 10.0. Security researchers at Bishop Fox demonstrated the full end-to-end exploit chain on version 5.0.6, proving that a single crafted HTTP request can yield a root shell without requiring credentials or user interaction.
Critical UniFi OS Vulnerabilities
The first stage of the attack leverages an Authentication Gateway Bypass (CVE-2026-34908 and CVE-2026-34909). The Nginx authentication handler evaluates the raw, percent-encoded request URI to determine if a route is exempt from authentication.
However, Nginx selects the upstream backend using the normalized URI, where %2f decodes to / and directory traversal sequences collapse. 
Attackers can bypass the gateway by prefixing a request with the auth-exempt /api/auth/validate-sso/ endpoint while using normalized paths to access authenticated internal routes, Bishopfox said.
Once past the gateway, the attacker triggers a Command Injection flaw in the package-update service (CVE-2026-34910). This service accepts user-supplied package names and passes them through fmt.Sprintf into a command string executed via sh -c. 
Due to a lack of input validation in version 5.0.6, shell metacharacters are interpreted directly. The injected command initially runs under the ucs-update service account. 
Because this account possesses passwordless sudo privileges for critical binaries like /usr/bin/dpkg and /bin/systemctl, attackers can easily escalate to full root privileges by installing a maliciously crafted .deb package.
Root access on a UniFi OS appliance yields complete control over an organization’s network management plane. Threat actors can extract the JWT signing key to forge persistent administrator sessions that survive patching, password resets, and system reboots.
Bishop Fox confirmed that a forged owner-scope JWT token minted from a stolen key will authenticate successfully even against fully patched 5.0.8 consoles.
The patch closes the initial entry point but leaves the underlying token verification model unchanged, meaning stolen keys will continue to generate valid sessions indefinitely.
Furthermore, attackers can exfiltrate TLS private keys, cloud access tokens, and the full PostgreSQL user database. The compromise extends beyond digital infrastructure into physical security environments.
In deployments using UniFi Access and UniFi Protect, attackers can unlock physical doors, clone NFC and facial-recognition credentials, monitor live camera feeds, and permanently delete surveillance footage.
Mitigation
Administrators must immediately update to UniFi OS Server 5.0.8 or the hardware-equivalent fixed version. Per SAB-064, most Cloud Gateways require version 5.1.12; the UNAS line requires 5.1.10; Dream Machine Beast requires 5.1.11; and UniFi Express requires 4.0.14.
Organizations must treat any externally exposed, unpatched instance as fully compromised and rebuild it from a known-good image rather than simply updating it.
- Restrict TCP port 11443 to a dedicated management VLAN and explicitly block all external access.
- Rotate the JWT signing key (/data/unifi-core/config/jwt.yaml ), TLS keys, cloud access tokens, RADIUS secrets, and database credentials before rebooting.
- Treat all biometric and NFC data as permanently disclosed, as these physical credentials cannot be safely rotated.
Follow us on Google News, LinkedIn, and X to Get Instant Updates and Set GBH as a Preferred Source in Google.

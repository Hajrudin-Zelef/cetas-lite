---
id: collect-261001-fortinet/fortinet/blog-fortigate-firewall-audit-checklist-india-2026-25a2d727-2
title: "Disable SSL VPN if not needed:"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: ["2026-04", "2026-07"]
keywords: ["benchmark", "benchmarks", "cybersecurity", "incident", "memory", "safeguards"]
source: docs/RAG/collect-261001-fortinet/blog-fortigate-firewall-audit-checklist-india-2026-25a2d727.md
source_anchor: ""
source_lines: [25, 146]
sha256: f0fdd318c6886d883985cb708fdd87ef8aec708f395f3e7dd081e77032ec0195
---

# Disable SSL VPN if not needed:

config system admin
  edit "admin"
    set password <strong-password>
  next
end
    2. HTTP/Telnet Management on WAN Interface
CRITICAL
Enabling HTTP or Telnet admin access on the WAN interface exposes credentials in cleartext. Attackers passively sniff traffic or brute-force the login. Management access should be HTTPS-only on internal interfaces with trusted host restrictions.
config system interface
  edit "wan1"
    set allowaccess ping https ssh
  next
end
    3. No Trusted Host Restriction for Admin Accounts
HIGH
Without trusted host configuration, admin accounts can log in from any source IP — including the internet. Trusted hosts restrict management access to specific subnets, making brute-force and credential-stuffing attacks from external IPs impossible.
config system admin
  edit "admin"
    set trusthost1 10.10.10.0 255.255.255.0
    set trusthost2 192.168.1.0 255.255.255.0
  next
end
    4. No Two-Factor Authentication for Admin Logins
HIGH
Single-factor admin authentication means a stolen or guessed password gives complete control of the firewall. FortiToken (hardware or mobile) adds a second factor that prevents credential-based attacks even if the password is compromised.
config system admin
  edit "admin"
    set two-factor fortitoken
    set fortitoken <serial-number>
  next
end
    5. SSL VPN Unpatched or Unnecessarily Exposed
CRITICAL
SSL VPN has been the attack surface for CVE-2023-27997, CVE-2024-21762, and the symlink persistence technique. If you are not actively using SSL VPN, disable it entirely. If you need it, ensure FortiOS is current and restrict portal access to specific source IPs or use ZTNA as a replacement.
# Disable SSL VPN if not needed:
config vpn ssl settings
  set status disable
end
# Or restrict source IPs:
config firewall address
  edit "SSLVPN-Allowed-IPs"
    set subnet 203.0.113.0 255.255.255.0
  next
end
    6. Overly Permissive Any/Any/Any Rules
HIGH
Policies with source=any, destination=any, service=any bypass the entire purpose of a firewall. These are often created during initial deployment as “temporary” rules and never removed. They allow unrestricted lateral movement and data exfiltration.
# Identify overly permissive rules:
diagnose firewall policy list | grep -i "all"
# Replace with specific src/dst/service policies
# Enable UTM inspection on every allow rule
    7. Unused, Shadow, and Redundant Policies
MEDIUM
Over time, firewall policies accumulate. Unused policies (zero hit counts) create confusion, shadow policies (never matched because a broader rule above catches the traffic first) mask intent, and redundant policies slow down processing. A clean policy table is essential for security and performance.
# Check policy hit counts:
diagnose firewall policy list
# Delete policies with zero hits after review:
config firewall policy
  delete <policy-id>
end
    8. FortiGuard Subscriptions Expired
HIGH
When FortiGuard subscriptions expire, IPS signatures stop updating, AV definitions go stale, web filtering categories freeze, and application control loses visibility into new apps. Your FortiGate effectively becomes a stateful packet filter — blind to modern threats. We find expired subscriptions in ~30% of audits.
# Check subscription status:
diagnose autoupdate status
diagnose autoupdate versions
# Verify FortiGuard connectivity:
exec ping guard.fortinet.com
exec telnet guard.fortinet.com 443
    9. Logging Disabled or Memory-Only
MEDIUM
Memory-only logging means logs are lost on reboot — exactly when you need them most (during an incident). CERT-In requires 180 days of log retention. Without syslog/FortiAnalyzer forwarding and disk logging, incident response and forensics are impossible, and you fail every compliance audit.
config log disk setting
  set status enable
  set maximum-log-age 180
end
config log fortianalyzer setting
  set status enable
  set server <FortiAnalyzer-IP>
end
    10. Firmware Not Updated or Running EOL FortiOS
CRITICAL
Running end-of-life FortiOS versions (6.0, 6.2, 6.4) means you receive no security patches — period. Known CVEs remain permanently exploitable. We routinely find devices running FortiOS 6.0.x in production, sometimes with internet-facing SSL VPN. The upgrade path exists; the risk of not upgrading is existential.
# Check current firmware:
get system status
# Upgrade path (always use Fortinet's tool):
# https://docs.fortinet.com/upgrade-tool
# FortiOS 6.4 → 7.0 → 7.2 → 7.4 → 7.6
    CIS FortiGate Benchmark — Top 10 Controls
The Center for Internet Security (CIS) publishes vendor-specific security benchmarks — consensus-based configuration guidelines developed by cybersecurity practitioners, government agencies, and technology vendors. The CIS FortiGate Benchmark is the gold standard for FortiOS hardening and is referenced by auditors worldwide.
Current versions: CIS FortiGate 7.0.x Benchmark v1.2.0 and CIS FortiGate 7.4.x Benchmark v1.0.0. Each contains 50+ individual controls across categories including management plane, data plane, logging, authentication, and network configuration.
- Disable HTTP and Telnet on all interfaces — Management access must use HTTPS and SSH only (CIS 1.1)
- Configure trusted hosts for all admin accounts — Restrict management access to known admin subnets (CIS 1.2)
- Set idle timeout to 5 minutes or less — Prevent session hijacking from unattended terminals (CIS 1.3)
- Enable two-factor authentication for all admins — FortiToken or RADIUS with MFA (CIS 1.4)
- Encrypt admin password storage using strong hashing — Use set admin-password-hash sha256 (CIS 1.5)
- Configure NTP with authentication — Accurate timestamps are essential for log correlation and forensics (CIS 2.1)
- Enable logging to external syslog or FortiAnalyzer — Memory-only logging is insufficient for incident response (CIS 3.1)
- Enable DNS filtering on all outbound policies — Block C2 callbacks and malicious domain resolution (CIS 4.1)
- Disable unused interfaces and services — Reduce attack surface by shutting down ports not in use (CIS 5.1)
- Configure firmware auto-update notifications — Ensure visibility into available security patches (CIS 6.1)
Ogma’s Firewall Audit benchmarks your FortiGate against the full CIS profile — not just the top 10. Our report maps each finding to the specific CIS control ID, severity level, and provides exact CLI remediation commands. Book a Firewall Compliance Audit.
India Compliance Landscape for Firewall Audits
Indian enterprises face an increasingly stringent compliance environment where firewall audits are no longer a best practice — they are a legal requirement. Here is how each regulatory framework applies to your FortiGate:
CERT-In (Indian Computer Emergency Response Team)
- 6-hour incident reporting to CERT-In for any cybersecurity incident
- 180-day log retention within Indian jurisdiction
- Annual cybersecurity audit by a CERT-In format auditors (mandatory from July 2026)
- Audit must cover firewall configurations, access controls, and vulnerability management
RBI IS Audit Framework (Banking & NBFCs)
- Annual IS audit mandated for all scheduled commercial banks and NBFCs
- Firewall rule review is a specific audit checklist item
- Must verify no any/any rules, proper segmentation, and logging
- Board-level reporting of audit findings required
SEBI CSCRF (Capital Markets)
- Effective April 2026 for stock brokers, depository participants, and MIIs
- Quarterly firewall rule review mandated
- Network segmentation between trading and back-office systems required
- Audit trail of all firewall policy changes
DPDPA 2023 & IT Act Section 43A
- Digital Personal Data Protection Act 2023 — significant data fiduciaries must implement “reasonable security safeguards”
- IT Act Section 43A — bodies corporate handling sensitive personal data must maintain “reasonable security practices”

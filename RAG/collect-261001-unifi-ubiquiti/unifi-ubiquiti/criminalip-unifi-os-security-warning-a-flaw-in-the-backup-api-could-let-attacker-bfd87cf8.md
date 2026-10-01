---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/criminalip-unifi-os-security-warning-a-flaw-in-the-backup-api-could-let-attacker-bfd87cf8
title: "criminalip-unifi-os-security-warning-a-flaw-in-the-backup-api-could-let-attacker-bfd87cf8"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["United States"]
dates: ["2025-11"]
keywords: ["attention", "distribution", "exploit"]
source: docs/RAG/collect-261001-unifi-ubiquiti/criminalip-unifi-os-security-warning-a-flaw-in-the-backup-api-could-let-attacker-bfd87cf8.md
source_anchor: ""
source_lines: [1, 39]
sha256: c11778a7494cb96c11c21ff1a290b36e0981b49aa408c7b218072eca9548b488
---

# criminalip-unifi-os-security-warning-a-flaw-in-the-backup-api-could-let-attacker-bfd87cf8

UniFi OS Security Warning: A flaw in the Backup API could let attackers run code without authentication (CVE-2025–52665)
In early November 2025, researchers disclosed a severe unauthenticated remote code execution vulnerability in Ubiquiti’s UniFi OS ecosystem. Tracked as CVE-2025–52665, the flaw originates in the input handling of the backup orchestration endpoint (/api/ucore/backup/export). Instances were observed where the backup API, intended to run only on the local loopback was instead reachable from external networks. The dir parameter supplied by the client is directly interpolated into a shell command chain, allowing meta-characters (e.g., ;, #) to inject arbitrary commands. Proof-of-concept testing confirmed that an attacker can exfiltrate system files and achieve remote control.
This is not merely a service bug — in real deployments the vulnerability can lead to full device compromise and even physical-security bypass. This article explains the vulnerability mechanics and exploitation flow, then shows how to detect exposed assets using Criminal IP and details immediate, practical remediation steps for security teams.
In-Depth Analysis of CVE-2025–52665: Command Injection in UniFi OS Backup API
The core issue is that the backup creation process calls several utilities (mktemp, chmod, tar, etc.) via the shell. The backup orchestrator directly substitutes a user- or API-client-provided path (dir) into the file-system operation chain. If input validation or escaping is not performed, an attacker can insert meta-characters and terminate the existing command to inject new commands. For example, if dir contains a semicolon, commands appended after the intended compress/move workflow will execute.
If this endpoint does not require authentication or is unintentionally exposed externally, an attacker can send a malicious JSON payload over the network without additional interaction. In other words, if network access is available, automated scanning can enable mass exploitation. In addition, similar unauthenticated API endpoints were reported (e.g., /api/v1/user_assets/nfc, /api/v1/user_assets/touch_pass/keys), and those endpoints may handle NFC credentials or mobile pass sensitive keys; exposure could therefore threaten physical security.
Attack Walkthrough — Real-World Exploit Chain for UniFi OS RCE
The attack is relatively simple.
- UniFi Device Identification: The attacker scans for UniFi devices exposed to the Internet on public ports (for example, 9780).
- Vulnerable API Call: The attacker sends an unauthenticated POST request to /api/ucore/backup/export and inserts a malicious string containing meta-characters (e.g.,; ,|) into thedir parameter.
- Command-Injection Execution: During processing by the backup script, the malicious command is executed, allowing the attacker to obtain a reverse shell or exfiltrate system files.
- Intrusion and Propagation: Using harvested SSH keys, API tokens, NFC keys, etc., the attacker moves laterally within the network and carries out further attacks. In environments integrated with UniFi Access, this can neutralize physical access controls (door operation, NFC issuance, etc.).
UniFi OS RCE Exposure Scan: Criminal IP Asset Identification
In Criminal IP Asset Search, search product: “unifi os” to quickly identify UniFi OS instances exposed to the Internet.
Criminal IP Search Query: product: “unifi os”
As of 4 November 2025, the number of UniFi OS instances searched by Criminal IP totals 90,884.
Criminal IP’s Maps function can visualize the query results on a world map, letting you view geographic distribution and concentration of exposed assets at a glance.
Get Criminal IP’s stories in your inbox
Join Medium for free to get updates from this writer.
From the same query results, of the 90,884 instances, the United States had 38,763 instances (the most), followed by Germany 9,999, the United Kingdom 5,414, and the Netherlands 3,478.
Criminal IP’s IP Report page shows which vulnerabilities affect a specific IP, as well as its exploitation history, open ports, IP location, WHOIS data, and other details. For example, clicking one IP from the product: “unifi os” results showed 13 open ports, one policy violation, and a remote address is true status.
How to Secure UniFi OS Against RCE Exploitation
- Immediate Patching (Top Priority)
Apply vendor-issued security patches immediately. Where possible, upgrade devices to the vendor-recommended versions. If patching is not possible, isolate vulnerable devices to a management network or keep them offline.
- Network Access Control
If management port 9780 is publicly exposed, block it with firewall or security group rules. Management should be performed only via VPN, bastion, or a management-only subnet.
- Endpoint Temporary Mitigation
If the backup API can be disabled, do so. If not, apply IP whitelisting. As an interim measure, apply WAF or network filtering rules to block shell meta-characters (semicolon, pipe, ampersand, etc.) to reduce the attack surface.
- Authentication And Authorization Review
 Verify that all API paths have authentication and authorization. Pay particular attention to endpoints that handle sensitive data such as/api/v1/user_assets/nfc or/api/v1/user_assets/touch_pass/keys .
- Credential Revocation And Reissuance
If compromise is suspected, immediately invalidate and reissue SSH keys, API tokens, NFC credentials, and certificates. If physical access credentials may have been leaked, those credentials must be revoked and reissued.
- Detection And Forensics Enhancement
 Add detection rules in SIEM/EDR for: abnormal POSTs to/api/ucore/backup/export , mass scanning to port 9780, unexpected execution ofmktemp /tar , and abnormal temporary-directory file creation. Isolate suspected hosts and acquire forensic images.
- Continuous Monitoring (ASM Integration)
 Use Criminal IP to monitor queries such as:product: "unifi os", cve_id:CVE-2025-52665, port:9780 Regular monitoring enables automatic detection of new exposures and rapid response prioritization.
Conclusion
CVE-2025–52665 is not merely a remote command execution vulnerability; in the presence of unauthenticated and publicly exposed endpoints, it can fully compromise devices in one step. Environments integrated with UniFi Access are at increased risk because physical security can also be compromised. The priority is vendor patching; before and after patching, apply network blocking, endpoint disabling, authentication strengthening, credential revocation, and ASM-based continuous monitoring. Using Criminal IP to measure exposure and automate prioritization allows limited resources to reduce the attack surface quickly.
In relation to this, you can refer to RediShell RCE Alert: Over 8,000 Redis Instances — Immediate Update Recommended.

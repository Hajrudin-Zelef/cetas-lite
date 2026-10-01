---
id: collect-261001-fortinet/fortinet/news-events-alerts-2026-06-18-cisa-urges-hardening-fortinet-devices-after-report-9bad69c5
title: "news-events-alerts-2026-06-18-cisa-urges-hardening-fortinet-devices-after-report-9bad69c5"
domain: fortinet
role: reference
task: reference
actors: ["Apple", "CISA"]
dates: ["2026-06-22"]
keywords: ["cyber", "disclosure"]
source: docs/RAG/collect-261001-fortinet/news-events-alerts-2026-06-18-cisa-urges-hardening-fortinet-devices-after-report-9bad69c5.md
source_anchor: ""
source_lines: [1, 20]
sha256: 45852b5a1e0b172273173818bbc7c649f459c7a79ef31a4b83dd3e0f921a5b24
---

# news-events-alerts-2026-06-18-cisa-urges-hardening-fortinet-devices-after-report-9bad69c5

CISA Urges Hardening Fortinet Devices After Reports of Credential Exposure
Update June 22, 2026:
CISA has updated this Alert to incorporate the link to Fortinet’s recent guidance on this activity. 
CISA is aware of global reports that malicious cyber actors have targeted internet-accessible Fortinet devices across government and private sector organizations using compromised credentials. This activity, referred to as FortiBleed, involves the exposure of leaked credentials associated with approximately 74,000 Fortinet devices, including firewalls and virtual private network (VPN) gateways.
To defend against this malicious cyber activity, CISA urges impacted Fortinet customers with FortiGate appliances and associated secure sockets layer (SSL) VPN gateways to immediately:
- Terminate sessions and reset credentials. Terminate all active SSL VPN and administrative sessions. Reset all Fortinet VPN and administrative passwords, especially on internet-facing systems, and enforce strong password policies.
- Ensure secure credential storage. Confirm your organization’s use of the Password-Based Key Derivation Function 2 (PBKDF2) algorithm to store administrator credentials and remove weaker legacy hashes per Fortinet’s guidance (see, Fortinet's Technical Tip: Enforcing PBKDF2 as hash function for administrator accounts in FortiOS v7.2.11 and later).
- Review logs. Review firewall, VPN, authentication, and domain controller logs for lateral movement, unusual access, suspicious accounts, or unauthorized configuration changes.
- Enable phishing-resistant multifactor authentication (MFA). Require phishing-resistant MFA on all remote access and administrative accounts and ensure it is enforced on all external gateways and administrative interfaces.
- Reduce the attack surface and lock down management access. Ensure the administration of your firewall is inaccessible from the public internet; restrict Fortinet management interfaces to trusted internal networks; and remove or disable any unauthorized or unnecessary accounts.
See the following resources to determine your organization’s potential impact and find additional guidance on the credentials compromised:
- Tech Times: Fortinet FortiGate Credential Leak Hits 73,932 Firewalls: Half the Internet-Facing Fleet
- SOCRadar: FortiBleed: The Compromise of 80,000+ Fortinet Firewalls
- Hudson Rock: FortiBleed: 75,000 Fortinet Firewalls Compromised: Global Enterprises Exposed – Claim Your Ethical Disclosure
- Arctic Wolf: Active FortiBleed Campaign Impacting Fortinet Devices Across 194 Countries
- Fortinet: Attacks at the Speed of AI
- Fortinet: Analysis of Reported Credential Compromise of FortiGate Devices
Disclaimer
The information in this report is being provided “as is” for informational purposes only. CISA does not endorse any commercial entity, product, company, or service, including any entities, products, or services linked within this document. Any reference to specific commercial entities, products, processes, or services by service mark, trademark, manufacturer, or otherwise, does not constitute or imply endorsement, recommendation, or favoring by CISA.
This product is provided subject to this Notification and this Privacy & Use policy.

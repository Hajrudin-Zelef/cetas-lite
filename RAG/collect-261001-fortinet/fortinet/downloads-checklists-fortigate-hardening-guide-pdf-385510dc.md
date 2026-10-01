---
id: collect-261001-fortinet/fortinet/downloads-checklists-fortigate-hardening-guide-pdf-385510dc
title: "downloads-checklists-fortigate-hardening-guide-pdf-385510dc"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/downloads-checklists-fortigate-hardening-guide-pdf-385510dc.md
source_anchor: ""
source_lines: [1, 33]
sha256: 3ee613db6696805fccf922ed38ac0aac29d86397330937a617ab98468b87a74d
---

# downloads-checklists-fortigate-hardening-guide-pdf-385510dc

FortiGate Hardening Checklist
BLACKHAWK DATA
BlackHawk Data  |  (877) 383-1845  |  info@blackhawk11.com
Page 1
Administrative Access
Default admin password changed
Trusted hosts configured for admin access
HTTPS-only admin access (HTTP disabled)
Admin session timeout configured (max 15 min)
Two-factor authentication for admin accounts
Login banner configured
Admin profiles with least privilege
Network Security
Unused interfaces disabled
Management interface on dedicated VLAN
DNS set to trusted resolvers (not ISP defaults)
NTP configured with authentication
SNMP community strings changed from defaults
SNMPv3 used instead of v1/v2c
Firewall Policy
Deny-all default policy in place
No any/any allow rules
Unused policies removed
Policy logging enabled for all rules
IPS profiles applied to relevant policies
Application control enabled
SSL deep inspection configured where appropriate
Logging & Monitoring
FortiAnalyzer or syslog configured
Log all traffic (or at minimum denied traffic)
Alert thresholds configured
Firmware on vendor-recommended version
Automatic signature updates enabled

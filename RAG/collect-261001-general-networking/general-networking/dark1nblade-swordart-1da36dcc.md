---
id: collect-261001-general-networking/general-networking/dark1nblade-swordart-1da36dcc
title: "dark1nblade-swordart-1da36dcc"
domain: general-networking
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["benchmark"]
source: docs/RAG/collect-261001-general-networking/dark1nblade-swordart-1da36dcc.md
source_anchor: ""
source_lines: [1, 17]
sha256: c6914163dd20fd4459a862606ac1a670f8787ea10b97db78bd0dcb424a12b8a5
---

# dark1nblade-swordart-1da36dcc

A Python-based CLI tool to audit FortiGate firewall configurations against the CIS Fortinet FortiOS Benchmark.
- Multiple Input Methods: Support for live devices (SSH/API) and offline config files.
- Comprehensive CIS Checks: Covers Management Access, Authentication, Logging, Network Services, Firewall Policies, VPN, and System Hardening.
- Flexible Reporting: Output results in Console, JSON, HTML, or CSV formats.
- Interactive Dashboard: Generate a visual HTML dashboard with Chart.js.
- CI/CD Ready: Exit codes for critical findings.
pip install .fortigate-cis-audit --file backup.conf --output consolefortigate-cis-audit --file backup.conf --dashboardfortigate-cis-audit --host 192.168.1.1 --user admin --password mypassword --output htmlfortigate-cis-audit --file backup.conf --fail-on-critical
- CIS-1.1: Ensure HTTPS-only admin access
- CIS-1.4: Check idle session timeout <= 5 minutes
- CIS-2.3: Ensure SSH v2 only
- CIS-3.1: Confirm syslog server is configured
- CIS-4.3: Ensure NTP is configured
- CIS-5.1: Flag any permit any any policies
- CIS-6.1: Enforce IKEv2 for IPsec tunnels
- CIS-7.2: Verify auto-install from USB is disabled
- fortigate_cis_audit/ : Main package
- tests/ : Unit tests

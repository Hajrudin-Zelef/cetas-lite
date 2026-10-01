---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/adurrr-adurrr-github-io-blob-head-content-post-sysadmin-opnsense-mastery-en-md-de80efd7-3
title: "/usr/local/opnsense/service/conf/actions.d/conf.d/"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["attention", "aws", "memory", "open source", "throughput"]
source: docs/RAG/collect-261001-opnsense-pfsense/adurrr-adurrr-github-io-blob-head-content-post-sysadmin-opnsense-mastery-en-md-de80efd7.md
source_anchor: ""
source_lines: [262, 376]
sha256: 78a786430feaf5519a7780d109a08e0f534bdbe41849ef5eea830b758726f0fe
---

# /usr/local/opnsense/service/conf/actions.d/conf.d/

It's tempting to think that a formal auditing framework is only for companies. But the reality is more pragmatic: an auditing framework is a checklist that people with more experience than you have already validated. Following it avoids the "I forgot to review X" problem that appears when the security routine depends only on memory.

There's an additional reason if you're in Spain: the National Security Scheme (ENS, Real Decreto 311/2022) is mandatory for the public sector and its technology providers<sup>1</sup>. This increasingly includes freelancers and small companies providing services to public administrations. Having the infrastructure aligned with the ENS, even at a basic level, isn't just good practice—it's a potential requirement.

ISO 27001:2022 organizes security controls in Annex A<sup>2</sup>. The ones that apply directly to what was configured in this series:

| Control | Description | OPNsense Implementation | 
|---|---|---|
| A.8.20 | Networks security | WAN deny-all firewall, IPS, CrowdSec, GeoIP | 
| A.8.22 | Networks segregation | 5 VLANs with restrictive inter-VLAN rules | 
| A.8.23 | Web filtering | Zenarmor DPI with per-category policies | 
| A.8.15 | Logging | Remote syslog with TLS, Suricata EVE JSON | 
| A.8.16 | Monitoring activities | Suricata alerts, CrowdSec dashboards, Zenarmor | 
| A.8.17 | Clock synchronization | NTP configured in System > General (essential for log correlation) | 
| A.5.15 | Access control | Least-privilege policy, dedicated admin | 
| A.5.17 | Authentication information | 16+ character passwords, OTP enabled | 
| A.5.18 | Access rights | Periodic review of users and permissions | 
| A.8.2 | Privileged access | Restricted root, separate admin, SSH keys only | 

The standard doesn't prescribe exact review frequencies, but auditors expect: firewall rule review at least quarterly, privileged access review quarterly, and continuous log monitoring with weekly manual review.

The ENS (RD 311/2022) classifies systems into three levels: basic, medium, and high. The relevant measures for a firewall:

| ENS Measure | Description | Minimum Level | Implementation | 
|---|---|---|---|
| mp.com.1 | Secure perimeter | Medium | WAN firewall, GeoIP, deny-all by default | 
| mp.com.2 | Confidentiality protection | Medium | WireGuard VPN, DoT for DNS | 
| mp.com.4 | Networks segregation | Medium | VLANs with inter-VLAN rules | 
| op.exp.2 | Security configuration | Basic | SSH hardening, disabled services | 
| op.exp.8 | Activity logging | Basic | Remote syslog (2-year retention for medium/high) | 
| op.acc.1 | Identification | Basic | Unique users, no shared accounts | 
| op.acc.5 | Authentication mechanism | Basic | Strong passwords; OTP for medium/high | 
| op.acc.7 | Remote access | Basic | WireGuard VPN mandatory for remote admin | 
| op.mon.1 | Intrusion detection | Medium | Suricata IPS + CrowdSec | 

The CCN-STIC guides from the National Cryptology Center provide detailed implementation instructions. Specifically, CCN-STIC-811 for system interconnection and CCN-STIC-408 for perimeter security are the most relevant for this context<sup>3</sup>.

An important ENS detail: log retention for medium and high levels is a minimum of two years. If the remote syslog doesn't have capacity for that, it needs to be planned. With Loki and compression, firewall logs from a homelab don't take up much space, but it needs to be considered from the design phase.

Combining ISO 27001 and ENS recommendations with what's realistic for a small environment:

| Frequency | Task | Type | 
|---|---|---|
| Daily | Automated review of IDS/IPS alerts and CrowdSec decisions | Automated | 
| Weekly | Manual log review: blocked events, failed login attempts, anomalous traffic | Manual | 
| Monthly | Firewall rule review and alias update. Verify threat feeds are updating | Manual | 
| Quarterly | Complete rule review: identify rules with no hits, overly permissive rules, forgotten temporary rules | Manual | 
| Quarterly | Access review: active users, permissions, valid SSH keys, API tokens | Manual | 
| Semi-annual | Basic exposure test: `nmap` from outside the network against the public IP | Manual | 
| Annual | Complete security posture review. Compare with previous year's state | Manual | 

For automated daily review, a basic script that runs with cron:

```
#!/bin/bash
# /root/scripts/daily_audit.sh
# Run with: crontab -e -> 0 7 * * * /root/scripts/daily_audit.sh
LOG="/var/log/daily_audit_$(date +%Y%m%d).log"
echo "=== Daily audit $(date) ===" > "$LOG"
# Critical Suricata alerts in the last 24h
echo -e "\n--- Suricata alerts (severity 1) ---" >> "$LOG"
cat /var/log/suricata/eve.json | \
  jq -r 'select(.event_type=="alert" and .alert.severity==1) | 
  "\(.timestamp) \(.alert.signature) src=\(.src_ip) dst=\(.dest_ip)"' | \
  tail -50 >> "$LOG"
# Active CrowdSec decisions
echo -e "\n--- CrowdSec active decisions ---" >> "$LOG"
cscli decisions list -o raw 2>/dev/null | wc -l >> "$LOG"
# Age of last backup
echo -e "\n--- Last backup ---" >> "$LOG"
LAST_BACKUP=$(ls -t /conf/backup/*.xml 2>/dev/null | head -1)
if [ -n "$LAST_BACKUP" ]; then
  stat -f "%Sm" "$LAST_BACKUP" >> "$LOG"
else
  echo "No backups found" >> "$LOG"
fi
# System status
echo -e "\n--- System resources ---" >> "$LOG"
echo "CPU: $(sysctl -n dev.cpu.0.temperature 2>/dev/null || echo 'N/A')" >> "$LOG"
echo "RAM: $(sysctl -n vm.stats.vm.v_free_count)" >> "$LOG"
echo "States: $(pfctl -si 2>/dev/null | grep 'current entries')" >> "$LOG"
# Send via email or copy to syslog
logger -t daily_audit "Audit completed. See $LOG"
```
This script isn't meant to be a SIEM. It's a first line of automated review that flags if there's something requiring attention. For serious log analysis, the Grafana + Loki combination or a dedicated Wazuh is appropriate.

Before continuing to invest time in OPNsense, it's worth comparing with the other serious open source firewall/router options. Not to migrate now, but to know what's out there and make informed decisions if needs change.

VyOS is a network operating system based on Debian with a CLI configuration model inspired by JunOS. It has no web interface.

What it does better than OPNsense:

- **Plain text configuration** . VyOS config is a readable text file, with atomic commits and native rollback. It's the IaC dream:`git diff` works directly on the configuration.
- **Official Terraform provider** (Foltik/vyos). Real declarative network infrastructure.
- **Advanced routing** . BGP, OSPF, IS-IS at scale. Supports routing tables with more than one million BGP prefixes.
- **Performance** . The VPP dataplane enables 10 Gbps+ throughput with appropriate hardware.
- **Cloud-native** . Direct deployment on AWS, Azure, and GCP with official support.

What it does worse:

- **No GUI** . Everything is CLI or API. The learning curve is steep if you don't come from enterprise networking.
- **Limited integrated security** . There's no equivalent to Suricata with GUI, no DPI like Zenarmor, no integrated CrowdSec. Suricata can be installed manually, but without the integration OPNsense offers.
- **Paid LTS** . Since 2024, stable LTS images require a subscription. Rolling releases are free but without stability guarantee.

OpenWrt is a Linux-based router operating system. Its strength is embedded hardware support.

What it does better than OPNsense:

- **Massive hardware support** . Runs on more than 500 commercial router models, plus x86.
- **Native WiFi** . Directly manages wireless interfaces, with excellent driver support and advanced AP configuration.
- **Lightweight** . Can run on 128 MB RAM and 16 MB flash on embedded hardware.
- **Package ecosystem** . More than 27,000 packages available.

What it does worse:


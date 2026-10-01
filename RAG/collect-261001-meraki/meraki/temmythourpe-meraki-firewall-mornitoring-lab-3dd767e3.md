---
id: collect-261001-meraki/meraki/temmythourpe-meraki-firewall-mornitoring-lab-3dd767e3
title: "temmythourpe-meraki-firewall-mornitoring-lab-3dd767e3"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["sandbox"]
source: docs/RAG/collect-261001-meraki/temmythourpe-meraki-firewall-mornitoring-lab-3dd767e3.md
source_anchor: ""
source_lines: [1, 49]
sha256: bd6b018e0e929e493dc35bd53f0cbd4b4f2c549d6ca05e41504fa0df7534f7a7
---

# temmythourpe-meraki-firewall-mornitoring-lab-3dd767e3

This project demonstrates how to configure and monitor enterprise firewalls using the Cisco Meraki Dashboard.
The lab was built with the Meraki DevNet Sandbox and Dashboard API, without requiring physical hardware.
The goal is to show:
- Firewall configuration (L3/L7 rules, content filtering, geo-blocking).
- System monitoring (alerts, webhooks, syslog).
- Automation via Postman and Python SDK.
- Policy-as-code approach (YAML → API push).
- Cisco Meraki DevNet Sandbox / Demo Dashboard
- Meraki Dashboard API + API Key
- Postman (Meraki collections)
- Python 3.x + meraki SDK (pip install meraki)
- Flask (for webhook receiver)
- GitHub for version control & documentation
- Layer 3 (L3): Deny RFC1918 outbound, block risky ports, allow HTTP/HTTPS.
- Layer 7 (L7): Block P2P, gaming, malware domains, and specific categories.
- Content Filtering: Talos security categories enabled.
See policy.yaml for rules-as-code.
- Webhooks: Alerts sent to a custom Flask receiver.
- Syslog: Security events forwarded to SIEM (optional: Wazuh).
- Dashboard: Client/device analytics, event logs, usage trends.
- Postman: Test APIs (list orgs, networks, firewall rules).
- Python:
  - scripts/list_assets.py → enumerate orgs/networks/devices.
  - scripts/push_firewall_policy.py → apply rules from YAML.
  - scripts/export_events.py → save security events to JSON/CSV.
This lab was built using the Cisco Meraki DevNet Always-On Sandbox.
The sandbox allows API authentication and network queries, but write operations (such as updating firewall policies) return 403 Forbidden because it only gives read-only for firewall updates.
This is expected behavior for the Always-On environment.
In a real Meraki environment (or in a Reserved Sandbox), the same automation would successfully apply the firewall rules.
- Sandbox API is temporary
- Webhook receiver uses sample events
- Real webhooks will only work while the sandbox exists.
- Scripts demonstrate automation and policy-as-code workflow
- Offline replay demonstrates monitoring functionality without live Meraki resources.
The Postman collection demonstrates a realistic automation workflow:
- Discover
GET List Organizations → verify available organizations.
GET List Networks → verify available networks.
- Configure Firewall Rules
PUT Push L3 Firewall Rules → apply L3 rules from policy.yaml.
PUT Push L7 Firewall Rules → apply L7 rules from policy.yaml.
- Monitor & Test Events
POST Send Sample Webhook Event → test the local Flask receiver (app.py).
Verify the events are logged locally in webhook_events.log.
** Notes: Due to the Meraki Always-On Sandbox: **
Firewall write operations return 403 Forbidden (L3 & L7).
Content Filtering API endpoints are unavailable (would return 404).
These limitations are expected and noted as part of the project.
I designed the project around a standard enterprise topology, featuring a LAN, Guest Wi-Fi, and DMZ, all protected behind a Meraki MX firewall. While the DevNet Sandbox didn’t allow me to spin up the real VLANs, I still implemented the policies as if those networks existed. That demonstrates my ability to map real-world network diagrams to API-driven automation workflows. Moreover, I utilized Postman as both a testing platform and a proof-of-concept for automation with Cisco Meraki. Instead of manually configuring firewalls in the dashboard, I leveraged the Meraki API via Postman to fetch organizations, list networks, and push firewall rules. This demonstrated how to replace repetitive manual work with API-driven automation for scalability, consistency, and integration with tools like SIEM or SOAR.

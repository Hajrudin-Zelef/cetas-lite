---
id: collect-261001-fortinet/fortinet/t5-fortigate-troubleshooting-tip-fortiguard-ddns-ip-update-fails-ta-p-198179-d377a834-2
title: "t5-fortigate-troubleshooting-tip-fortiguard-ddns-ip-update-fails-ta-p-198179-d377a834"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["cybersecurity"]
source: docs/RAG/collect-261001-fortinet/t5-fortigate-troubleshooting-tip-fortiguard-ddns-ip-update-fails-ta-p-198179-d377a834.md
source_anchor: ""
source_lines: [252, 256]
sha256: 613b653f917bec185a79fa5b18504596ea8f879282abf221bbc88902196c9dca
---

# t5-fortigate-troubleshooting-tip-fortiguard-ddns-ip-update-fails-ta-p-198179-d377a834

If the ISP link is getting changed with a new public-ip and the DDNS resolving to that entry also needs to be changed, reach out to the TAC team to delete the old DDNS entry from the database.

After that, under DDNS settings via CLI, delete the copy config for the previous one and delete that entry. Once done, paste the copied configuration and only change the attribute for set monitor-interface to the new WAN port.

The Fortinet Security Fabric brings together the concepts of convergence and consolidation to provide comprehensive cybersecurity protection for all users, devices, and applications and across all network edges.

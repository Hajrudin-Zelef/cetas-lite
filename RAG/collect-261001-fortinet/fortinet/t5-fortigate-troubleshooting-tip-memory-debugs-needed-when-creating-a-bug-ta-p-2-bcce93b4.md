---
id: collect-261001-fortinet/fortinet/t5-fortigate-troubleshooting-tip-memory-debugs-needed-when-creating-a-bug-ta-p-2-bcce93b4
title: "t5-fortigate-troubleshooting-tip-memory-debugs-needed-when-creating-a-bug-ta-p-2-bcce93b4"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["memory", "distribution"]
source: docs/RAG/collect-261001-fortinet/t5-fortigate-troubleshooting-tip-memory-debugs-needed-when-creating-a-bug-ta-p-2-bcce93b4.md
source_anchor: ""
source_lines: [1, 4]
sha256: 575da2ebe016cbec7a3d92a91e8b4ada82c1a769cc9ab24079da133d5bc1b710
---

# t5-fortigate-troubleshooting-tip-memory-debugs-needed-when-creating-a-bug-ta-p-2-bcce93b4

Troubleshooting Tip: Memory debugs needed when creating a bug ticket
| Description | This article lists the CLI commands that should be provided to developers when creating an internal Bug ID related to a suspected FortiGate memory leak. | 
| Scope | FortiGate v7.2.6+, FortiGate v7.4.0+, FortiGate v7.6.0+. | 
| Solution | CLI commands:  get system status fnsysctl date get sys perf status diagnose autoupdate version diagnose sys session stat diagnose hardware sysinfo memory diagnose hardware sysinfo slab diagnose sys vd list \| grep fib diagnose sys mpstat 1 5 diagnose sys top-all 2 30 diagnose sys top-mem 10 diagnose firewall packet distribution diagnose snmp ip frags fnsysctl df -k fnsysctl ls -l /tmp fnsysctl du -i /tmp fnsysctl du -a /tmp fnsysctl du -a / -d 1 fnsysctl du -i /dev/shm fnsysctl du -a /dev/shm fnsysctl du -i /node-scripts fnsysctl du -a /node-scripts  Note: Super Admin privileges are necessary to run the 'fnsysctl' command. Otherwise, FortiGate will return an error. For more details, see Technical Tip: fnsysctl command returns Unknown action 0.  It is also possible to attach the output of the commands from this related article: Technical Tip: Memory Debugs |

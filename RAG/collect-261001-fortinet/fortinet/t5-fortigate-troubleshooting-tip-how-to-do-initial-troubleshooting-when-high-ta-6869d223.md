---
id: collect-261001-fortinet/fortinet/t5-fortigate-troubleshooting-tip-how-to-do-initial-troubleshooting-when-high-ta-6869d223
title: "t5-fortigate-troubleshooting-tip-how-to-do-initial-troubleshooting-when-high-ta--6869d223"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["memory", "throughput"]
source: docs/RAG/collect-261001-fortinet/t5-fortigate-troubleshooting-tip-how-to-do-initial-troubleshooting-when-high-ta--6869d223.md
source_anchor: ""
source_lines: [1, 4]
sha256: 1991b7f7063afa0ef2afbcae488ee30d179d697f5534a0e8c5b71ec3ed80e7ff
---

# t5-fortigate-troubleshooting-tip-how-to-do-initial-troubleshooting-when-high-ta--6869d223

Troubleshooting Tip: How to do initial troubleshooting when high CPU usage is observed
| Description | This article describes general actions which can be taken and which information should be sent to Fortinet Support in the case of an unexpected increase in CPU usage. | 
| Scope | FortiGate, FortiProxy. | 
| Solution |   get system performance status CPU states: 8% user 3% system 0% nice 87% idle 2% iowait 0% irq 0% softirq CPU0 states: 8% user 3% system 0% nice 87% idle 2% iowait 0% irq 0% softirq Memory: 2005244k total, 816796k used (40.7%), 1030464k free (51.4%), 157984k freeable (7.9%) Average network usage: 120 / 18 kbps in 1 minute, 259 / 38 kbps in 10 minutes, 194 / 29 kbps in 30 minutes Maximal network usage: 804 / 146 kbps in 1 minute, 804 / 146 kbps in 10 minutes, 804 / 146 kbps in 30 minutes Average sessions: 73 sessions in 1 minute, 59 sessions in 10 minutes, 44 sessions in 30 minutes Maximal sessions: 105 sessions in 1 minute, 105 sessions in 10 minutes, 105 sessions in 30 minutes Average session setup rate: 2 sessions per second in last 1 minute, 4 sessions per second in last 10 minutes, 3 sessions per second in last 30 minutes Maximal session setup rate: 23 sessions per second in last 1 minute, 23 sessions per second in last 10 minutes, 23 sessions per second in last 30 minutes Average NPU sessions: 0 sessions in last 1 minute, 0 sessions in last 10 minutes, 0 sessions in last 30 minutes Maximal NPU sessions: 0 sessions in last 1 minute, 0 sessions in last 10 minutes, 0 sessions in last 30 minutes Virus caught: 0 total in 1 minute IPS attacks blocked: 0 total in 1 minute Uptime: 0 days,  0 hours,  0 minutes Run the command above a few times and compare patterns of CPU usage, throughput, and the sessions' setup rates.      Example screenshot:  In this particular case, eap_proxy (process) use 99.9% of CPU. The commands below will provide more CPU information related to the user process. In this case, 1130 is the process ID of eap_proxy:  diagnose sys process dump 1130 diagnose sys process pstack 1130 diagnose sys process trace 1130    diagnose sys profile cpumask X <----- Where X is the CPU core with the highest CPU usage in the system space. diagnose sys profile start  Wait 20-30 seconds:  diagnose sys profile stop diagnose sys profile show order    Average NPU sessions: 0 sessions in last 1 minute, 0 sessions in last 10 minutes, 0 sessions in last 30 minutes Maximal NPU sessions: 0 sessions in last 1 minute, 0 sessions in last 10 minutes, 0 sessions in last 30 minutes  Compare with the total sessions:  Average sessions: 73 sessions in 1 minute, 59 sessions in 10 minutes, 44 sessions in 30 minutes Maximal sessions: 105 sessions in 1 minute, 105 sessions in 10 minutes, 105 sessions in 30 minutes  Most of the sessions should be offloaded.  Run the command 'diagnose hardware sysinfo interrupts' multiple times.  Add the command 'diagnose sys profile report' on Teraterm or Auto Script for intermittent issues.  Attach all of the outputs to the support ticket.  Related articles: Troubleshooting Tip: Best use for the 'diagnose sys profile report' command Troubleshooting Tip: How high CPU usage should be investigated Troubleshooting Tip: FortiGate CPU Profiling |

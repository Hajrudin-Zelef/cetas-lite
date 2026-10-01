---
id: collect-261001-huawei/huawei/questions-22218-b12c1dac
title: "questions-22218-b12c1dac"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["attention", "ethernet"]
source: docs/RAG/collect-261001-huawei/questions-22218-b12c1dac.md
source_anchor: ""
source_lines: [1, 4]
sha256: 06b4ec7cbeee570de367c497cf831ac801d60ff48291be8be418cf42a88d049d
---

# questions-22218-b12c1dac

We need to monitor the download/upload bandwidth on interfaces on Huawei S5300 switches, and we need to do it with a specific monitoring system. This system can get information from a switch using SNMP OID and can draw historic graphs based on collected data.
The issue is that there is no standard OID for bandwidth. For instance, the well known IF-MIB and its ifTable (1.3.6.1.2.1.2.2.) can return values for the number of bytes going in/out of the interface (.16-ifOutOctets and .10-ifInOctets) but they aren't useful since they need to be processed in a formula to get the actual bandwidth. The monitoring system we use has not built in option to calculate bandwidth using the information from IF-MIB.
As a solution, we're looking for Huawei custom MIB that can give us data about the current bandwidth usage of an interface. Is there a way to get the "Last 300 seconds input rate" and "Last 300 seconds ourput rate" in bits/sec for Ethernet and GigabitEthernet interface on S5300 switches using SNMP? Practically, we need the MIB and SNMP OID that relates to the input/output rate that is shown using the command: "display interface GigabitEthernet 0/0/X", as shown at the image below:
Thank you for your attention to this matter.

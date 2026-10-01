---
id: collect-261001-huawei/huawei/questions-1075852-38fc9ff0
title: "questions-1075852-38fc9ff0"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/questions-1075852-38fc9ff0.md
source_anchor: ""
source_lines: [1, 11]
sha256: 6392ecf0f3c33041433f0efb3f88fb11f7cd0d04f09f5cae28f1b077cdc138e8
---

# questions-1075852-38fc9ff0

I'm working with NagiosCore to check the stats of my network's switchs (mainly S5700 (modulars and non-modulars) and S5600). Using SNMP (v2c) I'm able to get stuff defined by myself such as: DeviceName, Location, Contact, and Description of the ifaces (ports) I'm currently using, as too the status (up/down) of those mentioned ifaces.
Now I would love to also check CPU_usage, MEM_usage, Temperature, PowerSupply_Status, and Bandwith_usage. I can do it trough snmpget using the OIDs I get from Huawei Infoweb but only for the S5700, I can't find any information in that web (or in any other web) related to the S5600's OIDs, I don't know if the S5600 is discontinued or what, but it is not showing on any of the Huawei Infoweb webs...
S5600 not showing up in Huawei Infoweb
Sooo... I would like to know the S5600's OIDs related to:
- CPU_usage
- MEM_usage
- Temperature
- PowerSupply_Status
- Bandwith_usage
Thanks in advance
.1.3.6.1.2.1.25.3.3.1.2, but doesn't start with the right code, when should be:.1.3.6.1.4.1.2011.6.1.1.1.2.65536, starting with the.1.3.6.1.4.1.2011.. Thanks anyways!!!

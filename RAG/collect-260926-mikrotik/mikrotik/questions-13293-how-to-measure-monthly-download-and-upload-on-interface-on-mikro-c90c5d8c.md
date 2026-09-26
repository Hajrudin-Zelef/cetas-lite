---
id: collect-260926-mikrotik/mikrotik/questions-13293-how-to-measure-monthly-download-and-upload-on-interface-on-mikro-c90c5d8c
title: "questions-13293-how-to-measure-monthly-download-and-upload-on-interface-on-mikro-c90c5d8c"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["license"]
source: docs/RAG/lot-mikrotik/forum/misc/questions-13293-how-to-measure-monthly-download-and-upload-on-interface-on-mikro-c90c5d8c.md
source_anchor: ""
source_lines: [1, 8]
sha256: 6b7f048cf4ad950aefe57c7648da1c45d6720c1f6f3835bbe3db55957556f09c
---

# questions-13293-how-to-measure-monthly-download-and-upload-on-interface-on-mikro-c90c5d8c

While most networking equipment has the ability to see interface statistics (like traffic over a period of time), these counters are generally used for troubleshooting or while you're actively in the switch.  For longer term data from an interface like what you're looking for I strongly suggest a monitoring tool that utilizes SNMP for the job.
Mirkotik has a specific example using MRTG here:
http://wiki.mikrotik.com/wiki/SNMP_MRTG
I find this example to be overly complicated but not knowing your environment it's hard to say what you will think.  Personally I suggest, a free tool if you only have this one use case, like PRTG which can be downloaded for free, and remains free unless you monitor more than 30 things with it.
https://shop.paessler.com/shop/standalone_free_license/
Once you have your tool of choice installed, setting up SNMP can be as simple as setting a community string, and setting a target for SNMP traps.  Details on configuring SNMP for your device are here:
http://www.mikrotik.com/testdocs/ros/2.9/root/snmp.php
If you have need of more monitoring than just this one interface, you may want to investigate other tools to fit your specifications.

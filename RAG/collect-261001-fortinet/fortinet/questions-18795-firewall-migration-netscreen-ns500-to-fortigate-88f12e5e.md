---
id: collect-261001-fortinet/fortinet/questions-18795-firewall-migration-netscreen-ns500-to-fortigate-88f12e5e
title: "questions-18795-firewall-migration-netscreen-ns500-to-fortigate-88f12e5e"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-fortinet/questions-18795-firewall-migration-netscreen-ns500-to-fortigate-88f12e5e.md
source_anchor: ""
source_lines: [1, 16]
sha256: 041b111406a3c0bb658fba161cecc142071d153391677fb4aa9cad3abce08565
---

# questions-18795-firewall-migration-netscreen-ns500-to-fortigate-88f12e5e

Information Security is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
2
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have recently been tasked with doing a firewall migration from an old NetScreen NS500 box to a FortiGate device (I believe it will be a 60C). I am sure most of you aware that NetScreen boxes are no longer sold since Juniper took over, so finding documentation has been a little difficult for me. Basically, I am wondering if anyone has any advice or possibly tools to help with this migration process. Doing this manually is proving to be a nightmare!
Note that I have tried https://convert.fortinet.com however it does not support this old NetScreen config, but it does support the newer Juniper configs. If there was someway of going NetScreen -> Juniper that would allow me to use the FortiConverter, but that also means another step which inherently means that I will incur more variation/errors in the migration process.
Let me know what you guys think, really any advice at all would be good at this point.
When you running old, unsupported versions, the vendor usually doesn't want to know, especially when you're moving from them. However, if you have a support contract with Fortinet (sorry I'm not familiar with them much), I'd have thought that your Account Manager would've been helpful?
Regarding general firewall migrations, 9 times out of 10, I found manual to be better, although more tedious. When the upgrade was version upgrade within the same product line, it was much easier and straight-forward with the latest version/package from the vendor. Otherwise it was more painful and required copying each rule over (taking notes for the 5-tuple and any other information).
I would support the manual approach @MarkHillick mentioned. I have helped organisations (one with over 600 firewalls and many thousand rules) migrate to new firewalls, and the only way to do it which provides review, auditability, checks on ownership of rules and a straightforward backout plan is to do this manually.
It takes time, but trying to automate the process of moving from a legacy rulebase is likely to not only propagate existing weaknesses, but introduce more.
Try two steps process,
There might be support from juniper or in forums to migrate from older netscreen to latest netscreen configs and from latest netscreen config to fortinet .
This requires rigorous testing though for correctness

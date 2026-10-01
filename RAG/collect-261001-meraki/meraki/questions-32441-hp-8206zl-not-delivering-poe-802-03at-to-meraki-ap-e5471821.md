---
id: collect-261001-meraki/meraki/questions-32441-hp-8206zl-not-delivering-poe-802-03at-to-meraki-ap-e5471821
title: "questions-32441-hp-8206zl-not-delivering-poe-802-03at-to-meraki-ap-e5471821"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-meraki/questions-32441-hp-8206zl-not-delivering-poe-802-03at-to-meraki-ap-e5471821.md
source_anchor: ""
source_lines: [1, 9]
sha256: 1bebe180c25f4525d38435e17f5e150867d37272cd493a6c4bc57b01cb55e719
---

# questions-32441-hp-8206zl-not-delivering-poe-802-03at-to-meraki-ap-e5471821

I found the issue, it has to do with the device itself, specifically an Meraki MR42 AP. When the 802.03at Meraki APs first powers up, it does so in 802.03af mode, thus limiting the power it can draw. Then later in the boot sequence it tries to request more power to be delivered, and for some reason the HP switches does not understand to give it the full 33 watts it needs. The solution is to force the port to allocate power based on a configured value instead of what's negotiated over LLDP. The Meraki documentation shows you how to do it through the GUI, which in my book isn't that useful. Instead, the following commands does the exact thing but for several ports at once.
interface [port_range] poe-value 33
interface [port_range] poe-allocate-by value
interface [port_range] power-over-ethernet critical
The first command sets the "value" for the port to 33W, maximum under the IEEE 802.03at standard.
The second command forces the port to forget about LLDP and just give the device connected to the port whatever it chooses to consume, instead of limiting it.
The third command is probably not necessary, all it should do is to list the port as "critical", aka. tell the switch to shut something else down if it experiences shortage of available PoE power, but I left it in as Cisco Meraki specifically listed it in their fix. 
The most amusing thing is that the APs doesn't actually draw any more power after this, at least when they are on standby with no connected clients, but they probably operate better at higher loads, and you get rid of that annoying error in the Meraki dashboard.
https://documentation.meraki.com/MR/Monitoring_and_Reporting/MR34_Operates_in_Low_Power_Mode_on_HP_ProCurve_Switch

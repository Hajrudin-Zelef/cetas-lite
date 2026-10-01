---
id: collect-261001-huawei/huawei/questions-33541-350cfeb5
title: "questions-33541-350cfeb5"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/questions-33541-350cfeb5.md
source_anchor: ""
source_lines: [1, 10]
sha256: 5523fcf4086a128b74a1c1a143ce29fc8222b0822f9b3d88bef48635e3d123e6
---

# questions-33541-350cfeb5

How can I change a switched port into a routed port?
Compared to IOS, is there an equivalent of no switchport? If there is none, is there a workaround?
6.6 Switching an Interface to Layer 3 Mode
<Quidway> system-view
[Quidway] interface gigabitethernet 1/0/1
[Quidway-GigabitEthernet1/0/1] undo portswitch
[Quidway-GigabitEthernet1/0/1] ip address 10.10.10.10 255.255.255.0 
If you are asking how to turn an interface on a switch back to factory defaults, where you can hook up a cable and it will connect to the Native VLAN with no problems or issues, use the switchport mode access command to put the switch port into access mode. Access mode will negate any configured VLAN assignments and port security setup on it. To assign the switchport to a VLAN use switchport access vlan x command which of course will assign that port to whatever VLAN you replace with X.
If you are asking about something on a router then the switchport command doesn't really apply as you can't have switchport enabled on a router interface.
Vlanif.

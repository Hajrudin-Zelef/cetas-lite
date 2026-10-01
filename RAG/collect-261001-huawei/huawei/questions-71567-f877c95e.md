---
id: collect-261001-huawei/huawei/questions-71567-f877c95e
title: "questions-71567-f877c95e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["copyright", "ethernet", "memory"]
source: docs/RAG/collect-261001-huawei/questions-71567-f877c95e.md
source_anchor: ""
source_lines: [1, 30]
sha256: d77110815d0add95ea0bd3590a985b15777040e6486114e89e0be5fa0b87702d
---

# questions-71567-f877c95e

I can't set an IP address on a Huawei AR207V router because it rejects the "ip address" command.
It's a used but factory reset unit, I'm currently connected on the serial port no other connections.
[Huawei]display version
Huawei Versatile Routing Platform Software
VRP (R) software, Version 5.130 (AR200 V200R005C00SPC100)
Copyright (C) 2011-2013 HUAWEI TECH CO., LTD
Huawei AR207V Router uptime is 0 week, 0 day, 0 hour, 40 minutes
MPU 0(Master) : uptime is 0 week, 0 day, 0 hour, 40 minutes
SDRAM Memory Size    : 512     M bytes
Flash Memory Size    : 512     M bytes
MPU version information :
1. PCB      Version  : AR01SRU0B VER.B
2. MAB      Version  : 0
3. Board    Type     : AR207V
..snip
I can't update to latest firmware because it needs an IP first.
This is what happens when I try to set an IP on an interface -
Please configure the login password (<6-128>)
Enter password:
Confirm password:
<Huawei>system-view
Enter system view, return user view with Ctrl+Z.
[Huawei]interface Ethernet 0/0/0
[Huawei-Ethernet0/0/0]ip address 192.168.2.1 24
                     ^
Error: Unrecognized command found at '^' position.
There are plenty of instructions* which show setting an IP as above but this unit rejects the crucial command
What's going wrong here?
Thanks
(instructions for setting an IP - https://support.huawei.com/enterprise/en/doc/EDOC1100033773/403d9282/configuring-ip-addresses-for-interfaces )

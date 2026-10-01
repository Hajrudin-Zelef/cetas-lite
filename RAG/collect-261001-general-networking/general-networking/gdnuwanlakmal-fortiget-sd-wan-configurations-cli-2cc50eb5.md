---
id: collect-261001-general-networking/general-networking/gdnuwanlakmal-fortiget-sd-wan-configurations-cli-2cc50eb5
title: "gdnuwanlakmal-fortiget-sd-wan-configurations-cli-2cc50eb5"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/gdnuwanlakmal-fortiget-sd-wan-configurations-cli-2cc50eb5.md
source_anchor: ""
source_lines: [1, 36]
sha256: d6e9660c7f65f9d1114b3427e0f6b4cd31366658eaa82a22e743f05abe1aba2f
---

# gdnuwanlakmal-fortiget-sd-wan-configurations-cli-2cc50eb5

This guide provides CLI commands to configure SD-WAN on a FortiGate device, including adding ISP interfaces to WAN links.

```
config system interface 
edit "wan1"
    set alias to_ISP1
    set mode dhcp
    set distance 10
next
edit "wan2"
    set alias to_ISP2
    set ip 10.100.20.1 255.255.255.0
next
end
```
```
config system sdwan
    set status enable   
    config members
    
    edit 1
        set interface "wan1"
    next
    edit 2
        set interface "wan2"
        set gateway 10.100.20.2
    next
end
```
```
config router static
    edit 1
        set sdwan-zone "virtual-wan-link"
    next
end
```

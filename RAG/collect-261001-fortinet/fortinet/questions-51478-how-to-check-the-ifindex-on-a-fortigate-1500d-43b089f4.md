---
id: collect-261001-fortinet/fortinet/questions-51478-how-to-check-the-ifindex-on-a-fortigate-1500d-43b089f4
title: "How to check the ifIndex on a Fortigate 1500D?"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-fortinet/questions-51478-how-to-check-the-ifindex-on-a-fortigate-1500d-43b089f4.md
source_anchor: ""
source_lines: [1, 49]
sha256: 3b84e64e2d1c0c7c4357af2ebd514ec959265690e4891c96912775b6fdac8560
---

# How to check the ifIndex on a Fortigate 1500D?

*Score : 4 | Source : https://networkengineering.stackexchange.com/questions/51478/how-to-check-the-ifindex-on-a-fortigate-1500d*

What is the exact command to check the port ifIndex that is used by SNMP to reflect trap at SNMP host at 1500D Fortigate?
The command diagnose sys device list root displays the index which isn't unique for every port. 
The command diagnose hardware  deviceinfo nic  Portno also displays various hardware parameters related to port but not ifIndex.
How can I check the ifIndex on a Fortigate 1500D device?
Version: FortiGate-1500D v5.4.4,build1117,170209

---

### Reponse (acceptee) — score 4

The following should show you the snmp-index for each interface on the Fortigate. You can also modify this field if you so choose.
# config system interface (interface) # show
# config system interface (interface) # show

---

### Reponse — score 0

diagnose sys device list root
diagnose sys device list root
or
diagnose sys device list VDOM
diagnose sys device list VDOM
should do the trick.

---

### Reponse — score 0

Use the command:
diagnose netlink interface list
It will show you the index references for the interfaces.
Example:
if=port1 family=00 type=1 index=10 mtu=1500 link=0 master=0 ref=8 state=start present flags=up broadcast run promsic multicast

---

### Reponse — score 0

Use the command diag ip route list, where the "dev" is same "index"
tab=255 vf=0 scope=254 type=2 proto=2 prio=0 0.0.0.0/0.0.0.0/0->192.168.1.100/32 pref=192.168.1.100 gwy=0.0.0.0 dev=12(port10)
tab=255 vf=0 scope=253 type=3 proto=2 prio=0 0.0.0.0/0.0.0.0/0->192.168.1.255/32 pref=192.168.1.100 gwy=0.0.0.0 dev=12(port10)
tab=255 vf=0 scope=253 type=3 proto=2 prio=0 0.0.0.0/0.0.0.0/0->192.168.2.0/32 pref=192.168.2.127 gwy=0.0.0.0 dev=3(port1)
tab=255 vf=0 scope=254 type=2 proto=2 prio=0 0.0.0.0/0.0.0.0/0->192.168.2.127/32 pref=192.168.2.127 gwy=0.0.0.0 dev=3(port1)
tab=255 vf=0 scope=253 type=3 proto=2 prio=0 0.0.0.0/0.0.0.0/0->192.168.2.255/32 pref=192.168.2.127 gwy=0.0.0.0 dev=3(port1)

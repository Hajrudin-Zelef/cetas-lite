---
id: collect-261001-fortinet/fortinet/questions-23521-dell-powerconnect-vlan-trunking-with-a-fortigate-63be36f8
title: "questions-23521-dell-powerconnect-vlan-trunking-with-a-fortigate-63be36f8"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/questions-23521-dell-powerconnect-vlan-trunking-with-a-fortigate-63be36f8.md
source_anchor: ""
source_lines: [1, 4]
sha256: d36a3148cc7bd5235256980eee83297ebabf63e33206888e39ccf5062b615c6e
---

# questions-23521-dell-powerconnect-vlan-trunking-with-a-fortigate-63be36f8

I am using a Dell Powerconnect 6248 (layer 3 48 port gigabit Switch) and a Fortigate 310. I am attempting to use the switch in layer two mode and assign a vlan to each port where a server is attached and a "Trunk" line directly connected to the fortinet. on the Fortinet, I have virtual interfaces designed to match the vlan's for each port. The Goal here is to require any communications from server to serverto be evaluated by policiesdefined for the virual interfaces that corresponds to the vlans for the servers attempting to communicate.
So essentially, server A wants to talk to server B. Server A is in port 2 (defined as vlan 2) and Server B is in port 3 (defined as vlan 3). The firewall is attached to the switch port 1, and has virtual interface A (defined as vlan2) and virtual interface B (defined as vlan3). I have a policy that allows virtual interface A to communicate with virtual interface B. All of the server Apackets should flow up to the firewall and be passed back to the switch with a new vlan tag to go to the server B.
My question is, should port 1 on the switch be defined as a "TRUNK" and if so what is the method for doing so (docs are weak)? is there anything else I need to consider here?
THANKS!

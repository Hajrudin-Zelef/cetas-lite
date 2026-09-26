---
id: collect-260926-mikrotik/mikrotik/questions-1620380-mikrotik-router-how-can-i-to-setup-a-local-wifi-network-in-rou-e6aa4c49
title: "questions-1620380-mikrotik-router-how-can-i-to-setup-a-local-wifi-network-in-rou-e6aa4c49"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-1620380-mikrotik-router-how-can-i-to-setup-a-local-wifi-network-in-rou-e6aa4c49.md
source_anchor: ""
source_lines: [1, 10]
sha256: 39ed106e7743bb4b88858d3f7ca78c4cf49ee220ea93aa3c22d0b82fc0a76df7
---

# questions-1620380-mikrotik-router-how-can-i-to-setup-a-local-wifi-network-in-rou-e6aa4c49

Super User is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
Does anyone have any idea how to setup a local wifi network in RouterOS v6.44.6 (from the scratch)? If you have any questions, feel free to ask! I appreciate your help!
But I don't have the "Master Port tab", like it is shown in the instruction.
The instructions are out of date. Recent RouterOS versions do not have this option anymore; instead they automatically configure switching for ports which belong to the same bridge. So you should create a bridge via /interface bridge, then add all of your LAN ports to it – you'll see them showing up with an H flag indicating that hardware offloading is active (i.e. the bridging is done by a switch).
Don't forget to add the Wi-Fi interface to the bridge as well. (It won't have the H flag listed; that's normal.)
I also do not have a preconfigured wifi network like here:
You don't need to have "Virtual AP" interfaces. They're only needed if you want to host multiple SSIDs on the same radio, and you can add/remove them at any time (via the "+" button).
For a simple wireless network you only need to configure the actual hardware interface (the one that shows up with "Type: Wireless" – wlan1 in the picture).
(If your device is dual-band, it will have two "Wireless" interfaces and you'll need to configure both separately.)
However, if you do not have any interfaces with type "Wireless", then it is possible that your router simply doesn't have Wi-Fi built-in. Not all routers have Wi-Fi!

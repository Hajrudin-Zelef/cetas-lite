---
id: collect-260926-mikrotik/mikrotik/questions-704131-extend-wifi-coverage-with-mikrotik-ed482a74
title: "questions-704131-extend-wifi-coverage-with-mikrotik-ed482a74"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "research"]
source: docs/RAG/lot-mikrotik/forum/wifi-capsman/questions-704131-extend-wifi-coverage-with-mikrotik-ed482a74.md
source_anchor: ""
source_lines: [1, 13]
sha256: e4f50af588af7295d3b94cb3b0da815462e6334e32a26eee680ae71f6e35f383
---

# questions-704131-extend-wifi-coverage-with-mikrotik-ed482a74

Super User is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
1
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have two Mikrotik routers (951g-2hnd) connected to each other via ethernet.
As I understand if I will be using WDS for extending WiFi coverage, my routers will be using WiFi for communicate between themselves, but I want to use ethernet link instead of WiFi, could someone provide example configuration?
WDS don't do anything special. I just allow bridging APs without wires.
Just bridge all wireless cards to same LAN. And ensure, this LAN has only one DHCP server enabled.
For example interconnect this RB951 by wires between 2 - 5 ports. Look for not build loops or enable RSTP in bridges. And disable dhcp server on all except one.
Also is good to set different IP for every device.
You can configure wireless in dewices with same identical SSID and Security profile to form one ESSID.

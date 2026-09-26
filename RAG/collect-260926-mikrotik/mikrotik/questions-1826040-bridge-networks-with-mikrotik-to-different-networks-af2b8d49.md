---
id: collect-260926-mikrotik/mikrotik/questions-1826040-bridge-networks-with-mikrotik-to-different-networks-af2b8d49
title: "questions-1826040-bridge-networks-with-mikrotik-to-different-networks-af2b8d49"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "research"]
source: docs/RAG/lot-mikrotik/forum/vlan-bridge/questions-1826040-bridge-networks-with-mikrotik-to-different-networks-af2b8d49.md
source_anchor: ""
source_lines: [1, 24]
sha256: 2db35c8fbb3ffb00a3c505ca51cac7db6ba098e5555103f1f6c7bbaefff6e24b
---

# questions-1826040-bridge-networks-with-mikrotik-to-different-networks-af2b8d49

Super User is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
0
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I am trying to bridge a network across building to building using a MikroTik SXTac and while I can do this with a standard format in the example 10.0.0.1/24 > 10.0.0.11/24 <===> 10.0.0.12/24 > 10.0.0.2 this is working. What I would like though is to use different bridging IP's for the MikroTik's. How can I accomplish this?
Just to clarify the question a bit more - the endpoints are not just those single computers - anything on that networks I want to reach.
See script below the image .
interface wireless set MyWLan1 disabled=no ssid=MySiddyID band=5ghz-onlyn mode=station
ip address add address=10.0.0.12/24 interface=MyWLan1
mpls ldp set enabled=yes lsr-id=10.0.0.12 transport-address=10.0.0.12
mpls ldp interface add interface=MyWLan1
interface vpls add name=vpls1 remote-peer=10.0.0.10 vpls-id=2:2 disabled=no
interface bridge add name=MyBridge1
interface bridge port add bridge=MyBridge1 interface=ether1
interface bridge port add bridge=MyBridge1 interface=vpls1
Without VPLS: You do it exactly the same way as usual. Bridge IP addresses have no relationship with the traffic being bridged; they're only there for management of the bridge host (and for LDP in your case). So if you want to use addresses from a different prefix, there's nothing special about that; it's basically as if you had two subnets sharing the same ethernet.
With VPLS: You still do it the same way as usual, as the interfaces that have the IP addresses you want to change aren't even part of the bridge anyway. (Which means they should not have had addresses matching the bridged network in the first place! What you have right now is two entirely separate networks – the WLAN-WLAN underlay network and the ether1-VPLS-ether1 overlay network – having overlapping addresses...) So there should be no problem with changing them to 172.16.99./24 or something else.
That being said, RouterOS already has built-in support for L2 bridging over wireless interfaces (that's what the SXTs are built for), so if your setup has just those two devices talking VPLS to each other, then I would suggest a simpler setup (unless of course the current method is required by your workplace or something like that). I'm not sure whether the SXT default configuration uses station-bridge or station-wds (the former is Mikrotik proprietary mode, while the latter is a kind-of semi-standard "4addr" mode), but either should work.
interface wireless set MyWLan1 disabled=no ssid=MySiddyID <...> mode=station-bridge
interface bridge add name=MyBridge1
interface bridge port add bridge=MyBridge1 interface=ether1
interface bridge port add bridge=MyBridge1 interface=MyWLan1

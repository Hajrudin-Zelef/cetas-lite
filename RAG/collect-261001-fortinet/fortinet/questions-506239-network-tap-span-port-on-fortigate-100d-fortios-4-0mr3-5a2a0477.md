---
id: collect-261001-fortinet/fortinet/questions-506239-network-tap-span-port-on-fortigate-100d-fortios-4-0mr3-5a2a0477
title: "questions-506239-network-tap-span-port-on-fortigate-100d-fortios-4-0mr3-5a2a0477"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-fortinet/questions-506239-network-tap-span-port-on-fortigate-100d-fortios-4-0mr3-5a2a0477.md
source_anchor: ""
source_lines: [1, 42]
sha256: fed05fe9b5132ea0b45ac2030a3111c5f370a2eb6c4a0ab3b3495b5c0a51d27c
---

# questions-506239-network-tap-span-port-on-fortigate-100d-fortios-4-0mr3-5a2a0477

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
2
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I'm dealing with a FortiGate 100D for the first time, and am scratching my head as there doesn't seem to be an easy way to mirror ports in the switch; which is really a facility that I presumed it would provide.
Ideally, I want to mirror one (or more) ports to another port, so that I can track the traffic that is flowing through it.
i.e.
mirror WAN1 to an internal port
mirror an internal port to a different internal port
etc.
I could do it with a passive network tap, of course; but it seems really strange to me that the 100D doesn't seem to expose an easy way to do this.
I'm new to the hardware/FortiOS, though -- so possibly I am simply missing something obvious.
Many thanks if someone can point me in the direction of how to set this up on FortiOS/FortiGate.
From the FortiOS CLI reference, under system > switch-interface:
config system switch-interface
edit <group_name>
set member <iflist>
set span {enable | disable}
set span-dest-port <portnum>
set span-direction {rx | tx | both}
set span-source-port <portlist>
set type {hub | switch | hardware-switch}
set vdom <vdom_name>
end
The Switch Port Analyzer (SPAN) feature is now available for hardware switch interfaces on FortiGate models with built-in hardware switches (for example, the FortiGate-100D, 140D, and 200D etc.)
To enable SPAN on a hardware switch via the GUI, go to System > Network > Interfaces and edit a hardware switch interface.
By default the system may have a hardware switch interface called LAN. A new hardware switch interface can also be created.
Select the SPAN check box, then select a source port from which traffic will be mirrored.
Select the destination port to which the mirrored traffic is sent.
Select to mirror traffic received, traffic sent, or both.
SPAN can also be enabled in the CLI:
config system virtual-switch
edit <port>
set span enable
set span-source-port <port>
set span-dest-port <port>
set span-direction {both | Tx | Rx}
end
end
flow-inspectorand a robust flow capturing software by a company qosient calledargus. Both are free. Check them out.

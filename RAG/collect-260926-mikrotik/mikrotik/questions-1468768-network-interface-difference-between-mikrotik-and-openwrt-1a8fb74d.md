---
id: collect-260926-mikrotik/mikrotik/questions-1468768-network-interface-difference-between-mikrotik-and-openwrt-1a8fb74d
title: "questions-1468768-network-interface-difference-between-mikrotik-and-openwrt-1a8fb74d"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "research"]
source: docs/RAG/lot-mikrotik/forum/misc/questions-1468768-network-interface-difference-between-mikrotik-and-openwrt-1a8fb74d.md
source_anchor: ""
source_lines: [1, 21]
sha256: 7fda4a4b7de03a21983014263ae2e034f1aecfddb9d223117e77a6a747d0ea22
---

# questions-1468768-network-interface-difference-between-mikrotik-and-openwrt-1a8fb74d

Super User is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
1
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I've started using OpenWRT on one old dsl modem and I'm a little confused about how interfaces are defined.
On Mikrotik:
there is one or more switches which contain several interfaces. I can do whatever I want with them: assign them to bridge, leave them alone and add their own address etc. And the most important is I do not create VLAN to manage bridge!
On the other hand in OpenWRT:
It seems the whole idea of bridge is build on VLAN. There is only one physical interface for switch. I do not understand how can I manipulate with each port without VLAN tagging.
Let's say I have 4 ports in switch. I want to make bridge from 1-3 ethernet ports plus wlan and add own address on 4th port to be WAN interface. How can I manage this the same way as on Mikrotik without VLAN?
I understand they are two different approaches to the same function but which one is better and more native for linux enviroment?
Edit: Driver name of switch in OpenWRT is bcm63xx_enetsw
Mikrotik is a particular brand, and your link shows a particular kind of switch hardware (QCA8337). OpenWRT on the other hand is a generic operating system, runs on many different kinds of routers, and handles lots of different hardware that implement a switch.
Looking at the data sheet of the QCA8337, this chip also uses VLAN tagging, and looking at your picture, there's only a single link to the CPU, and I doubt that the CPU on a SoC has more than two ethernet ports. So my guess is that Mikrotik also uses VLAN tagging for your switch, it's just that the GUI hides this fact from you.
[In OpenWRT], it seems the whole idea of bridge is build on VLAN.
This is nonsense. You can use the switch hardware in whatever way the hardware works. If you are using OpenWRT on your Mikrotik, then your switch uses VLAN tagging, so you'll have to deal with that. That's actually more flexible than the Mikrotik software which hides this fact from you, because you can use swconfig and freely configure the switch hardware in whatever way you want, so you could, say, combine two ports each in the switch, without needing the CPU to handle that. And yes, if you want to "import" single ports into the CPU, they'll need to be tagged in the switch hardware, and detagged by the CPU.
A Linux kernel bridge can contain whatever network interfaces you like. It's not limited to VLAN interfaces. On hardware which directly exposes ports to the CPU, you can just include those into a bridge.
which one is better and more native for linux enviroment
They are exactly the same. It's just that the Mikrotik UI (very likely, I haven't seen the UI) manages the details for you, and doesn't let you see them.

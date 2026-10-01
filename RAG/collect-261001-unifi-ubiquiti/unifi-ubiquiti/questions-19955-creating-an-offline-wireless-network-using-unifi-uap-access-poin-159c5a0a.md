---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-19955-creating-an-offline-wireless-network-using-unifi-uap-access-poin-159c5a0a
title: "questions-19955-creating-an-offline-wireless-network-using-unifi-uap-access-poin-159c5a0a"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["research", "throughput"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-19955-creating-an-offline-wireless-network-using-unifi-uap-access-poin-159c5a0a.md
source_anchor: ""
source_lines: [1, 18]
sha256: 5a2d331442f71e00bc6b1540fa52386e7a25a168b80d2f951e91593e284afe26
---

# questions-19955-creating-an-offline-wireless-network-using-unifi-uap-access-poin-159c5a0a

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
3
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have to confess that I am by no means a network engineer I am a software engineer with a reasonable grasp of networking.
Here is my dilemma...
I have 20 iPads that I want to put onto a wireless network that has no internet connection. These iPads are connected to a Macbook Pro using Bonjour and they have data pushed/broadcast to them from said Macbook.
If I create a network with a home based wireless router/switch the devices all communicate flawlessly and the data is pushed to the iPads with ease and with an almost instant response time.
But... If I create a wireless mesh with 3 UniFi Access points which are connected to a Pro Curve switch. I can get all the iPads to connect to the network with static IPs but when I broadcast data to them the response is very slow. Some of the iPads randomly drop connection and ping times between devices can range from 100ms to 600ms.
I assume I possibly need something on the network to possibly help route the traffic but this is where my knowledge runs short.
Could anyone point me in the right direction? What am I doing wrong?
Mesh is an inaccurate term for how they work, by the way(unless you put non-factory firmware on them). They CAN be configured in a "single-leaf" wireless uplink mode, but should not be (it cuts throughput by at least half .vs. wiring them all.) You seem to be saying they are all wired into a switch, so there is no "mesh" going on, simply provision of the SSID over multiple devices.
A common problem (not clear from the level of detail in your question) is people deploying them with WAY too much power and overloading the input stages of the client devices (especially if they get the UAP-LR because it's advertised to have "more range") - in MANY cases, performance is improved by turning power down. The link does not work unless the radio can also hear the client device, and UAP-LRs have a VERY small field of appropriate application.
That's one fork of possible issues/resolutions. The other is:
A "home type router/switch" is a wireless access point, a switch, and a router, typically including a DHCP server, DNS Server, etc.
So, if it works better with a "home-type router/switch" you might need a device to play the role of "router" though in fact your issue may be more DHCP/DNS than "actual routing." I'd suggest pfSense for the role, but there are many possible options.

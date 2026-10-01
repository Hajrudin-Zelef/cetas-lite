---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-834673-wifi-hotspot-ubiquiti-configuring-bullet-m2-titanium-7575a659
title: "questions-834673-wifi-hotspot-ubiquiti-configuring-bullet-m2-titanium-7575a659"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "omni"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-834673-wifi-hotspot-ubiquiti-configuring-bullet-m2-titanium-7575a659.md
source_anchor: ""
source_lines: [1, 5]
sha256: 506d866306bea03ee199bc923fa2348509512f47d9628e95772ac535dc61e634
---

# questions-834673-wifi-hotspot-ubiquiti-configuring-bullet-m2-titanium-7575a659

Recently one of my clients wanted me to setup a Wifi Hotspot for him. This wifi hotspot will give access to internet to his users after some verifications and some other things. He also wanted small range on his Wifi Hotspot, so i bought the Bullet M2 Titanium along with a 7dbi Omni Antenna. He asked from me to create a Software for Windows Operating Systems, that it would control the users, and also control the pages the users will see before accessing the internet, So i though of using MyHotspot Software along with some custom applications i created my self to achieve this. So the basic idea is this.
2 Network Adapters on the Same Computer, 1 network adapter connecting to the Router, and one Connecting to the Bullet, And then bridge them with MyHotspot Software in order to control their access. However i'm not sure how shall i proceed with the Configuration of the Bullet M2, The default IP Address that MyHotspot Software gives to the adapter when you click the Bridge, is 192.168.137.1 , And it listens on that address on port 80, and redirects all traffic there.
Tested on Laptop, with Virtual Wifi Router Having the Ethernet for Internet and Wifi Adapter for New Connections and MyHotspot Software.
Can someone guide me to configure my setup in order to achieve this?
Users -> Bullet M2 -> Computer -> Router -> Internet.

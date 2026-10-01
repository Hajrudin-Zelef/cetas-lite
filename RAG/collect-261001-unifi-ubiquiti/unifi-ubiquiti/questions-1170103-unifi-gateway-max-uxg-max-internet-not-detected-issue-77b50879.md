---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1170103-unifi-gateway-max-uxg-max-internet-not-detected-issue-77b50879
title: "questions-1170103-unifi-gateway-max-uxg-max-internet-not-detected-issue-77b50879"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1170103-unifi-gateway-max-uxg-max-internet-not-detected-issue-77b50879.md
source_anchor: ""
source_lines: [1, 10]
sha256: c5b74966678a8acdcf73c0b6be619dc4184e69369ca339995cefe1ce4c07aa2d
---

# questions-1170103-unifi-gateway-max-uxg-max-internet-not-detected-issue-77b50879

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
0
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I am trying to replace the router in my house, going from Linksys WRT1900 to the UniFi Gateway Max (not Cloud). The ISP's modem is in bridge mode, and everything works fine with the old router, without any complicated setup (DHCP). When I connect UniFi, it seems to work fine for a short period of time, like 20 seconds - the setup tries to move forward and it even delivers internet to whatever is connected to its ports. But then (just after some seconds pass, not due to any actions) the picture changes to "No Internet Detected" message on both the android app and on the router's web interface, and no amount of restarts of the provider's modem would help. I tried factory-resetting the device, the ISP tried factory-resetting the modem and then gave me a completely new one (because they thought maybe the old one was a bit too slow or whatever). They see the UniFi's mac-address when it is connected to the modem.
Any ideas? I think I've exhausted all the combinations of what to restart in which order. Really wondering what criteria the UniFi device uses to detect "internet".
I got this working with some help from Ubiquiti's support. Apparently, with these "headless" devices (UXG-*, so, without the word "Cloud" in the product name), you are supposed to have the management server running and adopt the product into the Uniti system before it would connect. It might be obvious for anyone who dealt with Ubiquiti's products before, but it wasn't for me. Actually, I'd even say I was a bit tricked by how simple the "getting started" instructions are ("download app, plug in, follow instructions in app"). Moreover, it's a bit silly for the error message to be "no internet detected".
Just to clarify one more thing. The "non-cloud" edition of the device runs completely fine on its configuration without the management software constantly running in parallel. You need it running to make changes to your configuration, but you can just as well run it on your laptop that powers down regularly. At this moment I'm rather fascinated by how cool the software is, and kind of feel like I might like to have all the stats gathered constantly. But the price difference gives you one extra AP, for example. So, all in all i like Ubiquiti having these product options.

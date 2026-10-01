---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-71096-is-site-to-site-with-ubiquiti-udm-pro-usg-erpro-8-behind-isp-rou-f3a2dfeb
title: "questions-71096-is-site-to-site-with-ubiquiti-udm-pro-usg-erpro-8-behind-isp-rou-f3a2dfeb"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["latency", "research"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-71096-is-site-to-site-with-ubiquiti-udm-pro-usg-erpro-8-behind-isp-rou-f3a2dfeb.md
source_anchor: ""
source_lines: [1, 31]
sha256: c856d5e42ed04b99653d7be2def11bcf92b482c126dee117a224cd8abb71de3e
---

# questions-71096-is-site-to-site-with-ubiquiti-udm-pro-usg-erpro-8-behind-isp-rou-f3a2dfeb

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
0
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
Evaluating the capabilities of the UniFi Dream Machine Pro all-in-one enterprise security gateway & network appliance (UDM Pro), I was wondering whether this site-to-site setup is possible:
All three routers are behind ISP routers, which do support port-forwarding to the ubiquiti routers, but might not support bridge mode.
All three sites have dynamic IPs, referenced by dynamic DNS.
Setup goals
The three networks behind the Ubiquiti routers should be connected via site-to-site VPN, e.g. IPSec.
All UniFi devices, i.e. the Access Points (APs), the UDM Pro, and the USG, should be controlled by the UniFi controller on the UDM Pro.
Some observations
IPSec between several EdgeRouters only (without ISP routers, without UniFi routers) does work, but the UDM Pro interface did not allow to enter dynamic DNS names as IPSec peers.
When adding the ISP routers with port forwarding (UDP 500 and 4500), I think I would need to tell the EdgeRouters to use their dynamic public IPs when establishing IPSec for authentication, but I'm not sure if I can do this in the GUI.
From what I've seen, the USG and the UDM Pro would support dynamic DNS when using OpenVPN rather than IPSec, but the EdgeRouter does not support OpenVPN from the GUI.
The most uncertain thing to me is whether I can use the UDM's UniFi controller through the tunnel, especially because the UDM Pro is an appliance and I'm unsure whether it would support controlling multiple sites, or be controlled by an external UniFi controller.
It's possible, but first you need static IP addresses. You can ask your ISPs to give you static addresses (for a fee, of course).
The most uncertain thing to me is whether I can use the UDM's UniFi
controller through the tunnel...
Once you establish the tunnels, they essentially become transparent to the devices. Baring latency problems, your controllers can't tell the difference if you're using VPN tunnels or not.
As this question still receives views and comments, I just wanted to give a quick update. While I obviously agree with Ron that this is a non-desirable setup, we needed to implement something as a quick, but temporary measure to deal with a local ISP issue and some challenges of the home-office situation.
Solution 1: Manual Configuration
IPSec itself does, of course, support dynamic addresses. The issue is just that the UDM user interface does only allow static IPs as addresses. Configuring the IPSec connection manually on the UDM file system, does work however.
However, this configuration file is deleted or overwritten and lost on UDM reboot and firmware updates, which is why we've dropped this solution.
Solution 2: Adding ER-X
Because the Ubiquiti Edge Routers all support dynamic addresses, we've just replaced the USG with an ER-X-SFP, and also added an ER-X-SFP between the ISP router and the UDM. This way, the ER-X routers could handle the IPSec connections, and the rest (including the unifi controller) could be done on the UDM.
This requires no modifications of the configuration files at all; the IPSec connection can just be configured as usual through the Edge Routers' user interfaces. https://help.ui.com/hc/en-us/articles/115012831287
Remarks
I would personally consider the artificial requirement of static addresses a shortcoming of the UDM, but as this is not a common use case, I guess the UDM is just not made for this. (Usually, when requiring non-streamline configurations, I'm going for a pfSense device.)
What really went well, however, was the exposure of the unifi controller on the UDM through the IPSec VPN, making adding more and more unifi access points on all sites easy.

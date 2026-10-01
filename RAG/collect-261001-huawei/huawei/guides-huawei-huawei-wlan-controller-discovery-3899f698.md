---
id: collect-261001-huawei/huawei/guides-huawei-huawei-wlan-controller-discovery-3899f698
title: "guides-huawei-huawei-wlan-controller-discovery-3899f698"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/guides-huawei-huawei-wlan-controller-discovery-3899f698.md
source_anchor: ""
source_lines: [1, 15]
sha256: 8e5c2b0004d92905f28fe71da4456050c27197e7951f3da91b18cb55154c42be
---

# guides-huawei-huawei-wlan-controller-discovery-3899f698

Huawei access points can be operated in three ways: Fat, Fit, and as a Cloud AP. In this article, we will discuss the configuration in Fit mode and how the access point can discovery a WLAN controller (AC).
A Fit AP only works together with the WLAN controller (AC). There are two ways how the access point can find the WLAN controller: Static or Dynamic.
Static
The access point is configured with the IP address of the WLAN controller via serial interface:
<Huawei> system-view[Huawei] ap-address mode static[Huawei] ap-address static ac-list 192.168.1.10[Huawei] quit<Huawei> rebootSystem will reboot! Continue ? [y/n]:y 
Dynamic
In dynamic mode, there are three options: Broadcast, DNS or DHCP.
Broadcast
If the access points are in the same network as the WLAN controller, then the broadcast discovery will be successful and no further configuration is required. If the access points are not in the same network as the WLAN controller, then the next two options can be used.
DNS
The Huawei access points have a fixed DNS name integrated for the WLAN controller discovery. An attempt is made to resolve the DNS name "huawei-wlan-controller". If the access point has been given a domain name, e.g. via DHCP, then it searches for the FQDN, e.g. huawei-wlan-controller.netcamp.local.
DHCP
You can use the DHCP option 43 (for IPv4) and DHCP option 52 (for IPv6) to inform the access points about the IP address of the WLAN controller. The format of the DHCP option 43 is according to the following format:
In this example two IP addresses of WLAN controllers are supplied via DHCP. The option is configured in hex format on the DHCP server (e.g. Windows).
Leave a Comment

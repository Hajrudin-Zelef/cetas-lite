---
id: collect-260926-mikrotik/mikrotik/questions-19243-tablets-disconnecting-and-reconnecting-randomly-from-routerboard-d5f9b5ed-4
title: "questions-19243-tablets-disconnecting-and-reconnecting-randomly-from-routerboard-d5f9b5ed"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-19243-tablets-disconnecting-and-reconnecting-randomly-from-routerboard-d5f9b5ed.md
source_anchor: ""
source_lines: [244, 274]
sha256: 687f6ba422254758c5b4ba9f9bb3ed76cc479e8231cd54f1233d0cde99d981b4
---

# questions-19243-tablets-disconnecting-and-reconnecting-randomly-from-routerboard-d5f9b5ed

jun/15 18:56:43 wireless,info 0C:B3:19:29:08:13@wlan1_My_Companys_Name: disconnected, received deauth: class 3 frame received (7) 
jun/15 18:56:47 wireless,debug wlan1_My_Companys_Name: 0C:B3:19:29:08:13 attempts to associate 
jun/15 18:56:47 wireless,debug wlan1_My_Companys_Name: 0C:B3:19:29:08:13 in local ACL, accept 
jun/15 18:56:47 wireless,info 0C:B3:19:29:08:13@wlan1_My_Companys_Name: connected 
jun/15 18:57:02 wireless,debug wlan1_My_Companys_Name: 0C:B3:19:29:08:13 attempts to associate 
jun/15 18:57:02 wireless,info 0C:B3:19:29:08:13@wlan1_My_Companys_Name: reassociating 
jun/15 18:57:02 wireless,info 0C:B3:19:29:08:13@wlan1_My_Companys_Name: disconnected, ok 
jun/15 18:57:02 wireless,debug wlan1_My_Companys_Name: 0C:B3:19:29:08:13 in local ACL, accept 
jun/15 18:57:02 wireless,info 0C:B3:19:29:08:13@wlan1_My_Companys_Name: connected 
jun/15 18:57:03 wireless,debug wlan1_My_Companys_Name: 0C:B3:19:29:08:13 attempts to associate 
jun/15 18:57:03 wireless,info 0C:B3:19:29:08:13@wlan1_My_Companys_Name: reassociating 
jun/15 18:57:03 wireless,info 0C:B3:19:29:08:13@wlan1_My_Companys_Name: disconnected, ok 
jun/15 18:57:03 wireless,debug wlan1_My_Companys_Name: 0C:B3:19:29:08:13 in local ACL, accept 
jun/15 18:57:03 wireless,info 0C:B3:19:29:08:13@wlan1_My_Companys_Name: connected 
jun/15 18:57:03 wireless,info 0C:B3:19:29:08:13@wlan1_My_Companys_Name: disconnected, received deauth: class 3 frame received (7) 
Also, we did a spectrum analisis showing the usage of each channel (but only with networks with visible SSIDs) and the result was the following:
As the image shows we are currently working on channel 4 but previously when the problem started happening we were on channel 1.
We have no idea where to keep looking and why is this problem happening. Any help or guidance will be deeply appreciated.
Upate
I would like to add more detailed information about the usage of the wifi spectrum. At the beggining we were on channel one. Then we changed to channel 3 (see image on first post) but after using InSSIDer we ended changing to channel 10.
Next is the image of the test results when channel 3 was being used:
Next is the image of the test results after selecting the best wifi channel (which was channel 10):
I would love to hear more comments and thought. Everything that makes us do something we have not yet done.
EVEN MORE UPDATE
The next image presents the DHCP packets sent between the tablets and the Mikrotik router installed on one room:
The DHCP Discover and DHCP Request packets are sent from the tablets to the DHCP server (supposedly the Microtik's), which are not beeing answered with a DHCP Offer or a DHCP ACK. One could get confuzed and think that the packets #997, #995, #999 y #1004 respond to the requests that appear before but their transaction ID shows that they are from another set of requests. Same happens with the DHCP NAK. On the DHCP protocol the different messages from one communication flow will keep the transaction ID as shown in the next image:
Looking deeper on the DHCP Request messages sent by the tablets, we found one (packet #785) asking for the IP 172.22.198.134 to the DHCP server with the IP address 152.141.217.7, which is very weird because the only Wi-Fi network registrated on the test devices was the one with SSID CVNET62710 owned by the Mikrotik.
Additionally, the ARP packets were analized and we found that the device with a MAC address 00:00:0c:07:ac:64, that appears as All-HSRP-routers_64, reports that he has the IP address 172.22.198.1 and the IP address 192.168.89.1, being this last one the IP address of the Mikrotik's Wi-Fi interface that holds the DHCP server. Next you can see the evidence:
I'm thinking this duplicated IP is the reason the problem of the disconnection and reconnection on of the tablet of this room. Sadly, I haven't been able to take captures on the other rooms so I can't assure or discard that this could be happening on every one of them.
Also, today in the morning when we went to configure the IP address of the tablet (of the current room) as a static one, when we selected manual IP address we found that it was set to 192.168.1.128, which is not from the ranged delivered by the Mikrotik (192.168.89.X). It seems that it took the IP address of another network even though the only registrated Wi-Fi network is CVNET62710.
I would like to hear your thoughts on the matter and see if you think this could be causing the recurrent disconnection and reconnection of the devices.

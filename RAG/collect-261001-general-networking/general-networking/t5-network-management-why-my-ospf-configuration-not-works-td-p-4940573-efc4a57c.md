---
id: collect-261001-general-networking/general-networking/t5-network-management-why-my-ospf-configuration-not-works-td-p-4940573-efc4a57c
title: "t5-network-management-why-my-ospf-configuration-not-works-td-p-4940573-efc4a57c"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-network-management-why-my-ospf-configuration-not-works-td-p-4940573-efc4a57c.md
source_anchor: ""
source_lines: [1, 160]
sha256: 556f19fdaced2583de1e2b62baba911654ff8b23d1d983759c69ccdf34cf02ed
---

# t5-network-management-why-my-ospf-configuration-not-works-td-p-4940573-efc4a57c

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-15-2023 07:39 AM
So here is my topology. But I will not write each router's interface configuration, because I already put them in my .pkt file attached below this question. But I will show you my OSPF configuration for each router as following:
1. router0
router ospf 1
router-id 1.1.1.1
network 10.0.0.0 0.255.255.255 area 0
network 15.0.0.0 0.255.255.255 area 0
network 13.0.0.0 0.255.255.255 area 0
network 192.168.1.0 0.0.0.255 area 0
2. router1
router ospf 1
router-id 2.2.2.2
network 10.0.0.0 0.255.255.255 area 0
network 11.0.0.0 0.255.255.255 area 0
network 14.0.0.0 0.255.255.255 area 0
network 192.168.2.0 0.0.0.255 area 0
3. router2
router ospf 1
router-id 3.3.3.3
network 11.0.0.0 0.255.255.255 area 0
network 12.0.0.0 0.255.255.255 area 0
network 15.0.0.0 0.255.255.255 area 0
network 192.168.3.0 0.0.0.255 area 0
4. router3
router ospf 1
router-id 4.4.4.4
network 12.0.0.0 0.255.255.255 area 0
network 13.0.0.0 0.255.255.255 area 0
network 14.0.0.0 0.255.255.255 area 0
network 192.168.4.0 0.0.0.255 area 0
5. router10
router ospf 1
router-id 5.5.5.5
network 192.168.1.0 0.0.0.255 area 0
6. router11
router ospf 1
router-id 6.6.6.6
network 192.168.2.0 0.0.0.255 area 0
7. router12
router ospf 1
router-id 7.7.7.7
network 192.168.3.0 0.0.0.255 area 0
8. router13
router ospf 1
router-id 8.8.8.8
network 192.168.4.0 0.0.0.255 area 0
But whenever I tried to send simple PDU, it shows that it fails, this happens when I sent it from router 10 to router 11 and sometimes the router 0 to router 2:
In the picture above the route from router0 to router2 might successful, but sometimes they can be fail.
Also if I type the 'show ip ospf neighbor' command, it does not show anything. Below is the example from router 0.
But still.. I'm curious about what's wrong with my configuration? Is there any mismatch with the IP address?
Thankyou.
Solved! Go to Solution.
- Labels:
- 
						
							
		
			Network Management
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-15-2023 08:49 AM - edited 10-15-2023 08:50 AM
Hello @Arren
Thanks for your sharing file.
On Router1 2 and 3 network command is not configured under ospf procces.
Example on Router_3:
Configure network command on these Routers like you have done for Router_0.
Router_1:
Router_2:
Router_3:
After that OSFP on your topology should be good.
Example on Router_0:
.ı|ı.ı|ı. If This Helps, Please Rate .ı|ı.ı|ı.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-15-2023 07:52 AM
You config is correct' you need only to take double check the interface IP in each router' it can be missconfig IP in two router connect to each other
One use different subnet than other.
Check correct and try again
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-15-2023 08:19 AM
The Picture is not clear.
Follow below troubleshooting :
1. show ip interface brief (make sure all the connected interface up and running)
2. ping neighbour IP see device each other can ping easily
3. show ip ospf interface (show that OSPF enabled on the interface) - if not some IOS have default passive interfaces all (if i remember correctly) - so enable ospf no passive interface x/x
4. if the interface are p2p make sure you use ospf point to point
=====️ Preenayamo Vasudevam ️=====
***** Rate All Helpful Responses *****
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-15-2023 08:49 AM - edited 10-15-2023 08:50 AM
Hello @Arren
Thanks for your sharing file.
On Router1 2 and 3 network command is not configured under ospf procces.
Example on Router_3:
Configure network command on these Routers like you have done for Router_0.
Router_1:
Router_2:
Router_3:
After that OSFP on your topology should be good.
Example on Router_0:
.ı|ı.ı|ı. If This Helps, Please Rate .ı|ı.ı|ı.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-15-2023 10:32 PM
It works, thankyou
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-15-2023 10:38 PM
You're welcome @Arren
.ı|ı.ı|ı. If This Helps, Please Rate .ı|ı.ı|ı.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-15-2023 10:45 PM
Sorry but the config you attach is different than zip file ?
I couldnot open zip and check your config it correct!!!

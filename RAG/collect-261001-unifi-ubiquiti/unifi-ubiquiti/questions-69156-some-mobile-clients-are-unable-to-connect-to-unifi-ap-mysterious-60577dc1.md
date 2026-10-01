---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-69156-some-mobile-clients-are-unable-to-connect-to-unifi-ap-mysterious-60577dc1
title: "questions-69156-some-mobile-clients-are-unable-to-connect-to-unifi-ap-mysterious-60577dc1"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-69156-some-mobile-clients-are-unable-to-connect-to-unifi-ap-mysterious-60577dc1.md
source_anchor: ""
source_lines: [1, 37]
sha256: 83dfb4352929a56b720e1b7fad5024c00fd26b415230253b968983b143dbed46
---

# questions-69156-some-mobile-clients-are-unable-to-connect-to-unifi-ap-mysterious-60577dc1

We have two UniFi Controllers, one contains the main internet its directly connected to the fibre and then we give internet to the second controller using LAN connection, internet works fine in main controller but on the second controller which has 4 wireless networks, it sometimes gets connected and sometimes it does not even let the device connected in wireless it says Access Denied or sometimes tries to connect and then comes back.
I'm really new to networking and all that stuff i have attached. Logs in PNG. Here are a few technical stuff i can explain:
Main Unifi Controller ==>
WAN:
Router: 192.168.1.1
Subnet : /24
DHCP Mode Enabled:
DHCP gateway: 192.168.1.1
DNS: 1.1.1.1
DNS2: 1.0.0.1
domain: our.server
Wireless:
Radius EAP Auth enabled which authenticates clients to our NPAS in Windows server.
LAN:
Static IP: 192.168.1.250 (Assigned by our ISP the guy who setup the first router)
So internet works a but fine on laptop from all the wireless AP in the main Controller.
Now in the second Controller
WAN:
Router: 192.168.2.1 (This second controller)
Subnet : /24
DHCP Mode Enabled:
DHCP gateway: 192.168.2.1
DNS: 1.1.1.1
DNS2: 1.0.0.1
domain: our.server
Wireless:
WPA with key
LAN:
Static IP: 192.168.1.250 (Since Ip is being assigned by DHCP in Controller 1 via the LAN) .
The issue is mainly in the second controller's Wireless:
I have assigned static channels based on a forum here at UI.
Access Denied Cannot Connect Disabled
These are the messages we get while connecting mobiles to our wireless AP from the second controller.
I'm really new so maybe someone can help.
Thanks!
Regards
UPDATE: I have turned of CCMP encryption and now on TKIP mode clients can connect.

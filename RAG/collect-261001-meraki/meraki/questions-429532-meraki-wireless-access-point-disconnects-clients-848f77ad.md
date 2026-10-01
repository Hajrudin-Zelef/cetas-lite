---
id: collect-261001-meraki/meraki/questions-429532-meraki-wireless-access-point-disconnects-clients-848f77ad
title: "questions-429532-meraki-wireless-access-point-disconnects-clients-848f77ad"
domain: meraki
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["intel", "research"]
source: docs/RAG/collect-261001-meraki/questions-429532-meraki-wireless-access-point-disconnects-clients-848f77ad.md
source_anchor: ""
source_lines: [1, 27]
sha256: 116d102824c5d624afe1f1c7da362eaf4243e9ca90a1dc6f0de1d0bff0e7973d
---

# questions-429532-meraki-wireless-access-point-disconnects-clients-848f77ad

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
5
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
We have a Meraki MR16 Cloud Managed AP and it disconnects certain clients. The clients with Intel wireless cards work without any disconnects. The Meraki reports the follow in its event log:
Sep 4 09:55:47 WPA authentication
Sep 4 09:55:47 802.11 association channel: 11, rssi: 64
Sep 4 09:55:38 802.11 disassociation client has left AP
Sep 4 09:55:38 WPA deauthentication vap: 0, radio: 0, aid: 1633956416
An example wireless network card which the Meraki disconnects is Realtek RTL8191SE 802.11b/g/n WiFi Adapter. The realtek laptop is sat 2 meters away from the AP and has a lot of signal and the Meraki reports minimal interference.
Any ideas why it disconnects non-intel wireless network cards?
I have run into similar problems with our MR12 access point we got as a demo unit. Specifically it seems to be newer Intel WiFi Cards that have this problem. If these Intel WiFi laptops attempt to connect to an SSID using encryption like WEP or WPA2, they will appear to connect and then immediately disconnect and repeat that pattern forever. Only when using an open SSID (eliminating security all together) could those laptops connect.
Obviously, this presented a real problem for us. :) As we refresh computers we are getting more laptops with the affected Intel WiFi cards in them. I opened a ticket with Meraki and they blamed it on Intel drivers. Like in your situation they did a firmware update that gave us an extra option to specify 'WPA encryption mode' but that did nothing to help.
Here are the cards I found in our environment that were affected:
Card 1: Intel Centrino Advanced N 6235
Card 2: Intel Centrino Advanced N 6205
Card 3: Intel Centrino Ulitmate N 6300 AGN
This put a serious reservation in my mind about deploying Meraki WAPs in our environment as this is a pretty big problem. So far they have shown little motivation to actually fix the problem. We have $50 WAPs that seem to handle these WiFi cards just fine.
So anyway, they know about the problem but are doing nothing to fix it. :(
I am having exactly the same issue. Cisco Meraki is simply wasted my 3 months in blaming my network end that enforced me to change the switches, cabling, nodes positions etc etc etc. After three months I requested them to downgrade the firmware and all my clients started to authenticate without any issue. Today one of the nodes showed the same behavior of dis-association. I am again at the mercy of Meraki. :(
I was told that Meraki Cisco engineering team is aware of this issue with Intel Cards. When it will be fixed? who knows.
Update: Its been more than one year. MR16, MR24 keeps on disconnecting our clients with Intel(R) Centrino 6200 series cards. You can imagine how much loads of tech support would have been performed to find the core issue. No one at Meraki able to understand this. I am seriously looking into other options. :(
Save yourself and buy a cheaper node, I am sure it will outperform meraki.
A little update on this: I think the MR12 AP we have been using over the past year is the culprit. In the last 24 hours I managed to get my hands on an MR18 and found that laptops with the affected WiFi cards are able to access the network with encryption.
Anyway, there may be other APs in Meraki's lineup that have problems but at least we know one to avoid.

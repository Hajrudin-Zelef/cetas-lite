---
id: collect-261001-meraki/meraki/r-meraki-comments-tc1kfj-firewall-stopped-talking-to-wan-port-for-no-f27995a9
title: "r-meraki-comments-tc1kfj-firewall-stopped-talking-to-wan-port-for-no-f27995a9"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-meraki/r-meraki-comments-tc1kfj-firewall-stopped-talking-to-wan-port-for-no-f27995a9.md
source_anchor: ""
source_lines: [1, 35]
sha256: 609f3fd71cef41c262aed57293fa25524ea3586c31c762cc80e67b3d7356dde3
---

# r-meraki-comments-tc1kfj-firewall-stopped-talking-to-wan-port-for-no-f27995a9

Firewall stopped talking to WAN port for no reason yesterday? 
        
        
        
    
    
    So we had a real bizarre situation yesterday. Our internet suddenly stopped working yesterday around 9:30-10am CST 3/10/22.
Our ISP confirmed zero traffic out of the WAN port - couldn’t even see a MAC address. Some of our switches rebooted on their own as well.
Called meraki support and they seemed to understand the issue without me saying much - like they knew something was going on but wouldn’t provide any details. They vaguely said something to the effect that our firewall downloaded a corrupt config and stopped traffic from our WAN port.
To get it working again, I had to factory reset our firewall and go through the setup process/configure WANs etc. After a few minutes it sprang back to life and all was well.
We made no changes to our config - this happened out of no where. I asked what happened and they said since I factory reset the firewall there are no logs to look at… but aren’t they in the cloud?? I asked what we can do to prevent this from happening again and support said “oh don’t worry, you’re protected now, it won’t happen again”. I asked protected from what and he got flustered and said they “tagged our network” and it can’t do this anymore.
Did they get hacked or something? Is there something seriously going wonky here? I see all these posts about stuff being offline and something up in Europe as well as a Dallas data center being offline causing issues…. Just a real bizarre situation without any good explanation.
Anyone else have something like this happen?
Section des commentaires
Lots of people had issues yesterday
All week it seems…
Mx84 by chance? We had to do similar and roll back firmware a few months ago.
Mx450
I haven’t dealt with one of those. Biggest thing we use are 100’s.
We had this happen today, suddenly. MX100 randomly dropped both WAN connections completely. Shows offline in the Meraki Cloud dashboard.
WAN 1 Internet is not registering any connection or traffic at all, no lights. Upstream device registers link connection.
WAN 2 Internet registers link connection, solid green only but no activity. Upstream device registers link connection. Strangely enough, when I remove the cable, the solid green link light stays on.
Connectivity provider confirms both upstream devices are functioning properly and that neither upstream device was receiving a MAC. I bypassed the MX100 and verified connectivity myself.
The MX was functioning properly this morning and bam, lights out, literally.
On troubleshooting, I tried accessing the MX locally and the default local password wouldn't take at all. Resorted to factory reset and I was able to access the local config (verifies that LAN side is working). Tried to reconfigure the WAN ports again, but the MX was still reporting "The security appliance is trying to join a network or find a working ethernet connection"
FW: MX 18.107.2
Opened up a case, but it seems like an RMA is in order.
Anyone seen this before?
Have you made any VLAN changes lately? That's what they (support) said was set incorrectly but couldn't find the problem. We moved all our L3 termination points to a aggregation switch, then used a VLAN to connect the two and all the problems went away with our firewall. Perhaps their specs on capacity are inflated on the MX models when taking in all the new features of the newer firmware release features? I dunno.
None. Just crapped out suddenly. Support was taking a while to respond, so I called, reviewed with them and they RMA'd with not much additional questions.
I think it may have been due to this issue and you were just particularly unlucky and your config somehow was corrupted.
That thread seems to talk about vpn authentication - barely mentions what the issue was with fetching config but not sure why you were downvoted. Wish I understood more clearly what the issue was. It really porked us good when it happened.
The second post in the thread I shared is from a Meraki employee and acknowledges the problem as wider than just Meraki authentication and affecting config files is why It seems relevant. And the timeframe. Front line support probably had no clue what the real problem was or how wide it was, but it does seem significant enough that maybe something official comes from Cisco soon. I would assume employees are trained to not make authoritative statements about the nature of outages and for good reasons assuming they even know (not likely in the middle of an outage that they know exactly all the details about it).
Someone also reported issues for sfps with the latest firmware. Don't know if you have a similar issue.
https://www.reddit.com/r/meraki/comments/tc0380/meraki_mx_update_1616_broken_sfp/?utm_medium=android_app&utm_source=share

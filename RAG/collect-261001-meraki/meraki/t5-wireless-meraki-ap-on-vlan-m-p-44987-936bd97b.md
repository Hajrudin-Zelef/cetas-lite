---
id: collect-261001-meraki/meraki/t5-wireless-meraki-ap-on-vlan-m-p-44987-936bd97b
title: "t5-wireless-meraki-ap-on-vlan-m-p-44987-936bd97b"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/t5-wireless-meraki-ap-on-vlan-m-p-44987-936bd97b.md
source_anchor: ""
source_lines: [1, 200]
sha256: e9b4152034026a87a2b7ccb6b328dcfa0ec9f0bc2f7bc3d548512d795eed0db1
---

# t5-wireless-meraki-ap-on-vlan-m-p-44987-936bd97b

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-23-2019 09:19 PM
Last time my network only 1 segment with IP 192.168.0.0/24. and my Meraki AP works fine.
I have 2 SSID : bridge mode & NAT (DHCP Meraki)
Then we implement VLAN in our network, now Meraki AP port in switch already change to VLAN Client segment with different IP (10.10....)
Now the problem is my meraki AP status is offline, i change the Meraki AP static IP address to VLAN CLient segment, but still can not ping (offline)
Somebody help me ?
Solved! Go to Solution.
- Labels:
- 
						
							
		
			Meraki
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-24-2019 04:29 AM
@barth1 wrote:
With this config, The SSID internal will share IP VLAN Client (11) 192.168.15.0 /24 to user ?
Yes, with that config the packets coming in from the wireless clients will be tagged with VLAN 11 on the trunk, and inversely, the packets coming into the AP via the trunk on VLAN 11 will be sent out to the wireless clients untagged.
I think the problem is that you're actually using the native/untagged VLAN. In my setup you'll see that I don't actually use the native VLAN for anything. Try setting VLAN 11 to tagged as well, and use another random untagged VLAN that you're not actually using.
Also it might also be interesting to know that in NAT-mode, the packets will be tagged with the AP's native VLAN.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-23-2019 10:51 PM
What VLAN are you running in the switch as native? Does that match on the accessport?
@barth1 wrote:
Last time my network only 1 segment with IP 192.168.0.0/24. and my Meraki AP works fine.
I have 2 SSID : bridge mode & NAT (DHCP Meraki)
Then we implement VLAN in our network, now Meraki AP port in switch already change to VLAN Client segment with different IP (10.10....)
Now the problem is my meraki AP status is offline, i change the Meraki AP static IP address to VLAN CLient segment, but still can not ping (offline)
Somebody help me ?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-23-2019 11:04 PM
Yes, Native VLAN Client
"match on the accessport" what thats means ?
I can not reboot, "unreacable device "
I can not ping
any something configuration ?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-23-2019 11:57 PM
If the switch is PoE version you can turn off the PoE from port and then turn it on again to reboot the AP
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-24-2019 12:03 AM
Dear @mikko.tainio@blc.fi 
i know and i already try, but still can not ping the static IP, (offline)
What should i do? reset then setup again ?
Another question :
Is the all configuration (ssid, policy, etc) is save in cloud ? so when i reset the AP, the configuration will apply automatically again ?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-24-2019 12:13 AM
Hi @barth1
All the configuration is saved in Meraki cloud so reseting single AP doesn't affect the SSID's etc.
Is your AP configured to use IP from same VLAN that is configured as native VLAN on the switch port the AP is connected? And are all the VLANs in use allowed in the switch port configuration?
Another way is not to use native VLAN and configure the VLAN in APs IP configuration.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-24-2019 12:47 AM
My pc can ping to another pc but the AP still can not ping (even when i change the IP to static follow vlan client segment 10.10.10.0/24) but still can not...
But..
Why the ssid internal (bridge mode) is working. So the AP can share the vlan client segment to user.
SSid with NAT Mode (dhcp meraki) is not show.
I can not unplug or chage to default vlan 192.168.0.0/24.
It will share the old ip segmet to user...
I dont understand ur suggention, can yoy more detail ?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-24-2019 01:57 AM
be default, the maraki AP (MRxx) transfers its management traffic over native VLAN (or access port).
if you want your SSID(s) in a different network segment:
1) use Meraki NAT: nothing else is required on the switch (access mode/port is enough)
2) use bridge mode:
2a) configure additional VLAN(s) on your meraki AP switch port (trunk mode).
2b) make sure you have a DHCP server within the additional VLAN(s)
2c) setup the additional VLAN(s) in the 'access control' setting, section VLAN tagging / VLAN id
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-24-2019 03:00 AM
How is everything setup? Do you have a network drawing? This is how I've set it up for example:
Note that for the switch I set the VLAN both in Switch Settings > Management VLAN and in the IP settings of the individual switch.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-24-2019 03:38 AM
Hi, here the simple topology of my new network
Any miss config ?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-24-2019 04:17 AM
Hi @barth1
Are you using HP Procurve for switches as shown in the drawing?
Have you introduced the new VLAN to the switches and added it to all the necessary ports as untagged?
-Mikko
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-24-2019 04:24 AM
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-24-2019 04:14 AM
OTHER SOLUTION
With this config, The SSID internal will share IP VLAN Client (11) 192.168.15.0 /24 to user ?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-24-2019 04:29 AM
@barth1 wrote:
With this config, The SSID internal will share IP VLAN Client (11) 192.168.15.0 /24 to user ?
Yes, with that config the packets coming in from the wireless clients will be tagged with VLAN 11 on the trunk, and inversely, the packets coming into the AP via the trunk on VLAN 11 will be sent out to the wireless clients untagged.
I think the problem is that you're actually using the native/untagged VLAN. In my setup you'll see that I don't actually use the native VLAN for anything. Try setting VLAN 11 to tagged as well, and use another random untagged VLAN that you're not actually using.
Also it might also be interesting to know that in NAT-mode, the packets will be tagged with the AP's native VLAN.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-25-2019 06:20 AM
Its works now, i do line my other solution
Thank you for response my thread.
Bye

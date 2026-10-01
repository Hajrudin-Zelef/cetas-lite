---
id: collect-261001-general-networking/general-networking/t5-wireless-mandatory-dhcp-m-p-102573-57866f4f-2
title: "t5-wireless-mandatory-dhcp-m-p-102573-57866f4f"
domain: general-networking
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["latency", "memory"]
source: docs/RAG/collect-261001-general-networking/t5-wireless-mandatory-dhcp-m-p-102573-57866f4f.md
source_anchor: ""
source_lines: [20, 162]
sha256: 67533d1258917f44f6ecb9e6db872054c8a29d90395ffeb249339716e609ba82
---

# t5-wireless-mandatory-dhcp-m-p-102573-57866f4f

			Meraki
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-01-2020 12:18 AM
The Meraki documentation has now been updated on this feature to reflect the fact that it only works when roaming if the client refreshes its DHCP when the roam is complete.
https://documentation.meraki.com/MR/Access_Control#Mandatory_DHCP
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-11-2020 12:03 PM
DHCP shouldn’t be refreshing on every roam, so long as the AP has the same subnets/VLAN available it should do a Layer 2 roam with no DHCP refresh. If DHCP refreshes on every roam then it would impact latency sensitive protocols such as VoIP. I believe the Mandatory DHCP just forces the client to use a DHCP address (rather than the client setting one manually), it doesn’t force it to renew it on every roam.
How is your SSID set up, is it in Bridge mode, or Layer 3 Roaming? If both the APs have the same VLAN/subnet available then you should be able to use Bridge mode with Layer 2 roaming. If the subnets/VLANs are different you’ll need to use Layer 3 Roaming, and then traffic will be tunnelled back to the originating AP if a Layer 2 roam can’t be made.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-11-2020 12:30 PM
Hi, the SSID is set in bridge mode to the same VLAN on every AP. It's Meraki Support who have told me that the DHCP should refresh on every roam and like you I'm sure that's not correct.
The bit I don't understand is this: If manadatory DHCP is enabled on an SSID, how does the "roamed to" AP know that the DHCP exchange happened on the "roamed from" AP? Either the first AP has to tell all the others about a new client and the fact it has witnessed the DHCP exchange OR the "roamed to" AP has to query the dashboard or other APs about whether the mobile client is good to join from a DHCP perspective. Either way, that's the bit which seems to be failing and support can't tell me which of these it is!
If I disable mandatory DHCP then it works fine.
If I disconnect and then reconnect to the "roamed to" AP then that works as well. And from that point on I can roam back to it so it's obviously caching the DHCP status of the client.
I've got another call with support on Friday this week so they can collect yet more logs. As least I've managed to reproduce it in my lab now and not on the site which is 2.5 hours drive away!
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-11-2020 12:33 PM
I've got packet captures from the air and on the LAN. My tester is pinging the gateway. As soon as the roam occurs the ICMP packets are not bridged to the LAN by the "roamed to" AP. And there's no DHCP traffic either so Support seem to have got that one wrong.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-11-2020 12:36 PM
No ICMP until it roams back to the original AP!
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-11-2020 01:52 PM
Like you said, the only way I could see this working is if there is meant to be some sort of state exchange between the Meraki cloud and the MR to allow the 'roamed to' AP to recognise that the IP address is DHCP issued. Incidentally what model are the APs?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-11-2020 02:02 PM
The production APs are MR33 and MR74. There's no correlation on roaming between similar or different models either. It does the same in my lab which is an MR32 and an MR24. I'm on my third support engineer now and no-one seems to be able to answer in any detail how it's supposed to work.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-11-2020 02:03 PM
Surely I can't be the first person to turn this feature on?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-11-2020 02:07 PM
Maybe you are the first person to use it in anger... maybe there is a bug in it with roaming (that's not unheard of). On a side note it shouldn't actually work at all in your lab according to the doco, "All 802.11ac Wave 2 capable MR access points running MR 26.0 firmware or later support this feature", https://documentation.meraki.com/MR/Access_Control. But that said there is no reason it shouldn't be working with the MR33 and MR74, they're both Wave 2 device.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-11-2020 02:34 PM
Good point about the AP models. I had read that but forgot it soon after so thanks for the memory jog. I've got a customer who's got a box of MR42s they're not using at the moment while their building is being refurbished and they're upgrading to MR55s when it's done anyway. I'll be passing close to their site tomorrow so I think it might be a chance to pop in and borrow a couple.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-18-2020 04:46 PM
Not 100% aligned to your issue, though I also have had issues when enabling 'Mandatory DHCP'. In my case it was on an SSID which uses L3 Roaming (noting that yours is 'Bridge to LAN'). In my case all WAPs are MR52 running 26.8.1. WPA2-Enterprise SSID, ClearPass RADIUS, using AP Tags for client VLAN assignment. Disabling Mandatory DHCP fixed the issue. (Note, I'm not using VLAN Override via Group Policy in this case, though am for other use cases on the same SSID).
Client roams between WAPs with the same VLAN present (eg, WAPs on a given floor) were fine, however once the client roamed to another floor (WAP with a different VLAN), a L3 Roam occurs and the client loses all connectivity. This is because the 'roamed to' WAP didn't see the DHCP Request, therefore drops all client traffic.
I spent time troubleshooting this with support, and in the end was told this is expected behaviour. Unfortunately this is not mentioned anywhere (that I have found), so hopefully Meraki will either update the doco, or prevent the feature being enabled for an SSID with L3 Roaming enabled (which aligns with how other feature incompatibilities are handled in dash)...
IMO they should fix this as it is a good security feature, particularly liked by enterprise security folks...
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-19-2020 06:20 AM
I was hoping to enable this feature but I may wait a bit now, until summer anyway.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-26-2020 05:21 AM
Hi, ok I think I may finally have answered this one after many hours testing. I have tested four clients from scratch with a factory reset of the APs in between each test and using a fresh SSID profile on each client.
Apple iPad
- Roams between APs with a “reassociation frame”
- DHCP is updated on roaming
- LAN connection is maintained
Motorola Moto G6
- Roams between APs with a “reassociation frame”
- DHCP is updated on roaming
- LAN connection is maintained
Netally Aircheck G2
- Roams between APs with a “reassociation frame”
- DHCP does not update on roaming
- LAN connection is then blocked from wired to wireless i.e. I could see the ICMP echo from the wireless client on the wired capture and the ICMP reply from the target but this is NOT passed on to the client by the AP
Panasonic FZ-N1
- Roams between APs with a fresh “association frame” instead of a reassociation frame.
- DHCP does not update on roaming

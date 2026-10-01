---
id: collect-261001-general-networking/general-networking/t5-security-sd-wan-best-practices-for-native-vlan-configuration-m-p-93478-538e25b2-2
title: "t5-security-sd-wan-best-practices-for-native-vlan-configuration-m-p-93478-538e25b2"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-security-sd-wan-best-practices-for-native-vlan-configuration-m-p-93478-538e25b2.md
source_anchor: ""
source_lines: [21, 196]
sha256: d4bfe7eaecf776a576706499f84922cc9dbfcec2576d5f2918e77d88d23f4fce
---

# t5-security-sd-wan-best-practices-for-native-vlan-configuration-m-p-93478-538e25b2

			Meraki
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-28-2020 03:08 AM
I must have expressed myself poorly.
I always ensure that the options of ALL and DEFAULT are not used.
It doesn't matter what the default network is, avoid using it. The entire concept of a default value is insecure, inherently.
As far as the example you have given is concerned, I'd observe:
- Why call a VLAN Default (is it a mistake)?
- Why give the VLAN 192.168.3.0/24 an ID of 1?
You don't need a default VLAN. If you need to enter a VLAN value, when it is not required, use 101 (think about it).
Always reference VLANs explicitly, not generically. Refer to them by their unique VLAN ID. Individually declare the VLANs to be passed by a trunk port. Declare the specific VLAN ID on access ports.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-28-2020 03:29 AM
In case you‘re wondering why @Uberseehandel strongly advises against using Vlan1: historically, there has been something called „VLAN hopping“ that leveraged the default VLAN to „hop“ into others that should notbe accessible by the attacker.
Please find more information here: https://portunreachable.com/vlan-hopping-vulnerability-527ee506dae3
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-28-2020 05:17 AM
Thank you both. I’m kind of getting it. The thing i still don’t understand is how will my switches and APs know where to get their IP from? What determines which VLAN leases IPs to layer 2 devices?
Made many a network over the years, now de facto admin of a retreat center with some of this fine Meraki hardware.
Fortune 100 Tech veteran/refugee.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-28-2020 06:40 AM
The switches and APs are on the Management VLAN. They get their IP address from the Management VLAN DHCP server on the MX.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-28-2020 06:42 AM
Ok - thanks for going slow with me here. How does one define the "management VLAN"? Is it the VLAN labelled default? What makes it so? Right now, the IPs my switches and APs are leasing IPs from the range current ascribed to VLAN1 aka "Default". What designates a VLAN the management VLAN?
Made many a network over the years, now de facto admin of a retreat center with some of this fine Meraki hardware.
Fortune 100 Tech veteran/refugee.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-28-2020 06:58 AM
@Uberseehandel OK I think I figured it out. Tell me if I got this right. To make a management VLAN, Id create another VLAN (lets say VLAN 5), and then I would change this:
I would select VLAN 5 from the "native VLAN" dropdown. This would tell anything connected at Layer 2 to join that VLAN. Then I would change ALL VLANs under "allowed" to the ones needed over each trunk port. I think thats the piece I was missing - to change the native VLAN on the wired ports.
Made many a network over the years, now de facto admin of a retreat center with some of this fine Meraki hardware.
Fortune 100 Tech veteran/refugee.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-28-2020 07:25 AM
have you tried it ?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-28-2020 08:14 AM
No I’ll need to populate my other switches with the new vlan number so it all works. Did i get the right idea though??
Made many a network over the years, now de facto admin of a retreat center with some of this fine Meraki hardware.
Fortune 100 Tech veteran/refugee.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-28-2020 08:40 AM
What you are proposing looks very like what I have, so I suggest you proceed. Don't forget to explicitly declare the VLANs to be passed on the uplinks, and to pass the Management VLAN ID as well as the SSID VLAN IDs to the access points.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-28-2020 08:59 AM
Yes I think I have figured it out. I just had literally forgotten that the actual ports on the MX determine the native vlan and that I was to configure them accordingly. Thanks for your patience. It makes sense now. The APs and the switches will all grab their IP from whatever the native VLAN is set to on the trunk originating at the port on the MX.
Made many a network over the years, now de facto admin of a retreat center with some of this fine Meraki hardware.
Fortune 100 Tech veteran/refugee.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-24-2019 03:55 AM
@federicogalarza wrote:
I would like to know what are the best practices which you usually implement in the Meraki world.
Thanks!
Federico
The general rules I adhere to are:
- Tag all VLANs
- Set up a Management VLAN
- Set up an Isolated Guest VLAN (and SSID)
- Do not use the native LAN
- Create a faux VLAN for those cases where the configuration GUI requires a VLAN ID (make sure it goes nowhere)
- Explicitly declare the VLANs a given switch port should pass, avoid the all option on uplinks
- Make the firewall proscriptive rather than permissive.
There are some exceptions for certain network architectures, even so, they are usually very limited in scope. For example, at the moment we have a Meraki network within a third party network. The MX running the Meraki network has its WAN port on a native LAN that is connected to the LAN port of the external facing security appliance which uses PPPoE on its WAN uplink. This enables the dynamic external IP address supplied by the ISP to be passed to the MX and even to the Z3C connected to the MX.
It is simpler to set up than it is to write about it.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-05-2020 07:36 AM
I appreciate reading through your responses! Sorry to necro this thread but I figured it wasn't too old?
I was wondering if you would use your Management VLAN (changed from VLAN 1) as the native VLAN on Trunks for between switches?
Also, would you use the trunks native vlan to be the management vlan when connecting to a server with a virtual switch in it?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-05-2020 02:15 PM
I think my question would be how would you not use the native VLAN as the management VLAN? The native VLAN is what assigns IPs to switches and APs, no? So wouldn't you want to be on that VLAN to interface with all those pieces of hardware?
Made many a network over the years, now de facto admin of a retreat center with some of this fine Meraki hardware.
Fortune 100 Tech veteran/refugee.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-29-2022 03:52 PM
Hello RumorConsumer,
I think what you were missing was the Management VLAN setting.... Navigate to Switch -> Switch Settings, at the top you define the Management VLAN. All your APs and Switches will try to grab an address from this VLAN.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-06-2020 01:50 AM
@CharlesIsWorkin wrote:
. . .
I was wondering if you would use your Management VLAN (changed from VLAN 1) as the native VLAN on Trunks for between switches?
Also, would you use the trunks native vlan to be the management vlan when connecting to a server with a virtual switch in it?

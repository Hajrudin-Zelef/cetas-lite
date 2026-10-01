---
id: collect-261001-general-networking/general-networking/t5-security-sd-wan-mx84-swap-m-p-11483-0e69b3bb
title: "t5-security-sd-wan-mx84-swap-m-p-11483-0e69b3bb"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-security-sd-wan-mx84-swap-m-p-11483-0e69b3bb.md
source_anchor: ""
source_lines: [1, 104]
sha256: c5e2f0ab0b7ba0cdfcbe2913bed4302436e2a1e3a693348c7e7a47c4fd34eb2a
---

# t5-security-sd-wan-mx84-swap-m-p-11483-0e69b3bb

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-06-2018 12:35 PM
We received a replacement MX84 for the clock signal issue, and I am sure this is an easy process, but what do we do to swap in the new MX84 and remove the old one while retaining all the settings?
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
02-06-2018 03:00 PM
That article isn't bad. Basically I did the following when I went through this.
1. Plugin the MX to an internet connection so it can check-in and be managed
2. Add it to the same network as the existing MX84 that you are going to replace. It'll add as a spare and get the same config as the current MX84
3. Configure the new MX84 with any WAN interface static IPs. Those won't migrate over automatically. Anything you configure on the MX local config page won't migrate over automatically.
4. Rack the MX84 directly above below your current MX84 if possible
5. Once ready move cables port by port from the old MX84 to the new one (this will result in interruption but it should be minimal < 60 seconds).
If this was helpful click the Kudo button below
If my reply solved your issue, please mark it as a solution.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-06-2018 01:06 PM
See this link. I think this describes what you need.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-06-2018 03:00 PM
That article isn't bad. Basically I did the following when I went through this.
1. Plugin the MX to an internet connection so it can check-in and be managed
2. Add it to the same network as the existing MX84 that you are going to replace. It'll add as a spare and get the same config as the current MX84
3. Configure the new MX84 with any WAN interface static IPs. Those won't migrate over automatically. Anything you configure on the MX local config page won't migrate over automatically.
4. Rack the MX84 directly above below your current MX84 if possible
5. Once ready move cables port by port from the old MX84 to the new one (this will result in interruption but it should be minimal < 60 seconds).
If this was helpful click the Kudo button below
If my reply solved your issue, please mark it as a solution.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-12-2018 06:31 AM
BUT
When I went to gracefully do HA to swap the primary in the Dashboard the spare said it was up and fine but the who LAN lost internet. (I have noticed this before and mentioned it to Meraki. If i were to have just pulled the plug on the primary and let it do a fail over it would have come up fine. This issue Only happens when using the swap button on the dashboard. Meraki support told me this is something to do with the Virtual IPs but he wasn't sure why.)
Needless to say this caused Meraki support to tell me to do a cold swap and that is when all h*%l broke loose. Downtime of 1 hour and issues caused by the dashboard seeing them as dual DHCP servers.
All I can say is stick to the Hotswap Mentioned above. just don't use the swap button on the dashboard.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-02-2020 07:56 AM
I did the HA method to replace a MX400 from an RMA. However I had to swap back and forth a couple of times as the RMA did not send me the SFP+ module and the modules are NOT hot-swap. But it worked. Each time the network stopped for only about 10 seconds. The one issue we had that I had not seen mentioned anywhere is on the MS switch settings, I had the old MX mac address as the only one allowed for DHCP. You need to add both mac addresses (old and new MX) to this setting before starting the swap in order to not cause DHCP blockage.
I did not, and to make things worse, I initiated a MS firmware upgrade which meant that when the MS did a restart it did not get its address from the new MX and could therefore not reach the cloud to update even the DHCP mac allowed setting. I had to physically visit every switch and give it a static - once connected to the cloud, the new allowed DHCP mac address synced and then everyone could get back into the network. Simple mistake, bad consequences. Just add to your list of things to check.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-12-2018 11:20 AM
Besides static IP also note that a few things on the appliance status page get lost too:
- Device name
- address
- notes
I also had a strange issue yesterday replacing an MX84 where the internet 1 link light was dark for about 25 minutes, but apparently actually working as the MX came up after that time. I was just about to factory reset it though.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-31-2019 06:52 AM
Hi dudes,
What if I want to replace an MX64 with an MX84? Would I follow the same procedure?
Thanks!

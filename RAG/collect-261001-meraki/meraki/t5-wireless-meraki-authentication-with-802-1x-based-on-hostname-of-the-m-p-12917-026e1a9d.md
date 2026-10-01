---
id: collect-261001-meraki/meraki/t5-wireless-meraki-authentication-with-802-1x-based-on-hostname-of-the-m-p-12917-026e1a9d
title: "t5-wireless-meraki-authentication-with-802-1x-based-on-hostname-of-the-m-p-12917-026e1a9d"
domain: meraki
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/t5-wireless-meraki-authentication-with-802-1x-based-on-hostname-of-the-m-p-12917-026e1a9d.md
source_anchor: ""
source_lines: [1, 120]
sha256: a94e51beac7cebfc843f84ed1b47c83bfca9ad205ad230f87f909a6831d3a167
---

# t5-wireless-meraki-authentication-with-802-1x-based-on-hostname-of-the-m-p-12917-026e1a9d

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-01-2021 01:28 PM
I'm in the middle of an implementation that one of the needs is for an SSID with 802.1x authentication to be based in addition to the user, if the computer's hostname is within an AD group.
With wireshark, I can identify that the Meraki AP does not send this information. I was able to activate to send a NAS-ID, in this case the SSID name to be able to differentiate the polices but I do not identify the sending of the computer's hostname.
Can anyone help how to enable or how to configure 802.1x authentication based on the hostname of the computer that should be in an AD group?
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
09-15-2021 09:08 AM
Hi Philip,
Just to get back to you, I managed to solve it in an easier way. In access control, on radius advanced settings, i put on called-station-id, as custom, %n which in this way sends to the radius server as username, hostname.domain.local
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-01-2021 02:13 PM
You’re not going to be able to achieve this with Meraki Authentication. You need to use 802.1x with your own RADIUS server which is integrated with AD (examples are Microsoft NPS, Cisco ISE). Then you’ll be able to configure the client to authenticate using the host/computer rather than the user credentials.
If you want to authenticate both the user and computer you are going to need a more advanced RADIUS server than Microsoft NPS (probably something like Cisco ISE), and use something like EAP-TEAP, or use some of the profiling capabilities of ISE to determine if it’s the host you expect.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-02-2021 05:42 AM
Hi Bruce, how are you?
The customer currently has a wifi backbone with Cisco Aironet without Cisco ISE. It has an SSID with 802.1x authenticating to a Radius on Windows and the policy is based on an AD group that owns the computers.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-15-2021 09:10 AM
Hi Bruce,
Just to get back to you, I managed to solve it in an easier way. In access control, on radius advanced settings, i put on called-station-id, as custom, %n which in this way sends to the radius server as username, hostname.domain.local
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-15-2021 01:51 PM
@FelipeTeevo, thanks for the update. Glad to hear you found a way to achieve your desired outcome. That’s a relatively new feature, good to see it’s enabling people to use the Meraki kit in the way they want.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-01-2021 02:43 PM
I've confused by this one. You mention Meraki Authentication in the title but talk about the AP not sending the info, suggesting it is sending something to something else (perhaps RADIUS).
You can configure Windows 10 to do hostname authentication, username authentication, or send the hostname prior to login, and then the username after login. This requires a RADIUS server like Microsoft NPS.
You can configure this in group policy or under advanced settings for the WiFi connection. Check out step 3 in this link:
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-02-2021 05:44 AM
Hi Philip, how are you?
This configuration you sent is for 802.1x authentication on the wired network, correct? In this case I mean authentication with 802.1x, by wifi so that it is based on the hostname, authenticating in an AD group on a radius server running on a windows server. Currently the client has a simple structure with cisco aironet without cisco ISE and it works perfectly.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-02-2021 05:57 AM
Same deal - wired or WiFi. Just configure Windows 10 to authenticate with whatever you want.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-02-2021 06:53 AM
Phillip,
Tell me if you agree with me:
- in the current scenario, running a wireshark, I can capture from the same workstation I use for testing, in this case authenticating on the cisco aironet network, it receives the hostname as user as a parameter. So I think that this config you informed is already applied, right? Image below:
At Meraki, on the same workstation, authentication only arrives with the username:
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-15-2021 09:08 AM
Hi Philip,
Just to get back to you, I managed to solve it in an easier way. In access control, on radius advanced settings, i put on called-station-id, as custom, %n which in this way sends to the radius server as username, hostname.domain.local

---
id: collect-261001-general-networking/general-networking/t5-wireless-dhcp-snooping-amp-wlc-9800-td-p-5216456-da360f60
title: "t5-wireless-dhcp-snooping-amp-wlc-9800-td-p-5216456-da360f60"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-wireless-dhcp-snooping-amp-wlc-9800-td-p-5216456-da360f60.md
source_anchor: ""
source_lines: [1, 150]
sha256: ff3ead569382949a14a14017e03358187892e83674b7ae61e961afc7b2d43d58
---

# t5-wireless-dhcp-snooping-amp-wlc-9800-td-p-5216456-da360f60

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-28-2024 01:58 PM
Hi,
I'm using WLC 9800 local switching mode, I can see all the wifi client devices MAC coming from WLC ports (since the traffic is tunneled to the WLC), at the same time I'm configuring DHCP snooping in the Core switch to avoid any DHCP rogue servers from WIFI clients.
should I trust the WLC ports ? otherwise, my client can't get any IPs. Core switch logs show blocked DHCP packets coming from the WLC.
I'm using different SVIs in the WLC acting as dhcp relay, i'm just wondering if this is a correct implementation, or there is a way to untrust WLC ports (since all the clients MACs are seen from WLC ports).
Solved! Go to Solution.
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-28-2024 04:00 PM
Cisco does not recommend having SVI on the WLC side. Ideally, you should have the SVI on the core and use ip help-address on the SVI.
Cisco Catalyst 9800 Series Configuration Best Practices - Cisco
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-28-2024 04:33 PM
If Cisco doesn’t recommend SVIs on the WLC, then setting the SVI on the core switch and using the ip helper-address to relay DHCP requests is indeed the preferred approach. This setup reduces complexity and aligns with best practices, especially for DHCP snooping configurations. By relocating the SVI to the core, you can untrust the WLC port while maintaining DHCP functionality and security.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-28-2024 02:04 PM
Issue I think in op82 wlc add to dhcp packet' try use "" allow op82 in untrust port"" under port connect core to wlc9800.
MHM
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-28-2024 03:18 PM
Hi,
This command doesn't exit neither in the Core, nor in the WLC
If you mean in the core switch, I already did. Without trusting the WLC port it doesn't work
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-28-2024 10:00 PM
ip dhcp snooping information option allow-untrusted
This command I talk about, the wlc to SW port must config as untrust
MHM
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-28-2024 04:00 PM
Cisco does not recommend having SVI on the WLC side. Ideally, you should have the SVI on the core and use ip help-address on the SVI.
Cisco Catalyst 9800 Series Configuration Best Practices - Cisco
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-29-2024 09:01 AM - edited 10-29-2024 09:01 AM
As i'm reading the Best practices shared, I see that 9800 has built-in DHCP Snooping feature. I didn't find it anywhere in the settings "Cisco IOS XE has embedded security features such as Dynamic Host Configuration Protocol (DHCP) snooping"
Since all the clients MAC/IPs are seen as coming from the WLC ports, I want to untrust the WLC ports since it's acting as a big switch for all the Wireless clients to avoid any DHCP rogues among wireless devices.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-29-2024 09:44 AM
I believe you need to trust the WLC interface if you leave the SVI on the WLC side.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-29-2024 10:55 AM
Just removed them from WLC, untrusted the port.. everything is good now.. thanks to everyone
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-28-2024 04:33 PM
If Cisco doesn’t recommend SVIs on the WLC, then setting the SVI on the core switch and using the ip helper-address to relay DHCP requests is indeed the preferred approach. This setup reduces complexity and aligns with best practices, especially for DHCP snooping configurations. By relocating the SVI to the core, you can untrust the WLC port while maintaining DHCP functionality and security.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-29-2024 08:13 AM
SVI are located on a palo alto firewall. The core switch itself is acting as an aggregation layer for Access switches and the WLC.
So all I need is to disable SVIs and everything should be fine ? the only problem is the Mgmt interface, I should absolutely keep it ON in the WLC, and it will keep the ip helper role for the APs. At the same time the WLC ports will be untrusted. I'm confused a little bit about this one.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-06-2024 01:57 PM
So all I need is to disable SVIs and everything should be fine ? Yes
the only problem is the Mgmt interface, I should absolutely keep it ON in the WLC - Yes
and it will keep the ip helper role for the APs - not clear why you would have ip helper on the WLC for that?
If in doubt check your WLC config using the Config Analyzer using out of "show tech wireless" (not show tech) - link below.
Please click Helpful if this post helped you and Accept as Solution if this answered your query.
------------------------------
TAC recommended codes for AireOS WLC's and TAC recommended codes for 9800 WLC's
Best Practices for AireOS WLC's, Best Practices for 9800 WLC's and Cisco Wireless compatibility matrix
Check your 9800 WLC config with Wireless Config Analyzer using "show tech wireless" output or "config paging disable" then "show run-config" output on AireOS and use Wireless Debug Analyzer to analyze your WLC client debugs
Field Notice: FN63942 APs and WLCs Fail to Create CAPWAP Connections Due to Certificate Expiration
Field Notice: FN72424 Later Versions of WiFi 6 APs Fail to Join WLC - Software Upgrade Required
Field Notice: FN72524 IOS APs stuck in downloading state after 4 Dec 2022 due to Certificate Expired
- Fixed in 8.10.196.0, latest 9800 releases, 8.5.182.12 (8.5.182.13 for 3504) and 8.5.182.109 (IRCM, 8.5.182.111 for 3504)
Field Notice: FN70479 AP Fails to Join or Joins with 1 Radio due to Country Mismatch, RMA needed
Field Notice: FN74383 APs Running 17.12.4/5/6/6a May Run Out of Flash Space Preventing Upgrades
How to avoid boot loop due to corrupted image on Wave 2 and Catalyst 11ax Access Points (CSCvx32806)
Field Notice: FN74035 - Wave2 APs DFS May Not Detect Radar After Channel Availability Check Time
Leo's list of bugs affecting 2800/3800/4800/1560 APs
Default AP console baud rate from 17.12.x is 115200 - introduced by CSCwe88390
AP supported channel lookup: https://apchannels.cisco.com/

---
id: collect-261001-general-networking/general-networking/t5-switching-9200-trunk-port-goes-down-after-changing-allowed-vlan-ids-m-p-52145-af933047
title: "t5-switching-9200-trunk-port-goes-down-after-changing-allowed-vlan-ids-m-p-52145-af933047"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-general-networking/t5-switching-9200-trunk-port-goes-down-after-changing-allowed-vlan-ids-m-p-52145-af933047.md
source_anchor: ""
source_lines: [1, 180]
sha256: 1d59554aa667ee4458ee10bb4e8e08816e5e5688e77ae387c9689716be4ff88b
---

# t5-switching-9200-trunk-port-goes-down-after-changing-allowed-vlan-ids-m-p-52145-af933047

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-24-2024 07:17 AM
There is a trunk port which works fine and allows VLANs 2,4-6.
When I change the allowed VLAN IDs field, the port goes down - operational status red.
It doesn't matter what VLAN IDs I enter. It goes down even if I choose "All VLAN IDs".
Anyone has an idea?
Solved! Go to Solution.
- Labels:
- 
						
							
		
			Catalyst 9000
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-24-2024 08:52 AM
the trunk is entering in Suspended mode and this could be because the WLC and not the switch side. If could be due native vlan mismatch. The negociation between switch and WLC is broken when yo remove vlans.
Which WLC is it? is it Cisco? does it have Lag properly configured on the port?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-24-2024 07:55 AM
- Check logs on the 9200 when that happens ,
M.
-- ' Listen to the wind, it talks
Listen to the silence, it speaks
Listen to your heart, it knows
Ganado Mucho (1809 to 1893 ) Navajo Indian
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-24-2024 07:56 AM
What goes down, the vlan interface or the physical interface?
Can you post the configs from both sides of the trunk?
HTH
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-24-2024 08:07 AM
"show logging" shows nothing unfortunately.
I think it's the physical interface that goes down, please see attached screenshot, Gi1/0/3 has operational status red.
I can tell it's down because it is connected to a wireless controller (hence the trunk mode) and all wifi goes down after that.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-24-2024 08:13 AM
   + Are the extra VLAN ids , existing (meaning do those vlans exist)
   + The screenshot seems to show a GUI result ; can you try when configuring
      from CLI and use 'show interface status' too , because working through CLI
      may bring more informative and extra messages , 
M.
-- ' Listen to the wind, it talks
Listen to the silence, it speaks
Listen to your heart, it knows
Ganado Mucho (1809 to 1893 ) Navajo Indian
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-24-2024 08:31 AM
Yes, the VLAN ids do exist and work fine on ethernet.
But it does not have anything to do with these ids. It goes down even with "All IDs allowed" on the trunk setting.
The only difference I see on CLI is that it show the port as suspended.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-24-2024 08:46 AM
  - Can you check the logs on the WLC too , just after doing your operations
  - Is it possible that there could be a mismatch between the WLC-vlan-list on the trunk and the switch-vlan-list on the port ?
M.
-- ' Listen to the wind, it talks
Listen to the silence, it speaks
Listen to your heart, it knows
Ganado Mucho (1809 to 1893 ) Navajo Indian
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-24-2024 08:52 AM
the trunk is entering in Suspended mode and this could be because the WLC and not the switch side. If could be due native vlan mismatch. The negociation between switch and WLC is broken when yo remove vlans.
Which WLC is it? is it Cisco? does it have Lag properly configured on the port?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-24-2024 11:13 AM
The controller is a Cisco 3500.
When the switch port is suspended, the WLC reports the following in the logs:
*spamApTask7: Oct 24 17:58:48.994: %CAPWAP-3-DTLS_CLOSED_ERR: capwap_ac_sm.c:7126 c4:b3:6a:a5:b0:80: DTLS connection closed forAP 10:0:4:31 (5256), Controller: 10:0:4:10 (5246) AP Message Timeout
*spamApTask7: Oct 24 17:58:48.993: %CAPWAP-3-MAX_RETRANSMISSIONS_REACHED: capwap_ac_sm.c:7673 Max retransmissions reached on AP(c4:b3:6a:a5:b0:80),message (CAPWAP_CONFIGURATION_UPDATE_REQUEST
which I believe is the result of losing connection to the APs, 5 of them, the above message repeated 5 times.
As for the Lag... I don't know what this is.
I simply added a WLAN on the controller, WID 8, corresponding to VLAN 8 on the switch. And the problem appears when I add VLAN 8 on the trunk port. Or when I permit all VLANs on it. I am out of options here.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-24-2024 11:38 AM
          - Indeed the capwap messages are only because the APs loose the link with the controller.
            Could  you try the following ; use the CLI to configure or perform next tasks on the switch :
            Issue the command terminal monitor first
            -   shutdown the port gi1/0/3
            -   set it to default : default int  g1/0/3
            -   configure trunk and allow vlans settings again as needed
            issue no shut (when port configuration is finished)
    The idea of the initial terminal monitor command is to have the console messages too on the
    current CLI session , if they would arrive ; 
     perhaps this way it becomes possible to get extra messages as to why the port gets suspended
    when it is activated again
 M.
  
    
-- ' Listen to the wind, it talks
Listen to the silence, it speaks
Listen to your heart, it knows
Ganado Mucho (1809 to 1893 ) Navajo Indian
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-24-2024 11:41 AM - edited 10-24-2024 11:45 AM
I just looked up the LAG term = Link Aggregation. Yes, the WLC is setup with two interfaces connected on the catalyst switch.
The switch on the other hand, has its two physical ports grouped as Port-Channel 1. It seems that whatever change MUST be done on the Port-channel and not on the individual ports. When I add VLAN 8 on the Port-channel or perform any other change for that matter, both physical ports update their values automatically. And they do not get suspended, they remain active.
My problem has been solved. Thank you all immensely for the help.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-24-2024 11:53 AM
Great Jog @isquare

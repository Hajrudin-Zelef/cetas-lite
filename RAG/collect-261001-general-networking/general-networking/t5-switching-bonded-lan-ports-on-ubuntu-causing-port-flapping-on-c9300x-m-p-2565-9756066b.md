---
id: collect-261001-general-networking/general-networking/t5-switching-bonded-lan-ports-on-ubuntu-causing-port-flapping-on-c9300x-m-p-2565-9756066b
title: "t5-switching-bonded-lan-ports-on-ubuntu-causing-port-flapping-on-c9300x-m-p-2565-9756066b"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-general-networking/t5-switching-bonded-lan-ports-on-ubuntu-causing-port-flapping-on-c9300x-m-p-2565-9756066b.md
source_anchor: ""
source_lines: [1, 90]
sha256: 0595b0d2fc18eae2f20fcc9bc3e41707037d4d556f88850aef08ea18b644d7e1
---

# t5-switching-bonded-lan-ports-on-ubuntu-causing-port-flapping-on-c9300x-m-p-2565-9756066b

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-04-2024 04:43 AM
Have a bonded pair of 10GbE ports on a Ubuntu server, bonded in the Round Robin mode, plugged into 2 stacked C9300X switches. There is an alert constantly coming up about MAC address flapping.
Aggregating the 2 ports on the switches causes both ports to get blocked. Anyone who knows Ubuntu/Linux know what the bonding mode should be please?
The Windows servers are working fine with Switch Independent SET configured.
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
12-05-2024 02:51 PM
+1 on cmr's comment
Your Ubuntu/Linux box needs to run LACP. MS/Catalyst running Meraki mode use LACP for link aggregation. If the switch does not receive a reply to its LACP negotiation packets then it will block the port as "....LACP has disabled this port". Is this the error you are seeing on your switchport?
See these KB's below for more info
https://documentation.meraki.com/MS/Port_and_VLAN_Configuration/Switch_Ports#Link_Aggregation
Any doubts, you can raise a ticket with support to see why the aggregation is not forming correctly.
If you found this post helpful, please give it kudos. If my answer solved your problem, click "accept as solution" so that others can benefit from it.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-04-2024 07:19 AM
Hi SimonReach,
Stacking is not stable with CS9300X. I do face the same issue , this is a bug in the firmware.
Previous discussion with Meraki tac:
Once again, thank you for your patience during these past troubleshooting calls. I am writing this email to provide a summary of the various issues that resulted in instability with the C9300 stacks in this network, the changes made to address these issues, and our recommended next steps.
Originally, we observed that ports on the Core stack went into an LACP blocking state when a member of a downstream stack was rebooted. By taking packet captures on the Core stacks, we confirmed that the downstream stacks had stopped sending LACP PDUs following the reboot of their active member, resulting in the Core stack going into an LACP blocking state. Further investigation confirmed that this was due to an internal issue in which a C9300 stack running CS 16.x firmware may not properly reapply its LACP config following the reboot of the active member of the stack.
To address this CS 16.x LACP issue, we tested upgrading the Core stack and one downstream stack to Beta firmware version CS 17.1.2.1, as a fix for that issue is included in this firmware. After performing this upgrade and testing failover again, we observed the LACP instability issue did not persist, however, a separate issue occurred that caused the downstream stack to go offline for 10-20 minutes until the downstream stack automatically rebooted. After examining the boot reason, we confirmed that this behavior was caused by a separate internal issue in which a C9300 stack running CS 17.x firmware does not properly failover when a stack member is rebooted.
Due to the two issues above, the decision was made to revert the firmware to stable release CS 16.8 and remove LACP configs. Once this was completed and the network was stabilized, we again tested the failover by power cycling the active member of a downstream stack once more. After doing so, the failover time was approximately 4 minutes. Due to the necessary RSTP convergence needed for this failover, along with the expected behavior for C9300 stack failover times, this duration would be considered within normal range.
Going forward, we recommend continuing to monitor the network for stability in its current configuration. Once a stable firmware patch is released that includes a fix for both the LACP instability issue in CS 16 and the stack failover instability issue in CS 17, an additional update will be provided on this case.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-04-2024 07:58 AM
Thank you for the response VivekT, we've got a brand new dual C9300X stack running CS 17.1.4 and we've had the issue with port aggregation, we had issues with 1 of the switches running out of memory and falling over and taking out the entire stack for 5 hours before it came back up and we had issues with DNS/DHCP not working reliably at all. It had been in place for a day before it went down but luckily it went down at midnight rather than midday.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-05-2024 01:43 PM
@SimonPeart you should be using LACP from the server, not round robin. Is that an option in the version you are using?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-05-2024 02:51 PM
+1 on cmr's comment
Your Ubuntu/Linux box needs to run LACP. MS/Catalyst running Meraki mode use LACP for link aggregation. If the switch does not receive a reply to its LACP negotiation packets then it will block the port as "....LACP has disabled this port". Is this the error you are seeing on your switchport?
See these KB's below for more info
https://documentation.meraki.com/MS/Port_and_VLAN_Configuration/Switch_Ports#Link_Aggregation
Any doubts, you can raise a ticket with support to see why the aggregation is not forming correctly.
If you found this post helpful, please give it kudos. If my answer solved your problem, click "accept as solution" so that others can benefit from it.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-06-2024 01:48 AM
Thank you to you and @CMR, the bonding mode was changed to Mode 1, this is Active/Backup mode so is a switch independent mode so doesn't require the aggregation of the ports. I want to avoid port aggregation for the time being until all the faults with the switches are resolved.

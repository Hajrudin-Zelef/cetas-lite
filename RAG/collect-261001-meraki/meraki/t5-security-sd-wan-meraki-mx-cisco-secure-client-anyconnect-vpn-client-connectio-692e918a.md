---
id: collect-261001-meraki/meraki/t5-security-sd-wan-meraki-mx-cisco-secure-client-anyconnect-vpn-client-connectio-692e918a
title: "t5-security-sd-wan-meraki-mx-cisco-secure-client-anyconnect-vpn-client-connectio-692e918a"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/t5-security-sd-wan-meraki-mx-cisco-secure-client-anyconnect-vpn-client-connectio-692e918a.md
source_anchor: ""
source_lines: [1, 85]
sha256: 96a459f690acad2e045315c0e339130ba849a5059e0ba15a037d0fbda66ee28c
---

# t5-security-sd-wan-meraki-mx-cisco-secure-client-anyconnect-vpn-client-connectio-692e918a

Meraki MX Cisco Secure Client (AnyConnect) VPN client connection failover question
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-20-2023 07:54 AM
Our company has implemented Cisco AnyConnect VPN connection with our users successfully. Works great with the authentication with the Radius server and also the Certification Authentication. Below is the 2 questions need to see if anyone knows in the Meraki Community....
Question #1....
Is there any way to have the users Cisco AnyConnect client automatically be updated in our previous Cisco Firepower ASA appliance. We are currently deploying updates there ManageEngine but curious if anyone knows if Meraki will have the ability to update our clients AnyConnect Secure Client application on our users workstations?
Question #2....
We implemented "Load Balancing" with both our circuits and when one of our circuits fails(Ex: Atlanta) our users that our VPN into that network location will failover to the secondary location (Ex: Denver). Since we have "Load Balancing" configured in our Atlanta network location why wouldnt the secondary circuit take over and allow our VPN users to stay connected. Our users do experience about a 3-5 minutes outage to failover over to our secondary location. Is there any additional setting I am missing on the Meraki Dashboard to allow the VPN failover stay static in the Atlanta network location or not allow the 3-5 minute outage our users experience.
Thank you!
- Labels:
- 
						
							
		
			Meraki
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-20-2023 08:04 AM
Please, if this post was useful, leave your kudos and mark it as solved.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-20-2023 01:15 PM
Cisco is transitioning AnyConnect management for all platforms (Firepower, ASA and MX) to the Cisco SecureX portal. You can also manage the AnyConnect versions and perform profile management there. It can automatically push out new versions.
I use this system for maybe 90% of new AnyConnect deployments.
To use the SecureX platform, you need two Cisco security technologies licenced. If you don't have another, then the easiest way is to buy a single Cisco Umbrella licence.
You can start the process of signing up by going here:
Once you have setup Cisco SecureX, you go to the "Insights" tab to setup AnyConnect. It does not seem very intuitive to me to have it there - but that is where it is located.
I'm not sure about your question #2. I assume you have failover configured in the AnyConnect profile. I know it works, but I'm not sure on the expected timeframes.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-08-2024 02:02 PM
Hi Philip,
Now that SecureX is EoL, is there a replacement tool for Secure Client management?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-08-2024 06:16 PM
We have been told there is a replacement system but it has not been released yet.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-14-2024 07:22 AM
There will be a tool for client management, but stay tuned for that.

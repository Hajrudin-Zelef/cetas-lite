---
id: collect-261001-general-networking/general-networking/t5-wireless-dhcp-snooping-configuration-for-svi-m-p-4919223-f93ceb98
title: "t5-wireless-dhcp-snooping-configuration-for-svi-m-p-4919223-f93ceb98"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-wireless-dhcp-snooping-configuration-for-svi-m-p-4919223-f93ceb98.md
source_anchor: ""
source_lines: [1, 126]
sha256: 52a5e0c451f632090b7c7b76774b631b17c8af5f4d219f6a4a20e09b29d322fc
---

# t5-wireless-dhcp-snooping-configuration-for-svi-m-p-4919223-f93ceb98

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-06-2023 01:25 PM
Hello Professionals,
I have configured DHCP pool from my Cisco 9300L switch, but when I tried to get IP address through wireless, I'm getting weired IP address.
From Cisco 9300>>
- DHCP Pool: 10.28.28.0/23 (SVI: 10.28.28.1)
- IP dhcp excluded-address: 10.28.28.1 - 10.228.29.100
From Cisco 2504 WLC>>
- DHCP proxy: Disabled
- Primary DHCP server: 10.28.28.1
but client device getting IP like this:
- IP: 10.28.28.25 / 255.255.254.0 / 10.28.28.68
Client got an IP from 10.28.28.68 (NVR system), and when I send DHCP request, 3 more other devices responded me.
I would like to configure DHCP snooping to my switch - SVI (currently SVI is the one having a dhcp pool).
Can you please tell me how can I setup?
I appreciate your response.
Solved! Go to Solution.
- Labels:
- 
						
							
		
			Catalyst 9000
- 
						
							
		
			Catalyst Switch
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-08-2023 03:51 PM
You can enable "dhcp snooping" to prevent clients getting IP from rogue DHCP servers. In your case since 9300 switch itself the DHCP server, you have to test it out.
Below my blog post may help you to get it test
https://mrncciew.com/2012/12/27/understanding-dhcp-snooping/ 
HTH
Rasika
*** Pls rate all useful responses ***
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-07-2023 10:34 PM
Where is the 2504 WLC connected ? Is that connect to same 9300L switch ?
Rasika
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-08-2023 08:10 AM - edited 09-08-2023 08:12 AM
Hello @Rasika Nayanajith
Yes, WLC2504 is connected to 9300L switch. Also, I figured out 4 IP addresses are Rogue DHCP.
Those are NVR system, and they have a 'DHCP server' feature. For now, my issue has been resolved temporarily by disabling NVR's feature.
I was expecting WLC is sending the DHCP request to Core switch, cause each WLAN has IP and gateway(which is Core switch SVI). but looks like it only gets DHCP respond from rogue DHCP servers.
Could you recommend a command to manage rogue DHCP for preventing further issue?
Thank you for your comments.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-08-2023 03:51 PM
You can enable "dhcp snooping" to prevent clients getting IP from rogue DHCP servers. In your case since 9300 switch itself the DHCP server, you have to test it out.
Below my blog post may help you to get it test
https://mrncciew.com/2012/12/27/understanding-dhcp-snooping/ 
HTH
Rasika
*** Pls rate all useful responses ***
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-11-2023 09:14 AM
Hello Rasika,
I read your blogs and I believe that is very useful and helpful to me. 
I have to wait until our production break and then I can try some test followed your guide.
One thing I have a question for you is about trust port.
On my drwaing, two clients are connected each switches directly so I need to configure DHCP snooping on both switches.
If I'm right, Uplinks (Gi1/1/1 for both switches) will be the trusted ports.
However, what if rogue DHCP connected to Core switch like this? Should I configure DHCP snooping to Core switch?
Appreciate you to providing precious comments.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-12-2023 03:20 AM
In your case if core sw provide DHCP service, then you have to enable those uplink ports from downstream switch to configure as trusted ports.
To prevent any rougue DHCP servers connecting on core switch issuing IPs, you have to enable it on that switch as well (without trusting any ports).
Testing is the best way to verify the operation.
HTH
Rasika
*** Pls rate all useful responses ***
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-12-2023 06:39 AM
Thank you Sir, I will test.

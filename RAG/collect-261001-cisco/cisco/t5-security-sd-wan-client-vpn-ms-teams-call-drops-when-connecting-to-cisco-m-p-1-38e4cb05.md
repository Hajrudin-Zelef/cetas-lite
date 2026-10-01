---
id: collect-261001-cisco/cisco/t5-security-sd-wan-client-vpn-ms-teams-call-drops-when-connecting-to-cisco-m-p-1-38e4cb05
title: "t5-security-sd-wan-client-vpn-ms-teams-call-drops-when-connecting-to-cisco-m-p-1-38e4cb05"
domain: cisco
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/t5-security-sd-wan-client-vpn-ms-teams-call-drops-when-connecting-to-cisco-m-p-1-38e4cb05.md
source_anchor: ""
source_lines: [1, 121]
sha256: 23cd300051bf467b4a78d7478b5bfacb907aa1119eb487d164af0128c09f308b
---

# t5-security-sd-wan-client-vpn-ms-teams-call-drops-when-connecting-to-cisco-m-p-1-38e4cb05

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-19-2022 01:46 AM
When a user connects to Client VPN using Cisco AnyConnect client during a MS teams call the call drops for a few seconds and then reconnects.
Dynamic Client Routing is enabled with " Send all traffic except traffic going to these hostnames" enabled which has "microsoft.com" included
What can cause this call drop? Is there a way to fix it?
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
09-19-2022 02:03 AM
"microsoft.com" is not enough to exclude from the tunnel. Here is a Cisco Tech-Note on the needed exclusions:
If you found this post helpful, please give it Kudos. If my answer solves your problem, please click Accept as Solution so others can benefit from it.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-19-2022 02:03 AM
"microsoft.com" is not enough to exclude from the tunnel. Here is a Cisco Tech-Note on the needed exclusions:
If you found this post helpful, please give it Kudos. If my answer solves your problem, please click Accept as Solution so others can benefit from it.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-19-2022 02:16 AM
I just smashed this together. Copy and paste this into Powershell. It sends a query to Microsoft for the list of IP addresses used by Microsoft Teams. Exclude these from your full tunnel VPN.
$clientRequestId = [GUID]::NewGuid().Guid
$uri = "https://endpoints.office.com/endpoints/worldwide?NoIPv6=true&clientRequestId=$clientRequestId"
$endpointSets = Invoke-RestMethod -Uri ($uri)
$Optimize = $endpointSets | Where-Object { $_.category -eq "Optimize" }
$optimizeIpsv4 = $Optimize.ips | Where-Object { ($_).contains(".") } | Sort-Object -Unique
Write-Host $optimizeIpsv4
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-19-2022 02:28 AM
This one is great for Windows users!
If you found this post helpful, please give it Kudos. If my answer solves your problem, please click Accept as Solution so others can benefit from it.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-19-2022 02:17 AM
OR - don't use full tunnel. Specify only the list of subnets required that AnyConncet users will need to access inside of your organisation.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-19-2022 03:13 AM
...afterwards realize that you'll need another layer of security because there's no centralized endpoint protection anymore. Let me warmly welcome you to the funny world of SASE!
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-19-2022 03:25 AM
And what are you going to do for the user when they are not connected to VPN ... not much difference really is there?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-19-2022 03:36 AM
Adding a point to the question - The whole idea of using full tunnel with Cisco Any connect here is to whitelist the public IP to Azure resources like Virtual machines, SQL managed instances etc.
So that people get onto the VPN to connect to these resources & not really be able to connect from a different network
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-19-2022 06:40 AM
If that IP address is static you could include that in the AnyConnect encryption domain and not use full tunnel.
Use whichever approach is easier for you.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-20-2022 12:20 AM
Thank You All. Adding all the IPs from the below list helped.
Appreciate your help @Philip D'Ath @Christian_Ney @Karsten Iwen

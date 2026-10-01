---
id: collect-261001-cisco/cisco/t5-routing-and-sd-wan-how-to-configuration-default-route-and-nat-on-asa-firewall-80c64d2a
title: "t5-routing-and-sd-wan-how-to-configuration-default-route-and-nat-on-asa-firewall-80c64d2a"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["amd"]
source: docs/RAG/collect-261001-cisco/t5-routing-and-sd-wan-how-to-configuration-default-route-and-nat-on-asa-firewall-80c64d2a.md
source_anchor: ""
source_lines: [1, 98]
sha256: 51af306dfdc2823fbcda30688359e5e06c9fa1629ffb238032bafaa9354d9e70
---

# t5-routing-and-sd-wan-how-to-configuration-default-route-and-nat-on-asa-firewall-80c64d2a

How to configuration Default route And NAT on ASA firewall 5500 series
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-20-2023 05:48 AM
How to confugtion Defualt route Amd NAT on ASA firewall 5500 series , ISP device assigend to my LAN using VLAN , I Have not Public Ip address how to confugtion ISP 's VLAN address as Defualt ip And confugtion NAT ?
regards
Habtemariam
- Labels:
- 
						
							
		
			Routing Protocols
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-20-2023 06:00 AM
start from here :
good video :
https://www.youtube.com/watch?v=3EbfsIUA8sc
=====️ Preenayamo Vasudevam ️=====
***** Rate All Helpful Responses *****
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-20-2023 01:56 PM
Hello
You can use after-auto section 3 nat on the ASA.
route outside 0 0 x.x.x.x <next hop ip address>
nat (inside,outside) after-auto source dynamic any interface
Please rate and mark as an accepted solution if you have found any of the information provided useful.
This then could assist others on these forums to find a valuable answer and broadens the community’s global network.
Kind Regards
Paul
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-20-2023 01:59 PM
routing 
route OUT 0.0.0.0 0.0.0.0 <Router IP>
NAT
NAT(IN,OUT) dynamic interface 
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-20-2023 02:48 PM
Hello
@MHM Cisco World wrote:
NAT
NAT(IN,OUT) dynamic interface 
This will work however it will default to section 1 nat, and with the three nat sections of the ASA, section 1 takes preference, so if the requirement for more specific manual nat statement is added in the future it could become obsolete and not work if preference wasn't given to it within section 1 over the default pat statement, hence its suggested to append a default pat statement at the very end of the ASA nat order (section 3) 
Please rate and mark as an accepted solution if you have found any of the information provided useful.
This then could assist others on these forums to find a valuable answer and broadens the community’s global network.
Kind Regards
Paul
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-20-2023 03:07 PM
Yes you are right but I dont thing he use manual NAT for 1:1 NAT or static PAT.

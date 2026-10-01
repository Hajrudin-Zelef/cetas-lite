---
id: collect-261001-cisco/cisco/t5-vpn-vpn-from-asa-using-same-subnet-both-ends-m-p-2298419-5b038792
title: "t5-vpn-vpn-from-asa-using-same-subnet-both-ends-m-p-2298419-5b038792"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/t5-vpn-vpn-from-asa-using-same-subnet-both-ends-m-p-2298419-5b038792.md
source_anchor: ""
source_lines: [1, 63]
sha256: c4a98a6bdf8cdf4714430ac9a734d5dad08f4fe35558ba2330e1b2d30f0356b7
---

# t5-vpn-vpn-from-asa-using-same-subnet-both-ends-m-p-2298419-5b038792

VPN from ASA using same subnet both ends
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-01-2013 03:29 AM
Hi
I am looking for advice on a potential DR solution and hopefully this is the place, I am a newbie so please go easy on me:)
I am looking at a DR solution which we would ideally like to use the same IP range at dual locations due to replicating a large number of virtual machines and not wanting to change all their IP numbers in DR situation. The other scenario we have to account for is we have a number of physical boxes that have to talk to each other for replication.
I have a query around VPN access relating to potential DR changes we may make in future. I will give scenario first, apologies if this looks long but is not that bad I don’t think:
- •- We are looking at using same subnet IP range and VLAN’s at DR as our head office for our hosted DMZ’s
- •- Certain servers in each subnet will need to talk to a related server at each site which has the matching subnet
- •- We would give servers at both sites different IP numbers which would technically be in same IP range but at the different locations
So for example we would have the following servers in a 192.168.222.0/24 subnet that would need to communicate with each other:
ServerA at Head office 192.168.222.10 /24 IP Address
ServerB at DR 192.168.222.20 /24 IP Address
We would create a VPN at Headoffice with crypto of 192.168.222.10 Inside, 192.168.222.20 remote network and vice versa at DR.
At this point ServerA in Headoffice 192.168.222.0/24 subnet would not be able to contact ServerB at DR as it would think 192.168.222.20 is in its local subnet and would not send traffic to its default gateway to get to the remote server and the VPN would be inactive still.
Next step we create a NAT on the Headoffice ASA for ServerB on the outside interface as 192.168.222.20 to inside 192.168.222.20 which means the Headoffice asa would pickup any traffic on the inside aimed at ServerB IP address and bring up the VPN to ServerB at DR, we also create a NAT at DR for ServerA on the outside as 192.168.222.10 to inside 192.168.222.10 so when the server replies the DR ASA picks up traffic and send back over the VPN to Headoffice.
I have tested this using two ASA 5505’s and appears to work fine, it wasn’t as complicated to setup as it probably sounds and would mean our DR subnets and firewall rules could match live making change requests a ton easier along with managing DR rules and IP ranges etc.. We could then even look at refreshing the DR code to match the live on periodic basis if we felt the need knowing the DR rules are meant to match the live. This also means that the VM’s we sync at SAN level to DR can just be brought up and would run fine at DR without changing Ip number or anything so the benefits are pretty massive if there is no or little reason not to do it this way.
So my question is, is there any reason we should not do this or anything I missed? Thanks anyone for advice on thisand hope it makes sense.
- Labels:
- 
						
							
		
			VPN
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-04-2013 02:21 PM
It sounds like your heading in the right direction. If you build a site to site VPN tunnel then I would suggest using a NAT exempt for that subnet. Also I have seen a solution where a company's HQ is connected to the DR site via a MetroE connection.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-07-2013 01:59 AM
Thanks Rashid. My concern was/is that this may not be a good way of achieving this, am I right to think from your reply it would be and is a method used by others? Dues to my limited experience with Cisco networking at this level I am unsure of this and would hate to set something up that someone else comes in and asks what on earth you done it like this for.
Cheers
Steve

---
id: collect-261001-general-networking/general-networking/t5-network-security-ask-the-experts-configuring-and-troubleshooting-nat-and-fail-e73238e0
title: "t5-network-security-ask-the-experts-configuring-and-troubleshooting-nat-and-fail-e73238e0"
domain: general-networking
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-network-security-ask-the-experts-configuring-and-troubleshooting-nat-and-fail-e73238e0.md
source_anchor: ""
source_lines: [1, 127]
sha256: 47914d44e2f8969fad3476f46bd7b4e7939a572730453d3f3e59fb3b0b92c544
---

# t5-network-security-ask-the-experts-configuring-and-troubleshooting-nat-and-fail-e73238e0

ASK THE EXPERTS : Configuring and Troubleshooting NAT and Failover on ASA Firewall
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-20-2011 09:21 AM - edited 03-11-2019 01:47 PM
Welcome to the Cisco Support Community Ask the Expert conversation. This is an opportunity to learn expert tips on how to configure and troubleshoot Network Address Translation (NAT) and Failover on Cisco ASA Firewalls with Cisco Expert Amitashwa Agarwal. Amitashwa is a senior customer support engineer and technical lead at the Cisco Technical Assistance Center in Bangalore, India. He works with the Security Firewall team, where his areas of expertise include configuring and troubleshooting issues related to firewall, VPN, and AAA technology. He holds a bachelor's degree in computer science from the University of Pune, India, and holds CCSP and CCIE certifications in Security (#22164).
Remember to use the rating system to let Amitashwa know if you have received an adequate response.
Amitashwa might not be able to answer each question due to the volume expected during this event. Remember that you can continue the conversation on the shortly after the event. This event lasts through July 1st, 2011. Visit this forum often to view responses to your questions and the questions of other community members.
- Labels:
- 
						
							
		
			NGFW Firewalls
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-30-2011 12:04 PM
Hi,
Source based routing or Policy based routing is not supported on ASA. Therefore, based on the source of the traffic ASA cannot take any routing decision. Routing on ASA works as follows:
STEP 1: Check if there is any translation for the destination:
i) If yes, send the packet to the destination interface as per the translation and then do a route lookup on that interface for the real ip address of the destination.
ii) If no then go to STEP 2
STEP 2: Do a route lookup for the destination ip address.
(Usually, for all outbound traffic (going from high to low security level) it follows STEP 2 to route packets and for all inbound traffic (coming from low to high security level) it follows STEP 1 to route packets)
Therefore, keeping the above in mind a workaround to the issue could be:
To translate the 172.21.0.0/16 network behind the DMZ to some other unused subnet on the Outside and ask the users on 10.1.163.0/24 subnet to connect to that unused subnet (say 172.22.0.0/16) instead of 172.21.0.0/16 subnet to get to resources on the DMZ.
static (DMZ,Outside) 172.22.0.0 172.21.0.0 mask 255.255.0.0
access-list outside_in permit ip 10.1.163.0 255.255.255.0 172.22.0.0 255.255.0.0
access-group outside_in in interface Outside
and at the same time allow users on Outside to continue to use 172.21.0.0/24 to get to resources on the Management interface:
static (Management,Outside) 172.21.0.0 172.21.0.0 mask 255.255.0.0
access-list outside_in deny ip 10.1.163.0 255.255.255.0 172.21.0.0 255.255.0.0
access-list outside_in permit ip any 172.21.0.0 255.255.0.0
(Note: The above configuration will only work if the ASA has specific routes to get to the destination subnets in 172.21.0.0/16 subnet through the DMZ and the Management interface respectively. This is required as we cannot configure 2 routes to the same destination subnet through 2 different interfaces on the ASA, something like:
route DMZ 172.16.0.0 255.255.0.0 x.x.x.x
route Management 172.16.0.0 255.255.0.0 y.y.y.y)
Here is another example of overlapping subnets on ASA:
http://www.cisco.com/en/US/docs/security/asa/asa72/configuration/guide/cfgnat.html#wp1043610
Hope this helps.
Regards,
Amit
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-30-2011 12:22 AM
Hi Amit ,
Ali Ahmad again ! First of all, thanks for your response. Your scenario is Ok in case PC is on the local subnet but I want to remotely manage it for which secondary FW must have dynamic routes for remote subnets. I think that issue lies there ?
Can you elaborate how the secondary FW can obtain remote subnets through OSPF and whether it is possible. In case of any possibility , Kindly highlight a config. scenario like you did before. Thanks
Best Regards,
Ali Ahmad
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-30-2011 09:58 AM
Hi Ali,
In case of Active/Standby failover the standby unit cannot participate in dynamic routing nor the routes learnt by the active unit will get pushed over to it.
Here is a link that confirms the same:
http://www.cisco.com/en/US/docs/security/asa/asa72/configuration/guide/failover.html#wp1052476
Regards,
Amit
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-30-2011 04:28 AM
ASA's at our remote sites typically have at least three interfaces - INSIDE, MPLS, and OUTSIDE.
We seem to always have problem w/ NAT between INSIDE & MPLS due to overlapping subnets.
The 10.0.0.0/8 supernet routes through MPLS, while a more specific subnet such as 10.1.1.0/24 routes through INSIDE.
In the routing world, routers can easily identify what to do - more specific wins.
However, the same rule doesn't seem to apply to NAT'ing on the ASA's - it has problem NAT'ing between interfaces that have overlapping subnets, even though one is more specific than the other.
Is that just the nature of NAT'ing on the ASA's?
Do you have any tips or suggestions around this issue?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-30-2011 12:12 PM
Hi,
ASA takes routing decisions as specified below:
STEP 1: Check if there is any translation for the destination:
i) If yes, send the packet to the destination interface as per the translation and then do a route lookup on that interface for the real ip address of the destination.
ii) If no then go to STEP 2
STEP 2: Do a route lookup for the destination ip address.
(Usually, for all outbound traffic (going from high to low security level) it follows STEP 2 to route packets and for all inbound traffic (coming from low to high security level) it follows STEP 1 to route packets)
Here is an example of overlapping subnets configuration on ASA:
http://www.cisco.com/en/US/docs/security/asa/asa72/configuration/guide/cfgnat.html#wp1043610
See if this helps. Let me know in case of further questions or concerns.
Regards,
Amit
- « Previous
- Next »

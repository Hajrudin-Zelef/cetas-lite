---
id: collect-261001-meraki/meraki/t5-security-sd-wan-meraki-mx-high-availablity-m-p-116120-b8fbb6b5
title: "t5-security-sd-wan-meraki-mx-high-availablity-m-p-116120-b8fbb6b5"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["throughput"]
source: docs/RAG/collect-261001-meraki/t5-security-sd-wan-meraki-mx-high-availablity-m-p-116120-b8fbb6b5.md
source_anchor: ""
source_lines: [1, 152]
sha256: 57f05be6b9d24c295271d1bc5595efe75715b5360ae438d878418f1796e91fdb
---

# t5-security-sd-wan-meraki-mx-high-availablity-m-p-116120-b8fbb6b5

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-30-2020 09:08 PM
Dear Team,
We have two MX Meraki FW and Two Cisco Switches in network , 2 ISP
we need load sharing with High availability
SO i need to know what connectivity we should make to achive high availabity and what configuration need to made
I go through meraki document but it is confusing so i need help from your side .
Actually main thing to connect WAN
So how wan should be connect and LAN Should be conncet and both Meraki
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
07-01-2020 07:14 AM
If both ISPs give you a single cable then you need a basic dumb layer 2 switch for each (Dlink or similar) as each connection needs to be plugged into both MX firewalls.
Plug ISP1 to WAN1 on primary MX
Plug ISP2 to WAN2 on primary MX
Both ISPs need to give you four IP addresses in the same subnet, one for their device, one for each MX and one virtual IP to move between them, they will form an active/passive pair.
Configure the LAN IP for the MX (there is only one for both)
Configure WAN1 with an IP address from ISP1s subnet and set the gateway to the IP used on their device
Configure WAN2 with an IP address from ISP1s subnet and set the gateway to the IP used on their device
Both WAN ports should now show as up on the dashboard with public IPs as you have set
Click on Configure warm spare and select Use virtual uplink IPs
Enter the virtual IP for each ISP in the boxes called WAN X shared IP
Both WAN ports should now show as up on the dashboard with virtual public IPs as you have just set
You will now have Primary and Spare click on spare to access the warm spare MX and configure the two ISP public IP addresses and gateways using the remaining addresses from the ISP subnets.
Plug one of the LAN ports on the primary MX to one of your switches
Plug one of the LAN ports on the spare MX to the second switch with the switch port on the same VLAN and able to see the port on the first switch (trunk interconnect / stack etc.)
Plug in the ISP1 cable from the layer 2 switch to WAN1
Plug in the ISP2 cable from the layer 2 switch to WAN2
Make sure you have spanning tree configured on the switches, then connect a second LAN cable from each MX to the other switch (i.e. MX1 to switch2 and MX2 to switch1) where again both ports are in the same VLAN as the first two.
Config is now complete.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-01-2020 07:14 AM
If both ISPs give you a single cable then you need a basic dumb layer 2 switch for each (Dlink or similar) as each connection needs to be plugged into both MX firewalls.
Plug ISP1 to WAN1 on primary MX
Plug ISP2 to WAN2 on primary MX
Both ISPs need to give you four IP addresses in the same subnet, one for their device, one for each MX and one virtual IP to move between them, they will form an active/passive pair.
Configure the LAN IP for the MX (there is only one for both)
Configure WAN1 with an IP address from ISP1s subnet and set the gateway to the IP used on their device
Configure WAN2 with an IP address from ISP1s subnet and set the gateway to the IP used on their device
Both WAN ports should now show as up on the dashboard with public IPs as you have set
Click on Configure warm spare and select Use virtual uplink IPs
Enter the virtual IP for each ISP in the boxes called WAN X shared IP
Both WAN ports should now show as up on the dashboard with virtual public IPs as you have just set
You will now have Primary and Spare click on spare to access the warm spare MX and configure the two ISP public IP addresses and gateways using the remaining addresses from the ISP subnets.
Plug one of the LAN ports on the primary MX to one of your switches
Plug one of the LAN ports on the spare MX to the second switch with the switch port on the same VLAN and able to see the port on the first switch (trunk interconnect / stack etc.)
Plug in the ISP1 cable from the layer 2 switch to WAN1
Plug in the ISP2 cable from the layer 2 switch to WAN2
Make sure you have spanning tree configured on the switches, then connect a second LAN cable from each MX to the other switch (i.e. MX1 to switch2 and MX2 to switch1) where again both ports are in the same VLAN as the first two.
Config is now complete.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-04-2020 04:32 AM
Dear
I really appriciate with your reply with deep understanding
More few question which in my mind is
1. How many IP require (Same Subnet ) For Wan1, Wan2 and LAN --->
Correct me if i am wrong --> As per i understand in your reply that we should configure one IP for Device and One IP for Virtual so it will require only two --> and what for second device ?
2. Our requirement is we need active active for Meraki MX
Your given configuration is related to high availability right ?
so how can we do active active for MX ? is there any limitation ?
Thanks for your reply
Regards,
Harmesh Yadav
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-04-2020 05:11 AM
@harmesh88 glad to be of help, answers below to your further questions:
1) You are correct, one for primary MX, one for vIP and you do need a third for the second MX. The ISP also needs to use one on the equipment that the line terminates in so effectively the minimum subnet you need is a /29.
2) MX can only be active/active in when deployed purely as a VPN concentrator. Why do you need active/active, with the virtual IP it appears to one firewall to the outside and it always appears as one to the inside.
Cheers,
Charles
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-04-2020 10:17 AM
actually my client want active active
if i will go for active active should i follow same method which you provided right ?
or other ?
And same IP will be use for Active active
Regards,
Harmesh Yadav
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-04-2020 11:26 AM
If you want to use in routing mode then MXs cannot be active/active. Why would they want active/active, is the combined internet throughput over 6Gb/s?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-25-2021 11:58 PM
@CMR Hi, applauds to your support.
I would like to know one thing here in this...
what is the relation between "active/active configuration, and the combined internet throughput over 6Gb/s?"
Thanks in advance !
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-26-2021 12:23 AM
Hi @RAMESHKAKU49329,
The MX with the highest throughput, as of today, is the MX450 and it can manage 6Gb/s either singly or as an active/passive pair. That is what I was referring to above and shown below. If you have an active/active pair then the performance should be greater, but I personally do not have experience of this with Meraki. Note that if you use the advanced security features the performance decreases:

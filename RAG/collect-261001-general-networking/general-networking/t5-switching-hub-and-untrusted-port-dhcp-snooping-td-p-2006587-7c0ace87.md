---
id: collect-261001-general-networking/general-networking/t5-switching-hub-and-untrusted-port-dhcp-snooping-td-p-2006587-7c0ace87
title: "t5-switching-hub-and-untrusted-port-dhcp-snooping-td-p-2006587-7c0ace87"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-general-networking/t5-switching-hub-and-untrusted-port-dhcp-snooping-td-p-2006587-7c0ace87.md
source_anchor: ""
source_lines: [1, 113]
sha256: 61d30594a91f59eae1e3bfc3cbb0a79a4f2b92a0f9a78ecd52511d59137140cd
---

# t5-switching-hub-and-untrusted-port-dhcp-snooping-td-p-2006587-7c0ace87

hub and untrusted port,dhcp snooping
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-10-2012 09:15 AM - edited 03-07-2019 07:10 AM
Hi everybody
The following discussion only focuses on dhcp snooping feature. We further assume only dhcp snooping is configured for all the scenarios i.e no ip source guard feature etc is enabled.
Dhcp snooping combats two issues;
Detection of rogue dhcp server.
Prevention of dhcp pool exhaustion ( which essentially results in DOS attack )
Please consider the following scenario:
Hosts ,h1 and h2 are connected to untrusted port f1/1 of sw. Sw is cnofigured for dhcp snooping.
case 1:
h1 -----hub-------f1/1 sw-----------dhcp sever
h2
h1---mac1
h2----mac2
h2 is turned off. h1 powers up, after few dhcp messages exchange, h1 acquires ip 199.199.199.1. SW which is configured for dhcp snooping( f1/1 is untrusted port), creates an dhcp binding as:
mac 1 199.199.199.1 f1/1 vlan1.
Dhcp snooping mitigates dhcp pool exhaustion by checking dhcp message against the dhcp binding.
In our case, sw will check every dhcp message received on f1/1 to see if it contains client hardware address as mac1 in dhcp message. If not,such dhcp message is dropped
With the above concept in mind, we turn on h2.
H2 sends a dhcp discover message with client hardware address mac2.
What will switch do? Will it not drop the dhcp message?
If sw doesn't drop such dhcp message, the question is then how could we prevent dhcp pool exhaustion by rogue host?
If sw does drop pdhcp message, then we can not connect more than one host to untrusted access port?
================================================================================
case 2:
hosts-------sw------trunk---- f1/1( untrusted) SW2-------------dhcp server.
In above case, only SW2 is configured with dhcp snooping. The goal is to see if dhcp snooping creates multiple dhcp bindings on a untrusted trunk port.
Hosts justs powers up and multiple dhcp discover messages from hosts are sent to sw
The question is: will sw2 be able to create multiple dhcp bindings for all the hosts?
thanks and have a great weekend.
- Labels:
- 
						
							
		
			Other Switching
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-10-2012 10:08 AM
 You can have the switchport port-security maximum 
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-10-2012 11:02 AM
Thanks Karthikeyan
Even if we have port security maximum configured, it will not prevent dhcp pool exhaustion . For example
Dhcp server assigns ip address based on client hardware address in dhcp discover message. As long as dhcp receives different mac addresses in client hardware address field, dhcp server assigns ip addresses.
With the above concept in mind, Lets consider an example to see if port security could prevent dhcp pool exhaustion.
h1------ f1/1----Sw-----Dhcp server.
h1 mac address is mac1
Sw has port security enabled on f1/1 allowing only mac1 and maximum mac address 1.
h1 is rogue host. Various tools can be used to generate dhcp discover message each with different mac address in client hardware address field.
Rogure h1 creates a dhcp discover message with client mac address field set to mac2.
Next h1 simply encapsulates the the dhcp message in ip packet which is then encapsulated in ethernet frame with src mac= mac1 ( the mac address of h1).
Sw receives the frame. Sw does not find anything odd as mac address in src mac indeed matches the mac address allowed by port security on f1/1.
Sw simply forwards the frame to to dhcp sever.
Dhcp sever look at client mac address field and finds it is different mac address and assigns an ip address.
Rogue can continue to craft such dhcp discover messages ; each time with different spoofed mac address and thereby causing DHCP pool exhaustion.
In nutshell, port security can not alone prevent dhcp pool exhaustion.
--------------------------------------------------------------------------------------------------------------------------------------------------------------------------
The questions I am striving to find answers for, are:
Does a switch configured with dhcp snooping allow one single dhcp binding per untrusted access port?
If yes, then no more than one host can connect to such port.
Does a switch configured with dhcp snooping allow multiple dhcp bindings on untrusted trunk port?
If yes, then dhcp pool exhaustion by rogue host can not be prevented.
For example:
In the following example, only SW2 is configured with dhcp snooping while sw1 is left at default settings.
h1-----sw1------trunk--untrusted-f1/1-SW2----dhcp server.
H1 has mac addess mac1. After few dhcp messages exchanges, SW2 creates a dhcp binding;
mac1 199.199.199.1 vlan1 f1/1 1000(seconds)
H1 being rogue host creates another dhcp discover message with mac address ; mac 2
Again Sw2 creates a dhcp dinding after few dhcp message exchange.;
mac2 199.199.199.2 vlan1 f1/1
The process can continue until dhcp server has no ip address left in its pool to assign to genuine hosts.
On the other hand if a switch configured with dhcp snooping allows only one dhcp binding on untrusted trunk port, then not more than one host can connect to network
For example:
Again, Sw2 is configured with dhcp snooping while sw1 is left at default settings:
h1,h2------------vlan1-----sw1----trunk-- f1/1(untrusted)SW2--dhcp server.
SW2 is configured with dhcp snooping.
Let say h1 is first one to boot up.
After few dhcp messages exchanges, SW2 creates a dhcp binding for its untrusted tunk port f1/1 as;
mac1 199.199.199.1 vlan1 f1/1 1000( sec)
Since we assume sw2 can create only one binding for untrusted trunk port, therefore when h2 powers up and sends dhcp discover message with mac2 ,Sw2 will simply drop it.
Thanks and have a great weekend

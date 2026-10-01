---
id: collect-261001-general-networking/general-networking/t5-routing-and-sd-wan-eigrp-amp-hsrp-m-p-4186357-highlight-true-24344479
title: "t5-routing-and-sd-wan-eigrp-amp-hsrp-m-p-4186357-highlight-true-24344479"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-general-networking/t5-routing-and-sd-wan-eigrp-amp-hsrp-m-p-4186357-highlight-true-24344479.md
source_anchor: ""
source_lines: [1, 134]
sha256: fa680372fe08dd6857870bbf6a9773f69e2b75f0ed4eb58d30560125c7363ce1
---

# t5-routing-and-sd-wan-eigrp-amp-hsrp-m-p-4186357-highlight-true-24344479

EIGRP & HSRP
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-19-2020 09:56 PM
Hi All,
Can anyone pls clear the doubt.
Suppose there is a downstream core switch connected to 2 WAN routers (R1 & R2).
Switch is getting EIGRP routes from both WAN routers.
R1 is advertising better route to Switch than R2.
Now suppose, HSRP is configured between R1 & R2. R1 has set hsrp low priority and R2 has high hsrp priority.
Doubt : Which path will be taken by switch. Will EIGRP overide HSRP ?
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
11-20-2020 12:08 AM
If you pointing to VIP IP as default route Gateway - No IGP will not override the HSRP Paths - it always uses ACTIVE HSRP as Path router.
Edge Switch does not require to know all the Routes since it has only 1 exit point - so you need to summary the route make an adjustment on the priority and cost.(metric)
BFD is a good to use case here - high level.
=====️ Preenayamo Vasudevam ️=====
***** Rate All Helpful Responses *****
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-20-2020 01:06 AM
Hello @gauravpundir231 ,
>> Which path will be taken by switch. Will EIGRP overide HSRP ?
remember that most specific route is used first so if the HSRP VIP is used as next-hop for a default route with an standard AD the only route that HSRP can override is the default route.
All destinations with a specific EIGRP route in routing table will use the EIGRP best path/best next-hop.
Hope to help
Giuseppe
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-20-2020 02:09 AM
The original poster asks this questions "Which path will be taken by switch." The answer to this question is that it depends on whether the switch has some static routes or is using just the dynamic routing protocol. If the switch has some static routes, and if the static routes specify the HSRP address as next hop then the path taken will be the HSRP active router. For any routes in the routing table advertised by EIGRP the next hop will be the address of the router advertising the better metric.
Perhaps a different way of looking at this question is to understand that HSRP is a protocol to provide next hop redundancy. It is particularly appropriate for devices which depend on a configuration of a gateway - its focus is one hop away. Dynamic routing protocols such as EIGRP focus on destinations which are multiple hops away.
Rick
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-20-2020 04:27 AM
in Campus there is L2 L3 level line, where your L3 level Line end, it is end in R or in SW?
I think because of HSRP is config in boundary (L2/L3) then your SW see the HSRP not see EIGRP.
SW traffic is toward the Active HSRP.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-22-2020 06:47 AM
Hello,
just in case anybody is interested, I was trying to figure out a way to manipulate HSRP based on changes of the feasible distance on the primary HSRP router. Below is an EEM script that checks for changes of the FD value, and decreases the HSRP priority when the FD value changes. The script is based on the output of:
sh ip eigrp topology detail-links | include 8.8.8.8
P 8.8.8.8/32, 1 successors, FD is 2850, serno 2
For this to work, the standby HSRP router needs to be configured with 'standby x preempt' as well.
Primary HSRP Router
ip sla 1
icmp-echo 8.8.8.8 source-interface X
frequency 5
!
ip sla schedule 1 start-time now life forever
!
router eigrp 1
redistribute static
!
ip route 8.8.8.8 255.255.255.255 interface x
!
event manager applet FEASIBLE_DISTANCE_TRACKER_DOWN
event timer watchdog time 60
action 1.0 cli command "enable"
action 2.0 cli command "sh ip eigrp topology detail-links | include 8.8.8.8"
action 2.1 foreach result $_cli_result \n
action 3.0 regexp "[P]+\s+[8]+[.]+[8]+[.]+[8]+[.]+[8]+[\/]+[32]+[,]+\s+[1]+\s+[A-Za-z]+[,]+\s+[A-Za-z]+\s+[A-Za-z]+\s+([0-9]+)+[,]" "$result" match FD
action 4.0 if $FD gt 2850
action 5.0 cli command "conf t"
action 6.0 cli command "interface GigabitEthernet0/0"
action 7.0 cli command "standby 1 priority 90"
action 8.0 cli command "end"
action 9.0 end
!
event manager applet FEASIBLE_DISTANCE_TRACKER_UP
event timer watchdog time 60
action 1.0 cli command "enable"
action 2.0 cli command "sh ip eigrp topology detail-links | include 8.8.8.8"
action 2.1 foreach result $_cli_result \n
action 3.0 regexp "[P]+\s+[8]+[.]+[8]+[.]+[8]+[.]+[8]+[\/]+[32]+[,]+\s+[1]+\s+[A-Za-z]+[,]+\s+[A-Za-z]+\s+[A-Za-z]+\s+([0-9]+)+[,]" "$result" match FD
action 4.0 if $FD le 2850
action 5.0 cli command "conf t"
action 6.0 cli command "interface GigabitEthernet0/0"
action 7.0 cli command "standby 1 priority 110"
action 8.0 cli command "end"
action 9.0 end

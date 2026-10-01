---
id: collect-261001-general-networking/general-networking/t5-switching-multiple-ospf-processes-td-p-1335402-24efa88f-2
title: "t5-switching-multiple-ospf-processes-td-p-1335402-24efa88f"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-general-networking/t5-switching-multiple-ospf-processes-td-p-1335402-24efa88f.md
source_anchor: ""
source_lines: [56, 183]
sha256: 8e41277e4a5c38950a5f145f432c383edb12a5e29e19b15b0e277b364cb98260
---

# t5-switching-multiple-ospf-processes-td-p-1335402-24efa88f

			Other Switching
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-16-2009 03:51 AM
Joe
You are pretty close. Let me try to clarify one of the things that you say:"it is not the fact that the network statement under OSPF encompasses the whole range of subnets that determines which LSAs are created and sent, but the actual interfaces that are configured on each router."
It is not just the network statement and it is not just the interfaces on the router that are configured. It is the combination of the network statement and the configured interfaces that determine what is advertised by OSPF.
To explain in a bit more detail:
When the OSPF process starts it looks at its configured network statements (and the address ranges defined by the address and the mask) and it looks at every interface on the router (that is in up/up state) and if an interface falls into the range defined by a network statement then that interface is included into OSPF. Then OSPF looks at the subnet defined on that interface (including its mask) and advertises that subnet.
So to clarify a couple of points:
- the network statement does not tell OSPF what to advertise but tells OSPF what interfaces to process.
- the network statement does not tell OSPF to summarize (there are separate commands to control summarization).
- OSPF will determine what to advertise based on the configured subnets on the interfaces that it includes in its processing.
HTH
Rick
Rick
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-13-2009 09:42 PM
Why would someone run multiple OSPF processes on the same router? --you can do it but to able to achieve routing between them you have to use the redistribution.
Yes they have their own separate databases.
What about the RIB? Are there effectively separate routing tables? -yes
lets say a route is learned through both OSPF processes, what happens then....?
it will only happen when you redistribute.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-14-2009 03:21 AM
Joe
I would answer a couple of things a bit differently from Manish. There are several reasons why you might choose to run separate OSPF processes. Perhaps your router connects to customer A and you want to run OSPF with customer A so that you can support them. And perhaps your router also connects to customer B. And you want to run OSPF with customer B. But you do not want customer A to see customer B routes. The easy solution for this is to run an OSPF process on the interface connecting to customer A and to run a separate OSPF process on the interface connecting to customer B. In this case your router knows routes of both customers but will not advertise routes from customer B to customer A.
Or perhaps there is a situation where your router will learn certain routes in OSPF but you do not want to advertise some of those routes to other parts of your network. The easy way to accomplish this is to run separate OSPF processes, to redistribute routes between processes, and to filter the redistribution to allow only some routes to be redistributed and advertised by the other process.
The implications of running separate OSPF processes include these:
- an interface can be active in only 1 interface. So each OSPF process will have a unique set of neighbors.
- each OSPF process will learn its own prefixes and maintain those prefixes in its own database.
- each OSPF process will advertise only prefixes from its own data base. The only way to have one process advertise prefixes from the other process is to redistribute.
Yes the separate OSPF processes have separate OSPF data bases and the separate OSPF processes do not share any data with the other OSPF process unless you redistribute.
No there are not separate RIBs. There is a single routing table for the router which will contain routes from both OSPF processes. But each OSPF process will advertise only the routes that are contained in its own database.
If you run 2 OSPF processes and if each process learns the same prefix then each OSPF process will attempt to insert its route into the routing table. If both prefixes have the same metric them both routes will show up in the routing table (and the router will load share traffic to that destination). Note that since both processes can not run on the same interface, the 2 routes that it learns will have different next hop addresses - it is not possible for both OSPF processes to learn the same prefix with the same next hop.
OSPF process 599 may have a no passive-interface command, but if there is not a network statement that matches the address then it will not cause the interface to be included in the process. Having the no passive-interface command under OSPF 599 does no harm, but it does no good either. Perhaps at one time process 599 did include that interface and then it was moved to process 499 but someone forgot to remove the no passive-interface. Note that if both processes did have network statements that match the interface address then the interface can be active in one process but not the other.
HTH
Rick
Rick
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-14-2009 04:41 AM
Rick:
Thank you VERY much for that very informative and extremey helpful answer.
Much of what you said is precisely what I expected, but I couldnt find some good documentation on the subject to verify my thoughts. I know plenty must exist, but I couldnt find anyhing on the Internet and it was getting late....
I think I have all the information I need, but I MAY come back for a bit more.
Thanks again, really appreciate it....
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-14-2009 05:26 AM
Rick:
I have another question. That was fast, huh?
Let's please go back to that router with the OSPF 599 process and ignore 499 for a moment.
There are actually 2 routers running the same OSPF process and "advertising" the same summarized addresses.
The reason is that each router has about 120 /29 vlans in the SAME network ranges that you see under the OSPF process, but one router hosts odd-numbered vlans while the other hosts even-numbered vlans.
Example:
10.195.80.241/29 - VLAN 537 - ODD
10.195.80.249/29 - VLAN 410 - EVEN
Router 1:
router ospf 599
log-adjacency-changes
auto-cost reference-bandwidth 1000000
nsf
area 0 authentication message-digest
area 2 authentication message-digest
redistribute static
passive-interface default
no passive-interface GigabitEthernet4/1
no passive-interface Vlan98
no passive-interface Vlan99
network 10.195.48.8 0.0.0.3 area 2
network 10.195.48.32 0.0.0.7 area 2
network 10.195.48.248 0.0.0.7 area 0
network 10.195.49.0 0.0.0.127 area 2
network 10.195.50.0 0.0.0.127 area 2
network 10.195.64.0 0.0.15.255 area 2
network 10.195.80.0 0.0.0.255 area 2 <---- SAME network range
ROUTER 2:
router ospf 599
log-adjacency-changes
auto-cost reference-bandwidth 1000000
nsf
area 0 authentication message-digest
area 2 authentication message-digest
redistribute static
passive-interface default
no passive-interface GigabitEthernet4/1
no passive-interface Vlan98
no passive-interface Vlan99
network 10.195.48.0 0.0.0.3 area 2
network 10.195.48.32 0.0.0.7 area 2
network 10.195.48.248 0.0.0.7 area 0
network 10.195.49.0 0.0.0.127 area 2
network 10.195.50.0 0.0.0.127 area 2
network 10.195.64.0 0.0.15.255 area 2
network 10.195.80.0 0.0.0.255 area 2 <---- SAME network range
There is NO HSRP between the two.
There is also a routed connection between the two routers (vlan 99). See config.
These routers are also uplinked to 2 other core routers, which we could probably leave out of the dicussion for now.
OK, so the Question:

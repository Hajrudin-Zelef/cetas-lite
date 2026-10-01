---
id: collect-261001-cisco/cisco/questions-707949-cisco-catalyst-2960-switches-can-i-reverse-the-air-flow-1ead25bb
title: "questions-707949-cisco-catalyst-2960-switches-can-i-reverse-the-air-flow-1ead25bb"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-cisco/questions-707949-cisco-catalyst-2960-switches-can-i-reverse-the-air-flow-1ead25bb.md
source_anchor: ""
source_lines: [1, 22]
sha256: e6d89393b719fbc1cc137e22958e74a6ac67a4e81d7d9ed67e7d248dbc10714a
---

# questions-707949-cisco-catalyst-2960-switches-can-i-reverse-the-air-flow-1ead25bb

Cisco 2960 switches pull in cool air from the sides and exhaust to the rear.
Depth-wise they are 1/3 to 1/2 rack-depth (depending on exact model switch and rack).  
This leaves you very few options. If you mount them unmodified backside of the rack just about the entire switch is in the hot zone. That will only be OK if you have cold airflow running on the side of the rack so the switch can get sufficient cooling. Unfortunately this is usually not the case in a strict hot/cold isle setup.
When you mount on the front you will have to run most server-cabling from the back to the front (if your servers have most wiring on the back) which makes for messy cabling.
Reversing the fans is maybe possible. I have no idea if this is electrically feasible and how the firmware would react on a 2960. In most switches that I took apart the fans plug directly into a motherboard connector and are not reversible.
Then you could mount them, ports facing backwards, in the front of the rack. This makes cabling awkward because you will have to reach quite a long way into the rack to reach the RJ45's. It might be OK if you only need to (re-)patch them on very rare occasions. Be prepared to leave 1 U above and below a switch unused, just to give yourself some working room.
Precisely because of these issues nowadays we do it completely different in our bigger server-rooms. We avoid the problem altogether:
 - For each 42U rack we reserve the lower 30U for servers. (Not higher, to difficult to mount/unmount them.)
 - Next 6U is for switches: ports at the front.
On the side of each rack is filler plate, except at the U's with switches so there is some cold airflow to the side-intakes on the switches.
 - Top 6U is for patch-panels: Also ports at the front.
From the back of the patch-panels we just run 8 UTP cables (CAT7) to the back of each of the 30 U's. (Alternating 8 left-side, 8 right-side). That is 30x8= 240 ports which fits in 5x 48 port patch-panels. This cabling is a one-time fixed installation with all the cables made exactly to length and neatly placed in cable-guides/trays in the rack.
The top-most patch-panel is reserved for backbone-cabling to other racks (24x OM3 or OM4 fiber). We have another fiber patchpanel (in some racks 2) backside mounted in the topmost slot(s) for SAN cabling.
We simply hook up all UTP ports (used or not) on the back of the server to the corresponding block of patch-panel ports. (In the rare cases a server needs more than 8 UTP connections we take them from the U above it. Typically such servers are more than 1 U anyway.)
All UTP patching is done front-side. Fiber-SAN stays at the back of the rack.
This way cable management becomes easy: You don't have to thread new cable though the rack every time you change something. The cabling (except for the short patches in the front) is static and can be made exactly to length. There is no over-length to stuff in a corner inside the rack itself. That also helps airflow.  
It is so easy you can talk anybody (liek someone from local FCM on-site who has access to the server-room) through a re-wiring job on the phone if necessary:
Find rack number 5. Big yellow number on the front door. Open that front-door. About chest high and higher you'll see a bunch of cables in several colors. On the left and right there are numbers on the sides of the equipment. They go from 31 at the lowest piece of equipment that has cables in it, all the way up to 42 at the top. Find number 33 on the side, look for the cable in port 21 (should be a blue cable). Pull it loose (press on the little lip on the plug to unlock it) and plug it back in at height 35 in port number 17. Thank you for your help, don't forget to close the door of the rack on your way out.  
Initial cost of setting up a rack is higher, but you recoup that real quickly in labor and downtime when you need to swap servers later on.
Of course that totally depends on how many changes you expect. In our case about 1 server-replacement per 6 weeks in each rack and we deal with about 300 racks in 35 server-rooms on 21 locations all over Europe.
It really pays of in the long term if you don't need to physically go over to each site for small changes.
I get service-techs from HP, DELL, etc. which I just direct over the phone where to place the new server. As soon as the cables are in and I can see the ILO or DRAC on the LAN I can take it from there.

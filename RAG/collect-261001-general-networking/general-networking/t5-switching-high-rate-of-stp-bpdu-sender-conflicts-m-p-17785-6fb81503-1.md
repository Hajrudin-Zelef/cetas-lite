---
id: collect-261001-general-networking/general-networking/t5-switching-high-rate-of-stp-bpdu-sender-conflicts-m-p-17785-6fb81503-1
title: "t5-switching-high-rate-of-stp-bpdu-sender-conflicts-m-p-17785-6fb81503"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-switching-high-rate-of-stp-bpdu-sender-conflicts-m-p-17785-6fb81503.md
source_anchor: ""
source_lines: [1, 43]
sha256: 6f4fb09517212eda95199f8b10009478806775faa6d5e14d2949ce0a76cbdfc4
---

# t5-switching-high-rate-of-stp-bpdu-sender-conflicts-m-p-17785-6fb81503

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-24-2018 12:56 AM
Hi,
I'm experiencing an issue on my Meraki-based network, where a certain set of switches are almost constantly reporting STP BPDU sender conflicts.
Here is some info about my network:
We are situated on a large farm, therefore we have multiple points in different locations on the same network. I have no redundant links between any of my switches; there are only daisy-chains between some of them. Some of the locations are connected wirelessly, using Ubiquiti and Mimosa links.
The Mimosa links are merely two Point to Point links connecting to a facility, with two switches daisy-chained on the other side. Then I have a Ubiquiti Rocket M5 set up as an Access Point. Three different locations connect to them, each using a Ubiquiti Nanobeam M5. There is a Meraki switch on the other side of each of those stations. The Rocket M5 and Mimosa each connects directly to my Root Bridge.
Since I am only getting errors on the switches connected to the Rocket AP, I will refer only to them by name: Admin, Grape Office, Workshop and Controlroom.
Admin connects via a Nanobeam M5 to the Rocket M5.
Grape Office connects via a Nonobeam M5 to the Rocket M5.
Workshop connects via a Nanobeam M5 to the Rocket M5.
Controlroom connects via two point to point Nanobeam M5s to Switch C.
The following events occur frequently in the Event log:
| Apr 24 09:28:07 | Workshop |  | STP BPDU sender conflict | Port 6 received BPDU from ROOT BRIDGE, 19; expected Grape Office, 1 | 
| Apr 24 09:28:06 | Admin |  | STP BPDU sender conflict | Port 1 received BPDU from Grape Office, 1; expected ROOT BRIDGE, 19 | 
| Apr 24 09:28:06 | Workshop |  | STP BPDU sender conflict | Port 6 received BPDU from Grape Office, 1; expected ROOT BRIDGE, 19 | 
| Apr 24 09:28:05 | Admin |  | STP BPDU sender conflict | Port 1 received BPDU from ROOT BRIDGE, 19; expected Grape Office, 1 | 
| Apr 24 09:28:05 | Workshop |  | STP BPDU sender conflict | Port 6 received BPDU from ROOT BRIDGE, 19; expected Grape Office, 1 | 
| Apr 24 09:28:04 | Admin |  | STP BPDU sender conflict | Port 1 received BPDU from Grape Office, 1; expected ROOT BRIDGE, 19 | 
| Apr 24 09:28:04 | Grape Office |  | Port STP change | Port 1 root→designated | 
| Apr 24 09:28:04 | Workshop |  | STP BPDU sender conflict | Port 6 received BPDU from Grape Office, 1; expected ROOT BRIDGE 19 | 
| Apr 24 09:28:03 | Admin |  | STP BPDU sender conflict | Port 1 received BPDU from ROOT BRIDGE, 19; expected Grape Office, 1 | 
| Apr 24 09:28:01 | Workshop |  | STP BPDU sender conflict | Port 6 received BPDU from ROOT BRIDGE, 19; expected Grape Office, 1 | 
| Apr 24 09:28:00 | Admin |  | STP BPDU sender conflict | Port 1 received BPDU from Grape Office, 1; expected ROOT BRIDGE, 19 | 
| Apr 24 09:28:00 | Workshop |  | STP BPDU sender conflict | Port 6 received BPDU from Grape Office, 1; expected ROOT BRIDGE, 19 | 
| Apr 24 09:27:59 | Grape Office |  | Port STP change | Port 1 designated→root | 
It most commonly occurs that the Grape Office switch sends BPDUs to the other switches, but it is not limited to that. Sometimes Admin will send a BPDU to Workshop, Workshop to Grape Office, etc. The Port STP change, however, only ever happens on the Grape Office switch. The Port STP Changes happen 2-6 times each minute, where the BPDU errors happen every 1-3 seconds.
The Grape Office switch is a small 8-port switch. It uses 4 ports: Its Uplink via a Nanobeam M5; A VoIP phone; A PC; A Ubiquiti Unifi that provides WiFi in the office.
When I temporarily disconnect the Grape Office switch remotely, I get no more BPDU errors on the network for the entire time it is disconnected. I do not know why this switch would misbehave, as it is configured the exact same way as all other switches on the network.
Although the network does not come to a complete standstill when it happens, my users frequently complain about disconnecting from my main server or slow internet connections in general.
Any help regarding this issue would be greatly appreciated.
Solved! Go to Solution.
- Labels:
- 
						
							
		

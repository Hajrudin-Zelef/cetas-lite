---
id: collect-261001-general-networking/general-networking/t5-switching-has-anyone-had-issues-with-stacks-and-lacp-after-going-from-14-m-p-6879fe62-1
title: "t5-switching-has-anyone-had-issues-with-stacks-and-lacp-after-going-from-14-m-p--6879fe62"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-switching-has-anyone-had-issues-with-stacks-and-lacp-after-going-from-14-m-p--6879fe62.md
source_anchor: ""
source_lines: [1, 34]
sha256: f880caf3ae5cc02b336279b360351c62d3ca465520c4412cd89bc189b3062f3f
---

# t5-switching-has-anyone-had-issues-with-stacks-and-lacp-after-going-from-14-m-p--6879fe62

Has anyone had issues with Stacks and LACP after going from 14 to 15 code?
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-30-2023 06:55 AM
I'm working on a challenging issue with support right now. We have a somewhat complex network, with about 65 switches, multiple stacks in 30+ buildings all interconnected by 10G single mode fiber. We have a pair of MS425-32's at the core, and every link out to the and MS350-48FP stacks in the buildings is LACP spanning the two core switches. All inter-vlan routing is done at the core.
We were running 14.33.1 code and we went to 15.21.1 code because it attained the "Stable" rating. I upgraded remote sites first because they are simple with only one or two switches. Everything went smoothly so I let our main campus upgrade.
That one went so bad. Half of our switches went down and never came back, stacks were half alive with some switches working and others not. The core looked green, but Core 1 turned out to not be working once the network was put under load. I tried to reboot it and it went down and didn't come back. Unplugging the stack completely and plugging it back in got it to boot again. That became the fix for every switch we have.
Once force-rebooted I thought we were good, but now I have a couple switches that have dropped out of their stacks, then hours later they just pop back in. From the local status pages other switches in the stack show them as there, and the ports appear to be up, but the switch is gone and the light is red. In most cases the ports in the switch don't work. Also if I reboot the core some of my stacks will not reconnect to it until they are also rebooted. My other weird issue is that I'm now suddenly getting UDLD errors on my LACP links heading back specifically to Core 1, Core 2 never has an issue. This could be the stack acting up, or LACP acting up, I'm just not sure.
I'm trying to revert to 14.33.1 code now, but I'm curious if anyone else has had issues anything like this on 15.21.1 or any 15 code?
Thank you,
James
- Labels:
- 
						
							
		

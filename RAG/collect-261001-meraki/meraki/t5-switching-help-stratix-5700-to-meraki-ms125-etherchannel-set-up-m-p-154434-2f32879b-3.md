---
id: collect-261001-meraki/meraki/t5-switching-help-stratix-5700-to-meraki-ms125-etherchannel-set-up-m-p-154434-2f32879b-3
title: "t5-switching-help-stratix-5700-to-meraki-ms125-etherchannel-set-up-m-p-154434-2f32879b"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/t5-switching-help-stratix-5700-to-meraki-ms125-etherchannel-set-up-m-p-154434-2f32879b.md
source_anchor: ""
source_lines: [88, 115]
sha256: 710cc15029fbd7d429a12e61b08d3ef8cca9e3dcca8adf49645f733961a07df9
---

# t5-switching-help-stratix-5700-to-meraki-ms125-etherchannel-set-up-m-p-154434-2f32879b

So the misconfig checks for incoming BPDU's.  If you have a working port-channel on the neighbor switch it will only send BPDU's over one of the links, however if you mis uplink to another switch your switch will receive BPDU's from both switches and the bridge-id will not match.  Following this the spanning-tree etherchannel guard misconfig will fire and will put your ports into err-disable until your err-disable recovery kicks in and re-enables the ports and the cycle continues after receiving new BPDU's
In your case, you are presumably correctly uplinking your switches however you are still receiving BPDU's with different bridge id's how is this possible...?
Normal BPDU's from STP, RSTP and MSTP are sent with the normal destination MAC address 01:80:c2:00:00:00. These are understood by the Meraki MS switch and will be processed. However Cisco switches by default run PVST+ or RPVST and these use 01:00:0c:cc:cc:cd. Meraki switches do not run PVST+ or PVST so these frames are seen as regular multicast frames and are just forwarded on.
But you are running MSTP... so what gives..?
Well in a Cisco MSTP network whenever you have a boundary port. That is a port that links to a switches that runs some other flavor of STP it will do a PVST simulation and will send Cisco PVST+ or RPVST BPDU's on each VLAN on that port for compatibility. So if you connect two cisco switches to the same Meraki MS switch those switches will see both BPDU's from the Meraki switch but will also see each other just as if they were directly connected and this WILL cause problems.
So how to fix this:
2 options
- You turn OFF PVST simulation if the firmware on the Stratix switches allow sit. This way it will prevent those switches to send PVST frames on boundary ports
- If that is not an option or you have another device running PVST+ like a router with a switch module or ports in a bridge group then the other choice is to just disable the guard feature with following global config command: no spanning-tree etherchannel guard misconfig. This will turn off the check but please take care in using lacp where possible and be very carefull around static port-channels.
Good luck!
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-13-2022 11:39 AM
GldenJoe - You are spot on. The problem was BPDU Filtering. The 5700 switches do not have port level control of BPDU guard / filtering via the Device Manager interface. IT has to be enabled / disabled for the entire switch. Although I suspect the CLI interface would allow port level enable / disable functionality. Anyway, Those options were disabled on the Meraki side at the port level but they were left enabled on the 5700. Which was forcing "err-disable until your err-disable recovery kicks in and re-enables the ports and the cycle continues after receiving new BPDU's" . Your last comment set: "please take care in using lacp where possible and be very careful around static port-channels." What issues should I watch for with regard to LACP (active)? I completely understand the second half of that comment because I have witnessed first hand the packet loss associated with static port-channels when a link falls out then comes back up. Thank you for taking the time to respond.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-13-2022 12:51 PM
Oh the only thing I conveyed with that lacp thought is that you should use it as much as possible. Static port-channels are the ones to watch for

---
id: collect-261001-general-networking/general-networking/t5-switching-enabling-spanning-tree-loop-guard-and-bpdu-guard-at-the-same-td-p-2-8e7a7783
title: "t5-switching-enabling-spanning-tree-loop-guard-and-bpdu-guard-at-the-same-td-p-2-8e7a7783"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-switching-enabling-spanning-tree-loop-guard-and-bpdu-guard-at-the-same-td-p-2-8e7a7783.md
source_anchor: ""
source_lines: [1, 99]
sha256: 64d08e7f0f5c2c51082d4c36c63d58a2798281f5219632f6af54bad473d3e8c3
---

# t5-switching-enabling-spanning-tree-loop-guard-and-bpdu-guard-at-the-same-td-p-2-8e7a7783

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-29-2016 07:28 PM - edited 03-08-2019 08:44 AM
Good Evening,
I had some issues with eigrp not working and internet being down once in a while. It comes up once the port connected to the eigrp neighbor is shut down and no shutted. I think it is something related to spanning tree loop being created but none of the ports are blocked. We had loops few years back and I had seen ports blocked. We fixed it and everything was fine.
Recently 3 times, the internet went down and it was back up once the switch was restarted.
Question: Is it possible to enabled loop guard and bpdu guard at the same time on the root switch?
Spanning tree summary shows:
EtherChannel misconfig guard is enabled
Extended system ID is disabled
Portfast Default is disabled
Portfast Edge BPDU Guard Default is disabled
Portfast Edge BPDU Filter Default is disabled
Loopguard Default is disabled
 
Thanks!
Saji
Solved! Go to Solution.
- Labels:
- 
						
							
		
			LAN Switching
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-30-2016 05:30 AM
Hello Saji.
It is possible to have loop guard and bpdu guard enabled on the switch at the same time on the root; however, I would not reccomend enabling bpdu guard on the root switch since bpdu guard is usually enabled on edge ports (port where a single host connects). If you are thinking of enabling it because you want to stop another switch from becoming the root bridge then maybe root guard is what you are looking for. It is actually good practice to have your root bridge configured with root guard and loop guard
Hope that answers your question
Guillermo Fierro
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-30-2016 05:30 AM
Hello Saji.
It is possible to have loop guard and bpdu guard enabled on the switch at the same time on the root; however, I would not reccomend enabling bpdu guard on the root switch since bpdu guard is usually enabled on edge ports (port where a single host connects). If you are thinking of enabling it because you want to stop another switch from becoming the root bridge then maybe root guard is what you are looking for. It is actually good practice to have your root bridge configured with root guard and loop guard
Hope that answers your question
Guillermo Fierro
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-31-2016 10:49 PM
Hi Guillermo,
Thanks for the answer.
My core 6509 has one blade with fiber ports and all other are GB ports. The fiber ports are connected to the other switches and the GB ports are connected to the offices. However, sometimes people connect extra 4 port small netgear switches for an additional drop for a printer. Usually that works fine but if someone connects the end of the same cable on to the same switch, it eventually creates a loop. Which is what I am trying to avoid.
I think enabling root guard is good practice and I will do that shortly. Do you think Loop guard or BPDU guard can be enabled on the GB ports? All I am trying to do is to avoid STP loops.
Thanks!
Saji
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-02-2017 08:40 AM
Hello again Saji.
If you are expecting (and allowing) people to connect a switch on the GBIC's then I would not suggest enabling BPDU guard since it will send the port into an err-disable state. Assuming auto-recovery is not enabled, you would need to manually recover those ports
"All I am trying to do is to avoid STP loops". STP will take care of those loops without you having to do anything. However, the behavior of STP is usually manipulated in order to have the desired path. STP will do its best to select the best path and it is usually pretty good but it is not perfect. Based on what you describe, on the GBIC's I would only enable root guard since my only concern would be having that new switch becoming the root bridge.
Peter: You are right, saying "stop another switch from becoming the root" was a poor choice of words. When I said that I was thinking in a hierarchical design in which it is reccomended to have root guard enabled on every downstream port on every switch which would prevent a rogue switch from becoming the root, but it would not prevent a backup switch from doing so in the event of a failure.
Regards
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-01-2017 02:18 PM
Hello Guillermo and Saji,
Please allow me to join.
Saji: The combination of Loop Guard and BPDU Guard does not really make sense even though you might be able to have them configured at the same time. Loop Guard applies to Alternate and Root ports, and it will move them into LOOP_Inconsistent state if they suddenly stop receiving BPDUs. BPDU Guard, on the other hand, will err-disable a port as soon as it receives a BPDU. In essence, while Loop Guard expects BPDUs to continuously arrive on a port, BPDU Guard will shoot down the port as soon as a BPDU arrives - so these two features seek for complete opposites. Guillermo was very correct in stating that the BPDU Guard is intended to protect edge ports against connecting new unauthorized switches or creating loops. BPDU Guard has no place on internal ports between switches, and there is a very good reason for that: Your network may suffer such an outage of existing links and switches that in order to keep your network connected, it might be required to use exactly the link you have protected with the BPDU Guard - in which case, however, this link will likely go down, as it will be exchanging BPDUs.
Guillermo: You are writing: "If you are thinking of enabling it because you want to stop another switch from becoming the root bridge then maybe root guard is what you are looking for. It is actually good practice to have your root bridge configured with root guard and loop guard"
With all due respect, I have a different opinion. You cannot really stop another switch from becoming a root bridge in STP, not even with Root Guard. You can only isolate it so that it does not affect the rest of the network, and this is what the Root Guard does. It makes little sense, however, to configure the Root Guard on your intended root switch. If another switch tried to become the root, your current root switch running Root Guard would block its would-be Root port(s) and move them into ROOT_Inconsistent state, with the net result being that your root switch would isolate itself completely from the network, while the rest of the network would converge to the new root switch.
In addition, having Root Guard configured on internal ports of your network makes little sense, too: Your current root switch may fail, or the intended links may fail, and in order to keep your network connected and running, other switches or other links may need to become root switches and root links. If you had them configured with the Root Guard, you would effectively prevent them from being able to be used as alternative, redundant, backup paths.
To me, the Root Guard makes most sense in scenarios when your network needs to interoperate its STP with a different network where you are not the administrator (say, a customer's network) and you need to make sure that the customer does not attempt to hijack your root switch. The Root Guard here is well deserved: It will prevent you from any attempts of the customer to steal the role of the root switch, and will also prevent your network from healing itself through the customer's network in an unlikely case of partitioning.
My two cents...
Best regards,
Peter

---
id: collect-260926-mikrotik/mikrotik/questions-1640115-fiber-intel-x520-sfp-with-mikrotik-crs305-1g-4sin-sfp-troubles-a66b628d
title: "questions-1640115-fiber-intel-x520-sfp-with-mikrotik-crs305-1g-4sin-sfp-troubles-a66b628d"
domain: mikrotik
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["intel"]
source: docs/RAG/lot-mikrotik/forum/misc/questions-1640115-fiber-intel-x520-sfp-with-mikrotik-crs305-1g-4sin-sfp-troubles-a66b628d.md
source_anchor: ""
source_lines: [1, 9]
sha256: a51b4957598669cca03aefddc5880fbc259789e087f4cf00819cb0c275000d4b
---

# questions-1640115-fiber-intel-x520-sfp-with-mikrotik-crs305-1g-4sin-sfp-troubles-a66b628d

- I have an Intel 82599 X520 (dual SFP+) network adapter in my motherboard.
- I have a HP Procurve 1410-24g switch with SFP ports.
- I have a Mikrotik CRS305-1G-4S+IN switch with SFP+ ports
- I have A7EL-SN85-ADMA transceivers
- I believe I have an IBM 12R9914 (OFNP, MMF, LC-LC, orange) fiber optic cable. It was given to me for free, so I do not know where it came from.
I am using Windows 10, 64 bit. In the device manager, I see the dual SFP+ ports as "Intel(R) 82599 10 Gigabit Dual Port Network Connection," so I believe my NIC is installed in my computer correctly. I tried plugging the transceiver into the x520 NIC and into one of the SFP ports on the HP switch, and I plugged the fiber optic cable into both. Nothing seemed to happen (the respective status lights did not light up). I tried flipping the fiber optic on one side as someone had suggested, but no luck. I am guessing that this is due to the HP switch port being SFP and not SFP+ (thus needing a SFP transceiver). Also, people online say that HP hardware is picky...
Somewhere online, I read that the Intel x520 NIC is agnostic when it comes to transceivers and motherboard interfaces, so my understanding is that there is no problem with the NIC. I also read that the Mikrotik switch is also agnostic when it comes to using other vendors' transceivers, so I bought one and decided to use that instead of the HP Procurve 1410-24g switch with SFP ports. Also, the Mikrotik switch has SFP+ ports, so I believe that eliminates the possibility of needing 1G SFP transceivers.
Even so, it still does not work. The respective lights on the NIC or Mikrotik switch did not light up. Are the transceivers or fiber optic cable the issue here? I am not sure if it is a compatibility issue, due to all the different companies and manufacturers involved. I do not have experience with fiber...
Side note: I have a different machine that has a x520 NIC, and it was pretty much plug-and-play using a Direct Attach Copper (DAC) cable with a Mikrotik SFP+ switch. However, I do not have access to that DAC cable, so I cannot try it with this equipment yet...

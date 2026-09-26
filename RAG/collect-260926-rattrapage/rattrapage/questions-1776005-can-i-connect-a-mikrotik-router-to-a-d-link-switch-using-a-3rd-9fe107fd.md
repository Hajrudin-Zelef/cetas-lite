---
id: collect-260926-rattrapage/rattrapage/questions-1776005-can-i-connect-a-mikrotik-router-to-a-d-link-switch-using-a-3rd-9fe107fd
title: "questions-1776005-can-i-connect-a-mikrotik-router-to-a-d-link-switch-using-a-3rd-9fe107fd"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-rattrapage/ai-llm/questions-1776005-can-i-connect-a-mikrotik-router-to-a-d-link-switch-using-a-3rd-9fe107fd.md
source_anchor: ""
source_lines: [1, 17]
sha256: f4ebdbb5b003965657d427e5354e124448bb72ae796b669fd5b26b683ac35f90
---

# questions-1776005-can-i-connect-a-mikrotik-router-to-a-d-link-switch-using-a-3rd-9fe107fd

I have a Mikrotik CCR1036 with RJ-45 downlink ports and a D-Link DGS-1210-28XS/ME with SFP uplinks. May I connect D-Link to Mikrotik via UTP cable and SFP with RJ-45 connector like GR-S1-RJ?
1 Answer 1
Maybe, maybe not.
There's no such thing really as up- or downlink port as a separate port type. The ports leading from edge devices towards core are called uplinks, and from core to the edge downlinks. The four slots on the D-Link aren't SFP, they're SFP+. SFP+ is (normally) 1/10GB, so in theory they are backwards compatible with 1GB SFP transceivers.
SFP slots are just slots. They're usually used for fiber connectivity, but that's no rule. As long as the xceiver itself is compatible, it doesn't matter whether the medium is fiber or copper.
The problem is that manufacturers can be picky on whose SFPs are compatible. If they produce xceivers themselves the guaranteed compatibility is usually only with their own products. And both Mikrotik and D-Link do produce their own.
To have a fully qualified answer you really should contact D-Link support - and Mikrotik, if you plan to use same xceiver on their device. Chances are the response will be "we only support our own", which means a 3rd party may or may not work. If it doesn't, then it doesn't, and that's that. There's no configuration that can force a device to work with an unsupported part.
Therefore I'd recommend sticking with the network device manufacturer's own parts, unless the manufacturer guarantees compatibility. Especially in a business environment.
- 
        At least Mikrotik explicitly state that they don't have manufacturer restrictions, but the CCR1036 already has RJ45 ports anyway – whereas the D-Link doesn't have any (all of the other 24 ports are 1G SFP), which is why I'm assuming OP has been looking at the GR-S1-RJ in the first place... (Wouldn't a SFP direct attach cable be more common in this situation, though?)grawity– grawity2023-03-28 07:28:23 +00:00Commented Mar 28, 2023 at 7:28
- 
            
            
- 
        Good to know, I've not dealt with Mikrotik before, just with the major manufacturers, so couldn't say what's their stand on the issue :-) Whether or not DAC is used depends on the installation. I think in a Data Center or company's network device room I've seem mostly DACs and SFPs used for fiber. But if the cable run's >10m... I don't think I've ever seen a +10m DAC.Peregrino69– Peregrino692023-03-28 07:39:51 +00:00Commented Mar 28, 2023 at 7:39
- 
        Why I'm warning is that MFG:s own SFPs tend to be pricey. I've seen companies spend $2000 on 3rd party products and ending up replacing them with MFG:s own products for 5 times the price as the 3rd party SFPs didn't work... and they can't even be returned as they're not faulty as such, just incompatible...Peregrino69– Peregrino692023-03-28 07:47:47 +00:00Commented Mar 28, 2023 at 7:47

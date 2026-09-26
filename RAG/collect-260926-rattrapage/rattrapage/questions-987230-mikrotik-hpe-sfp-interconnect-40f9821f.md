---
id: collect-260926-rattrapage/rattrapage/questions-987230-mikrotik-hpe-sfp-interconnect-40f9821f
title: "questions-987230-mikrotik-hpe-sfp-interconnect-40f9821f"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "parameters", "research", "wavelength"]
source: docs/RAG/lot-rattrapage/ai-llm/questions-987230-mikrotik-hpe-sfp-interconnect-40f9821f.md
source_anchor: ""
source_lines: [1, 14]
sha256: 81da520f625110919065c8b8b13f1faa0cc901696fa175d0b8a16e20fadb5c75
---

# questions-987230-mikrotik-hpe-sfp-interconnect-40f9821f

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
0
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
Can I connect a Mikrotik Router (hEX S, for instance) to HP/Aruba switch (J9776A, for instance) via SPF? What modules should I use?
I can easily find compatibility table of HPE SFP modules and switches, of Mikrotik SFP modules and switches, but nothing about cross-brand compatibility. Is cross-brand compatibility a thing?
Is it enough for wavelength and mode, like 850nm multi-mode, to be the same for compatibility or should I verify any other parameters?
You are recommended to buy Mikrotik's module to install into Mikrotik, and HPE's module to install into HPE. Mikrotik is usually able to use any modules, but it is not guaranteed; HPE in general would only work with HPE branded module.
What you have to be sure is that: your both modules from different vendors should have same speed rating (1Gbps or 100Mbps), same optical wavelength (1550 nm or 1310 or so), same fiber designation (single mode or multi mode; also same core size, 9/120, 50/125 and so on). You also have to find a compartible patch cord. Mikrotik modules usually have dual LC connectors, so one side of your patch cord will be two LCs. What is on other side depends on HPE's module, we have used HPE's modules with dual LC (but didn't connected them to Mikrotik via fiber, though).
Ideally it should work. 1000BASE-LX is a standard the same way 1000BASE-T is, except that it uses fiber as the transmission medium.Copper interfaces with 1000BASE-T are interoperable since they are defined by the same 1000BASE-T standards. You can expect the same with fiber 1000BASE-LX interfaces – which are all defined by 1000BASE-LX. SFP is a multi-vendor standard that is specified by a multi-source agreement (MSA).
You can expect your fiber patch cable to lead a consistent link as long as you use modules of the same Ethernet protocol, cable type and working wavelength.
As far as the module goes, you need to check with your vendor to see if you can use a generic SFP module in that switch, or if you need a vendor-specific module. Mikrotik can usually use generic modules, while HP needs a vendor-specific (or compatible) module.

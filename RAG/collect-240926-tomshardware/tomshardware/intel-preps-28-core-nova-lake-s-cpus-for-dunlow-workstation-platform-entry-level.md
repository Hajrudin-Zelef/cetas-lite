---
id: collect-240926-tomshardware/tomshardware/intel-preps-28-core-nova-lake-s-cpus-for-dunlow-workstation-platform-entry-level
title: "intel-preps-28-core-nova-lake-s-cpus-for-dunlow-workstation-platform-entry-level"
domain: tomshardware
role: reference
task: reference
actors: ["AMD", "Google", "Intel"]
dates: []
keywords: ["intel", "amd", "compute", "energy", "memory", "packaging", "wafer"]
source: docs/RAG/clean_en/tomshardware/intel-preps-28-core-nova-lake-s-cpus-for-dunlow-workstation-platform-entry-level.md
source_anchor: ""
source_lines: [1, 98]
sha256: 880bf9108402fe84d38a82df1971c53078954a05cf9487f42c9edeeaada522f6
---

# intel-preps-28-core-nova-lake-s-cpus-for-dunlow-workstation-platform-entry-level

<!-- source: https://www.tomshardware.com/pc-components/cpus/intel-preps-28-core-nova-lake-s-cpus-for-dunlow-workstation-platform-entry-level-xeon-chip-features-lga1954-socket -->

Intel is working on a version of its Nova Lake-S processor platform codenamed Dunlow that will offer up to 28 cores and will target entry-level server and workstation applications, according to shipment manifests located in the NBD database by @x86deadandback.

Formally, Intel's codenamed Dunlow platform will succeed the company's Catlow platform with Xeon 6300P-series CPUs and will support Xeon E-class Nova Lake-S processors (presumably) with up to 28 cores that feature a dual-channel memory subsystem, come in an LGA1954 form-factor, and have a processor base power of 95W, according to shipments manifests at NBD data.

Intel's next-generation Core Ultra 400-series platforms for desktop computers, codenamed Nova Lake-S, allegedly feature up to 52 cores, which include up to 16 high-performance Coyote Cove cores and up to 32 energy-efficient Arctic Wolf cores in the compute tile, as well as four low-power Arctic Wolf cores presumably in the SoC tile. These Nova Lake-S CPUs are aimed at enthusiasts and reportedly pull up to 474W with a single purpose: to offer unbeatable performance and feature set to put Intel back on the map of enthusiast-grade platforms currently dominated by AMD.

By contrast, the Dunlow platform seems to be a completely different kind of animal. The CPU deliberately features 28 cores and up to 95W PBP (TDP). All Xeon processors except Xeon 6700E, Xeon 6+, and some Atom-based solutions for specialty applications released to date have only featured high-performance cores. Even Intel's Xeon 6300P-series 'Raptor Lake-E' based products feature up to 12 P-cores to offer higher sustained all-core frequencies. Therefore, unless Intel plans to offer energy-efficient cores in its next Xeon CPU aimed at entry-level servers and workstations, we may be dealing with a very special processor that features 28 P-cores that is designed to beat all desktop-grade platforms in demanding applications.

While, for now, 28 P-cores inside Nova Lake-S processors for the Dunlow platform is speculation, it should also be noted that 28 cores do not naturally derive from a 16P+32E desktop design and are impossible to derive from a notebook-grade 8P+16E design. Also, Intel typically does not create server/workstation products by fusing off nearly half a desktop die (it does not even matter whether it disables some P-cores and some E-cores, disabling 20 cores in a 48-core tile hardly makes a lot of sense).

A Nova Lake-S CPU for Dunlow featuring a compute tile with 28 P-cores would resemble the abandoned Raptor Lake-32C, which featured an all-P-core design aimed at workstations and entry servers before being canceled. It is also possible that this could be a derivative of a small Xeon die adapted to an LGA1954 packaging and dual-channel memory to reduce platform costs. At the end of the day, many server applications like storage or web hosting do not need extremely high memory bandwidth, so two DDR5 channels could be enough.

Another reason for Intel to release a Nova Lake-S CPU with up to 28 P-cores is to fill the gap between high-end enthusiast-grade desktops that feature up to 16 P-cores and expensive Xeon 6 server and workstation CPUs that may start at 16 cores, but feature an octa-channel memory subsystem that is costly and is an overkill for many applications. Also note that since Xeon 'Diamond Rapids' processors with an octa-channel memory subsystem have been canceled, the gap between desktop and high-end server CPUs just gets way too wide in 2028, making Nova Lake 28 P-core silicon a potentially viable option.

Get Tom's Hardware's best news and in-depth reviews, straight to your inbox.

*Follow* *Tom's Hardware on Google News**, or* *add us as a preferred source**, to get our latest news, analysis, & reviews in your feeds.*

**Note:** the embedded chart doesn't have a "zoom" button. If you click this one, you can zoom it to full size:
 
 
 Don't go creating silly rumors!The article said:Therefore, unless Intel plans to offer energy-efficient cores in its next Xeon CPU aimed at entry-level servers and workstations, we may be dealing with a very special processor that features 28 P-cores that is designed to beat all desktop-grade platforms in demanding applications.
 
 99% likelihood this is just a Nova Lake with single compute tile.*Especially* at 95W!!!
 
 
 How did you not get the memo that Nova Lake's die configuration is one of:The article said:it should also be noted that 28 cores do not naturally derive from a 16P+32E desktop design and are impossible to derive from a notebook-grade 8P+16E design.
 8P + 16E + 4LPE
 2 * (8P + 16E) + 4LPE???
 
 The first*clearly* shows how they arrive at 28 cores.
 
 **Source:** https://www.tomshardware.com/pc-components/cpus/intel-nova-lake-specs-leaked-up-to-52-cores-and-150w-of-tdp-for-intels-amd-zen-6-rival
 
 But, that wouldThe article said:the abandoned Raptor Lake-32C, which featured an all-P-core design aimed at workstations and entry servers before being canceled.*never* have gone in a desktop socket!! It must've been planned to use the same socket as Xeon-W2400.
 
 Either that, or it was a 0P + 32E chip.*That* I could easily see slotting into LGA1700 and would make a lot of sense for light-duty servers.
- 
It seemsReply*very* odd that Intel would be bringing hybrid design into Xeon unless they're rebranding the edge processors. The other super cynical guess would be dropping ECC support from desktop parts, but that doesn't seem particularly likely.
 
 Ibit_user said:But, that would*never* have gone in a desktop socket!! It must've been planned to use the same socket as Xeon-W2400.*think* the article was referring to this which I don't recall ever seeing any real clarification on: https://www.tomshardware.com/news/intels-unannounced-34-core-raptor-lake-cpus-displayed-on-wafer
- 
Reply
 Didn't Anton Shilov write this article? :Pbit_user said:**Note:** the embedded chart doesn't have a "zoom" button. If you click this one, you can zoom it to full size:
 
 
 Don't go creating silly rumors!
 
 99% likelihood this is just a Nova Lake with single compute tile.*Especially* at 95W!!!
 
 
 How did you not get the memo that Nova Lake's die configuration is one of:
 8P + 16E + 4LPE
 2 * (8P + 16E) + 4LPE???
 
 The first*clearly* shows how they arrive at 28 cores.
 
 **Source:** https://www.tomshardware.com/pc-components/cpus/intel-nova-lake-specs-leaked-up-to-52-cores-and-150w-of-tdp-for-intels-amd-zen-6-rival
 
 But, that would*never* have gone in a desktop socket!! It must've been planned to use the same socket as Xeon-W2400.
 
 Either that, or it was a 0P + 32E chip.*That* I could easily see slotting into LGA1700 and would make a lot of sense for light-duty servers.
- 
Reply
 In that article, it says:thestryker said:I*think* the article was referring to this which I don't recall ever seeing any real clarification on: https://www.tomshardware.com/news/intels-unannounced-34-core-raptor-lake-cpus-displayed-on-wafer
 "The die appears larger than the die that drops into the standard desktop PCs with the LGA 1700 socket, so it appears to be too large to fit inside the package for desktop PCs. That means this is likely a CPU destined for the workstation market. "
 So, it was some sort of HEDT or Xeon-W die, most likely. That's why I said it would never slot into LGA1700 and must've been destined for LGA4677.
 
This article is specifically talking about Dunlow being LGA1954. So, there's no way it's 28 P-cores. Not in 95W. That's just crazy talk.
- 
Reply
 I think Anton doesn't track desktop PC products as well as he tracks silicon engineering. Like everyone, he has his niches and areas of interest.Gururu said:Didn't Anton Shilov write this article? :P
 
The fact that he didn't know about the single compute tile Nova Lake die configuration is understandable, from that perspective, but still something that should've been caught by an editor.
- 
Reply
It's a wafer full of EMR tiles, but at the time I don't believe it had been announced. That was just the only thing I could think of which ever carried the "RPL-S" title which had 3x cores.bit_user said:So, it was some sort of HEDT or Xeon-W die, most likely. That's why I said it would never slot into LGA1700 and must've been destined for LGA4677.
- 
A: It is 8P + 16E + 4LPZReply
 B: No it is 3P + 4B + 12E + 4LPE+14XP
 A: At the maximum it is 8P + 12E + 4LPE + 4Z
 B: As to the maximum then it is 3P + 4B + 12E + 4LPE+14XP+3P + 4B + 12E + 4LPE+14XP
 A: This is what I spent 3 years on and uncovered that in my PhD thesis. And do your math fool.
B: And who are you then?
- 
Reply
 C: Don't quit your day job!Stomx said:A: It is 8P + 16E + 4LPZ
 B: No it is 3P + 4B + 12E + 4LPE+14XP
 A: At the maximum it is 8P + 12E + 4LPE + 4Z
 B: As to the maximum then it is 3P + 4B + 12E + 4LPE+14XP+3P + 4B + 12E + 4LPE+14XP
 A: This is what I spent 3 years on and uncovered that in my PhD thesis. And do your math fool.
B: And who are you then?
- 
Reply
 I have no opinion either way but why is 95W such an issue for you?!bit_user said:This article is specifically talking about Dunlow being LGA1954. So, there's no way it's 28 P-cores. Not in 95W. That's just crazy talk.
It's entry level so it could be the equivalent of the -T desktop models, significantly lowered PBP to run cooler and be much more efficient (and you can still give them the same amount of power as the normal ones to get the same speed and inefficiency) .

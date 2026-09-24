---
id: collect-240926-tomshardware/tomshardware/intel-s-next-gen-52-core-nova-lake-cpu-could-pull-up-to-474w-high-end-lga1954-mo
title: "intel-s-next-gen-52-core-nova-lake-cpu-could-pull-up-to-474w-high-end-lga1954-mo"
domain: tomshardware
role: reference
task: reference
actors: ["AMD", "Google", "Intel"]
dates: []
keywords: ["intel", "amd", "compute", "gpu", "gpus", "memory"]
source: docs/RAG/clean_en/tomshardware/intel-s-next-gen-52-core-nova-lake-cpu-could-pull-up-to-474w-high-end-lga1954-mo.md
source_anchor: ""
source_lines: [1, 72]
sha256: b901a473cb4ab285f222b46dbb64e5a3d41e6298a41d1b75409f7169a2054165
---

# intel-s-next-gen-52-core-nova-lake-cpu-could-pull-up-to-474w-high-end-lga1954-mo

<!-- source: https://www.tomshardware.com/pc-components/cpus/intels-next-gen-52-core-nova-lake-cpu-could-pull-up-to-474w-high-end-lga1954-motherboards-may-need-three-8-pin-power-connectors-to-feed-the-monster -->

Intel is expected to push the boundaries on power draw with its upcoming Nova Lake series processors, which will rival the best CPUs. According to newly leaked information, the flagship 52-core desktop variant is expected to feature a dual-compute tile architecture with a massive PL2 limit of 474W. The information was shared by LC Tech Leaks and confirmed by Jaykihn, who has a pretty solid track record with Intel hardware.

PL2, or Power Limit 2, represents the maximum power a CPU can draw during short boost periods. That said, a PL2 target of 474W remains quite demanding, although a previous rumor suggests Intel may also have a PL4 emergency power limit over 700W. It is important to note that these power limits may only apply to the top-end models with the dual-tile architecture.

Additionally, the leak also sheds light on the upcoming platform, including the previously rumored LGA1954 socket. We already know that Nova Lake-S will require a new generation of motherboards. Motherboard vendors are expected to classify their boards by sustained PL1 power levels, with configurations for 35W, 65W, 125W, and 175W CPUs. Enthusiast-grade motherboards, likely the Z990 series, are also rumored to feature three EPS 8-pin CPU power connectors instead of the traditional two. While vendors will have the option to include a third connector, its primary purpose would be to support extreme overclocking and would not affect the CPU's rated performance profile.

The upcoming Nova Lake-S lineup is expected to carry the ‘Core Ultra 400S’ moniker and will be Intel's biggest desktop CPU overhaul in years. We’ve previously reported leaked specifications indicating configurations ranging from 6 to 52 cores, with support for DDR5-8000 memory. The flagship 52-core model is expected to feature 16 performance cores, 32 efficiency cores, and a new Big Last Level Cache (bLLC) design to take on AMD's 3D V-Cache gaming dominance. The company is also rumored to introduce integrated Xe3 graphics, Thunderbolt 5, PCIe 5.0 connectivity, and an upgraded NPU for AI workloads.

While these specifications are unconfirmed, it is clear that Intel is targeting substantial gains in gaming, multi-threaded performance, and overall platform capabilities with its next-gen processors.

*Follow* *Tom's Hardware on Google News**, or* *add us as a preferred source**, to get our latest news, analysis, & reviews in your feeds.*

Get Tom's Hardware's best news and in-depth reviews, straight to your inbox.

Kunal Khullar is a contributing writer at Tom’s Hardware. He is a long time technology journalist and reviewer specializing in PC components and peripherals, and welcomes any and every question around building a PC.

- 
even if it’s just PL2 that’s a lot of heat you’re dumping into a room. For gaming add a GPU that’s doing something similar in power consumption (or worse) and a monitor doing its bit and you’re effectively not gaming in summer without an AC going full blast.Reply
 
 This seems kind of ridiculous by itself, and I’m not even thinking about the kind of cooling you’ll need for this to work well.
 
I could kind of get behind it if you’re using it for work to make a living, but for gaming this seems insane.
- 
Reply
 This is why I don't usually go for the highest core-count CPUs, at home. I don't do enough compiling on my personal machines to justify the cooling system needed for a big CPU, nor do I want that kind of heat output.VizzieTheViz said:even if it’s just PL2 that’s a lot of heat you’re dumping into a room.
 
 
 Most games don't scale well with core count. They're starting to do more of that, but you can get a pretty good idea by just looking at CPU power consumption while gaming to see that a lot of the cores spend most of the time idle.VizzieTheViz said:For gaming add a GPU that’s doing something similar in power consumption (or worse)
 
 https://www.techpowerup.com/review/amd-ryzen-9-9950x3d2-dual-edition/22.html
 The average gaming power consumption on the X3D2 is 128W, while average app power consumption is 153W and MT power consumption is 249W.
 
 
 That's mainly driven by your GPU.VizzieTheViz said:for gaming this seems insane.
 
BTW, I doubt gaming monitors use more than 100W, which is still less than old CRT monitors did (or plasma TVs, if you had one of those).
- 
Reply
 We don't know that and the graph you are showing isn't showing that either, you would need an infinite GPU to see if current CPUs scale in games or not. (or run games at postage stamp size windows)bit_user said:Most games don't scale well with core count. They're starting to do more of that, but you can get a pretty good idea by just looking at CPU power consumption while gaming to see that a lot of the cores spend most of the time idle.
 No review ever tries to show GPU usage during these tests to show anything (not that it would since a GPU has so many parts that can bottleneck).
 "We don't have GPUs that allow CPUs to scale very much in games" would be the correcterer thing to say, since we also don't know that for sure.
 
 
 It's in this article:usertests said:IIRC the 474W limit is only going to be hit by overclocked dual-tile CPUs, not out of the box. I'll look for a source.
 " It is important to note that these power limits may only apply to the top-end models with the dual-tile architecture."
 
 
 It depends on if they keep the PL2=PL1="max wattage" credo, if they revert to what was before that then PL1 will be the power draw and it will only go to PL2 if the CPU budgeted enough wattage over the previous time span to boost up to PL2 without increasing the overall average.VizzieTheViz said:even if it’s just PL2 that’s a lot of heat you’re dumping into a room. For gaming add a GPU that’s doing something similar in power consumption (or worse) and a monitor doing its bit and you’re effectively not gaming in summer without an AC going full blast.
 
 This seems kind of ridiculous by itself, and I’m not even thinking about the kind of cooling you’ll need for this to work well.
 
 I could kind of get behind it if you’re using it for work to make a living, but for gaming this seems insane.
 
Also the whole article is "expected" "leaked" "may" and so on, nothing of that might come true.
- 
Reply
 Key detail being: overclocked. Not hitting that out of the box even with 48-52 stressed cores.TerryLaze said:It's in this article:
" It is important to note that these power limits may only apply to the top-end models with the dual-tile architecture."
- 
Reply
 Since these are all speculations...does it even matter?usertests said:Key detail being: overclocked. Not hitting that out of the box even with 48-52 stressed cores.
If it's like previous gens, unlocked PL will be the same as the max overclock anyway and PL2 will probably be the max that every single CPU will be able to reach.
- 
Unless you're running six VMs so you have a complete copy of your server farm on your workstation, this all seems silly. Four performance cores and four economy cores should take care of 99% of what anyone does on a workstation (much less a laptop). With modern SSDs on the PCI that's already one buttload of computing power.Reply
 Ditto graphics, really, OK game fanatics ... whatever. I dropped out of that after Pac-Man.
If I need more than about 60 watts on a workstation, just do it in the cloud.
- 
Intel+W11 have gotten pretty good at managing hybrid cores. This Dual compute chip is new and will probably be sloppy at scheduling at first, and keeping both energized while only needing one should hurt light use efficiency. And performance by a bit. The single chip version won't have these issues.Reply
 
Maybe I'll be surprised and they will have it all worked out by release.

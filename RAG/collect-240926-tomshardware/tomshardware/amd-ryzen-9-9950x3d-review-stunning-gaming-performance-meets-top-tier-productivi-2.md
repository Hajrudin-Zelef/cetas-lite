---
id: collect-240926-tomshardware/tomshardware/amd-ryzen-9-9950x3d-review-stunning-gaming-performance-meets-top-tier-productivi-2
title: "amd-ryzen-9-9950x3d-review-stunning-gaming-performance-meets-top-tier-productivi"
domain: tomshardware
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: []
keywords: ["amd", "benchmarks", "chiplet", "compute", "gpu", "intel", "pricing"]
source: docs/RAG/clean_en/tomshardware/amd-ryzen-9-9950x3d-review-stunning-gaming-performance-meets-top-tier-productivi.md
source_anchor: ""
source_lines: [71, 124]
sha256: 80ca575f325f524dd721da10df8dbdc27b31346a3eb3e47e055860892e141b06
---

# amd-ryzen-9-9950x3d-review-stunning-gaming-performance-meets-top-tier-productivi

The Ryzen 9 9900X3D sports 12 cores, 24 threads,140 MB of total cache, and a 5.5 GHz boost. This chip has the same 120/162W TDP as its predecessor and its non-X3D counterpart, the Ryzen 9 9900X, so we expect significantly less performance than the 9950X3D across the board. AMD hasn't sampled this chip but says it has dramatically reduced the performance gulf between the two Ryzen 9 X3D models. We included the company's 9900X3D benchmarks at the bottom of the following page.

## AMD's Updated Chipset Drivers

We've covered AMD's chipset drivers in-depth in the past. The drivers have a suite of components that enable its dual-chiplet processors to operate as close as possible to the single-chiplet comparables, despite having only a single L3 cache die under one of the compute chiplets.

AMD's first dual-chiplet X3D processors employed a new version of thread targeting that works in tandem with putting unneeded cores to sleep, thus forcing game code to run on the single chiplet that houses the performance-boosting L3 cache. However, this implementation made an irreversible change to the operating system that could hamper performance if the chip were later swapped out for a single-CCD processor, with the only remedy being a complete reinstallation of the operating system.

As you can see in the slides above, AMD has now fixed that issue with an updated Provisioning Packages Service (the engine that manages core parking and thread targeting). After roughly 15 minutes of idle time, this driver now automatically detects when a new processor has been installed in the system and adjusts the provisioning accordingly, so there's no need for an operating system reinstall. Things are essentially plug-and-play now, as they should be. AMD also addressed a shortcoming with its 3D V-Cache Performance Optimizer, so it now works when Virtulization-Based Security (VBS) is enabled on Windows 10.

Despite multiple components working in concert to ensure that games run smoothly on the dual-CCD models, some game titles remain problematic. To fix this, AMD revived its Application Compatibility Database (ACD), a technology that debuted with the first Threadripper CPUs. The ACD is a list-based feature that detects when certain games are launched (listed in the image above). It then triggers a mechanism that reduces the number of threads, thus hiding them from the operating system and fully preventing the game code from running on the unoptimized chiplet.

Internally, AMD has affectionately nicknamed this 'Core Lie' because the feature lies with the operating system about the number of cores available. This mechanism assures optimal performance with several of the more stubborn titles, helping to once again reduce the difference between the single- and dual-chiplet X3D models.

Let's see what all of this looks like in our gaming benchmarks on the next page.

- 
ReplyCons: pricing The Ryzen 9 9950X3D scores another walk-in touchdown for AMD, easily earning its $699 price tag. Come on, give it a 5 already! There are no cons, it's not a $1000 CPU, it has the perfect price for the performance and features. There won't be another 5 star-worthy processor for a long time.
 
 (not being a fanboy, I'm just challenging the 5-stars-but-not-quite rating system)
 
https://i.imgflip.com/9mzje0.jpg
- 
I don't think I'll have to upgrade from my 7800X3D for a while, and I believe others with similar chips will concur. The 9950X3D is compelling, but not enough. I got my chip for $266(from a bundle), 80% of the performance of a 9950X3D for ~30% of the price was a damn good deal. If I ever do upgrade, it'll be far in the future when I can also afford faster RAM and a new Mobo, perhaps even a new GPU, but that's just me dreaming :smile:Reply
- 
Reply
 intel made 96.8bil in the last 10 years....oofdragon said:RIP Intel 10th consecutive year
 that's net income after all expenses including dividends and after all the losses they had.
 intel is RIP-ing all the way to bank, laughing.
 
 During the same last ten years amd made 8.5 bil.
 
10 years ago AMD was still on faildozer................................................................................................................................................
- 
Reply
 It's going to affect them long term. They are losing space in the gaming market as they haven't released anything competitive as of recent, and they are losing space in the server market as AMD's Threadripper and Epyc lines exist. Intel is subsiding on locked-in server owners who are already on intel's platform, along with the sale of older(relative)CPUs in the gaming market. At least, that's what I think is happening. For all we know, Intel could just be cooking up the hardest comeback ever with all that money! I don't care who has the best, competition lowers prices and that's what we all need.TerryLaze said:intel made 96.8bil in the last 10 years....
 that's net income after all expenses including dividends and after all the losses they had.
 intel is RIP-ing all the way to bank, laughing.
 
 During the same last ten years amd made 8.5 bil.
 
10 years ago AMD was still on faildozer................................................................................................................................................
- 
Reply
I honestly belive that, if Intel were cooking a comeback, they would be doing it since first-gen Zen. All they managed to do was to join two bad cores together, and disable hyperthreading. I'm not holding my breath for them, although the future might indeed bring surprises.Crazyy8 said:It's going to affect them long term. They are losing space in the gaming market as they haven't released anything competitive as of recent, and they are losing space in the server market as AMD's Threadripper and Epyc lines exist. Intel is subsiding on locked-in server owners who are already on intel's platform, along with the sale of older(relative)CPUs in the gaming market. At least, that's what I think is happening. For all we know, Intel could just be cooking up the hardest comeback ever with all that money! I don't care who has the best, competition lowers prices and that's what we all need.
- 
Reply
 You don't need to make a comeback if you are never gone in the first place...salgado18 said:I honestly belive that, if Intel were cooking a comeback, they would be doing it since first-gen Zen. All they managed to do was to join two bad cores together, and disable hyperthreading. I'm not holding my breath for them, although the future might indeed bring surprises.
 Amd is still barely making any money,
 
and things are only going to get worse, they will have to pay for more advanced nodes and they will have to add more cores and maybe add x3d to more CPUs all of that is going to eat into amds earnings.
- 
Reply
In August I upgraded my old desktop from an i7-4770k to an R7 9700X. I kept the GPU during the upgrade, RX6700XT, and I have noticed higher FPS and consistently better lower FPS. Depending on what your current CPU is with something like the 5070Ti will matter quite a bit as to if you see huge gains or not.Gururu said:I hate to ask the question, but will these wins be evident on 5070 ti cards and below? I just don't see the majority of people who go for this chip also forking for a 5090.

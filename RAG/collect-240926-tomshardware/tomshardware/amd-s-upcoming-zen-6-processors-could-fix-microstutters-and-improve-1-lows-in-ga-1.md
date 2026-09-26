---
id: collect-240926-tomshardware/tomshardware/amd-s-upcoming-zen-6-processors-could-fix-microstutters-and-improve-1-lows-in-ga-1
title: "amd-s-upcoming-zen-6-processors-could-fix-microstutters-and-improve-1-lows-in-ga"
domain: tomshardware
role: reference
task: reference
actors: ["AMD", "Google"]
dates: []
keywords: ["amd", "energy", "gpu", "gpus", "memory", "power delivery"]
source: docs/RAG/clean_en/tomshardware/amd-s-upcoming-zen-6-processors-could-fix-microstutters-and-improve-1-lows-in-ga.md
source_anchor: ""
source_lines: [1, 66]
sha256: d4fc2b4fd26699e3c469c7bf6d45039f2fede9c945cb4526530023ccdac74fe8
---

# amd-s-upcoming-zen-6-processors-could-fix-microstutters-and-improve-1-lows-in-ga

<!-- source: https://www.tomshardware.com/pc-components/cpus/amds-upcoming-zen-6-processors-could-fix-microstutters-and-improve-1-percent-lows-in-games-next-gen-cpus-tipped-to-feature-per-core-optimizations-for-thermal-and-power-budgets -->

AMD is expected to unveil its next-gen Zen 6 platform at CES 2026 with Ryzen 10000 series processors. While we're looking forward to IPC and clock speed improvements, it seems like more subtle, under-the-hood changes could end up upgrading the gaming experience in a big way. According to a tip received by Videocardz, AMD is implementing various per-core optimizations to ensure foreground tasks get priority over background applications. Each new feature is supposed to more smartly manage the power and thermal budgets at the silicon's disposal, so let's go over each of them.

First up, we have CPPC Performance Priority — CPPC stands for Collaborative Processor Performance Control, and it's responsible for communication between the silicon and the OS. It lets the firmware sitting in between decide the performance of each core separately. This feature has actually existed since the Ryzen 3000 series, but it doesn't work perfectly. Zen 6 CPUs are apparently supposed to apply a band-aid fix and improve their effectiveness.

Secondly, we have *FloorPerf*, which acts as a dynamic, targeted power delivery system. It sets a minimum clock speed for all cores that the silicon can clock down to in case of thermal throttling. For instance, imagine you're running a game in the foreground while Discord and Spotify sit in the background. The temperatures on your Zen 6 CPU rise to the point of throttling, but instead of reducing the speed of all cores, FloorPerf will target the cores running the background tasks first, in order to maintain the performance of your game.

We covered a recent fix like this for the Linux world as well, where a sudden spike in the background can cause the foreground task to choke. Linux has an issue with priority allocation across tasks, so that was a different situation, but similar logic applies here. On a Zen 6 CPU, instead of your game suddenly experiencing microstutters because of aggressive throttling, *FloorPerf* will try to limit whatever's running in the background first. The aforementioned CPPC Performance Priority should work hand in hand with this feature to maintain clocks on the important cores.

CPPC will get another boost in the form of *HighestFreq*, which will allow the OS to access more granular chip data to make better core management decisions. As the name suggests, it pertains to maintaining high clock speeds on the cores running a foreground task. The OS will be aware of exactly which core can boost the highest and maintain that speed the longest. This will allow it to assign, for instance, a game's main rendering thread to the fastest cores, while pushing background apps to more power-efficient ones.

In conjunction with this, Zen 6 is also seemingly getting per-core EPP boost. EPP stands for Energy Performance Preference, and it's supposed to fix core parking issues caused by a momentary bottleneck. For example, if the CPU is waiting for the GPU to render the next frame, the clock speeds for the active cores will drop in that moment and struggle to ramp back up in time, forcing all cores to boost when the frame is ready. EPP boost is supposed to identify the active cores and individually keep them in high-performance mode to eliminate any wait times.

Lastly, Zen 6 is reportedly introducing PQOS Global Bandwidth Enforcement and an updated IBS Memory Profiler. Both of these target memory bandwidth with the same goal of maximum stability as everything else discussed so far. The former is responsible for allocating RAM to background tasks, ensuring their usage stays under the limit for foreground priority. The latter will limit L3 cache access to background tasks when it detects that those tasks are slowing down the foreground application, which should help with frametime consistency in games.

Get Tom's Hardware's best news and in-depth reviews, straight to your inbox.

As you can tell, all of these optimizations are supposed to work together and harmonize into a more efficient processor at the end — one that maximizes the silicon's potential as much as possible. Of course, none of this is confirmed and there's a lot that could be gated behind product segmentation. We could see some features only on high-end parts, while some are limited to mobile; there's no telling at the moment, but it's exciting stuff, nonetheless. Seems like a battle for the ages is brewing between Nova Lake and Zen 6 next year.

*Follow* *Tom's Hardware on Google News**, or* *add us as a preferred source**, to get our latest news, analysis, & reviews in your feeds.*

Hassam Nasir is a die-hard hardware enthusiast with years of experience as a tech editor and writer, focusing on detailed CPU comparisons and general hardware news. When he’s not working, you’ll find him bending tubes for his ever-evolving custom water-loop gaming rig or benchmarking the latest CPUs and GPUs just for fun.

- 
Are 1% lows always present despite the average? If one were to lock framerate to an observed 1%, would the 1% disappear or would a new lower 1% appear?Reply
 For example:
 Average: 150 fps
 1%: 88 fps
 >>change settings to lock framerate to 88
 Average: 88 fps
1%:???
- 
Reply
 Consider rendering 1000 frames. The 1% low is the 10th slowest frame to render. There are 9 frames as slow or slower.Gururu said:Are 1% lows always present despite the average? If one were to lock framerate to an observed 1%, would the 1% disappear or would a new lower 1% appear?
 For example:
 Average: 150 fps
 1%: 88 fps
 >>change settings to lock framerate to 88
 Average: 88 fps
1%:???
 
Adjust the calculation to use FPS and then choose how you want to measure it (what are you averaging across) and it'll still hold true
- 
Reply
 So based on that logic, the slowest frame would remain 88 fps, thus the average and the lowest would be identical if one were to lock the framerate to the lowest observed.moon2 said:Consider rendering 1000 frames. The 1% low is the 10th slowest frame to render. There are 9 frames as slow or slower.
 
Adjust the calculation to use FPS and then choose how you want to measure it (what are you averaging across) and it'll still hold true
- 
Reply
 No, that's wrong. The explanation is in the message you quoted though.Gururu said:So based on that logic, the slowest frame would remain 88 fps, thus the average and the lowest would be identical if one were to lock the framerate to the lowest observed.
 
Remember 1% of 1000 is 10. There are still 9 frames slower, or as slow, as the 10th slowest.
- 
Reply
 So then, just to answer the original question, would a new lower 1% appear then? Or is it inconclusive?moon2 said:No, that's wrong. The explanation is in the message you quoted though.
 
Remember 1% of 1000 is 10. There are still 9 frames slower, or as slow, as the 10th slowest.
- 
ReplyGururu said:So then, just to answer the original question, would a new lower 1% appear then? Or is it inconclusive?
 Under ideal circumstances each frame would have an equal time slice, perfectly evenly divided over 1 second and that would provide you the FPS number which in that case would never see any up or downs.
 
 Practically slowdowns happen and the FPS is not a stable rate but fluctuates up and down.
 
 1% lows in games is the measurement chosen to display the effects of fluctuating FPS over time, and show the 1% of the lowest frame rates as a constant measurement over time.
 

---
id: collect-240926-tomshardware/tomshardware/intel-core-ultra-7-270k-plus-review-back-from-the-brink-2
title: "intel-core-ultra-7-270k-plus-review-back-from-the-brink"
domain: tomshardware
role: reference
task: reference
actors: ["AMD", "Intel", "Microsoft"]
dates: []
keywords: ["intel", "amd", "benchmark", "benchmarks", "consumer"]
source: docs/RAG/clean_en/tomshardware/intel-core-ultra-7-270k-plus-review-back-from-the-brink.md
source_anchor: ""
source_lines: [60, 109]
sha256: 808827643b6386a9953eee85f23b6d79910ef28354c278243594346375031aba
---

# intel-core-ultra-7-270k-plus-review-back-from-the-brink

Intel describes iBOT as translating “other x86” to “Intel x86.” You can think of it as a translation layer along the lines of something like Microsoft Prism, but we’re not moving from one ISA to another. Instead, Intel is optimizing instructions to better leverage a particular architecture. It’s able to do this using Hardware Profile Guided Optimization, or HWPGO. Within Arrow Lake Refresh chips — and Intel chips moving forward — there are registers to show what is happening when code is executing on the chip. That includes things like cache misses, branch mispredictions, and hardware interrupts.

When a developer is compiling their binary, there’s a toolchain of optimization that takes place where they look at these types of inefficiencies. Then, they can go back to the source code, make adjustments as necessary, and recompile. With iBOT, Intel is trying to eliminate those inefficiencies, but it’s doing so on a production binary. It doesn’t need to touch any source code. That’s because these “hooks,” as Intel calls them, work on shipping binaries. It’s able to see inefficiencies and make adjustments, but it does so at runtime on a production binary, not through source code.

Let’s use a cache miss as an example. Intel can see a cache miss happen, and it can investigate what went wrong. For instance, maybe a piece of data wasn’t tagged properly and was flushed from the cache. You’d have to go get that data again, and your performance would go down. iBOT allows Intel to tag that data properly so it doesn’t get ejected from the cache. Add up these small efficiency improvements, and you could squeeze out some extra performance. And, as Intel describes it, this would effectively increase IPC. Cache misses and branch mispredictions represent instructions that weren’t fully executed within a cycle, so fixing those issues makes IPC go up.

This post-ship optimization presents a lot of opportunities. I’ll tell you now that, in the handful of games iBOT is releasing with, you’re looking at somewhere in the high single-digits for an uplift. It’s not massive, but it’s an early demonstration that this concept has legs. Developers use different compilers and different toolchains, and those have evolved and will continue to evolve over time. iBOT allows Intel to take out at least some of the inefficiencies in those toolchains. It could apply to an older application running on a new architecture just as easily as it could to a newer application running on an older architecture.

iBOT is an opt-in feature; Intel tells me that it’s being cautious about rolling out the feature, trying to avoid claims that it’s playing dirty tricks to win favor in benchmarks. That isn’t the case, short of Geekbench, where Intel has a proof of concept for how iBOT can work outside of games. I’ll address that when we reach Geekbench in our productivity benchmarks.

Intel is modifying code running in real-time, and it’s been clear that multiplayer games aren’t initially supported in iBOT because of that fact. If there are broader security implications remains to be seen, but it’s something to keep in mind.

If there are security risks, they shouldn't reach deep. Intel says iBOT operates at the same level as user-mode applications. It doesn't have direct access to the hardware, and it's making system calls like any application would.

- **MORE:** **Best CPU for gaming**
- **MORE:** **CPU Benchmark Hierarchy**
- **MORE:** **Intel vs AMD**
- **MORE:** **How to Overclock a CPU**

Jake Roach is the Senior CPU Analyst at Tom’s Hardware, writing reviews, news, and features about the latest consumer and workstation processors.

- 
I think it's a little controversial to include iBOT in a hardware review, unless you at least test with it both on & off, so see how much it's contributing.Reply
 
 I'm not really surprised to see something like this come along. I figured we'd have it by now, but I thought it'd be accompanied by hardware changes that*required* it. Based on my understanding, it's not really different than what JIT-based emulators are doing, for instance like when you run x86 code on ARM CPUs. In this case, it just so happens to be doing x86 -> x86.
 
 I think the*real* story behind iBOT is APX and AVX10! Right now, we're just seeing a*hint* of what it will do for Nova Lake!
- 
The die-to-die frequency increase has helped it to perform much better than a typical refresh, although it clearly tanks efficiency and idle power consumption badly.Reply
 
 Combined with the price, while it's not magic, it's the best possible outcome for Arrow Lake.
 
It will be interesting to see if the other reviews are so generous with iBOT.
- 
Reviews across the board are painting it as an absolute best for value. HU complained about the temps but the major tiff everyone had was of course the platform, being DDR5 and dead end. If I waited this long to upgrade from a DDR4, I'd just wait for Nova or Zen 6. If I have an 1851 already, the performance bump doesn't warrant more spending. If I was buying for family or significant others who don't upgrade period, this is a no-brainer.Reply
- 
I know builder and tycoon games aren't popular, but if you really want to test out CPU performance, load an end game save from Factorio, Timberborn, Cities Skylines 2, or Transport Fever 2.Reply
 
 The path finding calculation will bring a 9850X3D to its knees, and you'll get to see the true value of an X3D processor.
 
Also, where is the i5 250K review?
- 
Reply
This is valid. The test configuration is a 5090FE at 1920x1080, where the CPU becomes the limiting factor. Otherwise if you were to put a 5090FE on pretty much any modern CPU you would get equally high frame rates even at higher resolutions just not as many as you would on a 9800X3D, 9850X3D or even 9700 where the 270 wouldn't keep up.TerryLaze said:114 minimum FPS on average in a suite of 17 games...."Struggles"
- 
Reply
 No Ryzen that isn't an X3D tested here can keep up with the 270K+ in games, the more expensive and vastly slower in everything else 9700X included.warezme said:This is valid. The test configuration is a 5090FE at 1920x1080, where the CPU becomes the limiting factor. Otherwise if you were to put a 5090FE on pretty much any modern CPU you would get equally high frame rates even at higher resolutions just not as many as you would on a 9800X3D, 9850X3D or even 9700 where the 270 wouldn't keep up.
 https://cdn.mos.cms.futurecdn.net/WVvc7x7yHrYgJKwdTp7WEo-1200-80.png.webpTom's should just add that the non X3D Ryzens are all a worse choice for gaming when they mention that better gaming chips are lackluster. That would be a way to seem unbiased.
 
Edit: You could call the 270K+ the $300 9950X.

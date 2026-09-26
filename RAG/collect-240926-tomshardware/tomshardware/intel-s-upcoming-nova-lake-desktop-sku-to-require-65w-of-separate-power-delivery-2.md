---
id: collect-240926-tomshardware/tomshardware/intel-s-upcoming-nova-lake-desktop-sku-to-require-65w-of-separate-power-delivery-2
title: "intel-s-upcoming-nova-lake-desktop-sku-to-require-65w-of-separate-power-delivery"
domain: tomshardware
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["gpu", "nvidia"]
source: docs/RAG/clean_en/tomshardware/intel-s-upcoming-nova-lake-desktop-sku-to-require-65w-of-separate-power-delivery.md
source_anchor: ""
source_lines: [57, 89]
sha256: 93cead968a191ad57d0a99b72214dc64acf76f62e2fa8f02bdbb51d784436643
---

# intel-s-upcoming-nova-lake-desktop-sku-to-require-65w-of-separate-power-delivery

They have the expertise, they have the technology. We'll see if their new desktop APUs survive in the face of Nvidia deal, or if they simply use Nvidia tiles later on.
- 
My question is: how far can this iGPU help increase the performance of a discrete GPU?Reply
 I'd always opt for a processor without an iGPU if possible, so it can run at lower power and/or heat. But recently, an article in Tom's throws a possibility to add a second GPU, a lesser one than your main GPU, to increase performance by handling other tasks (such as AI, physics, etc).
I wonder if such a case is possible with an iGPU, and if it is worth it ?
- 
Reply
 It is possible to devote certain applications to an iGPU and others to the dedicated GPU in Win11 settings. I have tried this and it was more trouble for me than what it was worth. I don't know if two dGPUs have a better outcome.samopa said:My question is: how far can this iGPU help increase the performance of a discrete GPU?
 I'd always opt for a processor without an iGPU if possible, so it can run at lower power and/or heat. But recently, an article in Tom's throws a possibility to add a second GPU, a lesser one than your main GPU, to increase performance by handling other tasks (such as AI, physics, etc).
 I wonder if such a case is possible with an iGPU, and if it is worth it ?
Though it is always advisable to get the CPU with integrated GPU for troubleshooting purposes, a very powerful iGPU is not needed. The iGPUs here are very powerful and would mostly be useful for casual gamers allowing one to save a few hundred bucks by not needing a low-mid tier dGPU. It will be interesting to see if this iGPU setup can be maximized for tasking like you suggest in combination with a dGPU.
- 
Reply
 Since something like that would need the game devs to do work I wouldn't hold my breath for it happening.samopa said:My question is: how far can this iGPU help increase the performance of a discrete GPU?
 I'd always opt for a processor without an iGPU if possible, so it can run at lower power and/or heat. But recently, an article in Tom's throws a possibility to add a second GPU, a lesser one than your main GPU, to increase performance by handling other tasks (such as AI, physics, etc).
 I wonder if such a case is possible with an iGPU, and if it is worth it ?
In theory yes the igpu could handle upscaling/frame generation or general filters but the game would have to be aware of the igpu and actually give these tasks to it.
- 
Reply
 I don't think game devs necessarily need to do any work, if the drivers can handle some of these dual-GPU applications that don't actually involve a game rendering on more than one GPU.TerryLaze said:Since something like that would need the game devs to do work I wouldn't hold my breath for it happening.
 In theory yes the igpu could handle upscaling/frame generation or general filters but the game would have to be aware of the igpu and actually give these tasks to it.
 
 Here are some I've seen:
 
 1. One GPU handles game, the other GPU handles Lossless Scaling.
 2. One GPU handles game, the other GPU handles PhysX. This became necessary for a handful of games after Nvidia deprecated 32-bit PhysX/CUDA support on 50-series, until they reversed it.
 3. When Nvidia first showed off DLSS 5, it was running on two 5090s.
 
 GDvfIbRIb3U
 ClmKqKrrJGA
 h4w_aObRzCc
 xfv6g7DIYf0
Whether any of this is a good or practical is up for debate, but dramatically faster iGPUs could open up more options, while working with boards with only one x16 slot.

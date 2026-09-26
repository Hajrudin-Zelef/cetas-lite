---
id: collect-240926-tomshardware/tomshardware/amd-confirms-low-power-cpu-cores-in-linux-kernel-patch-zen-6-chips-could-follow-2
title: "amd-confirms-low-power-cpu-cores-in-linux-kernel-patch-zen-6-chips-could-follow-"
domain: tomshardware
role: reference
task: reference
actors: ["AMD", "Intel", "TSMC"]
dates: []
keywords: ["amd", "compute", "intel"]
source: docs/RAG/clean_en/tomshardware/amd-confirms-low-power-cpu-cores-in-linux-kernel-patch-zen-6-chips-could-follow-.md
source_anchor: ""
source_lines: [62, 90]
sha256: ec16d9cf916d08596b6f964b5e352ff05be01edae6a840625f2ef102f51acb4b
---

# amd-confirms-low-power-cpu-cores-in-linux-kernel-patch-zen-6-chips-could-follow-

If they don't make it to desktop, various APUs will have them.
- 
Reply
 Yeah so they would have to update the I/O tile from n6 to n3, the whole thing, with tsmc having done how many price increases now...usertests said:If Zen 6 includes them they should use whatever node the I/O uses. And I don't think it would be possible for AMD to make Zen 6 LP cores on TSMC N6.
 
I think the node would have to be TSMC N3C or N3P, which is hardly ancient
- 
Reply
 Zen 6 CPUs could end up being quite expensive indeed, with the core count increases used to justify it. But AMD might keep the I/O on N6 to save money.TerryLaze said:Yeah so they would have to update the I/O tile from n6 to n3, the whole thing, with tsmc having done how many price increases now...
 
Either way, it should be different than the one used with Zen 4/5, in order to add things like CUDIMM support, and possibly an NPU.
- 
The concept of having e-cores or low-power cores is to stop interrupting the performance cores! So ironically having a few e-cores should give you better performance on the p-cores. Since it is still an additional core the power saving per unit time should be pretty small.Reply
- 
Reply
 Your operating system might have it by the time these come out, but you might have to make that mode yourself. I'm also not sure how well this would work on AMD vs Intel chips.usertests said:I've argued about the Zen 6 LP cores here back when they were a leak. Desktops should get them, and they can lower idle/low-intensity power consumption and save everybody (small amounts of) money. Governments and businesses will appreciate lowering power across millions of x86 machines.
 
 The E-cores/C-cores are there to maximize multi-threading performance per die area for the most part. LPE-cores and LP cores will be truly efficiency focused, especially if they can turn off the compute die(s) completely when they aren't needed.
 
 Having more of them could prevent tasks from spilling over to compute die(s). Every Nova Lake-S SKU will have 4c/4t of these. I think they should move up to 8c/8t in a future generation.
 
 **Your operating system could have a super power saving mode to manually disable CCDs and run everything on LPs. For example, if you're trying to run a mini PC off a battery pack, are working in a hot environment, etc.**
 Zen 6 mobile APUs will definitely get LP cores. Zen 6 Olympic Ridge desktop CPUs are still uncertain, and we've recently heard they may not include an iGPU, but will include a big NPU. If they do include Zen 6 LP cores, it will be 2c/4t. If they make the final cut, I suspect they may include a tiny amount of dedicated L3 cache, e.g. 4 MiB.
 There's a program some guy made and posted on Guru3d that gives you a GUI to enable a bunch of hidden power options. Called PowerSettingsExplorer : https://forums.guru3d.com/threads/windows-power-plan-settings-explorer-utility.416058/ You can enable this stuff typing long chains of text as well but that is a hassle.
 
 Right now Windows power plans already have controls for 3 types of cores and with my Raptor Lake CPUs I know at least 2 are independently functional. I can have the pc use P cores, E cores or both and independently limit the speeds of those core types using only Windows power plan settings. I can probably do a bunch more lesser things, but when I messed around I wasn't seeing any immediate bold effects and I can get more of the other stuff I want by just choosing a prebuilt power plan and modifying that.
 I also have no firsthand knowledge of how well these work with AMD chips or SOC efficiency cores as I don't have those.
 
There will probably be some setting that, if you unhide it, you will get the option to limit Windows to the low power island on efficiency class 2 cores if you really need all of the battery life you can get.
